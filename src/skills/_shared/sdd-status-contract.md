# SDD Status and Instructions Contract

Shared status contract for SDD commands and phase skills. Use this before acting on a change so orchestration does not guess state, artifact references, or edit scope.

## Purpose

Commands that select, continue, apply, verify, or archive an SDD change MUST first produce or consume structured status. The status is the handoff between orchestrator and phase executor.

## Change Selection

- If a change name is provided, use that exact change after confirming it exists in the selected artifact store.
- If no change name is provided, infer only when the active change is unambiguous from session state or there is exactly one active change.
- If multiple active changes match or the active change is unclear, ask the user to choose. Do not guess.
- If no active changes exist, report that no SDD change is active and suggest `/sdd-new <change>`.

## Native Engine

- This local status contract is the source of truth for this repository's complete SDD DAG. Status producers MUST include active phases `test-design`, `review`, and `review-security` before `verify` and `archive`.
- The `gentle-ai` binary is optional advisory input. When available, `gentle-ai sdd-status [change] --cwd <repo> --json --instructions` or `gentle-ai sdd-continue [change] --cwd <repo>` MAY be used to obtain compact OpenSpec status, artifact paths, task progress, and blocker hints.
- Native `gentle-ai` status MUST NOT override valid persisted state or this local contract. If native output is missing active repository phases, omits required review artifacts, routes `design -> tasks`, routes `apply -> verify`, or otherwise conflicts with the dependency rules below, treat the native route as incomplete and reconstruct/normalize status from local artifacts and persisted state.
- When `blockedReasons` is non-empty, do not proceed to terminal, archive, or apply work. Return or report `blockedReasons` and stop unless `nextRecommended` is `verify`, in which case verification may run only to remediate or refresh evidence for the blockers. When `nextRecommended` is `resolve-blockers`, always report `blockedReasons` and stop.
- `nextRecommended` is a bounded machine token for routing, not human prose. Route only by the locally normalized `nextRecommended` and dependency states.
- Human-readable explanation belongs in `blockedReasons`, not `nextRecommended`.
- If the binary is unavailable or incomplete for the active DAG, fall back to this prompt contract and the manual status schema below. Manual fallback status MUST stay shape-compatible with `sdd.status` JSON where fields overlap, while preserving this repository's additional phase fields.

## Routing Token Mapping

Native status uses bounded tokens for new changes: `propose`, `spec`, `design`, `test-design`, `tasks`, `apply`, `review`, `review-security`, `verify`, `archive`, `sdd-new`, `select-change`, `resolve-blockers`, and `none`. Legacy `security-design` and `security-applicability` may appear only when reading old or archived state and MUST NOT be emitted as active new-change successors.

When launching phase agents, normalize through this mapping:

| `nextRecommended` | Launch |
| --- | --- |
| `propose` | `sdd-propose` |
| `spec` | `sdd-spec` |
| `security-applicability` | No launch target; legacy/archive data compatibility only |
| `design` | `sdd-design` |
| `security-design` | No launch target for new changes; legacy/archive compatibility only |
| `test-design` | `sdd-test-design` |
| `tasks` | `sdd-tasks` |
| `apply` | `sdd-apply` |
| `review` | `sdd-review` |
| `review-security` | `sdd-review-security` |
| `verify` | `sdd-verify` |
| `archive` | `sdd-archive` |
| `sdd-new` | orchestrator workflow |
| `select-change` | ask user to choose |
| `resolve-blockers` | report blockers and stop |
| `none` | no next phase |

Post-resolution rule for `resolve-blockers`: after blocker resolution, the orchestrator MUST re-delegate per Blocker Resolution Protocol in `executor-boundary.md`.

Phase envelopes may return prefixed phase tokens such as `sdd-verify`; consumers MUST normalize them to the native token before comparing dependency states.

## Field Naming Across Contracts

| Contract | Field style | Routing field | Phase field | Completed phases field |
| --- | --- | --- | --- | --- |
| Phase result envelope | snake_case | `next_recommended` | N/A | N/A |
| Native/status JSON | camelCase | `nextRecommended` | dependency states | N/A |
| Persisted SDD state | camelCase | `nextRecommended` | `currentPhase` | `completedPhases` |

Rules:

- Phase agents may return `next_recommended` because `sdd-phase-common.md` defines the phase envelope.
- Status and state artifacts MUST use camelCase fields.
- Before routing or persisting state, normalize phase envelope `next_recommended` into status/state `nextRecommended`.
- Do not write `next_recommended`, `current_phase`, or `completed_phases` into persisted state.
- If legacy snake_case state is encountered, normalize it in memory and rewrite only camelCase state after validation succeeds.

## Status Schema

Return status as markdown with these fields, or as equivalent JSON when the host supports it:

