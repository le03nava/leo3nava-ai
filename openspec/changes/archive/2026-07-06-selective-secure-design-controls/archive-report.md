# Archive Report: Selective Secure Design Controls

```yaml
schemaName: sdd.archive-report
schemaVersion: 1
changeName: selective-secure-design-controls
artifactStore: openspec
archiveDate: 2026-07-06
verdict: ARCHIVED WITH WARNINGS
nextRecommended: none
archivePath: openspec/changes/archive/2026-07-06-selective-secure-design-controls/
```

## Summary

The OpenSpec change `selective-secure-design-controls` is ready for archive after completed proposal, delta specs, design, test design, tasks, apply progress, general review, security review, and verification evidence. The accepted deltas are already synchronized into the source specs, and this archive preserves the full audit trail under `openspec/changes/archive/2026-07-06-selective-secure-design-controls/`.

The archive is `ARCHIVED WITH WARNINGS` because runtime test runner, build, lint, typecheck, formatter, and coverage tooling are unavailable by repository configuration. That unavailable tooling is preserved as a warning only and is not treated as passing automated evidence.

## Readiness Gates

| Gate | Result | Evidence |
| --- | --- | --- |
| Required artifacts readable | PASS | `proposal.md`, five delta specs, `design.md`, `test-design.md`, `tasks.md`, `apply-progress.md`, `review-report.md`, `review-security-report.md`, `verify-report.md`, and `state.yaml` were read before archive. |
| Task completion | PASS | `tasks.md` records 11/11 revised implementation tasks checked complete and 0 unchecked implementation tasks. |
| General review | PASS WITH WARNINGS, non-blocking | `review-report.md` reports `Blocking failures: 0`, `Non-blocking findings: 1`, and routes to `review-security`. |
| Security review | PASS WITH WARNINGS, non-blocking | `review-security-report.md` reports `blockingFindings: 0`, `nonBlockingFindings: 1`, safe evidence, exact-once source-row coverage, and routes to `verify`. |
| Verification | PASS WITH WARNINGS | `verify-report.md` reports `criticalIssues: 0`, `warningIssues: 1`, `runtimeTooling: unavailable`, and `nextRecommended: archive`. |
| Legacy standalone security artifacts | Not required | New active archive readiness uses `design.md#secure-development-design`, `review-security-report.md`, and `verify-report.md`; standalone `security-design.md`, `security-applicability.md`, and `scripts/validate_security_design.ps1` are legacy/read-only only. |

## Specs Synced

Accepted delta specs were already synchronized into the source OpenSpec specs before archive. Archive re-read the delta and source specs and confirmed the expected requirements exist in source.

| Domain | Action | Details |
| --- | --- | --- |
| `sdd-design-workflow` | Updated | 4 modified requirements synced: `Embedded Secure Development Design`, `Secure Development Design Category Rule Contract`, `Secure Design Source ID Coverage`, and `Design Preserves Compact Summary`. |
| `sdd-test-design-workflow` | Updated | 3 modified requirements synced: `test-design.md Artifact Contract`, `Source Row Test Planning`, and `Source Row N/A and Warning Evidence`. |
| `sdd-review-security-workflow` | Updated | 5 modified requirements synced: `Security Review Artifact`, `Security Matrix Validation`, `Exhaustive Source Row Security Review`, `Source Row Blocking Rules`, and `Source Row Evidence Correlation`. |
| `sdd-security-guideline-catalog` | Updated | 5 modified requirements synced: `Compact Security Taxonomy`, `Catalog Boundary Preservation`, `Corporate Source Row Inventory`, `Safe Source Row Evidence`, and `Shared Security Contract Source Row Schema`. |
| `sdd-execution-persistence-contracts` | Updated | 3 modified requirements synced: `Conflict and Ambiguity Resolution`, `Mandatory Security Artifacts and Status`, and `Source Row Persistence Compatibility`. |

No destructive removals or ambiguous renames were present in the delta specs.

## Source-Row Audit Trail

Corporate source-row validation applies and is preserved by reference and summary rather than copying the full review-security matrix.

| Evidence | Archived Status |
| --- | --- |
| Catalog snapshot identity/path | `skills/_shared/security-guideline-catalog.md`; snapshot `security-guidelines-initial-user-snapshot-2026-06-30`; catalog version `1`; taxonomy version `1` from `verify-report.md#source-row-verdict--warning-consumption`. |
| Expected expanded Source ID count | `155`. |
| Exact Source ID coverage | `review-security-report.md` metadata: `sourceRowExpectedCount: 155`, `sourceRowActualCount: 155`, `sourceRowCoverage: exact-once`; verify consumed the same evidence. |
| Compact mappings | 8 compact controls represented exactly once: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001`. Row-level compact mappings are preserved in `review-security-report.md#corporate-source-row-validation`. |
| Safe evidence | `safeEvidence: pass`, `unsafeEvidenceRows: 0`; evidence uses safe paths, sections, summaries, and report references. |
| `N/A` evidence/justification | `unsupportedNaRows: 0`; omitted runtime categories are justified in report observations and missed-category validation. |
| Missing/duplicate/unknown/malformed blockers | None reported; verify confirms no source-row blockers remain. |
| Exceptions | None; no approved security exceptions required. |
| Non-blocking warnings | Runtime/build/lint/typecheck/formatter/coverage tooling unavailable; warning preserved in review, security review, verify, and this archive report. |
| Review-security matrix link | `review-security-report.md#corporate-source-row-validation` contains the full 155-row matrix and remains the authoritative exhaustive report. |

## Archive Contents

- `proposal.md` ✅
- `specs/` ✅ five delta specs
- `design.md` ✅ includes mandatory `## Secure Development Design`
- `test-design.md` ✅
- `tasks.md` ✅ 11/11 revised tasks complete
- `apply-progress.md` ✅
- `review-report.md` ✅ non-blocking PASS WITH WARNINGS
- `review-security-report.md` ✅ non-blocking PASS WITH WARNINGS and exact-once source-row evidence
- `verify-report.md` ✅ PASS WITH WARNINGS; only warning is unavailable tooling
- `archive-report.md` ✅ this report
- `state.yaml` ✅ final complete state

## Runtime / Quality Tooling Evidence

`openspec/config.yaml#testing` reports no configured runtime test runner, build command, linter, type checker, formatter, or coverage command. The archive preserves this as unavailable tooling evidence. The change relies on static/manual read-back evidence and must not be represented as having passed unavailable runtime tools.

## Final State

The change folder was archived to `openspec/changes/archive/2026-07-06-selective-secure-design-controls/`. The active change folder must be absent after the move, and source specs remain the source of truth for the accepted behavior.

The SDD cycle for `selective-secure-design-controls` is complete. Next recommended phase: `none`.
