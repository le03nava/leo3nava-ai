# sdd-security-applicability-workflow Specification

## Purpose

Define legacy security applicability compatibility after new-change classification moved into `design.md#secure-development-design`.

## Requirements

### Requirement: Legacy-Only Applicability Classification

The SDD workflow MUST NOT provide, launch, or require a repo-local `sdd-security-applicability` executor or skill for new changes. Applicability classification MUST live in `design.md#secure-development-design`. Historical `security-applicability.md` artifacts MAY be read by status, persistence, or archive display paths as data and MUST NOT create active routing authority, artifact production, or executor availability for new changes.

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

### Requirement: Blocking and Risk Rules

Security applicability blockers MUST be evaluated inside `sdd-design` secure development design for new changes. Legacy applicability blockers MAY remain visible when reading old artifacts, but MUST NOT block a new-change DAG edge.

#### Scenario: Design-changing information is missing

- GIVEN security-relevant scope is ambiguous in a way that could change required controls
- WHEN `sdd-design` evaluates classification
- THEN it MUST block or record risk there
- AND no `security-applicability.md` MUST be created.

#### Scenario: Minor evidence gap exists

- GIVEN a security impact is known but a non-decisive detail is incomplete
- WHEN `sdd-design` evaluates classification
- THEN it SHOULD continue when the gap is non-blocking
- AND it MUST record the gap as an embedded secure design risk.

### Requirement: Artifact and Routing Contract

New changes MUST NOT produce `security-applicability.md`. The canonical classification artifact for new changes MUST be `design.md#secure-development-design`, and routing MUST be `spec -> design -> test-design`. Legacy compatibility readers MAY continue resolving old applicability artifact paths as data references without making them authoritative or mapping them to a runnable phase.

#### Scenario: Artifact is not produced

- GIVEN a new change needs security classification
- WHEN planning artifacts are persisted
- THEN `design.md#secure-development-design` MUST contain classification
- AND `security-applicability.md` MUST be absent.

#### Scenario: Historical data reference is resolved

- GIVEN an old change has `security-applicability.md`
- WHEN a reader resolves historical artifacts
- THEN the path MAY be resolved as historical data
- AND no runnable applicability executor or skill MUST be required.

### Requirement: Narrative Design Classification and Review-Security Matrix Ownership

For new changes, `design.md#secure-development-design` MUST record changed-surface classification and applicable category rules as narrative Markdown. The exhaustive category/guideline matrix, Source ID expansion, lifecycle row status, and `N/A` decisions MUST be produced only by canonical `review-security-report.json`. Legacy applicability matrices MAY be parsed for archive readability only.

#### Scenario: Applicable categories are narrated

- GIVEN the catalog exposes supported taxonomy categories
- WHEN `sdd-design` writes its artifact
- THEN applicable categories MUST be represented as narrative rules with evidence expectations
- AND omitted categories MUST remain reviewable by `sdd-review-security`
- AND applicability artifacts MUST NOT be authoritative.

#### Scenario: Unknown decision is design-changing

- GIVEN a category is marked `unknown` with `blocking` severity
- WHEN `sdd-design` evaluates changed-surface classification
- THEN it MUST return blocked
- AND it MUST identify the missing evidence or decision.

### Requirement: Explicit No-Impact Rationale

No-impact proof for new changes MUST be recorded as narrative changed-surface rationale in `design.md#secure-development-design`; absence of standalone `security-design.md` MUST NOT prove or disprove impact. Exhaustive `not-applicable` / `N/A` matrix rows belong to canonical `review-security-report.json`. Legacy no-impact proof remains readable only for old artifacts.

#### Scenario: Valid no-impact artifact

- GIVEN the changed surface does not touch security-relevant behavior, data, or operational evidence
- WHEN the new workflow runs
- THEN `design.md#secure-development-design` MUST still be created
- AND `sdd-design` MUST NOT skip the embedded section
- AND exhaustive `N/A` rows MUST NOT be copied into design.

#### Scenario: Absence of evidence is insufficient

- GIVEN a change has no mapped guidelines but lacks changed-surface rationale for one relevant category
- WHEN `sdd-design` evaluates classification
- THEN `design.md#secure-development-design` MUST NOT classify the change as no-impact
- AND the missing proof MUST be recorded as a blocker or risk by severity.
