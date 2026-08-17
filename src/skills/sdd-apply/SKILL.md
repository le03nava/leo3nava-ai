---
name: sdd-apply
description: "Implement SDD tasks from specs and design. Trigger: orchestrator launches apply for one or more change tasks."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-apply` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for IMPLEMENTATION. You receive specific tasks from `tasks.md` and implement them by writing actual code. You follow the specs and design strictly.

## What You Receive

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.execution_mode` — controls blocking behavior
- `launch.status.nextRecommended` — must be `apply`; if `all_done` do not edit
- `launch.status.dependencies` — all planning phases must be completed
- `launch.artifacts.paths` or `launch.artifacts.refs` — refs for spec, design, test-design, tasks; do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute workspace root
- `launch.actionContext.allowedEditRoots` — CRITICAL: only edit files under these roots; return `blocked` if a needed edit is outside
- `launch.delivery_strategy` — resolved delivery decision; `null` means deferred (apply may still proceed for low-risk work)
- `launch.review.review_budget_lines` — resolved budget; default `400` when `null`
- `launch.chain_strategy` — `stacked-to-main | feature-branch-chain | null`
- `launch.review.size_exception` — `approved | pending | none | null`
- `launch.review.current_slice_boundary` — which tasks/PR slice to implement; `null` means all tasks
- `launch.skill_paths` — supplemental skills to load before work

The `## Phase Context` section contains:
- Which specific tasks are assigned to this batch (e.g., "Phase 1, tasks 1.1–1.3")
- Strict TDD instructions when active (`STRICT TDD MODE IS ACTIVE. Test runner: {cmd}.`)
- Apply-progress continuity instruction when a previous batch exists (`PREVIOUS APPLY-PROGRESS EXISTS. Read it first, merge, do not overwrite.`)

If `launch.actionContext.allowedEditRoots` is missing or empty, return `blocked` with `next_recommended: resolve-blockers` before reading any implementation file.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal, specs, design with mandatory `## Secure Development Design`, `test-design`, and tasks from the selected backend. Standalone `security-design.md` is legacy/read-only compatibility data only. |
| Produced artifact | Apply progress as `sdd/{change-name}/apply-progress`; in OpenSpec, the durable progress source is `openspec/changes/{change-name}/tasks.md` checkbox state plus the returned Section D evidence. |
| Mutates | Assigned implementation files inside `allowedEditRoots`; task progress in `sdd/{change-name}/tasks` / `openspec/changes/{change-name}/tasks.md`; apply-progress when the selected mode supports it. |
| Task progress semantics | Read previous progress first, preserve existing `[x]` marks, skip already-complete assigned tasks, and mark only completed assigned tasks. Hybrid writes must keep Engram progress and OpenSpec checkboxes aligned. |
| Apply evidence semantics | Record completed tasks, files changed, Standard/Strict TDD mode, test-design coverage or justified deviations, embedded secure-design evidence including N/A rationale where applicable, unavailable runtime/coverage/lint/typecheck/format tooling, issues, remaining tasks, workload/PR boundary, and persisted checkbox verification in `detailed_report` / apply-progress. |
| Deviation semantics | If implementation cannot follow design or `test-design.md`, record the deviation, rationale, replacement evidence, and downstream verify implication; do not silently drop mandatory planned evidence. |
| Conditional behavior | Strict TDD loads `strict-tdd.md` only when active. |
| Success routing | `next_recommended: apply` while implementation tasks remain; `next_recommended: review-parallel` only when all implementation tasks are visibly complete in the persisted task artifact. The orchestrator MUST launch `sdd-review` AND `sdd-review-security` in parallel; use `currentPhase: review-parallel`. |
| Block routing | `next_recommended: resolve-blockers` for unsafe workspace, unresolved workload decision, missing artifact, Strict TDD issue, partial persistence failure, or blocked task. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the implementation progress summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Assigned batch complete, implementation tasks remain**: return `next_recommended: apply` and report the next pending task/slice.
- **All implementation tasks complete**: return `next_recommended: review-parallel`. The orchestrator MUST launch `sdd-review` AND `sdd-review-security` in parallel (use `currentPhase: review-parallel`) and MUST NOT advance to `sdd-verify` until both phases appear in `completedPhases`. Verify and archive are never direct successors of apply.
- **Blocked apply**: return `next_recommended: resolve-blockers` and include the exact unsafe workspace, unresolved workload decision, missing artifact, Strict TDD issue, or blocked task in `risks` / `detailed_report`.
- **Partial persistence failure**: return `next_recommended: resolve-blockers` unless the same progress/task checkbox update can be safely retried without new user input.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| `applyState` is `blocked` | Return `blocked` with `next_recommended: resolve-blockers`; do not edit. |
| `applyState` is `all_done` | Do not edit; return `success` with `next_recommended: review-parallel`. The orchestrator MUST launch `sdd-review` AND `sdd-review-security` in parallel. Verify and archive are never direct successors of apply. |
| `workspace-planning` mode has no `allowedEditRoots` | Return `blocked` with `next_recommended: resolve-blockers`; treat linked repos and folders as read-only. |
| Needed edit is outside `allowedEditRoots` | Return `blocked` with `next_recommended: resolve-blockers` and the unsafe path. |
| Workload decision is required but unresolved | Return `blocked` with `next_recommended: resolve-blockers`, `Decision needed before apply: Yes`, `Chain strategy: pending`, and `Size exception: pending`; do not ask the user directly. |
| Assigned task is already complete | Skip it and continue with the first assigned incomplete task. |
| Task was not assigned to this apply batch | Do not implement it; return `blocked` with `next_recommended: resolve-blockers` if no assigned pending task remains. |
| Strict TDD is active without a runnable test command | Return `blocked` with `next_recommended: resolve-blockers` unless testing capabilities explicitly permit Standard Mode. |
| Persisted task artifact does not mark completed work `[x]` | Fix the artifact before returning. If it cannot be fixed, return `partial` or `blocked` with `next_recommended: resolve-blockers` and explain which persisted task state could not be updated. |

