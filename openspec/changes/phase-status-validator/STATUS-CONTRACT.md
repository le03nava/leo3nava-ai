# STATUS-CONTRACT.md — Frozen Wire Schema

> **Source of truth**: The golden file in `wire_test.go` (`goldenJSON` constant).
> This document is a human-readable reference; the test defines the byte-exact contract.

## Schema Identity

| Field | Type | Constant | Description |
|-------|------|----------|-------------|
| `schemaName` | `string` | `"phase-status-validator"` | Identifies this wire format |
| `schemaVersion` | `int` | `1` | Semver-compatible version counter |

## Document Fields

The frozen JSON wire document contains exactly **8 fields**:

```json
{
  "schemaName": "phase-status-validator",
  "schemaVersion": 1,
  "changeName": "<string>",
  "artifacts": { "<artifact-type>": "<artifact-state>", ... },
  "taskProgress": { "total": <int>, "checked": <int>, "percent": <float> },
  "dependencies": { "<phase-name>": "<dependency-state>", ... },
  "nextRecommended": "<next-token>",
  "blockedReasons": [ "<string>", ... ] | null
}
```

### Field Details

#### `changeName` (`string`)
- Inferred from the directory basename of the change root.
- Can be explicitly overridden via `Options.ChangeName`.

#### `artifacts` (`map[string]string`)
- Keys are artifact type labels (exactly 6):
  - `proposal`
  - `specs`
  - `design`
  - `tasks`
  - `verify-report`
  - `apply-progress`
- Values are artifact state labels:
  - `"missing"` — artifact file does not exist on disk
  - `"partial"` — file exists but fails its content validation rule
  - `"done"` — file exists and passes its content validation rule
- Map keys are sorted alphabetically for deterministic output.

#### `taskProgress` (`object`)
| Sub-field | Type | Description |
|-----------|------|-------------|
| `total` | `int` | Total checkbox lines found (`- [ ]` or `- [x]`) |
| `checked` | `int` | Lines with `- [x]` (completed) |
| `percent` | `float64` | `checked / total * 100`; `0.0` when total is `0` |

#### `dependencies` (`map[string]string`)
- Keys are phase gate names (exactly 6, in dependency order):
  1. `proposal`
  2. `specs`
  3. `design`
  4. `tasks`
  5. `apply`
  6. `verify`
- Values are dependency state labels:
  - `"blocked"` — prior gates not satisfied
  - `"ready"` — all prior gates satisfied, this phase is the next candidate
  - `"all_done"` — this gate and all prior gates are complete
- Map keys are sorted alphabetically for deterministic output.

#### `nextRecommended` (`string`)
- One of exactly 7 tokens:
  - `"proposal"` | `"specs"` | `"design"` | `"tasks"` | `"apply"` | `"verify"` | `"none"`
- Unknown tokens are **rejected at projection time** (error returned, no JSON emitted).

#### `blockedReasons` (`[]string | null`)
- Populated when dependency state is `"blocked"` due to anomaly detection.
- Empty array `[]` or `null` when no anomalies exist.
- Anomaly example: tasks artifact is `done` (100% checked) but apply-progress is missing or partial.

## Validation Rules

| Rule | Behavior |
|------|----------|
| Unknown `nextRecommended` token | `ProjectWire()` returns error; no JSON emitted |
| Missing fields | Not possible — `ProjectWire()` always produces all 8 fields |
| Non-deterministic output | Prevented by sorted map keys for `artifacts` and `dependencies` |
| Schema version increment | Required when field set, types, or semantics change |

## Wire Projection Source

The canonical projection logic lives in:
- `internal/phasestatus/wire.go` — `ProjectWire(Status) ([]byte, error)`
- `internal/phasestatus/wire_test.go` — `TestProjectWireGolden` (byte-exact golden test)
