# Delta for sdd-security-guideline-catalog

## ADDED Requirements

### Requirement: Review Control Cross-References

The security guideline catalog MUST support cross-references from `sdd-review` controls to applicable guideline identifiers. These references MUST help reviewers cite security standards in `review-report.md` without transferring security applicability, security design, or exception authority to review.

#### Scenario: Review cites a security guideline

- GIVEN a review checklist control evaluates a security concern
- WHEN the control maps to a catalog guideline
- THEN the review control SHOULD cite that guideline identifier
- AND security applicability/design MUST remain authoritative for applicability and required controls.

### Requirement: Review-Safe Security Evidence

The catalog SHOULD identify evidence types suitable for code review rows, including implementation reference, static inspection, test evidence, approved exception, or N/A evidence. N/A for platform-specific security controls MUST include evidence that the platform, framework, API, or data class is irrelevant to the change.

#### Scenario: Platform-specific security control is N/A

- GIVEN a security review control applies only to an unused platform
- WHEN review marks the control N/A
- THEN the evidence MUST show the platform is out of scope
- AND the comment MUST explain why no security design control is required.

### Requirement: Catalog Boundary Preservation

The catalog MUST remain the source for security guideline identifiers, taxonomy, mandatory evidence expectations, and exception fields. `sdd-review` MAY reference catalog entries but MUST NOT duplicate or redefine catalog guideline text.

#### Scenario: Catalog authority is preserved

- GIVEN review and security design both reference a guideline
- WHEN downstream verification compares evidence
- THEN catalog identifiers MUST remain consistent
- AND conflicts MUST be resolved in favor of security applicability/design outputs.
