 # Token Usage Capture Specification
 
 ## Purpose
 
 Intercepts outbound LLM API traffic via a local mitmproxy addon, extracts token usage fields from non-streaming and streaming (SSE) responses, and persists structured records to SQLite.
 
 ## Requirements
 
 ### Requirement: Proxy Interception
 
 The system MUST intercept HTTP and HTTPS outbound traffic via mitmproxy addon mode. The addon MUST hook the `response` event and filter flows by configurable LLM endpoint host patterns. Unmatched flows MUST pass through without modification or logging.
 
 #### Scenario: Non-streaming LLM response intercepted
 
 - GIVEN the proxy is running and a client app sends an HTTPS request to an LLM endpoint
 - WHEN the response Content-Type is `application/json` and body contains a `usage` object
 - THEN the addon extracts `prompt_tokens`, `completion_tokens`, `total_tokens` and persists a token event row
 - AND the original response is forwarded unmodified
 
 #### Scenario: Non-LLM traffic passes through
 
 - GIVEN the proxy is running
 - WHEN a request targets a host not matching any LLM endpoint pattern
 - THEN no token event is recorded and the flow is not modified
 
 ### Requirement: TLS Certificate Handling
 
 The system MUST support HTTPS interception via a mkcert-generated local CA installed once via `mkcert -install` into the Windows user trust store. The README MUST document this setup step.
 
 #### Scenario: TLS handshake succeeds after CA install
 
 - GIVEN `mkcert -install` has been run on the machine
 - WHEN the proxy intercepts an HTTPS connection to an LLM endpoint
 - THEN the TLS handshake completes without certificate errors
 
 ### Requirement: Streaming SSE Token Extraction
 
 The system MUST buffer all SSE data chunks per `flow.id`. On stream completion the addon MUST parse the final `data:` line as JSON and extract `usage`. If `usage` is absent the addon SHOULD log a warning and MUST NOT persist a partial record.
 
 #### Scenario: Streaming response — usage in final chunk
 
 - GIVEN a streaming SSE response is intercepted
 - WHEN the final `data:` chunk contains `"usage": {…}`
 - THEN one token event row is persisted with correct values
 
 #### Scenario: Streaming response — no usage in final chunk
 
 - GIVEN a streaming SSE response is intercepted
 - WHEN no `usage` key is present in the final chunk
 - THEN no token event row is persisted and a warning is logged
 
 ### Requirement: Agent Detection
 
 The system MUST derive the `agent` field from the `User-Agent` header. An optional `X-Agent-Name` header MUST take precedence when present. If neither yields a recognized label (`"opencode"`, `"claude-code"`, `"copilot-chat"`) the system MUST record `agent` as `"unknown"`.
 
 #### Scenario: Known agent via User-Agent
 
 - GIVEN a request with `User-Agent: opencode/…`
 - WHEN the response is processed
 - THEN the persisted row has `agent = "opencode"`
 
 #### Scenario: Unknown agent fallback
 
 - GIVEN an unrecognized `User-Agent` and no `X-Agent-Name` header
 - WHEN the response is processed
 - THEN the row has `agent = "unknown"` and `session_id = null`
 
 ### Requirement: SQLite Persistence
 
 The system MUST write each token event as a row in `token_events`. The database path MUST default to `~/.token-monitor/usage.db` and MUST be overridable via `--db-path`. A failed write MUST NOT crash the proxy; the addon MUST log the error and continue.
 
 #### Scenario: Row inserted on successful capture
 
 - GIVEN a valid token event is extracted
 - WHEN the storage layer writes to SQLite
 - THEN a row exists in `token_events` with all eight fields populated
 
 #### Scenario: DB write error does not crash proxy
 
 - GIVEN the SQLite file is inaccessible (permissions error)
 - WHEN the addon attempts to persist a token event
 - THEN the error is logged and the proxy continues intercepting flows
 
 ### Requirement: CLI Configuration
 
 The system MUST expose configuration via `argparse`. REQUIRED flags: `--port`, `--db-path`. OPTIONAL flag: `--filter-host` (repeatable host pattern).
 
 #### Scenario: Custom DB path accepted
 
 - GIVEN the user starts the proxy with `--db-path /tmp/test.db`
 - WHEN a token event is captured
 - THEN the row is written to `/tmp/test.db`, not the default path
 
 (End of file)
