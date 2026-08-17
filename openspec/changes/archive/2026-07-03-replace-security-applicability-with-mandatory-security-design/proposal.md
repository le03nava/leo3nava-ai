# Proposal: Mandatory Security Design for SDD Changes

## Intent

The current `sdd-security-applicability` phase runs before technical design, so classification happens without enough design context and splits security ownership across two planning artifacts. New changes should classify and plan security inside mandatory `sdd-security-design`, after `sdd-design` has established the technical approach.

## Scope

### In Scope
- Replace the active new-change DAG with: `explore? -> proposal -> spec -> design -> security-design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`.
- Make `sdd-design` always include baseline security best practices and always route to mandatory `security-design`.
- Make `sdd-security-design` always create `security-design.md` for new changes and own classification, category coverage, guideline matrix, exceptions, and archive gate expectations.
- Define the security matrix over all corporate security categories/guidelines with review-inspired fields: `Yes/No/N/A`, evidence location, observations, plus lifecycle status: `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, `blocked`.
- Add mandatory `sdd-review-security` after `sdd-review`; it validates the security matrix, adds implementation evidence, and writes `review-security-report.md`.
- Require non-blocking `review-security` evidence before `sdd-verify` and `sdd-archive` for new changes.
- Keep legacy read compatibility for archived or old `security-applicability.md` artifacts without using that phase for new changes.

### Out of Scope
- Retroactively revalidating archived changes or requiring new artifacts in old archives.
- Merging the general 96-control `sdd-review` matrix into `sdd-review-security`.
- Implementing code or contract edits during proposal.

## Capabilities

### New Capabilities
- `sdd-review-security-workflow`: mandatory post-review security evidence validation and `review-security-report.md` production.

### Modified Capabilities
- `sdd-security-applicability-workflow`: deprecate as active routing; retain legacy read-only compatibility.
- `sdd-security-design-workflow`: change from conditional design to mandatory classification, matrix, and security planning authority.
- `sdd-test-design-workflow`: consume mandatory security design before test planning.
- `sdd-review-workflow`: route from general review to security review before verify.
- `sdd-execution-persistence-contracts`: add mandatory `security-design`/`review-security` refs, status tokens, and archive gates.
- `sdd-security-guideline-catalog`: support canonical matrix/status guidance.

## Approach

Use a compatibility-first migration: make `security-design.md` canonical for new changes, add `sdd-review-security`, update orchestration/status/persistence/spec contracts, and keep old applicability artifacts readable only as legacy evidence.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `agents/sdd/*`, `skills/sdd-*` | Modified/New | DAG routing, phase contracts, new review-security phase |
| `skills/_shared/*`, `scripts/*` | Modified/New | Security schema, persistence/status, validator migration |
| `openspec/specs/*` | Modified/New | Source requirements for workflow behavior |
| `README.md` | Modified | Public phase order |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Dual authority between old and new artifacts | Med | State new changes use only `security-design.md`; legacy is read-only. |
| Dispatcher/status breakage | Med | Migrate tokens, refs, and dependencies together. |
| Matrix checkbox theater | Med | Require evidence locations, observations, statuses, and security review validation. |
| Review budget overrun | High | Plan stacked work units under the 400-line budget. |

## Migration Plan

Update specs and contracts first, then phase skills/agents, validators, README, and compatibility readers. Keep archived changes valid under their old artifact set.

## Rollback Plan

Revert the DAG/spec/skill changes to the prior conditional model, remove `review-security` routing, and continue accepting existing `security-applicability.md` artifacts as active inputs.

## Open Questions

- None blocking. Proposal assumes guideline-level matrix rows grouped by category and `review-security-report.md` as the review artifact name.

## Success Criteria

- [ ] New SDD changes no longer route through `sdd-security-applicability`.
- [ ] Every new change creates `security-design.md` and `review-security-report.md` before verify/archive.
- [ ] Archived legacy artifacts remain readable without retroactive revalidation.
