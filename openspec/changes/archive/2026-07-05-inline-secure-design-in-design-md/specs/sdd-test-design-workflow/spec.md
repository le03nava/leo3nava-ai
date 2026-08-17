# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: Mandatory Phase Order

The SDD workflow MUST run `sdd-test-design` after `sdd-design` succeeds for every new change. `sdd-design` MUST include embedded secure development design, and `sdd-tasks` MUST be blocked until `test-design.md` exists or the phase explicitly blocks.
(Previously: `sdd-test-design` required both `sdd-design` and mandatory `sdd-security-design`.)

#### Scenario: Design routes through mandatory security design

- GIVEN an SDD change has completed `sdd-design`
- WHEN the phase returns a successful envelope
- THEN the next recommended phase MUST be `test-design` or `sdd-test-design`
- AND `security-design` MUST NOT be the direct successor for new changes.

#### Scenario: Tasks requested too early

- GIVEN `test-design.md` is missing for an active change
- WHEN `sdd-tasks` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing `sdd-test-design` phase or artifact.

#### Scenario: Test design requires security design

- GIVEN `design.md` lacks `## Secure Development Design` for a new active change
- WHEN `sdd-test-design` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing embedded secure design section.

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-design.md` for every change. The artifact MUST map spec scenarios, design risks, and embedded `design.md` secure development rows to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document justified no-impact assessments.
(Previously: the artifact consumed the mandatory `security-design.md` matrix.)

#### Scenario: Behavior-impacting change

- GIVEN specs or design describe behavior, contracts, routing, or compatibility changes
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST list planned checks linked to those inputs
- AND each check MUST include type, severity, and expected evidence.

#### Scenario: Security matrix is consumed

- GIVEN `design.md` lists secure development rows and lifecycle statuses
- WHEN `sdd-test-design` runs
- THEN it MUST include checks or justified non-test evidence for applicable mandatory rows
- AND blocked security rows MUST remain blockers.

#### Scenario: No-impact matrix rows are handled

- GIVEN all security rows are `not-applicable` with evidence
- WHEN `sdd-test-design` runs
- THEN it MUST cite that assessment
- AND it MUST still produce `test-design.md`.

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