## Status and Workspace Guard

Before reading implementation files or writing code, consume the structured status provided by the orchestrator or build the equivalent status from artifacts.

- If `applyState` is `blocked`, STOP and return `blocked` with `next_recommended: resolve-blockers` plus the missing artifacts or unsafe context.
- If `applyState` is `all_done`, do not edit. Return `success` with `next_recommended: review-parallel`. The orchestrator MUST launch `sdd-review` AND `sdd-review-security` in parallel and MUST NOT launch `sdd-verify` until both appear in `completedPhases`. Verify and archive are never direct successors of apply.
- If `applyState` is `ready`, proceed only on the assigned pending tasks.
- Read context from `contextFiles` / `artifactPaths` instead of assuming fixed filenames. For spec-driven OpenSpec, these normally map to proposal, specs, design, and tasks.
- If `actionContext.mode` is `workspace-planning` and `allowedEditRoots` is empty, STOP before editing. Treat linked repos and folders as read-only planning context.
- If `allowedEditRoots` is present, edit only files under those roots. If a needed edit is outside the allowed roots, STOP and report the unsafe path.

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Read Context

Before writing ANY code:
1. Read the structured status and confirm `applyState: ready`
2. Read every applicable artifact path/topic from `contextFiles`, falling back to `artifactPaths` and `artifactRefs` according to `artifact_store.mode`
3. Read the specs — understand WHAT the code must do
4. Read the design — understand HOW to structure the code
5. Read mandatory `design.md#secure-development-design` — understand changed-surface classification, applicable narrative category rules, mandatory evidence expectations, residual risks, safe-evidence policy, and approved exceptions. Exhaustive matrix rows, `N/A` rationale, and lifecycle statuses belong to canonical `review-security-report.json` with derived Markdown compatibility; do not block new changes on missing standalone `security-design.md`.
6. Read `test-design.md` — understand planned automated, manual, static, and security-control checks and expected evidence
7. Read existing code in affected files — understand current patterns
8. Check the project's coding conventions from `config.yaml`

#### Step 2a: Enforce Review Workload Decision

Before implementing, inspect the tasks artifact for `Review Workload Forecast` and enforce `skills/_shared/sdd-phase-common.md#f-review-workload-guard`. This phase consumes the resolved delivery decision; it must not choose a missing strategy or approve an exception.

If the forecast says any of the following:

- `Review budget risk: High`
- `400-line budget risk: High`
- `Chained PRs recommended: Yes`
- `Decision needed before apply: Yes`

Then you MUST confirm the orchestrator/user provided a resolved delivery path:

1. For chained/stacked work, implement only the assigned work-unit slice, keep scope autonomous, and report the intended boundary.
2. For large single-PR work, continue only when the launch context records approved `size:exception` evidence and the tasks artifact says `Size exception: approved`.
3. For unresolved decisions, stop before writing code; the orchestrator owns user interaction.

Also check for `Chain strategy` in the tasks artifact. If present and not `pending`, follow the shared guard's rules for that strategy consistently.

Also check for `Size exception` in the tasks artifact or launch context:
- `approved`: a large single-PR path may proceed only when the prompt also records maintainer approval evidence.
- `pending`: STOP before writing code when review workload requires an exception.
- `none`: do not treat this as approval for over-budget single-PR work.

