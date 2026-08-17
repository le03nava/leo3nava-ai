# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: Mandatory Phase Order

The SDD workflow MUST run `sdd-test-design` after both `sdd-design` and mandatory `sdd-security-design` succeed for every new change. `sdd-tasks` MUST be blocked until `test-design.md` exists or the phase explicitly blocks.
(Previously: `sdd-test-design` waited for security design only for security-impacting changes.)

#### Scenario: Design routes through mandatory security design

- GIVEN an SDD change has completed `sdd-design`
- WHEN the phase returns a successful envelope
- THEN the next recommended phase MUST be `security-design` or `sdd-security-design`
- AND `test-design` MUST NOT be the direct successor.

#### Scenario: Test design requires security design

- GIVEN `security-design.md` is missing for a new active change
- WHEN `sdd-test-design` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing mandatory security-design artifact.

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-design.md` for every change. The artifact MUST map spec scenarios, design risks, and the mandatory `security-design.md` matrix to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document justified no-impact assessments. Common persistence behavior MUST remain delegated to the shared persistence authority.
(Previously: security-design controls were required only when a conditional artifact existed.)

#### Scenario: Security matrix is consumed

- GIVEN `security-design.md` lists guideline rows and lifecycle statuses
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST include checks or justified non-test evidence for applicable mandatory rows
- AND blocked security rows MUST remain blockers.

#### Scenario: No-impact matrix rows are handled

- GIVEN all security rows are `not-applicable` with evidence
- WHEN `sdd-test-design` runs
- THEN it MUST cite that assessment
- AND it MUST still produce `test-design.md`.

### Requirement: Downstream Consumption

Downstream phases MUST use `test-design.md` as the test-planning source of truth and MUST preserve links to mandatory security-design evidence. `sdd-verify` MUST compare evidence against both test cases and security evidence obligations.
(Previously: downstream consumption referenced required security design only conditionally.)

#### Scenario: Verify compares security evidence

- GIVEN test design references mandatory security-design evidence
- WHEN `sdd-verify` runs
- THEN it MUST evaluate matching evidence or justified exceptions
- AND missing mandatory evidence MUST remain blocking.
