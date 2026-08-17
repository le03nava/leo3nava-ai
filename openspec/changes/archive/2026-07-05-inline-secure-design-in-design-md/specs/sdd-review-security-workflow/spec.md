# Delta for sdd-review-security-workflow

## MODIFIED Requirements

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist `review-security-report.md`. The report MUST validate the embedded `design.md` secure development rows, include verdict, row-level evidence locations, observations, blocking findings, exceptions, and next recommendation, and MUST NOT depend on the active validator `scripts/validate_security_design.ps1`.
(Previously: the report required readable `security-design.md` as the security matrix source.)

#### Scenario: Report is persisted

- GIVEN `design.md` and `review-report.md` are readable
- WHEN security review completes
- THEN `review-security-report.md` MUST be written and read back
- AND it MUST state a blocking or non-blocking verdict.

#### Scenario: Legacy standalone artifact is read-only

- GIVEN a legacy archive has standalone `security-design.md`
- WHEN security review needs historical evidence
- THEN it MAY read that artifact as compatibility context
- AND it MUST NOT require that artifact for new changes.

### Requirement: Security Matrix Validation

Security review MUST validate every embedded `design.md` secure development row using `Yes`, `No`, or `N/A`, evidence location, observations, and lifecycle status: `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, or `blocked`.
(Previously: validation consumed rows from `security-design.md`.)

#### Scenario: Mandatory evidence is missing

- GIVEN an applicable mandatory guideline lacks implementation evidence or approved exception
- WHEN security review validates the matrix
- THEN the report MUST mark the row `No` and `blocked`
- AND verify/archive MUST be blocked.

#### Scenario: Not applicable row is justified

- GIVEN a guideline is marked `N/A`
- WHEN security review validates it
- THEN evidence MUST prove irrelevance
- AND observations MUST explain the scope decision.

## ADDED Requirements

### Requirement: Active Security Validator Retirement

The active new-change workflow MUST remove `scripts/validate_security_design.ps1` as a required validator. Any remaining references MUST be legacy/archive-only and MUST NOT block new-change routing, review-security, verify, or archive.

#### Scenario: New change does not invoke validator

- GIVEN `sdd-review-security` validates a new change
- WHEN it checks secure development evidence
- THEN it MUST validate against the corporate catalog and embedded `design.md` section
- AND it MUST NOT require `scripts/validate_security_design.ps1` execution.
