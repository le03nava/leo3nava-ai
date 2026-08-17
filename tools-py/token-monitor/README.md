# token-monitor

Local mitmproxy addon + exporter to capture token usage from LLM API traffic and store it in SQLite.

## Prerequisites

- Python 3.11+
- `mitmproxy` (installed from `requirements.txt`)
- `mkcert` (for local certificate trust setup)
  - Windows: `winget install FiloSottile.mkcert`
  - macOS: `brew install mkcert`
  - Linux: download binary from https://github.com/FiloSottile/mkcert/releases

## One-time setup (certificate trust)

1. Install mkcert from: https://github.com/FiloSottile/mkcert
2. Install local CA into your OS trust store:

```bash
mkcert -install
```

3. Install Python dependencies:

```bash
python -m pip install -r requirements.txt
python -m pip install -r requirements-dev.txt
```

## Run the proxy

From `token-monitor/`:

```bash
mitmdump -s token_monitor.py --set db_path=~/.token-monitor/usage.db --listen-port 8080
```

Optional host filtering (comma-separated):

```bash
mitmdump -s token_monitor.py --set db_path=~/.token-monitor/usage.db --listen-port 8080 --set filter_host=api.openai.com,api.githubcopilot.com
```

## Agent configuration

### OpenCode

- Configure the client to use the local proxy (`http://127.0.0.1:8080`)
- Typical environment variables:

```bash
set HTTP_PROXY=http://127.0.0.1:8080
set HTTPS_PROXY=http://127.0.0.1:8080
```

- Send header: `X-Agent-Name: opencode`
- Optional context headers:
  - `X-Project: <project-name>`
  - `X-Phase: <phase-name>`
  - `X-Session-ID: <session-id>`

### Claude Code

- Configure HTTPS proxy to `127.0.0.1:8080`
- Typical environment variable:

```bash
set HTTPS_PROXY=http://127.0.0.1:8080
```

- Send `X-Agent-Name: claude-code` (recommended)
- Without custom header, detection falls back to `User-Agent` substring `claude`

### GitHub Copilot

- Configure client proxy to `127.0.0.1:8080`
- Configure your IDE/network proxy settings to route Copilot HTTP(S) traffic through `127.0.0.1:8080`
- Send `X-Agent-Name: copilot-chat` when possible
- Without custom header, detection falls back to `User-Agent` containing `copilot` / `github-copilot`

## Export data

From `token-monitor/`:

### JSONL (stdout)

```bash
python export.py --db-path ~/.token-monitor/usage.db --format jsonl
```

### JSONL (file)

```bash
python export.py --db-path ~/.token-monitor/usage.db --format jsonl --output usage.jsonl
```

### CSV (stdout)

```bash
python export.py --db-path ~/.token-monitor/usage.db --format csv
```

### CSV (file + date range)

```bash
python export.py --db-path ~/.token-monitor/usage.db --format csv --output usage.csv --from 2026-07-13T00:00:00Z --to 2026-07-13T23:59:59Z
```

## SQLite schema reference

```sql
CREATE TABLE token_events (
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
```

## Custom headers reference

- `X-Agent-Name` (highest priority for agent identity)
- `X-Project`
- `X-Phase`
- `X-Session-ID`

Related header used for request correlation:

- `X-Request-Id` (read from response headers when available)

## Test

Run unit tests from `token-monitor/`:

```bash
pytest tests/
```

## Manual integration checks (non-blocking)

### TC-023 — Non-streaming end-to-end

1. Run proxy command above.
2. Send one non-streaming completion request through the proxy.
3. Verify exactly one new row appears in `token_events` with expected token fields.

### TC-024 — SSE end-to-end

1. Run proxy command above.
2. Send one streaming completion request through the proxy.
3. Verify exactly one new row appears after stream completion and token fields come from final SSE usage chunk.
