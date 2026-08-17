# Apply Progress — sdd-review-schema-v2

## Status

- Change: `sdd-review-schema-v2`
- Mode: Standard
- Overall: ✅ Completed
- Next recommended: `review`

## Completed Tasks

- [x] Phase 1 — Catalog Update (`review-control-catalog.json`)
- [x] Phase 2 — Schema Rewrite (`review-report.schema.json`)
- [x] Phase 3 — Template Partial Rewrite (`report-template.md`)
- [x] Phase 4 — Static Verification

## Files Changed

| File | Action | Summary |
|---|---|---|
| `skills/sdd-review/references/review-control-catalog.json` | Modified | Migrated catalog to schemaVersion 2, removed deprecated root blocks, simplified vocabulary, added root reviewer, removed per-control reviewer and placeholder notes. |
| `skills/sdd-review/references/review-report.schema.json` | Rewritten | Replaced with v2 canonical schema contract using 14 required root fields, new totals/matrix/artifactMetadata model, and minimal `$defs`. |
| `skills/sdd-review/references/report-template.md` | Modified | Kept preamble unchanged and replaced content from `## JSON Field Mapping` onward with v2 rendering contract and matrix/evidence rules. |
| `openspec/changes/sdd-review-schema-v2/tasks.md` | Modified | Marked all apply tasks as completed (`[x]`). |

## Static Verification Results

1. **Catalog checks**
   - `schemaVersion === 2` ✅
   - Root fields `generatedHumanViewRef`, `presentation`, `validation` absent ✅
   - `vocabulary` contains only `complies` and `severity` ✅
   - Root `reviewer: "sdd-review"` present ✅
   - Spot/manual + scripted checks confirm no control-level `reviewer` and no `notes: "Not provided"` ✅
   - `controls.length === 96` ✅

2. **Schema checks**
   - `required` has exactly 14 entries in requested order ✅
   - Legacy fields absent (`blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, `operationalEvidenceSummary`, `presentation`) ✅
   - `$defs` includes `ref` and `stateArtifactRef` (plus `matrixRow`) ✅
   - Root `additionalProperties: false` ✅

3. **Template checks**
   - `## JSON Field Mapping` present ✅
   - `## Required Structure` present ✅
   - `## Matrix` section present with join instruction ✅
   - `## Safe Evidence Rules` present ✅
   - `## Matrix Validation` absent ✅
   - `## Catalog Validation` absent ✅

4. **Additional checks**
   - `skills/sdd-review-security/` unchanged ✅
   - Tooling for this apply is static/manual validation only (no runtime test/lint/typecheck/coverage execution) ✅

## Risks / Notes

- Task list wording for schema `$defs.ref`/`$defs.stateArtifactRef` in `tasks.md` reflects older shape labels, but implemented schema follows the explicit apply envelope v2 contract and passes required structure checks.

## Remediation Batch Merge — Verification Failures TC-017, TC-018, TC-021, TC-024, TC-026

### Scope Guard

- Applied only within allowed remediation files for this batch.
- Persisted previous progress unchanged and merged this remediation evidence into the same artifact.

### Fix Results

1. **TC-017 — `$defs.ref` additionalProperties (false failure)**
   - No schema code change applied in `review-report.schema.json`.
   - `$defs.ref` remains `{ "type": "string", "minLength": 1 }`, which is correct for JSON Schema 2020-12.
   - `additionalProperties: false` is only applicable to object schemas; it is inapplicable to string schemas.
   - Verified object definitions already enforce `additionalProperties: false` where relevant (`root`, `totals`, `artifactMetadata`, `stateArtifactRef`, `matrixRow`).

2. **TC-018 + TC-026 — JSON Field Mapping completeness**
   - Replaced `## JSON Field Mapping` table in `report-template.md`.
   - Table now explicitly lists all 14 required root JSON fields with type and rendering destination.

3. **TC-021 — Matrix join instruction exact wording**
   - Updated `## Matrix` rendering rule line in `report-template.md` to exactly:
   - `Rendering rule: join `matrix[].item` with `catalog controls[].id` to get `Requirement`, `Standard`, `Category`, `Severity` columns.`

4. **TC-024 — Safe Evidence Rules verbatim**
   - Confirmed `## Safe Evidence Rules` section matches required 4-line canonical wording exactly.
   - No additional line content introduced beyond the required block.

### Remediation Files Touched

| File | Action | Summary |
|---|---|---|
| `skills/sdd-review/references/review-report.schema.json` | Verified (no change) | Confirmed TC-017 is verifier interpretation issue; string `$defs.ref` is spec-correct. |
| `skills/sdd-review/references/report-template.md` | Modified | Replaced JSON field mapping table, set matrix rule to exact required sentence, preserved verbatim safe evidence rules. |
| `openspec/changes/sdd-review-schema-v2/apply-progress.md` | Modified | Merged remediation batch evidence without overwriting prior apply history. |
