# sdd-security-design-workflow Specification

## Purpose

Define the conditional security design phase that translates applicable guidelines into design controls, evidence expectations, and downstream verification obligations.

## Requirements

### Requirement: Conditional Security Design Phase

The SDD workflow MUST run `sdd-security-design` only when `security-applicability.md` marks a change as security-impacting. The workflow MUST NOT require `security-design.md` for explicit no-impact changes.

#### Scenario: Applicable change requires security design

- GIVEN `security-applicability.md` marks the change as security-impacting
- WHEN technical design completes
- THEN the next required phase MUST be `sdd-security-design`
- AND task planning MUST remain blocked until `security-design.md` exists or the phase is blocked.

#### Scenario: No-impact change skips security design

- GIVEN `security-applicability.md` records no-impact evidence
- WHEN technical design completes
- THEN the workflow MUST skip `sdd-security-design`
- AND downstream phases MUST NOT treat missing `security-design.md` as a blocker.

### Requirement: Security Design Artifact Contract

`sdd-security-design` MUST create `security-design.md` for security-impacting changes. The artifact MUST map applicable guidelines to required controls, expected evidence, mandatory status, residual risks, and downstream test-design obligations.

#### Scenario: Guidelines become controls

- GIVEN applicability identified mandatory authentication or access-control guidelines
- WHEN `sdd-security-design` runs
- THEN `security-design.md` MUST define the required controls
- AND it MUST state evidence expected from apply, test-design, and verify phases.

#### Scenario: Applicability risks are carried forward

- GIVEN applicability recorded minor evidence gaps as risks
- WHEN security design is written
- THEN `security-design.md` MUST either resolve each risk or preserve it with an owner-facing evidence expectation.

### Requirement: Mandatory Evidence and Exceptions

Security design MUST identify mandatory applicable guideline evidence. Archive MUST be treated as blocked when mandatory evidence is missing unless an approved exception records approver, guideline, accepted-risk rationale, and mitigation or follow-up.

#### Scenario: Missing mandatory evidence blocks archive

- GIVEN an applicable mandatory guideline has no evidence or approved exception
- WHEN archive readiness is evaluated
- THEN archive MUST be blocked
- AND the blocker MUST name the guideline and missing evidence.

#### Scenario: Approved exception allows archive

- GIVEN mandatory evidence is missing but an exception records approver, guideline, accepted-risk rationale, and mitigation or follow-up
- WHEN archive readiness is evaluated
- THEN archive MAY proceed
- AND the exception MUST remain in the audit trail.
