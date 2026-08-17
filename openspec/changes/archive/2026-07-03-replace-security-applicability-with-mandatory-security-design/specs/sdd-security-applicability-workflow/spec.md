# Delta for sdd-security-applicability-workflow

## MODIFIED Requirements

### Requirement: Always-Run Applicability Classification

The SDD workflow MUST NOT run `sdd-security-applicability` for new changes. Applicability classification MUST move into mandatory `sdd-security-design` after technical design. Legacy `security-applicability.md` artifacts MAY be read only for archived or old changes and MUST NOT create active routing authority for new changes.
(Previously: the workflow always ran `sdd-security-applicability` after specs and before design.)

#### Scenario: New change excludes applicability phase

- GIVEN a new SDD change completes `sdd-spec`
- WHEN routing is computed
- THEN the next planning phase MUST be `sdd-design`
- AND `sdd-security-applicability` MUST NOT appear in the active DAG.

#### Scenario: Legacy artifact is read-only

- GIVEN an archived change contains `security-applicability.md`
- WHEN compatibility readers inspect the archive
- THEN they MAY read it as historical evidence
- AND they MUST NOT require rerunning the phase.

### Requirement: Blocking and Risk Rules

Security applicability blockers MUST be evaluated inside `sdd-security-design` for new changes. Legacy applicability blockers MAY remain visible when reading old artifacts, but MUST NOT block a new-change DAG edge.
(Previously: applicability itself blocked design-changing ambiguity.)

#### Scenario: New security ambiguity exists

- GIVEN design context reveals ambiguous security scope
- WHEN `sdd-security-design` evaluates classification
- THEN it MUST block or record risk there
- AND no `security-applicability.md` MUST be created.

### Requirement: Artifact and Routing Contract

New changes MUST NOT produce `security-applicability.md`. The canonical classification artifact for new changes MUST be `security-design.md`, and routing MUST be `spec -> design -> security-design -> test-design`.
(Previously: `security-applicability.md` drove conditional routing to security design.)

#### Scenario: Artifact is not produced

- GIVEN a new change needs security classification
- WHEN planning artifacts are persisted
- THEN `security-design.md` MUST contain classification
- AND `security-applicability.md` MUST be absent.

### Requirement: Complete Category Decision Matrix

For new changes, the complete category/guideline matrix MUST be recorded in `security-design.md`. Legacy applicability matrices MAY be parsed for archive readability only.
(Previously: `security-applicability.md` contained the complete category decision matrix.)

#### Scenario: New matrix location

- GIVEN the security catalog exposes supported categories
- WHEN `sdd-security-design` writes its artifact
- THEN every category/guideline MUST be represented there
- AND applicability artifacts MUST NOT be authoritative.

### Requirement: Explicit No-Impact Proof

No-impact proof for new changes MUST be recorded as `not-applicable` matrix rows in `security-design.md`; absence of `security-design.md` MUST NOT prove no impact. Legacy no-impact proof remains readable only for old artifacts.
(Previously: no-impact proof allowed security design to be skipped.)

#### Scenario: No-impact still creates security design

- GIVEN all security categories are not applicable
- WHEN the new workflow runs
- THEN `security-design.md` MUST still be created
- AND `sdd-security-design` MUST NOT be skipped.

### Requirement: Static Applicability Validator

Static validation for new changes MUST target `security-design.md`. Any applicability validator MAY remain only for legacy archive checks and MUST NOT be required for new phase success.
(Previously: `security-applicability.md` required static validation before success.)

#### Scenario: Legacy validator remains compatible

- GIVEN an archive contains old applicability evidence
- WHEN legacy validation is requested
- THEN the validator MAY check old schema fields
- AND it MUST NOT affect new-change routing.
