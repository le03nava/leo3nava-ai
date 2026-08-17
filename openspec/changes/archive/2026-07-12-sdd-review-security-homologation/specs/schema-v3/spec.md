# Delta for schema-v3

## Purpose

Delta spec for `review-security-report.schema.json`. Covers the full rewrite to the v3 flat-field pattern: 16 required root fields, `rows[]` at root, flat `sourceRow` $def, and removal of eliminated $defs and properties.

## ADDED Requirements

### Requirement: SCHEMA-001 — schemaVersion Is 3

The JSON Schema MUST declare `schemaVersion` as an integer field with `minimum: 3`.

#### Scenario: Valid report is validated against schema

- GIVEN a review-security-report.json with `schemaVersion: 3`
- WHEN the schema validator runs
- THEN the document MUST pass validation

#### Scenario: Report with schemaVersion below 3 is rejected

- GIVEN a report with `schemaVersion: 2`
- WHEN the schema validator runs
- THEN the document MUST fail validation

---

### Requirement: SCHEMA-002 — required[] Contains Exactly 16 Fields

The schema `required` array MUST contain exactly these 16 fields: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `generalReviewRef`, `rows`, `exceptions`, `artifactMetadata`.

#### Scenario: required array is validated for exact membership

- GIVEN the v3 schema file
- WHEN a validator reads the top-level `required` array
- THEN it MUST contain exactly those 16 field names and no others

---

### Requirement: SCHEMA-003 — Removed Properties Must Not Exist

The following properties MUST NOT exist in the schema: `sourceRefs`, `catalogRefs`, `generalReviewHandoff`, `sourceRowValidation`, `blockers`, `warnings`, `unsafeEvidenceRejections`, `warningCarryForward`.

#### Scenario: Schema is checked for removed property keys

- GIVEN the v3 schema file
- WHEN a validator inspects the top-level `properties` object
- THEN none of the 8 removed keys MUST be present

---

### Requirement: SCHEMA-004 — catalogRef Is a Constant

`catalogRef` MUST be defined as `const: "skills/sdd-review-security/references/security-guideline-catalog.operational.json"`.

#### Scenario: Report with correct catalogRef passes

- GIVEN a report with the exact const value for `catalogRef`
- WHEN the schema validates the report
- THEN validation MUST pass

#### Scenario: Report with different catalogRef fails

- GIVEN a report with a different `catalogRef` value
- WHEN the schema validates the report
- THEN validation MUST fail

---

### Requirement: SCHEMA-005 — generalReviewRef Is a Non-Empty String

`generalReviewRef` MUST be defined as `type: string` with `minLength: 1`.

#### Scenario: Report with generalReviewRef present passes

- GIVEN a report where `generalReviewRef` is a non-empty string
- WHEN the schema validates the report
- THEN validation MUST pass

#### Scenario: Empty generalReviewRef fails

- GIVEN a report where `generalReviewRef` is an empty string
- WHEN the schema validates the report
- THEN validation MUST fail

---

### Requirement: SCHEMA-006 — totals Has 8 Required Fields

The `totals` object MUST require exactly 8 fields: `sourceRowCount` (const 155), `validated`, `passing`, `failing`, `notApplicable`, `blockers`, `warnings`, `exceptions`.

#### Scenario: totals with all 8 fields passes validation

- GIVEN a report with a complete `totals` object
- WHEN the schema validates the report
- THEN validation MUST pass

#### Scenario: totals.sourceRowCount must be 155

- GIVEN a report with `totals.sourceRowCount` set to any value other than `155`
- WHEN the schema validates the report
- THEN validation MUST fail

---

### Requirement: SCHEMA-007 — rows[] Is at Root Level with 155 Items

`rows` MUST be defined at the report root as an array with `minItems: 155` and `maxItems: 155`.

#### Scenario: Report with 155 rows passes

- GIVEN a report with exactly 155 entries in `rows[]`
- WHEN the schema validates the report
- THEN validation MUST pass

#### Scenario: Report with fewer than 155 rows fails

- GIVEN a report with fewer than 155 rows
- WHEN the schema validates the report
- THEN validation MUST fail

---

### Requirement: SCHEMA-008 — sourceRow $def Excludes Catalog Fields

The `sourceRow` $def MUST NOT contain: `corporateSection`, `pciAlignment`, `guidelineText`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`.

#### Scenario: sourceRow $def is checked for excluded fields

- GIVEN the v3 schema file
- WHEN a validator inspects `$defs.sourceRow.properties`
- THEN none of the 8 catalog-only fields MUST be present

---

### Requirement: SCHEMA-009 — sourceRow $def Contains Exactly 10 Fields

The `sourceRow` $def MUST contain exactly: `sourceId`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, `route`.

#### Scenario: sourceRow $def has correct property set

- GIVEN the v3 schema file
- WHEN a validator reads `$defs.sourceRow.properties` keys
- THEN exactly those 10 keys MUST be present and no others

---

### Requirement: SCHEMA-010 — exception $def Does Not Contain status

The `exception` $def MUST NOT contain a `status` field.

#### Scenario: exception $def is checked for status field

- GIVEN the v3 schema file
- WHEN a validator inspects `$defs.exception.properties`
- THEN no `status` key MUST be present

---

### Requirement: SCHEMA-011 — exception $def required Contains 6 Fields

The `exception` $def `required` array MUST be exactly: `sourceId`, `approver`, `approvedAt`, `acceptedRiskRationale`, `mitigationOrFollowUp`, `evidenceGap`.

#### Scenario: exception required array is validated

- GIVEN the v3 schema file
- WHEN a validator reads `$defs.exception.required`
- THEN exactly those 6 fields MUST be listed and no others

---

### Requirement: SCHEMA-012 — Removed $defs Must Not Exist

`$defs` MUST NOT contain: `navigationGroup`, `groupedNaSummary`, `finding`.

#### Scenario: $defs is checked for removed definitions

- GIVEN the v3 schema file
- WHEN a validator reads the `$defs` keys
- THEN none of the 3 removed definition names MUST be present

---

### Requirement: SCHEMA-013 — $comment References rows[]

The schema `$comment` MUST reference `rows[]` and MUST NOT reference `sourceRowValidation.rows`.

#### Scenario: $comment content is validated

- GIVEN the v3 schema file
- WHEN a validator reads the top-level `$comment`
- THEN the text MUST reference `rows[]`
- AND MUST NOT contain the string `sourceRowValidation`

---

### Requirement: SCHEMA-014 — artifactMetadata Shape Matches review-report.schema.json

The `artifactMetadata` object shape MUST be identical to the one in `review-report.schema.json`, with the same 9 required fields.

#### Scenario: artifactMetadata required fields match the reference schema

- GIVEN the v3 schema file and the reference `review-report.schema.json`
- WHEN a validator compares `artifactMetadata.required` in both files
- THEN both arrays MUST contain exactly the same 9 field names
