# Apply Evidence WU4: Verify, Archive, Source Specs, Evidence

## Scope

Work Unit 4 updates verify/archive source-row preservation contracts, syncs accepted delta requirements into source OpenSpec specs, and records static/manual evidence for TD-016 through TD-020.

## Static / Manual Evidence

| Check | Evidence | Result |
| --- | --- | --- |
| TD-016 Verify Source Row Consumption | `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` now require verify to consume non-blocking `review-security-report.md` source-row evidence, preserve catalog snapshot/path, expected count, compact mappings, warnings, exceptions, evidence refs, and report links, and block unresolved source-row blockers without owning the full matrix. | Pass |
| TD-017 Archive Source Row Preservation | `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` now require archive to preserve summaries, catalog identity/path, expected count, compact mappings, warnings, exceptions, safe evidence refs, `N/A` status, and report links without requiring legacy standalone security artifacts or copying the full matrix. | Pass |
| TD-018 Proposal out of scope | Changed files are Markdown instruction/spec/evidence artifacts only. No runtime application code, package/build configuration, validation scripts, `security-design.md`, or `security-applicability.md` active dependency was introduced. | Pass |
| TD-019 Safe evidence | Evidence uses paths, section names, summaries, and sanitized observations only. No secrets, credentials, PAN, PII, tokens, private keys, connection strings, or confidential values were added. | Pass |
| TD-020 Warnings and advisory evidence | Verify/archive contracts preserve warning-only source-row findings and allow forward routing only when mandatory evidence is complete and no source-row blockers remain. | Pass |

## Source Spec Sync Evidence

Accepted requirements from `openspec/changes/slim-secure-design-artifacts/specs/*/spec.md` were synced into source specs:

| Source Spec | Synced Requirement Evidence |
| --- | --- |
| `openspec/specs/sdd-design-workflow/spec.md` | Slim design source coverage by catalog snapshot/path, `expectedSourceIdCount: 155`, grouped summaries, compact mappings, and no 96-control/155-row duplication. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Grouped source-row planning from slim design coverage, unavailable-runner handling, N/A/warning evidence preservation, and no full matrix requirement in design. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Exclusive `review-security-report.md` ownership for exhaustive exact-once Source ID materialization and citation-only general review behavior. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Catalog ownership of the authoritative 155 Source ID inventory, snapshot metadata, range expansion, slim artifact references, and shared ownership boundary. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Verify/archive preservation of catalog identity/count/mappings/warnings/exceptions/links without full matrix duplication or legacy standalone artifact dependency. |

## Tooling Availability

Per `openspec/config.yaml#testing`, no test runner, build command, linter, type checker, formatter, or coverage command is available. These are unavailable evidence, not passing evidence. WU4 therefore used static/manual read-back and changed-file inspection.

| Tooling Area | Available | Evidence Used |
| --- | --- | --- |
| Runtime tests | No | Static/manual contract inspection. |
| Build | No | Not applicable; no buildable runtime application exists. |
| Lint | No | Manual Markdown/frontmatter inspection. |
| Type check | No | Not applicable; no typed runtime source changed. |
| Format | No | Manual Markdown structure/readability inspection. |
| Coverage | No | Source-row coverage boundary checklist. |

## Changed-File Review

| File | Evidence Summary |
| --- | --- |
| `skills/sdd-verify/SKILL.md` | Verify consumes source-row summary evidence and preserves catalog/count/mapping/warning/exception/link evidence without owning the full matrix. |
| `agents/sdd/sdd-verify.md` | Adapter prompt mirrors verify preservation boundary and no-matrix-copy rule. |
| `skills/sdd-archive/SKILL.md` | Archive preserves source-row audit trail fields and avoids copying the full matrix. |
| `agents/sdd/sdd-archive.md` | Adapter prompt mirrors archive catalog/report-link preservation and no full matrix duplication. |
| `openspec/specs/*/spec.md` | Source specifications now include accepted slim secure-design artifact requirements. |
| `openspec/changes/slim-secure-design-artifacts/tasks.md` | WU4 checkboxes marked complete. |
| `openspec/changes/slim-secure-design-artifacts/apply-progress.md` | Cumulative WU1-WU4 progress preserved and updated. |
| `openspec/changes/slim-secure-design-artifacts/state.yaml` | State routes completed apply work to review. |

## Security Evidence

- `SEC-ACCESS-001`: verify/archive blocker routing remains denial-by-default for unresolved source-row blockers.
- `SEC-FILE-001`: file-based OpenSpec and skill artifacts preserve references and summaries without copying the full source-row matrix.
- `SEC-LOG-001`, `SEC-DATA-001`, `SEC-SECRET-001`: evidence remains review-safe and avoids raw sensitive values.
- `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DB-001`: compact mappings remain catalog-owned and preserved through verify/archive summaries and report links.

## Deviations

None. WU4 matches the planned static/manual evidence strategy. Runtime/build/lint/type/format/coverage commands are unavailable and were reported explicitly.

## Result

WU4 static/manual evidence passes. All 12 implementation tasks are now complete and the change is ready for `sdd-review`.
