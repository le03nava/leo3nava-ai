# Proposal: Inline Secure Development Design in `design.md`

## Intent

Remove `sdd-security-design` as an active new-change phase and make secure development design a mandatory section inside `design.md`. Keep `sdd-review-security` as the separate post-review validation gate against the corporate security catalog.

## Scope

### In Scope
- Update new-change DAG/routing so `sdd-design` routes directly to `sdd-test-design`.
- Require `design.md` to include `## Secure Development Design` with all 8 catalog rows: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, `SEC-LOG-001`.
- Define applicability, rationale, secure design decisions, controls, evidence owners, expected evidence, lifecycle status, residual risk, and exceptions in that section.
- Update downstream phases to read security obligations from `design.md` while preserving legacy `security-design.md` readability.
- Keep `sdd-review-security` after non-blocking review; it validates applicable rows and persists `review-security-report.md` with row-level evidence.

### Out of Scope
- Application runtime security changes.
- Removing archived legacy `security-design.md` artifacts.
- Replacing the corporate catalog with an external source.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-security-design-workflow`: retire standalone active phase semantics for new changes and move its artifact contract into `design.md`.
- `sdd-review-security-workflow`: validate the embedded `design.md` security section instead of standalone `security-design.md`.
- `sdd-test-design-workflow`: consume security evidence obligations from `design.md`.
- `sdd-execution-persistence-contracts`: update active artifact refs, DAG tokens, status dependencies, and legacy compatibility rules.
- `sdd-security-guideline-catalog`: align catalog usage wording with embedded design and review-security validation.

## Approach

Implement the exploration recommendation: make `sdd-design` own secure design alongside technical architecture. Update orchestrator, skills, shared contracts, validators, README, and OpenSpec specs together to avoid token/status drift. Preserve legacy compatibility for old archives, but prevent new changes from requiring or launching `sdd-security-design`.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `README.md`, `agents/sdd/sdd-orchestrator.md` | Modified | Phase list, DAG, routing, state/status recovery |
| `agents/sdd/sdd-design.md`, `skills/sdd-design/SKILL.md` | Modified | Add mandatory secure design section contract |
| `skills/_shared/*security*`, `persistence-contract.md`, `sdd-status-contract.md` | Modified | Embedded security schema and refs |
| `skills/sdd-test-design`, `sdd-tasks`, `sdd-apply`, `sdd-verify`, `sdd-archive` | Modified | Downstream input source changes |
| `skills/sdd-review-security/SKILL.md`, validator scripts | Modified | Validate embedded section and report matrix |
| `openspec/specs/*` | Modified | Source requirements updated |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Routing/status drift | High | Update DAG, state, specs, and skills in one sliced chain |
| Downstream phases read stale artifacts | Medium | Add explicit legacy-only compatibility rules |
| Larger `design.md` | Medium | Use compact table-first security section |

## Rollback Plan

Revert the OpenSpec change and restore the previous DAG: `design -> security-design -> test-design`, standalone `security-design.md` requirements, and existing validator input path.

## Dependencies

- `skills/_shared/security-guideline-catalog.md` remains the authoritative 8-guideline catalog.

## Success Criteria

- [ ] New changes no longer launch `sdd-security-design`.
- [ ] `design.md` requires the full 8-row secure development design section.
- [ ] `sdd-review-security` validates embedded design rows and writes evidence matrix.
- [ ] Legacy standalone `security-design.md` archives remain readable.
