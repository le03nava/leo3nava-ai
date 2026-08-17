# Test Design: Canonical Review JSON Excel Exporter

## Overview

This change adds a portable Python CLI that reads canonical SDD review JSON and writes a derived `.xlsx` workbook. Test coverage must prove schema-aware table selection, explicit table override behavior, fail-closed error handling, workbook structure/content/formatting, dependency policy, and safe-evidence constraints. The repository previously had no active test runner, but this change introduces pytest under `python/tests`; verification for this change is `python -m pytest python/tests`.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/canonical-review-json-excel-exporter/proposal.md` | Defines intent, scope, non-goals, dependency policy, schema defaults, formatting expectations, and success criteria. |
| Spec | `openspec/changes/canonical-review-json-excel-exporter/specs/canonical-review-json-excel-exporter/spec.md` | Defines RFC 2119 requirements and scenarios for layout, CLI behavior, table selection, workbook structure, formatting, and pytest verification. |
| Design | `openspec/changes/canonical-review-json-excel-exporter/design.md` | Defines single-file CLI architecture, helper-function contracts, data flow, file changes, formatting behavior, rollout, and testing strategy. |
| Secure Development Design | `openspec/changes/canonical-review-json-excel-exporter/design.md#secure-development-design` | Provides changed-surface classification, applicable file-handling, sensitive-logging, JSON parsing, dependency, safe-evidence rules, residual risks, and exception policy. |
| Testing Capabilities | `openspec/config.yaml` plus launch context | Baseline repo tooling has no pre-existing runner, coverage, lint, typecheck, or format commands. This change introduces pytest and verification command `python -m pytest python/tests`; coverage/lint/type/format remain unavailable evidence. |

## Source ID Coverage Baseline

