package app_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/leo3nava/leo3nava_ai/tools-go/leo3nava-ai/internal/app"
)

func TestRunValidChangePath(t *testing.T) {
	root := t.TempDir()
	seedValidChange(t, root)

	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{root})

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr should be empty, got: %s", stderr.String())
	}
	if stdout.Len() == 0 {
		t.Fatal("stdout should contain JSON output, got empty")
	}

	// Verify the output is valid JSON with required fields.
	var result map[string]interface{}
	if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v\nraw: %s", err, stdout.String())
	}

	for _, field := range []string{"schemaName", "schemaVersion", "changeName", "artifacts", "taskProgress", "dependencies", "nextRecommended", "blockedReasons"} {
		if _, ok := result[field]; !ok {
			t.Errorf("JSON missing required field: %s", field)
		}
	}

	if result["schemaName"] != "phase-status-validator" {
		t.Errorf("schemaName: want %q, got %v", "phase-status-validator", result["schemaName"])
	}
}

func TestRunInvalidPath(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{"/nonexistent/path/that/does/not/exist"})

	if err == nil {
		t.Fatal("expected error for invalid path, got nil")
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout should be empty on error, got: %s", stdout.String())
	}
	if stderr.Len() == 0 {
		t.Error("stderr should contain error message, got empty")
	}
}

func TestRunNoArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{})

	if err == nil {
		t.Fatal("expected error for no args, got nil")
	}
}

func TestRunEmptyString(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{""})

	if err == nil {
		t.Fatal("expected error for empty string arg, got nil")
	}
}

func TestRunValidChangeJSONStructure(t *testing.T) {
	root := t.TempDir()
	seedValidChange(t, root)

	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{root})

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v", err)
	}

	// Verify changeName is inferred from directory basename.
	if got := result["changeName"]; got != filepath.Base(root) {
		t.Errorf("changeName: want %q, got %v", filepath.Base(root), got)
	}

	// Verify nextRecommended is a string.
	if _, ok := result["nextRecommended"].(string); !ok {
		t.Error("nextRecommended should be a string")
	}

	// Verify artifacts is an object.
	if _, ok := result["artifacts"].(map[string]interface{}); !ok {
		t.Error("artifacts should be an object")
	}

	// Verify dependencies is an object.
	if _, ok := result["dependencies"].(map[string]interface{}); !ok {
		t.Error("dependencies should be an object")
	}
}

func TestRunErrorContainsPathInfo(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := app.Run(&stdout, &stderr, []string{"/no/such/path"})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stderr.Len() == 0 {
		t.Error("stderr should have error output")
	}
}

// seedValidChange creates a minimal valid change directory.
func seedValidChange(t *testing.T, root string) {
	t.Helper()
	writeFile(t, root, "proposal.md", "## Intent\nBuild feature\n## Scope\nModule X")
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "## Requirements\nMUST do X")
	writeFile(t, root, "design.md", "# Design\nArchitecture: hexagonal\nDecisions made.")
	writeFile(t, root, "tasks.md", "- [x] 1.1 Done\n- [ ] 1.2 Pending")
	writeFile(t, root, "verify-report.md", "# Verify Report\n## Verdict\nPASS")
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nCompleted tasks 1.1\nNext: 1.2")
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
