package main

import (
	"context"
	"encoding/json"
	"math"
	"testing"
)

func TestMCPCostQueryReturnsMatchingFilters(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for _, r := range []PhaseRecord{
		baseRecord("proj-a", "feat-x", "design", "s1"),
		baseRecord("proj-a", "feat-x", "apply", "s2"),
		baseRecord("proj-b", "feat-x", "design", "s3"),
	} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	args := map[string]any{"project": "proj-a", "change_name": "feat-x", "phase": "apply"}
	res := callTool(t, store, "cost_query", args)

	rows := decodeToolText[[]PhaseRecord](t, res)
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Project != "proj-a" || rows[0].Phase != "apply" {
		t.Fatalf("unexpected row: %+v", rows[0])
	}
}

func TestMCPCostSummaryReturnsAggregatedTotals(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	a := baseRecord("proj-a", "feat-x", "design", "s1")
	a.CostUSD = 0.10
	b := baseRecord("proj-a", "feat-x", "apply", "s2")
	b.CostUSD = 0.20

	for _, r := range []PhaseRecord{a, b} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	res := callTool(t, store, "cost_summary", map[string]any{"project": "proj-a"})
	summary := decodeToolText[[]ChangeSummary](t, res)
	if len(summary) != 1 {
		t.Fatalf("expected 1 row, got %d", len(summary))
	}
	if math.Abs(summary[0].CostUSD-0.30) > 1e-9 {
		t.Fatalf("expected cost 0.30, got %v", summary[0].CostUSD)
	}
}

func TestMCPCostQueryWithoutProjectFiltersAcrossAllProjects(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for _, r := range []PhaseRecord{
		baseRecord("proj-a", "feat-x", "design", "s1"),
		baseRecord("proj-b", "feat-y", "apply", "s2"),
	} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	res := callTool(t, store, "cost_query", map[string]any{})
	rows := decodeToolText[[]PhaseRecord](t, res)
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows across all projects, got %d", len(rows))
	}
}

func TestMCPCostSummaryWithoutProjectAggregatesAcrossAllProjects(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	a := baseRecord("proj-a", "feat-x", "design", "s1")
	a.CostUSD = 0.10
	b := baseRecord("proj-b", "feat-y", "apply", "s2")
	b.CostUSD = 0.20

	for _, r := range []PhaseRecord{a, b} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	res := callTool(t, store, "cost_summary", map[string]any{})
	summary := decodeToolText[[]ChangeSummary](t, res)
	if len(summary) != 2 {
		t.Fatalf("expected 2 summary rows across all projects, got %d", len(summary))
	}
}

func TestMCPCostQueryNoMatchReturnsEmptyList(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	if err := store.UpsertPhase(ctx, baseRecord("proj-a", "feat-x", "design", "s1")); err != nil {
		t.Fatalf("seed data: %v", err)
	}

	res := callTool(t, store, "cost_query", map[string]any{"project": "ghost-project"})
	rows := decodeToolText[[]PhaseRecord](t, res)
	if len(rows) != 0 {
		t.Fatalf("expected empty result, got %d rows", len(rows))
	}
}

func TestMCPInitializeHandshake(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	resp, ok := HandleMCPRequest(context.Background(), store, jsonRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
	})
	if !ok {
		t.Fatalf("expected response for initialize")
	}
	if resp.Error != nil {
		t.Fatalf("expected no error, got %+v", resp.Error)
	}

	result, ok := resp.Result.(map[string]any)
	if !ok {
		t.Fatalf("expected map result, got %T", resp.Result)
	}
	if result["protocolVersion"] != "2024-11-05" {
		t.Fatalf("unexpected protocolVersion: %v", result["protocolVersion"])
	}
}

func callTool(t *testing.T, store *Store, name string, args map[string]any) jsonRPCResponse {
	t.Helper()

	argRaw, err := json.Marshal(args)
	if err != nil {
		t.Fatalf("marshal args: %v", err)
	}
	paramsRaw, err := json.Marshal(map[string]any{
		"name":      name,
		"arguments": json.RawMessage(argRaw),
	})
	if err != nil {
		t.Fatalf("marshal params: %v", err)
	}

	resp, ok := HandleMCPRequest(context.Background(), store, jsonRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params:  paramsRaw,
	})
	if !ok {
		t.Fatalf("expected response")
	}
	if resp.Error != nil {
		t.Fatalf("unexpected rpc error: %+v", resp.Error)
	}
	return resp
}

func decodeToolText[T any](t *testing.T, resp jsonRPCResponse) T {
	t.Helper()

	result, ok := resp.Result.(map[string]any)
	if !ok {
		t.Fatalf("expected result map, got %T", resp.Result)
	}
	content, ok := result["content"].([]map[string]any)
	if !ok || len(content) == 0 {
		t.Fatalf("expected content array, got %T", result["content"])
	}
	text, _ := content[0]["text"].(string)

	var out T
	if err := json.Unmarshal([]byte(text), &out); err != nil {
		t.Fatalf("decode tool text: %v", err)
	}
	return out
}
