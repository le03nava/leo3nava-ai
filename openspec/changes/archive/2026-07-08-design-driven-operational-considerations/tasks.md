# Tasks: Design-Driven Operational Considerations

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 450-700 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 contract deletion/skills -> PR 2 active specs/evidence |
| Delivery strategy | null |
| Chain strategy | pending |
| Size exception | none |

Decision needed before apply: Yes
Chained PRs recommended: Yes
Chain strategy: pending
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Remove stale readiness authority and update phase/shared skills | PR 1 | Delete contract; update active skills and shared contracts. |
| 2 | Sync active specs and static/manual evidence | PR 2 | Depends on PR 1; verify TD-001..TD-015. |

## Phase 1: Remove Stale Contract Authority

- [x] 1.1 Delete `skills/_shared/sdd-operational-readiness-contract.md`; evidence: TD-001 file absence.
- [x] 1.2 Search active `skills/**` for `sdd-operational-readiness-contract.md`, `## Operational Readiness`, `readiness completeness`, and mandatory category language; record active-only results for TD-002 and TD-013.
- [x] 1.3 Update `skills/_shared/sdd-security-contract.md` and `skills/_shared/sdd-post-apply-gates.md` to preserve safe-evidence boundaries while consuming only actual operational evidence; covers TD-008, TD-009, TD-015, SEC-OP-001..006.

## Phase 2: Update Phase Skills

- [x] 2.1 Update `skills/sdd-design/SKILL.md` to make `## Operational Considerations` conditional/design-owned, keep `## Secure Development Design`, and remove mandatory completeness gates; covers TD-003.
- [x] 2.2 Update `skills/sdd-test-design/SKILL.md` to derive operational checks only from design evidence or `No aplica.` decisions, with unavailable-tooling reporting; covers TD-004, TD-014.
- [x] 2.3 Update `skills/sdd-tasks/SKILL.md` to generate operational tasks only from design/test-design evidence, never the deleted contract; covers TD-005.
- [x] 2.4 Update `skills/sdd-review/SKILL.md` to review present/planned evidence without altering the 96-control matrix or inventing data; covers TD-006.
- [x] 2.5 Update `skills/sdd-review-security/SKILL.md` to validate leakage only for existing evidence and accept exact safe placeholders without hiding security proof; covers TD-007, TD-009, TD-015.
- [x] 2.6 Update `skills/sdd-verify/SKILL.md` and `skills/sdd-archive/SKILL.md` to verify/archive actual evidence, gaps, warnings, and unavailable tooling without requiring readiness completeness or `sdd-operational-doc`; covers TD-008, TD-014.
- [x] 2.7 Update `skills/sdd-operational-doc/SKILL.md` to remain manual post-archive, consume archived evidence, preserve sections 1-9/R1-R4, and emit `No aplica.` or `Pendiente de confirmar:` without invention; covers TD-012.

## Phase 3: Sync Active OpenSpec Specs

- [x] 3.1 Update active specs: `sdd-design-workflow`, `sdd-test-design-workflow`, `sdd-review-workflow`, `sdd-review-security-workflow`, `sdd-execution-persistence-contracts`, `sdd-operational-readiness-workflow`, and `sdd-security-guideline-catalog`; covers TD-010, TD-011.
- [x] 3.2 Confirm monitoring wording in affected skills/specs is mechanism-oriented and not SQL-only; covers TD-010.

## Phase 4: Static and Manual Evidence

- [x] 4.1 Record static/manual evidence for TD-001..TD-015 in apply/verify artifacts, including restricted-data absence and final-document-only boundary searches.
- [x] 4.2 Cite `openspec/config.yaml#testing` to report unavailable runtime/build/lint/type/format/coverage tooling as unavailable, not passing.
- [x] 4.3 State archived `openspec/changes/archive/**` references are historical only and do not require active compatibility support.
