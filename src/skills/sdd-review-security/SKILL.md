---
name: sdd-review-security
description: "Validate embedded secure-design evidence and persist canonical review-security-report.json plus derived Markdown. Trigger: orchestrator launches review-security in parallel with sdd-review, after sdd-apply."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-review-security` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Activation Contract

Run in parallel with `sdd-review`, after `sdd-apply` completes implementation work, and before `sdd-verify`. Validate narrative `design.md#secure-development-design`, `test-design.md`, tasks/apply evidence, and changed-file context against the control catalog. Does not depend on or consume `sdd-review` output.

Persist two coordinated owned artifacts: canonical `review-security-report.json` plus derived Markdown compatibility view `review-security-report.md` / `sdd/{change-name}/review-security`.

This phase owns the security-review verdict, exact-once control ID validation, missed-omission checks, safe-evidence checks, row-level non-applicability decisions, blockers, warning carry-forward, and next routing. It does not replace general review, verification, archive readiness, or implementation fixes.

The orchestrator sends a structured `launch:` YAML envelope first, followed by a `## Phase Context` section. The envelope schema is defined in `skills/_shared/sdd-phase-common.md ## Launch Envelope Contract`.

Required fields to extract:
- `launch.changeName` — the change name
- `launch.artifact_store.mode` — backend to use for reads and persistence
- `launch.status.dependencies` — `apply: completed` must be present
- `launch.artifacts.paths` or `launch.artifacts.refs` — refs for design (with `## Secure Development Design`), test-design, apply-progress, and changed-file context; do NOT expect content inline
- `launch.actionContext.workspaceRoot` — absolute workspace path to inspect changed files
- `launch.skill_paths` — supplemental skills to load before work

If `launch.artifact_store.mode` is absent or `launch.status.dependencies` does not show `apply: completed`, return `blocked` with `next_recommended: resolve-blockers`.

## Control Authority Contract

- Canonical JSON is the only active security-review authority.
- `controls` in the report is the only active security matrix.
- Every active report MUST contain exactly 155 rows under `controls`.
- Every row MUST have a unique `id` matching a catalog control; the set MUST equal all 155 catalog `controls[].id` exactly once.
- Derived Markdown is presentation only. If JSON and Markdown disagree, JSON wins and the stale derived artifact routes to `resolve-blockers`.

## Phase Artifact Contract

Follow the shared contracts instead of duplicating their rules:

| Concern | Contract |
| --- | --- |
| Backend mechanics | `skills/_shared/sdd-phase-common.md` Sections B/C and `skills/_shared/persistence-contract.md`. |
| Post-apply gates | `skills/_shared/sdd-post-apply-gates.md`. |
| Security vocabulary/routing/safe evidence | `skills/_shared/sdd-security-contract.md`. |
| Catalog authority | `assets/review-security-control-catalog.json` is the canonical control catalog with 155 controls; each control has `id`, `category`, `guideline`, `severity`, `defaultComplies`, `finding`, and `evidenceHint`. `assets/review-security-control-catalog.md` is a derived human/audit view. |
| Markdown presentation | `assets/review-security-report.md` is the derived Markdown compatibility presentation generated from canonical JSON. |
| Required inputs | Structured status; narrative `design.md#secure-development-design`; tasks/apply evidence; changed-file context; `test-design.md`; security catalog; shared security contract. |
| Produced artifact | Canonical `sdd/{change-name}/review-security-report.json` plus derived `sdd/{change-name}/review-security` in Engram mode, or canonical `openspec/changes/{change-name}/review-security-report.json` plus derived `openspec/changes/{change-name}/review-security-report.md` in OpenSpec mode. |
| Mutates | None outside the produced security review report artifacts. |
| JSON authority | Build, validate, persist, and read back canonical JSON first. Render Markdown only from that JSON; Markdown never wins on conflict. |
| Matrix ownership | Only this phase owns security control validation. Do not copy or re-score the general 96-control matrix. |
| Success routing | Non-blocking verdict: `next_recommended: verify`. |
| Failure routing | Missing mandatory evidence, incomplete exceptions, blocked rows, or security blockers: `next_recommended: apply` or `resolve-blockers` according to remediation owner. |

## Decision Gates

