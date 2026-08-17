package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type upsertPhaseResponse struct {
	OK bool `json:"ok"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type listChangesResponse struct {
	Changes []string `json:"changes"`
}

type getChangePhasesResponse struct {
	Phases []PhaseRecord `json:"phases"`
}

type summaryResponse struct {
	Summary []ChangeSummary `json:"summary"`
}

type getCallsResponse struct {
	Calls []CallRecord `json:"calls"`
}

func NewRouter(store *Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth(store))
	mux.HandleFunc("/phases", handlePostPhases(store))
	mux.HandleFunc("/changes", handleListChanges(store))
	mux.HandleFunc("/changes/", handleGetChangePhases(store))
	mux.HandleFunc("/summary", handleSummary(store))
	mux.HandleFunc("/calls", handleCalls(store))
	return mux
}

func handleCalls(store *Store) http.HandlerFunc {
	post := handlePostCalls(store)
	get := handleGetCalls(store)

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			post(w, r)
		case http.MethodGet:
			get(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func handleHealth(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		if err := store.Health(r.Context()); err != nil {
			writeError(w, http.StatusInternalServerError, "health check failed")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func handlePostPhases(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		var record PhaseRecord
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&record); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON body")
			return
		}

		if err := requireNonEmpty("project", record.Project); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := requireNonEmpty("change_name", record.ChangeName); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := requireNonEmpty("phase", record.Phase); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := requireNonEmpty("session_id", record.SessionID); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		existed, err := sessionExists(r.Context(), store, record.SessionID)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to query existing phase")
			return
		}

		if err := store.UpsertPhase(r.Context(), record); err != nil {
			writeError(w, http.StatusInternalServerError, "failed to upsert phase")
			return
		}

		status := http.StatusCreated
		if existed {
			status = http.StatusOK
		}

		writeJSON(w, status, upsertPhaseResponse{OK: true})
	}
}

func handleListChanges(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		project := strings.TrimSpace(r.URL.Query().Get("project"))
		if project == "" {
			writeError(w, http.StatusBadRequest, "project is required")
			return
		}

		changes, err := store.ListChanges(r.Context(), project)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to list changes")
			return
		}

		writeJSON(w, http.StatusOK, listChangesResponse{Changes: changes})
	}
}

func handleGetChangePhases(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		project := strings.TrimSpace(r.URL.Query().Get("project"))
		if project == "" {
			writeError(w, http.StatusBadRequest, "project is required")
			return
		}

		prefix := "/changes/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			writeError(w, http.StatusNotFound, "not found")
			return
		}

		changeName := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, prefix))
		if changeName == "" {
			writeError(w, http.StatusBadRequest, "change name is required")
			return
		}

		phases, err := store.GetChangePhases(r.Context(), project, changeName)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to get change phases")
			return
		}

		writeJSON(w, http.StatusOK, getChangePhasesResponse{Phases: phases})
	}
}

func handleSummary(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		project := strings.TrimSpace(r.URL.Query().Get("project"))
		if project == "" {
			writeError(w, http.StatusBadRequest, "project is required")
			return
		}

		summary, err := store.GetSummary(r.Context(), project)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to get summary")
			return
		}

		writeJSON(w, http.StatusOK, summaryResponse{Summary: summary})
	}
}

func handlePostCalls(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record CallRecord
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&record); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON body")
			return
		}

		if err := requireNonEmpty("session_id", record.SessionID); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := store.InsertCall(r.Context(), record); err != nil {
			writeError(w, http.StatusInternalServerError, "failed to insert call")
			return
		}

		writeJSON(w, http.StatusCreated, upsertPhaseResponse{OK: true})
	}
}

func handleGetCalls(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := strings.TrimSpace(r.URL.Query().Get("session_id"))
		if sessionID == "" {
			writeError(w, http.StatusBadRequest, "session_id is required")
			return
		}

		calls, err := store.GetCallsBySession(r.Context(), sessionID)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to get calls")
			return
		}

		writeJSON(w, http.StatusOK, getCallsResponse{Calls: calls})
	}
}

func requireNonEmpty(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func sessionExists(ctx context.Context, store *Store, sessionID string) (bool, error) {
	const q = `SELECT 1 FROM phases WHERE session_id = ? LIMIT 1;`

	var marker int
	err := store.db.QueryRowContext(ctx, q, sessionID).Scan(&marker)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return false, err
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}
