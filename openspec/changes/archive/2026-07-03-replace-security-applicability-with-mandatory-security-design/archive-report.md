# Archive Report: Mandatory Security Design for SDD Changes

```yaml
schemaName: sdd.archive-report
schemaVersion: 1
changeName: replace-security-applicability-with-mandatory-security-design
artifactStore: openspec
verdict: SUCCESS
archivedAt: 2026-07-03T15:41:31.6033096-06:00
archiveDestination: openspec/changes/archive/2026-07-03-replace-security-applicability-with-mandatory-security-design
nextRecommended: none
```

## Readiness Validation

| Gate | Result | Evidence |
| --- | --- | --- |
| Implementation tasks | PASS | `tasks.md` has 12 checked implementation tasks and 0 unchecked implementation tasks. |
| General review | PASS WITH WARNINGS / non-blocking | `review-report.md` records `Blocking failures: 0` and only non-blocking warnings. |
| Security review | PASS WITH WARNINGS / non-blocking | `review-security-report.md` records no blocking findings and no exceptions. |
| Verification | PASS WITH WARNINGS / archive-ready | `verify-report.md` records 0 critical issues, 0 blocking issues, `archiveReady: true`, and `nextRecommended: archive`. |
| Mandatory security design | PASS | `security-design.md` exists and `scripts/validate_security_design.ps1 -Path ... -AllowManualPending` passed. |
| Legacy security applicability | PASS / intentionally absent | `securityApplicability: []` is expected for this change. The active DAG intentionally removes `sdd-security-applicability`; compatibility is preserved in synced legacy specs only. |

## Specs Synced

| Domain | Action | Details |
| --- | --- | --- |
| `sdd-execution-persistence-contracts` | Verified already synced | Source spec already contained the delta requirements and scenarios for the new review-security route and mandatory security refs. |
| `sdd-review-security-workflow` | Verified already synced | Source spec already matched the new mandatory security review workflow. |
| `sdd-review-workflow` | Verified already synced | Source spec already matched the apply -> review -> review-security routing and handoff requirements. |
| `sdd-security-applicability-workflow` | Verified already synced | Source spec already preserved legacy-only applicability behavior and new-change absence of `security-applicability.md`. |
| `sdd-security-design-workflow` | Verified already synced | Source spec already matched mandatory security design for every new change and direct classification from proposal/spec/design. |
| `sdd-security-guideline-catalog` | Updated | Normalized the N/A security review scenario wording from `comment` to `observations` to match the delta and review-security matrix vocabulary. |
| `sdd-test-design-workflow` | Updated | Renamed the security-design prerequisite scenario to match the delta and added explicit shared-persistence delegation wording while preserving existing unrelated requirements. |

No requirements were removed or renamed. No destructive merge was performed. Unrelated requirements and scenarios in the source specs were preserved.

## Archive Contents

The archive contains the full OpenSpec change folder at `openspec/changes/archive/2026-07-03-replace-security-applicability-with-mandatory-security-design/`, including:

- `proposal.md`
- `specs/` with seven delta spec domains
- `design.md`
- `security-design.md`
- `test-design.md`
- `tasks.md`
- `apply-evidence.md`
- `review-report.md`
- `review-security-report.md`
- `verify-report.md`
- `state.yaml`
- `archive-report.md`

No active `security-applicability.md` artifact is required or expected for this change.

## Review, Security Review, and Verify Verdicts

- General review: `PASS WITH WARNINGS`, 0 blocking failures, 2 non-blocking findings.
- Security review: `PASS WITH WARNINGS`, no blocking security findings, no approved exceptions needed.
- Verification: `PASS WITH WARNINGS`, 0 critical issues, 0 blocking issues, archive-ready, next recommendation `archive`.

## Security Design Evidence

- Mandatory security design artifact: `security-design.md`.
- Classification: `security-impacting` because the change redesigns SDD security governance, evidence gates, security matrix ownership, and archive readiness behavior.
- Catalog snapshot: `security-guidelines-initial-user-snapshot-2026-06-30`, catalog version 1, taxonomy version 1, source `skills/_shared/security-guideline-catalog.md`.
- Compact guideline rows: 8 total.
- Applicable mandatory controls verified through static/manual evidence and security review: `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-LOG-001`.
- N/A guideline rows preserved with rationale and evidence: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-FILE-001`, `SEC-DB-001`.
- Exceptions: none.
- Validator evidence: `powershell -NoProfile -ExecutionPolicy Bypass -File scripts/validate_security_design.ps1 -Path openspec/changes/replace-security-applicability-with-mandatory-security-design/security-design.md -AllowManualPending` returned `PASS: security design artifact is valid` before archive move.

## Unavailable Tooling Warnings

The repository has no configured runtime test runner, unit/integration/e2e layers, coverage command, linter, type checker, formatter, or build command per `openspec/config.yaml`. These tools are unavailable, not passing evidence. The accepted archive evidence is static/manual review plus the available PowerShell security-design validator.

Additional non-blocking warning carried forward: `git diff --check` passed during review/verify but emitted a CRLF normalization warning for `agents/sdd/sdd-design.md`.

## Source of Truth Updated

The following source specs now reflect the archived behavior:

- `openspec/specs/sdd-execution-persistence-contracts/spec.md`
- `openspec/specs/sdd-review-security-workflow/spec.md`
- `openspec/specs/sdd-review-workflow/spec.md`
- `openspec/specs/sdd-security-applicability-workflow/spec.md`
- `openspec/specs/sdd-security-design-workflow/spec.md`
- `openspec/specs/sdd-security-guideline-catalog/spec.md`
- `openspec/specs/sdd-test-design-workflow/spec.md`

## Archive Verification

Post-move verification confirmed the active change folder is gone and the archived contents are readable. The SDD cycle is complete and no next phase is recommended.
