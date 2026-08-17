# Test Design: Operational Readiness Evidence Flow

## Overview

This change updates Markdown SDD workflow contracts, not application runtime code. Test planning therefore uses static, documentary, and manual checks to prove that operational-readiness evidence is planned, validated, preserved, archived, and consumed safely without requiring real production operational data. `openspec/config.yaml` reports no executable test runner, build, linter, type checker, formatter, or coverage tooling, so those missing tools are unavailable evidence and must not be treated as passing evidence.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/operational-readiness-evidence-flow/proposal.md` | Scope, non-goals, affected phase list, restricted-data boundary, exact marker success criteria. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-operational-readiness-workflow/spec.md` | Mandatory readiness evaluation, safe SDD evidence boundary, phase ownership, manual operational document consumption. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-design-workflow/spec.md` | Required `design.md#Operational Readiness` section, exact marker behavior, and design safety constraints. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-test-design-workflow/spec.md` | Required readiness test planning, placeholder checks, monitoring coverage, and restricted-data checks. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-review-workflow/spec.md` | General review checks outside the fixed 96-control matrix and readiness handoff to security review. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-review-security-workflow/spec.md` | Leakage, secrets, restricted production identifiers, safe placeholders, and security-review boundaries. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-execution-persistence-contracts/spec.md` | Verify/archive evidence persistence, pending carry-forward, manual utility boundary, and final-document restricted-data boundary. |
| Spec | `openspec/changes/operational-readiness-evidence-flow/specs/sdd-security-guideline-catalog/spec.md` | Operational safe-evidence policy, monitoring mechanisms beyond SQL, restricted operational data classification. |
| Design | `openspec/changes/operational-readiness-evidence-flow/design.md` | Shared contract architecture, phase-local responsibilities, data flow, affected files, operational-readiness strategy, testing constraints. |
| Secure Development Design | `openspec/changes/operational-readiness-evidence-flow/design.md#secure-development-design` | Changed-surface classification, applicable narrative rules, evidence owners, safe-evidence policy, residual risks, and no-exception policy. |
| Testing Capabilities | `openspec/config.yaml#testing` | Confirms unavailable runtime, build, coverage, lint, type-check, formatter, unit, integration, and e2e tooling. |

## Source ID Coverage Baseline

Corporate source-row coverage applies because the design classifies this as security-impacting instruction work that governs operational evidence and leakage prevention. This test design consumes only narrative design rules and changed-surface context from `design.md#Secure Development Design`:

- Applicable narrative categories: sensitive operational identifiers, secrets, sensitive logging and monitoring evidence, files/generated artifacts, safe placeholders, final manual document boundary, and security exception evidence policy.
- Out-of-scope catalog rows such as authentication, sessions, permissions/access-control, and database-access remain reviewable omissions owned by `sdd-review-security`; this artifact does not create exhaustive `N/A` rows.
- `review-security-report.md` owns compact/source-row expansion, omitted-row validation, and matrix/source-row decisions.
- This artifact must not require design YAML, schema fields, compact controls, Source ID matrices, machine-readable applicability fields, all-row `N/A` bookkeeping, or the full source-row matrix.
- Evidence is static/manual/documentary because runtime/build/lint/type/format/coverage tooling is unavailable.

Unavailable evidence that must be reported downstream, not counted as passed:

