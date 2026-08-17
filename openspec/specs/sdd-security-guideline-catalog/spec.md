# sdd-security-guideline-catalog Specification

## Purpose

Define the in-repo corporate security guideline source-row snapshot, matrix vocabulary, and evidence model used by narrative secure development design, security review, verification, and archive checks.

## Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain a canonical source-row catalog snapshot for security review. The operational JSON MUST preserve 155 expanded Source IDs, source text or refs, corporate section, PCI alignment, control domain, repo/runtime/data applicability surfaces, evidence expectations, lifecycle vocabulary, owner phase, and route metadata. Compact `SEC-*` identifiers MUST NOT be required for active report validation, navigation, summaries, or grouped `N/A`. The Markdown view MUST preserve audit readability without becoming a second editable source of truth.

#### Scenario: Catalog snapshot is available

- GIVEN an SDD design or security review phase needs guideline context
- WHEN it reads the catalog
- THEN it MUST find all 155 source rows and grouping metadata in the canonical JSON catalog
- AND it MAY use the Markdown view for human/audit readability
- AND active report generation MUST NOT require compact `SEC-*` controls.

#### Scenario: Operational JSON is available

- GIVEN a script or phase needs row expansion, mappings, counts, or export data
- WHEN it reads `skills/sdd-review-security/references/security-guideline-catalog.operational.json`
- THEN it MUST find schema metadata, vocabularies, source sections, expanded Source IDs, source guideline text or refs, PCI alignment, grouping metadata, applicability surfaces, evidence expectations, and expected Source ID count
- AND it MUST be able to generate review-security coverage summaries, full audit matrices, Excel exports, or other documents without parsing Markdown tables.

#### Scenario: Source row vocabulary exists

- GIVEN a source row is expanded
- WHEN validation builds a report row
- THEN required row fields and allowed values MUST be resolvable from catalog vocabulary.

#### Scenario: Catalog source changes later

- GIVEN an official external guideline source becomes available
- WHEN the catalog is updated
- THEN the change MUST preserve migration/audit metadata
- AND existing archived evidence MUST remain understandable.

### Requirement: Compact Security Taxonomy

The catalog MUST expose a compact taxonomy for phase prompts. The taxonomy MUST include authentication, sessions, sensitive data or PAN, secrets, permissions or access control, files, database access, and sensitive logging categories. `sdd-design` MUST use the taxonomy selectively for changed-surface planning; `sdd-review-security` MUST use it exhaustively for applicability validation.

#### Scenario: Security design uses taxonomy

- GIVEN a change modifies session behavior
- WHEN design maps the change
- THEN it MUST be able to select the sessions category
- AND it SHOULD avoid injecting unrelated full guideline text.

#### Scenario: Multiple categories apply

- GIVEN a change touches secrets and database access
- WHEN design maps guidelines
- THEN it MUST record both categories
- AND downstream security review MUST receive both mappings.

### Requirement: Mandatory Evidence Model

Each catalog guideline MUST declare whether it is mandatory when applicable and what evidence is expected across design, implementation, test/static/manual checks, security review, archive, or approved exceptions.

#### Scenario: Mandatory guideline has evidence expectations

- GIVEN a guideline is mandatory when applicable
- WHEN embedded secure development design consumes it
- THEN the guideline MUST provide expected evidence types
- AND security review MUST be able to verify coverage.

#### Scenario: Exception fields are required

- GIVEN mandatory evidence is unavailable
- WHEN an exception is recorded
- THEN it MUST include approver, guideline, accepted-risk rationale, and mitigation or follow-up
- AND archive MUST reject incomplete exception evidence.

### Requirement: Review Control Cross-References

The security guideline catalog MUST support cross-references from `sdd-review` controls and `sdd-review-security` rows to applicable guideline identifiers. These references MUST help reviewers cite standards without transferring guideline authority away from embedded secure development design and security review.

#### Scenario: Review cites a security guideline

- GIVEN a review checklist control evaluates a security concern
- WHEN the control maps to a catalog guideline
- THEN the review control SHOULD cite that guideline identifier
- AND embedded design/review-security MUST remain authoritative for required controls and final security verdict.

### Requirement: Review-Safe Security Evidence

The catalog SHOULD identify evidence types suitable for review rows, including implementation reference, static inspection, test evidence, approved exception, or N/A evidence. N/A MUST include evidence proving the category, platform, framework, API, or data class is irrelevant.

