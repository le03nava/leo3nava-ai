# Tasks: sdd-review Schema v2 — Reference Artifact Cleanup

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 280–380 (catalog header ~15 removals + ~96×2 field drops ~192 lines; schema rewrite ~180 net; template partial rewrite ~80 net) |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | No |
| Suggested split | Single PR — three files, logically atomic |
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
| 1 | All three reference file edits + static verification evidence | PR 1 | Single atomic commit; catalog + schema + template change together |

---

## Phase 1: Catalog Update — review-control-catalog.json

- [x] 1.1 In `skills/sdd-review/references/review-control-catalog.json`: set `"schemaVersion": 2` (covers TC-001)
- [x] 1.2 Add root field `"reviewer": "sdd-review"` immediately after `expectedControlCount` (covers TC-002)
- [x] 1.3 Remove root field `"generatedHumanViewRef"` (covers TC-003)
- [x] 1.4 Remove root field `"presentation"` object (covers TC-004)
- [x] 1.5 Remove root field `"validation"` object (covers TC-005)
- [x] 1.6 Remove `vocabulary.safeEvidence`, `vocabulary.forbiddenEvidence`, and `vocabulary.reviewer` arrays from the `vocabulary` object; retain only `"complies"` and `"severity"` keys (covers TC-006)
- [x] 1.7 Iterate all 96 control entries: remove the `"reviewer"` field from every entry (covers TC-007)
- [x] 1.8 Iterate all 96 control entries: remove the `"notes"` field from every entry where `notes === "Not provided"`; retain meaningful notes (covers TC-008)
- [x] 1.9 Confirm `controls` array still has exactly 96 entries after edits (covers TC-009)

## Phase 2: Schema Rewrite — review-report.schema.json

- [x] 2.1 Full rewrite of `skills/sdd-review/references/review-report.schema.json`: write root object with `$schema`, `$id`, `title`, `description`, `$comment` (safe evidence rules), `type: object`, `additionalProperties: false`
- [x] 2.2 Write `required` array in exact order: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `matrix`, `artifactMetadata` (14 fields; covers TC-010)
- [x] 2.3 Write all 14 root property definitions per design.md specs (consts, enums, types, minLength constraints)
- [x] 2.4 Ensure NO legacy field properties exist: `blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, `operationalEvidenceSummary` (covers TC-011)
- [x] 2.5 Write `$defs.ref` object: required `artifact`, `ref`, `readable`; optional `notes`; `additionalProperties: false` (covers TC-012)
- [x] 2.6 Write `$defs.stateArtifactRef` object: required `artifact`, `ref`, `kind` (enum `canonical-json`/`derived-markdown`); `additionalProperties: false` (covers TC-012)
- [x] 2.7 Write `$defs.matrixRow` with required fields including `finding` as `type: string` (not enum); `item` pattern `^REV-CORP-(00[1-9]|0[1-8][0-9]|09[0-6])$`; `additionalProperties: false` (covers TC-013, TC-014)
- [x] 2.8 Write `totals` property as object with `blockingFailureCount` and `nonBlockingFindingCount` (integers ≥ 0); `additionalProperties: false` (covers TC-017)
- [x] 2.9 Write `artifactMetadata` property with `stateRegistration` array `minItems: 2`, items `$ref: #/$defs/stateArtifactRef`; `additionalProperties: false` (covers TC-015, TC-017)
- [x] 2.10 Confirm `additionalProperties: false` at schema root and all nested object definitions (`totals`, `matrixRow`, `stateArtifactRef`, `ref`, `artifactMetadata`) (covers TC-016, TC-017)

## Phase 3: Template Partial Rewrite — report-template.md

- [x] 3.1 In `skills/sdd-review/references/report-template.md`: keep lines 1–8 (preamble) completely unchanged (covers TC-025)
- [x] 3.2 Replace content from `## JSON Field Mapping` onward: write new `## JSON Field Mapping` table mapping all 14 v2 fields per design.md; no legacy field names in table (covers TC-018)
- [x] 3.3 Add `## Required Structure` section with full derived Markdown skeleton including 96-row matrix spec (covers TC-019)
- [x] 3.4 Matrix table header MUST be exactly: `Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes` (9 columns) (covers TC-020)
- [x] 3.5 Include `## Matrix Rules` section with matrix join instruction: `matrix[].item` joined with `catalog controls[].id` (covers TC-021)
- [x] 3.6 Ensure `## Matrix Validation` heading is ABSENT from the file (covers TC-022)
- [x] 3.7 Ensure `## Catalog Validation` heading is ABSENT from the file (covers TC-023)
- [x] 3.8 Include `## Safe Evidence Rules` section verbatim (all 4 lines from original lines 108–111) (covers TC-024)

## Phase 4: Static Verification Evidence

- [x] 4.1 Inspect `review-control-catalog.json`: confirm root keys match v2 header spec; confirm `controls.length === 96`; confirm zero occurrences of `"reviewer"` in any control entry; confirm zero occurrences of `"notes": "Not provided"` (covers TC-001–TC-009)
- [x] 4.2 Inspect `review-report.schema.json`: confirm `required` array has exactly 14 fields in exact order; confirm no legacy field name appears anywhere in the file (covers TC-010, TC-011)
- [x] 4.3 Inspect `review-report.schema.json $defs`: confirm `ref` and `stateArtifactRef` definitions present; confirm `matrixRow.finding` is `type: string` with no `enum`; confirm `item` pattern covers REV-CORP-001–096; confirm `stateRegistration.minItems === 2`; confirm `additionalProperties: false` on root and all 5 nested objects (covers TC-012–TC-017)
- [x] 4.4 Inspect `report-template.md`: confirm lines 1–8 unchanged; confirm `## JSON Field Mapping` table maps exactly 14 v2 fields; confirm `## Required Structure` present; confirm 9-column matrix header; confirm `## Matrix Rules` with join instruction; confirm `## Matrix Validation` absent; confirm `## Catalog Validation` absent; confirm `## Safe Evidence Rules` verbatim (covers TC-018–TC-025)
- [x] 4.5 Cross-inspect all three files: catalog `schemaVersion: 2` consistent with schema `schemaVersion const 2`; template field-mapping table references exactly 14 schema required fields; template matrix columns match `matrixRow` $def fields (covers TC-026)
- [x] 4.6 Confirm `skills/sdd-review-security/` directory has no modified files (`git diff` or equivalent) (covers TC-027)
- [x] 4.7 Record unavailable tooling: no test runner, unit, integration, e2e, coverage, linter, type checker, or formatter available — all checks in this change are static/manual per `openspec/config.yaml` (covers TC-028)
