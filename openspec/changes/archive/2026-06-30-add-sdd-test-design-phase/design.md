# Design: Add Mandatory SDD Test Design Phase

## Technical Approach

Add `sdd-test-design` as a first-class SDD executor between `sdd-design` and `sdd-tasks`. The phase writes `test-design.md`, a compact plan mapping spec scenarios/design risks to automated, manual, or static checks. Because this repository has no runtime test runner in `openspec/config.yaml`, verification for this change will rely on static/file-contract checks, artifact presence, and manual evidence descriptions.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
|---|---|---|---|
| Phase boundary | Create `agents/sdd/sdd-test-design.md` and `skills/sdd-test-design/SKILL.md` | Fold test cases into `design.md` | The spec requires a mandatory artifact before tasks; a separate executor keeps design HOW separate from test evidence planning. |
| Token naming | Use `test-design` as native/status token, `sdd-test-design` as phase/agent token, and `testDesign` for camelCase persisted state/artifact fields. | Use only `sdd-test-design`; use only `testDesign` | Matches existing native-token vs phase-token split while preserving camelCase state conventions. If native `gentle-ai` cannot accept `test-design`, prompts must document compatibility limits and route through prompt-level fallback until native support exists. |
| Artifact contract | Store as Engram `sdd/{change-name}/test-design` and OpenSpec `openspec/changes/{change-name}/test-design.md`. | Reuse `design` or `tasks` artifact keys | Keeps persistence resolver, state refs, and archive audit trail explicit. |
| Coverage semantics | Mandatory uncovered cases block/fail verify; non-mandatory uncovered cases warn. | Treat all uncovered cases as failures | Preserves the user's decision and supports advisory manual/static checks when automation is unavailable. |

## Data Flow

```text
proposal + spec + design
        │
        ▼
sdd-test-design ──→ test-design.md
        │              ├─ scenario/risk links
        │              ├─ check type: automated/manual/static
        │              ├─ severity: mandatory/non-mandatory
        │              └─ expected evidence
        ▼
sdd-tasks ──→ sdd-apply ──→ sdd-verify ──→ sdd-archive
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `agents/sdd/sdd-test-design.md` | Create | Executor prompt mirroring other SDD executors; reads `%USERPROFILE%/.config/opencode/skills/sdd-test-design/SKILL.md`; forbids delegation and `skill()` calls. |
| `skills/sdd-test-design/SKILL.md` | Create | Phase contract: inputs proposal/spec/design, writes `test-design.md`, validates cases, persists by artifact mode, returns `next_recommended: tasks`. |
| `agents/sdd/sdd-orchestrator.md` | Modify | Update description, flowchart, DAG, routing table, launch metadata, state schema, recovery fallback, context requirements, gatekeeper, and status handling for `design -> test-design -> tasks`. |
| `skills/_shared/openspec-convention.md` | Modify | Add `test-design.md` to directory structure, artifact paths, reading rules, and archive contents. |
| `skills/_shared/sdd-phase-common.md` | Modify | Add routing mapping `test-design` ↔ `sdd-test-design` and artifact naming guidance. |
| `skills/_shared/persistence-contract.md` | Modify | Add resolver row, minimum state `artifactRefs.testDesign`, `currentPhase`/`nextRecommended` token, recovery fallback, and sub-agent context rules. |
| `skills/_shared/sdd-status-contract.md` | Modify | Add token mapping, status schema fields, artifact status, dependency state, and context files for `testDesign`. |
| `skills/sdd-design/SKILL.md` | Modify | Successful design routes to `test-design`, not `tasks`. |
| `skills/sdd-tasks/SKILL.md` | Modify | Require `test-design`; derive testing/evidence tasks from it; block if missing. |
| `skills/sdd-apply/SKILL.md` | Modify | Read `test-design`; follow planned checks or document justified deviations in apply-progress. |
| `skills/sdd-verify/SKILL.md` and `skills/sdd-verify/references/report-format.md` | Modify | Compare evidence against `test-design`; fail missing mandatory cases and warn for uncovered non-mandatory cases. |
| `skills/sdd-archive/SKILL.md` | Modify | Include `test-design.md` in required archive artifacts unless an explicit partial archive exception applies. |
| `agents/sdd/sdd-onboard.md`, `skills/sdd-onboard/SKILL.md` | Modify | Update narrated phase order and artifact list. |

## Interfaces / Contracts

`test-design.md` should contain: Overview, Inputs, Test Cases table, No-Impact Assessment when applicable, Evidence Expectations, and Open Questions. Case rows use: `ID`, `Source`, `Check`, `Type`, `Severity`, `Expected Evidence`, `Notes`. Valid type values: `automated`, `manual`, `static`. Valid severity values: `mandatory`, `non-mandatory`.

State/status fields use `testDesign` for artifact refs/paths/context/artifact status. Native/status routing uses `test-design`; launch uses `phase: sdd-test-design`.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static contract | New files, frontmatter, phase order, token mappings, artifact refs, and archive contents | Inspect Markdown files and OpenSpec state/status examples; no runner is available. |
| Workflow coverage | `design -> test-design -> tasks` blocking and recovery paths | Manual/static check against orchestrator DAG, routing table, persistence contract, and status contract. |
| Verification semantics | Mandatory vs non-mandatory uncovered cases | Static check that verify skill/report format defines fail vs warning behavior. |

## Migration / Rollout

No data migration required. Existing active changes without `test-design.md` should be treated as needing the new phase before tasks; archive exceptions must be explicit.

## Open Questions

None.
