# Apply Progress: Source Row First Security Review

## Slice 1 Summary

**Status**: partial — Slice 1 implementation complete; overall change still has downstream propagation and Python exporter slices pending.

**Work unit**: PR 1 / Slice 1 — replace active `sdd-review-security` schema, catalog, template, and validation rules with source-row-first authority.

**Boundary respected**: This slice changed only `skills/sdd-review-security` active contract files plus this OpenSpec progress/task state. It did not implement shared verify/archive/agent propagation or Python exporter changes.

## Slice 2 Summary

**Status**: partial — Slice 2 implementation complete; overall change still has Python exporter slice pending.

**Work unit**: PR 2 / Slice 2 — propagate canonical source-row contract to shared contracts, verify/archive skills, agent wrappers, and the OpenSpec base-spec sync plan.

**Boundary respected**: This slice changed downstream/shared/agent contracts plus OpenSpec task/progress state. It did not implement Python exporter code, tests, or README. It also did not sync base specs directly because repository archive convention syncs accepted delta specs during `sdd-archive` after review/security-review/verify pass.

## Slice 3 Summary

**Status**: complete — Slice 3 implementation complete; all planned apply tasks for this change are now complete and ready for `sdd-review`.

**Work unit**: PR 3 / Slice 3 — update Python exporter default, pytest coverage, README, and final apply evidence.

**Boundary respected**: This slice changed only `python/json_report_to_excel.py`, `python/tests/test_json_report_to_excel.py`, `python/README.md`, and OpenSpec task/progress state. It preserved Slice 1 and Slice 2 changes and did not broaden SDD contract files.

## Completed Tasks

- [x] 1.1 Updated `skills/sdd-review-security/SKILL.md` so canonical JSON owns `sourceRowValidation.rows[155]` and legacy compact controls are forbidden as active authority.
- [x] 1.2 Updated `skills/sdd-review-security/references/review-security-report.schema.json` to require source-row validation with exactly 155 rows, required source-row fields, counts, grouping fields, and no `compactControlValidation` property.
- [x] 1.3 Updated `skills/sdd-review-security/references/validation-rules.md` for missing/duplicate/unknown Source IDs, required fields, safe validation errors, unsafe evidence rejection, and row-preserving `N/A` grouping.
- [x] 1.4 Transformed `skills/sdd-review-security/references/security-guideline-catalog.operational.json` and `.md` to source-row-first fields, expanded 155 concrete Source IDs, vocabularies, evidence expectations, owner phase, and route metadata.
- [x] 1.5 Simplified `skills/sdd-review-security/references/report-template.md` to lean JSON-derived Markdown with source-row navigation, grouped `N/A`, blockers/warnings, metadata, and full source-row matrix last.
- [x] 2.1 Updated `skills/_shared/sdd-security-contract.md`, `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-post-apply-gates.md`, and propagation-adjacent `skills/_shared/sdd-phase-common.md` to consume canonical source-row JSON summaries/refs and remove active compact-control authority from downstream contracts.
- [x] 2.2 Updated `skills/sdd-verify/SKILL.md` and `skills/sdd-archive/SKILL.md` to cite canonical source-row counts, warnings, exceptions, evidence refs, parity metadata, and to avoid copying or re-scoring the full matrix.
- [x] 2.3 Updated `agents/sdd/sdd-review-security.md`, `agents/sdd/sdd-verify.md`, and `agents/sdd/sdd-archive.md` to mirror source-row-first skill contracts without duplicating implementation rules.
- [x] 2.4 Prepared the accepted-spec sync plan for archive: archive must merge the active delta specs for `sdd-review-security-workflow`, `sdd-security-guideline-catalog`, `sdd-execution-persistence-contracts`, and `canonical-review-json-excel-exporter` into `openspec/specs/*/spec.md`; apply does not sync them early.
- [x] 3.1 Updated `python/json_report_to_excel.py` so `sdd-review-security.review-security-report` defaults to `sourceRowValidation.rows`; compact-only reports fail safely by default because that path is missing.
- [x] 3.2 Updated `python/tests/test_json_report_to_excel.py` for source-row default export, compact-only default failure, generic nested override, workbook read-back, sheet naming, flattening, and temporary output behavior.
- [x] 3.3 Updated `python/README.md` for source-row security-review default, virtualenv/dependency setup, `--table`, safe examples, and no compact active default.
- [x] 4.1 Static evidence recorded below for TD-001..TD-015, NA-001..NA-005, and MD-001..MD-005, including Slice 2 downstream TD-013/TD-014 propagation evidence.
- [x] 4.2 Ran `python -m pytest python/tests`; Python, pytest, and openpyxl were available and all exporter tests passed.
- [x] 4.3 Global runner/build/lint/type/format/coverage tooling unavailable per `openspec/config.yaml#testing`; no unavailable tooling was claimed as passing.
- [x] 4.4 Completed final safe-evidence inspection across apply evidence: paths, section anchors, sanitized summaries, command outcomes, and no restricted operational data or committed workbook bytes.

