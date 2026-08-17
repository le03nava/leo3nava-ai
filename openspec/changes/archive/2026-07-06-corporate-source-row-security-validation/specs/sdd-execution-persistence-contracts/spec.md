# Delta for sdd-execution-persistence-contracts

## ADDED Requirements

### Requirement: Source Row Persistence Compatibility

The SDD contracts MUST preserve corporate source-row evidence across OpenSpec, Engram, hybrid, and none modes according to the shared persistence contract. Backend behavior MUST NOT redefine source-row semantics, and source-row artifacts MUST remain recoverable through the established artifact keys or paths for design, test-design, review-security, verify, and archive.

#### Scenario: OpenSpec mode preserves rows

- GIVEN a change runs in OpenSpec mode
- WHEN source-row artifacts are persisted
- THEN rows MUST be stored in the established change files
- AND downstream phases MUST read those files as source-row evidence.

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

### Requirement: Verify Source Row Consumption

`sdd-verify` MUST consume non-blocking `review-security-report.md` source-row evidence and validate that no source-row blockers remain. Verify MUST cite the security review verdict without owning or duplicating the full source-row matrix.

#### Scenario: Security source blocker remains

- GIVEN review-security reports a blocking source row
- WHEN verify runs
- THEN verification MUST block
- AND it MUST route to apply or resolve-blockers according to the blocker cause.

#### Scenario: Warnings only after security review

- GIVEN review-security is non-blocking with warnings only
- WHEN verify records evidence
- THEN it MUST preserve the warnings
- AND verification MAY proceed if all mandatory evidence is complete.

### Requirement: Archive Source Row Preservation

`sdd-archive` MUST require passing verification plus non-blocking source-row security review for new changes. Archive MUST preserve source-row coverage, mappings, warnings, exceptions, and evidence references without requiring legacy standalone security artifacts.

#### Scenario: Archive checks no source blockers remain

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST confirm no source-row blockers remain
- AND missing mandatory source-row evidence MUST prevent archive.

#### Scenario: Archive preserves audit trail

- GIVEN archive completes
- WHEN future readers inspect the archived change
- THEN source-row coverage and evidence references MUST remain understandable
- AND compact `SEC-*` mappings MUST still be traceable.
