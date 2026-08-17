# Review Security Report: canonical-review-json-excel-exporter

## Verdict

| Field | Value |
| --- | --- |
| Change | `canonical-review-json-excel-exporter` |
| Status | success |
| Verdict | PASS WITH WARNINGS |
| Blocking findings | 0 |
| Non-blocking warnings | 2 |
| Next recommendation | verify |
| JSON authority | `openspec/changes/canonical-review-json-excel-exporter/review-security-report.json` |
| Markdown authority | derived compatibility view |

## Source References

- Secure design: `openspec/changes/canonical-review-json-excel-exporter/design.md#secure-development-design`
- Test design: `openspec/changes/canonical-review-json-excel-exporter/test-design.md`
- Tasks/apply evidence: `openspec/changes/canonical-review-json-excel-exporter/tasks.md`, `openspec/changes/canonical-review-json-excel-exporter/tasks.md#apply-blocker-remediation-evidence`
- Changed-file context: `python/requirements.txt`, `python/README.md`, `python/json_report_to_excel.py`, `python/tests/test_json_report_to_excel.py`
- General review JSON: `openspec/changes/canonical-review-json-excel-exporter/review-report.json`
- General review Markdown compatibility: `openspec/changes/canonical-review-json-excel-exporter/review-report.md`
- Catalog JSON: `skills/sdd-review-security/references/security-guideline-catalog.operational.json` snapshot `security-guidelines-initial-user-snapshot-2026-06-30`
- Catalog human view: `skills/sdd-review-security/references/security-guideline-catalog.md`

## General Review Handoff

General review rerun is non-blocking: PASS WITH WARNINGS, 0 blocking failures, 2 non-blocking findings, nextRecommended `review-security`. Canonical `review-report.json` is authoritative; Markdown is compatibility only. Pytest passed after nested table path remediation with 18 collected / 18 passed. Coverage, lint, typecheck, formatter, project build, and dedicated secret scanner remain unavailable and are not claimed passing.

Security-relevant handoff refs:

- `openspec/changes/canonical-review-json-excel-exporter/review-report.json#changedFileSecurityHandoff`
- `openspec/changes/canonical-review-json-excel-exporter/review-report.json#operationalEvidenceSummary`
- `openspec/changes/canonical-review-json-excel-exporter/review-report.json#runtimeChecks`

## Compact Control Validation

| Guideline ID | Category | Design Applies | Lifecycle Status | Complies | Evidence Location | Observations | Finding |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | N/A | `not-applicable` | N/A | `proposal.md#scope; design.md#secure-development-design; python/json_report_to_excel.py` | No authentication surface is introduced. | none |
| `SEC-SESS-001` | `sessions` | N/A | `not-applicable` | N/A | `proposal.md#scope; design.md#secure-development-design; python/json_report_to_excel.py` | No session/token surface is introduced. | none |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | `implemented` | Yes | `design.md#restricted-data-boundary; python/README.md; python/tests/test_json_report_to_excel.py#test_cli_exit_codes_and_safe_errors` | Workbooks inherit source JSON confidentiality; SDD evidence remains path/command/structure-only. | warning |
| `SEC-SECRET-001` | `secrets` | Yes | `implemented` | Yes | `python/requirements.txt; python/json_report_to_excel.py; python/tests/test_json_report_to_excel.py#test_cli_exit_codes_and_safe_errors` | No secret storage or credential constants are introduced; sentinel sensitive-looking fixture values are not emitted to stderr. | none |
| `SEC-ACCESS-001` | `permissions-access-control` | N/A | `not-applicable` | N/A | `proposal.md#scope; design.md#secure-development-design; python/json_report_to_excel.py` | No access-control surface is introduced. | none |
| `SEC-FILE-001` | `files` | Yes | `implemented` | Yes | `python/json_report_to_excel.py#DEFAULT_TABLE_BY_SCHEMA; python/json_report_to_excel.py#get_value_at_path; python/tests/test_json_report_to_excel.py#test_security_schema_exports_nested_compact_control_rows; python/README.md#table-selection` | Nested path remediation supports the current security-review JSON path `compactControlValidation.rows`; manual `--table` supports top-level names and nested dotted paths; missing/invalid paths fail before workbook save. | none |
| `SEC-DB-001` | `database-access` | N/A | `not-applicable` | N/A | `proposal.md#scope; design.md#secure-development-design; python/json_report_to_excel.py` | No database surface is introduced. | none |
| `SEC-LOG-001` | `sensitive-logging` | Yes | `implemented` | Yes | `design.md#sensitive-logging-rules; python/json_report_to_excel.py#main; python/tests/test_json_report_to_excel.py#test_cli_exit_codes_and_safe_errors; python/tests/test_json_report_to_excel.py#test_subprocess_has_no_traceback_on_user_error` | Concise stderr errors; no raw payload sentinel values or tracebacks are emitted for user-correctable failures. | none |

