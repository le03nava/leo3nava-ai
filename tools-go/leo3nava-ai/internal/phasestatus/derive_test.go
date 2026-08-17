package phasestatus

import "testing"

func TestDeriveTaskProgressAllChecked(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	content := map[string]string{
		"tasks": "- [x] 1.1 Do thing\n- [x] 1.2 Do other thing\n- [x] 1.3 Last thing",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 3 {
		t.Errorf("total: want 3, got %d", got.Total)
	}
	if got.Checked != 3 {
		t.Errorf("checked: want 3, got %d", got.Checked)
	}
	if got.Percent != 100.0 {
		t.Errorf("percent: want 100.0, got %f", got.Percent)
	}
}

func TestDeriveTaskProgressNoneChecked(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	content := map[string]string{
		"tasks": "- [ ] 1.1 Do thing\n- [ ] 1.2 Do other thing",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 2 {
		t.Errorf("total: want 2, got %d", got.Total)
	}
	if got.Checked != 0 {
		t.Errorf("checked: want 0, got %d", got.Checked)
	}
	if got.Percent != 0.0 {
		t.Errorf("percent: want 0.0, got %f", got.Percent)
	}
}

func TestDeriveTaskProgressMixed(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	content := map[string]string{
		"tasks": "- [x] 1.1 Done\n- [ ] 1.2 Pending\n- [x] 1.3 Done\n- [ ] 1.4 Pending\n- [x] 1.5 Done\n- [ ] 1.6 Pending",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 6 {
		t.Errorf("total: want 6, got %d", got.Total)
	}
	if got.Checked != 3 {
		t.Errorf("checked: want 3, got %d", got.Checked)
	}
	if got.Percent != 50.0 {
		t.Errorf("percent: want 50.0, got %f", got.Percent)
	}
}

func TestDeriveTaskProgressNoCheckboxes(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactPartial,
	}
	content := map[string]string{
		"tasks": "# Tasks\nJust some text, no checkboxes",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 0 {
		t.Errorf("total: want 0, got %d", got.Total)
	}
	if got.Checked != 0 {
		t.Errorf("checked: want 0, got %d", got.Checked)
	}
	if got.Percent != 0.0 {
		t.Errorf("percent: want 0.0, got %f", got.Percent)
	}
}

func TestDeriveTaskProgressTasksMissing(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactMissing,
	}
	content := map[string]string{}

	got := deriveTaskProgress(m, content)

	if got.Total != 0 {
		t.Errorf("total: want 0, got %d", got.Total)
	}
	if got.Percent != 0.0 {
		t.Errorf("percent: want 0.0, got %f", got.Percent)
	}
}

func TestDeriveTaskProgressSingleTask(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	content := map[string]string{
		"tasks": "- [x] 1.1 The only task",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 1 {
		t.Errorf("total: want 1, got %d", got.Total)
	}
	if got.Checked != 1 {
		t.Errorf("checked: want 1, got %d", got.Checked)
	}
	if got.Percent != 100.0 {
		t.Errorf("percent: want 100.0, got %f", got.Percent)
	}
}

// --- Dependency derivation tests ---

func TestDeriveDependenciesAllPhasesComplete(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactProposal:      ArtifactDone,
		ArtifactSpecs:         ArtifactDone,
		ArtifactDesign:        ArtifactDone,
		ArtifactTasks:         ArtifactDone,
		ArtifactApplyProgress: ArtifactDone,
		ArtifactVerifyReport:  ArtifactDone,
	}
	tp := TaskProgress{Total: 3, Checked: 3, Percent: 100.0}

	deps, reasons := deriveDependencies(m, tp)

	if len(reasons) != 0 {
		t.Errorf("blockedReasons: want empty, got %v", reasons)
	}
	for _, phase := range PhaseOrder {
		if state := deps[phase]; state != DepAllDone {
			t.Errorf("phase %s: want %s, got %s", phase, DepAllDone, state)
		}
	}
}

func TestDeriveDependenciesBlockedByMissingDesign(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactProposal: ArtifactDone,
		ArtifactSpecs:    ArtifactDone,
		ArtifactDesign:   ArtifactMissing,
		ArtifactTasks:    ArtifactMissing,
	}
	tp := TaskProgress{}

	deps, reasons := deriveDependencies(m, tp)

	if len(reasons) != 0 {
		t.Errorf("blockedReasons: want empty, got %v", reasons)
	}
	// proposal and specs are all_done
	if deps[PhaseProposal] != DepAllDone {
		t.Errorf("proposal: want %s, got %s", DepAllDone, deps[PhaseProposal])
	}
	if deps[PhaseSpecs] != DepAllDone {
		t.Errorf("specs: want %s, got %s", DepAllDone, deps[PhaseSpecs])
	}
	// design is ready (prior gates satisfied, artifact missing)
	if deps[PhaseDesign] != DepReady {
		t.Errorf("design: want %s, got %s", DepReady, deps[PhaseDesign])
	}
	// tasks is blocked (design not all_done)
	if deps[PhaseTasks] != DepBlocked {
		t.Errorf("tasks: want %s, got %s", DepBlocked, deps[PhaseTasks])
	}
}

func TestDeriveDependenciesEarlyMissingProposal(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactProposal: ArtifactMissing,
		ArtifactSpecs:    ArtifactMissing,
		ArtifactDesign:   ArtifactMissing,
	}
	tp := TaskProgress{}

	deps, reasons := deriveDependencies(m, tp)

	if len(reasons) != 0 {
		t.Errorf("blockedReasons: want empty, got %v", reasons)
	}
	// proposal is ready (first in chain, artifact missing)
	if deps[PhaseProposal] != DepReady {
		t.Errorf("proposal: want %s, got %s", DepReady, deps[PhaseProposal])
	}
	// everything else is blocked
	for _, phase := range PhaseOrder[1:] {
		if deps[phase] != DepBlocked {
			t.Errorf("phase %s: want %s, got %s", phase, DepBlocked, deps[phase])
		}
	}
}