| Situation | Action |
| --- | --- |
| Shared post-apply, safe-evidence, dependency, or persistence gate fails | Follow `skills/_shared/sdd-post-apply-gates.md`; keep verify/archive unavailable. |
| Embedded `design.md#secure-development-design` is missing | Return `blocked` with `next_recommended: resolve-blockers`. |
| Security-review JSON and Markdown disagree after rendering/read-back | JSON wins; mark Markdown stale/parity-failed, return `blocked` or `partial` with `next_recommended: resolve-blockers`, and do not let downstream phases consume stale Markdown. |
| New-change evidence requires YAML/schema/matrices in design or all-row design non-applicability bookkeeping | Treat as invalid for the active flow; validate from narrative design, catalog, and artifact evidence. |
| Control rows are missing, duplicated, unknown, or malformed | Persist no passing report; return `blocked` with `next_recommended: resolve-blockers` unless remediation clearly belongs to `apply`. |
| Changed files/apply/test-design evidence prove an omitted security concern applies | Persist a blocking report; route to `apply` for file/prompt/contract/task evidence remediation, otherwise `resolve-blockers`. |
| Applicable mandatory security evidence is absent and no complete approved exception exists | Mark row `No` / blocked and route to `apply` when implementation/evidence work is needed. |
| Operational evidence leaks restricted identifiers, secrets, raw logs/payloads, generated bytes, final-document-only values, or invented details | Persist a blocking report and route by remediation owner. |
| Exact `Pendiente de confirmar:` or `No aplica.` appears | Treat as safe placeholder text, but still require safe non-leakage proof when a security obligation applies. |
| Only non-blocking warnings remain with complete mandatory safe evidence | Persist `PASS WITH WARNINGS` and return `next_recommended: verify`. |

## Execution Steps

1. Load supplemental skills via `skills/_shared/skill-resolver.md` and `sdd-phase-common.md` Section A.
2. Apply common post-apply gates; resolve/read required inputs from the selected backend or explicit `contextFiles`.
3. Confirm embedded secure design exists and apply-progress is present from the selected backend.
4. Parse narrative secure-design classification, applicable category rules, evidence expectations, exceptions, residual risks, safe-evidence policy, and omitted categories. Do not require design YAML/schema/matrices.
5. Load the catalog from `assets/review-security-control-catalog.json`. Validate all 155 expected control IDs (`REV-SEC-001` to `REV-SEC-155`) exactly once — missing, duplicate, or unknown IDs block.
6. Validate each report row: required fields (`id`, `complies`, `evidenceLocation`, `justification`, `finding`); vocabulary (`complies`: Yes/No/N/A; `finding`: none/blocker/warning); non-applicability justification when `complies: N/A`; finding/verdict consistency; safe evidence; warning carry-forward; and exception completeness.
7. Validate operational evidence leakage boundaries when operational evidence exists; preserve exact safe placeholders while rejecting unsafe evidence.
8. Build canonical `review-security-report.json` with schema identity, `changeName`, verdict/status, `nextRecommended`, source refs, catalog refs, `controls` (exactly 155 rows, each with only `id`, `complies`, `finding`, `evidenceLocation`, `justification`), blockers, warnings, unsafe evidence rejections, warning carry-forward, exceptions, unavailable tooling, and artifact parity/read-back metadata.
9. Validate JSON before Markdown rendering: required top-level fields present, exactly 155 rows under `controls`, exact-once IDs matching catalog, each row has only the 5 required fields, safe-evidence rules checked, N/A justification completeness, exception completeness, and counts match verdict.
10. Render derived Markdown from canonical JSON using `assets/review-security-report.md`. Follow the Required Structure exactly: Verdict → Totals → Unavailable Tooling → Exceptions → Artifact Metadata → Summary (grouped by `category`) → Recommendation → Matrix. The `## Matrix` MUST render all 155 rows as a table with exactly these 7 columns in order: `ID | Guideline | Category | Complies | Finding | Evidence Location | Justification`. Join each row's `id` with catalog `controls[].id` to populate `guideline` and `category`. Do NOT truncate `justification` or `guideline` — render full text.
11. Persist and read back canonical JSON first, then derived Markdown, compare parity metadata, and register both artifacts in state/status handoff with JSON first using `authority: canonical` and Markdown second using `authority: derived`.

## Row Validation Rules

### Control Coverage

| Rule | Requirement |
| --- | --- |
| Exact control universe | Load catalog and validate all 155 control IDs (`REV-SEC-001` to `REV-SEC-155`) exactly once. Missing, duplicate, or unknown IDs block. |
| Exact report table | The report table path is `controls`; no parallel security matrix is permitted. |
| Count consistency | `totals.controlCount` (const 155) and `totals.validated` MUST both be `155` for non-blocking reports. |
| Required row fields | Every row MUST include `id`, `complies`, `evidenceLocation`, `justification`, and `finding`. Catalog fields (`guideline`, `category`) are joined from the catalog at render time. |
| Row values | `complies` MUST be `Yes`, `No`, or `N/A`; `finding` MUST be `none`, `blocker`, or `warning`. |
| Blocker consistency | Rows with `finding: blocker` make the verdict FAIL. A non-blocking report cannot contain unresolved blockers. |

### Missing, Duplicate, and Unknown Control IDs

- Missing catalog control IDs are blockers. The safe error MAY name missing IDs or counts, but MUST NOT echo row evidence payloads.
- Duplicate control IDs are blockers. The safe error MAY name the repeated `id` and report exact-once failure.
- Unknown control IDs are blockers. Unknown rows never compensate for missing catalog rows.

### Safe Evidence and Leakage Rules

