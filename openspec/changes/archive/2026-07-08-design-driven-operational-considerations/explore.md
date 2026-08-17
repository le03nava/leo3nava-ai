# Exploration: design-driven-operational-considerations

## Current State

The current SDD operational readiness model is contract-driven and mandatory across active phases. `skills/_shared/sdd-operational-readiness-contract.md` says every SDD change must evaluate all readiness categories and defines exact marker, safe-evidence, ownership, and handoff behavior. Active phase skills consume that contract directly from design through archive.

This conflicts with the desired model: `sdd-design` should design the change first, include operational considerations only when they apply, and later phases should consume what design/test-design/tasks actually planned or implemented rather than validating against a shared readiness contract.

## Affected Areas

### Active phase skills

- `skills/sdd-design/SKILL.md` — directly requires `skills/_shared/sdd-operational-readiness-contract.md`, mandates `## Operational Readiness`, requires every readiness category, and validates against the contract. This is the primary source of contract-driven behavior.
- `skills/sdd-test-design/SKILL.md` — directly reads the readiness contract and requires readiness checks when `design.md#Operational Readiness` exists. It should instead plan checks only from design content and applicable operational considerations.
- `skills/sdd-tasks/SKILL.md` — directly reads the readiness contract and requires concrete readiness tasks from the contract plus test-design. It should derive tasks from design/test-design only.
- `skills/sdd-review/SKILL.md` — directly references the contract for general readiness review. It should review designed/implemented operational evidence when present, outside the 96-control matrix, without imposing missing categories.
- `skills/sdd-review-security/SKILL.md` — directly references the contract for leakage and placeholder safety. It should validate restricted-data leakage for operational evidence that exists, not require contract completeness.
- `skills/sdd-verify/SKILL.md` — directly verifies readiness completeness from the contract. It should verify operational considerations only when planned/designed or implemented.
- `skills/sdd-archive/SKILL.md` — directly preserves readiness status as an archive-readiness input. It should archive operational evidence that exists and preserve explicit gaps/warnings without blocking on absent non-designed categories.
- `skills/sdd-operational-doc/SKILL.md` — already manual post-archive and does not directly reference the shared readiness contract. Its behavior mostly matches the desired model: read archived evidence first, use `Pendiente de confirmar:` for applicable missing fields, and `No aplica.` for inapplicable sections.

### Shared contracts

- `skills/_shared/sdd-post-apply-gates.md` — treats readiness evidence as common post-apply context and archive readiness, including required fields and missing-readiness blockers. This must change so post-apply phases consume operational evidence only when design/test-design/tasks/apply produced it.
- `skills/_shared/sdd-security-contract.md` — references `sdd-operational-readiness-contract.md` for restricted operational data. The safe-evidence/leakage boundary may remain, but it should not depend on mandatory readiness categories or contract-driven phase obligations.
- `skills/_shared/sdd-operational-readiness-contract.md` — likely legacy/delete candidate. If retained, it should be clearly marked historical/archive compatibility or transformed into a non-authoritative reference for manual operational-doc vocabulary. Active phases should not read it.

### Active OpenSpec specs encoding mandatory readiness

- `openspec/specs/sdd-operational-readiness-workflow/spec.md` — entirely encodes the mandatory readiness workflow. This is the strongest spec conflict and should be rewritten, deprecated, or replaced.
- `openspec/specs/sdd-design-workflow/spec.md` — requires every `design.md` to include `## Operational Readiness` and mandatory fields.
- `openspec/specs/sdd-test-design-workflow/spec.md` — requires operational-readiness test planning from design and specific category coverage.
- `openspec/specs/sdd-review-workflow/spec.md` — requires review to validate readiness fields cite evidence/markers.
- `openspec/specs/sdd-review-security-workflow/spec.md` — validates operational evidence leakage and safe placeholders; should keep leakage boundaries but stop implying mandatory readiness evidence.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` — requires persistence, verify, and archive to preserve/confirm readiness completeness.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` — includes operational safe-evidence policy. This can likely remain as restricted-data/security-boundary guidance, but must avoid implying mandatory operational readiness categories.

## Approaches

1. **Design-driven operational considerations** — Remove active-phase references to `skills/_shared/sdd-operational-readiness-contract.md`; make `sdd-design` include an optional `## Operational Considerations` section only when applicable; downstream phases consume the designed/test-designed/task evidence.
   - Pros: Matches the requested model, keeps phase ownership clean, prevents mandatory category inflation.
   - Cons: Requires coordinated edits across skills and specs to avoid stale blockers.
   - Effort: Medium.

2. **Retain contract as optional guidance** — Keep the shared file but rewrite it as non-binding guidance and continue references from phases.
   - Pros: Smaller conceptual change and preserves vocabulary in one place.
   - Cons: The user explicitly wants active phases to stop using the contract; references would keep the old authority boundary alive.
   - Effort: Low/Medium, but poor fit.

3. **Delete the contract entirely** — Remove `skills/_shared/sdd-operational-readiness-contract.md` and all references.
   - Pros: Strongest guarantee that active phases cannot use it.
   - Cons: Could make archived changes harder to understand and may break references in historical artifacts or docs.
   - Effort: Medium with compatibility risk.

## Recommendation

Use Approach 1 and treat `skills/_shared/sdd-operational-readiness-contract.md` as legacy/archive compatibility unless the proposal decides to delete it after reference impact review. The active model should be:

- `sdd-design` designs the change and may include operational considerations when applicable.
- Operational consideration categories are examples, not required checklist rows: logs/errors, monitoring, administration, reprocessing/recovery, backup/retention/cleanup/generated artifacts.
- `sdd-test-design` plans checks only for considerations present in design.
- `sdd-tasks` creates tasks only from design/test-design.
- `sdd-review`, `sdd-review-security`, `sdd-verify`, and `sdd-archive` evaluate and preserve what was actually designed/implemented, including explicit pending markers when present.
- `sdd-operational-doc` remains manual post-archive and generates from archived evidence without inventing missing data.

## Risks

- Stale OpenSpec specs may continue to require mandatory readiness even after skill edits unless updated in the same change.
- Removing contract usage from post-apply gates requires careful wording so safe-evidence/leakage protections remain intact.
- Historical archived artifacts reference the old contract; deletion without compatibility wording could make old evidence harder to interpret.
- If design omits operational considerations that truly apply, review/security-review still needs authority to flag the omission based on the implemented change, but not by requiring every category by default.

## Ready for Proposal

Yes. The proposal should define a design-driven operational-considerations model, explicitly remove active-phase dependency on `skills/_shared/sdd-operational-readiness-contract.md`, update the active OpenSpec specs listed above, and decide whether the old contract file is retained as legacy/archive compatibility or deleted.
