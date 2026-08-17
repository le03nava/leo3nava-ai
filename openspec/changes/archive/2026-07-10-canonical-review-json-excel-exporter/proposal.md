# Proposal: Canonical Review JSON Excel Exporter

## Intent

Provide a portable Python utility that converts canonical SDD review JSON reports into `.xlsx` workbooks for human review and sharing. JSON remains the source of truth; Markdown is never parsed.

## Scope

### In Scope
- Add `python/` with a minimal dependency file, usage README, CLI script, and pytest coverage.
- Export `review-report.json` table `reviewMatrix` and current `review-security-report.json` table path `compactControlValidation.rows` through schema-based default detection.
- Generate `.xlsx` files using `openpyxl` without Excel, pandas, or platform-specific dependencies.
- Validate behavior with `python -m pytest python/tests`.

### Out of Scope
- Changing review or security-review JSON schemas.
- Parsing Markdown reports or treating Markdown as authoritative.
- Adding workbook generation to SDD review phases automatically.
- Supporting arbitrary complex relational exports beyond one selected table/path plus scalar metadata.

## Capabilities

### New Capabilities
- `canonical-review-json-excel-exporter`: Converts canonical review JSON report tables into portable Excel workbooks.

### Modified Capabilities
- None.

## Approach

Implement `python/json_report_to_excel.py` as an argparse CLI accepting a required JSON path plus optional `--output` and `--table`. If `--table` is absent, map `schemaName` values to default table paths: `sdd-review.review-report` → `reviewMatrix`; `sdd-review-security.review-security-report` → `compactControlValidation.rows`. Unknown schemas fail clearly and ask for `--table`. Manual `--table` values may be top-level table names or nested dotted paths such as `compactControlValidation.rows`.

The workbook will include a `summary` sheet for scalar metadata and one table sheet. Rows will flatten scalar values directly, lists as `; `-joined text, and nested objects as compact JSON. Formatting will include bold headers, autofilter, `A2` freeze panes, reasonable column widths, and wrapped long text.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `python/` | New | Python utility package area for script, README, requirements, and tests. |
| `openspec/changes/canonical-review-json-excel-exporter/` | New | Planning artifacts for this change. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Future JSON table shape differs from current assumptions | Med | Keep table selection explicit via `--table` and use generic flattening. |
| Excel formatting expectations grow beyond basic export | Low | Limit first slice to readability basics and document non-goals. |
| Dependency drift | Low | Pin minimal portable dependency range in `python/requirements.txt`. |

## Rollback Plan

Remove `python/` and the OpenSpec change artifacts for `canonical-review-json-excel-exporter`. No runtime SDD workflow behavior or canonical report generation is modified.

## Dependencies

- Python environment capable of installing `openpyxl>=3.1,<4` and `pytest` from `python/requirements.txt`.

## Success Criteria

- [ ] `python python/json_report_to_excel.py openspec/changes/sdd-review-json-contract/review-report.json` creates a valid `.xlsx`.
- [ ] `python -m pytest python/tests` passes.
- [ ] Export logic supports schema defaults, manual `--table`, clear failures, and reusable flattening without pandas or Excel.