| Tooling Area | Availability | Planned Substitute Evidence |
| --- | --- | --- |
| Runtime test runner | Unavailable | Static/manual inspection of Markdown contracts and generated phase artifacts. |
| Build command | Unavailable | Static contract and path-reference validation. |
| Unit/integration/e2e layers | Unavailable | Scenario-to-check traceability table in this artifact and downstream review evidence. |
| Coverage command | Unavailable | Manual coverage mapping from specs/design risks to test-design cases. |
| Linter | Unavailable | Static Markdown inspection for required headings, exact marker strings, and forbidden data examples. |
| Type checker | Unavailable | Not applicable to Markdown-only contract changes; record as unavailable. |
| Formatter | Unavailable | Manual Markdown readability/section preservation review. |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Mandatory Operational Readiness Evaluation; Design: `## Operational Readiness` | Confirm every new/updated SDD design contract requires evaluating logs, monitoring mechanisms, administration, reprocessing, ownership, final documentation inputs, and unresolved gaps. | static | mandatory | Changed files and design/task/review references showing required readiness categories. | Evaluation is mandatory even when data disclosure is unsafe or unavailable. |
| TD-002 | Proposal success criteria; specs: exact markers | Validate exact marker contract uses `Pendiente de confirmar:` for unresolved applicable fields and `No aplica.` for inapplicable fields, with exact casing and punctuation. | static | mandatory | Static inspection of shared contract and affected phase skills for exact strings. | Synonyms, lowercase variants, missing colon, or translations fail. |
| TD-003 | Spec: Operational Readiness Test Planning | Confirm `sdd-test-design` requires planned static/documentary/manual checks for readiness evidence when no runtime runner exists. | static | mandatory | Updated `skills/sdd-test-design/SKILL.md` with unavailable-tooling note and static/manual check guidance. | Missing tooling is unavailable evidence, not passing evidence. |
| TD-004 | Design: centralized readiness semantics | Confirm the shared readiness contract exists with required sections: Purpose, Evidence Categories, Exact Marker Contract, Restricted Operational Data Boundary, Phase Ownership Boundary, Design Artifact Section Contract, Test/Task/Review/Verify/Archive Handoff, Manual Operational Document Boundary, and Safe Evidence Examples. | static | mandatory | `skills/_shared/sdd-operational-readiness-contract.md` section inspection. | Section names may be extended, but required meanings must remain present. |
| TD-005 | Spec: Phase Ownership Model | Confirm each phase consumes only its local readiness responsibility and references the shared contract rather than duplicating the full policy. | static | mandatory | Affected `skills/sdd-*` files cite shared contract and define phase-local checks only. | Prevents contract drift and over-broad phase ownership. |
| TD-006 | Design: monitoring mechanisms; catalog spec | Confirm monitoring evidence is mechanism-oriented and not SQL-only by default. Accept dashboards, alerts, jobs, traces, scripts, documented manual checks, and SQL only where appropriate. | static | mandatory | Shared readiness/security contract and phase skills mention non-SQL monitoring mechanisms. | SQL-only assumptions fail unless explicitly justified for the change context. |
| TD-007 | Safe SDD Evidence Boundary; Secure Development Design | Scan ordinary SDD evidence guidance, examples, tests, fixtures, and snippets for restricted operational data categories: production hostnames, IPs, ports, SID/service names, credentials, tokens, private keys, connection strings, payloads, full ID lists, generated file bytes, PAN, PII, or confidential values. | static | mandatory | Review/security-review planned evidence plus static changed-file inspection results. | Use pattern-based/manual inspection; do not require real production values. |
| TD-008 | Spec: Operational Data Safety Checks | Verify safe SDD evidence is separated from final operational document inputs and that final restricted values may appear only when explicitly user-provided for the final manual document. | static | mandatory | Shared contract and `sdd-operational-doc` skill boundary text. | SDD evidence must not be backfilled with final document values. |
| TD-009 | Review workflow spec; Design decision: separate quality review from leakage review | Confirm `sdd-review` adds readiness checks outside the fixed 96-control matrix while preserving the matrix shape. | static | mandatory | `skills/sdd-review/SKILL.md` and review-report contract include separate readiness section; 96-control matrix columns remain unchanged. | Readiness rows must not be inserted into the fixed matrix. |
| TD-010 | Review workflow: Operational Review Handoff | Confirm `sdd-review` hands off changed-file, evidence, placeholder, and unresolved-gap context to `sdd-review-security` without owning leakage verdicts. | static | mandatory | Review-report contract includes readiness refs/gaps handoff. | General review owns existence, traceability, no-invention, exact placeholders. |
| TD-011 | Review-security workflow spec | Confirm `sdd-review-security` owns leakage checks for restricted production identifiers, secrets, payloads, generated bytes, and equivalent operational data in ordinary SDD artifacts, code, tests, fixtures, and examples. | static | mandatory | `skills/sdd-review-security/SKILL.md` and security contract mention leakage categories and blocking behavior. | Security review does not require real operational data disclosure to pass. |
| TD-012 | Secure Development Design: review-security matrix ownership | Confirm `sdd-review-security` remains owner of compact/source-row expansion, matrix/source-row decisions, omitted-row validation, and `N/A` decisions. | static | mandatory | `skills/sdd-review-security/SKILL.md` and shared security contract preserve review-security ownership. | Test-design must not duplicate the matrix or demand Source IDs from design. |
| TD-013 | Execution/persistence spec: Verify checks readiness completion | Confirm `sdd-verify` validates every readiness field as safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.`, and carries unresolved pending fields forward instead of inventing data. | static | mandatory | `skills/sdd-verify/SKILL.md` includes completeness and pending carry-forward checks. | Verify must preserve unavailable-tooling notes. |
| TD-014 | Execution/persistence spec: Archive preserves handoff evidence | Confirm `sdd-archive` preserves readiness status, evidence refs, unresolved gaps, warning carry-forward, and final-document handoff boundaries. | static | mandatory | `skills/sdd-archive/SKILL.md` and archive report guidance include readiness preservation. | Archive must not require `sdd-operational-doc` execution. |
| TD-015 | Manual Operational Document Consumption | Confirm manual `sdd-operational-doc` reads archived readiness evidence first, preserves sections 1-9 and diagrams R1-R4, and writes unresolved fields as exact `Pendiente de confirmar:`. | manual | mandatory | Inspection of `skills/sdd-operational-doc/SKILL.md` and `assets/operational-document-template.md`. | Utility remains manual post-archive, not a DAG phase. |
| TD-016 | Proposal non-goal: no real operational data requirement | Confirm all affected instructions say evaluation is mandatory but disclosure of real production data is not required when unavailable or unsafe. | static | mandatory | Shared readiness contract and phase skill wording. | Prevents agents from inventing operational values. |
| TD-017 | Spec: Safe Placeholder Security Boundary | Confirm exact placeholders are accepted as safer than invented operational details while still requiring safe proof of non-leakage where security obligations apply. | static | mandatory | `sdd-review-security` guidance distinguishes placeholder safety from boundary proof. | Placeholder-only evidence cannot hide missing leakage checks. |
| TD-018 | Design data flow | Confirm readiness evidence flows from design through test-design, tasks, apply evidence, review, review-security, verify, archive, and then manual operational-doc consumption. | static | mandatory | Cross-file references and downstream phase contracts align with design data flow. | Missing phase handoff breaks downstream consumption. |
| TD-019 | Tasks phase implied by specs/design | Confirm `sdd-tasks` will create concrete collection/validation work for readiness evidence, placeholder validation, safe-evidence checks, archive handoff, and unavailable-tooling notes. | static | mandatory | `skills/sdd-tasks/SKILL.md` task planning requirements. | Ensures implementation work is not left implicit. |
| TD-020 | Markdown documentation quality | Confirm readiness guidance is chunked, headed, and reviewable with clear ownership boundaries and checklist/table-friendly evidence. | manual | non-mandatory | Manual review of changed Markdown for cognitive load and reviewer scanability. | Advisory; unresolved issues should be review warnings unless they obscure mandatory evidence. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-OP-001` | Ordinary SDD evidence must not contain production hostnames, IPs, ports, SID/service names, credentials, tokens, private keys, connection strings, sensitive payloads, full ID lists, generated file bytes, PAN, PII, or confidential values. | Yes | TD-007, TD-011 | covered | None |
| `SEC-OP-002` | Safe evidence may cite paths, section anchors, sanitized summaries, command summaries, redacted placeholders, and exact unresolved markers. | Yes | TD-001, TD-002, TD-008, TD-013, TD-014 | covered | None |
| `SEC-OP-003` | Monitoring evidence must be mechanism-oriented and must not assume SQL-only monitoring coverage. | Yes | TD-006 | covered | None |
| `SEC-OP-004` | Secrets and secret-like values must never appear in SDD artifacts, examples, tests, fixtures, review reports, verify reports, archive summaries, or generated evidence snippets. | Yes | TD-007, TD-011 | covered | None |
| `SEC-OP-005` | Raw logs, stack traces with sensitive payloads, user identifiers, credentials, and confidential operational context must not be copied into readiness evidence. | Yes | TD-006, TD-007, TD-011 | covered | None |
| `SEC-OP-006` | Generated file bytes, full exported file contents, and environment-specific payloads must be referenced or summarized safely, not embedded in ordinary SDD evidence. | Yes | TD-007, TD-014, TD-015 | covered | None |
| `SEC-OP-007` | Final operational documentation may include restricted operational values only when explicitly user-provided for the final manual document; ordinary SDD evidence must not be backfilled with those values. | Yes | TD-008, TD-015, TD-016 | covered | None |
| `SEC-OP-008` | `sdd-review-security` owns leakage validation plus compact/source-row expansion, omitted-row validation, and `N/A` decisions. | Yes | TD-011, TD-012 | covered | None |
| `SEC-OP-009` | No security exception is planned; any future exception must include approver, approval date, accepted-risk rationale, mitigation/follow-up, and exact evidence gap. | Yes | TD-017 plus downstream review-security exception checks | covered | None |

