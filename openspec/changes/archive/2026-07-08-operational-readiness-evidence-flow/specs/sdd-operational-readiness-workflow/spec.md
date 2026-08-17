# sdd-operational-readiness-workflow Specification

## Purpose

Define mandatory operational-readiness evaluation, safe evidence boundaries, phase ownership, archive handoff, and manual operational documentation consumption for SDD changes.

## Requirements

### Requirement: Mandatory Operational Readiness Evaluation

Every SDD change MUST evaluate operational readiness. Operational data disclosure is optional and MUST NOT be required when unavailable or unsafe. Every operational field MUST contain evidence, exact `Pendiente de confirmar:`, or exact `No aplica.` with optional rationale.

#### Scenario: Field is resolved safely

- GIVEN an operational readiness field is evaluated
- WHEN safe evidence exists
- THEN the field MUST cite the evidence without restricted data.

#### Scenario: Field remains unresolved or inapplicable

- GIVEN evidence is unavailable or the field is out of scope
- WHEN readiness is recorded
- THEN the field MUST use exact `Pendiente de confirmar:` or exact `No aplica.`.

### Requirement: Safe SDD Evidence Boundary

Ordinary SDD evidence, code, tests, fixtures, and examples MUST NOT include production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, or generated file bytes. Such operational identifiers MAY appear only in final operational documentation when explicitly provided by the user.

#### Scenario: Restricted value appears in SDD evidence

- GIVEN an SDD artifact includes restricted operational data
- WHEN review-security evaluates it
- THEN the evidence MUST be rejected as unsafe.

### Requirement: Phase Ownership Model

Design MUST define the operational-readiness strategy and expected evidence. Tasks, review, verify, and archive MUST confirm collection, traceability, completeness, and unresolved gaps. Review and review-security MUST keep distinct responsibilities.

#### Scenario: Design declares monitoring strategy

- GIVEN a change needs operational monitoring
- WHEN design records readiness
- THEN it MUST describe monitoring mechanisms, not SQL-only checks.

#### Scenario: Review responsibilities differ

- GIVEN both review gates run
- WHEN readiness is assessed
- THEN general review MUST validate existence and traceability
- AND security review MUST validate leakage and restricted-data boundaries.

### Requirement: Manual Operational Document Consumption

`sdd-operational-doc` MUST remain a manual post-archive utility. It MUST consume archived readiness evidence, MUST NOT invent missing data, and MUST preserve sections 1-9 and diagrams R1-R4.

#### Scenario: Archived evidence has gaps

- GIVEN archived readiness evidence has unresolved fields
- WHEN the operational document is generated
- THEN missing fields MUST remain `Pendiente de confirmar:`
- AND sections 1-9 plus diagrams R1-R4 MUST remain present.
