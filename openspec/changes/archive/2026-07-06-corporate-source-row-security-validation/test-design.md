# Test Design: Corporate Source Row Security Validation

## Overview

This change is validated through static/manual artifact checks because `openspec/config.yaml` declares no runtime test runner, build command, linter, formatter, type checker, or coverage command. The test design focuses on exhaustive corporate Source ID coverage, exact-once source-row validation, compact `SEC-*` mapping preservation, safe evidence handling, and phase-to-phase traceability from design through archive.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/corporate-source-row-security-validation/proposal.md` | Intent, scope, non-goals, safe-evidence policy, persistence compatibility, and success criteria. |
| Specs | `openspec/changes/corporate-source-row-security-validation/specs/**/spec.md` | Source-row catalog rules, design/test-design/review-security requirements, verify/archive consumption, and persistence compatibility. |
| Design | `openspec/changes/corporate-source-row-security-validation/design.md` | Architecture decisions, source-row schema, routing semantics, rollout work units, risks, and testing strategy. |
| Secure Development Design | `openspec/changes/corporate-source-row-security-validation/design.md#secure-development-design` | Eight compact controls, expected Source ID coverage, source-row groups, safe evidence rules, N/A policy, traceability, lifecycle statuses, and archive gates. |
| Testing Capabilities | `openspec/config.yaml#testing` | Confirms static/manual validation only; unavailable automation must be reported as unavailable evidence, not passing evidence. |

## Source ID Coverage Baseline

The expected source-row universe is 155 expanded Source IDs from `skills/_shared/security-guideline-catalog.md#full-corporate-guideline-snapshot`, as declared by `design.md#secure-development-design`.

