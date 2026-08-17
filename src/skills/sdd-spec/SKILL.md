---
name: sdd-spec
description: "Write SDD delta specs with requirements and scenarios. Trigger: orchestrator launches spec work for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-spec` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for writing SPECIFICATIONS. You take the proposal and produce delta specs — structured requirements and scenarios that describe what's being ADDED, MODIFIED, REMOVED, or RENAMED from the system's behavior.

## What You Receive

From the orchestrator:
- Change name
- Artifact store mode (`engram | openspec`)

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal `sdd/{change-name}/proposal` or `openspec/changes/{change-name}/proposal.md`; existing source specs when modifying OpenSpec capabilities. |
| Produced artifact | Engram single concatenated `sdd/{change-name}/spec`; OpenSpec domain files `openspec/changes/{change-name}/specs/{domain}/spec.md`. |
| Mutates | None outside produced delta/new spec artifacts. |
| Conditional behavior | Capabilities drive domain selection; `MODIFIED` requirements must contain full replacement blocks. |
| Success routing | `next_recommended: design`. |
| Block routing | `next_recommended: resolve-blockers` for missing proposal/capability/source spec or validation issue. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the specs summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Successful specs**: return `next_recommended: design`. The orchestrator normalizes this into state `nextRecommended: design` before routing or persisting state.
- **Blocked specs**: return `next_recommended: resolve-blockers` and include the exact missing proposal, capability, existing spec, or validation issue in `risks` / `detailed_report`.
- **Partial persistence failure**: return `next_recommended: resolve-blockers` unless the same artifact can be safely retried without new user input.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| Required proposal is missing | Return `blocked` with `next_recommended: resolve-blockers`; do not write specs. |
| Proposal has no `Capabilities` section and `Affected Areas` is ambiguous | Return `blocked` with `next_recommended: resolve-blockers` and the exact clarification needed. |
| `Modified Capabilities` references a missing existing spec | Return `blocked` with `next_recommended: resolve-blockers` unless the proposal explicitly says to create it as new. |
| `engram` mode | Do not create `openspec/`; persist only `sdd/{change-name}/spec`. |
| `openspec` mode | Write only `openspec/changes/{change-name}/specs/{domain}/spec.md`; do not call `mem_save`. |
| Spec draft fails validation | Fix it before persistence; if it cannot be fixed, return `blocked` with `next_recommended: resolve-blockers`. |

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Identify Affected Domains

Read the proposal's **Capabilities section** — this is your primary contract:

```
FOR EACH entry under "New Capabilities":
├── This becomes a NEW full spec artifact for this change
├── OpenSpec path, in openspec: openspec/changes/{change-name}/specs/<capability-name>/spec.md
├── Engram shape: a `# <capability-name>` domain section inside the returned/persisted spec artifact
└── Write a complete spec (not a delta) — archive later promotes it to openspec/specs/<capability-name>/spec.md

FOR EACH entry under "Modified Capabilities":
├── This becomes a DELTA spec in openspec: openspec/changes/{change-name}/specs/<capability-name>/spec.md
├── In engram, this becomes a modified domain section in the single spec artifact
└── In openspec, read existing openspec/specs/<capability-name>/spec.md first — your delta modifies it
```

If the proposal has no Capabilities section (older format), fall back to inferring from "Affected Areas". But always prefer the explicit Capabilities mapping when present. Do not create or read OpenSpec paths in `engram` mode.

### Step 3: Read Existing Specs

**IF mode is `openspec`:** If `openspec/specs/{domain}/spec.md` exists, read it to understand CURRENT behavior. Your delta specs describe CHANGES to this behavior.

**IF mode is `engram`:** Read `sdd/{change-name}/proposal` and check whether `sdd/{change-name}/spec` already exists before saving. If existing domain spec context is unavailable in Engram, state that limitation in the return envelope.


### Step 4: Write Delta Specs

**IF mode is `openspec`:** Create specs inside the change folder:

```
openspec/changes/{change-name}/
├── proposal.md              ← (already exists)
└── specs/
    └── {domain}/
        └── spec.md          ← Delta spec
```

**IF mode is `engram`:** Do NOT create any `openspec/` directories or files. Compose the spec content in memory; persist it only in Engram.

#### MODIFIED Requirements Workflow (CRITICAL — read before writing deltas)

When writing a `## MODIFIED Requirements` section, follow this exact workflow:

```
1. Locate the requirement in openspec/specs/{domain}/spec.md
2. COPY the ENTIRE requirement block — from `### Requirement:` through ALL its scenarios
3. PASTE it under `## MODIFIED Requirements`
4. EDIT the copy to reflect the new behavior
5. Add "(Previously: {one-line summary of what changed})" under the requirement text

Why copy-full-then-edit?
→ The archive step REPLACES the requirement in main specs with your MODIFIED block
→ If your block is partial, the archive will lose scenarios you didn't copy
→ Common pitfall: only writing the changed scenario and losing the rest
→ If adding NEW behavior WITHOUT changing existing behavior, use ADDED instead
```

#### Delta Spec Format

```markdown
# Delta for {Domain}

## ADDED Requirements

### Requirement: {Requirement Name}

{Description using RFC 2119 keywords: MUST, SHALL, SHOULD, MAY}

