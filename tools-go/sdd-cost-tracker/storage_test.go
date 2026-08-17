package main

import (
	"context"
	"testing"
)

func TestSchemaInit(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	const q = `PRAGMA table_info(phases);`
	rows, err := store.db.QueryContext(context.Background(), q)
	if err != nil {
		t.Fatalf("query table info: %v", err)
	}
	defer rows.Close()

	columns := map[string]bool{}
	for rows.Next() {
		var (
			cid       int
			name      string
			colType   string
			notNull   int
			defaultV  interface{}
			primaryPK int
		)
		if err := rows.Scan(&cid, &name, &colType, &notNull, &defaultV, &primaryPK); err != nil {
			t.Fatalf("scan pragma row: %v", err)
		}
		columns[name] = true
	}

	for _, col := range []string{
		"id", "project", "change_name", "phase", "session_id", "model_id", "provider_id",
		"tokens_input", "tokens_output", "tokens_reasoning", "tokens_cache_read", "tokens_cache_write",
		"cost_usd", "started_at", "completed_at",
	} {
		if !columns[col] {
			t.Fatalf("expected column %q to exist", col)
		}
	}

	const callsQ = `PRAGMA table_info(calls);`
	callsRows, err := store.db.QueryContext(context.Background(), callsQ)
	if err != nil {
		t.Fatalf("query calls table info: %v", err)
	}
	defer callsRows.Close()

	callsColumns := map[string]bool{}
	for callsRows.Next() {
		var (
			cid       int
			name      string
			colType   string
			notNull   int
			defaultV  interface{}
			primaryPK int
		)
		if err := callsRows.Scan(&cid, &name, &colType, &notNull, &defaultV, &primaryPK); err != nil {
			t.Fatalf("scan calls pragma row: %v", err)
		}
		callsColumns[name] = true
	}

	for _, col := range []string{
		"id", "session_id", "call_index", "model_id", "provider_id",
		"tokens_input", "tokens_output", "tokens_reasoning", "tokens_cache_read", "tokens_cache_write",
		"cost_usd", "recorded_at",
	} {
		if !callsColumns[col] {
			t.Fatalf("expected calls column %q to exist", col)
		}
	}
}

func TestInsertCallHappyPath(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	r := baseCallRecord("s1", 0)
	r.CostUSD = 0.12
	r.TokensInput = 100
	r.TokensOutput = 50

	if err := store.InsertCall(ctx, r); err != nil {
		t.Fatalf("insert call: %v", err)
	}

	count := rowCount(t, store, "SELECT COUNT(*) FROM calls")
	if count != 1 {
		t.Fatalf("expected 1 call row, got %d", count)
	}
}

func TestGetCallsBySessionHappyPath(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	r := baseCallRecord("s1", 0)
	r.CostUSD = 0.21
	r.TokensInput = 11
	r.TokensOutput = 7
	r.TokensReasoning = 3
	r.TokensCacheRead = 2
	r.TokensCacheWrite = 1

	if err := store.InsertCall(ctx, r); err != nil {
		t.Fatalf("insert call: %v", err)
	}

	rows, err := store.GetCallsBySession(ctx, "s1")
	if err != nil {
		t.Fatalf("get calls by session: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 call row, got %d", len(rows))
	}

	got := rows[0]
	if got.SessionID != "s1" || got.CallIndex != 0 || got.CostUSD != 0.21 ||
		got.TokensInput != 11 || got.TokensOutput != 7 || got.TokensReasoning != 3 ||
		got.TokensCacheRead != 2 || got.TokensCacheWrite != 1 {
		t.Fatalf("unexpected call row: %+v", got)
	}
}

func TestGetCallsBySessionEmpty(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	rows, err := store.GetCallsBySession(ctx, "ghost")
	if err != nil {
		t.Fatalf("get calls by session: %v", err)
	}
	if len(rows) != 0 {
		t.Fatalf("expected empty call rows, got %d", len(rows))
	}
}

func TestGetCallsBySessionMultiple(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for i := 0; i < 3; i++ {
		r := baseCallRecord("s1", i)
		if err := store.InsertCall(ctx, r); err != nil {
			t.Fatalf("insert call %d: %v", i, err)
		}
	}

	rows, err := store.GetCallsBySession(ctx, "s1")
	if err != nil {
		t.Fatalf("get calls by session: %v", err)
	}
	if len(rows) != 3 {
		t.Fatalf("expected 3 call rows, got %d", len(rows))
	}
}

func TestGetCallsBySessionOrdering(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for _, i := range []int{2, 0, 1} {
		r := baseCallRecord("s1", i)
		if err := store.InsertCall(ctx, r); err != nil {
			t.Fatalf("insert call %d: %v", i, err)
		}
	}

	rows, err := store.GetCallsBySession(ctx, "s1")
	if err != nil {
		t.Fatalf("get calls by session: %v", err)
	}
	if len(rows) != 3 {
		t.Fatalf("expected 3 call rows, got %d", len(rows))
	}

	if rows[0].CallIndex != 0 || rows[1].CallIndex != 1 || rows[2].CallIndex != 2 {
		t.Fatalf("expected ordered call indices [0,1,2], got [%d,%d,%d]",
			rows[0].CallIndex, rows[1].CallIndex, rows[2].CallIndex)
	}
}

func TestUpsertInsert(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	rec := baseRecord("proj-a", "change-1", "design", "s1")
	if err := store.UpsertPhase(ctx, rec); err != nil {
		t.Fatalf("upsert insert: %v", err)
	}

	count := rowCount(t, store, "SELECT COUNT(*) FROM phases")
	if count != 1 {
		t.Fatalf("expected 1 row, got %d", count)
	}
}