| Corporate Section | Expected Expanded Source IDs | Count | Expected Compact Mapping |
| --- | --- | ---: | --- |
| 1. Authentication | `1.1`-`1.10` | 10 | `SEC-AUTH-001`, `SEC-ACCESS-001` |
| 2. Passwords | `2.1`-`2.23` | 23 | `SEC-AUTH-001`, `SEC-SECRET-001` |
| 3. Access and Activity Logging | `3.1`-`3.11` | 11 | `SEC-LOG-001` |
| 4. Cryptography | `4.1`-`4.8` | 8 | `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-AUTH-001`, `SEC-SESS-001` |
| 5. Databases | `5.1`-`5.12` | 12 | `SEC-DB-001`, `SEC-ACCESS-001`, `SEC-SECRET-001` |
| 6. Coding | `6.1`-`6.14` | 14 | all compact `SEC-*` categories as applicable |
| 7. Session Management | `7.1`-`7.13` | 13 | `SEC-SESS-001` |
| 8. Error Handling | `8.1`-`8.5` | 5 | `SEC-LOG-001` |
| 9. File Handling | `9.1`-`9.12` | 12 | `SEC-FILE-001` |
| 10. Memory Management | `10.1`-`10.6` | 6 | `SEC-DATA-001`, `SEC-SECRET-001` |
| 11. Input Validation | `11.1`-`11.16` | 16 | `SEC-DB-001` |
| 12. Output Encoding | `12.1`-`12.5` | 5 | `SEC-DATA-001`, `SEC-DB-001` |
| 13. Data Protection | `13.1`-`13.9` | 9 | `SEC-DATA-001`, `SEC-ACCESS-001`, `SEC-SECRET-001` |
| 14. Access Control | `14.1`-`14.9` | 9 | `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-AUTH-001` |
| 15. PAN | `15.1`-`15.2` | 2 | `SEC-DATA-001` |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Proposal success criteria; Catalog spec: Corporate Source Row Inventory; Design sourceRowCoverage | Confirm every Source ID from the Full Corporate Guideline Snapshot is expanded before coverage checks and appears exactly once across catalog/design/review-security matrices. | static | mandatory | Artifact inspection showing 155 expanded Source IDs, no compressed ranges in validation matrices, and exact-once coverage. | Missing, duplicate, or unknown Source IDs route to `resolve-blockers`. |
| TD-002 | Catalog spec: Ranges expand before coverage | Validate that range notation such as `1.1-1.10` is expanded only into concrete IDs that exist in the preserved snapshot. | static | mandatory | Source-row inventory or matrix listing concrete IDs and rejecting IDs outside the snapshot. | Compressed ranges may remain only in human summaries, not validation rows. |
| TD-003 | Catalog spec: Compact SEC Mapping Coverage; Design mapping strategy | Confirm every source row maps to one or more existing compact `SEC-*` IDs and no source row replaces the compact eight-control taxonomy. | static | mandatory | Mapping table or row matrix with valid `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, or `SEC-LOG-001` values. | Unknown or missing mappings block. |
| TD-004 | Catalog spec: Shared Security Contract Source Row Schema; Design Source Row Schema | Validate required row fields exist: `sourceId`, `corporateSection`, `guidelineTextRef`, `pciAlignment`, `mappedCompactGuidelineIds`, `applies`, `complies`, `lifecycleStatus`, `evidenceLocation`, `observations`, `finding`, and route/status metadata where the phase owns routing. | static | mandatory | Schema section in `skills/_shared/sdd-security-contract.md` and compatible examples in affected phase artifacts. | Field omissions or malformed rows route to `resolve-blockers`. |
| TD-005 | Catalog spec: Shared Security Contract Source Row Schema; Design Source Row Schema | Validate allowed values: `applies` is `Yes`, `No`, or `N/A`; `complies` is phase-appropriate (`planned`, `Yes`, `No`, or `N/A`); lifecycle/status/finding/route values match the shared contract. | static | mandatory | Contract text and review-security matrix examples use only allowed values. | Design rows use planned/applicability semantics; review rows use verdict semantics. |
| TD-006 | Catalog spec: Safe Source Row Evidence; Proposal safe-evidence scope | Confirm applicable rows have review-safe evidence references and observations, not raw sensitive values. | manual | mandatory | Evidence locations cite paths, sections, command summaries, sanitized summaries, or redacted placeholders. | Applicable rows without safe evidence block. |
| TD-007 | Catalog spec: Safe Source Row Evidence; Test-design spec: Source Row N/A and Warning Evidence | Confirm each `N/A` source row has evidence and justification proving irrelevance by category, platform, API, data class, or workflow. | manual | mandatory | Matrix rows include an evidence reference plus N/A rationale; missing rationale is called out as blocking. | This design currently declares no `notApplicableGuidelines`; future N/A rows must still satisfy this check. |
| TD-008 | Proposal safe-evidence rule; Review-security spec: Evidence is unsafe | Reject evidence containing secrets, PII, PAN, tokens, connection strings, private keys, or confidential values. | manual | mandatory | Review-security/report guidance identifies unsafe evidence as a blocker and uses sanitized examples only. | Unsafe evidence routes to `resolve-blockers` unless remediation requires apply work. |
| TD-009 | Design spec: Secure Design Source ID Coverage | Confirm `design.md#secure-development-design` declares expected Source ID coverage, compact mappings, expected evidence owners, lifecycle status, and downstream traceability. | static | mandatory | Design section includes `sourceRowCoverage`, 155 expected rows, group mappings, compact control coverage, and traceability path. | Missing design coverage blocks test-design or review-security. |
| TD-010 | Test-design spec: Source Row Test Planning | Confirm `test-design.md` plans static/manual checks for every applicable source-row category and preserves no-runner limitations. | static | mandatory | This artifact includes source coverage baseline, mandatory test cases, security coverage table, and explicit unavailable automation statement. | Missing planned evidence for mandatory rows blocks. |
| TD-011 | Review-security spec: Exhaustive Source Row Security Review | Confirm review-security is planned to validate every row against design, test-design, apply evidence, changed files, and `review-report.md`, without duplicating the 96-control general review matrix. | static | mandatory | Review-security skill/report requirements cite all evidence sources and state general review is cited, not copied. | Missing artifacts, missing changed-file context, or duplicated general matrix are blockers. |
| TD-012 | Review-security spec: Source Row Evidence Correlation | Confirm rows cannot pass solely by being listed; evidence must corroborate applicability, compliance, or justified `N/A`. | manual | mandatory | Review-security matrix includes evidence source references and finding status per row. | Listed-only rows fail validation. |
| TD-013 | Execution persistence spec: Verify Source Row Consumption | Confirm verify consumes non-blocking source-row security review evidence and blocks unresolved source-row blockers without owning or duplicating the full source-row matrix. | static | mandatory | Verify contract/skill requires readable review-security evidence, cites verdict and warnings, and routes blockers to apply or resolve-blockers by cause. | Warnings-only may proceed when mandatory evidence is complete. |
| TD-014 | Execution persistence spec: Archive Source Row Preservation | Confirm archive preserves source-row coverage, mappings, warnings, exceptions, and safe evidence refs, and requires passing verification plus non-blocking security review. | static | mandatory | Archive readiness/report guidance retains source-row audit trail and denies archive when mandatory blockers remain. | Standalone legacy security artifacts are not required for new changes. |
| TD-015 | Execution persistence spec: Source Row Persistence Compatibility; Persistence contract | Validate OpenSpec, Engram, hybrid, and none modes preserve identical source-row semantics while respecting backend behavior. | static | mandatory | Contract text keeps OpenSpec in files, Engram under established SDD topic keys, hybrid in both with conflict policy, and none inline with recovery limits. | Backend choice must not change row schema or validation semantics. |
| TD-016 | Design testing strategy; `openspec/config.yaml#testing` | Confirm unavailable runtime/build/coverage/lint/type/format tooling is reported explicitly and not treated as passing evidence. | static | mandatory | Test-design, verify, and review-security reports include unavailable-tooling notes where relevant. | No automated runtime test command may be claimed for this repo. |
| TD-017 | Proposal affected areas; Design file changes | Confirm affected shared contracts, phase skills, agent prompts, and source specs are represented by tasks/apply evidence. | static | non-mandatory | Tasks/apply evidence links changed files to source-row requirements and compact controls. | Advisory for implementation completeness; becomes blocking only when required artifacts are missing. |
| TD-018 | Design tradeoffs and carried risks | Confirm reviewer load mitigation keeps compact summaries first and detailed source rows below them. | manual | non-mandatory | Review-facing docs place compact summaries before large matrices and preserve row details as operational evidence. | Supports review budget and cognitive load goals. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | Authentication/password Source IDs must expand, map, and validate under compact auth coverage. | Yes | TD-001, TD-002, TD-003, TD-009, TD-011 | covered | None |
| `SEC-SESS-001` | Session Source IDs must be expanded and mapped without replacing the compact session control. | Yes | TD-001, TD-002, TD-003, TD-011, TD-014 | covered | None |
| `SEC-DATA-001` | Data/PAN, crypto, memory, and output rows require safe source-row evidence. | Yes | TD-003, TD-006, TD-008, TD-011, TD-014 | covered | None |
| `SEC-SECRET-001` | Secret-related rows and evidence must never expose raw secret values. | Yes | TD-003, TD-006, TD-008, TD-013 | covered | None |
| `SEC-ACCESS-001` | Missing coverage, mappings, N/A evidence, or unsafe evidence must deny progression. | Yes | TD-001, TD-004, TD-005, TD-007, TD-011, TD-013 | covered | None |
| `SEC-FILE-001` | File-handling Source IDs must be mapped and validated for future file-related changes. | Yes | TD-001, TD-003, TD-006, TD-011 | covered | None |
| `SEC-DB-001` | Database/input-validation/output-to-query rows must map to compact DB coverage and require corroborating evidence. | Yes | TD-003, TD-006, TD-011, TD-012 | covered | None |
| `SEC-LOG-001` | Logging/error rows must validate audit usefulness while blocking sensitive evidence leakage. | Yes | TD-006, TD-008, TD-011, TD-014, TD-018 | covered | None |

