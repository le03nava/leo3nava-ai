# Test Design: Per-Call Granularity for sdd-cost-tracker

## Overview

This change adds a parallel `calls` storage path to `sdd-cost-tracker` — a new `calls` table in SQLite, `InsertCall`/`GetCallsBySession` store methods, `POST /calls` and `GET /calls` HTTP handlers, and a TypeScript plugin fire-and-forget dispatch. The testable surface is the Go layer only; the TypeScript plugin has no test infrastructure in this project and is validated manually. The security posture is classified **no-impact**: the `calls` path is structurally identical to the existing `phases` path, storing only numeric token counts and USD cost values to a localhost SQLite file.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/sdd-cost-tracker-calls/proposal.md` | Scope: add call-level cost tracking; defers MCP exposure |
| Spec | `openspec/changes/sdd-cost-tracker-calls/specs/call-level-cost-tracking/spec.md` | Requirements and scenarios for storage, HTTP, and plugin behavior |
| Spec | `openspec/changes/sdd-cost-tracker-calls/specs/sdd-cost-tracking/spec.md` | Backward compatibility and build integrity requirements |
| Design | `openspec/changes/sdd-cost-tracker-calls/design.md` | Architecture decisions, data flow, file changes, contracts, testing strategy |
| Secure Development Design | `openspec/changes/sdd-cost-tracker-calls/design.md#secure-development-design` | No-impact classification; no applicable security categories; changed surface is localhost numeric telemetry |
| Testing Capabilities | Phase context + `openspec/config.yaml` | Go: `go test ./...` in `tools-go/sdd-cost-tracker` with in-memory SQLite and `net/http/httptest` — available. TypeScript test runner: unavailable. Coverage command: unavailable. Lint/typecheck/format: unavailable. |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-001 | Spec: REQ-STORE-INSERT | `InsertCall` with valid `CallRecord` → row count in `calls` table is 1, no error returned | automated | mandatory | `go test -run TestInsertCallHappyPath ./tools-go/sdd-cost-tracker/` passes | Uses `newTestStore(t)` with `:memory:` SQLite |
| TC-002 | Spec: REQ-STORE-QUERY | `GetCallsBySession` returns the inserted record with matching fields | automated | mandatory | `go test -run TestGetCallsBySessionHappyPath` passes | Verify all scalar fields: session_id, call_index, cost_usd, token fields |
| TC-003 | Spec: REQ-STORE-QUERY-EMPTY | `GetCallsBySession` with unknown session_id returns `[]CallRecord{}` (empty, not nil/error) | automated | mandatory | `go test -run TestGetCallsBySessionEmpty` passes | Mirrors existing `TestGetSummaryEmpty` pattern |
| TC-004 | Spec: REQ-STORE-MULTI | Insert three calls with same session_id → `GetCallsBySession` returns all three | automated | mandatory | `go test -run TestGetCallsBySessionMultiple` passes | Verifies no deduplication or collision on same session |
| TC-005 | Spec: REQ-STORE-ORDER | Insert call_index 2, 0, 1 for same session → `GetCallsBySession` returns them ordered `[0, 1, 2]` | automated | mandatory | `go test -run TestGetCallsBySessionOrdering` passes | Validates `ORDER BY call_index ASC` in SQL |
| TC-006 | Spec: REQ-HTTP-POST-201 | `POST /calls` with full valid JSON body → HTTP 201, `Content-Type: application/json` | automated | mandatory | `go test -run TestHTTPPostCallsInsertReturnsCreated` passes | Use `performJSONRequest` helper; body must include `session_id` |
| TC-007 | Spec: REQ-HTTP-POST-400-MISSING | `POST /calls` without `session_id` field → HTTP 400 with `{"error": "..."}` payload | automated | mandatory | `go test -run TestHTTPPostCallsMissingSessionIDReturns400` passes | Mirrors `TestHTTPPostPhasesMissingProjectReturns400`; use `assertErrorPayload` |
| TC-008 | Spec: REQ-HTTP-POST-400-MALFORMED | `POST /calls` with non-JSON body → HTTP 400 with `{"error": "..."}` payload | automated | mandatory | `go test -run TestHTTPPostCallsMalformedJSONReturns400` passes | Send raw string `"not-json"` as body |
| TC-009 | Spec: REQ-HTTP-GET-200 | `GET /calls?session_id=X` after seeding one call → HTTP 200 and `{"calls": [...]}` array length 1 | automated | mandatory | `go test -run TestHTTPGetCallsValidSession` passes | Decode response, verify `calls` key and length |
| TC-010 | Spec: REQ-HTTP-GET-EMPTY | `GET /calls?session_id=ghost` with no seeded data → HTTP 200 and `{"calls": []}` | automated | mandatory | `go test -run TestHTTPGetCallsUnknownSession` passes | Mirrors `TestHTTPGetChangePhasesNoMatchReturnsEmptyArray` |
| TC-011 | Spec: REQ-HTTP-GET-400 | `GET /calls` without `session_id` query param → HTTP 400 with `{"error": "..."}` payload | automated | mandatory | `go test -run TestHTTPGetCallsMissingSessionIDReturns400` passes | Mirrors `TestHTTPListChangesMissingProject` |
| TC-012 | Spec: backward-compatibility | All 19 existing phases tests pass unchanged after calls table DDL and handlers are added | automated | mandatory | `go test ./tools-go/sdd-cost-tracker/...` passes with zero failures | No phases table or handler logic must be modified |
| TC-013 | Spec: build-integrity | `go build -o sdd-cost-tracker.exe .` in `tools-go/sdd-cost-tracker` exits 0 | automated | mandatory | Command exits with code 0; no compile errors | Must pass on Windows; binary name uses `.exe` suffix |
| TC-014 | Spec: REQ-PLUGIN-INDEX | Code review of `sdd-cost-tracker.ts`: `callIndices` map increments per session message, separate from `sessions` map, resets or is absent for new session keys | manual | non-mandatory | Reviewer confirms increment logic is correct and independent of `SessionContext` | No TypeScript test runner available in this project |
| TC-015 | Secure Development Design | Static review: `design.md#secure-development-design` confirms no-impact classification; changed surface stores only numeric values to localhost SQLite; no security categories apply | static | mandatory | Reviewer confirms classification rationale matches implementation surface | Evidence: design narrative section present and consistent with code surface |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Operational Considerations | Design explicitly states `No aplica.` for this local dev tool | static | `design.md#Operational Considerations` contains `No aplica.` marker | N/A — no monitoring, deployment, or backup applies |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| SEC-NO-IMPACT | No-impact classification validation | Yes | TC-015: static review of `design.md#secure-development-design` confirms no-impact rationale and changed-surface description | covered | None |
| SEC-OMITTED-CATS | Authentication, sessions, PII/PAN, secrets, permissions, files, database-access, sensitive-logging | No | Reviewed and omitted by `review-security` per `design.md` exception policy; no applicable categories | not-applicable | Omitted categories are reviewable by `review-security` as documented in `design.md` |

## Evidence Expectations

- Mandatory cases require implementation, execution, static/manual evidence, or a justified skip.
- Non-mandatory cases (TC-014) are advisory and become verification warnings when uncovered.
- TC-015 (security static review) is mandatory; the no-impact classification must be confirmed against the actual implementation surface before verification closes.
- TypeScript test runner is **unavailable** — TC-014 is planned as manual review only. This is not passing evidence; it is a reported constraint.
- Coverage command is **unavailable** — no coverage threshold is enforced.
- Lint, type checker, and formatter are **unavailable** — not claimed as passing evidence.
- Build check (TC-013) uses `go build`; this is available in the `tools-go/sdd-cost-tracker` module.
- Exhaustive source-row validation and `N/A` decisions remain owned by canonical `review-security-report.json`.

## Open Questions

- None
