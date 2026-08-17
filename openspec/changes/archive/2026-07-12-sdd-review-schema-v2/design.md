# Design: sdd-review Schema v2 — Reference Artifact Cleanup

## Technical Approach

Edit-in-place the three reference files in `skills/sdd-review/references/`. No new files, no runtime code, no application logic. The change is a JSON/Markdown artifact restructure that tightens the review contract, removes legacy fields, and adds a `Finding` column to the 96-control matrix. The spec defines the exact v2 field contract; this design specifies the structural transformations and target shapes.

## Architecture Decisions

### Decision: Full Rewrite of review-report.schema.json

**Choice**: Replace the entire schema document rather than incremental edits.
**Alternatives considered**: Patch individual properties and `$defs` in-place.
**Rationale**: 7 root fields are removed, 14 remain, `$defs` changes from 4 definitions to 2, and every object gains `additionalProperties: false`. A full rewrite is cleaner, less error-prone, and easier to review than a diff-heavy patch of the existing 221-line file.

### Decision: Catalog Header Cleanup Without Control Content Changes

**Choice**: Remove root-level legacy fields and per-control `reviewer`/blank-`notes` fields; preserve all control semantic content (id, sourceItem, category, artifactDeliverable, requirement, standard, severity, defaultComplies, evidenceHint, and meaningful notes).
**Alternatives considered**: Also rename `artifactDeliverable` or restructure controls into nested objects.
**Rationale**: Proposal explicitly scopes this to field removal only. Semantic restructuring is a separate change.

### Decision: Template Preserves Pre-Field-Mapping Content

**Choice**: Rewrite from `## JSON Field Mapping` onward; lines 1–8 of the current template remain untouched.
**Alternatives considered**: Full template rewrite.
**Rationale**: Spec requires content before `## JSON Field Mapping` to remain unchanged. Lines 1–8 define the template's purpose and JSON-authority contract.

## Data Flow

```
review-control-catalog.json (v2, static)
         │
         ▼
  sdd-review phase reads catalog + schema
         │
         ▼
review-report.schema.json (v2) ──validates──▶ review-report.json (per-change)
         │
         ▼
report-template.md (v2) ──guides──▶ review-report.md (derived Markdown)
```

No data flow changes; only the shape of the three reference artifacts changes.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/sdd-review/references/review-control-catalog.json` | Modify | Bump schemaVersion to 2; add root `reviewer`; remove 5 root fields and per-control `reviewer`/blank-`notes` |
| `skills/sdd-review/references/review-report.schema.json` | Modify (full rewrite) | 14 required fields; new `$defs` for `ref` and `stateArtifactRef`; `additionalProperties: false` everywhere |
| `skills/sdd-review/references/report-template.md` | Modify (partial rewrite) | Rewrite from `## JSON Field Mapping` onward; new matrix columns with Finding; remove validation sections |

### File 1: review-control-catalog.json — Header Before/After

**Before (v1 header):**
```json
{
  "schemaName": "sdd-review.control-catalog",
  "schemaVersion": 1,
  "catalogVersion": "2026-07-10",
  "snapshotId": "sdd-review-control-catalog-2026-07-10-rev-corp-001-096",
  "status": "canonical",
  "expectedControlCount": 96,
  "generatedHumanViewRef": "skills/sdd-review/references/control-catalog.md",
  "vocabulary": {
    "complies": ["Yes", "No", "N/A"],
    "reviewer": ["sdd-review"],
    "severity": ["Menor", "Media", "Mayor"],
    "safeEvidence": ["paths", "section anchors", "sanitized summaries", "unavailable-tooling statements"],
    "forbiddenEvidence": ["secrets", "credentials", "tokens", ...]
  },
  "presentation": { ... },
  "validation": { ... },
  "controls": [...]
}
```

**After (v2 header):**
```json
{
  "schemaName": "sdd-review.control-catalog",
  "schemaVersion": 2,
  "catalogVersion": "2026-07-10",
  "snapshotId": "sdd-review-control-catalog-2026-07-10-rev-corp-001-096",
  "status": "canonical",
  "expectedControlCount": 96,
  "reviewer": "sdd-review",
  "vocabulary": {
    "complies": ["Yes", "No", "N/A"],
    "severity": ["Menor", "Media", "Mayor"]
  },
  "controls": [...]
}
```

