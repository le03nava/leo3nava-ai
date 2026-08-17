# Design: Source Row First Security Review

## Technical Approach

This change is a breaking SDD contract redesign: `review-security-report.json` becomes source-row-first, with the 155 corporate Source ID rows as the only active security validation matrix. Compact `SEC-*` controls are removed from the active report contract entirely. They may remain only in historical archived artifacts or transient migration context during implementation review, but they must not appear in active validation, navigation, summaries, metadata, or grouped `N/A` authority.

The implementation will update the `sdd-review-security` skill, its schema/template/catalog references, shared downstream contracts, and the Python Excel exporter. The security-review phase will build canonical JSON first, validate exact-once source-row coverage, then generate lean Markdown from that JSON. Verify and archive will consume canonical JSON summaries, counts, warnings, exceptions, and evidence refs rather than copied matrices.

## Architecture Decisions

### Decision: Source rows are the only active security matrix

**Choice**: Make `sourceRowValidation.rows` mandatory with exactly 155 unique rows in every active `review-security-report.json`.
**Alternatives considered**: Keep compact rows as a parallel validation layer, or keep source rows optional/audit-only.
**Rationale**: Parallel compact/source validation creates two authorities and lets grouped compact `N/A` hide row-level gaps. Exact-once Source ID coverage gives one reviewable source of truth.

### Decision: Compact `SEC-*` is removed from active report authority

**Choice**: Remove `compactControlValidation`, compact mappings, compact summary sections, and compact navigation from active report schema/template/validation.
**Alternatives considered**: Keep compact IDs for summaries or metadata only.
**Rationale**: The proposal requires no compact active authority. Keeping compact IDs in summaries or metadata would still influence navigation and downstream interpretation.

### Decision: Grouping uses source-row category fields

**Choice**: Use `controlDomain`, `corporateSection`, `repoProfiles`, `runtimeSurface`, `dataSurface`, or another source-row field for navigation and `N/A` grouping.
**Alternatives considered**: Group by compact control or render only an ungrouped full table.
**Rationale**: Source-row fields preserve row authority while still giving humans concise navigation.

### Decision: Markdown is generated presentation only

**Choice**: Keep `report-template.md` as a lean JSON-field mapping and presentation contract with the full source-row matrix rendered last.
**Alternatives considered**: Duplicate validation rules in Markdown sections.
**Rationale**: Canonical JSON must own validation; Markdown is for human compatibility and must not drift into a second source of truth.

## Data Flow

