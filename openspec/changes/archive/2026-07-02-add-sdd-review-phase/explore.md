# Exploration: add-sdd-review-phase

## Current State

The SDD workflow currently routes `explore? -> proposal -> spec -> security-applicability -> design -> security-design? -> test-design -> tasks -> apply -> verify -> archive`.

`sdd-apply` records implementation progress and task checkbox evidence, while `sdd-verify` proves completion against specs, design, security evidence, test-design cases, tasks, and runtime/static evidence. The current verification report already has compliance matrices for specs, security evidence, test-design coverage, correctness, and design coherence, but it does not model a dedicated post-apply code review artifact or the requested 96-control validation matrix.

## Affected Areas

- `agents/sdd/sdd-orchestrator.md` — owns the DAG, dependency graph, routing tokens, launch envelope, result contract, state persistence, and required context by phase.
- `agents/sdd/sdd-review.md` — likely new phase executor prompt, mirroring existing executor boundary prompts.
- `skills/sdd-review/SKILL.md` — likely new phase skill defining review inputs, review matrix contract, routing, persistence, and severity semantics.
- `skills/_shared/sdd-status-contract.md` — status schema, artifact refs/paths, dependencies, routing-token mapping, and archive readiness would need a `review` / `reviewReport` addition.
- `skills/_shared/persistence-contract.md` — artifact resolver, state schema, backend references, and sub-agent context table would need a review-report artifact.
- `skills/_shared/openspec-convention.md` — OpenSpec directory layout and artifact table would need `review-report.md` between tasks/apply evidence and verify report.
- `skills/sdd-apply/SKILL.md` — success routing would change from direct `verify` to `review` when all tasks are complete.
- `skills/sdd-verify/SKILL.md` and `skills/sdd-verify/references/report-format.md` — verify should consume a passing review report as static/manual quality evidence, but should not own the 96-control review process.
- `skills/sdd-archive/SKILL.md` — archive readiness should require a passing `review-report.md` in addition to passing verification, complete tasks, test design, and applicable security evidence.
- `README.md` — phase list and phase order need the new phase.
- `openspec/specs/sdd-execution-persistence-contracts/spec.md` — source spec should define the new artifact and DAG behavior.
- `openspec/specs/sdd-security-guideline-catalog/spec.md` and `skills/_shared/security-guideline-catalog.md` — existing security catalog overlaps part of the requested checklist; review should reference these controls where possible instead of duplicating security authority blindly.

## Approaches

1. **Separate mandatory `sdd-review` gate between apply and verify**
   - Pros: clean ownership; produces the requested validation matrix as a first-class artifact; keeps `sdd-verify` focused on proving implementation/spec/test evidence; gives archive a durable review gate.
   - Cons: requires broad DAG/status/persistence/orchestrator updates; adds one more phase to the workflow.
   - Effort: High.

2. **Fold the review matrix into `sdd-verify`**
   - Pros: fewer routing/state changes; verify already produces matrices and quality verdicts.
   - Cons: mixes code review with verification; makes verify too large; risks treating static checklist review as runtime/spec proof; harder to rerun review independently after apply fixes.
   - Effort: Medium.

3. **Make review an optional `judgment-day`-style workflow outside the DAG**
   - Pros: minimal SDD DAG disruption; can reuse adversarial review ideas.
   - Cons: not a reliable gate; archive/status cannot require the artifact; less auditable and not aligned with the user's request for a structured validation matrix after apply.
   - Effort: Low to Medium.

## Recommendation

Add `sdd-review` as a separate mandatory review gate between `apply` and `verify`:

```text
explore? -> proposal -> spec -> security-applicability -> design -> security-design? -> test-design -> tasks -> apply -> review -> verify -> archive
```

This is the cleanest fit because the requested output is a code-review deliverable, not a verification report. `sdd-review` should read proposal/specs/design/security/test-design/tasks/apply evidence plus changed files, produce `review-report.md`, and route:

- `PASS` or `PASS WITH WARNINGS` with no blocking severity -> `next_recommended: verify`
- any CRITICAL/blocking control failure -> `next_recommended: apply`
- missing artifacts, unknown changed files, unsafe workspace, or persistence failure -> `next_recommended: resolve-blockers`

`sdd-verify` should then consume `review-report.md` as one input among others and confirm that no blocking review failures remain before it can recommend archive. `sdd-archive` should require both a passing review report and a passing verify report.

The validation matrix should be the primary artifact section and use the requested columns exactly:

| Item | Artifact/Deliverable | Requirement | Reviewer | Standard | Severity | Complies | Affected Requirement | Evidence Location | Observations/Comments |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |

The 96 controls should be normalized into stable control IDs and categories so future reports can be compared across runs. Security controls should cross-reference the existing security guideline catalog where applicable, while non-security controls should live in a review-control catalog owned by the new review phase or a shared `sdd-review-control-catalog.md` reference.

## Uncertainty / Risk

- The exact 96-control list was summarized in the launch prompt but not available as a complete row-by-row source artifact during exploration; proposal/spec should capture the full checklist with stable IDs before implementation.
- Some requested controls are language/platform-specific (for example ClassLoader/modules, log4j version, dynamic SQL, XML/HTML generation). The review phase must support `N/A` with evidence and avoid failing irrelevant controls.
- Overlap with `sdd-security-applicability`, `sdd-security-design`, and `sdd-verify` needs careful boundaries: review may evaluate compliance evidence, but security design remains the owner of applicable security controls and verify remains the owner of runtime/spec proof.
- Repository has no runtime test runner; review implementation will primarily be Markdown prompt/contract changes validated through static inspection and OpenSpec artifacts.
- Adding a mandatory DAG phase is a broad contract change likely above the 400-line review budget, so later tasks should plan stacked work units.

## Ready for Proposal

Yes. The proposal should define `sdd-review` as a mandatory post-apply/pre-verify gate, add the `review-report` artifact to persistence/status/OpenSpec conventions, and require the review matrix with the exact requested columns plus stable control IDs and severity/routing rules.
