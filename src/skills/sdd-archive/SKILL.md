---
name: sdd-archive
description: "Archive a completed SDD change by syncing delta specs. Trigger: orchestrator launches archive after implementation and verification."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-archive` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for ARCHIVING. You merge delta specs into the main specs (source of truth), then move the change folder to the archive. You complete the SDD cycle.

In `engram` mode, archive means lineage and closure report only; it does not promote specs into a shared filesystem source-of-truth tree.

## What You Receive

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads, spec sync, and archive
- `launch.status.dependencies` — `verify: completed` must be present with a passing verdict
- `launch.artifacts.paths` or `launch.artifacts.refs` — refs for all change artifacts (proposal, spec, design, test-design, tasks, review-report, review-security-report, verify-report); do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute workspace root for filesystem operations
- `launch.actionContext.allowedEditRoots` — must include `openspec/specs/` and `openspec/changes/`; return `blocked` if archive paths fall outside these roots
- `launch.skill_paths` — supplemental skills to load before work

The `## Phase Context` section contains:
- Any explicit intentional archive override or stale-checkbox reconciliation approval from the user/orchestrator

If `launch.artifact_store.mode` is absent, `launch.status.dependencies` does not show `verify: completed`, or `launch.actionContext.allowedEditRoots` is missing in `openspec` mode, return `blocked` with `next_recommended: resolve-blockers`.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`. Common archive readiness, review/security-review consumption, task-completion gates, safe-evidence rules, and routing defaults are defined in `skills/_shared/sdd-post-apply-gates.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal, specs, design with mandatory `## Secure Development Design`, `test-design`, tasks, non-blocking general review evidence from the selected backend (canonical `review-report.json` when present plus derived `review-report.md` / `sdd/{change-name}/review` compatibility view), non-blocking security-review evidence from the selected backend (canonical `review-security-report.json` when present plus derived `review-security-report.md` / `sdd/{change-name}/review-security` compatibility view), and passing `verify-report`. Standalone `security-design.md` is legacy/read-only compatibility data only. |
| Produced artifact | Archive report as `sdd/{change-name}/archive-report`; in OpenSpec, the archive audit-trail reference is `openspec/changes/archive/YYYY-MM-DD-{change-name}/`. |
| Mutates | OpenSpec source specs under `openspec/specs/{domain}/spec.md`; OpenSpec change folder location from active change to dated archive; Engram archive report lineage. |
| Spec sync semantics | Merge delta specs before moving the change folder. Preserve unrelated requirements; create missing main specs from full new specs; require explicit reason/migration for removals and explicit old/new names for renames. |
| Archive move semantics | Move the entire change folder to the dated archive destination, never overwrite an existing archive folder, and verify the active change folder is gone and archived contents are complete. |
| Destructive-delta warnings | Stop before destructive merges, large removals, unresolved removals, or ambiguous renames; return `confirmation_required: destructive-merge` for orchestrator-owned confirmation. |
| Audit-trail semantics | Record artifact refs/observation IDs or concrete paths, synced domains and counts, task completion status, canonical general-review JSON ref when present, derived Markdown compatibility ref, general review verdict/blocking state, canonical security-review JSON ref when present, derived security Markdown compatibility ref, security review verdict/blocking state, verify verdict, embedded secure-design validation metadata, secure design/control evidence and N/A rationale, archive destination, warnings, and any approved reconciliation. Canonical general-review and security-review JSON artifacts are authoritative over derived Markdown on conflict. |
| Source-row preservation | When corporate source-row validation applies, preserve canonical JSON refs before derived Markdown refs, source-row coverage summary, catalog snapshot identity/path, expected and validated Source ID counts, exact-once coverage status, non-blocking warnings, complete exceptions, safe evidence references, `N/A` evidence/justification status, parity metadata, and review-security/verify verdict links. Archive MUST NOT copy the full review-security source-row matrix into archive summaries. |
| Operational evidence preservation | Preserve operational evidence/status, refs, exact `Pendiente de confirmar:` gaps, exact `No aplica.` states, warning carry-forward, unavailable-tooling notes, exceptions, and manual operational document handoff boundaries when present. Archive MUST NOT require absent readiness categories or `sdd-operational-doc` execution. |
| Conditional behavior | Engram mode records lineage and closure without filesystem promotion. |
| Success routing | `next_recommended: none` after archive report persistence and selected-backend read-back verification succeed. |
| Block routing | `next_recommended: review`, `review-security`, `verify`, `apply`, or `resolve-blockers` according to missing/blocking review evidence, missing verify evidence, unchecked tasks, missing embedded secure design, unsafe context, destructive merge, destination conflict, or persistence failure. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the full archive summary in `detailed_report`.

