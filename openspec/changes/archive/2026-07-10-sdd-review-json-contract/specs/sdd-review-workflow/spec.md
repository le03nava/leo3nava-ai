# Delta for sdd-review-workflow

## ADDED Requirements

### Requirement: Canonical Review JSON Authority

`sdd-review` MUST treat canonical JSON as the source of truth for general review report facts and all 96 `REV-CORP-*` controls. The JSON MUST directly define control IDs, source mappings, required report facts, vocabulary, validation metadata, and presentation-ready fields. `control-catalog.md` MUST NOT remain the authoritative control source; it MAY be replaced or kept only as a derived/reference view. Security-review ownership and source-row security review MUST remain unchanged.

#### Scenario: JSON owns controls

- GIVEN the review phase needs the corporate controls
- WHEN it loads review control definitions
- THEN it MUST load them from canonical JSON
- AND every `REV-CORP-001..REV-CORP-096` control MUST appear exactly once.

#### Scenario: Markdown catalog is not authority

- GIVEN `control-catalog.md` exists
- WHEN its content differs from canonical JSON
- THEN JSON MUST win
- AND the mismatch MUST be treated as a derived-view repair issue, not a control decision.

### Requirement: Review Artifact Authority Boundary

Persistence and shared post-apply contracts MUST expose both canonical JSON and derived Markdown in OpenSpec mode. Downstream phases MUST consume the selected review artifact identity without treating Markdown as authoritative. The JSON SHOULD be presentation-friendly for future Excel generation, but this change MUST NOT implement Excel, Python, or spreadsheet output.

#### Scenario: Downstream consumes derived Markdown safely

- GIVEN review completed without blockers
- WHEN security review, verify, or archive consumes review evidence
- THEN it MAY read `review-report.md` for compatibility
- AND it MUST recognize canonical JSON as authoritative when facts conflict.

#### Scenario: Excel is deferred

- GIVEN canonical JSON includes presentation metadata
- WHEN this change is implemented
- THEN no Excel/Python generation MUST be added.

## MODIFIED Requirements

### Requirement: Review Report Artifact

`sdd-review` MUST persist canonical `review-report.json` and generate derived `review-report.md` in the same phase. OpenSpec mode MUST expose both under `openspec/changes/{change-name}/`. Missing required artifacts, unknown changed files, unsafe workspace context, JSON validation failure, Markdown generation/parity failure, or persistence failure MUST route to `resolve-blockers`.
(Previously: `sdd-review` only required `review-report.md` as the first-class persisted artifact.)

#### Scenario: Reports are persisted

- GIVEN review can resolve all required inputs
- WHEN review completes
- THEN it MUST write canonical `review-report.json`
- AND it MUST write derived `review-report.md` with verdict, matrix, evidence summary, and next recommendation.

#### Scenario: Markdown generation fails

- GIVEN canonical JSON is valid
- WHEN derived Markdown cannot be generated or read back
- THEN review MUST route to `resolve-blockers`
- AND downstream phases MUST NOT treat the stale Markdown as current evidence.

### Requirement: Code-Review Validation Matrix

The canonical JSON MUST define the review matrix source rows and validation rules. Derived `review-report.md` MUST preserve the exact matrix columns: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. The derived matrix MUST contain 96 rows, use stable `REV-CORP-*` IDs, preserve category without adding columns, and limit `Complies` to `Yes`, `No`, or `N/A`. `N/A` MUST include Evidence Location proving irrelevance and Observations/Comments explaining scope.
(Previously: matrix rows were sourced from a stable control catalog or Item ID prefix.)

#### Scenario: All controls are represented

- GIVEN canonical JSON contains 96 controls
- WHEN review writes JSON and derived Markdown
- THEN every control MUST appear once with a stable Item ID
- AND every Markdown row MUST use the exact required columns.

#### Scenario: Platform control is irrelevant

- GIVEN a control targets an unused platform or technology
- WHEN review marks it `N/A` in JSON
- THEN derived Markdown Evidence Location MUST prove irrelevance
- AND Observations/Comments MUST explain the scope decision.

#### Scenario: Derived Markdown preserves compatibility sections

- GIVEN canonical JSON has report facts
- WHEN Markdown is generated
- THEN it MUST preserve verdict, blocking summary, evidence summary, review matrix, applicable operational evidence, changed-file/security handoff, matrix validation, and recommendation sections.
