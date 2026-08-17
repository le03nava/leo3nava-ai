# Design: Harden Security Applicability Contract

## Technical Approach

Harden the SDD security applicability contract by making the existing Markdown/YAML artifacts auditable without changing phase identity or routing. The implementation will update the shared schema, applicability executor template, compact guideline catalog, security-design consumer contract, and add a static repository-local validator for `security-applicability.md` structure and references.

The repository has no configured runtime test runner in `openspec/config.yaml`, so validation evidence should be static/manual unless a later phase adds executable checks.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Schema evolution | Extend `schemaVersion: 1` fields in-place with catalog metadata, category decision rows, source refs, override metadata, and validator metadata. | Create a new artifact name or require `security-design.md` for all changes. | Preserves archived artifact identity, existing no-impact routing, and downstream resolver paths. |
| Validator location | Create `scripts/validate_security_applicability.ps1`. | Embed validation only in skill prose, or add a package-managed test runner. | Existing support scripts are PowerShell-only and the repo has no package manifest; a focused script is reviewable and runnable on the current Windows/OpenCode workflow. |
| Source mapping | Store formal `SEC-* -> Source IDs` coverage in `skills/_shared/security-guideline-catalog.md` compact rows/metadata, preserving the full corporate snapshot unchanged. | Duplicate full source rows into every artifact. | Keeps artifacts compact while retaining stable audit references to the snapshot. |
| Severity model | Use only `blocking`, `conditional`, and `advisory` for applicability operations. | Reuse review labels such as `Menor`, `Media`, `Mayor`. | Applicability severity controls phase blocking; review severity labels are a separate review concern. |
| Compatibility | Enriched fields are additive; no-impact still skips `sdd-security-design` when explicit proof and validator metadata are valid. | Make enriched fields security-impacting by default. | Specs require compatibility preservation for valid no-impact changes. |

## Data Flow

```text
proposal/specs/config
  -> sdd-security-applicability
  -> security-applicability.md
  -> scripts/validate_security_applicability.ps1
  -> design/security-design/test-design/tasks/verify/archive consumers
```

`sdd-security-applicability` reads the catalog snapshot and optional `rules.security-applicability` config. It writes an artifact containing catalog identity, every taxonomy category decision, applicable guideline/source refs, override metadata, no-impact proof when relevant, and validation metadata. The validator checks the artifact against the same catalog snapshot before phase success.

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/sdd-security-applicability/SKILL.md` | Modify | Add matrix template, validator requirement, safe override handling, no-impact proof rules, and blocking behavior for invalid artifacts. |
| `skills/_shared/sdd-security-contract.md` | Modify | Define additive applicability fields, decision statuses, operational severity, override metadata, validator metadata, and compatibility rules. |
| `skills/_shared/security-guideline-catalog.md` | Modify | Add catalog snapshot/version metadata, formal source coverage per compact `SEC-*`, operational severity, and validator contract fields. |
| `skills/sdd-security-design/SKILL.md` | Modify | Consume enriched fields and preserve catalog/source references in controls while keeping no-impact skip behavior. |
| `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` | Modify | Add compatibility notes for enriched applicability, mandatory security-design only when `securityImpact: true`, and validation metadata consumption. |
| `scripts/validate_security_applicability.ps1` | Create | Static validator for required fields, routing consistency, matrix completeness, guideline/source validity, override safety, severity vocabulary, and no-impact proof. |

## Interfaces / Contracts

Additive artifact fields should include:

```yaml
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  taxonomyVersion: 1
categoryDecisionMatrix:
  - category: authentication
    decision: applicable | not-applicable | unknown
    severity: blocking | conditional | advisory
    rationale: <text>
    evidenceRefs: []
    guidelineIds: []
    sourceIds: []
overridesApplied: []
validation:
  validator: scripts/validate_security_applicability.ps1
  status: pass | fail
  checkedAt: <iso-8601-or-manual>
```

Supported `rules.security-applicability` overrides: extra prompts, stricter source coverage, validator mode, and stricter category severity. Unsafe weakening must be rejected or ignored in favor of the base contract.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static | Validator accepts complete impacting/no-impact artifacts and rejects missing categories, bad source IDs, unsafe overrides, invalid severity, and routing mismatch. | PowerShell script examples or manual static runs; no configured test runner exists. |
| Manual | Skill contracts preserve phase routing and downstream compatibility. | Inspect affected Markdown contracts against specs. |
| Integration | Downstream phases consume enriched fields without requiring `security-design.md` for valid no-impact artifacts. | Static artifact walkthrough in `test-design.md` and verify evidence. |

## Migration / Rollout

No data migration required. Existing archived artifacts remain readable because artifact names, paths, `schemaName`, `schemaVersion`, `classification`, `securityImpact`, `SEC-*` IDs, and no-impact routing are preserved. New validation applies to newly produced or intentionally revalidated artifacts.

## Review Slicing

Use stacked-to-main slices near the 400-line review budget: PR 1 updates shared schema/applicability template/spec compatibility, PR 2 updates catalog source coverage/severity metadata, PR 3 adds validator and downstream compatibility notes.

## Open Questions

None.

## Security Applicability Routing

`security-applicability.md` classifies this change as `security-impacting` with `securityImpact: true`, covering all supported categories and `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001`. Next route: `security-design`.
