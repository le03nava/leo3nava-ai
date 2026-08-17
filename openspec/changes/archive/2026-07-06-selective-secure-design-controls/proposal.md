# Proposal: Selective Secure Design Controls

## Intent

Reduce `sdd-design` security noise while preserving audit strength. `design.md#secure-development-design` should plan only applicable security categories/controls for the planned design changes. `sdd-review-security` should own exhaustive compact-control and Source ID validation, including non-applicable decisions and missed applicable controls.

## Scope

### In Scope
- Update `sdd-design` skill and adapter prompt so design identifies changed security surface, applicable categories/controls, guidance, evidence owners, residual risk, and exceptions without enumerating all non-applicable compact controls or Source IDs.
- Update shared security contract to define design as selective planning and review-security as exhaustive validation authority.
- Update `sdd-review-security` skill and adapter prompt to expand the full catalog, decide/report applicable and non-applicable rows, validate omissions, and block missed applicable controls.
- Update OpenSpec source specs for design, review-security, security catalog boundary, test-design consumption, and persistence/status semantics where affected.
- Update test-design prompt/contract if it assumes all compact or `N/A` rows are present in design.

### Out of Scope / Non-Goals
- No runtime application code changes.
- No rewrite of archived changes or historical exhaustive design artifacts.
- No removal of the compact taxonomy, catalog inventory, Source ID count, or review-security exhaustive report.
- No weakening of mandatory evidence, exception, or safe-evidence requirements.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-design-workflow`: selective secure-design planning replaces all-row compact/N/A enumeration in design.
- `sdd-review-security-workflow`: exhaustive applicability validation, non-applicable reporting, and missed-control blocking become explicit responsibilities.
- `sdd-security-guideline-catalog`: catalog remains taxonomy/source inventory while selective design and exhaustive review-security consume it with distinct boundaries.
- `sdd-test-design-workflow`: test design consumes applicable design controls and preserves review-security as the exhaustive gate.
- `sdd-execution-persistence-contracts`: active security artifact semantics stay aligned with the new design/review-security boundary.

## Approach

Adopt selective design plus exhaustive review-security. Design must justify its classification and record enough catalog/context references for audit, but it does not perform N/A bookkeeping for every compact row or Source ID. Review-security expands the full matrix, validates evidence and omissions, reports non-applicable rows, and blocks unsafe or missed applicable controls.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-design/SKILL.md` | Modified | Selective secure-design contract. |
| `skills/_shared/sdd-security-contract.md` | Modified | Shared ownership boundary. |
| `skills/sdd-review-security/SKILL.md` | Modified | Exhaustive validation authority. |
| `skills/sdd-test-design/SKILL.md` | Modified | Consume applicable controls only. |
| `openspec/specs/*/spec.md` | Modified | Source capability requirements. |
| `agents/sdd/*.md` | Modified | Adapter prompt alignment. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Design becomes too terse and misses an applicable control | Medium | Require changed-surface classification; review-security blocks missed applicable controls. |
| Prompt/spec drift | Medium | Update skills, shared contract, specs, and adapters in the same change. |
| Review-security reports grow | Medium | Accept growth in review artifact, not design artifact. |

## Rollback Plan

Revert the proposal/spec/skill/prompt changes for this OpenSpec change. Archived artifacts remain untouched, and existing exhaustive designs stay readable as compatibility data.

## Dependencies

- Existing compact security catalog and Source ID inventory remain authoritative.
- `sdd-review-security` must continue to run after non-blocking general review and before verify.

## Success Criteria

- [ ] New designs include only applicable secure-design categories/controls plus classification context.
- [ ] Design no longer requires all eight compact controls or all non-applicable Source IDs.
- [ ] Security review exhaustively reports applicable/non-applicable compact and Source ID rows.
- [ ] Security review blocks when design omitted an applicable category/control.
- [ ] Skills, shared contract, specs, and adapter prompts state the same boundary.