Corporate source-row validation remains owned by `sdd-review-security` and its canonical `review-security-report.json`. This test design consumes only the narrative security rules in `design.md#secure-development-design`: local file input/output, no Markdown parsing, fail-closed JSON/table handling, dependency restrictions, and safe error/evidence behavior. It does not require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, exhaustive `N/A` rows, or the full source-row matrix.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Python utility layout and dependencies | Assert `python/requirements.txt` contains `openpyxl>=3.1,<4` and `pytest`, and does not contain `pandas`, Excel COM automation, or platform-specific Excel dependencies. | automated | mandatory | Pytest static assertion in `python/tests`; verification command `python -m pytest python/tests`. | Satisfies dependency/static assertions and security dependency rule. |
| TD-002 | Spec: CLI contract | Assert the CLI accepts one required JSON path, optional `--output`, and optional `--table`; without `--output`, the resolved path uses the input directory/base name with `.xlsx`. | automated | mandatory | Helper-level or CLI-level pytest using `tmp_path`. | Covers default output path and CLI contract. |
| TD-003 | Spec: Schema-aware table selection | Given `schemaName: sdd-review.review-report` and no `--table`, assert the selected/exported sheet is `reviewMatrix`. | automated | mandatory | `tmp_path` JSON fixture and openpyxl workbook read-back showing `reviewMatrix` sheet. | Required exact case: automatic `reviewMatrix` detection for `sdd-review.review-report`. |
| TD-004 | Spec: Schema-aware table selection | Given `schemaName: sdd-review-security.review-security-report` and compact rows at `compactControlValidation.rows`, assert the selected/exported sheet is `compactControlValidation.rows`. | automated | mandatory | `tmp_path` JSON fixture and openpyxl workbook read-back showing `compactControlValidation.rows` sheet. | Required exact case: automatic nested path detection for current `sdd-review-security.review-security-report`. |
| TD-005 | Spec: Unknown schema requires table | Given an unrecognized `schemaName` and no `--table`, assert the command/helper fails clearly and asks for `--table`. | automated | mandatory | Pytest assertion over raised error or subprocess stderr/exit code. | Required exact case: clear error for unknown schema without `--table`; fail-closed behavior. |
| TD-006 | Spec: Schema-aware table selection; Design: explicit table override | Given an unknown schema with `--table customRows` or `--table compactControlValidation.rows`, assert the exporter uses the requested top-level table or nested dotted path and creates an Excel-safe sheet named from that full path. | automated | mandatory | `tmp_path` JSON fixture and openpyxl workbook read-back. | Required exact case: manual `--table` supports both top-level and nested dotted paths; chosen sheet naming uses sanitized full path. |
| TD-007 | Spec: Workbook structure and flattening; Design: failure boundaries | Given a requested table name/path that does not exist, assert failure occurs before workbook save with a concise actionable error mentioning the requested table/path. | automated | mandatory | Pytest assertion on no output workbook plus safe error text. | Required exact case: error when table/path does not exist. |
| TD-008 | Spec: Workbook sheets are generated | Assert a valid report generates a real `.xlsx` file readable by `openpyxl.load_workbook`. | automated | mandatory | `tmp_path` output path exists and openpyxl loads the workbook. | Required exact cases: real `.xlsx` generation and openpyxl workbook read-back validation. |
| TD-009 | Spec: Workbook structure and flattening | Assert the `summary` sheet contains expected scalar metadata such as `schemaName`, `schemaVersion`, `changeName`, verdict/status-like scalar values, and excludes the exported table body. | automated | mandatory | openpyxl read-back assertions for `summary` sheet rows/cells. | Required exact case: expected content in `summary` sheet. |
| TD-010 | Spec: Workbook structure and flattening | Assert the selected matrix/table sheet contains expected headers and row values, preserving scalars directly, scalar lists as `; `-joined text, and nested objects/complex lists as compact JSON text. | automated | mandatory | openpyxl read-back assertions for the matrix/table sheet. | Required exact case: expected content in matrix/table sheet. |
| TD-011 | Spec: Pytest verification | All JSON fixtures used by exporter tests must be created under pytest `tmp_path` rather than relying on repository-local generated files. | automated | mandatory | Test source assertion or review evidence plus pytest execution. | Required exact case: tmp_path JSON fixtures. |
| TD-012 | Spec: Workbook readability formatting | Assert generated sheets have bold header cells, autofilter enabled, `freeze_panes == "A2"`, reasonable non-default column widths, and wrap text for long text cells. | automated | mandatory | openpyxl workbook read-back assertions. | Required formatting assertions: bold header, autofilter, freeze panes `A2`, reasonable widths, wrap text for long text. |
| TD-013 | Spec: Pytest verification | Run the complete test suite using `python -m pytest python/tests`. | automated | mandatory | Verify phase command output. | Verification-blocking command for this change. |
| TD-014 | Proposal/Design non-goal; Secure Design: Files Rules | Assert the exporter reads JSON only and has no Markdown parsing requirement or Markdown input path behavior. | automated | mandatory | Static pytest assertion over source or behavior test proving `.md` reports are not parsed as source. | Safe-evidence/security-oriented case: no Markdown parsing requirement. |
| TD-015 | Secure Design: Sensitive Logging Rules | Error tests should assert messages are concise and do not dump raw JSON payloads, generated workbook contents, stack traces, secrets, tokens, PII-like fixture values, or full table rows where practical. | automated | mandatory | Pytest failure-path assertions using sentinel sensitive-looking values in `tmp_path` fixtures that must not appear in stderr/error text. | Safe-evidence/security-oriented case: no payload dumps in errors if practical. |
| TD-016 | Secure Design: JSON Parsing and Dependency Rules | Assert malformed JSON, non-object top-level JSON, and selected top-level or nested-path tables that are not lists of objects fail closed without writing an output workbook. | automated | mandatory | Pytest assertions on raised errors/exit code and absent output file. | Complements unknown-schema fail-closed behavior. |
| TD-017 | Spec: User documentation | Assert or manually review that `python/README.md` documents virtualenv setup, dependency installation, generation from canonical `review-report.json`, `--table`, dependency policy, and `python -m pytest python/tests`. | static | mandatory | Static assertion over README text or documented review evidence in apply/verify. | Required by documentation scenario; can be automated with file-content assertions. |
| TD-018 | Design: local file output; Operational considerations | Assert output path behavior is documented and predictable: default path derives from input `.json`, explicit `--output` is honored, and failures do not delete unrelated files. | automated | non-mandatory | Pytest using `tmp_path` sentinel files plus README evidence. | Non-blocking unless implementation introduces unsafe side effects. |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Local CLI usage | Verify README usage examples and pytest command demonstrate local execution against canonical JSON and optional `--table`. | static | `python/README.md` content plus `python -m pytest python/tests` result. | No pre-existing repo runner; pytest is introduced by this change. |
| Generated artifact handling | Verify documentation states `.xlsx` is derived output from JSON and may overwrite the selected output path; tests should avoid preserving workbook bytes as SDD evidence. | static | README section and sanitized test assertions over paths/sheet metadata. | Coverage/lint/type/format tooling unavailable. |
| Runtime operations | Confirm no service, scheduler, database, monitoring, authentication, deployment, or background operational surface is introduced. | static | Review of created files limited to `python/` utility and docs. | No runtime/build/deploy tooling applicable. |
| Safe evidence boundary | Test and review evidence must cite commands, paths, workbook structure, and sanitized summaries only; no raw workbook bytes or full sensitive payload dumps. | static | Test source and verify/report evidence. | Missing dedicated secret scanner is unavailable evidence, not a pass. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-FILES-LOCAL-IO` | Read only the explicit JSON path and write only the resolved output workbook path; do not scan directories, fetch remote content, parse Markdown, require Excel, or use platform-specific automation. | Yes | TD-002, TD-014, TD-018 plus README/static review. | covered | None |
| `SEC-FILES-FAIL-CLOSED` | Fail before workbook save on unreadable JSON, invalid JSON, unknown schema without `--table`, missing table, or unsupported table shape. | Yes | TD-005, TD-007, TD-016. | covered | None |
| `SEC-LOGGING-SAFE-ERRORS` | Errors must be concise/actionable and must not dump raw payloads, workbook content, stack traces, secrets, tokens, PII, PAN, or confidential values. | Yes | TD-005, TD-007, TD-015, TD-016. | covered | None |
| `SEC-JSON-STDLIB-NO-EVAL` | Use standard JSON parsing semantics; never execute/evaluate JSON content or formulas, and serialize nested object values as compact text only. | Yes | TD-010, TD-016 plus static review during apply/review. | covered | None |
| `SEC-DEPS-MINIMAL` | Depend on `openpyxl` and `pytest` only for this utility; do not introduce pandas, Excel automation, or platform-specific Excel dependencies. | Yes | TD-001. | covered | None |
| `SEC-EVIDENCE-SAFE` | SDD evidence must cite sanitized commands, paths, summaries, and workbook structure checks; do not include generated workbook bytes or full exported sensitive content. | Yes | TD-009, TD-010, TD-012, TD-015 and operational safe-evidence check. | covered | None |

## No-Impact Assessment

Not applicable. The change has behavior and testability impact because it adds a new Python CLI, dependencies, documentation, tests, local JSON input handling, and `.xlsx` output generation.

## Evidence Expectations

- Mandatory cases require implementation, execution, static/manual evidence, or a justified skip.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- The verification command for this change is `python -m pytest python/tests`.
- Tests must use `tmp_path` JSON fixtures and read generated workbooks back with `openpyxl`.
- Formatting evidence must include bold headers, autofilter, `A2` freeze panes, reasonable widths, and long-text wrapping.
- Dependency evidence must prove `requirements.txt` contains `openpyxl>=3.1,<4` and `pytest`, and excludes pandas/Excel automation dependencies.
- Security validation evidence should cite embedded `design.md` narrative rules, owner phase, and planned automated/static evidence.
- Applicable narrative category rules require planned safe evidence. Exhaustive source-row validation coverage, `N/A` decisions, and missed-applicable validation remain owned by canonical `review-security-report.json`.
- Test-design consumes narrative design rules only and MUST NOT require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, or the full 155-row Source ID matrix; exhaustive validation coverage belongs to canonical `review-security-report.json`.
- Runtime tests are introduced for this change through pytest. Coverage, lint, typecheck, format, and pre-existing build commands remain unavailable evidence and must not be reported as passing checks.

## Open Questions

- None.
