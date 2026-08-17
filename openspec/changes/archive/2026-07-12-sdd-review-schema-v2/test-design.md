# Test Design: sdd-review Schema v2 — Reference Artifact Cleanup

## Overview

This change modifies three static reference files in `skills/sdd-review/references/`: the control catalog JSON, the report JSON schema, and the Markdown report template. There is no application runtime, no executable test runner, and no build tooling available in this repository (see `openspec/config.yaml`). All verification is static inspection of the produced artifacts against the v2 contracts defined in the spec and design. All 27 behavioral test cases are `static` or `manual` checks. No automated runner, coverage command, linter, type checker, or formatter is available; missing tooling is reported as an explicit constraint and is not treated as passing evidence.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/sdd-review-schema-v2/proposal.md` | Intent, scope, non-goals, affected areas, risks, success criteria |
| Spec | `openspec/changes/sdd-review-schema-v2/specs/sdd-review-workflow/spec.md` | RFC 2119 requirements and acceptance scenarios for catalog v2, schema v2, template v2 |
| Design | `openspec/changes/sdd-review-schema-v2/design.md` | Architecture decisions, file changes, exact v2 target shapes, data flow, migration notes |
| Secure Development Design | `openspec/changes/sdd-review-schema-v2/design.md#secure-development-design` | No-impact classification; all 8 security categories omitted as non-applicable; evidence safety rules preserved verbatim; known dependency risk on sdd-review-security |
| Testing Capabilities | `openspec/config.yaml` (testing section) | All tooling unavailable: test runner, unit, integration, e2e, coverage, linter, type checker, formatter |

## Source ID Coverage Baseline

**Classification**: No-impact (per `design.md#Secure Development Design`).

The changed artifacts are static JSON/Markdown reference files consumed as LLM instructional context. They do not handle authentication, sessions, sensitive data, secrets, permissions, files, database access, or sensitive logging. All eight corporate security categories are omitted as non-applicable for this surface.

