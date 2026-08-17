# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Conflict and Ambiguity Resolution

When SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless an approved spec change explicitly redesigns DAG, artifact, routing, status, or persistence semantics. Explicit redesigns MUST define compatibility rules for old artifacts and MUST NOT silently invalidate archives. Status tokens, resolver rows, and persisted refs MAY preserve historical `security-applicability` data for read/display behavior, but MUST NOT normalize it into a runnable phase, launchable agent, active security authority, or required new-change successor.
(Previously: the requirement preserved legacy data but did not explicitly exclude active security authority.)

#### Scenario: Explicit DAG redesign is applied

- GIVEN an approved spec changes phase order
- WHEN contracts are updated
- THEN the new DAG MUST be authoritative for new changes
- AND older artifacts MUST remain readable only under documented compatibility rules.

#### Scenario: Compatibility is preserved

- GIVEN existing persisted SDD state or artifacts use established keys, paths, routing tokens, or status fields
- WHEN the contracts are updated
- THEN existing consumers MUST continue to resolve those artifacts and states without migration.

#### Scenario: Historical token is not launchable

- GIVEN persisted state or status contains `security-applicability`
- WHEN routing or agent resolution is computed
- THEN the token MAY be interpreted as historical data state
- AND it MUST NOT map to `sdd-security-applicability` or any runnable successor.

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include `design.md` with embedded secure development rows and `review-security-report.md` refs, paths, dependency states, native/status token `review-security`, and archive gates. `security-design.md` and `security-applicability.md` MAY appear only as historical data refs needed for existing state or archive readability and MUST NOT be active dependencies, produced artifacts, phase-launch inputs, or active security authority.
(Previously: the requirement preserved legacy refs as data but did not explicitly exclude active authority.)

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include design and security review report slots
- AND active dependencies MUST NOT include `security-design` or `security-applicability`.

#### Scenario: Historical refs are preserved as data

- GIVEN an archived change contains `artifactRefs.securityDesign` or `artifactRefs.securityApplicability`
- WHEN status or continuation displays historical evidence
- THEN the ref MAY remain visible as read-only data
- AND continuation MUST route active work through `design` instead of security-design or applicability.
