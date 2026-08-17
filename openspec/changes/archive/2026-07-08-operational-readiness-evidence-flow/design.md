# Design: Operational Readiness Evidence Flow

## Technical Approach

Implement operational readiness as a shared SDD evidence contract plus phase-local responsibilities. The change adds `skills/_shared/sdd-operational-readiness-contract.md` as the canonical definition for readiness categories, exact placeholder behavior, safe-evidence boundaries, archive handoff, and manual operational document consumption. Phase skills then consume only the part they own: design plans expected evidence, test-design plans checks, tasks make collection work explicit, review checks traceability, review-security checks leakage, verify confirms completeness, and archive preserves the handoff.

This matches the proposal and specs by making readiness evaluation mandatory without requiring disclosure of real production details. The manual `sdd-operational-doc` remains a post-archive utility: it reads archived evidence first, keeps sections 1-9 and diagrams R1-R4, and writes unresolved fields with the exact `Pendiente de confirmar:` marker.

## Architecture Decisions

### Decision: Centralize readiness semantics in a shared contract

**Choice**: Create `skills/_shared/sdd-operational-readiness-contract.md` and reference it from affected phase skills.
**Alternatives considered**: Duplicate readiness rules in each phase skill.
**Rationale**: A shared contract prevents placeholder drift and keeps safe-evidence policy consistent while allowing each phase to own only its local gate.

### Decision: Preserve the SDD DAG and keep the operational document manual

**Choice**: Design readiness evidence through archive, then let `sdd-operational-doc` consume archived evidence manually after archive.
**Alternatives considered**: Add `sdd-operational-doc` as a required DAG phase.
**Rationale**: The specs explicitly require archive completion without operational-doc execution. The manual utility is a final documentation generator, not a lifecycle gate.

### Decision: Use exact human markers instead of machine schemas

**Choice**: Operational fields must contain safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.` with optional rationale.
**Alternatives considered**: Add YAML/JSON readiness blocks or machine-readable states.
**Rationale**: Existing SDD artifacts are Markdown instruction contracts. Exact markers are simple to validate statically and safe for downstream Spanish operational documentation.

### Decision: Separate quality review from leakage review

**Choice**: `sdd-review` validates existence, traceability, no invention, and exact placeholders outside the fixed 96-control matrix; `sdd-review-security` validates restricted operational data leakage.
**Alternatives considered**: Add readiness rows to the 96-control matrix or move all readiness checks into security review.
**Rationale**: The specs require preserving matrix shape and distinct responsibilities. Quality and safety checks need different evidence and routing.

## Data Flow

```text
proposal/specs
    -> design.md#Operational Readiness
    -> test-design.md readiness checks
    -> tasks.md collection/validation work
    -> apply evidence and changed files
    -> review-report.md traceability handoff
    -> review-security-report.md leakage validation
    -> verify-report.md completeness verdict
    -> archive audit trail
    -> manual sdd-operational-doc output
```

Operational values move as safe summaries and references through SDD artifacts. Restricted values do not move through ordinary SDD evidence. If the final operational document needs production hostnames, IPs, ports, SID/service names, credentials, tokens, sensitive payloads, full ID lists, or generated file bytes, those values must be explicitly provided by the user for the final document and must not be backfilled into design, tasks, review, verify, or archive evidence.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/_shared/sdd-operational-readiness-contract.md` | Create | Shared readiness contract covering categories, phase ownership, markers, restricted-data policy, safe summaries, archive handoff, and manual document boundary. |
| `skills/_shared/sdd-security-contract.md` | Modify | Extend safe-evidence policy to reference restricted operational data and the readiness contract without moving security matrices into design. |
| `skills/_shared/sdd-post-apply-gates.md` | Modify | Add readiness consumption expectations for review, review-security, verify, and archive gates. |
| `skills/sdd-design/SKILL.md` | Modify | Require `## Operational Readiness` in `design.md` and define design-owned readiness strategy fields. |
| `skills/sdd-test-design/SKILL.md` | Modify | Require static/documentary readiness checks, exact marker validation, monitoring mechanism coverage, and safe-evidence checks. |
| `skills/sdd-tasks/SKILL.md` | Modify | Require tasks for readiness evidence collection, placeholder validation, archive handoff, and unavailable-tooling notes. |
| `skills/sdd-review/SKILL.md` | Modify | Add readiness review outside the 96-control matrix and hand off evidence locations/gaps to security review. |
| `skills/sdd-review-security/SKILL.md` | Modify | Add operational leakage validation for restricted production identifiers, secrets, payloads, full ID lists, and generated bytes. |
| `skills/sdd-verify/SKILL.md` | Modify | Confirm each readiness field has safe evidence, exact `Pendiente de confirmar:`, or exact `No aplica.` before archive routing. |
| `skills/sdd-archive/SKILL.md` | Modify | Preserve readiness status, refs, unresolved gaps, and final-document handoff boundaries in archive evidence. |
| `skills/sdd-operational-doc/SKILL.md` | Modify | Make archived-readiness evidence the first input source while preserving manual post-archive activation. |
| `skills/sdd-operational-doc/assets/operational-document-template.md` | Modify | Tighten sections 1-9 and diagrams R1-R4 so missing/inapplicable fields use exact markers. |
| `openspec/specs/**/spec.md` | Modify/Create during archive | Persist accepted requirements for the readiness workflow and affected SDD phase capabilities. |

