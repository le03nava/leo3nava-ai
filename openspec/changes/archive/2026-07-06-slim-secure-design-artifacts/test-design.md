# Test Design: Slim Secure Design Artifacts

## Overview

This change has no runtime application surface; it changes SDD Markdown contracts, OpenSpec requirements, and adapter prompts. Test planning therefore relies on static/manual artifact checks. The key testability boundary is that the shared catalog owns the exhaustive 155 Source ID inventory, `design.md#secure-development-design` owns slim referenced coverage, `test-design.md` owns grouped evidence planning, `review-security-report.md` owns exact-once expansion, and verify/archive preserve summaries, warnings, exceptions, and links without copying the full matrix.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/slim-secure-design-artifacts/proposal.md` | Intent, scope, non-goals, artifact-boundary risks, success criteria. |
| Spec | `openspec/changes/slim-secure-design-artifacts/specs/sdd-design-workflow/spec.md` | Slim secure-design coverage, catalog reference, expected count, no exhaustive design matrix. |
| Spec | `openspec/changes/slim-secure-design-artifacts/specs/sdd-test-design-workflow/spec.md` | Grouped source-row test planning, unavailable runner handling, N/A and warning evidence expectations. |
| Spec | `openspec/changes/slim-secure-design-artifacts/specs/sdd-review-security-workflow/spec.md` | Exclusive exact-once full Source ID expansion in `review-security-report.md`. |
| Spec | `openspec/changes/slim-secure-design-artifacts/specs/sdd-security-guideline-catalog/spec.md` | Catalog ownership of the authoritative 155 Source ID inventory and compact mappings. |
| Spec | `openspec/changes/slim-secure-design-artifacts/specs/sdd-execution-persistence-contracts/spec.md` | Verify/archive source-row consumption and preservation boundaries. |
| Design | `openspec/changes/slim-secure-design-artifacts/design.md` | Architecture decisions, affected files, testing strategy, rollout, and risk boundaries. |
| Secure Development Design | `openspec/changes/slim-secure-design-artifacts/design.md#secure-development-design` | Security-impacting classification, compact `SEC-*` controls, catalog metadata, `expectedSourceIdCount: 155`, grouped coverage, lifecycle state, owners, safe evidence, N/A/exception policy. |
| Testing Capabilities | `openspec/config.yaml#testing` | Runtime/build/lint/type/format/coverage tooling unavailable; static/manual evidence required. |

## Source ID Coverage Baseline

Corporate source-row coverage applies.

