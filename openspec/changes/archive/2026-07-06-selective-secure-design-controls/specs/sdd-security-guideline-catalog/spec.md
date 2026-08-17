# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: Compact Security Taxonomy

The catalog MUST expose a compact taxonomy: authentication, sessions, sensitive data or PAN, secrets, permissions or access control, files, database access, and sensitive logging. `sdd-design` MUST consume it as human-readable category guidance. `sdd-review-security` MUST consume it as the authoritative machine matrix.
(Previously: downstream design consumption implied compact category evaluation in design.)

#### Scenario: Security design uses taxonomy

- GIVEN a change modifies authentication behavior
- WHEN design maps the changed surface
- THEN it MUST be able to write an Authentication rules section
- AND it SHOULD avoid unrelated guideline text or matrix rows.

### Requirement: Catalog Boundary Preservation

The catalog MUST remain the source for identifiers, taxonomy, evidence expectations, exceptions, statuses, matrix vocabulary, compact inventory, Source ID inventory, and safe-evidence policy. Active authority MUST remain split between narrative design rules and exhaustive `review-security-report.md` validation. The catalog MUST NOT require design YAML, schema, matrices, exhaustive applicability, or `N/A` decisions.
(Previously: active authority treated embedded design plus review-security as if design carried compact applicability evidence.)

#### Scenario: Catalog authority is preserved

- GIVEN design and security review reference a guideline category
- WHEN downstream verification compares evidence
- THEN catalog identifiers and statuses MUST remain consistent
- AND conflicts MUST resolve through narrative design context plus exhaustive review-security results.

### Requirement: Corporate Source Row Inventory

The catalog MUST own the authoritative corporate source-row inventory from `Full Corporate Guideline Snapshot`: 155 expanded Source IDs, snapshot metadata, grouping, guideline text, PCI alignment or `N/A`, mapped compact `SEC-*` IDs, matrix vocabulary, and safe-evidence expectations. Ranges MUST be expanded before validation. Design and test-design MUST reference applicable category guidance; review-security MUST expand the inventory exhaustively.
(Previously: design and test-design referenced full count and grouped coverage as design obligations.)

#### Scenario: Ranges expand before coverage

- GIVEN the snapshot contains a source range
- WHEN catalog coverage is validated
- THEN every concrete Source ID in the range MUST be represented exactly once
- AND missing, duplicate, or unknown Source IDs MUST block.

### Requirement: Safe Source Row Evidence

Source-row evidence MUST be review-safe. Evidence locations and observations MUST NOT include secrets, PII, PAN, tokens, connection strings, private keys, or confidential values. Design MUST state safe-evidence policy for applicable category rules. `review-security-report.md` MUST justify every `N/A` row.
(Previously: N/A proof obligations were also required for non-applicable design rows.)

#### Scenario: Applicable row lacks safe evidence

- GIVEN a source row applies
- WHEN evidence is missing or unsafe
- THEN the row MUST be non-compliant and blocking
- AND the report MUST identify the unsafe or absent evidence.

### Requirement: Shared Security Contract Source Row Schema

`skills/_shared/sdd-security-contract.md` MUST define the review-security report schema, allowed values, traceability, artifact boundary, narrative design obligations, exhaustive review-security obligations, safe-evidence policy, and routing. It MUST state that catalog owns inventory, design owns narrative category rules, test-design owns planned checks, review-security owns schema/matrix validation, and verify/archive preserve evidence without design YAML.
(Previously: the contract said design owned slim classification and references without separating narrative design from exhaustive review.)

#### Scenario: Contract consumers share schema

- GIVEN a phase writes or validates source rows
- WHEN it reads the shared security contract
- THEN it MUST use the same report fields, allowed values, and artifact boundary
- AND backend-specific storage MUST follow the persistence contract.
