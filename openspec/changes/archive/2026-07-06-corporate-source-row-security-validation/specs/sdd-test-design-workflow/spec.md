# Delta for sdd-test-design-workflow

## ADDED Requirements

### Requirement: Source Row Test Planning

`sdd-test-design` MUST plan static, manual, or automated checks for every applicable corporate source row declared by `design.md#secure-development-design`. Each planned check MUST reference Source ID, compact `SEC-*` mapping, expected evidence, mandatory/advisory status, and accepted no-runner limitations.

#### Scenario: Applicable source row receives a check

- GIVEN design marks a source row applicable
- WHEN test design is produced
- THEN `test-design.md` MUST include a planned check or justified non-test evidence
- AND the check MUST cite Source ID and compact `SEC-*` mapping.

#### Scenario: No runtime runner exists

- GIVEN repository testing capabilities report no runner
- WHEN source-row checks are planned
- THEN test design MUST use static or manual evidence where appropriate
- AND it MUST report unavailable runtime automation explicitly.

#### Scenario: Mandatory row lacks planned evidence

- GIVEN an applicable mandatory source row has no check or evidence plan
- WHEN `sdd-test-design` validates readiness
- THEN the phase MUST block
- AND the blocker MUST name the Source ID and missing evidence plan.

### Requirement: Source Row N/A and Warning Evidence

`sdd-test-design` MUST preserve evidence expectations for `N/A` source rows and warning-only rows. `N/A` requires proof of irrelevance; warnings require planned observation evidence but MUST NOT block solely when mandatory evidence is complete.

#### Scenario: N/A row is planned for review

- GIVEN design marks a Source ID `N/A`
- WHEN test design maps security evidence
- THEN it MUST include the irrelevance evidence to be checked
- AND missing justification MUST remain blocking.

#### Scenario: Warning-only row is tracked

- GIVEN a source row has warning classification only
- WHEN test design records checks
- THEN it MUST preserve the warning and expected evidence
- AND it MUST allow downstream routing when no blockers exist.
