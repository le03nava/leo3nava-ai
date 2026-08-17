# Apply Progress: Slim Secure Design Artifacts

## Implementation Progress

**Change**: `slim-secure-design-artifacts`  
**Mode**: Standard  
**Artifact store**: OpenSpec  
**Delivery strategy**: auto-chain  
**Chain strategy**: stacked-to-main  
**Current slice**: WU4 verify/archive/source specs/evidence only

### Completed Tasks

- [x] 1.1 Update `skills/_shared/sdd-security-contract.md` to state catalog/design/test-design/review-security/verify/archive ownership, blocker routing, safe evidence, N/A, warning, and exception rules.
- [x] 1.2 Update `skills/_shared/security-guideline-catalog.md` to declare the authoritative 155 Source ID inventory, snapshot metadata, expanded ranges, compact mappings, and expected count.
- [x] 1.3 Record WU1 static evidence for TD-003, TD-005, TD-011, and SEC-AUTH/SESS/DATA/SECRET/ACCESS/FILE/DB/LOG coverage.
- [x] 2.1 Update `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` to require slim `design.md#secure-development-design` references and prohibit 155-row or 96-control duplication.
- [x] 2.2 Update `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` to plan grouped static/manual checks from catalog references and report unavailable runtime tooling.
- [x] 2.3 Record WU2 evidence covering TD-001, TD-002, TD-004, TD-006-TD-013, safe evidence, no legacy standalone dependency, and no runtime/tooling changes.
- [x] 3.1 Update `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` so only `review-security-report.md` materializes all 155 Source IDs exactly once.
- [x] 3.2 Preserve general review citation-only behavior: security review may cite `review-report.md` rows but must not duplicate the 96-control matrix.
- [x] 3.3 Record WU3 evidence for TD-014, TD-015, duplicate/missing/unknown Source ID blockers, and unsafe evidence blockers.
- [x] 4.1 Update `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` to consume non-blocking review-security evidence without owning the full matrix.
- [x] 4.2 Update `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` to preserve source-row summaries, catalog identity, count, mappings, warnings, exceptions, and links.
- [x] 4.3 Sync accepted requirements into `openspec/specs/sdd-design-workflow/spec.md`, `sdd-test-design-workflow/spec.md`, `sdd-review-security-workflow/spec.md`, `sdd-security-guideline-catalog/spec.md`, and `sdd-execution-persistence-contracts/spec.md` when archive applies the delta.
- [x] 4.4 Record WU4 evidence for TD-016-TD-020, unavailable tests/lint/type/format/coverage, changed-file review, and archive-safe references.

### Files Changed