func TestUpsertConflict(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	rec := baseRecord("proj-a", "change-1", "design", "s1")
	rec.CostUSD = 0.10
	if err := store.UpsertPhase(ctx, rec); err != nil {
		t.Fatalf("first upsert: %v", err)
	}

	rec.CostUSD = 0.35
	rec.TokensInput = 999
	if err := store.UpsertPhase(ctx, rec); err != nil {
		t.Fatalf("second upsert: %v", err)
	}

	count := rowCount(t, store, "SELECT COUNT(*) FROM phases WHERE session_id = 's1'")
	if count != 1 {
		t.Fatalf("expected 1 row after conflict update, got %d", count)
	}

	var cost float64
	var input int64
	if err := store.db.QueryRowContext(ctx, "SELECT cost_usd, tokens_input FROM phases WHERE session_id = ?", "s1").Scan(&cost, &input); err != nil {
		t.Fatalf("query updated row: %v", err)
	}
	if cost != 0.35 {
		t.Fatalf("expected cost 0.35, got %v", cost)
	}
	if input != 999 {
		t.Fatalf("expected tokens_input 999, got %d", input)
	}
}

func TestListChanges(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	seed := []PhaseRecord{
		baseRecord("proj-a", "change-1", "design", "s1"),
		baseRecord("proj-a", "change-2", "apply", "s2"),
		baseRecord("proj-a", "change-1", "verify", "s3"),
		baseRecord("proj-b", "change-x", "apply", "s4"),
	}

	for _, r := range seed {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed upsert: %v", err)
		}
	}

	changes, err := store.ListChanges(ctx, "proj-a")
	if err != nil {
		t.Fatalf("list changes: %v", err)
	}

	if len(changes) != 2 {
		t.Fatalf("expected 2 distinct changes, got %d (%v)", len(changes), changes)
	}
	if changes[0] != "change-1" || changes[1] != "change-2" {
		t.Fatalf("unexpected change list order/content: %v", changes)
	}
}

func TestGetChangePhases(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	seed := []PhaseRecord{
		baseRecord("proj-a", "feat-x", "design", "s1"),
		baseRecord("proj-a", "feat-x", "apply", "s2"),
		baseRecord("proj-a", "feat-x", "verify", "s3"),
		baseRecord("proj-a", "feat-y", "design", "s4"),
		baseRecord("proj-b", "feat-x", "design", "s5"),
	}

	for _, r := range seed {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed upsert: %v", err)
		}
	}

	phases, err := store.GetChangePhases(ctx, "proj-a", "feat-x")
	if err != nil {
		t.Fatalf("get change phases: %v", err)
	}

	if len(phases) != 3 {
		t.Fatalf("expected 3 rows, got %d", len(phases))
	}
}

func TestGetSummaryAggregation(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	a := baseRecord("proj-a", "feat-x", "design", "s1")
	a.TokensInput = 100
	a.TokensOutput = 50
	a.TokensReasoning = 20
	a.TokensCacheRead = 5
	a.TokensCacheWrite = 2
	a.CostUSD = 0.10

	b := baseRecord("proj-a", "feat-x", "apply", "s2")
	b.TokensInput = 200
	b.TokensOutput = 80
	b.TokensReasoning = 30
	b.TokensCacheRead = 8
	b.TokensCacheWrite = 3
	b.CostUSD = 0.10

	c := baseRecord("proj-a", "feat-y", "design", "s3")
	c.TokensInput = 10
	c.TokensOutput = 10
	c.CostUSD = 0.05

	for _, r := range []PhaseRecord{a, b, c} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed upsert: %v", err)
		}
	}

	summary, err := store.GetSummary(ctx, "proj-a")
	if err != nil {
		t.Fatalf("get summary: %v", err)
	}

	if len(summary) != 2 {
		t.Fatalf("expected 2 summary rows, got %d", len(summary))
	}

	x := summary[0]
	if x.ChangeName != "feat-x" {
		t.Fatalf("expected first summary row feat-x, got %s", x.ChangeName)
	}
	if x.CostUSD != 0.20 || x.TokensInput != 300 || x.TokensOutput != 130 || x.TokensReasoning != 50 || x.TokensCacheRead != 13 || x.TokensCacheWrite != 5 {
		t.Fatalf("unexpected aggregated values for feat-x: %+v", x)
	}
}

func TestGetSummaryEmpty(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	summary, err := store.GetSummary(ctx, "ghost-project")
	if err != nil {
		t.Fatalf("get summary empty: %v", err)
	}
	if len(summary) != 0 {
		t.Fatalf("expected empty summary, got %d rows", len(summary))
	}
}

func TestHealth(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	if err := store.Health(ctx); err != nil {
		t.Fatalf("expected health to pass, got error: %v", err)
	}
}

func newTestStore(t *testing.T) *Store {
	t.Helper()
	store, err := NewStore(":memory:")
	if err != nil {
		t.Fatalf("new store: %v", err)
	}
	return store
}

func rowCount(t *testing.T, store *Store, q string) int {
	t.Helper()
	var count int
	if err := store.db.QueryRowContext(context.Background(), q).Scan(&count); err != nil {
		t.Fatalf("count query failed: %v", err)
	}
	return count
}

func baseRecord(project, changeName, phase, sessionID string) PhaseRecord {
	return PhaseRecord{
		Project:    project,
		ChangeName: changeName,
		Phase:      phase,
		SessionID:  sessionID,
	}
}

func baseCallRecord(sessionID string, callIndex int) CallRecord {
	return CallRecord{
		SessionID: sessionID,
		CallIndex: callIndex,
	}
}
