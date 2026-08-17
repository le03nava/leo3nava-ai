# Apply Progress: Active-Only Security Contract

## Workload / PR Boundary

| Field | Value |
| --- | --- |
| Mode | Standard apply; no Strict TDD |
| Delivery | Stacked PR slice to main |
| Current slice | Work Unit 3 / PR3 only: reader boundaries, source specs, and evidence |
| Completed tasks | 1.1, 1.2, 1.3, 2.1, 2.2, 2.3, 2.4, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6 |
| Out of scope | Archived OpenSpec changes; legacy security contract creation; runtime tooling introduction |

## Completed Tasks

- [x] 1.1 Updated `skills/_shared/sdd-security-contract.md` to remove active standalone security applicability schema/prose and active validator wording while preserving embedded secure design, review-security report, exception, lifecycle/status, and safe-evidence rules.
- [x] 1.2 Updated `skills/_shared/security-guideline-catalog.md` to keep snapshot metadata, all 8 SEC IDs, lifecycle/status vocabulary, and catalog boundaries centered on `design.md#secure-development-design` plus `review-security-report.md`.
- [x] 1.3 Performed summarized static safe-evidence checks on the changed shared files.
- [x] 2.1 Updated `skills/sdd-design/SKILL.md` so `## Secure Development Design` is the active authority, includes all 8 SEC IDs, and routes directly to `sdd-test-design` / `test-design` without producing a standalone security-design artifact.
- [x] 2.2 Updated `skills/sdd-review-security/SKILL.md` so `review-security-report.md` validates embedded design rows, blocks missing embedded design, and does not require `scripts/validate_security_design.ps1` or standalone security artifacts for new changes.
- [x] 2.3 Searched repo-local `skills/`, active workflow docs, and agent config surfaces for active `sdd-security-applicability` executor/skill or new-change DAG route.
- [x] 2.4 Manually read back changed phase docs for positive current-flow wording centered on embedded design and review-security evidence.
- [x] 3.1 Updated shared persistence/status contracts so historical `securityDesign` / `securityApplicability` refs are read-only data and not runnable phases, dependencies, successors, or active authority.
- [x] 3.2 Synced the five active source specs with the delta specs and removed active applicability override/static-validator requirements.
- [x] 3.3 Confirmed before and after apply changed paths do not include `openspec/changes/archive/**`.
- [x] 3.4 Confirmed no separate legacy security contract file was created.
- [x] 3.5 Reported unavailable runtime tests, lint, type-check, format, and coverage per `openspec/config.yaml#testing`; used static/manual evidence instead.
- [x] 3.6 Confirmed rollback remains a git revert of active contract/spec edits with no migration, runtime, or archive repair steps.

## Files Changed

| File | Action | What Was Done |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modified | Removed the active standalone applicability schema and compatibility note; narrowed vocabulary and evidence rules to the embedded design and review-security flow. |
| `skills/_shared/security-guideline-catalog.md` | Modified | Reframed catalog scope/usage/taxonomy/consistency wording around active embedded design and security review while preserving metadata and guideline rows. |
| `skills/sdd-design/SKILL.md` | Modified | Reframed design guidance so the embedded secure development section is the active authority, all 8 SEC IDs are required, and successful design routes directly to `sdd-test-design`. |
| `skills/sdd-review-security/SKILL.md` | Modified | Reframed security review guidance around embedded row validation, missing-section blocking, catalog/artifact evidence, and no active validator or standalone artifact dependency. |
| `skills/_shared/persistence-contract.md` | Modified | Narrowed historical security applicability/design resolver and state values to read-only data; removed them from active new-state phase/successor tokens. |
| `skills/_shared/sdd-status-contract.md` | Modified | Added explicit historical-token non-launchability and removed legacy tokens from new-change status successor enumeration. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modified | Synced active catalog scope and authority wording with embedded design plus review-security evidence. |
| `openspec/specs/sdd-security-applicability-workflow/spec.md` | Modified | Removed active applicability override/static-validator requirements and kept applicability as historical-read-only data. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modified | Synced report, embedded-design blocker, and retired validator requirements. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modified | Synced historical token/readability requirements and active artifact/status boundaries. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modified | Synced embedded secure design as active authority and direct routing/no-standalone-artifact requirements. |
| `openspec/changes/active-only-security-contract/tasks.md` | Modified | Preserved WU1/WU2 checkboxes and marked only assigned WU3 tasks 3.1-3.6 complete. |
| `openspec/changes/active-only-security-contract/apply-progress.md` | Modified | Merged WU3 apply evidence with prior WU1/WU2 evidence. |

