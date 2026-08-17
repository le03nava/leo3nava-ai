# Token Usage Export Specification

## Purpose

Reads captured token event rows from the SQLite database and converts them to JSONL or CSV format on demand via a separate export script.

## Requirements

### Requirement: JSONL Export

The system MUST provide an export script that reads all rows from `token_events` and writes one JSON object per line to stdout or a specified output file. Each line MUST contain all eight fields: `ts`, `agent`, `model`, `endpoint`, `prompt_tokens`, `completion_tokens`, `total_tokens`, `session_id`. The script MUST accept `--db-path` and `--output` flags.

#### Scenario: Full JSONL export succeeds

- GIVEN the database contains three token event rows
- WHEN the export script is run with `--format jsonl`
- THEN stdout contains exactly three newline-delimited JSON objects, each with all eight fields

#### Scenario: JSONL export to file

- GIVEN the user specifies `--output usage.jsonl`
- WHEN the export script runs
- THEN `usage.jsonl` is created with one JSON object per line and no trailing newline after the last line

### Requirement: CSV Export

The system MUST support CSV export with a header row matching the eight field names. The script MUST accept `--format csv` to select this output mode. Null values MUST be represented as empty strings in the CSV.

#### Scenario: Full CSV export succeeds

- GIVEN the database contains rows where some `session_id` values are null
- WHEN the export script is run with `--format csv`
- THEN stdout contains a header row followed by data rows, and null `session_id` cells appear as empty strings

#### Scenario: CSV written to file

- GIVEN the user specifies `--output report.csv --format csv`
- WHEN the export script runs
- THEN `report.csv` is created with a valid CSV header and one data row per token event

### Requirement: Export CLI Interface

The export script MUST be invocable as `python export.py [--db-path PATH] [--format jsonl|csv] [--output FILE]`. All flags MUST have sensible defaults: `--db-path` defaults to `~/.token-monitor/usage.db`, `--format` defaults to `jsonl`, `--output` defaults to stdout. The script MUST exit with a non-zero code and an error message if the database file does not exist.

#### Scenario: Missing database exits with error

- GIVEN no database file exists at the specified path
- WHEN the export script is invoked
- THEN the script exits with code 1 and prints an error message to stderr

## Non-Functional Requirements

### Requirement: Proxy Overhead

The addon SHOULD add less than 5 ms of latency per intercepted flow for non-streaming responses under normal workload. SSE buffering MUST NOT consume unbounded memory; the per-flow buffer SHOULD be released after stream completion and row persistence.

### Requirement: Reliability — No Token Event Loss on Error

The addon MUST NOT lose a token event due to an internal exception. Any error during parsing or persistence MUST be caught, logged, and allow the flow to continue. The database MUST use WAL mode to reduce lock contention when the proxy and export script run concurrently.

### Requirement: Portability

The system MUST run on Windows 10/11 as the primary target. Python MUST be ≥ 3.10. All paths MUST use `pathlib.Path` to handle Windows path separators correctly. The mkcert CA install step is a one-time OS-level operation and MUST be documented in the README.

#### Scenario: Default DB path resolves on Windows

- GIVEN no `--db-path` flag is provided
- WHEN the proxy starts on a Windows machine
- THEN the database is created at `%USERPROFILE%\.token-monitor\usage.db`

## Out of Scope

- Org-level GitHub API usage data
- Real-time streaming export or push-based delivery
- Export filtering by date range or agent (deferred to future enhancement)
