# Apply Evidence WU3: Review-Security Ownership

## Scope

WU3 updates only the review-security phase contract and adapter prompt for `slim-secure-design-artifacts`.

Completed tasks:

- 3.1 Update `skills/sdd-review-security/SKILL.md` and `agents/sdd/sdd-review-security.md` so only `review-security-report.md` materializes all 155 Source IDs exactly once.
- 3.2 Preserve general review citation-only behavior: security review may cite `review-report.md` rows but must not duplicate the 96-control matrix.
- 3.3 Record WU3 evidence for TD-014, TD-015, duplicate/missing/unknown Source ID blockers, and unsafe evidence blockers.

Out of scope for this slice: verify/archive edits, source spec sync, PR creation, runtime code changes, validation scripts, and legacy standalone `security-design.md` / `security-applicability.md` restoration.

## Static / Manual Evidence

| Evidence ID | Test-design case / control | Evidence | Result |
| --- | --- | --- | --- |
| WU3-TD-014 | TD-014 exclusive source-row expansion | `skills/sdd-review-security/SKILL.md` now states that `review-security-report.md` is the only active new-change artifact that materializes the exhaustive corporate Source ID matrix. | Pass |
| WU3-TD-014 | TD-014 expected count and exact-once rule | The review-security source-row contract and report format require expanding the catalog inventory, confirming `sourceRowExpectedCount: 155`, and materializing every expected Source ID exactly once in `review-security-report.md`. | Pass |
| WU3-TD-014 | TD-014 no duplicate ownership | The review-security boundary states that design, test-design, apply evidence, verify, archive, and general review artifacts may cite catalog refs, grouped coverage, summaries, evidence links, warnings, exceptions, or report rows only. | Pass |
| WU3-TD-015 | TD-015 general review citation-only | The review-security boundary, report format, and adapter prompt preserve `review-report.md` as supporting evidence only and prohibit duplicating or recreating the 96-control general review matrix. | Pass |
| WU3-SOURCE-BLOCKERS | Missing/duplicate/unknown Source ID blockers | Decision gates and routing rules preserve `resolve-blockers` routing for missing, duplicate, unknown, malformed, or unmapped Source IDs. | Pass |
| WU3-UNSAFE-EVIDENCE | Unsafe evidence blockers | Decision gates and source-row review rules reject secrets, credentials, PAN, PII, tokens, connection strings, private keys, or confidential values and route unsafe evidence to blockers unless implementation remediation is required. | Pass |
| WU3-ADAPTER | Adapter prompt alignment | `agents/sdd/sdd-review-security.md` now mirrors the exclusive 155-row report ownership, exact-once materialization, citation-only general review handoff, and duplicate/unsafe evidence blockers. | Pass |

## Security Control Evidence

| Guideline | WU3 evidence |
| --- | --- |
| `SEC-AUTH-001` | Review-security owns exact-once authentication Source ID expansion and validation from catalog mappings. |
| `SEC-SESS-001` | Session Source IDs are materialized only in `review-security-report.md` and validated against design/test-design/apply evidence. |
| `SEC-DATA-001` | Safe-evidence rules reject raw sensitive data, PAN, PII, credentials, tokens, and confidential values in review-security evidence. |
| `SEC-SECRET-001` | Secret/config evidence remains reference-only; unsafe evidence blocks source-row validation. |
| `SEC-ACCESS-001` | Missing, duplicate, unknown, malformed, unmapped, unsupported N/A, and unsafe source-row evidence remains denial-by-default. |
| `SEC-FILE-001` | File-based artifacts outside `review-security-report.md` may carry only references, summaries, and evidence links instead of the full 155-row matrix. |
| `SEC-DB-001` | Database/input/output Source IDs stay catalog-backed and are validated in the review-security report. |
| `SEC-LOG-001` | General review and logging/error evidence remain citation-only and review-safe; the 96-control matrix is not duplicated. |

## Tooling Evidence

`openspec/config.yaml#testing` reports no runtime test runner, build command, coverage command, linter, type checker, or formatter. Those tools are unavailable and are not treated as passing evidence. WU3 validation used static/manual artifact read-back only.

## Changed Files Read-Back

- `skills/sdd-review-security/SKILL.md` — read back after edit; contains exclusive `review-security-report.md` ownership for the 155-row Source ID matrix, exact-once count rules, blocker routing, safe-evidence blockers, and citation-only general review handoff.
- `agents/sdd/sdd-review-security.md` — read back after edit; adapter prompt mirrors exclusive source-row materialization and no 96-control duplication.
- `openspec/changes/slim-secure-design-artifacts/tasks.md` — read back after WU3 checkbox update.
- `openspec/changes/slim-secure-design-artifacts/apply-progress.md` — read back after cumulative progress merge.

## Deviations

None. WU3 follows the assigned review-security boundary and does not implement WU4.
