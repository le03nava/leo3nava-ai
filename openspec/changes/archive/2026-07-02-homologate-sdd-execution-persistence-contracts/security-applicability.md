# Security Applicability: Homologate SDD Execution and Persistence Contracts

```yaml
schemaName: sdd.security-applicability
schemaVersion: 1
changeName: homologate-sdd-execution-persistence-contracts
classification: no-impact
securityImpact: false
taxonomyCategories: []
applicableGuidelines: []
evidenceSummary:
  - "Proposal scope is limited to Markdown SDD contracts, phase skill prose, shared persistence wording, and OpenSpec specification artifacts."
  - "Proposal explicitly excludes changing persistence backends, default artifact store behavior, OpenSpec/Engram key formats, non-Markdown runtime code, and installed adapter copies as source of truth."
  - "Specs define execution/persistence contract ownership, artifact routing compatibility, and downstream SDD workflow consumption; they do not define authentication, session, sensitive data, secrets, access control, file handling, database access, or sensitive logging behavior."
  - "OpenSpec config identifies this repository as an AI agent/skill distribution made of Markdown instruction contracts, not an application runtime."
designChangingUnknowns: []
nonBlockingRisks: []
nextRecommended: design
```

## Classification Rationale

This change is classified as `no-impact` because the proposal and specs only change SDD Markdown contracts and OpenSpec documentation for execution/persistence responsibilities. The catalog taxonomy applies when a change touches authentication, sessions, sensitive data/PAN, secrets, permissions/access control, files, database access, or sensitive logging. None of those runtime or security-control areas are modified by this change.

The affected behavior is process/documentation behavior for SDD phase executors: artifact-store mode authority, phase-local artifact contracts, routing compatibility, and downstream consumption of SDD artifacts. The proposal also preserves existing artifact keys, paths, routing tokens, status fields, and backend behavior, and explicitly excludes runtime code and backend redesign. Therefore no catalog guideline is applicable and no security design artifact is required.

## Guideline Mapping

| Guideline ID | Category | Mandatory When Applicable | Evidence |
| --- | --- | --- | --- |
| Not applicable | Not applicable | Not applicable | No catalog taxonomy category applies to this Markdown-only SDD contract/documentation change. |

## No-Impact Evidence

- `proposal.md` limits the scope to `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-phase-common.md`, `skills/sdd-*/SKILL.md`, `skills/_shared/engram-convention.md`, and `openspec/specs/sdd-*` documentation/contract artifacts.
- `proposal.md` lists runtime/backend behavior changes as out of scope: no SDD DAG redesign, no backend changes, no default artifact-store behavior changes, no OpenSpec/Engram key-format changes, and no non-Markdown runtime code.
- `sdd-execution-persistence-contracts/spec.md` defines ownership boundaries for persistence semantics, executor-envelope behavior, phase artifact contracts, and compatibility preservation.
- `sdd-security-applicability-workflow/spec.md`, `sdd-security-design-workflow/spec.md`, and `sdd-test-design-workflow/spec.md` only preserve SDD artifact/routing contracts while delegating common persistence mechanics.
- `openspec/config.yaml` identifies the repository as an AI agent/skill distribution with Markdown instruction contracts, not an application runtime.
- The security catalog taxonomy was checked and no categories apply: authentication, sessions, sensitive-data-pan, secrets, permissions-access-control, files, database-access, and sensitive-logging are not touched.

## Blocking Unknowns

None.

## Security-Design Risks

None.
