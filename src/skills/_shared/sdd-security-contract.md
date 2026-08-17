# SDD Security Contract

Shared cross-phase contract for narrative secure design, downstream security evidence, routing boundaries, approved exceptions, and review-safe evidence. This file defines what every SDD phase may rely on. Machine-readable operational catalog inventory, Source ID expansion, PCI mappings, source-row validation, and report matrix schemas belong to `skills/sdd-review-security/references/`; the Markdown catalog remains the human/audit snapshot.

## Phase Ownership Boundary

| Artifact / phase | Owns | Must not own |
| --- | --- | --- |
| `design.md#secure-development-design` | Human-readable changed-surface classification, applicable security category rules, evidence owners, residual risks, exceptions, and safe-evidence policy. | YAML/JSON schemas, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, exhaustive `N/A` bookkeeping, or the 96-control general review matrix. |
| `test-design.md` | Evidence plan for applicable narrative category rules and changed-surface risks, including static/manual/automated coverage expectations and unavailable-tooling substitutions. | Security catalog expansion, design schema parsing, compact/source matrices, or exhaustive `N/A` planning. |
| `apply` evidence | Changed-file references and implementation/static/manual proof for the narrative rules and planned evidence. | Security-review verdicts, Source ID validation results, or exception approval. |
| `review-report.json` + derived `review-report.md` | General applied-change review and security handoff context. Canonical JSON is authoritative; Markdown is compatibility. | Security row verdicts, source-row validation results, full source matrices, or final verification. |
| `review-security-report.json` + derived `review-security-report.md` | Security review verdicts, `sourceRowValidation.rows` exact-once coverage, source-row counts, row-level evidence, `N/A` decisions, missed-applicable blockers, unsafe-evidence blockers, warning carry-forward, exceptions, evidence refs, and artifact parity/read-back metadata. Canonical JSON is authoritative; Markdown is compatibility. | Redefining shared phase boundaries, creating a second matrix authority, or duplicating/re-scoring the general 96-control review matrix. |
| `verify` / `archive` | Consume and preserve non-blocking review/security-review verdicts, warnings, exceptions, evidence refs, and report links. | Re-scoring review matrices, copying full Source ID matrices, or fixing implementation. |

Boundary rules:

- Design is narrative: it plans applicable category rules; it does not prove all omitted categories.
- Omitted categories are reviewable omissions. Canonical `review-security-report.json` validates omissions and blocks missed applicable controls; derived Markdown presents the JSON facts for compatibility.
- Exhaustive Source ID validation belongs to canonical `review-security-report.json`; downstream phases consume summaries, counts, warnings, exceptions, and evidence refs without copying or re-scoring the source-row matrix.
- Warning-only evidence may route forward only when mandatory evidence is complete and warnings stay visible to verify/archive.
- Missing, malformed, unsafe, unsupported, or conflicting security evidence blocks downstream phases until routed to `apply` or `resolve-blockers`.

## Shared Vocabulary

Use this vocabulary across design, test-design, apply evidence, review-security, verify, and archive.

| Field | Allowed values |
| --- | --- |
| `classification` | `security-impacting`, `no-impact` |
| `securityImpact` | `true`, `false` |
| `securityCategory` / `taxonomyCategory` | `authentication`, `sessions`, `sensitive-data-pan`, `secrets`, `permissions-access-control`, `files`, `database-access`, `sensitive-logging` |
| `decision` | `applicable`, `not-applicable`, `unknown` |
| `matrixAnswer` / `applies` / `complies` | `Yes`, `No`, `N/A` |
| `evidenceStatus` | `not-started`, `planned`, `implemented`, `verified`, `not-applicable`, `exception-approved`, `blocked` |
| `evidenceType` | `implementation-reference`, `static-inspection`, `test-evidence`, `approved-exception`, `n/a-evidence` |
| `finding` | `none`, `blocker`, `warning` |
| `ownerPhase` | `design`, `test-design`, `tasks`, `apply`, `review`, `review-security`, `verify`, `archive` |
| `route` | `apply`, `resolve-blockers`, `verify`, `archive` |
| `operationalSeverity` | `blocking`, `conditional`, `advisory` |

