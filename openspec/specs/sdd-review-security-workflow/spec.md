# sdd-review-security-workflow Specification

## Purpose

Define the mandatory post-review security evidence gate that parses narrative secure design and validates the 155 corporate Source ID rows as the only active security-review matrix before verification.

## Requirements

### Requirement: Mandatory Security Review Gate

The SDD workflow MUST run `sdd-review-security` after non-blocking `sdd-review` and before `sdd-verify` for every new change.

#### Scenario: General review routes to security review

- GIVEN canonical `review-report.json` has no blocking findings and derived `review-report.md` is readable or not required by the selected backend
- WHEN routing is computed
- THEN the next required phase MUST be `sdd-review-security`
- AND `sdd-verify` MUST remain blocked until security review evidence exists.

#### Scenario: General review is blocking

- GIVEN `sdd-review` reports blocking findings
- WHEN routing is computed
- THEN `sdd-review-security` MUST NOT run
- AND the workflow MUST route back to `sdd-apply` or `resolve-blockers`.

### Requirement: Security Review Artifact

`sdd-review-security` MUST persist canonical `review-security-report.json` as the only source of truth and MUST generate derived Markdown from that JSON. The JSON MUST expose `sourceRowValidation.rows` with exactly 155 unique Source ID rows, exact-once coverage, verdict, routing, source refs, blockers, warnings, unsafe evidence rejections, warning carry-forward, and parity/read-back metadata. The active report contract MUST NOT expose compact `SEC-*` controls for validation, navigation, summaries, grouped `N/A`, or metadata.

#### Scenario: Report is persisted

- GIVEN `design.md` and canonical `review-report.json` are readable
- WHEN security review completes
- THEN `review-security-report.json` MUST be written and read back first
- AND derived `review-security-report.md` MUST be generated from JSON, written, read back, and parity-checked
- AND both artifacts MUST state the same verdict, with JSON authoritative on conflict.

#### Scenario: Report is source-row authoritative

- GIVEN required review inputs are readable
- WHEN security review completes
- THEN canonical JSON MUST contain exactly 155 unique `sourceRowValidation.rows`
- AND derived Markdown MUST be generated from JSON and parity-checked.

#### Scenario: Compact controls are absent

- GIVEN an active security report is produced
- WHEN validation inspects validation, navigation, summaries, and `N/A` groups
- THEN no `SEC-*` compact-control contract data MUST be present.

#### Scenario: Embedded secure design is required

- GIVEN a new change lacks `design.md#secure-development-design`
- WHEN security review evaluates readiness
- THEN it MUST block with missing embedded design evidence
- AND verify/archive MUST remain unavailable.

### Requirement: Security Matrix Validation

Security review MUST validate every corporate source row exactly once using row fields `sourceId`, `corporateSection`, `pciAlignment`, `guidelineText` or `guidelineRefs`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, and `route`. Grouped `N/A` MAY be rendered only by source-row grouping fields and only when equivalent non-applicability remains preserved per row.

#### Scenario: Mandatory evidence is missing

- GIVEN an applicable mandatory guideline lacks implementation evidence or approved exception
- WHEN security review validates the matrix
- THEN the report MUST mark the row `No` and `blocked`
- AND verify/archive MUST be blocked.

#### Scenario: Exact-once row validation

- GIVEN the catalog defines 155 Source IDs
- WHEN JSON validation runs
- THEN missing, duplicate, or unknown Source IDs MUST block
- AND every row MUST include the required field set.

#### Scenario: Not applicable row is justified

- GIVEN security review marks a guideline or Source ID `N/A`
- WHEN it validates the row
- THEN evidence MUST prove irrelevance
- AND observations MUST explain the scope decision.

#### Scenario: Grouped N/A preserves rows

- GIVEN multiple rows share equivalent non-applicability
- WHEN the report groups them for humans
- THEN grouping MUST use `controlDomain`, `corporateSection`, or another source-row category
- AND each row MUST retain its own `applies`, `justification`, and evidence fields.

#### Scenario: Design omitted an applicable control

