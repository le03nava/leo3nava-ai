# Delta for sdd-review-workflow

## MODIFIED Requirements

### Requirement: Code-Review Validation Matrix

The canonical JSON MUST define the review matrix source rows and validation rules. Derived `review-report.md` MUST preserve the exact matrix columns: Item, Requirement, Standard, Category, Severity, Complies, Finding, Evidence Location, Notes. The derived matrix MUST contain 96 rows, use stable `REV-CORP-*` IDs, and limit `Complies` to `Yes`, `No`, or `N/A`. `N/A` MUST include Evidence Location proving irrelevance and Notes explaining scope. Each matrix row MUST include a `finding` free-text field in the canonical JSON that maps to the `Finding` column in derived Markdown. `finding` MUST be a free-text string and MUST NOT be a controlled enum.

(Previously: matrix columns were `Item | Artifact/Deliverable | Requirement | Reviewer | Standard | Severity | Complies | Affected Requirement | Evidence Location | Observations/Comments`; no `Finding` column existed; `finding` was not a required matrix row field)

#### Scenario: All controls are represented

- GIVEN canonical JSON contains 96 controls
- WHEN review writes JSON and derived Markdown
- THEN every control MUST appear once with a stable Item ID
- AND every Markdown row MUST use the exact required v2 columns.

#### Scenario: Platform control is irrelevant

- GIVEN a control targets an unused platform or technology
- WHEN review marks it `N/A` in JSON
- THEN derived Markdown Evidence Location MUST prove irrelevance
- AND Notes MUST explain the scope decision.

#### Scenario: Finding column is populated

- GIVEN a control receives a review result
- WHEN review records the matrix row in JSON
- THEN `matrix[].finding` MUST be present as a free-text string
- AND derived Markdown MUST render it in the `Finding` column.

#### Scenario: Derived Markdown preserves compatibility sections

- GIVEN canonical JSON has report facts
- WHEN Markdown is generated
- THEN it MUST preserve verdict, evidence summary, review matrix, and recommendation sections
- AND it MUST NOT render `## Blocking Summary`, `## Operational Evidence Summary`, `## Changed-File / Security Handoff`, `## Matrix Validation`, or `## Catalog Validation` sections.

### Requirement: Canonical Review JSON Authority

`sdd-review` MUST treat `skills/sdd-review/references/review-control-catalog.json` (v2) as the authoritative static source for all 96 general review controls. The catalog MUST carry `schemaVersion: 2`, a root `reviewer` field, and MUST NOT include `generatedHumanViewRef`, `presentation`, `validation`, `safeEvidence`, or `forbiddenEvidence` root fields. Each control entry MUST NOT carry a per-entry `reviewer` field, and MUST NOT carry a `notes` field when its value is exactly `"Not provided"`. The `vocabulary` object MUST retain only `complies` and `severity` arrays. `review-report.json` remains the authoritative per-change source for review facts.

(Previously: catalog was schemaVersion 1, had root `generatedHumanViewRef`, `presentation`, `validation`, and `vocabulary.safeEvidence`/`vocabulary.forbiddenEvidence` fields; each control carried a `reviewer` field and a `notes` field regardless of value)

#### Scenario: Catalog loads at v2

- GIVEN sdd-review resolves the canonical catalog
- WHEN it reads `review-control-catalog.json`
- THEN `schemaVersion` MUST be `2`
- AND root `reviewer`, `generatedHumanViewRef`, `presentation`, `validation`, `safeEvidence`, `forbiddenEvidence` MUST NOT be present outside `vocabulary`.

#### Scenario: Per-control reviewer field is absent

- GIVEN a catalog control entry is read
- WHEN any of the 96 entries is inspected
- THEN the entry MUST NOT contain a `reviewer` field.

#### Scenario: Blank notes are absent

- GIVEN a catalog control entry has no meaningful notes
- WHEN the entry is inspected
- THEN the entry MUST NOT contain a `notes` field whose value is `"Not provided"`.

#### Scenario: Vocabulary is minimal

- GIVEN the catalog `vocabulary` object is read
- WHEN its keys are enumerated
- THEN it MUST contain `complies` and `severity`
- AND it MUST NOT contain `safeEvidence` or `forbiddenEvidence`.

#### Scenario: JSON wins over Markdown

- GIVEN canonical JSON and derived Markdown disagree
- WHEN the workflow decides review facts, controls, counts, routing, or validation status
- THEN the JSON MUST win
- AND Markdown MUST be treated as the repair target.

### Requirement: Review Report Artifact

`sdd-review` MUST persist canonical `review-report.json` conforming to `review-report.schema.json` v2 and generate derived `review-report.md` in the same phase. The report MUST include the 14 required root fields: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `matrix`, `artifactMetadata`. The report MUST NOT include legacy fields `blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, or `operationalEvidenceSummary`. `artifactMetadata.stateRegistration` MUST contain at least 2 entries: ① canonical JSON review-report ref, ② derived Markdown review-report ref. All objects in the schema MUST enforce `additionalProperties: false`. Missing required artifacts, JSON validation failure, Markdown generation/parity failure, stale Markdown, or persistence failure MUST route to `resolve-blockers`.

(Previously: report was schemaVersion 1 with 16 required root fields including `blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, `operationalEvidenceSummary`, `reviewMatrix`, and `presentation`; no `artifactMetadata.stateRegistration` constraint; no `finding` column; `additionalProperties: false` only partially enforced)

