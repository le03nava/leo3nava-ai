---
name: sdd-verify
description: "Trigger: SDD verification phase, verify change. Execute tests and prove implementation matches specs, design, and tasks."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-verify` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Activation Contract

Run when the orchestrator launches verification for an SDD change. You are the quality gate: prove completion with source inspection plus real execution evidence.

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.status.dependencies` — `review: completed` and `review-security: completed` must both be present
- `launch.artifacts.paths` or `launch.artifacts.refs` — refs for spec, design, test-design, tasks, apply-progress, review-report, review-security-report; do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute workspace path to run tests and inspect code
- `launch.actionContext.allowedEditRoots` — scope for any test execution or file inspection
- `launch.skill_paths` — supplemental skills to load before work

The `## Phase Context` section contains:
- Strict TDD instructions when active (`STRICT TDD MODE IS ACTIVE. Test runner: {cmd}.`)

If `launch.artifact_store.mode` is absent, or `launch.status.dependencies` does not show both `review` and `review-security` as completed, return `blocked` with `next_recommended: resolve-blockers`.

## Hard Rules

- Follow `skills/_shared/sdd-post-apply-gates.md` for common post-apply dependency gates, review-evidence consumption, safe evidence, unavailable tooling, matrix ownership boundaries, routing defaults, and persistence/read-back expectations.
- Read all available status `contextFiles` before judging implementation. Full spec-driven verification reads proposal, specs, design with mandatory `## Secure Development Design`, test design, tasks, general review, and security review; partial artifact sets degrade as described below.
- Execute relevant tests; static analysis alone is never verification.
- A spec scenario is compliant only when a covering test passed at runtime.
- When repository testing capabilities explicitly report no runtime/build/lint/type/format/coverage runner, report those tools as unavailable evidence per `skills/_shared/sdd-post-apply-gates.md`.
- Compare specs first, design second, task completion third.
- Compare embedded secure-design controls and mandatory evidence before archive readiness can be claimed.
- Compare `test-design.md` planned cases against apply/verification evidence. Uncovered mandatory cases fail verification; uncovered non-mandatory cases are warnings only.
- Verify operational evidence from actual artifacts: applicable operational considerations, gaps, warnings, exact `Pendiente de confirmar:` / `No aplica.` markers, and unavailable-tooling notes must be preserved without requiring absent readiness categories.
- Consume review reports and source-row summaries according to `skills/_shared/sdd-post-apply-gates.md`; verification MUST NOT own, reproduce, or re-score either review matrix.
- Do not fix issues; report them for the orchestrator/user.
- Persist `verify-report` according to `skills/_shared/persistence-contract.md` and the phase artifact contract below.
- If Strict TDD is active, load `strict-tdd-verify.md` from this skill directory; if inactive, never load it.
- Apply any `rules.verify` from `openspec/config.yaml`.
- Return the Section D envelope from `skills/_shared/sdd-phase-common.md`.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`.
Common post-apply gates and review-evidence consumption rules are defined in `skills/_shared/sdd-post-apply-gates.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Structured status plus available proposal, specs, design with mandatory `## Secure Development Design`, mandatory `test-cases.json` and `test-design`, tasks, apply-progress/task checkbox evidence, non-blocking general review evidence from the selected backend (canonical `review-report.json` when present plus derived `review-report.md` / `sdd/{change-name}/review` compatibility view), and non-blocking security-review evidence from the selected backend (canonical `review-security-report.json` when present plus derived `review-security-report.md` / `sdd/{change-name}/review-security` compatibility view). Standalone `security-design.md` is legacy/read-only compatibility data only. |
| Produced artifact | `sdd/{change-name}/verify-report` or `openspec/changes/{change-name}/verify-report.md`, plus updated `sdd/{change-name}/test-cases` or `openspec/changes/{change-name}/test-cases.json` with final per-case statuses as the canonical lifecycle record. |
| Mutates | Canonical `test-cases.json` status fields during verification, plus the produced verification report artifact. |
| Test-design consumption | Read `test-cases.json` as canonical source of truth. For each case: assign `status: verified` when evidence passes, `status: warning` for uncovered cases with `mandatory: false`, `status: skipped` for justified deviations. Persist updated `test-cases.json` before returning success. `verify-report` MUST reference the JSON path/topic key and cite case statuses from it; it MUST NOT reproduce the full cases table. `test-design.md` remains a derived human-readable aid only. |
| Security consumption | Consume the `## Security Review Summary` section from the derived `review-security-report.md` (or canonical `review-security-report.json` when Markdown is unavailable): lift the by-domain table (passing/total + blockers per control domain) and the overall line. Do not reproduce the full 155-row matrix or build a row-by-row evidence matrix. Compare `design.md#secure-development-design` narrative rules, changed-surface rationale, residual risks, static/manual validation notes, canonical `review-security-report.json` row evidence, source-row verdict/warnings/evidence summary, artifact parity/read-back metadata, and mandatory evidence before archive readiness; complete approved exceptions are the only valid substitute for missing mandatory evidence. Derived Markdown is compatibility only. |
| Review consumption | Follow `skills/_shared/sdd-post-apply-gates.md`: resolve exactly one non-blocking general review identity and one non-blocking security review identity, cite summaries only, and never duplicate either matrix. When canonical general-review or security-review JSON is present, it is authoritative over derived Markdown for verdict, counts, routing, matrix facts, catalog identity, validation state, and artifact parity. Lift a `## General Review Summary` section from the consumed `review-report.md` (or from canonical `review-report.json` when Markdown is unavailable): include verdict, blocking count, non-blocking count, catalog identity, top blocking findings when present, and next route. Do not reproduce the full 96-control matrix. |
| Operational evidence verification | Confirm operational evidence from design/test-design/tasks/apply/review/security-review when present, including safe evidence, exact `Pendiente de confirmar:` / `No aplica.` markers, unresolved warnings, gaps, and unavailable-tooling notes. Never invent missing operational values or require absent categories. |
| Source-row prerequisite | Follow `skills/_shared/sdd-post-apply-gates.md` source-row consumption boundary and `skills/_shared/sdd-security-contract.md` routing semantics. |
| Runtime/static evidence | Execute configured commands when available; otherwise report unavailable tooling per `skills/_shared/sdd-post-apply-gates.md`. |
| Success routing | `next_recommended: archive` only for `PASS` or eligible `PASS WITH WARNINGS`. |
| Block/fail routing | `next_recommended: apply` for failed verification or unchecked tasks; `next_recommended: resolve-blockers` for unsafe workspace, missing selected change/artifacts, unresolved configuration, or persistence failure. |

