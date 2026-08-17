# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-design.md` for every change. The artifact MUST map spec scenarios, design risks, and required security-design controls to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document no-impact assessments when no behavior, security, or testability impact exists. The phase artifact contract MUST preserve mandatory artifact creation and downstream consumption while delegating common artifact-store mode semantics, artifact resolution, and persistence verification to the shared persistence authority.
(Previously: The requirement defined the mandatory test-design artifact content, but did not state the shared persistence boundary.)

#### Scenario: Behavior-impacting change

- GIVEN specs or design describe behavior, contracts, routing, or compatibility changes
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST list planned checks linked to those inputs
- AND each check MUST include type, severity, and expected evidence.

#### Scenario: Security-impacting change

- GIVEN `security-design.md` lists required controls or mandatory evidence
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST include planned checks or justified non-test evidence for each required control
- AND uncovered mandatory security evidence MUST be reported as a blocker.

#### Scenario: No-impact change

- GIVEN the change has no behavior, security, or testability impact
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST state a no-impact assessment
- AND downstream phases MUST treat the artifact as complete rather than absent.

#### Scenario: Persistence boundary is delegated

- GIVEN `sdd-test-design` writes or resolves `test-design.md` in any supported artifact-store mode
- WHEN backend behavior is required
- THEN it MUST follow the shared persistence authority
- AND it MUST keep the mandatory artifact contract local to the phase.

### Requirement: Downstream Consumption

Downstream phases MUST use `test-design.md` as the test-planning source of truth. `sdd-tasks` MUST derive implementation and verification tasks from it; `sdd-apply` SHOULD follow planned cases or report justified deviations; `sdd-verify` MUST compare evidence against it. Updates to persistence contract wording MUST NOT weaken this downstream consumption requirement or route tasks directly from design without the required test-design artifact.
(Previously: The requirement defined downstream consumption, but did not explicitly protect it from persistence-contract refactoring.)

#### Scenario: Tasks derive testing work

- GIVEN `test-design.md` contains planned checks
- WHEN `sdd-tasks` creates implementation tasks
- THEN tasks MUST include corresponding testing or evidence work
- AND omitted mandatory cases MUST be reported as blockers

#### Scenario: Apply deviates from plan

- GIVEN implementation cannot follow a planned case exactly
- WHEN `sdd-apply` records evidence
- THEN it SHOULD document the deviation and replacement evidence
- AND `sdd-verify` MUST evaluate that justification against the planned case

#### Scenario: Persistence refactor preserves consumption

- GIVEN common persistence rules move to a shared authority
- WHEN downstream phases resolve test-planning inputs
- THEN they MUST still consume `test-design.md` from the established key or path
- AND they MUST NOT treat the artifact as optional.