| File | Action | What Was Done |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modified | Added artifact ownership boundary and slim secure-design source-row contract; clarified catalog/design/test-design/apply/review-security/verify/archive ownership, routing, safe evidence, N/A, warning, and exception semantics. |
| `skills/_shared/security-guideline-catalog.md` | Modified | Declared authoritative 155 Source ID inventory ownership, snapshot metadata, expected count, inventory ownership contract, expanded range contract, compact mapping vocabulary, and canonical section count total. |
| `openspec/changes/slim-secure-design-artifacts/apply-evidence-wu1.md` | Created | Recorded WU1 static/manual evidence for TD-003, TD-005, TD-011, compact security coverage, tooling unavailability, and changed-file read-back. |
| `skills/sdd-design/SKILL.md` | Modified | Required slim catalog-referenced secure-design source-row coverage, `expectedSourceIdCount: 155`, grouped references, review-security full-expansion ownership, and no 96-control/155-row matrix duplication. |
| `agents/sdd/sdd-design.md` | Modified | Aligned design adapter prompt with slim source-row references and review-security expansion ownership. |
| `skills/sdd-test-design/SKILL.md` | Modified | Required grouped/catalog-backed source-row planning, unavailable-tooling reporting, warning/N/A preservation, and no full 155-row matrix requirement. |
| `agents/sdd/sdd-test-design.md` | Modified | Aligned test-design adapter prompt with grouped static/manual evidence and unavailable-tooling semantics. |
| `openspec/changes/slim-secure-design-artifacts/apply-evidence-wu2.md` | Created | Recorded WU2 static/manual evidence for TD-001, TD-002, TD-004, TD-006-TD-013, security controls, no legacy standalone dependency, and no runtime/tooling changes. |
| `skills/sdd-review-security/SKILL.md` | Modified | Made `review-security-report.md` the exclusive active artifact that materializes the exhaustive 155 Source ID matrix exactly once; preserved source-row blocker routing, safe-evidence rejection, and citation-only general review behavior. |
| `agents/sdd/sdd-review-security.md` | Modified | Aligned review-security adapter prompt with exclusive 155-row report ownership, exact-once materialization, and no 96-control general review duplication. |
| `openspec/changes/slim-secure-design-artifacts/apply-evidence-wu3.md` | Created | Recorded WU3 static/manual evidence for TD-014, TD-015, duplicate/missing/unknown Source ID blockers, unsafe evidence blockers, security controls, and unavailable tooling. |
| `skills/sdd-verify/SKILL.md` | Modified | Required verify to consume non-blocking source-row evidence, preserve catalog/count/mapping/warning/exception/link evidence, block unresolved source-row blockers, and avoid full matrix ownership. |
| `agents/sdd/sdd-verify.md` | Modified | Aligned verify adapter prompt with source-row preservation boundary and no full matrix duplication. |
| `skills/sdd-archive/SKILL.md` | Modified | Required archive to preserve source-row summaries, catalog identity/path, expected count, mappings, warnings, exceptions, safe evidence refs, N/A status, and report links without copying the full matrix. |
| `agents/sdd/sdd-archive.md` | Modified | Aligned archive adapter prompt with catalog/report-link preservation and no full matrix duplication. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modified | Synced accepted slim secure-design source coverage requirements into source specs. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Modified | Synced accepted grouped source-row test planning requirements into source specs. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modified | Synced accepted exclusive review-security expansion requirements into source specs. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modified | Synced accepted catalog inventory ownership and shared ownership-boundary requirements into source specs. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modified | Synced accepted verify/archive source-row preservation requirements into source specs. |
| `openspec/changes/slim-secure-design-artifacts/apply-evidence-wu4.md` | Created | Recorded WU4 static/manual evidence for TD-016-TD-020, unavailable tooling, changed-file review, and archive-safe references. |
| `openspec/changes/slim-secure-design-artifacts/apply-progress.md` | Modified | Merged cumulative apply progress for WU1 through WU4. |
| `openspec/changes/slim-secure-design-artifacts/tasks.md` | Modified | Preserved WU1-WU3 checkboxes and marked only tasks 4.1-4.4 complete. |
| `openspec/changes/slim-secure-design-artifacts/state.yaml` | Modified | Routed completed apply work to review. |

### Deviations from Design

None — implementation matches the completed WU1-WU3 slices and the assigned WU4 verify/archive/source-spec/evidence slice from `design.md` and `test-design.md`.

### Test-Design Evidence

