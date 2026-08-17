# Design: Token Monitor

## Technical Approach

A local mitmproxy addon intercepts outbound LLM API traffic, extracts token usage from JSON and SSE responses, and persists structured events to SQLite. A separate export script reads the DB on demand. The proxy is transparent — it forwards all traffic unmodified and only records usage metadata from matched LLM endpoints.

## Architecture Decisions

### Decision: mitmproxy addon vs raw HTTP proxy

**Choice**: mitmproxy addon mode (`mitmdump --scripts token_monitor.py`)
**Alternatives considered**: Raw asyncio CONNECT proxy, `http.server` + `ssl`, Node.js-based proxy
**Rationale**: mitmproxy handles TLS MITM, HTTP/2, connection pooling, and SSE streaming out of the box. Writing a raw proxy would require reimplementing certificate generation, CONNECT tunnel handling, and chunked transfer parsing. The addon API (`response` hook) gives direct access to decoded bodies with zero boilerplate.

### Decision: SQLite WAL vs flat JSONL as primary store

**Choice**: SQLite at `~/.token-monitor/usage.db` with WAL mode
**Alternatives considered**: Append-only JSONL file (original proposal approach)
**Rationale**: SQLite supports concurrent reads (export script) while the proxy writes. WAL mode eliminates lock contention. Queryable storage enables future filtering without loading entire files. JSONL remains available as an export format.

### Decision: WAL mode enabled at DB init

**Choice**: `PRAGMA journal_mode=WAL` set once at connection open
**Alternatives considered**: Default rollback journal, WAL per-connection toggle
**Rationale**: WAL allows the export script to read without blocking proxy writes. Set once at init avoids repeated PRAGMA overhead. WAL is persistent — subsequent connections inherit it.

### Decision: mkcert for TLS CA

**Choice**: mkcert-generated local CA, one-time `mkcert -install`
**Alternatives considered**: mitmproxy's built-in CA (`~/.mitmproxy/mitmproxy-ca-cert.pem`), manual OpenSSL CA
**Rationale**: mitmproxy's built-in CA works but requires manual trust-store import on Windows. mkcert automates trust-store installation cross-platform with a single command. However — mitmproxy already generates its own CA on first run. **Revised**: Use mitmproxy's native CA and instruct users to trust `~/.mitmproxy/mitmproxy-ca-cert.cer` via `certutil` or mkcert's trust mechanism. This avoids configuring mitmproxy to use a custom CA.

### Decision: Agent detection priority

**Choice**: `X-Agent-Name` header > `User-Agent` substring match > `"unknown"`
**Alternatives considered**: Config file mapping, regex-based UA parsing
**Rationale**: Simple priority chain covers known agents. `X-Agent-Name` is explicit and avoids ambiguous UA parsing. Fallback to `"unknown"` ensures no data loss.

## Data Flow

```
  Agent (OpenCode/Claude Code/Copilot)
       │
       │  HTTPS request → LLM endpoint
       ▼
  ┌─────────────────────────┐
  │   mitmproxy (mitmdump)  │
  │   token_monitor.py addon│
  │                         │
  │  response hook fires    │
  │  ┌───────────────────┐  │
  │  │ JSON? → extract   │  │
  │  │ SSE?  → buffer →  │  │
  │  │         extract    │  │
  │  └────────┬──────────┘  │
  │           │              │
  │           ▼              │
  │  storage.insert_event() │
  └─────────────────────────┘
              │
              ▼
  ┌─────────────────────┐
  │  SQLite (WAL mode)  │
  │  ~/.token-monitor/  │
  │  usage.db           │
  └─────────┬───────────┘
            │
            ▼ (on demand)
  ┌─────────────────────┐
  │  export.py          │
  │  → JSONL or CSV     │
  └─────────────────────┘
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `token-monitor/token_monitor.py` | Create | mitmproxy addon class with `response` hook, SSE buffer dict, agent detection, usage extraction |
| `token-monitor/storage.py` | Create | SQLite connection init, WAL setup, schema creation, `insert_event(event: dict)` |
| `token-monitor/export.py` | Create | Standalone argparse CLI, reads SQLite, writes JSONL/CSV to stdout or file |
| `token-monitor/requirements.txt` | Create | `mitmproxy>=10.0` |
| `token-monitor/README.md` | Create | Setup (mkcert/CA trust), proxy start, agent config, export usage |
| `token-monitor/tests/test_token_monitor.py` | Create | Unit tests for parsing, agent detection, SSE buffering, storage, export |

## Interfaces / Contracts

```python
# storage.py
class TokenStorage:
    def __init__(self, db_path: Path):
        """Open/create SQLite DB, enable WAL, ensure schema."""

    def insert_event(self, event: dict) -> None:
        """Insert a token_events row. Keys: ts, agent, model, endpoint,
        prompt_tokens, completion_tokens, total_tokens, session_id."""

    def close(self) -> None:
        """Close the connection."""

