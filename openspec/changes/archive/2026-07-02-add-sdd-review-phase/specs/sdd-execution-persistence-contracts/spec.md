# Delta for sdd-execution-persistence-contracts

## ADDED Requirements

### Requirement: Review Phase Artifact Contract

The SDD contract set MUST define `review-report.md` as the first-class artifact produced by `sdd-review` after apply and before verify. OpenSpec mode MUST store it at `openspec/changes/{change-name}/review-report.md`; Engram or hybrid modes MUST use the stable artifact key `sdd/{change-name}/review`. State and status contracts MUST expose review artifact refs, paths, verdict, blocking-failure state, and `nextRecommended` routing.

#### Scenario: Review artifact is resolved

- GIVEN a downstream phase needs review evidence
- WHEN it resolves artifacts for a change
- THEN it MUST find `review-report.md` or the matching backend artifact key
- AND missing review evidence MUST block verify or archive.

### Requirement: Apply Review Verify Routing

The SDD DAG MUST route `apply -> review -> verify -> archive`. Apply success MUST recommend review, review success without blocking failures MUST recommend verify, blocking review failures MUST recommend apply, and unsafe or incomplete review context MUST recommend `resolve-blockers`.

#### Scenario: Mandatory review route is enforced

- GIVEN implementation has completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of apply
- AND review MUST be the required successor.

#### Scenario: Review cannot safely run

- GIVEN required artifacts or changed-file context are missing
- WHEN review evaluates readiness
- THEN it MUST return `resolve-blockers`
- AND it MUST state the missing or unsafe input.

### Requirement: Verify and Archive Review Consumption

Verify MUST consume `review-report.md` as evidence and MUST NOT own the full 96-control review matrix. Archive MUST require both non-blocking review and passing verification before completing a change.

#### Scenario: Verify consumes review evidence

- GIVEN review produced a non-blocking report
- WHEN verify runs
- THEN it MUST cite the review report as evidence
- AND it MUST NOT duplicate the full review matrix.

#### Scenario: Archive checks review readiness

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST also require a non-blocking review report
- AND blocking review findings MUST prevent archive.