```yaml
schemaName: sdd.status
schemaVersion: 2
changeName: <change-name-or-null>
artifactStore: engram | openspec
planningHome:
  mode: repo-local | memory-local
  path: <absolute path to openspec, memory namespace, or null>
changeRoot: <absolute path to openspec/changes/<change>, memory namespace, or null>
artifactRefs:
  explore: [<topic keys, file paths, or inline refs>]
  proposal: [<topic keys, file paths, or inline refs>]
  specs: [<topic keys, file paths, or inline refs>]
  securityApplicability: [<topic keys, file paths, or inline refs>] # legacy/read-only
  design: [<topic keys, file paths, or inline refs>]
  securityDesign: [<topic keys, file paths, or inline refs>] # legacy/read-only
  testDesign: [<topic keys, file paths, or inline refs>]
  tasks: [<topic keys, file paths, or inline refs>]
  applyProgress: [<topic keys, file paths, or inline refs>]
  reviewReport: [<topic keys, file paths, or inline refs>]
  securityReviewReport: [<topic keys, file paths, or inline refs; canonical JSON first when present, derived Markdown second>]
  verifyReport: [<topic keys, file paths, or inline refs>]
  state: [<topic keys, file paths, or inline refs>]
artifactPaths:
  explore: [<absolute path>]
  proposal: [<absolute path>]
  specs: [<absolute paths>]
  securityApplicability: [<absolute path>] # legacy/read-only
  design: [<absolute path>]
  securityDesign: [<absolute path>] # legacy/read-only
  testDesign: [<absolute path>]
  tasks: [<absolute path>]
  applyProgress: [<absolute path>]
  reviewReport: [<absolute path>]
  securityReviewReport: [<absolute path; canonical JSON first when present, derived Markdown second>]
  verifyReport: [<absolute path>]
  state: [<absolute path>]
contextFiles:
  explore: [<absolute readable files>]
  proposal: [<absolute readable files>]
  specs: [<absolute readable files>]
  securityApplicability: [<absolute readable files>] # legacy/read-only
  design: [<absolute readable files>]
  securityDesign: [<absolute readable files>] # legacy/read-only
  testDesign: [<absolute readable files>]
  tasks: [<absolute readable files>]
  applyProgress: [<absolute readable files>]
  reviewReport: [<absolute readable files>]
  securityReviewReport: [<absolute readable files; canonical JSON first when present, derived Markdown second>]
  verifyReport: [<absolute readable files>]
  state: [<absolute readable files>]
artifacts:
  explore: missing | done | partial
  proposal: missing | done | partial
  specs: missing | done | partial
  securityApplicability: missing | done | partial | legacy
  design: missing | done | partial
  securityDesign: missing | done | partial | legacy
  testDesign: missing | done | partial
  tasks: missing | done | partial
  applyProgress: missing | done | partial
  reviewReport: missing | done | partial
  securityReviewReport: missing | done | partial
  verifyReport: missing | done | partial
  state: missing | done | partial
taskProgress:
  total: 0
  completed: 0
  pending: 0
  allComplete: false
dependencies:
  proposal: blocked | ready | all_done
  specs: blocked | ready | all_done
  securityApplicability: blocked | ready | all_done | legacy
  design: blocked | ready | all_done
  securityDesign: blocked | ready | all_done | legacy
  testDesign: blocked | ready | all_done
  tasks: blocked | ready | all_done
  apply: blocked | ready | all_done
  review: blocked | ready | all_done
  reviewSecurity: blocked | ready | all_done
  verify: blocked | ready | all_done
  archive: blocked | ready | all_done
applyState: blocked | all_done | ready
actionContext:
  mode: repo-local | workspace-planning | memory-local | ephemeral
  workspaceRoot: <absolute path>
  allowedEditRoots: [<absolute paths>]
relationships:
  dependsOn: []
  supersedes: []
  amends: []
  conflictsWith: []
  sameDomainActiveChanges: []
phaseInstructions:
  apply: [<instruction strings>]
  review: [<instruction strings>]
  verify: [<instruction strings>]
  archive: [<instruction strings>]
nextRecommended: propose | spec | design | test-design | tasks | apply | review | review-security | verify | archive | sdd-new | select-change | resolve-blockers | none
blockedReasons: []
```

`phaseInstructions` is optional and appears only when instructions are requested. `nextRecommended` is a required top-level routing token, never nested under `phaseInstructions`. Empty array fields MUST be arrays, not null. `changeName`, `planningHome.path`, and `changeRoot` are nullable; all other sections should be present in fallback output so consumers can parse native and manual status the same way.

Mode-specific status rules:

- **`engram`**: `artifactRefs` contain Engram topic keys. `artifactPaths` and `contextFiles` are empty arrays unless a real filesystem path is also known.
- **`openspec`**: `artifactRefs`, `artifactPaths`, and `contextFiles` contain OpenSpec file paths where applicable.


