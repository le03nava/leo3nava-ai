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
