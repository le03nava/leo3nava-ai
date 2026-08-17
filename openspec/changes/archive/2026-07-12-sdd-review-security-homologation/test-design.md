# Test Design: SDD Review Security Homologation

## Overview

This change is a pure reference-artifact homologation: 4 static files under
`skills/sdd-review-security/references/` are updated from the v2 structural
pattern to the v3 flat-field pattern. No runtime services, test runners, build
tools, linters, type checkers, or coverage tools exist in this repository. All
47 requirements across 4 spec files are verified exclusively through
**static inspection** — `jq` assertions on JSON artifacts and `grep`/manual-read
assertions on Markdown artifacts. No automated test framework is available or
needed.

The `design.md` Secure Development Design section classifies this change as
**No security impact**: no trust boundaries, no external inputs, no secrets, no
PII, no authentication/authorization surfaces. No security category rules apply.
The safe-evidence and matrix-rules sections added to the template are
self-reinforcing documentation, not new risk surfaces.

---

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/sdd-review-security-homologation/proposal.md` | Scope, success criteria, rollback plan, and risk register |
| Spec — catalog-v3 | `openspec/changes/sdd-review-security-homologation/specs/catalog-v3/spec.md` | 12 requirements / scenarios for catalog JSON |
| Spec — schema-v3 | `openspec/changes/sdd-review-security-homologation/specs/schema-v3/spec.md` | 14 requirements / scenarios for schema JSON |
| Spec — template-v3 | `openspec/changes/sdd-review-security-homologation/specs/template-v3/spec.md` | 11 requirements / scenarios for report template |
| Spec — validation-rules-v3 | `openspec/changes/sdd-review-security-homologation/specs/validation-rules-v3/spec.md` | 10 requirements / scenarios for validation rules |
| Design | `openspec/changes/sdd-review-security-homologation/design.md` | Architecture decisions, file changes, data flow, testing strategy |
| Secure Development Design | `openspec/changes/sdd-review-security-homologation/design.md#secure-development-design` | No-security-impact classification; no applicable category rules; no residual risks |
| Testing Capabilities | `openspec/config.yaml` § `testing` | All layers unavailable (no runner, no linter, no coverage, no type-checker, no formatter) |

---

## Source ID Coverage Baseline

**Not applicable.** The `design.md#secure-development-design` section classifies
this change as having **no security impact**. No corporate source-row category
rules apply to the changed surface (static LLM reference artifacts with no
runtime, no trust boundary, no data class beyond "internal skill reference").
No static/manual security-category checks are required.

The safe-evidence rules added to the template (TMPL-010) and the safe-evidence
section preserved in validation rules are documentation-level guidance, not
controls requiring test-design coverage.

---

## Test Cases

All test cases are `static` inspection type. Method column shows the tool and
query. All cases are `mandatory` (every spec requirement must be met before
verification passes).

### Catalog — `security-guideline-catalog.operational.json`

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-C-001 | CATALOG-001 | `jq '.schemaVersion'` equals `3` (integer) | static | mandatory | Output: `3` | Run against the written file |
| TC-C-002 | CATALOG-001 (pre-v3 scenario) | `jq '.schemaVersion < 3'` on a stale copy returns `true` | static | mandatory | Confirms the condition a consumer would reject; documented as design-risk check from proposal risk table | Verify by inspection of schema version field type |
| TC-C-003 | CATALOG-002 | `jq 'has("generatedHumanViewRef")'` equals `false` | static | mandatory | Output: `false` | Top-level key check |
| TC-C-004 | CATALOG-003 | `jq 'has("reporting")'` equals `false` | static | mandatory | Output: `false` | Top-level key check |
| TC-C-005 | CATALOG-004 | `jq 'has("authority")'` equals `false` | static | mandatory | Output: `false` | Top-level key check |
| TC-C-006 | CATALOG-005 | `jq '[.sourceRows[] \| select(.ownerPhase == null or has("ownerPhase") == false)] \| length'` equals `0` | static | mandatory | Output: `0` | All 155 rows must have ownerPhase |
| TC-C-007 | CATALOG-005 / CATALOG-009 | `jq '[.sourceRows[] \| select(has("defaultOwnerPhase"))] \| length'` equals `0` | static | mandatory | Output: `0` | Legacy field must be gone on all rows |
| TC-C-008 | CATALOG-006 | `jq '[.sourceRows[] \| select(.route == null or has("route") == false)] \| length'` equals `0` | static | mandatory | Output: `0` | All 155 rows must have route |
| TC-C-009 | CATALOG-006 / CATALOG-010 | `jq '[.sourceRows[] \| select(has("defaultRoute"))] \| length'` equals `0` | static | mandatory | Output: `0` | Legacy field must be gone on all rows |
| TC-C-010 | CATALOG-007 | `jq '[.sourceRows[] \| select(has("guidelineRefs"))] \| length'` equals `0` | static | mandatory | Output: `0` | Removed field absent from all rows |
| TC-C-011 | CATALOG-008 | `jq '[.sourceRows[] \| select(has("evidenceExpectation"))] \| length'` equals `0` | static | mandatory | Output: `0` | Removed field absent from all rows |
| TC-C-012 | CATALOG-011 | `jq '.sourceRows \| length'` equals `155`; `jq '[.sourceRows[] \| select(.guidelineText == null or .guidelineText == "")] \| length'` equals `0` | static | mandatory | Row count: `155`; empty-guidelineText count: `0` | Content preservation check |

