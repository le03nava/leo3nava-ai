# Review Security Report: Active-Only Security Contract

```yaml
schemaName: sdd.review-security-report
schemaVersion: 1
changeName: active-only-security-contract
verdict: PASS WITH WARNINGS
sourceSecureDesign: openspec/changes/active-only-security-contract/design.md#secure-development-design
sourceReviewReport: openspec/changes/active-only-security-contract/review-report.md
nextRecommended: verify
```

## Summary

Security review validates all 8 embedded secure-development rows from `design.md#secure-development-design` against apply evidence, task completion, test-design coverage, changed-file context, the shared security catalog, and the non-blocking general review handoff.

Verdict is **PASS WITH WARNINGS**. There are no security blockers and no approved exceptions. The only warning is unavailable runtime/lint/type/format/coverage tooling; static/manual evidence is the scoped validation method for this Markdown/spec contract change.

Evidence in this report is intentionally review-safe: paths, section references, line ranges, summaries, and negative scan summaries only. It does not reproduce raw secrets, PAN, PII, private keys, credentials, tokens, or confidential payloads.

## Security Row Validation

| Guideline ID | Category | Design Applies | Lifecycle Status | Complies | Evidence Location | Observations | Finding |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | N/A | `not-applicable` | N/A | `design.md:70`; `test-design.md:51`; `apply-progress.md:81`; `review-report.md:59` | Change is Markdown/spec workflow contract work only; no login, identity, MFA, recovery, impersonation, or credential-flow surface changed. | None |
| `SEC-SESS-001` | `sessions` | N/A | `not-applicable` | N/A | `design.md:71`; `test-design.md:52`; `apply-progress.md:82`; `review-report.md:46` | No cookie, bearer token, refresh, revocation, session lifetime, or cache/session result behavior changed. | None |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | `implemented` | Yes | `design.md:72`; `test-design.md:43,53`; `apply-progress.md:69,77`; `review-report.md:65,143` | Safe-evidence controls are preserved; implementation evidence reports no raw PAN/PII/confidential payload indicators in changed shared security files. | None |
| `SEC-SECRET-001` | `secrets` | Yes | `implemented` | Yes | `design.md:73`; `test-design.md:43,54`; `apply-progress.md:69,78`; `review-report.md:68,143`; `skills/_shared/sdd-security-contract.md:24,135` | Contract and evidence keep raw credentials, secret values, tokens, and private keys out of SDD artifacts; findings use only summaries and paths. | None |
| `SEC-ACCESS-001` | `permissions-access-control` | Yes | `implemented` | Yes | `design.md:74`; `test-design.md:55`; `apply-progress.md:79,110-119`; `review-report.md:55,111,135-140`; `skills/_shared/sdd-status-contract.md:27-48,192-216`; `agents/sdd/sdd-orchestrator.md:847` | Active routing authority remains embedded secure design plus review-security. Historical `securityDesign` / `securityApplicability` values are read-only compatibility data and are not runnable successors for new changes. | None |
| `SEC-FILE-001` | `files` | N/A | `not-applicable` | N/A | `design.md:75`; `test-design.md:56`; `apply-progress.md:83`; `review-report.md:45,82,84,139` | No runtime upload, download, generated-file, path traversal, or file metadata behavior changed; changed-file evidence excludes archive rewrites. | None |
| `SEC-DB-001` | `database-access` | N/A | `not-applicable` | N/A | `design.md:76`; `test-design.md:57`; `apply-progress.md:84`; `review-report.md:70` | No database queries, migrations, ORM filters, tenant isolation paths, or persistence engine code changed. | None |
| `SEC-LOG-001` | `sensitive-logging` | Yes | `implemented` | Yes | `design.md:77`; `test-design.md:58`; `apply-progress.md:80`; `review-report.md:48,66`; `skills/_shared/sdd-security-contract.md:137` | Audit/report evidence remains useful without raw secrets, PAN, credentials, tokens, or sensitive payloads. | None |