## Routing and Readiness Contract

- Apply `skills/_shared/sdd-post-apply-gates.md#archive-readiness` before syncing specs, moving folders, or claiming SDD cycle completion.
- Successful archive -> return `next_recommended: none`.
- Archive readiness failures route according to the shared Archive Readiness table.
- Archive-specific filesystem/spec-sync blockers route to `resolve-blockers`, except implementation task incompleteness routes to `apply` as defined by the shared archive readiness rules.
- Status `partial` after filesystem operations -> return `next_recommended: resolve-blockers` and include exact recovery steps in `detailed_report`.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Archive-Specific Decision Gates

| Condition | Action |
|---|---|
| Delta spec removal lacks `(Reason: ...)` or `(Migration: ...)` | Return `blocked` with `next_recommended: resolve-blockers`; do not delete from main specs. |
| Delta spec rename lacks explicit old and new requirement names | Return `blocked` with `next_recommended: resolve-blockers`; do not rename in main specs. |
| Merge would be destructive or remove large sections | Return `blocked` with `next_recommended: resolve-blockers` and `confirmation_required: destructive-merge`; the orchestrator owns confirmation. |
| Archive destination already exists | Return `blocked` with `next_recommended: resolve-blockers` unless the orchestrator provides an explicit alternate destination. |
| Archive operation would leave `allowedEditRoots` | Return `blocked` with `next_recommended: resolve-blockers` and report the offending path. |
| Archive verification fails after filesystem operations | Return `partial` with `next_recommended: resolve-blockers` and exact failed checks and recovery paths. |
| Applicable operational refs, gaps, warnings, unavailable-tooling notes, or manual-doc handoff are missing from archive evidence | Return `blocked` or route per `skills/_shared/sdd-post-apply-gates.md`; do not complete archive by dropping operational context that exists. |
| Archive criteria require running `sdd-operational-doc` | Fix the archive report before persistence; the utility is manual post-archive and not a DAG gate. |
| Archive report persistence fails | Return `partial` with `next_recommended: resolve-blockers` and the full archive report inline in `detailed_report`. |

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Sync Delta Specs to Main Specs

Do not start this step until shared Archive Readiness passes.

**IF mode is `engram`:** Skip filesystem sync — artifacts live in Engram only. The archive report (Step 5) records all observation IDs for traceability.


**IF mode is `openspec`:** For each delta spec in `openspec/changes/{change-name}/specs/`:

#### If Main Spec Exists (`openspec/specs/{domain}/spec.md`)

Read the existing main spec and apply the delta:

```
FOR EACH SECTION in delta spec:
├── ADDED Requirements → Append to main spec's Requirements section
├── MODIFIED Requirements → Replace the matching requirement in main spec
├── REMOVED Requirements → Delete the matching requirement from main spec after recording Reason/Migration
└── RENAMED Requirements → Rename the matching requirement while preserving scenarios unless the delta also modifies them
```

**Merge carefully:**
- Match requirements by name (e.g., "### Requirement: Session Expiration")
- Preserve all OTHER requirements that aren't in the delta
- Maintain proper Markdown formatting and heading hierarchy
- For REMOVED requirements, require `(Reason: ...)` and `(Migration: ...)` notes in the delta before deleting from main specs
- For RENAMED requirements, require the old and new requirement names to be explicit
- If the merge would be destructive, STOP before editing and return `blocked` with `next_recommended: resolve-blockers` and `confirmation_required: destructive-merge`

#### If Main Spec Does NOT Exist

The delta spec IS a full spec (not a delta). Copy it directly:

```bash
# Copy new spec to main specs
openspec/changes/{change-name}/specs/{domain}/spec.md
  → openspec/specs/{domain}/spec.md
```

### Step 3: Move to Archive

**IF mode is `engram`:** Skip — there are no `openspec/` directories to move. The archive report in Engram serves as the audit trail.


**IF mode is `openspec`:** Move the entire change folder to archive with date prefix:

```
openspec/changes/{change-name}/
  → openspec/changes/archive/YYYY-MM-DD-{change-name}/
```

Use today's date in ISO format (e.g., `2026-02-16`).
If the destination already exists, STOP and return `blocked` with the existing destination path. Do not overwrite or mutate an existing archive folder.