### Schema — `review-security-report.schema.json`

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-S-001 | SCHEMA-001 | `jq '.properties.schemaVersion.minimum'` equals `3` | static | mandatory | Output: `3` | Type constraint check |
| TC-S-002 | SCHEMA-002 | `jq '.required \| length'` equals `16`; `jq '.required \| sort'` matches exact 16-field list | static | mandatory | Length: `16`; sorted list matches spec | schemaName, schemaVersion, changeName, artifactKind, generatedAt, status, verdict, nextRecommended, totals, catalogRef, catalogSnapshotId, unavailableTooling, generalReviewRef, rows, exceptions, artifactMetadata |
| TC-S-003 | SCHEMA-003 | `jq '.properties \| keys \| map(select(. == "sourceRefs" or . == "catalogRefs" or . == "generalReviewHandoff" or . == "sourceRowValidation" or . == "blockers" or . == "warnings" or . == "unsafeEvidenceRejections" or . == "warningCarryForward")) \| length'` equals `0` | static | mandatory | Output: `0` | 8 removed keys must be absent |
| TC-S-004 | SCHEMA-004 | `jq '.properties.catalogRef.const'` equals `"skills/sdd-review-security/references/security-guideline-catalog.operational.json"` | static | mandatory | Output: exact const string | Const value must match exactly |
| TC-S-005 | SCHEMA-005 | `jq '.properties.generalReviewRef.type'` equals `"string"`; `jq '.properties.generalReviewRef.minLength'` equals `1` | static | mandatory | type: `"string"`, minLength: `1` | Non-empty string constraint |
| TC-S-006 | SCHEMA-006 | `jq '."$defs".totals.required \| length'` equals `8`; `jq '."$defs".totals.properties.sourceRowCount.const'` equals `155` | static | mandatory | required length: `8`; const: `155` | totals shape check |
| TC-S-007 | SCHEMA-007 | `jq '.properties.rows.minItems'` equals `155`; `jq '.properties.rows.maxItems'` equals `155` | static | mandatory | minItems: `155`, maxItems: `155` | rows[] at root with fixed bounds |
| TC-S-008 | SCHEMA-008 | `jq '."$defs".sourceRow.properties \| keys \| map(select(. == "corporateSection" or . == "pciAlignment" or . == "guidelineText" or . == "controlDomain" or . == "repoProfiles" or . == "runtimeSurface" or . == "dataSurface" or . == "appliesWhen")) \| length'` equals `0` | static | mandatory | Output: `0` | 8 catalog-only fields absent from sourceRow $def |
| TC-S-009 | SCHEMA-009 | `jq '."$defs".sourceRow.properties \| keys \| length'` equals `10`; key list matches spec | static | mandatory | Length: `10`; keys: sourceId, applies, complies, lifecycleStatus, evidenceType, evidenceLocation, justification, finding, ownerPhase, route | Exact property set check |
| TC-S-010 | SCHEMA-010 | `jq '."$defs".exception.properties \| has("status")'` equals `false` | static | mandatory | Output: `false` | status field must not exist in exception $def |
| TC-S-011 | SCHEMA-011 | `jq '."$defs".exception.required \| length'` equals `6`; list matches spec | static | mandatory | Length: `6`; fields: sourceId, approver, approvedAt, acceptedRiskRationale, mitigationOrFollowUp, evidenceGap | Exact required set check |
| TC-S-012 | SCHEMA-012 | `jq '."$defs" \| has("navigationGroup") or has("groupedNaSummary") or has("finding")'` equals `false` | static | mandatory | Output: `false` | 3 removed $defs must be absent |
| TC-S-013 | SCHEMA-013 | `jq '."$comment"'` contains `"rows[]"` and does not contain `"sourceRowValidation"` | static | mandatory | $comment includes `rows[]`; excludes `sourceRowValidation` | Path reference check |
| TC-S-014 | SCHEMA-014 | `jq '."$defs".artifactMetadata.required \| sort'` matches `review-report.schema.json` `."$defs".artifactMetadata.required \| sort` (9 fields) | static | mandatory | Both sorted arrays equal | Cross-file comparison; manual-read if jq not available |

