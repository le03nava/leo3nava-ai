# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Apply Review Verify Routing

The SDD DAG for new changes MUST route `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`. Apply success MUST recommend review, non-blocking review MUST recommend review-security, non-blocking security review MUST recommend verify, and blocking findings MUST route to apply or resolve-blockers.
(Previously: this requirement only defined the post-apply review/security-review/verify route.)

#### Scenario: Mandatory review route is enforced

- GIVEN implementation has completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of apply
- AND review MUST be the required successor.

#### Scenario: Mandatory security review route is enforced

- GIVEN implementation and general review have completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of review
- AND `review-security` MUST be the required successor.

#### Scenario: Design routes directly to test design

- GIVEN `design.md` includes `## Secure Development Design`
- WHEN phase routing is evaluated
- THEN `test-design` MUST be the direct successor
- AND `security-design` MUST NOT be an active new-change successor.

#### Scenario: Review cannot safely run

- GIVEN required artifacts or changed-file context are missing
- WHEN review evaluates readiness
- THEN it MUST return `resolve-blockers`
- AND it MUST state the missing or unsafe input.

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include `design.md` with embedded secure development rows and `review-security-report.md` refs, paths, dependency states, native/status token `review-security`, and archive gates. `security-design.md` and `security-applicability.md` MAY appear only as legacy archived data refs and MUST NOT be active dependencies, produced artifacts, or phase-launch inputs.
(Previously: new changes required mandatory `security-design.md` refs and security-applicability was legacy-only.)

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include design and security review report slots
- AND active dependencies MUST NOT include `security-design` or `security-applicability`.

#### Scenario: Legacy refs are preserved as data

- GIVEN an archived change contains `artifactRefs.securityDesign` or `artifactRefs.securityApplicability`
- WHEN status or continuation displays historical evidence
- THEN the refs MAY remain visible as read-only data
- AND continuation MUST route active work through `design` instead of security-design or applicability.

## ADDED Requirements

### Requirement: Active Security Validator Removal

New-change contracts MUST NOT require `scripts/validate_security_design.ps1`. If retained for archived artifacts, references MUST be explicitly marked legacy-only and MUST NOT participate in active status, continuation, verify, or archive gating.

#### Scenario: Validator absence does not block

- GIVEN the validator script is absent or retired
- WHEN a new change reaches review-security, verify, or archive
- THEN the workflow MUST use catalog and artifact evidence instead
- AND absence of the script MUST NOT be a blocker.
