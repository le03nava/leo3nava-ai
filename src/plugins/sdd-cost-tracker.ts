/**
 * sdd-cost-tracker plugin
 *
 * Automatically captures token/cost data for SDD phase sub-agents and posts
 * a PhaseRecord to the sdd-cost-tracker HTTP server when the child session ends.
 *
 * Capture flow:
 *   tool.execute.before (Task tool)
 *     → detect subagent_type matching sdd-*
 *     → extract changeName from the launch envelope
 *     → queue { phase, changeName } keyed by parent sessionID
 *
 *   event: session.created (child session with parentID)
 *     → associate childSessionID → { phase, changeName, project }
 *
 *   event: message.updated (AssistantMessage in a tracked child session)
 *     → accumulate tokens + cost by childSessionID
 *
 *   event: session.idle
 *     → POST /phases with the accumulated totals for that session
 */

import type { Plugin } from "@opencode-ai/plugin"

// ---------------------------------------------------------------------------
// Configuration
// ---------------------------------------------------------------------------

const TRACKER_URL = process.env.SDD_COST_TRACKER_URL ?? "http://127.0.0.1:7438"

// ---------------------------------------------------------------------------
// Internal types
// ---------------------------------------------------------------------------

/** Pending context queued from tool.execute.before, keyed by parent sessionID. */
interface PendingContext {
  phase: string
  changeName: string
}

/** Full tracking context attached to a live child session. */
interface SessionContext {
  phase: string
  changeName: string
  project: string
  modelId: string | null
  providerId: string | null
  tokensInput: number
  tokensOutput: number
  tokensReasoning: number
  tokensCacheRead: number
  tokensCacheWrite: number
  costUsd: number
  startedAt: number
}

// ---------------------------------------------------------------------------
// Regex helpers
// ---------------------------------------------------------------------------

const RE_SUBAGENT_TYPE = /subagent_type[:\s]+sdd-([a-z-]+)/i
const RE_CHANGE_NAME = /change_?name[:\s"']+([^\s"',\n]+)/i

function extractSddPhase(text: string): string | null {
  const m = RE_SUBAGENT_TYPE.exec(text)
  return m ? `sdd-${m[1]}` : null
}

function extractChangeName(text: string): string | null {
  const m = RE_CHANGE_NAME.exec(text)
  return m ? m[1] : null
}

// ---------------------------------------------------------------------------
// HTTP helper
// ---------------------------------------------------------------------------

async function postPhase(sessionId: string, ctx: SessionContext): Promise<void> {
  const now = Math.floor(Date.now() / 1000)

  const body = JSON.stringify({
    project: ctx.project,
    change_name: ctx.changeName,
    phase: ctx.phase,
    session_id: sessionId,
    model_id: ctx.modelId,
    provider_id: ctx.providerId,
    tokens_input: ctx.tokensInput,
    tokens_output: ctx.tokensOutput,
    tokens_reasoning: ctx.tokensReasoning,
    tokens_cache_read: ctx.tokensCacheRead,
    tokens_cache_write: ctx.tokensCacheWrite,
    cost_usd: ctx.costUsd,
    started_at: ctx.startedAt,
    completed_at: now,
  })

  try {
    const res = await fetch(`${TRACKER_URL}/phases`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body,
    })

    if (!res.ok) {
      const text = await res.text().catch(() => "(no body)")
      console.error(`[sdd-cost-tracker] POST /phases failed: HTTP ${res.status} — ${text}`)
    }
  } catch (err) {
    // Fire-and-forget: never crash the session on a tracking failure
    console.error(`[sdd-cost-tracker] POST /phases error: ${err}`)
  }
}

async function postCall(sessionId: string, callIndex: number, msg: any): Promise<void> {
  const tokens = msg?.tokens
  const costRaw = msg?.cost ?? msg?.costUSD ?? 0

  const body = JSON.stringify({
    session_id: sessionId,
    call_index: callIndex,
    model_id: msg?.modelID ?? null,
    provider_id: msg?.providerID ?? null,
    tokens_input: tokens?.input ?? tokens?.inputTokens ?? 0,
    tokens_output: tokens?.output ?? tokens?.outputTokens ?? 0,
    tokens_reasoning: tokens?.reasoning ?? tokens?.reasoningTokens ?? 0,
    tokens_cache_read: tokens?.cache?.read ?? tokens?.cacheRead ?? tokens?.cacheReadTokens ?? 0,
    tokens_cache_write:
      tokens?.cache?.write ?? tokens?.cacheWrite ?? tokens?.cacheWriteTokens ?? 0,
    cost_usd: typeof costRaw === "number" ? costRaw : 0,
  })

  try {
    const res = await fetch(`${TRACKER_URL}/calls`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body,
    })

    if (!res.ok) {
      const text = await res.text().catch(() => "(no body)")
      console.error(`[sdd-cost-tracker] POST /calls failed: HTTP ${res.status} — ${text}`)
    }
  } catch (err) {
    // Fire-and-forget: never crash the session on a tracking failure
    console.error(`[sdd-cost-tracker] POST /calls error: ${err}`)
  }
}

// ---------------------------------------------------------------------------
// Plugin factory
// ---------------------------------------------------------------------------

