# Design: phase-status-validator

## Technical Approach

Read-only Go CLI deriving SDD phase state from on-disk artifacts under `openspec/changes/{changeName}/`. Core logic in `internal/phasestatus` library; thin CLI wrapper calls `Resolve()` and prints frozen JSON to stdout. Stdlib only. Follows `tools-go/sdd-cost-tracker` patterns (table-driven tests, `t.TempDir()`, `go test ./...`).

## Architecture Decisions

### Decision: Package layout — internal/ library + cmd/ wrapper

| Option | Tradeoff | Decision |
|--------|----------|----------|
| Flat `main` package (like sdd-cost-tracker) | Simpler but no reusable library | Rejected — proposal requires `internal/phasestatus` as importable package |
| `internal/phasestatus` + `cmd/` + `internal/app/` | Standard Go project layout; library testable independently | **Chosen** — matches proposal §Approach and enables unit testing without CLI |

**Rationale**: The validator must be callable programmatically by the orchestrator, not only via CLI. Separating the library from the CLI wrapper enables direct `phasestatus.Resolve()` calls.

### Decision: Module location

| Option | Tradeoff | Decision |
|--------|----------|----------|
| `tools-py/leo3nava-ai/` | Unconventional for Go; deviates from existing Go tools | Rejected — inconsistent with `tools-go/` convention |
| `tools-go/leo3nava-ai/` | Consistent with existing Go tool module path (`tools-go/sdd-cost-tracker`) | **Chosen** — maintainer decision, keeps Go tools co-located |

**Rationale**: Maintainer override of proposal §Affected Areas. Existing Go tool lives at `tools-go/sdd-cost-tracker`; keeping Go modules under `tools-go/` avoids path-convention drift.

### Decision: Specs detection — glob vs recursive walk

| Option | Tradeoff | Decision |
|--------|----------|----------|
| `filepath.Walk` | Handles any depth; more code | Rejected — specs max 2 levels deep |
| `filepath.Glob("specs/**/spec.md")` + flat fallback | Simple, matches known layouts | **Chosen** — covers nested and flat per spec §On-Disk Detection |

### Decision: Content validation approach

| Option | Tradeoff | Decision |
|--------|----------|----------|
| Markdown AST parser (goldmark) | Robust; adds dependency | Rejected — stdlib-only constraint |
| `strings.Contains` / `regexp` | Simple; covers all 6 rules | **Chosen** — sufficient for headings, checkboxes, verdict sections |

### Decision: Wire rejection for unknown tokens

| Option | Tradeoff | Decision |
|--------|----------|----------|
| Warn + emit `null` | Lenient; risks silent schema drift | Rejected |
| Return error, no JSON | Strict; fails fast | **Chosen** — matches spec §Frozen Wire Projection |

## Data Flow

```
CLI args (change root path)
    │
    ▼
app.Run() ──→ phasestatus.Resolve(Options)
                    │
                    ├─→ detectArtifacts(root)     → map[ArtifactType]ArtifactState
                    ├─→ validateContent(root, m)  → map[ArtifactType]ArtifactState (may downgrade to partial)
                    ├─→ deriveTaskProgress(m)     → TaskProgress
                    ├─→ deriveDependencies(m, tp) → map[dep]DependencyState + blockedReasons
                    ├─→ routeNext(m, deps)        → NextRecommended
                    └─→ ProjectWire(status)       → frozen JSON bytes
    │
    ▼
stdout (JSON) or stderr (error, exit 1)
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `tools-go/leo3nava-ai/go.mod` | Create | Module `github.com/leo3nava/leo3nava_ai/tools-go/leo3nava-ai`, Go 1.24, stdlib only |
| `tools-go/leo3nava-ai/internal/phasestatus/status.go` | Create | Core types: `ArtifactState`, `DependencyState`, `NextRecommended` enum, phase constants |
| `tools-go/leo3nava-ai/internal/phasestatus/artifact.go` | Create | 6-artifact catalog table, `ArtifactType` enum |
| `tools-go/leo3nava-ai/internal/phasestatus/fs.go` | Create | On-disk detection per artifact type, content validators |
| `tools-go/leo3nava-ai/internal/phasestatus/derive.go` | Create | Dependency state derivation, task progress, routing decision table |
| `tools-go/leo3nava-ai/internal/phasestatus/wire.go` | Create | Frozen JSON wire projection, unknown-token rejection |
| `tools-go/leo3nava-ai/internal/phasestatus/resolve.go` | Create | `Resolve(Options)` entry point |
| `tools-go/leo3nava-ai/internal/phasestatus/*_test.go` | Create | Table-driven tests: fs detection, content validation, derivation, routing, wire golden |
| `tools-go/leo3nava-ai/cmd/phase-status-validator/main.go` | Create | CLI entry: parse args, call `app.Run()` |
| `tools-go/leo3nava-ai/internal/app/app.go` | Create | Thin `Run()` wrapper: validate path, call `Resolve()`, encode JSON to stdout |

## Interfaces / Contracts

```go
type Options struct {
    ChangeRoot string
    ChangeName *string // inferred from dir name if nil
}

func Resolve(opts Options) (Status, error)

type Status struct {
    SchemaName      string                        `json:"schemaName"`
    SchemaVersion   int                           `json:"schemaVersion"`
    ChangeName      string                        `json:"changeName"`
    Artifacts       map[ArtifactType]ArtifactState `json:"artifacts"`
    TaskProgress    TaskProgress                  `json:"taskProgress"`
    Dependencies    map[string]DependencyState     `json:"dependencies"`
    NextRecommended NextRecommended               `json:"nextRecommended"`
    BlockedReasons  []string                      `json:"blockedReasons"`
}
```

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit — fs | On-disk detection per type (present/missing/nested specs) | `t.TempDir()` seeded fixtures, table-driven |
| Unit — content | Per-type content validation (pass/partial per rule) | Table-driven: empty, minimal, malformed, valid content |
| Unit — derive | Dependency state derivation, task progress | Table-driven: all combinations of artifact states |
| Unit — routing | Exhaustive decision table for every branch | Table-driven: all 7 phases × artifact states |
| Unit — wire | Frozen JSON output, unknown token rejection | Golden file test with `-update` flag |
| Unit — resolve | End-to-end `Resolve()` with seeded temp dirs | Table-driven: all-artifacts-present, missing-artifact, anomaly scenarios |

## Migration / Rollout

No migration required. New standalone tool, no existing code modified. Delete new files to roll back.

## Open Questions

None. Maintainer decided `tools-go/leo3nava-ai/` (consistent with `tools-go/sdd-cost-tracker`).
