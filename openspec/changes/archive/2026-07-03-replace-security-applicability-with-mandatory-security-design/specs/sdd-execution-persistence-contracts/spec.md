# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Conflict and Ambiguity Resolution

When SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless an approved spec change explicitly redesigns DAG, artifact, routing, status, or persistence semantics. Explicit redesigns MUST define compatibility rules for old artifacts and MUST NOT silently invalidate archives.
(Previously: clarifications could not redesign the SDD DAG or artifact/status surface.)

#### Scenario: Explicit DAG redesign is applied

- GIVEN an approved spec changes phase order
- WHEN contracts are updated
- THEN the new DAG MUST be authoritative for new changes
- AND legacy artifacts MUST remain readable only under documented compatibility rules.

### Requirement: Review Phase Artifact Contract

The SDD contract set MUST define `review-report.md` as the first review artifact after apply and before `sdd-review-security`. State and status contracts MUST expose review refs, verdict, blocking-failure state, and routing to security review when non-blocking.
(Previously: review was the only review artifact before verify.)

#### Scenario: Review artifact is resolved

- GIVEN `sdd-review-security` needs review evidence
- WHEN it resolves artifacts for a change
- THEN it MUST find `review-report.md` or the matching backend key
- AND missing review evidence MUST block security review, verify, and archive.

### Requirement: Apply Review Verify Routing

The SDD DAG for new changes MUST route `apply -> review -> review-security -> verify -> archive`. Apply success MUST recommend review, non-blocking review MUST recommend review-security, non-blocking security review MUST recommend verify, and blocking findings MUST route to apply or resolve-blockers.
(Previously: the DAG routed `apply -> review -> verify -> archive`.)

#### Scenario: Mandatory security review route is enforced

- GIVEN implementation and general review have completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of review
- AND `review-security` MUST be the required successor.

### Requirement: Verify and Archive Review Consumption

Verify MUST consume both `review-report.md` and `review-security-report.md` as evidence and MUST NOT own either review matrix. Archive MUST require passing verification plus non-blocking general and security review reports for new changes.
(Previously: verify/archive required only non-blocking general review.)

#### Scenario: Verify consumes both review artifacts

- GIVEN both review reports are non-blocking
- WHEN verify runs
- THEN it MUST cite both reports as evidence
- AND it MUST NOT duplicate their full matrices.

## ADDED Requirements

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include mandatory `security-design.md` and `review-security-report.md` refs, paths, dependency states, native/status token `review-security`, and archive gates. `security-applicability.md` MAY appear only as a legacy archived artifact.

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include security design and security review report slots
- AND active dependencies MUST NOT include `security-applicability`.