Native status JSON is advisory when available. If native currently emits only `artifactStore: openspec`, use that output only as compact evidence for OpenSpec-mode artifact paths/task progress, then normalize through this local contract. Manual fallback MUST use `artifactStore` and `artifactRefs` from the active session preflight for non-OpenSpec modes.

## Apply State

- `blocked`: Required apply artifacts are missing, task selection is ambiguous, or action context makes edits unsafe.
- `all_done`: Tasks artifact exists and every implementation task is checked `[x]`.
- `ready`: Tasks artifact exists, at least one implementation task remains unchecked, and edit scope is safe.

## Dependency States

- `proposal`, `specs`, `design`, `testDesign`, and `tasks` report whether prerequisite artifacts are blocked, ready, or all done. `securityDesign` and `securityApplicability` are legacy/read-only for old or archived changes and MUST NOT be active new-change dependencies.
- `design` is `ready` only when proposal and specs are available. Missing `securityApplicability` or standalone `securityDesign` MUST NOT block technical design for new changes. A completed new-change design MUST contain `## Secure Development Design`.
- `securityDesign` is legacy/read-only for new changes. If present in old or archived state, it may be displayed as historical evidence but MUST NOT be required or launched as an active dependency.
- Historical `security-design` and `security-applicability` routing values MAY be recognized only when reading old or archived state. They MUST NOT be emitted for new-change status, normalized into runnable successors, mapped to launchable agents, active dependencies, or active security authority.
- `testDesign` is `ready` only when proposal, specs, and design with `## Secure Development Design` narrative rules are available; it is `all_done` when the `test-design` artifact exists and is readable.
- `tasks` is `ready` only when specs, design with `## Secure Development Design` narrative rules, and test design are available. Missing `testDesign` blocks task planning.
- `apply` is `ready` only when specs, design with `## Secure Development Design` narrative rules, test design, and tasks are available and task progress is not all done.
- `review` is `ready` when tasks exist, test design is available, embedded design security evidence is available, and either apply-progress exists or the tasks artifact shows all intended implementation work complete. Incomplete tasks remain blockers for full review.
- `reviewSecurity` is `ready` when `design.md#secure-development-design` exists, all intended implementation work is complete, and `reviewReport` exists with a non-blocking verdict. Missing, blocking, or unreadable general review evidence blocks security review.
- `verify` is `ready` when tasks exist, test design is available, embedded design security evidence is available, all intended implementation work is complete, `reviewReport` exists with a non-blocking verdict, and `securityReviewReport` exists with a non-blocking verdict. When canonical security-review JSON exists or is expected for current contracts, it is authoritative and derived Markdown must not be stale/parity-failed. Missing or blocking review evidence blocks full verification.
- `archive` is `ready` only when the Archive Readiness contract in `skills/_shared/sdd-post-apply-gates.md#archive-readiness` passes. Status evaluation MUST still confirm the required non-blocking review reports, passing verify report, completed task evidence, mandatory artifacts, security evidence or approved exceptions, and absence of blocker/negation signals before setting archive ready.

## Phase Routing Order

The active new-change DAG is `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`; standalone `security-design` is legacy/read-only. The implementation sub-DAG is `apply -> review -> review-security -> verify -> archive`.

- Completed apply work MUST recommend `review`, not direct `verify`.
- Non-blocking review evidence MUST recommend `review-security`, not direct `verify`.
- Blocking review findings MUST recommend `apply` and identify failed controls plus affected requirements.
- Non-blocking security review evidence MUST recommend `verify`.
- Blocking security review findings MUST recommend `apply` or `resolve-blockers` and identify affected guideline IDs, matrix rows, evidence gaps, and exceptions.
- Missing review inputs, unknown changed-file context, unsafe workspace context, or review-report persistence failure MUST recommend `resolve-blockers`.
- Archive readiness MUST follow `skills/_shared/sdd-post-apply-gates.md#archive-readiness` and require both non-blocking review reports plus a passing verify report.

## Action Context Guard

The orchestrator MUST carry `actionContext` into any phase launch.

- If manually reconstructed context cannot prove edit ownership or allowed edit roots, stop before editing.
- If `allowedEditRoots` is present, only edit files within those roots.
- If a command cannot prove a file is inside the authoritative workspace or allowed edit roots, stop and ask for clarification.

## Status Output

Every command that acts on a change MUST show status before launching an executor or performing archive work:

- Active change selection and how it was resolved.
- Artifact statuses and paths/topics used as context.
- Task progress and unchecked task list when tasks exist.
- Next recommended action.
- `blockedReasons` when `nextRecommended` is not `verify`, plus any edit-root blockers.
