# sdd-design-workflow Specification

## Purpose

Define `sdd-design` as the mandatory technical design phase that also owns secure development design for new changes.

## Requirements

### Requirement: Embedded Secure Development Design

`sdd-design` MUST persist `design.md` for every new change. The artifact MUST include a `## Secure Development Design` section that evaluates all catalog rows: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001`.

#### Scenario: Design contains all security rows

- GIVEN proposal and specs are readable
- WHEN `sdd-design` succeeds
- THEN `design.md` MUST include `## Secure Development Design`
- AND the section MUST contain all 8 catalog guideline IDs.

#### Scenario: No-impact change records rationale

- GIVEN no security category applies
- WHEN `design.md` is written
- THEN every row MUST be marked `N/A` or `not-applicable`
- AND each row MUST include rationale and evidence proving irrelevance.

### Requirement: Secure Development Design Row Contract

Each secure design row MUST state applicability, rationale, secure design decision, controls, evidence owner, expected evidence, lifecycle status, residual risk, and exception data when evidence cannot be produced.

#### Scenario: Applicable guideline has obligations

- GIVEN a guideline applies to the change
- WHEN `sdd-design` records the row
- THEN it MUST identify controls and expected downstream evidence
- AND unresolved mandatory evidence MUST remain visible to test design, review-security, verify, and archive.

#### Scenario: Exception is required

- GIVEN mandatory evidence cannot be planned
- WHEN the row records an exception
- THEN it MUST include approver, accepted-risk rationale, mitigation or follow-up, and evidence gap.

### Requirement: Direct Routing to Test Design

For new changes, successful `sdd-design` MUST route directly to `sdd-test-design`; it MUST NOT require or launch `sdd-security-design`.

#### Scenario: Design routes to test design

- GIVEN `design.md` includes the secure development section
- WHEN the phase returns a successful envelope
- THEN `next_recommended` MUST be `sdd-test-design` or `test-design`
- AND no standalone security-design artifact MUST be required for the new change.