## Implementation Evidence

| Evidence Area | Result | Review-Safe Evidence |
| --- | --- | --- |
| Embedded secure-design row count | PASS | `design.md:63-77` contains all 8 compact catalog IDs exactly once. |
| Catalog vocabulary | PASS | `skills/_shared/security-guideline-catalog.md:62-69` preserves all 8 guideline IDs; `skills/_shared/security-guideline-catalog.md:33-39` defines `Yes` / `No` / `N/A` and lifecycle status vocabulary. |
| Test-design coverage | PASS | `test-design.md:47-58` maps each SEC row to static/manual coverage; `test-design.md:64-73` defines evidence expectations and safe-evidence constraints. |
| Apply/task completion | PASS | `tasks.md:41-61` and `apply-progress.md:13-27` show all planned tasks complete; `apply-progress.md:47-71` records TD-001..TD-021 evidence. |
| Changed-file scope | PASS | `review-report.md:24` and current changed-file inspection show active contract/spec files only: `agents/sdd/sdd-orchestrator.md`, five `openspec/specs/sdd-*/spec.md` files, shared contracts/catalog, and design/review-security skills. No `openspec/changes/archive/**` path appears in the changed-file list. |
| Active authority | PASS | `skills/_shared/sdd-security-contract.md:26-107`, `skills/sdd-review-security/SKILL.md:28-38`, and `skills/_shared/sdd-status-contract.md:27-48,192-216` keep new-change authority on embedded design rows and review-security evidence. |
| Historical refs | PASS | `skills/_shared/sdd-status-contract.md:27-37,194-208`, `skills/_shared/persistence-contract.md`, and `agents/sdd/sdd-orchestrator.md:847` scope historical `security-design` / `security-applicability` values to read-only old/archive compatibility, not launchable phases. |
| Safe evidence | PASS | Static/manual evidence reports no raw secret, PAN, PII, private-key, credential, token, or confidential payload indicators in changed shared files; this report preserves the same path/summary-only boundary. |

## General Review Handoff

`review-report.md` is non-blocking and supports continuation to security review:

- Verdict: **PASS WITH WARNINGS** (`review-report.md:3-12`).
- Blocking failures: `0` (`review-report.md:8-10`).
- Blocking summary: no blocking follow-up; route to `review-security` (`review-report.md:13-18`).
- Handoff explicitly states security review remains the authority for row-level security validation (`review-report.md:26`).
- Remaining warning: automated runtime/lint/type/format/coverage evidence is unavailable and must be preserved in security review and verify (`review-report.md:27-28,152-153`).

This report cites general review as supporting evidence only and does not duplicate the 96-control matrix.

## Exceptions

None.

## Blockers and Non-Blocking Findings

### Blockers

None.

### Non-Blocking Findings

| ID | Guideline IDs | Owner | Route | Evidence | Finding |
| --- | --- | --- | --- | --- | --- |
| RS-WARN-001 | `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-LOG-001` | verify | verify | `openspec/config.yaml:17-46`; `review-report.md:27-28`; `apply-progress.md:86-96` | Runtime tests, coverage, lint, type-check, and format commands are unavailable. This is acceptable for this Markdown/spec contract review when preserved as a warning and not treated as passing automated evidence. |

## Unavailable Tooling

Per `openspec/config.yaml#testing`, no automated runtime test runner, coverage command, linter, type checker, formatter, or strict TDD command is configured:

| Capability | Configured | Evidence |
| --- | --- | --- |
| Runtime tests | No | `openspec/config.yaml:20-23` |
| Coverage | No | `openspec/config.yaml:34-36` |
| Lint | No | `openspec/config.yaml:37-40` |
| Type-check | No | `openspec/config.yaml:41-43` |
| Format | No | `openspec/config.yaml:44-46` |

Missing tooling is reported as unavailable evidence, not as passing evidence.

## Next Route

`nextRecommended: verify`
