# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Verify and Archive Review Consumption

Verify MUST consume canonical `review-report.json` and canonical `review-security-report.json` when present. Security-review JSON MUST be authoritative for verdict, routing, blockers, warnings, artifact parity, and `sourceRowValidation.rows` exact-once coverage. Verify and archive MUST NOT consume derived Markdown or compact `SEC-*` data as validation authority. (Previously: verify consumed compact mappings, warnings, exceptions, and source-row evidence.)

#### Scenario: Verify consumes review evidence

- GIVEN review and security-review reports are non-blocking
- WHEN verify runs
- THEN it MUST cite both canonical JSON reports
- AND it MUST NOT duplicate or reinterpret their matrices.

#### Scenario: Compact report data is ignored

- GIVEN a security-review artifact contains legacy compact `SEC-*` report data
- WHEN verify or archive evaluates a new change
- THEN that compact data MUST NOT satisfy active security validation.

### Requirement: Source Row Persistence Compatibility

The SDD contracts MUST preserve source-row-first security evidence across OpenSpec, Engram, hybrid, and none modes. Active security-review artifacts MUST persist `sourceRowValidation.rows` with exactly 155 unique rows and required row fields. Derived Markdown MUST remain a generated compatibility view with lean navigation/summary plus the full row matrix at the end. Backend behavior MUST NOT redefine row semantics. (Previously: reports could persist summary-mode coverage and full rows only on request.)

#### Scenario: OpenSpec mode preserves rows

- GIVEN a change runs in OpenSpec mode
- WHEN security-review artifacts are persisted
- THEN canonical JSON MUST contain all 155 source rows
- AND Markdown MUST be regenerated from JSON.

#### Scenario: Engram or hybrid mode preserves rows

- GIVEN Engram or hybrid mode is selected
- WHEN source-row evidence is persisted
- THEN backend refs MUST use shared naming
- AND hybrid disagreements MUST block before continuing.

### Requirement: Verify Source Row Consumption

`sdd-verify` MUST validate that non-blocking `review-security-report.json` has complete `sourceRowValidation.rows` coverage and no row-level blockers. It MUST cite catalog snapshot/count, warnings, exceptions, and evidence refs without owning validation logic or compact-control summaries. (Previously: verify cited compact mappings.)

#### Scenario: Security source blocker remains

- GIVEN review-security reports a blocking source row
- WHEN verify runs
- THEN verification MUST block
- AND it MUST route to apply or resolve-blockers according to cause.

#### Scenario: Complete source rows continue

- GIVEN all 155 rows are present and non-blocking
- WHEN verify records evidence
- THEN it MAY proceed
- AND it MUST preserve warning and exception refs.

### Requirement: Archive Source Row Preservation

`sdd-archive` MUST require passing verification plus non-blocking source-row security review for new changes. Archive MUST preserve canonical JSON and generated Markdown refs, catalog snapshot identity/path, expected and validated counts, warnings, exceptions, and evidence references. Archive summaries MUST NOT use compact `SEC-*` validation, navigation, summaries, or `N/A` grouping. (Previously: archive preserved compact mappings and avoided full matrix duplication.)

#### Scenario: Archive checks row completeness

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST confirm 155 expected and validated rows with no blockers
- AND missing mandatory row evidence MUST prevent archive.

#### Scenario: Archive preserves generated matrix

- GIVEN derived Markdown contains the full source-row matrix at the end
- WHEN archive completes
- THEN the matrix ref MUST remain readable
- AND archive MUST not create a second validation source.
