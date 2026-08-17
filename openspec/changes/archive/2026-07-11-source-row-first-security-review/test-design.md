# Test Design: Source Row First Security Review

## Overview

This change is a security-impacting SDD contract redesign. Testing must prove that canonical `review-security-report.json` is source-row-first, that `sourceRowValidation.rows` is the only active security matrix, and that downstream Markdown, verify/archive, and Excel export behavior consume canonical JSON without recreating or overriding validation. Repository-wide runtime/build/lint/type/format/coverage tooling is unavailable, so most phase evidence is static inspection, schema/sample validation planning, and documentary read-back; Python exporter pytest coverage is planned when dependencies are available.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/source-row-first-security-review/proposal.md` | Scope, non-goals, affected areas, risks, rollback, and success criteria. |
| Spec: security workflow | `openspec/changes/source-row-first-security-review/specs/sdd-review-security-workflow/spec.md` | Canonical report, exact-once source-row validation, generated Markdown, and JSON authority scenarios. |
| Spec: guideline catalog | `openspec/changes/source-row-first-security-review/specs/sdd-security-guideline-catalog/spec.md` | 155-row catalog inventory, source-row vocabulary, grouping metadata, and safe evidence policy. |
| Spec: persistence contracts | `openspec/changes/source-row-first-security-review/specs/sdd-execution-persistence-contracts/spec.md` | Verify/archive consumption rules, backend persistence behavior, and canonical/derived artifact authority. |
| Spec: Excel exporter | `openspec/changes/source-row-first-security-review/specs/canonical-review-json-excel-exporter/spec.md` | Default table path, compact-only default failure, manual nested override, README, and pytest expectations. |
| Design | `openspec/changes/source-row-first-security-review/design.md` | Architecture decisions, data flow, file changes, contracts, risks, rollout, and testing strategy. |
| Secure Development Design | `openspec/changes/source-row-first-security-review/design.md#secure-development-design` | Changed-surface classification, applicable narrative security rules, owner phases, safe-evidence rules, residual risks, and exception policy. |
| Testing Capabilities | `openspec/config.yaml#testing` | No global runner, build, coverage, linter, type checker, or formatter detected; Python pytest may be available only when dependencies are installed. |

## Source ID Coverage Baseline

The canonical source-row baseline is the corporate 155 Source ID snapshot carried by `skills/sdd-review-security/references/security-guideline-catalog.operational.json` after implementation. Test planning does not duplicate the full matrix here and does not require design-time source rows. Instead, implementation evidence must prove that the review-security validator/catalog expansion compares the active report to the expanded Source ID inventory exactly once.

Required coverage evidence:

