# Delta for sdd-execution-persistence-contracts

## MODIFIED Requirements

### Requirement: Operational Readiness Evidence Persistence

SDD persistence, status, verify, and archive contracts MUST preserve operational considerations, evidence, placeholders, gaps, and artifact references when they exist in design, test-design, tasks, apply, review, security review, verify, or archive evidence. Verify and archive MUST consume actual evidence rather than a shared operational-readiness contract. They MUST NOT require mandatory operational category completeness or disclosure of real operational data.
(Previously: persistence, verify, and archive preserved readiness evidence and confirmed readiness completeness.)

#### Scenario: Operational refs survive workflow

- GIVEN any SDD artifact records operational considerations or gaps
- WHEN status, verify, or archive resolves artifacts
- THEN those refs and unresolved gaps MUST remain readable.

#### Scenario: Verify checks applicable evidence

- GIVEN design or downstream artifacts make operational evidence applicable
- WHEN verify runs
- THEN each applicable item MUST have safe evidence, `Pendiente de confirmar:`, or `No aplica.`.

#### Scenario: No applicable evidence exists

- GIVEN design marks operational considerations not applicable or omits them safely
- WHEN verify or archive runs
- THEN missing readiness categories MUST NOT block completion.

### Requirement: Manual Operational Document Boundary

The DAG MUST NOT treat `sdd-operational-doc` as a required phase. The utility MUST remain manual, post-archive, and archive-consuming. It MUST generate from archived evidence, MUST NOT invent data, and MUST mark absent inapplicable values as `No aplica.` or unresolved applicable values as pending while preserving operational document sections 1-9 and diagrams R1-R4.
(Previously: the utility consumed archived readiness evidence and unresolved fields remained pending.)

#### Scenario: Archive completes without operational doc

- GIVEN verify passes and archive criteria are met
- WHEN archive runs
- THEN completion MUST NOT require `sdd-operational-doc` execution.

#### Scenario: Manual utility consumes archive

- GIVEN an archived change contains operational evidence or gaps
- WHEN `sdd-operational-doc` is invoked manually
- THEN it MUST read archived evidence first
- AND absent values MUST remain pending or `No aplica.` without invention.

### Requirement: Final Documentation Restricted Data Boundary

Production hostnames, IPs, ports, SID/service names, and similar operational identifiers MAY be included only in final operational documentation when explicitly provided by the user. Ordinary SDD evidence and examples MUST preserve safe placeholders or references for applicable operational considerations.
(Previously: the boundary was scoped to mandatory operational-readiness evidence.)

#### Scenario: User provides final operational values

- GIVEN the user explicitly provides production operational identifiers for the manual document
- WHEN documentation is generated
- THEN the final operational document MAY include them
- AND SDD evidence artifacts MUST NOT be backfilled with those values.
