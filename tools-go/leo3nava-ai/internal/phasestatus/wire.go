package phasestatus

import (
	"encoding/json"
	"fmt"
	"sort"
)

// validNextTokens is the set of recognized NextRecommended tokens.
var validNextTokens = map[NextRecommended]bool{
	NextProposal: true,
	NextSpecs:    true,
	NextDesign:   true,
	NextTasks:    true,
	NextApply:    true,
	NextVerify:   true,
	NextNone:     true,
}

// wireStatus is the JSON-serializable representation of Status.
// ArtifactType and DependencyState are projected as string labels
// for a stable, human-readable wire format.
type wireStatus struct {
	SchemaName      string            `json:"schemaName"`
	SchemaVersion   int               `json:"schemaVersion"`
	ChangeName      string            `json:"changeName"`
	Artifacts       map[string]string `json:"artifacts"`
	TaskProgress    wireTaskProgress  `json:"taskProgress"`
	Dependencies    map[string]string `json:"dependencies"`
	NextRecommended string            `json:"nextRecommended"`
	BlockedReasons  []string          `json:"blockedReasons"`
}

type wireTaskProgress struct {
	Total   int     `json:"total"`
	Checked int     `json:"checked"`
	Percent float64 `json:"percent"`
}

// ProjectWire serializes a Status into a frozen JSON wire document.
// Unknown NextRecommended tokens are rejected with an error.
func ProjectWire(s Status) ([]byte, error) {
	if !validNextTokens[s.NextRecommended] {
		return nil, fmt.Errorf("unknown nextRecommended token: %q", s.NextRecommended)
	}

	// Build artifacts map with sorted string keys for deterministic output.
	artifactKeys := make([]ArtifactType, 0, len(s.Artifacts))
	for k := range s.Artifacts {
		artifactKeys = append(artifactKeys, k)
	}
	sort.Slice(artifactKeys, func(i, j int) bool {
		return artifactKeys[i].String() < artifactKeys[j].String()
	})

	artifacts := make(map[string]string, len(s.Artifacts))
	for _, k := range artifactKeys {
		artifacts[k.String()] = s.Artifacts[k].String()
	}

	// Build dependencies map with sorted string keys.
	depKeys := make([]string, 0, len(s.Dependencies))
	for k := range s.Dependencies {
		depKeys = append(depKeys, k)
	}
	sort.Strings(depKeys)

	deps := make(map[string]string, len(s.Dependencies))
	for _, k := range depKeys {
		deps[k] = s.Dependencies[k].String()
	}

	wire := wireStatus{
		SchemaName:    s.SchemaName,
		SchemaVersion: s.SchemaVersion,
		ChangeName:    s.ChangeName,
		Artifacts:     artifacts,
		TaskProgress: wireTaskProgress{
			Total:   s.TaskProgress.Total,
			Checked: s.TaskProgress.Checked,
			Percent: s.TaskProgress.Percent,
		},
		Dependencies:    deps,
		NextRecommended: string(s.NextRecommended),
		BlockedReasons:  s.BlockedReasons,
	}

	return json.Marshal(wire)
}
