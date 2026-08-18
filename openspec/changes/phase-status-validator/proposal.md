# Proposal: phase-status-validator

## Intent

SDD orchestrator lacks a reusable, read-only status validator that derives phase state from on-disk artifacts alone. Today the orchestrator must ad-hoc probe files and rebuild logic each time a new phase or artifact is added. A single-source catalog + derivation engine eliminates this drift, makes routing deterministic, and produces a frozen wire document for CLI and programmatic consumers.

## Scope

### In Scope

- Go package (`tools-go/leo3nava-ai/internal/phasestatus`, with `go.mod`; CLI under `cmd/phase-status-validator`) with: artifact catalog (6 artifacts), on-disk detection, per-type minimum-content validators, dependency state derivation, decision-table routing, frozen wire projection, `Resolve(options)` entry point.
- CLI wrapper: `cmd/phase-status-validator/main.go` → thin `internal/app.Run()` → `phasestatus.Resolve()`.
- Table-driven tests: `t.TempDir()` seeded fixtures, golden test for frozen wire JSON, routing-table coverage for every branch, per-artifact content validation cases.
- `STATUS-CONTRACT.md` documenting the frozen wire schema.

### Out of Scope

- Review artifacts (reviewPolicy, reviewLedger, etc.) — not applicable to this repository.
- Engram store integration — read-only filesystem derivation only.
- Runtime ledger or review gate logic (present in gentle-ai reference, excluded here).
- Spec, design, tasks, apply, verify, or archive artifacts for this change — proposal only.

## Capabilities

### New Capabilities

- `phase-status-validator`: Read-only status validator deriving SDD phase state from on-disk artifact files. Covers artifact catalog, per-type content validation, dependency derivation, routing, wire projection, and CLI entry point.

### Modified Capabilities

None.

## Approach

1. **Core types + catalog** (`status.go`, `artifact.go`): Define `ArtifactState`, `DependencyState`, `NextRecommended` enum, phase constants, and the 6-artifact catalog table modeled on `gentle-ai/internal/sddstatus/artifact_states.go` (stripped of review/Engram fields).

2. **Filesystem detection + content validators** (`fs.go`): Per-artifact on-disk detection (single file for proposal/design/tasks/verifyReport/applyProgress; nested `specs/**/spec.md` or flat `spec.md` for specs). Each type has a strict minimum-content rule:
   - proposal: must contain `## Intent` and/or `## Scope`
   - specs: must contain `## Requirements`
   - tasks: at least one checkbox line (`- [ ]` or `- [x]`)
   - verify-report: verdict section present
   - design: substantive content beyond empty
   - apply-progress: substantive content (pragmatic — optional until tasks 100% complete)

3. **Derivation + decision table** (`derive.go`): Build artifact state map → derive `TaskProgress` → compute per-dependency state (`blocked | ready | all_done`) → route `nextRecommended` via exhaustive decision table. Anomalies produce `blocked` with `blockedReasons`; no `resolve-blockers` token.

4. **Wire projection** (`wire.go`): Frozen JSON document with schema identity, artifact states, task progress, dependencies, `nextRecommended`, and `blockedReasons`. Reject unknown `NextRecommended` tokens at projection time.

5. **Entry point** (`resolve.go`): `Resolve(Options)` — accepts change root, returns frozen status struct. Options: `ChangeRoot string`, `ChangeName *string`.

6. **CLI** (`cmd/phase-status-validator/main.go` + `internal/app/app.go`): Thin wrapper following `gentle-ai/cmd/gentle-ai/main.go` pattern — parse args, call `app.Run()`, print JSON to stdout.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `tools-go/leo3nava-ai/internal/phasestatus/` | New | Core validator package (6 files) |
| `tools-go/leo3nava-ai/cmd/phase-status-validator/` | New | CLI entry point |
| `tools-go/leo3nava-ai/internal/app/` | New | Thin app.Run() wrapper |
| `tools-go/leo3nava-ai/go.mod` | New | Module initialization (stdlib only) |
| `openspec/changes/phase-status-validator/` | New | Change artifacts |

## Decisions

1. **Minimal wire shape**: Only schemaName, schemaVersion, changeName, artifacts, taskProgress, dependencies, nextRecommended, blockedReasons. No planningHome, changeRoot, actionContext, relationships, remediation, review fields — unjustified by this domain.

2. **STRICT per-type minimum-content validation**: Each artifact type enforces its own content rule (see Approach §2). Files failing their rule report `partial`, not `done`.

3. **apply-progress.md pragmatic**: Optional until tasks.md is 100% complete. When tasks are complete AND no apply-progress exists → anomaly (dependency state `blocked`, `blockedReasons` populated).

4. **Routing on anomaly**: No `resolve-blockers` token. Genuine anomalies result in dependency state `blocked` with `blockedReasons` list. Missing planning artifacts route to the next missing phase directly.

5. **CLI included**: `cmd/phase-status-validator/main.go` thin wrapper + `internal/app.Run()` + core logic in `internal/phasestatus` package. Follows `gentle-ai/cmd/` + `internal/app/` pattern.

6. **Proposal question round completed**: All 5 shaping questions answered by user — decisions above are locked and treated as final. No questions remain.

## Non-Goals

- No review artifact support (reviewPolicy, reviewLedger, etc.)
- No Engram memory integration
- No runtime ledger or review gate logic
- No action-context, relationships, remediation, or phase-instructions wire fields
- No spec/design/tasks/apply/verify/archive artifacts for this change

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Content validators too strict for edge-case artifacts | Medium | Table-driven tests covering empty, minimal, and malformed content per type |
| Wire schema drift if future phases add fields | Low | Frozen wire projection with unknown-token rejection at compile time |
| Specs path ambiguity (nested vs flat) | Low | Both paths explicitly handled; test fixtures cover both layouts |

## Rollback Plan

Delete the new files under `tools-go/leo3nava-ai/internal/phasestatus/`, `tools-go/leo3nava-ai/cmd/phase-status-validator/`, and `tools-go/leo3nava-ai/internal/app/`. No existing code is modified. Remove the `go.mod`/`go.sum` if created.

## Dependencies

- Go 1.24+ (stdlib only — no external dependencies)
- On-disk SDD artifact structure under `openspec/changes/{changeName}/`

## Success Criteria

- [ ] `go test ./...` passes with all table-driven cases
- [ ] Frozen wire golden test produces byte-identical JSON across runs
- [ ] Routing table covers every branch (all 7 phases × artifact states)
- [ ] Per-artifact content validation rejects partial/empty files with correct `partial` state
- [ ] CLI outputs valid JSON to stdout for a real change directory
- [ ] `STATUS-CONTRACT.md` documents wire schema completely