## Interfaces / Contracts

The new shared contract should define these Markdown sections:

```markdown
# SDD Operational Readiness Contract

## Purpose
## Evidence Categories
## Exact Marker Contract
## Restricted Operational Data Boundary
## Phase Ownership Boundary
## Design Artifact Section Contract
## Test/Task/Review/Verify/Archive Handoff
## Manual Operational Document Boundary
## Safe Evidence Examples
```

Minimum readiness categories:

- Product ownership and support ownership.
- Logs and error evidence.
- Monitoring mechanisms, including dashboards, alerts, jobs, traces, scripts, documented manual checks, or SQL where appropriate; never SQL-only by default.
- Administration operations: quarantine/stop, restart/resume, periodic activities.
- Reprocessing and recovery.
- Backup, retention, cleanup, and generated artifacts.
- Final operational document inputs and unresolved gaps.

Marker behavior is exact:

- Use `Pendiente de confirmar:` when the field applies but safe evidence is unavailable, incomplete, or must be supplied later by the user.
- Use `No aplica.` when the field is genuinely out of scope. Optional rationale may follow, but the marker text and punctuation must remain exact.
- Do not replace either marker with synonyms, lowercase variants, missing colon, or translated alternatives.

## Operational Readiness

### Strategy

Operational readiness becomes a required design concern for every SDD change. Each design artifact must include `## Operational Readiness` and evaluate logs, monitoring mechanisms, administration, reprocessing, ownership, final documentation inputs, and unresolved gaps. Evaluation is mandatory; real production data disclosure is not.

### Expected Evidence by Phase

- `sdd-design`: records the strategy, expected safe evidence, exact unresolved markers, non-applicable rationale, restricted-data boundary, and downstream evidence owners.
- `sdd-test-design`: plans static, documentary, manual, or automated checks for marker exactness, evidence traceability, monitoring mechanism coverage, and restricted-data absence. Missing runtime tooling routes to static/documentary checks, not to a pass.
- `sdd-tasks`: creates concrete collection and validation tasks for readiness fields, unresolved gaps, safe-evidence checks, and archive handoff.
- `sdd-review`: validates existence, traceability, no-invention behavior, and exact marker usage in a readiness section outside the fixed 96-control matrix.
- `sdd-review-security`: validates leakage boundaries for production hostnames, IPs, ports, SID/service names, credentials, tokens, sensitive payloads, full ID lists, generated file bytes, and equivalent restricted operational data.
- `sdd-verify`: confirms every readiness field is evidenced, `Pendiente de confirmar:`, or `No aplica.` and preserves unavailable-tooling notes.
- `sdd-archive`: preserves readiness status, evidence refs, unresolved gaps, warning carry-forward, and manual-document handoff notes.
- `sdd-operational-doc`: reads archived readiness evidence after archive; it does not create archive readiness and does not become a required DAG phase.

### Safe-Evidence Boundary

Ordinary SDD evidence may cite paths, artifact sections, sanitized summaries, marker states, command summaries, redacted placeholders, and user-provided final-document references. It must not include production hostnames, IPs, ports, SID/service names, credentials, tokens, private keys, connection strings, sensitive payloads, full ID lists, generated file bytes, PAN, PII, or confidential values. If such details are needed for the final operational document, the SDD evidence records `Pendiente de confirmar:` and the manual utility may use explicit user-provided values only in the generated final document.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static/documentary | Shared contract sections, phase skill references, exact markers, and operational-readiness section requirements. | Manual/static inspection because `openspec/config.yaml` reports no test runner, linter, type checker, formatter, or coverage command. |
| Static/documentary | Review matrix preservation and readiness section placement. | Inspect `sdd-review` output contract to ensure readiness checks are outside the fixed 96-control matrix. |
| Static/documentary | Security leakage policy. | Inspect `sdd-review-security`, security contract, and readiness contract for restricted operational data prohibitions and safe placeholders. |
| Manual | `sdd-operational-doc` post-archive behavior. | Check that it consumes archive evidence, preserves sections 1-9 and diagrams R1-R4, and does not invent missing data. |

