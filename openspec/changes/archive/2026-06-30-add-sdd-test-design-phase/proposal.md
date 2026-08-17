# Proposal: Add Mandatory SDD Test Design Phase

## Intent

Add a mandatory `sdd-test-design` phase between `sdd-design` and `sdd-tasks` so every SDD change has explicit test-case planning before implementation tasks are created. This closes the current gap where concrete test coverage is spread across specs, design, tasks, apply evidence, and verification reports.

## Scope

### In Scope
- Add phase/agent `sdd-test-design` with artifact `test-design.md`.
- Recommended order: `explore? -> proposal -> spec -> design -> test-design -> tasks -> apply -> verify -> archive`.
- Require `test-design.md` for every change, including explicit no-impact assessments when behavior/testability is unaffected.
- Include automatable test cases plus manual/static checks when no runner exists or behavior cannot be automated.
- Define verification intent: uncovered mandatory cases fail `sdd-verify`; uncovered non-mandatory cases warn.

### Out of Scope
- Implementing the phase, agent, skill, dispatcher, or state changes now.
- Choosing final machine token syntax without spec/design validation.
- Adding a runtime test framework to this repository.

## Capabilities

### New Capabilities
- `sdd-test-design-workflow`: Defines mandatory test-design phase behavior, artifact expectations, downstream task usage, and verification coverage semantics.

### Modified Capabilities
- None. Current main specs cover opencode sync scripts, not the SDD phase workflow.

## Approach

Introduce `test-design.md` as the bridge between technical design and task planning. The artifact should map spec scenarios to planned automated, manual, or static checks; mark cases as mandatory or non-mandatory; document no-impact changes explicitly; and state expected evidence. Later `sdd-spec` and `sdd-design` must refine status/state/persistence contracts and bounded native dispatcher/status token support for `test-design` / `sdd-test-design`.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `agents/sdd/` | Modified/New | Orchestrator routing plus new executor prompt. |
| `skills/sdd-test-design/` | New | Phase skill and artifact contract. |
| `skills/sdd-design/`, `skills/sdd-tasks/`, `skills/sdd-apply/`, `skills/sdd-verify/` | Modified | Routing, dependencies, task derivation, and coverage checks. |
| `skills/_shared/` | Modified | OpenSpec, persistence, status, and phase-common contracts. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Native dispatcher/status token enum rejects `test-design` | Medium | Spec/design must confirm bounded token support and compatibility mapping before implementation. |
| State/persistence drift breaks continuation | Medium | Update state, artifact refs, routing, and recovery contracts together. |
| Test design becomes too large | Medium | Keep cases scenario-linked and evidence-focused. |

## Rollback Plan

Revert the proposal/spec/design/tasks and any later implementation files for the new phase. Restore routing from `sdd-design` directly to `sdd-tasks` and remove `test-design.md` artifact requirements from shared contracts.

## Dependencies

- Follow-up `sdd-spec` and `sdd-design` phases must define exact requirements, schemas, persistence refs, status tokens, and compatibility behavior.

## Success Criteria

- [ ] Proposal names `sdd-test-design`, `test-design.md`, and the correct phase position.
- [ ] Specs/design later define mandatory/no-impact behavior and fail-vs-warn verification semantics.
- [ ] Compatibility with native dispatcher/status tokens is explicitly resolved before implementation.
