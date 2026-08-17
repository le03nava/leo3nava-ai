# Tasks: Corporate Source Row Security Validation

## Review Workload Forecast

| Field | Value |
| --- | --- |
| estimated_changed_lines | 700-1,000 |
| review_budget_lines | 400 |
| review_budget_risk | High |
| chained_prs_recommended | Yes |
| decision_needed_before_apply | No |
| rationale | Shared contracts, five phase skills, five prompts, source specs, and static/manual evidence will likely exceed one focused review. |
| work_unit_boundaries | WU1 shared contracts; WU2 design/test-design; WU3 review-security; WU4 verify/archive/spec sync. |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Chain Plan

Stacked-to-main PRs recommended: PR1 WU1 -> PR2 WU2 -> PR3 WU3 -> PR4 WU4. Each PR keeps matching static/manual evidence with its changed contracts and stays independently reviewable.

## Work Unit 1: Shared Catalog and Security Contract

- [x] 1.1 Update `skills/_shared/security-guideline-catalog.md` with the 155 expanded Source ID inventory rules, PCI alignment preservation, range expansion, and compact mappings. Covers TD-001, TD-002, TD-003.
- [x] 1.2 Update `skills/_shared/sdd-security-contract.md` with source-row schema fields, allowed values, safe-evidence/N/A policy, traceability, routing, and OpenSpec/Engram/hybrid/none semantics. Covers TD-004-TD-008, TD-015.
- [x] 1.3 Add static/manual apply evidence showing no replacement of compact `SEC-*`, no duplicated 96-control matrix, and no legacy `validate_security_design.ps1` dependency. Covers TD-003, TD-011, TD-016.

### Apply Evidence — Work Unit 1

| Check | Evidence | Result |
| --- | --- | --- |
| TD-001 / TD-002 — expanded inventory | `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` declares 155 concrete Source IDs across 15 corporate sections, requires range expansion before validation, and preserves section PCI alignment or `N/A`. | Pass — static/manual artifact evidence. |
| TD-003 — compact mapping preservation | The eight compact checklist records remain in `skills/_shared/security-guideline-catalog.md#guideline-checklist-records`; WU1 added an operational source-row layer below them and maps rows only to existing compact IDs. | Pass — no replacement compact controls were introduced. |
| TD-004 / TD-005 — source-row schema and allowed values | `skills/_shared/sdd-security-contract.md#source-row-operational-layer` defines source-row fields, allowed values, lifecycle/finding/route vocabulary, and compact mapping constraints. | Pass — static/manual contract evidence. |
| TD-006 / TD-007 / TD-008 — safe evidence and N/A policy | `skills/_shared/sdd-security-contract.md#source-row-operational-layer` requires safe evidence for applicable rows, evidence plus `naJustification` for every `N/A`, and rejects secrets, PII, PAN, tokens, connection strings, private keys, and confidential values. | Pass — static/manual contract evidence. |
| TD-011 — no duplicated 96-control matrix | WU1 only updates shared source-row security contracts; it states review-security may cite `review-report.md` but MUST NOT duplicate the general 96-control matrix. | Pass — no general review matrix was copied. |
| TD-015 — persistence compatibility | `skills/_shared/sdd-security-contract.md#source-row-routing-and-persistence-semantics` keeps source-row semantics identical across OpenSpec, Engram, hybrid, and none while delegating storage to the shared persistence contract. | Pass — backend-neutral semantics preserved. |
| TD-016 — unavailable tooling / no legacy validator dependency | This repo has no configured runtime/build/coverage/lint/type/format runner in `openspec/config.yaml`; WU1 uses static/manual Markdown evidence and explicitly states active source-row validation does not require `scripts/validate_security_design.ps1`. | Pass — no runtime command or legacy validator was claimed as required. |

## Work Unit 2: Design and Test-Design Contracts

- [x] 2.1 Update `skills/sdd-design/SKILL.md` to require `design.md#secure-development-design` Source ID coverage, evidence owners, lifecycle status, compact summary first, and N/A rationale. Covers TD-009, TD-018.
- [x] 2.2 Update `skills/sdd-test-design/SKILL.md` to require static/manual source-row planning, unavailable-tooling notes, warning preservation, and mandatory evidence blockers. Covers TD-010, TD-016.
- [x] 2.3 Align `agents/sdd/sdd-design.md` and `agents/sdd/sdd-test-design.md` prompts with WU2 contracts. Covers TD-017.

### Apply Evidence — Work Unit 2