## No-Impact Assessment

Not applicable. This change modifies active SDD security contracts, phase skills, prompts, and source specs, so it has security workflow and validation impact.

## Evidence Expectations

- Mandatory cases are verification-blocking when uncovered, unsafe, malformed, or unsupported by evidence.
- Non-mandatory cases are advisory unless they reveal missing required artifacts, missing safe evidence, or missing implementation coverage.
- Source-row validation evidence must cite embedded `design.md` metadata, catalog snapshot/version, compact mapping, lifecycle status, apply evidence, changed files, review-report references, and review-security findings as applicable.
- Applicable rows require safe evidence. `N/A` rows require safe evidence plus justification; unsupported `N/A` routes to `resolve-blockers`.
- Unsafe evidence is rejected when it includes secrets, PII, PAN, tokens, connection strings, private keys, or confidential values.
- Review-security validates every expected row against design, this test design, apply evidence, changed files, and `review-report.md`; it cites the general review matrix instead of duplicating it.
- Verify consumes the non-blocking review-security verdict and must preserve warning-only evidence without duplicating the full matrix.
- Archive preserves source-row coverage, mappings, warnings, exceptions, and evidence references after passing verification and non-blocking security review.
- OpenSpec, Engram, hybrid, and none modes must preserve the same row semantics while following their backend persistence rules.
- Runtime tests, linters, type checkers, formatters, build commands, and coverage commands are unavailable for this repository and must be reported as unavailable evidence, never as passed checks.

## Open Questions

None.
