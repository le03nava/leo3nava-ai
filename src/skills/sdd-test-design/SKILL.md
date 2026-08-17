---
name: sdd-test-design
description: "Create the SDD test design artifact between technical design and task planning. Trigger: orchestrator launches test-design for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-test-design` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are a sub-agent responsible for TEST DESIGN. You take the proposal, specs, technical design with mandatory narrative `design.md#secure-development-design`, then produce `test-design.md` that maps scenarios, design risks, applicable security category rules, and behavior contracts to planned automated, manual, or static checks before task planning begins.

## What You Receive

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.execution_mode` — controls blocking behavior
- `launch.status.dependencies` — `proposal: completed`, `spec: completed`, and `design: completed` must be present
- `launch.artifacts.paths.proposal` or `launch.artifacts.refs.proposal` — do NOT expect content inline
- `launch.artifacts.paths.spec` or `launch.artifacts.refs.spec` — do NOT expect content inline
- `launch.artifacts.paths.design` or `launch.artifacts.refs.design` — must include `## Secure Development Design`; do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute workspace path
- `launch.skill_paths` — supplemental skills to load before work

If `launch.artifact_store.mode` is absent, or `launch.status.dependencies` does not show `proposal`, `spec`, and `design` as completed, return `blocked` with `next_recommended: resolve-blockers`.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | Proposal, specs, design with mandatory `## Secure Development Design`, and testing capabilities from the selected backend. Standalone `security-design.md` is legacy/read-only compatibility data only. |
| Produced artifact | Mandatory canonical `sdd/{change-name}/test-cases` / `openspec/changes/{change-name}/test-cases.json` plus derived `sdd/{change-name}/test-design` / `openspec/changes/{change-name}/test-design.md` for every change, including no-impact changes. |
| Mutates | None outside the produced test design artifact. |
| Mandatory artifact behavior | Do not route directly from design to tasks without a complete `test-design` artifact. No-impact changes still produce a concise no-impact assessment rather than omitting the artifact. |
| Planned evidence mapping | Preserve scenario, design-risk, applicable security category rule, check type, severity, expected evidence, no-impact assessment, and open-question mapping. Mandatory cases are verification-blocking when uncovered. |
| Operational considerations planning | Consume `design.md#Operational Considerations` or equivalent design evidence when present, and plan static, documentary, manual, or automated checks only for applicable operational evidence/gaps. If design omits operational considerations or marks them `No aplica.`, do not synthesize mandatory operational categories. |
| Source-row planning | When `design.md#secure-development-design` describes corporate security context, plan static/manual/automated checks from applicable narrative category rules and changed-surface context only. Test-design MUST NOT require design to carry YAML, schema fields, compact controls, Source IDs, matrices, machine-readable applicability fields, or exhaustive `N/A` rows; it MUST preserve `review-security` as the owner that validates every Source ID exactly once and reports coverage/focused findings. |
| Unavailable tooling | If runtime, build, coverage, lint, typecheck, or format tooling is unavailable, record explicit unavailable-tooling notes and plan static/manual evidence instead. Missing tooling is never passing evidence. |
| Downstream consumption | `sdd-tasks`, `sdd-apply`, `sdd-verify`, and archive readiness checks consume `test-cases.json` as canonical case data; `test-design.md` is a derived human-readable view only. |
| Success routing | `next_recommended: tasks`. |
| Block routing | `next_recommended: resolve-blockers` for missing proposal/spec/design, missing embedded secure development design, testability blockers, or persistence failure. Do not route new changes to standalone `security-design`. |

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the test-design summary in `detailed_report`.

Routing rules for `next_recommended`:
- **Successful test design**: return `next_recommended: tasks`. The orchestrator normalizes this into state `nextRecommended: tasks` before routing or persisting state.
- **Blocked test design**: return `next_recommended: resolve-blockers` and include the exact missing proposal, spec, design, testability decision, or artifact validation issue in `risks` / `detailed_report`.
- **Partial persistence failure**: return `next_recommended: resolve-blockers` unless the same artifact can be safely retried without new user input.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for status/state artifacts only.

## Decision Gates

