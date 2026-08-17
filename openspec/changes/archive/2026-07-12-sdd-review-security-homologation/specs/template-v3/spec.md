# Delta for template-v3

## Purpose

Delta spec for `report-template.md`. Covers the rewrite to the v3 9-section structure: updated field mapping table, new skeleton, correct join instruction for the matrix, and removal of obsolete sections.

## ADDED Requirements

### Requirement: TMPL-001 — Preamble Preserved

The first 3 paragraphs describing JSON authority MUST be preserved unchanged from the pre-v3 template.

#### Scenario: Preamble content is verified

- GIVEN the v3 template file
- WHEN a reviewer reads the opening preamble
- THEN the 3 JSON-authority paragraphs MUST be present and their content MUST match the original

---

### Requirement: TMPL-002 — JSON Field Mapping Table Has Exactly 9 Sections

The `## JSON Field Mapping` table MUST map exactly 9 sections: Verdict, Totals, General Review Reference, Unavailable Tooling, Exceptions, Artifact Metadata, Recommendation, Matrix, and one additional section if present in the v3 structure.

#### Scenario: Mapping table section count is validated

- GIVEN the v3 template's `## JSON Field Mapping` table
- WHEN a validator counts the mapped sections
- THEN exactly 9 sections MUST be listed

---

### Requirement: TMPL-003 — Required Structure Skeleton Shows 8 Sections

The `## Required Structure` section MUST show a skeleton with exactly these sections: Verdict, Totals, General Review Reference, Unavailable Tooling, Exceptions, Artifact Metadata, Recommendation, Matrix.

#### Scenario: Skeleton sections are validated

- GIVEN the `## Required Structure` section of the v3 template
- WHEN a validator lists all section headings in the skeleton
- THEN exactly those 8 section names MUST appear and no others

---

### Requirement: TMPL-004 — Verdict Table Includes Status and JSON Authority Rows

The `## Verdict` table MUST include both a Status row and a JSON authority row.

#### Scenario: Verdict table rows are validated

- GIVEN the `## Verdict` section of the v3 template
- WHEN a validator reads the table rows
- THEN a Status row MUST be present
- AND a JSON authority row MUST be present

---

### Requirement: TMPL-005 — Totals Table Shows 8 Fields Including sourceRowCount 155

The `## Totals` table MUST show all 8 totals fields, including one labeled "Total source rows (155)".

#### Scenario: Totals table completeness is validated

- GIVEN the `## Totals` section of the v3 template
- WHEN a validator counts the field rows in the table
- THEN exactly 8 rows MUST be present
- AND one row MUST reference "Total source rows (155)"

---

### Requirement: TMPL-006 — General Review Reference Is a Simple Ref Section

The `## General Review Reference` section MUST be a simple reference field, not a complex handoff object.

#### Scenario: Section is checked for simple structure

- GIVEN the `## General Review Reference` section of the v3 template
- WHEN a validator inspects the section content
- THEN the section MUST reference a single string field (`generalReviewRef`)
- AND MUST NOT describe nested object properties or a handoff structure

---

### Requirement: TMPL-007 — Matrix Section Shows Join Instruction

The `## Matrix` section MUST include an explicit join instruction: join `rows[].sourceId` with catalog `sourceRows[].sourceId`.

#### Scenario: Join instruction is present and correct

- GIVEN the `## Matrix` section of the v3 template
- WHEN a validator reads the section preamble
- THEN a join instruction referencing `rows[].sourceId` and `sourceRows[].sourceId` MUST be present

---

### Requirement: TMPL-008 — Matrix Table Has Exactly 15 Columns

The matrix table MUST have exactly 15 columns: Source ID, Corporate Section, Control Domain, PCI Alignment, Guideline, Applies When, Applies, Complies, Lifecycle Status, Evidence Type, Evidence Location, Justification, Finding, Owner Phase, Route.

#### Scenario: Matrix column count and names are validated

- GIVEN the matrix table header in the v3 template
- WHEN a validator reads the column headers
- THEN exactly those 15 column names MUST be present in that order

---

### Requirement: TMPL-009 — Matrix Rules Section Is Present

The `## Matrix Rules` section MUST be present in the template.

#### Scenario: Section presence is confirmed

- GIVEN the v3 template
- WHEN a validator scans section headings
- THEN `## Matrix Rules` MUST be present

---

### Requirement: TMPL-010 — Safe Evidence Rules Section Is Present

The `## Safe Evidence Rules` section MUST be present in the template.

#### Scenario: Section presence is confirmed

- GIVEN the v3 template
- WHEN a validator scans section headings
- THEN `## Safe Evidence Rules` MUST be present

---

### Requirement: TMPL-011 — Obsolete Sections Must Not Exist

The following sections MUST NOT exist in the v3 template: `## Source References`, `## General Review Handoff`, `## Source Row Navigation`, `## Source Row Summary`, `## Grouped Non-Applicability`, `## Blockers and Warnings`.

#### Scenario: Template is scanned for each obsolete section

- GIVEN the v3 template file
- WHEN a validator checks all second-level headings
- THEN none of the 6 obsolete section names MUST be found
