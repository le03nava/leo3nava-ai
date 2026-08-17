# Exploration: harden-security-applicability-contract

## Current State

`sdd-security-applicability` is an always-run phase after specs and before technical design. It writes `security-applicability.md`, classifies a change as `security-impacting` or `no-impact`, maps compact `SEC-*` guideline IDs when security applies, and routes to design. Downstream routing then uses `securityImpact` to require or skip `sdd-security-design` after technical design.

The current contract is coherent but still under-specified for audit-grade decisions:

- The artifact records `taxonomyCategories`, `applicableGuidelines`, `evidenceSummary`, `designChangingUnknowns`, and risks, but it does not require a category-by-category decision matrix.
- `no-impact` requires explicit evidence, but the schema does not separate positive no-impact evidence from mere absence of security evidence.
- The catalog has a compact `SEC-*` layer and a preserved corporate source snapshot, but compact rows do not declare one-to-many source ID mappings.
- Guideline severity/blocking behavior is mostly implicit through `mandatoryWhenApplicable`; review catalog rows already have stronger severity vocabulary that could inform this catalog.
- `rules.security-applicability` is referenced by the skill but no supported override schema is documented in `openspec/config.yaml` or specs.
- The repository has no executable test runner/linter. Validation opportunities should therefore be static Markdown/YAML validation or manual checklist validation unless a new script/tool is explicitly introduced later.

## Affected Areas

