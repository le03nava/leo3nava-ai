# Delta for sdd-review-workflow

## ADDED Requirements

### Requirement: Operational Readiness General Review

`sdd-review` MUST validate operational-readiness existence, traceability, no-invention behavior, and exact placeholder usage without changing the fixed 96-control matrix shape. Readiness findings MAY be reported outside the matrix. Review MUST NOT require real operational data disclosure.

#### Scenario: Readiness evidence is traceable

- GIVEN apply evidence and SDD artifacts are readable
- WHEN general review runs
- THEN it MUST verify operational-readiness fields cite evidence, `Pendiente de confirmar:`, or `No aplica.`.

#### Scenario: Review detects invented data

- GIVEN a readiness field contains unsupported operational details
- WHEN review validates traceability
- THEN the finding MUST be blocking or routed to `resolve-blockers`.

#### Scenario: Matrix shape is preserved

- GIVEN readiness review is added
- WHEN `review-report.md` is written
- THEN the 96-control matrix columns MUST remain unchanged
- AND readiness evidence MAY appear in a separate section.

### Requirement: Operational Review Handoff

Review MUST hand off enough changed-file, evidence, placeholder, and gap context for `sdd-review-security` to validate leakage boundaries without duplicating security-review responsibilities.

#### Scenario: Security review receives readiness context

- GIVEN general review is non-blocking
- WHEN security review starts
- THEN it MUST be able to read readiness evidence locations and unresolved gaps from review output.
