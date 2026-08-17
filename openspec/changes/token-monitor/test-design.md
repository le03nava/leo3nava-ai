# Test Design: Token Monitor

## Overview

`token-monitor` is a new standalone local Python tool. The testable surface covers three pure-Python modules — `token_monitor.py` (extraction + agent detection logic), `storage.py` (SQLite WAL persistence), and `export.py` (JSONL/CSV CLI) — all of which can be exercised with `pytest` and a temporary SQLite database without a live mitmproxy runtime. Integration testing that requires a running `mitmdump` process is documented as manual scope only; no automated integration runner is available in this repository. The project-level `openspec/config.yaml` records no test runner, but the new tool ships its own `requirements.txt` with `pytest` as a dev dependency, so unit tests run inside `token-monitor/` independently of the repository root.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/token-monitor/proposal.md` | Scope, success criteria, out-of-scope items (no dashboard, no auth) |
| Spec — capture | `openspec/changes/token-monitor/specs/token-usage-capture/spec.md` | Proxy interception, SSE buffering, agent detection, SQLite persistence, CLI config scenarios |
| Spec — export | `openspec/changes/token-monitor/specs/token-usage-export/spec.md` | JSONL/CSV export correctness, missing DB exit code, portability |
| Design | `openspec/changes/token-monitor/design.md` | Architecture decisions, data flow, file contracts, WAL rationale, testing strategy table |
| Secure Development Design | `openspec/changes/token-monitor/design.md#secure-development-design` | Classification: no-impact. Local-only tool, metadata-only storage, parameterized SQL, no secrets logging. No applicable security categories. |
| Testing Capabilities | `openspec/config.yaml` (testing section) | No repo-level test runner, no coverage command, no linter, no type checker, no formatter available at repo root. New tool introduces its own pytest environment. |

## Source ID Coverage Baseline

**Classification**: no-impact (local developer tool, no production surface, no auth/session/PII/secrets/ACL surface touched).

No corporate source-row security categories are applicable. The changed surface is limited to local filesystem writes (SQLite), localhost network binding, and one-time OS trust-store install. All writes use parameterized SQL by design. No sensitive data is persisted beyond integer token counts and agent labels.

