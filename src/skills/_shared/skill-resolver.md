# Skill Resolver — Universal Protocol

Any agent that **delegates work to sub-agents** MUST use this protocol to resolve relevant skills and pass them safely.

## Why This Exists

Sub-agents start with no project skill context. The registry gives delegators a cheap index of available skills without rewriting or summarizing those skills.

## When to Apply

Before every sub-agent launch that involves reading, writing, reviewing, testing, documenting, or creating project artifacts. Skip only for purely mechanical commands.

## The Protocol

### Step 1: Obtain the Skill Registry

The registry is an **index** of skill names, triggers, scopes, and exact `SKILL.md` paths. It is not a compact-rules bundle.

Resolution order:
1. Use the session cache if present.
2. `mem_search(query: "skill-registry", project: "{project}")` → `mem_get_observation(id)` for full content.
3. Fallback: read `.atl/skill-registry.md` from the project root.
4. No registry found → proceed without project skills and warn the user to run `gentle-ai skill-registry refresh`.

### Step 2: Match Relevant Skills

Match on two dimensions:

| Context | Match against |
| --- | --- |
| Code/files | Registry trigger/description mentions the language, framework, tool, or path context |
| Task/action | Registry trigger/description mentions actions like PR, review, docs, tests, Jira, comments, release |

Prefer the smallest useful set. If more than five skills match, keep the five most relevant and prioritize code context over task context.

### Step 3: Pass Skill Paths

Inject paths, not summaries:

```markdown
## Skills to load before work

Read these exact files before reading, writing, reviewing, testing, or creating artifacts:

- /absolute/path/to/skills/go-testing/SKILL.md
- /absolute/path/to/skills/typescript/SKILL.md
```

The sub-agent MUST read those files before task-specific work. `SKILL.md` is the runtime contract and source of truth.

Executor loading rules:

1. If a `## Skills to load before work` block is present, read only those exact paths for supplemental skills and ignore redundant `SKILL: Load` instructions.
2. If no injected paths are present but explicit `SKILL: Load` paths exist, read those exact files and report `mode: fallback-path`.
3. If neither injected paths nor explicit fallback paths exist, resolve the registry using Step 1, match relevant skills using Step 2, read the exact listed paths, and report `mode: fallback-registry`.
4. If no registry or relevant supplemental skills exist, proceed with the phase skill only and report `mode: none`.

Searching the registry for skill loading is not delegation. Executor phase agents still execute their own phase and MUST NOT launch sub-agents.

### Step 4: Report Resolution

Sub-agents MUST report `skill_resolution`:

- `paths-injected` — received exact skill paths from the delegator and loaded them.
- `fallback-registry` — no paths received, self-loaded paths from the registry.
- `fallback-path` — loaded an explicit fallback path outside the registry.
- `none` — no skills loaded.

If a sub-agent reports anything other than `paths-injected`, the orchestrator MUST re-read the registry before the next delegation.

Expected shape:

```yaml
skill_resolution:
  mode: paths-injected | fallback-registry | fallback-path | none
  loaded:
    - name: {skill-name}
      path: {absolute SKILL.md path-or-null}
  missing_required:
    - {skill-name}
  registry_source: session-cache | engram | atl-file | fallback-path | none
  notes: {text-or-null}
```

Required-skill checks:

- If the launch envelope included `skill_paths`, the result should report `mode: paths-injected` and list the loaded skills, unless the sub-agent explicitly explains why a path could not be read.
- If a task required `chained-pr`, `work-unit-commits`, testing, review, security, PR, or repo-specific skills, `missing_required` must be empty before dependent work continues.
- If `sdd-apply` runs with review-budget risk or chained delivery, missing `chained-pr` or `work-unit-commits` is a gate failure.
- If `sdd-verify` runs with Strict TDD or testing requirements, missing relevant testing/verification instructions is a gate failure.

Acceptance policy:

| `skill_resolution.mode` | Low-risk read-only task | SDD planning phase | `sdd-apply` / `sdd-verify` / PR / security-sensitive work |
| --- | --- | --- | --- |
| `paths-injected` | Accept | Accept | Accept if required skills loaded |
| `fallback-registry` | Accept with warning; refresh cache | Accept only if required skills loaded; refresh cache | Gate warning or failure if required skills were not loaded |
| `fallback-path` | Accept only for known allowed paths | Accept only for known required paths; refresh registry | Gate failure unless all required paths are known and loaded |
| `none` | Accept only if no relevant skills existed | Gate warning; re-check registry before next phase | Gate failure when skills were required |

Corrective loop:

- When required skills are missing, re-read the registry (`mem_search` -> `mem_get_observation`, fallback `.atl/skill-registry.md`) and re-run the same phase once with exact `SKILL.md` paths injected.
- If the retry still misses required skills, STOP and report the missing skills, registry source used, and the phase blocked.
- Persist a discovery when registry lookup is missing, stale, or repeatedly fails so future sessions do not silently repeat the problem.

## Compaction Safety

- The registry persists in Engram and `.atl/skill-registry.md`.
- Delegators can recover selected paths after compaction by re-reading the registry.
- Sub-agents receive exact files to read, so skill meaning is not degraded by generated summaries.

## Integration Points

- **ATL Orchestrator**: resolves paths for all SDD and non-SDD delegations.
- **judgment-day**: resolves paths before Judge A, Judge B, and Fix Agent.
- **pr-review and future delegators**: use this protocol when launching sub-agents.
