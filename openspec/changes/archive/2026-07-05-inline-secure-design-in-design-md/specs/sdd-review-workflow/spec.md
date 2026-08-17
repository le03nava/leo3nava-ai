# Delta for sdd-review-workflow

## MODIFIED Requirements

### Requirement: Mandatory Review Gate

The SDD workflow MUST route completed implementation through `sdd-review` before `sdd-review-security` and `sdd-verify`. Review MUST inspect applied changes, required new-change SDD artifacts, and task evidence, then route non-blocking results to security review rather than directly to verify. `security-applicability.md` and standalone `security-design.md` MUST NOT be default required new-change inputs.
(Previously: new-change review inputs included security design as a standalone artifact.)

#### Scenario: Apply routes to review

- GIVEN all apply tasks are complete
- WHEN apply returns a success route
- THEN the next phase MUST be `sdd-review`
- AND verify MUST NOT run until both review gates are non-blocking.

#### Scenario: Blocking review returns to apply

- GIVEN review finds any blocking failure
- WHEN the review phase completes
- THEN it MUST route to `sdd-apply`
- AND `sdd-review-security` MUST NOT run yet.

#### Scenario: Legacy applicability is optional evidence

- GIVEN a new change has proposal, specs, design with embedded secure development rows, test design, tasks, and apply evidence
- WHEN review resolves required inputs
- THEN missing `security-applicability.md` or standalone `security-design.md` MUST NOT block review
- AND any such artifact present MUST be treated only as legacy or archive evidence.

### Requirement: Security Boundary

Security-related review controls MUST cross-reference the security guideline catalog where applicable. Review MUST NOT replace embedded secure development design or mandatory security review ownership. Legacy applicability or standalone security-design evidence MAY be cited only for archived or old changes and MUST NOT displace `design.md` or `review-security-report.md` as new-change security authorities.
(Previously: `security-design.md` and `review-security-report.md` were the new-change security authorities.)

#### Scenario: Security control is reviewed

- GIVEN a review control maps to a catalog guideline
- WHEN review records the matrix row
- THEN Standard SHOULD cite the guideline identifier
- AND embedded `design.md` security rows and `review-security-report.md` MUST remain the security authorities.

#### Scenario: Applicability does not replace security evidence

- GIVEN a reviewer finds a legacy `security-applicability.md` or `security-design.md` artifact
- WHEN reviewing a new change
- THEN the reviewer MAY cite it only as compatibility context
- AND MUST require embedded `design.md` secure development rows for authoritative classification.

### Requirement: Security Review Handoff

`review-report.md` MUST provide enough changed-file, evidence, finding context, and `design.md` references for `sdd-review-security` to validate embedded secure development implementation evidence without duplicating the 96-control matrix.
(Previously: handoff targeted the security matrix without requiring `design.md` as the embedded source.)

#### Scenario: Handoff evidence is available

- GIVEN review completed without blockers
- WHEN security review starts
- THEN it MUST be able to read `review-report.md` and `design.md`
- AND missing readable review or design evidence MUST route to `resolve-blockers`.