# token_monitor.py — mitmproxy addon class
class TokenMonitor:
    def __init__(self):
        self.storage: TokenStorage
        self.sse_buffers: dict[str, list[bytes]]  # flow.id → chunks

    def response(self, flow: http.HTTPFlow) -> None:
        """Hook: extract usage from JSON responses."""

    def responseheaders(self, flow: http.HTTPFlow) -> None:
        """Hook: detect SSE streams, enable buffering."""
```

## Operational Considerations

No aplica. This is a local developer tool with no production deployment, no shared infrastructure, no monitoring targets, and no backup/retention requirements beyond the user's local filesystem.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | JSON usage extraction from various response shapes | pytest, mock flow objects |
| Unit | SSE chunk buffering and final-chunk detection | pytest, simulated chunk sequences |
| Unit | Agent detection priority logic | pytest, parametrized headers |
| Unit | Storage insert and WAL mode verification | pytest, temp SQLite DB |
| Unit | Export JSONL/CSV output correctness | pytest, pre-seeded DB, capture stdout |
| Integration | End-to-end: mitmdump + addon + real HTTP call | Manual / CI with local server stub |

## Migration / Rollout

No migration required. New standalone tool — no existing data, no shared state.

## Open Questions

None. All design decisions are confirmed.

## Secure Development Design

### Classification and Changed Surface

**Classification**: no-impact.

**Changed surface**: This change creates a brand-new standalone local tool (`token-monitor/`). It does not modify any existing application code, infrastructure, or deployed service. The tool runs on the developer's local machine only.

**Touched runtime surfaces**: Local filesystem (SQLite DB write), localhost network (proxy port binding), TLS certificate trust store (one-time mkcert install).

**Untouched runtime surfaces**: No authentication system, no session management, no user-facing application, no database with PII/PAN, no secrets management system, no access control layer, no production logging pipeline.

**Why no security category applies**:
- **Authentication/Sessions**: The proxy has no auth layer (out of scope per proposal). It binds to localhost only.
- **Sensitive data / PAN**: Token counts are operational metrics, not sensitive data. No PII, PAN, or credentials are stored — only integer counts and agent labels.
- **Secrets**: The tool does not store, transmit, or log API keys or credentials. Request bodies (which may contain prompts) are never persisted — only response `usage` metadata.
- **Permissions/Access control**: Single-user local tool. No multi-tenant access.
- **Files**: SQLite DB is user-owned, default permissions. No file upload/download surface.
- **Database access**: Local SQLite with no network exposure. No SQL injection vector — all writes use parameterized queries by design.
- **Sensitive logging**: The addon logs warnings to stderr only (missing usage fields). No request/response bodies, prompts, or completions are logged.

Omitted categories are reviewable omissions for downstream `review-security-report.json` validation.

### Exception and Evidence Policy

No exceptions are planned. No security categories are applicable. The tool's local-only, single-user, metadata-only nature means no security controls are required beyond standard coding hygiene (parameterized SQL queries, no secret logging), which are built into the implementation approach.
