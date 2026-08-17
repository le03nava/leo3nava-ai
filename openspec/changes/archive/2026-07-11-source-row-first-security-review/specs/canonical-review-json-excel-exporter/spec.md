# Delta for canonical-review-json-excel-exporter

## MODIFIED Requirements

### Requirement: Schema-aware table selection

The exporter MUST select reusable table data by top-level table name or nested dotted path. Without `--table`, schema `sdd-review.review-report` MUST default to `reviewMatrix`, and schema `sdd-review-security.review-security-report` MUST default to `sourceRowValidation.rows`. Unknown schemas without `--table` MUST fail clearly and ask the user to provide `--table`. (Previously: security-review defaulted to `compactControlValidation.rows`.)

#### Scenario: Known review schema

- GIVEN JSON with `schemaName: sdd-review.review-report`
- WHEN the CLI runs without `--table`
- THEN it exports the `reviewMatrix` table.

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

### Requirement: User documentation

`python/README.md` MUST document virtualenv setup, dependency installation, generation from canonical review JSON, `--table`, dependency policy, and the security-review default table `sourceRowValidation.rows`. It MUST NOT describe compact `SEC-*` controls as the active security-review export contract. (Previously: README documented `--table` generally and accepted the prior security default.)

#### Scenario: README explains security export

- GIVEN a user wants an Excel workbook from security-review JSON
- WHEN they read `python/README.md`
- THEN they can run the CLI without `--table` for `sourceRowValidation.rows`
- AND compact-control export is not presented as the active default.

### Requirement: Pytest verification

The tests MUST validate CLI defaults, schema defaults, unknown-schema failure, output derivation, workbook read-back, flattening, formatting, dependency constraints, nested dotted paths, and security-review default export from `sourceRowValidation.rows`. Tests MUST prove `compactControlValidation.rows` is not the active no-argument default. (Previously: tests validated the compact nested security default.)

#### Scenario: Test suite validates behavior

- GIVEN the implementation is complete
- WHEN `python -m pytest python/tests` is run
- THEN tests pass and validate generated `.xlsx` files through `openpyxl`.

#### Scenario: Security default regression is covered

- GIVEN a sample security-review JSON uses `sourceRowValidation.rows`
- WHEN exporter tests run
- THEN default table selection MUST export that path
- AND a compact-only active sample MUST fail without manual `--table`.
