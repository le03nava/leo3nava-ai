# Tasks: phase-status-validator

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 550–650 |
| 400-line budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 → PR 2 → PR 3 → PR 4 |
| Delivery strategy | force-chained |
| Chain strategy | feature-branch-chain |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: feature-branch-chain
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Core types + catalog + fs detection + content validators + tests | PR 1 | base: feature/phase-status-validator; foundation for all later work |
| 2 | Derivation engine + routing decision table + tests | PR 2 | base: PR 1 branch; depends on types from unit 1 |
| 3 | Wire projection + Resolve entry point + tests | PR 3 | base: PR 2 branch; depends on derivation from unit 2 |
| 4 | CLI wrapper + integration test + STATUS-CONTRACT.md | PR 4 | base: PR 3 branch; thin shell over Resolve |

## Phase 1: Foundation — Core Types + Catalog + Filesystem Detection

- [x] 1.1 Create `tools-go/leo3nava-ai/go.mod` — module `github.com/leo3nava/leo3nava_ai/tools-go/leo3nava-ai`, Go 1.24, stdlib only
- [x] 1.2 Create `tools-go/leo3nava-ai/internal/phasestatus/status.go` — define `ArtifactState` (missing/partial/done), `DependencyState` (blocked/ready/all_done), `NextRecommended` enum (proposal/specs/design/tasks/apply/verify/none), phase constants
- [x] 1.3 Create `tools-go/leo3nava-ai/internal/phasestatus/artifact.go` — define `ArtifactType` enum (proposal/specs/design/tasks/verify-report/apply-progress), 6-entry catalog table with detection rules
- [x] 1.4 RED: Create `tools-go/leo3nava-ai/internal/phasestatus/fs_test.go` — table-driven tests for on-disk detection: all artifacts present, missing artifact detected, nested specs layout, no specs files
- [x] 1.5 GREEN: Create `tools-go/leo3nava-ai/internal/phasestatus/fs.go` — implement `detectArtifacts(root string) map[ArtifactType]ArtifactState` with per-type file detection (single file for proposal/design/tasks/verify-report/apply-progress; glob `specs/**/spec.md` + flat `spec.md` fallback for specs)
- [x] 1.6 RED: Extend `fs_test.go` — table-driven tests for per-type content validation: proposal missing intent/scope → partial, specs missing requirements → partial, tasks with checkboxes → done, tasks without checkboxes → partial, design empty → partial, verify-report missing verdict → partial, apply-progress empty → partial
- [x] 1.7 GREEN: Extend `fs.go` — implement `validateContent(root string, m map[ArtifactType]ArtifactState) map[ArtifactType]ArtifactState` with per-type rules: proposal `## Intent`/`## Scope`, specs `## Requirements`, tasks `- [ ]`/`- [x]` checkbox, verify-report verdict section, design/apply-progress substantive content (non-empty beyond headings)

## Phase 2: Derivation Engine + Routing

- [x] 2.1 RED: Create `tools-go/leo3nava-ai/internal/phasestatus/derive_test.go` — table-driven tests for `deriveTaskProgress`: all checked → 100%, none checked → 0%, mixed → correct percentage
- [x] 2.2 GREEN: Create `tools-go/leo3nava-ai/internal/phasestatus/derive.go` — implement `deriveTaskProgress(m map[ArtifactType]ArtifactState, content map[string]string) TaskProgress` parsing checkbox lines from tasks content
- [x] 2.3 RED: Extend `derive_test.go` — table-driven tests for dependency state derivation: all phases complete → all_done, blocked by missing design → ready, early missing artifact → ready, anomaly (tasks done, no apply-progress) → blocked with blockedReasons
- [x] 2.4 GREEN: Extend `derive.go` — implement `deriveDependencies(m map[ArtifactType]ArtifactState, tp TaskProgress) (map[string]DependencyState, []string)` with ordered phase gates: proposal → specs → design → tasks → apply → verify
- [x] 2.5 RED: Extend `derive_test.go` — table-driven tests for routing: exhaustive decision table covering all 7 phases × artifact states, missing planning artifact routes directly, anomaly produces blocked + none
- [x] 2.6 GREEN: Extend `derive.go` — implement `routeNext(m map[ArtifactType]ArtifactState, deps map[string]DependencyState, blockedReasons []string) NextRecommended` with decision table

## Phase 3: Wire Projection + Entry Point

- [x] 3.1 RED: Create `tools-go/leo3nava-ai/internal/phasestatus/wire_test.go` — golden file test for frozen JSON output: byte-identical across runs, unknown token rejection returns error
- [x] 3.2 GREEN: Create `tools-go/leo3nava-ai/internal/phasestatus/wire.go` — implement `ProjectWire(status Status) ([]byte, error)` with JSON marshaling, schema identity fields, unknown-token rejection
- [x] 3.3 RED: Create `tools-go/leo3nava-ai/internal/phasestatus/resolve_test.go` — table-driven tests for `Resolve()`: all-artifacts-present → none, missing artifact → correct nextRecommended, anomaly → blocked, invalid path → error
- [x] 3.4 GREEN: Create `tools-go/leo3nava-ai/internal/phasestatus/resolve.go` — implement `Resolve(opts Options) (Status, error)` orchestrating detect → validate → derive → route → project

## Phase 4: CLI Wrapper + Documentation

- [x] 4.1 Create `tools-go/leo3nava-ai/internal/app/app.go` — implement `Run(args []string) error` parsing change root path, calling `phasestatus.Resolve()`, encoding JSON to stdout, error to stderr
- [x] 4.2 Create `tools-go/leo3nava-ai/cmd/phase-status-validator/main.go` — thin entry: parse args, call `app.Run()`, exit code on error
- [x] 4.3 RED: Create `tools-go/leo3nava-ai/internal/app/app_test.go` — integration test: valid change path → JSON stdout + exit 0, invalid path → stderr + exit 1
- [x] 4.4 GREEN: Extend `app.go` if needed to pass integration test
- [x] 4.5 Create `openspec/changes/phase-status-validator/STATUS-CONTRACT.md` — document frozen wire schema: 8 fields, types, validation rules
- [x] 4.6 Run `go test ./...` from `tools-go/leo3nava-ai/` — verify all tests pass
- [x] 4.7 Run `go vet ./...` and `gofmt -l .` — verify no issues
