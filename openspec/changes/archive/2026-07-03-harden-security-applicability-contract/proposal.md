# Proposal: Harden Security Applicability Contract

## Intent

Make `security-applicability.md` audit-grade: every security decision must show explicit category evaluation, no-impact proof, formal corporate source coverage, operational blocking metadata, supported overrides, and automatic static validation.

## Problem

Current applicability artifacts can classify impact, but they do not force category-by-category decisions, separate no-impact proof from missing evidence, or validate `SEC-*` mappings against corporate source IDs. This creates audit ambiguity and relies too much on agent discipline.

## Scope

### In Scope
- Harden the shared applicability schema with catalog identity, evaluated categories, decision matrix, source evidence refs, explicit no-impact evidence, and validation metadata.
- Treat `SEC-* -> Source IDs` mappings as formal corporate source coverage, not best-effort technical references.
- Add operational metadata using `blocking`, `conditional`, and `advisory`; do not reuse review labels such as `Menor`, `Media`, or `Mayor`.
- Document supported `rules.security-applicability` overrides, including strict source coverage and validator mode.
- Add an automatic static validator for `security-applicability.md` artifacts.
- Preserve no-impact routing compatibility and add downstream compatibility notes where needed.

### Out of Scope
- Replacing the in-repo corporate source snapshot with an external policy system.
- Renaming existing `SEC-*` IDs or making `sdd-review` authoritative for security applicability.
- Splitting the large security guideline catalog unless later sizing proves it necessary.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-security-applicability-workflow`: require auditable matrix/no-impact proof, config overrides, and validator use.
- `sdd-security-guideline-catalog`: require source coverage, catalog metadata, and operational severity/blocking semantics.
- `sdd-security-design-workflow`: consume enriched applicability fields without breaking no-impact routing.

## Approach

Use schema-first hardening while retaining the compact catalog. Update shared contracts, applicability template, catalog columns/sections, OpenSpec requirements, and add a small static validator/check command scoped to `security-applicability.md` structure and references.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-security-applicability/SKILL.md` | Modified | Artifact template, gates, overrides, validator expectations. |
| `skills/_shared/sdd-security-contract.md` | Modified | Authoritative schema fields and validation contract. |
| `skills/_shared/security-guideline-catalog.md` | Modified | Formal source coverage and blocking metadata. |
| `openspec/specs/sdd-security-*` | Modified | Delta requirements for workflow/catalog/design compatibility. |
| validation support path TBD by design | New | Static validator for applicability artifacts. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Incorrect source mappings create false audit confidence. | Med | Require formal source coverage evidence and validator checks. |
| Required fields make no-impact artifacts noisy. | Med | Keep template matrix compact and reviewer-scannable. |
| Validator scope grows beyond repo tooling. | Med | Limit to static Markdown/YAML structure and references. |

## Rollback Plan

Revert the change folder/spec deltas and any contract/catalog/validator edits from the stacked PR slices. Existing archived `security-applicability.md` artifacts remain readable because artifact paths, `SEC-*` IDs, and no-impact routing are preserved.

## Dependencies

- Corporate source IDs in the current in-repo snapshot are the authoritative mapping basis.

## Review Slicing

Use stacked-to-main slices: contract/template/specs, catalog/source metadata, then validator/downstream compatibility. Keep each slice near the 400-line review budget.

## Success Criteria

- [ ] Specs can require matrix completeness, no-impact proof, source coverage, blocking metadata, overrides, and validation.
- [ ] Next phases can design validation without changing established artifact identity or routing.
