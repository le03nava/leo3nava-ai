# Design: Canonical Review JSON Excel Exporter

## Technical Approach

Implement a small, portable Python CLI in `python/json_report_to_excel.py` that reads canonical SDD review JSON, selects one table by top-level name or nested dotted path, and writes a readable `.xlsx` workbook with `openpyxl`. The JSON file is the only source of truth; Markdown reports are never parsed. The exporter remains table-generic by using a schema-to-default-table-path map for known canonical report schemas and an explicit `--table` override for unknown or future tables.

The implementation should be a single script with focused pure functions so tests can exercise behavior without shelling out for every case. The CLI layer should only parse arguments, call the orchestration function, print concise success/error messages, and return deterministic exit codes.

## Architecture Decisions

### Decision: Single-file CLI with pure helper functions

**Choice**: Create `python/json_report_to_excel.py` with helpers for JSON loading, table resolution, metadata extraction, flattening, workbook writing, formatting, and CLI dispatch.
**Alternatives considered**: A package split into multiple modules, or embedding logic into SDD review phases.
**Rationale**: The spec asks for a portable standalone utility, not workflow automation. A single script minimizes setup while pure helpers keep tests precise and reusable.

### Decision: Schema defaults plus explicit table override

**Choice**: Use a constant such as `DEFAULT_TABLE_BY_SCHEMA = {"sdd-review.review-report": "reviewMatrix", "sdd-review-security.review-security-report": "compactControlValidation.rows"}` and let `--table` override it. The selected table reference may be a top-level table or nested dotted path.
**Alternatives considered**: Hard-code only `reviewMatrix`, infer the first array found, or require `--table` always.
**Rationale**: This satisfies current review JSON and current security-review JSON while avoiding reviewMatrix-only coupling. Unknown schemas fail clearly unless the user provides `--table`, preventing accidental export of the wrong array.

### Decision: Generic scalar/list/object flattening

**Choice**: Preserve scalar values directly, join lists with `; ` when all elements are scalar, and serialize nested objects or complex list elements as compact deterministic JSON.
**Alternatives considered**: Recursive column expansion or pandas normalization.
**Rationale**: The selected table path may evolve. Generic flattening keeps output readable without introducing pandas or a brittle schema-specific column model.

### Decision: openpyxl-only workbook generation

**Choice**: Use `openpyxl` directly for workbook creation, read-back validation in tests, and formatting.
**Alternatives considered**: pandas, Excel COM automation, or CSV-only export.
**Rationale**: `openpyxl` is portable, does not require Excel, and supports required formatting features: bold headers, filters, freeze panes, widths, and text wrapping.

## Data Flow

```text
CLI args ──→ resolve input/output paths ──→ load JSON
             │                              │
             │                              ├──→ resolve table path/data
             │                              ├──→ extract scalar metadata
             │                              └──→ flatten table rows/headers
             │
             └──────────────────────────────→ write workbook + formatting ──→ .xlsx
```

Failures should stop before workbook save when JSON is unreadable, invalid, unknown-schema-without-`--table`, missing table/path, or selected table/path is not a list of objects.

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `python/json_report_to_excel.py` | Create | Argparse CLI and exporter helpers. |
| `python/requirements.txt` | Create | Minimal dependencies: `openpyxl>=3.1,<4` and `pytest`. |
| `python/README.md` | Create | Setup, dependency policy, sample commands, default schema behavior, `--table`, and verification command. |
| `python/tests/test_json_report_to_excel.py` | Create | Pytest coverage using `tmp_path` and `openpyxl` read-back. |

## Interfaces / Contracts

### CLI

```text
python python/json_report_to_excel.py <json_path> [--output <xlsx_path>] [--table <table_name>]
```

- `json_path` is required and must point to a local JSON file.
- `--output` is optional; when omitted, the exporter writes next to the input with the same base name and `.xlsx` extension.
- `--table` is optional; when omitted, `schemaName` selects a default table path for known schemas. Manual values can be top-level table names or nested dotted paths such as `compactControlValidation.rows`.
- Exit code `0` means workbook created successfully. Exit code `1` means user-correctable input or write failure with a concise stderr message.

### Helper Functions

- `load_report(path: Path) -> dict`: read UTF-8 JSON and require a top-level object.
- `resolve_output_path(input_path: Path, output: str | None) -> Path`: derive or normalize the target workbook path.
- `resolve_table_name(report: dict, requested_table: str | None) -> str`: use `--table` or known schema defaults; raise a clear error for unknown schemas.
- `get_table_rows(report: dict, table_name: str) -> list[dict]`: require the selected top-level table or nested dotted path to be a list of objects.
- `make_excel_sheet_name(table_path: str) -> str`: sanitize the full table path into an Excel-safe sheet name, preserving readable dotted paths such as `compactControlValidation.rows` when valid.
- `extract_scalar_metadata(report: dict, table_name: str) -> dict`: include scalar top-level fields while excluding the exported table and complex report sections.
- `flatten_value(value) -> str | int | float | bool | None`: preserve scalars, join scalar lists, compact-JSON nested values.
- `flatten_rows(rows: list[dict]) -> tuple[list[str], list[list[Any]]]`: union headers in first-seen order and produce flattened row values.
- `write_workbook(metadata: dict, table_name: str, headers: list[str], rows: list[list[Any]], output_path: Path) -> None`: create `summary` and table/path sheets.
- `format_sheet(ws) -> None`: apply bold headers, autofilter, `A2` freeze panes, widths, and wrapped long text.