### Template — `report-template.md`

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-T-001 | TMPL-001 | `grep` / manual-read: first 3 paragraphs before `## JSON Field Mapping` are present and contain JSON authority language | static | mandatory | 3 non-empty paragraphs present; content not truncated | Read file from offset 1; confirm preamble before first `##` heading |
| TC-T-002 | TMPL-002 | `grep -c '^\|'` on `## JSON Field Mapping` table body equals `9` data rows | static | mandatory | Row count: `9` | Count pipe-delimited body rows in that section |
| TC-T-003 | TMPL-003 | `grep` on `## Required Structure` section: exactly 8 subsection names match: Verdict, Totals, General Review Reference, Unavailable Tooling, Exceptions, Artifact Metadata, Recommendation, Matrix | static | mandatory | All 8 section names present; none others | Manual-read or grep per section name |
| TC-T-004 | TMPL-004 | `grep "Status"` and `grep -i "JSON authority"` both match inside `## Verdict` section | static | mandatory | Both strings present in Verdict section | Scoped grep within section boundaries |
| TC-T-005 | TMPL-005 | `grep -c '^\|'` on `## Totals` table body equals `8` data rows; `grep "Total source rows (155)"` matches | static | mandatory | Row count: `8`; pattern found | Count and string check |
| TC-T-006 | TMPL-006 | `grep "generalReviewRef"` matches in `## General Review Reference` section; `grep -i "handoff"` does not match in that section | static | mandatory | `generalReviewRef` found; `handoff` absent | Scoped to section |
| TC-T-007 | TMPL-007 | `grep "rows\[\]\.sourceId"` and `grep "sourceRows\[\]\.sourceId"` both match in `## Matrix` section | static | mandatory | Both patterns found | Join instruction validation |
| TC-T-008 | TMPL-008 | Manual-read of matrix table header row: exactly 15 `|`-delimited column names match spec list in order | static | mandatory | 15 columns: Source ID, Corporate Section, Control Domain, PCI Alignment, Guideline, Applies When, Applies, Complies, Lifecycle Status, Evidence Type, Evidence Location, Justification, Finding, Owner Phase, Route | Count pipe separators: 16 pipes = 15 columns |
| TC-T-009 | TMPL-009 | `grep "^## Matrix Rules"` matches | static | mandatory | Pattern found | Section heading presence |
| TC-T-010 | TMPL-010 | `grep "^## Safe Evidence Rules"` matches | static | mandatory | Pattern found | Section heading presence |
| TC-T-011 | TMPL-011 | `grep` for each of 6 obsolete headings returns no match: `## Source References`, `## General Review Handoff`, `## Source Row Navigation`, `## Source Row Summary`, `## Grouped Non-Applicability`, `## Blockers and Warnings` | static | mandatory | All 6 greps return no match | Absence assertion for each heading |