## Decision Gates

| Condition | Action |
|---|---|
| Orchestrator says `STRICT TDD MODE IS ACTIVE` | Treat as authoritative. |
| Cached/config `strict_tdd: true` and runner exists | Strict TDD verify; load module. |
| Strict TDD false and runner exists | Standard verify; skip TDD protocol checks. |
| No runner and specs exist | Do not claim full spec compliance; missing runtime evidence is CRITICAL unless config explicitly allows manual verification. |
| No runner and only tasks exist | Verify task completion only; verdict may be `PASS WITH WARNINGS`. |
| Strict TDD active but `apply-progress` or TDD evidence is missing | FAIL; report CRITICAL missing TDD evidence. |
| Common post-apply dependency, review-evidence, context, safe-evidence, or persistence gate fails | Follow `skills/_shared/sdd-post-apply-gates.md`; do not claim archive readiness. |
| `actionContext.mode: workspace-planning` | Return `blocked` with `next_recommended: resolve-blockers`; full workspace implementation verification is not supported in this slice. |
| Changed files cannot be identified | Mark static correctness/design coherence as PARTIAL, record the skipped evidence, and return `next_recommended: resolve-blockers` unless runtime evidence still proves the required behavior. |
| Only tasks artifact exists | Verify task completion only; skip spec/design correctness and record skipped checks. |
| Tasks + specs exist | Verify completeness and correctness; skip design coherence and record skipped checks. |
| Proposal/specs/design/test-design/tasks exist | Verify all dimensions, including planned case coverage. |
| `design.md#secure-development-design` is missing for a new active change | CRITICAL blocker; return `next_recommended: resolve-blockers`. |
| Standalone `security-design.md` is missing for a new active change | Continue; do not require it. It is legacy/read-only compatibility data only. |
| Review report or security review report is missing, unreadable, ambiguous, malformed, blocking, missing canonical JSON when expected, or has stale/parity-failed derived Markdown | Follow `skills/_shared/sdd-post-apply-gates.md` review-evidence consumption and routing rules. |
| Security review source-row evidence has unresolved source blockers: missing/duplicate/unknown Source IDs, malformed schema, unsafe evidence, unsupported `N/A`, or missing mandatory source-row evidence | CRITICAL blocker. Route to `resolve-blockers` for schema/catalog/artifact/unsafe-evidence/unsupported-`N/A` causes and to `apply` when remediation is implementation, prompt, contract, or apply-evidence work. |
| Security review is non-blocking with source-row warnings only | Preserve the warning summary in `verify-report`; verification MAY proceed only if all mandatory evidence is complete and no CRITICAL issue exists. |
| Applicable operational evidence lacks safe evidence, exact `Pendiente de confirmar:`, exact `No aplica.`, or a carried gap after artifacts made it applicable | CRITICAL blocker; route to `apply` or `resolve-blockers` according to remediation ownership. |
| Unavailable operational tooling notes are dropped or treated as passing checks | CRITICAL blocker; preserve the unavailable-tooling state instead of claiming execution evidence. |
| Verification report duplicates the full 96-control review matrix, security review matrix, or corporate source-row matrix | Fix before persistence; matrix ownership is defined by `skills/_shared/sdd-post-apply-gates.md`. |
| Mandatory embedded secure-design control has no implementation, verification evidence, or complete approved exception | CRITICAL `SECURITY_EVIDENCE_MISSING` and verdict `FAIL`. |
| Mandatory security exception is incomplete | CRITICAL `SECURITY_EXCEPTION_INCOMPLETE` and verdict `FAIL`. |
| Task incomplete | CRITICAL for core task, WARNING for cleanup task. |
| Test command exits non-zero | CRITICAL. |
| Spec scenario has no passing covering test | CRITICAL `UNTESTED` or `FAILING`. |
| Mandatory test-design case has no matching implementation, execution, or justified skip evidence | CRITICAL `UNTESTED` and verdict `FAIL`. |
| Runtime test, linter, type-checker, formatter, or coverage command is unavailable | Report unavailable tooling according to `skills/_shared/sdd-post-apply-gates.md`. |
| Non-mandatory test-design case has no matching evidence | WARNING; do not fail solely because of this uncovered non-mandatory case. |
| After assigning statuses, `test-cases.json` persistence fails | Return `partial` with `next_recommended: resolve-blockers` and include inline case-status summary in `detailed_report`. |
| Design deviation exists | WARNING unless it breaks a spec. |
| Verification failure discovered | Report only; do not patch implementation. |
| Verify-report persistence fails | Return `partial` with `next_recommended: resolve-blockers` and the report inline in `detailed_report`. |