Security hygiene checks (TC-025, TC-026) are planned as static code review rather than automated tooling, because no linter or SAST runner is configured in this repository.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-001 | Spec capture: Non-streaming LLM response intercepted | Parse a fixture JSON body with `usage: {prompt_tokens, completion_tokens, total_tokens}` → extraction function returns correct int values | automated | mandatory | `pytest token-monitor/tests/test_token_monitor.py::test_extract_usage_nonstreaming` passes | See fixture `FIXTURE_NON_STREAMING` below |
| TC-002 | Spec capture: Proxy Interception (missing usage guard) | JSON body without a `usage` key → extraction returns `None` / skips persistence | automated | mandatory | `pytest` assertion: return value is `None` | Guard `if "usage" in data` required |
| TC-003 | Spec capture: Proxy Interception (malformed body) | Non-JSON body → extraction does not raise; returns `None` | automated | mandatory | `pytest` assertion: no exception propagated | `json.JSONDecodeError` must be caught |
| TC-004 | Spec capture: SSE usage in final chunk | Simulate 3 SSE chunks via buffer dict; only the last contains `"usage"` → `extract_sse_usage()` returns correct values | automated | mandatory | `pytest` assertion: token values match final chunk | See fixture `FIXTURE_SSE_CHUNKS` below |
| TC-005 | Spec capture: SSE no usage in final chunk | Final SSE chunk has no `"usage"` key → returns `None`, warning logged | automated | mandatory | `pytest` assertion: `None` returned; `caplog` captures warning | `logging.warning` or `mitmproxy.ctx.log.warn` mock required |
| TC-006 | Spec export (NFR): SSE buffer released after stream | After `extract_sse_usage()` processes a complete flow, `sse_buffers[flow_id]` is absent from dict | automated | mandatory | `pytest` assertion: `flow_id not in sse_buffers` | Verifies no per-flow memory leak |
| TC-007 | Spec capture: Agent Detection — X-Agent-Name priority | Call `detect_agent(headers={"X-Agent-Name": "copilot-chat", "User-Agent": "unknown-ua"})` → returns `"copilot-chat"` | automated | mandatory | `pytest` parametrized assertion | Must take precedence over User-Agent |
| TC-008 | Spec capture: Known agent via User-Agent | `detect_agent(headers={"User-Agent": "opencode/1.0"})` → returns `"opencode"` | automated | mandatory | `pytest` parametrized assertion | Substring match on known labels |
| TC-009 | Spec capture: Agent Detection — claude-code UA | `detect_agent(headers={"User-Agent": "claude-code/0.9"})` → returns `"claude-code"` | automated | mandatory | `pytest` parametrized assertion | |
| TC-010 | Spec capture: Unknown agent fallback | `detect_agent(headers={"User-Agent": "curl/7.x"})` → returns `"unknown"` | automated | mandatory | `pytest` assertion | |
| TC-011 | Spec capture: No headers at all | `detect_agent(headers={})` → returns `"unknown"`, no crash | automated | mandatory | `pytest` assertion: no exception | Nullable-header contract |
| TC-012 | Spec capture: X-Project header extracted | `extract_headers(headers={"X-Project": "my-project"})` → `project == "my-project"` | automated | mandatory | `pytest` assertion | |
| TC-013 | Spec capture: X-Phase header extracted | `extract_headers(headers={"X-Phase": "sdd-apply"})` → `phase == "sdd-apply"` | automated | mandatory | `pytest` assertion | |
| TC-014 | Spec capture: X-Session-ID header extracted | `extract_headers(headers={"X-Session-ID": "sess-abc"})` → `session_id == "sess-abc"` | automated | mandatory | `pytest` assertion | |
| TC-015 | Spec capture: X-Request-Id header extracted | `extract_headers(headers={"X-Request-Id": "req-123"})` → `request_id == "req-123"` | automated | mandatory | `pytest` assertion | |
| TC-016 | Spec capture: All nullable headers absent | `extract_headers(headers={})` → `project`, `phase`, `session_id`, `request_id` all `None` | automated | mandatory | `pytest` assertion: all None, no crash | Core nullable contract |
| TC-017 | Spec capture: Row inserted on successful capture | Call `TokenStorage.insert_event(event_dict)` on a temp DB → `SELECT * FROM token_events` returns 1 row with all 15 columns matching input | automated | mandatory | `pytest` with `tmp_path` fixture; assertion on all columns | Uses authoritative schema from change request |
| TC-018 | Spec capture: WAL mode enabled at init | After `TokenStorage(tmp_path)`, execute `PRAGMA journal_mode` → result is `"wal"` | automated | mandatory | `pytest` assertion on PRAGMA result | Critical for concurrent access |
| TC-019 | Spec capture: DB write error does not crash | Drop `token_events` table after init, then call `insert_event()` → no exception propagates; error is logged | automated | mandatory | `pytest` assertion: no exception; `caplog` captures error | Proxy resilience contract |
| TC-020 | Spec export: Full JSONL export | Pre-seed DB with 3 rows, run `export.py --format jsonl` via `subprocess` or CLI runner → stdout has exactly 3 valid JSON lines, each with all 15 fields | automated | mandatory | `pytest` output assertion | Fields must match authoritative schema |
| TC-021 | Spec export: CSV with NULL session_id | Pre-seed DB with rows where `session_id` is NULL → `--format csv` output has empty strings in that column, not "None" or "null" | automated | mandatory | `pytest` assertion on CSV rows | RFC-compliant NULL → empty string |
| TC-022 | Spec export: Missing DB exits with code 1 | Invoke `export.py --db-path /nonexistent/path.db` → process exits with code 1, stderr has error message | automated | mandatory | `pytest` `subprocess.run` check: `returncode == 1`, `stderr` non-empty | |
| TC-023 | Spec capture (integration, document-only): Non-streaming end-to-end | Running `mitmdump` + addon, mock LLM server returns non-streaming JSON → exactly 1 row in SQLite | manual | non-mandatory | Manual test log noting row count and field values | Requires live mitmproxy; no CI runner |
| TC-024 | Spec capture (integration, document-only): SSE end-to-end with buffering | Running `mitmdump` + addon, mock server streams N SSE chunks → buffer accumulates all → 1 row after stream end | manual | non-mandatory | Manual test log noting row count | Requires live mitmproxy; no CI runner |
| TC-025 | design.md#secure-development-design: parameterized SQL | Code review of `storage.py` confirms all `INSERT`/`SELECT` use `?` placeholders; no f-string or `%`-format SQL | static | mandatory | Code review sign-off; no `f"INSERT` or `f"SELECT` strings present | SQL injection hygiene |
| TC-026 | design.md#secure-development-design: no sensitive data logged | Code review of `token_monitor.py`, `storage.py`, `export.py` confirms no request body, prompt content, completion text, or API key is persisted or logged | static | mandatory | Code review sign-off | Request body MUST never reach SQLite columns |
| TC-027 | Spec export: pathlib.Path in storage.py | Code review confirms `storage.py` uses `pathlib.Path` for all path operations; no bare string concatenation for paths | static | mandatory | Code review sign-off | Windows `\` separator portability |
| TC-028 | Spec export: pathlib.Path in export.py | Code review confirms `export.py` uses `pathlib.Path` for all path operations | static | mandatory | Code review sign-off | Windows portability |

## Test Fixtures

The following fixtures MUST be defined in `token-monitor/tests/conftest.py` or at the top of the test module:

### `FIXTURE_NON_STREAMING`

```json
{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "model": "gpt-4o",
  "choices": [
    {
      "index": 0,
      "message": {"role": "assistant", "content": "Hello!"},
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 42,
    "completion_tokens": 7,
    "total_tokens": 49
  }
}
```

Expected extraction: `prompt_tokens=42`, `completion_tokens=7`, `total_tokens=49`.

### `FIXTURE_SSE_CHUNKS`

A sequence of SSE data lines simulating a streaming response. The first two chunks carry delta content only; the final chunk carries the `usage` object:

```python
FIXTURE_SSE_CHUNKS = [
    b'data: {"id":"x","choices":[{"delta":{"content":"Hi"},"finish_reason":null}]}\n\n',
    b'data: {"id":"x","choices":[{"delta":{"content":"!"},"finish_reason":null}]}\n\n',
    b'data: {"id":"x","choices":[{"delta":{},"finish_reason":"stop"}],"usage":{"prompt_tokens":10,"completion_tokens":2,"total_tokens":12}}\n\n',
    b'data: [DONE]\n\n',
]
```

Expected extraction from final chunk: `prompt_tokens=10`, `completion_tokens=2`, `total_tokens=12`.

### `FIXTURE_SSE_NO_USAGE`

Same structure as `FIXTURE_SSE_CHUNKS` but the final chunk omits the `"usage"` key. Used for TC-005.

### `FIXTURE_DB_ROWS` (pytest fixture using `tmp_path`)

```python
@pytest.fixture
def seeded_db(tmp_path):
    db_path = tmp_path / "test_usage.db"
    storage = TokenStorage(db_path)
    for i in range(3):
        storage.insert_event({
            "ts": f"2026-07-13T00:0{i}:00Z",
            "agent": "opencode",
            "model": "gpt-4o",
            "endpoint": "api.openai.com",
            "prompt_tokens": 10 * (i + 1),
            "completion_tokens": 5 * (i + 1),
            "total_tokens": 15 * (i + 1),
            "session_id": None if i == 0 else f"sess-{i}",
            "request_id": f"req-{i}",
            "duration_ms": 200 + i * 10,
            "status_code": 200,
            "project": "my-project",
            "phase": "sdd-apply",
        })
    storage.close()
    return db_path
```

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Operational scope | Verify `design.md#Operational Considerations` is marked `No aplica.` and no production monitoring, backup, or retention checks are required | static | Section reads `No aplica.` — confirmed during inputs read | N/A |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| SEC-SQL | Parameterized SQL queries | Yes | TC-025 (static code review) | covered | None |
| SEC-LOG | No sensitive data in logs or DB | Yes | TC-026 (static code review) | covered | None |
| SEC-CLASSIFY | Design classification assessment | Yes | `design.md#secure-development-design` classification: no-impact; verified during inputs read | covered | None |

All other corporate security categories (auth, sessions, PAN, secrets, ACL, file upload, network exposure) are not applicable per the no-impact classification in `design.md#secure-development-design`. Exhaustive `N/A` row decisions are owned by canonical `review-security-report.json`.

## Coverage Targets

| Module | Target | Rationale |
| --- | --- | --- |
| `token_monitor.py` (extraction + agent detection functions) | ≥ 90 % line coverage | Core capture logic; all branches (JSON vs SSE, known/unknown agent, missing usage) are unit-testable with mock flows |
| `storage.py` | ≥ 90 % line coverage | All paths (init, WAL pragma, insert success, insert error) covered by TC-017–019 |
| `export.py` | ≥ 85 % line coverage | JSONL + CSV + missing DB paths covered by TC-020–022; `argparse` parsing branches are lower-risk |

> **Tooling note**: `pytest-cov` is not configured at the repository root. Coverage commands (`pytest --cov=token-monitor`) MUST be added to `token-monitor/` requirements/dev dependencies and run locally. No CI pipeline is currently configured; coverage cannot be claimed as passing evidence until a runner is wired up.

## Evidence Expectations

- Mandatory cases (TC-001–022, TC-025–028) require implementation, execution, and passing test output or code review sign-off before verification is considered complete.
- Non-mandatory cases (TC-023, TC-024) are integration-scope documentation items; they do not block verification but SHOULD be run manually before the first production-like use of the tool.
- Static checks (TC-025–028) are satisfied by code review sign-off during `sdd-apply` or `sdd-review`; no automated SAST tool is available.
- `pytest` test runner is not yet configured in this repository's root `openspec/config.yaml`. The tool introduces its own test environment under `token-monitor/tests/`; the tasks phase MUST include a `requirements-dev.txt` or equivalent that pins `pytest` and optionally `pytest-cov`.
- Runtime/build/lint/type/format/coverage tooling is unavailable at the repository level. This is reported as a constraint, not as passing evidence.

## Open Questions

- None.
