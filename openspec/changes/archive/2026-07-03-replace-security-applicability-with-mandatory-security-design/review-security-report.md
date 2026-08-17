# Review Security Report: Mandatory Security Design for SDD Changes

```yaml
schemaName: sdd.review-security-report
schemaVersion: 1
changeName: replace-security-applicability-with-mandatory-security-design
verdict: PASS WITH WARNINGS
sourceSecurityDesign: openspec/changes/replace-security-applicability-with-mandatory-security-design/security-design.md
sourceReviewReport: openspec/changes/replace-security-applicability-with-mandatory-security-design/review-report.md
nextRecommended: verify
```

## Summary

Security review validated all 8 compact security guideline rows from `security-design.md` against the catalog, security contract, completed tasks, apply evidence, test-design coverage, non-blocking general review handoff, and changed-file context. No blocking security findings were identified. The report returns `PASS WITH WARNINGS` because configured runtime test/lint/type/format/coverage/build tooling is unavailable and evidence is static/manual plus PowerShell validator output.

Evidence in this report is intentionally review-safe: paths, section references, command summaries, and high-level observations are used instead of raw secrets, credentials, PAN, PII, tokens, or unnecessary confidential values.

## Security Row Validation

| Guideline ID | Category | Design Applies | Lifecycle Status | Complies | Evidence Location | Observations | Finding |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | N/A | `not-applicable` | N/A | `security-design.md#non-applicable-rationale`; `design.md#technical-approach` | The change does not modify login, identity proofing, MFA, credential validation, impersonation, account recovery, or protected URL behavior. N/A rationale and evidence are present. | None |
| `SEC-SESS-001` | `sessions` | N/A | `not-applicable` | N/A | `security-design.md#non-applicable-rationale`; `design.md#technical-approach` | The change does not modify cookies, bearer tokens, refresh tokens, server-side sessions, logout, revocation, or session lifetime behavior. N/A rationale and evidence are present. | None |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | `implemented` | Yes | `security-design.md#security-matrix`; `apply-evidence.md#security-control-evidence`; `test-design.md#security-control-coverage`; `skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls`; `skills/_shared/security-guideline-catalog.md#review-phase-cross-reference-guidance` | Security evidence contracts and apply evidence require paths, sections, summaries, command names, and redacted placeholders instead of raw PAN, PII, credentials, or confidential values. Static sensitive-literal inspection found only documented placeholder examples in the test plan, not actual sensitive values. | Warning: verify must repeat or cite safe-evidence inspection because runtime tooling is unavailable. |
| `SEC-SECRET-001` | `secrets` | Yes | `implemented` | Yes | `security-design.md#security-matrix`; `apply-evidence.md#security-control-evidence`; `test-design.md#security-control-coverage`; `skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls`; `scripts/validate_security_design.ps1` | Contracts, catalog guidance, review-security contract text, and validator expectations prohibit committing, echoing, or reproducing secret values in SDD artifacts. Evidence uses artifact paths, command summaries, and placeholder policy only. | Warning: verify must preserve the same no-secret evidence boundary. |
| `SEC-ACCESS-001` | `permissions-access-control` | Yes | `implemented` | Yes | `apply-evidence.md#td-evidence-summary`; `tasks.md#phase-1-contracts-pr-1`; `tasks.md#phase-2-routing-and-phases-pr-2`; `agents/sdd/sdd-orchestrator.md`; `skills/_shared/persistence-contract.md`; `openspec/specs/sdd-execution-persistence-contracts/spec.md` | The active workflow now enforces denial-by-default phase progression through mandatory `security-design` and `review-security` before verify/archive. General review confirmed routing evidence and no blocking findings. | None |
| `SEC-FILE-001` | `files` | N/A | `not-applicable` | N/A | `security-design.md#non-applicable-rationale`; `design.md#file-changes` | The change edits repository Markdown/contracts/scripts only and does not add upload, download, generated export, user-controlled path, path traversal, file metadata, or file authorization behavior. N/A rationale and evidence are present. | None |
| `SEC-DB-001` | `database-access` | N/A | `not-applicable` | N/A | `security-design.md#non-applicable-rationale`; `openspec/config.yaml#context`; `design.md#technical-approach` | The repository has no database runtime and this change adds no queries, migrations, persistence access paths, tenant boundaries, or data jobs. N/A rationale and evidence are present. | None |
| `SEC-LOG-001` | `sensitive-logging` | Yes | `implemented` | Yes | `security-design.md#security-matrix`; `apply-evidence.md#security-control-evidence`; `test-design.md#security-control-coverage`; `skills/sdd-review-security/SKILL.md#report-format`; `skills/_shared/security-guideline-catalog.md#review-phase-cross-reference-guidance` | Security review and archive/report contracts require audit-useful evidence through locations and observations without raw secrets, PAN, credentials, tokens, or unnecessary sensitive operational context. This report follows that boundary. | Warning: verify/archive must continue citing locations and summaries only. |

