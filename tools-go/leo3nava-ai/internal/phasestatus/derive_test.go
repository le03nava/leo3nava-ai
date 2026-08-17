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
