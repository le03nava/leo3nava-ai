# sdd-review-workflow Specification

## Purpose

Define `sdd-review` as the mandatory post-apply code-review gate that produces canonical `review-report.json` and derived `review-report.md` before security review and verification.

## Requirements

### Requirement: Mandatory Review Gate

The SDD workflow MUST route completed implementation through `sdd-review` before `sdd-review-security` and `sdd-verify`. Review MUST inspect applied changes, required new-change SDD artifacts, and task evidence, then route non-blocking results to security review rather than directly to verify. `security-applicability.md` and standalone `security-design.md` MUST NOT be default required new-change inputs.

#### Scenario: Apply routes to review

- GIVEN all apply tasks are complete
- WHEN apply returns a success route
- THEN the next phase MUST be `sdd-review`
- AND verify MUST NOT run until both review gates are non-blocking.

#### Scenario: Blocking review returns to apply

- GIVEN review finds any blocking failure
- WHEN the review phase completes
- THEN it MUST route to `sdd-apply`
- AND `sdd-review-security` MUST NOT run yet.

#### Scenario: Legacy applicability is optional evidence

- GIVEN a new change has proposal, specs, design with `## Secure Development Design` narrative rules, test design, tasks, and apply evidence
- WHEN review resolves required inputs
- THEN missing `security-applicability.md` or standalone `security-design.md` MUST NOT block review
- AND any such artifact present MUST be treated only as legacy or archive evidence.

### Requirement: Review Report Artifact

`sdd-review` MUST persist canonical `review-report.json` and generate derived `review-report.md` in the same phase. OpenSpec mode MUST expose both under `openspec/changes/{change-name}/`. Missing required artifacts, unknown changed files, unsafe workspace context, JSON validation failure, Markdown generation/parity failure, stale Markdown, or persistence failure MUST route to `resolve-blockers`.

#### Scenario: Reports are persisted

- GIVEN review can resolve all required inputs
- WHEN review completes
- THEN it MUST write canonical `review-report.json`
- AND it MUST write derived `review-report.md` with verdict, matrix, evidence summary, and next recommendation.

#### Scenario: Markdown generation fails

- GIVEN canonical JSON is valid
- WHEN derived Markdown cannot be generated or read back
- THEN review MUST route to `resolve-blockers`
- AND downstream phases MUST NOT consume stale Markdown as current evidence.

### Requirement: Code-Review Validation Matrix

The canonical JSON MUST define the review matrix source rows and validation rules. Derived `review-report.md` MUST preserve the exact matrix columns: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. The derived matrix MUST contain 96 rows, use stable `REV-CORP-*` IDs, preserve category without adding columns, and limit `Complies` to `Yes`, `No`, or `N/A`. `N/A` MUST include Evidence Location proving irrelevance and Observations/Comments explaining scope.

#### Scenario: All controls are represented

- GIVEN canonical JSON contains 96 controls
- WHEN review writes JSON and derived Markdown
- THEN every control MUST appear once with a stable Item ID
- AND every Markdown row MUST use the exact required columns.

#### Scenario: Platform control is irrelevant

- GIVEN a control targets an unused platform or technology
- WHEN review marks it `N/A` in JSON
- THEN derived Markdown Evidence Location MUST prove irrelevance
- AND Observations/Comments MUST explain the scope decision.

#### Scenario: Derived Markdown preserves compatibility sections

- GIVEN canonical JSON has report facts
- WHEN Markdown is generated
- THEN it MUST preserve verdict, blocking summary, evidence summary, review matrix, applicable operational evidence, changed-file/security handoff, matrix validation, and recommendation sections.

### Requirement: Canonical Review JSON Authority

`sdd-review` MUST treat canonical JSON as the source of truth for general review report facts and all 96 `REV-CORP-*` controls. The JSON MUST directly define control IDs, source mappings, required report facts, vocabulary, validation metadata, and presentation-ready fields. `control-catalog.md` MUST NOT remain the authoritative control source; it MAY be replaced or kept only as a derived/reference view. Security-review ownership and source-row security review MUST remain unchanged.

#### Scenario: JSON owns controls

- GIVEN the review phase needs the corporate controls
- WHEN it loads review control definitions
- THEN it MUST load them from canonical JSON
- AND every `REV-CORP-001..REV-CORP-096` control MUST appear exactly once.

#### Scenario: Markdown catalog is not authority

- GIVEN `control-catalog.md` exists
- WHEN its content differs from canonical JSON
- THEN JSON MUST win
- AND the mismatch MUST be treated as a derived-view repair issue, not a control decision.

### Requirement: Review Artifact Authority Boundary

Persistence and shared post-apply contracts MUST expose both canonical JSON and derived Markdown in OpenSpec mode. Downstream phases MUST consume the selected review artifact identity without treating Markdown as authoritative. The JSON SHOULD be presentation-friendly for future Excel generation, but this change MUST NOT implement Excel, Python, or spreadsheet output.

#### Scenario: Downstream consumes derived Markdown safely

- GIVEN review completed without blockers
- WHEN security review, verify, or archive consumes review evidence
- THEN it MAY read `review-report.md` for compatibility
- AND it MUST recognize canonical JSON as authoritative when facts conflict.

