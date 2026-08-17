# SDD Phase — Common Protocol

Boilerplate identical across all SDD phase skills. Sub-agents MUST load this alongside their phase-specific SKILL.md.

Executor boundary: every SDD phase agent is an EXECUTOR, not an orchestrator. Do the phase work yourself. Do NOT launch sub-agents, do NOT call `delegate`/`task`, and do NOT bounce work back unless the phase skill explicitly says to stop and report a blocker.

## Launch Envelope Contract

Every SDD phase executor receives a structured `launch` envelope from the orchestrator before task-specific instructions. Treat it as the authoritative launch metadata for this run.

```yaml
launch:
  phase: sdd-<phase>
  changeName: {change-name}
  artifact_store:
    mode: engram | openspec
  execution_mode: interactive | auto
  delivery_strategy: ask-on-risk | auto-chain | single-pr | exception-ok | null
  chain_strategy: stacked-to-main | feature-branch-chain | null
  model_assignment:
    requested_agent: sdd-<phase>
    resolved_model: {provider/model-or-runtime-default-or-unknown}
    resolved_model_source: explicit-phase | explicit-profile | orchestrator-synthesis | runtime-default | unknown
    escalation_reason: {reason-or-null}
  status:
    nextRecommended: {token}
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
    blockedReasons: []
  artifacts:
    refs:
      projectContext: []
      testingCapabilities: []
      state: []
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
    paths:
      projectContext: []
      testingCapabilities: []
      state: []
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
  actionContext:
    mode: repo-local | workspace-planning | memory-local | ephemeral
    workspaceRoot: {absolute-path-or-null}
    allowedEditRoots: []
  review:
    review_budget_lines: {number|null}
    current_slice_boundary: {text-or-null}
    size_exception: {approved/evidence-or-null}
  skill_paths:
    - {absolute SKILL.md path}
```

## Launch Envelope Examples

Use these as canonical examples when constructing launch payloads. The orchestrator must still resolve real refs/paths from the selected backend before delegation.

### Example A — `openspec` launch (design)

```yaml
launch:
  phase: sdd-design
  changeName: fix-blocker-redelegation-protocol
  artifact_store:
    mode: openspec
  execution_mode: interactive
  delivery_strategy: null
  chain_strategy: null
  model_assignment:
    requested_agent: sdd-design
    resolved_model: runtime-default
    resolved_model_source: runtime-default
    escalation_reason: null
  status:
    nextRecommended: design
    dependencies:
      proposal: all_done
      specs: all_done
      securityApplicability: legacy
      design: ready
      securityDesign: legacy
      testDesign: blocked
      tasks: blocked
      apply: blocked
      review: blocked
      reviewSecurity: blocked
      verify: blocked
      archive: blocked
    blockedReasons: []
  artifacts:
    refs:
      projectContext:
        - openspec/config.yaml
      testingCapabilities:
        - openspec/config.yaml
      state:
        - openspec/changes/fix-blocker-redelegation-protocol/state.yaml
      explore:
        - openspec/changes/fix-blocker-redelegation-protocol/explore.md
      proposal:
        - openspec/changes/fix-blocker-redelegation-protocol/proposal.md
      specs:
        - openspec/changes/fix-blocker-redelegation-protocol/spec.md
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
    paths:
      projectContext:
        - openspec/config.yaml
      testingCapabilities:
        - openspec/config.yaml
      state:
        - openspec/changes/fix-blocker-redelegation-protocol/state.yaml
      explore:
        - openspec/changes/fix-blocker-redelegation-protocol/explore.md
      proposal:
        - openspec/changes/fix-blocker-redelegation-protocol/proposal.md
      specs:
        - openspec/changes/fix-blocker-redelegation-protocol/spec.md
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
  actionContext:
    mode: repo-local
    workspaceRoot: C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai
    allowedEditRoots:
      - C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai\\openspec\\changes\\fix-blocker-redelegation-protocol
  review:
    review_budget_lines: null
    current_slice_boundary: null
    size_exception: null
  skill_paths:
    - C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai\\skills\\project-conventions\\SKILL.md
```

### Example B — `engram` launch (design)