- `sourceRowValidation.rows` is required and contains exactly 155 entries.
- `sourceRowValidation.expectedCount` and `validatedCount` are `155` for non-blocking reports.
- The row `sourceId` set equals the expanded catalog Source ID set with no missing, duplicate, or unknown IDs.
- Each row has the required field set: `sourceId`, `corporateSection`, `pciAlignment`, `guidelineText` or `guidelineRefs`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, and `route`.
- Source-row fields such as `controlDomain`, `corporateSection`, `repoProfiles`, `runtimeSurface`, or `dataSurface` are the only active grouping/navigation fields.
- Legacy compact-control identifiers may be mentioned only as forbidden legacy input in tests/docs; they must not be active validation, navigation, summary, metadata, or grouped non-applicability authority.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Security Review Artifact; Design: Canonical JSON shape | Inspect `review-security-report.schema.json` to confirm `sourceRowValidation.rows` is required, has min/max/exact count semantics for 155 rows, includes source-row summary fields, and does not expose a parallel compact validation object. | static | mandatory | Static inspection notes citing schema path and relevant properties. | JSON Schema may need custom validator support for catalog set equality. |
| TD-002 | Spec: Security Matrix Validation; Design: Validation invariants | Validate positive sample JSON with 155 unique Source IDs and required row fields. | automated | mandatory | Schema/sample validation command result when available, or static validator walkthrough if no runner exists. | Sample data should use sanitized placeholder evidence only. |
| TD-003 | Spec: Exact-once row validation | Validate negative sample with one missing Source ID. | automated | mandatory | Failing validation evidence naming missing Source ID location/count without echoing unsafe payloads. | Failure should block downstream phases. |
| TD-004 | Spec: Exact-once row validation | Validate negative sample with one duplicate Source ID. | automated | mandatory | Failing validation evidence naming duplicate `sourceId` and exact-once failure. | Prefer validator rule over schema-only assertion if schema cannot express uniqueness against catalog. |
| TD-005 | Spec: Exact-once row validation | Validate negative sample with one unknown Source ID outside the catalog. | automated | mandatory | Failing validation evidence naming unknown `sourceId` safely. | Unknown rows cannot compensate for missing catalog rows. |
| TD-006 | Spec: Security Matrix Validation; Design: Source row fields | Validate negative samples missing each required row-field category, including the `guidelineText`/`guidelineRefs` either-or rule. | automated | mandatory | Schema/sample validation failures plus field checklist. | Can be table-driven tests or fixture mutation tests. |
| TD-007 | Proposal success criteria; Design: Compact authority removal | Static grep/inspection proves active schema, validation rules, template, skill, shared contracts, verify/archive skills, and adapter prompts do not use legacy compact-control identifiers as report authority. | static | mandatory | File-by-file inspection evidence for changed surfaces. | Historical/archive mentions are allowed only when clearly non-authoritative. |
| TD-008 | Spec: Grouped N/A preserves rows; Design: Safe evidence | Validate `N/A` rows require row-level `applies`, `complies`, `justification`, `evidenceType`, `evidenceLocation`, `finding`, `ownerPhase`, and `route`. | automated | mandatory | Positive and negative sample validation evidence. | Group summaries cannot replace row-level fields. |
| TD-009 | Spec: Safe Source Row Evidence; Secure Development Design | Validate unsafe row evidence is rejected or represented through `unsafeEvidenceRejections` plus safe row findings. | automated/static | mandatory | Validator evidence or static rule inspection showing secret/PII/PAN/token/raw-log payloads are not copied into reports. | Missing runtime tooling requires static/manual proof, not a pass. |
| TD-010 | Spec: Grouping cannot hide gaps; Design: Grouping uses source-row category fields | Validate grouped non-applicability only uses source-row grouping fields and fails or splits rows when any grouped row lacks equivalent justification. | automated/static | mandatory | Sample validation failures or validation-rule inspection. | Grouped text must not become validation authority. |
| TD-011 | Spec: Lean generated Markdown; Design: Markdown Template Simplification Strategy | Inspect `report-template.md` to confirm Markdown is generated from canonical JSON, has lean navigation/summaries, renders the full source-row matrix last, and does not duplicate validation logic. | static | mandatory | Template inspection evidence and section-order checklist. | Markdown is compatibility only. |
| TD-012 | Spec: JSON wins; Design: Markdown parity | Create or inspect Markdown parity tests comparing JSON and generated Markdown for verdict, route, counts, source refs, blockers/warnings, exceptions, unavailable tooling, and artifact metadata. | automated/static | mandatory | Parity test evidence or static generator inspection. | Disagreement routes to blockers because canonical JSON wins. |
| TD-013 | Spec: Verify consumes review evidence; Design: Downstream Verify / Archive Consumption | Inspect `skills/sdd-verify/SKILL.md`, shared contracts, and `agents/sdd/sdd-verify.md` to confirm verify consumes canonical JSON summaries/refs and does not copy or rescore the source-row matrix. | static | mandatory | Static inspection evidence with changed-file refs. | Verify may check completeness status/counts without owning validation logic. |
| TD-014 | Spec: Archive Source Row Preservation; Design: Downstream Verify / Archive Consumption | Inspect `skills/sdd-archive/SKILL.md`, shared contracts, and `agents/sdd/sdd-archive.md` to confirm archive preserves canonical refs/summaries and does not create a second matrix authority. | static | mandatory | Static inspection evidence with changed-file refs. | Archive should preserve refs before derived Markdown where structured refs support authority. |
| TD-015 | Spec: Catalog snapshot is available; Design: Catalog Transformation Approach | Inspect transformed operational catalog to confirm expanded 155 concrete Source IDs, source-row grouping metadata, safe-evidence expectations, owner phase, and route vocabulary are present. | static/automated | mandatory | Catalog self-check result or static count/evidence notes. | Range notation must be expanded before validation. |
| TD-016 | Spec: Schema-aware table selection; Design: Excel Exporter Impact | Pytest proves schema `sdd-review-security.review-security-report` defaults to `sourceRowValidation.rows` without `--table`. | automated | mandatory | `python -m pytest python/tests` result when dependencies are available; otherwise unavailable-tooling note plus static code inspection. | Pytest is scoped to Python exporter, not global repo tests. |
| TD-017 | Spec: Legacy compact path is not default; Proposal success criteria | Pytest proves compact-only security-review JSON fails by default because `sourceRowValidation.rows` is missing. | automated | mandatory | Pytest failure-case assertion and sanitized CLI error text. | Manual override remains a generic exporter feature, not active contract. |
| TD-018 | Spec: Manual nested table path; Design: Excel Exporter Impact | Pytest preserves generic `--table` nested dotted path behavior, including a manual historical compact path override if implemented. | automated | non-mandatory | Pytest fixture showing selected nested list exports successfully. | Do not document manual compact override as active default. |
| TD-019 | Spec: Pytest verification; Design: Excel Exporter Impact | Pytest validates workbook sheet naming, row flattening, formatting, generated output path, and workbook read-back through `openpyxl`. | automated | mandatory | Pytest result and read-back assertions. | Generated workbooks should use temp paths and not be committed. |
| TD-020 | Spec: User documentation; Secure Development Design | Inspect `python/README.md` to confirm source-row default, virtualenv/dependency setup, `--table`, and dependency policy are documented with safe examples. | static | mandatory | README inspection evidence. | README must not present legacy compact-control export as active contract. |
| TD-021 | Design: Operational Considerations | Inspect implementation evidence for safe artifact refs, generated Markdown refs, exporter evidence, contract propagation refs, and absence of restricted operational data. | static/manual | mandatory | Sanitized evidence table in apply/verify/review-security artifacts. | No hostnames, credentials, payloads, logs, or workbook bytes are required. |
| TD-022 | Testing capabilities | Carry forward unavailable global tooling as explicit unavailable evidence. | manual/static | mandatory | Verify/review notes stating no global runner/build/lint/type/format/coverage command was available. | Missing tooling is not passing evidence. |