## Files Changed

| File | Action | What changed |
| --- | --- | --- |
| `skills/sdd-review-security/SKILL.md` | Modified | Reframed phase authority around canonical JSON and `sourceRowValidation.rows[155]`; removed active legacy compact-control ownership from the phase contract. |
| `skills/sdd-review-security/references/review-security-report.schema.json` | Modified | Replaced active matrix schema with source-row validation, 155-row exact count constraints, source-row grouping fields, row required fields, and source-row exceptions. |
| `skills/sdd-review-security/references/validation-rules.md` | Modified | Replaced compact-control rules with exact Source ID coverage, required fields, safe errors, unsafe evidence, and row-preserving non-applicability rules. |
| `skills/sdd-review-security/references/security-guideline-catalog.operational.json` | Modified | Converted catalog to source-row-first operational metadata with 155 expanded `sourceRows`, source-row vocabularies, grouping metadata, safe evidence expectation, owner phase, and route metadata. |
| `skills/sdd-review-security/references/security-guideline-catalog.md` | Modified | Regenerated human/audit view around source-row inventory, source-row grouping, vocabularies, and safe evidence expectations. |
| `skills/sdd-review-security/references/report-template.md` | Modified | Simplified derived Markdown structure and placed the full 155-row source matrix as the final major section. |
| `openspec/changes/source-row-first-security-review/tasks.md` | Modified | Marked Slice 1 task checkboxes complete and recorded 4.1 as partial evidence for this slice only. |
| `openspec/changes/source-row-first-security-review/apply-progress.md` | Created | Recorded Slice 1 progress, validation evidence, limitations, and next recommendation. |
| `skills/_shared/sdd-security-contract.md` | Modified | Replaced active compact-control downstream authority language with canonical `sourceRowValidation.rows` summary/ref consumption, source-row exception fields, and safe-evidence ownership. |
| `skills/_shared/persistence-contract.md` | Modified | Updated security-review persistence/read-back rules so downstream phases consume canonical source-row counts, warnings, exceptions, evidence refs, and parity metadata from JSON. |
| `skills/_shared/sdd-post-apply-gates.md` | Modified | Updated post-apply gates, source-row consumption boundary, and archive readiness blockers to remove active compact-control authority and preserve source-row summaries/refs only. |
| `skills/_shared/sdd-phase-common.md` | Modified | Updated shared status/evidence wording to reference source-row validation summaries rather than compact/source-row summaries. |
| `skills/sdd-verify/SKILL.md` | Modified | Updated verify decision gates, execution steps, and output contract to consume canonical source-row counts/evidence refs/parity metadata without matrix ownership. |
| `skills/sdd-archive/SKILL.md` | Modified | Updated archive source-row preservation/audit rules to preserve canonical refs, counts, warnings, exceptions, safe evidence refs, parity metadata, and no full-matrix copies. |
| `agents/sdd/sdd-review-security.md` | Modified | Added source-row-first wrapper reminder pointing to canonical JSON and derived Markdown compatibility. |
| `agents/sdd/sdd-verify.md` | Modified | Added verify wrapper reminder to consume canonical summaries/refs and avoid copying or re-scoring source-row matrices. |
| `agents/sdd/sdd-archive.md` | Modified | Updated archive wrapper reminder to preserve canonical refs, counts, warnings, exceptions, evidence refs, and parity metadata without full-matrix copies. |
| `openspec/changes/source-row-first-security-review/tasks.md` | Modified | Marked Slice 2 tasks 2.1-2.4, 4.1, and 4.3 complete; left 4.4 partial because exporter evidence remains pending. |
| `openspec/changes/source-row-first-security-review/apply-progress.md` | Modified | Merged Slice 2 progress with existing Slice 1 progress without dropping prior evidence. |
| `python/json_report_to_excel.py` | Modified | Changed the security-review schema default table path from the legacy compact path to `sourceRowValidation.rows`. |
| `python/tests/test_json_report_to_excel.py` | Modified | Added/updated tests for source-row default export, compact-only failure, generic nested override, workbook read-back, sheet naming, flattening, temp output, README assertions, and safe CLI errors. |
| `python/README.md` | Modified | Documented source-row security-review default, setup/dependencies, `--table`, safe examples, and generated workbook safety guidance without presenting compact export as active contract. |
| `openspec/changes/source-row-first-security-review/tasks.md` | Modified | Marked Slice 3 tasks 3.1-3.3, 4.2, and 4.4 complete. |
| `openspec/changes/source-row-first-security-review/apply-progress.md` | Modified | Merged Slice 3 progress with Slice 1 and Slice 2 evidence. |

