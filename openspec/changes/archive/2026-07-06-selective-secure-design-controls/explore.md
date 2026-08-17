# Exploration: selective-secure-design-controls

## Current State

The active SDD security flow already keeps the exhaustive corporate Source ID matrix out of design: `review-security-report.md` is the only active artifact that materializes all 155 Source IDs, while `design.md#secure-development-design` cites catalog identity, expected count, grouped coverage, compact mappings, and downstream ownership. That boundary was established by the archived `slim-secure-design-artifacts` change and is reflected in the current shared contract, design skill, review-security skill, and OpenSpec specs.

However, `sdd-design` still requires too much non-applicable evidence at the compact-control layer. Current contracts require `design.md#secure-development-design` to include all eight compact `SEC-*` IDs exactly once and require every `N/A` / `not-applicable` row to include evidence and rationale before design can succeed. The source specs mirror that behavior: `openspec/specs/sdd-design-workflow/spec.md` says design must evaluate all catalog rows and, for no-impact changes, mark every row as `N/A` or `not-applicable` with rationale. `skills/_shared/sdd-security-contract.md` repeats the same all-row design requirement.

`sdd-review-security` already owns stronger validation semantics: it runs after general review, consumes design, test-design, apply evidence, changed files, and the catalog, expands all Source IDs when applicable, validates `N/A` evidence, and blocks missing applicable evidence. This makes it the natural owner for exhaustive applicability decisions across the full compact/security matrix and Source ID inventory.

## Affected Areas

- `skills/sdd-design/SKILL.md` — currently requires all eight compact rows, N/A rationale/evidence for every non-applicable row, and source-row coverage metadata in design. It should shift to selective design: identify categories/controls that apply to the planned design changes, provide guidance and evidence owners for those applicable controls, and record enough classification/catalog context for review-security to audit omissions.
- `skills/_shared/sdd-security-contract.md` — currently defines `design.md#secure-development-design` as owning all compact rows and no-impact N/A rows. It should redefine design as an applicable-control planning artifact while assigning exhaustive non-applicability decisions and omission validation to `review-security-report.md`.
- `skills/sdd-review-security/SKILL.md` — already validates compact rows and all 155 Source IDs, but its wording assumes embedded design rows are exhaustive. It should explicitly expand the full compact/security catalog itself, decide/report non-applicable compact controls and Source IDs, validate design omissions, and block when design missed an applicable category/control.
- `openspec/specs/sdd-design-workflow/spec.md` — must change requirements/scenarios that require all eight compact rows and every no-impact N/A row in design.
- `openspec/specs/sdd-review-security-workflow/spec.md` — should add requirements for exhaustive applicability validation, non-applicable reporting, and missed-applicable-control blockers in review-security.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` — likely needs only boundary wording: the catalog remains the taxonomy/source inventory; design selectively consumes it; review-security exhaustively evaluates it.
- `openspec/specs/sdd-test-design-workflow/spec.md` — should consume applicable design controls only, while preserving review-security as the exhaustive validation gate. Existing no-impact scenarios that expect all security rows in design may need revision.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` — likely minor updates to active security artifact semantics if the state/status descriptions imply design owns exhaustive compact rows.
- `agents/sdd/sdd-design.md` and `agents/sdd/sdd-review-security.md` — adapter prompts mirror the skill contracts and must stay aligned. `agents/sdd/sdd-test-design.md` may also need adjustment if it assumes design carries all compact or N/A rows.
- Archived `slim-secure-design-artifacts` — useful context only; no migration should rewrite archived artifacts. Compatibility wording should keep older exhaustive design rows readable without making them mandatory for new changes.

## Approaches

1. **Selective design, exhaustive review-security** — Design records classification, catalog snapshot, changed security surface, applicable compact categories/controls, design guidance, evidence owners, residual risks, exceptions, and safe-evidence policy. It does not enumerate every non-applicable compact control or Source ID. Review-security expands the full compact catalog and Source ID inventory, reports applicable/non-applicable decisions, validates design omissions, and blocks missed applicable controls.
   - Pros: Best matches the request, reduces design noise, keeps security auditability, and puts exhaustive validation in the phase with full implementation evidence.
   - Cons: Requires careful contract language so design omissions are treated as reviewable, not silently acceptable.
   - Effort: Medium.

2. **Keep all compact rows in design, move only Source ID N/A to review-security** — Preserve the current eight-row design matrix but avoid Source ID-level N/A obligations in design.
   - Pros: Smaller change and preserves existing downstream assumptions.
   - Cons: Does not satisfy the core request because design still marks every non-applicable compact control as `N/A`.
   - Effort: Low, but incomplete.

3. **Move all security applicability out of design** — Design would only include a generic security note, and review-security would determine all applicability later.
   - Pros: Smallest design artifact.
   - Cons: Too late for design-time secure architecture guidance; test-design/tasks would lack planned security evidence for applicable controls.
   - Effort: Medium with high workflow risk.

## Recommendation

Use approach 1. The proposal should define `design.md#secure-development-design` as a selective secure-design planning section, not an exhaustive matrix. Design must still prove its classification from proposal/specs/code context and must identify all categories/controls that appear applicable at design time. For those applicable categories, design should provide concrete guidance, downstream evidence owners, expected evidence, residual risk, and exception policy.

`sdd-review-security` should become the explicit exhaustive authority for the full compact/security matrix and corporate Source ID universe. It should expand the catalog, decide and report which compact controls and Source IDs are applicable or non-applicable, validate that non-applicable decisions have safe rationale/evidence, compare against the selective design section, and block when an omitted category/control should have been designed.

This keeps security design useful at the architecture point while moving audit-heavy `N/A` bookkeeping to the phase designed to validate evidence.

## Risks

- If design becomes too terse, applicable controls may be missed until late review. Mitigation: require design to include a changed-surface inventory and explicit applicability rationale for included controls, while review-security blocks missed applicable controls.
- `sdd-test-design` currently expects embedded security rows and N/A rationale. Its contract must be updated so it plans checks from applicable design controls and records no-impact assessment without requiring a full N/A matrix.
- Adapter prompts can drift from skill contracts. Update `agents/sdd/*` copies with the same boundary language during apply.
- Archived changes may contain all-eight-row designs. Compatibility language must keep them readable and valid as historical artifacts while changing only the new-change contract.
- Review-security reports may grow because they now report compact non-applicability decisions as well as Source ID non-applicability decisions. This is acceptable because review-security already owns exhaustive validation.

## Ready for Proposal

Yes. The proposal should target Markdown/OpenSpec contract changes only. No runtime code or executable tests are expected in this repository; verification should use static/manual artifact checks and read-back evidence.
