package phasestatus

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectArtifactsAllPresent(t *testing.T) {
	root := t.TempDir()

	// Seed all artifact files.
	writeFile(t, root, "proposal.md", "## Intent\nBuild something\n## Scope\nFull")
	writeFile(t, root, "design.md", "# Design\nArchitecture here")
	writeFile(t, root, "tasks.md", "- [ ] 1.1 Do thing")
	writeFile(t, root, "verify-report.md", "# Verify Report\nVerdict: PASS")
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nDone stuff")
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "# Auth Spec\n## Requirements\n...")

	got := detectArtifacts(root)

	for _, meta := range Catalog {
		state, ok := got[meta.Type]
		if !ok {
			t.Fatalf("artifact %s missing from result map", meta.Name)
		}
		if state != ArtifactDone {
			t.Errorf("artifact %s: want %s, got %s", meta.Name, ArtifactDone, state)
		}
	}
}

func TestDetectArtifactsMissingDesign(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "proposal.md", "## Intent\nBuild something\n## Scope\nFull")
	writeFile(t, root, "tasks.md", "- [ ] 1.1 Do thing")
	writeFile(t, root, "verify-report.md", "# Verify Report\nVerdict: PASS")
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nDone stuff")
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "# Auth Spec\n## Requirements\n...")
	// design.md intentionally omitted

	got := detectArtifacts(root)

	if state := got[ArtifactDesign]; state != ArtifactMissing {
		t.Errorf("design: want %s, got %s", ArtifactMissing, state)
	}

	// All others should be present (at file level; content validation happens later).
	for _, meta := range Catalog {
		if meta.Type == ArtifactDesign {
			continue
		}
		state, ok := got[meta.Type]
		if !ok {
			t.Fatalf("artifact %s missing from result map", meta.Name)
		}
		if state == ArtifactMissing {
			t.Errorf("artifact %s: expected present, got missing", meta.Name)
		}
	}
}

func TestDetectArtifactsNestedSpecs(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "# Auth Spec\n## Requirements\n...")
	writeFile(t, root, filepath.Join("specs", "routing", "spec.md"), "# Routing Spec\n## Requirements\n...")

	got := detectArtifacts(root)

	if state := got[ArtifactSpecs]; state != ArtifactDone {
		t.Errorf("specs: want %s, got %s", ArtifactDone, state)
	}
}

func TestDetectArtifactsNoSpecsFiles(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "proposal.md", "## Intent\nSomething\n## Scope\nFull")

	got := detectArtifacts(root)

	if state := got[ArtifactSpecs]; state != ArtifactMissing {
		t.Errorf("specs: want %s, got %s", ArtifactMissing, state)
	}
}

func TestDetectArtifactsEmptyDirAllMissing(t *testing.T) {
	root := t.TempDir()

	got := detectArtifacts(root)

	for _, meta := range Catalog {
		state, ok := got[meta.Type]
		if !ok {
			t.Fatalf("artifact %s missing from result map", meta.Name)
		}
		if state != ArtifactMissing {
			t.Errorf("artifact %s: want %s on empty dir, got %s", meta.Name, ArtifactMissing, state)
		}
	}
}

func TestDetectArtifactsFlatSpecFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "spec.md", "# Spec\n## Requirements\n...")

	got := detectArtifacts(root)

	if state := got[ArtifactSpecs]; state != ArtifactDone {
		t.Errorf("specs (flat spec.md): want %s, got %s", ArtifactDone, state)
	}
}

func TestValidateContentProposalMissingIntentScope(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "proposal.md", "# Proposal\nJust a title here")

	state := map[ArtifactType]ArtifactState{
		ArtifactProposal: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactProposal] != ArtifactPartial {
		t.Errorf("proposal without Intent/Scope: want %s, got %s", ArtifactPartial, got[ArtifactProposal])
	}
}

func TestValidateContentProposalValid(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "proposal.md", "## Intent\nBuild feature\n## Scope\nModule X")

	state := map[ArtifactType]ArtifactState{
		ArtifactProposal: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactProposal] != ArtifactDone {
		t.Errorf("proposal with Intent+Scope: want %s, got %s", ArtifactDone, got[ArtifactProposal])
	}
}

