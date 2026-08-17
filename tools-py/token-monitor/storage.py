"""SQLite storage helpers for token-monitor.

This module is stdlib-only and Windows-friendly (pathlib-based).
"""

from __future__ import annotations

import sqlite3
import sys
from pathlib import Path
from typing import Any, Mapping

SCHEMA_VERSION = 1

TABLE_SQL = """
CREATE TABLE IF NOT EXISTS token_events (
  id                INTEGER PRIMARY KEY AUTOINCREMENT,
  ts                TEXT NOT NULL,
  agent             TEXT,
  model             TEXT,
  endpoint          TEXT,
  prompt_tokens     INTEGER,
  completion_tokens INTEGER,
  total_tokens      INTEGER,
  session_id        TEXT,
  request_id        TEXT,
  duration_ms       INTEGER,
  status_code       INTEGER,
  project           TEXT,
  phase             TEXT
)
"""


def _resolve_db_path(db_path: str | Path) -> Path:
    path = Path(db_path).expanduser()
    return path


def create_connection(db_path: str | Path) -> sqlite3.Connection:
    """Create a SQLite connection with WAL mode enabled."""
    path = _resolve_db_path(db_path)
    path.parent.mkdir(parents=True, exist_ok=True)

    conn = sqlite3.connect(path)
    conn.execute("PRAGMA journal_mode=WAL")
    return conn


def init_db(db_path: str | Path) -> sqlite3.Connection:
    """Initialize DB schema (version 1) and return an open connection."""
    conn = create_connection(db_path)

    current_version = int(conn.execute("PRAGMA user_version").fetchone()[0])

    conn.execute(TABLE_SQL)

    if current_version < SCHEMA_VERSION:
        conn.execute(f"PRAGMA user_version={SCHEMA_VERSION}")
    elif current_version > SCHEMA_VERSION:
        print(
            (
                "[token-monitor] Warning: database user_version "
                f"({current_version}) is newer than supported ({SCHEMA_VERSION})."
            ),
            file=sys.stderr,
        )

    conn.commit()
    return conn


def insert_event(conn: sqlite3.Connection, event: Mapping[str, Any]) -> None:
    """Insert one token event row.

    Never raises: catches and logs all exceptions to stderr.
    Missing nullable fields default to None.
    """
    try:
        conn.execute(
            """
            INSERT INTO token_events (
              ts,
              agent,
              model,
              endpoint,
              prompt_tokens,
              completion_tokens,
              total_tokens,
              session_id,
              request_id,
              duration_ms,
              status_code,
              project,
              phase
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (
                event["ts"],
                event.get("agent", None),
                event.get("model", None),
                event.get("endpoint", None),
                event.get("prompt_tokens", None),
                event.get("completion_tokens", None),
                event.get("total_tokens", None),
                event.get("session_id", None),
                event.get("request_id", None),
                event.get("duration_ms", None),
                event.get("status_code", None),
                event.get("project", None),
                event.get("phase", None),
            ),
        )
        conn.commit()
    except Exception as exc:  # noqa: BLE001 - required by contract (never raise)
        print(f"[token-monitor] Failed to insert token event: {exc}", file=sys.stderr)