## Corporate Source Row Validation

Expected Source ID count: `155`. Validated Source ID count: `155`. Coverage: `complete`. Security review validated every expected Source ID exactly once. Summary mode is used; the full 155-row matrix is intentionally omitted.

### Coverage Summary

| Corporate Section | Expected | Validated | Blockers | Warnings | N/A | Notes |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| 1. Authentication | 10 | 10 | 0 | 0 | 10 | No authentication or protected web-resource surface is introduced. |
| 2. Passwords | 23 | 23 | 0 | 0 | 23 | No password or credential-store behavior is introduced. |
| 3. Access and Activity Logging | 11 | 11 | 0 | 0 | 9 | Sensitive-output expectations apply to CLI stderr only. |
| 4. Cryptography | 8 | 8 | 0 | 0 | 8 | No cryptographic storage, transmission, certificate, random generation, or key-management surface is introduced. |
| 5. Databases | 12 | 12 | 0 | 0 | 12 | No database or connection-string surface is introduced. |
| 6. Coding | 14 | 14 | 0 | 0 | 8 | Fail-closed validation, stdlib JSON parsing, no eval, no remote fetch, no exposed config, minimal dependencies, and review evidence are satisfied. |
| 7. Session Management | 13 | 13 | 0 | 0 | 13 | No session surface is introduced. |
| 8. Error Handling | 5 | 5 | 0 | 0 | 2 | Deterministic ExportError handling and no traceback/payload dumping. |
| 9. File Handling | 12 | 12 | 0 | 1 | 5 | Explicit local JSON input/output and current nested security-review table export are satisfied; generated workbook confidentiality remains a warning. |
| 10. Memory Management | 6 | 6 | 0 | 0 | 6 | No native buffer or explicit sensitive-memory management surface is introduced. |
| 11. Input Validation | 16 | 16 | 0 | 0 | 16 | No database surface; JSON/table shape validation is covered under file/coding evidence. |
| 12. Output Encoding | 5 | 5 | 0 | 0 | 5 | No browser, XML, LDAP, SQL, or command-output encoding surface is introduced; workbook text is generated through openpyxl. |
| 13. Data Protection | 9 | 9 | 0 | 1 | 8 | Generated workbooks inherit source JSON confidentiality. |
| 14. Access Control | 9 | 9 | 0 | 0 | 9 | No access-control surface is introduced. |
| 15. PAN — Primary Account Number | 2 | 2 | 0 | 0 | 2 | No PAN-specific storage, display, masking, or truncation surface is introduced. |

### Focused Source Row Details

