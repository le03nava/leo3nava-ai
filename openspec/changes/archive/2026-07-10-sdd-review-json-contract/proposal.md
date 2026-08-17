# Proposal: SDD Review JSON Contract

## Intent

Make `sdd-review` use a canonical JSON contract as the source of truth for the general review report and its 96 controls. Markdown must be generated from JSON, not maintained as the authoritative control/report definition. The JSON should also be presentation-friendly for future Excel generation without implementing Excel output now.

## Scope

### In Scope
- Define canonical `review-report.json` with schema metadata, verdict/routing, evidence summaries, blocking summary, vocabulary, validation metadata, and all 96 `REV-CORP-*` controls directly in JSON.
- Generate `review-report.md` from the canonical JSON during `sdd-review`, preserving the existing human-readable report sections and required matrix columns.
- Update `sdd-review` contracts, templates, OpenSpec requirements, and shared post-apply/persistence wording so JSON is authoritative and Markdown is derived.
- Migrate or replace `control-catalog.md` semantics so it is generated/derived/reference-only, not the control source of truth.

### Out of Scope
- Excel generation, Python scripts, or spreadsheet formatting.
- Changing `sdd-review-security` ownership, security-review matrices, or embedded secure-design authority.
- Removing `review-report.md` as a downstream compatibility artifact.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-review-workflow`: General review report persistence changes from Markdown-authored controls to canonical JSON controls with derived Markdown presentation.

## Approach

Use `skills/sdd-review-security/references/security-guideline-catalog.operational.json` as the model for a machine-readable contract. Add a general-review JSON catalog/report contract under `skills/sdd-review/references/`, require `sdd-review` to validate the JSON row count, IDs, vocabulary, and derived Markdown parity, and keep `review-report.md` at the existing path for downstream phases.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-review/SKILL.md` | Modified | JSON source-of-truth execution and validation rules. |
| `skills/sdd-review/references/*` | Modified/New | Canonical JSON contract; Markdown template/catalog become generated or derived views. |
| `openspec/specs/sdd-review-workflow/spec.md` | Modified | Requirements for JSON authority and Markdown derivation. |
| `skills/_shared/persistence-contract.md` | Modified | Review artifact refs include canonical JSON plus derived Markdown. |
| `skills/_shared/sdd-post-apply-gates.md` | Modified | Downstream phases consume derived Markdown while recognizing JSON authority. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| JSON/Markdown drift | Medium | Generate Markdown from JSON and validate read-back parity. |
| Downstream breakage | Medium | Keep `review-report.md` path and required sections/columns stable. |
| Overfitting to Excel | Low | Store presentation metadata only; defer Python/Excel rendering. |

## Compatibility Boundaries

- JSON is authoritative for all general review controls and report facts.
- `review-report.md` remains a derived compatibility artifact for humans and existing downstream phases.
- Historical Markdown catalog authority does not need preservation.

## Rollback Plan

Revert the proposal/spec/design/apply changes and restore `sdd-review` to writing only `review-report.md` from the Markdown catalog/template. Existing downstream phases remain protected because the derived Markdown path is preserved during the change.

## Dependencies

- Existing 96-control IDs `REV-CORP-001..REV-CORP-096` and required review matrix columns.
- Existing security JSON pattern as a reference, not a dependency to modify.

## Success Criteria

- [ ] `sdd-review` has a canonical JSON contract containing all 96 general controls.
- [ ] `review-report.md` is generated from JSON and preserves current downstream compatibility.
- [ ] Specs and shared contracts state JSON authority and Excel generation remains deferred.
