# Delta for sdd-security-guideline-catalog

## MODIFIED Requirements

### Requirement: In-Repo Guideline Snapshot

The repository MUST maintain a canonical source-row catalog snapshot for security review. The operational JSON MUST preserve 155 expanded Source IDs, source text or refs, corporate section, PCI alignment, control domain, repo/runtime/data applicability surfaces, evidence expectations, lifecycle vocabulary, owner phase, and route metadata. Compact `SEC-*` identifiers MUST NOT be required for active report validation, navigation, summaries, or grouped `N/A`. (Previously: the catalog preserved compact controls, Source ID inventory, and compact mappings.)

#### Scenario: Catalog snapshot is available

- GIVEN security review needs guideline context
- WHEN it reads the catalog
- THEN it MUST find all 155 source rows and grouping metadata
- AND active report generation MUST NOT require compact `SEC-*` controls.

#### Scenario: Source row vocabulary exists

- GIVEN a source row is expanded
- WHEN validation builds a report row
- THEN required row fields and allowed values MUST be resolvable from catalog vocabulary.

### Requirement: Corporate Source Row Inventory

The canonical JSON catalog MUST own the exhaustive 155-row corporate Source ID inventory. Range notation MUST be expanded before validation. Each row MUST preserve `sourceId`, `corporateSection`, `pciAlignment`, guideline text or refs, `controlDomain`, applicable surfaces, safe-evidence expectations, lifecycle values, owner phase, and route metadata. `controlDomain`, `corporateSection`, or equivalent source-row category fields MUST support report navigation and grouping. (Previously: rows also carried mapped compact `SEC-*` guideline IDs and full-row output was audit-only.)

#### Scenario: Ranges expand before coverage

- GIVEN the snapshot contains a source range
- WHEN catalog coverage is validated
- THEN every concrete Source ID MUST be represented exactly once
- AND missing, duplicate, or unknown Source IDs MUST block.

#### Scenario: Grouping fields are present

- GIVEN a report needs navigation or `N/A` grouping
- WHEN it reads source rows
- THEN `controlDomain`, `corporateSection`, or another source-row category MUST be available.

### Requirement: Safe Source Row Evidence

Source-row evidence MUST be review-safe and row-preserving. Every `N/A` row MUST keep its own non-applicability decision, justification, evidence type/location, finding, owner phase, and route even when a generated report groups equivalent rows. Evidence MUST NOT include secrets, PII, PAN, tokens, connection strings, private keys, or confidential values. (Previously: canonical JSON justified every `N/A` row while compact mappings remained traceable.)

#### Scenario: N/A row is justified

- GIVEN review-security marks a source row `N/A`
- WHEN validation reviews the row
- THEN evidence MUST prove irrelevance
- AND missing row-level justification MUST block.

#### Scenario: Grouping cannot hide gaps

- GIVEN a grouped `N/A` summary is rendered
- WHEN any grouped row lacks equivalent justification
- THEN validation MUST block or split the row from the group.

## REMOVED Requirements

### Requirement: Compact SEC Mapping Coverage

(Reason: Active source-row-first security review no longer uses compact `SEC-*` mappings as validation, navigation, summary, or grouped `N/A` metadata.)
(Migration: Use source-row fields such as `controlDomain` and `corporateSection` for grouping and route decisions.)

### Requirement: Formal Source Coverage Mapping

(Reason: Coverage is now direct Source ID exact-once validation, not coverage through compact `SEC-*` guidelines.)
(Migration: Validate the 155-row Source ID inventory directly.)
