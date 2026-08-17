"""Export token monitor SQLite events into JSONL or CSV."""

from __future__ import annotations

import argparse
import csv
import json
import sqlite3
import sys
from pathlib import Path
from typing import Any, Iterable

DEFAULT_DB_PATH = Path("~/.token-monitor/usage.db").expanduser()


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Export token events from SQLite")
    parser.add_argument("--db-path", type=Path, default=DEFAULT_DB_PATH, help="SQLite database path")
    parser.add_argument("--format", dest="fmt", choices=("jsonl", "csv"), default="jsonl")
    parser.add_argument("--output", type=Path, default=None, help="Output file path; defaults to stdout")
    parser.add_argument("--from", dest="from_ts", default=None, help="Optional lower bound ISO8601 timestamp")
    parser.add_argument("--to", dest="to_ts", default=None, help="Optional upper bound ISO8601 timestamp")
    return parser


def _fetch_rows(db_path: Path, from_ts: str | None = None, to_ts: str | None = None) -> tuple[list[str], list[dict[str, Any]]]:
    where: list[str] = []
    params: list[Any] = []
    if from_ts:
        where.append("ts >= ?")
        params.append(from_ts)
    if to_ts:
        where.append("ts <= ?")
        params.append(to_ts)

    where_clause = f" WHERE {' AND '.join(where)}" if where else ""
    query = f"SELECT * FROM token_events{where_clause} ORDER BY ts ASC"

    conn = sqlite3.connect(db_path)
    try:
        cursor = conn.execute(query, tuple(params))
        columns = [desc[0] for desc in cursor.description]
        rows = [dict(zip(columns, row)) for row in cursor.fetchall()]
        return columns, rows
    finally:
        conn.close()


def _write_jsonl(columns: list[str], rows: Iterable[dict[str, Any]], output_stream: Any) -> None:
    _ = columns
    first = True
    for row in rows:
        if not first:
            output_stream.write("\n")
        output_stream.write(json.dumps(row, ensure_ascii=False))
        first = False


def _write_csv(columns: list[str], rows: Iterable[dict[str, Any]], output_stream: Any) -> None:
    writer = csv.DictWriter(output_stream, fieldnames=columns)
    writer.writeheader()
    for row in rows:
        normalized = {k: ("" if v is None else v) for k, v in row.items()}
        writer.writerow(normalized)


def main(argv: list[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    db_path = Path(args.db_path).expanduser()

    if not db_path.exists():
        print(f"Database file not found: {db_path}", file=sys.stderr)
        return 1

    columns, rows = _fetch_rows(db_path=db_path, from_ts=args.from_ts, to_ts=args.to_ts)

    if args.output:
        output_path = Path(args.output).expanduser()
        output_path.parent.mkdir(parents=True, exist_ok=True)
        with output_path.open("w", encoding="utf-8", newline="") as stream:
            if args.fmt == "jsonl":
                _write_jsonl(columns, rows, stream)
            else:
                _write_csv(columns, rows, stream)
        return 0

    if args.fmt == "jsonl":
        _write_jsonl(columns, rows, sys.stdout)
    else:
        _write_csv(columns, rows, sys.stdout)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
