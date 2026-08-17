# Delta for sdd-security-design-workflow

## MODIFIED Requirements

### Requirement: Enriched Applicability Consumption

For new changes, `sdd-security-design` MUST consume proposal, spec, and technical design context directly rather than relying on `security-applicability.md` or a `sdd-security-applicability` executor. Legacy applicability artifacts MAY be read only as compatibility context for old or archived changes. Mandatory `security-design.md` MUST remain the classification authority for every new change.
(Previously: new changes did not require `security-applicability.md`, but the retired executor was not explicitly excluded as a dependency.)

#### Scenario: Direct classification

- GIVEN proposal, specs, and design are readable
- WHEN security design runs for a new change
- THEN it MUST classify categories from those inputs
- AND it MUST NOT require `security-applicability.md`.

#### Scenario: Conditional obligation becomes design predicate

- GIVEN an applicable guideline has `conditional` severity
- WHEN security design maps controls
- THEN it MUST state the predicate and required evidence
- AND unresolved true predicates MUST remain trackable downstream.

#### Scenario: Retired executor is not consulted

- GIVEN a new change reaches security design
- WHEN classification inputs are resolved
- THEN `sdd-security-design` MUST NOT launch or depend on `sdd-security-applicability`
- AND the absence of that executor MUST NOT block classification.

### Requirement: No-Impact Compatibility Preservation

No-impact compatibility MUST apply only to legacy artifacts. For new changes, no-impact means every category/guideline row is justified as `not-applicable` inside `security-design.md`; it MUST NOT skip the phase or artifact, and MUST NOT be inferred from missing or retired applicability artifacts.
(Previously: no-impact compatibility was legacy-only but did not explicitly cover retired applicability sources.)

#### Scenario: Legacy no-impact is readable

- GIVEN an old archive lacks `security-design.md` but has valid no-impact applicability
- WHEN archive compatibility is evaluated
- THEN the archive MAY remain valid
- AND new changes MUST still require `security-design.md`.

#### Scenario: Invalid no-impact does not silently skip

- GIVEN no-impact matrix evidence is incomplete for a new change
- WHEN downstream routing is evaluated
- THEN the workflow MUST NOT treat missing `security-design.md` as no-impact proof
- AND orchestration MUST route to resolve the security-design blocker.

#### Scenario: Missing applicability is not no-impact proof

- GIVEN no `security-applicability.md` artifact or executor exists for a new change
- WHEN no-impact classification is evaluated
- THEN `security-design.md` MUST still provide row-level rationale
- AND absence of applicability data MUST NOT prove no impact.
