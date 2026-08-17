---
name: sdd-new
description: "Start a new SDD change. Trigger: orchestrator coordinates a new change through exploration and proposal."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
  delegate_only: false
---

> **ORCHESTRATOR GATE (INLINE)**: Follow `skills/_shared/executor-boundary.md`.
> This workflow skill is intentionally coordinated inline by the orchestrator.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

Coordinate the start of a new SDD change. This workflow uses cached preflight state, launches exploration when needed, collects proposal-shaping input in interactive mode, then launches proposal creation for the same change.

## What You Receive

From the user or orchestrator:
- Change name or raw change request
- Current project/workspace identity
- Cached SDD session preflight: execution mode and artifact store mode. Delivery strategy, review budget, and chain strategy may be `null` until `sdd-tasks` forecast or an explicit delivery decision resolves them.

## Decision Gates

| Situation | Action |
| --- | --- |
| Change name or request is missing | Ask for one concrete change request and stop. |
| SDD session preflight is incomplete | Ask the exact orchestrator preflight prompt and stop; do not launch phases in the same turn. |
| Init is missing or partial for the selected artifact store mode | Run `sdd-init` first, then resume `sdd-new`. |
| Artifact store mode is unknown | Resolve it through preflight; do not hardcode Engram. |
| Exploration returns `blocked` | Present the blocker and stop before proposal. |
| Interactive mode is active before proposal | Ask/collect proposal-shaping answers, or record explicit user skip/approval, before launching `sdd-propose`. |
| Proposal-shaping input is unavailable and product facts are needed | Stop with a blocker instead of launching proposal blindly. |
| Proposal returns `blocked` | Present the needed user decision or blocker and stop before specs/design. |
| Interactive mode is active after exploration or proposal | Present the phase summary and ask whether to adjust or continue. |

## What to Do

### Step 1: Confirm Preflight

Ensure the SDD session has execution mode and artifact store mode cached. Preserve delivery strategy, review budget, and chain strategy as `null` when they have not yet been resolved.

### Step 2: Ensure Init

Confirm initialization exists in the selected artifact store using the orchestrator Init Guard: `sdd-init/{project}` plus testing capabilities for `engram`, or `openspec/config.yaml` with project context plus testing capabilities for `openspec`. If the selected mode is missing or partial, delegate `sdd-init` and resume only after it succeeds.

### Step 3: Launch Exploration

Delegate to the dedicated `sdd-explore` sub-agent with the change name/request, project, artifact store mode, and relevant skill paths.

### Step 4: Collect Proposal-Shaping Input

In interactive mode, the orchestrator asks focused product/PRD questions after exploration when applicable, or after init when exploration is skipped. Record the user's answers, corrections, or explicit skip/approval before proposal launch.

Questions should cover product facts such as the business problem, target users, business rules, outcomes, current-state gap, impact, edge cases, first-slice boundaries, non-goals, constraints, and tradeoffs. Do not ask delivery mechanics such as PR shape, review budget, chain strategy, or test commands unless the user explicitly requests delivery planning.

In automatic mode, or when the user skips the round, launch proposal only if the request/exploration contains enough product facts to avoid material speculation. If missing facts would force guessing, stop with `status: blocked` and name the missing facts.

### Step 5: Launch Proposal

If exploration succeeds or the user provided enough direct input, delegate to the dedicated `sdd-propose` sub-agent with the exploration artifact/reference and the same artifact store mode.

### Step 6: Present Next Action

Present the proposal summary and recommend `sdd-spec` as the next planning phase. `sdd-design` runs only after specs pass.

## Output Contract

Return a concise orchestrator summary, not a phase artifact. Include:

- `status`: `success`, `partial`, or `blocked`
- `executive_summary`: what exploration/proposal produced or why the workflow stopped
- `artifacts`: exploration and proposal artifact references, when created
- `next_recommended`: `sdd-spec`, `resolve-blockers`, or `none`
- `risks`: relevant blockers, unclear product decisions, or `None`

## Rules

- Do not execute exploration or proposal work inline; delegate to their dedicated sub-agents.
- Do not create artifacts directly from this workflow; sub-agents own persistence using the selected artifact store.
- In `engram`, sub-agents save phase artifacts under `sdd/{change-name}/{artifact}` topic keys.
- Do not proceed from exploration to proposal when exploration reports missing context that would make the proposal speculative.
- Do not proceed to proposal in automatic mode when material product facts are missing; report a blocker instead of guessing.
