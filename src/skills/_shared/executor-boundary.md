# Executor Boundary

Use this shared boundary from SDD phase skills to prevent the orchestrator from executing phase work inline.

## Delegated Phase Skills

If you loaded a delegated phase skill via the `skill()` tool, or read `skills/sdd-*/SKILL.md` for a delegated phase, you are the ORCHESTRATOR and this is a boundary violation. Stop using that loaded content for orchestration decisions. Delegate to the dedicated phase sub-agent using the launch envelope and platform delegation primitive. Delegated phase skills are for EXECUTORS only.

If you are the dedicated phase sub-agent, the orchestrator gate does not apply to you. Continue with the phase work. Do not delegate. Do not call the Skill tool. You are the executor; execute the phase yourself.

## Inline Skills

Some SDD skills are intentionally inline. When a skill says `ORCHESTRATOR GATE (INLINE)`, the orchestrator may execute it directly. Inline skills still must not silently delegate unless their phase-specific instructions explicitly say to do so.

## Coordinator Exceptions

`sdd-onboard` is a coordinator workflow, not a normal executor phase. It may coordinate narrated phase launches only when it is explicitly configured as a coordinator exception by the orchestrator/runtime. Do not treat this exception as permission for ordinary phase executors (`sdd-apply`, `sdd-verify`, etc.) to delegate.

## Blocker Resolution Protocol

When a delegated phase returns `status: blocked` with `next_recommended: resolve-blockers`, the orchestrator MUST report blockers and stop. After the user resolves the blocker, the orchestrator MUST re-delegate the blocked phase to its dedicated executor sub-agent using the launch envelope contract, and MUST NOT execute phase logic inline.

Re-delegated executors SHOULD resume from the last persisted checkpoint artifact instead of restarting from scratch.

- **`sdd-apply` special case**: The executor MUST read existing `apply-progress` evidence and resume from the first unchecked task in persisted task state.
- **`sdd-archive` special case**: If a destructive-merge operation was partially applied, the orchestrator MUST surface partial-archive state and require explicit user confirmation (or repair instruction) before re-delegating archive. The orchestrator MUST NOT silently re-run archive from scratch in this state.
- **Partial artifact rule (all phases)**: Executors MUST read existing partial artifacts before resuming and MUST return `blocked` if required partial artifacts are unreadable or corrupt.
- **No-checkpoint rule**: If no readable checkpoint exists, the executor MAY start from the beginning and MUST NOT assume prior partial work.
