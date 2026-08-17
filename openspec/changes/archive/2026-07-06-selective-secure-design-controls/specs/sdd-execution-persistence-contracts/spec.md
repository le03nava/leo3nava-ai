# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Conflict and Ambiguity Resolution

When SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless an approved spec explicitly redesigns DAG, artifact, routing, status, or persistence semantics. Historical exhaustive secure-design matrices MAY remain readable as archives, but new changes MUST treat narrative `design.md#secure-development-design` plus machine-readable `review-security-report.md` as the active boundary. Historical `security-applicability` refs MAY display as data, but MUST NOT normalize into a runnable phase, launchable agent, active authority, or required successor.
(Previously: compatibility did not distinguish historical exhaustive design matrices from the new narrative design boundary.)

#### Scenario: Explicit DAG redesign is applied

- GIVEN an approved spec changes phase order
- WHEN contracts are updated
- THEN the new DAG MUST be authoritative for new changes
- AND older artifacts MUST remain readable only under documented compatibility rules.

#### Scenario: Compatibility is preserved

- GIVEN existing persisted SDD state or artifacts use established keys, paths, routing tokens, exhaustive design rows, or status fields
- WHEN the contracts are updated
- THEN existing consumers MUST continue to resolve those artifacts and states without migration.

#### Scenario: Historical token is not launchable

- GIVEN persisted state or status contains `security-applicability`
- WHEN routing or agent resolution is computed
- THEN the token MAY be interpreted as historical data state
- AND it MUST NOT map to `sdd-security-applicability` or any runnable successor.

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include `design.md` with narrative secure development rules and `review-security-report.md` refs, dependency states, `review-security` token, and archive gates. Design MUST persist classification rationale, changed-surface inventory, applicable category rules, evidence owners, residual risks, exceptions, and safe-evidence policy. It MUST NOT require YAML, schemas, compact matrices, Source ID matrices, exhaustive applicability, or `N/A` rows. Those machine-readable artifacts MUST persist in `review-security-report.md`. `security-design.md` and `security-applicability.md` MAY appear only as historical refs.
(Previously: security artifact status expected embedded secure development rows without separating narrative design from exhaustive review-security.)

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include design and security review report slots
- AND active dependencies MUST NOT include `security-design` or `security-applicability`.

#### Scenario: Legacy refs are preserved as data

- GIVEN an archived change contains `artifactRefs.securityDesign`, `artifactRefs.securityApplicability`, or exhaustive design rows
- WHEN status or continuation displays historical evidence
- THEN the ref or row MAY remain visible as read-only data
- AND continuation MUST route active work through narrative design and review-security.

### Requirement: Source Row Persistence Compatibility

The SDD contracts MUST preserve corporate source-row evidence across OpenSpec, Engram, hybrid, and none modes according to the shared persistence contract. Backend behavior MUST NOT redefine source-row semantics. Source-row artifacts MUST remain recoverable through established review-security, verify, and archive keys/paths. Persistence MUST allow narrative designs and archived exhaustive designs to coexist without migration; verify/archive MUST require narrative design evidence plus the review-security report schema, not design YAML.
(Previously: compatibility preserved source rows but did not state coexistence for narrative and historical exhaustive designs.)

#### Scenario: OpenSpec mode preserves rows

- GIVEN a change runs in OpenSpec mode
- WHEN source-row artifacts are persisted
- THEN rows MUST be stored in `review-security-report.md`
- AND downstream phases MUST read that report as source-row evidence.

#### Scenario: Engram or hybrid mode preserves rows

- GIVEN Engram or hybrid mode is selected
- WHEN source-row artifacts are persisted
- THEN Engram keys MUST use the shared artifact naming contract
- AND hybrid mode MUST reconcile backend disagreements before continuing.

#### Scenario: None mode is explicit

- GIVEN none mode is selected
- WHEN source-row evidence is produced inline
- THEN no project files or Engram observations MUST be written
- AND downstream recovery limits MUST be reported.
