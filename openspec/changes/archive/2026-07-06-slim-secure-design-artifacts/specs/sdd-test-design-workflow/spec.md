# Delta for sdd-test-design-workflow

## MODIFIED Requirements

### Requirement: Source Row Test Planning

`sdd-test-design` MUST plan static, manual, or automated checks from the slim `design.md#secure-development-design` coverage contract. It MUST consume catalog snapshot identity/path, `expectedSourceIdCount: 155`, section/group coverage, compact `SEC-*` mappings, applicability, evidence owners, lifecycle states, and N/A/exception policy. It MUST NOT require design to carry the full 155-row matrix, but MUST preserve enough planned evidence for review-security to expand every Source ID exactly once.
(Previously: test-design planned checks for every applicable corporate source row declared directly by design.)

#### Scenario: Applicable source group receives checks

- GIVEN design marks a source group or compact mapping applicable
- WHEN test design is produced
- THEN `test-design.md` MUST include planned checks or justified non-test evidence
- AND each check MUST cite the catalog reference and compact `SEC-*` mapping.

#### Scenario: No runtime runner exists

- GIVEN repository testing capabilities report no runner
- WHEN source-row checks are planned
- THEN test design MUST use static or manual evidence where appropriate
- AND it MUST report unavailable runtime automation explicitly.

#### Scenario: Mandatory coverage lacks planned evidence

- GIVEN applicable mandatory source coverage has no check or evidence plan
- WHEN `sdd-test-design` validates readiness
- THEN the phase MUST block
- AND the blocker MUST name the affected Source ID, group, or compact mapping and missing evidence plan.

### Requirement: Source Row N/A and Warning Evidence

`sdd-test-design` MUST preserve evidence expectations for `N/A` source coverage and warning-only coverage from the slim design contract. `N/A` requires proof of irrelevance; warnings require planned observation evidence but MUST NOT block solely when mandatory evidence is complete.
(Previously: evidence expectations were preserved for explicit `N/A` source rows and warning-only rows.)

#### Scenario: N/A coverage is planned for review

- GIVEN design marks a source group or compact mapping `N/A`
- WHEN test design maps security evidence
- THEN it MUST include the irrelevance evidence to be checked
- AND missing justification MUST remain blocking.

#### Scenario: Warning-only coverage is tracked

- GIVEN source coverage has warning classification only
- WHEN test design records checks
- THEN it MUST preserve the warning and expected evidence
- AND it MUST allow downstream routing when no blockers exist.
