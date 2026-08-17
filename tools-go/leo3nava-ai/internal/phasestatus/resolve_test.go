package phasestatus

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveAllArtifactsPresent(t *testing.T) {
	root := t.TempDir()
	seedAllArtifacts(t, root)

	status, err := Resolve(Options{ChangeRoot: root})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if status.NextRecommended != NextNone {
		t.Errorf("nextRecommended: want %s, got %s", NextNone, status.NextRecommended)
	}
	if len(status.BlockedReasons) != 0 {
		t.Errorf("blockedReasons: want empty, got %v", status.BlockedReasons)
	}
	// All dependencies should be all_done.
	for _, phase := range PhaseOrder {
		if state := status.Dependencies[phase]; state != DepAllDone {
			t.Errorf("dependency %s: want %s, got %s", phase, DepAllDone, state)
		}
	}
	// All artifacts should be done.
	for _, meta := range Catalog {
		if state := status.Artifacts[meta.Type]; state != ArtifactDone {
			t.Errorf("artifact %s: want %s, got %s", meta.Name, ArtifactDone, state)
		}
	}
}

func TestResolveMissingDesign(t *testing.T) {
	root := t.TempDir()
	seedAllArtifacts(t, root)
	// Remove design.md to simulate missing artifact.
	os.Remove(filepath.Join(root, "design.md"))

	status, err := Resolve(Options{ChangeRoot: root})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if status.NextRecommended != NextDesign {
		t.Errorf("nextRecommended: want %s, got %s", NextDesign, status.NextRecommended)
	}
	if state := status.Artifacts[ArtifactDesign]; state != ArtifactMissing {
		t.Errorf("design artifact: want %s, got %s", ArtifactMissing, state)
	}
	// Design dependency should be ready (prior gates satisfied).
	if state := status.Dependencies[PhaseDesign]; state != DepReady {
		t.Errorf("design dependency: want %s, got %s", DepReady, state)
	}
}

func TestResolveAnomalyTasksDoneNoApply(t *testing.T) {
	root := t.TempDir()
	seedAllArtifactsComplete(t, root) // all done including tasks with 100%
	// Remove apply-progress to create anomaly.
	os.Remove(filepath.Join(root, "apply-progress.md"))

	status, err := Resolve(Options{ChangeRoot: root})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if status.NextRecommended != NextNone {
		t.Errorf("nextRecommended: want %s, got %s", NextNone, status.NextRecommended)
	}
	if len(status.BlockedReasons) == 0 {
		t.Error("blockedReasons: want non-empty for anomaly, got empty")
	}
	// Apply dependency should be blocked.
	if state := status.Dependencies[PhaseApply]; state != DepBlocked {
		t.Errorf("apply dependency: want %s, got %s", DepBlocked, state)
	}
}

func TestResolveInvalidPath(t *testing.T) {
	_, err := Resolve(Options{ChangeRoot: "/nonexistent/path/that/does/not/exist"})
	if err == nil {
		t.Error("expected error for invalid path, got nil")
	}
}

func TestResolveChangeNameInferred(t *testing.T) {
	root := t.TempDir()
	seedAllArtifacts(t, root)

	status, err := Resolve(Options{ChangeRoot: root})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	// ChangeName should be inferred from directory name.
	dirName := filepath.Base(root)
	if status.ChangeName != dirName {
		t.Errorf("changeName: want %q, got %q", dirName, status.ChangeName)
	}
}

func TestResolveChangeNameExplicit(t *testing.T) {
	root := t.TempDir()
	seedAllArtifacts(t, root)
	name := "my-explicit-name"

	status, err := Resolve(Options{ChangeRoot: root, ChangeName: &name})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if status.ChangeName != name {
		t.Errorf("changeName: want %q, got %q", name, status.ChangeName)
	}
}

// seedAllArtifacts creates valid artifact files for a change in progress
// (not all tasks checked, so no anomaly).
func seedAllArtifacts(t *testing.T, root string) {
	t.Helper()
	writeFile(t, root, "proposal.md", "## Intent\nBuild feature\n## Scope\nModule X")
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "## Requirements\nMUST do X")
	writeFile(t, root, "design.md", "# Design\nArchitecture: hexagonal pattern\nDecisions made.")
	writeFile(t, root, "tasks.md", "- [x] 1.1 Done\n- [ ] 1.2 Pending\n- [ ] 1.3 Pending")
	writeFile(t, root, "verify-report.md", "# Verify Report\n## Verdict\nPASS")
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nCompleted tasks 1.1\nNext: 1.2")
}

// seedAllArtifactsComplete creates valid artifact files with all tasks checked
// (100% complete, suitable for testing anomalies when apply-progress is missing).
func seedAllArtifactsComplete(t *testing.T, root string) {
	t.Helper()
	writeFile(t, root, "proposal.md", "## Intent\nBuild feature\n## Scope\nModule X")
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "## Requirements\nMUST do X")
	writeFile(t, root, "design.md", "# Design\nArchitecture: hexagonal pattern\nDecisions made.")
	writeFile(t, root, "tasks.md", "- [x] 1.1 Done\n- [x] 1.2 Done\n- [x] 1.3 Done")
	writeFile(t, root, "verify-report.md", "# Verify Report\n## Verdict\nPASS")
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nCompleted tasks 1.1-1.3\nAll done")
}
