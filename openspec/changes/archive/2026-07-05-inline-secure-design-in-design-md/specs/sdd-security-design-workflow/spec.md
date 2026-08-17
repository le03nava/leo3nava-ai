# Delta for sdd-security-design-workflow

## MODIFIED Requirements

### Requirement: Mandatory Security Design Phase

For new changes, the SDD workflow MUST NOT run `sdd-security-design` as an active phase. Secure development design MUST be completed inside `design.md`; standalone `security-design.md` MAY be read only for legacy/archive compatibility.
(Previously: every new change routed from `sdd-design` to mandatory `sdd-security-design` before `sdd-test-design`.)

#### Scenario: Every new change requires security design

- GIVEN `sdd-design` has completed for a new change
- WHEN routing is computed
- THEN the next required phase MUST be `sdd-test-design`
- AND `sdd-security-design` MUST NOT be launched.

#### Scenario: No-impact still creates artifact

- GIVEN no security guideline applies
- WHEN `sdd-design` writes `design.md`
- THEN the embedded security rows MUST record `not-applicable` evidence
- AND no standalone `security-design.md` artifact MUST be created for that reason.

### Requirement: Security Design Artifact Contract

For new changes, `design.md` MUST own security classification, catalog snapshot identity, guideline matrix, controls, expected evidence, lifecycle statuses, risks, exceptions, and archive gate expectations. Standalone `security-design.md` MUST remain legacy-only readable evidence.
(Previously: `sdd-security-design` created and owned `security-design.md` for every new change.)

#### Scenario: Classification and controls are unified

- GIVEN design context is available
- WHEN `sdd-design` writes `design.md`
- THEN it MUST classify security impact in `## Secure Development Design`
- AND it MUST map applicable guidelines to controls and expected evidence.

#### Scenario: Applicability risks are carried forward

- GIVEN prior security risks exist in proposal or specs
- WHEN design is written
- THEN `design.md` MUST resolve each risk or preserve it with an evidence owner.

#### Scenario: Persistence boundary is delegated

- GIVEN `sdd-design` persists or resolves `design.md`
- WHEN backend behavior is needed
- THEN the phase MUST follow the shared persistence authority.

### Requirement: Enriched Applicability Consumption

For new changes, secure development design MUST classify proposal, spec, and technical design context directly inside `design.md`. Legacy `security-applicability.md` and `security-design.md` MAY be read only as compatibility context for old or archived changes.
(Previously: `sdd-security-design` consumed proposal, spec, and technical design directly.)

#### Scenario: Direct classification

- GIVEN proposal, specs, and design inputs are readable
- WHEN `sdd-design` runs for a new change
- THEN it MUST classify security categories from those inputs
- AND it MUST NOT require `security-applicability.md` or standalone `security-design.md`.

#### Scenario: Conditional obligation becomes design predicate

- GIVEN an applicable guideline has `conditional` severity
- WHEN `sdd-design` maps controls
- THEN it MUST state the predicate and required evidence
- AND unresolved true predicates MUST remain trackable downstream.

#### Scenario: Retired executor is not consulted

- GIVEN a new change reaches design
- WHEN security inputs are resolved
- THEN `sdd-security-design` MUST NOT launch or block the workflow.
