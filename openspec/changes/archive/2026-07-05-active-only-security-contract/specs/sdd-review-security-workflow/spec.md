# Delta for sdd-review-security-workflow

## MODIFIED Requirements

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist `review-security-report.md`. The report MUST validate embedded `design.md#secure-development-design` rows, include verdict, row-level evidence locations, observations, blocking findings, exceptions, and next recommendation, and MUST NOT depend on `scripts/validate_security_design.ps1` or standalone security artifacts for active new-change evidence.
(Previously: the requirement allowed a standalone security-design read scenario inside the active workflow.)

#### Scenario: Report is persisted

- GIVEN `design.md` and `review-report.md` are readable
- WHEN security review completes
- THEN `review-security-report.md` MUST be written and read back
- AND it MUST state a blocking or non-blocking verdict.

#### Scenario: Embedded secure design is required

- GIVEN a new change lacks `design.md#secure-development-design`
- WHEN security review evaluates readiness
- THEN it MUST block with missing embedded design evidence
- AND verify/archive MUST remain unavailable.

### Requirement: Active Security Validator Retirement

The active new-change workflow MUST remove `scripts/validate_security_design.ps1` as a required validator. Any remaining parser or archive references MUST NOT block new-change routing, review-security, verify, or archive.
(Previously: retained references were framed as legacy/archive-only validator references.)

#### Scenario: New change does not invoke validator

- GIVEN `sdd-review-security` validates a new change
- WHEN it checks secure development evidence
- THEN it MUST validate against the corporate catalog and embedded `design.md` section
- AND it MUST NOT require `scripts/validate_security_design.ps1` execution.
