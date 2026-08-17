# Proposal: Slim Secure Design Artifacts

## Intent

Reduce secure-design artifact bloat while preserving full auditability. `design.md#secure-development-design` should remain the slim security contract by reference; the catalog should own the 155 corporate Source ID inventory; `review-security-report.md` should own exhaustive row validation.

## Scope

### In Scope
- Define the catalog/design/test-design/review-security/archive artifact boundary.
- Update SDD workflow specs and skills so design records compact `SEC-*` decisions plus catalog snapshot, expected count, grouped coverage, evidence owners, and N/A/exception policy.
- Require `review-security` to be the only phase that materializes and validates the full 155-row Source ID matrix.
- Keep adapter prompts aligned with changed skill/spec contracts.

### Out of Scope
- Runtime application changes or new validation scripts.
- Reintroducing standalone `security-design.md` or `security-applicability.md` as active new-change artifacts.
- Migrating or rewriting archived changes that already contain detailed security artifacts.

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- `sdd-design-workflow`: slim secure-design contract by reference instead of exhaustive Source ID duplication.
- `sdd-test-design-workflow`: consume slim design coverage summaries and plan Source ID checks without requiring design to carry the full matrix.
- `sdd-review-security-workflow`: preserve/strengthen exhaustive Source ID expansion and exact-once validation ownership.
- `sdd-security-guideline-catalog`: clarify catalog ownership of the authoritative 155 Source ID inventory and mappings.
- `sdd-execution-persistence-contracts`: preserve archive/verify boundaries and final evidence references without legacy artifact dependencies.

## Approach

Use the exploration recommendation: design keeps the security decision layer slim, with compact eight-control rows and explicit references to catalog snapshot/path, `expectedSourceIdCount: 155`, section/group coverage, compact mappings, evidence owners, lifecycle state, and N/A/exception policy. Test-design plans evidence from those obligations. Review-security expands every Source ID exactly once and records the full validation matrix. Archive preserves summaries, mappings, warnings, exceptions, and links.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared/sdd-security-contract.md` | Modified | Shared artifact boundary and source-row semantics. |
| `skills/sdd-design`, `skills/sdd-test-design`, `skills/sdd-review-security`, `skills/sdd-archive` | Modified | Phase contracts consume/produce the right detail level. |
| `openspec/specs/*security*`, `sdd-design-workflow`, `sdd-test-design-workflow`, `sdd-execution-persistence-contracts` | Modified | Delta requirements for artifact ownership. |
| `agents/sdd/*.md` | Modified | Prompt copies stay aligned. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Slim design becomes too vague | Med | Require catalog identity, count 155, grouped coverage, mappings, and exact-once downstream rule. |
| Prompt/spec drift | Med | Update skill and adapter prompt copies in same work units. |
| Archive duplicates matrices | Low | Limit archive to summaries, evidence refs, warnings, and exceptions. |

## Rollback Plan

Revert the proposal/spec/skill/prompt documentation changes. Existing archived artifacts remain readable because legacy standalone security artifacts stay compatibility-only.

## Dependencies

- Existing catalog inventory in `skills/_shared/security-guideline-catalog.md`.
- Existing 155 Source ID validation behavior from `corporate-source-row-security-validation`.

## Success Criteria

- [ ] Design contracts prohibit copying the exhaustive 155-row inventory when catalog references are available.
- [ ] Review-security remains responsible for exact-once validation of all 155 Source IDs.
- [ ] Test-design, verify, and archive consume source-row evidence without weakening blockers.
