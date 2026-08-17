## Verification Report

**Change**: `source-row-first-security-review`  
**Version**: OpenSpec delta change  
**Mode**: Standard (`strict_tdd: false`)  
**Artifact store**: OpenSpec  
**Verdict**: PASS WITH WARNINGS  
**Next recommendation**: archive

### Completeness

| Metric | Value |
| --- | --- |
| Tasks total | 17 |
| Tasks complete | 17 |
| Tasks incomplete | 0 |
| Required artifacts readable | proposal, 4 delta specs, design, test-design, tasks, apply-progress, review-report JSON/Markdown, review-security-report JSON/Markdown |
| General review prerequisite | PASS WITH WARNINGS; 0 blocking findings; canonical JSON authoritative |
| Security review prerequisite | PASS WITH WARNINGS; 0 blockers; canonical JSON authoritative |

### Review Evidence Consumption

| Report | Canonical ref | Derived compatibility ref | Verdict / route | Verification finding |
| --- | --- | --- | --- | --- |
| General review | `openspec/changes/source-row-first-security-review/review-report.json` | `openspec/changes/source-row-first-security-review/review-report.md` | PASS WITH WARNINGS / `review-security` | Non-blocking. Canonical JSON reports `blockingFailureCount: 0`; warnings are limited to unavailable global tooling and downstream security-review ownership. Matrix not reproduced here. |
| Security review | `openspec/changes/source-row-first-security-review/review-security-report.json` | `openspec/changes/source-row-first-security-review/review-security-report.md` | PASS WITH WARNINGS / `verify` | Non-blocking. Canonical JSON reports `status: success`, `verdict: PASS WITH WARNINGS`, no blockers, no unsafe evidence rejections, and parity/read-back passed. Full source-row matrix not reproduced here. |

### Source-Row Security Review Consumption

| Check | Expected | Observed | Result |
| --- | ---: | ---: | --- |
| `sourceRowValidation.rows` count | 155 | 155 | PASS |
| `sourceRowValidation.expectedCount` | 155 | 155 | PASS |
| `sourceRowValidation.validatedCount` | 155 | 155 | PASS |
| Unique Source IDs | 155 | 155 | PASS |
| Duplicate Source IDs | 0 | 0 | PASS |
| Rows missing required field set | 0 | 0 | PASS |
| Coverage status | complete | complete | PASS |
| Exact once | true | true | PASS |
| Row routes | verify | verify | PASS |
| Blockers | 0 | 0 | PASS |
| Unsafe evidence rejections | 0 | 0 | PASS |
| Exceptions | 0 | 0 | PASS |
| JSON/Markdown parity | passed | passed | PASS |

Catalog evidence: `skills/sdd-review-security/references/security-guideline-catalog.operational.json`, snapshot `security-guidelines-initial-user-snapshot-2026-06-30`, expected source row count `155`. Grouping fields are `corporateSection` and `controlDomain`; Markdown navigation is generated from source-row fields only.

Security-review warning preserved: repository-wide build, lint, type-check, format, and coverage tooling is unavailable; scoped Python exporter pytest passed and unavailable tooling is carried forward without being treated as passing evidence.

### Build & Tests Execution

| Command / capability | Availability | Result | Evidence |
| --- | --- | --- | --- |
| `python -m pytest python/tests` | Available | PASS | 20 tests passed in 0.99s on Python 3.12.6 / pytest 9.1.1. |
| `git diff --check` | Available | PASS WITH WARNINGS | No whitespace errors reported; Git emitted CRLF-to-LF normalization warnings for touched files. |
| Global build command | Unavailable | WARNING | `openspec/config.yaml#testing` has no build command; not treated as passing evidence. |
| Global lint command | Unavailable | WARNING | No repository-wide linter detected; not treated as passing evidence. |
| Global type-check command | Unavailable | WARNING | No repository-wide type checker detected; not treated as passing evidence. |
| Global format command | Unavailable | WARNING | No repository-wide formatter detected; not treated as passing evidence. |
| Global coverage command | Unavailable | WARNING | No coverage runner detected; not treated as passing evidence. |

