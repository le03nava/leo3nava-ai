# Design: SDD Review JSON Contract

## Technical Approach

`sdd-review` will move from Markdown-authored review data to a split JSON contract: a static canonical control catalog under `skills/sdd-review/references/` and a per-change canonical `review-report.json` under `openspec/changes/{change-name}/`. The Markdown report remains mandatory, but it becomes a generated compatibility presentation created from `review-report.json` in the same phase. This maps directly to the delta spec requirements for JSON authority, exact 96-control coverage, unchanged Markdown columns, and downstream compatibility.

## Architecture Decisions

### Decision: Split static catalog JSON from per-change report JSON

**Choice**: Create `skills/sdd-review/references/review-control-catalog.json` for the stable 96 `REV-CORP-*` controls and `openspec/changes/{change-name}/review-report.json` for evaluated per-change facts.
**Alternatives considered**: A single reference JSON containing both catalog and report shape; a JSON block embedded in Markdown; JSON-only output.
**Rationale**: Splitting static definitions from per-change evidence is clearer for maintenance and future Excel consumption. The static catalog changes rarely and can be counted/versioned independently, while each review report records verdict, evidence, routing, and 96 evaluated rows. Embedding JSON in Markdown is fragile for automation, and JSON-only output would break security review, verify, archive, and human audit flows that still read `review-report.md`.

### Decision: JSON wins over Markdown on conflict

**Choice**: `review-report.json` is authoritative for per-change facts, and `review-control-catalog.json` is authoritative for control definitions. `review-report.md` and `control-catalog.md` are derived/reference-only views.
**Alternatives considered**: Keep Markdown catalog as fallback authority or allow either artifact to win by timestamp.
**Rationale**: The proposal explicitly requires JSON as the contract. Timestamp conflict resolution would make review outcomes ambiguous. If Markdown differs, the repair target is Markdown generation, not control semantics.

### Decision: Generate Markdown in-process from canonical JSON

**Choice**: `sdd-review` will first construct and validate `review-report.json`, then render `review-report.md` from that JSON using `references/report-template.md` as the presentation contract.
**Alternatives considered**: Manually fill Markdown first and convert it to JSON, or produce Markdown in a later phase.
**Rationale**: Markdown-first preserves the current drift problem. A later rendering phase creates stale-artifact risk. Same-phase generation makes persistence atomic enough for downstream phases and gives read-back validation a single source of truth.

## Data Flow

```text
review-control-catalog.json ──┐
proposal/spec/design/test/tasks/apply evidence ──→ sdd-review evaluator ──→ review-report.json
changed-file context ─────────┘                                      │
                                                                     └──→ review-report.md
```

`sdd-review` loads the static catalog JSON, validates catalog identity, vocabulary, expected count, and the exact `REV-CORP-001..REV-CORP-096` sequence. It then evaluates each control against the selected change evidence and writes the 96 evaluated rows into `review-report.json.reviewMatrix`. The Markdown renderer reads only `review-report.json` and emits the existing sections and matrix columns.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/sdd-review/SKILL.md` | Modify | Change produced artifact contract to canonical JSON plus derived Markdown; load controls from JSON; validate JSON, row count, unique IDs, vocabulary, Markdown parity, and conflict authority. |
| `skills/sdd-review/references/review-control-catalog.json` | Create | Canonical static catalog with schema metadata, vocabulary, presentation metadata, expected count, and 96 control definitions. |
| `skills/sdd-review/references/review-report.schema.json` | Create | JSON Schema-style contract for per-change `review-report.json` facts, matrix rows, validation metadata, and derived Markdown refs. |
| `skills/sdd-review/references/control-catalog.md` | Modify | Mark as derived/reference-only human view generated from or synchronized to `review-control-catalog.json`; not authoritative. |
| `skills/sdd-review/references/report-template.md` | Modify | Define Markdown rendering from `review-report.json` fields while preserving current sections and matrix header. |
| `openspec/specs/sdd-review-workflow/spec.md` | Modify | Sync base workflow requirements to JSON authority, derived Markdown compatibility, conflict handling, and deferred Excel scope. |
| `skills/_shared/persistence-contract.md` | Modify | Add canonical review JSON artifact refs and state that downstream Markdown remains compatible while JSON wins on conflict. |
| `skills/_shared/sdd-post-apply-gates.md` | Modify | Update post-apply consumption wording so security review, verify, and archive may consume Markdown but recognize canonical JSON authority. |

## Interfaces / Contracts

### Static catalog JSON

Location: `skills/sdd-review/references/review-control-catalog.json`.

Required top-level fields: `schemaName`, `schemaVersion`, `catalogVersion`, `snapshotId`, `status`, `expectedControlCount`, `generatedHumanViewRef`, `vocabulary`, `presentation`, `validation`, and `controls[]`.

Each `controls[]` entry defines: `id`, `sourceItem`, `artifactDeliverable`, `requirement`, `reviewer`, `standard`, `severity`, `defaultComplies`, `evidenceHint`, `notes`, and optional `category` for JSON/Excel grouping only. The Markdown review matrix must not add a Category column.

### Per-change canonical report JSON

Location in OpenSpec mode: `openspec/changes/{change-name}/review-report.json`.

Required top-level fields: `schemaName`, `schemaVersion`, `changeName`, `artifactKind`, `generatedAt`, `sourceCatalogRef`, `sourceCatalogSnapshotId`, `derivedMarkdownRef`, `verdict`, `blockingFailureCount`, `nonBlockingFindingCount`, `nextRecommended`, `inputsInspected[]`, `runtimeChecks`, `blockingSummary[]`, `evidenceSummary`, `operationalEvidenceSummary`, `changedFileSecurityHandoff`, `reviewMatrix[]`, `validation`, and `presentation`.

Each `reviewMatrix[]` row contains the exact facts required to render the Markdown columns: `item`, `artifactDeliverable`, `requirement`, `reviewer`, `standard`, `severity`, `complies`, `affectedRequirement`, `evidenceLocation`, and `observationsComments`. JSON may include extra non-rendered metadata such as `sourceItem`, `category`, or `findingType` for validation and future Excel generation, but the Markdown columns remain unchanged.

### Validation rules

- Static catalog validation must prove exactly 96 controls, unique IDs, the complete `REV-CORP-001..REV-CORP-096` sequence, source items `1..96`, allowed `defaultComplies`, and required presentation metadata.
- Report validation must prove `reviewMatrix.length == 96`, every row maps to one catalog control, `complies` is only `Yes`, `No`, or `N/A`, every `N/A` has evidence and scope rationale, blocking failures match verdict/routing, and `derivedMarkdownRef` points to the Markdown generated in the same phase.
- Markdown parity validation must read back `review-report.md`, confirm required sections, confirm the exact matrix header, confirm 96 rendered rows, and compare key rendered facts back to `review-report.json`.

## Operational Considerations

### Strategy

No runtime deployment, data migration, monitoring, administration, reprocessing, backup, retention, or cleanup behavior is introduced. The operational concern is artifact integrity: `sdd-review` must avoid stale Markdown by writing and reading back both JSON and Markdown before returning success. Future Excel generation is intentionally deferred and must not add Python, spreadsheet output, or generated workbook artifacts in this change.

### Evidence Plan

| Category | Expected safe evidence | Owner phase | Status |
| --- | --- | --- | --- |
| Catalog integrity | Static/manual read-back proving 96 unique JSON controls and expected IDs | apply/review | planned |
| Report integrity | `review-report.json` read-back with schema identity, verdict, routing, and 96 evaluated rows | review | planned |
| Markdown compatibility | `review-report.md` read-back with existing sections and exact matrix columns generated from JSON | review | planned |
| Tooling availability | `openspec/config.yaml#testing` unavailable-runner statement; no invented test command | test-design/verify | planned |

