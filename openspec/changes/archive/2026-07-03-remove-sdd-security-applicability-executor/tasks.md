# Tasks: Remove SDD Security Applicability Executor

## Review Workload Forecast

| Field | Value |
|---|---|
| estimated_changed_lines | 500-700 additions + deletions |
| review_budget_lines | 400 |
| review_budget_risk | High |
| 400-line budget risk | High |
| chained_prs_recommended | Yes |
| decision_needed_before_apply | No |
| rationale | Deletes two launchable sources and updates multiple shared contracts/specs; slicing protects review focus. |
| work_unit_boundaries | WU1 active repo launch/contracts; WU2 review/source spec sync and final evidence. |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|---|---|---|---|
| 1 | Remove launchable executor/skill and active routing authority. | PR 1 to main | Covers TD-001..TD-008, TD-012, SDD-GOV-001/002. |
| 2 | Update review/source specs and capture verification hygiene. | PR 2 to main after PR 1 | Covers TD-009..TD-015, SDD-GOV-002/003. |

## Phase 1: Remove Active Launch Surface

- [x] 1.1 Delete `agents/sdd/sdd-security-applicability.md`; verify repo-local executor absence with TD-002 and SDD-GOV-001 evidence.
- [x] 1.2 Delete `skills/sdd-security-applicability/SKILL.md` and remove the empty directory if present; verify no replacement launchable skill is added (TD-002).
- [x] 1.3 Update `agents/sdd/sdd-orchestrator.md` so `security-applicability` is legacy/archive data only and never an active DAG successor or launch mapping (TD-001, TD-003, TD-005, SDD-GOV-001).

## Phase 2: Preserve Legacy Data Compatibility

- [x] 2.1 Update `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md`, `skills/_shared/sdd-security-contract.md`, and `skills/_shared/openspec-convention.md` to preserve `securityApplicability` refs only as legacy/read-only archive data (TD-004, TD-006, SDD-GOV-002).
- [x] 2.2 Verify `skills/sdd-security-design/SKILL.md` and related shared contracts keep direct classification from proposal/specs/design and never consult the retired executor or artifact (TD-007, TD-008, SDD-GOV-001).
- [x] 2.3 Confirm `scripts/validate_security_applicability.ps1` remains archive-only and is not introduced as a new-change gate (TD-012, SDD-GOV-002).

## Phase 3: Review and Source Spec Alignment

- [x] 3.1 Update `skills/sdd-review/references/report-template.md` so missing `security-applicability.md` is not a required new-change input and present applicability is optional legacy evidence (TD-009, TD-010).
- [x] 3.2 Sync active source specs under `openspec/specs/**/spec.md` with the delta requirements after implementation, without editing `openspec/changes/archive/**` (TD-011, SDD-GOV-002).
- [x] 3.3 Document repo-local scope and stale `%USERPROFILE%` copy caveat in active contracts/docs without claiming global cleanup (TD-013).

## Phase 4: Static Verification Evidence

- [x] 4.1 Capture targeted search/read-back evidence for deleted launch files, absent new `security-applicability.md`, mandatory `security-design.md`, and no active successor mapping (TD-001..TD-007, TD-009..TD-012).
- [x] 4.2 Inspect changed artifacts for secrets, credentials, production identifiers, sensitive runtime data, and copied environment values (TD-014, SDD-GOV-003).
- [x] 4.3 In apply/verify reports, state runtime tests, coverage, lint, typecheck, format, and build tooling are unavailable rather than passed (TD-015).
