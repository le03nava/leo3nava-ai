# Test Design: Design-Driven Operational Considerations

## Overview

This change updates repository instruction contracts, not application runtime behavior. Testing is therefore static and manual: validate that the old mandatory operational-readiness contract is deleted, active skills/specs no longer consume it, operational considerations are conditional and design-owned, downstream phases consume only actual artifact evidence, and restricted-data/safe-evidence protections remain in place. `openspec/config.yaml` reports no test runner, build command, linter, type checker, formatter, or coverage command, so unavailable tooling must be carried as explicit evidence rather than treated as passing evidence.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/design-driven-operational-considerations/proposal.md` | Scope, deletion decision, non-goals, risks, and success criteria. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-design-workflow/spec.md` | Conditional `## Operational Considerations`, safe design evidence, and no mandatory completeness gates. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-test-design-workflow/spec.md` | Checks derived only from design operational evidence and safe-evidence validation when evidence exists. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-review-workflow/spec.md` | General review validation for present/planned operational evidence without altering the 96-control matrix. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-review-security-workflow/spec.md` | Leakage review for existing operational evidence and exact placeholder handling. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-execution-persistence-contracts/spec.md` | Verify/archive persistence of actual operational evidence and manual operational-doc boundary. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-operational-readiness-workflow/spec.md` | Replacement model: design-driven operational considerations and removal of mandatory readiness evaluation. |
| Spec | `openspec/changes/design-driven-operational-considerations/specs/sdd-security-guideline-catalog/spec.md` | Safe-evidence policy and restricted operational data classification without mandatory categories. |
| Design | `openspec/changes/design-driven-operational-considerations/design.md` | Architecture decisions, affected files, design-owned applicability, evidence flow, testing strategy, and migration/rollout. |
| Secure Development Design | `openspec/changes/design-driven-operational-considerations/design.md#secure-development-design` | Security-impacting changed surface, safe-evidence policy, restricted operational data rules, final-document boundary, residual risks, and evidence owners. |
| Testing Capabilities | `openspec/config.yaml#testing` | Confirms no runtime/build/coverage/lint/type/format tooling is available; plan static/manual evidence only. |

## Source ID Coverage Baseline

