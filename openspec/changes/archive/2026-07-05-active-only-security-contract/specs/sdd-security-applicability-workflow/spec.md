# Delta for sdd-security-applicability-workflow

## MODIFIED Requirements

### Requirement: Legacy-Only Applicability Classification

The SDD workflow MUST NOT provide, launch, or require a repo-local `sdd-security-applicability` executor or skill for new changes. Applicability classification MUST live in `design.md#secure-development-design`. Historical `security-applicability.md` artifacts MAY be read by status, persistence, or archive display paths as data and MUST NOT create active routing authority, artifact production, or executor availability for new changes.
(Previously: the requirement used broader legacy compatibility wording.)

#### Scenario: New change excludes applicability phase

- GIVEN a new SDD change completes `sdd-spec`
- WHEN routing is computed
- THEN the next planning phase MUST be `sdd-design`
- AND `sdd-security-applicability` MUST NOT appear in the active DAG.

#### Scenario: Historical artifact is read-only

- GIVEN an archived change contains `security-applicability.md`
- WHEN compatibility readers inspect the archive
- THEN they MAY read it as historical evidence
- AND they MUST NOT require rerunning the phase.

#### Scenario: Executor and skill are absent

- GIVEN active repo-local SDD agents and skills are enumerated
- WHEN launchable security phases are resolved
- THEN no `sdd-security-applicability` executor or skill MUST be offered
- AND historical artifacts MUST remain readable without launching one.

## REMOVED Requirements

### Requirement: Supported Applicability Overrides

(Reason: New-change secure design overrides belong to `design.md#secure-development-design`; keeping an applicability-specific override contract makes the retired workflow look active.)
(Migration: Preserve only parser/read behavior for old artifacts in status or persistence contracts.)

### Requirement: Static Applicability Validator

(Reason: Static validation for active new changes MUST target embedded secure design and review-security evidence, not an applicability workflow.)
(Migration: Use catalog-backed `design.md#secure-development-design` and `review-security-report.md` checks for active changes.)
