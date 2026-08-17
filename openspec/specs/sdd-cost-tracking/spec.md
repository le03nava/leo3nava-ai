# Delta for sdd-cost-tracking

**Change**: sdd-cost-tracker
**Capability**: sdd-cost-tracking
**Type**: new

## Purpose

Records and queries token consumption and USD cost per SDD phase, grouped by project and change name. Exposed via MCP tools (stdio) and an HTTP REST API from a single Go binary with no CGO or external runtime dependencies.

---

## ADDED Requirements

### Requirement: HTTP Server Startup

The binary MUST start an HTTP server on a configurable port (default 7438). The port MUST be settable via a CLI flag or environment variable. The server MUST be ready to accept requests before the process signals readiness.

#### Scenario: Default port binding

- GIVEN the binary is started with no port argument
- WHEN the server initializes
- THEN it listens on port 7438

#### Scenario: Custom port via flag

- GIVEN the binary is started with `--port 9000`
- WHEN the server initializes
- THEN it listens on port 9000

---

### Requirement: Health Endpoint

The server MUST expose `GET /health`. It MUST return HTTP 200 with a non-empty body indicating liveness. No authentication is required.

#### Scenario: Liveness check on fresh DB

- GIVEN the server has started and the DB is initialized
- WHEN a client sends `GET /health`
- THEN the response status is 200

---

### Requirement: Phase Upsert Endpoint

The server MUST expose `POST /phases`. It MUST insert a new row if `session_id` is not present, or replace the row if it already exists (upsert by `session_id`). The request body MUST be JSON matching the phases schema. The endpoint MUST return HTTP 201 on new insert, HTTP 200 on upsert update, and HTTP 400 on a malformed or missing required field.

#### Scenario: Insert new phase

- GIVEN no row with `session_id = "s1"` exists
- WHEN `POST /phases` is called with a valid JSON body including `session_id = "s1"`
- THEN the server returns 201 and the row is stored

#### Scenario: Upsert existing phase

- GIVEN a row with `session_id = "s1"` already exists
- WHEN `POST /phases` is called again with `session_id = "s1"` and updated `cost_usd`
- THEN the server returns 200 and the row reflects the new `cost_usd`

#### Scenario: Reject missing required fields

- GIVEN a request body with no `project` field
- WHEN `POST /phases` is called
- THEN the server returns 400

---

### Requirement: List Changes Endpoint

The server MUST expose `GET /changes?project=X`. It MUST return all distinct `change_name` values for the given project. It MUST return HTTP 400 if `project` is absent.

#### Scenario: List changes for known project

- GIVEN phases for project `"myproj"` exist with two distinct `change_name` values
- WHEN `GET /changes?project=myproj` is called
- THEN the response contains both change names

#### Scenario: Missing project param

- GIVEN no `project` query param
- WHEN `GET /changes` is called
- THEN the server returns 400

---

### Requirement: Get Change Phases Endpoint

The server MUST expose `GET /changes/:name?project=X`. It MUST return all phase rows for the given change within the given project. It MUST return HTTP 400 if `project` is absent.

#### Scenario: Phases for a known change

- GIVEN three phases exist for `change_name = "feat-x"` in `project = "myproj"`
- WHEN `GET /changes/feat-x?project=myproj` is called
- THEN the response contains all three rows

---

### Requirement: Cost Summary Endpoint

The server MUST expose `GET /summary?project=X`. It MUST return aggregated `tokens_input`, `tokens_output`, `tokens_reasoning`, `tokens_cache_read`, `tokens_cache_write`, and `cost_usd` per `change_name` for the given project. It MUST return HTTP 400 if `project` is absent.

#### Scenario: Aggregated summary

- GIVEN two phases for `change_name = "feat-x"` with `cost_usd = 0.10` each
- WHEN `GET /summary?project=myproj` is called
- THEN the response shows `cost_usd = 0.20` for `"feat-x"`

---

### Requirement: MCP cost_query Tool

The binary MUST expose a `cost_query` MCP tool over stdio. It MUST accept `project`, `change_name`, and `phase` as optional filter parameters and return matching phase rows as JSON. It MUST return an empty list (not an error) when no rows match.

#### Scenario: Query by project

- GIVEN phases exist for `project = "myproj"` and `project = "other"`
- WHEN `cost_query` is called with `project = "myproj"`
- THEN only rows for `"myproj"` are returned

#### Scenario: No matches

- GIVEN no rows exist for `project = "ghost"`
- WHEN `cost_query` is called with `project = "ghost"`
- THEN the result is an empty list with no error

---

### Requirement: MCP cost_summary Tool

The binary MUST expose a `cost_summary` MCP tool over stdio. It MUST return aggregated cost and token totals per `change_name`, grouped by `project`. It MAY accept an optional `project` filter.

#### Scenario: Summary grouped by project

- GIVEN phases for two projects with different change names
- WHEN `cost_summary` is called with no filter
- THEN the result groups totals by project and change name

---

### Requirement: SQLite Storage Schema

The binary MUST create the following table on first startup if it does not exist:

```sql
CREATE TABLE phases (
  id                   INTEGER PRIMARY KEY AUTOINCREMENT,
  project              TEXT    NOT NULL,
  change_name          TEXT    NOT NULL,
  phase                TEXT    NOT NULL,
  session_id           TEXT    NOT NULL UNIQUE,
  model_id             TEXT,
  provider_id          TEXT,
  tokens_input         INTEGER DEFAULT 0,
  tokens_output        INTEGER DEFAULT 0,
  tokens_reasoning     INTEGER DEFAULT 0,
  tokens_cache_read    INTEGER DEFAULT 0,
  tokens_cache_write   INTEGER DEFAULT 0,
  cost_usd             REAL    DEFAULT 0,
  started_at           INTEGER,
  completed_at         INTEGER
);
```

The DB MUST be stored at a configurable path (default `~/.sdd-cost-tracker/db.sqlite`). The path MUST be settable via CLI flag or environment variable. The driver MUST be `modernc.org/sqlite` (pure-Go, no CGO).

#### Scenario: Auto-create schema on fresh DB

- GIVEN no DB file exists at the configured path
- WHEN the binary starts
- THEN the `phases` table is created automatically

#### Scenario: Custom DB path

- GIVEN `--db-path /tmp/test.sqlite` is passed
- WHEN the binary starts
- THEN the DB file is created at `/tmp/test.sqlite`

---

### Requirement: Non-Functional — Single Binary, No CGO

The binary MUST compile with `CGO_ENABLED=0`. It MUST produce a single self-contained executable with no external runtime dependencies. It MUST support cross-compilation for Linux, macOS, and Windows (amd64/arm64).

#### Scenario: CGO-free build

- GIVEN `CGO_ENABLED=0 go build ./...` is run
- WHEN the build completes
- THEN no CGO or dynamic link errors occur and a binary is produced

---

### Requirement: Unit Test Coverage

Co-located tests under `tools-go/sdd-cost-tracker/` MUST cover: DB schema initialization, `POST /phases` insert, `POST /phases` upsert, `GET /health`, `GET /changes`, `GET /summary`, `cost_query` tool, `cost_summary` tool, and the new call-level tests: `calls` table creation, `InsertCall` success and missing-field rejection, `GetCallsBySession` with results and with no results, `POST /calls` 201/400 cases, and `GET /calls` 200/400 cases. All tests MUST pass with `go test ./...`.

#### Scenario: Full test suite green

- GIVEN the module is built with no CGO
- WHEN `go test ./...` is run
- THEN all tests pass with zero failures
