# Exploration: replace-security-applicability-with-mandatory-security-design

## Current State

The current SDD workflow treats security in two separate planning gates:

```text
explore? -> proposal -> spec -> security-applicability -> design -> security-design? -> test-design -> tasks -> apply -> review -> verify -> archive
```

- `agents/sdd/sdd-orchestrator.md` encodes this DAG in the flowchart, routing table, dependency graph, no-skip rules, state schema, recovery examples, and phase context tables. It currently launches `sdd-security-applicability` after specs, then lets `sdd-design` route either to `security-design` or directly to `test-design` based on `securityImpact`.
- `skills/sdd-security-applicability/SKILL.md` owns the existing independent classification phase and writes `security-applicability.md`. The artifact has a hardened schema: catalog snapshot identity, every supported taxonomy category, no-impact proof, validation metadata, and static validator support.
- `skills/sdd-security-design/SKILL.md` is currently conditional. It creates `security-design.md` only when applicability is security-impacting and explicitly skips artifact creation for valid no-impact changes.
- `skills/sdd-design/SKILL.md` depends on `security-applicability.md` and uses it for post-design routing. It does not currently own baseline security design; it documents a `Security Applicability Routing` section.
- `sdd-test-design`, `sdd-tasks`, `sdd-apply`, `sdd-review`, `sdd-verify`, and `sdd-archive` already consume security applicability and required security design, but most wording says `security-design` is required only for security-impacting changes.
- `skills/sdd-review/references/control-catalog.md` and `report-template.md` define the 96-control review matrix with `Complies: Yes | No | N/A`, evidence location, and observations. Review can cite security guideline IDs, but it is explicitly not security authority.
- `skills/_shared/sdd-security-contract.md` defines security vocabularies: `classification`, `securityImpact`, category decisions, operational severity, `evidenceStatus` values (`not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, `blocked`), owner phases, validation status, and exception fields.
- `skills/_shared/security-guideline-catalog.md` contains the compact security taxonomy and eight current `SEC-*` guideline records, each with source IDs, mandatory flags, default complies, evidence hints, and notes.
- `scripts/validate_security_applicability.ps1` validates only `security-applicability.md` and hardcodes `schemaName: sdd.security-applicability`, `nextRecommended: design`, the eight taxonomy categories, guideline IDs, source IDs, no-impact proof, overrides, and validation metadata.

The current OpenSpec source specs intentionally codify the existing conditional model:

- `openspec/specs/sdd-security-applicability-workflow/spec.md` requires always-run independent applicability after specs and before design.
- `openspec/specs/sdd-security-design-workflow/spec.md` requires `sdd-security-design` only when `security-applicability.md` is security-impacting and preserves no-impact skipping.
- `openspec/specs/sdd-test-design-workflow/spec.md` requires test design after design and after required security design only for impacting changes.
- `openspec/specs/sdd-review-workflow/spec.md` defines the post-apply review gate and preserves security applicability/design authority.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` defines artifact resolver, review artifact, and apply-review-verify-archive routing; it currently includes security applicability/design artifacts in the established persistence surface.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` defines the catalog as the authority for security guideline IDs, taxonomy, evidence, exception fields, source coverage, and operational severity.

## Affected Areas

### Orchestration and status contracts

- `agents/sdd/sdd-orchestrator.md` — main DAG, flowchart, `/sdd-ff` route, status routing table, phase readiness, no-skip rules, no-parallel dependent planning, naming conventions, result minimums, state schema, recovery examples, and launch context table must change from `security-applicability -> design -> security-design?` to `design -> security-design -> test-design`, with `sdd-review-security` inserted after `sdd-review` and before `sdd-verify`.
- `skills/_shared/sdd-status-contract.md` — routing token mapping, status schema, artifact refs/paths/context files, artifact states, dependency states, status tokens, and phase routing order need `review-security` / `securityReviewReport` (or equivalent) plus a new security-design-mandatory model.
- `skills/_shared/persistence-contract.md` — artifact resolver, state schema, currentPhase/completedPhases/nextRecommended enums, artifactRefs, and review/verify/archive resolver rules must add mandatory `security-design` and `review-security` artifacts while deciding how to handle legacy `security-applicability` refs.
- `skills/_shared/openspec-convention.md` — active change layout and artifact path table need `security-design.md` as mandatory and a new `review-security-report.md` path after `review-report.md`.
- `README.md` — phase list and phase order currently document `sdd-security-applicability` and conditional `sdd-security-design`.

### Phase agents and skills

- `agents/sdd/sdd-security-applicability.md` and `skills/sdd-security-applicability/SKILL.md` — likely deprecated or converted into compatibility documentation only, unless retained as an internal helper concept under `sdd-security-design`.
- `agents/sdd/sdd-security-design.md` and `skills/sdd-security-design/SKILL.md` — primary redesign target. It should become mandatory after `sdd-design`, always create `security-design.md`, own classification duties, evaluate additional security design needs, and produce the full corporate security matrix.
- New `agents/sdd/sdd-review-security.md` and `skills/sdd-review-security/SKILL.md` — new mandatory phase after `sdd-review`, responsible for validating the security matrix and adding implementation evidence without duplicating the full code-review matrix.
- `skills/sdd-design/SKILL.md` — should stop requiring `security-applicability.md`, always include baseline security best-practice considerations, and route to mandatory `security-design`.
- `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-review/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` — all must consume mandatory `security-design.md`; verify/archive must also require `review-security` evidence.

### Shared security/review contracts and validators

- `skills/_shared/sdd-security-contract.md` — needs a new `sdd.security-design` schema shape that covers classification, category matrix, guideline/control matrix, statuses, expected evidence, implementation/review-security evidence, exceptions, and archive gates in one mandatory artifact.
- `skills/_shared/security-guideline-catalog.md` — likely remains the source of taxonomy and `SEC-*` IDs, but may need explicit matrix-column guidance aligned with review conventions: `Complies`/applicability answer, evidence location, observations, plus security-specific `evidenceStatus`.
- `skills/sdd-review/references/control-catalog.md` and `skills/sdd-review/references/report-template.md` — useful as a shape reference only; avoid making review the security authority.
- `scripts/validate_security_applicability.ps1` — likely deprecated, replaced, or wrapped by a new validator such as `scripts/validate_security_design.ps1`; current hardcoded schema and `nextRecommended: design` are incompatible with the target DAG.

### OpenSpec source specs likely affected

- `openspec/specs/sdd-security-applicability-workflow/spec.md` — deprecate or rewrite as compatibility/legacy behavior.
- `openspec/specs/sdd-security-design-workflow/spec.md` — rewrite from conditional artifact to mandatory classification + matrix + security-design decision artifact.
- `openspec/specs/sdd-review-workflow/spec.md` — add the boundary between general review and security review, and route `review -> review-security -> verify`.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` — add `review-security` artifact, status token, persistence, and archive readiness rules.
- `openspec/specs/sdd-test-design-workflow/spec.md` — change conditional security-design dependency to mandatory security-design dependency and consume security matrix statuses.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` — add status/matrix guidance if the matrix shape becomes catalog-level.

## Approaches

### 1. Big-bang DAG replacement

Replace `sdd-security-applicability` with mandatory `sdd-security-design` across all contracts in one coordinated change, add `sdd-review-security`, and update specs/status/persistence in a single implementation sequence.

- Pros: Clean final model quickly; no long-lived dual semantics; easier to explain once complete.
- Cons: High blast radius across orchestrator, shared contracts, phase skills, specs, README, and validators; likely exceeds the 400-line review budget; higher chance of breaking continuation/status for archived or active changes.
- Effort: High.

### 2. Compatibility-first migration with mandatory security-design as canonical

Introduce mandatory `security-design.md` as the canonical artifact and `sdd-review-security` as a new phase, while keeping legacy `security-applicability` as a recognized archived/read-only artifact during transition. Update status/persistence to normalize old artifacts into the new readiness model when reading archives, but require new changes to use the new DAG.

- Pros: Safer for archived changes; supports staged PRs; lets validators and downstream consumers migrate without losing existing evidence semantics; best fit for this repo's artifact-heavy contracts.
- Cons: Requires clear compatibility text to avoid dual authority; temporary complexity in status and archive rules.
- Effort: Medium-High.

### 3. Add security-review first, then collapse applicability later

Add `sdd-review-security` and security matrix evidence after review while leaving `sdd-security-applicability` and conditional `sdd-security-design` unchanged initially. Later, merge classification into mandatory security-design.

- Pros: Smaller first slice; improves post-apply security evidence without immediately rewriting planning DAG.
- Cons: Does not satisfy the core intent; prolongs the current split-brain model where classification is separate; creates rework because review-security would initially consume conditional artifacts.
- Effort: Medium now, higher total.

## Recommendation

Proceed with **Approach 2: Compatibility-first migration with mandatory security-design as canonical**.

Recommended target DAG for new changes:

```text
explore? -> proposal -> spec -> design -> security-design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive
```

Recommended design principles:

1. `sdd-design` always includes baseline security best practices in technical design, but it does not own the full corporate security matrix.
2. `sdd-security-design` becomes mandatory for every change and always writes `security-design.md`.
3. Classification duties move into `security-design.md`; the artifact records whether additional security controls are applicable, not whether the phase is skipped.
4. The matrix should cover every supported category/guideline with a stable status model:
   - Applicability/status shape inspired by review: `applies` / `not-applicable` and reviewer-facing evidence location + observations.
   - Evidence lifecycle from the security contract: `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, `blocked`.
   - Planning/implementation outcome statuses requested by the user: `applies`, `not-applicable`, `planned`, `implemented`, `verified`, `exception`, `blocking` can be represented as matrix-facing labels backed by the contract vocabulary.
5. `sdd-review-security` should not replace `sdd-review`; it should validate the security-design matrix after implementation, add implementation evidence, and produce a security review report consumed by verify/archive.
6. `sdd-verify` should consume both `review-report.md` and `review-security-report.md`, verify mandatory security evidence, and avoid duplicating either matrix in full.
7. `sdd-archive` should require mandatory `security-design.md` and non-blocking `review-security` evidence for new changes; legacy archived changes may remain readable under compatibility rules.

## Migration and Compatibility Concerns

- **Archived changes**: existing archived OpenSpec changes contain `security-applicability.md` and may omit `security-design.md` for no-impact changes. The new contracts should not make historical archives invalid retroactively. Archive/status readers should treat legacy folders as valid if they match the old spec version and archive report.
- **Active state schema**: `currentPhase`, `completedPhases`, `artifactRefs`, `dependencies`, and `nextRecommended` currently include `security-applicability` and no `review-security`. A state migration or compatibility reader is needed for in-flight changes.
- **Native/status tokens**: `sdd-status-contract.md` bounds the token set. Adding `review-security` requires updating token mapping and any native dispatcher assumptions. Removing `security-applicability` outright may break status outputs unless a deprecated token is still recognized for legacy state.
- **Artifact paths**: `security-applicability.md` is currently a first-class OpenSpec path. New changes should not create it, but OpenSpec conventions should document it as a legacy archived artifact or compatibility input.
- **Validator continuity**: `validate_security_applicability.ps1` is useful but schema-specific. A new validator should target `security-design.md`; the old validator can be retained for legacy artifact checks or deprecated after migration.
- **Review boundary**: `sdd-review` owns the 96-control code-review matrix. `sdd-review-security` should own the security matrix evidence update, not move those controls into general review or verify.
- **Review budget**: this redesign will touch many Markdown contracts and possibly add a validator/script and new phase files. It should be planned as stacked work units under the 400-line budget.

## Key Risks

- **Dual authority risk**: keeping legacy applicability readable while moving authority to mandatory security-design can confuse downstream phases unless contracts explicitly say new changes use only `security-design.md` for classification.
- **Matrix vocabulary drift**: review uses `Yes/No/N/A`; security contract uses lifecycle statuses. The design must map these deliberately instead of mixing incompatible columns.
- **False security confidence**: a full corporate matrix with `not-applicable` rows must require concrete evidence locations and observations, otherwise it becomes checkbox theater.
- **Dispatcher incompatibility**: adding `review-security` and removing `security-applicability` may break native/status routing unless token normalization and persisted state are updated together.
- **Archive invalidation**: archive readiness rules must distinguish legacy completed changes from new mandatory artifacts.
- **Scope size**: the likely change touches orchestrator, shared contracts, multiple phase skills, new phase files, specs, README, and validators; it should not be implemented as one large PR.

## Open Questions for Proposal

- Should the matrix rows be category-level only (8 rows) or guideline/source-level (at least 8 compact `SEC-*` rows plus source coverage details)? Recommendation: guideline-level matrix with category grouping, because archive evidence and exceptions are guideline-specific.
- What exact persisted artifact name should the new phase use: `review-security-report.md`, `security-review-report.md`, or update `security-design.md` in place? Recommendation: create `review-security-report.md` and keep `security-design.md` as planning authority.
- Should `security-applicability.md` remain accepted for active in-flight changes, or only for archived changes? Recommendation: accept it only as legacy compatibility input; new changes should not create it.

## Ready for Proposal

Yes. The redesign is coherent and should proceed to `sdd-propose` with Approach 2, scoped as a compatibility-first DAG migration that makes `security-design.md` mandatory, introduces `sdd-review-security`, and preserves legacy archived evidence without treating old artifacts as authority for new changes.
