# Delta for sdd-security-applicability-workflow

## MODIFIED Requirements

### Requirement: Legacy-Only Applicability Classification

The SDD workflow MUST NOT provide, launch, or require a repo-local `sdd-security-applicability` executor or skill for new changes. Applicability classification MUST live in mandatory `sdd-security-design`. Legacy `security-applicability.md` artifacts MAY be read as historical data for archived or old changes and MUST NOT create active routing authority, artifact production, or executor availability for new changes.
(Previously: new changes skipped the phase, but legacy wording still allowed an executor/skill to exist as a compatibility reader.)

#### Scenario: New change excludes applicability phase

- GIVEN a new SDD change completes `sdd-spec`
- WHEN routing is computed
- THEN the next planning phase MUST be `sdd-design`
- AND `sdd-security-applicability` MUST NOT appear in the active DAG.

#### Scenario: Legacy artifact is read-only

- GIVEN an archived change contains `security-applicability.md`
- WHEN compatibility readers inspect the archive
- THEN they MAY read it as historical evidence
- AND they MUST NOT require rerunning the phase.

#### Scenario: Executor and skill are absent

- GIVEN active repo-local SDD agents and skills are enumerated
- WHEN launchable security phases are resolved
- THEN no `sdd-security-applicability` executor or skill MUST be offered
- AND legacy artifacts MUST remain readable without launching one.

### Requirement: Artifact and Routing Contract

New changes MUST NOT produce `security-applicability.md`. The canonical classification artifact for new changes MUST be `security-design.md`, and routing MUST be `spec -> design -> security-design -> test-design`. Legacy compatibility readers MAY continue resolving old applicability artifact paths as data references without making them authoritative or mapping them to a runnable phase.
(Previously: the persistence scenario named `sdd-security-applicability` as an executor that reads or writes the artifact.)

#### Scenario: Artifact is not produced

- GIVEN a new change needs security classification
- WHEN planning artifacts are persisted
- THEN `security-design.md` MUST contain classification
- AND `security-applicability.md` MUST be absent.

#### Scenario: Legacy data reference is resolved

- GIVEN an old change has `security-applicability.md`
- WHEN a reader resolves historical artifacts
- THEN the path MAY be resolved as legacy data
- AND no runnable applicability executor or skill MUST be required.
