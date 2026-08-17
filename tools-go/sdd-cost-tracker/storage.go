package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

const createPhasesTableSQL = `
CREATE TABLE IF NOT EXISTS phases (
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
`

const createCallsTableSQL = `
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
`

type Store struct {
	db *sql.DB
}

type PhaseRecord struct {
	ID               int64   `json:"id"`
	Project          string  `json:"project"`
	ChangeName       string  `json:"change_name"`
	Phase            string  `json:"phase"`
	SessionID        string  `json:"session_id"`
	ModelID          *string `json:"model_id,omitempty"`
	ProviderID       *string `json:"provider_id,omitempty"`
	TokensInput      int64   `json:"tokens_input"`
	TokensOutput     int64   `json:"tokens_output"`
	TokensReasoning  int64   `json:"tokens_reasoning"`
	TokensCacheRead  int64   `json:"tokens_cache_read"`
	TokensCacheWrite int64   `json:"tokens_cache_write"`
	CostUSD          float64 `json:"cost_usd"`
	StartedAt        *int64  `json:"started_at,omitempty"`
	CompletedAt      *int64  `json:"completed_at,omitempty"`
}

type CallRecord struct {
	ID               int64   `json:"id"`
	SessionID        string  `json:"session_id"`
	CallIndex        int     `json:"call_index"`
	ModelID          *string `json:"model_id,omitempty"`
	ProviderID       *string `json:"provider_id,omitempty"`
	TokensInput      int64   `json:"tokens_input"`
	TokensOutput     int64   `json:"tokens_output"`
	TokensReasoning  int64   `json:"tokens_reasoning"`
	TokensCacheRead  int64   `json:"tokens_cache_read"`
	TokensCacheWrite int64   `json:"tokens_cache_write"`
	CostUSD          float64 `json:"cost_usd"`
	RecordedAt       *int64  `json:"recorded_at,omitempty"`
}

type ChangeSummary struct {
	Project          string  `json:"project"`
	ChangeName       string  `json:"change_name"`
	TokensInput      int64   `json:"tokens_input"`
	TokensOutput     int64   `json:"tokens_output"`
	TokensReasoning  int64   `json:"tokens_reasoning"`
	TokensCacheRead  int64   `json:"tokens_cache_read"`
	TokensCacheWrite int64   `json:"tokens_cache_write"`
	CostUSD          float64 `json:"cost_usd"`
}