| Situation | Action |
| --- | --- |
| Required proposal, spec, or design is missing | Return `blocked` with `next_recommended: resolve-blockers`; do not write test design. |
| `design.md#secure-development-design` is missing for a new active change | Return `blocked` with `next_recommended: resolve-blockers`; name the missing embedded section and do not write test design. |
| Standalone `security-design.md` is missing for a new active change | Continue; do not require it. It is legacy/read-only compatibility data only. |
| Specs/design have no behavior or testability impact | Write a no-impact assessment in `test-design.md`; do not treat the artifact as absent. |
| A mandatory spec scenario or design risk has no planned check and no justified skip | Return `blocked` or fix the draft before persistence. |
| A mandatory applicable narrative security rule has no planned check, non-test evidence, or complete approved exception | Return `blocked` with `next_recommended: resolve-blockers`; uncovered mandatory security rules cannot proceed to tasks. |
| Design includes applicable operational considerations but corresponding checks are missing | Fix the draft before persistence; applicable operational evidence cannot be left unplanned. |
| Missing runtime/build/lint/type/format/coverage tooling is treated as a pass | Fix the draft before persistence; unavailable tooling requires static/documentary/manual substitute evidence and explicit notes. |
| Test-design attempts to require design YAML, schema fields, compact controls, Source IDs, matrices, machine-readable applicability fields, or exhaustive `N/A` rows | Fix the draft before persistence; these are review-security-owned concerns, not test-design prerequisites. |
| An applicable narrative rule lacks planned safe evidence | Return `blocked` with `next_recommended: resolve-blockers`; unsupported applicable coverage cannot proceed to tasks. |
| Test-design draft fails validation | Fix it before persistence; if it cannot be fixed, return `blocked` with `next_recommended: resolve-blockers`. |

## What to Do

### Step 1: Load Skills

Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.

### Step 2: Read Required Inputs

Before writing the artifact, read:
- Proposal: user intent, scope, non-goals, risks, and success criteria.
- Specs: requirements and scenarios that need coverage.
- Design: architecture decisions, data flow, file changes, contracts, and testing strategy.
- Embedded secure development design: `design.md#secure-development-design` changed-surface classification, applicable narrative category rules, required controls, mandatory evidence expectations, residual risks, safe-evidence policy, and approved exceptions. Historical exhaustive `N/A` rows may be read for compatibility, but new active test-design MUST NOT require them.
- Operational considerations design: `design.md#Operational Considerations` or equivalent design evidence for strategy, evidence plan, exact marker usage, restricted-data boundary, monitoring mechanisms, unresolved gaps, owner phases, and manual operational document handoff when applicable.
- Testing capabilities when available:
  - Engram: `sdd/{project}/testing-capabilities`
  - OpenSpec: `openspec/config.yaml` `testing` section

### Step 3: Identify Test Design Inputs

Collect planned checks from:
- Spec scenarios and RFC 2119 requirements.
- Design risks, compatibility decisions, routing/state/persistence contracts, migrations, and rollout notes.
- Applicable embedded secure-development narrative rules, mandatory evidence expectations, carried risks, and archive-gate notes.
- Changed-surface context and applicable narrative category rules from `design.md#secure-development-design`; do not parse or require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, or exhaustive `N/A` rows.
- Operational considerations and handoffs from design evidence: logs/error evidence, monitoring mechanisms, administration, reprocessing/recovery, ownership, backup/retention/cleanup/generated artifacts, final-document inputs, exact markers, safe evidence, and unresolved gaps when applicable.
- Testing capability constraints such as unavailable runners, missing coverage tooling, or static-only repositories.

When runtime test runner, coverage, linter, type checker, or formatter commands are unavailable, plan static/manual evidence explicitly. Missing tooling is a reported constraint, not passing evidence.

Operational checks MUST include static/documentary/manual validation for exact `Pendiente de confirmar:` and `No aplica.` markers, evidence traceability, non-SQL-only monitoring mechanisms, restricted operational data absence, separation of ordinary SDD evidence from final operational document inputs, and unavailable-tooling carry-forward when the design makes those concerns applicable.

Mandatory narrative-rule evidence blockers MUST stay visible: unsafe evidence or missing evidence for an applicable mandatory category rule blocks test-design. Exhaustive omitted-row and `N/A` validation belongs to canonical `review-security-report.json`; warning-only coverage remains tracked as warnings and may route forward only when mandatory evidence is complete.

If there is no behavior or testability impact, write a concise no-impact assessment instead of inventing checks.

### Step 4: Write test-cases.json and derive test-design.md

**IF mode is `openspec`:** Create both artifacts under the change folder:

```text
openspec/changes/{change-name}/
├── proposal.md
├── specs/
├── design.md
├── test-cases.json         ← You create this first (canonical)
└── test-design.md          ← You derive this from JSON (human-readable view)
```

