# Proposal: Corporate Source Row Security Validation

## Intent

Make SDD security validation audit every corporate Source ID from `skills/_shared/security-guideline-catalog.md#full-corporate-guideline-snapshot` while preserving the eight compact `SEC-*` architectural controls.

### In Scope
- Add exhaustive source-row validation with range expansion such as `1.1-1.10` before coverage checks.
- Define a source-row matrix schema: `sourceId`, corporate section, guideline text, PCI alignment if present, `mappedCompactGuidelineId`, applies `Yes|No|N/A`, complies `Yes|No|N/A`, lifecycle status, evidence location, observations, finding/blocker/warning.
- Block missing, duplicate, or unknown `sourceId`; missing compact mapping; and missing evidence when a row applies.
- Require evidence plus justification for every `N/A`; evidence must contain no secrets, PII, PAN, tokens, connection strings, or confidential values.
- Preserve traceability: Source ID -> `SEC-*` -> design/test/apply/review-security/verify/archive evidence.
- Update shared catalog/contract, phase skills, agent prompts, and source specs for design, test-design, review-security, verify, archive, and persistence compatibility.

### Out of Scope
- Replacing compact controls, duplicating the 96-control `sdd-review` matrix, requiring legacy `validate_security_design.ps1`, or dropping OpenSpec/Engram/hybrid/none compatibility.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `sdd-security-guideline-catalog`: require exhaustive corporate Source ID expansion, mapping, and safe-evidence rules.
- `sdd-design-workflow`: require source-row planning and Source ID traceability in secure design.
- `sdd-test-design-workflow`: require static/manual test planning for source-row evidence.
- `sdd-review-security-workflow`: validate every source row without duplicating `sdd-review`.
- `sdd-execution-persistence-contracts`: preserve source-row evidence through verify/archive across persistence modes.

## Approach

Add a bounded source-row operational matrix under compact controls. Design owns planning/mapping; test-design owns checks; review-security gates rows; verify consumes the verdict; archive preserves evidence. Routing: implementation blockers -> apply; missing mappings/schema/artifacts/N/A evidence -> resolve-blockers; warnings only -> verify.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared/security-guideline-catalog.md` | Modified | Source-row schema, expansion, mappings, safe evidence. |
| `skills/_shared/sdd-security-contract.md` | Modified | Schema, routing, traceability. |
| `skills/sdd-{design,test-design,review-security,verify,archive}/SKILL.md` | Modified | Phase gates. |
| `agents/sdd/` | Modified | Prompt reminders aligned to expanded contracts. |
| `openspec/specs/` | Modified | Delta specs for catalog, phase workflows, and persistence. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Large matrices increase review load. | Med | Keep compact summaries first. |
| Some Source IDs map imperfectly to compact controls. | Med | Allow one-or-more mappings without adding new compact controls. |
| Evidence may leak sensitive values. | Low | Require sanitized references only and block unsafe evidence. |

## Rollback Plan

Revert this change's proposal/spec/design/tasks/implementation files; archived historical evidence remains untouched.

## Dependencies

- Existing OpenSpec specs and active security contracts.
- No runtime test runner is available; validation evidence will be static/manual.

## Success Criteria

- [ ] Every corporate Source ID is expanded, mapped, and validated exactly once unless explicitly justified.
- [ ] Missing/duplicate/unknown IDs, missing mappings, missing required evidence, and unsupported `N/A` rows block with the correct route.
- [ ] The eight compact `SEC-*` controls and persistence-mode compatibility remain intact.
