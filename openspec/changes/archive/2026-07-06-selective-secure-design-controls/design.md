# Design: Selective Secure Design Controls

## Technical Approach

Regenerate the design around the revised boundary: `design.md#secure-development-design` is a narrative development-rules artifact, while `review-security-report.md` owns machine-readable schema, compact-control matrices, Source ID matrices, exhaustive applicability decisions, and all-row `N/A` bookkeeping.

The change updates SDD instruction contracts, shared security wording, adapter prompts, and OpenSpec source specs. It does not change runtime application code, APIs, databases, authentication flows, session handling, file handling, or observability implementation. Because this repository has no configured runner, build, linter, type checker, formatter, or coverage command, validation is static/manual read-back against Markdown contracts and specs.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Narrative secure design | Make secure development design a human-readable changed-surface and category-rules section only. | Keep YAML, schema fields, compact matrices, Source ID matrices, or all-row `N/A` evidence in design. | The revised specs require design to reduce noise and avoid machine-readable security matrices. |
| Exhaustive validation owner | Put compact-control, Source ID, missed-category, and `N/A` validation in `review-security-report.md`. | Split exhaustive validation across design, test-design, and review-security. | A single report owner prevents drift and preserves audit strength. |
| Test-design consumption | Let `sdd-test-design` consume applicable narrative category rules and changed-surface rationale. | Require test-design to parse design YAML or infer omitted rows. | The test-design specs now reject dependencies on design schemas or exhaustive `N/A` rows. |
| Compatibility | Keep archived exhaustive designs readable without migrating them. | Rewrite historical artifacts. | Archives are audit evidence; the new boundary applies to active new changes. |

## Data Flow

```text
proposal + revised specs + catalog context
        │
        ▼
sdd-design writes narrative changed-surface rules
        │
        ▼
sdd-test-design plans checks from applicable narrative rules
        │
        ▼
apply evidence + general review
        │
        ▼
sdd-review-security owns exhaustive schema/matrices, omissions, and blockers
        │
        ▼
verify and archive preserve report evidence
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modify | Define artifact ownership so design owns narrative category rules and review-security owns machine-readable exhaustive validation. |
| `skills/sdd-design/SKILL.md` | Modify | Remove design-time YAML/schema/matrix/all-row requirements and require changed-surface classification plus applicable category rules. |
| `skills/sdd-review-security/SKILL.md` | Modify | Require exhaustive compact-control and Source ID validation, missed-category blockers, report-owned schema, safe evidence, and `N/A` decisions. |
| `skills/sdd-test-design/SKILL.md` | Modify | Consume applicable narrative rules and changed-surface context without requiring design schema or matrices. |
| `agents/sdd/sdd-design.md` | Modify | Align adapter instructions with narrative secure-design output. |
| `agents/sdd/sdd-review-security.md` | Modify | Align adapter instructions with exhaustive security report ownership. |
| `agents/sdd/sdd-test-design.md` | Modify | Align adapter instructions with narrative-rule test planning. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modify | Sync accepted source requirements for narrative secure design. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modify | Sync accepted source requirements for report-owned matrices and blockers. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modify | Preserve catalog authority while preventing design from owning schema/matrix obligations. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Modify | Sync applicable-rule consumption and unavailable-tooling behavior. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modify | Preserve active narrative design plus review-security report semantics across storage modes. |

## Interfaces / Contracts

The active contract for new changes is prose-first:

- `design.md#secure-development-design` must classify the changed surface and list only applicable security category rules.
- Design must not contain security YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, or all-row `N/A` bookkeeping.
- Omitted categories remain reviewable omissions. They are not design-time passing rows.
- `test-design.md` must map applicable narrative category rules to static, manual, or automated checks and must report unavailable automation explicitly.
- `review-security-report.md` must own the report schema, exhaustive compact-control validation, exhaustive Source ID validation, non-applicable decisions, missed applicable controls, blockers, exceptions, and next recommendation.
- Historical exhaustive design artifacts remain readable as compatibility evidence but are not the output shape for new active changes.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static contract | Skills and shared contract state the same narrative-design / exhaustive-review boundary. | Manual diff/read-back inspection. |
| Static source specs | Accepted deltas are synchronized into source OpenSpec specs. | Manual comparison of change specs and source specs. |
| Adapter prompts | `agents/sdd/*.md` mirrors match skill contracts. | Manual prompt inspection. |
| Security artifact shape | Secure development design is narrative only and review-security owns matrices. | Static review of this design and downstream artifacts. |
| Runtime | Not applicable. | `openspec/config.yaml` reports no configured test runner, build, linter, type checker, formatter, or coverage command. |

## Migration / Rollout

No data migration required. This is a documentation and instruction-contract change. Existing archived exhaustive security designs remain compatible read-only evidence. New active changes use narrative `design.md#secure-development-design` plus machine-readable `review-security-report.md`.

## Open Questions

None.

## Secure Development Design

### Classification and Changed Surface

