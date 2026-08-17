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

When existing SDD rules are duplicated, ambiguous, or conflicting, the system MUST preserve current behavior unless a clarification is required to make the rule well-defined. Clarifications MUST identify the authoritative owner and MUST NOT redesign the SDD DAG, backend formats, artifact keys, OpenSpec paths, routing tokens, status fields, or native state names.

#### Scenario: Duplicated rule is homologated

- GIVEN a phase skill and shared contract describe the same persistence behavior differently
- WHEN the rule is homologated
- THEN the shared persistence contract MUST become authoritative
- AND the phase skill MUST retain only phase-local artifact obligations.

#### Scenario: Compatibility is preserved

- GIVEN existing persisted SDD state or artifacts use established keys, paths, routing tokens, or status fields
- WHEN the contracts are updated
- THEN existing consumers MUST continue to resolve those artifacts and states without migration.