- GIVEN proposal, specs, changed files, or evidence show a category applies
- WHEN design did not include that category/control
- THEN security review MUST report a missed applicable control
- AND the verdict MUST be blocking.

### Requirement: Boundary with General Review

Security review MUST NOT replace `sdd-review` or duplicate the 96-control matrix. It MUST focus on security guideline evidence and MAY cite general review findings as supporting evidence.

#### Scenario: General review evidence is reused

- GIVEN a review row supports a security guideline
- WHEN security review records evidence
- THEN it MAY cite that review row
- AND it MUST keep the security verdict in canonical `review-security-report.json`.

### Requirement: Exhaustive Source Row Security Review

`sdd-review-security` MUST render lean Markdown from canonical JSON with verdict, handoff, navigation/summary by source-row categories, grouped `N/A` summaries, blockers/warnings, and the full 155-row matrix at the end. Markdown MUST NOT duplicate validation logic or reintroduce compact-control sections.

#### Scenario: Lean generated Markdown

- GIVEN canonical JSON validates successfully
- WHEN Markdown is rendered
- THEN it MUST contain summary/navigation first and the full source-row matrix last
- AND it MUST NOT contain `## Compact Control Validation`.

#### Scenario: JSON wins

- GIVEN Markdown and JSON disagree
- WHEN downstream consumers resolve security evidence
- THEN canonical JSON MUST be authoritative.

### Requirement: Source Row Evidence Correlation

Security review MUST correlate source rows with narrative design rules, test-design checks, apply evidence, changed files, and review findings. A row MUST NOT pass solely because it is listed or omitted from design; evidence MUST support applicability, compliance, justified `N/A`, or approved exception.

#### Scenario: Listed row has no corroboration

- GIVEN a source row appears in the matrix
- WHEN no supporting design, test, apply, changed-file, review, or exception evidence exists
- THEN the row MUST fail validation
- AND the report MUST identify the missing corroboration.

#### Scenario: Evidence is unsafe

- GIVEN evidence contains secret-like or confidential values
- WHEN security review validates the row
- THEN it MUST reject the evidence as unsafe
- AND route to `resolve-blockers` unless implementation remediation is required.

### Requirement: Operational Evidence Leakage Review

`sdd-review-security` MUST validate leakage boundaries for operational evidence that exists in design, test-design, tasks, apply evidence, review output, code, tests, fixtures, examples, or archived evidence. It MUST focus on restricted production identifiers, secrets, payloads, full ID lists, and generated bytes. It MUST NOT require operational-readiness category completeness when no applicable evidence exists.

#### Scenario: Restricted production identifier is found

- GIVEN ordinary SDD evidence contains a production hostname, IP, port, SID/service name, credential, token, payload, full ID list, or generated file bytes
- WHEN security review evaluates operational evidence
- THEN it MUST report unsafe evidence as blocking
- AND route to `resolve-blockers` unless implementation remediation is required.

#### Scenario: Operational document is the target

- GIVEN the user explicitly provides production operational values for final documentation
- WHEN security review evaluates ordinary SDD artifacts
- THEN those values MUST remain outside SDD evidence, code, tests, fixtures, and examples.

#### Scenario: No operational evidence exists

- GIVEN design and review evidence show no operational considerations apply
- WHEN security review runs
- THEN it MUST NOT block for missing readiness categories.

### Requirement: Safe Placeholder Security Boundary

Security review MUST accept exact `Pendiente de confirmar:` and exact `No aplica.` placeholders as safer than invented operational details for applicable operational evidence. It MUST NOT require disclosure of real operational data to pass. Placeholders MUST NOT be used to hide required security proof when non-leakage evidence is applicable.

#### Scenario: Missing value uses safe placeholder

- GIVEN operational data is unavailable
- WHEN security review validates the evidence boundary
- THEN exact placeholder usage MUST NOT be treated as a leakage failure.

#### Scenario: Placeholder hides required security evidence

- GIVEN a security obligation still requires proof of non-leakage
- WHEN only a placeholder is present
- THEN security review MUST require safe evidence of the boundary, not real operational values.