```text
security-guideline-catalog.operational.json
        │
        ▼
catalog transformer / source-row vocabulary
        │
        ▼
sdd-review-security validation
        │
        ├── reads design/test-design/tasks/apply/review JSON evidence
        ├── builds sourceRowValidation.rows[155]
        ├── validates exact-once Source IDs and safe evidence
        ▼
review-security-report.json (canonical)
        │
        ├── generated review-security-report.md (derived)
        ├── verify consumes summaries/counts/warnings/exceptions/evidence refs
        ├── archive preserves canonical refs and summaries
        └── Excel exporter defaults to sourceRowValidation.rows
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/sdd-review-security/SKILL.md` | Modify | Rewrite phase authority, execution steps, and output contract around source-row-first JSON and no active compact data. |
| `skills/sdd-review-security/references/review-security-report.schema.json` | Modify | Remove `compactControlValidation`; require `sourceRowValidation.rows` with exactly 155 source rows and source-row grouping fields. |
| `skills/sdd-review-security/references/validation-rules.md` | Modify | Replace compact validation rules with exact-once Source ID validation, row field validation, safe evidence, and unsupported `N/A` blockers. |
| `skills/sdd-review-security/references/report-template.md` | Modify | Simplify generated Markdown: verdict, refs, handoff, source-row navigation/summaries, grouped `N/A`, blockers/warnings, metadata, full matrix last. |
| `skills/sdd-review-security/references/security-guideline-catalog.operational.json` | Modify | Transform catalog metadata to source-row-first fields: `corporateSection`, `controlDomain`, profiles/surfaces, evidence expectations, owner/route vocabulary. |
| `skills/sdd-review-security/references/security-guideline-catalog.md` | Modify | Align human/audit catalog view with source-row-first vocabulary and remove compact-control authority language. |
| `skills/_shared/sdd-security-contract.md` | Modify | Update phase boundaries so canonical security JSON owns source-row validation only; verify/archive consume summaries and refs. |
| `skills/_shared/persistence-contract.md` | Modify | Update security-review artifact verification/read-back requirements to require `sourceRowValidation.rows` exact-once coverage and no compact authority. |
| `skills/_shared/sdd-post-apply-gates.md` | Modify | Ensure post-apply gates route on canonical source-row blockers/warnings rather than compact validation. |
| `skills/sdd-verify/SKILL.md` | Modify | Consume canonical source-row security review summaries/counts/warnings/exceptions/evidence refs without revalidating or copying the matrix. |
| `skills/sdd-archive/SKILL.md` | Modify | Preserve canonical JSON/derived Markdown refs and source-row summary evidence; do not preserve compact validation as active evidence. |
| `agents/sdd/sdd-review-security.md` | Modify | Adapter prompt mirror of skill contract. |
| `agents/sdd/sdd-verify.md` | Modify | Adapter prompt mirror of verify consumption contract. |
| `agents/sdd/sdd-archive.md` | Modify | Adapter prompt mirror of archive preservation contract. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modify | Sync accepted delta after archive. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modify | Sync accepted delta after archive. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modify | Sync accepted delta after archive. |
| `openspec/specs/canonical-review-json-excel-exporter/spec.md` | Modify | Sync accepted delta after archive. |
| `python/json_report_to_excel.py` | Modify | Change security-review default table to `sourceRowValidation.rows`; fail compact-only reports without `--table`. |
| `python/tests/test_json_report_to_excel.py` | Modify | Add source-row default tests, compact-only default failure, nested override preservation, workbook read-back. |
| `python/README.md` | Modify | Document source-row security-review export default and remove compact active-default language. |

## Interfaces / Contracts

### Canonical JSON shape

The active JSON contract keeps top-level review metadata, source refs, handoff, findings, exceptions, unavailable tooling, and artifact metadata. The active security matrix is only:

```json
{
  "schemaName": "sdd-review-security.review-security-report",
  "schemaVersion": 2,
  "changeName": "source-row-first-security-review",
  "artifactKind": "canonical-security-review-report",
  "status": "success|blocked|partial",
  "verdict": "PASS|PASS WITH WARNINGS|FAIL",
  "nextRecommended": "verify|apply|resolve-blockers",
  "sourceRefs": {},
  "catalogRefs": {
    "operationalJson": "skills/sdd-review-security/references/security-guideline-catalog.operational.json",
    "humanMarkdown": "skills/sdd-review-security/references/security-guideline-catalog.md",
    "snapshotId": "...",
    "expectedSourceRowCount": 155
  },
  "generalReviewHandoff": {},
  "sourceRowValidation": {
    "expectedCount": 155,
    "validatedCount": 155,
    "coverageStatus": "complete|incomplete",
    "exactOnce": true,
    "groupingFields": ["controlDomain", "corporateSection"],
    "rows": []
  },
  "blockers": [],
  "warnings": [],
  "unsafeEvidenceRejections": [],
  "warningCarryForward": [],
  "exceptions": [],
  "unavailableTooling": [],
  "artifactMetadata": {}
}
```

### Source row fields

Each row must include: `sourceId`, `corporateSection`, `pciAlignment`, either `guidelineText` or `guidelineRefs`, `controlDomain`, `repoProfiles`, `runtimeSurface`, `dataSurface`, `appliesWhen`, `applies`, `complies`, `lifecycleStatus`, `evidenceType`, `evidenceLocation`, `justification`, `finding`, `ownerPhase`, and `route`.