## Validation Evidence

### Commands Run

| Command | Result | Evidence |
| --- | --- | --- |
| `python -m json.tool skills\sdd-review-security\references\review-security-report.schema.json` | Passed | Schema JSON parses successfully. |
| `python -m json.tool skills\sdd-review-security\references\security-guideline-catalog.operational.json` | Passed | Operational catalog JSON parses successfully. |
| Python static assertion command for schema/catalog invariants | Passed | Reported `schemaVersion=2 sourceRows=155 uniqueSourceIds=155 missingFields=0 sections=15 sectionTotal=155`; asserted schema min/max rows 155 and absence of `compactControlValidation`. |
| Grep tool search for `compactControlValidation`, `SEC-*`, `compactGuidelines`, `compactMapping`, and `Compact Control` under `skills/sdd-review-security` | Passed | No matches found in active Slice 1 security-review skill/reference files. |
| `rg ... skills/sdd-review-security` | Not available | Local shell reported `rg` is not recognized; replaced with the dedicated Grep tool. |
| Grep tool search for active compact-control authority markers under `skills/_shared`, `skills/sdd-verify`, `skills/sdd-archive`, and `agents/sdd` | Passed with one allowed forbidden-legacy reminder | No matches remained in shared/verify/archive skill files. Agent wrappers only mention legacy compact-control identifiers as forbidden non-authority text. |
| `git diff --check` | Passed with line-ending warnings | No whitespace errors; Git warned some working-copy files will normalize CRLF to LF when touched. |
| `git diff --stat` | Informational | Confirmed Slice 2 edits stayed in shared/verify/archive/agent/OpenSpec progress files; diff also includes pre-existing Slice 1 changes. |
| `python -m pytest python/tests` | Passed | 20 tests passed in 1.02s; Python 3.12.6, pytest 9.1.1, openpyxl available. |
| Grep tool search for `sourceRowValidation.rows` under `python/` | Passed | Found default mapping in exporter, README source-row default, and tests asserting the active default/read-back behavior. |
| Grep tool search for `compactControlValidation.rows` under `python/` | Passed with legacy-test-only matches | Matches remain only in tests for explicit negative/manual legacy path behavior and README assertion that the compact path is not documented. |
| Glob search for `**/*.xlsx` | Passed | No generated workbook files were found in the repository; tests use temporary output paths. |

### Test-Design Coverage: TD Cases