This change is security-impacting because it changes how SDD security obligations are planned, tested, reviewed, verified, and archived. The changed surface is limited to Markdown instruction contracts, OpenSpec source specs, and adapter prompts for SDD phases. It does not modify runtime authentication, session management, file handling, database access, application APIs, credential storage, or production logging behavior.

The applicable security categories for development rules are Sensitive Data/PAN, Secrets, Permissions/Access Control, and Sensitive Logging. Authentication, Sessions, Files, and Database Access are mentioned here only as changed-surface rationale: those runtime behaviors are not implemented or modified by this change, and exhaustive omission validation belongs to `review-security-report.md`, not design.

Catalog context remains available through `skills/_shared/security-guideline-catalog.md` for category names, safe-evidence expectations, identifiers, and later security-review validation. Design must not copy catalog matrices or Source ID inventories.

### Sensitive Data/PAN Rules

Development must preserve review-safe evidence rules anywhere the SDD workflow discusses sensitive data, PAN, PII, or confidential values. Updated skills, shared contracts, specs, adapters, test design guidance, review-security guidance, verify guidance, and archive guidance must require evidence by path, section, summary, command summary, sanitized example, or redacted placeholder.

Development must not require or encourage agents to paste raw PAN, PII, credentials, tokens, private keys, connection strings, or confidential payloads into design, test design, review reports, verify reports, archive reports, or memory observations. If a phase needs to prove a sensitive-data rule, it must cite the artifact location and describe the control without reproducing the sensitive value.

Evidence owners are `test-design` for planned static/manual checks, `apply` for changed-file references, `review-security` for exhaustive validation and unsafe-evidence blockers, `verify` for preserving non-blocking evidence, and `archive` for retaining the final audit trail. Residual risk is that a narrative design could omit an applicable sensitive-data rule; `review-security` must block missed applicable categories or unsafe evidence.

### Secrets Rules

Development must keep the no-secret policy explicit across all changed SDD contracts. Prompts and skills must instruct agents never to commit, echo, log, or reproduce secret values, API keys, passwords, certificates, tokens, encryption keys, private keys, or connection strings as evidence. Required evidence must use secret/config names, file paths, owner notes, redacted placeholders, or sanitized summaries only.

The design and test-design phases must not demand secret values to prove compliance. The review-security phase must treat unsafe secret evidence as blocking and route the change to remediation or blocker resolution before verify/archive. Exceptions for missing mandatory secret evidence must remain complete and visible through downstream phases; incomplete exceptions are not acceptable development output.

Evidence owners are `apply` for contract and prompt changes that preserve no-secret wording, `review-security` for unsafe-evidence rejection, `verify` for checking the report is non-blocking, and `archive` for preserving safe evidence. Residual risk is limited to prompt/spec drift; test design and review-security must explicitly check for drift.

### Permissions / Access Control Rules

Development must preserve denial-by-default workflow routing for security evidence. Missed applicable categories, missing mandatory evidence, malformed security report content, unsupported `N/A` decisions in the report, unsafe evidence, missing artifacts, and incomplete exceptions must not silently pass to verify or archive.

The design phase must not attempt to prove every non-applicable row. Instead, it must provide enough changed-surface context for downstream phases, and review-security must validate whether omissions are safe. If proposal, specs, changed files, test design, apply evidence, or review findings show that an omitted category actually applies, review-security must report a missed applicable category/control and block.

Evidence owners are `test-design` for planned checks covering omission handling and blocker routing, `apply` for implemented contract wording, `review-security` for authoritative blocker decisions, `verify` for requiring non-blocking security review before archive, and `archive` for preserving final routing evidence. Residual risk is that narrative wording becomes too terse; the mitigation is mandatory missed-category validation in review-security.

### Sensitive Logging Rules

Development must treat SDD review, security review, verify, archive, and memory/report artifacts as audit logs that need useful evidence without unsafe payloads. Reports must preserve warnings, unavailable-tooling notes, changed-file references, and manual inspection summaries, but must not include raw secrets, PAN, PII, tokens, credentials, connection strings, private keys, or unnecessary confidential values.

When runtime tooling is unavailable, phases must say so explicitly instead of inventing passing evidence. Static/manual evidence must cite readable artifacts and describe what was inspected. Security review must reject unsafe report evidence and keep warning-only findings visible to verify and archive when mandatory evidence is otherwise complete.

Evidence owners are `apply` for report-format and contract wording, `review-security` for safe audit evidence validation, `verify` for preserving non-blocking warnings, and `archive` for final evidence retention. Residual risk is that concise reports omit warning context; downstream verification and archive must preserve non-blocking warnings rather than dropping them.

### Exception and Evidence Policy

No approved exceptions are planned for this design. If implementation discovers missing mandatory security evidence, the exception must name the affected category, approver, approval date, accepted-risk rationale, mitigation or follow-up, and evidence gap. Incomplete exceptions must block.

All evidence for this change must be review-safe and must remain in English technical artifacts. The expected validation method is static/manual read-back of changed Markdown contracts, source specs, adapter prompts, this design, test design, apply evidence, general review, security review, verify, and archive artifacts.
