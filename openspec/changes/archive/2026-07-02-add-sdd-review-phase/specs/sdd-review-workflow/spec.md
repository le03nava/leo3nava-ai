# sdd-review-workflow Specification

## Purpose

Define `sdd-review` as the mandatory post-apply code-review gate that produces `review-report.md` before verification.

## Requirements

### Requirement: Mandatory Review Gate

The SDD workflow MUST route completed implementation through `sdd-review` before `sdd-verify`. Review MUST inspect applied changes, required SDD artifacts, and task evidence.

#### Scenario: Apply routes to review

- GIVEN all apply tasks are complete
- WHEN apply returns a success route
- THEN the next phase MUST be `sdd-review`
- AND verify MUST NOT run until review is non-blocking.

#### Scenario: Blocking review returns to apply

- GIVEN review finds any blocking failure
- WHEN the review phase completes
- THEN it MUST route to `sdd-apply`
- AND it MUST identify the failed controls and affected requirements.

### Requirement: Review Report Artifact

`sdd-review` MUST persist `review-report.md` as a first-class artifact at `openspec/changes/{change-name}/review-report.md`. Missing required artifacts, unknown changed files, unsafe workspace context, or persistence failure MUST route to `resolve-blockers`.

#### Scenario: Report is persisted

- GIVEN review can resolve all required inputs
- WHEN review completes
- THEN it MUST write `review-report.md`
- AND the report MUST include verdict, matrix, evidence summary, and next recommendation.

### Requirement: Code-Review Validation Matrix

`review-report.md` MUST include a matrix with exactly these columns: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. The 96 source checklist controls MUST have stable IDs and categories. Category MUST be preserved by a stable control catalog or by the Item ID prefix without adding matrix columns. `Complies` MUST support Yes, No, and N/A. N/A MUST include evidence proving the control is irrelevant, especially for platform-specific controls.

#### Scenario: All controls are represented

- GIVEN the source checklist contains 96 controls
- WHEN review writes the matrix
- THEN every control MUST appear once with a stable Item ID
- AND every row MUST use the exact required columns.

#### Scenario: Platform control is irrelevant

- GIVEN a control targets an unused platform or technology
- WHEN review marks it N/A
- THEN Evidence Location MUST prove irrelevance
- AND Observations/Comments MUST explain the scope decision.

### Requirement: Severity and Routing Semantics

Review MUST classify each failed control severity. Critical or explicitly blocking failures MUST prevent verify. Non-blocking findings MAY proceed to verify as warnings.

#### Scenario: Non-blocking review proceeds

- GIVEN review has no blocking failures
- WHEN it returns its envelope
- THEN next_recommended MUST be `sdd-verify`
- AND verify MUST consume the review report as evidence.

### Requirement: Security Boundary

Security-related review controls MUST cross-reference the security guideline catalog where applicable. Review MUST NOT replace security applicability or security design ownership.

#### Scenario: Security control is reviewed

- GIVEN a review control maps to a catalog guideline
- WHEN review records the matrix row
- THEN Standard SHOULD cite the guideline identifier
- AND applicability/design remain the security authorities.