#### Scenario: Reports are persisted

- GIVEN review can resolve all required inputs
- WHEN review completes
- THEN it MUST write canonical `review-report.json` with all 14 required root fields
- AND it MUST write derived `review-report.md` with verdict, matrix, evidence summary, and next recommendation.

#### Scenario: Legacy fields are rejected

- GIVEN a report document includes `blockingSummary`, `runtimeChecks`, or any other removed field
- WHEN it is validated against `review-report.schema.json` v2
- THEN validation MUST fail
- AND the phase MUST route to `resolve-blockers`.

#### Scenario: stateRegistration has two entries

- GIVEN review persists the canonical report
- WHEN `artifactMetadata.stateRegistration` is inspected
- THEN it MUST contain at minimum: the canonical JSON ref and the derived Markdown ref
- AND the array MUST have `minItems: 2`.

#### Scenario: Markdown generation fails

- GIVEN canonical JSON is valid
- WHEN derived Markdown cannot be generated or read back
- THEN review MUST route to `resolve-blockers`
- AND downstream phases MUST NOT consume stale Markdown as current evidence.

## REMOVED Requirements

### Requirement: Operational Readiness General Review

(Reason: The v2 schema removes `operationalEvidenceSummary` from the report contract. Operational evidence coverage is no longer a required field in the canonical JSON, and the corresponding Markdown section is removed from the template.)
(Migration: Operational evidence MAY be documented in the `matrix[].finding` free-text field or in evidence location fields. No separate section or JSON field is required.)

### Requirement: Operational Review Handoff

(Reason: The `operationalEvidenceSummary` handoff field is removed from the v2 schema. The corresponding template section `## Operational Evidence Summary` is removed.)
(Migration: Any applicable operational context SHOULD be recorded in matrix finding or evidence location fields.)

## ADDED Requirements

### Requirement: Review Schema v2 Reference Contract

The three reference artifacts in `skills/sdd-review/references/` MUST collectively define the v2 review contract: `review-control-catalog.json` (catalog), `review-report.schema.json` (schema), and `report-template.md` (Markdown presentation template). All three MUST be updated atomically as part of `sdd-review-schema-v2`. `sdd-review-security/` MUST NOT be modified in this change; a follow-on change is required before running `sdd-review-security` against a v2 report.

#### Scenario: Reference artifacts are coherent

- GIVEN the three reference files have been updated
- WHEN sdd-review runs against a new change
- THEN it MUST load catalog v2, validate the report against schema v2, and render Markdown using the updated template
- AND the three artifacts MUST be consistent with each other.

#### Scenario: sdd-review-security is deferred

- GIVEN `sdd-review-schema-v2` is applied
- WHEN `sdd-review-security` runs against a v2 report
- THEN the phase MUST fail or route to `resolve-blockers` until the follow-on update is applied
- AND `skills/sdd-review-security/` MUST remain unchanged by this change.

### Requirement: Report Template v2 Structure

The `report-template.md` MUST reflect the v2 field contract from `## JSON Field Mapping` onward. The field-mapping table MUST map only the 14 v2 required fields. The required structure skeleton MUST show a full 96-row matrix spec with columns `Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes`. The sections `## Matrix Validation` and `## Catalog Validation` MUST be removed. The `## Safe Evidence Rules` section MUST be preserved verbatim (current lines 108–111 of `report-template.md`). Content before `## JSON Field Mapping` MUST remain unchanged.

#### Scenario: Field mapping table is updated

- GIVEN the template is read
- WHEN the `## JSON Field Mapping` section is parsed
- THEN it MUST list only v2 required root fields
- AND it MUST NOT reference `blockingSummary`, `inputsInspected`, `runtimeChecks`, `evidenceSummary`, `operationalEvidenceSummary`, `changedFileSecurityHandoff`, `validation`, or `presentation`.

#### Scenario: Validation sections are absent

- GIVEN the updated template is read
- WHEN its sections are enumerated
- THEN `## Matrix Validation` and `## Catalog Validation` MUST NOT be present.

#### Scenario: Safe Evidence Rules are verbatim

- GIVEN the current `## Safe Evidence Rules` block (lines 108–111)
- WHEN the updated template is read
- THEN that block MUST appear verbatim with no modifications.

#### Scenario: Matrix columns include Finding

- GIVEN the required structure skeleton shows a matrix header
- WHEN its columns are parsed
- THEN the column sequence MUST be: Item, Requirement, Standard, Category, Severity, Complies, Finding, Evidence Location, Notes.

## Non-Requirements (Explicitly Out of Scope)

- `skills/sdd-review/SKILL.md` — no behavior changes.
- `skills/sdd-review-security/` — any update is deferred to a follow-on change.
- Existing archived `review-report.json` files — no migration of reports already archived.
- `openspec/specs/sdd-review-workflow/spec.md` and `openspec/specs/sdd-review-security-workflow/spec.md` — workflow behavior specs are unchanged.
- Excel, Python, script, spreadsheet, or workbook generation — deferred; not in scope.
