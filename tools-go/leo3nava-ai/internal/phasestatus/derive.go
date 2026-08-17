package phasestatus

import (
	"strings"
)

// deriveTaskProgress parses checkbox lines from tasks content and computes
// the total number of tasks, how many are checked, and the completion percentage.
func deriveTaskProgress(m map[ArtifactType]ArtifactState, content map[string]string) TaskProgress {
	if m[ArtifactTasks] == ArtifactMissing {
		return TaskProgress{}
	}

	tasksContent, ok := content["tasks"]
	if !ok {
		return TaskProgress{}
	}

	var total, checked int
	for _, line := range strings.Split(tasksContent, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- [x]") {
			total++
			checked++
		} else if strings.HasPrefix(trimmed, "- [ ]") {
			total++
		}
	}

	if total == 0 {
		return TaskProgress{}
	}

	return TaskProgress{
		Total:   total,
		Checked: checked,
		Percent: float64(checked) / float64(total) * 100,
	}
}

// phaseArtifact maps each phase name to the artifact type that gates it.
var phaseArtifact = map[string]ArtifactType{
	PhaseProposal: ArtifactProposal,
	PhaseSpecs:    ArtifactSpecs,
	PhaseDesign:   ArtifactDesign,
	PhaseTasks:    ArtifactTasks,
	PhaseApply:    ArtifactApplyProgress,
	PhaseVerify:   ArtifactVerifyReport,
}

// deriveDependencies computes per-phase dependency state by walking the ordered
// phase gates: proposal → specs → design → tasks → apply → verify.
// A gate is all_done when its artifact is done, ready when all prior gates
// are all_done and the artifact is not done, and blocked otherwise.
// Anomalies (e.g., all tasks checked but apply-progress missing) produce
// blockedReasons and override the phase to blocked.
func deriveDependencies(m map[ArtifactType]ArtifactState, tp TaskProgress) (map[string]DependencyState, []string) {
	deps := make(map[string]DependencyState, len(PhaseOrder))
	var blockedReasons []string
	priorAllDone := true

	for _, phase := range PhaseOrder {
		artifact, ok := phaseArtifact[phase]
		if !ok {
			deps[phase] = DepBlocked
			continue
		}

		artifactDone := m[artifact] == ArtifactDone

		switch {
		case artifactDone:
			deps[phase] = DepAllDone
		case priorAllDone:
			// Detect anomaly: all tasks checked but apply-progress not done.
			if phase == PhaseApply && tp.Total > 0 && tp.Checked == tp.Total {
				deps[phase] = DepBlocked
				blockedReasons = append(blockedReasons,
					"tasks complete but apply-progress artifact is missing or partial")
			} else {
				deps[phase] = DepReady
			}
		default:
			deps[phase] = DepBlocked
		}

		priorAllDone = priorAllDone && artifactDone
	}

	return deps, blockedReasons
}
