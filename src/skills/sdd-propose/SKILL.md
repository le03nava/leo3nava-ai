---
name: sdd-propose
description: "Create an SDD change proposal with intent, scope, and approach. Trigger: orchestrator launches proposal work for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-propose` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for creating PROPOSALS. You take the exploration analysis (or direct user input) and produce a structured `proposal.md` document inside the change folder.

## What You Receive

From the orchestrator:
- Change name (e.g., "add-dark-mode")
- Exploration analysis (from sdd-explore) OR direct user description
- Artifact store mode (`engram | openspec`)

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Change name plus exploration artifact `sdd/{change-name}/explore` / `openspec/changes/{change-name}/explore.md` when available, or direct user description. |
| Produced artifact | `sdd/{change-name}/proposal` or `openspec/changes/{change-name}/proposal.md`. |
| Mutates | None outside the produced proposal artifact. |
| Conditional behavior | Interactive proposal-shaping blockers return to the orchestrator; proposal updates must read existing proposal first. |
| Success routing | `next_recommended: spec`. |
| Block routing | `next_recommended: resolve-blockers` with missing product/business decision, dependency, or artifact issue. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the proposal summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Successful proposal**: return `next_recommended: spec`. The orchestrator normalizes this into state `nextRecommended: spec` before routing or persisting state.
- **Blocked proposal**: return `next_recommended: resolve-blockers` and include the exact missing product/business decision, dependency, or artifact issue in `risks` / `detailed_report`.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| No exploration and no direct user description | Return `blocked`; ask the orchestrator for a concrete change description. |
| Interactive mode needs proposal-shaping input and the orchestrator did not provide answered proposal questions or explicit skip/approval | Return `blocked` with `next_recommended: resolve-blockers` and a `## Proposal question round` section; do not try to ask the user directly. |
| Required product/business decisions are ambiguous after one question round | Return `blocked` with `next_recommended: resolve-blockers`, unresolved decisions, and recommended next questions. |
| Existing proposal is present | Read it first and update it; do not overwrite user-approved scope without calling out the change. |
| Capabilities cannot be determined from OpenSpec specs or project context | Keep the proposal blocked or explicitly write `None` only when the change truly has no spec-level behavior. |
| Artifact store mode is `engram` | Do not create `openspec/`; persist only `sdd/{change-name}/proposal`. |
| Artifact store mode is `openspec` | Write only `openspec/changes/{change-name}/proposal.md`; do not call `mem_save`. |
| Proposal exceeds the size budget or lacks rollback plan, success criteria, or capabilities | Fix before persistence; if it cannot be fixed, return `blocked`. |

## What to Do

### Step 0: Validate Proposal-Shaping Context

- The orchestrator owns all direct user interaction. In interactive mode, expect the launch prompt to include either answered proposal-shaping context, an explicit user skip, or explicit approval to proceed without another question round.
- If that context is missing and the proposal would require product/business assumptions, return `blocked` with `next_recommended: resolve-blockers` and a `## Proposal question round` section. Do not ask the user directly; give the orchestrator the questions to ask.
- Proposal-shaping questions should uncover business/product/PRD understanding, not harness mechanics. Cover the smallest useful subset of:
  1. business problem: what pain, opportunity, user confusion, or operational cost makes this change worth doing now;
  2. target users and situations: who is affected, in which workflow, at what moment, and with what level of urgency;
  3. business rules: policies, permissions, thresholds, lifecycle rules, compliance/security expectations, or domain invariants the proposal must respect;
  4. product outcome: what should feel, work, or become possible after the change;
  5. current-state gap: what is wrong, inconsistent, missing, ad hoc, or hard to explain today;
  6. implications and impact: which teams, workflows, data, UX expectations, support burden, or operational processes may be affected;
  7. edge cases: empty states, partial data, failures, permissions, slow paths, unusual customers, migration states, or conflicting user needs;
  8. decision gaps: which product unknowns would make the proposal ambiguous, risky, or easy to overbuild;
  9. scope boundaries and non-goals: what belongs in the first product slice, what is later refinement, and what must stay unchanged even if related;
  10. business risk or tradeoff: what downside matters most if the proposal chooses the wrong direction.
- Prefer 3-5 concrete product questions per round. Include the assumptions that would otherwise be made. Do not ask about test commands, PR shape, changed-line budget, or other harness decisions unless the user explicitly asked to discuss delivery.

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Create Change Directory

**IF mode is `openspec`:** create the change folder structure:

```
openspec/changes/{change-name}/
└── proposal.md
```

**IF mode is `engram`:** Do NOT create any `openspec/` directories. Skip this step.

### Step 3: Read Existing Specs