| Case | Evidence |
| --- | --- |
| TD-003 | Catalog declares snapshot metadata, inventory authority, expected count `155`, section grouping, expanded Source IDs, and compact mappings. |
| TD-005 | Shared contract and catalog restrict compact mapping vocabulary to the eight existing `SEC-*` IDs and preserve blocker routing for unknown/missing mappings. |
| TD-011 | Shared contract requires source-row planning by catalog reference, grouped coverage, compact mappings, owner phases, lifecycle state, and exact-once review-security handoff. |
| TD-001 | Design and test-design phase contracts now preserve the artifact boundary: catalog owns inventory, design stays slim, test-design plans grouped evidence, review-security expands, and verify/archive consume downstream evidence. |
| TD-002 | Design and test-design instructions prohibit duplicating the general 96-control review matrix and the exhaustive 155-row Source ID matrix. |
| TD-004 | Design and test-design instructions preserve catalog snapshot/path and `expectedSourceIdCount: 155`. |
| TD-006-TD-010 | Grouped source-row planning is required by catalog-backed source section/group references, compact mappings, counts, lifecycle status, evidence owners, safe evidence, and exact-once review-security handoff. |
| TD-012 | Test-design instructions report unavailable runtime/build/lint/type/format/coverage tooling explicitly and require static/manual substitutes. |
| TD-013 | Design and test-design instructions preserve evidence-backed `N/A` checks and warning-only evidence tracking. |
| TD-014 | Review-security instructions and adapter prompt now require `review-security-report.md` to be the only active artifact that materializes all 155 expected Source IDs exactly once from the catalog inventory. |
| TD-015 | Review-security instructions and adapter prompt preserve citation-only handoff from `review-report.md` and prohibit copying or recreating the 96-control general review matrix. |
| TD-016 | Verify instructions and adapter prompt consume non-blocking source-row evidence, preserve catalog/count/mapping/warning/exception/report-link evidence, and avoid full source-row matrix ownership. |
| TD-017 | Archive instructions and adapter prompt preserve source-row summaries, catalog identity/path, expected count, compact mappings, warnings, exceptions, evidence refs, and report links without copying the full matrix. |
| TD-018 | Changed-file review confirms no runtime application code, package/build config, validation script, or legacy active standalone security artifact was introduced. |
| TD-019 | Static evidence remains review-safe and contains no raw secrets, credentials, PAN, PII, tokens, private keys, connection strings, or confidential values. |
| TD-020 | Verify/archive contracts preserve warning-only findings and allow forward routing only when mandatory evidence is complete and no source-row blockers remain. |
| Source-row blockers | Review-security, verify, and archive decision gates preserve blocker routing for missing, duplicate, unknown, malformed, unmapped, unsupported `N/A`, and unsafe Source ID evidence. |

Runtime/build/lint/type/format/coverage tooling is unavailable per `openspec/config.yaml#testing`; WU1 through WU4 used static/manual read-back evidence.

### Security Evidence

- `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001` remain valid compact mappings in the catalog.
- `skills/_shared/sdd-security-contract.md` states catalog authority, design/test-design slim reference boundaries, review-security exact-once expansion ownership, and verify/archive preservation rules.
- `skills/_shared/security-guideline-catalog.md` explicitly owns the 155 concrete Source ID inventory and maps each expanded section to compact controls.
- Safe evidence policy remains review-safe: paths, anchors, summaries, command summaries, sanitized observations, or redacted placeholders only.
- `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` require slim catalog-referenced design coverage and prohibit 96-control/155-row matrix duplication.
- `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` plan grouped/catalog-backed static/manual evidence, report unavailable tooling as unavailable evidence, preserve warning/N/A semantics, and avoid full matrix duplication.
- `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` require only `review-security-report.md` to materialize all 155 expected Source IDs exactly once.
- `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` consume source-row security review evidence, preserve catalog/count/mapping/warning/exception/link evidence, and do not own the full row matrix.
- `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` preserve source-row audit summaries and report links without requiring legacy standalone artifacts or copying the full matrix.
- General review evidence remains citation-only: `review-report.md` rows may support source-row evidence, but the 96-control matrix must not be duplicated.
- Unsafe evidence containing secrets, credentials, PAN, PII, tokens, connection strings, private keys, or confidential values is rejected and routed to blockers unless implementation remediation is required.

### Issues Found

None.

### Remaining Tasks

None. All 12 implementation tasks are complete.

### Workload / PR Boundary

- Mode: stacked PR slice.
- Current work unit: WU4 verify/archive/source specs/evidence.
- Boundary: starts from completed WU1 shared contract/catalog, WU2 design/test-design contracts, and WU3 review-security ownership; ends with verify/archive preservation contracts, source OpenSpec spec sync, WU4 evidence, and apply routing to review.
- Estimated review budget impact: autonomous docs/contract slice intended to remain below the 400-line review budget; no PR was created by this apply executor.

### Status

12/12 tasks complete. Ready for `sdd-review`.