func TestDeriveDependenciesAnomalyTasksDoneNoApplyProgress(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactProposal:      ArtifactDone,
		ArtifactSpecs:         ArtifactDone,
		ArtifactDesign:        ArtifactDone,
		ArtifactTasks:         ArtifactDone,
		ArtifactApplyProgress: ArtifactMissing,
		ArtifactVerifyReport:  ArtifactMissing,
	}
	tp := TaskProgress{Total: 5, Checked: 5, Percent: 100.0}

	deps, reasons := deriveDependencies(m, tp)

	// All gates before apply should be all_done
	if deps[PhaseProposal] != DepAllDone {
		t.Errorf("proposal: want %s, got %s", DepAllDone, deps[PhaseProposal])
	}
	if deps[PhaseSpecs] != DepAllDone {
		t.Errorf("specs: want %s, got %s", DepAllDone, deps[PhaseSpecs])
	}
	if deps[PhaseDesign] != DepAllDone {
		t.Errorf("design: want %s, got %s", DepAllDone, deps[PhaseDesign])
	}
	if deps[PhaseTasks] != DepAllDone {
		t.Errorf("tasks: want %s, got %s", DepAllDone, deps[PhaseTasks])
	}
	// apply is blocked (anomaly: tasks 100% but apply-progress missing)
	if deps[PhaseApply] != DepBlocked {
		t.Errorf("apply: want %s, got %s", DepBlocked, deps[PhaseApply])
	}
	// verify is blocked (apply not all_done)
	if deps[PhaseVerify] != DepBlocked {
		t.Errorf("verify: want %s, got %s", DepBlocked, deps[PhaseVerify])
	}
	// Anomaly should produce blockedReasons
	if len(reasons) == 0 {
		t.Error("blockedReasons: want non-empty for anomaly, got empty")
	}
}

func TestDeriveDependenciesAnomalyTasksDoneApplyPartial(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactProposal:      ArtifactDone,
		ArtifactSpecs:         ArtifactDone,
		ArtifactDesign:        ArtifactDone,
		ArtifactTasks:         ArtifactDone,
		ArtifactApplyProgress: ArtifactPartial,
		ArtifactVerifyReport:  ArtifactMissing,
	}
	// Tasks not 100% — not an anomaly, just normal apply in progress
	tp := TaskProgress{Total: 5, Checked: 3, Percent: 60.0}

	deps, reasons := deriveDependencies(m, tp)

	// All prior gates all_done
	for _, phase := range PhaseOrder[:4] {
		if deps[phase] != DepAllDone {
			t.Errorf("phase %s: want %s, got %s", phase, DepAllDone, deps[phase])
		}
	}
	// apply is ready (prior gates satisfied, tasks not 100% — normal flow)
	if deps[PhaseApply] != DepReady {
		t.Errorf("apply: want %s, got %s", DepReady, deps[PhaseApply])
	}
	// verify is blocked
	if deps[PhaseVerify] != DepBlocked {
		t.Errorf("verify: want %s, got %s", DepBlocked, deps[PhaseVerify])
	}
	if len(reasons) != 0 {
		t.Errorf("blockedReasons: want empty (normal flow), got %v", reasons)
	}
}

func TestDeriveTaskProgressPercentRounding(t *testing.T) {
	m := map[ArtifactType]ArtifactState{
		ArtifactTasks: ArtifactDone,
	}
	content := map[string]string{
		"tasks": "- [x] 1.1 Done\n- [ ] 1.2 Pending\n- [ ] 1.3 Pending\n- [ ] 1.4 Pending",
	}

	got := deriveTaskProgress(m, content)

	if got.Total != 4 {
		t.Errorf("total: want 4, got %d", got.Total)
	}
	if got.Checked != 1 {
		t.Errorf("checked: want 1, got %d", got.Checked)
	}
	if got.Percent != 25.0 {
		t.Errorf("percent: want 25.0, got %f", got.Percent)
	}
}
