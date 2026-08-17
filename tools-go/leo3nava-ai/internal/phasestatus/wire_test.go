package phasestatus

import (
	"testing"
)

// goldenJSON is the canonical frozen wire output for a fully populated Status.
// It must be byte-identical across runs — this is the contract.
const goldenJSON = `{"schemaName":"phase-status-validator","schemaVersion":1,"changeName":"my-feature","artifacts":{"apply-progress":"done","design":"done","proposal":"done","specs":"done","tasks":"done","verify-report":"done"},"taskProgress":{"total":5,"checked":5,"percent":100},"dependencies":{"apply":"all_done","design":"all_done","proposal":"all_done","specs":"all_done","tasks":"all_done","verify":"all_done"},"nextRecommended":"none","blockedReasons":[]}`

func TestProjectWireGolden(t *testing.T) {
	status := Status{
		SchemaName:    "phase-status-validator",
		SchemaVersion: 1,
		ChangeName:    "my-feature",
		Artifacts: map[ArtifactType]ArtifactState{
			ArtifactProposal:      ArtifactDone,
			ArtifactSpecs:         ArtifactDone,
			ArtifactDesign:        ArtifactDone,
			ArtifactTasks:         ArtifactDone,
			ArtifactApplyProgress: ArtifactDone,
			ArtifactVerifyReport:  ArtifactDone,
		},
		TaskProgress: TaskProgress{
			Total:   5,
			Checked: 5,
			Percent: 100.0,
		},
		Dependencies: map[string]DependencyState{
			PhaseProposal: DepAllDone,
			PhaseSpecs:    DepAllDone,
			PhaseDesign:   DepAllDone,
			PhaseTasks:    DepAllDone,
			PhaseApply:    DepAllDone,
			PhaseVerify:   DepAllDone,
		},
		NextRecommended: NextNone,
		BlockedReasons:  []string{},
	}

	got, err := ProjectWire(status)
	if err != nil {
		t.Fatalf("ProjectWire returned error: %v", err)
	}

	if string(got) != goldenJSON {
		t.Errorf("wire output mismatch\ngot:\n%s\nwant:\n%s", got, goldenJSON)
	}
}

func TestProjectWireDeterministic(t *testing.T) {
	// Run twice and confirm byte-identical output.
	status := Status{
		SchemaName:    "phase-status-validator",
		SchemaVersion: 1,
		ChangeName:    "my-feature",
		Artifacts: map[ArtifactType]ArtifactState{
			ArtifactProposal:      ArtifactDone,
			ArtifactSpecs:         ArtifactDone,
			ArtifactDesign:        ArtifactDone,
			ArtifactTasks:         ArtifactDone,
			ArtifactApplyProgress: ArtifactDone,
			ArtifactVerifyReport:  ArtifactDone,
		},
		TaskProgress: TaskProgress{
			Total:   5,
			Checked: 5,
			Percent: 100.0,
		},
		Dependencies: map[string]DependencyState{
			PhaseProposal: DepAllDone,
			PhaseSpecs:    DepAllDone,
			PhaseDesign:   DepAllDone,
			PhaseTasks:    DepAllDone,
			PhaseApply:    DepAllDone,
			PhaseVerify:   DepAllDone,
		},
		NextRecommended: NextNone,
		BlockedReasons:  []string{},
	}

	first, err := ProjectWire(status)
	if err != nil {
		t.Fatalf("first call error: %v", err)
	}

	second, err := ProjectWire(status)
	if err != nil {
		t.Fatalf("second call error: %v", err)
	}

	if string(first) != string(second) {
		t.Errorf("non-deterministic output:\nfirst:\n%s\nsecond:\n%s", first, second)
	}
}

func TestProjectWireUnknownToken(t *testing.T) {
	status := Status{
		SchemaName:      "phase-status-validator",
		SchemaVersion:   1,
		ChangeName:      "test",
		Artifacts:       map[ArtifactType]ArtifactState{},
		TaskProgress:    TaskProgress{},
		Dependencies:    map[string]DependencyState{},
		NextRecommended: NextRecommended("unknown-token"),
		BlockedReasons:  []string{},
	}

	_, err := ProjectWire(status)
	if err == nil {
		t.Error("expected error for unknown token, got nil")
	}
}

func TestProjectWireBlockedReasons(t *testing.T) {
	status := Status{
		SchemaName:    "test",
		SchemaVersion: 1,
		ChangeName:    "test",
		Artifacts:     map[ArtifactType]ArtifactState{},
		TaskProgress:  TaskProgress{},
		Dependencies:  map[string]DependencyState{},
		BlockedReasons: []string{
			"tasks complete but apply-progress artifact is missing or partial",
		},
		NextRecommended: NextNone,
	}

	got, err := ProjectWire(status)
	if err != nil {
		t.Fatalf("ProjectWire returned error: %v", err)
	}

	// Verify the blockedReasons array is present and contains the reason.
	want := `"blockedReasons":["tasks complete but apply-progress artifact is missing or partial"]`
	if !contains(string(got), want) {
		t.Errorf("expected blockedReasons in output, got:\n%s", got)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
