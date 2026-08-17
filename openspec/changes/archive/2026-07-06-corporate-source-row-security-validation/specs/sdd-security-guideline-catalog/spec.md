# Delta for sdd-security-guideline-catalog

## ADDED Requirements

### Requirement: Corporate Source Row Inventory

The catalog MUST define an exhaustive corporate source-row inventory derived from `Full Corporate Guideline Snapshot`. Range notation such as `1.1-1.10` MUST be expanded before validation. Each row MUST include `sourceId`, corporate section, guideline text, PCI alignment or `N/A`, mapped compact `SEC-*` guideline IDs, applies, complies, lifecycle status, evidence location, observations, and finding classification.

#### Scenario: Ranges expand before coverage

- GIVEN the snapshot contains a source range
- WHEN catalog coverage is validated
- THEN every concrete Source ID in the range MUST be represented exactly once
- AND missing, duplicate, or unknown Source IDs MUST block.

#### Scenario: PCI alignment is preserved

- GIVEN a corporate section heading includes PCI alignment
- WHEN source rows are produced
- THEN each row from that section MUST preserve the PCI requirement
- AND rows without alignment MUST state `N/A`.

### Requirement: Compact SEC Mapping Coverage

Every corporate source row MUST map to one or more existing compact `SEC-*` guideline IDs. The compact eight-control taxonomy MUST remain the architectural control layer and MUST NOT be replaced by per-source controls.

#### Scenario: Source row maps to compact controls

- GIVEN a Source ID is in the inventory
- WHEN catalog validation runs
- THEN it MUST resolve at least one valid compact `SEC-*` mapping
- AND unknown or missing mappings MUST block.

#### Scenario: Compact taxonomy is preserved

- GIVEN source rows are validated
- WHEN downstream phases consume the catalog
- THEN they MUST trace Source ID to compact `SEC-*`
- AND they MUST NOT create replacement compact controls.

### Requirement: Safe Source Row Evidence

Source-row evidence MUST be review-safe. Evidence locations and observations MUST NOT include secrets, PII, PAN, tokens, connection strings, private keys, or confidential values. `N/A` rows MUST include evidence and justification proving irrelevance.

#### Scenario: Applicable row lacks safe evidence

- GIVEN a source row applies
- WHEN evidence is missing or unsafe
- THEN the row MUST be non-compliant and blocking
- AND the report MUST identify the unsafe or absent evidence.

#### Scenario: N/A row is justified

- GIVEN a source row is marked `N/A`
- WHEN validation reviews the row
- THEN evidence MUST prove irrelevance
- AND missing justification MUST block.

### Requirement: Shared Security Contract Source Row Schema

`skills/_shared/sdd-security-contract.md` MUST define the source-row schema, allowed values, traceability, and routing semantics used by design, test-design, review-security, verify, and archive. The contract MUST remain compatible with OpenSpec, Engram, hybrid, and none modes by delegating persistence mechanics to the shared persistence contract.

#### Scenario: Contract consumers share schema

- GIVEN a phase writes or validates source rows
- WHEN it reads the shared security contract
- THEN it MUST use the same row fields and allowed values
- AND backend-specific storage MUST follow the persistence contract.

#### Scenario: Routing semantics are defined

- GIVEN source-row validation finds blockers or warnings
- WHEN a phase computes next routing
- THEN implementation blockers MUST route to apply, contract/artifact/N/A evidence blockers MUST route to resolve-blockers, and warnings-only results MAY route forward.
