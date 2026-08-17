# Apply Evidence WU2: Design and Test-Design Contracts

## Scope

WU2 updates only the design and test-design phase contracts plus their adapter prompts for `slim-secure-design-artifacts`.

Completed tasks:

- 2.1 Update `skills/sdd-design/SKILL.md` and `agents/sdd/sdd-design.md` to require slim `design.md#secure-development-design` references and prohibit 155-row or 96-control duplication.
- 2.2 Update `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` to plan grouped static/manual checks from catalog references and report unavailable runtime tooling.
- 2.3 Record WU2 evidence covering TD-001, TD-002, TD-004, TD-006-TD-013, safe evidence, no legacy standalone dependency, and no runtime/tooling changes.

Out of scope for this slice: review-security ownership edits, verify/archive edits, source spec sync, PR creation, runtime code changes, validation scripts, and legacy standalone `security-design.md` / `security-applicability.md` restoration.

## Static / Manual Evidence

| Evidence ID | Test-design case / control | Evidence | Result |
| --- | --- | --- | --- |
| WU2-TD-001 | TD-001 artifact boundary | `skills/sdd-design/SKILL.md` now states design owns slim catalog-referenced coverage while `review-security` owns full expansion; `skills/sdd-test-design/SKILL.md` plans checks from that slim contract. | Pass |
| WU2-TD-002 | TD-002 no exhaustive duplication | Design skill and design adapter prompt explicitly prohibit duplicating the general 96-control review matrix and the exhaustive 155-row Source ID matrix in design. Test-design skill and prompt prohibit requiring or duplicating the full 155-row matrix. | Pass |
| WU2-TD-004 | TD-004 expected count preservation | Design skill requires catalog snapshot/path and `expectedSourceIdCount: 155`; test-design skill consumes and plans evidence from the same expected count. | Pass |
| WU2-TD-006 | TD-006 authentication/password groups | Design skill requires grouped section references, counts, compact mappings, lifecycle status, and evidence owners; test-design skill plans source groups or catalog-backed references without per-row duplication. | Pass |
| WU2-TD-007 | TD-007 logging/error groups | Test-design contract preserves warning-only source coverage and safe observation evidence; design contract preserves `SEC-LOG-001` safe-evidence expectations through compact controls and grouped references. | Pass |
| WU2-TD-008 | TD-008 data/secrets/PAN groups | Design and test-design contracts preserve safe-evidence policy and require review-safe paths, sections, summaries, or redacted placeholders instead of sensitive raw values. | Pass |
| WU2-TD-009 | TD-009 database/input/output groups | Test-design contract requires planned checks for source-row groups, compact mappings, or catalog-backed references; no runtime DB code is introduced. | Pass |
| WU2-TD-010 | TD-010 coding/session/file/access groups | Design and test-design contracts preserve lifecycle state, evidence owners, blocker routing visibility, and `review-security` full-expansion ownership. | Pass |
| WU2-TD-011 | TD-011 grouped planning coverage | Test-design contract requires every mandatory source-row group, compact mapping, or catalog-backed reference to have static/manual/automated evidence or a justified non-test plan. | Pass |
| WU2-TD-012 | TD-012 unavailable tooling | `skills/sdd-test-design/SKILL.md` and `agents/sdd/sdd-test-design.md` require unavailable runtime/build/lint/type/format/coverage tooling to be reported as unavailable evidence, not passing evidence. | Pass |
| WU2-TD-013 | TD-013 N/A and warnings | Design and test-design contracts preserve evidence-backed `N/A` policy and warning-only evidence preservation; unsupported `N/A` remains blocking. | Pass |
| WU2-LEGACY | No legacy standalone dependency | Design and test-design skills continue to treat standalone `security-design.md` / `security-applicability.md` as non-required legacy compatibility data for new changes. | Pass |
| WU2-NO-RUNTIME | Runtime/tooling non-change | Changed files are Markdown instruction contracts and OpenSpec apply evidence only; no package/build/test config, runtime source, or validation script was created. | Pass |

## Security Control Evidence

| Guideline | WU2 evidence |
| --- | --- |
| `SEC-AUTH-001` | Design/test-design now require authentication-related Source ID coverage by catalog snapshot/path, grouped references, expected count, compact mappings, and review-security exact-once handoff. |
| `SEC-SESS-001` | Session coverage remains catalog-owned, grouped, planned in test-design, and expanded only by review-security. |
| `SEC-DATA-001` | Safe-evidence policy is preserved for sensitive data/PAN groups; raw sensitive values are not requested in design/test-design artifacts. |
| `SEC-SECRET-001` | Secret/config evidence remains reference-only and unsafe evidence remains blocking. |
| `SEC-ACCESS-001` | Missing group coverage, missing mappings, unsupported `N/A`, unsafe evidence, and mandatory evidence gaps remain blocking in design/test-design. |
| `SEC-FILE-001` | File-based artifacts carry grouped references and summaries only; full matrix materialization is deferred to `review-security-report.md`. |
| `SEC-DB-001` | Database/input/output source groups remain planned by catalog-backed references without runtime DB changes. |
| `SEC-LOG-001` | Warning-only and logging/error evidence remain tracked as review-safe observations. |

## Tooling Evidence

`openspec/config.yaml#testing` reports no runtime test runner, build command, coverage command, linter, type checker, or formatter. Those tools are unavailable and are not treated as passing evidence. WU2 validation used static/manual artifact read-back only.

## Changed Files Read-Back

- `skills/sdd-design/SKILL.md` — read back after edit; contains slim catalog-referenced source-row coverage, `expectedSourceIdCount: 155`, grouped references, full-expansion owner, and no 96-control/155-row matrix duplication rules.
- `agents/sdd/sdd-design.md` — read back after edit; adapter prompt mirrors the slim design boundary and review-security expansion ownership.
- `skills/sdd-test-design/SKILL.md` — read back after edit; contains grouped/catalog-backed planning rules, unavailable-tooling reporting, warning/N/A preservation, and no full-matrix duplication requirement.
- `agents/sdd/sdd-test-design.md` — read back after edit; adapter prompt mirrors grouped static/manual planning and unavailable-tooling evidence semantics.
- `openspec/changes/slim-secure-design-artifacts/tasks.md` — read back after WU2 checkbox update.
- `openspec/changes/slim-secure-design-artifacts/apply-progress.md` — read back after cumulative progress merge.

## Deviations

None. WU2 follows the assigned design/test-design boundary and does not implement WU3-WU4.
