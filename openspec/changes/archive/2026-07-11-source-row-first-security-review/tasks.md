# Tasks: Source Row First Security Review

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 900-1,400 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 source-row security contract -> PR 2 downstream SDD propagation -> PR 3 Python exporter/tests/docs |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: Resolved by launch envelope for Slice 1
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Replace active security-review schema/catalog/template/rules with source-row-first authority. | PR 1 | Base main/chain root; includes TD-001..TD-012, NA-001..NA-005, MD-001..MD-005 evidence. |
| 2 | Propagate canonical source-row contract to shared contracts, verify/archive skills, agents, and OpenSpec base-spec sync plan. | PR 2 | Depends on PR 1; includes TD-013, TD-014, TD-021, TD-022. |
| 3 | Update Python exporter default, pytest fixtures, read-back tests, and README. | PR 3 | Depends on PR 1; includes TD-016..TD-020, PY-001..PY-005. |

## Phase 1: Source-Row Security Contract

- [x] 1.1 Update `skills/sdd-review-security/SKILL.md` so canonical JSON owns `sourceRowValidation.rows[155]` and compact `SEC-*` is never active authority.
- [x] 1.2 Update `skills/sdd-review-security/references/review-security-report.schema.json` for exact 155 rows, required fields, counts, grouping fields, and no `compactControlValidation`.
- [x] 1.3 Update `skills/sdd-review-security/references/validation-rules.md` for missing/duplicate/unknown Source IDs, required fields, safe errors, unsafe evidence, and row-preserving `N/A` grouping.
- [x] 1.4 Transform `skills/sdd-review-security/references/security-guideline-catalog.operational.json` and `.md` to expanded 155 Source IDs, source-row vocabularies, evidence expectations, owner phase, and route metadata.
- [x] 1.5 Simplify `skills/sdd-review-security/references/report-template.md` to lean JSON-derived Markdown, source-row navigation, grouped `N/A`, blockers/warnings, metadata, and full matrix last.

## Phase 2: Propagation and Persistence Contracts

- [x] 2.1 Update `skills/_shared/sdd-security-contract.md`, `persistence-contract.md`, and `sdd-post-apply-gates.md` to consume source-row JSON summaries/refs only.
- [x] 2.2 Update `skills/sdd-verify/SKILL.md` and `skills/sdd-archive/SKILL.md` to cite canonical counts, warnings, exceptions, evidence refs, and never rescore/copy matrices.
- [x] 2.3 Update `agents/sdd/sdd-review-security.md`, `agents/sdd/sdd-verify.md`, and `agents/sdd/sdd-archive.md` to mirror the skill contract.
- [x] 2.4 Prepare accepted-spec sync for `openspec/specs/sdd-review-security-workflow`, `sdd-security-guideline-catalog`, `sdd-execution-persistence-contracts`, and `canonical-review-json-excel-exporter` during archive.

## Phase 3: Python Exporter

- [x] 3.1 Update `python/json_report_to_excel.py` default for `sdd-review-security.review-security-report` to `sourceRowValidation.rows`; compact-only default must fail safely.
- [x] 3.2 Update `python/tests/test_json_report_to_excel.py` for source-row default, compact-only failure, nested override, workbook read-back, sheet naming, flattening, and temp output.
- [x] 3.3 Update `python/README.md` for source-row default, setup/dependencies, `--table`, safe examples, and no compact active default.

## Phase 4: Validation and Evidence

- [x] 4.1 Record static validation for TD-001..TD-015, NA-001..NA-005, MD-001..MD-005: exact 155 rows, unknown/missing/duplicate IDs, required fields, safe `N/A`, Markdown parity, and no compact authority. (Slice 1 evidence plus Slice 2 downstream propagation evidence recorded in `apply-progress.md`.)
- [x] 4.2 Run `python -m pytest python/tests` when dependencies are available; otherwise record unavailable Python/pytest/openpyxl evidence plus static substitute inspection for PY-001..PY-005.
- [x] 4.3 Record unavailable global runner/build/lint/type/format/coverage tooling explicitly; missing tooling is not passing evidence.
- [x] 4.4 Inspect apply/review evidence for safe paths, section anchors, sanitized summaries, command outcomes, redacted placeholders, and no restricted operational data or workbook bytes.
