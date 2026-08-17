# Exploration: corporate-source-row-security-validation

## Current State

The repository is a Markdown instruction-contract workspace for SDD agents and skills. OpenSpec config reports no runtime test/build/coverage runner and `strict_tdd: false`, so this change will be validated by static/manual evidence unless future tooling is added.

The active security flow is already compact and active-only:

`design.md#secure-development-design` -> `test-design.md` -> `review-security-report.md` -> `verify-report.md` -> `archive-report.md`

Current security authority validates the eight compact catalog rows only: `SEC-AUTH-001`, `SEC-SESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-FILE-001`, `SEC-DB-001`, and `SEC-LOG-001`. `skills/_shared/security-guideline-catalog.md` preserves the full corporate snapshot with stable dotted Source IDs and range notation such as `1.1-1.10`, but those source rows are currently used as traceability (`Source IDs` / `sourceRefs`) rather than as one-to-one operational validation rows.

The current catalog has an important gap: compact `SEC-*` source coverage does not cover every source ID in the `Full Corporate Guideline Snapshot`. Examples currently unmapped by compact source refs include `4.3-4.7`, `6.5-6.11`, `6.13-6.14`, `10.1-10.6`, `12.1`, and `12.3-12.5`. This makes the requested stricter source-row validation a real contract expansion, not just wording cleanup.

Current `sdd-review-security` validates every compact guideline ID exactly once and explicitly avoids duplicating the 96-control general review matrix. It does not yet require a corporate source-row matrix with per-source fields, missing/duplicate/unknown source ID checks, complete mapping checks, or per-source N/A evidence.

## Findings

- The compact eight `SEC-*` controls should remain the architectural control layer; replacing them would break the active design/review/verify/archive contracts and existing source specs.
- The full corporate snapshot already provides stable Source IDs and preserved guideline text, but there is no normalized operational source-row schema with `sourceId`, corporate section, guideline text, PCI alignment, mapped compact guideline ID, applies/complies, lifecycle status, evidence, observations, and finding classification.
- Range notation is documented as allowed only when every ID in the range exists, but no current contract requires expansion and validation of ranges before evaluating coverage.
- Current `Source Snapshot Checklist Template` is optional and too small for the requested gate: it lacks corporate section, PCI alignment, mapped compact guideline ID, lifecycle status, applies/complies separation, blocker/warning classification, and explicit safe-evidence/N/A requirements.
- Current phase guidance validates compact rows in design, test-design, review-security, verify, and archive. It does not yet preserve explicit traceability from every corporate Source ID through compact `SEC-*` controls to downstream evidence.
- Agent prompt copies under `agents/sdd/` are thin wrappers around global skill files. They likely need only targeted reminder updates if phase skill contracts change materially; the skill files remain the primary source of truth.

## Current Gaps

- No mandatory source-row inventory covering every source in `Full Corporate Guideline Snapshot`.
- No schema field in `design.md#secure-development-design` for source-row validation rows or a normalized source mapping table.
- No explicit rule that missing, duplicate, unknown, or unmapped source IDs block.
- No explicit rule that applicable source rows with missing evidence block and route to `apply` when remediation is implementation work.
- No explicit rule that missing mappings/schema/artifacts/N/A evidence route to `resolve-blockers`.
- No phase-level row shape requiring `applies`, `complies`, lifecycle status, evidence location, observations, and `finding/blocker/warning` per source row.
- No explicit PCI alignment extraction/preservation per corporate section, even though headings contain values such as `aligned to PCI Req 6.5.8`.
- No verification/archive gate proving source-row coverage remains complete after review-security.
- Existing source specs only require compact guideline Source IDs, not exhaustive corporate source-row validation.

## Affected Areas

- `skills/_shared/security-guideline-catalog.md` — add a normalized corporate source-row contract, range-expansion rules, exhaustive coverage checks, safe-evidence rules, and likely update compact mappings so every snapshot Source ID maps to exactly one or more compact `SEC-*` controls without replacing the eight controls.
- `skills/_shared/sdd-security-contract.md` — extend embedded secure-design and review-security schemas to include corporate source-row coverage, mapping validation, N/A evidence/justification, safe-evidence restrictions, and routing semantics.
- `skills/sdd-design/SKILL.md` — require design to plan source-row coverage and preserve Source ID -> compact guideline -> evidence traceability while still presenting the compact eight-control summary.
- `skills/sdd-test-design/SKILL.md` — require test design to cover mandatory source-row evidence expectations with static/manual checks when no runner exists.
- `skills/sdd-review-security/SKILL.md` — make the security review report validate every corporate source row, block on coverage/schema/mapping/evidence failures, and avoid duplicating the general 96-control matrix.
- `skills/sdd-verify/SKILL.md` — consume review-security source-row evidence and fail/block when mandatory row coverage or evidence is missing.
- `skills/sdd-archive/SKILL.md` — require archived audit trail to preserve source-row coverage and block missing mandatory source-row evidence unless complete approved exceptions exist.
- `agents/sdd/sdd-design.md`, `agents/sdd/sdd-test-design.md`, `agents/sdd/sdd-review-security.md`, `agents/sdd/sdd-verify.md`, `agents/sdd/sdd-archive.md` — prompt copies may need concise alignment notes if the phase contracts are expanded.
- `openspec/specs/sdd-security-guideline-catalog/spec.md`, `openspec/specs/sdd-design-workflow/spec.md`, `openspec/specs/sdd-test-design-workflow/spec.md`, `openspec/specs/sdd-review-security-workflow/spec.md`, and `openspec/specs/sdd-execution-persistence-contracts/spec.md` — source specs should receive delta requirements during the spec phase so the implementation has durable acceptance criteria.