## Execution Steps

1. Load relevant supplemental skills according to `skills/_shared/skill-resolver.md` and the shared SDD Section A executor minimum.
2. Apply `skills/_shared/sdd-post-apply-gates.md`, then retrieve artifacts via shared Section B for the active persistence mode or read the concrete `contextFiles` from structured status.
3. Resolve testing/TDD mode from cached capabilities, config, or project files. Prefer `sdd/{project}/testing-capabilities`, then `openspec/config.yaml`, then project files.
4. Count completed and incomplete tasks. Any unchecked implementation task is CRITICAL and blocks archive readiness.
5. Resolve the general review artifact identity and security review artifact identity; prefer canonical JSON for each when present and treat derived Markdown as compatibility only. Cite each verdict, blocking summary, evidence summary, artifact parity/read-back metadata, and next recommendation. Lift a `## General Review Summary` from the consumed `review-report.md` (or canonical `review-report.json` when Markdown is unavailable): verdict, blocking count, non-blocking count, catalog identity, top blocking findings when present, and next route. If either report is missing, ambiguous, unreadable, blocking, or has stale/parity-failed Markdown, classify it according to Decision Gates before continuing.
5a. When corporate source-row validation applies, consume the security review source-row summary: catalog snapshot identity/path, expected and validated row counts, exact-once coverage statement, safe-evidence status, `N/A` justification status, warnings, exceptions, evidence references, parity metadata, blocker list, and source-row next route. Block unresolved source-row blockers before archive readiness.
6. If specs exist, map each spec requirement/scenario to implementation evidence and tests.
7. Validate mandatory embedded secure-design evidence: every new change requires `design.md#secure-development-design`, while no-impact appears as narrative changed-surface rationale in design and source-row validation coverage, `N/A`, and `not-applicable` decisions belong to canonical `review-security-report.json` with derived Markdown compatibility.
8. Map each mandatory control to implementation references, test-design cases, review-security row evidence, verification evidence, archive evidence fields, or complete approved exceptions; classify gaps as CRITICAL.
9. If test design exists, map each planned case to implementation, execution, apply-progress, validation metadata, security evidence, review evidence summary, unavailable-runtime-tooling report, or justified skip evidence and classify mandatory gaps as CRITICAL and non-mandatory gaps as WARNING.
9a. Map applicable operational considerations and planned operational checks to safe evidence, exact `Pendiente de confirmar:`, exact `No aplica.`, carried gaps, or a complete blocker. Confirm review/security-review operational warnings and unavailable-tooling notes are carried forward without inventing values or requiring absent categories.
10. If design exists, check design decisions against changed code. If design is missing, skip design coherence and record why.
11. Run test, build/type-check, and coverage commands when available. For full spec verification, require runtime evidence when tooling is available: source inspection alone does not prove spec scenario compliance unless the repository configuration explicitly limits the change to static/manual evidence. Report unavailable runtime/build/lint/type/format/coverage tooling explicitly.
12. Lift the `## Security Review Summary` section from the consumed `review-security-report.md` (or derive it from canonical `review-security-report.json` when Markdown is unavailable): grouped by `controlDomain`, passing/total + blockers per domain. Do not build a row-by-row security matrix. Build the behavioral compliance matrix from actual test results and evidence when specs/scenarios exist; cite review and security-review summary evidence separately, preserve catalog identity/count/mapping/warning/exception/report-link evidence, and do not duplicate either review matrix or the full source-row matrix.
13. Validate the report before persistence: completeness table present, general review and security review evidence citations present, source-row verdict/warning consumption present when applicable, every spec scenario has a status, every security control has evidence/exception status when required, every test-design case has a coverage status, validation metadata is covered, unavailable tooling is explicit, runtime evidence includes command and result when available, skipped dimensions are listed, Strict TDD sections are present when active, and any CRITICAL issue forces verdict `FAIL`.
14. Persist updated `test-cases.json` first using the canonical key/path (`sdd/{change-name}/test-cases` / `openspec/changes/{change-name}/test-cases.json`) with final per-case statuses, then persist and return the verification report, including skipped dimensions for missing artifacts.

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put `## Verification Report` in `detailed_report` with change, mode, completeness table, general review summary (lifted from `## General Review Summary` in the consumed review-report: verdict, blocking/non-blocking counts, catalog identity, top blocking findings when present, and next route), security review evidence citations, security review summary by control domain (lifted from `## Security Review Summary` in the consumed review-security-report), operational evidence/gaps/warnings table when applicable, source-row verdict/warnings/evidence-summary consumption when applicable, catalog snapshot identity/path, expected and validated Source ID counts, exact-once coverage status, exceptions, evidence refs, parity metadata, report links, build/tests/coverage evidence, unavailable-tooling report, embedded secure-design validation/no-impact row evidence, spec compliance matrix, test-design coverage matrix, correctness table, design coherence table, skipped/degraded dimensions, issues grouped as CRITICAL/WARNING/SUGGESTION, and final verdict `PASS`, `PASS WITH WARNINGS`, or `FAIL`. Cite canonical `review-report.json` when present plus derived `review-report.md` / `sdd/{change-name}/review` compatibility refs, and canonical `review-security-report.json` when present plus derived `review-security-report.md` / `sdd/{change-name}/review-security` compatibility refs; summarize verdict/blocking state only and do not reproduce either review matrix or the full corporate source-row matrix.

