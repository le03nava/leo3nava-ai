# Delta for sdd-review-workflow

## MODIFIED Requirements

### Requirement: Mandatory Review Gate

The SDD workflow MUST route completed implementation through `sdd-review` before `sdd-review-security` and `sdd-verify`. Review MUST inspect applied changes, required SDD artifacts, and task evidence, then route non-blocking results to security review rather than directly to verify.
(Previously: non-blocking review routed directly to `sdd-verify`.)

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

### Requirement: Severity and Routing Semantics

Review MUST classify each failed control severity. Critical or explicitly blocking failures MUST prevent security review and verify. Non-blocking findings MAY proceed to `sdd-review-security` as warnings.
(Previously: non-blocking findings proceeded to verify.)

#### Scenario: Non-blocking review proceeds

- GIVEN review has no blocking failures
- WHEN it returns its envelope
- THEN next_recommended MUST be `sdd-review-security`
- AND the security review phase MUST consume the review report as evidence.

### Requirement: Security Boundary

Security-related review controls MUST cross-reference the security guideline catalog where applicable. Review MUST NOT replace mandatory security design or mandatory security review ownership.
(Previously: review did not include a separate security review successor.)

#### Scenario: Security control is reviewed

- GIVEN a review control maps to a catalog guideline
- WHEN review records the matrix row
- THEN Standard SHOULD cite the guideline identifier
- AND `security-design.md` and `review-security-report.md` MUST remain the security authorities.

## ADDED Requirements

### Requirement: Security Review Handoff

`review-report.md` MUST provide enough changed-file, evidence, and finding context for `sdd-review-security` to validate security matrix implementation evidence without duplicating the 96-control matrix.

#### Scenario: Handoff evidence is available

- GIVEN review completed without blockers
- WHEN security review starts
- THEN it MUST be able to read `review-report.md`
- AND missing readable review evidence MUST route to `resolve-blockers`.
