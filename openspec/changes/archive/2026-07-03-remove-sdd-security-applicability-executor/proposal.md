# Proposal: Remove SDD Security Applicability Executor

## Intent

Remove the repo-local launchable `sdd-security-applicability` executor and skill. The active SDD DAG already uses mandatory `sdd-security-design`; keeping a runnable legacy applicability phase creates conflicting guidance and may cause sync operations to publish obsolete prompts/skills as if they were active.

## Scope

### In Scope
- Delete `agents/sdd/sdd-security-applicability.md`.
- Delete `skills/sdd-security-applicability/SKILL.md` and its empty directory if applicable.
- Update active contracts/docs so `security-applicability.md` is legacy data-only compatibility, not a runnable phase.
- Preserve legacy `security-applicability.md` readability and archive-only validation support.
- Document the stale-copy caveat for global opencode destinations without changing broad sync deletion semantics.

### Out of Scope
- Editing archived OpenSpec change folders except normal archive mechanics for this change.
- Removing `scripts/validate_security_applicability.ps1` unless a later change intentionally migrates archive validation.
- Deleting stale global `%USERPROFILE%` opencode copies.
- Redesigning sync scripts to delete destination files with no current repo source.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-security-applicability-workflow`: make applicability a legacy artifact readability/data compatibility contract only.
- `sdd-execution-persistence-contracts`: ensure state, resolver, and routing contracts do not expose an active applicability phase dependency.
- `sdd-security-design-workflow`: preserve mandatory security-design classification as the new-change authority.
- `sdd-review-workflow`: treat legacy applicability as optional/archive evidence, not a default new-change input.

## Approach

Delete the obsolete repo-local executor/skill, then update active source contracts that still imply launchability. Keep archive compatibility language, `securityApplicability` state fields as legacy/read-only data, OpenSpec legacy path references, and the archive-only validator.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `agents/sdd/sdd-security-applicability.md` | Removed | Delete launchable executor prompt. |
| `skills/sdd-security-applicability/SKILL.md` | Removed | Delete deprecated executor skill. |
| `agents/sdd/sdd-orchestrator.md` | Modified | Remove active launch mapping/context. |
| `skills/_shared/*` | Modified | Reword status, persistence, OpenSpec, and security contracts around data-only legacy compatibility. |
| `openspec/specs/*` | Modified | Update delta specs for affected SDD workflow capabilities. |
| `skills/sdd-review/references/report-template.md` | Modified | Remove or label applicability as legacy/optional input. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Stale global copies remain launchable | Medium | Document repo-only scope and stale-copy caveat. |
| Archived mentions confuse searches | Medium | Preserve archives but clarify active contracts. |
| Validator name implies active phase | Low | Keep it explicitly archive-only. |

## Rollback Plan

Revert the proposal implementation commit to restore the repo-local executor/skill and previous contract wording. Archived artifacts are not modified, so no archive rollback is required.

## Dependencies

- Existing exploration artifact: `openspec/changes/remove-sdd-security-applicability-executor/explore.md`.

## Success Criteria

- [ ] Repo no longer contains launchable `sdd-security-applicability` agent or skill sources.
- [ ] Active contracts route new changes through mandatory `sdd-security-design` only.
- [ ] Legacy applicability artifacts remain readable as archive data.
- [ ] Archive-only validator remains clearly separated from new-change routing.