Removed from root: `generatedHumanViewRef`, `presentation`, `validation`.
Removed from `vocabulary`: `reviewer`, `safeEvidence`, `forbiddenEvidence`.
Added to root: `"reviewer": "sdd-review"`.

**Per-control transformation**: For each of the 96 controls, remove the `"reviewer": "sdd-review"` field. Remove the `"notes"` field only when its value is exactly `"Not provided"`. Controls with meaningful `notes` (e.g., REV-CORP-005, REV-CORP-010, REV-CORP-012, REV-CORP-022, REV-CORP-095) retain their `notes` field.

### File 2: review-report.schema.json — Complete v2 Schema

The full rewrite defines these structures:

**Root**: `$schema`, `$id`, `title`, `description`, `$comment` (safe evidence rules), `schemaName` (const `sdd-review.review-report.schema`), `schemaVersion` (const 2), `type: object`, `additionalProperties: false`.

**14 required root fields** (in order):
1. `schemaName` — const `"sdd-review.review-report"`
2. `schemaVersion` — integer, const 2
3. `changeName` — string, minLength 1
4. `artifactKind` — const `"canonical-review-report"`
5. `generatedAt` — string, format date-time
6. `status` — enum `["completed", "blocked"]`
7. `verdict` — enum `["PASS", "PASS WITH WARNINGS", "FAIL"]`
8. `nextRecommended` — enum `["review-security", "apply", "resolve-blockers"]`
9. `totals` — object with `blockingFailureCount` (integer ≥ 0), `nonBlockingFindingCount` (integer ≥ 0); `additionalProperties: false`
10. `catalogRef` — const `"skills/sdd-review/references/review-control-catalog.json"`
11. `catalogSnapshotId` — string, minLength 1
12. `unavailableTooling` — string (free-text note about unavailable runtime tools)
13. `matrix` — array, minItems 96, maxItems 96, items `$ref: #/$defs/matrixRow`
14. `artifactMetadata` — object with `stateRegistration` array (minItems 2, items `$ref: #/$defs/stateArtifactRef`); `additionalProperties: false`

**`$defs`**:

`ref` — object with required `artifact` (string), `ref` (string), `readable` (boolean); optional `notes` (string); `additionalProperties: false`.

`stateArtifactRef` — object with required `artifact` (string, minLength 1), `ref` (string, minLength 1), `kind` (enum `["canonical-json", "derived-markdown"]`); `additionalProperties: false`.

`matrixRow` — object with required fields: `item` (string, pattern `^REV-CORP-(00[1-9]|0[1-8][0-9]|09[0-6])$`), `sourceItem` (integer 1–96), `category` (string), `requirement` (string, minLength 1), `standard` (string, minLength 1), `severity` (string, minLength 1), `complies` (enum `["Yes", "No", "N/A"]`), `finding` (string — free-text, not a controlled enum), `evidenceLocation` (string), `notes` (string); `additionalProperties: false`.

**stateRegistration[2] expected entries**:
- Entry 0: `{ "artifact": "review-report", "ref": "openspec/changes/{change-name}/review-report.json", "kind": "canonical-json" }`
- Entry 1: `{ "artifact": "review-report", "ref": "openspec/changes/{change-name}/review-report.md", "kind": "derived-markdown" }`

### File 3: report-template.md — Replacement Content from `## JSON Field Mapping`

Lines 1–8 remain unchanged. From line 9 onward, the new content:

**`## JSON Field Mapping` table** — maps 14 v2 fields:

| Markdown Section | Source JSON Field(s) |
| --- | --- |
| `# Review Report: {Change Title}` | `changeName` |
| `## Verdict` | `changeName`, `status`, `verdict`, `totals.blockingFailureCount`, `totals.nonBlockingFindingCount`, `nextRecommended` |
| `## Evidence Summary` | `catalogRef`, `catalogSnapshotId`, `unavailableTooling` |
| `## Review Matrix` | `matrix[]` — columns: Item, Requirement, Standard, Category, Severity, Complies, Finding, Evidence Location, Notes |
| `## Recommendation` | `verdict`, `nextRecommended`, totals |
| `## Artifact Metadata` | `artifactMetadata.stateRegistration[]` |

**`## Required Structure` skeleton** — shows the full derived Markdown shape including a 96-row matrix spec with the v2 column sequence: `Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes`.

**`## Matrix Rules`** — updated to reference v2 columns; matrix join rule: `matrix[].item` joined with `catalog controls[].id`.

