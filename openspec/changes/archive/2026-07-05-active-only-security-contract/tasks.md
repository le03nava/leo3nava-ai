# Tasks: Active-Only Security Contract

## Review Workload Forecast

| Field | Value |
|---|---|
| estimated_changed_lines | 650-900 |
| review_budget_lines | 400 |
| review_budget_risk | High |
| chained_prs_recommended | Yes |
| decision_needed_before_apply | No |
| delivery_strategy | auto-chain |
| chain_strategy | stacked-to-main |
| size_exception | none |
| rationale | Markdown/spec contract cleanup spans shared contracts, phase skills, status/persistence boundaries, source specs, and static/manual evidence. |
| work_unit_boundaries | PR1 shared security contract/catalog; PR2 design/review-security phase docs; PR3 status/persistence/spec sync and verification evidence. |

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
| 1 | Make shared security authority active-only. | PR 1 -> main | Modify `skills/_shared/sdd-security-contract.md` and `skills/_shared/security-guideline-catalog.md`; verify TD-001..TD-004, TD-019. |
| 2 | Align phase guidance to embedded design and review-security. | PR 2 -> main after PR 1 | Modify `skills/sdd-design/SKILL.md` and `skills/sdd-review-security/SKILL.md`; verify TD-005..TD-009, TD-014..TD-015, TD-020. |
| 3 | Preserve historical readers and sync specs/evidence. | PR 3 -> main after PR 2 | Modify `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md`, `openspec/specs/sdd-*/spec.md`; verify TD-010..TD-013, TD-016..TD-018, TD-021. |

## Trace

- Proposal: active current-flow contract, no separate legacy contract, no archive rewrite.
- Specs: `sdd-security-guideline-catalog`, `sdd-review-security-workflow`, `sdd-security-applicability-workflow`, `sdd-execution-persistence-contracts`, `sdd-design-workflow`.
- Design: active authority, historical readability, positive cleanup style.
- Test design: mandatory static/manual TD-001..TD-020; advisory TD-021.

## Phase 1: Shared Contract Slice

- [x] 1.1 Update `skills/_shared/sdd-security-contract.md` to remove active standalone `security-design.md` / `security-applicability.md` schema and validator requirements; keep embedded design, review-security report, exceptions, lifecycle/status vocabulary, and safe-evidence rules (TD-001, TD-002).
- [x] 1.2 Update `skills/_shared/security-guideline-catalog.md` to preserve snapshot IDs, all 8 SEC IDs, lifecycle/status vocabulary, and catalog boundary favoring `design.md#secure-development-design` plus `review-security-report.md` (TD-003, TD-004).
- [x] 1.3 Static-check changed shared files for raw secrets, tokens, private keys, PAN, PII, or confidential payloads; report summarized safe-evidence results only (SEC-DATA-001, SEC-SECRET-001, SEC-LOG-001, TD-019).

## Phase 2: Active Phase Guidance Slice

- [x] 2.1 Update `skills/sdd-design/SKILL.md` so `## Secure Development Design` is the active authority, includes all 8 SEC IDs, and routes directly to `sdd-test-design` without standalone security-design production (TD-014, TD-015).
- [x] 2.2 Update `skills/sdd-review-security/SKILL.md` so `review-security-report.md` validates embedded design rows, blocks missing embedded design, and does not require `scripts/validate_security_design.ps1` or standalone artifacts for new changes (TD-005..TD-007).
- [x] 2.3 Search repo-local `skills/`, active workflow docs, and agent config to confirm no active `sdd-security-applicability` executor/skill or new-change DAG route is offered (TD-008, TD-009).
- [x] 2.4 Manual read-back changed phase docs for positive current-flow wording instead of legacy-warning framing (TD-020).

## Phase 3: Reader Boundary, Specs, and Evidence Slice

- [x] 3.1 If needed, update `skills/_shared/persistence-contract.md` and `skills/_shared/sdd-status-contract.md` so historical `securityDesign` / `securityApplicability` refs are read-only data and never runnable phases, dependencies, or active authority (TD-010, TD-012, TD-013).
- [x] 3.2 Sync source specs under `openspec/specs/sdd-*/spec.md` to match the five delta specs and remove active applicability override/static-validator requirements (TD-011, proposal success criteria).
- [x] 3.3 No-archive-rewrite check: before and after apply, inspect changed paths and confirm nothing under `openspec/changes/archive/**` changed (TD-016).
- [x] 3.4 Confirm no separate legacy security contract file was created, including candidate names like `legacy-security-contract` or `security-design-contract` (TD-017).
- [x] 3.5 Verification evidence task: report unavailable runtime tests, lint, type-check, format, and coverage from `openspec/config.yaml#testing`; use static/manual evidence instead (TD-018).
- [x] 3.6 Advisory check: confirm rollback remains git revert of active contract/spec edits with no migration, runtime, or archive repair steps (TD-021).
