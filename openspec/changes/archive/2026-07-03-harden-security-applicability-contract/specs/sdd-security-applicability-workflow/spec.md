# Delta for sdd-security-applicability-workflow

## ADDED Requirements

### Requirement: Complete Category Decision Matrix

`security-applicability.md` MUST include the catalog identity, supported taxonomy version, and a decision matrix that evaluates every supported security taxonomy category as `applicable`, `not-applicable`, or `unknown`. Each row MUST include rationale, evidence references, and operational severity using only `blocking`, `conditional`, or `advisory`.

#### Scenario: Every category is evaluated

- GIVEN the catalog exposes supported taxonomy categories
- WHEN `sdd-security-applicability` writes `security-applicability.md`
- THEN every category MUST appear exactly once in the decision matrix
- AND missing or duplicate categories MUST make the artifact invalid.

#### Scenario: Unknown decision is design-changing

- GIVEN a category is marked `unknown` with `blocking` severity
- WHEN the applicability phase completes
- THEN it MUST return blocked
- AND it MUST identify the missing evidence or decision.

### Requirement: Explicit No-Impact Proof

A no-impact classification MUST be supported by positive no-impact proof and MUST NOT be inferred from absent security evidence. No-impact proof MUST show every supported category as `not-applicable`, include rationale and evidence references, and have no design-changing unknowns.

#### Scenario: Valid no-impact artifact

- GIVEN every category is `not-applicable` with evidence and rationale
- WHEN the artifact classifies the change as no-impact
- THEN downstream routing MUST treat the artifact as complete
- AND `sdd-security-design` MUST remain skipped.

#### Scenario: Absence of evidence is insufficient

- GIVEN a change has no mapped guidelines but lacks no-impact rationale for one category
- WHEN applicability is evaluated
- THEN the artifact MUST NOT classify the change as no-impact
- AND the missing proof MUST be recorded as a blocker or risk by severity.

### Requirement: Supported Applicability Overrides

`openspec/config.yaml` MAY define `rules.security-applicability` overrides only for extra prompts, stricter source coverage, validator mode, and stricter category severity. Overrides MUST NOT disable required categories, weaken formal source coverage, downgrade `blocking` obligations, or bypass no-impact proof.

#### Scenario: Safe override is applied

- GIVEN config adds an extra design-changing unknown prompt
- WHEN applicability evaluates a change
- THEN the prompt MUST be considered in the matrix
- AND the artifact MUST record the override source.

#### Scenario: Unsafe weakening is rejected

- GIVEN config attempts to disable source coverage or remove a required category
- WHEN applicability loads overrides
- THEN the override MUST be rejected
- AND the phase MUST continue with the stricter base contract or block if ambiguity remains.

### Requirement: Static Applicability Validator

The workflow MUST require automatic static validation for `security-applicability.md` artifacts before reporting success. The validator MUST check required schema fields, classification/routing consistency, matrix completeness, no-impact proof, guideline ID validity, source reference validity, supported overrides, and severity vocabulary.

#### Scenario: Validator accepts a complete artifact

- GIVEN an artifact contains valid fields, matrix rows, guideline IDs, source refs, and routing
- WHEN static validation runs
- THEN validation MUST pass
- AND the artifact MUST record validation metadata.

#### Scenario: Validator blocks invalid artifact

- GIVEN an artifact omits a category or uses an unsupported severity
- WHEN static validation runs
- THEN validation MUST fail
- AND the phase MUST NOT report success until the artifact is corrected.
