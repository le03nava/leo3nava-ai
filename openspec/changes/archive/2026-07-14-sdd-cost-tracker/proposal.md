# Proposal: SDD Cost Tracker

## Intent

SDD phases consume tokens across multiple models and providers, but there is no structured way to query cost and token breakdown per phase or per change. This change builds a Go binary MCP server that records and exposes token consumption and USD cost, grouped by project, change name, and phase — becoming the canonical cost-tracking tool for this repo.

## Scope

### In Scope
- Go binary MCP server (`tools-go/sdd-cost-tracker/`)
- Embedded SQLite via `modernc.org/sqlite` (no CGO)
- HTTP server on configurable port (default 7438)
- MCP tools: `cost_query`, `cost_summary`
- HTTP endpoints: `POST /phases`, `GET /health`, `GET /changes`, `GET /changes/:name`, `GET /summary`
- Exact SQLite schema as specified in the change request
- DB stored at `~/.sdd-cost-tracker/db.sqlite`
- Co-located tests and README with build/run instructions

### Out of Scope
- Migration from `tools-py/token-monitor` (deprecated separately)
- Real-time streaming or WebSocket push
- Multi-user / multi-machine sync
- Authentication or access control on HTTP endpoints
- Dashboard UI

## Capabilities

### New Capabilities
- `sdd-cost-tracking`: Records and queries token consumption and USD cost per SDD phase, grouped by project and change name, exposed via MCP tools and HTTP API.
