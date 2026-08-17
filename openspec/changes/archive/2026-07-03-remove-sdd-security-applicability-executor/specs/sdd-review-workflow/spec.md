# Delta for sdd-review-workflow

## MODIFIED Requirements

### Requirement: Mandatory Review Gate

The SDD workflow MUST route completed implementation through `sdd-review` before `sdd-review-security` and `sdd-verify`. Review MUST inspect applied changes, required new-change SDD artifacts, and task evidence, then route non-blocking results to security review rather than directly to verify. `security-applicability.md` MUST NOT be a default required new-change input.
(Previously: the review gate required artifacts generally but did not explicitly exclude legacy applicability from default inputs.)

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

- GIVEN a new change has proposal, specs, design, security design, test design, tasks, and apply evidence
- WHEN review resolves required inputs
- THEN missing `security-applicability.md` MUST NOT block review
- AND any applicability artifact present MUST be treated only as legacy or archive evidence.

### Requirement: Security Boundary

Security-related review controls MUST cross-reference the security guideline catalog where applicable. Review MUST NOT replace mandatory security design or mandatory security review ownership. Legacy applicability evidence MAY be cited only for archived or old changes and MUST NOT displace `security-design.md` or `review-security-report.md` as new-change security authorities.
(Previously: security design and security review were authorities, but legacy applicability citation rules were not explicit.)

#### Scenario: Security control is reviewed

- GIVEN a review control maps to a catalog guideline
- WHEN review records the matrix row
- THEN Standard SHOULD cite the guideline identifier
- AND `security-design.md` and `review-security-report.md` MUST remain the security authorities.

#### Scenario: Applicability does not replace security evidence

- GIVEN a reviewer finds a legacy `security-applicability.md` artifact
- WHEN reviewing a new change
- THEN the reviewer MAY cite it only as compatibility context
- AND MUST require `security-design.md` for authoritative classification.