## Test-Design Evidence

| Case | Evidence |
| --- | --- |
| TD-001 | Read-back of `sdd-security-contract.md` confirms active sections remain for embedded secure design, review-security report, approved exceptions, and safe-evidence rules. |
| TD-002 | Targeted search of the two WU1 shared files found no active standalone `security-design.md`, `security-applicability.md`, `validate_security_design`, legacy compatibility, or applicability schema references in the changed WU1 authority surfaces. |
| TD-003 | Read-back of `security-guideline-catalog.md` confirms snapshot ID `security-guidelines-initial-user-snapshot-2026-06-30`, catalog version `1`, taxonomy version `1`, matrix vocabulary, lifecycle status vocabulary, and all 8 compact SEC IDs remain present. |
| TD-004 | Catalog boundary text now favors `design.md#secure-development-design` plus `review-security-report.md`; `sdd-review` remains a citation surface, not a duplicate authority. |
| TD-005 | Read-back of `skills/sdd-review-security/SKILL.md` confirms `review-security-report.md` validates embedded `design.md#secure-development-design` rows with row-level evidence, observations, blockers, exceptions, and next routing. |
| TD-006 | Decision gates in `skills/sdd-review-security/SKILL.md` block when `design.md#secure-development-design` is missing or unreadable and keep verify/archive unavailable until embedded design evidence exists. |
| TD-007 | `skills/sdd-review-security/SKILL.md` now states active new-change validation uses catalog, embedded rows, and artifact evidence; it does not invoke or require `scripts/validate_security_design.ps1` or standalone security artifacts. |
| TD-008 | `glob` found no repo-local `agents/sdd/*sdd-security-applicability*` or `skills/sdd-security-applicability/**`; active agent/orchestrator read-back scopes `sdd-security-applicability` to retired old/archive data only. |
| TD-009 | Active workflow read-back confirms `README.md`, `agents/sdd/sdd-orchestrator.md`, status contracts, and source specs route new changes through `design -> test-design` and do not offer `sdd-security-applicability` in the active DAG. |
| TD-014 | Read-back of `skills/sdd-design/SKILL.md` confirms `## Secure Development Design` is the active security authority and the template/rules list all 8 SEC IDs. |
| TD-015 | Read-back of `skills/sdd-design/SKILL.md` confirms successful design returns `next_recommended: test-design` and routes directly to `sdd-test-design` after persistence without standalone security-design production. |
| TD-010 | Read-back of persistence/status contracts confirms historical `security-applicability` data may be resolved/displayed but must not be required, rerun, or normalized into an active successor. |
| TD-011 | Active source spec sync removed `Supported Applicability Overrides` and `Static Applicability Validator` requirements from `openspec/specs/sdd-security-applicability-workflow/spec.md`; active classification remains in `design.md#secure-development-design`. |
| TD-012 | Read-back of `persistence-contract.md` and `sdd-status-contract.md` confirms historical `security-design` / `security-applicability` values are compatibility data only and must not map to runnable phases, launchable agents, active authority, or required successors. |
| TD-013 | Source specs and shared contracts keep new-change dependencies on embedded `design.md` rows and `review-security-report.md`; standalone security artifacts remain historical refs only. |
| TD-016 | `git status --short` before apply showed only active WU1/WU2 files plus the active change folder; after apply, changed paths still contain no `openspec/changes/archive/**` entries. |
| TD-017 | Targeted search for `legacy-security-contract`, `security-design-contract`, `security-applicability-contract`, and `legacy security contract` found no new active contract file; matches were only existing proposal/test/task wording or archived historical change names. |
| TD-018 | `openspec/config.yaml#testing` states no test runner, coverage, linter, type checker, or formatter commands are configured; this slice uses static/manual read-back only and does not claim automated checks passed. |
| TD-019 | Static safe-evidence scan over changed shared files found zero matches for private key blocks, AWS access-key shape, JWT-like tokens, quoted secret assignments, PAN-like 13-19 digit sequences, or email addresses. |
| TD-020 | Manual read-back of changed phase docs confirmed positive current-flow wording: active authority in embedded design, review-security row validation, direct test-design routing, and no validator/standalone-artifact dependency for new changes. |
| TD-021 | Rollback remains `git revert` of active contract/spec and active change artifact edits only; no migration, runtime, or archive repair steps were introduced. |

