# call-level-cost-tracking Specification

## Purpose

Records individual LLM invocations (one row per assistant message) with token and cost metrics,
queryable by session. Provides storage schema, Go methods, and HTTP endpoints alongside the
existing phase-level data.

---

## Requirements

### Requirement: Calls Table Schema

The binary MUST create the `calls` table on startup using `CREATE TABLE IF NOT EXISTS`:

```sql
CREATE TABLE IF NOT EXISTS calls (
  id                   INTEGER PRIMARY KEY AUTOINCREMENT,
  session_id           TEXT    NOT NULL,
  call_index           INTEGER NOT NULL,
  model_id             TEXT,
  provider_id          TEXT,
  tokens_input         INTEGER DEFAULT 0,
  tokens_output        INTEGER DEFAULT 0,
  tokens_reasoning     INTEGER DEFAULT 0,
  tokens_cache_read    INTEGER DEFAULT 0,
  tokens_cache_write   INTEGER DEFAULT 0,
  cost_usd             REAL    DEFAULT 0,
  recorded_at          INTEGER
);
```

No FK constraint SHALL be enforced between `calls.session_id` and `phases.session_id`.

#### Scenario: Auto-create on fresh DB

- GIVEN no DB file exists
- WHEN the binary starts
- THEN both `phases` and `calls` tables are created without error

#### Scenario: Idempotent init on existing DB

- GIVEN the `calls` table already exists
- WHEN the binary restarts
- THEN startup succeeds and no duplicate table is created

---

### Requirement: InsertCall Storage Method

The storage layer MUST expose `InsertCall(ctx context.Context, r CallRecord) error`.
It MUST insert one row per call. It MUST return a non-nil error if `session_id` is empty.

#### Scenario: Successful insert

- GIVEN a `CallRecord` with `session_id = "s1"` and `call_index = 0`
- WHEN `InsertCall` is called
- THEN the row is persisted and nil error is returned

#### Scenario: Reject missing session_id

- GIVEN a `CallRecord` with an empty `session_id`
- WHEN `InsertCall` is called
- THEN a non-nil error is returned and no row is inserted

---

### Requirement: GetCallsBySession Storage Method

The storage layer MUST expose `GetCallsBySession(ctx context.Context, sessionID string) ([]CallRecord, error)`.
It MUST return rows ordered by `call_index` ascending. It MUST return an empty slice (not an error)
when no rows match.

#### Scenario: Returns rows in order

- GIVEN three calls for `session_id = "s1"` with indices 0, 1, 2
- WHEN `GetCallsBySession(ctx, "s1")` is called
- THEN a slice of 3 records in ascending `call_index` order is returned

#### Scenario: No matching rows

- GIVEN no calls for `session_id = "ghost"`
- WHEN `GetCallsBySession(ctx, "ghost")` is called
- THEN an empty slice and nil error are returned

---

### Requirement: POST /calls Endpoint

The server MUST expose `POST /calls`. The request body MUST be JSON matching `CallRecord`.
It MUST return 201 on success, 400 on malformed body or missing `session_id`, 500 on storage error.

#### Scenario: Insert call returns 201

- GIVEN a valid JSON body with `session_id = "s1"` and `call_index = 0`
- WHEN `POST /calls` is called
- THEN the server returns 201 and the row is stored

#### Scenario: Missing session_id returns 400

- GIVEN a JSON body without `session_id`
- WHEN `POST /calls` is called
- THEN the server returns 400

---

### Requirement: GET /calls Endpoint

The server MUST expose `GET /calls?session_id=X`. It MUST return a JSON array of `CallRecord`
values ordered by `call_index` ascending. It MUST return 400 if `session_id` is absent.
It MUST return an empty JSON array when no records match.

#### Scenario: Returns records for session

- GIVEN two calls stored for `session_id = "s1"`
- WHEN `GET /calls?session_id=s1` is called
- THEN the server returns 200 with a JSON array of 2 records

#### Scenario: Missing session_id returns 400

- GIVEN no `session_id` query parameter
- WHEN `GET /calls` is called
- THEN the server returns 400

---

### Requirement: Call-Level Test Coverage

Tests MUST cover: `calls` table creation, `InsertCall` success and rejection, `GetCallsBySession`
with and without results, `POST /calls` 201/400, `GET /calls` 200/400. All MUST pass with `go test ./...`.

#### Scenario: Test suite green

- GIVEN the module builds with no CGO
- WHEN `go test ./...` is run
- THEN all call-level tests pass with zero failures
