# Delta for sdd-execution-persistence-contracts

## ADDED Requirements

### Requirement: Operational Readiness Evidence Persistence

SDD persistence, status, verify, and archive contracts MUST preserve operational-readiness evidence, placeholders, gaps, and artifact references across supported artifact-store modes. Verify and archive MUST confirm readiness completeness without requiring disclosure of real operational data.

#### Scenario: Readiness refs survive workflow

- GIVEN design, test-design, tasks, apply, review, and security review record readiness evidence
- WHEN status, verify, or archive resolves artifacts
- THEN readiness refs and unresolved gaps MUST remain readable.

#### Scenario: Verify checks readiness completion

- GIVEN mandatory readiness fields exist
- WHEN verify runs
- THEN each field MUST have evidence, exact `Pendiente de confirmar:`, or exact `No aplica.`.

#### Scenario: Archive preserves handoff evidence

- GIVEN verification passes
- WHEN archive completes
- THEN archived evidence MUST preserve readiness status, evidence refs, and unresolved gaps for manual documentation.

### Requirement: Manual Operational Document Boundary

The DAG MUST NOT treat `sdd-operational-doc` as a required phase. The utility MUST remain manual, post-archive, and archive-consuming. It MUST NOT invent data and MUST preserve operational document sections 1-9 and diagrams R1-R4.

#### Scenario: Archive completes without operational doc

- GIVEN verify passes and archive criteria are met
- WHEN archive runs
- THEN completion MUST NOT require `sdd-operational-doc` execution.

#### Scenario: Manual utility consumes archive

- GIVEN an archived change contains readiness evidence
- WHEN `sdd-operational-doc` is invoked manually
- THEN it MUST read archived evidence first
- AND unresolved fields MUST remain `Pendiente de confirmar:`.

### Requirement: Final Documentation Restricted Data Boundary

Production hostnames, IPs, ports, SID/service names, and similar operational identifiers MAY be included only in final operational documentation when explicitly provided by the user. Ordinary SDD evidence and examples MUST preserve safe placeholders or references instead.

#### Scenario: User provides final operational values

- GIVEN the user explicitly provides production operational identifiers for the manual document
- WHEN documentation is generated
- THEN the final operational document MAY include them
- AND SDD evidence artifacts MUST NOT be backfilled with those values.
