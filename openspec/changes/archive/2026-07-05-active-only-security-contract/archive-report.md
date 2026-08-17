# Archive Report: Active-Only Security Contract

```yaml
schemaName: sdd.archive-report
schemaVersion: 1
changeName: active-only-security-contract
artifactStore: openspec
archivedAt: 2026-07-05T21:00:00Z
verdict: PASS WITH WARNINGS
nextRecommended: none
archiveDestination: openspec/changes/archive/2026-07-05-active-only-security-contract/
```

## Summary

The OpenSpec change `active-only-security-contract` was archived after a completed SDD lifecycle. General review, security review, and verification all returned `PASS WITH WARNINGS` with no blockers or CRITICAL issues. The remaining warning is unavailable runtime, build, lint, type-check, formatter, and coverage tooling; this is non-blocking for this Markdown/spec contract change and is recorded as unavailable evidence, not passing automated evidence.

## Included Artifacts

| Artifact | Path | Status |
| --- | --- | --- |
| Proposal | `openspec/changes/archive/2026-07-05-active-only-security-contract/proposal.md` | Preserved |
| Delta specs | `openspec/changes/archive/2026-07-05-active-only-security-contract/specs/` | Preserved |
| Design | `openspec/changes/archive/2026-07-05-active-only-security-contract/design.md` | Preserved |
| Embedded secure design | `openspec/changes/archive/2026-07-05-active-only-security-contract/design.md#secure-development-design` | Preserved and validated |
| Test design | `openspec/changes/archive/2026-07-05-active-only-security-contract/test-design.md` | Preserved |
| Tasks | `openspec/changes/archive/2026-07-05-active-only-security-contract/tasks.md` | 13/13 complete |
| Apply progress | `openspec/changes/archive/2026-07-05-active-only-security-contract/apply-progress.md` | Preserved |
| General review | `openspec/changes/archive/2026-07-05-active-only-security-contract/review-report.md` | PASS WITH WARNINGS; 0 blockers |
| Security review | `openspec/changes/archive/2026-07-05-active-only-security-contract/review-security-report.md` | PASS WITH WARNINGS; no blockers; all 8 SEC rows validated |
| Verification | `openspec/changes/archive/2026-07-05-active-only-security-contract/verify-report.md` | PASS WITH WARNINGS; no CRITICAL issues |
| State | `openspec/changes/archive/2026-07-05-active-only-security-contract/state.yaml` | Updated to complete |
| Archive report | `openspec/changes/archive/2026-07-05-active-only-security-contract/archive-report.md` | Persisted and read back |

## Specs Synced

| Domain | Action | Details |
| --- | --- | --- |
| `sdd-security-guideline-catalog` | Updated | 2 modified requirements synced: `In-Repo Guideline Snapshot`, `Catalog Boundary Preservation`; 0 added, 0 removed, 0 renamed. |
| `sdd-security-applicability-workflow` | Updated | 1 modified requirement synced: `Legacy-Only Applicability Classification`; 2 removed requirements already absent from source spec: `Supported Applicability Overrides`, `Static Applicability Validator`; 0 added, 0 renamed. |
| `sdd-review-security-workflow` | Updated | 2 modified requirements synced: `Security Review Artifact`, `Active Security Validator Retirement`; 0 added, 0 removed, 0 renamed. |
| `sdd-execution-persistence-contracts` | Updated | 2 modified requirements synced: `Conflict and Ambiguity Resolution`, `Mandatory Security Artifacts and Status`; 0 added, 0 removed, 0 renamed. |
| `sdd-design-workflow` | Updated | 2 modified requirements synced: `Embedded Secure Development Design`, `Direct Routing to Test Design`; 0 added, 0 removed, 0 renamed. |

## Source of Truth Updated

The following source specs now reflect the archived behavior:

- `openspec/specs/sdd-security-guideline-catalog/spec.md`
- `openspec/specs/sdd-security-applicability-workflow/spec.md`
- `openspec/specs/sdd-review-security-workflow/spec.md`
- `openspec/specs/sdd-execution-persistence-contracts/spec.md`
- `openspec/specs/sdd-design-workflow/spec.md`

No destructive merge was performed. The delta removals include explicit reason and migration notes, and the corresponding removed requirements were already absent from the source spec at archive time.

## Task Completion

`tasks.md` contains no unchecked implementation tasks. The archived task artifact records all 13 tasks complete: 1.1, 1.2, 1.3, 2.1, 2.2, 2.3, 2.4, 3.1, 3.2, 3.3, 3.4, 3.5, and 3.6.

## Review and Verification Evidence

| Gate | Verdict | Blocking State | Archive Decision |
| --- | --- | --- | --- |
| General review | PASS WITH WARNINGS | 0 blockers | Non-blocking; archived |
| Security review | PASS WITH WARNINGS | No blockers; 0 approved exceptions | Non-blocking; archived |
| Verification | PASS WITH WARNINGS | No CRITICAL issues | Non-blocking; archived |

## Secure Development Evidence

`design.md#secure-development-design` is preserved as the active security authority. Security review validated all 8 catalog rows:

| Guideline ID | Category | Evidence Status | Exception State |
| --- | --- | --- | --- |
| `SEC-AUTH-001` | authentication | N/A with rationale; no auth-flow changes | None |
| `SEC-SESS-001` | sessions | N/A with rationale; no session/token/cookie changes | None |
| `SEC-DATA-001` | sensitive-data-pan | Implemented; review-safe evidence only | None |
| `SEC-SECRET-001` | secrets | Implemented; no raw secrets/tokens/private keys in evidence | None |
| `SEC-ACCESS-001` | permissions-access-control | Implemented; historical tokens are read-only and non-launchable | None |
| `SEC-FILE-001` | files | N/A with rationale; no runtime file-handling changes | None |
| `SEC-DB-001` | database-access | N/A with rationale; no database/query/migration changes | None |
| `SEC-LOG-001` | sensitive-logging | Implemented; sanitized audit evidence preserved | None |

Catalog snapshot identity and lifecycle vocabulary are preserved by `skills/_shared/security-guideline-catalog.md` and validated through `review-security-report.md`. No mandatory security evidence gap or incomplete exception was archived.

## Tooling Availability

| Capability | Status | Evidence |
| --- | --- | --- |
| Build | Unavailable | `openspec/config.yaml rules.verify.build_command` is empty. |
| Runtime tests | Unavailable | `openspec/config.yaml#testing.test_runner.available: false`. |
| Coverage | Unavailable | `openspec/config.yaml#testing.coverage.available: false`. |
| Lint | Unavailable | `openspec/config.yaml#testing.quality.linter.available: false`. |
| Type-check | Unavailable | `openspec/config.yaml#testing.quality.type_checker.available: false`. |
| Formatter | Unavailable | `openspec/config.yaml#testing.quality.formatter.available: false`. |

Unavailable tooling is preserved as a warning and was not treated as passing evidence.

## Archive Verification

- Source specs were checked against the delta specs and synced before archive movement.
- The archive destination did not exist before movement.
- The archive destination is inside the allowed edit root `openspec/changes/archive`.
- The active change folder was moved to `openspec/changes/archive/2026-07-05-active-only-security-contract/`.
- The archived folder contains proposal, specs, design, test-design, tasks, apply-progress, review, security review, verification, state, and archive report artifacts.
- The archived design contains `## Secure Development Design` and all 8 SEC IDs.
- The active change folder no longer exists after movement.

## Final Status

The SDD cycle is complete. The change is fully planned, implemented, reviewed, security-reviewed, verified, synced to source specs, and archived.