#### Scenario: Excel is deferred

- GIVEN canonical JSON includes presentation metadata
- WHEN this change is implemented
- THEN no Excel/Python generation MUST be added.

### Requirement: Review JSON Authority and Derived Compatibility

`sdd-review` MUST treat `skills/sdd-review/references/review-control-catalog.json` as the authoritative static source for all 96 general review controls and `review-report.json` as the authoritative per-change source for review facts. `control-catalog.md` and `review-report.md` MAY be read for human or downstream compatibility, but they MUST NOT override canonical JSON. JSON SHOULD include presentation metadata useful to future renderers, but this workflow MUST NOT add Excel, Python, script, spreadsheet, or workbook generation.

#### Scenario: JSON wins over Markdown

- GIVEN canonical JSON and derived Markdown disagree
- WHEN the workflow decides review facts, controls, counts, routing, or validation status
- THEN the JSON MUST win
- AND Markdown MUST be treated as the repair target.

#### Scenario: Spreadsheet generation remains deferred

- GIVEN canonical JSON includes presentation metadata
- WHEN `sdd-review` runs or shared contracts consume its artifacts
- THEN no Excel, Python, script, spreadsheet, or workbook output MUST be generated by this workflow.

### Requirement: Severity and Routing Semantics

Review MUST classify each failed control severity. Critical or explicitly blocking failures MUST prevent security review and verify. Non-blocking findings MAY proceed to `sdd-review-security` as warnings.

#### Scenario: Non-blocking review proceeds

- GIVEN review has no blocking failures
- WHEN it returns its envelope
- THEN next_recommended MUST be `sdd-review-security`
- AND the security review phase MUST consume the review report as evidence.

### Requirement: Security Boundary

Security-related review controls MUST cross-reference the security guideline catalog where applicable. Review MUST NOT replace embedded secure development design or mandatory security review ownership. Legacy applicability or standalone security-design evidence MAY be cited only for archived or old changes and MUST NOT displace `design.md` or canonical `review-security-report.json` as new-change security authorities.

#### Scenario: Security control is reviewed

- GIVEN a review control maps to a catalog guideline
- WHEN review records the matrix row
- THEN Standard SHOULD cite the guideline identifier
- AND embedded `design.md` security rows and canonical `review-security-report.json` MUST remain the security authorities.

#### Scenario: Applicability does not replace security evidence

- GIVEN a reviewer finds a legacy `security-applicability.md` or `security-design.md` artifact
- WHEN reviewing a new change
- THEN the reviewer MAY cite it only as compatibility context
- AND MUST require embedded `design.md` secure development rows for authoritative classification.

### Requirement: Security Review Handoff

The canonical review JSON and derived `review-report.md` MUST provide enough changed-file, evidence, finding context, and `design.md` references for `sdd-review-security` to validate embedded secure development implementation evidence without duplicating the 96-control matrix. Downstream phases MAY read derived Markdown for compatibility, but when canonical JSON is present it remains authoritative for general review facts.

#### Scenario: Handoff evidence is available

- GIVEN review completed without blockers
- WHEN security review starts
- THEN it MUST be able to read `review-report.json` or derived `review-report.md`, plus `design.md`
- AND missing readable review or design evidence MUST route to `resolve-blockers`.

### Requirement: Operational Readiness General Review

`sdd-review` MUST validate operational evidence only when design, test-design, tasks, apply evidence, or archived context says operational considerations apply. Review MUST assess traceability, no-invention behavior, and placeholder usage for present or planned evidence without changing the fixed 96-control matrix shape. Review MUST NOT enforce mandatory operational category completeness or require real operational data disclosure.

#### Scenario: Operational evidence is traceable

- GIVEN applicable operational evidence is present in SDD artifacts
- WHEN general review runs
- THEN it MUST verify the evidence cites safe sources, `Pendiente de confirmar:`, or `No aplica.`.

#### Scenario: Review detects invented data

- GIVEN operational evidence contains unsupported operational details
- WHEN review validates traceability
- THEN the finding MUST be blocking or routed to `resolve-blockers`.

#### Scenario: Operational evidence is absent by design

- GIVEN design marks operational considerations not applicable
- WHEN review runs
- THEN absence of readiness categories MUST NOT be a blocker.

#### Scenario: Matrix shape is preserved

- GIVEN operational review findings are recorded
- WHEN `review-report.md` is generated from `review-report.json`
- THEN the 96-control matrix columns MUST remain unchanged
- AND operational evidence MAY appear in a separate section.

### Requirement: Operational Review Handoff

Review MUST hand off changed-file, evidence, placeholder, and gap context for operational considerations that exist. It MUST NOT create a handoff requirement from an absent shared readiness contract.

#### Scenario: Security review receives applicable context

- GIVEN general review finds applicable operational evidence
- WHEN security review starts
- THEN it MUST be able to read evidence locations and unresolved gaps from review output.

#### Scenario: No applicable context exists

- GIVEN no operational considerations apply
- WHEN security review starts
- THEN missing readiness handoff MUST NOT block the phase.
