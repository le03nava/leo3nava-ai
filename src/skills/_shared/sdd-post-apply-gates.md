# SDD Post-Apply Gates Contract

Shared execution contract for phases that run after implementation: `sdd-review`, `sdd-review-security`, `sdd-verify`, and `sdd-archive`.

This contract centralizes common dependency gates, review-evidence handling, archive readiness, routing, matrix ownership boundaries, safe-evidence rules, unavailable-tooling language, and persistence expectations. Phase skills still own their phase-specific catalogs, report schemas, validation logic, final verdict semantics, and archive filesystem/spec-sync mechanics.

## Phase Scope

| Phase | Owns | Consumes | Must not own |
| --- | --- | --- | --- |
| `sdd-review` | General applied-change review, canonical `review-report.json`, and derived 96-control `review-report.md` compatibility matrix. | Status, design/test-design/tasks/apply evidence, changed-file context. | Security-review row verdicts, exhaustive compact security matrix, exhaustive Source ID matrix, final runtime verification, Excel/Python/workbook generation. |
| `sdd-review-security` | Security review, canonical `review-security-report.json`, derived `review-security-report.md` / `sdd/{change-name}/review-security` compatibility view, and exhaustive `sourceRowValidation.rows` exact-once Source ID validation. | Runs in parallel with `sdd-review`; shared base inputs only: `design.md#secure-development-design`, test-design, tasks/apply evidence, changed-file context. Does not consume general review evidence. | The general 96-control review matrix, final runtime verification, archive readiness. |
| `sdd-verify` | Final implementation/spec/test-design/security evidence verification and `verify-report.md`. | Non-blocking general review evidence: canonical `review-report.json` when present plus derived `review-report.md` / `sdd/{change-name}/review`, non-blocking security-review evidence with canonical `review-security-report.json` when present plus derived Markdown compatibility view, implementation evidence, runtime/static command evidence. | Reproducing or re-scoring either review matrix, owning exhaustive Source ID rows, fixing implementation. |
| `sdd-archive` | Final source-of-truth sync, archive move/report, and audit-trail preservation. | Passing `verify-report`, non-blocking review reports, completed tasks/apply evidence, mandatory security evidence, status/action context. | Re-scoring reviews, re-running verification, fixing implementation, or owning review matrices. |

Operational evidence is design-driven. Post-apply phases consume operational considerations, gaps, warnings, unavailable-tooling notes, evidence refs, exact `Pendiente de confirmar:` markers, exact `No aplica.` markers, and archive/manual-document handoff boundaries only when those items are present or planned in design/test-design/tasks/apply/review/archive artifacts.

## Common Required Context

Every post-apply phase MUST resolve the structured SDD status and confirm the selected change before judging evidence.

Common inputs:

- selected `changeName`
- selected `artifact_store.mode`
- structured status compatible with `skills/_shared/sdd-status-contract.md`
- `design.md#secure-development-design` for active new changes
- `test-design.md`
- tasks and apply-progress or completed task evidence
- operational considerations, unresolved gaps, warning carry-forward, and unavailable-tooling notes when the change includes applicable operational evidence
- changed-file context when the phase judges changed-file evidence directly
- safe `actionContext` / workspace context
- selected backend references following `skills/_shared/persistence-contract.md`

Additional phase-specific inputs:

| Phase | Additional required input |
| --- | --- |
| `sdd-review` | Proposal/specs when available, 96-control review catalog. |
| `sdd-review-security` | Security guideline catalog, security contract. |
| `sdd-verify` | Non-blocking general review evidence: canonical `review-report.json` when present and derived `review-report.md` / `sdd/{change-name}/review`, non-blocking security-review evidence with canonical `review-security-report.json` when present and derived `review-security-report.md` / `sdd/{change-name}/review-security`, testing capabilities/config. |
| `sdd-archive` | Passing `verify-report`, complete task state, state artifact, archive destination context, and any explicit partial/stale-checkbox reconciliation approval. |

## Common Blocking Rules

Return `status: blocked` with `next_recommended: resolve-blockers` when the phase cannot safely judge evidence because of infrastructure, context, or artifact integrity:

- required artifacts are missing, unreadable, ambiguous, or materially conflicting
- changed-file context is absent or ambiguous for a phase that judges changed-file evidence directly
- workspace/action context is unsafe, outside allowed roots, or planning-only for a phase that needs repo evidence
- required report evidence is missing, unreadable, malformed, or ambiguous
- persistence or read-back verification fails after no useful durable artifact can be trusted
- evidence contains secrets, credentials, PAN, PII, tokens, private keys, connection strings, or unnecessary confidential values
- operational evidence contains restricted production identifiers, raw logs/payloads, generated file bytes, final-document-only values outside the final manual document, or invented operational details
- applicable operational evidence is missing after design, test-design, tasks, apply evidence, or reviews explicitly made it applicable, and the gap is not carried as safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.`
- unavailable tooling is treated as passing evidence instead of being reported as unavailable

When a useful report was produced but persistence/read-back failed, return `status: partial` with `next_recommended: resolve-blockers` and include the safe report content or recovery path in `detailed_report`.

## Common Routing Rules

Use these routes consistently:

| Condition | Route |
| --- | --- |
| Remediation requires code, prompt, task, contract, or apply-evidence work | `next_recommended: apply` |
| Remediation requires missing context, artifact repair, backend reconciliation, unsafe evidence cleanup, schema/catalog repair, or configuration repair | `next_recommended: resolve-blockers` |
| Current phase produced non-blocking review evidence | next DAG phase only: `verify` or `archive` |
| Current phase found CRITICAL/blocking issues | never advance to a downstream phase |
| Archive readiness fails because general review is missing or unreadable | `next_recommended: review` |
| Archive readiness fails because security review is missing or unreadable | `next_recommended: review-security` |
| Archive readiness fails because verification is missing or non-passing | `next_recommended: verify` |
| Archive readiness fails because implementation tasks/evidence are incomplete | `next_recommended: apply` |

Phase envelopes use snake_case `next_recommended`. Persisted status/state uses camelCase `nextRecommended` only where required by `skills/_shared/sdd-status-contract.md`.

## Review Evidence Consumption

- `sdd-review` and `sdd-review-security` run in parallel after `sdd-apply`. They share the same base inputs (`design.md#secure-development-design`, test-design, tasks/apply evidence, changed-file context) and are independent of each other's output. The orchestrator MUST wait until both appear in `completedPhases` before launching `sdd-verify`.
- `sdd-review-security` does NOT consume `review-report.json` and MUST NOT require it as a prerequisite. It owns its security verdict independently.
- `sdd-review-security` MUST produce canonical `review-security-report.json` first and derived Markdown second. The canonical JSON owns security-review verdict/status, routing, source-row counts, exact-once coverage, blocker/warning summaries, unsafe evidence rejections, warning carry-forward, exceptions, evidence refs, and artifact parity/read-back metadata.
- `sdd-verify` MUST consume exactly one non-blocking general review report identity from the selected backend, including canonical JSON when present plus the derived Markdown compatibility view, and exactly one non-blocking security review report identity from the selected backend, including canonical `review-security-report.json` when present plus derived Markdown compatibility view.
- `sdd-archive` MUST consume exactly one non-blocking general review report identity from the selected backend, including canonical JSON when present plus the derived Markdown compatibility view, and exactly one non-blocking security review report identity from the selected backend, including canonical `review-security-report.json` when present plus derived Markdown compatibility view, plus a passing verify report.
- Missing, ambiguous, unreadable, or blocking review evidence blocks dependent phases.
- Downstream phases cite report verdicts, blocker summaries, warning summaries, evidence summaries, catalog identity/counts, validation/parity status, exceptions, and report links. They do not copy full matrices.
- When canonical JSON and derived Markdown disagree for either review report, canonical JSON wins and downstream phases MUST route stale/parity-failed Markdown to `resolve-blockers` rather than re-score or repair the matrix themselves.
- These gates do not introduce Excel, Python, script, spreadsheet, or workbook generation. Presentation metadata in JSON may be consumed as evidence, but generating workbooks remains outside this contract.

Operational evidence consumption:

| Phase | Must consume | Must preserve or hand off |
| --- | --- | --- |
| `sdd-review` | Present/planned operational considerations from design/test-design/tasks/apply evidence, exact markers, changed-file refs, unavailable-tooling notes, and unresolved gaps. | Operational findings outside the 96-control matrix plus refs/gaps for `sdd-review-security` when evidence exists. |
| `sdd-review-security` | Secure design rules, safe-evidence policy, changed files, and apply evidence when operational evidence exists. | Leakage verdicts for restricted operational identifiers, secrets, logs/payloads, generated bytes, final-doc-only values, safe placeholders, and future exception evidence. |
| `sdd-verify` | Non-blocking review and security-review operational evidence, warning summaries, unavailable-tooling notes, and implementation evidence when present. | Verification that applicable operational evidence/gaps/warnings remain preserved, without requiring absent categories. |
| `sdd-archive` | Passing verify operational evidence, non-blocking review/security-review summaries, completed task state, warnings, exceptions, and unavailable-tooling notes when present. | Operational status, evidence refs, unresolved gaps, warning carry-forward, final-document handoff boundaries, and confirmation that `sdd-operational-doc` remains manual post-archive when relevant. |

## Archive Readiness

`sdd-archive` MUST prove final readiness before syncing specs, moving folders, or recording cycle completion.

Archive may proceed only when all required evidence is readable in the selected backend:

