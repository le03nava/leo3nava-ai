# Canonical Review JSON Excel Exporter Specification

## Purpose

Define a portable Python utility that exports canonical SDD review JSON tables to `.xlsx` workbooks while keeping JSON as the source of truth.

## Requirements

### Requirement: Python utility layout and dependencies

The system MUST add a root `python/` folder containing the exporter CLI, README, dependency file, and pytest tests. `python/requirements.txt` MUST include `openpyxl>=3.1,<4` and `pytest`. The exporter MUST NOT require Excel, pandas, or platform-specific dependencies.

#### Scenario: Dependencies are installable

- GIVEN a fresh Python virtual environment
- WHEN `pip install -r python/requirements.txt` is run
- THEN `openpyxl` in the `>=3.1,<4` range and `pytest` are installed
- AND no pandas or Excel dependency is required

### Requirement: User documentation

`python/README.md` MUST document virtualenv setup, dependency installation, generation from canonical review JSON, `--table`, dependency policy, and the security-review default table `sourceRowValidation.rows`. It MUST NOT describe compact `SEC-*` controls as the active security-review export contract.

#### Scenario: README explains normal generation

- GIVEN a user wants an Excel workbook from canonical JSON
- WHEN they read `python/README.md`
- THEN they can create a virtualenv, install requirements, run the CLI against `review-report.json`, and override the table with `--table`

#### Scenario: README explains security export

- GIVEN a user wants an Excel workbook from security-review JSON
- WHEN they read `python/README.md`
- THEN they can run the CLI without `--table` for `sourceRowValidation.rows`
- AND compact-control export is not presented as the active default.

### Requirement: CLI contract

`python/json_report_to_excel.py` MUST expose an argparse CLI with one required JSON path, optional `--output`, and optional `--table`. If `--output` is omitted, the workbook path MUST use the same directory and base name as the JSON input with `.xlsx` extension.

#### Scenario: Default output path

- GIVEN `reports/review-report.json`
- WHEN the CLI runs without `--output`
- THEN it writes `reports/review-report.xlsx`

#### Scenario: Explicit output path

- GIVEN a valid report JSON and `--output custom.xlsx`
- WHEN the CLI runs
- THEN it writes the workbook to `custom.xlsx`

### Requirement: Schema-aware table selection

The exporter MUST select reusable table data by top-level table name or nested dotted path. Without `--table`, schema `sdd-review.review-report` MUST default to `reviewMatrix`, and schema `sdd-review-security.review-security-report` MUST default to `sourceRowValidation.rows`. Unknown schemas without `--table` MUST fail clearly and ask the user to provide `--table`.

#### Scenario: Known review schema

- GIVEN JSON with `schemaName: sdd-review.review-report`
- WHEN the CLI runs without `--table`
- THEN it exports the `reviewMatrix` table

#### Scenario: Known security-review schema

- GIVEN JSON with `schemaName: sdd-review-security.review-security-report`
- AND source rows at `sourceRowValidation.rows`
- WHEN the CLI runs without `--table`
- THEN it exports the `sourceRowValidation.rows` table path.

#### Scenario: Legacy compact path is not default

- GIVEN security-review JSON contains only `compactControlValidation.rows`
- WHEN the CLI runs without `--table`
- THEN it MUST fail clearly because active security-review JSON requires `sourceRowValidation.rows`.

#### Scenario: Manual nested table path

- GIVEN JSON with rows at a nested dotted path
- WHEN the CLI runs with `--table` for that path
- THEN it exports that nested table path if it is a list of objects.

#### Scenario: Missing nested path fails clearly

- GIVEN JSON without the selected nested path
- WHEN the CLI runs with that nested `--table` path
- THEN it exits before saving with a clear error mentioning the requested table/path

#### Scenario: Invalid nested path value fails clearly

- GIVEN JSON where the selected nested path does not resolve to a list of objects
- WHEN the CLI runs with that nested `--table` path
- THEN it exits before saving with a clear error mentioning that the selected table/path must be a list of objects

#### Scenario: Unknown schema requires table

- GIVEN JSON with an unrecognized `schemaName`
- WHEN the CLI runs without `--table`
- THEN it exits with a clear error requesting `--table`

### Requirement: Workbook structure and flattening

The workbook MUST contain a `summary` sheet with scalar JSON metadata and one table sheet named after the exported table/path using an Excel-safe sheet name. For nested dotted paths, the sheet SHOULD use the sanitized full path, for example `compactControlValidation.rows`, rather than only the last segment, as long as it is valid for Excel. Table rows MUST flatten scalar values directly, lists as `; `-joined text, and nested objects as compact JSON. The implementation SHOULD remain generic and MUST NOT be hard-coupled only to `reviewMatrix`.

#### Scenario: Workbook sheets are generated

- GIVEN a valid report with scalar metadata and a selected table
- WHEN the exporter creates a workbook
- THEN the workbook has `summary` and the selected table sheet
- AND flattened cells preserve scalar, list, and nested object values readably

### Requirement: Workbook readability formatting

The exporter MUST apply bold headers, autofilter, `A2` freeze panes, reasonable column widths, and wrapped long text to generated sheets.

#### Scenario: Formatting is present

- GIVEN a generated workbook
- WHEN it is opened with `openpyxl`
- THEN headers are bold, filters are enabled, panes freeze at `A2`, widths are readable, and long text wraps

### Requirement: Pytest verification

The tests MUST validate CLI defaults, schema defaults, unknown-schema failure, output derivation, workbook read-back, flattening, formatting, dependency constraints, nested dotted paths, and security-review default export from `sourceRowValidation.rows`. Tests MUST prove `compactControlValidation.rows` is not the active no-argument default.

#### Scenario: Test suite validates behavior

- GIVEN the implementation is complete
- WHEN `python -m pytest python/tests` is run
- THEN tests pass and validate generated `.xlsx` files through `openpyxl`

#### Scenario: Security default regression is covered

- GIVEN a sample security-review JSON uses `sourceRowValidation.rows`
- WHEN exporter tests run
- THEN default table selection MUST export that path
- AND a compact-only active sample MUST fail without manual `--table`.