### Exporter Verification

| Requirement | Evidence | Result |
| --- | --- | --- |
| Security-review default table path is `sourceRowValidation.rows` | `python/json_report_to_excel.py#DEFAULT_TABLE_BY_SCHEMA`; pytest `test_security_schema_exports_nested_source_row_rows_by_default` | PASS |
| Compact-only security reports fail by default | pytest `test_compact_only_security_report_fails_without_table_override` | PASS |
| Generic nested `--table` override remains supported | pytest `test_nested_table_override_and_excel_safe_sheet_name` | PASS |
| Workbook read-back validates sheet name, headers, flattening, formatting, and temp output behavior | pytest workbook read-back tests with `openpyxl` | PASS |
| README documents source-row default and not legacy compact active default | `python/README.md`; pytest `test_readme_documents_required_usage_and_policies` | PASS |
| Generated `.xlsx` artifacts committed | Glob `**/*.xlsx` | PASS: no files found |

### Spec Compliance Matrix

| Capability / requirement | Scenarios verified | Covering evidence | Result |
| --- | --- | --- | --- |
| `sdd-review-security-workflow`: Security Review Artifact | Report is source-row authoritative; compact controls are absent | Canonical `review-security-report.json` has schema v2, `sourceRowValidation.rows[155]`, no blockers, parity passed; active changed contracts grep found no compact authority except legacy-negative tests/context. | COMPLIANT |
| `sdd-review-security-workflow`: Security Matrix Validation | Exact-once row validation; grouped N/A preserves rows | Static JSON check: 155 rows, 155 unique Source IDs, required row fields present, `coverageStatus: complete`, `exactOnce: true`; Markdown grouped summaries state row-level preservation. | COMPLIANT |
| `sdd-review-security-workflow`: Exhaustive Source Row Security Review | Lean generated Markdown; JSON wins | `review-security-report.md` cites JSON authority, source-row navigation, source-row summary, grouped non-applicability, and full matrix last; artifact metadata reports parity passed. | COMPLIANT |
| `sdd-security-guideline-catalog`: In-Repo Guideline Snapshot | Catalog snapshot available; source-row vocabulary exists | Catalog refs in canonical security review; apply/review evidence reports 155 expanded unique source rows, grouping metadata, owner phase, route metadata, and evidence expectations. | COMPLIANT |
| `sdd-security-guideline-catalog`: Corporate Source Row Inventory | Ranges expand before coverage; grouping fields present | Source-row validation summary and navigation by `corporateSection`/`controlDomain`; exact-once check passes. | COMPLIANT |
| `sdd-security-guideline-catalog`: Safe Source Row Evidence | N/A row justified; grouping cannot hide gaps | Canonical rows preserve per-row evidence fields; grouped N/A summaries are derived from row-level JSON; unsafe evidence rejections are empty. | COMPLIANT |
| `sdd-execution-persistence-contracts`: Verify and Archive Review Consumption | Verify consumes review evidence; compact report data ignored | Verify consumed canonical JSON summaries only and did not reproduce/re-score review matrices; active changed downstream contracts route on source-row summaries. | COMPLIANT |
| `sdd-execution-persistence-contracts`: Source Row Persistence Compatibility | OpenSpec preserves rows; generated Markdown from JSON | OpenSpec canonical JSON and derived Markdown are persisted/readable; JSON/Markdown parity passed. | COMPLIANT |
| `sdd-execution-persistence-contracts`: Verify Source Row Consumption | Security source blocker remains; complete source rows continue | No source-row blockers; counts complete; warning carry-forward preserved; route eligible for archive. | COMPLIANT |
| `sdd-execution-persistence-contracts`: Archive Source Row Preservation | Archive checks row completeness; archive preserves generated matrix refs | Verify confirms archive can consume canonical refs/counts/warnings without creating a second validation source. | COMPLIANT |
| `canonical-review-json-excel-exporter`: Schema-aware table selection | Known review schema; known security schema; legacy compact not default; manual nested path | Python tests passed and code maps security review schema to `sourceRowValidation.rows`. | COMPLIANT |
| `canonical-review-json-excel-exporter`: User documentation | README explains security export | README inspected by pytest and static read-back; source-row default documented; legacy compact path not documented as active. | COMPLIANT |
| `canonical-review-json-excel-exporter`: Pytest verification | Test suite validates behavior; security default regression covered | `python -m pytest python/tests` passed with 20 tests. | COMPLIANT |

