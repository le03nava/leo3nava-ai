# Delta for sdd-review-security-workflow

## MODIFIED Requirements

### Requirement: Exhaustive Source Row Security Review

`sdd-review-security` MUST be the only active phase that materializes the exhaustive corporate source-row validation matrix for a new change. It MUST expand the catalog inventory to every expected Source ID exactly once and validate each row against catalog inventory, slim `design.md#secure-development-design`, `test-design.md`, apply evidence, changed files, and `review-report.md`. It MUST generate the full security source-row matrix in `review-security-report.md` without duplicating the 96-control general review matrix.
(Previously: security review generated the full source-row matrix, but exclusive ownership over materializing the 155 rows was not explicit.)

#### Scenario: Full matrix is generated

- GIVEN design, test-design, apply evidence, changed files, and review report are readable
- WHEN security review succeeds
- THEN `review-security-report.md` MUST include every expected Source ID exactly once
- AND each row MUST show mapping, status, evidence, observations, and finding.

#### Scenario: Design remains slim

- GIVEN design cites the catalog snapshot, expected count, grouped coverage, and compact mappings
- WHEN security review expands source rows
- THEN the exhaustive row matrix MUST be written only in `review-security-report.md`
- AND missing design references MUST block as contract evidence gaps.

#### Scenario: General review is cited, not duplicated

- GIVEN a general review finding supports a source row
- WHEN security review records evidence
- THEN it MAY cite the review-report row
- AND it MUST NOT reproduce the full 96-control matrix.
