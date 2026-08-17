# Delta for sdd-cost-tracking

## MODIFIED Requirements

### Requirement: Unit Test Coverage

Co-located tests under `tools-go/sdd-cost-tracker/` MUST cover: DB schema initialization,
`POST /phases` insert, `POST /phases` upsert, `GET /health`, `GET /changes`, `GET /summary`,
`cost_query` tool, `cost_summary` tool, `calls` table creation, `InsertCall` success and
missing-field rejection, `GetCallsBySession` with results and with no results,
`POST /calls` 201/400 cases, and `GET /calls` 200/400 cases.
All tests MUST pass with `go test ./...`.

(Previously: covered only phases-related tests; call-level test cases were not required.)

#### Scenario: Full test suite green

- GIVEN the module is built with no CGO
- WHEN `go test ./...` is run
- THEN all tests pass with zero failures

---

## ADDED Requirements

### Requirement: Plugin callIndex Counter

The TypeScript plugin MUST maintain a per-session `callIndex` counter initialized to 0 when
a new session starts. The counter MUST increment by 1 after each successful `postCall` dispatch.
The counter MUST be independent of any existing `postPhase` logic.

#### Scenario: Counter increments per message

- GIVEN the plugin has dispatched two call records for the current session
- WHEN a third assistant message with token data arrives
- THEN `postCall` is called with `call_index = 2`

#### Scenario: Counter resets on new session

- GIVEN a new session starts
- WHEN the first assistant message arrives
- THEN `postCall` is called with `call_index = 0`

---

### Requirement: Plugin postCall Fire-and-Forget

The plugin MUST dispatch `POST /calls` on every `message.updated` event where the message role
is `assistant` and token data is present. The payload MUST contain the individual message's
token and cost metrics — NOT accumulated totals. The dispatch MUST be fire-and-forget: errors
MUST be caught and silently discarded; a failed POST MUST NOT throw, reject, or crash the agent
session. The `postPhase` path MUST remain unmodified.

#### Scenario: Dispatches on assistant message with tokens

- GIVEN a `message.updated` event with role `assistant` and non-zero `tokens_input`
- WHEN the plugin handler fires
- THEN `POST /calls` is called with that message's individual token values and current `callIndex`

#### Scenario: Skips non-assistant messages

- GIVEN a `message.updated` event with role `user`
- WHEN the plugin handler fires
- THEN `POST /calls` is NOT called

#### Scenario: Skips messages without token data

- GIVEN a `message.updated` event with role `assistant` but no token fields
- WHEN the plugin handler fires
- THEN `POST /calls` is NOT called

#### Scenario: Network error does not crash session

- GIVEN the tracker server is unreachable
- WHEN `postCall` is dispatched
- THEN the error is silently discarded and the agent session continues normally

---

### Requirement: Build Integrity After Changes

The Go binary MUST compile successfully with `go build -o sdd-cost-tracker.exe .` after all
schema, handler, and storage additions. No existing HTTP handler or storage method signature
MUST be altered. The build MUST complete with zero errors.

#### Scenario: Binary compiles after adding calls support

- GIVEN all new files and modifications are applied
- WHEN `go build -o sdd-cost-tracker.exe .` is run
- THEN the build exits with code 0 and produces the executable
