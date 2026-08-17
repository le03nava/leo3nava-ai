from __future__ import annotations

import csv
import io
import json
import sqlite3
from dataclasses import dataclass, field
from pathlib import Path
from typing import Any

import pytest

import export
import storage
import token_monitor


@dataclass
class DummyRequest:
    headers: dict[str, str] = field(default_factory=dict)
    pretty_host: str = "api.openai.com"
    host: str = "api.openai.com"
    timestamp_start: float = 1.0


@dataclass
class DummyResponse:
    headers: dict[str, str] = field(default_factory=dict)
    content: bytes = b""
    timestamp_end: float = 1.2
    status_code: int = 200

    def json(self) -> Any:
        return json.loads(self.content.decode("utf-8"))


@dataclass
class DummyFlow:
    id: str
    request: DummyRequest
    response: DummyResponse


# TC-001
def test_extract_usage_nonstreaming_success(FIXTURE_NON_STREAMING: dict[str, object]) -> None:
    usage = token_monitor.extract_usage(json.dumps(FIXTURE_NON_STREAMING).encode("utf-8"))
    assert usage == {"prompt_tokens": 42, "completion_tokens": 7, "total_tokens": 49}


# TC-002
def test_extract_usage_nonstreaming_missing_usage_returns_none() -> None:
    payload = {"id": "abc", "choices": []}
    assert token_monitor.extract_usage(json.dumps(payload).encode("utf-8")) is None


# TC-003
def test_extract_usage_nonstreaming_invalid_json_returns_none() -> None:
    assert token_monitor.extract_usage(b"this is not valid json") is None


# TC-004
def test_extract_sse_usage_from_final_chunk(FIXTURE_SSE_CHUNKS: list[str]) -> None:
    usage = token_monitor.extract_sse_usage(FIXTURE_SSE_CHUNKS)
    assert usage == {"prompt_tokens": 10, "completion_tokens": 2, "total_tokens": 12}


# TC-005
def test_extract_sse_usage_missing_usage_logs_warning(
    caplog: pytest.LogCaptureFixture,
    FIXTURE_SSE_NO_USAGE: list[str],
) -> None:
    with caplog.at_level("WARNING", logger="token-monitor"):
        usage = token_monitor.extract_sse_usage(FIXTURE_SSE_NO_USAGE)
    assert usage is None
    assert "without usage metadata" in caplog.text


# TC-006
def test_sse_buffer_released_after_response_processed(
    tmp_path: Path,
    FIXTURE_SSE_CHUNKS: list[str],
) -> None:
    addon = token_monitor.TokenMonitorAddon(db_path=tmp_path / "usage.db")
    flow = DummyFlow(
        id="flow-1",
        request=DummyRequest(headers={"User-Agent": "opencode/1.0"}),
        response=DummyResponse(
            headers={"Content-Type": "text/event-stream", "X-Request-Id": "req-1"},
            content="".join(FIXTURE_SSE_CHUNKS).encode("utf-8"),
        ),
    )
    addon.responseheaders(flow)
    assert "flow-1" in addon.sse_buffers

    addon.response(flow)

    assert "flow-1" not in addon.sse_buffers
    addon.done()


# TC-007
def test_detect_agent_prefers_x_agent_name() -> None:
    headers = {"X-Agent-Name": "copilot-chat", "User-Agent": "curl/8.0"}
    assert token_monitor.detect_agent(headers) == "copilot-chat"


# TC-008
def test_detect_agent_from_user_agent_opencode() -> None:
    assert token_monitor.detect_agent({"User-Agent": "OpenCode/2.0"}) == "opencode"


# TC-009
def test_detect_agent_from_user_agent_claude() -> None:
    assert token_monitor.detect_agent({"User-Agent": "claude-code/0.9"}) == "claude-code"


# TC-010
def test_detect_agent_unknown_user_agent_fallback() -> None:
    assert token_monitor.detect_agent({"User-Agent": "curl/7.88"}) == "unknown"


# TC-011
def test_detect_agent_without_headers() -> None:
    assert token_monitor.detect_agent({}) == "unknown"


# TC-012
def test_extract_headers_project_present() -> None:
    meta = token_monitor.extract_headers({"X-Project": "my-project"})
    assert meta["project"] == "my-project"


# TC-013
def test_extract_headers_phase_present() -> None:
    meta = token_monitor.extract_headers({"X-Phase": "sdd-apply"})
    assert meta["phase"] == "sdd-apply"


# TC-014
def test_extract_headers_session_id_present() -> None:
    meta = token_monitor.extract_headers({"X-Session-ID": "sess-abc"})
    assert meta["session_id"] == "sess-abc"


# TC-015
def test_extract_headers_request_id_present() -> None:
    meta = token_monitor.extract_headers({"X-Request-Id": "req-123"})
    assert meta["request_id"] == "req-123"


# TC-016
def test_extract_headers_absent_are_none() -> None:
    meta = token_monitor.extract_headers({})
    assert meta == {"project": None, "phase": None, "session_id": None, "request_id": None}


