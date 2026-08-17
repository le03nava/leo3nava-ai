# Proposal: SDD Review Security Homologation

## Intent

The 4 reference files of the `sdd-review-security` skill were designed before the `sdd-review-schema-v2` simplification pattern was established. They carry nested substructures, redundant catalog fields in every row, and header objects (`authority{}`, `reporting{}`) that the v2 pattern eliminated. This creates a schema drift between the two review skills: downstream consumers (sdd-verify, sdd-archive) and the SKILL.md itself must handle two structural dialects. Homologating all 4 files to the v3 pattern closes that drift without changing the catalog's 155 rows or the runtime validation rules.

## Scope

### In Scope
- Bump `schemaVersion` to `3` in the catalog; rewrite `review-security-report.schema.json` to v3
- Remove `generatedHumanViewRef`, `reporting{}`, `authority{}` from catalog header
- Remove `guidelineRefs`, `evidenceExpectation` from all 155 `sourceRows`; rename `defaultOwnerPhase`→`ownerPhase` and `defaultRoute`→`route`
- Rewrite `review-security-report.schema.json`: 16 flat required fields, `rows[]` at root, `sourceRow` $def without catalog fields, drop removed $defs (`navigationGroup`, `groupedNaSummary`, `finding`)
- Rewrite `report-template.md` field mapping table and skeleton (9 sections; join instruction for matrix columns)
- Update `validation-rules.md`: `rows[]` path, updated required-row fields (no catalog fields), remove Row-Preserving Non-Applicability section and `guidelineText`/`guidelineRefs` rule, update checklist

### Out of Scope
- No changes to the 155 `guidelineText` values or `corporateSection`/`pciAlignment`/`controlDomain`/`repoProfiles`/`runtimeSurface`/`dataSurface`/`appliesWhen` catalog columns (these stay in the catalog for join at render time)
- No changes to `SKILL.md` (narrative update is a separate change)
- No changes to any `openspec/specs/` capability files
- No changes to `sdd-review` skill references

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- None

> This change is a pure reference-artifact homologation (schema version bump + structural cleanup). No spec-level behavior changes. No new capabilities introduced.

## Approach

Four targeted file rewrites under `skills/sdd-review-security/references/`. The catalog's content (155 rows, all `guidelineText` values) is preserved; only the structural fields identified in the exploration are removed or renamed. The schema rewrite is a full replacement of the JSON Schema object following the flat-field pattern already validated in `sdd-review-schema-v2`. Template and rules files are updated to reference the new paths and drop eliminated sections.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-review-security/references/security-guideline-catalog.operational.json` | Modified | Header cleanup + 155-row field rename; no content loss |
| `skills/sdd-review-security/references/review-security-report.schema.json` | Modified | Full rewrite to v3 flat schema (16 required fields) |
| `skills/sdd-review-security/references/report-template.md` | Modified | New 9-section field mapping; updated matrix join instruction |
| `skills/sdd-review-security/references/validation-rules.md` | Modified | Updated row path, required fields, removed obsolete sections |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Existing `review-security-report.json` artifacts in prior changes reference v2 schema fields | Low | Archived reports are immutable; only future runs consume the new schema |
| 155-row rename (`defaultOwnerPhase`→`ownerPhase`) missed on any row | Low | Apply phase must run a count assertion: `jq '[.sourceRows[] | select(has("defaultOwnerPhase"))] | length'` must return `0` |
| `report-template.md` join instruction ambiguous for future agents | Low | Include explicit column-source table in the template skeleton |

## Rollback Plan

All 4 files are in the repository. If the homologation introduces regressions, revert the 4 files via `git revert` or `git checkout HEAD~1 -- skills/sdd-review-security/references/`. No runtime state is mutated.

## Dependencies

- None. The `sdd-review-schema-v2` exploration is already complete and its pattern is the reference.

## Success Criteria

- [ ] `security-guideline-catalog.operational.json` `schemaVersion` is `3`; no `generatedHumanViewRef`, `reporting`, or `authority` keys present at top level
- [ ] No `guidelineRefs`, `evidenceExpectation`, `defaultOwnerPhase`, or `defaultRoute` keys present in any of the 155 `sourceRows`
- [ ] `review-security-report.schema.json` `required[]` contains exactly the 16 v3 fields; `rows[]` appears at the report root; no `sourceRowValidation` wrapper
- [ ] `report-template.md` field mapping covers the 9 v3 sections with catalog join instruction
- [ ] `validation-rules.md` references `rows[]` (not `sourceRowValidation.rows`), lists v3 required row fields only, and contains no Row-Preserving Non-Applicability section
