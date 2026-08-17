# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain the corporate security guideline catalog as an in-repo snapshot based on the user-provided text. The catalog MUST preserve source text, identifiers, version metadata, and matrix vocabulary needed by embedded secure development design in `design.md#secure-development-design` and `review-security-report.md` until an official external versioned source replaces it.
(Previously: the catalog scope also named standalone security-design compatibility.)

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

### Requirement: Catalog Boundary Preservation

The catalog MUST remain the source for security guideline identifiers, taxonomy, mandatory evidence expectations, exception fields, lifecycle statuses, and matrix vocabulary. `sdd-review` MAY reference catalog entries but MUST NOT duplicate, redefine, or replace guideline text. Active security authority MUST remain with `design.md#secure-development-design` plus `review-security-report.md`.
(Previously: conflict wording implied broader active authority boundaries.)

#### Scenario: Catalog authority is preserved

- GIVEN embedded design and security review reference a guideline
- WHEN downstream verification compares evidence
- THEN catalog identifiers and statuses MUST remain consistent
- AND conflicts MUST be resolved in favor of `design.md#secure-development-design` plus `review-security-report.md`.