### Validation invariants

- `sourceRowValidation.rows` is required, has exactly 155 entries, and is the only active security matrix.
- The set of row `sourceId` values must equal the expanded catalog Source ID set exactly once: no missing, duplicate, or unknown IDs.
- `sourceRowValidation.expectedCount` and `validatedCount` must both be 155 for non-blocking reports.
- Compact `SEC-*` data must be absent from active report properties, validation summaries, navigation metadata, grouped `N/A` authority, and derived Markdown sections.
- Every `N/A` row retains its own `applies`, `complies`, `justification`, `evidenceType`, `evidenceLocation`, `finding`, `ownerPhase`, and `route` even when Markdown groups equivalent `N/A` rows.
- Evidence must be review-safe; unsafe evidence is a blocker and must be represented through `unsafeEvidenceRejections` plus row findings.
- Derived Markdown parity checks compare generated content to canonical JSON for verdict, routing, counts, blocker/warning totals, source refs, and artifact metadata.

## Catalog Transformation Approach

The current operational catalog already preserves snapshot identity, expected count, taxonomy, source sections, and concrete source rows. Implementation should transform it as follows:

1. Promote source-row fields to first-class automation fields: `sourceId`, `corporateSection` from current `section`, `pciAlignment`, source text/ref, and grouping fields.
2. Add/derive `controlDomain` from taxonomy/category intent and corporate section semantics. For example, authentication/password rows group under identity/authentication domains; logging/error rows group under observability/logging domains; file rows group under file-handling domains.
3. Add `repoProfiles`, `runtimeSurface`, and `dataSurface` vocabularies for source-row applicability decisions in this repository type. Since this repository is an AI agent/skill distribution, likely profiles include `sdd-contracts`, `agent-prompts`, `python-exporter`, and `documentation`.
4. Preserve corporate guideline text or stable refs to the human catalog. Long text can stay in the catalog while report rows cite `guidelineRefs` when needed.
5. Remove compact mappings from active report generation. If retained temporarily in the catalog during implementation for reviewer migration context, they must not be read by active report validation or rendered as authority.
6. Add a catalog self-check that expands any remaining ranges before validation and asserts exactly 155 concrete Source IDs.

## Markdown Template Simplification Strategy

`report-template.md` should become a lean presentation mapping:

- Verdict and routing.
- Source refs and catalog snapshot identity.
- General review handoff summary without duplicating the 96-control matrix.
- Source-row navigation by `controlDomain`, `corporateSection`, or another source-row grouping field.
- Count summaries for expected/validated rows, blockers, warnings, exceptions, unsafe evidence, and `N/A` groups.
- Grouped `N/A` summaries only when row-level JSON preserves each row decision and justification.
- Blockers, warnings, unsafe evidence rejections, warning carry-forward, exceptions, unavailable tooling, and artifact metadata.
- Full 155-row matrix at the end, generated from `sourceRowValidation.rows`.

The template must delete `## Compact Control Validation` and must not define independent validation logic. It should describe rendering from JSON and parity/read-back expectations only.

## Excel Exporter Impact

`python/json_report_to_excel.py` should change `DEFAULT_TABLE_BY_SCHEMA["sdd-review-security.review-security-report"]` from `compactControlValidation.rows` to `sourceRowValidation.rows`. Existing nested dotted-path support remains valid. Without `--table`, compact-only security-review JSON must fail because the default path is missing. Manual `--table compactControlValidation.rows` may continue as a generic exporter override for historical files, but README must not present that as the active contract.

Tests should cover source-row default selection, compact-only default failure, manual nested path override, workbook sheet naming, row flattening, workbook read-back through `openpyxl`, and README contract text.

## Downstream Verify / Archive Consumption

Verify and archive should resolve canonical `review-security-report.json` as the security authority and read derived Markdown only for human compatibility/parity evidence. They consume:

