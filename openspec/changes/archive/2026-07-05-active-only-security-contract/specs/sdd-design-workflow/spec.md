# Delta for sdd-design-workflow

## MODIFIED Requirements

### Requirement: Embedded Secure Development Design

`sdd-design` MUST persist `design.md` for every new change. The artifact MUST include a `## Secure Development Design` section that is the active classification and secure-design authority and evaluates all catalog rows: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001`.
(Previously: the requirement required the embedded section but did not explicitly name it as the active authority.)

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

### Requirement: Direct Routing to Test Design

For new changes, successful `sdd-design` MUST route directly to `sdd-test-design`; it MUST NOT require, launch, or produce a standalone security-design phase or artifact.
(Previously: the requirement blocked a standalone phase but did not explicitly prohibit producing a standalone artifact.)

#### Scenario: Design routes to test design

- GIVEN `design.md` includes the secure development section
- WHEN the phase returns a successful envelope
- THEN `next_recommended` MUST be `sdd-test-design` or `test-design`
- AND no standalone security-design artifact MUST be required for the new change.