- `skills/sdd-security-applicability/SKILL.md` — primary executor contract, artifact template, decision gates, validation checklist, and config override behavior.
- `skills/_shared/sdd-security-contract.md` — shared schema should become the authority for new applicability metadata such as catalog snapshot identity, evaluated categories, decision matrix, source evidence references, and explicit no-impact evidence.
- `skills/_shared/security-guideline-catalog.md` — compact guidelines should expose catalog version/snapshot identity, source ID traceability, evidence obligation severity/blocking metadata, and clearer reviewer-facing sections.
- `openspec/specs/sdd-security-applicability-workflow/spec.md` — behavior spec should add requirements for auditable decision matrices, no-impact proof, blocking unknown prompts, and config overrides.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` — catalog spec should add source traceability, version metadata, severity/blocking semantics, and static validation expectations.
- `openspec/specs/sdd-security-design-workflow/spec.md` — downstream security design may need to consume new applicability fields without breaking no-impact routing.
- Downstream consumers (`skills/sdd-design`, `skills/sdd-security-design`, `skills/sdd-test-design`, `skills/sdd-tasks`, `skills/sdd-review`, `skills/sdd-verify`, `skills/sdd-archive`) — likely need small compatibility notes so they preserve or inspect the enriched fields instead of redefining authority.

## Approaches

1. **Schema-first hardening with compact catalog retained** — add explicit fields to the shared contract and applicability template, enrich the current catalog in place, and add spec deltas for the stronger behavior.
   - Pros: Minimal structural churn; best review fit for the 400-line budget if split into stacked slices; preserves current compact `SEC-*` IDs and archived evidence compatibility.
   - Cons: The catalog file remains large; static validation is initially a documented/manual contract unless a script is added in a later slice.
   - Effort: Medium.

2. **Catalog split plus schema hardening** — split compact checklist, source snapshot, and mapping metadata into separate reference files while updating contracts.
   - Pros: Better readability and future maintainability; easier targeted validation.
   - Cons: Higher blast radius; path/reference updates across security and review skills; greater archive compatibility risk.
   - Effort: High.

3. **Validation-tooling-first** — define or add a validator for `security-applicability.md` before broad schema/catalog edits.
   - Pros: Makes enforcement less dependent on prompt discipline; catches missing matrix/source refs early.
   - Cons: This repo currently has no executable toolchain; introducing scripts may exceed the intended SDD skill-contract scope and increase review burden.
   - Effort: Medium to High.

## Recommendation

Proceed to proposal with **Approach 1: schema-first hardening with compact catalog retained**.

Recommended change scope:

- Add required applicability artifact metadata: `catalogSnapshotId`, `catalogVersion`, `evaluatedCategories`, `decisionMatrix`, and `sourceEvidenceRefs`.
- Require a category decision matrix where every supported taxonomy category is evaluated as `applicable`, `not-applicable`, or `unknown`, with evidence and rationale.
- Separate explicit no-impact evidence from absence of evidence. A no-impact classification should be valid only when every supported category is evaluated as `not-applicable` with source evidence/rationale and no design-changing unknowns remain.
- Add compact-to-source traceability for each `SEC-*` row, likely as `Source IDs` and/or `Source Coverage` columns that map to preserved corporate source IDs such as `1.1`, `7.12`, or `14.5`.
- Add severity/blocking metadata per compact guideline or evidence obligation, reusing review catalog severity vocabulary where practical but keeping security applicability/design as the authority.
- Define supported `openspec/config.yaml` `rules.security-applicability` overrides explicitly. Suggested supported overrides: category additions disabled unless contract/spec updated, category severity overrides, extra design-changing unknown prompts, required source coverage strictness, and validator mode (`manual-only` initially).
- Improve the artifact template for reviewer scanability: lead with classification, compact matrix, blockers, guideline mapping, then detailed evidence.
- Keep the large catalog in one file for this change, but reorganize with clearer sections before considering a split in a future change.
- Specify static validation opportunities as acceptance criteria first: schema field presence, classification/securityImpact consistency, all categories evaluated, guideline IDs valid, source refs valid against the snapshot, no-impact evidence complete, blocking unknowns prevent success.

Non-goals for the first proposal:

- Do not replace the corporate source snapshot with an external policy system.
- Do not rename existing `SEC-*` IDs.
- Do not make `sdd-review` authoritative for security applicability or exceptions.
- Do not require executable validation tooling unless the proposal explicitly accepts adding repo tooling.
- Do not split the catalog file in the first slice unless review size analysis proves the in-place edit is worse.

First-slice boundary recommendation:

- PR 1: shared contract + applicability skill/template + applicability workflow spec deltas.
- PR 2: catalog metadata/source mappings/severity fields + catalog spec deltas.
- PR 3, only if needed: downstream compatibility notes and validation/checklist docs. If executable linting is chosen later, make it a separate change or final slice.

## Risks

- Mapping compact `SEC-*` rows to corporate source IDs requires judgment; inaccurate mappings would create false audit confidence.
- Adding too many required fields can make no-impact classifications noisy unless the template stays compact and reviewer-friendly.
- Severity terminology may conflict with existing corporate review severity values (`Menor`, `Media`, `Mayor`) unless normalized carefully.
- Config overrides can undermine security consistency if they permit disabling mandatory categories or source traceability.
- Any schema version bump must preserve downstream compatibility for archived artifacts and no-impact routing.

## Orchestrator-Owned Questions

- Does the user want source ID mappings to be best-effort for the current compact catalog, or must each `SEC-*` row have formally approved corporate source coverage before proposal/spec completion?
- Should severity vocabulary reuse the existing review catalog values (`Menor`, `Media`, `Mayor`) or introduce security-specific blocking levels such as `blocking`, `conditional`, and `advisory`?
- Is executable static validation in scope for this change, or should this change only define validation rules for agents/reviewers?

## Relevant Files Consulted

- `openspec/config.yaml`
- `skills/sdd-security-applicability/SKILL.md`
- `skills/_shared/sdd-security-contract.md`
- `skills/_shared/security-guideline-catalog.md`
- `openspec/specs/sdd-security-applicability-workflow/spec.md`
- `openspec/specs/sdd-security-guideline-catalog/spec.md`
- `openspec/specs/sdd-security-design-workflow/spec.md`
- `skills/sdd-security-design/SKILL.md`
- `skills/sdd-review/references/control-catalog.md`
- `skills/_shared/openspec-convention.md`

## Ready for Proposal

Yes. The change is coherent and should proceed to `sdd-propose`. The proposal should keep the scope focused on contract hardening, auditable applicability evidence, catalog traceability, and documented validation rules, with any executable validator treated as an optional later slice unless explicitly approved.
