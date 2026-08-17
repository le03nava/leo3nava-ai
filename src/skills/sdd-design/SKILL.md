---
name: sdd-design
description: "Create the SDD technical design and architecture approach. Trigger: orchestrator launches design for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-design` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for TECHNICAL DESIGN and narrative secure development design. You take the proposal and specs, then produce a `design.md` that captures HOW the change will be implemented — architecture decisions, data flow, file changes, changed-surface security classification, applicable security category rules, expected evidence, and technical rationale.

## What You Receive

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.execution_mode` — controls blocking behavior
- `launch.status.dependencies` — `proposal: completed` and `spec: completed` must be present
- `launch.artifacts.paths.proposal` or `launch.artifacts.refs.proposal` — do NOT expect content inline
- `launch.artifacts.paths.spec` or `launch.artifacts.refs.spec` — do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute path to read the codebase from
- `launch.actionContext.allowedEditRoots` — directories this phase may write to
- `launch.skill_paths` — supplemental skills to load before work

If `launch.artifact_store.mode` is absent, or `launch.status.dependencies` does not show both `proposal` and `spec` as completed, return `blocked` with `next_recommended: resolve-blockers`.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal and specs: `sdd/{change-name}/{proposal|spec}` or OpenSpec change-folder equivalents, plus `skills/_shared/sdd-security-contract.md`. |
| Produced artifact | `sdd/{change-name}/design` or `openspec/changes/{change-name}/design.md`. |
| Mutates | None outside the produced design artifact. |
| Secure design authority | For every new active change, `design.md#secure-development-design` is the active design-time classification and planning authority. It MUST classify the changed surface using the shared security contract, and include only human-readable rules for applicable security categories, evidence owners, residual risks, exceptions, and safe-evidence policy. Omitted security categories are reviewable omissions; exhaustive validation coverage and `N/A` decisions belong to `sdd-review-security`. |
| Source-row coverage | The embedded section MAY mention applicable category context in prose when needed. It MUST NOT read or copy the exhaustive 155-row Source ID inventory, create Source ID matrices, encode machine-readable applicability fields, or require per-row `N/A` evidence for omitted rows. |
| Secure design shape | The embedded section MUST be narrative Markdown only: headings, paragraphs, and bullet lists are allowed. It MUST NOT contain security YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, or all-row `N/A` bookkeeping. Historical schema-based exhaustive artifacts remain readable compatibility data only. |
| Operational considerations authority | `design.md` owns operational applicability. Include `## Operational Considerations` only when the change needs operational behavior planning, or state `No aplica.` with rationale when useful. Downstream phases derive checks only from actual design/test-design evidence and MUST NOT require absent operational categories. |
| Success routing | `next_recommended: test-design`; the next active phase is `sdd-test-design`. |
| Block routing | `next_recommended: resolve-blockers` for missing proposal/specs, missing code context, testing capability ambiguity, or unresolved architecture decisions. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the design summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Successful design**: return `next_recommended: test-design`. The orchestrator normalizes this into state `nextRecommended: test-design` before routing or persisting state.
- **Blocked design**: return `next_recommended: resolve-blockers` and include the exact missing proposal, spec, code context, testing capability, or architecture decision in `risks` / `detailed_report`.
- **Partial design**: return `next_recommended: resolve-blockers` unless the same design artifact can be safely retried without new user input.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| Required proposal is missing | Return `blocked` with `next_recommended: resolve-blockers`; do not write design. |
| Specs are missing | Return `blocked` with `next_recommended: resolve-blockers`; do not write design. |
| Affected code cannot be identified | Return `partial` with `next_recommended: resolve-blockers` if useful design context was produced; otherwise return `blocked` with `next_recommended: resolve-blockers`. Do not invent paths. |
| Open questions block implementation decisions | Return `blocked` with `next_recommended: resolve-blockers` and list the blocking questions. |
| Security classification cannot be proven from proposal/specs/code context | Return `blocked` with `next_recommended: resolve-blockers`; do not guess `security-impacting` or `no-impact`. |
| Data sensitivity is unknown for touched inputs, outputs, storage, logs, exports, or external interfaces | Treat as a design-changing unknown for `SEC-DATA-001`; return `blocked` unless the design can prove safe classification with review-safe evidence. |
| Changed-surface classification lacks enough proposal/spec/code-context evidence to justify selected applicable controls and omitted categories | Return `blocked` with `next_recommended: resolve-blockers`; do not guess. |
| Any applicable mandatory guideline lacks control, downstream evidence owner, expected evidence, residual-risk statement, or archive expectation | Fix before persistence, or return `blocked` with `next_recommended: resolve-blockers`. |
| Applicable security category narrative lacks evidence owners or safe-evidence expectations | Fix before persistence, or return `blocked` with `next_recommended: resolve-blockers`. |
| Any exception lacks approver, approvedAt, accepted-risk rationale, mitigation/follow-up, or evidence gap | Return `blocked`; incomplete exceptions cannot satisfy design or archive readiness. |
| Applicable operational behavior lacks `## Operational Considerations` evidence or an explicit gap | Fix before persistence, or return `blocked`; do not silently skip operational implications that the proposal/spec/code context makes applicable. |
| Operational evidence needs unavailable or unsafe operational values | Use exact `Pendiente de confirmar:` or safe evidence refs; do not invent hostnames, owners, dashboards, logs, payloads, credentials, ports, SID/service names, full ID lists, or generated bytes. |
| Monitoring strategy is SQL-only by default | Fix before persistence; monitoring evidence must be mechanism-oriented and may include dashboards, alerts, jobs, traces, scripts, documented manual checks, or SQL only where appropriate. |
| `engram` mode | Do not create `openspec/`; persist only `sdd/{change-name}/design`. |
| `openspec` mode | Write only `openspec/changes/{change-name}/design.md`; do not call `mem_save`. |
| Design draft fails validation | Fix it before persistence; if it cannot be fixed, return `blocked` with `next_recommended: resolve-blockers`. |

