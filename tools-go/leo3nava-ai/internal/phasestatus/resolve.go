package phasestatus

import (
	"fmt"
	"os"
	"path/filepath"
)

// Options configures the Resolve function.
type Options struct {
	// ChangeRoot is the absolute path to the change directory.
	ChangeRoot string
	// ChangeName overrides the inferred change name. If nil, the directory
	// basename is used.
	ChangeName *string
}

// Resolve derives the full phase status for a change directory.
// It orchestrates detect → validate → derive → route and returns a Status.
func Resolve(opts Options) (Status, error) {
	if opts.ChangeRoot == "" {
		return Status{}, fmt.Errorf("changeRoot is required")
	}

	// Verify the path exists and is a directory.
	info, err := os.Stat(opts.ChangeRoot)
	if err != nil {
		return Status{}, fmt.Errorf("changeRoot: %w", err)
	}
	if !info.IsDir() {
		return Status{}, fmt.Errorf("changeRoot is not a directory: %s", opts.ChangeRoot)
	}

	// 1. Detect on-disk artifacts.
	detected := detectArtifacts(opts.ChangeRoot)

	// 2. Validate content (may downgrade present → partial).
	artifacts := validateContent(opts.ChangeRoot, detected)

	// 3. Read artifact content for derivation.
	content := readAllContent(opts.ChangeRoot)

	// 4. Derive task progress.
	taskProgress := deriveTaskProgress(artifacts, content)

	// 5. Derive dependency states.
	deps, blockedReasons := deriveDependencies(artifacts, taskProgress)

	// 6. Route next recommended.
	next := routeNext(artifacts, deps, blockedReasons)

	// Infer change name.
	changeName := opts.ChangeName
	name := ""
	if changeName != nil {
		name = *changeName
	} else {
		name = filepath.Base(opts.ChangeRoot)
	}

	return Status{
		SchemaName:      "phase-status-validator",
		SchemaVersion:   1,
		ChangeName:      name,
		Artifacts:       artifacts,
		TaskProgress:    taskProgress,
		Dependencies:    deps,
		NextRecommended: next,
		BlockedReasons:  blockedReasons,
	}, nil
}

// readAllContent reads the content of each artifact type for derivation.
func readAllContent(root string) map[string]string {
	content := make(map[string]string, len(Catalog))
	for _, meta := range Catalog {
		c := readArtifactContent(root, meta)
		if c != "" {
			content[meta.Name] = c
		}
	}
	return content
}