### Security Evidence Matrix

| Narrative rule / control area | Expected evidence | Observed evidence | Result |
| --- | --- | --- | --- |
| Sensitive data and evidence handling | Safe paths, section anchors, summaries, command outcomes; no secrets/PII/PAN/tokens/log payloads copied | Design safe-evidence policy; review-security unsafe evidence rejections empty; exporter safe CLI error tests passed. | COMPLIANT |
| Files and exported artifacts | Source-row default export, temp workbook read-back, no committed workbook bytes | Pytest uses temporary paths; glob found no `.xlsx` artifacts; README warns against committing generated workbooks. | COMPLIANT |
| Artifact authority / access-control boundary | Canonical JSON wins; derived Markdown/exporter output are presentation only | Security-review artifact metadata reports JSON canonical, Markdown derived, parity passed; verify consumes summaries only. | COMPLIANT |
| Sensitive logging and error evidence | Errors mention paths/counts safely and do not echo unsafe payloads | Exporter tests assert no traceback and no unsafe payload echoes in CLI error output. | COMPLIANT |
| Schema and input validation | Exact-once 155 rows, required fields, allowed routes, safe N/A, no compact authority | Static JSON check confirms counts/required fields/routes; schema/template/rules inspected; security review has no blockers. | COMPLIANT |
| Exception and evidence policy | No incomplete exceptions; warnings preserved | Canonical security review has `exceptions: []`; warning carry-forward preserved for unavailable global tooling. | COMPLIANT |

### Test-Design Coverage Matrix

| Case range | Severity | Expected evidence | Observed evidence | Result |
| --- | --- | --- | --- | --- |
| TD-001..TD-006 | mandatory | Schema/source-row validation, exact-once rows, missing/duplicate/unknown/required-field handling | Schema and validation rules inspected by apply/review; verify static JSON check confirmed 155 rows, uniqueness, required fields, complete coverage. | COMPLIANT |
| TD-007 | mandatory | No active compact-control authority in changed contracts | Grep checks over active changed skill/shared/agent/python surfaces found no active compact authority; remaining matches are forbidden legacy tests/context, archived artifacts, or accepted specs pending archive sync. | COMPLIANT |
| TD-008..TD-010 / NA-001..NA-005 | mandatory | Row-level N/A evidence and source-row grouping only | Canonical JSON preserves row-level N/A evidence; Markdown groups by source-row fields and states groups do not override rows. | COMPLIANT |
| TD-011..TD-012 / MD-001..MD-005 | mandatory | Lean Markdown, full matrix last, JSON wins, parity checks | Derived Markdown has source-row summary/navigation, no compact-control validation section, full matrix last, parity passed. | COMPLIANT |
| TD-013..TD-014 | mandatory | Verify/archive consume canonical summaries and do not copy/re-score matrices | `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md`, shared contracts, and agent wrappers updated; verify report consumes summaries only. | COMPLIANT |
| TD-015 | mandatory | Catalog snapshot has 155 source rows and grouping metadata | Security review catalog refs and source-row validation counts confirm expected snapshot and 155-row inventory. | COMPLIANT |
| TD-016..TD-017 / PY-001..PY-002 | mandatory | Exporter defaults to source rows; compact-only default fails | Pytest passed. | COMPLIANT |
| TD-018 / PY-003 | non-mandatory | Generic nested override remains supported | Pytest passed. | COMPLIANT |
| TD-019 / PY-004 | mandatory | Workbook read-back through `openpyxl` | Pytest passed. | COMPLIANT |
| TD-020 / PY-005 | mandatory | README documents active source-row default and safe use | Pytest and static README inspection passed. | COMPLIANT |
| TD-021 | mandatory | Operational evidence uses safe refs and no restricted data | Review/security-review evidence uses paths, summaries, command outcomes, and unavailable-tooling notes only. | COMPLIANT |
| TD-022 | mandatory | Unavailable global tooling carried forward explicitly | Unavailable global build/lint/type/format/coverage preserved as warnings, not passing checks. | COMPLIANT WITH WARNING |

