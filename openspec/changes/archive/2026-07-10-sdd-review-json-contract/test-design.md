# Test Design: SDD Review JSON Contract

## Overview

This change converts `sdd-review` from Markdown-authored review data to canonical JSON authority for the static 96-control catalog and each per-change review report, while keeping `review-report.md` as a generated compatibility presentation. The repository has no executable test runner, build command, linter, type checker, formatter, or coverage tool, so verification is planned as static/manual artifact inspection and read-back evidence rather than invented commands.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Exploration | `openspec/changes/sdd-review-json-contract/explore.md` | Confirms the JSON-source-of-truth approach, derived Markdown compatibility, downstream compatibility, and deferred Excel scope. |
| Proposal | `openspec/changes/sdd-review-json-contract/proposal.md` | Defines intent, in-scope JSON contract work, Markdown generation, non-goals, risks, compatibility boundaries, and success criteria. |
| Spec | `openspec/changes/sdd-review-json-contract/specs/sdd-review-workflow/spec.md` | Provides requirements and scenarios for JSON authority, artifact persistence, matrix contract, N/A evidence, downstream consumption, and deferred Excel/Python generation. |
| Design | `openspec/changes/sdd-review-json-contract/design.md` | Defines architecture decisions, JSON/static catalog split, report JSON contract, Markdown generation/parity validation, operational considerations, and testing strategy. |
| Secure Development Design | `openspec/changes/sdd-review-json-contract/design.md#secure-development-design` | Supplies changed-surface classification, safe-evidence rules, file-scope constraints, evidence owners, residual risks, and exception policy. |
| Existing review skill | `skills/sdd-review/SKILL.md` | Baseline review behavior, current Markdown-only artifact contract, matrix ownership, routing, and current catalog dependency to be changed. |
| Existing review template | `skills/sdd-review/references/report-template.md` | Required Markdown sections, exact matrix header, vocabulary, N/A requirements, and routing expectations to preserve in generated Markdown. |
| Existing review catalog | `skills/sdd-review/references/control-catalog.md` | Current 96-control source to migrate into canonical JSON and demote to derived/reference-only status. |
| Shared persistence contract | `skills/_shared/persistence-contract.md` | Downstream artifact resolution and persistence/read-back expectations that need canonical JSON awareness without breaking Markdown compatibility. |
| Shared post-apply gates | `skills/_shared/sdd-post-apply-gates.md` | Review/security/verify/archive consumption boundaries, safe-evidence restrictions, unavailable-tooling handling, and routing semantics. |
| Testing Capabilities | `openspec/config.yaml#testing` | Confirms no executable test runner, coverage, linter, type checker, formatter, or build command is available. |

## Source ID Coverage Baseline

Corporate general-review coverage applies to the 96 `sdd-review` controls, not to the 155-row security source catalog owned by `sdd-review-security`. This test design plans static/manual checks proving that `skills/sdd-review/references/review-control-catalog.json` is the canonical source for the general review catalog and contains exactly 96 unique controls with IDs `REV-CORP-001` through `REV-CORP-096` and source items `1` through `96`.