**IF mode is `engram`:** Do NOT create any `openspec/` directories or files. Compose both artifacts in memory; persist in Engram.

#### Step 4a: Write canonical test-cases.json first

Create `test-cases.json` as the canonical source of truth. Every generated case MUST start with `status: pending`.

Use this schema contract inline:

```json
{
  "schemaName": "sdd.test-cases",
  "schemaVersion": 1,
  "changeName": "{change-name}",
  "generatedBy": "sdd-test-design",
  "updatedBy": "sdd-test-design",
  "updatedAt": "ISO-8601",
  "cases": [
    {
      "id": "TC-001",
      "description": "Planned check description",
      "type": "automated | manual | static",
      "mandatory": true,
      "status": "pending | blocked | applied | verified | skipped | warning",
      "evidence": null,
      "spec": "spec-domain: Scenario reference",
      "category": "category-label"
    }
  ]
}
```

Required per-case fields: `id`, `description`, `type`, `mandatory`, `status`, `evidence`, `spec`, `category`.

Valid `status` values: `pending | blocked | applied | verified | skipped | warning`.

#### Step 4b: Derive test-design.md from test-cases.json

Generate `test-design.md` from the persisted JSON cases table. `test-design.md` is a human-readable projection and is not authoritative case data.

#### Test Design Document Format