## What to Do

### Step 1: Load Skills
Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Read the Codebase

Before making claims about existing behavior, read the actual code that will be affected:
- Entry points and module structure
- Existing patterns and conventions
- Dependencies and interfaces
- Test infrastructure (if any)

Build a compact design-time security inventory before writing the secure section:
- Trust boundaries and external inputs crossed by the change.
- Data classes touched: public, internal, confidential, PII, PAN, credentials, tokens, keys, or unknown.
- Authentication, session, authorization, file, database, logging, error-reporting, export, and configuration/secrets touchpoints.
- Existing controls that are preserved, changed, removed, or newly required.
- Review-safe contract/context references that justify the changed-surface classification without exposing secrets, PAN, PII, tokens, private keys, connection strings, or confidential values.

Also read before writing:
- `skills/_shared/sdd-security-contract.md` for the mandatory narrative secure development design contract, category vocabulary, ownership boundaries, safe-evidence rules, and downstream review-security handoff.
- Operational considerations from proposal/spec/code context when applicable: logs/errors, monitoring mechanisms, administration operations, reprocessing/recovery, backup/retention/cleanup/generated artifacts, final-document boundaries, evidence owners, and unresolved gaps.
- For new changes, classify security impact and plan evidence only in `design.md#secure-development-design`; the design phase does not create a separate security-design artifact and does not emit security schemas or matrices.

Also read testing capabilities when available:
- Engram: `sdd/{project}/testing-capabilities`
- OpenSpec: `openspec/config.yaml` `testing` section

If affected code or testing capabilities cannot be found, state that limitation in the return envelope instead of guessing.

### Step 3: Write design.md

**IF mode is `openspec`:** Create the design document:

```
openspec/changes/{change-name}/
├── proposal.md
├── specs/
└── design.md              ← You create this
```

**IF mode is `engram`:** Do NOT create any `openspec/` directories or files. Compose the design content in memory; persist it in Engram.

#### Design Document Format

````markdown
# Design: {Change Title}

## Technical Approach

