# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: Mandatory Phase Order

The SDD workflow MUST run `sdd-test-design` after `sdd-design` succeeds and after required `sdd-security-design` succeeds for security-impacting changes. `sdd-tasks` MUST be blocked until the required test-design artifact exists or the phase is explicitly blocked.
(Previously: `sdd-test-design` ran directly after `sdd-design` and before `sdd-tasks` without security-design dependency handling.)

#### Scenario: Successful design routes to security design or test design

- GIVEN an SDD change has completed `sdd-design`
- WHEN the phase returns a successful envelope
- THEN the next recommended phase MUST be `security-design` or `sdd-security-design` when security applicability is impacting
- AND otherwise the next recommended phase MUST be `test-design` or `sdd-test-design`.

#### Scenario: Tasks requested too early

- GIVEN `test-design.md` is missing for an active change
- WHEN `sdd-tasks` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing `sdd-test-design` phase or artifact.

#### Scenario: Security design required before test design

- GIVEN security applicability marks a change as security-impacting and `security-design.md` is missing
- WHEN `sdd-test-design` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing `sdd-security-design` phase or artifact.

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-design.md` for every change. The artifact MUST map spec scenarios, design risks, and required security-design controls to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document no-impact assessments when no behavior, security, or testability impact exists.
(Previously: the artifact mapped specs and design risks but did not require security-control coverage.)

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
