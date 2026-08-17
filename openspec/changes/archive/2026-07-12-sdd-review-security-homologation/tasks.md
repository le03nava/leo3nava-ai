# Tasks: SDD Review Security Homologation

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | ~350–450 (4 static reference files; no runtime code) |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | null |
| Chain strategy | pending |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: pending
Size exception: none
Review budget lines: 400
Review budget risk: Medium
400-line budget risk: Medium

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | All 4 reference file rewrites + static verification | PR 1 | All files are independent; single coherent reviewable unit |

---

## Phase 1: Catalog Rewrite — `security-guideline-catalog.operational.json`

- [x] 1.1 Read `skills/sdd-review-security/references/security-guideline-catalog.operational.json` and confirm current `schemaVersion`, header keys (`generatedHumanViewRef`, `reporting`, `authority`), and per-row fields (`defaultOwnerPhase`, `defaultRoute`, `guidelineRefs`, `evidenceExpectation`).
- [x] 1.2 Rewrite catalog file: set `schemaVersion: 3`; remove top-level keys `generatedHumanViewRef`, `reporting`, `authority`; on every sourceRow rename `defaultOwnerPhase` → `ownerPhase`, `defaultRoute` → `route`; remove `guidelineRefs` and `evidenceExpectation` from every row. Preserve all other fields on all 155 rows.
- [x] 1.3 Verify TC-C-001: `jq '.schemaVersion'` equals `3`.
- [x] 1.4 Verify TC-C-002: confirm pre-v3 detection (`jq '.schemaVersion < 3'` on stale copy returns `true`); document by inspection.
- [x] 1.5 Verify TC-C-003/004/005: `jq 'has("generatedHumanViewRef")'`, `has("reporting")`, `has("authority")` all equal `false`.
- [x] 1.6 Verify TC-C-006/007: `ownerPhase` present on all 155 rows (count = 0 nulls); `defaultOwnerPhase` count = 0.
- [x] 1.7 Verify TC-C-008/009: `route` present on all 155 rows (count = 0 nulls); `defaultRoute` count = 0.
- [x] 1.8 Verify TC-C-010/011: `guidelineRefs` count = 0; `evidenceExpectation` count = 0.
- [x] 1.9 Verify TC-C-012: `jq '.sourceRows | length'` equals `155`; empty `guidelineText` count = `0`.

---

## Phase 2: Schema Rewrite — `review-security-report.schema.json`

- [x] 2.1 Read `skills/sdd-review-security/references/review-security-report.schema.json` and identify all fields to remove (`sourceRefs`, `catalogRefs`, `generalReviewHandoff`, `sourceRowValidation`, `blockers`, `warnings`, `unsafeEvidenceRejections`, `warningCarryForward`) and `$defs` to remove (`navigationGroup`, `groupedNaSummary`, `finding`).
- [x] 2.2 Rewrite schema file: new `required[]` with exactly 16 fields; add `totals`, `catalogRef` (const path), `catalogSnapshotId`, `generalReviewRef`, `rows[]` (root-level, min/maxItems 155); update `$comment` to reference `rows[]` not `sourceRowValidation`; `additionalProperties: false` on all objects.
- [x] 2.3 Rewrite `$defs`: `totals` (8 required, `sourceRowCount` const 155); `sourceRow` (10 fields only: sourceId, applies, complies, lifecycleStatus, evidenceType, evidenceLocation, justification, finding, ownerPhase, route); `exception` (6 required, no `status`); retain `stateArtifactRef` and `artifactMetadata` matching `review-report.schema.json`.
- [x] 2.4 Verify TC-S-001: `jq '.properties.schemaVersion.minimum'` equals `3`.
- [x] 2.5 Verify TC-S-002: `jq '.required | length'` equals `16`; sorted list matches spec.
- [x] 2.6 Verify TC-S-003: 8 removed property keys absent (count = 0).
- [x] 2.7 Verify TC-S-004/005: `catalogRef.const` equals exact path string; `generalReviewRef` type `string` minLength `1`.
- [x] 2.8 Verify TC-S-006/007: `totals` required length = 8; `sourceRowCount` const = 155; `rows` minItems/maxItems = 155.
- [x] 2.9 Verify TC-S-008/009: 8 catalog-only fields absent from `sourceRow $def`; `sourceRow` property count = 10.
- [x] 2.10 Verify TC-S-010/011: `exception` has no `status`; `exception.required` length = 6 with exact fields.
- [x] 2.11 Verify TC-S-012: `navigationGroup`, `groupedNaSummary`, `finding` absent from `$defs`.
- [x] 2.12 Verify TC-S-013: `$comment` contains `"rows[]"`; does not contain `"sourceRowValidation"`.
- [x] 2.13 Verify TC-S-014: `artifactMetadata.required` sorted matches `review-report.schema.json` (9 fields); manual-read if `jq` cross-file comparison unavailable.