Operational severity is not general review severity. Labels such as `Menor`, `Media`, or `Mayor` may appear in general review findings, but they MUST NOT control security routing.

## Operational Evidence Security Boundary

Operational evidence is design-driven. When `design.md`, `test-design.md`, tasks, apply evidence, reviews, verify reports, archive reports, or archived artifacts include operational considerations, ordinary SDD evidence MUST preserve the restricted-data boundaries below without requiring absent operational categories.

| Category | Restricted in ordinary SDD evidence | Safe handling |
| --- | --- | --- |
| Operational identifiers | Production hostnames, IP addresses, ports, SID/service names, environment-specific service identifiers, and equivalent infrastructure values. | Cite paths, section anchors, sanitized summaries, redacted placeholders, or exact `Pendiente de confirmar:`. |
| Secrets and secret-like values | Credentials, tokens, private keys, connection strings, secret values, API keys, and equivalent authentication material. | Cite secret/config names, storage mechanisms, ownership, or redacted placeholders only; never values. |
| Logs, traces, errors, and payloads | Raw logs, raw traces, stack traces with sensitive payloads, user identifiers, confidential operational context, and sensitive payloads. | Cite log categories, redaction expectations, sanitized summaries, or safe evidence locations. |
| Generated bytes and exported contents | Generated file bytes, full exported files, environment-specific payloads, binary/document payload copies, and full ID lists. | Cite file type summaries, safe paths, checksums only when non-sensitive, or unresolved markers. |
| Final-document-only values | Production operational values explicitly provided by the user for the final manual operational document. | Use only in the final manual document; never backfill into design, tasks, apply, review, verify, archive, examples, tests, or fixtures. |
| Future exception evidence | Any planned exception to mandatory security evidence or applicable operational evidence. | Require approver, approval date, accepted-risk rationale, mitigation/follow-up, and exact evidence gap before archive readiness. |

Operational placeholders are safer than invented data. Exact `Pendiente de confirmar:` and exact `No aplica.` MUST NOT be treated as leakage by themselves, but placeholder-only evidence cannot hide a missing non-leakage check when operational evidence exists. `sdd-review-security` owns the final leakage verdict, source-row expansion, omitted-row validation, and `N/A` decisions.

## Narrative Secure Design Contract

Every new active `design.md` MUST include `## Secure Development Design`.

Recommended shape:

- `### Classification and Changed Surface` — security-impacting or no-impact classification; touched artifacts/behaviors; untouched runtime surfaces; evidence supporting the classification; omitted categories left for review-security validation.
- `### <Applicable Security Category> Rules` — one subsection per applicable category. Each subsection states development rules, prohibited unsafe patterns, evidence owners, expected review-safe evidence, residual risks, exception handling, and safe-evidence policy.
- `### Exception and Evidence Policy` — complete exceptions when needed, otherwise state that no exceptions are planned.

Rules:

- New active designs MUST be human-readable Markdown prose only.
- New active designs MUST NOT contain security YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, exhaustive Source ID inventories, or all-row `N/A` bookkeeping.
- No-impact designs record changed-surface rationale in prose; they do not prove no-impact with exhaustive `N/A` rows.
- Unknown data sensitivity, auth/session/access boundaries, secret/config handling, file/database behavior, or sensitive logging implications block design success until clarified.
- Applicable category rules MUST identify controls, evidence owners, expected evidence, residual risks, and archive expectations.
- Incomplete exceptions do not satisfy design, verification, or archive readiness.

## Review-Security Contract

`sdd-review-security` runs after non-blocking general review and before verify. It owns the operational security catalog and report schema through its local `references/` files.

Rules:

- It validates narrative design rules, test-design coverage, apply evidence, changed files, and review-report handoff evidence.
- It validates canonical `sourceRowValidation.rows` in `review-security-report.json` with exact-once Source ID coverage. Derived Markdown is generated from JSON and must not become authority.
- It consumes canonical `review-report.json` as general-review handoff evidence and MUST NOT duplicate, recreate, or re-score the general 96-control review matrix.
- If canonical security-review JSON and derived Markdown disagree, JSON wins and stale/parity-failed Markdown routes to `resolve-blockers`.
- Row presence alone is not compliance evidence; rows require review-safe corroborating evidence or complete `N/A` justification.
- Missed applicable controls, unsafe evidence, unsupported `N/A`, malformed rows, or incomplete mandatory evidence block downstream phases.
- Non-blocking security review verdicts route to `verify`; implementation/security evidence gaps route to `apply`; artifact/schema/context repair routes to `resolve-blockers`.

## Approved Exception Contract

Archive MAY proceed with missing mandatory evidence only when the exception is complete and remains visible through verify/archive.

Required fields:

```yaml
exception:
  status: exception-approved
  sourceId: <source-row-id-or-n/a>
  guidelineRef: <catalog-ref-or-category>
  approver: <human-or-authoritative-role>
  approvedAt: <iso-8601-or-date>
  acceptedRiskRationale: <why risk is accepted>
  mitigationOrFollowUp: <mitigation, issue, expiry, or follow-up plan>
  evidenceGap: <missing mandatory evidence>
```

Exception rules:

- Exceptions apply only after authoritative approval is recorded.
- Exceptions MUST NOT remove category rows, source refs, validation metadata, or no-impact proof requirements.
- Incomplete exceptions are blockers.

## Safe-Evidence Rules

Evidence in SDD artifacts, review reports, verify reports, and archive reports MUST be review-safe.

Allowed evidence:

- artifact paths and section anchors
- changed-file references
- command summaries or sanitized outputs
- data-flow summaries
- redacted placeholders
- issue/follow-up references
- exact operational evidence markers: `Pendiente de confirmar:` and `No aplica.`
- mechanism-oriented monitoring summaries such as dashboards, alerts, jobs, traces, scripts, documented manual checks, or SQL where appropriate

Forbidden evidence:

- raw secrets, credentials, tokens, private keys, connection strings
- PAN, PII, confidential values, or unnecessary sensitive operational context
- raw sensitive payloads in logs/errors/traces
- production hostnames, IP addresses, ports, SID/service names, or equivalent restricted operational identifiers
- full ID lists, generated file bytes, full exported file contents, or environment-specific payload copies
- final operational document values unless they appear only in the explicitly user-provided final manual document

Mandatory category reminders:

- `sensitive-data-pan`: cite data-flow, masking/encryption, retention, or redacted examples only.
- `secrets`: cite secret/config names, storage mechanisms, or ownership notes only; never values.
- `permissions-access-control`: cite workflow gates and denial-by-default evidence without exposing hidden policy data.
- `sensitive-logging`: cite redaction rules, log locations, or sanitized summaries only.
- `files`: cite artifact paths, file type summaries, or safe checksums only when non-sensitive; never embed generated bytes or full exported contents.
- `operational evidence`: cite safe summaries, exact markers, and evidence refs when operational considerations exist; never require restricted operational values to pass ordinary SDD evidence checks.

## Routing Semantics

| Condition | Route |
| --- | --- |
| Remediation requires code, prompt, task, contract, or apply-evidence work | `apply` |
| Remediation requires missing context, artifact repair, schema/catalog repair, backend reconciliation, configuration repair, or unsafe-evidence cleanup | `resolve-blockers` |
| Mandatory evidence is complete and only warning rows remain | `verify` |
| Verify confirms non-blocking security evidence | `archive` |

Downstream phases MUST NOT advance past unresolved blockers or CRITICAL security evidence gaps.
