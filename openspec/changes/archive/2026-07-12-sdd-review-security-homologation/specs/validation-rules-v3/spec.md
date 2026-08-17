# Delta for validation-rules-v3

## Purpose

Delta spec for `validation-rules.md`. Covers path corrections from `sourceRowValidation.rows` to `rows[]`, updated required-row fields (catalog fields removed), removal of the Row-Preserving Non-Applicability section and obsolete rules, and checklist updates.

## ADDED Requirements

### Requirement: VALRULES-001 — Active Flow Boundary References rows[]

The Active Flow Boundary section MUST reference `rows[]` as the canonical report table path and MUST NOT reference `sourceRowValidation.rows`.

#### Scenario: Active flow boundary path is validated

- GIVEN the `validation-rules.md` file
- WHEN a validator reads the Active Flow Boundary section
- THEN all row-path references MUST use `rows[]`
- AND no reference to `sourceRowValidation.rows` MUST be present

---

### Requirement: VALRULES-002 — Source Row Coverage Table Uses rows[]

The Source Row Coverage Validation table's "Exact report table" row MUST reference `rows[]` and MUST NOT reference `sourceRowValidation.rows`.

#### Scenario: Coverage table row is validated

- GIVEN the Source Row Coverage Validation table in `validation-rules.md`
- WHEN a validator reads the "Exact report table" row
- THEN the cell content MUST reference `rows[]`
- AND MUST NOT contain the string `sourceRowValidation.rows`

---

### Requirement: VALRULES-003 — Count Consistency Row Uses totals Fields

The "Count consistency" row in the Source Row Coverage Validation table MUST reference `totals.sourceRowCount` and `totals.validated` and MUST NOT reference `sourceRowValidation.expectedCount`, `sourceRowValidation.validatedCount`, `coverageStatus`, or `exactOnce`.

#### Scenario: Count consistency cell references are validated

- GIVEN the Source Row Coverage Validation table
- WHEN a validator reads the "Count consistency" row
- THEN the cell MUST reference `totals.sourceRowCount` and `totals.validated`
- AND MUST NOT contain any of the 4 forbidden field names

---

### Requirement: VALRULES-004 — Required Row Fields List Excludes Catalog Fields

The "Required row fields" list MUST include exactly: `sourceId`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, `route`. Catalog-origin fields MUST NOT be listed.

#### Scenario: Required row fields list is validated

- GIVEN the required row fields section of `validation-rules.md`
- WHEN a validator reads the listed fields
- THEN exactly those 10 field names MUST be present
- AND no catalog-origin fields (e.g. `guidelineText`, `corporateSection`) MUST appear

---

### Requirement: VALRULES-005 — Source-Row Grouping Row Removed from Coverage Table

The "Source-row grouping" row MUST NOT exist in the Source Row Coverage Validation table.

#### Scenario: Coverage table is checked for removed row

- GIVEN the Source Row Coverage Validation table
- WHEN a validator scans all row labels
- THEN no row labeled "Source-row grouping" or equivalent MUST be present

---

### Requirement: VALRULES-006 — guidelineText/guidelineRefs Blocking Rule Removed

The line "A row with neither guidelineText nor guidelineRefs blocks" MUST NOT exist in `validation-rules.md`.

#### Scenario: File is searched for the removed rule

- GIVEN the v3 `validation-rules.md`
- WHEN a validator searches the full text for "guidelineText" or "guidelineRefs" in a blocking-rule context
- THEN no such line MUST be found

---

### Requirement: VALRULES-007 — Row-Preserving Non-Applicability Section Removed

The section "Row-Preserving Non-Applicability" MUST NOT exist in `validation-rules.md`.

#### Scenario: File is scanned for the removed section heading

- GIVEN the v3 `validation-rules.md`
- WHEN a validator scans all section headings
- THEN no heading matching "Row-Preserving Non-Applicability" MUST be found

---

### Requirement: VALRULES-008 — Report Validation Checklist Lists 12 v3 Fields

The Report Validation Checklist item about `review-security-report.json` MUST list exactly these fields: `schemaName`, `changeName`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `generalReviewRef`, `rows[]` (exactly 155), `exceptions[]`, `artifactMetadata`.

#### Scenario: Checklist item field list is validated

- GIVEN the Report Validation Checklist in `validation-rules.md`
- WHEN a validator reads the field list for the `review-security-report.json` item
- THEN exactly those 13 entries MUST be listed and no others

---

### Requirement: VALRULES-009 — Checklist Item Updated from sourceRowValidation.rows to rows[]

The checklist item previously reading "sourceRowValidation.rows contains exactly 155 rows" MUST be replaced with "rows[] contains exactly 155 rows".

#### Scenario: Checklist is checked for old and new wording

- GIVEN the v3 `validation-rules.md`
- WHEN a validator searches the checklist for "sourceRowValidation.rows"
- THEN no such item MUST be found
- AND an item referencing "rows[] contains exactly 155 rows" MUST be present

---

### Requirement: VALRULES-010 — Checklist Items for Removed Fields Must Not Exist

Checklist items referencing `sourceRowValidation.groupedNaSummaries`, `sourceRowValidation.navigationSummary`, or grouped N/A summaries MUST NOT exist in `validation-rules.md`.

#### Scenario: Checklist is scanned for each removed item

- GIVEN the v3 `validation-rules.md`
- WHEN a validator scans the full checklist
- THEN no item referencing `groupedNaSummaries`, `navigationSummary`, or "grouped N/A" MUST be present
