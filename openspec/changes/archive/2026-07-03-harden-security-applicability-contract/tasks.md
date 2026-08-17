# Tasks: Harden Security Applicability Contract

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 850-1200 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 -> PR 2 -> PR 3 |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

Rationale: scope crosses shared schema, applicability executor, catalog metadata, validator script, downstream phase contracts, and evidence; auto-chain allows apply to start PR 1.

### Suggested Work Units

| Unit | Goal | Likely PR | Work unit boundaries |
|------|------|-----------|----------------------|
| 1 | Shared contract and applicability template | PR 1 | Base main; finish when schema/template/spec compatibility is inspectable. |
| 2 | Catalog source/severity metadata | PR 2 | Base PR 1; finish when all compact `SEC-*` rows expose Source IDs and severity. |
| 3 | Validator and downstream evidence | PR 3 | Base PR 2; finish when validator, compatibility notes, and evidence are complete. |

## Phase 1: Shared Contract and Applicability Template / PR 1

- [x] 1.1 Update `skills/_shared/sdd-security-contract.md` with additive `catalog`, `categoryDecisionMatrix`, `overridesApplied`, `validation`, no-impact proof, exception, and severity rules for TD-001..TD-009, TD-022, SEC-AUTH-001, SEC-DATA-001, SEC-SECRET-001.
- [x] 1.2 Update `skills/sdd-security-applicability/SKILL.md` template/gates to require every taxonomy category once, safe `rules.security-applicability` overrides, validator metadata, and blocked unknowns for TD-001..TD-009, SEC-AUTH-001, SEC-SECRET-001, SEC-FILE-001.
- [x] 1.3 Update `openspec/specs/sdd-security-applicability-workflow/spec.md` and prompt-copy compatibility notes only if source specs/contracts need immediate apply evidence before archive.
- [x] 1.4 Record PR 1 static evidence in this file for TD-001..TD-009 and unavailable runtime tooling per TD-026.

### PR 1 Static Evidence

| Case | Evidence |
| --- | --- |
| TD-001 | `skills/_shared/sdd-security-contract.md` and `skills/sdd-security-applicability/SKILL.md` now require `catalog` identity and exactly one `categoryDecisionMatrix` row for each supported category. |
| TD-002 | The shared contract and applicability validation checklist state missing or duplicate category rows make the artifact invalid. Executable negative validation remains PR 3 scope. |
| TD-003 | Both contracts require `blocking` + `unknown` category decisions to block phase success and identify missing evidence. |
| TD-004 | No-impact routing remains compatible only when every category is `not-applicable`, proof is complete, unknowns are absent, and validation metadata is valid. |
| TD-005 | The shared contract and applicability skill explicitly reject absent evidence, empty guideline mappings, or missing rationale as no-impact proof. |
| TD-006 | `rules.security-applicability` safe overrides are documented as extra prompts, stricter source coverage, validator mode, and stricter category severity, with `overridesApplied` metadata. |
| TD-007 | Unsafe overrides that disable categories, weaken source coverage, downgrade `blocking`, or bypass no-impact proof must be rejected or ignored in favor of the base contract. |
| TD-008 | Applicability template now records validator path, status, `checkedAt`, and notes. Executable validator pass evidence is intentionally deferred to PR 3. |
| TD-009 | Validation gates now block success on invalid schema fields, incomplete matrix, unsupported severity, unsafe overrides, blocking unknowns, invalid guideline/source refs, or validation failure. Executable failure evidence remains PR 3 scope. |
| TD-022 | Shared contract states schema hardening is additive and preserves artifact identity, `schemaVersion: 1`, existing `SEC-*` IDs, and no-impact routing compatibility. |
| TD-026 | Runtime test runner, linter, type checker, formatter, and coverage remain unavailable per `openspec/config.yaml`; PR 1 evidence is static/manual only and does not claim missing tooling as pass evidence. |

PR 1 compatibility note: no canonical `openspec/specs/sdd-security-applicability-workflow/spec.md` or prompt-copy edit was required in this slice. The change delta specs already define the new requirements, and `agents/sdd/sdd-security-applicability.md` delegates to the updated skill file, so immediate apply evidence is inspectable without archive-time spec sync or downstream prompt-copy changes.

## Phase 2: Catalog Source Coverage and Severity / PR 2

