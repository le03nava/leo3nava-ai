# Delta for sdd-security-design-workflow

## MODIFIED Requirements

### Requirement: Security Design Artifact Contract

`sdd-security-design` MUST create `security-design.md` for security-impacting changes. The artifact MUST map applicable guidelines to required controls, expected evidence, mandatory status, residual risks, and downstream test-design obligations. The phase artifact contract MUST preserve conditional artifact creation and existing artifact identity while delegating common artifact-store mode semantics, artifact resolution, and persistence verification to the shared persistence authority.
(Previously: The requirement defined conditional security-design content, but did not state the shared persistence boundary.)

#### Scenario: Guidelines become controls

- GIVEN applicability identified mandatory authentication or access-control guidelines
- WHEN `sdd-security-design` runs
- THEN `security-design.md` MUST define the required controls
- AND it MUST state evidence expected from apply, test-design, and verify phases.

#### Scenario: Applicability risks are carried forward

- GIVEN applicability recorded minor evidence gaps as risks
- WHEN security design is written
- THEN `security-design.md` MUST either resolve each risk or preserve it with an owner-facing evidence expectation.

#### Scenario: Conditional artifact behavior is preserved

- GIVEN `security-applicability.md` records explicit no-impact evidence
- WHEN the workflow evaluates required artifacts
- THEN it MUST NOT require `security-design.md`
- AND this compatibility MUST hold across supported artifact-store modes.

#### Scenario: Persistence boundary is delegated

- GIVEN `sdd-security-design` persists or resolves `security-design.md`
- WHEN backend behavior is needed
- THEN the phase MUST follow the shared persistence authority
- AND phase-local rules MUST only define conditional creation, contents, and routing obligations.
