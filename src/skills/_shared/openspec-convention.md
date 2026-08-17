# OpenSpec File Convention (backend reference only)

NOTE: This document is a backend-specific reference for OpenSpec paths, directory layout, and file-format details. The authoritative cross-mode persistence rules live in `skills/_shared/persistence-contract.md`; phase skills MUST follow that contract for mode resolution, read/write behavior, and persistence verification.

Use this file only when an OpenSpec-specific path, folder shape, or archive-file convention is needed. Do not treat it as a competing source for artifact-store mode behavior.

## Directory Structure

```text
openspec/
├── config.yaml              <- Project-specific SDD config
├── specs/                   <- Source of truth (main specs)
│   └── {domain}/
│       └── spec.md
└── changes/                 <- Active changes
    ├── archive/             <- Completed changes (YYYY-MM-DD-{change-name}/)
    └── {change-name}/       <- Active change folder
        ├── state.yaml       <- DAG state (survives compaction)
        ├── explore.md       <- (optional) from sdd-explore
        ├── proposal.md      <- from sdd-propose
        ├── specs/           <- from sdd-spec
        │   └── {domain}/
        │       └── spec.md  <- Delta spec
        ├── security-applicability.md <- legacy/read-only; old or archived changes only
        ├── design.md        <- from sdd-design
        ├── security-design.md <- legacy/read-only; old or archived changes only
        ├── test-design.md   <- from sdd-test-design
        ├── tasks.md         <- from sdd-tasks (updated by sdd-apply)
        ├── review-report.json <- canonical from sdd-review
        ├── review-report.md <- derived compatibility view from sdd-review
        ├── review-security-report.json <- canonical from sdd-review-security
        ├── review-security-report.md <- derived compatibility view from sdd-review-security
        └── verify-report.md <- from sdd-verify
```

## Artifact File Paths

| Skill | Creates / Reads | Path |
| --- | --- | --- |
| orchestrator | Creates/Updates | `openspec/changes/{change-name}/state.yaml` |
| sdd-init | Creates | `openspec/config.yaml`, `openspec/specs/`, `openspec/changes/`, `openspec/changes/archive/` |
| sdd-explore | Creates (optional) | `openspec/changes/{change-name}/explore.md` |
| sdd-propose | Creates | `openspec/changes/{change-name}/proposal.md` |
| sdd-spec | Creates | `openspec/changes/{change-name}/specs/{domain}/spec.md` |
| Legacy security applicability data | Read-only archive compatibility only; no launchable phase | `openspec/changes/{change-name}/security-applicability.md` |
| sdd-design | Creates | `openspec/changes/{change-name}/design.md` |
| Legacy security design data | Read-only archive compatibility only; no launchable phase for new changes | `openspec/changes/{change-name}/security-design.md` |
| sdd-test-design | Creates | `openspec/changes/{change-name}/test-design.md` |
| sdd-tasks | Creates | `openspec/changes/{change-name}/tasks.md` |
| sdd-apply | Updates | `openspec/changes/{change-name}/tasks.md` (marks `[x]`) |
| sdd-review | Creates | `openspec/changes/{change-name}/review-report.json` canonical plus `openspec/changes/{change-name}/review-report.md` derived |
| sdd-review-security | Creates | `openspec/changes/{change-name}/review-security-report.json` canonical plus `openspec/changes/{change-name}/review-security-report.md` derived |
| sdd-verify | Creates | `openspec/changes/{change-name}/verify-report.md` |
| sdd-archive | Moves | `openspec/changes/{change-name}/` → `openspec/changes/archive/YYYY-MM-DD-{change-name}/` |
| sdd-archive | Updates | `openspec/specs/{domain}/spec.md` (merges deltas into main specs, including review workflow deltas such as `sdd-review-workflow`, `sdd-execution-persistence-contracts`, and `sdd-security-guideline-catalog`) |

## Reading Artifacts