## Approaches

1. **Add source-row coverage as a second operational matrix under the compact controls**
   - Keep the eight `SEC-*` rows as architectural controls.
   - Add `corporateSourceRows` / source-row matrix to design and review-security artifacts.
   - Validate range expansion, coverage completeness, mapping, evidence, lifecycle status, N/A evidence, and safe-evidence rules.
   - Pros: preserves current architecture, creates direct auditability, avoids duplicating the general review matrix.
   - Cons: larger artifacts; proposal/spec/design must define how much source-row detail belongs in design vs review-security.
   - Effort: Medium/High.

2. **Expand compact SEC controls into many SEC-like controls**
   - Convert every corporate source point into a first-class control.
   - Pros: simple one-row validation model.
   - Cons: violates the explicit restriction not to replace the compact `SEC-*` controls, increases prompt noise, and makes current contracts/changelogs harder to maintain.
   - Effort: High.

3. **Delegate source coverage to the general 96-control review matrix**
   - Reuse general review matrix mechanics for corporate source rows.
   - Pros: familiar report shape.
   - Cons: violates the restriction not to duplicate the general review matrix and blurs ownership between `sdd-review` and `sdd-review-security`.
   - Effort: Medium, but architecturally wrong.

## Recommendation

Use approach 1. Treat corporate source rows as an operational validation layer below the compact eight `SEC-*` architectural controls:

`Corporate Source ID -> mapped compact SEC-* -> design/test/apply/review-security/verify/archive evidence`

The proposal should define a source-row schema with at least:

| Field | Purpose |
| --- | --- |
| `sourceId` | Stable dotted ID from the full corporate snapshot after range expansion. |
| `corporateSection` | Snapshot section, including the title. |
| `guidelineText` | Preserved corporate guideline text or exact source reference. |
| `pciAlignment` | PCI requirement from section heading when present, otherwise `N/A`. |
| `mappedCompactGuidelineId` | One or more existing `SEC-*` IDs; missing/unknown mappings block. |
| `applies` | `Yes`, `No`, or `N/A`; `N/A` requires evidence and justification. |
| `complies` | `Yes`, `No`, or `N/A`; applicable missing evidence is `No`. |
| `lifecycleStatus` | Existing catalog lifecycle vocabulary. |
| `evidenceLocation` | Review-safe artifact/file/section/command/manual evidence location. |
| `observations` | Safe reviewer notes without secrets/PII/PAN/tokens/connection strings. |
| `finding` | `None`, `blocker`, or `warning`, with routing derived from the cause. |

Design should own source-row planning and mapping validation. Test design should plan source-row evidence checks. Review-security should be the main row-level implementation validation gate. Verify should consume the security review source-row verdict without reproducing the full matrix. Archive should preserve source-row coverage and block missing mandatory evidence.

## Routing Guidance

- Source-row implementation blockers -> `apply`.
- Missing mappings, malformed schema, missing artifacts, duplicate/unknown/missing Source IDs, or missing N/A evidence/justification -> `resolve-blockers`.
- Warnings only with complete required evidence -> `verify` after `sdd-review-security`.
- Successful exploration is ready for proposal -> `sdd-propose` / `propose`.

## Risks and Uncertainty

- Mapping every currently unmapped source ID to the compact eight controls requires careful design. Some rows (memory management, output encoding, generic coding controls, cryptography details) do not fit cleanly into today’s taxonomy and may require multi-mapping or broadened compact-control definitions without adding new compact controls.
- Full source-row matrices can become large. Use progressive disclosure: keep compact summaries visible, put exhaustive row detail in clearly bounded sections/tables, and avoid copying unnecessary sensitive values.
- Artifacts must remain review-safe. Evidence should cite locations and sanitized summaries, never raw secrets, PII, PAN, tokens, connection strings, private keys, or confidential values.
- With no runtime runner, verification evidence will be static/manual unless future config adds commands. Missing tooling must be reported, not treated as passing evidence.
- The change likely exceeds a small diff and should be planned in reviewable work units under the 400-line budget.

## Ready for Proposal

Yes. The next phase should create a proposal for an active security-contract expansion that preserves the eight compact `SEC-*` controls while adding exhaustive corporate Source ID row validation across design, test-design, review-security, verify, and archive.
