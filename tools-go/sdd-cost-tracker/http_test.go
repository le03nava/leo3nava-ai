package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPPostPhasesInsertReturnsCreated(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	body := map[string]any{
		"project":     "proj-a",
		"change_name": "feat-x",
		"phase":       "apply",
		"session_id":  "s-1",
		"cost_usd":    0.12,
	}

	rec := performJSONRequest(t, router, http.MethodPost, "/phases", body)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	assertContentTypeJSON(t, rec)
}

func TestHTTPPostPhasesUpsertReturnsOK(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	body := map[string]any{
		"project":     "proj-a",
		"change_name": "feat-x",
		"phase":       "apply",
		"session_id":  "s-1",
		"cost_usd":    0.10,
	}

	first := performJSONRequest(t, router, http.MethodPost, "/phases", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("expected first status 201, got %d", first.Code)
	}

	body["cost_usd"] = 0.25
	second := performJSONRequest(t, router, http.MethodPost, "/phases", body)
	if second.Code != http.StatusOK {
		t.Fatalf("expected second status 200, got %d", second.Code)
	}

	rows, err := store.GetChangePhases(ctx, "proj-a", "feat-x")
	if err != nil {
		t.Fatalf("get change phases: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].CostUSD != 0.25 {
		t.Fatalf("expected cost 0.25, got %v", rows[0].CostUSD)
	}
}

func TestHTTPPostPhasesMissingProjectReturns400(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	body := map[string]any{
		"change_name": "feat-x",
		"phase":       "apply",
		"session_id":  "s-1",
	}

	rec := performJSONRequest(t, router, http.MethodPost, "/phases", body)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	assertErrorPayload(t, rec)
}

func TestHTTPPostPhasesMissingRequiredFieldsReturns400(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	cases := []map[string]any{
		{"project": "proj-a", "phase": "apply", "session_id": "s-1"},
		{"project": "proj-a", "change_name": "feat-x", "session_id": "s-1"},
		{"project": "proj-a", "change_name": "feat-x", "phase": "apply"},
	}

	for i, body := range cases {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			rec := performJSONRequest(t, router, http.MethodPost, "/phases", body)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d", rec.Code)
			}
			assertErrorPayload(t, rec)
		})
	}
}

func TestHTTPPostCallsInsertReturnsCreated(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	body := map[string]any{
		"session_id":         "s-calls-1",
		"call_index":         0,
		"model_id":           "gpt-5",
		"provider_id":        "openai",
		"tokens_input":       12,
		"tokens_output":      34,
		"tokens_reasoning":   5,
		"tokens_cache_read":  2,
		"tokens_cache_write": 1,
		"cost_usd":           0.01,
	}

	rec := performJSONRequest(t, router, http.MethodPost, "/calls", body)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
	assertContentTypeJSON(t, rec)
}

func TestHTTPPostCallsMissingSessionIDReturns400(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	body := map[string]any{
		"call_index": 1,
	}

	rec := performJSONRequest(t, router, http.MethodPost, "/calls", body)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	assertErrorPayload(t, rec)
}

func TestHTTPPostCallsMalformedJSONReturns400(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodPost, "/calls", bytes.NewReader([]byte("not-json")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	assertErrorPayload(t, rec)
}

func TestHTTPGetCallsValidSession(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	if err := store.InsertCall(ctx, CallRecord{SessionID: "s-get-1", CallIndex: 0, CostUSD: 0.03}); err != nil {
		t.Fatalf("seed call: %v", err)
	}

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/calls?session_id=s-get-1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Calls []CallRecord `json:"calls"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(payload.Calls))
	}
}

func TestHTTPGetCallsUnknownSession(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/calls?session_id=ghost", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Calls []CallRecord `json:"calls"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Calls) != 0 {
		t.Fatalf("expected 0 calls, got %d", len(payload.Calls))
	}
}

func TestHTTPGetCallsMissingSessionIDReturns400(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/calls", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	assertErrorPayload(t, rec)
}

func TestHTTPHealth(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	assertContentTypeJSON(t, rec)

	var payload map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected status=ok, got %q", payload["status"])
	}
}

func TestHTTPListChanges(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for _, r := range []PhaseRecord{
		baseRecord("proj-a", "feat-x", "design", "s1"),
		baseRecord("proj-a", "feat-y", "apply", "s2"),
		baseRecord("proj-b", "other", "verify", "s3"),
	} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/changes?project=proj-a", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Changes []string `json:"changes"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Changes) != 2 {
		t.Fatalf("expected 2 changes, got %d", len(payload.Changes))
	}
}

func TestHTTPListChangesMissingProject(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/changes", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestHTTPGetChangePhases(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	for _, r := range []PhaseRecord{
		baseRecord("proj-a", "feat-x", "design", "s1"),
		baseRecord("proj-a", "feat-x", "apply", "s2"),
		baseRecord("proj-a", "feat-x", "verify", "s3"),
	} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/changes/feat-x?project=proj-a", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Phases []PhaseRecord `json:"phases"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Phases) != 3 {
		t.Fatalf("expected 3 phases, got %d", len(payload.Phases))
	}
}

func TestHTTPGetChangePhasesNoMatchReturnsEmptyArray(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/changes/unknown?project=proj-a", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Phases []PhaseRecord `json:"phases"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Phases) != 0 {
		t.Fatalf("expected empty phases, got %d", len(payload.Phases))
	}
}

func TestHTTPSummary(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)
	defer store.Close()

	a := baseRecord("proj-a", "feat-x", "design", "s1")
	a.CostUSD = 0.10
	b := baseRecord("proj-a", "feat-x", "apply", "s2")
	b.CostUSD = 0.15

	for _, r := range []PhaseRecord{a, b} {
		if err := store.UpsertPhase(ctx, r); err != nil {
			t.Fatalf("seed data: %v", err)
		}
	}

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/summary?project=proj-a", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var payload struct {
		Summary []ChangeSummary `json:"summary"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Summary) != 1 {
		t.Fatalf("expected 1 summary row, got %d", len(payload.Summary))
	}
	if payload.Summary[0].CostUSD != 0.25 {
		t.Fatalf("expected aggregated cost 0.25, got %v", payload.Summary[0].CostUSD)
	}
}

func TestHTTPSummaryMissingProject(t *testing.T) {
	store := newTestStore(t)
	defer store.Close()

	router := NewRouter(store)
	req := httptest.NewRequest(http.MethodGet, "/summary", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func performJSONRequest(t *testing.T, handler http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal body: %v", err)
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func assertErrorPayload(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()
	assertContentTypeJSON(t, rec)

	var payload map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode error payload: %v", err)
	}
	if payload["error"] == "" {
		t.Fatalf("expected non-empty error message")
	}
}

func assertContentTypeJSON(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", got)
	}
}