- verdict/status/next route;
- expected/validated source-row counts and exact-once status;
- catalog snapshot identity/path;
- blocker and warning summaries;
- unsafe evidence rejections;
- exception records;
- warning carry-forward;
- evidence refs and artifact parity/read-back metadata.

They must not copy the full matrix into verify/archive reports, recalculate row compliance, or accept compact `SEC-*` data as active validation authority.

## Operational Considerations

### Strategy

This is a repository contract/tooling change, not an application runtime change. Operational behavior is limited to artifact authority, report generation, validation routing, and exported workbook generation. No production hostnames, ports, credentials, dashboards, logs, payloads, or environment-specific values are required.

### Evidence Plan

| Category | Expected safe evidence | Owner phase | Status |
| --- | --- | --- | --- |
| Artifact authority | Read-back refs for JSON and derived Markdown; schema/template paths. | review-security / verify | planned |
| Exporter behavior | Pytest results or unavailable-tooling note plus static file inspection. | apply / verify | planned |
| Contract propagation | Changed-file refs for skills, agents, shared contracts, specs, README. | apply / review | planned |

### Restricted Data Boundary

Ordinary SDD evidence must use paths, section anchors, sanitized summaries, command summaries, and redacted placeholders only. No restricted operational data or final-document-only values are needed or allowed.

### Unresolved Gaps

- None blocking at design time.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static contract inspection | Schema/template/skill/shared-contract removal of active compact authority and source-row-first requirements. | Manual/static review because repository init reports no executable global test runner. |
| Python unit tests | Exporter default path, compact-only default failure, manual override, workbook read-back, README assertions. | Run `python -m pytest python/tests` when dependencies are available. |
| Integration-like artifact checks | Generated security-review contract consistency across schema, validation rules, template, verify/archive contracts, and adapter prompts. | Static grep/read-back checks and review-security phase evidence. |
| Security evidence checks | Exact-once 155-row invariants, row-level `N/A`, safe evidence rejection, warning carry-forward, exceptions. | Static schema/rules inspection plus generated report sample if implementation creates one. |

## Migration / Rollout

No data migration is required and no historical compatibility is required for active new reports. Roll out as one coordinated contract update across security-review skill references, shared contracts, adapter prompts, OpenSpec base specs, and Python exporter docs/tests. Existing archived reports remain historical records but must not satisfy new active validation.

Rollback is to revert this OpenSpec change before archive, or revert the coordinated implementation files together after implementation. Partial rollback is unsafe because schema, template, validator, downstream contracts, and exporter defaults must agree.

## Open Questions

- None blocking.

## Secure Development Design

### Classification and Changed Surface

Classification: security-impacting.

The changed surface is the SDD security evidence pipeline itself: security-review report schema, validation rules, catalog transformation, generated Markdown, shared verify/archive contracts, adapter prompts, and the Excel exporter. The change does not touch an application runtime, authentication flow, session mechanism, database connection, production file upload/download path, or secret storage runtime. It does change how security evidence is validated, persisted, exported, and consumed.

Catalog context considered: the current operational catalog contains the 155-row corporate Source ID snapshot, source sections, PCI alignment, taxonomy, vocabulary, and safe-evidence expectations. This design cites that context only at category level and leaves exact Source ID expansion, row decisions, and `N/A` validation to canonical `review-security-report.json`.

Applicable narrative categories are sensitive data/evidence handling, files/exported artifacts, permissions/access-control boundaries for artifact authority, sensitive logging/evidence leakage, and database/input-style validation as schema/data validation. Authentication, sessions, runtime secrets, and production database access are omitted from design-time rules because the changed files are repository contracts and exporter tooling, not application auth/session/database runtime surfaces. Omitted categories remain reviewable by `sdd-review-security`; they are not design-time `N/A` rows.

### Sensitive Data and Evidence Handling Rules

Source-row evidence must be review-safe. Report rows, Markdown summaries, verify/archive summaries, exporter README examples, and tests must cite paths, section anchors, sanitized summaries, command outcomes, or redacted placeholders only. They must not include secrets, credentials, tokens, private keys, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, full exported contents, or generated bytes.

