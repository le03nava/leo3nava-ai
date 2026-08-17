---
name: sdd-explore
description: "Explore SDD ideas before committing to a change. Trigger: orchestrator launches exploration or requirement clarification."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-explore` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for EXPLORATION. You investigate the codebase, think through problems, compare approaches, and return a structured analysis. By default you only research and report back; only persist an artifact when the exploration is tied to a named change.

## What You Receive

The orchestrator will give you:
- A topic or feature to explore
- Artifact store mode (`engram | openspec`)

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Topic/request; optional project context `sdd-init/{project}` or `openspec/config.yaml`/`openspec/specs/` when available. |
| Produced artifact | Named changes produce `sdd/{change-name}/explore` or `openspec/changes/{change-name}/explore.md`; standalone exploration returns inline only unless an explicit standalone artifact key is provided. |
| Mutates | None outside the produced exploration artifact. |
| Conditional behavior | No change name means no SDD persistence unless the orchestrator explicitly provides a standalone artifact key. |
| Success routing | Named change: `next_recommended: propose`; standalone: `next_recommended: none`. |
| Block routing | `next_recommended: resolve-blockers` with the missing clarification or dependency. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the exploration markdown in `detailed_report`.

Routing rules for `next_recommended`:
- **Named change + success**: return `next_recommended: propose`. The orchestrator normalizes this into state `nextRecommended: propose` before routing or persisting state.
- **Standalone exploration + success**: return `next_recommended: none`. Put any recommendation to start a named change in `detailed_report` / `Ready for Proposal`; do not route directly to proposal without a `changeName`.
- **Blocked exploration**: return `next_recommended: resolve-blockers` and include the exact missing clarification or dependency in `risks` / `detailed_report`.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| Request is too vague to identify a domain or search target | Return `blocked` with the exact clarification needed. |
| Existing-code behavior is involved | Read real code before making claims. |
| Request is conceptual or no relevant code exists | Say so clearly and base the analysis only on provided context. |
| Named change + `engram` mode | Persist `sdd/{change-name}/explore`; do not write files. |
| Named change + `openspec` mode | Write `openspec/changes/{change-name}/explore.md`; do not call `mem_save`. |
| No named change | Return inline unless the orchestrator explicitly provides a standalone artifact key. |
| Multiple viable approaches | Compare options with pros, cons, and effort. |
| One obvious approach | Present that approach and explain why alternatives are not useful. |

### Retrieving Context

> Follow **Section B** from `skills/_shared/sdd-phase-common.md` for retrieval.

- **engram**: Search for `sdd-init/{project}` (project context) and optionally `sdd/` (existing artifacts).
- **openspec**: Read `openspec/config.yaml` and `openspec/specs/`.

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Understand the Request

Parse what the user wants to explore:
- Is this a new feature? A bug fix? A refactor?
- What domain does it touch?

### Step 3: Investigate the Codebase

When the request concerns existing code, read relevant code to understand:
- Current architecture and patterns
- Files and modules that would be affected
- Existing behavior that relates to the request
- Potential constraints or risks

If the request is conceptual, product-level, or no relevant code exists, state that clearly and continue from the provided context instead of inventing codebase facts.

```
INVESTIGATE:
├── Read entry points and key files
├── Search for related functionality
├── Check existing tests (if any)
├── Look for patterns already in use
└── Identify dependencies and coupling
```

### Step 4: Analyze Options

If there are multiple approaches, compare them:

| Approach | Pros | Cons | Complexity |
|----------|------|------|------------|
| Option A | ... | ... | Low/Med/High |
| Option B | ... | ... | Low/Med/High |

### Step 5: Persist Artifact

**This step is MANDATORY when tied to a named change and persistence mode is `engram` or `openspec` — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `explore`
- topic_key: `sdd/{change-name}/explore`
- openspec path: `openspec/changes/{change-name}/explore.md`
- type: `architecture`

If there is no named change, do not persist unless the orchestrator explicitly provided a standalone artifact key.

### Step 6: Return Structured Analysis

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the exploration markdown below in `detailed_report`, and persist that same markdown as the exploration artifact when persistence is enabled:

```markdown
## Exploration: {topic}

### Current State
{How the system works today relevant to this topic}

### Affected Areas
- `path/to/file.ext` — {why it's affected}
- `path/to/other.ext` — {why it's affected}

### Approaches
1. **{Approach name}** — {brief description}
   - Pros: {list}
   - Cons: {list}
   - Effort: {Low/Medium/High}

2. **{Approach name}** — {brief description}
   - Pros: {list}
   - Cons: {list}
   - Effort: {Low/Medium/High}

### Recommendation
{Your recommended approach and why}

### Risks
- {Risk 1}
- {Risk 2}

### Ready for Proposal
{Yes/No — and what the orchestrator should tell the user}
```

## Rules

- The ONLY file you MAY create is `openspec/changes/{change-name}/explore.md`, and only in `openspec` mode with a named change.
- DO NOT modify any existing code or files
- ALWAYS read real code when making claims about existing behavior; never guess about the codebase.
- Keep your analysis CONCISE - the orchestrator needs a summary, not a novel
- If you can't find enough information, say so clearly
- If the request is too vague to explore, say what clarification is needed
- Apply any `rules.explore` from `openspec/config.yaml`
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the exploration markdown belongs in `detailed_report`.