Corporate source-row expansion is not planned in this phase. The changed surface is Markdown instruction contracts under `skills/`, shared safe-evidence wording, and active OpenSpec specs. Test design consumes only the secure-development narrative in `design.md`: restricted operational data rules, sensitive logging/generated artifact rules, final documentation boundary rules, no-exception policy, owner phases, and safe evidence expectations. Exhaustive Source ID matrices, omitted-row validation, and all-row `N/A` bookkeeping remain owned by `review-security-report.md` if that review applies.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Proposal success criteria; Design file changes | Verify `skills/_shared/sdd-operational-readiness-contract.md` is deleted. | static | mandatory | File absence from the active workspace, e.g. `Test-Path -LiteralPath "skills/_shared/sdd-operational-readiness-contract.md"` returns false, plus changed-file evidence showing deletion. | The deleted path is the subject of validation only, not a workflow input. |
| TD-002 | Proposal risk: stale references; Design interface contract | Search active phase skills and shared contracts for references to `skills/_shared/sdd-operational-readiness-contract.md`. | static | mandatory | Repository search output scoped to active `skills/` files shows no references outside historical/archive artifacts or this change's planning evidence. | Archive/history references are allowed only as audit trail, not active instructions. |
| TD-003 | `sdd-design-workflow` spec; Design decisions | Verify `skills/sdd-design/SKILL.md` uses conditional `## Operational Considerations` and no longer requires mandatory `## Operational Readiness` completeness gates. | static | mandatory | Static inspection of `skills/sdd-design/SKILL.md` showing design-owned conditional wording and absence of mandatory all-category readiness requirements. | Must keep `## Secure Development Design` requirements intact. |
| TD-004 | `sdd-test-design-workflow` spec | Verify `skills/sdd-test-design/SKILL.md` derives operational checks only from `design.md#Operational Considerations` or equivalent design evidence. | static | mandatory | Static inspection showing no dependency on the deleted contract and no creation of mandatory checks when design omits or marks considerations `No aplica.`. | This test-design artifact follows that target behavior despite the installed skill's stale wording. |
| TD-005 | `sdd-operational-readiness-workflow` phase ownership model | Verify `skills/sdd-tasks/SKILL.md` creates operational tasks only from design/test-design evidence, not from legacy mandatory categories. | static | mandatory | Static inspection of task skill wording and search results for removed contract-driven task generation. | Tasks may still preserve applicable unresolved gaps from actual artifacts. |
| TD-006 | `sdd-review-workflow` spec | Verify `skills/sdd-review/SKILL.md` reviews present/planned operational evidence without requiring mandatory category completeness or adding columns to the 96-control matrix. | static | mandatory | Static inspection showing traceability/no-invention/placeholder checks only when applicable, with fixed matrix shape preserved. | Missing applicable evidence remains reviewable if design/artifacts make it applicable. |
| TD-007 | `sdd-review-security-workflow` spec; Secure Development Design | Verify `skills/sdd-review-security/SKILL.md` validates leakage and restricted-data boundaries only for operational evidence that exists. | static | mandatory | Static inspection showing blocking behavior for unsafe evidence and no blocker for absent readiness categories when design says operational considerations do not apply. | Exact `Pendiente de confirmar:` and `No aplica.` remain safe states but cannot hide required security proof. |
| TD-008 | `sdd-execution-persistence-contracts` spec | Verify `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md`, and `skills/_shared/sdd-post-apply-gates.md` consume actual design/test-design/tasks/apply/review/archive evidence instead of a readiness completeness table. | static | mandatory | Search and inspection evidence showing verify/archive preserve applicable refs, gaps, warnings, and unavailable-tooling notes without requiring all legacy categories. | `sdd-operational-doc` execution must not be an archive requirement. |
| TD-009 | Secure Development Design: secrets and restricted operational data rules | Verify safe-evidence/restricted-data protections remain in security and phase wording where operational evidence exists. | static | mandatory | Static inspection of `skills/_shared/sdd-security-contract.md`, affected phase skills, and relevant specs showing production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, raw logs, full ID lists, generated bytes, PAN/PII, and confidential operational context are restricted in ordinary SDD evidence. | Do not require real operational data disclosure to pass. |
| TD-010 | `sdd-security-guideline-catalog` spec | Verify active OpenSpec security guideline catalog preserves operational safe-evidence policy without implying mandatory operational categories. | static | mandatory | Comparison of `openspec/specs/sdd-security-guideline-catalog/spec.md` against the delta spec after apply. | Monitoring evidence should remain mechanism-oriented and not SQL-only. |
| TD-011 | All changed delta specs | Verify active OpenSpec specs reflect design-driven operational considerations after implementation. | static | mandatory | For each affected active spec under `openspec/specs/`, static comparison against the corresponding change delta: design, test-design, review, review-security, execution/persistence, operational-readiness workflow, and security guideline catalog. | This is spec sync validation, not archived artifact rewriting. |
| TD-012 | Design operational considerations; `sdd-operational-doc` boundary | Verify `skills/sdd-operational-doc/SKILL.md` remains a manual post-archive utility and handles absent archived operational data as `No aplica.` or `Pendiente de confirmar:` without inventing data. | static | mandatory | Static inspection showing manual invocation, archive-consuming behavior, preserved sections 1-9 and diagrams R1-R4, and no DAG/verify/archive requirement to run it. | Final-document-only user-provided values must not be backfilled into ordinary SDD artifacts. |
| TD-013 | Proposal out of scope; Design migration | Verify archived historical artifacts may retain old references without forcing active compatibility support. | manual | mandatory | Manual review statement that `openspec/changes/archive/**` references are historical only and no active skill/spec depends on them. | Do not rewrite archived change artifacts for this change. |
| TD-014 | Testing capabilities | Verify implementation evidence reports unavailable runtime/build/lint/type/format/coverage tooling explicitly. | manual | mandatory | Apply/verify evidence cites `openspec/config.yaml#testing` and reports unavailable commands rather than treating them as passed checks. | Static/manual evidence is required substitute evidence. |
| TD-015 | Secure Development Design: final documentation boundary | Verify final operational-document-only values remain isolated from design, tasks, review, verify, archive, examples, tests, or fixtures. | static | mandatory | Search/inspection evidence showing ordinary SDD artifacts use paths, section names, sanitized excerpts, exact placeholders, or redacted summaries only. | Unsafe evidence or invented operational data blocks downstream approval. |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Deleted stale authority | Confirm the old shared contract file is absent and no active workflow reads it. | static | TD-001 and TD-002 evidence. | Runtime/build/lint/type/format/coverage unavailable; use filesystem and search evidence. |
| Design-owned applicability | Confirm `sdd-design` uses conditional operational considerations and downstream phases do not synthesize missing categories. | static | TD-003 through TD-008 evidence. | Static contract inspection only. |
| Evidence traceability | Confirm applicable operational evidence flows from actual artifacts: design, test-design, tasks, apply evidence, reviews, verify, archive, or archived evidence. | static | TD-004 through TD-008 and TD-012 evidence. | No automated runner available. |
| Safe placeholders | Confirm exact `Pendiente de confirmar:` and `No aplica.` remain accepted safe states for missing applicable or inapplicable operational content. | static | TD-007, TD-009, TD-012 evidence. | Static wording validation only. |
| Restricted-data absence | Confirm ordinary SDD evidence excludes restricted production identifiers, secrets, payloads, raw logs, full ID lists, generated bytes, PAN/PII, and confidential operational context. | static | TD-009 and TD-015 evidence. | No scanner configured; use repository search and manual inspection. |
| Monitoring mechanisms | Confirm monitoring wording remains mechanism-oriented and does not default to SQL-only checks. | static | TD-010 and affected skill/spec inspection. | Static inspection only. |
| Manual operational document boundary | Confirm `sdd-operational-doc` remains manual post-archive and does not block archive/verify. | static | TD-012 evidence. | No runtime documentation generator test available. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-OP-001` | Remove references to the deleted readiness contract without weakening safe-evidence policy. | Yes | TD-001, TD-002, TD-009 | covered | None |
| `SEC-OP-002` | Ordinary SDD evidence must not contain restricted operational identifiers, secrets, payloads, raw logs, full ID lists, generated bytes, PAN/PII, or confidential operational context. | Yes | TD-009, TD-015 | covered | None |
| `SEC-OP-003` | Operational evidence, when present, must use safe summaries, paths, section names, sanitized excerpts, redacted placeholders, `Pendiente de confirmar:`, or `No aplica.`. | Yes | TD-007, TD-009, TD-012 | covered | None |
| `SEC-OP-004` | Final operational-document-only values may appear only when explicitly supplied by the user for the manual final document and must not be copied back into ordinary SDD artifacts. | Yes | TD-012, TD-015 | covered | None |
| `SEC-OP-005` | Review-security owns leakage validation and any exhaustive security/source-row expansion; test-design does not require design YAML, matrices, Source IDs, or all-row `N/A` bookkeeping. | Yes | Source ID Coverage Baseline; TD-007 | covered | None |
| `SEC-OP-006` | No planned exception may approve unsafe evidence or incomplete proof after deleting the old contract. | Yes | TD-009, TD-015 | covered | None |

## No-Impact Assessment

Not applicable. This change has no application runtime behavior impact, but it has material SDD workflow, review, verification, archive, and evidence-safety impact. Static/manual checks are required.

## Evidence Expectations

- Mandatory cases require implementation, static/manual evidence, or a justified skip.
- Non-mandatory cases should be reported as warnings when uncovered, but no non-mandatory cases are planned for this change.
- Security validation evidence must cite embedded `design.md` narrative rules, owner phase, and planned static/manual evidence.
- Applicable narrative category rules require planned safe evidence. Exhaustive `N/A` source coverage and missed-applicable validation remain owned by `review-security-report.md`.
- Test-design consumes narrative design rules only and MUST NOT require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, or the full Source ID matrix; exhaustive row materialization belongs only to `review-security-report.md`.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands are unavailable and must be reported as unavailable evidence, not passing evidence.
- Archived references to the old contract are historical audit evidence only; active workflow files must not depend on them.

## Open Questions

None.