**Move semantics — follow this exact sequence:**

> **PATH CRITICAL**: The archive destination is `openspec/changes/archive/YYYY-MM-DD-{change-name}/` — NEVER `openspec/archive/`.

1. Confirm `openspec/changes/archive/` exists; create it if not.
2. Confirm the destination `openspec/changes/archive/YYYY-MM-DD-{change-name}/` does NOT exist; if it does, STOP and return `blocked`.
3. Copy ALL files and subdirectories from `openspec/changes/{change-name}/` to the destination. Do not skip any file, including JSON reports, Markdown files, test-cases, and apply-progress.
4. Verify the destination contains the same file count as the source before deletion.
5. Delete the ENTIRE source folder `openspec/changes/{change-name}/` — not individual files, the whole directory.
6. Verify immediately that `openspec/changes/{change-name}/` no longer exists. If it still exists after deletion, return `status: partial` with `next_recommended: resolve-blockers` and include the exact recovery steps: "Delete openspec/changes/{change-name}/ manually — archive copy at openspec/changes/archive/YYYY-MM-DD-{change-name}/ is complete."

**Never** leave the source folder partially populated. A partial state (some files deleted, some remaining) is worse than no deletion — it creates ambiguity about which copy is authoritative.

### Step 4: Verify Archive

**IF mode is `openspec`:** Confirm:
- [ ] Main specs updated correctly
- [ ] Change folder moved to archive
- [ ] Archive contains all artifacts (proposal, specs, design, test-design, tasks)
- [ ] Archive contains design artifact with mandatory `## Secure Development Design`
- [ ] Embedded secure-design validation metadata or static/manual notes are preserved and non-failing, or archive stops with a blocker
- [ ] Archive contains canonical `review-report.json` when present plus derived `review-report.md` / review compatibility artifact with a non-blocking verdict; JSON wins on conflict
- [ ] Archive contains canonical `review-security-report.json` when present plus derived `review-security-report.md` / security review compatibility artifact with a non-blocking verdict; JSON wins on conflict
- [ ] Missing `test-design.md` is blocked unless an explicit partial archive exception is provided and recorded in the archive report
- [ ] Mandatory applicable security evidence is verified or covered by complete approved exceptions recorded in the audit trail
- [ ] Applicable operational status, evidence refs, unresolved gaps, warning carry-forward, unavailable-tooling notes, exceptions, and manual operational document handoff boundaries are preserved without requiring absent readiness categories or `sdd-operational-doc` execution
- [ ] Archived `tasks.md` has no unchecked implementation tasks, unless the orchestrator explicitly approved archive-time stale-checkbox reconciliation backed by apply-progress/verify-report proof
- [ ] Active changes directory no longer has this change — verify by checking that `openspec/changes/{change-name}/` does NOT exist; if it still exists, return `status: partial` with `next_recommended: resolve-blockers` immediately, do not continue to Step 5
- [ ] Archive report lists all synced domains, archive destination, verification verdict, and any intentional-with-warnings reason

**IF mode is `engram`:** Confirm all artifact observation IDs are recorded in the archive report and the tasks observation has no unchecked implementation tasks unless the orchestrator explicitly approved archive-time stale-checkbox reconciliation backed by apply-progress/verify-report proof.


### Step 5: Persist Archive Report

This step is mandatory for `engram` and `openspec`.