| Boundary | Planned Validation |
| --- | --- |
| Inventory authority | Static/manual check that `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` remains the authoritative inventory. |
| Expected count | Static/manual check that slim design and updated contracts preserve `expectedSourceIdCount: 155`. |
| Compact mappings | Static/manual check that only `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001` are accepted compact mappings. |
| Grouped coverage | Static/manual checks cover sections 1-15 by group counts totaling 155, without duplicating all 155 rows here or in design. |
| Lifecycle status | Static/manual check that source coverage remains `planned` until implementation/review evidence updates it. |
| Evidence owners | Static/manual check that ownership flows through design, test-design, apply, review-security, verify, and archive. |
| Exact-once expansion | Static/manual check that only `sdd-review-security` materializes and validates every Source ID exactly once in `review-security-report.md`. |
| N/A policy | Static/manual check that any future `N/A` source coverage requires positive irrelevance evidence and justification. This design declares no not-applicable guidelines. |
| Warning preservation | Static/manual check that warning-only source rows remain observable and are preserved by verify/archive when mandatory evidence is complete. |
| Safe evidence | Static/manual check that artifacts cite paths, sections, summaries, and sanitized observations only; raw secrets, credentials, PAN, PII, tokens, private keys, connection strings, and confidential values are prohibited. |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Proposal success criteria; design decisions | Confirm the artifact boundary is explicit: catalog owns inventory, design stays slim, test-design plans grouped evidence, review-security expands, verify/archive preserve summaries and links. | static | mandatory | Diff/read-back of shared contract, phase skills, adapter prompts, and OpenSpec deltas. | Blocks if any phase owns the wrong detail level. |
| TD-002 | Spec: sdd-design-workflow, Design Preserves Compact Summary | Confirm design/test-design do not duplicate the general 96-control matrix or the full 155-row Source ID matrix. | static | mandatory | Markdown inspection of `design.md`, this file, updated skills/specs/prompts. | Full matrix belongs only in `review-security-report.md`. |
| TD-003 | Spec: sdd-security-guideline-catalog, Corporate Source Row Inventory | Confirm the catalog is declared as the authoritative inventory with 155 expanded Source IDs and compact mappings. | manual | mandatory | Reviewer checklist citing catalog path, snapshot/version metadata, expected count, and section grouping. | This test-design references groups only. |
| TD-004 | Secure design `sourceRowCoverage.expectedSourceIdCount` | Confirm design and downstream contracts preserve `expectedSourceIdCount: 155`. | static | mandatory | Read-back evidence from `design.md#secure-development-design` and changed contract files. | Count mismatch blocks. |
| TD-005 | Secure design compact mapping rule | Confirm only the eight compact `SEC-*` IDs are accepted and unknown/missing/empty mappings route to blockers. | static | mandatory | Diff/read-back of shared contract, design/test-design/review-security skills, and adapter prompts. | Covers mapping vocabulary. |
| TD-006 | Secure design groups: 1 Authentication, 2 Passwords | Plan source group coverage for Source ID refs `1.1-1.10` and `2.1-2.23`, mapped to `SEC-AUTH-001`, `SEC-ACCESS-001`, and `SEC-SECRET-001`. | manual | mandatory | Reviewer checklist verifying grouped refs, counts 10 + 23, compact mappings, owner phases, and review-security exact-once handoff. | No per-row expansion here. |
| TD-007 | Secure design groups: 3 Logging, 8 Error Handling | Plan source group coverage for `3.1-3.11` and `8.1-8.5`, mapped to `SEC-LOG-001`. | manual | mandatory | Reviewer checklist verifying safe logging/error evidence and warning preservation expectations. | Evidence must avoid sensitive operational detail. |
| TD-008 | Secure design groups: 4 Cryptography, 10 Memory Management, 13 Data Protection, 15 PAN | Plan source group coverage for data/secrets/PAN groups mapped to `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-AUTH-001`, and `SEC-SESS-001`. | manual | mandatory | Reviewer checklist verifying safe-evidence policy, no raw sensitive values, compact mappings, and review-security handoff. | Unsafe evidence blocks. |
| TD-009 | Secure design groups: 5 Databases, 11 Input Validation, 12 Output Encoding | Plan database/input/output coverage mapped to `SEC-DB-001`, `SEC-ACCESS-001`, `SEC-SECRET-001`, and `SEC-DATA-001`. | manual | mandatory | Reviewer checklist verifying catalog references, grouped counts 12 + 16 + 5, and exact-once review-security ownership. | Runtime DB code is not changed. |
| TD-010 | Secure design groups: 6 Coding, 7 Session Management, 9 File Handling, 14 Access Control | Plan coding/session/file/access coverage mapped across all relevant compact IDs. | manual | mandatory | Reviewer checklist verifying grouped counts 14 + 13 + 12 + 9, lifecycle state, evidence owners, and blocker routing. | These groups exercise the broadest mapping surface. |
| TD-011 | Spec: Source Row Test Planning, Applicable source group receives checks | Confirm every applicable source group has a planned static/manual check or non-test evidence plan citing catalog reference and compact mapping. | static | mandatory | This `test-design.md` table plus read-back of source coverage baseline. | Blocks if any applicable group is missing. |
| TD-012 | Spec: Source Row Test Planning, No runtime runner exists | Confirm unavailable runtime automation is reported explicitly and replaced with static/manual evidence, not treated as passing evidence. | static | mandatory | Read-back of `openspec/config.yaml#testing` and this file's Testing Tooling Constraints section. | Applies to runtime, build, lint, type, format, and coverage. |
| TD-013 | Spec: Source Row N/A and Warning Evidence | Confirm `N/A` policy requires positive irrelevance evidence and warning-only rows remain tracked without blocking when mandatory evidence is complete. | static | mandatory | Read-back of shared contract and phase skill wording after apply; this file declares no current N/A groups. | Unsupported N/A remains blocking. |
| TD-014 | Spec: Exhaustive Source Row Security Review | Confirm `sdd-review-security` is the only active phase that materializes the 155-row matrix and validates each Source ID exactly once. | static | mandatory | Diff/read-back of review-security skill, adapter prompt, and OpenSpec delta. | Missing/duplicate/unknown Source IDs block review-security. |
| TD-015 | Spec: General review is cited, not duplicated | Confirm security review may cite `review-report.md` rows but does not duplicate the full 96-control general review matrix. | static | mandatory | Review-security contract inspection. | Preserves slim security report boundary. |
| TD-016 | Spec: Verify Source Row Consumption | Confirm verify consumes non-blocking security review evidence, preserves catalog snapshot/count/mappings/warnings/exceptions/links, and blocks remaining source-row blockers. | static | mandatory | Verify skill/spec/prompt diff and final verify report read-back during verification phase. | Verify does not own or duplicate the full matrix. |
| TD-017 | Spec: Archive Source Row Preservation | Confirm archive preserves source-row summaries, catalog identity/path, expected count, compact mappings, warnings, exceptions, and evidence refs without copying the full review-security matrix. | static | mandatory | Archive skill/spec/prompt diff and archive report read-back during archive phase. | Archive must not require legacy standalone artifacts. |
| TD-018 | Proposal out of scope | Confirm no runtime application code, package/build configuration, validation scripts, or legacy active `security-design.md` / `security-applicability.md` dependencies are introduced. | static | mandatory | Changed-file review and artifact read-back. | Runtime/tooling changes are out of scope. |
| TD-019 | Design safe-evidence policy | Confirm generated evidence remains safe: paths, anchors, summaries, sanitized observations, or redacted placeholders only. | manual | mandatory | Reviewer checklist during apply/review-security/verify/archive. | Raw sensitive values block. |
| TD-020 | Warnings and advisory evidence | Confirm warning-only coverage, if encountered later, is preserved as observation evidence and routed forward only when mandatory evidence is complete. | manual | non-mandatory | Review-security/verify/archive report inspection. | Advisory by itself; mandatory gaps still block. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | Design/test-design reference authentication Source IDs by catalog snapshot and grouped coverage; review-security expands each concrete Source ID exactly once. | Yes | TD-003, TD-004, TD-006, TD-014 | covered | None |
| `SEC-SESS-001` | Session Source ID coverage remains catalog-owned and planned by group; exhaustive verdicts are deferred to review-security. | Yes | TD-006, TD-008, TD-010, TD-014 | covered | None |
| `SEC-DATA-001` | Artifacts cite safe paths, sections, and summaries only; no raw PAN, PII, credentials, tokens, connection strings, or confidential values. | Yes | TD-008, TD-009, TD-019 | covered | None |
| `SEC-SECRET-001` | Evidence references secret/config handling requirements without reproducing secret values; unsafe evidence blocks source-row validation. | Yes | TD-005, TD-006, TD-008, TD-009, TD-019 | covered | None |
| `SEC-ACCESS-001` | Workflow routing remains denial-by-default for missing, duplicate, malformed, unmapped, unsafe, or unsupported N/A source-row evidence. | Yes | TD-005, TD-010, TD-013, TD-016, TD-017 | covered | None |
| `SEC-FILE-001` | File-based OpenSpec and skill artifacts carry references and summaries only; review-security owns exhaustive row materialization. | Yes | TD-001, TD-002, TD-010, TD-014, TD-018 | covered | None |
| `SEC-DB-001` | Database Source IDs remain covered by catalog mappings and are validated through review-security expansion even though runtime database code is unchanged. | Yes | TD-009, TD-014, TD-018 | covered | None |
| `SEC-LOG-001` | Generated evidence remains review-safe and avoids raw secrets, PAN, credentials, tokens, or unnecessary sensitive operational context. | Yes | TD-007, TD-019, TD-020 | covered | None |