When the forecast requires a workload decision, if no resolved delivery decision, real chain strategy, or approved size exception is present, STOP before writing code and return `blocked` with `next_recommended: resolve-blockers`: `Workload decision required before apply: estimated work may exceed the session review budget. Decision needed before apply: Yes. Chain strategy: pending. Size exception: pending.` The orchestrator owns user interaction. Low-risk, under-budget work with `Decision needed before apply: No` does not require a chain strategy or size exception.

#### Step 2b: Read Previous Apply-Progress (if exists)

Before starting work, check for existing apply-progress using the selected artifact store:

1. `engram`: `mem_search(query: "sdd/{change-name}/apply-progress", project: "{project}")`; if found, `mem_get_observation(id)` and read the full content.
2. `openspec`: read `openspec/changes/{change-name}/tasks.md` checkbox state and any `applyProgress` entries from structured status.
3. If previous progress is unclear from the selected backend evidence, return `blocked` before editing.
5. Parse which tasks are already marked complete and skip those tasks — start from the first incomplete assigned task.
6. When saving your apply-progress in Step 7, MERGE: include all previously completed tasks PLUS your newly completed tasks in a single combined artifact.

**CRITICAL**: If the orchestrator told you previous progress exists, you MUST read it. If you overwrite without reading, completed work from prior batches is permanently lost.

### Step 3: Read Testing Capabilities and Resolve Mode

Read the cached testing capabilities to determine implementation mode:

```
Read testing capabilities from:
├── engram: mem_search("sdd/{project}/testing-capabilities") -> mem_get_observation(id)
├── openspec: openspec/config.yaml -> strict_tdd + testing section
└── Fallback only when persistence lacks capabilities: check project files directly (package.json, go.mod, etc.)

Resolve mode:
├── IF strict_tdd: true AND test runner exists
│   └── STRICT TDD MODE → Load and follow strict-tdd.md module
│       (read the file: skills/sdd-apply/strict-tdd.md)
│
├── IF strict_tdd: false OR no test runner
│   └── STANDARD MODE → use Step 4 below (no TDD module loaded)
│
└── Cache the resolved mode for the return summary
```

**Key principle**: If Strict TDD Mode is not active, ZERO TDD instructions are loaded. The `strict-tdd.md` module is never read, never processed, never consumes tokens.

#### Hard Gate (Strict TDD Only)

If Strict TDD Mode is active (either from orchestrator injection or self-discovery):
- You MUST produce a **TDD Cycle Evidence** table in your apply-progress artifact
- Each task row MUST have: RED (test written first) → GREEN (implementation passes) → REFACTOR columns
- If you complete a task WITHOUT writing tests first, mark it as FAILED in the evidence table
- The verify phase WILL reject your work if the TDD Evidence table is missing or incomplete

**There is no silent fallback.** If you resolved Strict TDD as active, you follow it or you report failure. You do NOT quietly switch to Standard Mode.

### Step 4: Implement Tasks (Standard Workflow)

This step is used when Strict TDD Mode is NOT active:

```
FOR EACH TASK:
├── Read the task description
├── Read relevant spec scenarios (these are your acceptance criteria)
├── Read the design decisions (these constrain your approach)
├── Read embedded secure-design controls and mandatory evidence when required
├── Read planned cases from test-design.md (these constrain evidence and checks)
├── Read existing code patterns (match the project's style)
├── Write the code
├── Follow planned checks or document justified deviations with replacement evidence
├── Mark task as complete [x] in the persisted tasks artifact immediately
└── Note any issues or deviations
```

### Step 5: Mark Tasks Complete

Update `tasks.md` — change `- [ ]` to `- [x]` for completed tasks:

```markdown
## Phase 1: Foundation

- [x] 1.1 Create `internal/auth/middleware.go` with JWT validation
- [x] 1.2 Add `AuthConfig` struct to `internal/config/config.go`
- [ ] 1.3 Add auth routes to `internal/server/server.go`  ← still pending
```

### Step 6: Validate Progress

Before persisting or returning, verify:

- Only assigned tasks were implemented.
- Every file edit is inside `allowedEditRoots` when roots are provided.
- Completed tasks are marked `[x]` in the persisted tasks artifact for `engram` and `openspec` modes.
- Planned `test-design.md` checks for the assigned slice were followed, or every deviation has a justification and replacement evidence in apply-progress.
- Security evidence for applicable controls in the assigned slice is recorded with guideline IDs, file references, evidence status, or complete approved exception details.
- Previous apply-progress was merged when it existed.
- Strict TDD mode includes the required TDD Cycle Evidence table.
- Workload / PR Boundary is reported.
- Do not report `Ready for review` if assigned work is incomplete or persisted artifacts are not updated.

