---
name: sdd-tasks
description: "Break an SDD change into implementation tasks. Trigger: orchestrator launches task planning for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-tasks` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for creating the TASK BREAKDOWN. You take the proposal, specs, design with embedded secure development design, and test design, then produce a `tasks.md` with concrete, actionable implementation steps organized by phase.

## What You Receive

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.execution_mode` — controls blocking behavior
- `launch.status.dependencies` — `spec: completed`, `design: completed`, and `test-design: completed` must be present
- `launch.artifacts.paths.spec` or `launch.artifacts.refs.spec` — do NOT expect content inline
- `launch.artifacts.paths.design` or `launch.artifacts.refs.design` — must include `## Secure Development Design`; do NOT expect content inline
- `launch.artifacts.paths.test-design` or `launch.artifacts.refs.test-design` — canonical `test-cases.json` plus derived `test-design.md`; do NOT expect content inline
- `launch.delivery_strategy` — `ask-on-risk | auto-chain | single-pr | exception-ok | null`; `null` means deferred
- `launch.review.review_budget_lines` — default `400` when `null` because delivery planning is deferred
- `launch.chain_strategy` — `stacked-to-main | feature-branch-chain | pending | null`
- `launch.review.size_exception` — `approved | pending | none | null`
- `launch.actionContext.workspaceRoot` — absolute workspace path
- `launch.skill_paths` — supplemental skills to load before work

If `launch.artifact_store.mode` is absent, or `launch.status.dependencies` does not show `spec`, `design`, and `test-design` as completed, return `blocked` with `next_recommended: resolve-blockers`.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal, specs, design with mandatory `## Secure Development Design`, mandatory `test-cases.json` (canonical) and `test-design.md` (derived), delivery context, and testing capabilities from the selected backend. Standalone `security-design.md` is legacy/read-only compatibility data only. |
| Produced artifact | `sdd/{change-name}/tasks` or `openspec/changes/{change-name}/tasks.md`. |
| Mutates | None outside the produced tasks artifact. |
| Test-design consumption | Tasks must derive implementation, testing, static/manual evidence, validation-metadata checks, verification, and warning work from `test-cases.json` (canonical) and `test-design.md` (derived); omitted mandatory planned cases are blockers. |
| Security consumption | Applicable narrative security rules, review-security expectations, and evidence expectations from `design.md#secure-development-design` and `test-design.md` must be represented as tasks or complete approved exceptions. Exhaustive `N/A` rationale and lifecycle row status belong to canonical `review-security-report.json` with derived Markdown compatibility. |
| Operational considerations consumption | Tasks must represent concrete operational evidence collection, exact-marker validation, safe-evidence checks, non-SQL monitoring coverage, unavailable-tooling notes, unresolved-gap carry-forward, and archive/manual-document handoff only when derived from `design.md`, `test-design.md`, or equivalent artifact evidence. |
| Review workload behavior | Preserve the Review Workload Forecast guard lines, delivery strategy (`null` when deferred), chain strategy, size-exception field, and reviewable work-unit split. |
| Success routing | `next_recommended: apply`, including when the workload guard requires the orchestrator to resolve apply-time decisions. |
| Block routing | `next_recommended: resolve-blockers` for missing required inputs, missing embedded secure development design, test-design gaps, testing capability blockers, or persistence failure. Do not route new changes to standalone `security-design`. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the tasks summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Successful tasks with no blocking workload decision**: return `next_recommended: apply`. The orchestrator normalizes this into state `nextRecommended: apply` before routing or persisting state.
- **Tasks created but workload decision is required**: return `next_recommended: apply`, include `Decision needed before apply: Yes`, and leave the blocker for the orchestrator's Review Workload Guard. Do not ask the user directly.
- **Blocked tasks**: return `next_recommended: resolve-blockers` and include the exact missing proposal, spec, embedded secure development design, test-design artifact, testing capability, or task validation issue in `risks` / `detailed_report`.
- **Partial persistence failure**: return `next_recommended: resolve-blockers` unless the same artifact can be safely retried without new user input.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| Required proposal, spec, design, or test-design is missing | Return `blocked` with `next_recommended: resolve-blockers`; do not write tasks. |
| `design.md#secure-development-design` is missing for a new active change | Return `blocked` with `next_recommended: resolve-blockers`; do not write tasks. |
| Standalone `security-design.md` is missing for a new active change | Continue; do not require it. It is legacy/read-only compatibility data only. |
| Mandatory security control evidence from `design.md#secure-development-design` or `test-design.md` is not represented in tasks | Return `blocked` with `next_recommended: resolve-blockers`; do not drop mandatory security evidence. |
| Applicable operational evidence or checks from design/test-design are not represented in tasks | Return `blocked` with `next_recommended: resolve-blockers`; applicable operational work must not remain implicit. |
| Task draft contains vague, non-verifiable, or oversized tasks | Fix it before persistence; if it cannot be fixed, return `blocked` with `next_recommended: resolve-blockers`. |
| Review workload risk is `High` against the received review budget and chain strategy is missing | Set `Decision needed before apply: Yes` and `Chain strategy: pending`; do not ask the user directly. |
| Strict TDD is active | Include RED/GREEN/REFACTOR task ordering for affected behavior. |
| Any case with `mandatory: true` in `test-cases.json` has no covering task | Return `blocked` with `next_recommended: resolve-blockers`; do not persist tasks. |

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Analyze the Design