## No-Impact Assessment

Not applicable. The change intentionally affects SDD lifecycle contracts, phase responsibilities, operational evidence handling, security leakage review, archive handoff, and the manual operational document utility. Static/documentary/manual evidence is required because the repository has no executable tooling.

## Evidence Expectations

- Mandatory cases require implementation, static/manual evidence, or a justified skip approved by the owning phase; uncovered mandatory readiness or security-boundary evidence blocks downstream verification.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Missing runtime, build, unit, integration, e2e, coverage, lint, type-check, and formatter tooling must be recorded as unavailable evidence, never as passed evidence.
- Static evidence should cite exact files, headings, section anchors, changed lines, or artifact references; it should not copy restricted operational values.
- Manual evidence should state what was inspected and whether sections, markers, diagrams, and phase boundaries were preserved.
- `sdd-review` readiness checks belong outside the fixed 96-control matrix and should report existence, traceability, no invention, exact placeholders, evidence refs, and unresolved gaps.
- `sdd-review-security` owns leakage validation, safe-evidence boundaries, secrets/restricted identifier checks, and review-security matrix/source-row decisions.
- `sdd-verify` must confirm every readiness field is evidenced, `Pendiente de confirmar:`, or `No aplica.`, and must carry unresolved pending fields forward.
- `sdd-archive` must preserve readiness status, evidence references, unresolved gaps, warning carry-forward, and manual-document handoff notes.
- Manual `sdd-operational-doc` must consume archived evidence after archive, preserve sections 1-9 and diagrams R1-R4, and must not invent missing operational data.
- Test-design consumes narrative design rules only and must not require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, exhaustive `N/A` rows, or the full source-row matrix; exhaustive expansion belongs to `review-security-report.md`.

## Open Questions

None.