func NewStore(dbPath string) (*Store, error) {
	if dbPath == "" {
		return nil, errors.New("dbPath is required")
	}

	if dbPath != ":memory:" {
		if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
			return nil, fmt.Errorf("create db directory: %w", err)
		}
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite db: %w", err)
	}

	db.SetMaxOpenConns(1)

	store := &Store{db: db}
	if err := store.init(context.Background()); err != nil {
		_ = db.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) init(ctx context.Context) error {
	if _, err := s.db.ExecContext(ctx, "PRAGMA journal_mode=WAL;"); err != nil {
		return fmt.Errorf("set journal_mode WAL: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, "PRAGMA busy_timeout=5000;"); err != nil {
		return fmt.Errorf("set busy_timeout: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, createPhasesTableSQL); err != nil {
		return fmt.Errorf("create phases table: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, createCallsTableSQL); err != nil {
		return fmt.Errorf("create calls table: %w", err)
	}

	return nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Health(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *Store) UpsertPhase(ctx context.Context, record PhaseRecord) error {
	const q = `
INSERT INTO phases (
  project,
  change_name,
  phase,
  session_id,
  model_id,
  provider_id,
  tokens_input,
  tokens_output,
  tokens_reasoning,
  tokens_cache_read,
  tokens_cache_write,
  cost_usd,
  started_at,
  completed_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(session_id) DO UPDATE SET
  project = excluded.project,
  change_name = excluded.change_name,
  phase = excluded.phase,
  model_id = excluded.model_id,
  provider_id = excluded.provider_id,
  tokens_input = excluded.tokens_input,
  tokens_output = excluded.tokens_output,
  tokens_reasoning = excluded.tokens_reasoning,
  tokens_cache_read = excluded.tokens_cache_read,
  tokens_cache_write = excluded.tokens_cache_write,
  cost_usd = excluded.cost_usd,
  started_at = excluded.started_at,
  completed_at = excluded.completed_at;
`

	_, err := s.db.ExecContext(
		ctx,
		q,
		record.Project,
		record.ChangeName,
		record.Phase,
		record.SessionID,
		record.ModelID,
		record.ProviderID,
		record.TokensInput,
		record.TokensOutput,
		record.TokensReasoning,
		record.TokensCacheRead,
		record.TokensCacheWrite,
		record.CostUSD,
		record.StartedAt,
		record.CompletedAt,
	)
	if err != nil {
		return fmt.Errorf("upsert phase: %w", err)
	}

	return nil
}

func (s *Store) ListChanges(ctx context.Context, project string) ([]string, error) {
	const q = `
SELECT DISTINCT change_name
FROM phases
WHERE project = ?
ORDER BY change_name ASC;
`

	rows, err := s.db.QueryContext(ctx, q, project)
	if err != nil {
		return nil, fmt.Errorf("list changes: %w", err)
	}
	defer rows.Close()

	changes := make([]string, 0)
	for rows.Next() {
		var changeName string
		if err := rows.Scan(&changeName); err != nil {
			return nil, fmt.Errorf("scan change name: %w", err)
		}
		changes = append(changes, changeName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate changes: %w", err)
	}

	return changes, nil
}

func (s *Store) GetChangePhases(ctx context.Context, project, changeName string) ([]PhaseRecord, error) {
	const q = `
SELECT
  id,
  project,
  change_name,
  phase,
  session_id,
  model_id,
  provider_id,
  tokens_input,
  tokens_output,
  tokens_reasoning,
  tokens_cache_read,
  tokens_cache_write,
  cost_usd,
  started_at,
  completed_at
FROM phases
WHERE project = ? AND change_name = ?
ORDER BY id ASC;
`

	rows, err := s.db.QueryContext(ctx, q, project, changeName)
	if err != nil {
		return nil, fmt.Errorf("get change phases: %w", err)
	}
	defer rows.Close()

	return scanPhaseRecords(rows)
}

func (s *Store) QueryPhases(ctx context.Context, project, changeName, phase string) ([]PhaseRecord, error) {
	const q = `
SELECT
  id,
  project,
  change_name,
  phase,
  session_id,
  model_id,
  provider_id,
  tokens_input,
  tokens_output,
  tokens_reasoning,
  tokens_cache_read,
  tokens_cache_write,
  cost_usd,
  started_at,
  completed_at
FROM phases
WHERE (? = '' OR project = ?)
  AND (? = '' OR change_name = ?)
  AND (? = '' OR phase = ?)
ORDER BY id ASC;
`

	rows, err := s.db.QueryContext(ctx, q, project, project, changeName, changeName, phase, phase)
	if err != nil {
		return nil, fmt.Errorf("query phases: %w", err)
	}
	defer rows.Close()

	return scanPhaseRecords(rows)
}

func (s *Store) InsertCall(ctx context.Context, r CallRecord) error {
	if strings.TrimSpace(r.SessionID) == "" {
		return errors.New("session_id is required")
	}

	const q = `
INSERT INTO calls (
  session_id,
  call_index,
  model_id,
  provider_id,
  tokens_input,
  tokens_output,
  tokens_reasoning,
  tokens_cache_read,
  tokens_cache_write,
  cost_usd,
  recorded_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
`

	_, err := s.db.ExecContext(
		ctx,
		q,
		r.SessionID,
		r.CallIndex,
		r.ModelID,
		r.ProviderID,
		r.TokensInput,
		r.TokensOutput,
		r.TokensReasoning,
		r.TokensCacheRead,
		r.TokensCacheWrite,
		r.CostUSD,
		r.RecordedAt,
	)
	if err != nil {
		return fmt.Errorf("insert call: %w", err)
	}

	return nil
}

func (s *Store) GetCallsBySession(ctx context.Context, sessionID string) ([]CallRecord, error) {
	const q = `
SELECT
  id,
  session_id,
  call_index,
  model_id,
  provider_id,
  tokens_input,
  tokens_output,
  tokens_reasoning,
  tokens_cache_read,
  tokens_cache_write,
  cost_usd,
  recorded_at
FROM calls
WHERE session_id = ?
ORDER BY call_index ASC, id ASC;
`

	rows, err := s.db.QueryContext(ctx, q, sessionID)
	if err != nil {
		return nil, fmt.Errorf("get calls by session: %w", err)
	}
	defer rows.Close()

	calls := make([]CallRecord, 0)
	for rows.Next() {
		var r CallRecord
		if err := rows.Scan(
			&r.ID,
			&r.SessionID,
			&r.CallIndex,
			&r.ModelID,
			&r.ProviderID,
			&r.TokensInput,
			&r.TokensOutput,
			&r.TokensReasoning,
			&r.TokensCacheRead,
			&r.TokensCacheWrite,
			&r.CostUSD,
			&r.RecordedAt,
		); err != nil {
			return nil, fmt.Errorf("scan call record: %w", err)
		}
		calls = append(calls, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate call rows: %w", err)
	}

	return calls, nil
}

func (s *Store) GetSummary(ctx context.Context, project string) ([]ChangeSummary, error) {
	const q = `
SELECT
  project,
  change_name,
  COALESCE(SUM(tokens_input), 0) AS tokens_input,
  COALESCE(SUM(tokens_output), 0) AS tokens_output,
  COALESCE(SUM(tokens_reasoning), 0) AS tokens_reasoning,
  COALESCE(SUM(tokens_cache_read), 0) AS tokens_cache_read,
  COALESCE(SUM(tokens_cache_write), 0) AS tokens_cache_write,
  COALESCE(SUM(cost_usd), 0) AS cost_usd
FROM phases
WHERE project = ?
GROUP BY project, change_name
ORDER BY change_name ASC;
`

	rows, err := s.db.QueryContext(ctx, q, project)
	if err != nil {
		return nil, fmt.Errorf("get summary: %w", err)
	}
	defer rows.Close()

	summary := make([]ChangeSummary, 0)
	for rows.Next() {
		var item ChangeSummary
		if err := rows.Scan(
			&item.Project,
			&item.ChangeName,
			&item.TokensInput,
			&item.TokensOutput,
			&item.TokensReasoning,
			&item.TokensCacheRead,
			&item.TokensCacheWrite,
			&item.CostUSD,
		); err != nil {
			return nil, fmt.Errorf("scan summary: %w", err)
		}
		summary = append(summary, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate summary rows: %w", err)
	}

	return summary, nil
}

func (s *Store) GetSummaryAllProjects(ctx context.Context) ([]ChangeSummary, error) {
	const q = `
SELECT
  project,
  change_name,
  COALESCE(SUM(tokens_input), 0) AS tokens_input,
  COALESCE(SUM(tokens_output), 0) AS tokens_output,
  COALESCE(SUM(tokens_reasoning), 0) AS tokens_reasoning,
  COALESCE(SUM(tokens_cache_read), 0) AS tokens_cache_read,
  COALESCE(SUM(tokens_cache_write), 0) AS tokens_cache_write,
  COALESCE(SUM(cost_usd), 0) AS cost_usd
FROM phases
GROUP BY project, change_name
ORDER BY project ASC, change_name ASC;
`

	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("get summary all projects: %w", err)
	}
	defer rows.Close()

	summary := make([]ChangeSummary, 0)
	for rows.Next() {
		var item ChangeSummary
		if err := rows.Scan(
			&item.Project,
			&item.ChangeName,
			&item.TokensInput,
			&item.TokensOutput,
			&item.TokensReasoning,
			&item.TokensCacheRead,
			&item.TokensCacheWrite,
			&item.CostUSD,
		); err != nil {
			return nil, fmt.Errorf("scan summary: %w", err)
		}
		summary = append(summary, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate summary rows: %w", err)
	}

	return summary, nil
}

func scanPhaseRecords(rows *sql.Rows) ([]PhaseRecord, error) {
	records := make([]PhaseRecord, 0)

	for rows.Next() {
		var record PhaseRecord
		if err := rows.Scan(
			&record.ID,
			&record.Project,
			&record.ChangeName,
			&record.Phase,
			&record.SessionID,
			&record.ModelID,
			&record.ProviderID,
			&record.TokensInput,
			&record.TokensOutput,
			&record.TokensReasoning,
			&record.TokensCacheRead,
			&record.TokensCacheWrite,
			&record.CostUSD,
			&record.StartedAt,
			&record.CompletedAt,
		); err != nil {
			return nil, fmt.Errorf("scan phase record: %w", err)
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate phase rows: %w", err)
	}

	return records, nil
}