| Case | Evidence |
| --- | --- |
| TD-001 | `review-security-report.schema.json` requires `sourceRowValidation`, `rows` min/max 155, source-row grouping fields, and no `compactControlValidation`. |
| TD-002 | Positive sample execution was not created in this slice; static schema/catalog invariant assertions prove required shape and catalog count. Full sample validation remains review-security/verify work. |
| TD-003 | Validation rules now define missing Source IDs as blockers with safe count/ID errors. |
| TD-004 | Validation rules now define duplicate Source IDs as blockers with exact-once failure. |
| TD-005 | Validation rules now define unknown Source IDs as blockers that cannot compensate for missing catalog rows. |
| TD-006 | Schema requires source-row field categories and enforces `guidelineText`/`guidelineRefs` either-or; validation rules block missing required fields. |
| TD-007 | Grep under active `skills/sdd-review-security` found no legacy compact-control identifiers or active compact authority markers. |
| TD-008 | Schema requires row-level `applies`, `complies`, `justification`, `evidenceType`, `evidenceLocation`, `finding`, `ownerPhase`, and `route`; rules preserve these for `N/A`. |
| TD-009 | Validation rules reject unsafe evidence and require `unsafeEvidenceRejections` plus safe findings without copying unsafe values. |
| TD-010 | Validation rules and template restrict grouped non-applicability to source-row fields only. |
| TD-011 | `report-template.md` is lean, generated from JSON, avoids independent validation logic, and renders the full matrix last. |
| TD-012 | Template/rules state JSON wins and parity/read-back failures route to `resolve-blockers`. Runtime generator/parity tests are pending downstream execution. |
| TD-013 | Slice 2 updated `skills/sdd-verify/SKILL.md`, `skills/_shared/sdd-post-apply-gates.md`, `skills/_shared/persistence-contract.md`, and `agents/sdd/sdd-verify.md` so verify consumes canonical JSON summaries/refs, expected/validated counts, warnings, exceptions, evidence refs, and parity metadata without copying or re-scoring `sourceRowValidation.rows`. |
| TD-014 | Slice 2 updated `skills/sdd-archive/SKILL.md`, `skills/_shared/sdd-post-apply-gates.md`, `skills/_shared/persistence-contract.md`, and `agents/sdd/sdd-archive.md` so archive preserves canonical JSON refs before derived Markdown refs, counts, warnings, exceptions, evidence refs, and parity metadata without creating a second matrix authority. |
| TD-015 | Operational catalog parses, contains 155 expanded unique Source IDs, 15 source sections, section total 155, grouping metadata, owner phase, route metadata, and evidence expectations. |
| TD-016 | `python/json_report_to_excel.py` maps `sdd-review-security.review-security-report` to `sourceRowValidation.rows`; pytest covers default selection. |
| TD-017 | Pytest covers compact-only security-review JSON failing without `--table` because `sourceRowValidation.rows` is missing; no workbook is created. |
| TD-018 | Pytest preserves generic nested dotted-path override behavior using a neutral `artifactExport.rows` fixture. Legacy compact path remains generic only in negative/manual test coverage, not active documentation. |
| TD-019 | Pytest reads generated workbooks through `openpyxl`, validating sheet names, headers, row count, flattening, formatting, and temp output behavior. |
| TD-020 | README inspection through pytest and static grep confirms setup/dependencies, `--table`, `sourceRowValidation.rows`, safe examples, and no documented active compact default. |

### Safe N/A Grouping Coverage

| Case | Evidence |
| --- | --- |
| NA-001 | Schema and validation rules require every `N/A` row to preserve row-level applicability, compliance, justification, evidence, finding, owner phase, and route fields. |
| NA-002 | Schema/template/rules restrict grouped non-applicability to `controlDomain`, `corporateSection`, `repoProfiles`, `runtimeSurface`, or `dataSurface`. |
| NA-003 | Validation rules require splitting or blocking a group when rows lack equivalent safe justification or compatible owner/route. |
| NA-004 | Template requires grouped summaries to cite row IDs/counts and states groups do not override row-level JSON. |
| NA-005 | Grep found no legacy compact-control authority markers under active `skills/sdd-review-security` files. |

### Markdown Coverage

| Case | Evidence |
| --- | --- |
| MD-001 | Template maps verdict, route, refs, general review handoff, source-row navigation, summaries, blockers/warnings, unsafe evidence, exceptions, unavailable tooling, artifact metadata, and next recommendation from JSON. |
| MD-002 | Template places `## Full Source Row Matrix` as the final major content section. |
| MD-003 | Template contains no compact-control validation section. |
| MD-004 | Template states Markdown is generated presentation only and must not define independent validation logic. |
| MD-005 | Template/rules route JSON/Markdown parity failures to `resolve-blockers` because JSON is authoritative. |

### Python Exporter Coverage