Validation must use static/manual read-back because the repository has no executable test runner. The planned checks must not require design YAML, compact matrices, security Source ID matrices, machine-readable security applicability fields, all-row `N/A` bookkeeping, or a generated Excel workbook. Full security source-row validation remains owned by `review-security-report.md`; this change only defines and validates the general review JSON/Markdown contract.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Canonical Review JSON Authority / Scenario: JSON owns controls | Verify `skills/sdd-review/references/review-control-catalog.json` exists as the canonical static control catalog and is the catalog loaded by `sdd-review`. | static | mandatory | File read-back plus `skills/sdd-review/SKILL.md` wording stating JSON is authoritative for controls. | Blocks if Markdown remains the loading authority. |
| TD-002 | Design: Static catalog JSON validation rules | Verify the static catalog has `expectedControlCount: 96`, exactly 96 `controls[]` entries, unique IDs, complete `REV-CORP-001..REV-CORP-096`, and complete `sourceItem` values `1..96`. | static | mandatory | Manual/static JSON inspection with counted rows and ID/source-item range evidence. | No executable command is required or available. |
| TD-003 | Design: Static catalog JSON interface | Verify static catalog required metadata fields exist: `schemaName`, `schemaVersion`, `catalogVersion`, `snapshotId`, `status`, `expectedControlCount`, `generatedHumanViewRef`, `vocabulary`, `presentation`, `validation`, and `controls[]`. | static | mandatory | Read-back of `review-control-catalog.json` showing required top-level fields. | Invalid or missing metadata routes to `resolve-blockers`. |
| TD-004 | Spec: Code-Review Validation Matrix | Verify catalog/report vocabulary allows only `Yes`, `No`, and `N/A` for `defaultComplies` and per-report `complies`. | static | mandatory | JSON vocabulary metadata and row inspection evidence. | Any other value blocks report validation. |
| TD-005 | Design: Static catalog JSON interface | Verify each control contains required row fields: `id`, `sourceItem`, `artifactDeliverable`, `requirement`, `reviewer`, `standard`, `severity`, `defaultComplies`, `evidenceHint`, and `notes`; verify optional `category` is allowed only for JSON/Excel grouping metadata. | static | mandatory | Catalog row sample plus field-presence inventory. | Missing required row fields block. |
| TD-006 | Spec: Code-Review Validation Matrix / Design: Markdown generation | Verify optional `category` metadata is not rendered as a Markdown matrix column and that the Markdown header remains exactly `Item`, `Artifact/Deliverable`, `Requirement`, `Reviewer`, `Standard`, `Severity`, `Complies`, `Affected Requirement`, `Evidence Location`, `Observations/Comments`. | static | mandatory | `report-template.md` and generated `review-report.md` contract read-back showing exact header and no Category column. | Preserves downstream compatibility. |
| TD-007 | Spec: Markdown catalog is not authority | Verify `control-catalog.md`, if present, is explicitly marked derived/reference-only and JSON wins on any conflict. | static | mandatory | `control-catalog.md` and `sdd-review/SKILL.md` read-back showing derived/reference-only wording and JSON conflict authority. | Mismatches are repair issues, not control decisions. |
| TD-008 | Spec: Review Report Artifact / Design: Per-change canonical report JSON | Verify `openspec/changes/{change-name}/review-report.json` contract includes schema/report identity fields: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `sourceCatalogRef`, `sourceCatalogSnapshotId`, and `derivedMarkdownRef`. | static | mandatory | `review-report.schema.json` or skill contract read-back plus sample/report shape evidence when generated. | Required for downstream identity and audit. |
| TD-009 | Spec: Review Artifact Authority Boundary | Verify report JSON includes verdict/routing fields: `verdict`, `blockingFailureCount`, `nonBlockingFindingCount`, and `nextRecommended`. | static | mandatory | Schema/contract read-back plus review report example or generated artifact evidence. | Routing mismatch blocks. |
| TD-010 | Proposal: canonical report shape / Design: Per-change report JSON | Verify report JSON includes evidence and handoff sections: `inputsInspected[]`, `runtimeChecks`, `blockingSummary[]`, `evidenceSummary`, `operationalEvidenceSummary`, `changedFileSecurityHandoff`, `validation`, and `presentation`. | static | mandatory | Schema/contract read-back proving field requirements. | Operational sections are required when applicable and safely empty/marked when not applicable. |
| TD-011 | Spec: Reports are persisted / Design: Report validation | Verify per-change `review-report.json.reviewMatrix[]` contains exactly 96 evaluated rows, each mapped to one static catalog control. | static | mandatory | Read-back row count and mapping evidence for generated report JSON. | Missing or extra rows route to `resolve-blockers`. |
| TD-012 | Spec: Platform control is irrelevant | Verify every JSON/Markdown `N/A` row has non-empty `evidenceLocation` proving irrelevance and `observationsComments` explaining the scope decision. | static | mandatory | Generated `review-report.json` and derived `review-report.md` row evidence for all `N/A` rows. | `N/A` without evidence/rationale blocks. |
| TD-013 | Spec: Derived Markdown preserves compatibility sections | Verify generated `review-report.md` preserves exact required sections: Verdict, Blocking Summary, Evidence Summary, Review Matrix, applicable operational evidence, changed-file/security handoff, matrix validation, and recommendation. | static | mandatory | Read-back of generated Markdown sections. | Section names must remain stable for downstream readers. |
| TD-014 | Spec: Code-Review Validation Matrix / Design: Markdown parity validation | Verify generated `review-report.md` has the exact matrix header and exactly 96 rendered data rows, and that key rendered facts compare back to `review-report.json`. | static | mandatory | Markdown read-back, row count, exact header check, and JSON-to-Markdown parity notes. | Stale Markdown must not be treated as current evidence. |
| TD-015 | Spec: Markdown generation fails | Verify Markdown generation/parity failure routes to `resolve-blockers` and does not let downstream phases consume stale Markdown. | static | mandatory | `sdd-review/SKILL.md`, `report-template.md`, and shared contract wording. | Persistence/read-back failure is not a pass. |
| TD-016 | Design: Validation rules / Review routing | Verify blocking or critical `No` rows produce verdict `FAIL` and route to `apply` when remediation is implementation/contract work, or `resolve-blockers` when remediation is artifact/schema/context repair. | static | mandatory | Skill routing rules and report validation contract read-back. | Must not route directly to verify. |
| TD-017 | Review routing / Shared post-apply gates | Verify non-blocking `No` rows are preserved as warnings and may route to `review-security` only when mandatory evidence is complete. | static | mandatory | `sdd-review/SKILL.md`, `report-template.md`, and `sdd-post-apply-gates.md` wording. | Warnings must not hide failed mandatory controls. |
| TD-018 | Spec: Downstream consumes derived Markdown safely | Verify `sdd-review-security`, `sdd-verify`, and `sdd-archive` can continue reading `review-report.md` for compatibility but recognize canonical JSON authority when present. | static | mandatory | Shared persistence/post-apply contract wording plus affected phase skill wording if updated. | Downstream phases should not duplicate the 96-control matrix. |
| TD-019 | Design: Restricted Data Boundary / Secure Development Design: Sensitive Logging Rules | Verify JSON and Markdown evidence rules forbid secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, generated bytes, and final-document-only values. | static | mandatory | Safe-evidence wording in `design.md`, `sdd-review/SKILL.md`, report template/schema, and shared gates. | Unsafe evidence blocks or routes to cleanup. |
| TD-020 | Proposal: Out of Scope / Spec: Excel is deferred / Design: Operational Considerations | Verify this change does not add Excel, Python, script, or workbook generation and does not require spreadsheet output in any contract. | static | mandatory | File list/static review evidence showing no Excel/Python/script/workbook generation contract or artifacts were added. | Explicit negative test; future Excel remains out of scope. |
| TD-021 | Config: Testing capabilities | Verify all unavailable tooling is recorded as unavailable evidence: no test runner, build, coverage, linter, type checker, formatter, or runtime command. | static | mandatory | `openspec/config.yaml#testing` citation and verify/review wording that no invented commands were run. | Missing tooling is a constraint, not passing evidence. |
| TD-022 | Design: review-report.schema.json | Verify `review-report.schema.json` or equivalent contract rejects missing required fields and invalid schema/catalog metadata. | static | mandatory | Schema/contract read-back with required fields and validation metadata expectations. | No runtime schema validator exists, so evidence is static/manual. |
| TD-023 | Design: generated Markdown from JSON | Verify `report-template.md` documents which JSON facts render the Markdown report and that Markdown is generated from `review-report.json`, not hand-authored as authority. | static | mandatory | Template/skill contract read-back. | Prevents JSON/Markdown drift. |
| TD-024 | Design: catalog integrity / Proposal success criteria | Verify `review-control-catalog.json` includes presentation metadata sufficient for Markdown and future presentation use without adding Excel-specific generation. | static | non-mandatory | Catalog `presentation` metadata read-back. | Future Excel-friendly data is allowed; generation is not. |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Artifact integrity | Confirm `sdd-review` writes and reads back both `review-report.json` and generated `review-report.md` before returning success. | static | Skill/shared contract wording and generated artifact read-back evidence during review/verify. | No executable runner exists; use manual/static read-back. |
| Catalog integrity | Confirm canonical catalog JSON count, unique IDs, source items, vocabulary, and snapshot identity are validated before report generation. | static | Catalog JSON read-back and validation metadata evidence. | No automated JSON validator command is available. |
| Markdown compatibility | Confirm generated Markdown keeps exact required sections, exact matrix header, and 96 rows. | static | Markdown read-back and row-count evidence. | No Markdown renderer test command is available. |
| Tooling availability | Confirm unavailable test/build/lint/type/format/coverage tooling is cited explicitly. | documentary | `openspec/config.yaml#testing` plus review/verify artifact notes. | Runtime/build/lint/type/format/coverage unavailable. |
| Safe evidence | Confirm report evidence cites paths, section anchors, sanitized summaries, and unavailable-tooling notes only. | static | Safe-evidence rule read-back in design, skill, template/schema, and generated reports. | No secret scanner is available; manual/static inspection is required. |
| Deferred Excel scope | Confirm no Excel/Python/script/workbook artifacts or generation instructions are added. | static | File list/static review evidence and contract wording. | No script execution expected or allowed for this change. |
| Operational categories not introduced | Confirm no runtime deployment, migration, monitoring, administration, reprocessing, backup, retention, or cleanup evidence is synthesized when design marks them not introduced. | documentary | `design.md#Operational Considerations` citation and `No aplica.`/not-applicable rationale when needed. | No runtime system exists in this repository. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SDD-SAFE-EVIDENCE` | JSON and Markdown review evidence must not include secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, generated bytes, or final-document-only values. | Yes | TD-019 plus Operational Considerations safe-evidence check. | covered | None |
| `SDD-FILE-SCOPE` | The change must create/update repository artifacts only and must not require arbitrary filesystem reads or generated Excel files. | Yes | TD-001, TD-008, TD-020, and TD-023. | covered | None |
| `SDD-EVIDENCE-INTEGRITY` | JSON authority and same-phase Markdown parity validation must prevent stale Markdown from hiding review findings. | Yes | TD-007, TD-011, TD-013, TD-014, and TD-015. | covered | None |
| `SDD-TOOLING-TRANSPARENCY` | Missing runtime/build/lint/type/format/coverage tooling must be reported as unavailable evidence and never as passing evidence. | Yes | TD-021 and Operational Considerations tooling availability check. | covered | None |
| `SDD-EXCEPTION-POLICY` | Any future exception to mandatory safe-evidence, catalog integrity, or report persistence rules must include approver, approval date, accepted-risk rationale, mitigation/follow-up, and exact evidence gap. | Yes | Static review of exception policy wording in design and implementation artifacts. | covered | None |
| `SDD-SECURITY-OWNERSHIP-BOUNDARY` | This change must not transfer security-review source-row ownership from `sdd-review-security` to `sdd-review`. | Yes | Source ID Coverage Baseline plus TD-018 and shared post-apply boundary evidence. | covered | None |

## No-Impact Assessment

Not applicable. This change has direct testability impact because it changes the durable review evidence contract, canonical catalog authority, report validation, Markdown generation/parity behavior, and downstream review evidence consumption.

## Evidence Expectations

- Mandatory cases require implementation evidence, static/manual read-back, or a justified skip; uncovered mandatory cases block downstream verification.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Static/manual verification is expected because `openspec/config.yaml#testing` states no executable test runner, coverage, linter, type checker, formatter, or build command is available.
- Missing runtime/build/lint/type/format/coverage tooling must be reported as unavailable evidence, not treated as passed checks.
- Security validation evidence should cite embedded `design.md#secure-development-design` narrative rules, owner phase, and planned static/manual evidence.
- Evidence must remain safe: cite paths, section anchors, sanitized summaries, and unavailable-tooling statements only; do not include secrets, credentials, PAN, PII, tokens, raw logs, production identifiers, generated bytes, or final-document-only values.
- JSON authority must be explicit: `review-control-catalog.json` owns static control definitions, `review-report.json` owns per-change report facts, and Markdown artifacts are generated/derived compatibility views.
- Exhaustive security source-row validation coverage, all-row `N/A` decisions, and missed-applicable security validation remain owned by `review-security-report.md`.
- Excel/Python/script/workbook generation is explicitly out of scope; any such generated artifact or instruction added by this change is a test failure.

## Open Questions

None.
