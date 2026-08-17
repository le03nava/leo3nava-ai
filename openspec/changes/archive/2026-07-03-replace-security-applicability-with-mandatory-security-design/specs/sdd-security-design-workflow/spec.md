# Delta for sdd-security-design-workflow

## MODIFIED Requirements

### Requirement: Conditional Security Design Phase

The SDD workflow MUST run `sdd-security-design` for every new change after `sdd-design` succeeds and before `sdd-test-design`. The phase MUST always create `security-design.md`, including for no-impact changes.
(Previously: security design ran only for security-impacting changes and was skipped for no-impact changes.)

#### Scenario: Every new change requires security design

- GIVEN `sdd-design` has completed for a new change
- WHEN routing is computed
- THEN the next required phase MUST be `sdd-security-design`
- AND `sdd-test-design` MUST remain blocked until `security-design.md` exists or the phase blocks.

#### Scenario: No-impact still creates artifact

- GIVEN no security guideline applies
- WHEN `sdd-security-design` runs
- THEN it MUST write `security-design.md`
- AND the matrix rows MUST record `not-applicable` evidence.

### Requirement: Security Design Artifact Contract

`sdd-security-design` MUST create `security-design.md` for every new change. The artifact MUST own classification, catalog snapshot identity, category/guideline matrix, controls, expected evidence, lifecycle statuses, risks, exceptions, and archive gate expectations while delegating backend persistence mechanics to the shared persistence authority.
(Previously: artifact creation was conditional and consumed `security-applicability.md`.)

#### Scenario: Classification and controls are unified

- GIVEN design context is available
- WHEN `sdd-security-design` writes the artifact
- THEN it MUST classify security impact inside `security-design.md`
- AND it MUST map applicable guidelines to controls and expected evidence.

### Requirement: Mandatory Evidence and Exceptions

Security design MUST identify mandatory evidence for every applicable guideline. Archive MUST be blocked when mandatory evidence is missing unless an approved exception records approver, guideline, accepted-risk rationale, and mitigation or follow-up.
(Previously: mandatory evidence applied only when conditional security design existed.)

#### Scenario: Missing mandatory evidence blocks archive

- GIVEN an applicable mandatory guideline has no evidence or approved exception
- WHEN archive readiness is evaluated
- THEN archive MUST be blocked
- AND the blocker MUST name the guideline and missing evidence.

#### Scenario: Approved exception allows archive

- GIVEN mandatory evidence is missing but an exception is complete
- WHEN archive readiness is evaluated
- THEN archive MAY proceed
- AND the exception MUST remain in the audit trail.

### Requirement: Enriched Applicability Consumption

For new changes, `sdd-security-design` MUST consume proposal, spec, and technical design context directly rather than relying on `security-applicability.md`. Legacy applicability artifacts MAY be read only as compatibility context for old or archived changes.
(Previously: security design consumed enriched applicability fields for security-impacting changes.)

#### Scenario: Direct classification

- GIVEN proposal, specs, and design are readable
- WHEN security design runs for a new change
- THEN it MUST classify categories from those inputs
- AND it MUST NOT require `security-applicability.md`.

### Requirement: No-Impact Compatibility Preservation

No-impact compatibility MUST apply only to legacy artifacts. For new changes, no-impact means every category/guideline row is justified as `not-applicable` inside `security-design.md`; it MUST NOT skip the phase or artifact.
(Previously: valid no-impact proof skipped `sdd-security-design`.)

#### Scenario: Legacy no-impact is readable

- GIVEN an old archive lacks `security-design.md` but has valid no-impact applicability
- WHEN archive compatibility is evaluated
- THEN the archive MAY remain valid
- AND new changes MUST still require `security-design.md`.
