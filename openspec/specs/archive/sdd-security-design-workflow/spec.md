# sdd-security-design-workflow Specification

## Purpose

Define the retired standalone security-design workflow and the legacy/archive compatibility rules for old `security-design.md` artifacts. New changes perform secure development design inside `design.md`.

## Requirements

### Requirement: Mandatory Security Design Phase

For new changes, the SDD workflow MUST NOT run `sdd-security-design` as an active phase. Secure development design MUST be completed inside `design.md`; standalone `security-design.md` MAY be read only for legacy/archive compatibility.

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

### Requirement: Mandatory Evidence and Exceptions

Embedded secure development design MUST identify mandatory evidence for every applicable guideline. Archive MUST be blocked when mandatory evidence is missing unless an approved exception records approver, guideline, accepted-risk rationale, and mitigation or follow-up.

#### Scenario: Missing mandatory evidence blocks archive

- GIVEN an applicable mandatory guideline has no evidence or approved exception
- WHEN archive readiness is evaluated
- THEN archive MUST be blocked
- AND the blocker MUST name the guideline and missing evidence.

#### Scenario: Approved exception allows archive

- GIVEN mandatory evidence is missing but an exception records approver, guideline, accepted-risk rationale, and mitigation or follow-up
- WHEN archive readiness is evaluated
- THEN archive MAY proceed
- AND the exception MUST remain in the audit trail.

### Requirement: Enriched Applicability Consumption

For new changes, secure development design MUST classify proposal, spec, and technical design context directly inside `design.md`. Legacy `security-applicability.md` and `security-design.md` MAY be read only as compatibility context for old or archived changes.

#### Scenario: Direct classification

- GIVEN proposal, specs, and design are readable
- WHEN `sdd-design` runs for a new change
- THEN it MUST classify categories from those inputs
- AND it MUST NOT require `security-applicability.md` or standalone `security-design.md`.

#### Scenario: Conditional obligation becomes design predicate

- GIVEN an applicable guideline has `conditional` severity
- WHEN `sdd-design` maps controls
- THEN it MUST state the predicate and required evidence
- AND unresolved true predicates MUST remain trackable downstream.

#### Scenario: Retired executor is not consulted

- GIVEN a new change reaches design
- WHEN classification inputs are resolved
- THEN `sdd-security-design` MUST NOT launch or block the workflow.

### Requirement: No-Impact Compatibility Preservation

No-impact compatibility MUST apply only to legacy artifacts. For new changes, no-impact means every category/guideline row is justified as `not-applicable` inside `design.md#secure-development-design`; it MUST NOT create or require standalone `security-design.md`, and MUST NOT be inferred from missing or retired applicability artifacts.

#### Scenario: Legacy no-impact is readable

- GIVEN an old archive lacks `security-design.md` but has valid no-impact applicability
- WHEN archive compatibility is evaluated
- THEN the archive MAY remain valid
- AND new changes MUST still require embedded `design.md` secure development rows.

#### Scenario: Invalid no-impact does not silently skip

- GIVEN no-impact matrix evidence is incomplete for a new change
- WHEN downstream routing is evaluated
- THEN the workflow MUST NOT treat missing `security-design.md` as a blocker or no-impact proof
- AND orchestration MUST route to resolve the embedded design evidence blocker.

#### Scenario: Missing applicability is not no-impact proof

- GIVEN no `security-applicability.md` artifact or executor exists for a new change
- WHEN no-impact classification is evaluated
- THEN `design.md#secure-development-design` MUST still provide row-level rationale
- AND absence of applicability data MUST NOT prove no impact.