No executable runtime, build, lint, type-check, format, or coverage tooling is available in this repository. Downstream phases must report that explicitly and use static/manual evidence.

## Migration / Rollout

No runtime migration required. This is a Markdown contract and skill-instruction change. Rollout is by applying the shared contract and phase skill updates together so downstream phase contracts do not reference a missing readiness contract.

## Open Questions

None.

## Secure Development Design

### Classification and Changed Surface

This change is security-impacting. It changes SDD instructions that govern how operational evidence, logs, monitoring details, administrative procedures, reprocessing details, archive evidence, and final operational documentation are collected and reviewed. The repository is an AI agent/skill distribution, not an application runtime, so it does not directly add authentication flows, sessions, database queries, runtime file uploads, or application authorization logic. The security impact is on evidence handling and leakage prevention in Markdown artifacts and generated operational documentation.

The changed surface includes shared SDD contracts, phase skill instructions, post-apply gates, the security contract, the security review workflow, and the manual operational document utility. The relevant catalog context is the compact safe-evidence categories for sensitive data, secrets, sensitive logging, and files/generated artifacts. Omitted categories such as authentication, sessions, permissions/access-control, and database-access are left as reviewable omissions for `sdd-review-security`, which owns compact/source-row validation and any `N/A` decisions.

### Sensitive Data and Operational Identifier Rules

Development must treat production hostnames, IPs, ports, SID/service names, sensitive payloads, full ID lists, PAN, PII, confidential values, and generated file bytes as unsafe for ordinary SDD evidence. The control is to preserve safe summaries, artifact refs, section anchors, and exact placeholders instead of copying raw operational details. Evidence owners are `sdd-design` for the planned boundary, `sdd-test-design` for static/documentary checks, `sdd-review` for traceability/no-invention checks, `sdd-review-security` for leakage rejection, `sdd-verify` for completeness confirmation, and `sdd-archive` for preserving unresolved gaps. Residual risk is that final operational documents may legitimately need environment-specific details; mitigation is to allow those values only when explicitly user-provided for the final manual document and never backfill them into SDD evidence.

### Secrets Rules

Credentials, tokens, private keys, connection strings, and secret values must never appear in SDD artifacts, code examples, tests, fixtures, review reports, verify reports, archive summaries, or generated evidence snippets. The implementation must describe secret ownership and storage mechanisms only through names, paths, or redacted placeholders. Evidence owners are `sdd-review-security` for unsafe-evidence blocking and `sdd-archive` for preserving safe refs without copying values. No exception is planned; any exception would require complete approval details before archive readiness.

### Sensitive Logging and Monitoring Rules

Readiness evidence may describe log categories, monitoring mechanisms, dashboards, alerts, jobs, traces, scripts, or documented manual checks, but must not copy raw logs, stack traces with sensitive payloads, user identifiers, credentials, or confidential operational context. Monitoring evidence must be mechanism-oriented and not limited to SQL-only checks. Evidence owners are `sdd-design` and `sdd-test-design` for planned monitoring coverage, `sdd-review` for traceability, and `sdd-review-security` for leakage validation. Residual risk is over-collection in examples; mitigation is to require sanitized summaries and exact pending markers where safe evidence is unavailable.

### Files and Generated Artifact Rules

Generated file bytes, full exported file contents, and environment-specific file payloads must not be stored as ordinary SDD evidence. The safe pattern is to cite artifact paths, filenames when non-sensitive, checksums only when safe, file type summaries, or `Pendiente de confirmar:` for missing final-document inputs. Evidence owners are `sdd-test-design` for planned static checks, `sdd-review-security` for leakage rejection, and `sdd-archive` for preserving references rather than bytes. Residual risk is accidental inclusion of generated outputs during manual documentation; mitigation is explicit utility guidance to summarize and link rather than copy bytes.

### Exception and Evidence Policy

No security exceptions are planned. If an exception is later needed, it must include approver, approval date, accepted-risk rationale, mitigation or follow-up, and the exact evidence gap. Safe evidence must cite paths, sections, sanitized command summaries, redacted placeholders, or unresolved markers only. Design and test-design must not emit security YAML, JSON, schemas, compact matrices, Source ID matrices, machine-readable applicability fields, or all-row `N/A` bookkeeping; `review-security-report.md` owns those review-time decisions.
