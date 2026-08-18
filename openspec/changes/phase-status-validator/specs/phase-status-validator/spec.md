# Phase Status Validator Specification

## Purpose

Read-only SDD phase status validator that derives phase state from on-disk artifact files alone. Provides a deterministic catalog, per-type content validation, dependency derivation, routing, frozen wire projection, and a CLI entry point. No Engram, no review artifacts, no runtime ledger.

## Requirements

### Requirement: Artifact Catalog

The system MUST maintain a catalog of exactly 6 artifact types: `proposal`, `specs`, `design`, `tasks`, `verify-report`, `apply-progress`. Each type SHALL have a defined on-disk detection rule.

#### Scenario: All artifacts present

- GIVEN a change directory containing proposal.md, specs/\*\*/\*.md, design.md, tasks.md, verify-report.md, apply-progress.md
- WHEN the catalog is evaluated
- THEN all 6 artifacts report `present`

#### Scenario: Missing artifact detected

- GIVEN a change directory missing design.md
- WHEN the catalog is evaluated
- THEN design reports `missing` and all other present artifacts report `present`

### Requirement: On-Disk Detection

The system MUST detect each artifact on-disk without Engram or runtime state. Proposal, design, tasks, verify-report, and apply-progress SHALL be detected as single files. Specs SHALL be detected via nested `specs/**/spec.md` or flat `spec.md` under the change directory.

#### Scenario: Nested specs layout

- GIVEN a change with `specs/auth/spec.md` and `specs/routing/spec.md`
- WHEN specs detection runs
- THEN specs state is `present`

#### Scenario: No specs files exist

- GIVEN a change directory with no `specs/` directory or `spec.md`
- WHEN specs detection runs
- THEN specs state is `missing`

### Requirement: Per-Type Content Validation

Each artifact type MUST enforce a strict minimum-content rule. Files present but failing their rule SHALL report `partial`, not `done`.

| Artifact | Rule |
|----------|------|
| proposal | Must contain `## Intent` and/or `## Scope` |
| specs | Must contain `## Requirements` |
| tasks | Must have at least one checkbox line (`- [ ]` or `- [x]`) |
| verify-report | Must have a verdict section |
| design | Must have substantive content (non-empty beyond headings) |
| apply-progress | Must have substantive content |

#### Scenario: Proposal with no intent or scope

- GIVEN proposal.md exists but lacks `## Intent` and `## Scope`
- WHEN content validation runs
- THEN proposal state is `partial`

#### Scenario: Tasks with checkboxes

- GIVEN tasks.md exists with at least one `- [ ]` line
- WHEN content validation runs
- THEN tasks state is `done`

#### Scenario: Apply-progress absent while tasks incomplete

- GIVEN tasks.md has no `- [x]` completing all items and apply-progress.md does not exist
- WHEN content validation runs
- THEN apply-progress state is `missing`

### Requirement: Dependency State Derivation

The system MUST derive per-dependency state as one of `blocked`, `ready`, or `all_done`. Dependencies SHALL be modeled as ordered phase gates: proposal → specs → design → tasks → apply → verify. Each gate is `all_done` when its artifact is `done`, `ready` when all prior gates are `all_done` and the artifact is `missing`/`partial`, and `blocked` otherwise.

#### Scenario: All phases complete

- GIVEN all 6 artifacts are `done`
- WHEN dependencies are derived
- THEN every dependency is `all_done` and nextRecommended is `none`

#### Scenario: Tasks blocked by missing design

- GIVEN proposal and specs are `done`, design is `missing`, tasks is `missing`
- WHEN dependencies are derived
- THEN the design dependency is `ready` and nextRecommended routes to design

### Requirement: Routing Decision Table

The system MUST route `nextRecommended` via an exhaustive decision table covering all artifact state combinations. The token SHALL be one of: `proposal`, `specs`, `design`, `tasks`, `apply`, `verify`, or `none`.

#### Scenario: Anomaly produces blocked

- GIVEN tasks is `done` but apply-progress is `missing` (anomaly: tasks complete without apply-progress)
- WHEN routing runs
- THEN dependency state is `blocked`, blockedReasons is populated, and nextRecommended is `none`

#### Scenario: Missing planning artifact routes directly

- GIVEN proposal is `missing` and all subsequent artifacts are `missing`
- WHEN routing runs
- THEN nextRecommended is `proposal` and no blocked state is reported

### Requirement: Frozen Wire Projection

The system MUST produce a frozen JSON wire document with exactly these fields: `schemaName`, `schemaVersion`, `changeName`, `artifacts`, `taskProgress`, `dependencies`, `nextRecommended`, `blockedReasons`. Unknown `NextRecommended` tokens MUST be rejected at projection time.

#### Scenario: Golden wire output

- GIVEN a fully populated status for a change named `my-feature`
- WHEN wire projection runs
- THEN the JSON output contains exactly the 8 declared fields and is byte-identical across runs

#### Scenario: Unknown token rejected

- GIVEN a status struct with nextRecommended set to an unrecognized token
- WHEN wire projection runs
- THEN an error is returned and no JSON is emitted

### Requirement: CLI Entry Point

The system MUST provide a CLI at `cmd/phase-status-validator/main.go` with a thin `internal/app.Run()` wrapper that accepts a change root path and prints the frozen wire JSON to stdout.

#### Scenario: Valid change path

- GIVEN a change root path pointing to a valid change directory
- WHEN the CLI is invoked
- THEN valid JSON is printed to stdout with exit code 0

#### Scenario: Invalid or missing path

- GIVEN a non-existent directory path
- WHEN the CLI is invoked
- THEN an error message is printed to stderr with exit code 1