func TestValidateContentSpecsMissingRequirements(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "# Auth Spec\nJust a title")

	state := map[ArtifactType]ArtifactState{
		ArtifactSpecs: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactSpecs] != ArtifactPartial {
		t.Errorf("specs without Requirements: want %s, got %s", ArtifactPartial, got[ArtifactSpecs])
	}
}

func TestValidateContentSpecsValid(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, filepath.Join("specs", "auth", "spec.md"), "## Requirements\nMUST do X")

	state := map[ArtifactType]ArtifactState{
		ArtifactSpecs: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactSpecs] != ArtifactDone {
		t.Errorf("specs with Requirements: want %s, got %s", ArtifactDone, got[ArtifactSpecs])
	}
}

func TestValidateContentTasksWithCheckboxes(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "tasks.md", "## Tasks\n- [ ] 1.1 Do thing\n- [x] 1.2 Done thing")

	state := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactTasks] != ArtifactDone {
		t.Errorf("tasks with checkboxes: want %s, got %s", ArtifactDone, got[ArtifactTasks])
	}
}

func TestValidateContentTasksWithoutCheckboxes(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "tasks.md", "## Tasks\nJust some text\nNo checkboxes here")

	state := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactTasks] != ArtifactPartial {
		t.Errorf("tasks without checkboxes: want %s, got %s", ArtifactPartial, got[ArtifactTasks])
	}
}

func TestValidateContentDesignEmpty(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "design.md", "# Design\n")

	state := map[ArtifactType]ArtifactState{
		ArtifactDesign: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactDesign] != ArtifactPartial {
		t.Errorf("design empty beyond headings: want %s, got %s", ArtifactPartial, got[ArtifactDesign])
	}
}

func TestValidateContentDesignSubstantive(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "design.md", "# Design\nArchitecture: hexagonal pattern\nDecisions made.")

	state := map[ArtifactType]ArtifactState{
		ArtifactDesign: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactDesign] != ArtifactDone {
		t.Errorf("design with substantive content: want %s, got %s", ArtifactDone, got[ArtifactDesign])
	}
}

func TestValidateContentVerifyReportMissingVerdict(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "verify-report.md", "# Verify Report\nResults here but no verdict")

	state := map[ArtifactType]ArtifactState{
		ArtifactVerifyReport: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactVerifyReport] != ArtifactPartial {
		t.Errorf("verify-report without verdict: want %s, got %s", ArtifactPartial, got[ArtifactVerifyReport])
	}
}

func TestValidateContentApplyProgressEmpty(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "apply-progress.md", "# Apply Progress\n")

	state := map[ArtifactType]ArtifactState{
		ArtifactApplyProgress: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactApplyProgress] != ArtifactPartial {
		t.Errorf("apply-progress empty: want %s, got %s", ArtifactPartial, got[ArtifactApplyProgress])
	}
}

func TestValidateContentApplyProgressSubstantive(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "apply-progress.md", "# Apply Progress\nCompleted tasks 1.1-1.3\nNext: 1.4")

	state := map[ArtifactType]ArtifactState{
		ArtifactApplyProgress: ArtifactDone,
	}
	got := validateContent(root, state)

	if got[ArtifactApplyProgress] != ArtifactDone {
		t.Errorf("apply-progress with content: want %s, got %s", ArtifactDone, got[ArtifactApplyProgress])
	}
}

func TestValidateContentMissingArtifactsUnchanged(t *testing.T) {
	root := t.TempDir()

	state := map[ArtifactType]ArtifactState{
		ArtifactDesign: ArtifactMissing,
		ArtifactTasks:  ArtifactMissing,
	}
	got := validateContent(root, state)

	if got[ArtifactDesign] != ArtifactMissing {
		t.Errorf("missing design should stay missing, got %s", got[ArtifactDesign])
	}
	if got[ArtifactTasks] != ArtifactMissing {
		t.Errorf("missing tasks should stay missing, got %s", got[ArtifactTasks])
	}
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
