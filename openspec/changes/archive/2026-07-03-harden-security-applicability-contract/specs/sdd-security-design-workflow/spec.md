# Delta for sdd-security-design-workflow

## ADDED Requirements

### Requirement: Enriched Applicability Consumption

`sdd-security-design` MUST consume enriched applicability fields for security-impacting changes, including catalog snapshot identity, evaluated categories, decision matrix rows, source evidence refs, operational severity, and validation metadata. It MUST translate applicable `blocking` and `conditional` obligations into controls, evidence expectations, risks, or approved exceptions.

#### Scenario: Security-impacting artifact provides enriched fields

- GIVEN `security-applicability.md` marks a change as security-impacting
- WHEN `sdd-security-design` runs
- THEN it MUST use the decision matrix and source refs to build controls
- AND it MUST preserve catalog snapshot identity in `security-design.md`.

#### Scenario: Conditional obligation becomes design predicate

- GIVEN an applicable guideline has `conditional` severity
- WHEN security design maps controls
- THEN it MUST state the predicate and required evidence
- AND unresolved true predicates MUST remain trackable downstream.

### Requirement: No-Impact Compatibility Preservation

The workflow MUST preserve no-impact routing compatibility. Enriched applicability fields MUST NOT cause `security-design.md` to be required when `security-applicability.md` contains valid explicit no-impact proof and routes security design as skipped.

#### Scenario: Valid no-impact still skips security design

- GIVEN applicability records complete no-impact proof and validation metadata
- WHEN technical design completes
- THEN the workflow MUST skip `sdd-security-design`
- AND missing `security-design.md` MUST NOT block tasks, verify, or archive.

#### Scenario: Invalid no-impact does not silently skip

- GIVEN no-impact proof is incomplete or validator metadata reports failure
- WHEN downstream routing is evaluated
- THEN the workflow MUST NOT rely on the no-impact classification
- AND orchestration MUST route to resolve the applicability blocker before design proceeds.