### Validation Rules — `validation-rules.md`

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TC-V-001 | VALRULES-001 | `grep "sourceRowValidation\.rows"` returns no match in file; `grep "rows\[\]"` matches in Active Flow Boundary section | static | mandatory | Forbidden pattern absent; new path present | Full-file and scoped grep |
| TC-V-002 | VALRULES-002 | `grep "sourceRowValidation\.rows"` in Source Row Coverage Validation table returns no match; `grep "rows\[\]"` in that table matches | static | mandatory | Forbidden pattern absent; new path present | Scoped to coverage table |
| TC-V-003 | VALRULES-003 | `grep "totals\.sourceRowCount"` and `grep "totals\.validated"` both match in count-consistency row; `grep -E "sourceRowValidation\.(expectedCount|validatedCount)|coverageStatus|exactOnce"` returns no match | static | mandatory | Both new fields present; 4 forbidden names absent | String presence + absence |
| TC-V-004 | VALRULES-004 | Manual-read of required row fields list: exactly 10 field names present matching spec; `grep -E "guidelineText|corporateSection"` in that list returns no match | static | mandatory | 10 fields listed; catalog-origin fields absent | Count and exclusion check |
| TC-V-005 | VALRULES-005 | `grep -i "source-row grouping"` in Source Row Coverage Validation table returns no match | static | mandatory | Pattern not found | Removed row label absence |
| TC-V-006 | VALRULES-006 | `grep -E "guidelineText.*guidelineRefs.*block\|guidelineRefs.*guidelineText.*block"` returns no match; `grep "guidelineRefs"` in blocking-rule context returns no match | static | mandatory | No blocking-rule mention of guidelineText/guidelineRefs | Full-file grep |
| TC-V-007 | VALRULES-007 | `grep "Row-Preserving Non-Applicability"` returns no match | static | mandatory | Pattern not found | Section heading absence |
| TC-V-008 | VALRULES-008 | Manual-read of Report Validation Checklist item for `review-security-report.json`: exactly 13 entries listed, matching spec field list | static | mandatory | 13 checklist entries: schemaName, changeName, status, verdict, nextRecommended, totals, catalogRef, catalogSnapshotId, unavailableTooling, generalReviewRef, rows[] (155), exceptions[], artifactMetadata | Count check |
| TC-V-009 | VALRULES-009 | `grep "sourceRowValidation\.rows"` in checklist returns no match; `grep "rows\[\] contains exactly 155 rows"` matches | static | mandatory | Old wording absent; new wording present | Wording replacement check |
| TC-V-010 | VALRULES-010 | `grep -E "groupedNaSummaries\|navigationSummary\|grouped N/A"` in checklist section returns no match | static | mandatory | All 3 patterns not found | Removed checklist item absence |

---

## Operational Considerations Checks

`design.md#Operational Considerations` states: **"No aplica."** This change
touches only static reference artifacts consumed by an LLM skill at design/review
time. No runtime services, logs, monitoring, administration, reprocessing, or
recovery operations are affected. No operational checks are required.

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Operational scope | Manual-read `design.md` "Operational Considerations" section confirms `No aplica.` marker is present | manual | Section reads "No aplica." | N/A — no tooling required for this check |

---

## Security Control Coverage

`design.md#secure-development-design` classifies this change as **No security
impact**. No trust boundaries are crossed, no external inputs are processed, no
data classes beyond "internal skill reference" are touched, and no existing
security controls are changed. No mandatory narrative category rules apply.

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| N/A | No security category rules applicable — classification: No security impact | No | Design classification in `design.md#secure-development-design` read and confirmed | not-applicable | None |

---

## Evidence Expectations

- All 47 test cases are **mandatory**; every case must be `applied` and
  `verified` before the verification gate passes.
- All test methods are **static inspection** (no runtime runner, no linter, no
  type-checker, no formatter, no coverage tool — all confirmed unavailable per
  `openspec/config.yaml`).
- `jq` is the primary tool for JSON assertions. `grep` and manual-read are the
  primary tools for Markdown assertions. Where `jq` is unavailable in the verify
  environment, manual-read of the file structure is acceptable substitute
  evidence provided the result is documented.
- Evidence for each case at verify time must record the actual command output or
  manual-read observation against the expected result.
- Absent runtime tooling is **not** passing evidence — it must be reported as
  unavailable if encountered.
- The `design.md#secure-development-design` no-security-impact classification is
  the justification for zero security-category checks. Any future change that
  modifies the safe-evidence or matrix-rules sections must re-evaluate this
  classification.
- Downstream phases (`sdd-tasks`, `sdd-apply`, `sdd-verify`) consume
  `test-cases.json` as the canonical case source. `test-design.md` is a
  human-readable projection only.

---

## Unavailable Tooling

| Tool | Status | Impact |
| --- | --- | --- |
| Test runner (any framework) | Unavailable | No automated test execution; all checks are static inspection |
| Unit / integration / E2E layers | Unavailable | Not applicable to this change |
| Coverage tool | Unavailable | Not applicable to this change |
| Linter | Unavailable | Not applicable to this change |
| Type checker | Unavailable | Not applicable to this change |
| Formatter | Unavailable | Not applicable to this change |
| `jq` | Available in most verify environments | Primary JSON assertion tool; manual-read fallback documented per case |

---

## Open Questions

None.
