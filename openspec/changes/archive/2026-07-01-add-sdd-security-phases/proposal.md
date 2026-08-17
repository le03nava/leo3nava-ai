# Proposal: Add SDD Security Phases

## Intent

Add security applicability triage and conditional security design so SDD changes are mapped to corporate guidelines and blocked from archive when mandatory evidence is missing without an approved exception.

## Scope

### In Scope
- Add `sdd-security-applicability`, producing `security-applicability.md` with guideline mapping or explicit no-impact evidence.
- Add conditional `sdd-security-design`, producing `security-design.md` only for security-impacting changes.
- Add an in-repo guideline catalog snapshot and compact taxonomy for phase prompts.
- Update routing, status, persistence, downstream phase, and archive contracts.

### Out of Scope
- Migrating guidelines to an external official versioned document.
- Full policy content beyond the initial user-provided snapshot.
- Changing repository runtime behavior; this is an AI agent/skill workflow change.

## Capabilities

### New Capabilities
- `sdd-security-applicability-workflow`: always-run impact classification, blocking rules, artifact contract, and token mapping.
- `sdd-security-design-workflow`: conditional translation of applicable guidelines into design controls and evidence expectations.
- `sdd-security-guideline-catalog`: in-repo catalog snapshot, taxonomy, and mandatory evidence model.

### Modified Capabilities
- `sdd-test-design-workflow`: test design must follow required security design and include security-control evidence.

## Approach

Insert `security-applicability` after spec and `security-design` after design when applicable. Applicability blocks only when missing information could change security design decisions across auth, sessions, sensitive data/PAN, secrets, permissions, files, DB, or sensitive logging. Minor evidence gaps continue as security-design risks. Archive blocks unevidenced mandatory applicable guidelines unless an approved exception records approver, guideline, rationale, and mitigation.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `agents/sdd/` | New/Modified | Add security phase agents and update orchestrator routing. |
| `skills/` | New/Modified | Add security phase skills and update existing SDD phase contracts. |
| `skills/_shared/` | Modified | Update status, persistence, OpenSpec, and phase common contracts. |
| `openspec/` | Modified | Add new capability specs and update test-design workflow. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Native tooling lacks new tokens | Med | Define bounded native/status token mapping before apply. |
| Guideline catalog increases prompt size | Med | Use compact taxonomy plus references, not full prompt injection. |
| Always-run applicability adds overhead | Low | Keep no-impact artifact concise and evidence-based. |

## Rollback Plan

Revert the new change artifacts and restore previous SDD routing/status contracts. Archived prior specs remain unchanged until this change is archived.

## Dependencies

- Initial corporate security guideline text supplied by the user.
- Existing SDD test-design phase contract.

## Success Criteria

- [ ] Every SDD change records security applicability or no-impact evidence.
- [ ] Security design runs only for applicable changes.
- [ ] Archive blocks missing mandatory security evidence unless an approved exception exists.