## JSON Schema and Sample Validation Strategy

| Sample / Fixture | Purpose | Expected Result | Evidence Owner |
| --- | --- | --- | --- |
| `valid-source-row-report` | 155 unique source rows, required fields, valid counts, valid grouping fields, safe evidence. | Passes schema and custom exact-once validation. | apply / review-security |
| `missing-source-row` | 154 catalog rows plus otherwise valid metadata. | Blocks with missing Source ID/count failure. | apply / review-security |
| `duplicate-source-row` | 155 rows with one duplicate and one catalog row displaced. | Blocks with duplicate and missing row failure. | apply / review-security |
| `unknown-source-row` | 155 rows with an ID outside the expanded catalog. | Blocks with unknown and missing row failure. | apply / review-security |
| `missing-required-field` | Row missing each required field category in turn. | Blocks with safe field-location errors. | apply / review-security |
| `missing-guideline-reference` | Row contains neither `guidelineText` nor `guidelineRefs`. | Blocks. | apply / review-security |
| `unsafe-evidence` | Row evidence contains a secret-like, token-like, PAN/PII-like, raw log, or sensitive payload placeholder. | Blocks or records `unsafeEvidenceRejections` without echoing the raw payload. | review-security |
| `compact-only-security-report` | Security report contains only legacy compact row data. | Fails active security validation and fails exporter default. | apply / verify |
| `markdown-json-mismatch` | Generated Markdown intentionally disagrees with canonical JSON counts or route. | Blocks parity/read-back. | review-security / verify |

Implementation may express exact 155 count partly in JSON Schema and partly in a validator because schema alone may not compare the row set against the expanded catalog inventory. Tests must prove the combined contract, not only JSON Schema syntax.

## Safe N/A Grouping Tests

| ID | Check | Type | Severity | Expected Evidence |
| --- | --- | --- | --- | --- |
| NA-001 | A row marked non-applicable keeps row-level `applies`, `complies`, `justification`, `evidenceType`, `evidenceLocation`, `finding`, `ownerPhase`, and `route`. | automated/static | mandatory | Positive and negative fixture evidence. |
| NA-002 | Grouped non-applicability summaries use `controlDomain`, `corporateSection`, or another source-row grouping field. | static/automated | mandatory | Template/validator evidence showing source-row grouping fields only. |
| NA-003 | Grouping fails or splits when any grouped row lacks equivalent safe justification. | automated/static | mandatory | Negative fixture or rule inspection evidence. |
| NA-004 | Markdown grouped summaries cite row IDs/counts and safe rationale but do not override canonical row fields. | static/automated | mandatory | Markdown parity inspection. |
| NA-005 | Grouped non-applicability never uses legacy compact-control mapping as authority. | static | mandatory | Static inspection evidence across template/rules/schema. |

