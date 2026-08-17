# Design: Mandatory Security Design for SDD Changes

## Technical Approach

Migrate the active SDD DAG from conditional security applicability to mandatory security design:

`explore? -> proposal -> spec -> design -> security-design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`

`security-design.md` becomes the canonical security planning/classification artifact for new changes. `security-applicability.md` remains readable only for legacy/archived changes and MUST NOT gate new-change routing. The change is documentation/prompt/contract driven: update agents, skills, shared contracts, OpenSpec specs, README, and PowerShell validators; no application runtime exists.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Active DAG | Remove `security-applicability` from new-change routing and make `security-design` mandatory after `design`. | Keep conditional applicability, or add security review without DAG migration. | Classification needs technical design context and one canonical authority. |
| Canonical artifact | Put classification, full category/guideline matrix, controls, evidence, exceptions, lifecycle statuses, and archive gates in `security-design.md`. | Keep split `security-applicability.md` + conditional `security-design.md`. | Avoids dual authority and makes no-impact a documented matrix result, not a skipped phase. |
| Security review | Add `sdd-review-security` producing `review-security-report.md` after non-blocking `sdd-review`. | Fold into 96-control review or verify. | Preserves general review boundary while validating implementation evidence for the security matrix. |
| Compatibility | Keep legacy applicability resolver/validator as archive compatibility only. | Delete applicability support. | Archived old changes must remain auditable without retroactive revalidation. |

## Data Flow

```text
proposal/specs -> sdd-design -> design.md
design.md + proposal/specs + catalog -> sdd-security-design -> security-design.md
security-design.md -> test-design/tasks/apply
apply -> review-report.md -> sdd-review-security -> review-security-report.md
review reports + verify evidence -> verify-report.md -> archive
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `agents/sdd/sdd-orchestrator.md` | Modify | Update descriptions, `/sdd-ff`, flowchart, dependency graph, readiness, no-skip rules, native routing guard, state examples, and launch context to the new DAG; add `review-security`; keep `security-applicability` only as legacy token. |
| `agents/sdd/sdd-design.md`, `skills/sdd-design/SKILL.md` | Modify | Required inputs become proposal/specs only; include baseline security best-practice considerations; always route `next_recommended: security-design`; do not block on missing applicability. |
| `agents/sdd/sdd-security-design.md`, `skills/sdd-security-design/SKILL.md` | Modify | Mandatory phase; reads proposal/specs/design/catalog/security contract; always writes `security-design.md`; owns classification, matrix, controls, expected evidence, exceptions, validation, and `next_recommended: test-design`. |
| `agents/sdd/sdd-security-applicability.md`, `skills/sdd-security-applicability/SKILL.md` | Modify/Deprecate | Mark non-active for new changes; document read-only legacy compatibility and no new artifact production. |
| `agents/sdd/sdd-review-security.md`, `skills/sdd-review-security/SKILL.md` | Create | New executor/skill; consumes `security-design.md`, `review-report.md`, changed-file/task/apply evidence; writes `review-security-report.md`; routes non-blocking to verify, blocking to apply/resolve-blockers. |
| `skills/_shared/persistence-contract.md` | Modify | Add `review-security` resolver key/path, state refs (`securityReviewReport`), tokens/enums, mandatory new-change `securityDesign`, legacy applicability notes, verify/archive requirements for both review reports. |
| `skills/_shared/sdd-status-contract.md` | Modify | Add `review-security` token mapping, artifact refs/paths/context, dependency state, routing order `review -> review-security -> verify`, and remove active `securityApplicability` dependency for new changes. |
| `skills/_shared/openspec-convention.md` | Modify | Active layout: mandatory `security-design.md`, `review-security-report.md`, legacy `security-applicability.md` only in archives/old changes. |
| `skills/_shared/sdd-security-contract.md` | Modify | Replace conditional security-design schema with mandatory schema: classification, matrix rows (`Yes/No/N/A`), lifecycle statuses, evidence locations, observations, exceptions, archive gates. |
| `skills/_shared/security-guideline-catalog.md` | Modify | Add matrix vocabulary, lifecycle statuses, mandatory evidence expectations, and review-security cross-reference guidance. |
| `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md` | Modify | Consume mandatory `security-design.md`; remove applicability gates; carry planned/applied evidence for security matrix rows. |
| `skills/sdd-review/SKILL.md` | Modify | Route non-blocking verdicts to `review-security`; cite catalog IDs without owning security matrix. |
| `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` | Modify | Require non-blocking `review-report.md` and `review-security-report.md`; consume but do not duplicate review matrices; archive requires mandatory security-design evidence or complete exceptions. |
| `scripts/validate_security_design.ps1` | Create | Static validator for mandatory `security-design.md` schema, catalog snapshot, guideline IDs, matrix completeness, values/statuses, evidence/exception fields, and `nextRecommended: test-design`. |
| `scripts/validate_security_applicability.ps1` | Modify | Retain as legacy archive compatibility validator only; do not require for new phase success. |
| `openspec/specs/*`, `README.md` | Modify | Sync final workflow/source specs and public phase order after archive. |

## Interfaces / Contracts

Status/native tokens include `review-security`. New-change state includes `artifactRefs.securityDesign` and `artifactRefs.securityReviewReport`; `securityApplicability` is legacy-only. `security-design.md` uses schema `sdd.security-design` and must record classification, catalog identity, every category/guideline row, answer `Yes|No|N/A`, evidence location, observations, lifecycle status, controls, exceptions, and archive gates. `review-security-report.md` records verdict, validated rows, implementation evidence, blocking findings, exceptions, and `nextRecommended: verify|apply|resolve-blockers`.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static contract | DAG tokens, artifact paths, required inputs, routing text, legacy wording removal | Read-back/grep checks over `agents/`, `skills/`, `openspec/specs/`, and `README.md`. |
| Validator | `security-design.md` schema and legacy applicability compatibility | PowerShell smoke checks with fixture artifacts in temp folders; old validator runs only when explicitly targeting legacy artifacts. |
| Integration | Status/persistence/readiness consistency across shared contracts | Manual/static cross-file checklist because repo has no package manifest, linter, or test runner. |

## Migration / Rollout

Implement in stacked slices: shared contracts/status/OpenSpec paths first; phase skills/agents second; validator and review-security third; README/spec sync last. In-flight changes with `security-applicability.md` can be completed through documented legacy compatibility or manually migrated by creating `security-design.md`. Rollback reverts DAG/contracts and removes `review-security` routing while keeping legacy applicability active.

## Open Questions

None blocking.

## Security Applicability Routing

Legacy `sdd-security-applicability` is removed from active new-change routing. `sdd-design` must not require `security-applicability.md`; it always routes to mandatory `sdd-security-design`. `sdd-security-design` owns classification and records no-impact as justified `not-applicable`/`N/A` matrix evidence inside `security-design.md`.