- Evidence locations, observations, and summaries cite paths, section refs, changed-file refs, command summaries, sanitized summaries, exact safe placeholders, or redacted placeholders only.
- Reject raw secrets, credentials, tokens, private keys, connection strings, PAN, PII, raw logs/payloads, restricted production identifiers, generated bytes, final-document-only values, and invented operational details.
- Unsafe row evidence MUST be represented through `unsafeEvidenceRejections` plus a safe row finding; do not copy the unsafe value into blockers, warnings, Markdown, or command output.
- Exact `Pendiente de confirmar:` and exact `No aplica.` are safe marker states but do not replace required safe non-leakage proof when an applicable security obligation exists.
- Missing tools or unavailable automation are not passing evidence. Report them as unavailable tooling.

### Report Validation Checklist

Before returning success, verify:

- Phase envelope uses snake_case `next_recommended`; camelCase `nextRecommended` only inside artifact/state metadata.
- Canonical `review-security-report.json` includes `schemaName`, `changeName`, `status`, `verdict`, `nextRecommended`, `totals`, `catalogRef`, `catalogSnapshotId`, `unavailableTooling`, `controls[]` (exactly 155), `exceptions[]`, and `artifactMetadata`.
- `controls` contains exactly 155 rows with exact-once IDs matching catalog `controls[].id`; each row has only `id`, `complies`, `finding`, `evidenceLocation`, and `justification`.
- Derived Markdown `## Matrix` rendered from `controls[]` joined with catalog by `id` to get `guideline` and `category`; no catalog metadata stored in the report rows.
- Every row contains exactly the required fields: `id`, `complies`, `finding`, `evidenceLocation`, `justification`.
- Every `N/A` row has row-level justification and safe evidence.
- Every exception has complete approved-exception fields from the shared security contract.
- Safe-evidence rules were enforced; no unsafe values in blockers, warnings, Markdown, or output.
- Derived Markdown rendered from canonical JSON, read back, parity-checked, full control matrix rendered last.
- Routing consistent with blockers, warnings, persistence result, and remediation owner.

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put `## Security Review Report Summary` in `detailed_report` with change, inputs inspected, verdict, canonical JSON ref, derived Markdown ref, control counts, blockers, non-blockers, operational evidence leakage verdict when evidence exists, placeholder safety notes, control ID ownership validation, safe-evidence notes, unavailable tooling, artifact parity/read-back metadata, security review summary by category (passing/total + blockers per category, format: `category: N/total — X blockers`), and next route. The `artifacts` array MUST list canonical JSON before derived Markdown; use notes such as `authority: canonical` and `authority: derived` until the shared envelope supports a dedicated authority field.

## Routing Contract

- `PASS` or `PASS WITH WARNINGS` with no blockers -> `status: success`, `next_recommended: verify`.
- Implementation/security evidence gaps -> persist `FAIL`, route `next_recommended: apply` when remediation is file, prompt, contract, task, or apply-evidence work.
- Missing/malformed/unsafe/conflicting context, catalog/schema repair, backend repair, unsupported non-applicability, or unsafe evidence cleanup -> `status: blocked`, `next_recommended: resolve-blockers` unless remediation clearly belongs to `apply`.
- Persistence failure after a useful report -> `status: partial`, `next_recommended: resolve-blockers`, include safe recovery details.
- Do not return camelCase `nextRecommended` from the phase envelope. CamelCase is for artifact/status/state metadata only.

## Rules

- ALWAYS run in parallel with `sdd-review`, after `sdd-apply`, and before `sdd-verify` for new changes.
- ALWAYS require narrative `design.md#secure-development-design` for active new changes.
- NEVER require design/test-design to materialize control IDs, schema fields, YAML/JSON, or exhaustive non-applicability rows.
- NEVER consume, duplicate, recreate, or re-score the 96-control general review matrix. This phase is independent of `sdd-review` output.
- ALWAYS treat canonical `review-security-report.json` as the security-review source of truth. Derived Markdown exists for human/downstream compatibility only.
- Control presence alone is not compliance evidence; each row needs safe corroborating evidence, justified non-applicability, or complete approved exception.
- Return `next_recommended: verify` only for non-blocking security-review verdicts.

## References

- `../_shared/skill-resolver.md` — supplemental skill loading and `skill_resolution` protocol.
- `../_shared/sdd-phase-common.md` — phase retrieval, persistence, and return envelope.
- `../_shared/persistence-contract.md` — artifact keys, backend behavior, and read-back verification.
- `assets/review-security-control-catalog.json` — canonical control catalog with 155 controls (`REV-SEC-001` to `REV-SEC-155`); each entry has `id`, `category`, `guideline`, `severity`, `defaultComplies`, `finding`, and `evidenceHint`.
- `assets/review-security-control-catalog.md` — derived human/audit view generated from the canonical JSON catalog.
- `assets/review-security-report.md` — derived `review-security-report.md` presentation contract generated from `review-security-report.json`.
- `../_shared/sdd-security-contract.md` — narrative secure-design and review-security report schema contracts.
- `../_shared/sdd-post-apply-gates.md` — common post-apply gates, review evidence consumption, routing defaults, safe evidence, and matrix ownership boundaries.