Evidence owners: `apply` provides changed-file refs and implementation/static proof; `review-security` validates row evidence and unsafe evidence rejections; `verify` preserves warnings/exceptions/evidence refs without revalidating the matrix; `archive` preserves canonical refs and summaries.

Residual risk: long corporate guideline text may include business-sensitive policy detail. Mitigation is to keep source text in the catalog and allow report rows to cite stable refs where full text is not necessary. Archive expectation: canonical report refs, counts, warnings, exceptions, and evidence refs remain readable without leaking restricted data.

### Files and Exported Artifact Rules

The Python Excel exporter may generate `.xlsx` workbooks from canonical JSON. The active default must export `sourceRowValidation.rows`; compact-only reports fail without explicit manual table selection. Tests must avoid committing generated workbook outputs as ordinary artifacts unless explicitly intended; workbook validation should use temporary paths and read-back checks.

Evidence owners: `apply` updates exporter code/tests/docs; `verify` cites test results or unavailable-tooling/static inspection; `archive` preserves report/exporter contract evidence, not generated workbook bytes.

Residual risk: exported workbooks can contain whatever evidence the JSON includes. Mitigation is upstream safe-evidence validation before export and README guidance that the exporter does not sanitize unsafe source JSON.

### Artifact Authority and Access-Control Boundary Rules

Canonical JSON is the only active security authority. Derived Markdown and Excel output are compatibility/presentation artifacts and must never override JSON. Verify and archive must consume canonical JSON summaries and refs; they must not copy, rescore, or repair the row matrix.

Evidence owners: `review-security` records artifact parity/read-back metadata; `verify` checks canonical refs, non-blocking verdict, row counts, blockers/warnings/exceptions, and parity metadata; `archive` preserves canonical JSON before derived Markdown in refs.

Residual risk: humans may still read Markdown first. Mitigation is explicit template wording, artifact metadata, and downstream contract language that JSON wins on disagreement and stale/parity-failed Markdown routes to `resolve-blockers`.

### Sensitive Logging and Error Evidence Rules

Validation errors, exporter errors, and review findings must be clear but safe. They may name missing paths such as `sourceRowValidation.rows`, missing counts, duplicate/unknown Source IDs, and schema fields, but must not echo unsafe evidence payloads from report rows.

Evidence owners: `apply` implements safe error text; `review-security` records unsafe evidence rejections; `verify` cites sanitized command summaries and unavailable tooling notes.

Residual risk: failing JSON may contain unsafe strings. Mitigation is to summarize failure locations and reject unsafe evidence rather than copying raw row values into logs or reports.

### Schema and Input Validation Rules

The schema and validation rules must enforce exact-once Source ID coverage, required row fields, allowed vocabularies, safe `N/A` justification, finding/route consistency, and absence of compact active authority. Missing, duplicate, unknown, malformed, unsafe, unsupported, or unsupported `N/A` rows block downstream phases.

Evidence owners: `apply` updates schema/rules/template; `review-security` validates generated reports; `verify` confirms non-blocking canonical JSON and preserves warnings/exceptions; `archive` stores refs and summaries.

Residual risk: JSON Schema alone may not express set equality against the catalog. Mitigation is a validator/static review rule that compares row IDs to the expanded catalog inventory before reporting success.

### Exception and Evidence Policy

No exceptions are planned. Any future exception must include approver, approved date, accepted-risk rationale, mitigation or follow-up, and the exact evidence gap. Incomplete exceptions block verification and archive readiness.

Safe evidence policy applies to every phase artifact: cite file paths, section anchors, sanitized summaries, command summaries, redacted placeholders, and exact safe markers only. Do not include raw secrets, credentials, tokens, private keys, connection strings, PAN, PII, confidential values, raw logs/payloads, production infrastructure identifiers, full ID lists, generated file bytes, or final-document-only values.