## General Review Handoff

`review-report.md` reports `PASS WITH WARNINGS`, `Blocking failures: 0`, `Non-blocking findings: 2`, and `Next recommendation: review-security`. Supporting evidence relevant to this security review:

- General review inspected OpenSpec config, proposal, specs, design, `security-design.md`, `test-design.md`, completed `tasks.md`, `apply-evidence.md`, changed-file status/stat context, SDD review catalog, and review report template.
- General review confirms all 12 implementation tasks are checked complete and no unchecked implementation tasks remain.
- General review confirms `scripts/validate_security_design.ps1 -Path .../security-design.md -AllowManualPending` passed.
- General review records unavailable runtime test/lint/type/format/coverage/build tooling and a non-blocking line-ending normalization warning for `agents/sdd/sdd-design.md`.
- The 96-control matrix remains in `review-report.md` and is not duplicated here.

## Exceptions

None.

## Blockers and Non-Blocking Findings

### Blocking Findings

None.

### Non-Blocking Findings

| Guideline ID | Owner | Route | Evidence | Finding |
| --- | --- | --- | --- | --- |
| `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-LOG-001` | `verify` | `verify` | `openspec/config.yaml#testing`; `apply-evidence.md#runtime-tooling-availability`; `review-report.md#non-blocking-findings` | Runtime test, lint, type-check, format, coverage, and build tooling are unavailable by repository configuration. Security evidence is static/manual plus PowerShell validator evidence; missing runtime tooling is not treated as passing evidence. |
| `SEC-ACCESS-001` | `verify` | `verify` | `review-report.md#non-blocking-findings`; command summary from general review | `git diff --check` passed but emitted a line-ending normalization warning for `agents/sdd/sdd-design.md`. No security blocker was identified. |

## Unavailable Tooling

Per `openspec/config.yaml#testing`, the repository has no configured runtime test runner, unit/integration/e2e layer, coverage command, linter, type checker, formatter, or build command. These tools are unavailable, not passed. Review-security evidence is therefore based on artifact read-back, changed-file context, general review handoff, task/apply evidence, security-design validation, and static/manual inspection.

During this phase, `rg` was also unavailable in the shell environment; static sensitive-literal inspection used the platform content-search tooling instead. This is not counted as runtime test evidence.

## Validation Notes

- Report schema is `sdd.review-security-report` with `schemaVersion: 1`.
- `nextRecommended` is `verify` because there are no blocking security findings.
- All 8 compact SEC rows appear exactly once in the Security Row Validation matrix.
- Matrix answers use only `Yes` or `N/A`; no `No` rows are present.
- Lifecycle statuses use catalog vocabulary: `implemented` for applicable controls with apply evidence and `not-applicable` for justified N/A rows.
- N/A rows include evidence locations and scope rationale.
- No approved exceptions are required.
- Safe-evidence rules were followed: no raw secrets, credentials, PAN, PII, tokens, private keys, or unnecessary confidential values are copied into this report.