Before persistence, validate the archive report includes:
- Change name and artifact store mode
- Observation IDs for Engram artifacts, or concrete OpenSpec paths for filesystem artifacts
- `test-design` artifact ref/path, or explicit partial archive exception text when intentionally omitted
- canonical `review-report.json` ref/path when present, derived Markdown compatibility ref/path, and confirmation that the authoritative review verdict is non-blocking
- canonical `review-security-report.json` ref/path when present, derived Markdown compatibility ref/path, and confirmation that the authoritative security-review verdict is non-blocking
- Mandatory `design.md#secure-development-design` ref/path, including no-impact changed-surface rationale when applicable; source-row validation coverage and `N/A` decisions belong to canonical `review-security-report.json` and are presented through derived Markdown compatibility
- Embedded secure-design validation metadata or static/manual notes: source section, status, catalog snapshot identity, lifecycle vocabulary, and validation notes
- Mandatory security evidence status and complete approved exception details for any accepted gaps
- Archive evidence fields for applicable controls: guideline IDs, taxonomy categories, source refs, operational severity, expected evidence status, residual risks, and exception state
- Task completion status and any stale-checkbox reconciliation proof
- Review verdict and confirmation that no blocking review findings were archived
- Verification verdict and confirmation that no CRITICAL issues were archived
- Runtime test runner/linter/typechecker/formatter/coverage availability from verify evidence; unavailable tools must be recorded explicitly rather than treated as passing evidence
- Corporate source-row audit trail when applicable: canonical JSON ref, derived Markdown compatibility ref, catalog snapshot identity/path, expected and validated expanded Source ID counts, exact-once coverage status, safe evidence refs, `N/A` evidence/justification status, warning-only findings, complete exceptions, parity metadata, review-security source-row verdict, verify source-row consumption, report links, and confirmation that no source-row blockers remain
- Operational evidence audit trail when applicable: operational status, evidence refs, unresolved `Pendiente de confirmar:` gaps, `No aplica.` states, warning carry-forward, unavailable-tooling notes, complete exceptions if any, and manual-document handoff boundary confirming `sdd-operational-doc` remains manual post-archive
- Confirmation that active new-change archive readiness uses embedded secure design, review-security, verify, and safe audit references
- Specs synced by domain with created/updated/removed/renamed counts
- Archive destination or inline-only closure reason
- Any intentional-with-warnings approval text and reason

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `archive-report`
- topic_key: `sdd/{change-name}/archive-report`
- openspec reference: `openspec/changes/archive/YYYY-MM-DD-{change-name}/` (the moved archive folder is the OpenSpec archive-report/audit-trail reference)
- type: `architecture`

### Step 6: Return Envelope

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the full archive summary in `detailed_report`:

```markdown
## Change Archived

**Change**: {change-name}
**Archived to**: `openspec/changes/archive/{YYYY-MM-DD}-{change-name}/` (openspec) | Engram archive report (engram)

### Specs Synced
| Domain | Action | Details |
|--------|--------|---------|
| {domain} | Created/Updated | {N added, M modified, K removed requirements} |

### Archive Contents
- proposal.md ✅
- specs/ ✅
- design.md ✅
- design.md#secure-development-design ✅
- test-design.md ✅
- review-report.json ✅ when present; derived review-report.md / review compatibility view ✅ (non-blocking; JSON authoritative)
- review-security-report.json ✅ when present; derived review-security-report.md / security review compatibility view ✅ (non-blocking; JSON authoritative)
- operational evidence ✅ when applicable (status, refs, gaps, warnings, unavailable tooling, manual-doc handoff)
- tasks.md ✅ ({N}/{N} tasks complete)

### Source of Truth Updated
The following specs now reflect the new behavior:
- `openspec/specs/{domain}/spec.md`

### SDD Cycle Complete
The change has been fully planned, implemented, verified, and archived.
Ready for the next change.
```

## Rules

- NEVER archive a change that has blocking review findings or CRITICAL issues in its verification report
- NEVER archive missing mandatory security evidence unless every missing item has a complete approved exception in the audit trail
- NEVER archive unresolved source-row blockers, unsafe evidence, unsupported `N/A`, malformed source-row schema, or missing mandatory source-row evidence.
- ALWAYS preserve source-row coverage, catalog identity, expected and validated counts, exact-once coverage status, warnings, exceptions, safe evidence references, parity metadata, report links, and unavailable-tooling notes in the audit trail when corporate source-row validation applies; cite canonical refs and summaries rather than copying the full source-row matrix.
- ALWAYS preserve applicable operational status, refs, gaps, exact markers, warnings, unavailable-tooling notes, and manual operational-document handoff boundaries; never require absent readiness categories or `sdd-operational-doc` execution to complete archive.
- If the user explicitly approves a non-critical partial archive or stale-checkbox reconciliation, record the exact reason in the archive report and mark the archive as intentional-with-warnings
- NEVER archive completed work while `tasks.md` / the tasks observation still shows stale unchecked implementation tasks
- ALWAYS sync delta specs BEFORE moving to archive
- When merging into existing specs, PRESERVE requirements not mentioned in the delta
- Use ISO date format (YYYY-MM-DD) for archive folder prefix
- If the merge would be destructive (removing large sections), return `blocked` with `next_recommended: resolve-blockers` and `confirmation_required: destructive-merge`; do not ask the user directly
- The archive is an AUDIT TRAIL — never delete or modify archived changes
- If `openspec/changes/archive/` doesn't exist, create it
- Apply any `rules.archive` from `openspec/config.yaml`
- Use embedded design, review-security, verify, and safe audit references as the authoritative active new-change archive gates.