| Check | Evidence | Result |
| --- | --- | --- |
| TD-009 — secure design Source ID coverage | `skills/sdd-design/SKILL.md` now requires `design.md#secure-development-design` to declare the expected Source ID universe, compact mappings, lifecycle status, evidence owners, downstream traceability, safe-evidence policy, and evidence-backed `N/A` policy when corporate source-row coverage applies. | Pass — static/manual contract evidence. |
| TD-018 — compact summary first | `skills/sdd-design/SKILL.md` now requires the compact eight-control summary to remain visible before large source-row detail and forbids duplicating the general 96-control review matrix. | Pass — review-load mitigation preserved. |
| TD-010 — source-row test planning | `skills/sdd-test-design/SKILL.md` now requires source-row planning from embedded design coverage, including Source ID/group, compact `SEC-*` mappings, lifecycle status, owner-phase evidence, mandatory/advisory status, warning preservation, and N/A checks. | Pass — static/manual planning contract evidence. |
| TD-016 — unavailable tooling notes | `skills/sdd-test-design/SKILL.md` now requires unavailable runtime/build/coverage/lint/typecheck/format tooling to be reported explicitly and never claimed as passing evidence. | Pass — no runtime command or legacy validator was claimed. |
| TD-017 — prompt alignment | `agents/sdd/sdd-design.md` and `agents/sdd/sdd-test-design.md` now remind executors to follow WU2 source-row contracts, preserve compact controls, use static/manual evidence when tooling is unavailable, and avoid legacy validators or duplicated 96-control matrices. | Pass — agent prompts aligned with phase skills. |

## Work Unit 3: Review-Security Contract

- [x] 3.1 Update `skills/sdd-review-security/SKILL.md` to validate every Source ID exactly once against design, test-design, apply evidence, changed files, and `review-report.md` without copying the 96-control matrix. Covers TD-011, TD-012.
- [x] 3.2 Encode blocker/warning routing for missing rows, malformed schema, missing mappings, unsafe evidence, unsupported N/A, missing implementation evidence, and warnings-only progression. Covers TD-001, TD-004-TD-008.
- [x] 3.3 Align `agents/sdd/sdd-review-security.md` with exhaustive source-row review and review-report citation boundaries. Covers TD-011, TD-017.

### Apply Evidence — Work Unit 3

| Check | Evidence | Result |
| --- | --- | --- |
| TD-011 — exhaustive source-row review | `skills/sdd-review-security/SKILL.md` now requires every expanded Source ID from `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` to be validated exactly once against `design.md#secure-development-design`, `test-design.md`, completed tasks/apply evidence, changed-file context, and `review-report.md`. | Pass — static/manual contract evidence. |
| TD-012 — evidence correlation | `skills/sdd-review-security/SKILL.md` now states source rows must not pass merely because they are listed; each row must correlate safe evidence for applicability, compliance, or justified `N/A`. | Pass — listed-only rows fail validation. |
| TD-001 / TD-004 / TD-005 — coverage, schema, and allowed values | `skills/sdd-review-security/SKILL.md` now blocks missing/duplicate/unknown Source IDs, malformed source-row schema, unsupported allowed values, and missing or unknown compact `SEC-*` mappings. | Pass — route is `resolve-blockers`. |
| TD-006 / TD-007 / TD-008 — safe evidence and N/A routing | `skills/sdd-review-security/SKILL.md` now rejects unsafe evidence, requires evidence plus `naJustification` for `N/A`, routes unsupported `N/A` to `resolve-blockers`, routes implementation evidence gaps to `apply`, and permits warnings-only progression to `verify` only when mandatory evidence is complete. | Pass — static/manual routing evidence. |
| TD-017 — prompt alignment | `agents/sdd/sdd-review-security.md` now reminds the executor to perform exhaustive source-row review, cite `review-report.md` only as supporting evidence, avoid copying the 96-control matrix, and follow source-row blocker routing. | Pass — agent prompt aligned with phase skill. |

## Work Unit 4: Verify, Archive, Spec Sync, Evidence

- [x] 4.1 Update `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` to consume non-blocking review-security source-row verdicts, block unresolved source blockers, and report unavailable tooling. Covers TD-013, TD-016.
- [x] 4.2 Update `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` to preserve coverage, mappings, warnings, exceptions, safe evidence refs, and no legacy standalone security artifact requirement. Covers TD-014.
- [x] 4.3 Sync source specs under `openspec/specs/` for catalog, design, test-design, review-security, and execution persistence contracts. Covers TD-015, TD-017.
- [x] 4.4 Record apply-time static/manual evidence in `apply-evidence-wu4.md`: 155-row coverage, compact mapping, safe evidence, N/A policy, no runtime/build/lint/type/format/coverage runners available. Covers TD-001-TD-018.

### Apply Evidence — Work Unit 4