- proposal, specs, design with `## Secure Development Design`, test-design, tasks, apply-progress or completed task evidence, state, non-blocking general review report, non-blocking security review report with canonical JSON when present, and passing `verify-report`
- all implementation tasks are complete in the persisted task artifact, unless explicit stale-checkbox reconciliation is approved and backed by apply-progress plus verify-report proof
- mandatory applicable security evidence is complete or covered by complete approved exceptions
- applicable operational evidence, unresolved gaps, warning carry-forward, unavailable-tooling notes, and manual-document handoff boundaries that exist are preserved without requiring absent readiness categories or `sdd-operational-doc` execution
- review/security-review reports contain no blocking findings, and verify contains no CRITICAL issues or failing verdict
- action context is repo-safe for archive operations; workspace-planning mode cannot move folders or edit linked repos

Allowed archive exceptions are narrow:

- **Intentional partial archive**: allowed only for missing non-critical planning artifacts when explicitly approved and recorded in the archive report.
- **Stale-checkbox reconciliation**: allowed only when the orchestrator provides explicit approval and apply-progress plus verify-report prove every stale unchecked task is complete.
- **Approved security exception**: allowed only when every required exception field from `skills/_shared/sdd-security-contract.md` is complete.

Archive MUST block on:

| Condition | Route |
| --- | --- |
| Missing/unreadable/ambiguous `review-report` | `review` |
| Blocking general review finding or verdict | `apply` |
| Missing/unreadable/ambiguous `review-security-report`, missing canonical `review-security-report.json` when expected, or stale/parity-failed derived security Markdown | `review-security` or `resolve-blockers` according to artifact ownership |
| Blocking security review finding, unresolved source-row blocker, unsafe evidence, unsupported `N/A`, missing mandatory source-row evidence, or malformed source-row schema | `apply` or `resolve-blockers` according to remediation ownership |
| Missing `verify-report` | `verify` |
| Verify verdict `FAIL`, CRITICAL issue, or non-passing report | `apply` |
| Missing applicable operational refs, unresolved gaps not carried forward, missing unavailable-tooling notes, or final-document handoff boundary absent when an artifact made them applicable | `apply` or `verify` according to remediation ownership |
| Operational evidence contains restricted operational data, invented values, raw logs/payloads, generated bytes, or final-document-only values outside the final manual document | `apply` or `resolve-blockers` according to remediation ownership |
| Unchecked persisted implementation tasks without approved stale-checkbox reconciliation | `apply` |
| Missing mandatory embedded secure design, test-design, or mandatory security evidence without complete approved exception | `resolve-blockers` |
| Unsafe action context, destination conflict, backend conflict, or archive operation outside allowed roots | `resolve-blockers` |

## Matrix Ownership Boundary

- `sdd-review` is the only active owner of the general 96-control review matrix.
- `sdd-review-security` is the only active owner of exhaustive 155 Source ID validation in canonical `sourceRowValidation.rows`.
- `sdd-verify` validates that the review reports are present, non-blocking, internally credible, and consumed as prerequisites. It MUST NOT reproduce or re-score either matrix.
- `sdd-archive` preserves links, verdicts, warning summaries, exceptions, and audit references only.

## Source-Row Consumption Boundary

For corporate source-row changes:

- `sdd-review-security` expands and validates every expected Source ID exactly once, then reports coverage metadata and focused findings by default.
- `sdd-verify` consumes the security-review source-row verdict, warning summary, catalog snapshot identity/path, expected and validated Source ID counts, exact-once coverage status, safe-evidence status, `N/A` justification status, exceptions, evidence references, blocker absence, parity metadata, and report links.
- `sdd-archive` preserves source-row coverage summaries, catalog identity/path, exact-once coverage status, expected and validated counts, warnings, exceptions, safe evidence references, `N/A` evidence/justification status, review-security verdict, verify consumption, parity metadata, and report links.
- `sdd-verify` and `sdd-archive` block unresolved source-row blockers but do not duplicate the full Source ID matrix.
- Source-row routing follows `skills/_shared/sdd-security-contract.md#routing-semantics`.

## Safe Evidence and Unavailable Tooling

Evidence must be review-safe:

- cite paths, section anchors, changed-file refs, command summaries, sanitized summaries, or redacted placeholders
- never copy raw secrets, credentials, PAN, PII, connection strings, private keys, tokens, or unnecessary confidential data
- never copy production hostnames, IP addresses, ports, SID/service names, sensitive payloads, full ID lists, generated file bytes, or final-document-only values into ordinary SDD evidence
- preserve applicable operational gaps with exact `Pendiente de confirmar:` and inapplicable operational content with exact `No aplica.`
- report unavailable runtime/build/lint/type/format/coverage tooling explicitly
- never invent commands or mark unavailable tooling as passing evidence

## Persistence and Return Envelope

Every post-apply phase MUST:

1. Load supplemental skills according to `skills/_shared/skill-resolver.md` and `skills/_shared/sdd-phase-common.md` Section A executor minimum.
2. Resolve/read dependencies through `skills/_shared/persistence-contract.md`.
3. Persist only its owned artifact.
4. Read back its artifact before reporting `status: success` in persistent modes.
5. Return the Section D phase envelope from `skills/_shared/sdd-phase-common.md`.

Do not call `mem_session_summary` from a phase executor. The final output must be text containing the phase envelope, not a final tool call.
