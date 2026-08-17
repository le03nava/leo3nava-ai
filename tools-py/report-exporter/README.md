# Canonical Review JSON Excel Exporter

This folder contains a small local utility that converts canonical SDD review JSON reports into derived `.xlsx` workbooks. JSON remains the source of truth; the workbook is only a reviewer-friendly export.

## Setup

Create and activate a virtual environment from the repository root:

```bash
python -m venv .venv
```

Windows PowerShell:

```powershell
.\.venv\Scripts\Activate.ps1
```

macOS/Linux:

```bash
source .venv/bin/activate
```

Install dependencies:

```bash
pip install -r python/requirements.txt
```

## Generate Excel from canonical review JSON

Prefer an active canonical review report when it exists:

```bash
python python/json_report_to_excel.py openspec/changes/example-change/review-report.json
```

If the active report is not available, use any local canonical JSON report path, including an archived or sample report:

```bash
python python/json_report_to_excel.py path/to/review-report.json
```

When `--output` is omitted, the workbook is written next to the JSON input using the same base name and `.xlsx` extension. For example, `reports/review-report.json` produces `reports/review-report.xlsx`.

Use `--output` to select an explicit workbook path:

```bash
python python/json_report_to_excel.py path/to/review-report.json --output exports/review-report.xlsx
```

Saving to an existing output path may overwrite that derived workbook. The utility does not delete unrelated files, scan directories, parse Markdown reports, fetch remote content, or require Microsoft Excel.

## Table selection

Known schemas choose a default table path:

- `sdd-review.review-report` exports `reviewMatrix`
- `sdd-review-security.review-security-report` exports `sourceRowValidation.rows`

Active security-review reports are source-row-first. A security-review JSON file that only contains legacy compact-control rows fails by default because the active table path is missing.

For future or unknown schemas, provide the table explicitly. The value can be a top-level table name or a nested dotted path:

```bash
python python/json_report_to_excel.py path/to/report.json --table customRows
python python/json_report_to_excel.py path/to/report.json --table artifactExport.rows
```

The selected table/path must resolve to a list of objects. Nested dotted paths use the full sanitized path as the worksheet name, for example `artifactExport.rows`, so the workbook remains understandable and Excel-safe.

The exporter does not sanitize unsafe source JSON. Only export reports that have already passed the relevant SDD safe-evidence checks, and avoid committing generated workbook files as ordinary repository artifacts.

## Dependency policy

`python/requirements.txt` intentionally contains only:

- `openpyxl>=3.1,<4`
- `pytest`

The exporter does not use pandas, Excel automation, platform-specific Excel dependencies, Markdown parsing, network access, or JSON evaluation.

## Verification

Run the test suite from the repository root:

```bash
python -m pytest python/tests
```
