# Exploration: active-only-security-contract

## Current State

The active security flow is already implemented conceptually as:

`design.md#secure-development-design` -> `review-security-report.md` -> `verify` -> `archive`

OpenSpec config confirms this repository is a Markdown instruction-contract workspace with no configured test runner and `strict_tdd: false`. Active SDD contracts now consistently say new changes must not create or require standalone `security-applicability.md` or `security-design.md` artifacts. However, active surfaces still contain legacy prose, legacy resolver rows, and compatibility labels mixed into current-flow docs.

The highest-noise file is `skills/_shared/sdd-security-contract.md`: it still includes a full legacy `security-applicability.md` schema, legacy/read-only headings, and PR compatibility notes. That conflicts with the intent that this shared contract should describe only the active/current security flow.

## Affected Areas

- `skills/_shared/sdd-security-contract.md` — primary target. It should become active-only: shared vocabulary, embedded secure design schema, review-security contract, exception fields, and safe-evidence rules. Remove the legacy applicability schema and compatibility notes from this active contract.
- `skills/_shared/security-guideline-catalog.md` — likely needs wording cleanup where the catalog describes its active scope as including legacy standalone compatibility. Keep stable guideline IDs/source snapshot audit language; prefer current-flow wording.
- `skills/sdd-design/SKILL.md` — contains safe current rules but several lines are framed as legacy compatibility. Keep the behavior (do not require standalone artifacts) while rewriting as positive current-flow rules.
- `skills/sdd-review-security/SKILL.md` — contains active behavior plus legacy read-only notes. Replace legacy notes with positive requirements: require embedded secure design and ignore absent standalone artifacts for new changes.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` — active source spec includes legacy standalone compatibility in purpose/scope. It likely needs a delta to align catalog authority with current-flow evidence only.
- `openspec/specs/sdd-review-security-workflow/spec.md` — active source spec has a legacy standalone scenario. It can be removed or reframed as out-of-scope for active workflow.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` and `skills/_shared/persistence-contract.md` / `skills/_shared/sdd-status-contract.md` — higher risk. These contain legacy refs/tokens for status and persisted-state readability. Do not blindly delete fields/tokens unless the design proves no reader depends on them; preserve runtime read behavior without active legacy prose where possible.
- `skills/_shared/openspec-convention.md`, `README.md`, `agents/sdd/sdd-orchestrator.md`, and phase skills (`sdd-apply`, `sdd-review`, `sdd-test-design`, `sdd-tasks`, `sdd-verify`, `sdd-archive`) — likely secondary cleanup only if proposal scope allows; many references are explanatory rather than contract authority.

Archived artifacts under `openspec/changes/archive/**` are out of scope and should not be edited.

## Approaches

1. **Active-contract cleanup with compatibility isolated to status/persistence internals**
   - Remove legacy schemas and legacy notes from `sdd-security-contract.md` and current-flow docs.
   - Keep any minimum compatibility tokens/fields in status and persistence contracts if needed for reading old state, but describe them as non-active internal/read behavior rather than part of the security contract.
   - Pros: matches user intent, keeps archive/status recovery safe, smaller diff.
   - Cons: some legacy words may remain in low-level persistence/status internals if required for safe parsing.
   - Effort: Medium.

2. **Aggressive removal from every active surface**
   - Delete all visible legacy mentions from skills, specs, status, persistence, README, and orchestrator prompts.
   - Pros: cleanest reading surface.
   - Cons: high risk of breaking status parsing, continuation, or archive display for old state fields such as `securityDesign` and `securityApplicability`.
   - Effort: Medium/High.

## Recommendation

Use approach 1. Treat `skills/_shared/sdd-security-contract.md` as the active-only security contract and remove its legacy applicability schema entirely. Preserve archive history by not editing archived folders, and preserve runtime behavior by leaving only minimal compatibility handling in generic status/persistence/read contracts when those fields are needed to parse old state.

This should be a **medium SDD change**. It touches several Markdown contracts and specs, but the implementation is documentation/contract cleanup rather than runtime code. Keep work units under the 400-line review budget by slicing along contract authority:

1. **Core active security contract**: update `skills/_shared/sdd-security-contract.md` and aligned catalog wording.
2. **Active phase consumers**: update `sdd-design` and `sdd-review-security` wording, plus any direct active consumers needed for consistency.
3. **Source specs and low-level contracts**: update OpenSpec specs and status/persistence wording only where necessary, preserving compatibility parsing without active legacy prose.

## Risks

- Removing `securityApplicability` / `securityDesign` fields or resolver rows from status/persistence contracts could break reading old persisted state, archived status output, or recovery. Mitigation: do not remove parser fields unless design proves they are unused; mark them as internal historical-read fields if needed.
- Removing all legacy notes from active surfaces may obscure why archived artifacts still contain older filenames. Mitigation: keep archive history in archived artifacts and, if required, generic archive/status docs only—not in the active security contract.
- Specs may currently require legacy readability. The proposal/spec phase should explicitly distinguish “no separate legacy contract” from “old artifacts remain readable as historical files.”
- No automated test runner exists. Verification will be static/manual inspection plus targeted search checks.

## Ready for Proposal

Yes. Proposal should define an active-only security contract cleanup with non-goals: no separate legacy contract, no edits under `openspec/changes/archive/**`, and no removal of parser/read behavior required for old state unless safely replaced.