```yaml
launch:
  phase: sdd-design
  changeName: fix-blocker-redelegation-protocol
  artifact_store:
    mode: engram
  execution_mode: interactive
  delivery_strategy: null
  chain_strategy: null
  model_assignment:
    requested_agent: sdd-design
    resolved_model: runtime-default
    resolved_model_source: runtime-default
    escalation_reason: null
  status:
    nextRecommended: design
    dependencies:
      proposal: all_done
      specs: all_done
      securityApplicability: legacy
      design: ready
      securityDesign: legacy
      testDesign: blocked
      tasks: blocked
      apply: blocked
      review: blocked
      reviewSecurity: blocked
      verify: blocked
      archive: blocked
    blockedReasons: []
  artifacts:
    refs:
      projectContext:
        - sdd-init/leo3nava_ai
      testingCapabilities:
        - sdd/leo3nava_ai/testing-capabilities
      state:
        - sdd/fix-blocker-redelegation-protocol/state
      explore:
        - sdd/fix-blocker-redelegation-protocol/explore
      proposal:
        - sdd/fix-blocker-redelegation-protocol/proposal
      specs:
        - sdd/fix-blocker-redelegation-protocol/spec
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
    paths:
      projectContext: []
      testingCapabilities: []
      state: []
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
  actionContext:
    mode: repo-local
    workspaceRoot: C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai
    allowedEditRoots:
      - C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai\\openspec\\changes\\fix-blocker-redelegation-protocol
  review:
    review_budget_lines: null
    current_slice_boundary: null
    size_exception: null
  skill_paths:
    - C:\\Leo\\Proyectos\\leo3nava_ai\\leo3nava_ai\\skills\\project-conventions\\SKILL.md
```

Consumption rules for executors:

- Read artifacts from `artifacts.refs` / `artifacts.paths` according to `artifact_store.mode`; do not reconstruct missing artifacts from prose or broad memory searches.
- Treat `status`, `dependencies`, `blockedReasons`, delivery/chain fields, review fields, and `actionContext` as launch constraints. If they are inconsistent with readable artifacts, return `blocked` instead of guessing.
- Respect `actionContext.allowedEditRoots` for any repository operation. If roots are missing or unsafe for the requested phase, return `blocked` before reading broadly, testing, editing, verifying, or archiving.
- Load supplemental skills from `skill_paths` using Section A. These are supplemental only; do not launch sub-agents or call skill tools to perform phase work.
- Unknown or unresolved fields must remain `null` or `unknown`. Do not invent model IDs, artifact refs, workspace roots, delivery decisions, review approvals, or missing state.

## A. Skill Loading

`skills/_shared/skill-resolver.md` is the source of truth for supplemental skill lookup, path injection, fallback loading, and `skill_resolution` reporting.

Executor minimum:

1. Prefer exact `skill_paths` / `## Skills to load before work` injected by the orchestrator.
2. Read those exact `SKILL.md` files before task-specific work.
3. If no paths were injected, follow the fallback rules in `skill-resolver.md`.
4. Report `skill_resolution` using the shape and acceptance semantics from `skill-resolver.md#step-4-report-resolution`.

Loading supplemental skills is not delegation. SDD phase executors still execute their own phase and MUST NOT launch sub-agents.

## B. Artifact Retrieval

Follow `skills/_shared/persistence-contract.md` for artifact-store mode resolution, artifact references, backend read behavior, Engram preview handling, and OpenSpec paths.

Phase skills remain responsible for naming their required inputs and reading every required dependency before producing phase output.

## B1. SDD Phase Context Contract

SDD phase agents read artifacts directly from the selected backend using exact refs from the launch envelope. They do not perform broad memory searches, infer active changes, or reconstruct missing artifacts from prose.

Mode-specific context rules:

- `engram`: read only the provided topic keys via `mem_search` / `mem_get_observation`; persist completed artifacts with `capture_prompt: false` when supported.
- `openspec`: read/write only the provided OpenSpec paths and paths defined by `skills/_shared/openspec-convention.md`.

Required context by phase:

| Phase | Required refs/context | Optional refs/context | Writes | Block if missing |
| --- | --- | --- | --- | --- |
| `sdd-explore` | user request, project context, testing capabilities when available | related prior context summary | `explore` | project context is unavailable and exploration would speculate |
| `sdd-propose` | user request or explore result sufficient to avoid material speculation | answered proposal questions, prior product decisions | `proposal` | product/business facts are missing and would require guessing |
| `sdd-spec` | proposal | explore, product assumptions | `spec` | proposal missing/unreadable |
| `sdd-design` | proposal + spec | architecture conventions, related files summary, baseline security considerations | `design` | proposal/spec missing/unreadable |
| `sdd-test-design` | proposal + spec + design with `## Secure Development Design` | testing capabilities, design risks, embedded security controls | `test-design` | proposal/spec/design or embedded secure design section missing/unreadable |
| `sdd-tasks` | spec + design with `## Secure Development Design` narrative rules + test-design | proposal, review budget, delivery/chain preferences | `tasks` | spec/design/test-design missing/unreadable |
| `sdd-apply` | tasks + spec + design with `## Secure Development Design` narrative rules + test-design + actionContext + Review Workload Guard result | apply-progress, chain plan, strict TDD instructions | `apply-progress` | required artifacts, safe edit roots, or review guard are missing |
| `sdd-review` | design with `## Secure Development Design` narrative rules + test-design + tasks + apply-progress/evidence + changed-file context + actionContext | review catalog, changed files summary | `review-report` | review inputs, changed-file context, safe workspace context, or persistence capability is missing |
| `sdd-review-security` | design with `## Secure Development Design` narrative rules + apply-progress/evidence + changed-file context | tasks, test-design, catalog | `review-security-report` canonical JSON plus derived Markdown | embedded security design, changed-file context, or persistence capability is missing |

> **Parallel execution**: `sdd-review` and `sdd-review-security` run in parallel after `sdd-apply`. Both consume the same base inputs independently. The orchestrator MUST NOT launch `sdd-verify` until both phases appear in `completedPhases`.
| `sdd-verify` | spec + design with `## Secure Development Design` narrative rules + test-design + tasks + apply-progress/evidence + non-blocking review-report canonical JSON/derived Markdown identity + non-blocking review-security-report canonical JSON/derived Markdown identity + testing capabilities | strict TDD evidence, changed files summary | `verify-report` | review/security-review evidence, verification evidence, or required artifacts are missing |
| `sdd-archive` | proposal + spec + design with `## Secure Development Design` narrative rules + test-design + tasks + apply-progress + non-blocking review-report canonical JSON/derived Markdown identity + non-blocking review-security-report canonical JSON/derived Markdown identity + verify-report + state | chain plan, size exception, stale-checkbox reconciliation approval, partial archive exception | `archive-report` | review-report/review-security-report missing/blocking, verify-report missing/non-passing, mandatory security evidence missing, or required artifacts unavailable |

Context integrity checks:

- Confirm `changeName`, `artifact_store.mode`, `currentPhase`/`nextRecommended`, and `stateRevision` when provided before doing phase work.
- If an artifact ref points to a different change, stale state, wrong mode, or missing backend, return `blocked`; do not silently continue.
- If `apply-progress`, tasks, chain plan, review budget, or actionContext changed compared with the launch envelope, return `blocked` or report the mismatch in the result envelope.
- If a required dependency is missing, do not create a placeholder artifact downstream. Route back to the owning earlier phase.

## C. Artifact Persistence

Every phase that produces or mutates an artifact MUST persist and verify it according to `skills/_shared/persistence-contract.md`. Skipping persistence or read-back verification BREAKS the pipeline because downstream phases cannot trust the artifact references.

Phase skills remain responsible for defining the artifact key/path, artifact content, phase-local mutations, and any conditional persistence requirements.

## D. Return Envelope

> **CRITICAL — Response ordering**: Your FINAL output MUST be a fenced YAML block and nothing else: open with ````yaml` on its own line, output the full envelope, then close with ``` on its own line. Do NOT output any text, prose, headers, or analysis before the opening fence or after the closing fence. All analysis belongs inside `detailed_report`. Do NOT end on a tool call; if you need to save to Engram (`mem_save`), do it BEFORE your final text response. Do NOT call `mem_session_summary` — that's for top-level agents only. **Why**: When a sub-agent's last action is a tool call, the parent agent receives only the tool result — your text response (the actual analysis) is lost. When a sub-agent wraps the envelope in prose, the orchestrator receives unstructured text instead of a parseable envelope.

Every phase MUST return a structured envelope to the orchestrator. This section is the authoritative phase result contract.

Required envelope fields:

| Field | Required shape / meaning |
| --- | --- |
| `status` | One of `success`, `partial`, or `blocked` only. Do not invent values such as `hold`, `failed`, `done`, or `ok`. |
| `executive_summary` | 1-3 sentence human summary of what happened and why it matters. |
| `detailed_report` | Full phase output, or an explicit note that the full output is already inline and intentionally small. Required for every phase result. |
| `artifacts` | Array/list of artifacts produced, updated, read, or intentionally omitted. Use the artifact entry shape below. |
| `next_recommended` | Next bounded routing token or phase token, normalized by the orchestrator through `skills/_shared/sdd-status-contract.md`. |
| `risks` | Structured risk entries using the shape below, or `None`. |
| `skill_resolution` | How supplemental skills were loaded. Return the structured shape from `skills/_shared/skill-resolver.md#step-4-report-resolution` so the orchestrator can gate missing required skills without guessing. |

Status semantics:

- `success`: the phase reached its objective, persisted or returned the expected artifact according to `artifact_store.mode`, and the next phase may be considered after gatekeeper validation.
- `partial`: the phase produced useful output but could not fully persist, verify, archive, or complete a required operation. Include recovery steps in `detailed_report` and route to `resolve-blockers` unless the phase-specific contract defines a safe retry.
- `blocked`: the phase could not safely proceed because it needs user input, dependency repair, artifact reconciliation, safe edit context, or another external decision. Set `next_recommended: resolve-blockers` unless a phase-specific contract routes to an earlier SDD phase for remediation. When unblocking a blocked phase, the orchestrator MUST re-delegate to the phase executor — never execute phase work inline. See Blocker Resolution Protocol in `executor-boundary.md`.
- If a phase needs user input or orchestration, return `blocked`; do not ask the user directly from the sub-agent.

Artifact entry shape:

```yaml
artifacts:
  - type: explore | proposal | spec | design | test-design | tasks | apply-progress | review-report | review-security-report | verify-report | archive-report | state | other
    mode: engram | openspec
    ref: "topic key, file path, or inline ref"
    persisted: true | false
    readable: true | false
    notes: "optional concise context"
```

Artifact rules:

- `success` requires every expected artifact for that phase to be `persisted: true` and `readable: true`.
- `partial` must name which artifact or write failed and include the inline artifact content when safe.
- Artifacts must use the naming conventions in this file and `skills/_shared/persistence-contract.md`.

Risk entry shape:

```yaml
risks:
  - severity: CRITICAL | WARNING | SUGGESTION
    blocker: true | false
    message: "specific risk"
    evidence: "artifact ref, file path, command output, or rationale"
    owner: orchestrator | user | phase-agent | reviewer
```

Risk rules:

- Use `risks: None` only when there are no material risks.
- Any `CRITICAL` risk or `blocker: true` prevents dependent phases from launching.
- Warnings may continue only when explicitly non-blocking and consistent with the phase-specific rules.

Use `mode: none` only when no supplemental skills were required or no registry/paths were available. If the orchestrator injected `## Skills to load before work`, prefer `mode: paths-injected` and list every path actually loaded according to `skill-resolver.md`.

Routing token and field naming rules live in `skills/_shared/sdd-status-contract.md`. Phase envelopes return `next_recommended`; consumers normalize it through the status contract before routing, comparing successors, or persisting state.

Phase-specific minimum details:

| Phase | Minimum required detail |
| --- | --- |
| `sdd-explore` | Findings, relevant files/artifacts consulted, recommendation, uncertainty/risk. |
| `sdd-propose` | Proposal scope, non-goals, assumptions, unresolved product questions if any. |
| `sdd-spec` | Requirements/scenarios produced, domain/spec refs, proposal traceability. |
| `sdd-design` | Architecture approach, tradeoffs, affected components, spec traceability. |
| `sdd-test-design` | Test-design ref, planned cases, mandatory/non-mandatory coverage expectations, and no-impact assessment when applicable. |
| `sdd-tasks` | Task list refs plus Review Workload Forecast, estimated changed lines, chain/exception recommendation, and test-design traceability. |
| `sdd-apply` | Apply-progress ref, completed/pending task summary, files changed, verification run or reason not run, next slice boundary if chained. |
| `sdd-review` | Final review verdict, blocking/non-blocking summary, matrix validation summary, review-report ref, and next remediation or verification route. |
| `sdd-review-security` | Final security review verdict, blocking/non-blocking summary, source-row validation summary, canonical `review-security-report.json` ref, derived Markdown ref, artifact parity/read-back metadata, and next remediation or verification route. |
| `sdd-verify` | Final verdict `PASS`, `PASS WITH WARNINGS`, or `FAIL`; evidence table; CRITICAL/WARNING/SUGGESTION issues; verify-report ref. |
| `sdd-archive` | Archive destination/ref, included artifacts, final status, recovery path if partial. |

