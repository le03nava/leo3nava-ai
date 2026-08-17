# sdd-cost-tracker

Local SDD phase cost tracker implemented as a Go module with SQLite (`modernc.org/sqlite`, no CGO).

## Build

```bash
CGO_ENABLED=0 go build -o sdd-cost-tracker .
```

## Run

### HTTP mode

```bash
./sdd-cost-tracker --port 7438
```

Server binds to `127.0.0.1` only.

### MCP stdio mode

```bash
./sdd-cost-tracker --mcp
```

### MCP registration (opencode)

Example `opencode.json` snippet:

```json
{
  "mcpServers": {
    "sdd-cost-tracker": {
      "command": "/path/to/sdd-cost-tracker",
      "args": ["--mcp"]
    }
  }
}
```

## Configuration

| CLI Flag | Env Var | Default | Description |
|---|---|---|---|
| `--port` | `SDD_COST_PORT` | `7438` | HTTP port for localhost server |
| `--db-path` | `SDD_COST_DB_PATH` | `~/.sdd-cost-tracker/db.sqlite` | SQLite database file path |

CLI flags override env vars.

## MCP Tools

### `cost_query`

Returns matching `PhaseRecord` rows.

Parameters:

- `project` (string, optional; when omitted, queries across all projects)
- `change_name` (string, optional)
- `phase` (string, optional)

Example arguments:

```json
{
  "project": "leo3nava_ai",
  "change_name": "sdd-cost-tracker",
  "phase": "apply"
}
```

### `cost_summary`

Returns aggregated `ChangeSummary` rows for a project.

Parameters:

- `project` (string, optional; when omitted, returns summary across all projects)

Example arguments:

```json
{
  "project": "leo3nava_ai"
}
```

## HTTP Endpoints

All responses include `Content-Type: application/json`.

### `GET /health`

Response:

```json
{"status":"ok"}
```

### `POST /phases`

Request body:

```json
{
  "project": "leo3nava_ai",
  "change_name": "sdd-cost-tracker",
  "phase": "apply",
  "session_id": "ses_123",
  "tokens_input": 100,
  "tokens_output": 80,
  "tokens_reasoning": 20,
  "tokens_cache_read": 5,
  "tokens_cache_write": 2,
  "cost_usd": 0.17
}
```

Responses:

- Insert (new `session_id`): HTTP `201` + `{"ok":true}`
- Upsert update (existing `session_id`): HTTP `200` + `{"ok":true}`
- Validation error: HTTP `400` + `{"error":"..."}`

### `GET /changes?project=X`

Response:

```json
{
  "changes": ["sdd-cost-tracker", "token-monitor"]
}
```

Missing `project` returns `400` with `{"error":"project is required"}`.

### `GET /changes/{name}?project=X`

Response:

```json
{
  "phases": [
    {
      "id": 1,
      "project": "leo3nava_ai",
      "change_name": "sdd-cost-tracker",
      "phase": "apply",
      "session_id": "ses_123",
      "tokens_input": 100,
      "tokens_output": 80,
      "tokens_reasoning": 20,
      "tokens_cache_read": 5,
      "tokens_cache_write": 2,
      "cost_usd": 0.17
    }
  ]
}
```

If there is no match, returns `{"phases":[]}` (HTTP 200).

### `GET /summary?project=X`

Response:

```json
{
  "summary": [
    {
      "project": "leo3nava_ai",
      "change_name": "sdd-cost-tracker",
      "tokens_input": 250,
      "tokens_output": 180,
      "tokens_reasoning": 50,
      "tokens_cache_read": 10,
      "tokens_cache_write": 4,
      "cost_usd": 0.42
    }
  ]
}
```

Missing `project` returns `400` with `{"error":"project is required"}`.
