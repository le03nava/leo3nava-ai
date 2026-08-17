# Delta for sdd-design-workflow

## MODIFIED Requirements

### Requirement: Operational Readiness Planning

`sdd-design` MAY include `## Operational Considerations` when the change design makes operational behavior applicable. The section MUST be design-owned and conditional, not a mandatory readiness gate. When present, it SHOULD cover applicable logs/errors, monitoring, administration operations, reprocessing/recovery, and backup/retention/cleanup/generated-artifact concerns. It MUST define expected downstream evidence only for applicable concerns and MUST use safe evidence, `Pendiente de confirmar:`, or `No aplica.` without inventing data.
(Previously: every design had to include mandatory `## Operational Readiness` and evaluate all readiness fields.)

#### Scenario: Applicable operational behavior is designed

- GIVEN proposal and specs show operational behavior applies
- WHEN `sdd-design` succeeds
- THEN `design.md` SHOULD include `## Operational Considerations`
- AND it MUST describe expected evidence for applicable concerns only.

#### Scenario: Operational behavior does not apply

- GIVEN operational considerations are out of scope
- WHEN `design.md` is written
- THEN the section MAY be omitted or state `No aplica.` with rationale
- AND downstream phases MUST NOT require readiness completeness.

#### Scenario: Production data is unavailable

- GIVEN an applicable operational field lacks user-provided data
- WHEN design records the consideration
- THEN it MUST use `Pendiente de confirmar:` instead of inventing data.

### Requirement: Operational Evidence Safety in Design

Design MUST keep safe SDD evidence separate from final operational documentation. It MUST NOT include production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, or generated file bytes unless the user explicitly provided them for final operational documentation. This safety boundary applies only to operational evidence that the design includes or references.
(Previously: safety applied inside a mandatory operational-readiness section.)

#### Scenario: Restricted identifier is needed later

- GIVEN production connection details are needed for handoff
- WHEN design lacks explicit user-provided values
- THEN design MUST record the gap safely
- AND final documentation MAY receive those values only from the user.
