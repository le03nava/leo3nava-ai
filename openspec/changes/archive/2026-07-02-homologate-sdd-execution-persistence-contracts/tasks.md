# Tasks: Homologate SDD Execution and Persistence Contracts

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 900-1,300 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 -> PR 2 -> PR 3 -> PR 4 -> PR 5 |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Shared persistence boundary | PR 1 | Base `main`; TD-001, TD-004, TD-005. |
| 2 | Backend conventions and core contracts | PR 2 | Base `main` after PR 1; TD-002, TD-003, TD-006. |
| 3 | Mutation-heavy contracts | PR 3 | Base `main` after PR 2; TD-007, TD-008, TD-009. |
| 4 | Security/test workflow contracts | PR 4 | Base `main` after PR 3; TD-010-TD-014. |
| 5 | Specs/docs and evidence | PR 5 | Base `main` after PR 4; TD-015, TD-016. |

## Phase 1: Shared Contract Boundary

- [x] 1.1 Update `skills/_shared/persistence-contract.md` as authority for modes, resolver, hybrid conflicts, state persistence, and verification; record TD-001.
- [x] 1.2 Trim `skills/_shared/sdd-phase-common.md` persistence prose; preserve executor boundary, skill loading, Section D, routing, naming, and review guard; record TD-004/TD-005.
- [x] 1.3 Confirm `openspec/config.yaml` still has no runtime test/lint/build tooling and no invented commands appear; record TD-015.

## Phase 2: Phase Artifact Contracts

- [x] 2.1 Scope `skills/_shared/engram-convention.md` and `skills/_shared/openspec-convention.md` as backend references only; record TD-003.
- [x] 2.2 Add compact `## Phase Artifact Contract` tables to non-mutating `skills/sdd-*/SKILL.md` files, delegating backend mechanics; record TD-002/TD-006.
- [x] 2.3 Preserve artifact keys, OpenSpec paths, produced artifacts, success routing, and block routing in updated phase skills; record TD-009.

## Phase 3: Mutation-Heavy Contracts

- [x] 3.1 Update `skills/sdd-apply/SKILL.md` with task progress, apply evidence, and deviation semantics; record TD-007/TD-014.
- [x] 3.2 Update `skills/sdd-archive/SKILL.md` with archive moves, spec sync, destructive-delta warnings, and audit-trail semantics; record TD-007/TD-008.
- [x] 3.3 Update `skills/sdd-init/SKILL.md` with initialization and local support artifact mutations preserved; record TD-007/TD-009.

## Phase 4: Security and Test Workflow Compatibility

