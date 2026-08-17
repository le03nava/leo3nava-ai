# Delta for sdd-design-workflow

## MODIFIED Requirements

### Requirement: Secure Design Source ID Coverage

`sdd-design` MUST require `design.md#secure-development-design` to declare slim expected corporate Source ID coverage alongside compact `SEC-*` rows. The section MUST record catalog snapshot identity/path, `expectedSourceIdCount: 155`, section or group coverage references, compact mapping summary, applicability, evidence owners, lifecycle status, residual risk, and planned exception or `N/A` policy. It MUST NOT copy the exhaustive 155-row Source ID inventory when the catalog is available.
(Previously: design had to show expected Source ID coverage with row-level traceability, which allowed exhaustive inventory duplication.)

#### Scenario: Design declares expected coverage

- GIVEN proposal and specs require security validation
- WHEN `sdd-design` succeeds
- THEN `design.md#secure-development-design` MUST cite the catalog snapshot and `expectedSourceIdCount: 155`
- AND it MUST summarize coverage by section/group and compact `SEC-*` mappings.

#### Scenario: Applicable group has evidence plan

- GIVEN one or more source rows apply to the change
- WHEN design records the slim coverage summary
- THEN it MUST name expected downstream evidence and owners for the affected group or compact mapping
- AND unresolved mandatory evidence MUST remain visible to test-design and review-security.

#### Scenario: N/A coverage has design rationale

- GIVEN a section, group, or compact mapping is expected to be not applicable
- WHEN design records the coverage summary
- THEN it MUST include evidence and justification for irrelevance
- AND missing `N/A` evidence MUST block design completion.

### Requirement: Design Preserves Compact Summary

`sdd-design` MUST keep the compact eight-control summary readable while adding source-row coverage only by reference, count, grouped summary, mappings, and downstream obligations. It MUST NOT duplicate the general 96-control review matrix, replace compact guidelines with per-source controls, or materialize the exhaustive 155-row review-security matrix.
(Previously: design preserved the compact summary while adding source-row detail as a bounded operational matrix.)

#### Scenario: Compact and source layers coexist

- GIVEN a change has multiple corporate Source IDs
- WHEN reviewers inspect secure design
- THEN the compact `SEC-*` summary MUST remain visible
- AND source coverage MUST be traceable through catalog references and grouped summaries.

#### Scenario: Exhaustive matrices are not copied

- GIVEN secure design documents source coverage
- WHEN design is persisted
- THEN it MUST avoid reproducing the 96-control review matrix and the full 155-row source matrix
- AND it MUST point exhaustive validation ownership to `review-security-report.md`.