#### Scenario: Security row is N/A

- GIVEN a guideline row is marked N/A
- WHEN security review validates it
- THEN evidence MUST show why it is out of scope
- AND observations MUST explain why no security design control is required.

### Requirement: Catalog Boundary Preservation

The catalog MUST remain the source for security source rows, taxonomy, mandatory evidence expectations, exception fields, lifecycle statuses, matrix vocabulary, Source ID inventory, and safe-evidence policy. `sdd-review` MAY reference catalog entries but MUST NOT duplicate, redefine, or replace guideline text. Active authority MUST remain split between narrative design rules and exhaustive canonical `review-security-report.json` validation. The catalog MUST NOT require design YAML, schema, matrices, exhaustive applicability, or `N/A` decisions outside canonical review-security evidence.

#### Scenario: Catalog authority is preserved

- GIVEN embedded design and security review reference a guideline
- WHEN downstream verification compares evidence
- THEN catalog identifiers and statuses MUST remain consistent
- AND conflicts MUST resolve through narrative design context plus exhaustive review-security results.

### Requirement: Operational Severity Vocabulary

The catalog MUST define operational applicability severity with only `blocking`, `conditional`, and `advisory`. `blocking` obligations MUST prevent phase success when unresolved, `conditional` obligations MUST apply when stated predicates are true, and `advisory` obligations SHOULD be preserved as downstream risk or guidance. The catalog MUST NOT use review labels such as `Menor`, `Media`, or `Mayor` for applicability blocking behavior.

#### Scenario: Blocking obligation prevents success

- GIVEN an applicable guideline has unresolved `blocking` evidence
- WHEN embedded design or security review evaluates completion
- THEN the phase MUST block
- AND the blocker MUST name the guideline and missing evidence.

#### Scenario: Conditional obligation predicate is false

- GIVEN a guideline is `conditional` and its predicate is not met
- WHEN embedded design evaluates the guideline
- THEN review-security MAY mark it not applicable
- AND it MUST record the predicate rationale.

### Requirement: Catalog Validator Contract

The catalog MUST support static validation of guideline IDs, taxonomy categories, Source IDs, severity values, lifecycle statuses, matrix values, mandatory evidence fields, and exception fields used by narrative `design.md` secure development rules and canonical `review-security-report.json`. Validation MUST compare artifact references against the same catalog snapshot identity recorded in the artifact.

#### Scenario: Artifact references current catalog snapshot

- GIVEN a security artifact records a catalog snapshot identity
- WHEN static validation checks guideline and source references
- THEN references MUST resolve within that snapshot
- AND mismatched or unknown references MUST fail validation.

#### Scenario: Advisory evidence is preserved

- GIVEN an advisory guideline is applicable
- WHEN design records it
- THEN downstream phases SHOULD preserve it as risk or guidance
- AND archive evidence MUST remain understandable even if it does not block.

### Requirement: Security Matrix Vocabulary

