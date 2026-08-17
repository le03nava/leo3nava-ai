# Proposal: Active-Only Security Contract

## Intent

Keep active SDD security surfaces focused on the current flow: `design.md#secure-development-design` -> `review-security-report.md` -> `verify` -> `archive`. Active contracts still mix current rules with legacy standalone artifact prose, which makes the current workflow harder to understand and easier to misuse.

## Scope

### In Scope
- Remove legacy standalone security sections, schemas, and compatibility prose from active security contract surfaces, especially `skills/_shared/sdd-security-contract.md`.
- Rewrite active phase guidance as positive current-flow rules: embedded secure design, security review report, catalog-backed evidence, and exception fields.
- Preserve required status/persistence parser behavior for historical refs only when needed for old state readability.
- Update source specs whose active requirements still describe legacy standalone security behavior.

### Out of Scope
- No archive rewrite: do not edit `openspec/changes/archive/**`.
- Do not create a separate legacy security contract.
- Do not remove historical artifact readability if persistence/status readers still need it.
- Do not introduce runtime tooling or a test runner.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-security-guideline-catalog`: align catalog purpose/scope with current-flow evidence only.
- `sdd-review-security-workflow`: remove active legacy standalone scenarios and require embedded design plus review-security evidence.
- `sdd-design-workflow`: keep embedded secure development design as the active classification/design authority.
- `sdd-execution-persistence-contracts`: preserve historical-read behavior without making legacy security artifacts active dependencies.
- `sdd-security-applicability-workflow`: keep as legacy-readability boundary only if specs still need to state non-active behavior.

## Approach

Use active-contract cleanup with compatibility isolated to low-level status/persistence readers. Remove legacy schemas and prose from active security docs; retain only minimal historical-read fields/tokens where deletion could break old state or archive display. Prefer current-flow wording over “legacy compatibility” wording in active surfaces.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared/sdd-security-contract.md` | Modified | Make the shared contract active-only. |
| `skills/_shared/security-guideline-catalog.md` | Modified | Clean active scope wording. |
| `skills/sdd-design/SKILL.md` | Modified | Positive embedded-design rules. |
| `skills/sdd-review-security/SKILL.md` | Modified | Positive review-security requirements. |
| `openspec/specs/*/spec.md` | Modified | Delta specs for active contract alignment. |
| `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md` | Conditional | Preserve parser/read behavior only if required. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Removing historical refs breaks old state reads | Medium | Keep minimal parser/read behavior outside active security contract. |
| Active docs still leak legacy framing | Medium | Verify with targeted searches for legacy filenames and compatibility wording. |
| No automated runner exists | High | Use static inspection and read-back verification. |

## Rollback Plan

Revert the Markdown/spec changes for this change. Because archives are not edited and runtime behavior is not introduced, rollback is a git revert of active contract/spec edits.

## Dependencies

- Existing exploration: `openspec/changes/active-only-security-contract/explore.md`.
- Existing OpenSpec specs under `openspec/specs/` for capability deltas.

## Success Criteria

- [ ] Active security contract surfaces describe only the current embedded-design and review-security flow.
- [ ] No new or separate legacy contract exists.
- [ ] Archives remain untouched.
- [ ] Required historical state/artifact readability is preserved where needed.