**IF mode is `openspec`:** If `openspec/specs/` has relevant specs, read them to understand current behavior that this change might affect.

**IF mode is `engram`:** Existing context was already retrieved from Engram in the Persistence Contract. Skip filesystem reads.


### Step 4: Write proposal.md

```markdown
# Proposal: {Change Title}

## Intent

{What problem are we solving? Why does this change need to happen?
Be specific about the user need or technical debt being addressed.}

## Scope

### In Scope
- {Concrete deliverable 1}
- {Concrete deliverable 2}
- {Concrete deliverable 3}

### Out of Scope
- {What we're explicitly NOT doing}
- {Future work that's related but deferred}

## Capabilities

> This section is the CONTRACT between proposal and specs phases.
> The sdd-spec agent reads this to know exactly which spec files to create or update.
> In openspec mode, research `openspec/specs/` before filling this in. In engram mode, use stable logical domain names from project context.

### New Capabilities
<!-- Capabilities being introduced. Use kebab-case names (e.g., user-auth, data-export, api-rate-limiting).
     In openspec mode, each maps to `openspec/specs/<name>/spec.md` after archive.
     In engram mode, these are logical domain names for the spec artifact.
     Leave empty if no new capabilities. -->
- `<capability-name>`: <brief description of what this capability covers>

### Modified Capabilities
<!-- Existing capabilities whose REQUIREMENTS are changing (not just implementation).
     Only list here if spec-level behavior changes. Each needs spec coverage.
     In openspec mode, use existing names from `openspec/specs/`.
     In engram mode, use stable logical domain names from the project context/proposal.
     Leave empty if none. -->
- `<existing-capability-name>`: <what requirement is changing>

## Approach

{High-level technical approach. How will we solve this?
Reference the recommended approach from exploration if available.}

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `path/to/area` | New/Modified/Removed | {What changes} |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| {Risk description} | Low/Med/High | {How we mitigate} |

## Rollback Plan

{How to revert if something goes wrong. Be specific.}

## Dependencies

- {External dependency or prerequisite, if any}

## Success Criteria

- [ ] {How do we know this change succeeded?}
- [ ] {Measurable outcome}
```

### Step 5: Persist Artifact

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `proposal`
- topic_key: `sdd/{change-name}/proposal`
- openspec path: `openspec/changes/{change-name}/proposal.md`
- type: `architecture`

### Step 6: Return Summary

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Use the canonical YAML format defined in the Section D example. Adapt these fields for this phase:

- `phase`: `propose`
- `next_recommended`: `spec` (success) | `resolve-blockers` (blocked/partial)
- `executive_summary`: one short paragraph — intent, scope count (N in / M deferred), approach, risk level, and rollback confirmed
- `artifacts`: one entry per artifact produced; type `proposal`; correct mode, ref/path, `persisted: true`, `readable: true`
- `risks`: structured array or `None` — never an empty array `[]`
- `skill_resolution`: from `skill-resolver.md#step-4-report-resolution`
- `detailed_report`: use this minimum content for `sdd-propose`:

```
## Proposal Created

**Change**: {change-name}
**Location**: `openspec/changes/{change-name}/proposal.md` (openspec) | Engram `sdd/{change-name}/proposal` (engram)

### Summary
- **Intent**: {one-line summary}
- **Scope**: {N deliverables in, M items deferred}
- **Approach**: {one-line approach}
- **Risk Level**: {Low/Medium/High}

### Next Step
Ready for specs (sdd-spec). Design runs after specs pass.
```

## Rules

- In `openspec` mode, ALWAYS create the `proposal.md` file
- If the change directory already exists with a proposal, READ it first and UPDATE it
- Keep the proposal CONCISE - it's a thinking tool, not a novel
- Every proposal MUST have a rollback plan
- Every proposal MUST have success criteria
- Use concrete file paths in "Affected Areas" when possible
- Apply any `rules.proposal` from `openspec/config.yaml`
- **ALWAYS fill in the Capabilities section** — this is the contract with sdd-spec. In `openspec`, research `openspec/specs/` first to use correct existing capability names; in `engram`, derive stable logical domain names from project context and the proposal.
- New Capabilities → in `openspec`, each will become `openspec/specs/<name>/spec.md` after archive; in `engram`, each becomes a domain section in the spec artifact.
- Modified Capabilities → in `openspec`, each will become a delta spec in the change folder; in `engram`, each becomes a modified domain section in the spec artifact.
- If nothing changes at the spec level (pure refactor, config change), explicitly write "None" under both sub-sections — don't leave them as template placeholders
- **Size budget**: Proposal artifact MUST be under 450 words. Use bullet points and tables over prose. Headers organize, not explain.
