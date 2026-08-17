# Persistence Contract (shared across all SDD skills)

This file is the authoritative SDD persistence contract. SDD phase skills and shared execution contracts MUST delegate detailed artifact-store mode behavior, backend read/write semantics, artifact reference resolution, state persistence, and persistence verification to this file.

Phase skills MAY define phase-local artifact obligations such as required inputs, produced artifacts, mutations, conditional behavior, validation, and routing. They MUST NOT redefine the common persistence mechanics below.

## Mode Resolution

The orchestrator passes `artifact_store.mode` with one of: `engram | openspec`.

The orchestrator asks/caches artifact mode during minimal preflight before the first mutating SDD route in a session. Status/read-only routes may bypass preflight and must not mutate artifacts just to resolve mode.

User-facing artifact options are OpenSpec and Engram. Map them internally to `openspec` and `engram` respectively.

Default (if user doesn't specify): `openspec`.

## Mode Roles

- **`engram`**: Working memory between sessions. Upserts overwrite - no iteration history. Local only, not shareable.
- **`openspec`**: Source of truth. Files in repo, git history, team-shareable, full audit trail.

### Mode Comparison

| Capability | `engram` | `openspec` |
|------------|----------|------------|
| Cross-session recovery | Yes | No (needs git) |
| Compaction survival | Yes | No |
| Shareable with team | No (local DB) | Yes (committed files) |
| Full iteration history | No (upsert overwrites) | Yes (git history) |
| Audit trail (archive) | Partial (report only) | Yes (full folder) |
| Project files created | Never | Yes |

### `engram` mode limitation

Engram uses `topic_key`-based upserts. Re-running a phase for the same change **overwrites** the previous version - no revision history is kept. The archive phase saves a summary report, not the full artifact folder. For iteration history or team collaboration, use `openspec`.

## Behavior Per Mode

| Mode | Read from | Write to | Project files |
|------|-----------|----------|---------------|
| `engram` | Engram | Engram | Never |
| `openspec` | Filesystem | Filesystem | Yes |

Mode behavior is intentionally centralized here. If a phase-specific instruction appears to conflict with this table, keep the phase-local artifact obligation only and resolve backend behavior from this contract.

## Artifact Reference Resolver

Use this resolver whenever checking dependencies, launching sub-agents, validating artifacts, continuing a change, or recovering state.

| Artifact | `engram` reference | `openspec` reference |
| --- | --- | --- |
| Project context | `sdd-init/{project}` | `openspec/config.yaml` |
| Testing capabilities | `sdd/{project}/testing-capabilities` | `openspec/config.yaml` |
| State | `sdd/{change-name}/state` | `openspec/changes/{change-name}/state.yaml` |
| Exploration | `sdd/{change-name}/explore` | `openspec/changes/{change-name}/explore.md` |
| Proposal | `sdd/{change-name}/proposal` | `openspec/changes/{change-name}/proposal.md` |
| Spec | `sdd/{change-name}/spec` | `openspec/changes/{change-name}/specs/{domain}/spec.md` |
| Security applicability (historical/read-only) | `sdd/{change-name}/security-applicability` | `openspec/changes/{change-name}/security-applicability.md` |
| Design | `sdd/{change-name}/design` | `openspec/changes/{change-name}/design.md` |
| Security design (historical/read-only) | `sdd/{change-name}/security-design` | `openspec/changes/{change-name}/security-design.md` |
| Test design | `sdd/{change-name}/test-design` | `openspec/changes/{change-name}/test-design.md` |
| Test cases | `sdd/{change-name}/test-cases` | `openspec/changes/{change-name}/test-cases.json` |
| Tasks | `sdd/{change-name}/tasks` | `openspec/changes/{change-name}/tasks.md` |
| Apply progress | `sdd/{change-name}/apply-progress` | `openspec/changes/{change-name}/tasks.md` checkbox state plus status evidence |
| Review report | Canonical JSON: `sdd/{change-name}/review-report.json`; derived Markdown compatibility view: `sdd/{change-name}/review` | Canonical JSON: `openspec/changes/{change-name}/review-report.json`; derived Markdown compatibility file: `openspec/changes/{change-name}/review-report.md` |
| Security review report | Canonical JSON: `sdd/{change-name}/review-security-report.json`; derived Markdown compatibility view: `sdd/{change-name}/review-security` | Canonical JSON: `openspec/changes/{change-name}/review-security-report.json`; derived Markdown compatibility file: `openspec/changes/{change-name}/review-security-report.md` |
| Verify report | `sdd/{change-name}/verify-report` | `openspec/changes/{change-name}/verify-report.md` |
| Archive report | `sdd/{change-name}/archive-report` | `openspec/changes/archive/YYYY-MM-DD-{change-name}/` |

Resolver rules:

- `engram`: use `mem_search` to find the topic key, then `mem_get_observation` for full content.
- `openspec`: read/write only paths defined by `openspec-convention.md`.

Resolver verification:

- Treat `mem_search` results as previews only; call `mem_get_observation` before using Engram artifact content.
- In `openspec`, read artifacts from the paths in `openspec-convention.md` or the structured status artifact paths. Do not infer alternate paths.
- For review evidence, downstream phases MUST resolve exactly one selected-backend identity for each required review report.
- Canonical general-review JSON wins over derived Markdown whenever both are present.
- Canonical security-review JSON wins over derived Markdown whenever both are present.
- The review phase MUST route to `resolve-blockers` when required artifacts, changed-file context, safe workspace context, or review-report persistence evidence are missing.

## State Persistence (Orchestrator)

The orchestrator persists DAG state after each phase transition to enable SDD recovery after compaction.

State is an index and recovery pointer, not a replacement for artifacts. It points to artifact refs/paths and summarizes routing state; do not duplicate full proposal/spec/design/tasks bodies into state.

Persisted SDD state uses camelCase fields. Phase result envelopes may return `next_recommended`, but the orchestrator MUST normalize that value into state `nextRecommended` before writing state. Do not write `next_recommended`, `current_phase`, or `completed_phases` into state artifacts. If legacy snake_case state is encountered, normalize it in memory and rewrite only camelCase state after validation succeeds.

Persist state after:

- Every phase `success` before launching the next phase
- Every `partial` result that produced useful artifacts, recovery steps, or changed task/apply/archive progress
- Every `blocked` result that creates or changes `blockedReasons`, required decisions, selected change, chain plan, or recovery instructions
- Every delivery decision that changes `deliveryStrategy`, `chainStrategy`, `chainPlanRef`, `sizeException`, or `reviewBudgetLines`

| Mode | Persist State | Recover State |
|------|--------------|---------------|
| `engram` | `mem_save(topic_key: "sdd/{change-name}/state", capture_prompt: false*)` | `mem_search("sdd/*/state")` -> `mem_get_observation(id)` |
| `openspec` | Write `openspec/changes/{change-name}/state.yaml` | Read `openspec/changes/{change-name}/state.yaml` |

*For state automated artifacts, set `capture_prompt: false` when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

Minimum state fields:

```yaml
schemaName: sdd.state
schemaVersion: 1
changeName: {change-name}
artifactStore: engram | openspec
currentPhase: explore | propose | spec | design | test-design | tasks | apply | review | review-security | review-parallel | verify | archive | blocked | complete
completedPhases: []
  artifactRefs:
  explore: []
  proposal: []
  specs: []
  securityApplicability: []
  design: []
  securityDesign: []
  testDesign: []
  tasks: []
  applyProgress: []
  reviewReport: []
  securityReviewReport: []
  verifyReport: []
  archiveReport: []
  state: []
delivery:
  deliveryStrategy: ask-on-risk | auto-chain | single-pr | exception-ok | null
  reviewBudgetLines: {number|null}
  chainStrategy: stacked-to-main | feature-branch-chain | null
  chainPlanRef: {ref|null}
  sizeException:
    approved: true | false
    approver: {name-or-null}
    rationale: {text-or-null}
nextRecommended: propose | spec | design | test-design | tasks | apply | review | review-security | verify | archive | sdd-new | select-change | resolve-blockers | none
blockedReasons:
  - code: {machine-readable-code}
    message: {human-readable-summary}
    owner: orchestrator | user | phase-agent | reviewer
    requiredAction: {next-action}
stateRevision: {monotonic-integer-or-iso-timestamp}
updatedAt: {iso-8601-timestamp}
```

State update order: validate the phase result, persist/update state in the selected backend, read state back for verification, then launch the next phase or report completion. If state persistence fails in `engram` or `openspec`, stop before dependent work.

## State Recovery (Orchestrator)

Use this when recovering after compaction, mid-session context drift, or `/sdd-continue` when state is not already trusted.

Recovery sources by mode:

| Mode | Recovery source |
|------|-----------------|
| `engram` | `mem_search(query: "sdd/*/state" or "sdd/{change-name}/state", project: "{project}")` -> `mem_get_observation(id)` |
| `openspec` | Read `openspec/changes/{change-name}/state.yaml` after selecting or confirming the active change |

## Artifact Persistence Verification

Every phase that produces or mutates a persistent artifact MUST verify the selected backend before reporting success.

| Mode | Required verification |
|------|-----------------------|
| `engram` | Save with the expected `topic_key`, then retrieve the observation by ID or topic and confirm the artifact identity and content needed by downstream phases. |
| `openspec` | Re-read the written file path and confirm the expected artifact identity, required sections, and task/status mutations are present. |

Verification rules:

- Do not report `success` for persistent modes if the artifact cannot be read back from the selected backend.
- For `sdd-apply`, completed work MUST be visible in the selected task/progress artifact before returning success or partial completion.
- For `sdd-review`, canonical `review-report.json` and derived Markdown MUST be readable after persistence in persistent modes.
- For `sdd-review-security`, canonical `review-security-report.json` and derived Markdown MUST be readable after persistence in persistent modes.
- For `sdd-archive`, archive movement and spec synchronization MUST be verified from the filesystem and/or Engram topic keys selected by mode.

## Common Rules

- `engram` -> do NOT write any project files; persist to Engram and return observation IDs
- `openspec` -> write files ONLY to paths defined in `openspec-convention.md`
- NEVER force `openspec/` creation unless orchestrator explicitly passed `openspec`
- If unsure which mode to use, default to `openspec`

## Cross-Contract Boundaries

This file owns persistence mechanics only. Keep adjacent orchestration and executor behavior in their dedicated contracts:

| Concern | Source of truth |
| --- | --- |
| SDD phase context, dependency retrieval, artifact persistence handoff, and final response ordering | `skills/_shared/sdd-phase-common.md` |
| Supplemental skill registry lookup, `## Skills to load before work`, and `skill_resolution` reporting | `skills/_shared/skill-resolver.md` |
| OpenSpec paths, folder layout, and delta-spec file conventions | `skills/_shared/openspec-convention.md` |
| Engram topic naming and Engram-specific call examples | `skills/_shared/engram-convention.md` |

Persistence-specific reminders retained here:

- SDD phase agents read required artifacts from the selected backend using exact refs/paths supplied by status or launch context; they must not reconstruct missing artifacts from prose.
- Every persistent SDD artifact write must follow the selected mode and pass read-back verification before a phase reports `status: success`.
- For automated SDD artifacts saved to Engram, use `capture_prompt: false` when the tool schema supports it; omit the field rather than failing on older schemas.
