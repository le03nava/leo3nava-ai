# Proposal: Operational Readiness Evidence Flow

## Intent

Move operational-readiness evidence into the SDD lifecycle so logs, monitoring, administration, reprocessing, ownership, and unresolved gaps are designed, validated, verified, and archived before the manual operational document is generated.

## Scope

### In Scope
- Add a shared Operational Readiness contract for evidence categories, phase ownership, safe summaries, exact placeholders, and restricted operational data.
- Update design, test-design, tasks, review, review-security, verify, and archive behavior so readiness evidence is planned, checked, and preserved.
- Update the manual `sdd-operational-doc` utility to consume archived evidence while preserving sections 1-9 and diagrams R1-R4.
- Normalize evidence states: every operational field is evidenced, `Pendiente de confirmar:`, or exact `No aplica.` with optional rationale.

### Out of Scope
- Requiring real operational data when it is unavailable; evaluation is mandatory, disclosure is not.
- Making `sdd-operational-doc` a required DAG phase.
- Storing production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, or generated file bytes in ordinary SDD evidence.

## Capabilities

### New Capabilities
- `sdd-operational-readiness-workflow`: Defines the cross-phase operational readiness evidence model, placeholders, safe-evidence boundary, archive handoff, and manual operational document consumption.

### Modified Capabilities
- `sdd-design-workflow`: Add `## Operational Readiness` planning and expected evidence strategy.
- `sdd-test-design-workflow`: Add static/documentary checks for readiness evidence and placeholder compliance.
- `sdd-review-workflow`: Validate existence, traceability, no invention, and exact placeholders without changing the 96-control matrix shape.
- `sdd-review-security-workflow`: Validate operational evidence leakage, secrets, and restricted production identifier boundaries.
- `sdd-execution-persistence-contracts`: Preserve readiness evidence through verify/archive gates and artifact refs.
- `sdd-security-guideline-catalog`: Extend safe-evidence policy for operational logs, monitoring mechanisms, and restricted operational data.

## Approach

Create one shared readiness contract consumed by phase skills. Design defines strategy and expected evidence; test-design plans static/documentary checks; tasks make collection/validation explicit; review and review-security gate quality and safety; verify confirms completeness; archive preserves status, references, and gaps. The operational document generator reads archived evidence first and writes `Pendiente de confirmar:` for unresolved fields.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/_shared/` | New/Modified | Shared readiness and post-apply/security contracts. |
| `skills/sdd-*` | Modified | Phase-local consumption for design through archive. |
| `skills/sdd-operational-doc/` | Modified | Archived-evidence-first generation and template tightening. |
| `openspec/specs/` | Modified/New | Requirements for readiness workflow and affected phases. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Sensitive operational data leaks into SDD artifacts | Med | Centralize safe-evidence restrictions and security review checks. |
| Placeholder drift | Med | Require exact markers in design/test/review/verify. |
| Review matrix contract breaks | Low | Add readiness section outside the fixed 96-control matrix. |

## Rollback Plan

Revert the OpenSpec change and related skill/template edits. Existing SDD phases and the manual operational document utility continue with their prior behavior.

## Dependencies

- Existing OpenSpec specs for SDD design, test-design, review, review-security, execution/persistence, and security catalog.

## Success Criteria

- [ ] Operational readiness evidence is planned by design and checked by test-design/tasks/review/verify/archive.
- [ ] Missing fields use exact `Pendiente de confirmar:` and non-applicable fields use exact `No aplica.`.
- [ ] Restricted operational data remains out of ordinary SDD evidence.
- [ ] Manual operational document sections 1-9 and diagrams R1-R4 are preserved.
