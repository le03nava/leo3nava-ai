# Proposal: Source Row First Security Review

## Intent

Redesign `sdd-review-security` as a breaking contract change so the 155 corporate `Source ID` rows are the only canonical security validation matrix. Compact `SEC-*` controls must be removed from the active report contract; navigation and grouping must use `controlDomain`, `corporateSection`, or another source-row grouping field instead.

## Scope

### In Scope
- Replace `review-security-report.json` validation authority with `sourceRowValidation.rows` containing exactly 155 unique rows.
- Require per-row fields: `sourceId`, `corporateSection`, `pciAlignment`, `guidelineText` or `guidelineRefs`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, and `route`.
- Remove compact controls from active navigation, summaries, and grouped `N/A` justifications; use source-row categories such as `controlDomain` and `corporateSection` instead.
- Simplify derived `report-template.md`: JSON-authoritative, lean human navigation only, with verdict, handoff, summaries, `N/A` groups, blockers/warnings, and the full row matrix at the end.
- Update specs, shared contracts, schema/template/validation/catalog references, exporter defaults/tests/docs, and relevant SDD skill instructions.

### Out of Scope
- Historical compatibility or migration for old compact-control reports.
- Changing the general `sdd-review` 96-control matrix.
- Adding a new runtime beyond existing Python exporter validation.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-review-security-workflow`: source-row-first canonical validation, routing, Markdown rendering, and category/domain-based `N/A` grouping rules.
- `sdd-security-guideline-catalog`: Source ID inventory, row schema vocabulary, category/domain grouping metadata, and safe `N/A` policy.
- `sdd-execution-persistence-contracts`: downstream security-review artifact authority and verify/archive consumption.
- `canonical-review-json-excel-exporter`: default security-review table changes to `sourceRowValidation.rows` with tests/docs.

## Approach

Make canonical JSON the only validation authority. Validation must prove exact-once coverage of all 155 Source IDs, reject missing/duplicate/unknown rows, and enforce safe row-level justification for every `N/A`. Grouped `N/A` is allowed only when equivalent non-applicability is preserved in row-level JSON. Markdown is generated from JSON and must not duplicate validation logic.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-review-security/` | Modified | Skill, schema, template, catalog, and validation rules. |
| `skills/_shared/` | Modified | Security and post-apply contracts. |
| `openspec/specs/` | Modified | Workflow/catalog/exporter/persistence delta specs. |
| `python/` | Modified | Excel exporter default table path, tests, README. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Large contract update misses a consumer | Med | Spec all affected capabilities and verify downstream refs. |
| Markdown becomes a second source of truth | Med | Keep template lean and generated only from JSON. |
| `N/A` grouping hides row gaps | Med | Require row-level preservation and exact-once validation. |

## Rollback Plan

Revert this OpenSpec change before archive, or after implementation revert the modified skill/contracts/spec/exporter files together. No compatibility migration is required.

## Dependencies

- Existing corporate 155-row Source ID catalog snapshot and recommended vocabularies supplied by the user.

## Success Criteria

- [ ] `review-security-report.json` validates exactly 155 `sourceRowValidation.rows`, each once.
- [ ] Compact `SEC-*` controls are removed from the active report contract and are not used for validation, navigation, summaries, or `N/A` grouping.
- [ ] Lean Markdown removes `## Compact Control Validation` and renders the full matrix last.
- [ ] Exporter default/tests/docs use `sourceRowValidation.rows`.
