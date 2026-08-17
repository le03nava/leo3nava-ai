# Apply Evidence: Mandatory Security Design for SDD Changes

## Scope

Work Unit 4 / PR4 updated public documentation and source OpenSpec specs, then recorded readiness evidence for review. This apply slice is intentionally limited to `README.md`, source `openspec/specs/*`, this evidence artifact, and task/state progress. It does not create review, security-review, verify, or archive artifacts.

## Workload / PR Boundary

| Field | Evidence |
| --- | --- |
| Delivery strategy | `auto-chain` |
| Chain strategy | `stacked-to-main` |
| Current slice | PR4 / Work Unit 4: README, source OpenSpec specs, and final apply-readiness evidence only |
| Prior dependency | PR1 contracts, PR2 phase routing, PR3 validators/evidence are represented by existing completed task checkboxes 1.1-3.2 |
| Out of scope | Runtime implementation, review reports, security-review reports, verify reports, archive artifacts |

## Static Commands

| Command | Result | Notes |
| --- | --- | --- |
| `powershell -NoProfile -ExecutionPolicy Bypass -File scripts/validate_security_design.ps1 -Path openspec/changes/replace-security-applicability-with-mandatory-security-design/security-design.md -AllowManualPending` | PASS | `security-design.md` validates against the mandatory security-design schema. |
| `git diff --check` | PASS with warning | No whitespace errors were reported. Git emitted a line-ending warning for a previously changed agent file in the stacked working tree. |

## Runtime Tooling Availability

`openspec/config.yaml` states this repository has no package manifest, build file, or executable test runner. The testing section reports unavailable runtime tooling:

- Test runner: unavailable; command empty.
- Unit/integration/E2E layers: unavailable.
- Coverage: unavailable; command empty.
- Linter: unavailable; command empty.
- Type checker: unavailable; command empty.
- Formatter: unavailable; command empty.

No runtime test, lint, type-check, format, coverage, or build command was invented or run for this apply slice. Evidence is static/manual plus the PowerShell security-design validator.

## TD Evidence Summary

| ID | Apply Evidence |
| --- | --- |
| TD-001 | README and source specs now document the active DAG as `explore? -> propose -> spec -> design -> security-design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`. |
| TD-002 | `openspec/specs/sdd-security-applicability-workflow/spec.md` scopes `sdd-security-applicability` to legacy/archive compatibility and forbids new-change active routing or artifact production. |
| TD-003 | `openspec/specs/sdd-security-design-workflow/spec.md` requires every new change to route from design into mandatory security design before test design. |
| TD-004 | Source security-design spec requires `security-design.md` to own classification, catalog identity, matrix rows, controls, evidence, statuses, risks, exceptions, and archive gates. |
| TD-005 | Source test-design spec blocks test design/task planning until mandatory `security-design.md` exists. |
| TD-006 | Source test-design spec requires checks or justified non-test evidence for mandatory security-design rows and preserved N/A evidence. |
| TD-007 | Source review spec routes non-blocking `sdd-review` to `sdd-review-security`, not directly to verify. |
| TD-008 | Added source `openspec/specs/sdd-review-security-workflow/spec.md` requiring persisted `review-security-report.md`. |
| TD-009 | Review-security source spec validates every security-design row with matrix answer, lifecycle status, evidence, and observations. |
| TD-010 | Review-security source spec preserves the boundary with the 96-control general review matrix. |
| TD-011 | Execution/persistence source spec requires `security-design.md`, `review-security-report.md`, `review-security` token, refs, paths, dependency states, and no active `securityApplicability` dependency. |
| TD-012 | Execution/persistence source spec requires verify/archive to consume both review artifacts without owning their matrices. |
| TD-013 | Source specs and README align OpenSpec layout expectations around mandatory `security-design.md`, `test-design.md`, and `review-security-report.md`; legacy applicability is archive-only. |
| TD-014 | This slice synchronizes the main OpenSpec specs with the phase/contract work completed in prior PR slices. |
| TD-015 | README and source OpenSpec specs were synchronized for the final public workflow and compatibility model. |
| TD-016 | Validator evidence uses `scripts/validate_security_design.ps1`; source specs no longer require applicability validation for new changes. |
| TD-017 | Validator command passed for this change's `security-design.md` artifact. |
| TD-018 | Negative fixture execution is not part of PR4; validator behavior contract was implemented in PR3 and remains covered by source spec requirements. |
| TD-019 | Applicability source spec and README preserve legacy/archive-only compatibility. |
| TD-020 | Static grep/read-back over README and source specs found no active route requiring `security-applicability` before design; remaining references are legacy/archive compatibility or historical change artifacts. |
| TD-021 | Execution/persistence source spec uses valid native tokens and camelCase artifact refs for `securityDesign`, `testDesign`, and `securityReviewReport`. |
| TD-022 | Source specs require `review-security-report.md` before verify/archive for new changes. |
| TD-023 | Security catalog source spec requires review-safe evidence using paths, summaries, static inspection, approved exceptions, or N/A evidence rather than raw sensitive values. |
| TD-024 | Security catalog and security-design source specs preserve secret-safe evidence expectations; this artifact records only paths, summaries, and command names. |
| TD-025 | Source specs enforce denial-by-default phase gates through mandatory security design and security review before downstream phases. |
| TD-026 | Review-security and execution/persistence source specs require audit evidence in report artifacts without duplicating unrelated matrices or sensitive payloads. |
| TD-027 | Security-design source spec requires no-impact rows inside `security-design.md`; absence of `security-design.md` is not no-impact proof. |
| TD-028 | This artifact explicitly reports unavailable runtime test, lint, type, format, and coverage tooling from `openspec/config.yaml`. |

## Security Control Evidence

| Guideline ID | Evidence |
| --- | --- |
| `SEC-DATA-001` | Evidence in README/specs/apply artifact uses file paths, workflow summaries, and command summaries only; no PAN/PII examples are included. |
| `SEC-SECRET-001` | No secret values are recorded. Evidence uses artifact paths and command names; placeholder/secret policy remains in source catalog and security-design contracts. |
| `SEC-ACCESS-001` | Source specs now enforce mandatory security-design and review-security gates before test-design/verify/archive. |
| `SEC-LOG-001` | Review-security/report contracts require useful audit evidence through locations and observations without raw sensitive payloads. |

## Read-Back Checklist

- `README.md` documents mandatory `sdd-security-design`, `sdd-review-security`, and the new phase order.
- `openspec/specs/sdd-security-design-workflow/spec.md` requires `security-design.md` for every new change.
- `openspec/specs/sdd-security-applicability-workflow/spec.md` limits applicability to legacy/archive compatibility.
- `openspec/specs/sdd-test-design-workflow/spec.md` consumes the mandatory security-design matrix.
- `openspec/specs/sdd-review-workflow/spec.md` routes non-blocking review to security review.
- `openspec/specs/sdd-review-security-workflow/spec.md` defines the new security review gate.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` defines review-security routing, refs, verify/archive consumption, and compatibility.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` defines matrix vocabulary and review-safe evidence expectations.

## Result

Tasks 4.1 and 4.2 are complete. The change is ready for `sdd-review` after this apply phase.