## Markdown Generation and Parity Tests

| ID | Check | Type | Severity | Expected Evidence |
| --- | --- | --- | --- | --- |
| MD-001 | Generated Markdown contains verdict, route, source refs, catalog snapshot identity, general review handoff, source-row navigation, summaries, blockers/warnings, unsafe evidence, exceptions, unavailable tooling, artifact metadata, and next recommendation from JSON. | automated/static | mandatory | Generator/template read-back evidence. |
| MD-002 | Full 155-row matrix appears as the final major content section. | automated/static | mandatory | Section-order check. |
| MD-003 | Markdown does not contain a compact-control validation section. | static/automated | mandatory | Template/generator inspection. |
| MD-004 | Markdown does not define independent validation rules; it only documents rendering/parity/read-back expectations. | static | mandatory | Template inspection. |
| MD-005 | JSON/Markdown parity failures route to blockers because JSON is authoritative. | automated/static | mandatory | Mismatch fixture or rule inspection. |

## Python Exporter Pytest Coverage

| ID | Check | Type | Severity | Expected Evidence | Unavailable Tooling Handling |
| --- | --- | --- | --- | --- | --- |
| PY-001 | `DEFAULT_TABLE_BY_SCHEMA["sdd-review-security.review-security-report"]` selects `sourceRowValidation.rows`. | automated/static | mandatory | Pytest assertion plus static code inspection. | If pytest dependencies are unavailable, record static inspection and unavailable pytest note. |
| PY-002 | Compact-only security report fails without `--table`. | automated | mandatory | Pytest failure-case assertion with safe CLI error. | If unavailable, record as pending mandatory automated evidence. |
| PY-003 | Manual nested table override remains generic and works for a list of objects. | automated | non-mandatory | Pytest success case. | If unavailable, document generic override inspection. |
| PY-004 | Workbook read-back through `openpyxl` validates sheet name, flattened fields, headers, and row count. | automated | mandatory | Pytest read-back assertion using temp output path. | If unavailable, record dependency/tooling gap. |
| PY-005 | README documents source-row default and does not describe legacy compact export as active. | static | mandatory | README inspection. | Static evidence is sufficient. |

Planned command when dependencies are available:

```text
python -m pytest python/tests
```

If Python, pytest, openpyxl, or local dependencies are unavailable, verification must report unavailable tooling explicitly and use static code/README inspection as substitute evidence. The unavailable command itself does not satisfy mandatory automated checks.

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Artifact authority | Confirm canonical JSON refs are treated as authority before derived Markdown and exporter output. | static/manual | Shared contract, verify/archive, and skill/agent inspection. | Global runner unavailable; static evidence required. |
| Exporter behavior | Run Python pytest when possible; otherwise inspect default table mapping, CLI error handling, temp workbook behavior, and README. | automated/static | Pytest result or unavailable-tooling note plus static inspection. | Python dependencies may be unavailable. |
| Contract propagation | Confirm skill, schema, rules, template, shared contracts, adapter prompts, specs, and README align. | static | File-by-file changed-surface checklist. | No lint/type/format tooling detected. |
| Safe evidence boundary | Confirm examples and errors use paths, section anchors, sanitized summaries, command outcomes, and redacted placeholders only. | static/manual | Evidence review notes; unsafe sample rejection evidence. | No runtime logs or sensitive payloads are needed. |
| Generated artifacts | Confirm workbook outputs are temporary test artifacts unless explicitly intended, and archive stores refs/summaries rather than generated bytes. | static/manual | Pytest tmp-path usage or static test inspection. | No artifact bytes should be required for SDD evidence. |

## Narrative Security Rule Coverage

