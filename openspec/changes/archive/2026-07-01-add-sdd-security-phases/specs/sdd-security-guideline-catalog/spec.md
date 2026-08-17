# sdd-security-guideline-catalog Specification

## Purpose

Define the in-repo corporate security guideline snapshot, compact taxonomy, and evidence model used by security applicability, security design, verification, and archive checks.

## Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain the initial corporate security guideline catalog as an in-repo snapshot based on the user-provided text. The catalog MUST preserve enough source text, identifiers, and version metadata for auditability until an official external versioned source replaces it.

#### Scenario: Catalog snapshot is available

- GIVEN an SDD security phase needs guideline context
- WHEN it reads the catalog
- THEN it MUST find guideline identifiers, source snapshot metadata, and applicable guideline text or summaries
- AND it MUST NOT depend on unavailable external policy documents.

#### Scenario: Catalog source changes later

- GIVEN an official external guideline source becomes available
- WHEN the catalog is updated
- THEN the change MUST preserve migration/audit metadata
- AND existing archived evidence MUST remain understandable.

### Requirement: Compact Security Taxonomy

The catalog MUST expose a compact taxonomy for phase prompts. The taxonomy MUST include authentication, sessions, sensitive data or PAN, secrets, permissions or access control, files, database access, and sensitive logging categories.

#### Scenario: Applicability uses taxonomy

- GIVEN a change modifies session behavior
- WHEN applicability maps the change
- THEN it MUST be able to select the sessions category
- AND it SHOULD avoid injecting unrelated full guideline text.

#### Scenario: Multiple categories apply

- GIVEN a change touches secrets and database access
- WHEN applicability maps guidelines
- THEN it MUST record both categories
- AND security design MUST receive both mappings.

### Requirement: Mandatory Evidence Model

Each catalog guideline MUST declare whether it is mandatory when applicable and what evidence is expected. Mandatory evidence MAY include design controls, implementation references, tests, static checks, manual review, or explicit approved exceptions.

#### Scenario: Mandatory guideline has evidence expectations

- GIVEN a guideline is mandatory when applicable
- WHEN security design consumes it
- THEN the guideline MUST provide expected evidence types
- AND downstream phases MUST be able to verify coverage.

#### Scenario: Exception fields are required

- GIVEN mandatory evidence is unavailable
- WHEN an exception is recorded
- THEN it MUST include approver, guideline, accepted-risk rationale, and mitigation or follow-up
- AND archive MUST reject incomplete exception evidence.
