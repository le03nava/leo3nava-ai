package phasestatus

// ArtifactState represents the on-disk state of an SDD artifact.
type ArtifactState int

const (
	// ArtifactMissing indicates the artifact file does not exist on disk.
	ArtifactMissing ArtifactState = iota
	// ArtifactPartial indicates the artifact exists but fails its content validation rule.
	ArtifactPartial
	// ArtifactDone indicates the artifact exists and passes its content validation rule.
	ArtifactDone
)

// String returns a human-readable label for an ArtifactState.
func (s ArtifactState) String() string {
	switch s {
	case ArtifactMissing:
		return "missing"
	case ArtifactPartial:
		return "partial"
	case ArtifactDone:
		return "done"
	default:
		return "unknown"
	}
}

// DependencyState represents whether a phase gate is satisfied.
type DependencyState int

const (
	// DepBlocked indicates the phase cannot proceed because prior gates are not satisfied.
	DepBlocked DependencyState = iota
	// DepReady indicates prior gates are satisfied and this phase is the next candidate.
	DepReady
	// DepAllDone indicates this phase gate and all prior gates are complete.
	DepAllDone
)

// String returns a human-readable label for a DependencyState.
func (s DependencyState) String() string {
	switch s {
	case DepBlocked:
		return "blocked"
	case DepReady:
		return "ready"
	case DepAllDone:
		return "all_done"
	default:
		return "unknown"
	}
}

// NextRecommended is the routing decision token.
type NextRecommended string

const (
	NextProposal NextRecommended = "proposal"
	NextSpecs    NextRecommended = "specs"
	NextDesign   NextRecommended = "design"
	NextTasks    NextRecommended = "tasks"
	NextApply    NextRecommended = "apply"
	NextVerify   NextRecommended = "verify"
	NextNone     NextRecommended = "none"
)

// Phase names used as dependency keys.
const (
	PhaseProposal = "proposal"
	PhaseSpecs    = "specs"
	PhaseDesign   = "design"
	PhaseTasks    = "tasks"
	PhaseApply    = "apply"
	PhaseVerify   = "verify"
)

// PhaseOrder defines the canonical dependency gate order.
var PhaseOrder = []string{
	PhaseProposal,
	PhaseSpecs,
	PhaseDesign,
	PhaseTasks,
	PhaseApply,
	PhaseVerify,
}

// TaskProgress holds checkbox parsing results for tasks artifact.
type TaskProgress struct {
	Total   int
	Checked int
	Percent float64
}

// Status is the top-level structured state for a change.
type Status struct {
	SchemaName      string                         `json:"schemaName"`
	SchemaVersion   int                            `json:"schemaVersion"`
	ChangeName      string                         `json:"changeName"`
	Artifacts       map[ArtifactType]ArtifactState `json:"artifacts"`
	TaskProgress    TaskProgress                   `json:"taskProgress"`
	Dependencies    map[string]DependencyState     `json:"dependencies"`
	NextRecommended NextRecommended                `json:"nextRecommended"`
	BlockedReasons  []string                       `json:"blockedReasons"`
}