## Routing Contract

- Verdict `PASS` -> return `next_recommended: archive`.
- Verdict `PASS WITH WARNINGS` -> return `next_recommended: archive` only when there are no CRITICAL issues, all implementation tasks are complete, and warnings are explicitly non-blocking.
- Verdict `FAIL`, any CRITICAL issue, failing tests, missing required runtime evidence, missing Strict TDD evidence, or unchecked implementation tasks -> return `next_recommended: apply`.
- Status `blocked` due unsafe workspace, missing selected change, missing required artifacts, missing/ambiguous/unreadable non-blocking review evidence, or unresolved configuration -> return `next_recommended: resolve-blockers` and do not claim archive readiness.
- Status `partial` because verify-report persistence failed -> return `next_recommended: resolve-blockers` unless the inline report clearly proves a retry can safely persist the same evidence.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Graceful Artifact Handling

- **Tasks only**: verify objective task completion only. Do not claim spec correctness or design coherence. If all tasks are checked and no runtime evidence is available, verdict may be `PASS WITH WARNINGS` for task completion only.
- **Tasks + specs**: verify task completeness and requirement/scenario correctness. Runtime test evidence is still required for full spec scenario compliance; missing covering tests are CRITICAL for required scenarios unless project config explicitly allows manual verification.
- **Full artifacts**: verify completeness, correctness, review evidence consumption, security evidence, test-design planned case coverage, and coherence.
- **Unchecked tasks**: always remain CRITICAL, even when other artifacts are missing or warnings-only.

## Rules

- ALWAYS update and persist `test-cases.json` with final case statuses before returning success. The JSON is the canonical lifecycle record.
- NEVER reproduce the full cases table in verify-report — reference `test-cases.json` instead.

## References

- [assets/verify-report.md](assets/verify-report.md) — full report template, compliance statuses, and command evidence fields.
- [strict-tdd-verify.md](strict-tdd-verify.md) — load only when Strict TDD is active.
- `../_shared/skill-resolver.md` — supplemental skill loading and `skill_resolution` protocol.
- `../_shared/sdd-phase-common.md` — phase retrieval, persistence, and return envelope.
- `../_shared/sdd-post-apply-gates.md` — common post-apply gates, review evidence consumption, source-row consumption, safe evidence, and matrix ownership boundaries.
