# Review Security Report Template

This file is the derived Markdown presentation contract for `sdd-review-security`. The canonical security-review artifact is `review-security-report.json`; `review-security-report.md` / `sdd/{change-name}/review-security` is generated from that JSON and kept for human/downstream compatibility.

`sourceRowValidation.rows` in canonical JSON is the only active security matrix. Markdown MUST NOT be hand-authored as the source of truth. If JSON and Markdown disagree, JSON wins and Markdown generation/parity must be repaired. Stale or parity-failed Markdown MUST route to `resolve-blockers` and MUST NOT be consumed as current evidence by `sdd-verify` or `sdd-archive`.

## JSON Field Mapping

| Markdown Section | Source JSON Field(s) |
| --- | --- |
| `# Review Security Report: {changeName}` | `changeName` |
| `## Verdict` | `changeName`, `status`, `verdict`, `nextRecommended`, `artifactMetadata.canonicalJsonRef` |
| `## Totals` | `totals` |
| `## Unavailable Tooling` | `unavailableTooling[]` |
| `## Exceptions` | `exceptions[]` |
| `## Artifact Metadata` | `artifactMetadata` |
| `## Summary` | `controls[]` grouped by catalog `category` (join by `id`); `totals` |
| `## Recommendation` | `verdict`, `nextRecommended`, `totals.blockers` |
| `## Matrix` | `controls[].id`, `controls[].complies`, `controls[].finding`, `controls[].evidenceLocation`, `controls[].justification` joined with catalog `controls[]` by `id` to get `guideline` and `category` |

## Required Structure

````markdown
# Review Security Report: {changeName}

## Verdict

| Field | Value |
| --- | --- |
| Change | `{changeName}` |
| Status | success \| blocked \| partial |
| Verdict | PASS \| PASS WITH WARNINGS \| FAIL |
| Next recommendation | verify \| apply \| resolve-blockers |
| JSON authority | `{artifactMetadata.canonicalJsonRef}` |
| Markdown authority | derived compatibility view |

## Totals

| Metric | Value |
| --- | --- |
| Total controls (155) | `{totals.controlCount}` |
| Validated | `{totals.validated}` |
| Passing | `{totals.passing}` |
| Failing | `{totals.failing}` |
| N/A | `{totals.notApplicable}` |
| Blockers | `{totals.blockers}` |
| Warnings | `{totals.warnings}` |
| Exceptions | `{totals.exceptions}` |

## Unavailable Tooling

{List from `unavailableTooling[]` or "None"}

## Exceptions

| ID | Approver | Approved At | Accepted Risk | Mitigation | Evidence Gap |
| --- | --- | --- | --- | --- | --- |
| `{id}` | `{approver}` | `{approvedAt}` | `{acceptedRiskRationale}` | `{mitigationOrFollowUp}` | `{evidenceGap}` |

(or "None")

## Artifact Metadata

| Check | Result |
| --- | --- |
| Canonical JSON ref | `{artifactMetadata.canonicalJsonRef}` |
| Derived Markdown ref | `{artifactMetadata.derivedMarkdownRef}` |
| JSON persisted / read back | `{artifactMetadata.jsonPersisted}` / `{artifactMetadata.jsonReadBack}` |
| Markdown generated / persisted / read back | `{artifactMetadata.markdownGenerated}` / `{artifactMetadata.markdownPersisted}` / `{artifactMetadata.markdownReadBack}` |
| JSON/Markdown parity | `{artifactMetadata.parityStatus}` |
| JSON authority | canonical |
| Markdown authority | derived |

## Summary

Summary by category. Derived from `controls[]` joined with catalog `category` by `id`. Passing = rows where `complies: Yes` or `complies: N/A` with justification. Blockers = rows where `finding: blocker`.

| Category | Passing/Total | Blockers |
| --- | --- | --- |
| authorization-access-control | `{N}/{total}` | `{count or 0}` |
| credential-secrets | `{N}/{total}` | `{count or 0}` |
| cryptography-data-protection | `{N}/{total}` | `{count or 0}` |
| database-access | `{N}/{total}` | `{count or 0}` |
| file-handling | `{N}/{total}` | `{count or 0}` |
| identity-authentication | `{N}/{total}` | `{count or 0}` |
| input-validation | `{N}/{total}` | `{count or 0}` |
| memory-safety | `{N}/{total}` | `{count or 0}` |
| observability-logging | `{N}/{total}` | `{count or 0}` |
| output-encoding | `{N}/{total}` | `{count or 0}` |
| pan-test-data | `{N}/{total}` | `{count or 0}` |
| safe-error-handling | `{N}/{total}` | `{count or 0}` |
| secure-coding | `{N}/{total}` | `{count or 0}` |
| sensitive-data-protection | `{N}/{total}` | `{count or 0}` |
| session-management | `{N}/{total}` | `{count or 0}` |

**Overall**: `{totals.passing}/{totals.controlCount}` passing · `{totals.blockers}` blockers · `{totals.warnings}` warnings · `{totals.exceptions}` exceptions

## Recommendation

- Next: `{nextRecommended}`
- {route-specific follow-up text}

## Matrix

Full 155-row table rendered last. For each row in `controls[]`, join `id` with catalog `controls[].id` to get `guideline` and `category`. All other columns come directly from the JSON row.

| ID | Guideline | Category | Complies | Finding | Evidence Location | Justification |
| --- | --- | --- | --- | --- | --- | --- |
| `{controls[].id}` | `{catalog join: guideline}` | `{catalog join: category}` | `{controls[].complies}` | `{controls[].finding}` | `{controls[].evidenceLocation}` | `{controls[].justification}` |
````

## Matrix Rules

- Header must match exactly: ID | Guideline | Category | Complies | Finding | Evidence Location | Justification
- 155 rows required, all control IDs (`REV-SEC-001` to `REV-SEC-155`) exactly once
- `Complies` must be `Yes`, `No`, or `N/A`
- `N/A` rows require non-empty `justification` and `evidenceLocation`
- Rows with `finding: blocker` make verdict FAIL unless an approved exception exists
- `justification` MUST NOT be truncated — render the full text from the JSON row
- `guideline` is joined from the catalog by `id`; it MUST NOT be truncated
- Rendered Markdown must be read back and compared to JSON for verdict/routing/counts and control coverage

## Safe Evidence Rules

- Evidence may cite paths, section anchors, sanitized summaries, command outcomes, and redacted placeholders only
- Must not include secrets, credentials, tokens, connection strings, PAN, PII, raw logs, sensitive payloads, production identifiers, generated bytes, or final-document-only values
- Missing runtime tooling is recorded in `unavailableTooling[]`, never as passing evidence
