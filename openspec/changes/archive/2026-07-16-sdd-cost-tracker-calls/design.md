# Design: Per-Call Granularity for sdd-cost-tracker

## Technical Approach

Add a parallel `calls` storage path mirroring the existing `phases` architecture: a new `CallRecord` struct, `InsertCall`/`GetCallsBySession` on `Store`, two HTTP handlers registered in `NewRouter`, and a `postCall` fire-and-forget path in the TypeScript plugin with a per-session `callIndex` counter. No changes to MCP tools, phases table, or `postPhase` logic.

## Architecture Decisions

### Decision: CallRecord as flat struct (no FK)

**Choice**: Flat `CallRecord` struct with `session_id` as a plain TEXT field; no FK to `phases`.
**Alternatives considered**: FK enforcement, composite key `(session_id, call_index)`.
**Rationale**: Matches existing no-FK pattern in the codebase. Calls may arrive before the phase row exists (fire-and-forget timing). Composite unique key adds complexity with no query benefit since we always filter by `session_id`.

### Decision: No MCP tool additions

**Choice**: Defer MCP exposure of call data.
**Alternatives considered**: Add `cost_calls` MCP tool alongside `cost_query`.
**Rationale**: Proposal explicitly defers MCP. HTTP endpoints suffice for now; MCP can be added later without breaking changes.

### Decision: Plugin uses separate callIndex map, not SessionContext mutation

**Choice**: A dedicated `callIndices: Map<string, number>` alongside the existing `sessions` map.
**Alternatives considered**: Adding `callIndex` field to `SessionContext`.
**Rationale**: Keeps call tracking orthogonal to phase accumulation. `callIndex` lifecycle differs — it increments on every message, while `SessionContext` accumulates totals. Separate map is cleaner and avoids coupling.

### Decision: Per-message token extraction (NOT accumulated)

**Choice**: Read `msg.tokens` fields directly as the individual message's values, not compute deltas from accumulated totals.
**Alternatives considered**: Track previous totals and compute deltas.
**Rationale**: The spec explicitly states "individual message's token and cost metrics — NOT accumulated totals". The `message.updated` event carries per-message token data on each assistant turn, not running totals (verified from existing accumulation pattern using `+=`). We send the raw per-event values.

## Data Flow

```
Assistant Message Event
       │
       ▼
Plugin: message.updated handler
       │
       ├──→ Accumulate into SessionContext (existing postPhase path, unchanged)
       │
       └──→ postCall(sessionID, callIndex, msg.tokens, msg.cost)
                 │
                 ▼
            POST /calls → http.go handler → store.InsertCall → SQLite calls table
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `tools-go/sdd-cost-tracker/storage.go` | Modify | Add `createCallsTableSQL`, `CallRecord` struct, `InsertCall`, `GetCallsBySession`; add calls DDL exec in `init()` |
| `tools-go/sdd-cost-tracker/http.go` | Modify | Add `handlePostCalls`, `handleGetCalls`; register `/calls` in `NewRouter` |
| `tools-go/sdd-cost-tracker/mcp.go` | None | No changes |
| `src/plugins/sdd-cost-tracker.ts` | Modify | Add `callIndices` map, `postCall` helper, fire-and-forget dispatch in `message.updated` handler |

## Interfaces / Contracts

### Go: CallRecord struct

```go
type CallRecord struct {
    ID               int64   `json:"id"`
    SessionID        string  `json:"session_id"`
    CallIndex        int     `json:"call_index"`
    ModelID          *string `json:"model_id,omitempty"`
    ProviderID       *string `json:"provider_id,omitempty"`
    TokensInput      int64   `json:"tokens_input"`
    TokensOutput     int64   `json:"tokens_output"`
    TokensReasoning  int64   `json:"tokens_reasoning"`
    TokensCacheRead  int64   `json:"tokens_cache_read"`
    TokensCacheWrite int64   `json:"tokens_cache_write"`
    CostUSD          float64 `json:"cost_usd"`
    RecordedAt       *int64  `json:"recorded_at,omitempty"`
}
```

### Go: Store methods

```go
func (s *Store) InsertCall(ctx context.Context, r CallRecord) error
func (s *Store) GetCallsBySession(ctx context.Context, sessionID string) ([]CallRecord, error)
```

### HTTP Endpoints

- `POST /calls` — JSON body matching `CallRecord`. Returns 201/400/500.
- `GET /calls?session_id=X` — Returns `{"calls": [...]}` ordered by `call_index` ASC. 400 if missing param.

### Plugin: postCall signature

```typescript
async function postCall(sessionId: string, callIndex: number, msg: any): Promise<void>
```

## Operational Considerations

No aplica. This is a local development tool with no production deployment, no monitoring, no backup requirements, and no administration operations.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | `InsertCall` success + empty session_id rejection | `storage_test.go` with in-memory SQLite |
| Unit | `GetCallsBySession` ordering + empty result | `storage_test.go` |
| Integration | `POST /calls` 201/400, `GET /calls` 200/400 | `http_test.go` with `httptest.NewServer` |
| Unit | Plugin `callIndex` increment + reset on new session | Manual verification (no TS test infra in project) |

## Migration / Rollout

No migration required. `CREATE TABLE IF NOT EXISTS` handles fresh and existing databases identically.

## Open Questions

None.

## Secure Development Design

### Classification and Changed Surface

**Classification: no-impact.**

Changed artifacts: `storage.go`, `http.go`, `sdd-cost-tracker.ts`. All changes add a new data path (`calls`) that stores token counts and cost metrics — the same non-sensitive operational telemetry already stored in the `phases` table.

**Touched runtime surfaces**: Local SQLite writes via new `InsertCall`/`GetCallsBySession`; local HTTP endpoints on `127.0.0.1:7438`; plugin fire-and-forget fetch to localhost.

**Untouched runtime surfaces**: No authentication (tool is localhost-only, no auth exists or is needed). No sessions. No PII/PAN. No secrets handling. No file operations beyond SQLite. No sensitive logging. No external network calls. No access control.

**Why no security category applies**: The tool records numeric token counts and USD cost values for local developer use. It binds to localhost only, stores no credentials or user data, processes no external input beyond the plugin's own POST bodies, and has no production deployment. The existing `phases` path already establishes this pattern with identical security posture. The new `calls` path is structurally identical.

### Exception and Evidence Policy

No exceptions planned. No applicable security categories require evidence. Omitted categories (authentication, sessions, sensitive-data-pan, secrets, permissions-access-control, files, database-access, sensitive-logging) are reviewable omissions for `review-security` validation.
