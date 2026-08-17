# Delta for sdd-review-workflow

## MODIFIED Requirements

### Requirement: Operational Readiness General Review

`sdd-review` MUST validate operational evidence only when design, test-design, tasks, apply evidence, or archived context says operational considerations apply. Review MUST assess traceability, no-invention behavior, and placeholder usage for present or planned evidence without changing the fixed 96-control matrix shape. Review MUST NOT enforce mandatory operational category completeness or require real operational data disclosure.
(Previously: review always validated operational-readiness existence and exact placeholder usage.)

#### Scenario: Operational evidence is traceable

- GIVEN applicable operational evidence is present in SDD artifacts
- WHEN general review runs
- THEN it MUST verify the evidence cites safe sources, `Pendiente de confirmar:`, or `No aplica.`.

#### Scenario: Review detects invented data

- GIVEN operational evidence contains unsupported operational details
- WHEN review validates traceability
- THEN the finding MUST be blocking or routed to `resolve-blockers`.

#### Scenario: Operational evidence is absent by design

- GIVEN design marks operational considerations not applicable
- WHEN review runs
- THEN absence of readiness categories MUST NOT be a blocker.

#### Scenario: Matrix shape is preserved

- GIVEN operational review findings are recorded
- WHEN `review-report.md` is written
- THEN the 96-control matrix columns MUST remain unchanged
- AND operational evidence MAY appear in a separate section.

### Requirement: Operational Review Handoff

Review MUST hand off changed-file, evidence, placeholder, and gap context for operational considerations that exist. It MUST NOT create a handoff requirement from an absent shared readiness contract.
(Previously: review always handed off readiness context to security review.)

#### Scenario: Security review receives applicable context

- GIVEN general review finds applicable operational evidence
- WHEN security review starts
- THEN it MUST be able to read evidence locations and unresolved gaps from review output.

#### Scenario: No applicable context exists

- GIVEN no operational considerations apply
- WHEN security review starts
- THEN missing readiness handoff MUST NOT block the phase.