Response ordering:

- Persist artifacts according to `artifact_store.mode` before returning the final envelope.
- The final output must be one fenced YAML block (` ```yaml ... ``` `), not a tool result and not prose. No text before the opening fence and no text after the closing fence. If the final action is a persistence tool call, the orchestrator loses the analysis.
- Do not call `mem_session_summary`; session summaries are reserved for the top-level agent.

Example (using `sdd-propose` as the reference phase — adapt `phase`, `next_recommended`, `artifacts.type`, and `detailed_report` content per phase):

```yaml
status: success
phase: propose
change: sdd-cost-tracker-calls
next_recommended: spec
executive_summary: |
  Proposal created for `sdd-cost-tracker-calls`. Defined intent, scope (5 deliverables in,
  phases table deferred), rollback plan, and capabilities contract for sdd-spec.
  Risk level is low — additive only, no breaking changes.
artifacts:
  - type: proposal
    mode: openspec
    ref: openspec/changes/sdd-cost-tracker-calls/proposal.md
    persisted: true
    readable: true
    notes: null
risks: None
skill_resolution:
  mode: paths-injected
  loaded:
    - skills/project-conventions/SKILL.md
detailed_report: |
  ## Proposal Created

  **Change**: sdd-cost-tracker-calls
  **Location**: `openspec/changes/sdd-cost-tracker-calls/proposal.md`

  ### Summary
  - **Intent**: Add per-call granularity to sdd-cost-tracker
  - **Scope**: 5 deliverables in scope; phases table and MCP tools deferred
  - **Approach**: Mirror existing patterns; fire-and-forget postCall branch in plugin
  - **Risk Level**: Low

  ### Next Step
  Ready for specs (sdd-spec). Design runs after specs pass.
```

Rules for the envelope:
- `status` MUST be one of `success`, `partial`, or `blocked`. Never `completed`, `done`, `ok`, `hold`, or `failed`.
- `phase` is the short phase token (e.g. `propose`, `spec`, `design`, `test-design`, `tasks`, `apply`, `review`, `review-security`, `verify`, `archive`).
- `change` is the kebab-case change name from the launch envelope `changeName`.
- `executive_summary` is a short human-readable paragraph (3-5 lines) summarizing what happened and why it matters. Write it as plain text or a single-line YAML scalar. Do NOT use bullet lists here.
- `artifacts` is always an array, even for a single artifact. Use the artifact entry shape above.
- `risks` is a structured array using the risk entry shape, or the literal word `None`. Never use an empty array `[]` — use `None` when there are no risks.
- `skill_resolution` follows `skills/_shared/skill-resolver.md#step-4-report-resolution`. Always include it; use `mode: none` only when no supplemental skills were required.
- `detailed_report` is a fenced YAML scalar (use `|`). Put the full phase-specific summary here — requirements tables, coverage, decisions, forecast, etc. See the phase-specific minimum detail table above.
- `next_recommended` uses snake_case phase tokens only (e.g. `spec`, `design`, `test-design`, `tasks`, `apply`, `resolve-blockers`). Never camelCase.

See `skills/_shared/skill-resolver.md` for the full `skill_resolution` shape and allowed modes.

## E. Artifact Naming Convention

This section only covers phase/artifact naming gotchas for phase envelopes and persisted artifacts. Routing-token mapping and camelCase/snake_case field normalization live in `skills/_shared/sdd-status-contract.md`; backend artifact refs live in `skills/_shared/persistence-contract.md`.

- `sdd-propose` and `sdd-spec` are phase/agent names.
- `proposal` and `spec` are singular Engram artifact keys: `sdd/{change-name}/proposal`, `sdd/{change-name}/spec`.
- `specs` is the OpenSpec/status collection name because file-based mode may write multiple domain files under `openspec/changes/{change-name}/specs/{domain}/spec.md`.
- `sdd-test-design` is the phase/agent token; `test-design` is the native/status token and artifact key: `sdd/{change-name}/test-design` or `openspec/changes/{change-name}/test-design.md`.
- Persisted camelCase state/status fields use `testDesign` when a field name cannot contain hyphens.
- Do not use `sdd-proposal`, `sdd-proposals`, `sdd-specs`, `sdd-test-design`, or `testDesign` as Engram artifact keys.

