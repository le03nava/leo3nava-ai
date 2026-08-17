# Proposal: Token Monitor

## Intent

AI API calls are invisible at cost level — there is no per-agent, per-session breakdown of token usage across GitHub Copilot and other LLM endpoints. Without structured logging at the network layer, cost attribution and quota alerting are impossible. This tool intercepts outbound LLM traffic via a local proxy, captures token counts from API responses, and stores them as queryable JSONL records.

## Scope

### In Scope
- `mitmproxy` addon that captures token usage from LLM API responses (OpenAI-compatible schema)
- TLS interception via `mkcert` self-signed CA installed in Windows user trust store
- SSE/streaming support: buffer all chunks per flow, extract `usage` from final chunk
- Agent detection via `User-Agent` + optional `X-Agent-Name` header
- JSONL output: one record per request with `ts`, `agent`, `model`, `endpoint`, `prompt_tokens`, `completion_tokens`, `total_tokens`, `session_id`
- CSV export script (separate from the addon)
- CLI configuration via `argparse` (mirrors `tools-py/report-exporter/` conventions)
- Unit tests under `token-monitor/tests/`

### Out of Scope
- Dashboard or web UI
- Database storage (SQLite, Postgres, etc.)
- Node.js or Go runtimes — Python only
- Real-time alerting or quota enforcement
- Multi-machine or remote proxy deployment
- Authentication/authorization on the proxy port

## Capabilities

### New Capabilities
- `token-usage-capture`: Intercepts LLM API traffic via mitmproxy addon, extracts token usage fields, and appends structured records to a JSONL log file.
- `token-usage-export`: Reads the JSONL log and converts it to CSV via a separate export script.

### Modified Capabilities
None.

## Approach

1. **Proxy layer**: `mitmproxy` in addon mode (`mitmdump --scripts token_monitor.py`). The addon hooks `response` events, filters on LLM endpoint patterns, and parses JSON bodies (or buffers SSE chunks) to extract `usage`.
2. **TLS**: `mkcert` generates a local CA; `mkcert -install` adds it to the Windows trust store once. Python's `requests`/`httpx` and the OS both trust it natively.
3. **Streaming**: The addon accumulates SSE data chunks per `flow.id`; on stream end it parses the last `data:` line for `usage`.
4. **Output**: Each captured event is appended as a JSON line to a configurable output file (default: `token_usage.jsonl`).
5. **Config**: `argparse` flags for `--output`, `--port`, `--filter-host`; mirrors `tools-py/report-exporter/` pattern.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `token-monitor/token_monitor.py` | New | mitmproxy addon — core capture logic |
| `token-monitor/export_csv.py` | New | JSONL → CSV converter |
| `token-monitor/requirements.txt` | New | `mitmproxy`, `mkcert` (CLI dep note) |
| `token-monitor/README.md` | New | Setup and usage |
| `token-monitor/tests/` | New | Unit tests for parsing and export logic |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| mkcert CA not trusted by target process | Low | Document install step; test with curl before running addon |
| SSE final chunk has no `usage` field | Low | Guard with `if "usage" in data`; log warning, skip record |
| mitmproxy addon API changes between versions | Low | Pin version in `requirements.txt` |
| False-positive agent detection from User-Agent | Medium | Fall back to `"unknown"` agent label; session_id logged as null |

## Rollback Plan

Tool is fully standalone — no changes to existing code or infra. Rollback = stop the proxy process and delete `token-monitor/`. No database migrations, no config changes, no side effects on other tooling.

## Dependencies

- `mitmproxy` ≥ 10.x (Python package)
- `mkcert` (system binary, one-time install for TLS CA)
- Python ≥ 3.10

## Success Criteria

- [ ] Proxy captures `prompt_tokens`, `completion_tokens`, `total_tokens` from a non-streaming GitHub Copilot API call
- [ ] Proxy captures token usage from a streaming (SSE) response
- [ ] JSONL file contains correctly structured records with all 8 fields
- [ ] CSV export produces a valid CSV from a sample JSONL file
- [ ] Unit tests pass for parsing and export logic
- [ ] README documents mkcert install, proxy start, and export usage
