# Proposal: Design-Driven Operational Considerations

## Intent

Replace the mandatory operational-readiness contract model with a design-driven model. SDD phases should plan, test, review, verify, and archive operational evidence only when the design, test-design, tasks, or archived artifacts make it applicable.

## Scope

### In Scope
- Delete `skills/_shared/sdd-operational-readiness-contract.md` completely.
- Remove active phase references to the deleted contract and mandatory cross-phase readiness completeness gates.
- Update `sdd-design` to include conditional `## Operational Considerations` for applicable logs/errors, monitoring, administration operations, reprocessing/recovery, and backup/retention/cleanup/generated-artifact concerns.
- Update downstream SDD phases and shared post-apply/security contracts to consume actual design/test-design/tasks/apply/archive evidence without requiring absent categories.
- Update active OpenSpec specs that encode mandatory readiness evaluation.

### Out of Scope
- Preserving the deleted file for legacy or archive compatibility.
- Rewriting archived change artifacts that already reference the old contract.
- Removing safe-evidence or restricted-data safeguards that still apply to operational evidence.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-design-workflow`: make operational considerations conditional and design-owned.
- `sdd-test-design-workflow`: derive operational checks only from design content.
- `sdd-review-workflow`: review present/planned operational evidence without mandatory category completeness.
- `sdd-review-security-workflow`: keep leakage/safe-evidence validation for existing operational evidence only.
- `sdd-execution-persistence-contracts`: persist, verify, and archive operational evidence that exists, not contract completeness.
- `sdd-operational-readiness-workflow`: replace mandatory readiness workflow with design-driven operational considerations or retire its mandatory requirements.
- `sdd-security-guideline-catalog`: preserve restricted-data safeguards without implying mandatory operational categories.

## Approach

Remove the shared contract file and all active references. Reword phase skills and shared contracts so `design.md`, `test-design.md`, `tasks.md`, apply evidence, and archive evidence are the authority. `sdd-operational-doc` remains post-archive and reports missing applicable data as pending or inapplicable as `No aplica.` without inventing data.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared` | Removed/Modified | Delete readiness contract; adjust post-apply/security wording. |
| `skills/sdd-*` | Modified | Remove contract-driven operational readiness requirements. |
| `openspec/specs/*/spec.md` | Modified | Remove mandatory readiness evaluation requirements. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Stale references keep old gates alive | Med | Search all skills/specs for contract and readiness-completeness language. |
| Over-removal weakens safe-evidence controls | Med | Keep restricted-data safeguards in security contracts. |
| Archived artifacts reference deleted file | Low | Treat archived refs as historical only; do not preserve active compatibility file. |

## Rollback Plan

Revert the change files and restore `skills/_shared/sdd-operational-readiness-contract.md` from git history if the design-driven model breaks SDD phase execution.

## Dependencies

- Existing exploration artifact: `openspec/changes/design-driven-operational-considerations/explore.md`.
- User decisions in the SDD launch envelope.

## Success Criteria

- [ ] No active phase skill references `skills/_shared/sdd-operational-readiness-contract.md`.
- [ ] `sdd-design` owns conditional operational considerations.
- [ ] Downstream phases consume actual artifacts instead of mandatory readiness categories.
- [ ] OpenSpec specs no longer require mandatory operational-readiness evaluation.