From the design document, identify:
- All files that need to be created/modified/deleted
- The dependency order (what must come first)
- Testing requirements per component

From `test-cases.json` (canonical source of truth), identify:
- Planned automated, manual, and static checks
- Mandatory and non-mandatory cases from case metadata
- Every case with `mandatory: true` that MUST be covered by at least one implementation, testing, or evidence task
- Expected evidence that `sdd-apply` and `sdd-verify` will later consume

From `test-design.md` (derived human-readable view), identify:
- Planned automated, manual, and static checks
- Mandatory cases that must become implementation, testing, or evidence tasks
- Non-mandatory cases that should become advisory evidence tasks when feasible
- Expected evidence that `sdd-apply` and `sdd-verify` will later consume
- Operational checks for marker exactness, safe evidence, traceability, monitoring mechanism coverage, unavailable tooling, final-document boundaries, and restricted-data absence when planned by design/test-design evidence

From mandatory `design.md#secure-development-design`, identify:
- Changed-surface classification, applicable narrative category rules, static/manual validation notes, catalog context, safe-evidence policy, residual risks, and complete approved exceptions.
- Expected evidence owners, review-security expectations, and implementation, verification, or archive evidence obligations for applicable rules.
- Implementation, apply-evidence, review-security, verification, or archive-evidence tasks needed to satisfy mandatory controls.

From `design.md#Operational Considerations`, equivalent design evidence, and `test-design.md`, identify concrete tasks for applicable evidence collection, exact `Pendiente de confirmar:` / `No aplica.` validation, safe-evidence inspection, non-SQL-only monitoring mechanism coverage, unresolved-gap ownership, archive preservation, manual operational-document handoff, and unavailable-tooling reporting. Do not create tasks that require real restricted operational values or absent operational categories.

Also read testing capabilities when available:
- Engram: `sdd/{project}/testing-capabilities`
- OpenSpec: `openspec/config.yaml` `testing` section

Use testing capabilities to decide whether test-first RED/GREEN/REFACTOR tasks are required. Do not invent TDD tasks when Strict TDD is unavailable or disabled.

If runtime test, coverage, linter, type-checker, or formatter commands are unavailable, create static/manual evidence tasks and require later apply/verify reports to state the unavailable tooling explicitly.

### Step 3: Write tasks.md

**IF mode is `openspec`:** Create the task file:

```
openspec/changes/{change-name}/
├── proposal.md
├── specs/
├── design.md
├── test-design.md
└── tasks.md               ← You create this; consumes design.md#secure-development-design
```

**IF mode is `engram`:** Do NOT create any `openspec/` directories or files. Compose the tasks content in memory; persist it in Engram.

#### Task File Format

```markdown
# Tasks: {Change Title}

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | <rough estimate or range> |
| Review budget lines | <review_budget_lines from preflight> |
| 400-line budget risk | Low / Medium / High |
| Review budget risk | Low / Medium / High |
| Chained PRs recommended | Yes / No |
| Suggested split | <single PR or PR 1 → PR 2 → PR 3> |
| Delivery strategy | <ask-on-risk / auto-chain / single-pr / exception-ok / null> |
| Chain strategy | <stacked-to-main / feature-branch-chain / pending> |
| Size exception | <approved / pending / none> |

Decision needed before apply: <Yes|No>
Chained PRs recommended: <Yes|No>
Chain strategy: <stacked-to-main|feature-branch-chain|pending>
Size exception: <approved|pending|none>
Review budget lines: <number>
Review budget risk: <Low|Medium|High>
400-line budget risk: <Low|Medium|High>

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | <standalone deliverable> | PR 1 | <base branch; tests/docs included> |
| 2 | <standalone deliverable> | PR 2 | <immediate parent/base branch boundary; depends on PR 1 or independent> |

## Phase 1: {Phase Name} (e.g., Infrastructure / Foundation)

- [ ] 1.1 {Concrete action — what file, what change}
- [ ] 1.2 {Concrete action}
- [ ] 1.3 {Concrete action}

## Phase 2: {Phase Name} (e.g., Core Implementation)

- [ ] 2.1 {Concrete action}
- [ ] 2.2 {Concrete action}
- [ ] 2.3 {Concrete action}
- [ ] 2.4 {Concrete action}

## Phase 3: {Phase Name} (e.g., Testing / Verification)

- [ ] 3.1 {Write tests for ...}
- [ ] 3.2 {Write tests for ...}
- [ ] 3.3 {Verify integration between ...}

## Phase 4: {Phase Name} (e.g., Cleanup / Documentation)

- [ ] 4.1 {Update docs/comments}
- [ ] 4.2 {Remove temporary code}
```

### Task Writing Rules

Each task MUST be:

| Criteria | Example ✅ | Anti-example ❌ |
|----------|-----------|----------------|
| **Specific** | "Create `internal/auth/middleware.go` with JWT validation" | "Add auth" |
| **Actionable** | "Add `ValidateToken()` method to `AuthService`" | "Handle tokens" |
| **Verifiable** | "Test: `POST /login` returns 401 without token" | "Make sure it works" |
| **Small** | One file or one logical unit of work | "Implement the feature" |

### Review Workload Forecast Rules

Follow the shared Review Workload / Delivery Guard in `skills/_shared/sdd-phase-common.md#f-review-workload-guard`. This phase produces the forecast; it does not ask the user or approve delivery exceptions.

Before finalizing tasks, estimate whether implementation is likely to exceed the resolved session `review_budget_lines` (`additions + deletions`). This is a planning guard, not an exact diff count. If the orchestrator did not pass a budget because delivery planning is still deferred, use `400`.

Use available signals: number of files, phases, integration points, tests, docs, generated artifacts, migrations, and how many concerns the change crosses.

If the estimate is **High** or likely above the session review budget:

1. Mark `Chained PRs recommended` as `Yes`.
2. Split tasks into **work units** that can become chained or stacked PRs.
3. Each suggested PR must have a clear start, clear finish, verification, and autonomous scope.
4. Do not ask the user directly. If chain strategy is missing, set `Decision needed before apply: Yes` and `Chain strategy: pending`; the orchestrator owns user interaction. Valid strategies and size-exception semantics are defined by the shared guard.
5. Use the received delivery strategy and any orchestrator-provided chain strategy to set `Decision needed before apply`:
   - `null`: follow the shared deferred-decision rule.
   - `ask-on-risk`, `auto-chain`, `single-pr`, `exception-ok`: derive `Decision needed before apply`, `Chain strategy`, and `Size exception` from the shared guard and the orchestration context.

