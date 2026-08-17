# Tasks: Slim Secure Design Artifacts

## Review Workload Forecast

| Field | Value |
|---|---|
| estimated_changed_lines | 900-1,300 |
| review_budget_lines | 400 |
| review_budget_risk | High |
| chained_prs_recommended | Yes |
| decision_needed_before_apply | No |
| rationale | Cross-cuts shared contracts, phase skills, adapter prompts, source spec sync, and static evidence; split by phase ownership. |
| work_unit_boundaries | WU1 shared catalog/contract; WU2 design/test-design; WU3 review-security; WU4 verify/archive/spec sync/evidence. |
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
|---|---|---|---|
| 1 | Establish catalog/shared source-row boundary. | PR 1 | Base `main`; verifies TD-003, TD-005, TD-011, SEC-ACCESS-001. |
| 2 | Make design/test-design slim by reference. | PR 2 | Stacks on PR 1; verifies TD-001, TD-002, TD-004, TD-006-TD-013. |
| 3 | Keep exhaustive expansion only in review-security. | PR 3 | Stacks on PR 2; verifies TD-014, TD-015. |
| 4 | Preserve verify/archive evidence and source specs. | PR 4 | Stacks on PR 3; verifies TD-016-TD-020 and archive readiness. |

## Phase 1: Shared Contract and Catalog

- [x] 1.1 Update `skills/_shared/sdd-security-contract.md` to state catalog/design/test-design/review-security/verify/archive ownership, blocker routing, safe evidence, N/A, warning, and exception rules.
- [x] 1.2 Update `skills/_shared/security-guideline-catalog.md` to declare the authoritative 155 Source ID inventory, snapshot metadata, expanded ranges, compact mappings, and expected count.
- [x] 1.3 Record WU1 static evidence for TD-003, TD-005, TD-011, and SEC-AUTH/SESS/DATA/SECRET/ACCESS/FILE/DB/LOG coverage.

## Phase 2: Design and Test-Design Contracts

- [x] 2.1 Update `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` to require slim `design.md#secure-development-design` references and prohibit 155-row or 96-control duplication.
- [x] 2.2 Update `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` to plan grouped static/manual checks from catalog references and report unavailable runtime tooling.
- [x] 2.3 Record WU2 evidence covering TD-001, TD-002, TD-004, TD-006-TD-013, safe evidence, no legacy standalone dependency, and no runtime/tooling changes.

## Phase 3: Review-Security Ownership

- [x] 3.1 Update `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` so only `review-security-report.md` materializes all 155 Source IDs exactly once.
- [x] 3.2 Preserve general review citation-only behavior: security review may cite `review-report.md` rows but must not duplicate the 96-control matrix.
- [x] 3.3 Record WU3 evidence for TD-014, TD-015, duplicate/missing/unknown Source ID blockers, and unsafe evidence blockers.

## Phase 4: Verify, Archive, Source Specs, Evidence

- [x] 4.1 Update `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` to consume non-blocking review-security evidence without owning the full matrix.
- [x] 4.2 Update `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` to preserve source-row summaries, catalog identity, count, mappings, warnings, exceptions, and links.
- [x] 4.3 Sync accepted requirements into `openspec/specs/sdd-design-workflow/spec.md`, `sdd-test-design-workflow/spec.md`, `sdd-review-security-workflow/spec.md`, `sdd-security-guideline-catalog/spec.md`, and `sdd-execution-persistence-contracts/spec.md` when archive applies the delta.
- [x] 4.4 Record WU4 evidence for TD-016-TD-020, unavailable tests/lint/type/format/coverage, changed-file review, and archive-safe references.
