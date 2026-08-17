# Tasks: Canonical Review JSON Excel Exporter

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 260-360 |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | null |
| Chain strategy | pending |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: pending
Size exception: none
Review budget lines: 400
Review budget risk: Medium
400-line budget risk: Medium

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Add exporter, docs, and pytest coverage | PR 1 | Single slice; monitor changed lines. |

## Phase 1: Python Utility Foundation

- [x] 1.1 Create `python/` with `python/tests/`, `python/json_report_to_excel.py`, `python/README.md`, and `python/requirements.txt`.
- [x] 1.2 Add `python/requirements.txt` with only `openpyxl>=3.1,<4` and `pytest`; exclude pandas/Excel automation/platform Excel deps (TD-001, SEC-DEPS-MINIMAL).
- [x] 1.3 Document in `python/README.md` virtualenv, install, `review-report.json`, `--table`, overwrite/derived-output policy, dependency policy, and `python -m pytest python/tests` (TD-017).
- [x] 1.4 In README examples, prefer active `openspec/changes/sdd-review-json-contract/review-report.json` when available; archived/sample paths are fallback inputs only.

## Phase 2: Exporter Helpers and CLI

- [x] 2.1 Implement `python/json_report_to_excel.py` helpers for JSON load/top-level-object validation, output path, scalar metadata, and argparse dispatch (TD-002, TD-016).
- [x] 2.2 Add `DEFAULT_TABLE_BY_SCHEMA`: `sdd-review.review-report -> reviewMatrix`, `sdd-review-security.review-security-report -> compactControlValidation.rows`; let `--table` override with a top-level table or nested dotted path and unknown schemas ask for `--table` (TD-003-TD-006).
- [x] 2.3 Validate selected table/path as a list of objects; fail before save for missing/invalid table shapes while mentioning the requested table/path (TD-007, TD-016, SEC-FILES-FAIL-CLOSED).
- [x] 2.4 Flatten scalars directly, scalar lists with `; `, and nested/complex values as compact deterministic JSON without eval/formula execution (TD-010, SEC-JSON-STDLIB-NO-EVAL).
- [x] 2.5 Create workbook with `summary` and table sheets, scalar metadata, union headers, bold headers, autofilter, `A2` freeze panes, widths, and wrapped long text (TD-008-TD-012).
- [x] 2.6 Return CLI exit `0` on success and `1` on user-correctable failures with concise stderr and no payload/workbook/stack trace/secrets/token/PII/full-row dumps (TD-015, SEC-LOGGING-SAFE-ERRORS).
- [x] 2.7 Read only the explicit JSON path and write only the resolved output path; no directory scans, Markdown parsing, remote fetch, unrelated deletes, or Excel requirement (TD-014, TD-018, SEC-FILES-LOCAL-IO).

## Phase 3: Tests and Verification Evidence

- [x] 3.1 Create `python/tests/test_json_report_to_excel.py` with `tmp_path` JSON fixtures and `openpyxl.load_workbook`; no repo generated files as inputs (TD-011).
- [x] 3.2 Cover dependency constraints, CLI/default output, schema defaults, `--table`, fail-closed cases, workbook content/formatting, no Markdown parsing, safe errors, and unrelated-file preservation where practical (TD-001-TD-018).
- [x] 3.3 Run `python -m pytest python/tests`; record sanitized command/result/path/sheet-structure evidence and report coverage/lint/typecheck/format as unavailable, not passing.

## Phase 4: Apply/Review Evidence Boundaries

- [x] 4.1 During apply/review, cite only safe evidence: files, sections, commands, sanitized summaries, and sheet metadata; no raw payloads/workbook bytes/secrets/tokens/PII/PAN/full exports (SEC-EVIDENCE-SAFE).
- [x] 4.2 Confirm this remains a manual local utility with no service, scheduler, database, monitoring, auth, deployment, or automated SDD phase surface.

## Apply Evidence

- Command: `python -m py_compile python/json_report_to_excel.py python/tests/test_json_report_to_excel.py` — Result: passed (syntax-only check).
- Command: `python -m pytest python/tests` — Result: blocked in current environment: `No module named pytest`; no test pass is claimed.
- Dependency availability check: `openpyxl` is also unavailable in the current global interpreter (`ModuleNotFoundError: No module named 'openpyxl'`).
- Created paths: `python/json_report_to_excel.py`, `python/requirements.txt`, `python/README.md`, `python/tests/test_json_report_to_excel.py`.
- Planned workbook metadata covered by tests: `summary` sheet plus selected table/path sheet (`reviewMatrix`, `compactControlValidation.rows`, or explicit `--table`), bold headers, autofilter, `A2` freeze panes, widths, and wrapped long text.
- Safe-evidence boundary: evidence records commands, paths, and sheet metadata only; no raw workbook bytes, full exports, secrets, tokens, PII, PAN, or raw payload dumps are included.
- Unavailable tooling: coverage, lint, typecheck, and format commands are unavailable for this change and are not reported as passing.

## Post-Apply Unblock Evidence

- Command: `python -m pytest python/tests` — Result: passed; 15 tests collected, 15 passed.
- Acceptance sample: active `openspec/changes/sdd-review-json-contract/review-report.json` was unavailable; archived sample `openspec/changes/archive/2026-07-10-sdd-review-json-contract/review-report.json` generated a valid workbook during unblock verification.
- Generated acceptance workbook was treated as derived local output evidence and is not required as a persisted SDD artifact.

## Apply Remediation Evidence

- Superseded remediation note: an earlier remediation updated the automatic security-review export schema name from historical `sdd.review-security-report` to current `sdd-review-security.review-security-report` without retaining a historical alias, but still targeted the wrong top-level table for the current JSON shape. The follow-up blocker remediation below is now authoritative for table/path selection.
- Command: `python -m pytest python/tests` — Result: passed; 15 tests collected, 15 passed.
- Stale evidence note: existing general review and security-review reports predate this remediation and were intentionally not edited; route back to `review` after apply remediation.

## Apply Blocker Remediation Evidence

- Remediation: updated table references to support either top-level table names or nested dotted paths, and changed the current security-review schema default to `compactControlValidation.rows`.
- Sheet naming decision: nested table paths use the sanitized full path as the worksheet name, e.g. `compactControlValidation.rows`, rather than only the last segment, because the full path is Excel-safe, under 31 characters, and more understandable for security-review exports.
- Failure behavior: missing nested paths fail before workbook save with an error mentioning the requested table/path; nested values that are not a list of objects fail before workbook save.
- Documentation/planning updates: proposal, spec, design, test-design, tasks, README, and pytest coverage now describe the nested-path behavior.
- Command: `python -m pytest python/tests` — Result: passed; 18 tests collected, 18 passed.
- Stale evidence note: existing general review and security-review reports predate this blocker remediation and were intentionally not edited; route back to `review` after apply remediation.
