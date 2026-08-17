# sdd-review-security-workflow Specification

## Purpose

Define the mandatory post-review security evidence gate that validates the security-design matrix and produces `review-security-report.md` before verification.

## Requirements

### Requirement: Mandatory Security Review Gate

The SDD workflow MUST run `sdd-review-security` after non-blocking `sdd-review` and before `sdd-verify` for every new change.

#### Scenario: General review routes to security review

- GIVEN `review-report.md` has no blocking findings
- WHEN routing is computed
- THEN the next required phase MUST be `sdd-review-security`
- AND `sdd-verify` MUST remain blocked until security review evidence exists.

#### Scenario: General review is blocking

- GIVEN `sdd-review` reports blocking findings
- WHEN routing is computed
- THEN `sdd-review-security` MUST NOT run
- AND the workflow MUST route back to `sdd-apply` or `resolve-blockers`.

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist `review-security-report.md`. The report MUST include verdict, validated matrix rows, implementation evidence locations, observations, blocking findings, exceptions, and next recommendation.

#### Scenario: Report is persisted

- GIVEN `security-design.md` and `review-report.md` are readable
- WHEN security review completes
- THEN `review-security-report.md` MUST be written and read back
- AND it MUST state a blocking or non-blocking verdict.

### Requirement: Security Matrix Validation

Security review MUST validate every category/guideline row from `security-design.md` using `Yes`, `No`, or `N/A`, evidence location, observations, and lifecycle status: `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, or `blocked`.

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

### Requirement: Boundary with General Review

Security review MUST NOT replace `sdd-review` or duplicate the 96-control matrix. It MUST focus on security guideline evidence and MAY cite general review findings as supporting evidence.

#### Scenario: General review evidence is reused

- GIVEN a review row supports a security guideline
- WHEN security review records evidence
- THEN it MAY cite that review row
- AND it MUST keep the security verdict in `review-security-report.md`.