### Restricted Data Boundary

Review evidence must cite paths, section anchors, sanitized summaries, and unavailable-tooling notes only. It must not include secrets, credentials, tokens, connection strings, PAN, PII, raw logs, generated file bytes, production identifiers, or final-document-only values.

### Unresolved Gaps

- None for design. Implementation must still choose whether to hand-maintain the derived Markdown catalog or regenerate it mechanically from the JSON catalog during apply.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static contract | Catalog JSON has 96 unique `REV-CORP-*` controls, source item coverage, allowed vocabulary, and required metadata. | Manual/static read-back and row-count validation; no repository test runner exists. |
| Review report contract | Per-change `review-report.json` contains schema identity, verdict/routing, summaries, validation metadata, and 96 evaluated rows. | Manual/static schema-field inspection and JSON row-count validation evidence in apply/verify. |
| Markdown generation | `review-report.md` is generated from JSON and preserves verdict, blocking summary, evidence summary, review matrix, operational evidence, changed-file/security handoff, matrix validation, and recommendation sections. | Read-back parity check comparing Markdown sections/header/row count to JSON. |
| Downstream compatibility | Security review, verify, and archive can keep reading Markdown while treating JSON as authoritative on conflict. | Static contract review across shared docs and phase skills. |

## Migration / Rollout

No data migration required. Rollout is a contract update: add JSON reference/schema files, update skill/shared docs, and keep `review-report.md` path stable. Existing archived Markdown artifacts remain readable as historical evidence; new reviews must use JSON authority. Rollback is to restore Markdown-only catalog/report behavior while retaining downstream Markdown paths.

## Open Questions

- None blocking.

## Secure Development Design

### Classification and Changed Surface

Classification: security-impacting. The change does not add application runtime code, user authentication, sessions, database access, network APIs, or production file handling, but it changes review evidence contracts that downstream security review, verify, and archive rely on. The touched surface is Markdown/JSON SDD artifacts under `skills/sdd-review/`, shared SDD contracts, and OpenSpec workflow specs. The relevant security context is evidence integrity, safe evidence handling, and preventing Markdown/JSON drift from hiding review findings.

Untouched runtime surfaces include application request handling, credentials, tokens, user data storage, logging pipelines, deployment infrastructure, and generated Excel output. Omitted categories remain reviewable by `sdd-review-security`; this design does not create exhaustive security matrices or per-row `N/A` bookkeeping.

### Sensitive Logging Rules

The implementation must keep review evidence review-safe. JSON and Markdown reports may cite paths, sections, row IDs, sanitized summaries, and unavailable-tooling statements, but must not embed raw secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, generated bytes, or final-document-only values. Evidence owners are `apply` for contract edits, `review-security` for leakage validation, `verify` for preserving warnings, and `archive` for audit retention. Residual risk is that future reviewers may paste unsafe evidence into JSON fields; this remains mitigated by explicit safe-evidence rules and downstream security review.

### Files Rules

The change creates and updates repository artifacts only. JSON schema/catalog/report files must use stable relative paths and must not require reading arbitrary filesystem locations. Future Excel files are out of scope. Evidence owners are `apply` for file creation, `review` for read-back of `review-report.json` and `review-report.md`, and `archive` for preserving both when produced. Residual risk is stale derived Markdown; JSON authority and same-phase parity validation mitigate it.

### Exception and Evidence Policy

No security exceptions are planned. Any future exception to mandatory safe-evidence, catalog integrity, or report persistence rules must record approver, approval date, accepted-risk rationale, mitigation or follow-up, and exact evidence gap before archive readiness. Evidence must remain review-safe in every SDD artifact.