## Operational Considerations

### Strategy

This change adds a manual local utility only. It does not modify SDD phase execution, background jobs, services, monitoring, databases, authentication, or deployment behavior. Operational use is limited to developers or reviewers running the CLI against local canonical JSON files and sharing the generated workbook when appropriate.

### Evidence Plan

| Category | Expected safe evidence | Owner phase | Status |
| --- | --- | --- | --- |
| Local CLI usage | `python/README.md` usage examples and `python -m pytest python/tests` result | apply / verify | planned |
| Generated artifact handling | Documentation that `.xlsx` output is derived from JSON and may overwrite the selected output path | apply / verify | planned |
| Runtime operations | No aplica. No service, scheduler, database, or deployment surface is introduced. | design | complete |

### Restricted Data Boundary

Do not include generated workbook bytes, raw report payloads, secrets, tokens, PII, PAN, or environment-specific operational values in SDD evidence. Evidence should cite paths, commands, sanitized summaries, and workbook structure assertions only.

### Unresolved Gaps

- None.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Unit | Output path derivation, schema default resolution, `--table` override, unknown-schema failure, scalar/list/object flattening, metadata extraction | Import helper functions in pytest and use small in-memory fixtures. |
| Integration | End-to-end workbook generation for `reviewMatrix` and `compactControlValidation.rows`-style tables | Use `tmp_path` JSON inputs, run the exporter entry point or subprocess, and read generated `.xlsx` with `openpyxl`. |
| Formatting | Headers, autofilter, freeze panes, widths, wrapped text | Read back workbook with `openpyxl` and assert expected workbook/sheet properties. |
| Dependency constraints | `requirements.txt` includes `openpyxl>=3.1,<4` and `pytest`, and excludes pandas/Excel-specific dependencies | Static pytest assertion over the requirements file. |

Verification command: `python -m pytest python/tests`.

## Migration / Rollout

No migration required. This is an additive local utility under `python/` and does not change canonical report generation or SDD workflow behavior.

## Open Questions

- None.

## Secure Development Design

### Classification and Changed Surface

Classification: security-impacting due to local file input/output and export behavior. The changed surface is limited to a new manual Python CLI, its dependency file, README, and tests under `python/`. The CLI reads user-selected local JSON files and writes derived local `.xlsx` files. It does not introduce authentication, sessions, permissions logic, database access, network calls, secret storage, or automated SDD phase execution.

Relevant security catalog context is the files category for safe local file handling and the sensitive-logging category for safe error output. Sensitive-data and secrets handling are not intentionally introduced because the utility treats report JSON as local user-provided input and must not inspect, log, or transform sensitive values beyond writing the selected derived workbook. Omitted categories remain reviewable by `sdd-review-security`; design does not create all-row `N/A` bookkeeping or source-row matrices.

### Files Rules

The implementation must only read the explicit JSON path provided by the user and write the resolved output workbook path. It must not scan directories, parse Markdown, fetch remote content, require Excel, or depend on platform-specific automation. If `--output` is omitted, the output path must be derived by replacing the input suffix with `.xlsx`; if `--output` is provided, the exact requested target is used after normal path normalization.

Output overwrite behavior should be intentional and documented: saving to an existing output path may replace that derived workbook, but the utility must not delete unrelated files or create hidden side effects. Parent directories should be handled predictably: either require them to exist with a clear error, or create only the explicitly requested parent path if implementation chooses to document that behavior. Evidence owners are `test-design` for planned cases, `apply` for code/README evidence, `review-security` for final safe-evidence validation, and `verify` for the pytest command result. Residual risk: users can select a JSON file containing sensitive report content and generate a derived workbook; documentation and evidence must treat workbook contents as user-controlled local output, not safe public evidence.

### Sensitive Logging Rules

Errors must be concise and actionable without dumping raw JSON payloads, generated workbook contents, stack traces, secrets, tokens, PII, PAN, or confidential values. Failure messages may include safe path references and high-level reasons such as unreadable JSON, unknown schema, missing table, invalid table shape, or write failure. Tests should assert clear failure behavior without embedding sensitive fixtures. Evidence owners are `apply` for CLI error handling, `review-security` for checking safe-evidence expectations, and `verify` for pytest results. Residual risk: filesystem errors may include OS-provided path text; paths are acceptable evidence, but payload contents and credentials are not.

### JSON Parsing and Dependency Rules

JSON parsing must use Python's standard `json` module, require a top-level object, and fail closed on malformed input or unsupported selected-table shapes. The exporter must not execute content from JSON, evaluate formulas, parse Markdown, import pandas, or use Excel automation. Nested object values should be serialized as compact JSON text only for workbook readability. Dependency evidence belongs to `apply` and `verify` through `python/requirements.txt` and tests confirming `openpyxl`/`pytest` only.

### Exception and Evidence Policy

No security exceptions are planned. If a later implementation needs an exception, it must include approver, approved date, accepted-risk rationale, mitigation or follow-up, and exact evidence gap before archive readiness. All evidence must be review-safe: cite files, sections, commands, sanitized summaries, or workbook structure checks; never embed raw secrets, credentials, tokens, private keys, connection strings, PAN, PII, raw logs, sensitive payloads, generated workbook bytes, or full exported contents.