## No-Impact Assessment

Not applicable. This is a documentation/contract-only change with no runtime behavior impact, but it has security-process impact and therefore requires static/manual evidence planning.

## Testing Tooling Constraints

`openspec/config.yaml#testing` reports the following tooling as unavailable. These are unavailable evidence, not passing evidence:

| Tooling Area | Available | Command | Planned Substitute |
| --- | --- | --- | --- |
| Test runner | No | None | Static/manual artifact inspection. |
| Unit tests | No | None | Static/manual contract coverage checks. |
| Integration tests | No | None | Static/manual cross-artifact boundary checks. |
| E2E tests | No | None | Static/manual SDD phase-flow checks. |
| Coverage | No | None | Source group coverage checklist and review-security exact-once expansion. |
| Linter | No | None | Manual Markdown/frontmatter inspection. |
| Type checker | No | None | Not applicable; no typed runtime source changed. |
| Formatter | No | None | Manual Markdown readability and structure review. |
| Build command | No | None | Not applicable; no buildable runtime application exists. |

## Evidence Expectations

- Mandatory cases require implementation, execution where tooling exists, static/manual evidence, or a complete approved exception.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Security validation evidence must cite `design.md#secure-development-design`, catalog snapshot/version/path, `expectedSourceIdCount: 155`, compact mappings, lifecycle status, owner phase, and planned static/manual evidence.
- Applicable source groups require planned safe evidence. `N/A` source coverage requires evidence plus irrelevance justification checks. Unsupported `N/A` remains blocking.
- Warning-only source rows must be preserved with expected observation evidence and may proceed only when mandatory evidence is complete.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands are unavailable and must be reported as unavailable evidence, never passing evidence.
- Design and test-design must not materialize the full 155-row Source ID matrix or the 96-control general review matrix.
- `review-security-report.md` is the only active artifact expected to contain the exhaustive Source ID matrix for this change.
- Archive must preserve source-row summaries, catalog identity/path, expected count, compact mappings, warning/exception summaries, report links, and safe evidence references without requiring standalone `security-design.md` or `security-applicability.md`.

## Open Questions

None.
