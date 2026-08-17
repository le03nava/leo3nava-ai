# Design: Inline Secure Development Design in `design.md`

## Technical Approach

Move secure development design from a standalone new-change phase into the existing `design.md` contract. The active DAG becomes `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`; `sdd-review-security` remains the post-review evidence gate and validates embedded rows against `skills/_shared/security-guideline-catalog.md`. Legacy `security-design.md` remains readable only for old/archive compatibility.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Inline security authority | `sdd-design` writes `## Secure Development Design` with all 8 SEC rows. | Keep `sdd-security-design` active. | One design artifact prevents duplicated classification and matches the target specs. |
| Routing | Successful new-change design recommends `test-design`. | Route to `security-design`. | Specs require removing the active phase while preserving security obligations. |
| Validation | `sdd-review-security` validates embedded rows and writes a matrix. | Require `scripts/validate_security_design.ps1`. | The script targets standalone `security-design.md`; active validation should use catalog + artifact evidence. |
| Compatibility | Keep legacy refs/readers for archived `security-design.md`. | Delete all mentions immediately. | Archives need audit readability; only active dependencies change. |

## Data Flow

```text
proposal/specs/catalog
        -> sdd-design: design.md + embedded SEC rows
        -> sdd-test-design: scenario + SEC evidence plan
        -> tasks/apply
        -> sdd-review -> sdd-review-security: row validation matrix
        -> verify -> archive
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `README.md` | Modify | Update phase list and DAG. |
| `agents/sdd/sdd-orchestrator.md` | Modify | Remove active `security-design` successor/dependencies; preserve legacy display. |
| `agents/sdd/sdd-design.md`, `skills/sdd-design/SKILL.md` | Modify | Require embedded secure design and route to `test-design`. |
| `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` | Modify | Consume security obligations from `design.md`. |
| `skills/sdd-review-security/SKILL.md` | Modify | Validate embedded rows and produce `review-security-report.md`. |
| `skills/_shared/*.md`, `openspec/specs/*` | Modify | Sync persistence/status/security contracts and source specs. |
| `scripts/validate_security_design.ps1` | Delete or mark legacy-only | Remove from active new-change obligations; retain only if implementation proves archive-only use is needed. |

## Interfaces / Contracts

`design.md` MUST include `## Secure Development Design` with catalog snapshot `security-guidelines-initial-user-snapshot-2026-06-30`, taxonomy version `1`, `securityImpact`, all 8 guideline IDs exactly once, applies value (`Yes`/`N/A`), rationale, design decision/control, evidence owner, expected evidence, lifecycle status, residual risk, and exception details when needed. New-change status MUST NOT require `artifactRefs.securityDesign` as an active dependency.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static/manual | DAG, status refs, artifact contracts, legacy wording | Inspect changed Markdown/contracts; no runner is configured. |
| Contract | Embedded 8-row section and downstream consumers | `test-design.md` should plan grep/static checks for IDs, routing tokens, and retired validator references. |
| Regression | Archive compatibility | Verify old `security-design.md` refs remain read-only context. |

## Migration / Rollout

No data migration required. Roll out as stacked-to-main slices: routing/contracts first, phase skill consumers second, validator removal/docs/spec sync last.

## Open Questions

None.

## Secure Development Design

**Classification**: `security-impacting` for SDD workflow governance, not application runtime security.  
**Catalog**: `security-guidelines-initial-user-snapshot-2026-06-30`, catalog version `1`, taxonomy version `1`.

| Guideline | Applies / lifecycle | Rationale | Secure design decision / control | Evidence owner / expected evidence | Residual risk / exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `N/A` / `not-applicable` | No login, identity, credential, MFA, recovery, or impersonation behavior changes. | Preserve N/A row with source refs. | `review-security`: cite proposal/spec/design scope. | None / no exception. |
| `SEC-SESS-001` | `N/A` / `not-applicable` | No cookies, tokens, refresh, revocation, or session lifetime behavior changes. | Preserve N/A row. | `review-security`: cite no runtime session surface. | None / no exception. |
| `SEC-DATA-001` | `N/A` / `not-applicable` | No PAN, PII, retention, storage, masking, or transmission changes. | Evidence must stay review-safe. | `review-security`: cite artifact-only scope. | None / no exception. |
| `SEC-SECRET-001` | `N/A` / `not-applicable` | No secret creation, storage, rotation, or exposure path changes. | Do not include raw secret values in SDD evidence. | `review-security` and `verify`: inspect docs for no secret values. | None / no exception. |
| `SEC-ACCESS-001` | `Yes` / `planned` | Workflow gates and dependencies are authorization-like progression controls. | Deny-by-default: missing embedded rows or review evidence blocks downstream phases. | `test-design`, `apply`, `review-security`: static routing/dependency evidence. | Drift risk if tokens not updated everywhere / no exception. |
| `SEC-FILE-001` | `N/A` / `not-applicable` | No upload, download, generated export, path traversal, or file content handling feature. | Deleting/retiring validator is repository maintenance, not runtime file handling. | `review-security`: cite changed-file list. | None / no exception. |
| `SEC-DB-001` | `N/A` / `not-applicable` | No database queries, migrations, persistence, or tenant data paths. | Preserve N/A row. | `review-security`: cite no DB files/config. | None / no exception. |
| `SEC-LOG-001` | `Yes` / `planned` | Review/verify/archive reports are audit evidence and must not leak sensitive values. | Safe-evidence rule: paths, sections, summaries, redacted placeholders only. | `test-design`, `review-security`, `verify`: static/manual report evidence. | Evidence quality risk if reviewers paste sensitive context / no exception. |
