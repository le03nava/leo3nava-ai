# sdd-execution-persistence-contracts Specification

## Purpose

Define the authoritative boundary between shared SDD execution contracts, shared persistence contracts, and phase-local artifact contracts.

## Requirements

### Requirement: Authoritative Persistence Boundary

The SDD contract set MUST define one shared persistence authority for artifact-store modes, backend read/write semantics, artifact resolution, hybrid conflict handling, and persistence verification. Phase and execution contracts MUST reference that authority instead of redefining detailed mode behavior.

#### Scenario: Shared persistence owns mode behavior

- GIVEN an SDD phase runs in `engram`, `openspec`, `hybrid`, or `none` mode
- WHEN the executor resolves read or write behavior
- THEN it MUST follow the shared persistence contract for backend semantics
- AND phase-local text MUST NOT introduce contradictory mode rules.

#### Scenario: Backend convention files remain scoped

- GIVEN backend convention files document Engram or OpenSpec details
- WHEN those files describe artifact storage
- THEN they SHOULD stay backend-specific references
- AND they MUST NOT compete with the shared persistence authority.

### Requirement: Execution Contract Boundary

The shared SDD phase contract MUST own executor boundary, supplemental skill loading, return envelope shape, routing-token conventions, artifact naming reminders, and review workload guard. It MUST NOT duplicate detailed persistence algorithms already owned by the persistence authority.

#### Scenario: Executor returns a stable envelope

- GIVEN a phase completes, partially completes, or blocks
- WHEN it returns to the orchestrator
- THEN it MUST use the shared Section D envelope fields
- AND routing tokens MUST remain normalizable through existing native/status mappings.

### Requirement: Phase Artifact Contracts

Each phase skill MUST keep a compact artifact contract that states required inputs, produced artifacts, artifact keys, OpenSpec paths, mutations, conditional behavior, and success/block routing. Phase contracts MAY define phase-specific validation or mutation semantics, but MUST delegate common persistence mechanics to the shared persistence authority.

#### Scenario: Phase-specific mutation is preserved

- GIVEN `sdd-apply`, `sdd-archive`, or `sdd-init` mutates existing artifacts or project state
- WHEN its phase contract is written
- THEN mutation semantics MUST remain explicit in that phase
- AND generic backend persistence mechanics MUST remain delegated.

### Requirement: Conflict and Ambiguity Resolution

When SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless an approved spec change explicitly redesigns DAG, artifact, routing, status, or persistence semantics. Explicit redesigns MUST define compatibility rules for old artifacts and MUST NOT silently invalidate archives. Historical exhaustive secure-design matrices MAY remain readable as archives, but new changes MUST treat narrative `design.md#secure-development-design` plus canonical `review-security-report.json` as the active boundary. Status tokens, resolver rows, and persisted refs MAY preserve historical `security-applicability` data for read/display behavior, but MUST NOT normalize it into a runnable phase, launchable agent, active security authority, or required new-change successor.

#### Scenario: Explicit DAG redesign is applied

- GIVEN an approved spec changes phase order
- WHEN contracts are updated
- THEN the new DAG MUST be authoritative for new changes
- AND older artifacts MUST remain readable only under documented compatibility rules.

#### Scenario: Compatibility is preserved

- GIVEN existing persisted SDD state or artifacts use established keys, paths, routing tokens, exhaustive design rows, or status fields
- WHEN the contracts are updated
- THEN existing consumers MUST continue to resolve those artifacts and states without migration.

#### Scenario: Historical token is not launchable

- GIVEN persisted state or status contains `security-applicability`
- WHEN routing or agent resolution is computed
- THEN the token MAY be interpreted as historical data state
- AND it MUST NOT map to `sdd-security-applicability` or any runnable successor.

### Requirement: Review Phase Artifact Contract