Static/manual evidence for the `additionalProperties: false` structural tightening (which reduces — not increases — the schema's unexpected-field surface) is covered by TC-016 and TC-017.

The `## Safe Evidence Rules` verbatim preservation check (TC-024) is the primary evidence safety control for this change.

No source-row matrix expansion is required; exhaustive Source ID validation belongs to `review-security-report.json`.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-001 | Spec: Scenario: Catalog loads at v2 | Inspect `review-control-catalog.json` — confirm `schemaVersion` equals `2` | static | mandatory | File read; `"schemaVersion": 2` present at root | No runtime validator; visual/tool-assisted JSON inspection |
| TC-002 | Spec: Requirement: Canonical Review JSON Authority | Inspect `review-control-catalog.json` root — confirm `"reviewer": "sdd-review"` field present | static | mandatory | File read; `"reviewer": "sdd-review"` at root | — |
| TC-003 | Spec: Scenario: Catalog loads at v2 | Inspect `review-control-catalog.json` root — confirm NO `generatedHumanViewRef` field | static | mandatory | Field absent from root keys | — |
| TC-004 | Spec: Scenario: Catalog loads at v2 | Inspect `review-control-catalog.json` root — confirm NO `presentation` field | static | mandatory | Field absent from root keys | — |
| TC-005 | Spec: Scenario: Catalog loads at v2 | Inspect `review-control-catalog.json` root — confirm NO `validation` field | static | mandatory | Field absent from root keys | — |
| TC-006 | Spec: Scenario: Vocabulary is minimal | Inspect `vocabulary` object — confirm keys are only `complies` and `severity`; no `safeEvidence`, `forbiddenEvidence`, or `reviewer` | static | mandatory | `vocabulary` object has exactly 2 keys | — |
| TC-007 | Spec: Scenario: Per-control reviewer field is absent | Inspect all 96 control entries — confirm none contains a `reviewer` field | static | mandatory | Zero occurrences of `"reviewer"` key in any control entry | Inspect full controls array |
| TC-008 | Spec: Scenario: Blank notes are absent | Inspect all 96 control entries — confirm none contains `"notes": "Not provided"` | static | mandatory | Zero occurrences of `"notes": "Not provided"` in any control entry | Controls with meaningful notes are allowed to retain their `notes` field |
| TC-009 | Spec: Scenario: All controls are represented | Count entries in `controls` array — confirm exactly 96 | static | mandatory | `controls.length === 96` | — |
| TC-010 | Spec: Requirement: Review Report Artifact | Inspect `review-report.schema.json` `required` array — confirm exactly 14 fields in order: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `matrix`, `artifactMetadata` | static | mandatory | `required` array length 14; exact field names in exact order | — |
| TC-011 | Spec: Scenario: Legacy fields are rejected | Inspect `review-report.schema.json` — confirm NO occurrence of: `blockingSummary`, `evidenceSummary`, `runtimeChecks`, `inputsInspected`, `validation`, `changedFileSecurityHandoff`, `operationalEvidenceSummary` | static | mandatory | Zero occurrences of any legacy field name as a property key | Search full file content |
| TC-012 | Spec: Requirement: Review Report Artifact — $defs | Inspect `review-report.schema.json` `$defs` — confirm both `ref` and `stateArtifactRef` definitions are present | static | mandatory | `$defs.ref` and `$defs.stateArtifactRef` objects exist | — |
| TC-013 | Spec: Scenario: All controls are represented | Inspect `matrixRow` `$def` `item` pattern — confirm it matches `REV-CORP-001` through `REV-CORP-096` | static | mandatory | Pattern `^REV-CORP-(00[1-9]\|0[1-8][0-9]\|09[0-6])$` or equivalent present | — |
| TC-014 | Spec: Scenario: Finding column is populated | Inspect `matrixRow` `$def` — confirm `finding` field is defined as `type: string` (not an enum) | static | mandatory | `finding` present with `"type": "string"`, no `enum` constraint | — |
| TC-015 | Spec: Scenario: stateRegistration has two entries | Inspect `artifactMetadata` definition — confirm `stateRegistration` array has `minItems: 2` | static | mandatory | `stateRegistration.minItems === 2` | — |
| TC-016 | Spec: Requirement: Review Report Artifact — additionalProperties | Inspect `review-report.schema.json` root — confirm `additionalProperties: false` | static | mandatory | `"additionalProperties": false` at root object | — |
| TC-017 | Spec: Requirement: Review Report Artifact — additionalProperties | Inspect all nested object definitions (`totals`, `matrixRow`, `stateArtifactRef`, `ref`, `artifactMetadata`) — confirm each has `additionalProperties: false` | static | mandatory | `"additionalProperties": false` in each nested object `$def` | Enumerate all object definitions in the schema |
| TC-018 | Spec: Scenario: Field mapping table is updated | Inspect `report-template.md` — confirm `## JSON Field Mapping` section present with table mapping all 14 v2 fields; confirm NO legacy fields (`blockingSummary`, `inputsInspected`, `runtimeChecks`, `evidenceSummary`, `operationalEvidenceSummary`, `changedFileSecurityHandoff`, `validation`, `presentation`) | static | mandatory | Section present; table has 14 rows; no legacy field name in table | — |
| TC-019 | Spec: Requirement: Report Template v2 Structure | Inspect `report-template.md` — confirm `## Required Structure` section present showing full matrix skeleton | static | mandatory | Section heading `## Required Structure` present with matrix spec content | — |
| TC-020 | Spec: Scenario: Matrix columns include Finding | Inspect matrix table header in `report-template.md` — confirm exact column sequence: `Item \| Requirement \| Standard \| Category \| Severity \| Complies \| Finding \| Evidence Location \| Notes` | static | mandatory | Markdown table header row matches exact column sequence (9 columns) | No extra or renamed columns |
| TC-021 | Spec: Requirement: Report Template v2 Structure — Matrix Rules | Inspect `report-template.md` — confirm matrix join instruction is present: `matrix[].item` joined with `catalog controls[].id` | static | mandatory | Text referencing `matrix[].item` and `catalog controls[].id` join rule present | — |
| TC-022 | Spec: Scenario: Validation sections are absent | Inspect `report-template.md` section headings — confirm `## Matrix Validation` is absent | static | mandatory | No `## Matrix Validation` heading in file | Search full file |
| TC-023 | Spec: Scenario: Validation sections are absent | Inspect `report-template.md` section headings — confirm `## Catalog Validation` is absent | static | mandatory | No `## Catalog Validation` heading in file | Search full file |
| TC-024 | Spec: Scenario: Safe Evidence Rules are verbatim | Inspect `report-template.md` — confirm `## Safe Evidence Rules` section present and content verbatim: (1) Evidence may cite paths, section anchors, sanitized summaries, and unavailable-tooling statements only. (2) Evidence MUST NOT include secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production IDs/production identifiers, generated bytes, or final-document-only values. (3) Missing runtime/build/lint/type-check/format/coverage tooling is recorded as unavailable evidence, never as passing evidence. (4) This template does not define or require Excel, Python, script, spreadsheet, or workbook generation. | static | mandatory | Section present; all 4 lines match exactly | Exact verbatim match required |
| TC-025 | Spec: Requirement: Report Template v2 Structure — pre-field-mapping content unchanged | Inspect `report-template.md` lines before `## JSON Field Mapping` — confirm content is identical to v1 (lines 1–8 unchanged) | static | mandatory | Pre-field-mapping content unchanged from v1 | Compare against git diff of the file |
| TC-026 | Spec: Scenario: Reference artifacts are coherent | Cross-inspect all three files: catalog `schemaVersion: 2` consistent with schema `schemaVersion const 2`; template `## JSON Field Mapping` references exactly the 14 schema `required` fields; template matrix columns match `matrixRow` $def fields | static | mandatory | No contradictions between catalog, schema, and template | Cross-artifact inspection |
| TC-027 | Spec: Scenario: sdd-review-security is deferred | Inspect `skills/sdd-review-security/` directory — confirm no files were modified (verify via git status or file timestamps) | static | mandatory | `git diff` or file-level inspection shows no changes in `skills/sdd-review-security/` | — |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Operational considerations | Design explicitly states `No aplica.` — no runtime services, logs, deployments, monitoring, or administration operations exist | static | `design.md#Operational Considerations` confirms `No aplica.` | N/A — no tooling applicable |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| N/A — no-impact classification | All 8 security categories (auth, sessions, sensitive-data-pan, secrets, permissions, files, database, sensitive-logging) are omitted | No | Design classifies changed surface as no-impact; static reference files with no runtime data handling | not-applicable | Design-documented no-impact classification with rationale; `additionalProperties: false` tightening confirmed by TC-016/TC-017; Safe Evidence Rules preservation confirmed by TC-024 |

## Evidence Expectations

- All 27 mandatory cases require static inspection during `sdd-apply` and `sdd-verify` phases. No runtime test runner exists.
- TC-028 (non-mandatory tooling constraint record) confirms unavailable tooling and is not claimed as passing evidence.
- Static inspection expected evidence: file read + field presence/absence check, JSON key enumeration, Markdown section/heading search, content comparison.
- Security validation: no-impact classification from `design.md#Secure Development Design` is the evidence. Evidence safety rules preservation is confirmed by TC-024.
- `sdd-review-security` dependency break is an explicit, acknowledged risk — not a test-design blocker. TC-027 confirms the deferred surface is untouched.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands are ALL unavailable (per `openspec/config.yaml`). They are reported here as unavailable, not as passed checks.

## Open Questions

- None. All design decisions are resolved.