The system {MUST/SHALL/SHOULD} {do something specific}.

#### Scenario: {Happy path scenario}

- GIVEN {precondition}
- WHEN {action}
- THEN {expected outcome}
- AND {additional outcome, if any}

#### Scenario: {Edge case scenario}

- GIVEN {precondition}
- WHEN {action}
- THEN {expected outcome}

## MODIFIED Requirements

### Requirement: {Existing Requirement Name}

{Full updated requirement text — replaces the existing one entirely}
(Previously: {what it was before, in one line})

#### Scenario: {Unchanged scenario — keep if still valid}

- GIVEN {precondition}
- WHEN {action}
- THEN {outcome}

#### Scenario: {Updated or new scenario}

- GIVEN {updated precondition}
- WHEN {updated action}
- THEN {updated outcome}

## REMOVED Requirements

### Requirement: {Requirement Being Removed}

(Reason: {why this requirement is being deprecated/removed})
(Migration: {what replaces it, or "None" if no migration is needed})

## RENAMED Requirements

### Requirement: {Old Requirement Name} → {New Requirement Name}

(Reason: {why the requirement is being renamed})
(Migration: {how references/tests/docs should update, or "None" if no migration is needed})
```

#### For NEW Specs (No Existing Spec)

If this is a completely new domain, create a FULL spec (not a delta):

```markdown
# {Domain} Specification

## Purpose

{High-level description of this spec's domain.}

## Requirements

### Requirement: {Name}

The system {MUST/SHALL/SHOULD} {behavior}.

#### Scenario: {Name}

- GIVEN {precondition}
- WHEN {action}
- THEN {outcome}
```

### Step 5: Validate Specs

Before persisting or returning, verify:

- Every requirement has at least one scenario.
- Every scenario uses Given/When/Then.
- Every requirement uses RFC 2119 keywords.
- Every `MODIFIED` requirement is the full replacement block, not a partial diff.
- Every `REMOVED` requirement includes `Reason`.
- Every `RENAMED` requirement states old and new names explicitly.
- The spec artifact stays under the 650-word size budget.

### Step 6: Persist Artifact

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `spec`
- topic_key: `sdd/{change-name}/spec`
- openspec path: `openspec/changes/{change-name}/specs/{domain}/spec.md`
- type: `architecture`

### Step 7: Return Summary

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Use the canonical YAML format defined in the Section D example. Adapt these fields for this phase:

- `phase`: `spec`
- `next_recommended`: `design` (success) | `resolve-blockers` (blocked/partial)
- `executive_summary`: one short paragraph — domains covered, total requirements added/modified, total scenarios, and coverage summary (happy paths, edge cases, error states)
- `artifacts`: one entry per spec file produced; type `spec`; correct mode, ref/path, `persisted: true`, `readable: true`. In openspec, include one entry per domain file.
- `risks`: structured array or `None` — never an empty array `[]`
- `skill_resolution`: from `skill-resolver.md#step-4-report-resolution`
- `detailed_report`: use this minimum content for `sdd-spec`:

```
## Specs Created

**Change**: {change-name}

### Specs Written
| Domain | Type | Requirements | Scenarios |
|--------|------|-------------|-----------|
| {domain} | Delta/New | {N added, M modified, K removed} | {total scenarios} |

### Coverage
- Happy paths: {covered/missing}
- Edge cases: {covered/missing}
- Error states: {covered/missing}

### Next Step
Ready for design (sdd-design).
```

## Rules

- ALWAYS use Given/When/Then format for scenarios
- ALWAYS use RFC 2119 keywords (MUST, SHALL, SHOULD, MAY) for requirement strength
- Read the proposal's **Capabilities section** first — it tells you exactly which spec files to create
- If existing specs exist, write DELTA specs (ADDED/MODIFIED/REMOVED sections)
- If NO existing specs exist for the domain, write a FULL spec
- Every requirement MUST have at least ONE scenario
- Include both happy path AND edge case scenarios
- Keep scenarios TESTABLE — someone should be able to write an automated test from each one
- DO NOT include implementation details in specs — specs describe WHAT, not HOW
- **MODIFIED requirements MUST be the FULL block** — copy entire requirement + all scenarios from main spec, then edit. Partial MODIFIED blocks lose content at archive time.
- If adding new behavior without changing existing behavior → use ADDED, not MODIFIED
- REMOVED requirements MUST include Reason and SHOULD include Migration when consumers, persisted behavior, docs, or tests are affected
- RENAMED requirements MUST state both old and new names explicitly and SHOULD include Migration guidance for references/tests/docs
- Apply any `rules.specs` from `openspec/config.yaml`
- **Size budget**: Spec artifact MUST be under 650 words. Prefer requirement tables over narrative descriptions. Each scenario: 3-5 lines max.
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the specs summary belongs in `detailed_report`.

## RFC 2119 Keywords Quick Reference

| Keyword | Meaning |
|---------|---------|
| **MUST / SHALL** | Absolute requirement |
| **MUST NOT / SHALL NOT** | Absolute prohibition |
| **SHOULD** | Recommended, but exceptions may exist with justification |
| **SHOULD NOT** | Not recommended, but may be acceptable with justification |
| **MAY** | Optional |