The SDD contract set MUST define canonical `review-report.json` plus derived `review-report.md` as the first review artifact identity after apply and before `sdd-review-security`. OpenSpec mode MUST store them at `openspec/changes/{change-name}/review-report.json` and `review-report.md`; Engram or hybrid modes MUST use canonical key `sdd/{change-name}/review-report.json` plus stable derived key `sdd/{change-name}/review`. State and status contracts MUST expose review refs, verdict, blocking-failure state, and routing to security review when non-blocking.

#### Scenario: Review artifact is resolved

- GIVEN a downstream phase needs review evidence
- WHEN it resolves artifacts for a change
- THEN it MUST find canonical `review-report.json` when present plus the matching derived backend artifact key
- AND missing review evidence MUST block verify or archive.

#### Scenario: Security review resolves review artifact

- GIVEN `sdd-review-security` needs review evidence
- WHEN it resolves artifacts for a change
- THEN it MUST find canonical `review-report.json` when present plus the matching derived backend key
- AND missing review evidence MUST block security review, verify, and archive.

### Requirement: Apply Review Verify Routing

The SDD DAG for new changes MUST route `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`. Apply success MUST recommend review, non-blocking review MUST recommend review-security, non-blocking security review MUST recommend verify, and blocking findings MUST route to apply or resolve-blockers.

#### Scenario: Mandatory review route is enforced

- GIVEN implementation has completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of apply
- AND review MUST be the required successor.

#### Scenario: Mandatory security review route is enforced

- GIVEN implementation and general review have completed
- WHEN phase routing is evaluated
- THEN verify MUST NOT be the direct successor of review
- AND `review-security` MUST be the required successor.

#### Scenario: Design routes directly to test design

- GIVEN `design.md` includes `## Secure Development Design`
- WHEN phase routing is evaluated
- THEN `test-design` MUST be the direct successor
- AND `security-design` MUST NOT be an active new-change successor.

#### Scenario: Review cannot safely run

- GIVEN required artifacts or changed-file context are missing
- WHEN review evaluates readiness
- THEN it MUST return `resolve-blockers`
- AND it MUST state the missing or unsafe input.

### Requirement: Verify and Archive Review Consumption

Verify MUST consume canonical `review-report.json` and canonical `review-security-report.json` when present. Security-review JSON MUST be authoritative for verdict, routing, blockers, warnings, artifact parity, and `sourceRowValidation.rows` exact-once coverage. Verify and archive MUST NOT consume derived Markdown or compact `SEC-*` data as validation authority.

#### Scenario: Verify consumes review evidence

- GIVEN review and security-review reports are non-blocking
- WHEN verify runs
- THEN it MUST cite both canonical JSON reports
- AND it MUST NOT duplicate or reinterpret their matrices.

#### Scenario: Verify consumes both review artifacts

- GIVEN both review reports are non-blocking
- WHEN verify runs
- THEN it MUST cite both reports as evidence
- AND it MUST NOT duplicate their full matrices.

#### Scenario: Compact report data is ignored

- GIVEN a security-review artifact contains legacy compact `SEC-*` report data
- WHEN verify or archive evaluates a new change
- THEN that compact data MUST NOT satisfy active security validation.

#### Scenario: Archive checks review readiness

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST also require a non-blocking review report
- AND blocking review findings MUST prevent archive.

### Requirement: Mandatory Security Artifacts and Status

For new changes, persistence and status contracts MUST include `design.md` with narrative secure development rules and `review-security-report.json` refs, dependency states, `review-security` token, and archive gates. Design MUST persist classification rationale, changed-surface inventory, applicable category rules, evidence owners, residual risks, exceptions, and safe-evidence policy. It MUST NOT require YAML, schemas, compact matrices, Source ID matrices, exhaustive applicability, or `N/A` rows. Those machine-readable artifacts MUST persist in canonical `review-security-report.json`; Markdown is a derived compatibility view. `security-design.md` and `security-applicability.md` MAY appear only as historical refs.

#### Scenario: New state exposes security refs

