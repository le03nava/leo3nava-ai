from __future__ import annotations

from pathlib import Path

import pytest

import storage


@pytest.fixture
def FIXTURE_NON_STREAMING() -> dict[str, object]:
    return {
        "id": "chatcmpl-abc123",
        "choices": [{"message": {"content": "Hello"}}],
        "usage": {"prompt_tokens": 42, "completion_tokens": 7, "total_tokens": 49},
        "model": "gpt-4o",
    }


@pytest.fixture
def FIXTURE_SSE_CHUNKS() -> list[str]:
    return [
        'data: {"id":"x","choices":[{"delta":{"content":"He"}}]}\n\n',
        'data: {"id":"x","choices":[{"delta":{"content":"llo"}}]}\n\n',
        'data: {"id":"x","choices":[],"usage":{"prompt_tokens":10,"completion_tokens":2,"total_tokens":12}}\n\n',
        'data: [DONE]\n\n',
    ]


@pytest.fixture
def FIXTURE_SSE_NO_USAGE() -> list[str]:
    return [
        'data: {"id":"x","choices":[{"delta":{"content":"Hi"}}]}\n\n',
        'data: [DONE]\n\n',
    ]


@pytest.fixture
def seeded_db(tmp_path: Path) -> Path:
    db_path = tmp_path / "test_usage.db"
    conn = storage.init_db(db_path)
    for i in range(3):
        storage.insert_event(
            conn,
            {
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
            },
        )
    conn.close()
    return db_path
