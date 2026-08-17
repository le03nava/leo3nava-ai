# Delta for sdd-design-workflow

## ADDED Requirements

### Requirement: Secure Design Source ID Coverage

`sdd-design` MUST require `design.md#secure-development-design` to declare expected corporate Source ID coverage alongside compact `SEC-*` rows. The section MUST show traceability from Source ID to mapped compact guideline, expected evidence, applicability, lifecycle status, residual risk, and any planned exception or `N/A` justification.

#### Scenario: Design declares expected coverage

- GIVEN proposal and specs require security validation
- WHEN `sdd-design` succeeds
- THEN `design.md#secure-development-design` MUST include expected Source ID coverage
- AND each listed Source ID MUST reference one or more compact `SEC-*` IDs.

#### Scenario: Applicable row has evidence plan

- GIVEN a Source ID applies to the change
- WHEN design records the row
- THEN it MUST name expected downstream evidence and owner
- AND unresolved mandatory evidence MUST remain visible to test-design and review-security.

#### Scenario: N/A row has design rationale

- GIVEN a Source ID is expected to be not applicable
- WHEN design records the row
- THEN it MUST include evidence and justification for irrelevance
- AND missing `N/A` evidence MUST block design completion.

### Requirement: Design Preserves Compact Summary

`sdd-design` MUST keep the compact eight-control summary readable while adding source-row detail as a bounded operational matrix. It MUST NOT duplicate the general 96-control review matrix or replace compact guidelines with per-source controls.

#### Scenario: Compact and source layers coexist

- GIVEN a change has multiple corporate Source IDs
- WHEN reviewers inspect secure design
- THEN the compact `SEC-*` summary MUST remain visible
- AND detailed source rows MUST be traceable below it.

#### Scenario: General review matrix is not copied

- GIVEN secure design documents source coverage
- WHEN design is persisted
- THEN it MUST avoid reproducing the 96-control review matrix
- AND it MUST focus on security source-row planning only.
