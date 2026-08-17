# Delta for sdd-review-security-workflow

## MODIFIED Requirements

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist `review-security-report.md`. The report MUST own the machine-readable schema, exhaustive compact and Source ID matrices, `Yes`/`No`/`N/A` decisions, evidence, blockers, exceptions, missed-category validation, and next recommendation. It MAY parse narrative category rules from design, but MUST NOT require design YAML, schema, matrices, or machine-readable fields.
(Previously: the report validated embedded design planning without explicitly owning all schema/matrix duties.)

#### Scenario: Report is persisted

- GIVEN `design.md` and `review-report.md` are readable
- WHEN security review completes
- THEN `review-security-report.md` MUST be written and read back
- AND it MUST state a verdict using its own report schema.

### Requirement: Security Matrix Validation

Security review MUST expand the full compact catalog and validate every compact control and Source ID using `Yes`, `No`, or `N/A`, evidence, observations, and lifecycle status. It MUST decide/report non-applicable rows, compare the matrix against narrative design rules, validate missed categories, and block applicable omissions.
(Previously: security review validated embedded design rows and their `N/A` rationale.)

#### Scenario: Mandatory evidence is missing

- GIVEN an applicable mandatory guideline lacks evidence or approved exception
- WHEN security review validates the matrix
- THEN the report MUST mark the row `No` and `blocked`
- AND verify/archive MUST be blocked.

#### Scenario: Design omitted an applicable category

- GIVEN proposal, specs, changed files, or evidence show a category applies
- WHEN design did not include narrative rules for that category
- THEN security review MUST report a missed applicable category/control
- AND the verdict MUST be blocking.

### Requirement: Exhaustive Source Row Security Review

`sdd-review-security` MUST be the only active phase that materializes the exhaustive corporate source-row matrix for a new change. It MUST expand every expected Source ID exactly once and validate rows against catalog inventory, narrative design rules, `test-design.md`, apply evidence, changed files, and `review-report.md`. It MUST write the full matrix in `review-security-report.md` without duplicating the 96-control matrix.
(Previously: design was expected to carry compact applicability evidence.)

#### Scenario: Full matrix is generated

- GIVEN design, test-design, apply evidence, changed files, and review report are readable
- WHEN security review succeeds
- THEN `review-security-report.md` MUST include every expected Source ID exactly once
- AND each row MUST show mapping, status, evidence, observations, and finding.

### Requirement: Source Row Blocking Rules

Security review MUST block missing, duplicate, or unknown Source IDs; missing compact mappings; malformed report schema; missing artifacts; unsafe evidence; missing `N/A` justification in the report matrix; and missed applicable design categories/controls. Design MUST NOT be blocked for lacking YAML, schema, or matrix fields.
(Previously: missing N/A evidence in design could block before security review.)

#### Scenario: Coverage or schema blocker exists

- GIVEN report source-row coverage is incomplete or malformed
- WHEN security review validates rows
- THEN the verdict MUST be blocking
- AND next recommendation MUST be `resolve-blockers`.

### Requirement: Source Row Evidence Correlation

Security review MUST correlate source rows with narrative design rules, test-design checks, apply evidence, changed files, and review findings. A row MUST NOT pass solely because it is listed or omitted from design; evidence MUST support applicability, compliance, justified `N/A`, or approved exception.
(Previously: correlation assumed embedded design row coverage was exhaustive.)

#### Scenario: Listed row has no corroboration

- GIVEN a source row appears in the matrix
- WHEN no supporting design, test, apply, changed-file, review, or exception evidence exists
- THEN the row MUST fail validation
- AND the report MUST identify the missing corroboration.
