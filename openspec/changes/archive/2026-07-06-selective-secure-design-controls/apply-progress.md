# Apply Progress: Selective Secure Design Controls

## Implementation Progress

**Change**: selective-secure-design-controls  
**Mode**: Standard (Strict TDD inactive; no runtime test runner configured)  
**Delivery**: auto-chain, stacked-to-main  
**Current Work Unit**: corrective no-YAML delta across revised tasks 2.1-5.3

### Completed Tasks

- [x] 1.1 Preserve earlier applied WU1-WU4 context in `apply-progress.md`: selective boundary, design/test-design planning, review-security authority, and initial source-spec sync.
- [x] 1.2 Treat prior `review-report.md` as stale for the old implementation shape; regenerate review evidence after corrective apply.
- [x] 2.1 Updated `skills/_shared/sdd-security-contract.md` so design owns narrative category rules only, while review-security owns schema, compact matrices, Source ID matrices, and exhaustive `N/A` decisions.
- [x] 2.2 Updated `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` to prohibit YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, and all-row `N/A` bookkeeping for new active design.
- [x] 3.1 Updated `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` to consume changed-surface rationale and applicable narrative rules only; design schema/matrix dependencies are rejected.
- [x] 3.2 Static/manual check covered changed design/test-design contracts for obsolete design-owned YAML/schema/matrix requirements; remaining matches are prohibitions, review-security ownership, historical compatibility, or unrelated orchestrator/review/verify/archive content outside this slice.
- [x] 4.1 Updated `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` to parse narrative design and own report schema, compact matrix, exact-once 155 Source ID expansion, missed-category blockers, unsafe-evidence blockers, and supported `N/A` validation.
- [x] 4.2 Preserved safe evidence, denial-by-default routing, warning visibility, and historical compatibility behavior across security review outputs.
- [x] 5.1 Synced revised deltas into `openspec/specs/sdd-design-workflow/spec.md`, `sdd-test-design-workflow/spec.md`, `sdd-review-security-workflow/spec.md`, `sdd-security-guideline-catalog/spec.md`, and `sdd-execution-persistence-contracts/spec.md`.
- [x] 5.2 Updated this `apply-progress.md` with corrective WU evidence, unavailable tooling notes from `openspec/config.yaml#testing`, read-back summaries, and safe grep results.
- [x] 5.3 Marked prior `review-report.md` stale for the previous implementation shape; general review must be regenerated next, followed by review-security, verify, and archive against the new narrative-design boundary.

### Files Changed

| File | Action | What Was Done |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modified | Replaced design-owned schema/matrix wording with a narrative secure-design contract; kept report schema/source-row example under review-security ownership. |
| `skills/sdd-design/SKILL.md` | Modified | Removed design YAML/schema/table/matrix instructions and replaced them with narrative category-rule output and explicit prohibitions. |
| `agents/sdd/sdd-design.md` | Modified | Aligned adapter prompt with narrative design output and review-security-owned matrices/`N/A` decisions. |
| `skills/sdd-test-design/SKILL.md` | Modified | Updated test-design to consume narrative rules only and reject design schema/matrix dependencies. |
| `agents/sdd/sdd-test-design.md` | Modified | Aligned adapter prompt with narrative-rule consumption and static/manual unavailable-tooling evidence. |
| `skills/sdd-review-security/SKILL.md` | Modified | Updated review-security to parse narrative design and own report schema, compact matrix, Source ID matrix, omission validation, blockers, and `N/A` evidence. |
| `agents/sdd/sdd-review-security.md` | Modified | Aligned adapter prompt with report-owned schema/matrices and narrative design parsing. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modified | Synced revised narrative secure-design requirements and prohibitions against YAML/schema/matrix/all-row `N/A` output. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Modified | Synced revised narrative-rule test-design consumption and rejection of design schema/matrix dependencies. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modified | Synced review-security report-owned schema/matrix/exhaustive validation and narrative design parsing requirements. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modified | Synced catalog boundary so design uses category guidance while review-security consumes the authoritative machine matrix. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modified | Synced persistence/status compatibility for narrative design plus review-security report semantics. |
| `openspec/changes/selective-secure-design-controls/tasks.md` | Modified | Marked revised tasks 2.1-5.3 complete. |
| `openspec/changes/selective-secure-design-controls/apply-progress.md` | Modified | Merged previous progress context with corrective no-YAML evidence. |
| `openspec/changes/selective-secure-design-controls/review-report.md` | Modified | Marked the previous general review report as stale/superseded for the old implementation shape. |

### Deviations from Design

None — implementation follows the revised narrative-design boundary. The only review-report mutation was a stale notice; the next SDD phase must regenerate full general review evidence.

### Test-Design Evidence

