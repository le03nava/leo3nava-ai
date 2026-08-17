# Exploration: operational-readiness-evidence-flow

## Current State

The repository is an SDD workflow/prompt distribution, not an application runtime. OpenSpec is the active artifact store.

Current operational documentation exists only as a manual post-archive utility:

- `skills/sdd-operational-doc/SKILL.md`
  - Generates a Spanish operational handoff document after archive.
  - Reads archived evidence and fills missing data with `Pendiente de confirmar:`.
  - Keeps sections 1-9 and diagrams R1-R4.
  - Does not currently make operational evidence mandatory during earlier SDD phases.
- `skills/sdd-operational-doc/assets/operational-document-template.md`
  - Already includes monitoring, logs, administration, and reprocessing sections.
  - Some placeholders still need tightening so missing data uses exactly `Pendiente de confirmar:` and non-applicable data uses exactly `No aplica.`.
- `skills/sdd-design/SKILL.md`, `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-review/*`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md`
  - Already carry secure-development evidence through the lifecycle.
  - They do not yet define a mandatory Operational Readiness evidence flow for logs, monitoring, admin procedures, and reprocessing.
- `skills/_shared/sdd-post-apply-gates.md`
  - Centralizes post-apply evidence gates.
  - Good place to consume operational-readiness gates for review/verify/archive without duplicating rules.
- `skills/_shared/sdd-security-contract.md`
  - Already defines safe-evidence restrictions.
  - Needs expansion for this change's operational-data restrictions: no tokens, credentials, sensitive payloads, complete operational ID lists, generated file bytes, or production hostnames/IPs/ports/SID/service names outside explicitly provided operational documentation.

## Affected Areas

- `skills/_shared/sdd-operational-readiness-contract.md` — recommended new shared contract for mandatory Operational Readiness evidence, exact placeholder language, restricted production identifiers, and phase ownership.
- `skills/sdd-design/SKILL.md` — add required `## Operational Readiness` section to `design.md`.
- `skills/sdd-test-design/SKILL.md` — add documentary/static checklist coverage for operational readiness evidence.
- `skills/sdd-tasks/SKILL.md` — require tasks for collecting/validating operational evidence and preserving unresolved gaps.
- `skills/sdd-review/SKILL.md` — verify operational readiness handoff without altering the fixed 96-control matrix shape.
- `skills/sdd-review-security/SKILL.md` — validate safe logging/monitoring evidence and restricted operational data handling.
- `skills/sdd-verify/SKILL.md` — require final evidence checks before archive readiness.
- `skills/sdd-archive/SKILL.md` — preserve operational readiness evidence/status in archive audit trail.
- `skills/sdd-operational-doc/SKILL.md` — consume archived readiness evidence first; do not discover critical operational gaps only at doc generation.
- `skills/sdd-operational-doc/assets/operational-document-template.md` — tighten exact `Pendiente de confirmar:` / `No aplica.` usage and required logs/monitoring/admin/reprocessing fields.
- `skills/_shared/sdd-post-apply-gates.md` — add shared post-apply operational readiness gates.
- `skills/_shared/sdd-security-contract.md` — extend safe-evidence restrictions.
- `openspec/specs/*` — add or update specs for the new cross-phase operational readiness workflow.
- `README.md` — optional update explaining that operational-doc remains manual but now consumes archived readiness evidence.

## Recommended Scope

Add a dedicated Operational Readiness evidence model instead of embedding duplicated checklists in every phase.

Minimum evidence categories:

1. **Logs**
   - Location by environment.
   - Format.
   - Retention policy.
   - Diagnostic signals.
   - Security restrictions:
     - do not log tokens,
     - credentials,
     - sensitive payloads,
     - complete operational ID lists,
     - generated file bytes,
     - production hostnames/IPs/ports/SID/service names outside explicitly user-provided operational documentation.

2. **Monitoring**
   - SQL/queries for OK process.
   - SQL/queries for error process.
   - Health-checks/dashboards.
   - Nominal thresholds.
   - Trigger thresholds.
   - Evidence whether monitoring exists or remains pending.

3. **Operational Administration**
   - Quarantine/stop procedure or commands.
   - Restart/rehabilitation procedure or commands.
   - Expected validation after stop/restart.
   - Operational owner when evidenced.

4. **Reprocessing**
   - Identify candidate requests/records.
   - Correct inputs.
   - Re-invoke entrypoint.
   - Validate pending equals zero.
   - If no persistence/staging exists, say exactly `No aplica.`.

## Approaches

1. **Shared Operational Readiness contract plus phase-local consumption** — create a new shared contract and have each SDD phase consume or validate its owned part.
   - Pros: one source of truth, matches existing shared-contract pattern, avoids duplicating operational rules across phase skills.
   - Cons: touches several phase contracts and requires careful wording to keep phase ownership clear.
   - Effort: Medium/High.

2. **Inline checklist in every affected phase skill** — add operational-readiness bullets directly to each phase skill without a shared contract.
   - Pros: lower upfront file count and simple local edits.
   - Cons: high duplication, higher risk of drift, harder to enforce exact placeholders and restricted-data rules consistently.
   - Effort: Medium now, higher maintenance later.

## Recommendation

Use a shared Operational Readiness contract plus phase-local consumption.

Design should own the first `## Operational Readiness` section. Test-design should own static/documentary validation planning. Tasks should turn missing or required evidence into implementation/documentation tasks. Review should validate evidence presence and handoff. Review-security should validate safe operational evidence and restricted data handling. Verify should check evidence completeness and unresolved pending/non-applicable markers. Archive should preserve the readiness evidence summary and status. `sdd-operational-doc` should generate the final Spanish document from archived evidence, using exact pending/non-applicable language.

## Important Design Constraint

Production hostnames, IPs, ports, SID/service names, and similarly restricted operational identifiers should not be archived in ordinary SDD evidence. The SDD evidence should preserve safe summaries, presence/status, source references, and `Pendiente de confirmar:` gaps. The final operational document may include explicit production details only when the user provides them for that operational document.

## Risks

- Operational data could leak into prompts, specs, examples, or archive reports if the safe-evidence boundary is not centralized.
- Adding operational checks to the fixed 96-control review matrix could accidentally break its exact row/header contract; prefer an additional review-readiness section.
- `Pendiente de confirmar:` and `No aplica.` must be exact, including punctuation, or downstream operational-doc generation may become inconsistent.
- The operational-doc template currently has some placeholder prose that should be normalized during implementation.
- No automated test runner exists; verification will rely on static/manual contract checks.

## Ready for Proposal

Yes. Context is sufficient for proposal. The proposal should define the shared Operational Readiness contract, the phase ownership model, exact placeholder rules, restricted-data policy, and OpenSpec spec updates.
