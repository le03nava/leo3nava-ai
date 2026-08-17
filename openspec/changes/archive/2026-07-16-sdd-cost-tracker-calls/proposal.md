# Proposal: Per-Call Granularity for sdd-cost-tracker

## Intent

The current tracker stores one aggregated row per SDD phase session. There is no record of individual LLM invocations within a session, making it impossible to audit which calls drove cost spikes, compare call-level token patterns, or correlate model behavior with phase step counts. This change adds a `calls` table and supporting endpoints so each assistant message maps to exactly one persisted call record.

## Scope

### In Scope
- New `calls` SQLite table (`CREATE TABLE IF NOT EXISTS`, no migration risk)
- `InsertCall` + `GetCallsBySession` storage methods in Go
- `POST /calls` (insert one call record) and `GET /calls?session_id=X` endpoints
- TypeScript plugin: per-message POST on `message.updated` with a `callIndex` counter; fire-and-forget
- Existing tests must stay green; new tests cover the new table, endpoints, and plugin behavior

### Out of Scope
- No changes to the `phases` table or aggregation logic
- No FK enforcement at DB level
- No MCP tool additions (unless they naturally expose call data — deferred)
- No changes to `postPhase` behavior in the plugin

## Capabilities

### New Capabilities
- `call-level-cost-tracking`: Records individual LLM invocations with token/cost metrics, queryable by session.

### Modified Capabilities
- `sdd-cost-tracking`: Plugin gains a `callIndex` counter and a `postCall` fire-and-forget path; Go binary gains two new HTTP endpoints and a new storage schema.

## Approach

Add a `calls` table alongside `phases` using the same `CREATE TABLE IF NOT EXISTS` pattern. Mirror the HTTP handler style from `http.go` for two new routes. In the TypeScript plugin, intercept `message.updated` events to POST each assistant message immediately (not accumulated), using a per-session `callIndex` counter. All plugin POSTs are fire-and-forget to prevent tracking from crashing the agent session.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `tools-go/sdd-cost-tracker/storage.go` | Modified | Add `calls` table DDL, `InsertCall`, `GetCallsBySession` |
| `tools-go/sdd-cost-tracker/http.go` | Modified | Add `POST /calls` and `GET /calls` handlers |
| `tools-go/sdd-cost-tracker/storage_test.go` | Modified | Tests for new table and methods |
| `tools-go/sdd-cost-tracker/http_test.go` | Modified | Tests for new endpoints |
| `src/plugins/sdd-cost-tracker.ts` | Modified | Add `callIndex`, `postCall` on `message.updated` |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Plugin logic regression (accumulate → per-call) | Med | Keep `postPhase` path untouched; add `postCall` as separate branch |
| Increased POST volume breaks local dev tool | Low | Local-only; fire-and-forget with no retry |
| `calls` table DDL conflicts with future schema changes | Low | `IF NOT EXISTS` guard + no FK enforcement |

## Rollback Plan

Remove `calls` table DDL and the two HTTP handlers from Go; remove `postCall` path and `callIndex` counter from the plugin. No data migrations needed — `phases` table is untouched.

## Dependencies

- None. All changes are internal to the `sdd-cost-tracker` binary and its plugin.

## Success Criteria

- [ ] `go build -o sdd-cost-tracker.exe .` succeeds with no errors
- [ ] All existing tests pass (`go test ./...`)
- [ ] `POST /calls` inserts a row and returns 201; `GET /calls?session_id=X` returns the rows
- [ ] Plugin POSTs one call record per assistant message without crashing the agent session
- [ ] `calls` table is created automatically on fresh DB start