| Case | Evidence |
| --- | --- |
| TD-001 | Read-back of `openspec/changes/selective-secure-design-controls/design.md#secure-development-design` confirms the section is narrative headings/paragraphs only and contains no YAML, JSON, schema block, control table, compact matrix, Source ID matrix, or all-row `N/A` bookkeeping. |
| TD-002 | Read-back of `skills/sdd-design/SKILL.md` confirms new active designs require narrative category rules and explicitly prohibit YAML/JSON/schema/control-table/compact-matrix/Source ID matrix/machine-readable applicability/all-row `N/A` design output. |
| TD-003 | Read-back of `agents/sdd/sdd-design.md` confirms the adapter mirrors narrative-only design behavior and sends compact/source matrices plus `N/A` decisions to `sdd-review-security`. |
| TD-004 | Read-back of `skills/_shared/sdd-security-contract.md` confirms artifact ownership: catalog owns inventory, design owns narrative category rules, test-design owns planned checks, and review-security owns report schema/matrix validation. |
| TD-005 | Source specs for design, catalog, and persistence were synced to remove design YAML/schema/matrix/Source ID coverage requirements for new active changes; remaining machine-readable content is review-security, catalog, state, or historical compatibility scoped. |
| TD-006 | Read-back of `skills/sdd-test-design/SKILL.md` confirms test-design consumes changed-surface context and applicable narrative category rules only. |
| TD-007 | Read-back of `agents/sdd/sdd-test-design.md` confirms the adapter does not instruct agents to parse design schema/YAML/matrices or demand full Source ID coverage from design. |
| TD-008 | Read-back of `skills/sdd-review-security/SKILL.md` confirms report-owned schema, compact matrix, exact-once 155 Source ID expansion, `Yes`/`No`/`N/A` decisions, blockers, exceptions, and next recommendation. |
| TD-009 | Read-back of `agents/sdd/sdd-review-security.md` confirms design YAML/schema/matrices are not required and report-owned exhaustive validation remains in review-security. |
| TD-010 | Review-security skill/adapter now correlate narrative design, test-design, apply evidence, changed files, and review findings; missed applicable controls, unsupported `N/A`, unsafe evidence, and incomplete exceptions block. |
| TD-011 | Safe-evidence rules remain in the shared contract and review-security paths: raw PAN, PII, secrets, tokens, credentials, connection strings, private keys, and confidential payloads are prohibited in evidence. |
| TD-012 | Denial-by-default routing remains preserved for missed applicable controls, malformed reports, unsupported `N/A`, unsafe evidence, missing artifacts, and incomplete exceptions. |
| TD-013 | Warning-only evidence remains visible and may proceed only when mandatory evidence is complete; unavailable tooling is reported as unavailable, not passing. |
| TD-014 | Historical `security-design.md`, `security-applicability.md`, and exhaustive design rows remain compatibility data and are not active new-change dependencies or launchable successors. |
| TD-015 | `openspec/config.yaml#testing` confirms no test runner, build, lint, typecheck, formatter, or coverage command is configured. |
| TD-016 | Targeted grep/read-back found design/test-design matches only as prohibitions, review-security ownership, historical compatibility, or unrelated non-target artifacts; no active design/test-design contract requires YAML/schema/matrix/all-row `N/A` for new changes. |
| TD-017 | Catalog and review-security specs preserve expected Source ID count 155, exact-once report expansion, missing/duplicate/unknown/malformed/mapping/unsafe/unsupported-`N/A` blockers. |
| TD-018 | Advisory/warning-only evidence remains visible to review-security, verify, and archive when mandatory evidence is complete. |

### Security Evidence

- `SEC-DATA-001`: Safe evidence remains mandatory; evidence must cite paths, sections, summaries, masking/encryption decisions, or redacted examples rather than sensitive values.
- `SEC-SECRET-001`: No-secret evidence policy remains mandatory; artifacts must not commit, echo, log, or reproduce secret values.
- `SEC-ACCESS-001`: Denial-by-default routing explicitly includes missed applicable controls from narrative design omission plus malformed/source-row/unsupported-`N/A` blockers.
- `SEC-LOG-001`: Review/verify/archive evidence remains useful audit data without raw secrets, PAN, credentials, tokens, or unnecessary sensitive operational context.
- Review-security is the only active owner of machine-readable compact and Source ID matrices, exhaustive applicability decisions, and report-level `N/A` evidence.

### Static / Manual Verification

- Read back changed shared/design/test-design/review-security skill and adapter files after patching.
- Read back revised source specs under `openspec/specs/` for design, test-design, review-security, catalog, and execution-persistence contracts.
- Grep pattern over design/test-design targets: `YAML|schema|matrix|Source ID|N/A|compact|machine-readable|control table`.
- Grep classification: design/test-design matches are acceptable prohibitions, review-security ownership statements, or historical compatibility notes; no new active design/test-design requirement remains for design-owned YAML/schema/matrix/all-row `N/A` output.
- Read-back of `openspec/changes/selective-secure-design-controls/design.md#secure-development-design` confirms narrative-only content from line 85 onward: headings and prose rules only, no YAML/schema/matrix content.
- `review-report.md` now includes a stale notice at the top and must be regenerated by the next `sdd-review` phase.
- Runtime tools unavailable from `openspec/config.yaml`: test runner, build, lint, typecheck, formatter, coverage.

### Issues Found

None.

### Remaining Tasks

None in apply. Next phase must regenerate general review before security review/verify/archive.

### Workload / PR Boundary

- Mode: stacked PR slice
- Current work unit: corrective no-YAML delta across WU5-WU8 / revised tasks 2.1-5.3
- Boundary: Shared security contract, design/test-design/review-security skills and adapters, five source specs, tasks/apply-progress, and stale review marker only. No commits or PRs were created.
- Estimated review budget impact: Markdown-only corrective slice; original forecast remains high and uses auto-chain stacked-to-main.

### Status

11/11 revised tasks complete. Ready for regenerated general review (`sdd-review`).
