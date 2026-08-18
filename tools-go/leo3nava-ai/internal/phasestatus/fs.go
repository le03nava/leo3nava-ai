package phasestatus

import (
	"os"
	"path/filepath"
	"strings"
)

// detectArtifacts scans the change root directory and returns the on-disk state
// of each artifact type. Single-file types are checked by exact filename.
// Specs uses a glob for nested layout with a flat spec.md fallback.
func detectArtifacts(root string) map[ArtifactType]ArtifactState {
	result := make(map[ArtifactType]ArtifactState, len(Catalog))

	for _, meta := range Catalog {
		result[meta.Type] = detectOne(root, meta)
	}

	return result
}

// detectOne checks whether a single artifact type exists on disk.
func detectOne(root string, meta ArtifactMeta) ArtifactState {
	if meta.SingleFile != "" {
		path := filepath.Join(root, meta.SingleFile)
		if fileExists(path) {
			return ArtifactDone
		}
		return ArtifactMissing
	}

	if meta.GlobPattern != "" {
		// Try nested glob first.
		pattern := filepath.Join(root, meta.GlobPattern)
		matches, err := filepath.Glob(pattern)
		if err == nil && len(matches) > 0 {
			return ArtifactDone
		}

		// Flat fallback: specs/ subdirectory with any .md file.
		specsDir := filepath.Join(root, "specs")
		if dir, err := os.Stat(specsDir); err == nil && dir.IsDir() {
			entries, err := os.ReadDir(specsDir)
			if err == nil {
				for _, e := range entries {
					if !e.IsDir() && filepath.Ext(e.Name()) == ".md" {
						return ArtifactDone
					}
				}
			}
		}

		// Flat fallback: spec.md at root.
		flatPath := filepath.Join(root, "spec.md")
		if fileExists(flatPath) {
			return ArtifactDone
		}

		return ArtifactMissing
	}

	return ArtifactMissing
}

// fileExists returns true if the path exists and is a regular file.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

// validateContent reads each present artifact and enforces its content rule.
// Artifacts that are missing stay missing. Artifacts that exist but fail
// their rule are downgraded to partial.
func validateContent(root string, m map[ArtifactType]ArtifactState) map[ArtifactType]ArtifactState {
	result := make(map[ArtifactType]ArtifactState, len(m))
	for k, v := range m {
		result[k] = v
	}

	for _, meta := range Catalog {
		if result[meta.Type] == ArtifactMissing {
			continue
		}
		content := readArtifactContent(root, meta)
		if !passesContentRule(meta.Type, content) {
			result[meta.Type] = ArtifactPartial
		}
	}

	return result
}

// readArtifactContent reads the file content for an artifact type.
// For specs, it reads the first matching spec file found.
func readArtifactContent(root string, meta ArtifactMeta) string {
	if meta.SingleFile != "" {
		data, err := os.ReadFile(filepath.Join(root, meta.SingleFile))
		if err != nil {
			return ""
		}
		return string(data)
	}

	if meta.GlobPattern != "" {
		pattern := filepath.Join(root, meta.GlobPattern)
		matches, err := filepath.Glob(pattern)
		if err == nil && len(matches) > 0 {
			data, err := os.ReadFile(matches[0])
			if err == nil {
				return string(data)
			}
		}
		// Flat fallback.
		data, err := os.ReadFile(filepath.Join(root, "spec.md"))
		if err == nil {
			return string(data)
		}
	}

	return ""
}

// passesContentRule checks whether an artifact's content satisfies its type rule.
func passesContentRule(typ ArtifactType, content string) bool {
	switch typ {
	case ArtifactProposal:
		return strings.Contains(content, "## Intent") || strings.Contains(content, "## Scope")
	case ArtifactSpecs:
		return strings.Contains(content, "## Requirements")
	case ArtifactTasks:
		return strings.Contains(content, "- [ ]") || strings.Contains(content, "- [x]")
	case ArtifactVerifyReport:
		return hasVerdictSection(content)
	case ArtifactDesign, ArtifactApplyProgress:
		return hasSubstantiveContent(content)
	default:
		return true
	}
}

// hasVerdictSection checks for a verdict heading or keyword in the content.
func hasVerdictSection(content string) bool {
	lower := strings.ToLower(content)
	return strings.Contains(lower, "## verdict") || strings.Contains(lower, "# verdict") ||
		strings.Contains(lower, "verdict:")
}

// hasSubstantiveContent checks that content has meaningful text beyond headings.
func hasSubstantiveContent(content string) bool {
	lines := strings.Split(content, "\n")
	nonEmpty := 0
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			nonEmpty++
		}
	}
	return nonEmpty > 1
}