- GIVEN a new change is persisted
- WHEN status or continuation reads state
- THEN artifact refs MUST include design and security review report slots
- AND active dependencies MUST NOT include `security-design` or `security-applicability`.

#### Scenario: Legacy refs are preserved as data

- GIVEN an archived change contains `artifactRefs.securityDesign`, `artifactRefs.securityApplicability`, or exhaustive design rows
- WHEN status or continuation displays historical evidence
- THEN the ref MAY remain visible as read-only data
- AND continuation MUST route active work through narrative design and review-security.

### Requirement: Active Security Validator Removal

New-change contracts MUST use catalog and artifact evidence for security validation. Active status, continuation, review-security, verify, and archive gating MUST NOT depend on retired validator scripts.

#### Scenario: Retired validators do not participate

- GIVEN retired validator scripts are absent from the active workflow
- WHEN a new change reaches review-security, verify, or archive
- THEN the workflow MUST use catalog and artifact evidence instead
- AND retired validator availability MUST NOT be a blocker.

### Requirement: Source Row Persistence Compatibility

The SDD contracts MUST preserve source-row-first security evidence across OpenSpec, Engram, hybrid, and none modes. Active security-review artifacts MUST persist `sourceRowValidation.rows` with exactly 155 unique rows and required row fields. Derived Markdown MUST remain a generated compatibility view with lean navigation/summary plus the full row matrix at the end. Backend behavior MUST NOT redefine row semantics.

#### Scenario: OpenSpec mode preserves rows

- GIVEN a change runs in OpenSpec mode
- WHEN source-row artifacts are persisted
- THEN canonical JSON MUST contain all 155 source rows
- AND Markdown MUST be regenerated from JSON.

#### Scenario: Engram or hybrid mode preserves rows

- GIVEN Engram or hybrid mode is selected
- WHEN source-row artifacts are persisted
- THEN Engram keys MUST use the shared artifact naming contract
- AND hybrid mode MUST reconcile backend disagreements before continuing.

#### Scenario: None mode is explicit

- GIVEN none mode is selected
- WHEN source-row evidence is produced inline
- THEN no project files or Engram observations MUST be written
- AND downstream recovery limits MUST be reported.

### Requirement: Verify Source Row Consumption

`sdd-verify` MUST validate that non-blocking `review-security-report.json` has complete `sourceRowValidation.rows` coverage and no row-level blockers. It MUST cite catalog snapshot/count, warnings, exceptions, and evidence refs without owning validation logic or compact-control summaries.

#### Scenario: Security source blocker remains

- GIVEN review-security reports a blocking source row
- WHEN verify runs
- THEN verification MUST block
- AND it MUST route to apply or resolve-blockers according to the blocker cause.

#### Scenario: Complete source rows continue

- GIVEN all 155 rows are present and non-blocking
- WHEN verify records evidence
- THEN it MAY proceed
- AND it MUST preserve warning and exception refs.

### Requirement: Archive Source Row Preservation

`sdd-archive` MUST require passing verification plus non-blocking source-row security review for new changes. Archive MUST preserve canonical JSON and generated Markdown refs, catalog snapshot identity/path, expected and validated counts, warnings, exceptions, and evidence references. Archive summaries MUST NOT use compact `SEC-*` validation, navigation, summaries, or `N/A` grouping.

#### Scenario: Archive checks row completeness

- GIVEN verification passes
- WHEN archive evaluates readiness
- THEN it MUST confirm 155 expected and validated rows with no blockers
- AND missing mandatory source-row evidence MUST prevent archive.

#### Scenario: Archive preserves generated matrix

- GIVEN derived Markdown contains the full source-row matrix at the end
- WHEN archive completes
- THEN the matrix ref MUST remain readable
- AND archive MUST not create a second validation source.

#### Scenario: Archive avoids matrix duplication

