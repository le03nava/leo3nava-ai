# Archive Report: sdd-review-schema-v2

## Summary

This archive captures the completed SDD change `sdd-review-schema-v2` and the synced delta spec for `sdd-review-workflow` into the active OpenSpec specs where required.

## Synced Specs

- Synced delta `openspec/changes/archive/2026-07-12-sdd-review-schema-v2/specs/sdd-review-workflow/spec.md` into active spec `openspec/specs/sdd-review-workflow/spec.md` by applying non-destructive narrative updates that preserve existing workflow behavior while promoting the v2 contract for report schema, catalog, and template.

## Moved Files

All change artifacts were moved to `openspec/changes/archive/2026-07-12-sdd-review-schema-v2/`. Implemented reference artifacts under `skills/sdd-review/references/` were left in place as permanent deliverables.

## Evidence

- Verify report: `openspec/changes/archive/2026-07-12-sdd-review-schema-v2/verify-report.md`
- Applied delta spec (archived copy): `openspec/changes/archive/2026-07-12-sdd-review-schema-v2/specs/sdd-review-workflow/spec.md`
- Implemented references: `skills/sdd-review/references/review-control-catalog.json`, `skills/sdd-review/references/review-report.schema.json`, `skills/sdd-review/references/report-template.md`

## Archive Authorization

Archive performed under reconciliation reason: "Verify passed PASS WITH WARNINGS 28/28 cases verified; apply-progress complete". Stale checkboxes approval present.

## Notes and Risks

- Risk: Follow-on `sdd-review-security` changes required for full v2 support; currently documented as a warning in verification.
- Risk: Unavailable runtime/test tooling preserved as unavailable evidence — no automated test runner ran.
- No source-row blockers or unsafe evidence found.

## Next Recommended

none

## Artifacts

- `archive-report.md`
