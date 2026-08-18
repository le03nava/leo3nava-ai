package phasestatus

// ArtifactType enumerates the six SDD artifact types.
type ArtifactType int

const (
	ArtifactProposal ArtifactType = iota
	ArtifactSpecs
	ArtifactDesign
	ArtifactTasks
	ArtifactVerifyReport
	ArtifactApplyProgress
)

// String returns the canonical name for an ArtifactType.
func (t ArtifactType) String() string {
	switch t {
	case ArtifactProposal:
		return "proposal"
	case ArtifactSpecs:
		return "specs"
	case ArtifactDesign:
		return "design"
	case ArtifactTasks:
		return "tasks"
	case ArtifactVerifyReport:
		return "verify-report"
	case ArtifactApplyProgress:
		return "apply-progress"
	default:
		return "unknown"
	}
}

// ArtifactMeta holds the detection rule for an artifact type.
type ArtifactMeta struct {
	Type        ArtifactType
	Name        string
	SingleFile  string // non-empty means detect by exact filename
	GlobPattern string // non-empty means detect by glob
}

// Catalog is the ordered list of artifact types and their detection rules.
var Catalog = []ArtifactMeta{
	{Type: ArtifactProposal, Name: "proposal", SingleFile: "proposal.md"},
	{Type: ArtifactSpecs, Name: "specs", GlobPattern: "specs/**/spec.md"},
	{Type: ArtifactDesign, Name: "design", SingleFile: "design.md"},
	{Type: ArtifactTasks, Name: "tasks", SingleFile: "tasks.md"},
	{Type: ArtifactVerifyReport, Name: "verify-report", SingleFile: "verify-report.md"},
	{Type: ArtifactApplyProgress, Name: "apply-progress", SingleFile: "apply-progress.md"},
}

// ArtifactNames maps ArtifactType to its human-readable name.
var ArtifactNames = map[ArtifactType]string{
	ArtifactProposal:      "proposal",
	ArtifactSpecs:         "specs",
	ArtifactDesign:        "design",
	ArtifactTasks:         "tasks",
	ArtifactVerifyReport:  "verify-report",
	ArtifactApplyProgress: "apply-progress",
}