```text
Proposal:   openspec/changes/{change-name}/proposal.md
Specs:      openspec/changes/{change-name}/specs/  (all domain subdirectories)
Security applicability: openspec/changes/{change-name}/security-applicability.md (legacy/read-only for old or archived changes)
Design:     openspec/changes/{change-name}/design.md
Security design: openspec/changes/{change-name}/design.md#secure-development-design (new changes); openspec/changes/{change-name}/security-design.md is legacy/read-only for old or archived changes
Test design: openspec/changes/{change-name}/test-design.md
Tasks:      openspec/changes/{change-name}/tasks.md
Review:     openspec/changes/{change-name}/review-report.json (canonical), openspec/changes/{change-name}/review-report.md (derived compatibility)
Security review: openspec/changes/{change-name}/review-security-report.json (canonical), openspec/changes/{change-name}/review-security-report.md (derived compatibility)
Verify:     openspec/changes/{change-name}/verify-report.md
Config:     openspec/config.yaml
Main specs: openspec/specs/{domain}/spec.md
```

## Writing Rules

- Always create the change directory before writing artifacts
- If a file already exists, READ it first and UPDATE it (don't overwrite blindly)
- If the change directory already exists with artifacts, the change is being CONTINUED
- Use `openspec/config.yaml` `rules` section for project-specific constraints per phase

## Delta Spec Sections

Delta specs MAY include these sections:

```markdown
## ADDED Requirements
## MODIFIED Requirements
## REMOVED Requirements
## RENAMED Requirements
```

- `ADDED` appends new requirements to the main spec.
- `MODIFIED` replaces the full matching requirement block in the main spec. The delta MUST contain the entire updated requirement, including unchanged scenarios that must be preserved.
- `REMOVED` deletes the matching requirement from the main spec. Each removed requirement MUST include `(Reason: ...)` and SHOULD include `(Migration: ...)` when consumers or persisted behavior are affected.
- `RENAMED` changes a requirement heading/name without changing behavior unless the delta also includes a `MODIFIED` block for the new requirement. Each rename MUST state old and new names explicitly.

## Config File Reference

```yaml
# openspec/config.yaml
schema: spec-driven

context: |
  Tech stack: {detected}
  Architecture: {detected}
  Testing: {detected}
  Style: {detected}

rules:
  proposal:
    - Include rollback plan for risky changes
  specs:
    - Use Given/When/Then for scenarios
    - Use RFC 2119 keywords (MUST, SHALL, SHOULD, MAY)
  design:
    - Include sequence diagrams for complex flows
    - Document architecture decisions with rationale
  tasks:
    - Group by phase, use hierarchical numbering
    - Keep tasks completable in one session
  apply:
    - Follow existing code patterns
    tdd: false           # Set to true to enable RED-GREEN-REFACTOR
    test_command: ""
  verify:
    test_command: ""
    build_command: ""
    coverage_threshold: 0
  archive:
    - Warn before merging destructive deltas
```

## Archive Structure

When archiving, the change folder moves to:

```text
openspec/changes/archive/YYYY-MM-DD-{change-name}/
```

Use today's date in ISO format. The archive is an AUDIT TRAIL — never delete or modify archived changes.

Archived change folders preserve every produced phase artifact. New-change archive readiness is defined by `skills/_shared/sdd-post-apply-gates.md#archive-readiness`; OpenSpec archives preserve the artifacts and explicit exceptions required by that contract. Legacy archives may include `security-applicability.md` or standalone `security-design.md` as historical read-only evidence; new changes must not create or require those standalone artifacts.

Archive sync behavior is owned by `skills/sdd-archive/SKILL.md`. Relevant OpenSpec main-spec target paths include `openspec/specs/{domain}/spec.md`; known SDD workflow domains include `openspec/specs/sdd-review-workflow/spec.md`, `openspec/specs/sdd-execution-persistence-contracts/spec.md`, and `openspec/specs/sdd-security-guideline-catalog/spec.md` when matching deltas exist under `openspec/changes/{change-name}/specs/`.
