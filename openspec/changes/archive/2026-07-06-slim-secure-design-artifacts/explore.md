# Exploration: slim-secure-design-artifacts

## Current State

The active SDD security flow already treats `design.md#secure-development-design` as the new-change security authority, with standalone `security-design.md` and `security-applicability.md` limited to legacy/archive compatibility. The current contract, however, still pushes design toward carrying too much operational detail: `openspec/specs/sdd-design-workflow/spec.md` requires expected Source ID coverage in design, and `skills/sdd-design/SKILL.md` asks the design YAML to include source-row coverage metadata and groups. The shared security contract partially supports a slimmer design by allowing source rows to be summarized by corporate section when the full expanded inventory already lives in the catalog.

The authoritative source-row inventory is already centralized in `skills/_shared/security-guideline-catalog.md`: it declares 155 expanded Source IDs, section-level mappings, exact-once validation rules, compact `SEC-*` mappings, safe-evidence rules, and immutable source snapshot text. `review-security` is the phase that should expand and validate every Source ID exactly once; `test-design` should plan coverage from the design contract; `archive` should preserve final evidence and report coverage without requiring legacy artifacts or duplicating all design details.

## Affected Areas

- `skills/_shared/sdd-security-contract.md` — should sharpen the boundary: design declares compact controls plus source-row coverage references/summary; catalog owns the 155-row inventory; review-security owns expansion/validation rows.
- `skills/sdd-design/SKILL.md` — currently permits grouped source-row coverage but still includes large YAML source-row structures and can be made slimmer by requiring references, expected count, section/group summary, and downstream obligations instead of exhaustive matrices.
- `openspec/specs/sdd-design-workflow/spec.md` — current requirements say design must declare expected Source ID coverage and traceability; they should explicitly prohibit duplicating the exhaustive 155-row inventory when the catalog is available.
- `skills/sdd-test-design/SKILL.md` and `openspec/specs/sdd-test-design-workflow/spec.md` — should consume the slim design contract and plan coverage by Source ID groups/compact mapping, preserving blockers for missing mandatory evidence.
- `skills/sdd-review-security/SKILL.md` and `openspec/specs/sdd-review-security-workflow/spec.md` — already own full Source ID expansion and validation; wording should remain/strengthen this as the only phase that materializes the full row-level review matrix.
- `skills/sdd-archive/SKILL.md` and `openspec/specs/sdd-execution-persistence-contracts/spec.md` — should preserve final coverage summaries, exact count, mappings, warnings, exceptions, and links to review-security/verify evidence without copying the full matrix into design.
- `agents/sdd/sdd-design.md`, `agents/sdd/sdd-test-design.md`, `agents/sdd/sdd-review-security.md`, `agents/sdd/sdd-archive.md`, and `agents/sdd/sdd-orchestrator.md` — adapter prompt copies include the current source-row contract and must stay aligned after skill/spec changes.
- `README.md` — likely minor wording update only if the public workflow summary needs to mention the catalog/design/review-security boundary.

## Approaches

1. **Slim design by reference, expand in review-security** — Keep `design.md#secure-development-design` as the contract authority, but require only compact eight-control rows plus source-row coverage baseline: catalog snapshot ID/path, expected count `155`, section/group references, compact mapping summary, N/A policy, evidence owners, lifecycle state, and downstream traceability. `review-security-report.md` expands every Source ID exactly once.
   - Pros: Lowest cognitive load in design, preserves auditability, matches existing catalog authority, avoids 155-row duplication in every design.
   - Cons: Requires clear enough references that downstream phases cannot claim rows were omitted.
   - Effort: Medium.

2. **Move all source-row detail out of design** — Design would only classify compact `SEC-*` controls; test-design/review-security would derive all Source ID obligations directly from the catalog.
   - Pros: Smallest possible design artifact.
   - Cons: Weakens the design contract because applicability, N/A rationale, and evidence ownership are no longer visible at design time; downstream phases may rediscover scope inconsistently.
   - Effort: Medium with higher correctness risk.

3. **Keep current detailed design contract** — Continue allowing grouped or detailed source-row coverage in design.
   - Pros: Minimal change.
   - Cons: Does not solve the artifact bloat problem; design can drift into duplicating catalog/review matrices and increases reviewer burden.
   - Effort: Low but misses the goal.

## Recommendation

Use approach 1. The proposal should define a slim secure-design artifact contract: design owns classification, compact `SEC-*` decisions, evidence owners, N/A/exception policy, catalog snapshot identity, expected Source ID universe reference, and grouped coverage summary; the catalog remains the authoritative 155 Source ID inventory; test-design plans checks from those obligations; review-security expands and validates the complete matrix; archive preserves final evidence summaries and links rather than duplicating everything in design.

This keeps design as the security contract for the change without turning it into a catalog clone. It also aligns with the cognitive-doc-design principle of progressive disclosure: design gives reviewers the decision and traceability path first, while exhaustive evidence lives in the phase that validates it.

## Risks

- If the slim design contract is too vague, downstream phases may not know whether all 155 Source IDs are in scope. Mitigation: require explicit `expectedSourceIdCount: 155`, catalog snapshot ID, section/group coverage refs, and exact-once rule in design.
- If adapter prompts are not updated with skills/specs, OpenCode/Copilot agents may continue producing bloated design artifacts.
- If archive wording is too aggressive, it may duplicate review-security matrices into archive reports. Keep archive focused on summaries, verdict links, exceptions, warnings, and evidence refs.
- Existing archived changes may contain older detailed or standalone security artifacts; compatibility language must keep them readable without making them active dependencies.

## Ready for Proposal

Yes. The proposal should target contract/documentation changes only, with no runtime code expected. It should explicitly name the artifact boundary: catalog = inventory, design = slim contract/reference, test-design = coverage plan, review-security = exhaustive validation, archive = final evidence preservation.
