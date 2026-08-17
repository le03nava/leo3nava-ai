# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Conflict and Ambiguity Resolution

When SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless an approved spec change explicitly redesigns DAG, artifact, routing, status, or persistence semantics. Explicit redesigns MUST define compatibility rules for old artifacts and MUST NOT silently invalidate archives. Status tokens, resolver rows, and persisted refs MAY preserve legacy `security-applicability` data, but MUST NOT normalize it into a runnable phase, launchable agent, or required new-change successor.
(Previously: compatibility preserved old routing tokens and artifact refs without explicitly forbidding a runnable applicability mapping.)

#### Scenario: Explicit DAG redesign is applied

- GIVEN an approved spec changes phase order
- WHEN contracts are updated
- THEN the new DAG MUST be authoritative for new changes
- AND legacy artifacts MUST remain readable only under documented compatibility rules.

#### Scenario: Compatibility is preserved

- GIVEN existing persisted SDD state or artifacts use established keys, paths, routing tokens, or status fields
- WHEN the contracts are updated
- THEN existing consumers MUST continue to resolve those artifacts and states without migration.

#### Scenario: Legacy token is not launchable

- GIVEN persisted state or status contains `security-applicability`
- WHEN routing or agent resolution is computed
- THEN the token MAY be interpreted as legacy data state
- AND it MUST NOT map to `sdd-security-applicability` or any runnable successor.

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include mandatory `security-design.md` and `review-security-report.md` refs, paths, dependency states, native/status token `review-security`, and archive gates. `security-applicability.md` and `securityApplicability` fields MAY appear only as legacy archived data refs and MUST NOT be active dependencies, produced artifacts, or phase-launch inputs.
(Previously: new state excluded active applicability dependencies, but legacy refs were not explicitly limited to non-launchable data.)

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include security design and security review report slots
- AND active dependencies MUST NOT include `security-applicability`.

#### Scenario: Legacy refs are preserved as data

- GIVEN an archived change contains `artifactRefs.securityApplicability`
- WHEN status or continuation displays historical evidence
- THEN the ref MAY remain visible as read-only data
- AND continuation MUST route active work through `design` or `security-design` instead of applicability.