- [x] 2.1 Update `skills/_shared/security-guideline-catalog.md` with snapshot metadata, taxonomy version, Source IDs, operational severity, predicates, validator fields, and exception fields for TD-010..TD-016, SEC-AUTH-001..SEC-LOG-001.
- [x] 2.2 Verify all compact guidelines `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, `SEC-LOG-001` map to valid Source IDs from snapshot `security-guidelines-initial-user-snapshot-2026-06-30` for TD-010, TD-011, TD-015.
- [x] 2.3 Ensure catalog severity uses only `blocking`, `conditional`, `advisory` and excludes `Menor`, `Media`, `Mayor` for TD-012..TD-014, SEC-LOG-001.
- [x] 2.4 Record PR 2 manual catalog evidence and residual Source ID review risk in this file for TD-010..TD-016.

### PR 2 Static Evidence

| Case | Evidence |
| --- | --- |
| TD-010 | `skills/_shared/security-guideline-catalog.md` compact rows now expose formal Source IDs for `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001` under snapshot `security-guidelines-initial-user-snapshot-2026-06-30`. |
| TD-011 | Manual inspection confirmed no compact `SEC-*` guideline row is missing Source IDs; executable strict-source negative validation remains PR 3 scope. |
| TD-012 | All compact guideline rows use `operationalSeverity` value `blocking`, and the Operational Severity Contract states unresolved blocking evidence prevents phase success unless a complete approved exception exists. |
| TD-013 | The catalog now defines `conditional` behavior, predicate rationale requirements, and validator checks even though current compact rows use `blocking` with predicate `N/A`. |
| TD-014 | Manual grep/static inspection found no `Menor`, `Media`, or `Mayor` labels in `skills/_shared/security-guideline-catalog.md`; the catalog explicitly limits applicability severity to `blocking`, `conditional`, and `advisory`. |
| TD-015 | Snapshot metadata now records catalog version `1`, taxonomy version `1`, Source ID pattern, and validator checks requiring artifact snapshot identity to match the catalog snapshot before guideline/source resolution. |
| TD-016 | The catalog now defines advisory preservation requirements so advisory obligations remain downstream-visible as risk, guidance, or archive-readable evidence when non-blocking. |

Manual Source ID review evidence: each mapped Source ID/range was checked against the preserved snapshot tables in `skills/_shared/security-guideline-catalog.md`: `1.1-1.10`, `2.1-2.23`, `3.1-3.11`, `4.1`, `4.2`, `4.8`, `5.1-5.12`, `5.5`, `6.1`, `6.2-6.4`, `6.3`, `6.12`, `7.1-7.13`, `8.1-8.5`, `9.1-9.12`, `11.1-11.16`, `12.2`, `13.1-13.9`, `13.5`, `14.1-14.9`, `14.7`, `14.8`, and `15.1-15.2` are present in the recorded snapshot.

Residual Source ID review risk: Source mappings are manually reviewed from the initial in-repo snapshot and are now inspectable, but executable source-reference validation is intentionally deferred to PR 3 with `scripts/validate_security_applicability.ps1`.

## Phase 3: Validator, Downstream Compatibility, and Evidence / PR 3

- [x] 3.1 Create `scripts/validate_security_applicability.ps1` for static Markdown/YAML checks: schema fields, routing, matrix completeness, no-impact proof, guideline/source refs, overrides, severity, and validation metadata for TD-001..TD-015, TD-020, TD-021, SEC-DB-001.
- [x] 3.2 Update `skills/sdd-security-design/SKILL.md` to consume enriched applicability metadata only when `securityImpact: true` and preserve no-impact skip compatibility for TD-017..TD-020, SEC-SESS-001, SEC-ACCESS-001.
- [x] 3.3 Update `skills/sdd-test-design/SKILL.md`, `skills/sdd-tasks/SKILL.md`, `skills/sdd-verify/SKILL.md`, and `skills/sdd-archive/SKILL.md` with validation metadata consumption, missing/invalid no-impact blockers, archive evidence fields, and unavailable-runtime-test reporting for TD-019, TD-023..TD-026, SEC-ACCESS-001, SEC-LOG-001.
- [x] 3.4 Run or manually inspect validator positive/negative examples and record static/manual evidence for TD-001..TD-026; keep TD-027 as advisory review-slicing evidence.

### PR 3 Static Evidence

| Case | Evidence |
| --- | --- |
| TD-001 | `scripts/validate_security_applicability.ps1` checks catalog snapshot identity, taxonomy version, and exactly one matrix row for each supported category. Positive fixture validation exited `0`. |
| TD-002 | Negative fixture validation exited non-zero and reported missing matrix categories. |
| TD-003 | Validator treats `decision: unknown` with `severity: blocking` as a failure and names the affected category. |
| TD-004 | Validator accepts no-impact only when every category is `not-applicable`, proof is complete, evidence refs exist, validation metadata is non-failing, and unknowns are absent; downstream skills preserve this skip compatibility. |
| TD-005 | Validator rejects incomplete no-impact proof and empty category evidence refs; downstream skills block missing/invalid no-impact proof instead of treating missing security evidence as no-impact. |
| TD-006 | Validator recognizes supported override keys: `extraPrompts`, `strictSourceCoverage`, `validatorMode`, and `categorySeverity`; applicability/contract PR 1 prose remains the source for safe override production. |
| TD-007 | Validator rejects accepted override entries that appear to weaken, disable, downgrade, bypass, or remove base requirements. |
| TD-008 | Positive fixture with valid schema fields, matrix rows, guideline/source refs, routing, and `validation.status: pass` exited `0`. |
| TD-009 | Negative fixture with missing categories, unsupported severity, incomplete no-impact proof, and failing validation metadata exited non-zero with actionable errors. |
| TD-010 | Validator contains the compact guideline map for `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001` with PR 2 Source IDs. |
| TD-011 | Validator fails unknown guideline IDs and missing Source IDs for applicable category rows; manual PR 2 catalog evidence covers complete compact-row source coverage. |
| TD-012 | Validator preserves operational severity vocabulary and fails unresolved blocking unknowns; downstream security-design/verify/archive instructions keep blocking obligations as archive gates. |
| TD-013 | `sdd-security-design` now requires true conditional obligations to become controls, evidence expectations, risks, or exceptions, and preserves advisory/conditional semantics downstream. |
| TD-014 | Validator allows only `blocking`, `conditional`, and `advisory`; the negative fixture using `Mayor` failed. |
| TD-015 | Validator requires artifact snapshot `security-guidelines-initial-user-snapshot-2026-06-30`, taxonomy version `1`, and Source IDs that resolve in the preserved snapshot ranges. |
| TD-016 | `sdd-security-design`, `sdd-verify`, and `sdd-archive` now preserve advisory obligations as downstream-visible risk/guidance/archive evidence instead of dropping non-blocking items. |
| TD-017 | `skills/sdd-security-design/SKILL.md` now consumes enriched catalog, source refs, operational severity, decision-matrix evidence, and validation metadata only when `securityImpact: true`. |
| TD-018 | `skills/sdd-security-design/SKILL.md` now requires true conditional obligations to carry predicate evidence downstream. |
| TD-019 | `sdd-security-design`, `sdd-test-design`, `sdd-tasks`, `sdd-verify`, and `sdd-archive` preserve valid no-impact skip compatibility when proof and validation metadata are complete. |
| TD-020 | Validator and downstream skills block incomplete no-impact proof or failed validation metadata instead of silently skipping `security-design.md`. |
| TD-021 | `scripts/validate_security_applicability.ps1` is a repository-local PowerShell static validator scoped to Markdown/YAML structure and references; it introduces no package-managed runner or runtime security scanner. |
| TD-022 | PR 1 additive schema compatibility remains intact; PR 3 downstream edits preserve `schemaVersion: 1`, artifact identity, existing `SEC-*` IDs, and no-impact routing. |
| TD-023 | Manual inspection confirms `sdd-security-design`, `sdd-test-design`, `sdd-tasks`, `sdd-verify`, and `sdd-archive` consume enriched metadata consistently without making `sdd-review` authoritative. |
| TD-024 | `skills/sdd-archive/SKILL.md` now requires archive evidence fields: catalog snapshot identity, guideline IDs, taxonomy categories, source refs, operational severity, evidence status, residual risks, and exception state. |
| TD-025 | `sdd-security-design`, `sdd-verify`, and `sdd-archive` continue to require complete approved exception fields before missing mandatory evidence can satisfy readiness. |
| TD-026 | Runtime test runner/linter/typechecker/formatter/coverage remain unavailable per `openspec/config.yaml`; PR 3 evidence used static PowerShell validation and manual inspection only, and downstream skills require unavailable tooling to be reported explicitly. |
| TD-027 | Advisory review slicing evidence remains the three stacked-to-main work units in this file; PR 3 stayed scoped to validator/downstream compatibility. |

Validator execution evidence: transient positive and negative fixture files were created under `openspec/changes/harden-security-applicability-contract/`, validated, then removed to keep the persisted slice limited to the validator script, downstream skills, and this task artifact. Command: `powershell -NoProfile -ExecutionPolicy Bypass -File scripts\validate_security_applicability.ps1 -Path <fixture>`. Results: positive fixture `positive_exit=0`; negative fixture `negative_exit=1` with failures for incomplete matrix, invalid severity `Mayor`, missing supported categories, incomplete no-impact proof, and `validation.status: fail`.

### Review Remediation Evidence

Blocking review finding `REV-CORP-096` reported that this change's own `security-applicability.md` still used the pre-hardened shape and failed `scripts/validate_security_applicability.ps1` because it lacked `catalog`, `categoryDecisionMatrix`, `noImpactProof`, and `validation` metadata. Remediation updated only `openspec/changes/harden-security-applicability-contract/security-applicability.md` with the hardened schema while preserving `classification: security-impacting` and `securityImpact: true`.

Validator command: `powershell -NoProfile -ExecutionPolicy Bypass -File scripts\validate_security_applicability.ps1 -Path openspec\changes\harden-security-applicability-contract\security-applicability.md`.

Result: `PASS: security applicability artifact is valid: openspec\changes\harden-security-applicability-contract\security-applicability.md`.