| Rule Category | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- |
| Sensitive data and evidence handling | Yes | TD-009, TD-020, TD-021 verify sanitized paths/summaries/placeholders, unsafe evidence rejection, README safety language, and no raw secrets/PII/PAN/tokens/log payloads in report/export evidence. | covered | None planned. |
| Files and exported artifacts | Yes | TD-016 through TD-020 and PY-001 through PY-005 verify exporter defaults, compact-only failure, temp workbook read-back, README guidance, and no committed workbook bytes. | covered | None planned. |
| Artifact authority and access-control boundary | Yes | TD-001, TD-007, TD-011 through TD-014 verify canonical JSON authority, derived Markdown presentation-only behavior, downstream ref consumption, and no second matrix authority. | covered | None planned. |
| Sensitive logging and error evidence | Yes | TD-003 through TD-006, TD-009, TD-017 verify safe validation/exporter errors name paths/counts/IDs/fields without echoing unsafe row payloads. | covered | None planned. |
| Schema and input validation | Yes | TD-001 through TD-010 and JSON sample strategy verify exact-once coverage, required fields, safe non-applicability, allowed grouping fields, safe evidence, and blocker routing. | covered | None planned. |
| Exception and evidence policy | Yes | TD-012 through TD-014 and TD-021 verify warnings/exceptions/evidence refs are preserved and incomplete exceptions remain blocking. | covered | None planned. |

Authentication, sessions, runtime secrets, and production database access are not design-time mandatory categories for this repository-contract/tooling change. Review-security remains responsible for final source-row coverage decisions and any omitted-category validation in canonical JSON.

## Acceptance Evidence Matrix

| Source Requirement / Decision | Planned Evidence |
| --- | --- |
| Proposal: `review-security-report.json` validates exactly 155 rows, each once. | TD-001 through TD-006; JSON fixture strategy; catalog self-check TD-015. |
| Proposal: legacy compact controls removed from active report contract. | TD-007; MD-003; NA-005; exporter default failure TD-017. |
| Proposal: lean Markdown removes compact validation section and renders full matrix last. | TD-011; MD-001 through MD-005. |
| Proposal: exporter default/tests/docs use `sourceRowValidation.rows`. | TD-016 through TD-020; PY-001 through PY-005. |
| Spec: Security Review Artifact. | TD-001, TD-002, TD-007, TD-011, TD-012. |
| Spec: Security Matrix Validation. | TD-002 through TD-010; NA-001 through NA-005. |
| Spec: Exhaustive Source Row Security Review. | TD-011, TD-012; MD-001 through MD-005. |
| Spec: In-Repo Guideline Snapshot and Corporate Source Row Inventory. | TD-015; JSON fixture strategy. |
| Spec: Safe Source Row Evidence. | TD-008 through TD-010; NA-001 through NA-005; narrative security coverage. |
| Spec: Verify and Archive Review Consumption. | TD-013, TD-014; operational artifact authority checks. |
| Spec: Source Row Persistence Compatibility. | TD-001, TD-011 through TD-014; persistence contract inspection. |
| Spec: Schema-aware table selection and pytest verification. | TD-016 through TD-019; PY-001 through PY-004. |
| Design decision: Source rows are the only active security matrix. | TD-001 through TD-006; TD-015. |
| Design decision: legacy compact data is removed from active report authority. | TD-007; TD-017; MD-003; NA-005. |
| Design decision: Grouping uses source-row category fields. | TD-010; NA-002 through NA-005; MD-001. |
| Design decision: Markdown is generated presentation only. | TD-011, TD-012; MD-004, MD-005. |
| Secure development design: safe evidence and unavailable tooling. | TD-009, TD-021, TD-022; operational checks. |

## Evidence Expectations

- Mandatory cases require implementation, execution, static/manual evidence, or a justified unavailable-tooling note with substitute inspection evidence.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Source-row exact-once coverage, required row fields, safe non-applicability, safe evidence, and canonical JSON authority are verification-blocking when uncovered.
- Test-design consumes narrative design rules only and does not require design YAML, schema fields, compact control tables, Source ID matrices, machine-readable applicability fields, exhaustive non-applicability rows, or the full 155-row matrix in this artifact.
- Exhaustive source-row validation coverage, omitted-row validation, and row-level non-applicability decisions belong to canonical `review-security-report.json`; derived Markdown is compatibility output.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands that are unavailable must be reported as unavailable evidence, never as passing evidence.
- Python exporter tests are required when dependencies are available; if they are unavailable, static inspection and explicit unavailable-tooling notes must carry forward until verification can run them.

## No-Impact Assessment

Not applicable. This change directly affects security-review validation, artifact authority, generated Markdown, verify/archive contracts, and Excel exporter behavior.

## Open Questions

None.
