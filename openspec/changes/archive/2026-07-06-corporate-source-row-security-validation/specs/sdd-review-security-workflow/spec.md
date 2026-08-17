# Delta for sdd-review-security-workflow

## ADDED Requirements

### Requirement: Exhaustive Source Row Security Review

`sdd-review-security` MUST validate every corporate source row against catalog inventory, `design.md#secure-development-design`, `test-design.md`, apply evidence, changed files, and `review-report.md`. It MUST generate a full security source-row matrix without duplicating the 96-control general review matrix.

#### Scenario: Full matrix is generated

- GIVEN design, test-design, apply evidence, changed files, and review report are readable
- WHEN security review succeeds
- THEN `review-security-report.md` MUST include every expected Source ID exactly once
- AND each row MUST show mapping, status, evidence, observations, and finding.

#### Scenario: General review is cited, not duplicated

- GIVEN a general review finding supports a source row
- WHEN security review records evidence
- THEN it MAY cite the review-report row
- AND it MUST NOT reproduce the full 96-control matrix.

### Requirement: Source Row Blocking Rules

Security review MUST block missing, duplicate, or unknown Source IDs; missing compact mappings; malformed source-row schema; missing design, test-design, apply, changed-file, or review artifacts; unsafe evidence; and missing `N/A` evidence or justification.

#### Scenario: Coverage or schema blocker exists

- GIVEN source-row coverage is incomplete or malformed
- WHEN security review validates rows
- THEN the verdict MUST be blocking
- AND next recommendation MUST be `resolve-blockers`.

#### Scenario: Implementation evidence is missing

- GIVEN a source row applies and remediation requires code or instruction changes
- WHEN review finds missing implementation evidence
- THEN the row MUST be blocking
- AND next recommendation MUST be `apply`.

#### Scenario: Warnings only remain

- GIVEN all mandatory rows have safe evidence and only warnings remain
- WHEN security review computes routing
- THEN the verdict MAY be non-blocking
- AND next recommendation MUST be `verify`.

### Requirement: Source Row Evidence Correlation

Security review MUST correlate source rows with design expectations, test-design checks, apply evidence, changed files, and review findings. A row MUST NOT pass solely because it is listed; evidence must support applicability, compliance, or justified `N/A`.

#### Scenario: Listed row has no corroboration

- GIVEN a source row appears in the matrix
- WHEN no supporting design, test, apply, changed-file, review, or exception evidence exists
- THEN the row MUST fail validation
- AND the report MUST identify the missing corroboration.

#### Scenario: Evidence is unsafe

- GIVEN evidence contains secret-like or confidential values
- WHEN security review validates the row
- THEN it MUST reject the evidence as unsafe
- AND route to `resolve-blockers` unless implementation remediation is required.