---

## Phase 3: Template Partial Replacement — `report-template.md`

- [x] 3.1 Read `skills/sdd-review-security/references/report-template.md`; identify the 3-paragraph preamble and the boundary at `## JSON Field Mapping`.
- [x] 3.2 Replace from `## JSON Field Mapping` onward: write 9-row mapping table, 8-section `## Required Structure` skeleton (Verdict, Totals, General Review Reference, Unavailable Tooling, Exceptions, Artifact Metadata, Recommendation, Matrix), join instruction (`rows[].sourceId` with `sourceRows[].sourceId`), 15-column matrix table, `## Matrix Rules` section, `## Safe Evidence Rules` section. Remove 6 obsolete sections (`## Source References`, `## General Review Handoff`, `## Source Row Navigation`, `## Source Row Summary`, `## Grouped Non-Applicability`, `## Blockers and Warnings`).
- [x] 3.3 Verify TC-T-001: preamble (3 paragraphs before first `##`) present and contains JSON authority language; content not truncated.
- [x] 3.4 Verify TC-T-002/003: `## JSON Field Mapping` table has 9 data rows; `## Required Structure` lists exactly 8 named sections.
- [x] 3.5 Verify TC-T-004/005: Verdict section has Status and JSON authority rows; Totals table has 8 rows including `Total source rows (155)`.
- [x] 3.6 Verify TC-T-006: `## General Review Reference` contains `generalReviewRef`; does not contain `handoff`.
- [x] 3.7 Verify TC-T-007/008: Matrix section contains join instruction (`rows[].sourceId` and `sourceRows[].sourceId`); matrix table header has exactly 15 columns in spec order.
- [x] 3.8 Verify TC-T-009/010: `## Matrix Rules` heading present; `## Safe Evidence Rules` heading present.
- [x] 3.9 Verify TC-T-011: all 6 obsolete headings absent (grep each, no match).

---

## Phase 4: Validation Rules Targeted Edits — `validation-rules.md`

- [x] 4.1 Read `skills/sdd-review-security/references/validation-rules.md`; locate Active Flow Boundary section, Source Row Coverage Validation table, Required Field and Vocabulary Validation section, Row-Preserving Non-Applicability section, and Report Validation Checklist.
- [x] 4.2 Replace all `sourceRowValidation.rows` references with `rows[]` throughout the file.
- [x] 4.3 Update Source Row Coverage Validation table: `Exact report table` row → `rows[]`; Count consistency row → `totals.sourceRowCount` and `totals.validated`; Required row fields row → 10 v3 fields only. Remove `Source-row grouping` row.
- [x] 4.4 Remove `guidelineText`/`guidelineRefs` blocking rule line from Required Field and Vocabulary Validation section.
- [x] 4.5 Remove entire `Row-Preserving Non-Applicability` section.
- [x] 4.6 Update Report Validation Checklist: 13 v3 entries for `review-security-report.json`; replace `sourceRowValidation.rows` item with `rows[] contains exactly 155 rows`; remove items for `groupedNaSummaries`, `navigationSummary`.
- [x] 4.7 Verify TC-V-001/002: `sourceRowValidation.rows` absent everywhere; `rows[]` present in Active Flow Boundary and Coverage table.
- [x] 4.8 Verify TC-V-003: `totals.sourceRowCount` and `totals.validated` present in count-consistency row; 4 forbidden field names absent.
- [x] 4.9 Verify TC-V-004: required row fields list has exactly 10 entries; `guidelineText` and `corporateSection` absent.
- [x] 4.10 Verify TC-V-005/006/007: `Source-row grouping` row absent; `guidelineText`/`guidelineRefs` blocking rule absent; `Row-Preserving Non-Applicability` section heading absent.
- [x] 4.11 Verify TC-V-008/009/010: Checklist has 13 entries; `sourceRowValidation.rows` item absent, `rows[] contains exactly 155 rows` item present; `groupedNaSummaries`, `navigationSummary`, `grouped N/A` absent from checklist.

---

## Apply Progress Tracking

> Update each task checkbox (`- [x]`) as work completes. All 47 test cases (TC-C-001..012, TC-S-001..014, TC-T-001..011, TC-V-001..010) must be `applied` in `test-cases.json` before `sdd-verify` runs. All verification is static inspection — record actual `jq`/`grep` output or manual-read observation per case.