## Security Evidence

| Guideline | Evidence status |
| --- | --- |
| `SEC-DATA-001` | Implemented for WU1 docs: safe-evidence text preserved and static scan reported no raw PAN/PII/confidential payload indicators. |
| `SEC-SECRET-001` | Implemented for WU1 docs: safe-evidence text preserves no raw credentials/tokens/private keys; static scan found no secret/token/private-key indicators. |
| `SEC-ACCESS-001` | Implemented across WU1-WU3: active contract/catalog authority favors embedded design and review-security; design/review-security/status/persistence/source specs keep new-change launch authority on embedded design rows, review-security evidence, and direct test-design routing. Historical `securityDesign` / `securityApplicability` refs remain read-only data only. |
| `SEC-LOG-001` | Implemented for WU1 docs: evidence/reporting guidance keeps audit records useful without raw sensitive payloads; static scan found no unsafe evidence indicators. |
| `SEC-AUTH-001` | WU2 remains documentation/workflow-only: changed phase docs do not alter login, identity, MFA, recovery, or credential-flow behavior. |
| `SEC-SESS-001` | WU2 remains documentation/workflow-only: changed phase docs do not alter cookies, tokens, sessions, or revocation behavior. |
| `SEC-FILE-001` | WU2 changed Markdown phase guidance only; no runtime upload/download/generated-file behavior changed. |
| `SEC-DB-001` | WU2 changed Markdown phase guidance only; no database, query, migration, or persistence engine behavior changed. |

## Tooling Availability

Per `openspec/config.yaml#testing`, no automated test runner, linter, type checker, formatter, or coverage command is configured. This apply slice used static/manual checks only and does not claim automated tests passed.

| Capability | Configured | Evidence |
| --- | --- | --- |
| Runtime tests | No | `testing.test_runner.available: false`, command empty. |
| Lint | No | `testing.quality.linter.available: false`, command empty. |
| Type-check | No | `testing.quality.type_checker.available: false`, command empty. |
| Format | No | `testing.quality.formatter.available: false`, command empty. |
| Coverage | No | `testing.coverage.available: false`, command empty. |

## Deviations

None. The implementation stayed inside the assigned WU3 file boundary and followed the active-only design.

## Remaining Tasks

- None. Assigned WU3 tasks 3.1-3.6 are complete.

## Issues Found

- `rg` was not used for this slice; targeted content checks used the provided grep/read/glob tools and `git status`/`git diff --name-only` for changed-path evidence.

## Review Remediation: REV-CORP-096

| Field | Evidence |
| --- | --- |
| Blocking finding | `REV-CORP-096`: active orchestrator state schema examples still exposed `security-design` as active `currentPhase` / `nextRecommended` tokens. |
| Source remediation | Updated `agents/sdd/sdd-orchestrator.md` minimum state schema so active `currentPhase` and `nextRecommended` enumerations no longer list `security-design`. |
| Historical scope | Kept `artifactRefs.securityDesign` only as `legacy/archive compatibility only; read-only for new changes`, and added explicit historical compatibility text for old or archived `security-design` state reads. |
| Validation | Read-back confirmed active schema examples omit `security-design`; targeted grep for `currentPhase:.*security-design` and `nextRecommended:.*security-design` now matches only the historical compatibility paragraph, not active emitted schema enumerations. |
| Tooling | Runtime tests, lint, type-check, format, and coverage remain unavailable per `openspec/config.yaml#testing`; remediation evidence is static/manual read-back. |
| Scope control | No archive folders, legacy contract files, or broader source areas were edited. |
