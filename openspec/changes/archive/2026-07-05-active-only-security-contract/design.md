# Design: Active-Only Security Contract

## Technical Approach

Make the active SDD security contract describe only the current flow: `design.md#secure-development-design` -> `review-security-report.md` -> `verify` -> `archive`. Remove legacy schemas and notes from active security authority surfaces, while keeping old artifact read/display behavior only in low-level status and persistence contracts where existing state or archives may still expose those fields.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Active authority | `skills/_shared/sdd-security-contract.md` becomes active-only: embedded design schema, review-security contract, exception fields, and safe-evidence rules. | Add a separate legacy contract. | A separate legacy contract is explicitly out of scope and would keep retired workflows prominent. |
| Historical readability | Keep minimal `securityApplicability` / `securityDesign` resolver and status fields only in `persistence-contract.md` and `sdd-status-contract.md` if needed. | Delete all historical tokens everywhere. | Old state/archive reads must not break, but historical tokens must not become launchable phases. |
| Cleanup style | Rewrite active phase guidance as positive current-flow rules instead of legacy warnings. | Leave legacy notes as explanatory prose. | Active surfaces should teach the current DAG without making retired artifacts look required. |

## Data Flow

```text
proposal/specs
  -> sdd-design writes design.md#secure-development-design
  -> sdd-test-design/tasks/apply produce planned evidence
  -> sdd-review-security writes review-security-report.md
  -> verify/archive check current-flow evidence

old state/archive refs
  -> status/persistence readers only
  -> no active launch target
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modify | Remove legacy applicability schema and PR compatibility notes; keep only active embedded design, review-security, exception, and safe-evidence contracts. |
| `skills/_shared/security-guideline-catalog.md` | Modify | Remove active-scope wording that names standalone legacy compatibility; preserve snapshot IDs and audit continuity. |
| `skills/sdd-design/SKILL.md` | Modify | Express embedded secure design as the only active classification/design authority and route directly to test-design. |
| `skills/sdd-review-security/SKILL.md` | Modify | Require embedded design and review-security evidence without standalone artifact dependencies. |
| `skills/_shared/persistence-contract.md` | Conditional modify | Preserve historical resolver/state slots only as read/display data, not active dependencies. |
| `skills/_shared/sdd-status-contract.md` | Conditional modify | Preserve bounded legacy tokens only as non-launchable historical data if required. |
| `openspec/specs/sdd-*/spec.md` | Modify | Sync active source specs with current-flow requirements and historical-read boundaries. |

## Interfaces / Contracts

No runtime API is introduced. Contract changes are Markdown SDD contracts. Active new-change successors remain `design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`; `security-applicability` and `security-design` may appear only as historical read data.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static/manual | Legacy filenames/prose removed from active authority surfaces. | Targeted search/read-back of changed Markdown. |
| Static/manual | Status/persistence still expose old refs only as non-launchable read data. | Inspect resolver/routing tables and dependency text. |
| Static/manual | Specs match proposal deltas. | Read changed OpenSpec source specs. |

`openspec/config.yaml` reports no runner, no lint/type/format commands, and `strict_tdd: false`; no automated unit, integration, or E2E tests are planned.

## Migration / Rollout

No migration required. Archives under `openspec/changes/archive/**` remain untouched. Rollout is a Markdown/spec contract update. Rollback is a git revert of the active contract/spec edits for this change.

## Open Questions

None.

## Secure Development Design

**Classification**: `security-impacting` because this changes the security evidence contract and routing authority, though not application runtime behavior.
**Catalog**: `security-guidelines-initial-user-snapshot-2026-06-30`, catalog version `1`, taxonomy version `1`.

| Guideline | Applies / lifecycle | Rationale | Secure design decision / control | Evidence owner / expected evidence | Residual risk / exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `N/A` / `not-applicable` | No login, identity, MFA, recovery, or credential flow changes; evidence is proposal scope and affected Markdown paths. | Preserve no auth-flow behavior changes. | review-security: cite changed files and no runtime auth surfaces. | None. |
| `SEC-SESS-001` | `N/A` / `not-applicable` | No cookies, tokens, sessions, or revocation behavior changes. | Preserve no session behavior changes. | review-security: static evidence from file list and specs. | None. |
| `SEC-DATA-001` | `Yes` / `planned` | Evidence contracts mention sensitive data handling. | Require review-safe evidence paths/summaries/redacted placeholders only. | test-design/review-security/verify: confirm no PAN/PII/raw confidential values in reports. | None. |
| `SEC-SECRET-001` | `Yes` / `planned` | Safe-evidence rules must continue forbidding secret disclosure. | Keep no raw credentials/tokens/private keys in SDD evidence. | review-security/verify: static search/read-back for redacted-only evidence. | None. |
| `SEC-ACCESS-001` | `Yes` / `planned` | Routing authority controls protected workflow progression. | Historical tokens are non-launchable; active gates require embedded design and review-security. | review-security/verify: inspect status/persistence routing tables. | Risk: accidental runnable legacy wording; mitigated by targeted search. |
| `SEC-FILE-001` | `N/A` / `not-applicable` | No upload/download/generated-file handling is changed; Markdown artifacts only. | Preserve no file-handling runtime behavior. | review-security: cite affected Markdown-only paths. | None. |
| `SEC-DB-001` | `N/A` / `not-applicable` | No database, query, migration, or persistence engine code changes. | Preserve no DB behavior changes. | review-security: cite no DB files or commands changed. | None. |
| `SEC-LOG-001` | `Yes` / `planned` | Reports and evidence are audit records. | Keep useful audit evidence without raw secrets, PAN, credentials, tokens, or sensitive payloads. | review-security/archive: cite sanitized report sections and unavailable-tooling note. | None. |
