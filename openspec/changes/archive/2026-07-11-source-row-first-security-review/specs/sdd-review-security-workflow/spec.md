# Delta for sdd-review-security-workflow

## MODIFIED Requirements

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist canonical `review-security-report.json` as the only source of truth and MUST generate derived Markdown from that JSON. The JSON MUST expose `sourceRowValidation.rows` with exactly 155 unique Source ID rows, exact-once coverage, verdict, routing, source refs, blockers, warnings, unsafe evidence rejections, warning carry-forward, and parity/read-back metadata. The active report contract MUST NOT expose compact `SEC-*` controls for validation, navigation, summaries, grouped `N/A`, or metadata. (Previously: JSON owned compact and Source ID validation, with full source rows optional.)

#### Scenario: Report is source-row authoritative

- GIVEN required review inputs are readable
- WHEN security review completes
- THEN canonical JSON MUST contain exactly 155 unique `sourceRowValidation.rows`
- AND derived Markdown MUST be generated from JSON and parity-checked.

#### Scenario: Compact controls are absent

- GIVEN an active security report is produced
- WHEN validation inspects validation, navigation, summaries, and `N/A` groups
- THEN no `SEC-*` compact-control contract data MUST be present.

### Requirement: Security Matrix Validation

Security review MUST validate every corporate source row exactly once using row fields `sourceId`, `corporateSection`, `pciAlignment`, `guidelineText` or `guidelineRefs`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, and `route`. Grouped `N/A` MAY be rendered only by source-row grouping fields and only when equivalent non-applicability remains preserved per row. (Previously: validation expanded compact controls and Source IDs.)

#### Scenario: Exact-once row validation

- GIVEN the catalog defines 155 Source IDs
- WHEN JSON validation runs
- THEN missing, duplicate, or unknown Source IDs MUST block
- AND every row MUST include the required field set.

#### Scenario: Grouped N/A preserves rows

- GIVEN multiple rows share equivalent non-applicability
- WHEN the report groups them for humans
- THEN grouping MUST use `controlDomain`, `corporateSection`, or another source-row category
- AND each row MUST retain its own `applies`, `justification`, and evidence fields.

### Requirement: Exhaustive Source Row Security Review

`sdd-review-security` MUST render lean Markdown from canonical JSON with verdict, handoff, navigation/summary by source-row categories, grouped `N/A` summaries, blockers/warnings, and the full 155-row matrix at the end. Markdown MUST NOT duplicate validation logic or reintroduce compact-control sections. (Previously: full matrix materialization was audit-only and compact/source summaries were rendered.)

#### Scenario: Lean generated Markdown

- GIVEN canonical JSON validates successfully
- WHEN Markdown is rendered
- THEN it MUST contain summary/navigation first and the full source-row matrix last
- AND it MUST NOT contain `## Compact Control Validation`.

#### Scenario: JSON wins

- GIVEN Markdown and JSON disagree
- WHEN downstream consumers resolve security evidence
- THEN canonical JSON MUST be authoritative.

## REMOVED Requirements

### Requirement: Source Row Blocking Rules

(Reason: Replaced by source-row-first validation where blockers are computed from exact-once row coverage, required row fields, row evidence, and route.)
(Migration: Express blocking behavior under the modified row validation/report artifact requirements.)
