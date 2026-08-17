# Delta for sdd-design-workflow

## ADDED Requirements

### Requirement: Operational Readiness Planning

`sdd-design` MUST include `## Operational Readiness` when producing `design.md`. The section MUST define operational strategy and expected evidence for logs, monitoring mechanisms, administration, reprocessing, ownership, final documentation inputs, and unresolved gaps. Evaluation is mandatory; disclosure of real operational data is not. Every field MUST contain safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.` with optional rationale.

#### Scenario: Readiness strategy is designed

- GIVEN proposal and specs are readable
- WHEN `sdd-design` succeeds
- THEN `design.md` MUST include `## Operational Readiness`
- AND it MUST describe expected evidence and monitoring mechanisms.

#### Scenario: Production data is unavailable

- GIVEN a required operational field lacks user-provided data
- WHEN design records readiness
- THEN it MUST use exact `Pendiente de confirmar:` instead of inventing data.

#### Scenario: Field is not applicable

- GIVEN an operational field is out of scope
- WHEN design records readiness
- THEN it MUST use exact `No aplica.` with optional rationale.

### Requirement: Operational Evidence Safety in Design

Design MUST keep safe SDD evidence separate from final operational documentation. It MUST NOT include production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, or generated file bytes unless the user explicitly provided them for final operational documentation.

#### Scenario: Restricted identifier is needed later

- GIVEN production connection details are needed for handoff
- WHEN design lacks explicit user-provided values
- THEN design MUST record the gap safely
- AND final documentation MAY receive those values only from the user.