- GIVEN derived `review-security-report.md` contains the exhaustive source-row matrix generated from canonical JSON
- WHEN archive writes final records
- THEN it MUST link or summarize the matrix instead of duplicating it
- AND archived evidence MUST remain readable through embedded secure design and review-security evidence.

### Requirement: Operational Readiness Evidence Persistence

SDD persistence, status, verify, and archive contracts MUST preserve operational considerations, evidence, placeholders, gaps, and artifact references when they exist in design, test-design, tasks, apply, review, security review, verify, or archive evidence. Verify and archive MUST consume actual evidence rather than a shared operational-readiness contract. They MUST NOT require mandatory operational category completeness or disclosure of real operational data.

#### Scenario: Operational refs survive workflow

- GIVEN any SDD artifact records operational considerations or gaps
- WHEN status, verify, or archive resolves artifacts
- THEN those refs and unresolved gaps MUST remain readable.

#### Scenario: Verify checks applicable evidence

- GIVEN design or downstream artifacts make operational evidence applicable
- WHEN verify runs
- THEN each applicable item MUST have safe evidence, `Pendiente de confirmar:`, or `No aplica.`.

#### Scenario: No applicable evidence exists

- GIVEN design marks operational considerations not applicable or omits them safely
- WHEN verify or archive runs
- THEN missing readiness categories MUST NOT block completion.

### Requirement: Manual Operational Document Boundary

The DAG MUST NOT treat `sdd-operational-doc` as a required phase. The utility MUST remain manual, post-archive, and archive-consuming. It MUST generate from archived evidence, MUST NOT invent data, and MUST mark absent inapplicable values as `No aplica.` or unresolved applicable values as pending while preserving operational document sections 1-9 and diagrams R1-R4.

#### Scenario: Archive completes without operational doc

- GIVEN verify passes and archive criteria are met
- WHEN archive runs
- THEN completion MUST NOT require `sdd-operational-doc` execution.

#### Scenario: Manual utility consumes archive

- GIVEN an archived change contains operational evidence or gaps
- WHEN `sdd-operational-doc` is invoked manually
- THEN it MUST read archived evidence first
- AND absent values MUST remain pending or `No aplica.` without invention.

### Requirement: Manual Technical Document Boundary

The DAG MUST NOT treat `sdd-technical-doc` as a required phase. The utility MUST remain manual, post-archive, and archive-consuming, analogous to `sdd-operational-doc`. It MUST NOT change phase order, status routing, dependency graph, verify gates, archive gates, or required persistence artifacts. It MUST generate from archived evidence only, MUST NOT invent data, MUST mark inapplicable sections as `No aplica.`, and MUST mark unavailable applicable information explicitly.

#### Scenario: Archive completes without technical document

- GIVEN verify passes and archive criteria are met
- WHEN archive runs
- THEN completion MUST NOT require `sdd-technical-doc` execution
- AND missing technical documentation MUST NOT block archive.

#### Scenario: Status does not expose a required phase

- GIVEN status or continuation evaluates an active SDD change
- WHEN successor phases or required artifacts are computed
- THEN `sdd-technical-doc` MUST NOT appear as a required phase, DAG token, status dependency, verify input, or archive input.

#### Scenario: Manual utility consumes archived evidence

- GIVEN an archived change exists
- WHEN `sdd-technical-doc` is invoked manually
- THEN it MUST read archived evidence as its source of truth
- AND absent or inapplicable values MUST be represented without invention.

### Requirement: Final Documentation Restricted Data Boundary

Production hostnames, IPs, ports, SID/service names, and similar operational identifiers MAY be included only in final operational documentation when explicitly provided by the user. Ordinary SDD evidence and examples MUST preserve safe placeholders or references for applicable operational considerations.

#### Scenario: User provides final operational values

- GIVEN the user explicitly provides production operational identifiers for the manual document
- WHEN documentation is generated
- THEN the final operational document MAY include them
- AND SDD evidence artifacts MUST NOT be backfilled with those values.
