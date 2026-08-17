# Security Applicability: Add SDD Review Phase

```yaml
schemaName: sdd.security-applicability
schemaVersion: 1
changeName: add-sdd-review-phase
classification: no-impact
securityImpact: false
taxonomyCategories: []
applicableGuidelines: []
evidenceSummary:
  - Proposal scope adds an SDD review workflow, review-report artifact, DAG routing, and shared phase contracts; it explicitly excludes runtime application tests/analyzers and replacement of security applicability/design ownership.
  - Specs add `sdd-review` as a mandatory post-apply/pre-verify process gate and require review rows to cross-reference security guideline identifiers where applicable.
  - The security catalog delta changes evidence/catalog workflow behavior only: cross-references, review-safe evidence types, and ownership boundary preservation.
  - The OpenSpec project context identifies this repository as AI agent/skill instruction contracts, not an application runtime.
  - No requirement changes authentication, sessions, sensitive data/PAN handling, secrets, permissions/access control, files, database access, or sensitive logging runtime behavior.
designChangingUnknowns: []
nonBlockingRisks:
  - Review implementation must preserve the security ownership boundary: `sdd-review` may cite catalog IDs, but security applicability/design remain authoritative for applicability, required controls, and exceptions.
  - Review control mapping must use existing catalog identifiers and must not duplicate or redefine security guideline text.
nextRecommended: design
```

## Classification Rationale

This change is classified as `no-impact` under the local SDD security contract because it does not touch any design-changing security taxonomy area: `authentication`, `sessions`, `sensitive-data-pan`, `secrets`, `permissions-access-control`, `files`, `database-access`, or `sensitive-logging`.

The change is security-relevant as workflow governance because it modifies how the future review phase references the security guideline catalog. However, the catalog contract only makes a change `security-impacting` when an applicable taxonomy category and existing guideline ID are in scope. Here, the new behavior preserves security authority boundaries instead of changing runtime security controls or assigning new mandatory security controls.

## Guideline Mapping

| Guideline ID | Category | Mandatory When Applicable | Evidence |
| --- | --- | --- | --- |
| Not applicable | Not applicable | Not applicable | No catalog taxonomy category applies to this process-only SDD workflow change. Review controls will reference applicable guideline IDs during future reviews, but this change does not itself implement or modify those security controls. |

## No-Impact Evidence

- Proposal lines 17-20 keep application runtime tests/analyzers out of scope and state that review may reference security controls but does not replace security applicability/design ownership.
- Proposal lines 24-33 describe new/modified SDD workflow capabilities, shared contracts, and review artifact handling rather than runtime auth, session, data, file, database, or logging behavior.
- `sdd-review-workflow` requires security-related review controls to cross-reference the catalog while preserving security applicability/design as the authorities.
- `sdd-execution-persistence-contracts` only adds `review-report.md`, DAG routing, state/status fields, and downstream consumption rules.
- `sdd-security-guideline-catalog` only adds review cross-reference support, review-safe evidence types, and catalog boundary preservation.
- `openspec/config.yaml` describes this repository as Markdown instruction contracts for SDD agents and skills, not an application runtime.

## Blocking Unknowns

None.

## Security-Design Risks

- The technical design must avoid moving security applicability, security design, or exception authority into `sdd-review`.
- The technical design must ensure review controls cite existing catalog IDs without duplicating or redefining catalog guideline text.
