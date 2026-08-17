# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: Corporate Source Row Inventory

The catalog MUST own the authoritative exhaustive corporate source-row inventory derived from `Full Corporate Guideline Snapshot`. The inventory MUST contain 155 expanded Source IDs, snapshot identity/version metadata, section grouping, guideline text, PCI alignment or `N/A`, mapped compact `SEC-*` guideline IDs, matrix vocabulary, and safe-evidence expectations. Range notation such as `1.1-1.10` MUST be expanded in the catalog before downstream validation. Design and test-design MUST reference this inventory rather than duplicate it.
(Previously: the catalog defined the inventory but did not explicitly own the 155-row artifact boundary for slim design references.)

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

#### Scenario: Slim artifacts reference catalog inventory

- GIVEN design or test-design needs source-row coverage
- WHEN it records coverage obligations
- THEN it MUST cite the catalog snapshot and expected count instead of copying all 155 rows
- AND review-security MUST use the catalog to expand the exhaustive matrix.

### Requirement: Shared Security Contract Source Row Schema

`skills/_shared/sdd-security-contract.md` MUST define the source-row schema, allowed values, traceability, artifact boundary, and routing semantics used by design, test-design, review-security, verify, and archive. The contract MUST state that the catalog owns inventory, design owns slim classification and references, test-design owns planned evidence, review-security owns exhaustive expansion and validation, and verify/archive preserve evidence without redefining source-row semantics.
(Previously: the shared contract defined schema and routing semantics without the full slim-artifact ownership boundary.)

#### Scenario: Contract consumers share schema

- GIVEN a phase writes or validates source-row evidence
- WHEN it reads the shared security contract
- THEN it MUST use the same row fields, allowed values, and artifact ownership boundary
- AND backend-specific storage MUST follow the persistence contract.

#### Scenario: Routing semantics are defined

- GIVEN source-row validation finds blockers or warnings
- WHEN a phase computes next routing
- THEN implementation blockers MUST route to apply, contract/artifact/N/A evidence blockers MUST route to resolve-blockers, and warnings-only results MAY route forward.