```markdown
# Test Design: {Change Title}

## Overview

{One paragraph summarizing the behavior/testability surface and testing constraints.}

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | {path-or-topic} | {scope/intent summary} |
| Spec | {path-or-topic} | {requirements/scenarios summary} |
| Design | {path-or-topic} | {architecture/risk summary} |
| Secure Development Design | {design path-or-topic}#secure-development-design | {changed-surface classification, applicable narrative rules, evidence obligations, residual risks, and exceptions} |
| Testing Capabilities | {path-or-topic} | {available and unavailable runtime/build/coverage/lint/type/format tooling} |

## Source ID Coverage Baseline

{Required when corporate source-row coverage applies. State changed-surface context, applicable narrative category rules, unavailable tooling, and whether validation uses static/manual/automated evidence. Do not require design schema fields, duplicate the full 155-row matrix, or create exhaustive `N/A` rows here.}

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: {requirement/scenario} | {planned check} | automated/manual/static | mandatory/non-mandatory | {command, artifact, file-contract check, or manual evidence} | {constraints/dependencies} |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| {operational concern} | {marker/traceability/safe-evidence/monitoring/final-doc-boundary check} | static/manual/documentary | {path, section, sanitized summary, marker, or N/A rationale} | {runtime/build/lint/type/format/coverage unavailable note when applicable} |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-...` | {rule from design.md#secure-development-design} | Yes/No | {test case ID, manual/static evidence, or complete exception} | covered/blocked/not-applicable | {None or complete approved exception} |

## No-Impact Assessment

{Required only when no behavior or testability impact exists. State why no checks are needed.}

## Evidence Expectations

- Mandatory cases require implementation, execution, static/manual evidence, or a justified skip.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Security validation evidence should cite embedded `design.md` narrative rules, owner phase, and planned static/manual evidence.
- Applicable narrative category rules require planned safe evidence. Exhaustive source-row validation coverage, `N/A` decisions, and missed-applicable validation remain owned by canonical `review-security-report.json`.
- Warning-only source coverage must be preserved with expected observation evidence and may proceed only when mandatory evidence is complete.
- Test-design consumes narrative design rules only and MUST NOT require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, or the full 155-row Source ID matrix; exhaustive validation coverage belongs to canonical `review-security-report.json`, while full matrix output is audit-only unless explicitly requested.
- No-impact routing is valid only when justified by changed-surface classification inside mandatory `design.md#secure-development-design`; absence of standalone `security-design.md` is not a blocker for new changes.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands that are unavailable must be reported as unavailable evidence, not treated as passed checks.

## Open Questions

- [ ] {Only questions that block reliable test planning. Say "None" if there are no blockers.}
```

Valid `Type` values: `automated`, `manual`, `static`.
Valid `Severity` values: `mandatory`, `non-mandatory`.

### Step 5: Validate Test Design

Before persisting or returning, verify:
- Every behavior-impacting spec scenario or design risk has a linked test case, or a justified omission.
- `design.md#secure-development-design` is present and every mandatory applicable narrative security rule has a planned check, justified non-test evidence, or complete approved exception.
- When source-row coverage applies, every applicable narrative category rule has a planned static/manual/automated check or justified non-test evidence; missing safe evidence remains blocking.
- Uncovered mandatory security evidence is a blocker; do not persist a successful `test-design.md` that leaves mandatory controls uncovered.
- Warning-only rows are preserved as warning evidence and are not silently dropped.
- Each test case has `ID`, `Source`, `Check`, `Type`, `Severity`, `Expected Evidence`, and `Notes`.
- `Type` is one of `automated`, `manual`, or `static`.
- `Severity` is one of `mandatory` or `non-mandatory`.
- Mandatory cases are verification-blocking when uncovered.
- Non-mandatory cases are advisory and become verification warnings when uncovered.
- Static/manual checks are allowed when no runtime test runner exists.
- Unavailable runtime/build/coverage/lint/typecheck/format tooling is reported explicitly and never claimed as passing evidence.
- No-impact changes include the no-impact assessment and still produce the artifact.
- Operational considerations checks cover exact markers, traceability, safe evidence, monitoring mechanism coverage, final-document boundary, and restricted-data absence when design evidence makes those concerns applicable.
- Blocking open questions set the return status to `blocked`.

### Step 6: Persist Artifact

**This step is MANDATORY for `engram` and `openspec` modes — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- Persist canonical artifact first:
  - artifact: `test-cases`
  - topic_key: `sdd/{change-name}/test-cases`
  - openspec path: `openspec/changes/{change-name}/test-cases.json`
  - type: `architecture`
- Then persist derived artifact:
  - artifact: `test-design`
  - topic_key: `sdd/{change-name}/test-design`
  - openspec path: `openspec/changes/{change-name}/test-design.md`
  - type: `architecture`

### Step 7: Return Summary

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Use the canonical YAML format defined in the Section D example. Adapt these fields for this phase:

- `phase`: `test-design`
- `next_recommended`: `tasks` (success) | `resolve-blockers` (blocked/partial)
- `executive_summary`: one short paragraph — total cases planned (mandatory/non-mandatory), automated/manual/static split, security posture (no-impact or applicable categories covered), and testing constraints summary
- `artifacts`: two entries — canonical `test-cases` (JSON) first, then derived `test-design` (MD); both with correct mode, ref/path, `persisted: true`, `readable: true`
- `risks`: structured array or `None` — never an empty array `[]`
- `skill_resolution`: from `skill-resolver.md#step-4-report-resolution`
- `detailed_report`: use this minimum content for `sdd-test-design`:

```
## Test Design Created

**Change**: {change-name}
**Location**: `openspec/changes/{change-name}/test-design.md` (openspec) | Engram `sdd/{change-name}/test-design` (engram)
**Canonical**: `openspec/changes/{change-name}/test-cases.json` (openspec) | Engram `sdd/{change-name}/test-cases` (engram)

### Summary
- **Inputs**: {Proposal, N specs, and design read. Embedded design.md#secure-development-design consumed.}
- **Security Inputs**: {No-impact classification confirmed / Applicable categories: list}
- **Cases Planned**: {N mandatory, M non-mandatory — X automated, Y manual, Z static}
- **No-Impact Assessment**: {Not applicable (behavioral changes present; security posture is no-impact) / present}
- **Testing Constraints**: {list detected runner/static/manual constraints}

### Open Questions
{List any unresolved questions, or "None"}

### Next Step
Ready for tasks (sdd-tasks).
```

## Rules

- ALWAYS read proposal, specs, and design before writing test design.
- ALWAYS read mandatory `design.md#secure-development-design` before writing test design; standalone `security-design.md` is legacy/read-only only.
- ALWAYS generate `test-cases.json` first; derive `test-design.md` from that JSON. `test-design.md` is read-only for humans and MUST NOT be used as source of truth by downstream phases.
- ALWAYS cover each mandatory applicable narrative security rule before returning success. Do not require justified `N/A` rows from design/test-design for omitted catalog rows.
- DO NOT implement tests or code; this phase plans evidence only.
- DO NOT skip the artifact for no-impact changes; document the no-impact assessment.
- Prefer scenario-linked, evidence-focused checks over broad testing wishes.
- Match testing capabilities; do not invent unavailable runners or commands.
- Apply any `rules.test-design` from `openspec/config.yaml` if present.
- Return the Section D envelope per `skills/_shared/sdd-phase-common.md`; the test-design summary belongs in `detailed_report`.