| Check | Evidence | Result |
| --- | --- | --- |
| TD-013 — verify consumes source-row verdicts | `skills/sdd-verify/SKILL.md` now requires verify to consume the non-blocking `review-security-report.md` source-row verdict, warnings, exceptions, exact-once coverage, compact mapping status, safe-evidence status, and `N/A` justification status without owning or duplicating the full source-row matrix. `agents/sdd/sdd-verify.md` carries the same executor reminder. | Pass — static/manual contract evidence. |
| TD-013 — unresolved source blockers block verify | `skills/sdd-verify/SKILL.md` now routes missing/duplicate/unknown Source IDs, malformed schema, missing compact mappings, unsafe evidence, unsupported `N/A`, and missing mandatory source-row evidence to `apply` or `resolve-blockers` by cause. | Pass — blocker routing preserved. |
| TD-016 — unavailable tooling reported | `skills/sdd-verify/SKILL.md`, `agents/sdd/sdd-verify.md`, `skills/sdd-archive/SKILL.md`, and `agents/sdd/sdd-archive.md` now require runtime/build/lint/type/format/coverage tooling unavailability to be reported explicitly and never treated as passing evidence. `openspec/config.yaml#testing` confirms no such runners are configured. | Pass — no runtime/build/lint/type/format/coverage command was claimed. |
| TD-014 — archive preserves source-row audit trail | `skills/sdd-archive/SKILL.md` now requires archive reports to preserve source-row coverage summary, exact Source ID count, compact `SEC-*` mappings, warnings, complete exceptions, safe evidence refs, `N/A` evidence/justification status, review-security verdict, verify source-row consumption, and confirmation that no blockers remain. `agents/sdd/sdd-archive.md` mirrors this reminder. | Pass — archive preservation contract evidence. |
| TD-014 — no legacy standalone artifact requirement | `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` state active new-change archive readiness does not require standalone `security-design.md` or `scripts/validate_security_design.ps1`; embedded design, review-security, verify, and safe audit references are authoritative. | Pass — legacy validator/artifact remains non-blocking. |
| TD-015 / TD-017 — source specs synced | `openspec/specs/sdd-security-guideline-catalog/spec.md`, `openspec/specs/sdd-design-workflow/spec.md`, `openspec/specs/sdd-test-design-workflow/spec.md`, `openspec/specs/sdd-review-security-workflow/spec.md`, and `openspec/specs/sdd-execution-persistence-contracts/spec.md` now include the source-row inventory, design, test-design, review-security, verify, archive, and persistence requirements from this change. | Pass — source specs updated in-place without archiving the change folder. |
| TD-001 / TD-002 — 155-row expanded coverage | The synced catalog and design contracts preserve the expected corporate source-row universe of 155 expanded Source IDs and require range expansion before validation. This WU4 evidence cites the compact source specs and does not copy the full matrix. | Pass — exact coverage contract preserved. |
| TD-003 / TD-018 — compact mapping and review-load mitigation | WU4 keeps the eight compact `SEC-*` taxonomy as the control layer; verify/archive consume or preserve compact mappings and summaries without creating replacement compact controls or duplicating the full 96-control general review matrix. | Pass — compact layer preserved. |
| TD-006 / TD-007 / TD-008 — safe evidence and `N/A` policy | Verify/archive contracts require safe evidence refs only, reject unsafe evidence, preserve complete exceptions, and require evidence plus justification for `N/A`; unsupported `N/A` remains a blocker. | Pass — no secrets, PII, PAN, tokens, connection strings, private keys, or confidential values were added. |

### Final Apply-Time Static/Manual Evidence

| Evidence Area | Static/Manual Result |
| --- | --- |
| Source-row coverage | Expected source-row universe remains 155 expanded Source IDs across 15 corporate sections, with range expansion required before validation and exact-once coverage required by review-security, verify, and archive contracts. |
| Compact mapping | Every source row must map to one or more existing compact IDs: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, or `SEC-LOG-001`; no replacement compact controls were introduced. |
| Safe evidence | Evidence is limited to paths, section refs, sanitized summaries, command summaries, or redacted placeholders; unsafe raw secrets, PII, PAN, tokens, connection strings, private keys, and confidential values remain blockers. |
| N/A policy | `N/A` source rows require safe evidence plus `naJustification` proving irrelevance by category, platform, API, data class, or workflow; unsupported `N/A` routes to `resolve-blockers`. |
| Review/report boundaries | Verify and archive consume/cite review-security source-row summaries and warnings; they must not duplicate the full corporate source-row matrix or the general 96-control review matrix. |
| Tooling availability | `openspec/config.yaml#testing` reports no runtime test runner, build command, linter, type checker, formatter, or coverage command. Evidence is static/manual, and unavailable tooling is reported as unavailable rather than passed. |
| Persistence compatibility | Source-row semantics are identical across OpenSpec, Engram, hybrid, and none modes; backend choice changes storage/recovery behavior only. |

Additional WU4 apply evidence is preserved in `openspec/changes/corporate-source-row-security-validation/apply-evidence-wu4.md`. This is non-formal apply-time evidence; the formal `sdd-verify` phase must still run later after non-blocking `review-report.md` and `review-security-report.md` exist.