| Source ID | Corporate Section | PCI Alignment | Guideline Ref | Compact Mapping | Applies | Complies | Lifecycle Status | Evidence Type | Evidence Location | Finding | Owner Phase | Route |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `6.5` | `6. Coding` | `PCI Req 6.5.8, 6.5.9` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 6.5)` | `SEC-SECRET-001; SEC-ACCESS-001; SEC-AUTH-001; SEC-DATA-001; SEC-DB-001; SEC-SESS-001; SEC-FILE-001; SEC-LOG-001` | Yes | Yes | implemented | static-inspection | `python/json_report_to_excel.py#DEFAULT_TABLE_BY_SCHEMA; python/tests/test_json_report_to_excel.py#test_security_schema_exports_nested_compact_control_rows` | none | review-security | verify |
| `9.12` | `9. File Handling` | `PCI Req 6.5.8` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 9.12)` | `SEC-FILE-001` | Yes | Yes | implemented | test-evidence | `python/tests/test_json_report_to_excel.py#test_security_schema_exports_nested_compact_control_rows; python/tests/test_json_report_to_excel.py#test_missing_nested_table_fails_before_save_and_mentions_path` | none | review-security | verify |
| `13.1` | `13. Data Protection` | `PCI Req 6.3.c, 6.5.4` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 13.1)` | `SEC-DATA-001; SEC-ACCESS-001` | Yes | Yes | implemented | implementation-reference | `design.md#secure-development-design; python/README.md` | warning | verify | verify |
| `3.10` | `3. Access and Activity Logging` | `N/A` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 3.10)` | `SEC-LOG-001` | Yes | Yes | implemented | test-evidence | `python/tests/test_json_report_to_excel.py#test_cli_exit_codes_and_safe_errors` | none | review-security | verify |
| `8.1` | `8. Error Handling` | `PCI Req 6.3.c, 6.5.5` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 8.1)` | `SEC-LOG-001` | Yes | Yes | implemented | implementation-reference | `python/json_report_to_excel.py#main` | none | review-security | verify |
| `9.7` | `9. File Handling` | `PCI Req 6.5.8` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 9.7)` | `SEC-FILE-001` | N/A | N/A | not-applicable | n/a-evidence | `design.md#secure-development-design; python/README.md` | none | review-security | verify |
| `1.1` | `1. Authentication` | `PCI Req 6.5.8, 6.5.10` | `skills/sdd-review-security/references/security-guideline-catalog.operational.json#sourceRows (Source ID 1.1)` | `SEC-AUTH-001; SEC-ACCESS-001` | N/A | N/A | not-applicable | n/a-evidence | `proposal.md#scope; design.md#secure-development-design` | none | review-security | verify |

### Full Source Row Matrix — Audit Mode Only

Omitted in summary mode; `sourceRowValidation.validatedCount` proves complete validation coverage.

## Source Row Findings

### Blockers

None.

### Warnings

- `SEC-DATA-001-W01`: generated workbooks inherit the confidentiality of user-selected canonical JSON content; do not persist workbook bytes or full exported payloads as SDD evidence.
- `GEN-REVIEW-W01`: coverage, lint, typecheck, formatter, project build, and dedicated secret scanner remain unavailable and are not claimed passing; acceptance sample limitations remain non-blocking evidence context.

### N/A Justifications

N/A decisions are justified by changed-surface evidence: this change is a manual local JSON-to-XLSX CLI with no authentication, session, database, web service, deployment, access-control, cryptographic storage/transmission, or PAN-specific surface.

### Missing Evidence Rows

None for local file I/O, nested `compactControlValidation.rows` support, fail-closed JSON/table validation, safe stderr, stdlib JSON/no eval, dependency minimization, or safe evidence.

### Unsafe Evidence Rejections

None.

### Warning Carry-Forward

- Carry forward generated-workbook confidentiality warning to verify/archive.
- Carry forward unavailable-tooling and acceptance-sample limitation warnings.

## Exceptions

None.

## Blockers and Non-Blocking Findings

- Blockers: None.
- Non-blocking warnings: generated workbook confidentiality; unavailable tooling and acceptance sample limitations.

## Unavailable Tooling

- Coverage command/tooling unavailable and not claimed passing.
- Lint command/tooling unavailable and not claimed passing.
- Typecheck command/tooling unavailable and not claimed passing.
- Formatter command/tooling unavailable and not claimed passing.
- Project build command unavailable for this standalone Python utility and not claimed passing.
- No dedicated secret scanner was available; safe-evidence judgment used design, static inspection, tests, and general-review handoff only.

## Artifact Metadata

| Check | Result |
| --- | --- |
| Canonical JSON ref | `openspec/changes/canonical-review-json-excel-exporter/review-security-report.json` |
| Derived Markdown ref | `openspec/changes/canonical-review-json-excel-exporter/review-security-report.md` |
| JSON persisted/read back | true / true |
| Markdown generated/read back | true / true |
| JSON/Markdown parity | passed |
| JSON authority | canonical |
| Markdown authority | derived |

## Recommendation

- Next recommendation: `verify`
- Follow-up: proceed to `sdd-verify` with warnings carried forward. Verification must not treat coverage/lint/typecheck/formatter/build/dedicated secret scanner as passing, and must preserve the generated-workbook confidentiality warning.
