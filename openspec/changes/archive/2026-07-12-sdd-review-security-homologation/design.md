# Design: SDD Review Security Homologation

## Technical Approach

Homologate 4 reference files under `skills/sdd-review-security/references/` to the v3 flat-field pattern established by `sdd-review-schema-v2`. Each file is an independent artifact with no cross-file runtime dependency at edit time. All edits are structural (field renames, removals, schema rewrites, section replacements) — no runtime code, no Python, no test runners. Verification is static inspection only.

## Architecture Decisions

### Decision: Full File Rewrite for Catalog and Schema

**Choice**: Read → transform in-memory → write full file for catalog JSON and schema JSON.
**Alternatives considered**: Surgical patches (find-and-replace individual fields).
**Rationale**: Catalog has 155 rows each needing 4 field changes — surgical patches risk partial application and are harder to verify. Schema changes are too deep (new `required[]`, new `$defs`, removed `$defs`) for patching. Full rewrite guarantees structural consistency.

### Decision: Partial Replacement for Template and Validation Rules

**Choice**: Preserve preamble/valid sections, replace from known boundary onward (template) or targeted section edits (validation rules).
**Alternatives considered**: Full rewrite of both files.
**Rationale**: Template preamble (3 paragraphs) is unchanged per TMPL-001 — preserving it avoids accidental content drift. Validation rules has multiple valid sections that need only path updates (`sourceRowValidation.rows` → `rows[]`), not full rewrites. Targeted edits are safer and more reviewable.

### Decision: All 4 Files Edited During Apply

**Choice**: Apply phase edits all 4 files. No file is skipped.
**Alternatives considered**: Splitting catalog into a separate change.
**Rationale**: The catalog is owned by `sdd-review-security` skill, but this change IS the reference homologation — all 4 files are in scope per proposal. The apply workflow must not skip file 1 thinking it's "owned" elsewhere.

## Data Flow

No runtime data flow. All artifacts are static reference files consumed at review-time by the `sdd-review-security` skill.

```
Catalog JSON ──(join sourceId)──→ Report Template (matrix render)
Schema JSON  ──(validates)──────→ review-security-report.json instances
Validation Rules ──(guides)─────→ sdd-review-security SKILL.md runtime
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/sdd-review-security/references/security-guideline-catalog.operational.json` | Modify | `schemaVersion` 2→3; remove header keys (`generatedHumanViewRef`, `reporting`, `authority`); per-row: remove `guidelineRefs`, `evidenceExpectation`; rename `defaultOwnerPhase`→`ownerPhase`, `defaultRoute`→`route`. 155 rows preserved. |
| `skills/sdd-review-security/references/review-security-report.schema.json` | Modify | Full rewrite: 16 flat required fields; `rows[]` at root (min/max 155); new `$defs`: `totals` (8 fields, `sourceRowCount` const 155), `sourceRow` (10 fields, no catalog fields), `exception` (6 required, no `status`), `stateArtifactRef`; remove `$defs`: `navigationGroup`, `groupedNaSummary`, `finding`; `additionalProperties: false` on all objects. |
| `skills/sdd-review-security/references/report-template.md` | Modify | Preserve 3-paragraph preamble; replace from `## JSON Field Mapping` onward: 9-section field mapping table, 8-section skeleton, join instruction (`rows[].sourceId` with `sourceRows[].sourceId`), 15-column matrix, `## Matrix Rules`, `## Safe Evidence Rules`. Remove 6 obsolete sections. |
| `skills/sdd-review-security/references/validation-rules.md` | Modify | Update `sourceRowValidation.rows` → `rows[]` throughout; required row fields = 10 (no catalog fields); remove Row-Preserving Non-Applicability section; remove `guidelineText`/`guidelineRefs` blocking rule; update checklist to v3 fields (13 entries); remove checklist items for `groupedNaSummaries`, `navigationSummary`. |

## Interfaces / Contracts

No new interfaces. The 4 files define the contract consumed by `sdd-review-security` SKILL.md at review time:

- **Catalog**: `sourceRows[].sourceId` is the join key; each row has `ownerPhase` and `route` (renamed from `default*` prefix)
- **Schema**: validates `review-security-report.json` — `rows[]` at root, `sourceRow` $def with 10 fields only
- **Template**: renders the Markdown report from JSON; matrix joins report `rows[]` with catalog `sourceRows[]` by `sourceId`
- **Validation rules**: guides the SKILL.md review logic; `rows[]` path, 10 required row fields

## Operational Considerations

No aplica. This change touches only static reference artifacts consumed by an LLM skill at design/review time. No runtime services, logs, monitoring, administration, or recovery operations are affected.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static | Catalog v3 structure | `jq` assertions: `schemaVersion == 3`, no banned header/row fields, 155 rows, `ownerPhase`/`route` on every row |
| Static | Schema v3 structure | `jq` assertions: 16 required fields, `$defs` keys, `sourceRow` 10 properties, `exception` 6 required/no `status`, no removed `$defs` |
| Static | Template v3 content | String/grep assertions: preamble preserved, 9 mapping sections, 8 skeleton sections, 15 matrix columns, join instruction present, `## Matrix Rules` and `## Safe Evidence Rules` present, 6 obsolete sections absent |
| Static | Validation rules v3 content | String/grep assertions: `rows[]` not `sourceRowValidation.rows`, 10 required fields listed, Row-Preserving Non-Applicability section absent, `guidelineText`/`guidelineRefs` blocking rule absent, checklist updated |

No unit, integration, or E2E tests — all verification is static file inspection (no runtime tooling).

## Migration / Rollout

No migration required. Archived `review-security-report.json` files from prior changes are immutable and not affected. Only future review-security runs consume the new schema. Rollback: `git revert` or `git checkout HEAD~1 -- skills/sdd-review-security/references/`.

## Open Questions

None.

## Secure Development Design

### Classification and Changed Surface

**Classification: No security impact.**

This change modifies 4 static reference files under `skills/sdd-review-security/references/`. These are LLM skill reference artifacts — not production code, not user-facing, not deployed to any runtime environment. They contain no secrets, credentials, PII, PAN, tokens, or confidential data. They define no authentication, authorization, session, cryptography, file I/O, database, network, or logging surfaces.

**Changed artifacts**: JSON catalog (field renames/removals), JSON Schema (structural rewrite), Markdown template (section replacement), Markdown validation rules (targeted edits).

**Untouched runtime surfaces**: None exist — these files have no runtime. The `sdd-review-security` SKILL.md (the consumer) is explicitly out of scope per proposal.

**Why no security category applies**: No trust boundaries are crossed. No external inputs are processed. No data classes beyond "internal skill reference" are touched. No existing security controls are changed, removed, or newly required. The Safe Evidence Rules section added to the template (TMPL-010) and preserved in validation rules actually strengthens evidence safety guidance — it does not introduce risk.

Omitted categories are reviewable omissions for downstream `review-security-report.json` validation.

### Exception and Evidence Policy

No exceptions are planned. No sensitive evidence is produced or consumed by this change. Safe-evidence rules are self-reinforcing: the template's new `## Safe Evidence Rules` section and the validation rules' preserved `Safe Evidence and Leakage Rules` section document the same prohibitions they enforce.