Do not bury this in prose. Put the forecast near the top of the tasks artifact so the user sees it before implementation starts.

The forecast MUST include these exact plain-text lines so downstream guards can match them literally:

```text
Decision needed before apply: Yes|No
Chained PRs recommended: Yes|No
Chain strategy: stacked-to-main|feature-branch-chain|pending
Size exception: approved|pending|none
Review budget lines: <number>
Review budget risk: Low|Medium|High
400-line budget risk: Low|Medium|High
```

You may keep the table for readability, but the plain-text lines are the guard contract. Keep `400-line budget risk` as a legacy compatibility line even when the active review budget is not 400; `Review budget risk` is authoritative for the current session.

For `feature-branch-chain`, suggested work units SHOULD name the intended base boundary: PR #1 base = feature/tracker branch; PR #2 base = PR #1 branch; PR #3 base = PR #2 branch. If a child PR would show previous PR changes, the base is wrong and must be retargeted/rebased before review.

### Phase Organization Guidelines

```
Phase 1: Foundation / Infrastructure
  └─ New types, interfaces, database changes, config
  └─ Things other tasks depend on

Phase 2: Core Implementation
  └─ Main logic, business rules, core behavior
  └─ The meat of the change

Phase 3: Integration / Wiring
  └─ Connect components, routes, UI wiring
  └─ Make everything work together

Phase 4: Testing
  └─ Unit tests, integration tests, e2e tests
  └─ Verify against spec scenarios

Phase 5: Cleanup (if needed)
  └─ Documentation, remove dead code, polish
```

### Step 4: Validate Tasks

Before persisting or returning, verify:

- Every task references concrete file paths or specific spec scenarios.
- Every task is specific, actionable, verifiable, and small enough for one session.
- Tasks are ordered by dependency.
- Testing/evidence tasks reference specific planned cases from `test-design.md` and scenarios from the specs.
- Security evidence tasks reference guideline IDs and controls from `design.md#secure-development-design` when required.
- Every mandatory planned case in `test-design.md` is represented by implementation, testing, or evidence work; omitted mandatory cases are blockers.
- Every case with `mandatory: true` in `test-cases.json` is traceable to at least one task. If any mandatory case has no corresponding task, return `blocked` with `next_recommended: resolve-blockers` and name the uncovered case IDs.
- Every mandatory applicable security guideline has implementation, test-design, apply, verify, archive evidence, or a complete approved exception represented in tasks.
- Static/manual validation notes, review-security evidence, archive evidence fields, and unavailable runtime/coverage/lint/typecheck/format reporting from `design.md#secure-development-design` / `test-design.md` are represented in tasks when applicable.
- Operational evidence collection, marker validation, safe-evidence checks, monitoring mechanism coverage, archive/manual-document handoff, and unavailable-tooling notes are represented as concrete tasks when applicable from design/test-design evidence.
- The Review Workload Forecast includes the required plain-text guard lines.
- If `Review budget risk` or `400-line budget risk` is `High`, Suggested Work Units are present.
- If `feature-branch-chain` is selected, work units name the intended base boundaries.
- `size:exception` is represented only in `Size exception`, not as a `Chain strategy` value.
- The tasks artifact stays under the 530-word size budget.

### Step 5: Persist Artifact

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `tasks`
- topic_key: `sdd/{change-name}/tasks`
- openspec path: `openspec/changes/{change-name}/tasks.md`
- type: `architecture`

### Step 6: Return Summary

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Use the canonical YAML format defined in the Section D example. Adapt these fields for this phase:

- `phase`: `tasks`
- `next_recommended`: `apply` (success, including when workload decision is needed) | `resolve-blockers` (blocked/partial)
- `executive_summary`: one short paragraph — total tasks, phase count, review budget risk level, whether a decision is needed before apply, and chain strategy if resolved
- `artifacts`: one entry; type `tasks`; correct mode, ref/path, `persisted: true`, `readable: true`
- `risks`: structured array or `None` — never an empty array `[]`. Include a WARNING risk entry when `Decision needed before apply: Yes`.
- `skill_resolution`: from `skill-resolver.md#step-4-report-resolution`
- `detailed_report`: use this minimum content for `sdd-tasks`. **The Review Workload Forecast guard lines are consumed literally by the orchestrator — preserve them exactly:**

```
## Tasks Created

**Change**: {change-name}
**Location**: `openspec/changes/{change-name}/tasks.md` (openspec) | Engram `sdd/{change-name}/tasks` (engram)

### Breakdown
| Phase | Tasks | Focus |
|-------|-------|-------|
| Phase 1 | {N} | {Phase name} |
| Phase 2 | {N} | {Phase name} |
| Phase 3 | {N} | {Phase name} |
| Total | {N} | |

### Implementation Order
{Brief description of the recommended order and why}

### Review Workload Forecast
- Estimated changed lines: {estimate or range}
- Review budget lines: {review_budget_lines}
- Review budget risk: {Low | Medium | High}
- 400-line budget risk: {Low | Medium | High}
- Chained PRs recommended: {Yes | No}
- Delivery strategy: {ask-on-risk | auto-chain | single-pr | exception-ok | null}
- Chain strategy: {stacked-to-main | feature-branch-chain | pending}
- Size exception: {approved | pending | none}
- Decision needed before apply: {Yes | No}
- Suggested work-unit PR split: {brief list or "Not needed"}

Decision needed before apply: {Yes|No}
Chained PRs recommended: {Yes|No}
Chain strategy: {stacked-to-main|feature-branch-chain|pending}
Size exception: {approved|pending|none}
Review budget lines: {number}
Review budget risk: {Low|Medium|High}
400-line budget risk: {Low|Medium|High}

### Next Step
{Ready for implementation (sdd-apply) | Decision needed before apply — orchestrator must resolve delivery strategy before sdd-apply.}
```

## Rules

- ALWAYS reference concrete file paths in tasks
- Tasks MUST be ordered by dependency — Phase 1 tasks shouldn't depend on Phase 2
- Testing/evidence tasks should reference planned case IDs from `test-design.md` plus specific scenarios from the specs
- Security evidence tasks should reference guideline IDs from `design.md#secure-development-design` plus planned case IDs from `test-design.md`
- Mandatory `test-design.md` cases MUST NOT be omitted. If they cannot be represented in implementation tasks, return `blocked` and name the missing cases.
- Mandatory security controls MUST NOT be omitted. Missing mandatory evidence without a complete approved exception blocks task creation.
- Each task should be completable in ONE session (if a task feels too big, split it)
- Use hierarchical numbering: 1.1, 1.2, 2.1, 2.2, etc.
- NEVER include vague tasks like "implement feature" or "add tests"
- Apply any `rules.tasks` from `openspec/config.yaml`
- If Strict TDD is active, integrate test-first tasks: RED task (write failing test) → GREEN task (make it pass) → REFACTOR task (clean up)
- **Size budget**: Tasks artifact MUST be under 530 words. Each task: 1-2 lines max. Use checklist format, not paragraphs.
- **Review workload guard**: ALWAYS include the Review Workload Forecast. If likely above the session review budget, recommend chained PRs and honor the received delivery strategy for whether a decision/exception is needed before apply. If delivery strategy is `null`, surface apply-time decisions through the forecast; do not invent a strategy.
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the tasks summary belongs in `detailed_report`.
