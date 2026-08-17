# Tasks: Selective Secure Design Controls

## Review Workload Forecast

| Field | Value |
| --- | --- |
| Estimated changed lines | 500-750 corrective delta |
| Review budget lines | 400 |
| Review budget risk | High |
| 400-line budget risk | High |
| Chained PRs recommended | Yes |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |
| Current slice boundary | Corrective no-YAML delta only |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
| --- | --- | --- | --- |
| 1-4 | Prior selective boundary | Completed | Shared contract, design/test-design, review-security, source-spec sync were applied for the earlier selective-control boundary. |
| 5 | Narrative shared/design contract | PR 5 to main | Remove design YAML/schema/matrix duties. |
| 6 | Narrative test-design consumption | PR 6 to main | Depends on WU5. |
| 7 | Review-security machine-readable ownership | PR 7 to main | Depends on WU5-WU6. |
| 8 | Source spec sync and reports | PR 8 to main | Depends on WU5-WU7. |

## Phase 1: Prior Completed Baseline

- [x] 1.1 Preserve earlier applied WU1-WU4 context in `apply-progress.md`: selective boundary, design/test-design planning, review-security authority, and initial source-spec sync.
- [x] 1.2 Treat prior `review-report.md` as stale for the old implementation shape; regenerate review evidence after corrective apply.

## Phase 2: Narrative Design Contract

- [x] 2.1 Update `skills/_shared/sdd-security-contract.md` so design owns narrative category rules only, while review-security owns schema, compact matrices, Source ID matrices, and exhaustive `N/A` decisions. Traces: TD-004, TD-016.
- [x] 2.2 Update `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` to prohibit YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, and all-row `N/A` bookkeeping. Traces: TD-001-TD-003.

## Phase 3: Narrative Test-Design Consumption

- [x] 3.1 Update `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` to consume changed-surface rationale and applicable narrative rules only; reject design schema/matrix dependencies. Traces: TD-006, TD-007.
- [x] 3.2 Static/manual check changed design/test-design contracts for obsolete design-owned YAML/schema/matrix requirements; classify remaining matches as historical or review-security-only. Traces: TD-015, TD-016.

## Phase 4: Review-Security Ownership

- [x] 4.1 Update `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` to parse narrative design and own report schema, compact matrix, exact-once 155 Source ID expansion, missed-category blockers, unsafe-evidence blockers, and supported `N/A` validation. Traces: TD-008-TD-010, TD-017.
- [x] 4.2 Preserve safe evidence, denial-by-default routing, warning visibility, and historical compatibility behavior across security review outputs. Traces: TD-011-TD-014, TD-018.

## Phase 5: Source Specs and Evidence

- [x] 5.1 Sync revised deltas into `openspec/specs/sdd-design-workflow/spec.md`, `sdd-test-design-workflow/spec.md`, `sdd-review-security-workflow/spec.md`, `sdd-security-guideline-catalog/spec.md`, and `sdd-execution-persistence-contracts/spec.md`.
- [x] 5.2 Update `apply-progress.md` with corrective WU evidence, unavailable tooling notes from `openspec/config.yaml#testing`, read-back summaries, and safe grep results.
- [x] 5.3 Mark prior `review-report.md` stale for the previous implementation shape; regenerate general review in the next `sdd-review` phase, then run review-security/verify/archive against the new narrative-design boundary.
