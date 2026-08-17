# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-design.md` for every change. The artifact MUST map spec scenarios, design risks, and applicable narrative category rules from `design.md#secure-development-design` to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document justified no-impact assessments from design classification. It MUST NOT require design to include YAML, schema fields, compact controls, Source IDs, matrices, or all-row `N/A` evidence. Common persistence behavior MUST remain delegated to the shared persistence authority.
(Previously: test-design consumed secure development rows as if they could include matrix-style evidence.)

#### Scenario: Behavior-impacting change

- GIVEN specs or design describe behavior, contracts, routing, or compatibility changes
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST list planned checks linked to those inputs
- AND each check MUST include type, severity, and expected evidence.

#### Scenario: Applicable category rules are consumed

- GIVEN `design.md` lists applicable secure development category rules
- WHEN `sdd-test-design` runs
- THEN `test-design.md` MUST include checks or justified non-test evidence for mandatory rules
- AND blocked security obligations MUST remain blockers.

#### Scenario: No-impact assessment is handled

- GIVEN design classifies the change as no security impact with changed-surface rationale
- WHEN `sdd-test-design` runs
- THEN it MUST cite that assessment
- AND it MUST still produce `test-design.md` without requiring YAML, schema, matrix, or all-row `N/A` content.

### Requirement: Source Row Test Planning

`sdd-test-design` MUST plan static, manual, or automated checks from applicable narrative category rules and changed-surface context in `design.md#secure-development-design`. It MUST consume catalog snapshot/context only as needed to understand categories, evidence owners, lifecycle states, residual risks, exceptions, and safe-evidence policy. It MUST NOT require design to carry compact control fields, Source IDs, or `N/A` rows for omitted controls, and MUST preserve `sdd-review-security` as the exhaustive machine-readable validation owner.
(Previously: source-row test planning consumed compact/source mappings from design.)

#### Scenario: Applicable category receives checks

- GIVEN design defines Sensitive Data/PAN development rules
- WHEN test design is produced
- THEN `test-design.md` MUST include planned checks or justified non-test evidence
- AND each check MUST cite the narrative rule or category context.

#### Scenario: No runtime runner exists

- GIVEN repository testing capabilities report no runner
- WHEN security checks are planned
- THEN test design MUST use static or manual evidence where appropriate
- AND it MUST report unavailable runtime automation explicitly.

#### Scenario: Mandatory coverage lacks planned evidence

- GIVEN an applicable mandatory category rule has no check or evidence plan
- WHEN `sdd-test-design` validates readiness
- THEN the phase MUST block
- AND the blocker MUST name the affected category rule and missing evidence plan.

### Requirement: Source Row N/A and Warning Evidence

`sdd-test-design` MUST preserve evidence expectations for applicable narrative rules, warning-only coverage, residual risks, and exceptions from design. It MUST NOT create or require exhaustive `N/A` evidence for every non-applicable compact control or Source ID. Exhaustive `N/A` decisions, compact/Source ID matrices, and missed-applicable validation MUST remain owned by `review-security-report.md`.
(Previously: test-design preserved N/A expectations from design for source coverage.)

#### Scenario: Omitted coverage remains reviewable

- GIVEN design omits a category as non-applicable
- WHEN test design maps security evidence
- THEN it MUST rely on design classification context only
- AND review-security MUST later validate whether the omission was safe.

#### Scenario: Warning-only coverage is tracked

- GIVEN applicable narrative rules have warning classification only
- WHEN test design records checks
- THEN it MUST preserve the warning and expected evidence
- AND it MUST allow downstream routing when no blockers exist.
