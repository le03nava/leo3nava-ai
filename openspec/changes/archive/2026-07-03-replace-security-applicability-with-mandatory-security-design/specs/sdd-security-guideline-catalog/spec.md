# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain the initial corporate security guideline catalog as an in-repo snapshot based on the user-provided text. The catalog MUST preserve source text, identifiers, version metadata, and matrix vocabulary needed by `security-design.md` and `review-security-report.md` until an official external versioned source replaces it.
(Previously: the catalog primarily served security applicability, security design, verification, and archive checks.)

#### Scenario: Catalog snapshot is available

- GIVEN an SDD security phase needs guideline context
- WHEN it reads the catalog
- THEN it MUST find guideline identifiers, source snapshot metadata, and applicable summaries
- AND it MUST support mandatory security-design and security-review matrices.

### Requirement: Mandatory Evidence Model

Each catalog guideline MUST declare whether it is mandatory when applicable and what evidence is expected across design, implementation, test/static/manual checks, security review, archive, or approved exceptions.
(Previously: evidence expectations were not explicitly tied to security review.)

#### Scenario: Mandatory guideline has evidence expectations

- GIVEN a guideline is mandatory when applicable
- WHEN security design consumes it
- THEN the guideline MUST provide expected evidence types
- AND security review MUST be able to verify coverage.

### Requirement: Review Control Cross-References

The security guideline catalog MUST support cross-references from `sdd-review` controls and `sdd-review-security` rows to applicable guideline identifiers. These references MUST help reviewers cite standards without transferring guideline authority away from security design and security review.
(Previously: only general review cross-references were described.)

#### Scenario: Review cites a security guideline

- GIVEN a review checklist control evaluates a security concern
- WHEN the control maps to a catalog guideline
- THEN the review control SHOULD cite that guideline identifier
- AND security design/review MUST remain authoritative for required controls and final security verdict.

### Requirement: Review-Safe Security Evidence

The catalog SHOULD identify evidence types suitable for review rows, including implementation reference, static inspection, test evidence, approved exception, or N/A evidence. N/A MUST include evidence proving the category, platform, framework, API, or data class is irrelevant.
(Previously: N/A guidance focused on platform-specific review controls.)

#### Scenario: Security row is N/A

- GIVEN a guideline row is marked N/A
- WHEN security review validates it
- THEN evidence MUST show why it is out of scope
- AND observations MUST explain why no security control is required.

### Requirement: Catalog Boundary Preservation

The catalog MUST remain the source for security guideline identifiers, taxonomy, mandatory evidence expectations, exception fields, lifecycle statuses, and matrix vocabulary. `sdd-review` MAY reference catalog entries but MUST NOT duplicate or redefine guideline text.
(Previously: conflicts resolved in favor of applicability/design outputs.)

#### Scenario: Catalog authority is preserved

- GIVEN security design and security review reference a guideline
- WHEN downstream verification compares evidence
- THEN catalog identifiers and statuses MUST remain consistent
- AND conflicts MUST be resolved in favor of `security-design.md` plus `review-security-report.md`.

### Requirement: Catalog Validator Contract

The catalog MUST support static validation of guideline IDs, taxonomy categories, Source IDs, severity values, lifecycle statuses, matrix values, mandatory evidence fields, and exception fields used by `security-design.md` and `review-security-report.md`.
(Previously: validation targeted `security-applicability.md`.)

#### Scenario: Artifact references current catalog snapshot

- GIVEN a security artifact records a catalog snapshot identity
- WHEN static validation checks references
- THEN references MUST resolve within that snapshot
- AND mismatched or unknown references MUST fail validation.

## ADDED Requirements

### Requirement: Security Matrix Vocabulary

The catalog MUST define matrix-facing values `Yes`, `No`, and `N/A` plus lifecycle statuses `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, and `blocked`.

#### Scenario: Matrix row uses valid vocabulary

- GIVEN security design or security review writes a matrix row
- WHEN validation runs
- THEN its answer and lifecycle status MUST use catalog vocabulary
- AND unsupported values MUST fail validation.