### Correctness (Static Evidence)

| Requirement | Status | Notes |
| --- | --- | --- |
| All tasks complete | PASS | `tasks.md` has all checkboxes complete and `apply-progress.md` records no remaining apply tasks. |
| Canonical security JSON source-row-first | PASS | JSON has schema v2, 155 rows, exact-once complete coverage, no blockers. |
| Exporter default/tests/docs align with `sourceRowValidation.rows` | PASS | Code, tests, and README align; scoped pytest passed. |
| Active compact-control authority removed from changed active contracts | PASS | No active changed-contract authority remains; legacy mentions are confined to negative tests, proposal/design/spec context, archived artifacts, and accepted specs awaiting archive sync. |
| Generated workbooks absent | PASS | `**/*.xlsx` returned no files. |

### Coherence (Design)

| Decision | Followed? | Notes |
| --- | --- | --- |
| Source rows are the only active security matrix | Yes | Canonical security review uses `sourceRowValidation.rows`; verify consumes summary/count evidence only. |
| Compact `SEC-*` removed from active report authority | Yes | No active compact-control validation/navigation/summary is used in changed active contracts. |
| Grouping uses source-row category fields | Yes | Security review navigation/grouped N/A use `corporateSection` and `controlDomain`. |
| Markdown is generated presentation only | Yes | Markdown states JSON authority; parity/read-back passed. |
| Excel exporter defaults to source rows | Yes | Code/tests/README verify `sourceRowValidation.rows`. |

### Operational Evidence / Gaps / Warnings

| Category | Status | Evidence |
| --- | --- | --- |
| Artifact authority | PASS | Canonical JSON refs are authoritative before derived Markdown; artifact metadata read-back/parity passed. |
| Exporter behavior | PASS | Scoped Python tests passed; generated workbook outputs are temporary and none are committed. |
| Contract propagation | PASS | Changed skills/shared contracts/agents/spec deltas align with source-row-first consumption. |
| Restricted operational data | PASS | Evidence uses safe paths/summaries/commands only; no production identifiers, credentials, raw logs, payloads, or workbook bytes recorded. |
| Global tooling | WARNING | Global build/lint/type/format/coverage unavailable and preserved as unavailable evidence. |

### Skipped / Degraded Dimensions

| Dimension | Status | Reason |
| --- | --- | --- |
| Strict TDD protocol | Skipped | `strict_tdd: false`; no Strict TDD instruction active. |
| Repository-wide build | Unavailable | No build command configured in `openspec/config.yaml`. |
| Repository-wide lint/type/format/coverage | Unavailable | No commands configured; unavailable tooling is warning evidence only. |
| Full review/security matrices | Intentionally not reproduced | Matrix ownership belongs to `sdd-review` and `sdd-review-security`; verify consumes summaries/counts/refs only. |
| Accepted spec sync | Deferred | Base specs under `openspec/specs/*` are updated during `sdd-archive` per apply-progress archive sync plan. |

### Issues Found

**CRITICAL**: None.

**WARNING**:
- Repository-wide build, lint, type-check, format, and coverage tooling are unavailable; do not treat them as passing evidence.
- `git diff --check` emitted CRLF-to-LF normalization warnings for touched files, but no whitespace errors.
- Active accepted specs still contain the previous exporter compact default until archive sync applies the delta specs; this is expected pre-archive and not active changed-contract authority.

**SUGGESTION**: None.

### Verdict

PASS WITH WARNINGS.

All implementation tasks are complete, scoped Python exporter tests pass, canonical general review and security review are non-blocking, source-row validation is complete and exact-once with 155 rows, no generated `.xlsx` artifacts are present, and warnings are non-blocking carry-forward items. Proceed to `sdd-archive`.
