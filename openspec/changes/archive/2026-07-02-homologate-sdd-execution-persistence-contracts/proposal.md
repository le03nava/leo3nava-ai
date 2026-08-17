# Proposal: Homologate SDD Execution and Persistence Contracts

## Intent

Define one authoritative boundary between SDD phase execution rules and artifact persistence rules. Today, shared contracts and phase skills repeat mode semantics with slightly different wording, which makes behavior hard to verify and easy to drift during future SDD changes.

## Scope

### In Scope
- Make the shared persistence contract authoritative for artifact-store modes, reads/writes, hybrid conflict handling, artifact resolution, and persistence verification.
- Keep `sdd-phase-common` focused on executor boundary, skill loading, return envelope, routing tokens, naming reminders, and review workload guard.
- Replace duplicated phase-local persistence prose with compact phase artifact contracts that preserve required inputs, produced artifacts, keys, paths, mutations, and next routing.
- Preserve existing SDD DAG, artifact keys, OpenSpec paths, routing tokens, status fields, and backend behavior unless current rules are ambiguous, duplicated, or conflicting.

### Out of Scope
- Redesigning the SDD DAG, security workflow, test-design workflow, review workload guard, or native status schema.
- Changing persistence backends, default artifact store behavior, or OpenSpec/Engram key formats.
- Implementing non-Markdown runtime code or editing installed adapter copies as source of truth.

## Capabilities

### New Capabilities
- `sdd-execution-persistence-contracts`: shared SDD contract boundaries for execution responsibilities, persistence responsibilities, artifact resolution, and phase-local artifact contracts.

### Modified Capabilities
- `sdd-security-applicability-workflow`: align artifact/routing contract language with the authoritative persistence boundary.
- `sdd-security-design-workflow`: preserve conditional artifact behavior while referencing shared persistence authority.
- `sdd-test-design-workflow`: preserve mandatory artifact and downstream consumption behavior while referencing shared persistence authority.

## Approach

Use `skills/_shared/persistence-contract.md` as the single detailed persistence authority. Convert repeated mode rules in phase skills into compact artifact contract tables. Keep phase-specific validation, artifact formats, mutations, and routing local. Update only directly related shared/reference wording that would otherwise contradict the new boundary.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared/persistence-contract.md` | Modified | Authoritative persistence and resolver contract. |
| `skills/_shared/sdd-phase-common.md` | Modified | Execution/envelope glue only; persistence sections become references. |
| `skills/sdd-*/SKILL.md` | Modified | Replace duplicated persistence prose with phase artifact contracts. |
| `skills/_shared/engram-convention.md` | Modified | Remove wording that competes with persistence authority. |
| `openspec/specs/sdd-*` | Modified/New | Specify behavior-level contract boundaries and compatibility requirements. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Behavior drift during partial migration | Med | Use chained slices and verify old/new contract equivalence. |
| Over-centralizing phase-specific mutations | Med | Keep apply/archive/init mutation semantics local. |
| Review overload | High | Split implementation into stacked reviewable work units. |

## Rollback Plan

Revert the Markdown contract changes for the affected slice. Because artifact keys, paths, and backends are preserved, rollback does not require data migration.

## Dependencies

- Existing OpenSpec specs and exploration artifact for this change.

## Success Criteria

- [ ] One shared contract owns detailed persistence behavior.
- [ ] Phase skills retain explicit artifact contracts without duplicating mode semantics.
- [ ] Existing artifact keys, paths, routing tokens, and conditional security/test-design behavior remain unchanged.
- [ ] Specs cover contract boundaries, compatibility preservation, and ambiguity resolution.