const SddCostTrackerPlugin: Plugin = async ({ project }) => {
  // project may be a string name or a full project object depending on opencode version
  const projectName: string =
    typeof project === "string"
      ? project
      : typeof (project as any)?.name === "string"
        ? (project as any).name
        : typeof (project as any)?.worktree === "string"
          ? (project as any).worktree.replace(/\\/g, "/").split("/").pop() ?? ""
          : ""

  /**
   * Queued contexts from tool.execute.before, waiting to be associated with a
   * child session once session.created fires with a parentID.
   * Key: parentSessionID
   */
  const pendingByParent = new Map<string, PendingContext>()

  /**
   * Active tracking state for child sessions.
   * Key: childSessionID
   */
  const sessions = new Map<string, SessionContext>()

  /**
   * Per-session call index for call-level tracking.
   * Key: childSessionID
   */
  const callIndices = new Map<string, number>()

  return {
    // -----------------------------------------------------------------------
    // Intercept Task tool launches to detect SDD sub-agent invocations
    // -----------------------------------------------------------------------
    "tool.execute.before": async (input, output) => {
      if (input.tool !== "task") return

      const args = (output as any).args ?? {}
      const prompt: string =
        typeof args.prompt === "string" ? args.prompt : JSON.stringify(args)

      const phaseFromArg: string | null =
        typeof args.subagent_type === "string" && args.subagent_type.startsWith("sdd-")
          ? args.subagent_type
          : null

      const phase = phaseFromArg ?? extractSddPhase(prompt)
      const changeName = extractChangeName(prompt)

      if (!phase) return
      if (!changeName) {
        console.error(`[sdd-cost-tracker] phase=${phase} but changeName not found in prompt`)
        return
      }

      const parentSessionID: string | undefined = (input as any).sessionID
      if (!parentSessionID) {
        console.error(`[sdd-cost-tracker] parentSessionID missing on input`)
        return
      }

      pendingByParent.set(parentSessionID, { phase, changeName })
    },

    // -----------------------------------------------------------------------
    // React to all bus events
    // -----------------------------------------------------------------------
    event: async ({ event }: { event: { type: string; properties?: any } }) => {
      const { type, properties } = event

      // -------------------------------------------------------------------
      // session.created — associate child session with pending context
      // -------------------------------------------------------------------
      if (type === "session.created") {
        const childID: string | undefined = properties?.info?.id
        const parentID: string | undefined = properties?.info?.parentID

        if (!childID || !parentID) return

        const pending = pendingByParent.get(parentID)
        if (!pending) return

        pendingByParent.delete(parentID)

        sessions.set(childID, {
          phase: pending.phase,
          changeName: pending.changeName,
          project: projectName,
          modelId: null,
          providerId: null,
          tokensInput: 0,
          tokensOutput: 0,
          tokensReasoning: 0,
          tokensCacheRead: 0,
          tokensCacheWrite: 0,
          costUsd: 0,
          startedAt: Math.floor(Date.now() / 1000),
        })
        callIndices.set(childID, 0)
        return
      }

      // -------------------------------------------------------------------
      // message.updated — accumulate tokens & cost for tracked sessions
      // -------------------------------------------------------------------
      if (type === "message.updated") {
        const msg = properties?.info ?? properties?.message
        if (!msg || msg.role !== "assistant") return

        const sessionID: string | undefined =
          properties?.info?.sessionID ?? properties?.sessionID
        if (!sessionID) return

        const ctx = sessions.get(sessionID)
        if (!ctx) return

        if (!ctx.modelId && msg.modelID) ctx.modelId = msg.modelID
        if (!ctx.providerId && msg.providerID) ctx.providerId = msg.providerID

        const t = msg.tokens
        const hasTokenData =
          t &&
          [
            t.input,
            t.inputTokens,
            t.output,
            t.outputTokens,
            t.reasoning,
            t.reasoningTokens,
            t.cache?.read,
            t.cacheRead,
            t.cacheReadTokens,
            t.cache?.write,
            t.cacheWrite,
            t.cacheWriteTokens,
          ].some((v) => typeof v === "number")

        if (hasTokenData) {
          const callIndex = callIndices.get(sessionID) ?? 0
          void postCall(sessionID, callIndex, msg)
          callIndices.set(sessionID, callIndex + 1)
        }

        if (t) {
          ctx.tokensInput += t.input ?? t.inputTokens ?? 0
          ctx.tokensOutput += t.output ?? t.outputTokens ?? 0
          ctx.tokensReasoning += t.reasoning ?? t.reasoningTokens ?? 0
          ctx.tokensCacheRead += t.cache?.read ?? t.cacheRead ?? t.cacheReadTokens ?? 0
          ctx.tokensCacheWrite += t.cache?.write ?? t.cacheWrite ?? t.cacheWriteTokens ?? 0
        }

        const cost = msg.cost ?? msg.costUSD ?? 0
        ctx.costUsd += typeof cost === "number" ? cost : 0
        return
      }

      // -------------------------------------------------------------------
      // session.idle — flush accumulated data to the tracker
      // -------------------------------------------------------------------
      if (type === "session.idle") {
        const sessionID: string | undefined =
          properties?.sessionID ?? properties?.info?.id

        if (!sessionID) return

        const ctx = sessions.get(sessionID)
        if (!ctx) return

        sessions.delete(sessionID)
        callIndices.delete(sessionID)

        if (ctx.tokensInput === 0 && ctx.tokensOutput === 0 && ctx.costUsd === 0) return

        await postPhase(sessionID, ctx)
      }
    },
  }
}

export default SddCostTrackerPlugin