## F. Review Workload Guard

SDD must protect reviewer cognitive load, not only generate tasks.

- **Authority boundary:** this section is the shared Review Workload / Delivery Guard contract. The orchestrator enforces it, `sdd-tasks` produces the forecast, and `sdd-apply` consumes the resolved decision. Phase skills may add phase-local validation, but they MUST NOT redefine the common delivery policy here.
- **Budget semantics:** `review_budget_lines` means changed lines (`additions + deletions`), not net delta or file count. Default to **400 changed lines** unless the orchestrator resolved a different budget. Treat work that likely exceeds about one focused review session as review-budget risk even when the numeric estimate is under budget.
- **Deferred decision:** the orchestrator may keep `delivery_strategy`, `review_budget_lines`, and `chain_strategy` as `null` until `sdd-tasks` produces a Review Workload Forecast or the user explicitly asks to decide delivery earlier. Pass `delivery_strategy: null` to `sdd-tasks` while deferred; pass the resolved decision to `sdd-apply`.
- **Forecast producer:** `sdd-tasks` MUST estimate whether implementation may exceed the resolved budget and include exact plain-text guard lines: `Decision needed before apply: Yes|No`, `Chained PRs recommended: Yes|No`, `Chain strategy: stacked-to-main|feature-branch-chain|pending`, `Size exception: approved|pending|none`, `Review budget lines: <number>`, `Review budget risk: Low|Medium|High`, and legacy `400-line budget risk: Low|Medium|High`.
- **Forecast validity:** the orchestrator MUST reject a missing, stale, ambiguous, or incomplete forecast before `sdd-apply`. Re-run `sdd-tasks` once with corrective feedback when the forecast is malformed. Recompute the forecast when scope, tasks, delivery strategy, chain plan, review budget, or artifact dependencies changed after the forecast.
- **Forecast handling:** low-risk work under budget may proceed as one reviewable unit; medium/near-budget work may proceed only with work-unit implementation and changed-line monitoring; high-risk, over-budget, `Chained PRs recommended: Yes`, or `Decision needed before apply: Yes` requires a resolved delivery path before implementation.
- **Delivery strategies:** `ask-on-risk` asks the user when the forecast requires a delivery decision; `auto-chain` proceeds with chained/stacked slices but still requires an explicit `chain_strategy`; `single-pr` requires approved `size:exception` before over-budget apply; `exception-ok` proceeds only when size-exception evidence is recorded.
- **Chain strategies:** valid chain strategies are `stacked-to-main` and `feature-branch-chain`. `size:exception` is a delivery/approval decision, not a chain strategy. Keep `Chain strategy` limited to `stacked-to-main`, `feature-branch-chain`, or `pending`.
- **Chain plan requirements:** when chaining is selected, persist a chain plan before `sdd-apply` with ordered slices, current slice boundary, dependency diagram, per-slice review estimate, verification plan, rollback scope, and out-of-scope/follow-up work. Each slice must have a clear start, clear finish, autonomous scope, verification, and reasonable rollback.
- **Feature Branch Chain:** PR #1 targets the feature/tracker branch; later child PRs target the immediate previous child branch; only the tracker is intended to merge to `main`. If a child diff includes previous slices, retarget/rebase before review.
- **Size exception evidence:** record approver/instruction, rationale for not splitting or accepting a large PR, accepted risk, verification plan, rollback plan, and follow-up work. A vague "ok" or phase approval is not `size:exception` approval.
- **Apply boundary:** `sdd-apply` MUST NOT start oversized work unless the delivery strategy resolves to a chained/stacked slice or explicitly accepted `size:exception`. If actual changed lines exceed the resolved budget or invalidate the forecast, stop before PR creation, archive, or another implementation slice and re-apply this guard.
- **PR authorization boundary:** delivery strategy shapes implementation and review; it does not authorize creating, pushing, or opening PRs unless the user explicitly requested PR work or the active command/phase contract includes PR creation.
- **Required supplemental skills:** when chaining is selected or forecasted work exceeds 400 changed lines, orchestrator launches for `sdd-tasks` and `sdd-apply` must include resolved `chained-pr` and `work-unit-commits` skill paths when available. Missing required skills are a gate failure for apply/PR-shaped work.

This guard exists to reduce reviewer burnout and keep implementation delivery safe. Do not treat it as optional process noise.