### Step 7: Persist Progress

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `apply-progress`
- topic_key: `sdd/{change-name}/apply-progress`
- openspec task path: `openspec/changes/{change-name}/tasks.md`
- type: `architecture`
- Also update the tasks artifact with `[x]` marks via `mem_update` (engram) or file edit (openspec).

#### Merge Protocol

When saving apply-progress:
1. If you read previous progress in Step 2b, your artifact MUST include ALL previously completed tasks (copy their status and evidence) PLUS your new completions.
2. In `openspec`, the persisted progress source is `tasks.md` checkbox state; do not invent a separate apply-progress file.
3. The final artifact should show the cumulative state of ALL tasks across ALL batches.
4. Format: keep the same structure but ensure no completed task is lost from prior batches.

### Step 8: Return Summary

Before returning, re-read the persisted tasks artifact in `engram` and `openspec` modes and confirm every task you report as completed is marked `[x]` there. If the artifact still shows a completed task as `- [ ]`, fix the checkbox before returning. Do not report `Ready for review` while completed work is only reflected in internal todos or apply-progress.

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put this implementation progress summary in `detailed_report`:

```markdown
## Implementation Progress

**Change**: {change-name}
**Mode**: {Strict TDD | Standard}

### Completed Tasks
- [x] {task 1.1 description}
- [x] {task 1.2 description}

### Files Changed
| File | Action | What Was Done |
|------|--------|---------------|
| `path/to/file.ext` | Created | {brief description} |
| `path/to/other.ext` | Modified | {brief description} |

{IF Strict TDD Mode → include TDD Cycle Evidence table from strict-tdd.md}

### Deviations from Design
{List any places where the implementation deviated from design.md and why.
If none, say "None — implementation matches design."}

### Test-Design Evidence
{List planned case IDs covered by this apply batch, the evidence/check performed, and any justified deviations with replacement evidence. If none apply to this slice, state why.}

### Security Evidence
{List applicable guideline IDs from `design.md#secure-development-design`, implementation references, evidence status, N/A row preservation, and approved exceptions if any.}

### Issues Found
{List any problems discovered during implementation.
If none, say "None."}

### Remaining Tasks
- [ ] {next task}
- [ ] {next task}

### Workload / PR Boundary
- Mode: {single PR | chained PR slice | stacked PR slice | size:exception}
- Current work unit: {unit name or "N/A"}
- Boundary: {what this apply batch starts from and ends with}
- Estimated review budget impact: {brief note}

### Status
{N}/{total} tasks complete. {Ready for next batch / Ready for parallel review (sdd-review ∥ sdd-review-security) / Blocked by X}
```

## Rules

- ALWAYS read specs before implementing — specs are your acceptance criteria
- ALWAYS follow the design decisions — don't freelance a different approach
- ALWAYS read `test-design.md` before implementing — planned cases are the evidence contract for apply and verify
- ALWAYS read mandatory `design.md#secure-development-design` before implementing — mandatory controls are evidence contracts for apply, review-security, verify, and archive; standalone `security-design.md` is legacy/read-only only
- ALWAYS match existing code patterns and conventions in the project
- ALWAYS consume or produce structured status before implementation; do not infer readiness from conversation alone
- STOP on `applyState: blocked` and do not edit; STOP on unsafe `actionContext` or edit roots
- In `openspec` mode, mark tasks complete in `tasks.md` AS you go, not at the end
- Before returning, re-read the persisted tasks artifact and ensure completed tasks are visibly marked `[x]`; internal todos are not completion evidence
- If you discover the design is wrong or incomplete, NOTE IT in your return summary — don't silently deviate
- If you cannot follow a planned `test-design.md` case exactly, document the deviation and replacement evidence in apply-progress; do not silently drop planned checks
- If you cannot produce mandatory security evidence, document a complete approved exception or stop as blocked; incomplete exceptions do not satisfy archive readiness
- If a task is blocked by something unexpected, STOP and report back
- If workload forecast requires a decision and none was provided, STOP before writing code
- When applying a chained/stacked PR slice, keep the batch autonomous: one deliverable scope, verification included, and clear rollback boundary
- When applying `size:exception`, state it explicitly in apply-progress and the return summary
- NEVER implement tasks that weren't assigned to you
- Skill loading is handled in Step 1 — follow any loaded skills strictly when writing code
- Apply any `rules.apply` from `openspec/config.yaml`
- If Strict TDD Mode is active (Step 3), load `strict-tdd.md` and follow its cycle INSTEAD of Step 4
- When Strict TDD is active, the `strict-tdd.md` module's rules OVERRIDE Step 4 entirely
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the implementation progress summary belongs in `detailed_report`.
