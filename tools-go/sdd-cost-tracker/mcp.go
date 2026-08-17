package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type jsonRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type jsonRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      any           `json:"id,omitempty"`
	Result  any           `json:"result,omitempty"`
	Error   *jsonRPCError `json:"error,omitempty"`
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type mcpTool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

type mcpToolCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

type costQueryArgs struct {
	Project    string `json:"project"`
	ChangeName string `json:"change_name,omitempty"`
	Phase      string `json:"phase,omitempty"`
}

type costSummaryArgs struct {
	Project string `json:"project"`
}

func RunMCP(store *Store) error {
	return RunMCPWithIO(store, os.Stdin, os.Stdout)
}

func RunMCPWithIO(store *Store, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	encoder := json.NewEncoder(out)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var req jsonRPCRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			resp := jsonRPCResponse{
				JSONRPC: "2.0",
				Error:   &jsonRPCError{Code: -32700, Message: "parse error"},
			}
			if err := encoder.Encode(resp); err != nil {
				return fmt.Errorf("write parse error response: %w", err)
			}
			continue
		}

		resp, shouldRespond := HandleMCPRequest(context.Background(), store, req)
		if !shouldRespond {
			continue
		}

		if err := encoder.Encode(resp); err != nil {
			return fmt.Errorf("write json-rpc response: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read mcp input: %w", err)
	}

	return nil
}

func HandleMCPRequest(ctx context.Context, store *Store, req jsonRPCRequest) (jsonRPCResponse, bool) {
	if req.JSONRPC != "2.0" {
		return jsonRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error:   &jsonRPCError{Code: -32600, Message: "invalid request"},
		}, true
	}

	switch req.Method {
	case "initialize":
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Result: map[string]any{
			"protocolVersion": "2024-11-05",
			"serverInfo": map[string]any{
				"name":    "sdd-cost-tracker",
				"version": "0.1.0",
			},
			"capabilities": map[string]any{
				"tools": map[string]any{},
			},
		}}, true

	case "tools/list":
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Result: map[string]any{
			"tools": []mcpTool{
				{
					Name:        "cost_query",
					Description: "Query SDD phase cost and token data with optional filters",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"project":     map[string]any{"type": "string", "description": "Project name"},
							"change_name": map[string]any{"type": "string", "description": "Optional change name filter"},
							"phase":       map[string]any{"type": "string", "description": "Optional phase filter"},
						},
					},
				},
				{
					Name:        "cost_summary",
					Description: "Get aggregated cost and token totals per change",
					InputSchema: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"project": map[string]any{"type": "string", "description": "Optional project name filter"},
						},
					},
				},
			},
		}}, true

	case "tools/call":
		return handleMCPToolCall(ctx, store, req), true

	case "notifications/initialized":
		return jsonRPCResponse{}, false

	default:
		return jsonRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error:   &jsonRPCError{Code: -32601, Message: "method not found"},
		}, true
	}
}

func handleMCPToolCall(ctx context.Context, store *Store, req jsonRPCRequest) jsonRPCResponse {
	var params mcpToolCallParams
	if err := json.Unmarshal(req.Params, &params); err != nil {
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32602, Message: "invalid params"}}
	}

	switch params.Name {
	case "cost_query":
		var args costQueryArgs
		if err := json.Unmarshal(params.Arguments, &args); err != nil {
			return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32602, Message: "invalid params"}}
		}
		rows, err := store.QueryPhases(ctx, strings.TrimSpace(args.Project), strings.TrimSpace(args.ChangeName), strings.TrimSpace(args.Phase))
		if err != nil {
			return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32000, Message: "failed to query phases"}}
		}

		return toolTextResult(req.ID, rows)

	case "cost_summary":
		var args costSummaryArgs
		if err := json.Unmarshal(params.Arguments, &args); err != nil {
			return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32602, Message: "invalid params"}}
		}
		project := strings.TrimSpace(args.Project)
		var summary []ChangeSummary
		var err error
		if project == "" {
			summary, err = store.GetSummaryAllProjects(ctx)
		} else {
			summary, err = store.GetSummary(ctx, project)
		}
		if err != nil {
			return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32000, Message: "failed to get summary"}}
		}

		return toolTextResult(req.ID, summary)

	default:
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: &jsonRPCError{Code: -32601, Message: "tool not found"}}
	}
}

func toolTextResult(id any, payload any) jsonRPCResponse {
	raw, _ := json.Marshal(payload)
	return jsonRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result: map[string]any{
			"content": []map[string]any{{
				"type": "text",
				"text": string(raw),
			}},
		},
	}
}
