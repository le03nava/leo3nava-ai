# Proposal: sdd-review Schema v2 — Reference Artifact Cleanup

## Intent

The three reference artifacts in `skills/sdd-review/references/` have accumulated legacy fields (`blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `changedFileSecurityHandoff`, `operationalEvidenceSummary`, `generatedHumanViewRef`, `presentation`, `validation`, `safeEvidence`, `forbiddenEvidence`, reviewer-level `notes`) that do not contribute to review correctness and create maintenance noise. The v2 schema removes those fields, tightens object shapes with `additionalProperties: false`, introduces a `Finding` column in the 96-control matrix, and aligns the report template with the new field contract. The goal is a leaner, stricter schema that reduces the surface area reviewers must maintain per change.

## Scope

### In Scope
- `skills/sdd-review/references/review-control-catalog.json` — bump `schemaVersion` to 2, add root `reviewer` field, remove `generatedHumanViewRef`, `presentation`, `validation`, `safeEvidence`, `forbiddenEvidence`; drop per-control `reviewer` and `notes` (when value is `"Not provided"`) from all 96 entries
- `skills/sdd-review/references/review-report.schema.json` — full rewrite to 14 required root fields; remove `blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, `operationalEvidenceSummary`; add `$defs` for `ref` and `stateArtifactRef`; add `artifactMetadata` with `stateRegistration` (minItems: 2); enforce `additionalProperties: false` on all objects
- `skills/sdd-review/references/report-template.md` — rewrite from `## JSON Field Mapping` onward: new field-mapping table, new required-structure skeleton with full 96-row matrix spec, updated matrix columns (`Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes`), remove `## Matrix Validation` and `## Catalog Validation` sections, preserve `## Safe Evidence Rules` verbatim (lines 108–111)

### Out of Scope
- `skills/sdd-review/SKILL.md` — no behavior changes in this change
- `skills/sdd-review-security/` — see **Dependency Note** below; coordinated update is deferred to a separate change
- Any `openspec/changes/*/review-report.json` files already archived — existing reports are not migrated
- `openspec/specs/sdd-review-workflow/spec.md` and `openspec/specs/sdd-review-security-workflow/spec.md` — spec delta coverage only if required; behavior spec is unchanged

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- `sdd-review-workflow`: schema contract changes (`review-report.schema.json` v2, catalog v2, template rewrite); the review output shape changes but the review workflow behavior does not
- `sdd-review-security-workflow`: **out-of-scope update deferred** — fields removed from the schema (`changedFileSecurityHandoff`, `operationalEvidenceSummary`) are currently consumed by `sdd-review-security`. That skill must be updated in a follow-on change. This change explicitly breaks that contract and acknowledges the dependency.

## Approach

Edit-in-place on the three reference files. No new files. No runtime code changes.

1. **Catalog** — `schemaVersion: 2`, add `"reviewer": "sdd-review"` at root, remove `generatedHumanViewRef`, `presentation`, `validation`, `safeEvidence`, `forbiddenEvidence`; iterate all 96 controls and drop `reviewer` field and `notes` where value is exactly `"Not provided"`.
2. **Schema** — full rewrite of `review-report.schema.json` keeping the 14 required root fields; new `$defs` for `ref` and `stateArtifactRef`; `stateRegistration` array with `minItems: 2` expecting exactly: ① canonical JSON review-report ref, ② derived Markdown review-report ref (per the existing persistence-contract `artifactRefs.reviewReport` order); `additionalProperties: false` everywhere.
3. **Template** — rewrite from `## JSON Field Mapping` onward; matrix columns include `Finding` as a new rendered column that maps to `matrix[].finding` in the JSON schema (a free-text string, not a controlled enum); keep `## Safe Evidence Rules` block verbatim.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-review/references/review-control-catalog.json` | Modified | schemaVersion 2, remove legacy fields, 96-control cleanup |
| `skills/sdd-review/references/review-report.schema.json` | Modified (full rewrite) | 14 required fields, new `$defs`, `artifactMetadata`, `Finding` column, strict `additionalProperties: false` |
| `skills/sdd-review/references/report-template.md` | Modified (partial rewrite) | New field mapping table, updated matrix skeleton, remove Validation sections |
| `skills/sdd-review-security/` | **Not touched** | Consumes removed fields — coordinated update is a follow-on change |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| `sdd-review-security` breaks because `changedFileSecurityHandoff` / `operationalEvidenceSummary` are removed | High (certain if not coordinated) | Explicitly out of scope; follow-on change required before running `sdd-review-security` against a v2 report |
| Existing archived `review-report.json` files fail validation against v2 schema | Low | Schema is reference-only; no runtime validator enforces it retroactively |
| `stateRegistration minItems: 2` misconfigured | Low | Proposal defines the two expected entries (canonical JSON ref + derived Markdown ref per persistence-contract order) |
| `report-template.md` `## Safe Evidence Rules` section accidentally dropped | Low | Explicit instruction to preserve lines 108–111 verbatim |

## Rollback Plan

All three files are in `skills/sdd-review/references/` under git. Revert via `git revert` or restore from the pre-change commit SHA. No runtime state is affected.

## Dependencies

- Follow-on change required: update `skills/sdd-review-security/` to remove its dependency on `changedFileSecurityHandoff` and `operationalEvidenceSummary` before any `sdd-review-security` run against a v2 report is attempted.

## Success Criteria

- [ ] `review-control-catalog.json` has `schemaVersion: 2`, root `reviewer` field, no `safeEvidence`/`forbiddenEvidence`, no `generatedHumanViewRef`/`presentation`/`validation` objects, and all 96 controls have `reviewer` and blank-value `notes` removed
- [ ] `review-report.schema.json` validates a minimal v2 report with the 14 required fields, rejects extra root properties, and enforces `stateRegistration` with exactly 2 entries
- [ ] `report-template.md` shows updated `## JSON Field Mapping`, full 96-row matrix skeleton with `Finding` column, no `## Matrix Validation` or `## Catalog Validation` sections, and `## Safe Evidence Rules` matches the current verbatim text
- [ ] `skills/sdd-review-security/` is **unchanged** and a follow-on change is registered
