# Tasks: Per-Call Granularity for sdd-cost-tracker

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | ~150–220 (additions + deletions across 4 files) |
| Review budget lines | 400 |
| 400-line budget risk | Low |
| Review budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | single-pr |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: Low
400-line budget risk: Low

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | All Go + TypeScript changes with tests | PR 1 | Single PR to main; all files in scope |

---

## Phase 1: Storage Foundation (`storage.go`)

- [x] 1.1 Add `createCallsTableSQL` constant with `CREATE TABLE IF NOT EXISTS calls` DDL (columns: id, session_id, call_index, model_id, provider_id, tokens_*, cost_usd, recorded_at) in `tools-go/sdd-cost-tracker/storage.go`
- [x] 1.2 Add `CallRecord` struct with JSON tags matching the DDL columns in `storage.go`
- [x] 1.3 Execute `createCallsTableSQL` inside `init()` alongside the existing `phases` DDL in `storage.go`
- [x] 1.4 Implement `func (s *Store) InsertCall(ctx context.Context, r CallRecord) error` — rejects empty `session_id` with error in `storage.go`
- [x] 1.5 Implement `func (s *Store) GetCallsBySession(ctx context.Context, sessionID string) ([]CallRecord, error)` — returns `[]CallRecord{}` (not nil) for unknown session; orders by `call_index ASC` in `storage.go`

## Phase 2: HTTP Layer (`http.go`)

- [x] 2.1 Add `handlePostCalls` handler: decode JSON body into `CallRecord`, validate `session_id` present, call `store.InsertCall`, return 201 on success / 400 on missing field or malformed JSON / 500 on store error in `tools-go/sdd-cost-tracker/http.go`
- [x] 2.2 Add `handleGetCalls` handler: require `session_id` query param (400 if absent), call `store.GetCallsBySession`, return `{"calls": [...]}` with HTTP 200 in `http.go`
- [x] 2.3 Register `POST /calls` and `GET /calls` routes in `NewRouter` in `http.go`

## Phase 3: Plugin (`sdd-cost-tracker.ts`)

- [x] 3.1 Add `callIndices: Map<string, number>` at module scope, separate from `sessions` map, in `src/plugins/sdd-cost-tracker.ts`
- [x] 3.2 Implement `async function postCall(sessionId: string, callIndex: number, msg: any): Promise<void>` — builds `CallRecord` JSON from per-message `msg.tokens` and `msg.cost` values (not accumulated totals), fire-and-forget POST to `http://127.0.0.1:7438/calls` in `sdd-cost-tracker.ts`
- [x] 3.3 In the `message.updated` handler: increment `callIndices.get(sessionId)` (or init to 0), then call `postCall(sessionId, callIndex, msg)` as fire-and-forget (do not await in the critical path) in `sdd-cost-tracker.ts`

## Phase 4: Tests

- [x] 4.1 Write `TestInsertCallHappyPath` in `storage_test.go` — covers TC-001 (REQ-STORE-INSERT)
- [x] 4.2 Write `TestGetCallsBySessionHappyPath` in `storage_test.go` — covers TC-002 (REQ-STORE-QUERY)
- [x] 4.3 Write `TestGetCallsBySessionEmpty` in `storage_test.go` — covers TC-003 (REQ-STORE-QUERY-EMPTY)
- [x] 4.4 Write `TestGetCallsBySessionMultiple` in `storage_test.go` — covers TC-004 (REQ-STORE-MULTI)
- [x] 4.5 Write `TestGetCallsBySessionOrdering` in `storage_test.go` — covers TC-005 (REQ-STORE-ORDER); insert call_index 2, 0, 1 and assert order `[0, 1, 2]`
- [x] 4.6 Write `TestHTTPPostCallsInsertReturnsCreated` in `http_test.go` — covers TC-006 (REQ-HTTP-POST-201)
- [x] 4.7 Write `TestHTTPPostCallsMissingSessionIDReturns400` in `http_test.go` — covers TC-007 (REQ-HTTP-POST-400-MISSING); use `assertErrorPayload`
- [x] 4.8 Write `TestHTTPPostCallsMalformedJSONReturns400` in `http_test.go` — covers TC-008 (REQ-HTTP-POST-400-MALFORMED); send raw `"not-json"` body
- [x] 4.9 Write `TestHTTPGetCallsValidSession` in `http_test.go` — covers TC-009 (REQ-HTTP-GET-200)
- [x] 4.10 Write `TestHTTPGetCallsUnknownSession` in `http_test.go` — covers TC-010 (REQ-HTTP-GET-EMPTY)
- [x] 4.11 Write `TestHTTPGetCallsMissingSessionIDReturns400` in `http_test.go` — covers TC-011 (REQ-HTTP-GET-400)

## Phase 5: Verification

- [x] 5.1 Run `go test ./...` in `tools-go/sdd-cost-tracker` — all 19 existing phases tests plus new calls tests must pass; covers TC-012 (backward-compatibility)
- [x] 5.2 Run `go build -o sdd-cost-tracker.exe .` in `tools-go/sdd-cost-tracker` — must exit 0; covers TC-013 (build-integrity)
- [x] 5.3 Manual code review of `sdd-cost-tracker.ts`: confirm `callIndices` map increments per session message, is independent of `SessionContext`, and `postCall` reads per-message values (not accumulated) — advisory TC-014 (REQ-PLUGIN-INDEX); note: TypeScript test runner unavailable
- [x] 5.4 Static review: confirm `design.md#secure-development-design` no-impact rationale matches implementation surface (localhost SQLite, numeric values only, no auth/PII/secrets) — mandatory TC-015 (SEC-NO-IMPACT)
- [x] 5.5 Static check: confirm `design.md#Operational Considerations` contains `No aplica.` marker — operational considerations check