The catalog MUST define matrix-facing values `Yes`, `No`, and `N/A` plus lifecycle statuses `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, and `blocked` for review-security rows.

#### Scenario: Matrix row uses valid vocabulary

- GIVEN security review writes a matrix row
- WHEN validation runs
- THEN its answer and lifecycle status MUST use catalog vocabulary
- AND unsupported values MUST fail validation.

### Requirement: Corporate Source Row Inventory

The canonical JSON catalog MUST own the exhaustive 155-row corporate Source ID inventory. Range notation MUST be expanded before validation. Each row MUST preserve `sourceId`, `corporateSection`, `pciAlignment`, guideline text or refs, `controlDomain`, applicable surfaces, safe-evidence expectations, lifecycle values, owner phase, and route metadata. `controlDomain`, `corporateSection`, or equivalent source-row category fields MUST support report navigation and grouping.

#### Scenario: Ranges expand before coverage

- GIVEN the snapshot contains a source range
- WHEN catalog coverage is validated
- THEN every concrete Source ID in the range MUST be represented exactly once
- AND missing, duplicate, or unknown Source IDs MUST block.

#### Scenario: Grouping fields are present

- GIVEN a report needs navigation or `N/A` grouping
- WHEN it reads source rows
- THEN `controlDomain`, `corporateSection`, or another source-row category MUST be available.

#### Scenario: PCI alignment is preserved

- GIVEN a corporate section heading includes PCI alignment
- WHEN source rows are produced
- THEN each row from that section MUST preserve the PCI requirement
- AND rows without alignment MUST state `N/A`.

#### Scenario: Selective artifacts reference catalog inventory

- GIVEN design or test-design needs source-row context
- WHEN it records coverage obligations
- THEN it MUST cite catalog snapshot/context and applicable category guidance instead of copying all 155 rows
- AND review-security MUST use the catalog to expand the exhaustive validation universe.

### Requirement: Safe Source Row Evidence

Source-row evidence MUST be review-safe and row-preserving. Every `N/A` row MUST keep its own non-applicability decision, justification, evidence type/location, finding, owner phase, and route even when a generated report groups equivalent rows. Evidence MUST NOT include secrets, PII, PAN, tokens, connection strings, private keys, or confidential values.

#### Scenario: Applicable row lacks safe evidence

- GIVEN a source row applies
- WHEN evidence is missing or unsafe
- THEN the row MUST be non-compliant and blocking
- AND the report MUST identify the unsafe or absent evidence.

#### Scenario: N/A row is justified

- GIVEN review-security marks a source row `N/A`
- WHEN validation reviews the row
- THEN evidence MUST prove irrelevance
- AND missing justification MUST block.

#### Scenario: Grouping cannot hide gaps

- GIVEN a grouped `N/A` summary is rendered
- WHEN any grouped row lacks equivalent justification
- THEN validation MUST block or split the row from the group.

### Requirement: Shared Security Contract Source Row Schema

`skills/_shared/sdd-security-contract.md` MUST define the review-security report schema, allowed values, traceability, artifact boundary, narrative design obligations, exhaustive review-security obligations, safe-evidence policy, and routing. It MUST state that catalog owns inventory, design owns narrative category rules, test-design owns planned checks, review-security owns validation and coverage reporting, full matrix output is audit-only unless requested, and verify/archive preserve evidence without design YAML.

#### Scenario: Contract consumers share schema

- GIVEN a phase writes or validates source rows
- WHEN it reads the shared security contract
- THEN it MUST use the same row fields, allowed values, and artifact ownership boundary
- AND backend-specific storage MUST follow the persistence contract.

#### Scenario: Routing semantics are defined

- GIVEN source-row validation finds blockers or warnings
- WHEN a phase computes next routing
- THEN implementation blockers MUST route to apply, contract/artifact/N/A evidence blockers MUST route to resolve-blockers, and warnings-only results MAY route forward.

### Requirement: Operational Safe-Evidence Policy

The security guideline catalog MUST define safe-evidence policy for operational considerations when they apply. It SHOULD cover applicable logs/errors, monitoring mechanisms, administration operations, reprocessing/recovery, backup/retention/cleanup/generated artifacts, unresolved gaps, and final documentation boundaries. Monitoring evidence MUST be mechanism-oriented and MUST NOT be limited to SQL-only checks. The policy MUST NOT imply mandatory operational categories for every change.

#### Scenario: Monitoring evidence is categorized

- GIVEN a change requires operational monitoring evidence
- WHEN the catalog is consulted
- THEN it MUST support monitoring mechanisms such as dashboards, alerts, jobs, traces, scripts, or documented manual checks
- AND it MUST NOT require SQL-only evidence.

#### Scenario: Operational gap is safe

- GIVEN operational evidence is unavailable
- WHEN catalog policy is applied
- THEN exact `Pendiente de confirmar:` MUST be an accepted safe state for ordinary SDD evidence.

#### Scenario: Operational category is not applicable

- GIVEN design states an operational concern does not apply
- WHEN catalog policy is applied
- THEN exact `No aplica.` with optional rationale MUST satisfy the safe-evidence state.

### Requirement: Restricted Operational Data Classification

The catalog MUST classify production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, generated file bytes, and equivalent environment-specific operational details as restricted for ordinary SDD evidence, code, tests, fixtures, and examples. Final operational documentation MAY include user-provided restricted operational values. This classification MUST protect present operational evidence without creating mandatory readiness completeness gates.

#### Scenario: Restricted data appears in evidence

- GIVEN an SDD artifact cites restricted operational data
- WHEN catalog-based review-security validation runs
- THEN the evidence MUST be unsafe unless it is scoped to final user-provided operational documentation.

#### Scenario: Safe evidence is enough

- GIVEN a field is inapplicable
- WHEN catalog policy is applied
- THEN exact `No aplica.` with optional rationale MUST satisfy the safe-evidence state.
