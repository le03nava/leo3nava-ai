# Tasks: Inline Secure Development Design in `design.md`

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 700-1,000 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR1 contracts -> PR2 consumers -> PR3 docs/specs |
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
| 1 | Routing/state/shared contracts | PR 1 | base `main`. |
| 2 | Phase consumers/review-security | PR 2 | stacked on PR 1. |
| 3 | Validator, docs, specs, evidence | PR 3 | stacked on PR 2. |

## Phase 1: Work Unit 1 - Routing and Shared Contracts

- [x] 1.1 Update `agents/sdd/sdd-orchestrator.md`, `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md`: DAG `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`; `securityDesign` refs legacy/read-only. Covers TD-008, TD-011, TD-012.
- [x] 1.2 Update `agents/sdd/sdd-design.md`, adapter prompt copies, `skills/sdd-design/SKILL.md`: require `design.md#secure-development-design`, 8 SEC IDs, row/exception fields, N/A rationale, catalog snapshot, lifecycle vocabulary, safe evidence, direct `test-design`. Covers TD-001..TD-007, TD-023..TD-027, TD-034.
- [x] 1.3 Update shared security/catalog contracts for all SEC rows: owners, controls, evidence, residual risks, archive expectations. Covers TD-002..TD-006, TD-025..TD-027.

## Phase 2: Work Unit 2 - Phase Consumers and Gates

- [x] 2.1 Update `skills/sdd-test-design/SKILL.md`: consume embedded rows, block if missing, plan spec/SEC checks without standalone `security-design.md`. Covers TD-028..TD-031.
- [x] 2.2 Update `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md`: read embedded obligations, keep legacy readability, report unavailable runtime/coverage/lint/typecheck/format tools. Covers TD-013, TD-015, TD-017, TD-030, TD-031, TD-035.
- [x] 2.3 Update `skills/sdd-review/SKILL.md`: require proposal/specs/design/test-design/tasks/apply evidence, route non-blocking review to `review-security`, include changed-file/security handoff. Covers TD-009, TD-010, TD-019..TD-021.
- [x] 2.4 Update `skills/sdd-review-security/SKILL.md`: validate embedded rows against `skills/_shared/security-guideline-catalog.md`; persist/read back `review-security-report.md`; mark missing evidence `No`/`blocked`; justify N/A rows. Covers TD-014..TD-018, TD-032, SEC-ACCESS-001, SEC-LOG-001.

## Phase 3: Work Unit 3 - Validator, Docs, Specs, and Evidence

- [x] 3.1 Delete `scripts/validate_security_design.ps1`; if retained, mark archive-only and remove active status/review-security/verify/archive invocation. Covers TD-013.
- [x] 3.2 Update `README.md`, skill registry/status refs, downstream refs, and legacy/archive compatibility text. Covers TD-012, TD-022, TD-032, TD-033.
- [x] 3.3 Update source specs under `openspec/specs/` for design, execution persistence, review, review-security, security-design, security-guideline-catalog, and test-design. Covers TD-001..TD-035.
- [x] 3.4 Produce static/manual evidence: grep/read-back TD-001..TD-035, cite unavailable tooling from `openspec/config.yaml`, scan for raw secrets/sensitive values. Covers TD-032..TD-035.