{Concise description of the overall technical strategy.
How does this map to the proposal's approach? Reference specs.}

## Architecture Decisions

### Decision: {Decision Title}

**Choice**: {What we chose}
**Alternatives considered**: {What we rejected}
**Rationale**: {Why this choice over alternatives}

### Decision: {Decision Title}

**Choice**: {What we chose}
**Alternatives considered**: {What we rejected}
**Rationale**: {Why this choice over alternatives}

## Data Flow

{Describe how data moves through the system for this change.
Use ASCII diagrams when helpful.}

    Component A ──→ Component B ──→ Component C
         │                              │
         └──────── Store ───────────────┘

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `path/to/new-file.ext` | Create | {What this file does} |
| `path/to/existing.ext` | Modify | {What changes and why} |
| `path/to/old-file.ext` | Delete | {Why it's being removed} |

## Interfaces / Contracts

{Define any new interfaces, API contracts, type definitions, or data structures.
Use code blocks with the project's language.}

## Operational Considerations

### Strategy

{When applicable, describe product/support ownership, logs and error evidence, monitoring mechanisms, administration operations, reprocessing/recovery, backup/retention/cleanup/generated artifacts, final operational document inputs, and unresolved gaps. If not applicable, state `No aplica.` with concise rationale. Real operational data disclosure is not required.}

### Evidence Plan

| Category | Expected safe evidence | Owner phase | Status |
| --- | --- | --- | --- |
| {operational concern} | {path/section/sanitized summary/redacted placeholder/`Pendiente de confirmar:`/`No aplica.`} | {phase} | {state} |

### Restricted Data Boundary

{State that restricted operational data stays out of ordinary SDD evidence and final-document-only values are not backfilled.}

### Unresolved Gaps

- `Pendiente de confirmar:` {field and downstream owner, when known}

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | {What} | {How} |
| Integration | {What} | {How} |
| E2E | {What} | {How} |

## Migration / Rollout

{If this change requires data migration, feature flags, or phased rollout, describe the plan.
If not applicable, state "No migration required."}

## Open Questions

- [ ] {Any unresolved technical question}
- [ ] {Any decision that needs team input}

## Secure Development Design

### Classification and Changed Surface

{Narrative security-impacting or no-impact classification with changed artifacts, touched behaviors, untouched runtime surfaces, security-contract context considered, and why omitted categories are left for review-security validation.}

### {Applicable Security Category} Rules

{Human-readable development rules for the applicable category. Include prohibited unsafe patterns, evidence owners, expected review-safe evidence, residual risks, exception handling, and safe-evidence policy. Repeat only for applicable categories.}

### Exception and Evidence Policy

{State complete exception requirements when needed and safe evidence rules. If no exceptions are planned, say so.}
````

Secure development design content rules:

- Include only applicable narrative security category rules for new active designs. Do not require all security controls, all Source IDs, or per-row `N/A` rationale for omitted rows.
- Treat `## Secure Development Design` as the active design-time security planning authority for changed-surface classification, applicable category rules, evidence owners, residual risks, exceptions, and safe-evidence policy.
- Do NOT emit security YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, all-row `N/A` bookkeeping, the general 96-control `sdd-review` matrix, or the exhaustive 155-row Source ID matrix in design.
- Omitted categories are reviewable omissions. Canonical `review-security-report.json` owns compact and Source ID validation coverage, `Yes`/`No`/`N/A` decisions, non-applicable evidence, missed-applicable blockers, and exact-once Source ID expansion/validation; derived Markdown presents those JSON facts for compatibility.
- Applicable category rules MUST identify secure design decisions/controls, downstream evidence owners, expected evidence, residual risks, and archive expectations in prose.
- No-impact designs MUST still record changed-surface rationale, touched/untouched runtime surfaces, and why no security category applies. They MUST NOT use all-row `N/A` bookkeeping to prove no-impact status.
- Exceptions MUST include approver, approvedAt, accepted-risk rationale, mitigation/follow-up, and evidence gap. Incomplete exceptions do not satisfy archive readiness.
- Evidence MUST be review-safe: cite paths, sections, summaries, command outputs, or redacted placeholders; never include raw secrets, credentials, PAN, PII, tokens, private keys, connection strings, or confidential values.

### Step 4: Validate Design

Before persisting or returning, verify the design as a gate, not as another source of exhaustive security matrices:

General design checks:

- Every architecture decision has a rationale.
- File changes use concrete paths, or explicitly mark paths as new/proposed.
- Testing strategy matches detected testing capabilities or explains unavailable layers.
- Migration / Rollout states `No migration required.` when not applicable.
- Blocking open questions set the return status to `blocked`.

Secure design shape checks:

- `## Secure Development Design` is present.
- The section records changed-surface classification and applicable category rules in narrative form.
- The section does not include security YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, all-row `N/A` bookkeeping, the general 96-control review matrix, or the exhaustive 155-row Source ID matrix.
- When corporate source-row review applies downstream, the secure design cites only high-level category context in prose and leaves exact-once compact/Source ID validation plus `N/A` decisions to canonical `review-security-report.json`.
- Omitted categories are clearly treated as reviewable omissions for canonical `review-security-report.json`, not as passing design-time `N/A` rows.

Classification and evidence checks:

- Security classification is evidence-backed; unknown data sensitivity, auth/session/access boundaries, secret/config handling, file/database behavior, or sensitive logging implications block success rather than being guessed.
- No-impact classification includes changed-surface rationale, touched/untouched runtime surfaces, and evidence explaining why no security category applies.
- Applicable category rules include controls, evidence owners, expected evidence, residual risk, and archive expectations.
- Every applicable category rule and compact guideline referenced in prose has a downstream evidence owner in `test-design`, `tasks`, `apply`, `review-security`, `verify`, or `archive`.
- Exceptions include complete exception fields: approver, approvedAt, accepted-risk rationale, mitigation/follow-up, and evidence gap.
- Safe-evidence rules are followed for `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, and `SEC-LOG-001`.
- `## Operational Considerations` is present when proposal/spec/code context makes operational behavior applicable, or absence/inapplicability is justified by the design context.
- Every applicable operational concern uses safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.`; unsupported or unsafe details are not invented.
- Monitoring evidence is mechanism-oriented and not SQL-only by default.
- Restricted operational data and final-document-only values are excluded from ordinary SDD evidence.
- Keep narrative design content compact. The 800-word budget is a preference, not permission to omit required classification rationale, applicable category rules, exception data, or evidence obligations.

### Step 5: Persist Artifact

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `design`
- topic_key: `sdd/{change-name}/design`
- openspec path: `openspec/changes/{change-name}/design.md`
- type: `architecture`

After persistence succeeds, route directly to `sdd-test-design` / `test-design`. Do not produce a standalone security-design artifact as part of the new-change DAG.

### Step 6: Return Summary

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Use the canonical YAML format defined in the Section D example. Adapt these fields for this phase:

- `phase`: `design`
- `status`: MUST be `success`, `partial`, or `blocked`. **Never** `completed`, `done`, `ok`, or any other value.
- `next_recommended`: `test-design` (success) | `resolve-blockers` (blocked/partial)
- `executive_summary`: one short paragraph — technical approach, number of key decisions, files affected count, security classification (impacting/no-impact), and whether open questions were resolved
- `artifacts`: one entry per artifact produced; type `design`; correct mode, ref/path, `persisted: true`, `readable: true`
- `risks`: structured array or `None` — never an empty array `[]`
- `skill_resolution`: from `skill-resolver.md#step-4-report-resolution`
- `detailed_report`: use this minimum content for `sdd-design`:

```
## Design Created

**Change**: {change-name}
**Location**: `openspec/changes/{change-name}/design.md` (openspec) | Engram `sdd/{change-name}/design` (engram)

### Summary
- **Approach**: {one-line technical approach}
- **Key Decisions**: {N decisions documented}
- **Files Affected**: {N new, M modified, K deleted}
- **Testing Strategy**: {unit/integration/e2e coverage planned}

### Open Questions
{List any unresolved questions, or "None"}

### Next Step
Ready for test design (sdd-test-design).
```

## Rules

- Read real code before making claims about existing behavior; never guess about the codebase.
- For new changes, keep security classification and secure design evidence inside `design.md#secure-development-design`; the next active phase after successful design is `sdd-test-design`.
- Every decision MUST have a rationale (the "why")
- Include concrete file paths, not abstract descriptions
- Use the project's ACTUAL patterns and conventions, not generic best practices
- If you find the codebase uses a pattern different from what you'd recommend, note it but FOLLOW the existing pattern unless the change specifically addresses it
- Keep ASCII diagrams simple — clarity over beauty
- Apply any `rules.design` from `openspec/config.yaml`
- On successful design only, return `next_recommended: test-design`; blocked or partial results MUST route to `resolve-blockers` unless the same design artifact can be safely retried without new user input.
- If you have open questions that BLOCK the design, say so clearly — don't guess
- **Size budget**: Keep narrative sections around 800 words when practical. Architecture decisions as tables (option | tradeoff | decision). Code snippets only for non-obvious patterns. Do not omit applicable category rules, evidence obligations, or exception data to satisfy the narrative budget.
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the design summary belongs in `detailed_report`.