| Case | Evidence |
| --- | --- |
| PY-001 | `DEFAULT_TABLE_BY_SCHEMA["sdd-review-security.review-security-report"]` is `sourceRowValidation.rows`; covered by pytest and static grep. |
| PY-002 | Compact-only security-review JSON fails by default with a safe missing-path error and does not create an `.xlsx`; covered by pytest. |
| PY-003 | Generic nested dotted-path override remains supported for a list of objects; covered by pytest with `artifactExport.rows`. |
| PY-004 | Workbook read-back through `openpyxl` validates source-row sheet naming, headers, flattened scalar lists/objects, formatting, and temp output paths. |
| PY-005 | README documents source-row default and does not describe legacy compact export as active; covered by pytest and static grep. |

## Deviations and Limitations

- This is a static/contract slice. No global runtime/build/lint/type/format/coverage runner exists in `openspec/config.yaml`; none was claimed as passing.
- Positive/negative sample JSON fixture execution was not added in this slice to keep the work unit focused and under the stacked PR boundary. The schema and rules now define the contract; generated report fixture validation remains appropriate for review-security/verify follow-up.
- TD-013 and TD-014 are now covered by Slice 2 static propagation evidence. Runtime verify/archive execution remains a downstream phase responsibility.
- Python exporter checks TD-016..TD-020 and PY-001..PY-005 are now covered by Slice 3 tests/static inspection.
- Diff size is still above the nominal 400-line review budget mainly because the catalog transformation touches the full 155-row source inventory. This is the smallest coherent Slice 1 boundary because schema/rules/template/catalog must agree.

## Security Evidence

- Safe evidence policy is explicitly recorded in `SKILL.md`, `validation-rules.md`, schema comments, catalog evidence expectations, and `report-template.md`.
- Unsafe evidence is rejected without copying raw payloads into reports or Markdown.
- Source-row grouping/navigation uses `controlDomain`, `corporateSection`, `repoProfiles`, `runtimeSurface`, and `dataSurface`.
- Canonical JSON remains the only active authority; derived Markdown is compatibility output only.
- Verify/archive/shared contracts now consume canonical `review-security-report.json` summaries, counts, warnings, exceptions, evidence refs, and parity metadata; they do not copy, reproduce, or re-score the full source-row matrix.
- Slice 2 evidence uses safe paths, section anchors, sanitized summaries, and command outcomes only. No restricted operational data, credentials, raw logs/payloads, production identifiers, full ID lists, or generated bytes were added.
- Slice 3 exporter tests use pytest temporary paths and `openpyxl` read-back; no generated workbook bytes were committed or recorded as SDD evidence.
- README safe-use guidance states the exporter does not sanitize unsafe source JSON and generated workbooks should not be committed as ordinary artifacts.

## Archive Sync Plan

Base spec synchronization is intentionally deferred to `sdd-archive`, matching `skills/sdd-archive/SKILL.md` and `openspec/config.yaml#rules.archive`.

During archive, merge these active delta specs into the corresponding accepted specs before moving the change folder:

| Delta spec | Accepted spec target | Archive action |
| --- | --- | --- |
| `openspec/changes/source-row-first-security-review/specs/sdd-review-security-workflow/spec.md` | `openspec/specs/sdd-review-security-workflow/spec.md` | Apply MODIFIED and REMOVED requirements for source-row-first canonical JSON and no active compact-control authority. |
| `openspec/changes/source-row-first-security-review/specs/sdd-security-guideline-catalog/spec.md` | `openspec/specs/sdd-security-guideline-catalog/spec.md` | Apply MODIFIED and REMOVED requirements for source-row catalog vocabulary and no active compact-control validation/navigation/grouping dependency. |
| `openspec/changes/source-row-first-security-review/specs/sdd-execution-persistence-contracts/spec.md` | `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Apply verify/archive source-row consumption and persistence compatibility requirements. |
| `openspec/changes/source-row-first-security-review/specs/canonical-review-json-excel-exporter/spec.md` | `openspec/specs/canonical-review-json-excel-exporter/spec.md` | Apply exporter default/readme/test requirements after Slice 3 implements Python code/tests/docs. |

## Remaining Tasks

- None for apply. Review, review-security, verify, and archive remain downstream SDD phases.

## Next Recommendation

Proceed to `sdd-review`. All apply tasks are complete in `tasks.md`, Python exporter tests passed, and cumulative apply evidence has been merged across all three slices.
