# Engram Artifact Convention (backend reference only)

NOTE: This document is a backend-specific reference for Engram topic names and call shapes. The authoritative cross-mode persistence rules live in `skills/_shared/persistence-contract.md`; phase skills MUST follow that contract for mode resolution, read/write behavior, state schema, recovery, and persistence verification.

Use this file only when an Engram-specific key or call shape is needed. Do not treat it as a competing source for artifact-store mode behavior, SDD state shape, recovery routing, or lifecycle policy.

## Naming Rules

All SDD artifacts persisted to the Engram backend MUST follow this deterministic naming:

```
title:     sdd/{change-name}/{artifact-type}
topic_key: sdd/{change-name}/{artifact-type}
type:      architecture
project:   {detected or current project name}
scope:     project
capture_prompt: false
```

Set `capture_prompt: false` when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

### Artifact Types

| Artifact Type | Produced By | Description |
|---------------|-------------|-------------|
| `explore` | sdd-explore | Exploration analysis |
| `proposal` | sdd-propose | Change proposal |
| `spec` | sdd-spec | Delta specifications (all domains concatenated) |
| `design` | sdd-design | Technical design |
| `test-design` | sdd-test-design | Test design and evidence plan |
| `tasks` | sdd-tasks | Task breakdown |
| `apply-progress` | sdd-apply | Implementation progress (one per batch) |
| `review` | sdd-review | General review report |
| `review-security` | sdd-review-security | Security review report |
| `verify-report` | sdd-verify | Verification report |
| `archive-report` | sdd-archive | Archive closure with lineage |
| `state` | orchestrator | DAG state for recovery after compaction |

### Project Init Artifacts

`sdd-init` persists project-level context separately from change artifacts:

- `sdd-init/{project}` (`architecture`): detected stack, architecture, conventions, and persistence context.
- `sdd/{project}/testing-capabilities` (`config`): Strict TDD mode, test commands, test layers, coverage, and quality tools.
- `skill-registry` (`config`): skill index with exact `SKILL.md` paths for orchestrator injection.

### State Artifact Key

State uses the normal naming rule with `artifact-type: state`:

```text
topic_key: sdd/{change-name}/state
title:     sdd/{change-name}/state
```

The state schema, recovery order, reconciliation behavior, and read-back requirements are defined in `skills/_shared/persistence-contract.md#state-persistence-orchestrator` and `skills/_shared/persistence-contract.md#state-recovery-orchestrator`.

## Reading Artifacts (2 steps)

```
Step 1: mem_search(query: "sdd/{change-name}/{artifact-type}", project: "{project}") → truncated preview + ID
Step 2: mem_get_observation(id: {observation-id}) → complete content
```

When retrieving multiple artifacts, group all searches first, then all retrievals:

```
STEP A — SEARCH (get IDs only):
  mem_search(query: "sdd/{change-name}/proposal", ...) → save ID
  mem_search(query: "sdd/{change-name}/spec", ...) → save ID
  mem_search(query: "sdd/{change-name}/design", ...) → save ID

STEP B — RETRIEVE FULL CONTENT (mandatory):
  mem_get_observation(id: {proposal_id})
  mem_get_observation(id: {spec_id})
  mem_get_observation(id: {design_id})
```

Project context key:

```text
mem_search(query: "sdd-init/{project}", project: "{project}") → get ID
mem_get_observation(id) → full project context
```

Testing capabilities key:

```text
mem_search(query: "sdd/{project}/testing-capabilities", project: "{project}") → get ID
mem_get_observation(id) → full testing context
```

## Writing Artifacts

Standard write:
```
mem_save(
  title: "sdd/{change-name}/{artifact-type}",
  topic_key: "sdd/{change-name}/{artifact-type}",
  type: "architecture",
  project: "{project}",
  capture_prompt: false,
  content: "{full markdown content}"
)
```

Concrete example — saving a proposal for `add-dark-mode`:
```
mem_save(
  title: "sdd/add-dark-mode/proposal",
  topic_key: "sdd/add-dark-mode/proposal",
  type: "architecture",
  project: "my-app",
  capture_prompt: false,
  content: "## Proposal\n\nAdd dark mode toggle..."
)
```

`capture_prompt: false` is REQUIRED for SDD artifacts when the Engram tool schema supports it. Engram v1.15.3 captures user prompts by default for human/proactive saves, but SDD artifacts are automated pipeline outputs. Do not infer this from `type` because both SDD artifacts and human architecture decisions use `architecture`. If an older schema rejects or does not expose `capture_prompt`, omit it rather than failing.

Update existing artifact (when you have the observation ID):
```
mem_update(id: {observation-id}, content: "{updated full content}")
```

Use `mem_update` when you have the exact ID. Use `mem_save` with same `topic_key` for upserts.

### Browsing All Artifacts for a Change

```
mem_search(query: "sdd/{change-name}/", project: "{project}")
→ Returns all artifacts for that change
```

## Project Name Resolution

Engram normally auto-detects the project name from the git remote. Use the current detected project consistently for all SDD artifact saves so topic-key upserts target the same project namespace.

## Upsert Behavior

Same `topic_key` + `project` + `scope` -> UPDATE (overwrite), not INSERT. Previous content is lost - `revision_count` increments but old content is NOT saved. This is by design - engram is working memory, not an audit trail. For iteration history or team collaboration, use `openspec`.

## Why This Convention

- Deterministic titles → recovery works by exact match
- `topic_key` → enables upserts without duplicates
- `sdd/` prefix → namespaces all SDD artifacts
- Two-step recovery → search previews are always truncated; `mem_get_observation` is the only way to get full content
- Lineage → archive-report includes all observation IDs for complete traceability
