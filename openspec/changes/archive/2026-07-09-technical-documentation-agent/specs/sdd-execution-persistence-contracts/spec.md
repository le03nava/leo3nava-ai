# Delta for sdd-execution-persistence-contracts

## ADDED Requirements

### Requirement: Manual Technical Document Boundary

The DAG MUST NOT treat `sdd-technical-doc` as a required phase. The utility MUST remain manual, post-archive, and archive-consuming, analogous to `sdd-operational-doc`. It MUST NOT change phase order, status routing, dependency graph, verify gates, archive gates, or required persistence artifacts. It MUST generate from archived evidence only, MUST NOT invent data, MUST mark inapplicable sections as `No aplica.`, and MUST mark unavailable applicable information explicitly.

#### Scenario: Archive completes without technical document

- GIVEN verify passes and archive criteria are met
- WHEN archive runs
- THEN completion MUST NOT require `sdd-technical-doc` execution
- AND missing technical documentation MUST NOT block archive.

#### Scenario: Status does not expose a required phase

- GIVEN status or continuation evaluates an active SDD change
- WHEN successor phases or required artifacts are computed
- THEN `sdd-technical-doc` MUST NOT appear as a required phase, DAG token, status dependency, verify input, or archive input.

#### Scenario: Manual utility consumes archived evidence

- GIVEN an archived change exists
- WHEN `sdd-technical-doc` is invoked manually
- THEN it MUST read archived evidence as its source of truth
- AND absent or inapplicable values MUST be represented without invention.
