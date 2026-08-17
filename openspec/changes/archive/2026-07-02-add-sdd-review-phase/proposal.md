# Proposal: Add SDD Review Phase

## Intent

Add `sdd-review` as a mandatory code-review gate after implementation and before verification. The current workflow can prove specs and tasks, but it lacks a first-class review artifact that records per-control validation for the requested 96-item checklist.

## Scope

### In Scope
- Add mandatory DAG order: `apply -> review -> verify -> archive`.
- Create `sdd-review` phase contracts, executor prompt, skill, status routing, and persistence/OpenSpec artifact handling.
- Produce `review-report.md` as a first-class OpenSpec artifact containing the exact matrix columns: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments.
- Define severity/routing semantics: blocking review failures return to `apply`; non-blocking reports proceed to `verify`.
- Make `sdd-verify` consume `review-report.md` without owning the full review matrix.
- Make `sdd-archive` require a passing/non-blocking review report plus passing verify report.

### Out of Scope
- Implementing application runtime tests or language-specific analyzers.
- Replacing security applicability/design ownership; review may reference security controls but does not become the security authority.
- Changing the 96 controls beyond normalizing them into stable IDs and categories.

## Capabilities

### New Capabilities
- `sdd-review-workflow`: Defines the review phase, review-report artifact, matrix contract, severity semantics, and apply/verify/archive integration.

### Modified Capabilities
- `sdd-execution-persistence-contracts`: Adds `review-report.md`, review routing, state/status fields, and archive readiness requirements.
- `sdd-security-guideline-catalog`: Cross-references applicable checklist security controls without duplicating catalog authority.

## Approach

Introduce `sdd-review` as a separate executor and skill. Update shared SDD contracts so apply routes to review, verify reads review results as evidence, and archive gates on both review and verification. Store review output at `openspec/changes/{change-name}/review-report.md`.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `agents/sdd/` | New/Modified | Add review executor and route orchestration. |
| `skills/sdd-review/` | New | Define review inputs, matrix, routing, and persistence. |
| `skills/_shared/` | Modified | Add review artifact/status/OpenSpec contracts. |
| `skills/sdd-apply`, `skills/sdd-verify`, `skills/sdd-archive` | Modified | Integrate review gate. |
| `openspec/specs/` | Modified/New | Add review workflow spec and modify shared SDD contracts. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Broad DAG contract change exceeds review budget | High | Split into stacked work units. |
| Checklist overlaps security workflow | Medium | Cross-reference security catalog and preserve ownership boundaries. |
| Platform-specific controls create false failures | Medium | Require `N/A` with evidence for irrelevant controls. |

## Rollback Plan

Revert the review phase artifacts and contract changes, restore `apply -> verify` routing, remove `review-report.md` requirements from status/OpenSpec/archive readiness, and leave existing verify/archive behavior unchanged.

## Dependencies

- Existing OpenSpec SDD execution/persistence and security guideline catalog specs.
- Full 96-control checklist captured during specs with stable IDs.

## Success Criteria

- [ ] `sdd-review` is a mandatory gate between apply and verify.
- [ ] `review-report.md` is produced with the exact requested matrix columns.
- [ ] Verify consumes review results without duplicating the matrix.
- [ ] Archive requires non-blocking review and passing verification.
