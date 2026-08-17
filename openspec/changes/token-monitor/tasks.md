# Tasks: Token Monitor

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 380–480 |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 (foundation + storage) → PR 2 (addon + export + tests) |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: Medium
400-line budget risk: Medium

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Storage layer + project scaffold | PR 1 | Base: main; includes `storage.py`, `requirements*.txt`, schema |
| 2 | Addon + export + full test suite | PR 2 | Base: PR 1 branch; includes `token_monitor.py`, `export.py`, `tests/`, `README.md` |

---

## Phase 1: Foundation

- [x] 1.1 Create `token-monitor/requirements.txt` — pin `mitmproxy>=10.0`
- [x] 1.2 Create `token-monitor/requirements-dev.txt` — pin `pytest>=8.0`, `pytest-cov`
- [x] 1.3 Create `token-monitor/storage.py` — `TokenStorage` class: `__init__(db_path: Path)` opens SQLite, sets `PRAGMA journal_mode=WAL`, creates `token_events` table with the 15-column authoritative schema; `insert_event(event: dict)` uses `?` parameterized INSERT; `close()` closes connection; all path ops via `pathlib.Path` (TC-025, TC-027)
- [x] 1.4 Create `token-monitor/tests/__init__.py` (empty marker file)

---

## Phase 2: Core Implementation

- [x] 2.1 Create `token-monitor/token_monitor.py` — standalone functions: `extract_usage(body: bytes) -> dict | None` handles JSON extraction + `JSONDecodeError` guard + missing-usage guard (TC-001, TC-002, TC-003); `extract_sse_usage(chunks: list[bytes]) -> dict | None` scans for final chunk with `"usage"`, logs warning if absent (TC-004, TC-005)
- [x] 2.2 Add to `token_monitor.py` — `detect_agent(headers: dict) -> str`: `X-Agent-Name` → UA substring match (`opencode`, `claude-code`, `copilot`) → `"unknown"` (TC-007–TC-011)
- [x] 2.3 Add to `token_monitor.py` — `extract_headers(headers: dict) -> dict`: extracts `session_id`, `request_id`, `project`, `phase` from `X-Session-ID`, `X-Request-Id`, `X-Project`, `X-Phase`; absent → `None` (TC-012–TC-016)
- [x] 2.4 Add to `token_monitor.py` — `TokenMonitor` mitmproxy addon class: `__init__` initialises `TokenStorage` and `sse_buffers: dict`; `responseheaders` hook detects SSE (`Content-Type: text/event-stream`) and enables response buffering; `response` hook dispatches to `extract_usage` or `extract_sse_usage`, builds event dict, calls `storage.insert_event`, releases `sse_buffers[flow.id]` (TC-006); no request/response body stored in DB (TC-026)
- [x] 2.5 Create `token-monitor/export.py` — argparse CLI: `--db-path` (default `~/.token-monitor/usage.db`), `--format {jsonl,csv}` (default `jsonl`), `--output` (default stdout); reads `token_events` via `SELECT *`; JSONL writes one JSON line per row; CSV uses `csv.DictWriter` with `None → ""` mapping; missing DB exits code 1 with stderr message; all paths via `pathlib.Path` (TC-020–TC-022, TC-028)

---

## Phase 3: Testing

- [ ] 3.1 Create `token-monitor/tests/conftest.py` — define `FIXTURE_NON_STREAMING`, `FIXTURE_SSE_CHUNKS`, `FIXTURE_SSE_NO_USAGE` constants and `seeded_db(tmp_path)` pytest fixture per test-design spec
- [x] 3.2 Add unit tests for `extract_usage` and `extract_sse_usage` — covers TC-001, TC-002, TC-003, TC-004, TC-005
- [x] 3.3 Add unit test for SSE buffer release — TC-006: assert `flow_id not in sse_buffers` after processing
- [x] 3.4 Add parametrized unit tests for `detect_agent` — covers TC-007, TC-008, TC-009, TC-010, TC-011
- [x] 3.5 Add parametrized unit tests for `extract_headers` — covers TC-012, TC-013, TC-014, TC-015, TC-016
- [x] 3.6 Add `TokenStorage` unit tests using `tmp_path` — WAL mode assertion (TC-018), full-row insert assertion all 15 columns (TC-017), error-resilience test drop table then insert (TC-019)
- [x] 3.7 Add `export.py` unit tests — JSONL 3-row output (TC-020), CSV NULL→empty-string (TC-021), missing DB exit code 1 (TC-022); use `seeded_db` fixture and `subprocess.run` or `CliRunner`

---

## Phase 4: Static Evidence + Documentation

- [x] 4.1 Static review sign-off — confirm `storage.py` uses only `?` parameterized SQL; no f-string SQL (TC-025)
- [x] 4.2 Static review sign-off — confirm `token_monitor.py`, `storage.py`, `export.py` never persist request body, prompt, or API key in DB columns or stderr (TC-026)
- [x] 4.3 Static review sign-off — confirm `storage.py` uses `pathlib.Path` for all path ops (TC-027)
- [x] 4.4 Static review sign-off — confirm `export.py` uses `pathlib.Path` for all path ops (TC-028)
- [x] 4.5 Create `token-monitor/README.md` — mitmproxy CA trust setup (`certutil` / mkcert), `mitmdump --scripts` invocation, `X-Agent-Name` header config, export CLI usage, Windows notes
- [x] 4.6 Document manual integration tests TC-023 and TC-024 as run instructions in `README.md` (non-blocking; required before first production-like use)
