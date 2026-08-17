# Delta for sdd-operational-readiness-workflow

## ADDED Requirements

### Requirement: Design-Driven Operational Considerations

Operational considerations MUST be conditional and design-driven. Active phases MUST NOT depend on `skills/_shared/sdd-operational-readiness-contract.md` or mandatory readiness completeness gates. When the design says operational considerations apply, phases SHOULD track applicable logs/errors, monitoring, administration operations, reprocessing/recovery, and backup/retention/cleanup/generated-artifact concerns through actual design, test-design, tasks, apply, review, verify, and archive evidence.

#### Scenario: Design marks considerations applicable

- GIVEN design includes operational considerations
- WHEN downstream phases run
- THEN they MUST consume the design-owned evidence expectations
- AND they MUST preserve safe placeholders and unresolved gaps.

#### Scenario: Design marks considerations not applicable

- GIVEN design omits the section or records `No aplica.`
- WHEN downstream phases validate operational evidence
- THEN they MUST NOT require mandatory readiness category completeness.

## MODIFIED Requirements

### Requirement: Safe SDD Evidence Boundary

Ordinary SDD evidence, code, tests, fixtures, and examples MUST NOT include production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, generated file bytes, or equivalent restricted operational details. Such identifiers MAY appear only in final operational documentation when explicitly provided by the user. This boundary applies wherever operational evidence exists.
(Previously: this safety boundary supported a mandatory operational-readiness workflow.)

#### Scenario: Restricted value appears in SDD evidence

- GIVEN an SDD artifact includes restricted operational data
- WHEN review-security evaluates it
- THEN the evidence MUST be rejected as unsafe.

### Requirement: Phase Ownership Model

Design MUST own whether operational considerations apply and define expected evidence when they do. Test-design, tasks, review, verify, and archive MUST consume actual design/test-design/tasks/apply/archive evidence without requiring a shared readiness contract. Review and review-security MUST keep distinct responsibilities.
(Previously: all phases confirmed collection, traceability, completeness, and unresolved gaps for mandatory readiness.)

#### Scenario: Design declares monitoring strategy

- GIVEN a change needs operational monitoring
- WHEN design records operational considerations
- THEN it MUST describe applicable monitoring mechanisms, not SQL-only checks.

#### Scenario: Review responsibilities differ

- GIVEN both review gates run and operational evidence exists
- WHEN the evidence is assessed
- THEN general review MUST validate existence and traceability
- AND security review MUST validate leakage and restricted-data boundaries.

### Requirement: Manual Operational Document Consumption

`sdd-operational-doc` MUST remain a manual post-archive utility. It MUST consume archived evidence, MUST NOT invent missing data, and MUST preserve sections 1-9 and diagrams R1-R4. Missing applicable values MUST remain pending; absent inapplicable values MUST be `No aplica.`.
(Previously: it consumed archived readiness evidence from the mandatory workflow.)

#### Scenario: Archived evidence has gaps

- GIVEN archived evidence has unresolved or inapplicable fields
- WHEN the operational document is generated
- THEN missing applicable values MUST remain pending
- AND inapplicable values MUST be `No aplica.`.

## REMOVED Requirements

### Requirement: Mandatory Operational Readiness Evaluation

(Reason: operational readiness is no longer mandatory for every SDD change; applicability is owned by design.)
(Migration: use `Design-Driven Operational Considerations` and downstream evidence from actual artifacts.)
