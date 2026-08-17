# Exploration: Inline secure development design in `design.md`

## Current State

The repository is an SDD workflow asset repo, not an application runtime. Current documented phase order is `explore? -> propose -> spec -> design -> security-design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive` (`README.md`, `agents/sdd/sdd-orchestrator.md`).

Security design is currently a dedicated active new-change phase:

- `agents/sdd/sdd-design.md` and `skills/sdd-design/SKILL.md` route successful design to `security-design`.
- `agents/sdd/sdd-security-design.md` and `skills/sdd-security-design/SKILL.md` create mandatory `security-design.md` after technical design.
- `skills/_shared/persistence-contract.md` and `skills/_shared/sdd-status-contract.md` model `securityDesign` as a separate artifact/dependency.
- `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, and `skills/sdd-archive/SKILL.md` require a separate mandatory `security-design` input.
- `skills/sdd-review-security/SKILL.md` validates `security-design.md` rows after general review and before verification.
- Source specs `openspec/specs/sdd-security-design-workflow/spec.md`, `openspec/specs/sdd-review-security-workflow/spec.md`, and `openspec/specs/sdd-test-design-workflow/spec.md` encode this separate-phase model.

The corporate catalog already has the desired 8 compact guideline IDs: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001` in `skills/_shared/security-guideline-catalog.md`.

## Affected Areas

- `README.md` — phase list and workflow explanation must remove active `sdd-security-design` and describe secure design as part of `design.md`.
- `agents/sdd/sdd-orchestrator.md` — DAG, routing tokens, dependency graph, status recovery, required context table, command docs, and artifact naming all currently include `security-design` as an active planning phase.
- `agents/sdd/sdd-design.md` and `skills/sdd-design/SKILL.md` — design must become responsible for secure development design, the 8-guideline applicability matrix, evidence expectations, and routing directly to `test-design`.
- `agents/sdd/sdd-security-design.md` and `skills/sdd-security-design/SKILL.md` — likely retire from active new-change routing. Keep only archive/legacy compatibility if needed, or delete/deprecate in a later proposal.
- `agents/sdd/sdd-review-security.md` and `skills/sdd-review-security/SKILL.md` — keep phase, but switch source input from separate `security-design.md` to the mandatory security section inside `design.md`.
- `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` — downstream phases must consume security controls/evidence from `design.md` instead of `security-design.md`.
- `skills/_shared/sdd-security-contract.md` — redefine the active security schema as an embedded `design.md` section while preserving legacy `security-design.md` compatibility for old archives.
- `skills/_shared/security-guideline-catalog.md` — scope wording should point to secure design inside `design.md` plus `review-security-report.md`.
- `skills/_shared/persistence-contract.md` and `skills/_shared/sdd-status-contract.md` — remove `securityDesign` as an active required artifact/dependency for new changes, while keeping legacy fields readable.
- `scripts/validate_security_design.ps1` — currently validates standalone `security-design.md`; either replace with a `design.md` secure-section validator or update it to validate embedded fenced metadata.
- Source specs under `openspec/specs/` — update separate security-design workflow spec into embedded secure design requirements; update review-security and test-design specs accordingly.

## Approaches

1. **Inline secure design section in `design.md` and retire active `sdd-security-design`** — `sdd-design` writes technical design plus a mandatory `## Secure Development Design` section containing catalog identity, all 8 guideline rows, applicability, rationale, controls, evidence expectations, exceptions, and archive-gate notes. The next phase becomes `test-design`.
   - Pros: removes one planning phase, keeps security decisions with architecture decisions, reduces routing and artifact overhead, matches the requested workflow.
   - Cons: larger `design.md`; many contracts and downstream readers must be updated together to avoid broken continuation/status behavior.
   - Effort: High.

2. **Keep `sdd-security-design` but make it update `design.md` instead of writing `security-design.md`** — preserve phase token and DAG edge while changing the produced artifact location.
   - Pros: smaller routing/status changes; less native-token churn.
   - Cons: does not satisfy the request to remove `sdd-security-design` as an active new-change phase; still forces a separate phase.
   - Effort: Medium.

3. **Alias `security-design` to `design` in status while keeping old files** — status treats `securityDesign` as complete when `design.md` has the section, but old phase/skill remains present.
   - Pros: compatibility bridge during migration.
   - Cons: ambiguous authority; easy for downstream phases to keep reading stale `security-design.md`; not clean enough as the final model.
   - Effort: Medium.

## Recommendation

Use Approach 1 as the target design, with explicit legacy compatibility. The proposal should define a new mandatory `design.md` section named `## Secure Development Design` that includes:

- Catalog snapshot metadata from `skills/_shared/security-guideline-catalog.md`.
- One row for each of the 8 guideline IDs.
- `Applies` values using `Yes` or `N/A` at design time; reserve `No` for review/security evidence failures.
- Rationale and evidence refs for every N/A row.
- Required secure design decision, control, evidence owner, expected evidence, lifecycle status, residual risk, and exception fields for every applicable row.
- Safe-evidence rules for `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, and `SEC-LOG-001`.

`sdd-review-security` should remain a separate post-review validation phase. It should validate the embedded `design.md` secure section against the corporate catalog, changed files, apply/task evidence, `test-design.md`, and non-blocking `review-report.md`, then persist `review-security-report.md` with row-level `Yes` / `No` / `N/A` validation and evidence locations.

## Risks

- Native/status token drift: `security-design` is currently a bounded routing token in orchestrator/status/state contracts. Removing it requires synchronized changes across state schema, status recovery, DAG tables, and source specs.
- Downstream breakage: test-design, tasks, apply, verify, and archive all currently require separate `security-design.md`; partial updates would block continuation.
- Validator mismatch: `scripts/validate_security_design.ps1` currently targets standalone `security-design.md` and must be replaced or adapted.
- Legacy archive compatibility: archived changes still contain `security-design.md`; readers should preserve compatibility while preventing new changes from requiring that artifact.
- Artifact size: `skills/sdd-design/SKILL.md` currently budgets design under 800 words; embedding 8 guideline rows may require either a larger budget or compact table-first format.

## Ready for Proposal

Yes. The proposal should scope this as an SDD workflow redesign affecting prompts, skills, shared contracts, validators, README, and OpenSpec source specs. The implementation should explicitly avoid modifying application runtime files because this repository is an AI workflow asset repository.