**Removed sections**: `## Matrix Validation`, `## Catalog Validation` (and the legacy `## Operational Evidence Summary`, `## Changed-File / Security Handoff` sections from the required structure skeleton).

**Preserved verbatim**: `## Safe Evidence Rules` (current lines 108–111):
```
- Evidence may cite paths, section anchors, sanitized summaries, and unavailable-tooling statements only.
- Evidence MUST NOT include secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production IDs/production identifiers, generated bytes, or final-document-only values.
- Missing runtime/build/lint/type-check/format/coverage tooling is recorded as unavailable evidence, never as passing evidence.
- This template does not define or require Excel, Python, script, spreadsheet, or workbook generation.
```

## Interfaces / Contracts

The v2 contract changes the shape consumed by `sdd-review` when generating `review-report.json`. The schema is a reference document; no runtime validator enforces it. The `sdd-review` skill reads these files as instructional context when producing review artifacts.

Key contract change: `changedFileSecurityHandoff` and `operationalEvidenceSummary` are removed from the schema. `sdd-review-security` currently reads these fields — that dependency breaks intentionally and is deferred to a follow-on change.

## Operational Considerations

No aplica. This change modifies static reference JSON/Markdown files consumed by LLM agents during review artifact generation. There are no runtime services, no deployments, no monitoring, no logs, no administration operations, and no backup/retention concerns.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | N/A | No test runner available (`openspec/config.yaml`: `test_runner.available: false`) |
| Integration | N/A | No executable integration tests for reference JSON/Markdown |
| E2E | N/A | No runtime to test |
| Manual/Static | Catalog v2 has 96 controls, correct header, no legacy fields | Static inspection during apply/verify |
| Manual/Static | Schema v2 has 14 required fields, `additionalProperties: false` everywhere, correct `$defs` | Static inspection during apply/verify |
| Manual/Static | Template v2 has correct field mapping, v2 matrix columns, no removed sections, Safe Evidence Rules verbatim | Static inspection during apply/verify |

## Migration / Rollout

No migration required. Reference files are replaced atomically via git commit. Existing archived `review-report.json` files are not retroactively validated against v2. The `sdd-review-security` follow-on change must be applied before running security review against a v2 report.

## Open Questions

None. All design decisions are resolved by the proposal and spec.

## Secure Development Design

### Classification and Changed Surface

**Classification**: No-impact.

**Changed artifacts**: Three static JSON/Markdown reference files in `skills/sdd-review/references/`. These files are consumed by LLM agents (`sdd-review` skill) as instructional context during review artifact generation. They are not executed by application runtime, do not process user input, do not handle authentication/authorization, do not touch databases, do not manage secrets, and do not produce logs.

**Untouched runtime surfaces**: There is no application runtime in this repository. The repository contains AI agent/skill distribution artifacts (Markdown instruction contracts, JSON reference schemas). No auth, session, data, secret, permission, file-processing, database, or logging surfaces exist.

**Why no security category applies**: The changed files define the shape of review reports generated by LLM agents. They do not introduce, modify, or remove any code path that handles sensitive data, credentials, access control, file operations, database queries, or logging. The `additionalProperties: false` constraint in the new schema is a structural tightening that reduces — not increases — the surface for unexpected fields.

**Omitted categories**: All eight security categories (authentication, sessions, sensitive-data-pan, secrets, permissions-access-control, files, database-access, sensitive-logging) are omitted as reviewable omissions for `review-security-report.json` validation.

### Evidence Safety in Changed Artifacts

The new schema retains the `$comment` field with safe evidence rules inherited from v1. The report template preserves `## Safe Evidence Rules` verbatim (lines 108–111). These rules prohibit secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, generated bytes, and final-document-only values in any evidence field.

The `additionalProperties: false` constraint on all schema objects prevents injection of unexpected fields in generated reports, strengthening the evidence boundary.

### Known Dependency Risk

Removing `changedFileSecurityHandoff` and `operationalEvidenceSummary` from the schema breaks the current `sdd-review-security` dependency on those fields. This is an explicit, acknowledged break — not accidental. Mitigation: the proposal registers a follow-on change for `sdd-review-security` update. Until that follow-on is applied, running `sdd-review-security` against a v2 report will fail or route to `resolve-blockers`.

### Exception and Evidence Policy

No exceptions are planned. No sensitive data is handled by these reference files.