- [x] 4.1 Update `skills/sdd-security-applicability/SKILL.md` to preserve artifact identity, classification, evidence, mapping, risks, and routing; record TD-010/TD-011.
- [x] 4.2 Update `skills/sdd-security-design/SKILL.md` to preserve conditional `security-design.md` creation and future control/evidence mapping; record TD-011/TD-012.
- [x] 4.3 Update `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, and `skills/sdd-verify/SKILL.md` to preserve mandatory test-design consumption; record TD-013/TD-014.

## Phase 5: Specs and Static Evidence

- [x] 5.1 Update affected `openspec/specs/sdd-*` files without changing DAG, keys, paths, routing tokens, camelCase fields, or backend behavior; record TD-008/TD-009.
- [x] 5.2 Refresh `.atl/skill-registry.md` or `README.md` only if stale references are introduced; otherwise record no-change evidence.
- [x] 5.3 Add apply/verify static evidence notes for TD-001 through TD-016, unavailable tooling, and stacked-to-main boundaries.

## Apply Static Evidence Notes

### Source Spec Synchronization

| Evidence | Result |
| --- | --- |
| New source spec | Added `openspec/specs/sdd-execution-persistence-contracts/spec.md` from the accepted change spec so the execution/persistence boundary has a source-of-truth capability before archive. |
| Modified source specs | Updated `openspec/specs/sdd-security-applicability-workflow/spec.md`, `openspec/specs/sdd-security-design-workflow/spec.md`, and `openspec/specs/sdd-test-design-workflow/spec.md` with the persistence-boundary clauses and scenarios from this change. |
| Compatibility guard | No DAG order, artifact keys, OpenSpec paths, routing tokens, camelCase state/status fields, or backend behavior were renamed or redesigned. Changes are requirement/scenario text only. |
| TD coverage | Covers TD-008 by naming the shared persistence contract as authoritative for duplicated/ambiguous persistence rules; covers TD-009 by preserving established keys, paths, routing tokens, state/status fields, DAG order, and backend behavior. |

### Registry and README Evidence

`.atl/skill-registry.md` exists, but this slice did not change skill names, paths, frontmatter, triggers, or README-visible workflow phase names. The registry is an index of non-SDD project skills by design, and no stale reference was introduced by the source-spec/evidence updates. `README.md` already lists the current SDD phases and phase order, so no registry or README refresh was necessary for Work Unit 5.

### TD-001 through TD-016 Apply/Verify Evidence Handoff

| TD | Static/manual evidence | Status |
| --- | --- | --- |
| TD-001 | `skills/_shared/persistence-contract.md` declares itself authoritative for artifact-store modes, backend read/write semantics, artifact resolution, hybrid conflict handling, state persistence, and persistence verification. | Ready for verify |
| TD-002 | Phase artifact contracts delegate common backend mechanics to `skills/_shared/persistence-contract.md`; no phase-local backend algorithm changes were introduced in this slice. | Ready for verify |
| TD-003 | `skills/_shared/engram-convention.md` and `skills/_shared/openspec-convention.md` remain backend/reference contracts; this slice did not modify them. | Ready for verify |
| TD-004 | `skills/_shared/sdd-phase-common.md` keeps executor boundary, skill loading, Section D envelope, routing-token convention, naming reminders, and review workload guard while pointing retrieval/persistence to the shared persistence contract. | Ready for verify |
| TD-005 | Section D envelope and routing-token mapping remain stable; phase envelopes use `next_recommended`, and persisted state/status uses `nextRecommended`. | Ready for verify |
| TD-006 | Modified SDD phase skills contain compact `## Phase Artifact Contract` sections with required inputs, produced artifacts, mutation/conditional behavior, and routing concerns. | Ready for verify |
| TD-007 | `sdd-apply`, `sdd-archive`, and `sdd-init` retain explicit phase-local mutation semantics for task progress, archive/spec sync, state/local support artifacts, and initialization. | Ready for verify |
| TD-008 | Source specs now identify the authoritative owner for delegated persistence behavior and preserve behavior when duplicated or ambiguous rules are homologated. | Ready for verify |
| TD-009 | Source specs and task evidence preserve established Engram keys, OpenSpec paths, routing tokens, camelCase state/status fields, native state names, DAG order, and backend behavior without migration. | Ready for verify |
| TD-010 | `sdd-security-applicability` preserves `security-applicability.md` artifact identity, classification/evidence/guideline/risk content, artifact-local `nextRecommended`, and routing behavior while delegating backend mechanics. | Ready for verify |
| TD-011 | No-impact applicability continues to skip `security-design.md`; missing `security-design.md` is not a blocker for explicit no-impact changes. | Ready for verify |
| TD-012 | `sdd-security-design` continues to create `security-design.md` only for security-impacting changes and preserves controls/evidence/risk/downstream obligations. | Ready for verify |
| TD-013 | `sdd-test-design` remains mandatory for every change and preserves planned check mapping, evidence expectations, and no-impact assessments. | Ready for verify |
| TD-014 | `sdd-tasks`, `sdd-apply`, and `sdd-verify` continue to consume `test-design.md` for planning, implementation evidence, deviation handling, and verification. | Ready for verify |
| TD-015 | `openspec/config.yaml` records no runtime test runner, linter, formatter, type checker, build, coverage, or test command; verification must use static/manual evidence only. | Ready for verify |
| TD-016 | Tasks are split into five stacked-to-main work units under the configured 400-line review budget; this final slice covers source specs/docs/evidence only. | Ready for verify |

### Unavailable Tooling and Verification Boundary

Runtime test, lint, format, type-check, build, and coverage commands are unavailable in `openspec/config.yaml`; no invented command was run or recorded. Verification should inspect the changed Markdown contracts, this evidence section, and the final diff manually/static-only.

### Stacked-to-Main Boundary

Work Unit 5 / PR 5 starts from `main` after PR 4 and ends with source OpenSpec spec synchronization plus final static evidence. It intentionally does not alter already updated skill contracts unless a stale reference blocks tasks 5.1-5.3; no such stale reference was found.
