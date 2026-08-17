# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Verify Source Row Consumption

`sdd-verify` MUST consume non-blocking `review-security-report.md` source-row evidence and validate that no source-row blockers remain. Verify MUST cite the security review verdict, catalog snapshot/count, compact mappings, warnings, exceptions, and evidence references without owning or duplicating the full source-row matrix.
(Previously: verify consumed security review source-row evidence without explicitly preserving slim-design boundary evidence.)

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

#### Scenario: Verify preserves boundary evidence

- GIVEN review-security is non-blocking and cites slim design coverage
- WHEN verify records final evidence
- THEN it MUST preserve catalog identity, expected count, compact mappings, and report links
- AND it MUST NOT require standalone `security-design.md` or `security-applicability.md`.

### Requirement: Archive Source Row Preservation

`sdd-archive` MUST require passing verification plus non-blocking source-row security review for new changes. Archive MUST preserve source-row coverage summaries, catalog snapshot identity/path, expected count, compact `SEC-*` mappings, warnings, exceptions, and evidence references without requiring legacy standalone security artifacts or copying the full review-security matrix into design/archive summaries.
(Previously: archive preserved source-row coverage, mappings, warnings, exceptions, and evidence references without legacy standalone artifacts.)

#### Scenario: Archive checks no source blockers remain

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST confirm no source-row blockers remain
- AND missing mandatory source-row evidence MUST prevent archive.

#### Scenario: Archive preserves audit trail

- GIVEN archive completes
- WHEN future readers inspect the archived change
- THEN source-row coverage summaries and evidence references MUST remain understandable
- AND compact `SEC-*` mappings MUST still be traceable.

#### Scenario: Archive avoids matrix duplication

- GIVEN `review-security-report.md` contains the exhaustive source-row matrix
- WHEN archive writes final records
- THEN it MUST link or summarize the matrix instead of duplicating it
- AND archived evidence MUST remain readable without standalone legacy security artifacts.