# TC-017
def test_storage_insert_event_persists_full_row(tmp_path: Path) -> None:
    db_path = tmp_path / "usage.db"
    conn = storage.init_db(db_path)
    event = {
        "ts": "2026-07-13T01:00:00Z",
        "agent": "opencode",
        "model": "gpt-4o",
        "endpoint": "api.openai.com",
        "prompt_tokens": 42,
        "completion_tokens": 7,
        "total_tokens": 49,
        "session_id": "sess-1",
        "request_id": "req-1",
        "duration_ms": 215,
        "status_code": 200,
        "project": "my-project",
        "phase": "sdd-apply",
    }
    storage.insert_event(conn, event)

    row = conn.execute("SELECT * FROM token_events").fetchone()
    columns = [d[0] for d in conn.execute("SELECT * FROM token_events").description]
    assert row is not None
    row_map = dict(zip(columns, row))
    assert row_map["id"] == 1
    for key, value in event.items():
        assert row_map[key] == value
    conn.close()


# TC-018
def test_storage_init_enables_wal_mode(tmp_path: Path) -> None:
    conn = storage.init_db(tmp_path / "usage.db")
    mode = conn.execute("PRAGMA journal_mode").fetchone()[0]
    assert str(mode).lower() == "wal"
    conn.close()


# TC-019
def test_storage_insert_error_is_caught_and_logged(tmp_path: Path, capsys: pytest.CaptureFixture[str]) -> None:
    conn = storage.init_db(tmp_path / "usage.db")
    conn.execute("DROP TABLE token_events")
    conn.commit()

    storage.insert_event(
        conn,
        {
            "ts": "2026-07-13T01:00:00Z",
            "agent": "opencode",
        },
    )
    captured = capsys.readouterr()
    assert "Failed to insert token event" in captured.err
    conn.close()


# TC-020
def test_export_jsonl_outputs_all_rows_with_all_fields(seeded_db: Path) -> None:
    stream = io.StringIO()
    columns, rows = export._fetch_rows(seeded_db)
    export._write_jsonl(columns, rows, stream)
    lines = [line for line in stream.getvalue().splitlines() if line.strip()]
    assert len(lines) == 3

    parsed = [json.loads(line) for line in lines]
    required = {
        "id",
        "ts",
        "agent",
        "model",
        "endpoint",
        "prompt_tokens",
        "completion_tokens",
        "total_tokens",
        "session_id",
        "request_id",
        "duration_ms",
        "status_code",
        "project",
        "phase",
    }
    for item in parsed:
        assert required.issubset(item.keys())


# TC-021
def test_export_csv_outputs_empty_for_null_values(seeded_db: Path) -> None:
    stream = io.StringIO()
    columns, rows = export._fetch_rows(seeded_db)
    export._write_csv(columns, rows, stream)

    reader = csv.DictReader(io.StringIO(stream.getvalue()))
    read_rows = list(reader)
    assert len(read_rows) == 3
    assert read_rows[0]["session_id"] == ""


# TC-022
def test_export_missing_db_returns_code_1_and_stderr(tmp_path: Path, capsys: pytest.CaptureFixture[str]) -> None:
    missing = tmp_path / "does-not-exist.db"
    code = export.main(["--db-path", str(missing)])
    captured = capsys.readouterr()
    assert code == 1
    assert "Database file not found" in captured.err


# Extra required-by-slice behavior coverage from request constraints
def test_addon_non_streaming_inserts_event_with_header_fallbacks(
    tmp_path: Path,
    FIXTURE_NON_STREAMING: dict[str, object],
) -> None:
    addon = token_monitor.TokenMonitorAddon(db_path=tmp_path / "usage.db", filter_hosts=["openai"])
    payload = json.dumps(FIXTURE_NON_STREAMING).encode("utf-8")
    flow = DummyFlow(
        id="flow-2",
        request=DummyRequest(
            headers={
                "User-Agent": "github-copilot/1.0",
                "X-Project": "my-project",
                "X-Phase": "sdd-apply",
                "X-Session-ID": "sess-1",
            }
        ),
        response=DummyResponse(
            headers={"Content-Type": "application/json", "X-Request-Id": "req-99"},
            content=payload,
        ),
    )

    addon.response(flow)
    conn = sqlite3.connect(tmp_path / "usage.db")
    row = conn.execute(
        "SELECT agent, project, phase, session_id, request_id, prompt_tokens, completion_tokens, total_tokens FROM token_events"
    ).fetchone()
    assert row == ("copilot-chat", "my-project", "sdd-apply", "sess-1", "req-99", 42, 7, 49)
    conn.close()
    addon.done()


def test_addon_filter_hosts_skips_unmatched(
    tmp_path: Path,
    FIXTURE_NON_STREAMING: dict[str, object],
) -> None:
    addon = token_monitor.TokenMonitorAddon(db_path=tmp_path / "usage.db", filter_hosts=["api.openai.com"])
    flow = DummyFlow(
        id="flow-3",
        request=DummyRequest(headers={"User-Agent": "opencode/1.0"}, pretty_host="example.com", host="example.com"),
        response=DummyResponse(
            headers={"Content-Type": "application/json"},
            content=json.dumps(FIXTURE_NON_STREAMING).encode("utf-8"),
        ),
    )
    addon.response(flow)

    conn = sqlite3.connect(tmp_path / "usage.db")
    count = conn.execute("SELECT COUNT(*) FROM token_events").fetchone()[0]
    assert count == 0
    conn.close()
    addon.done()
