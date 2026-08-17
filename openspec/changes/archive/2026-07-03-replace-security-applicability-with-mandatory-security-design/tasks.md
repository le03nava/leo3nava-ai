# Tasks: Mandatory Security Design for SDD Changes

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 900-1400 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR1 contracts -> PR2 phases -> PR3 validators/evidence -> PR4 docs/specs |
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
| 1 | Shared status/persistence/OpenSpec/security contracts | PR 1 | Base `main`; TD-011, TD-013, TD-021, SEC-ACCESS |
| 2 | Orchestrator and phase skills/agents | PR 2 | Stacked on PR 1; TD-001..TD-010, TD-014, TD-022 |
| 3 | Validators plus static/manual evidence | PR 3 | Stacked on PR 2; TD-016..TD-020, TD-023..TD-028 |
| 4 | README/source specs and final readiness | PR 4 | Stacked on PR 3; TD-015, archive evidence |

## Phase 1: Contracts (PR 1)

- [x] 1.1 Update `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md`, and state examples with `securityDesign`, `securityReviewReport`, `review-security`, `apply -> review -> review-security -> verify`, and legacy-only `securityApplicability` (TD-011, TD-021, SEC-ACCESS-001).
- [x] 1.2 Update `skills/_shared/openspec-convention.md`, `skills/_shared/sdd-security-contract.md`, and `skills/_shared/security-guideline-catalog.md` for mandatory `security-design.md`, `review-security-report.md`, matrix values/statuses, evidence, exceptions, N/A rationale, and SEC-DATA/SECRET/ACCESS/LOG safe-evidence rules (TD-004, TD-009, TD-013, TD-023..TD-027).

## Phase 2: Routing and Phases (PR 2)

- [x] 2.1 Update `agents/sdd/sdd-orchestrator.md` for the new DAG, readiness/native routing, `/sdd-ff`, launch envelopes, and stacked apply handoff (TD-001, TD-002, TD-025).
- [x] 2.2 Update `agents/sdd/sdd-design.md` and `skills/sdd-design/SKILL.md` to require proposal/specs only, include baseline security considerations, and always route to `security-design` (TD-003).
- [x] 2.3 Update `agents/sdd/sdd-security-design.md` and `skills/sdd-security-design/SKILL.md` to always persist/read back `security-design.md` with classification, catalog matrix, controls, exceptions, validation metadata, and `nextRecommended: test-design` (TD-004, SEC-DATA/SECRET/ACCESS/LOG).
- [x] 2.4 Deprecate `agents/sdd/sdd-security-applicability.md` and `skills/sdd-security-applicability/SKILL.md` as legacy/archive-only; no new-change artifact or route (TD-002, TD-019).
- [x] 2.5 Create `agents/sdd/sdd-review-security.md` and `skills/sdd-review-security/SKILL.md`; consume `security-design.md` plus `review-report.md`, validate rows, persist `review-security-report.md`, and route blockers/non-blockers (TD-008..TD-010, TD-022, SEC-LOG-001).
- [x] 2.6 Update downstream `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-review/SKILL.md`, `skills/sdd-verify/SKILL.md`, and `skills/sdd-archive/SKILL.md` for mandatory security-design evidence, review-security routing, unavailable-tooling reporting, and archive gates (TD-005..TD-007, TD-012, TD-028).

## Phase 3: Validators and Evidence (PR 3)

- [x] 3.1 Create `scripts/validate_security_design.ps1` for schema/version/changeName, catalog snapshot, SEC IDs, taxonomy, `Yes/No/N/A`, lifecycle statuses, evidence/exception/archive fields, and `nextRecommended: test-design` (TD-016..TD-018).
- [x] 3.2 Mark `scripts/validate_security_applicability.ps1` legacy/archive-only and capture evidence for stale active routing, direct `design -> test-design`, direct `review -> verify`, unsafe sensitive examples, secret-looking literals, N/A rationale, safe audit evidence, and unavailable PowerShell/runtime tooling (TD-019..TD-028).

## Phase 4: Docs and Readiness (PR 4)

- [x] 4.1 Update `README.md` and source `openspec/specs/*` for mandatory security-design/review-security workflow and compatibility model (TD-015).
- [x] 4.2 Record apply/verify evidence covering TD-001..TD-028, SEC-DATA-001/SEC-SECRET-001/SEC-ACCESS-001/SEC-LOG-001, no runtime test/lint/type/format/coverage commands, and stacked-to-main boundaries.
