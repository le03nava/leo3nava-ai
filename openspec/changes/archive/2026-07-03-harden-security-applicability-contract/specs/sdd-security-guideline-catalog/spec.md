# Delta for sdd-security-guideline-catalog

## ADDED Requirements

### Requirement: Formal Source Coverage Mapping

Each compact `SEC-*` guideline MUST declare formal corporate source coverage through stable Source IDs from the in-repo snapshot. Source coverage MUST be treated as an audit obligation, not best-effort commentary, and each mapping MUST preserve catalog snapshot identity and version metadata.

#### Scenario: Guideline maps to corporate sources

- GIVEN a compact `SEC-*` guideline is listed in the catalog
- WHEN security applicability references it
- THEN the guideline MUST expose one or more valid Source IDs
- AND the artifact MUST be able to cite those IDs as evidence refs.

#### Scenario: Source mapping is missing

- GIVEN a compact guideline lacks Source IDs
- WHEN catalog validation is performed
- THEN validation MUST fail for strict source coverage
- AND the missing mapping MUST be reported by guideline ID.

### Requirement: Operational Severity Vocabulary

The catalog MUST define operational applicability severity with only `blocking`, `conditional`, and `advisory`. `blocking` obligations MUST prevent phase success when unresolved, `conditional` obligations MUST apply when stated predicates are true, and `advisory` obligations SHOULD be preserved as downstream risk or guidance. The catalog MUST NOT use review labels such as `Menor`, `Media`, or `Mayor` for applicability blocking behavior.

#### Scenario: Blocking obligation prevents success

- GIVEN an applicable guideline has unresolved `blocking` evidence
- WHEN applicability or security design evaluates completion
- THEN the phase MUST block
- AND the blocker MUST name the guideline and missing evidence.

#### Scenario: Conditional obligation predicate is false

- GIVEN a guideline is `conditional` and its predicate is not met
- WHEN applicability evaluates the guideline
- THEN the artifact MAY mark it not applicable
- AND it MUST record the predicate rationale.

### Requirement: Catalog Validator Contract

The catalog MUST support static validation of guideline IDs, taxonomy categories, Source IDs, severity values, mandatory evidence fields, and exception fields used by `security-applicability.md`. Validation MUST compare artifact references against the same catalog snapshot identity recorded in the artifact.

#### Scenario: Artifact references current catalog snapshot

- GIVEN `security-applicability.md` records a catalog snapshot identity
- WHEN static validation checks guideline and source references
- THEN references MUST resolve within that snapshot
- AND mismatched or unknown references MUST fail validation.

#### Scenario: Advisory evidence is preserved

- GIVEN an advisory guideline is applicable
- WHEN applicability records it
- THEN downstream phases SHOULD preserve it as risk or guidance
- AND archive evidence MUST remain understandable even if it does not block.
