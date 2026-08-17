# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain the initial corporate security guideline catalog as an in-repo snapshot based on the user-provided text. The catalog MUST preserve source text, identifiers, version metadata, and matrix vocabulary needed by embedded secure development design in `design.md` and `review-security-report.md` until an official external versioned source replaces it.
(Previously: the catalog described support for `security-design.md` and `review-security-report.md`.)

#### Scenario: Catalog snapshot is available

- GIVEN an SDD design or security review phase needs guideline context
- WHEN it reads the catalog
- THEN it MUST find guideline identifiers, source snapshot metadata, and applicable summaries
- AND it MUST support embedded secure design and security-review matrices.

#### Scenario: Catalog source changes later

- GIVEN an official external guideline source becomes available
- WHEN the catalog is updated
- THEN the change MUST preserve migration/audit metadata
- AND existing archived evidence MUST remain understandable.

### Requirement: Catalog Validator Contract

The catalog MUST support static validation of guideline IDs, taxonomy categories, Source IDs, severity values, lifecycle statuses, matrix values, mandatory evidence fields, and exception fields used by embedded `design.md` secure development rows and `review-security-report.md`. Validation MUST compare artifact references against the same catalog snapshot identity recorded in the artifact.
(Previously: validation targeted `security-design.md` and `review-security-report.md`.)

#### Scenario: Artifact references current catalog snapshot

- GIVEN a security artifact records a catalog snapshot identity
- WHEN static validation checks guideline and source references
- THEN references MUST resolve within that snapshot
- AND mismatched or unknown references MUST fail validation.

#### Scenario: Advisory evidence is preserved

- GIVEN an advisory guideline is applicable
- WHEN design records it
- THEN downstream phases SHOULD preserve it as risk or guidance
- AND archive evidence MUST remain understandable even if it does not block.

### Requirement: Security Matrix Vocabulary

The catalog MUST define matrix-facing values `Yes`, `No`, and `N/A` plus lifecycle statuses `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, and `blocked` for embedded design rows and review-security rows.
(Previously: vocabulary applied to security design or security review rows without embedding in `design.md`.)

#### Scenario: Matrix row uses valid vocabulary

- GIVEN design or security review writes a matrix row
- WHEN validation runs
- THEN its answer and lifecycle status MUST use catalog vocabulary
- AND unsupported values MUST fail validation.
