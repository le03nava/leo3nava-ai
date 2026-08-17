# Test Design: Add SDD Review Phase

## Overview

This change adds `sdd-review` as a mandatory SDD phase between apply and verify, with durable `review-report.md` evidence, strict 96-control matrix rules, downstream verify/archive gates, and a security-catalog cross-reference boundary. The repository has no configured runtime test runner, linter, typechecker, formatter, coverage, unit, integration, or e2e tool, so planned evidence is limited to static and manual contract inspection.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/add-sdd-review-phase/proposal.md` | Intent, scope, success criteria, routing expectations, review artifact requirements, and non-goals. |
| Spec | `openspec/changes/add-sdd-review-phase/specs/sdd-review-workflow/spec.md` | Mandatory review gate, review-report persistence, exact matrix contract, routing severity semantics, and security boundary scenarios. |
| Spec | `openspec/changes/add-sdd-review-phase/specs/sdd-execution-persistence-contracts/spec.md` | Artifact resolver/state/status contract, `apply -> review -> verify -> archive` DAG, verify consumption, and archive readiness requirements. |
| Spec | `openspec/changes/add-sdd-review-phase/specs/sdd-security-guideline-catalog/spec.md` | Security guideline cross-reference behavior, review-safe evidence types, N/A evidence, and catalog authority preservation. |
| Security Applicability | `openspec/changes/add-sdd-review-phase/security-applicability.md` | Confirms `classification: no-impact`, `securityImpact: false`, and no required `security-design.md`. |
| Design | `openspec/changes/add-sdd-review-phase/design.md` | Architecture decisions, data flow, affected files, interface contracts, static/manual testing strategy, rollout, and sync compatibility expectations. |
| Security Design | Not required | Security applicability is no-impact; no mandatory security-design controls exist for this change. |
| Testing Capabilities | `openspec/config.yaml` | Confirms no automated runtime, lint, typecheck, formatting, coverage, unit, integration, or e2e commands are configured. |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: `sdd-review-workflow` Mandatory Review Gate; Design data flow | Inspect orchestrator/shared routing contracts to confirm completed apply routes to `review`, not directly to `verify`. | static | mandatory | Markdown contract references showing `apply -> review` as the apply success successor. | Verification MUST treat direct `apply -> verify` routing as a blocker. |
| TD-002 | Spec: `sdd-execution-persistence-contracts` Apply Review Verify Routing; Design data flow | Inspect phase DAG/status normalization to confirm the required route is `apply -> review -> verify -> archive`. | static | mandatory | Routing table or state/status contract text containing all four phases in order. | Covers the mandatory route requested for apply, review, verify, and archive. |
| TD-003 | Spec: Review Report Artifact; Design artifact identity | Confirm `sdd-review` persists `review-report.md` at `openspec/changes/{change-name}/review-report.md` in OpenSpec mode. | static | mandatory | Persistence/OpenSpec contract text naming the exact file path. | Missing path support MUST route to `resolve-blockers`. |
| TD-004 | Spec: Review Phase Artifact Contract | Confirm Engram/hybrid compatibility uses the stable artifact key `sdd/{change-name}/review` while OpenSpec uses `review-report.md`. | static | mandatory | Persistence/status contract text showing both backend identities without ambiguous aliases. | Prevents downstream resolver drift. |
| TD-005 | Spec: Review Report Artifact; Design interfaces | Inspect the `sdd-review` skill/report template to confirm `review-report.md` includes verdict, blocking summary, evidence summary, next recommendation, and matrix. | static | mandatory | Report template or contract section containing each required top-level report component. | This checks artifact persistence plus resolution content, not execution output. |
| TD-006 | Spec: Code-Review Validation Matrix; Proposal scope | Confirm the review matrix has exactly these columns, in this order: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. | static | mandatory | Report template table header matching the exact required columns and no extra columns. | Category MUST NOT be added as a matrix column. |
| TD-007 | Spec: Code-Review Validation Matrix; Design control catalog shape | Confirm the 96 source checklist controls are represented exactly once with stable Item IDs. | manual | mandatory | Manual catalog inspection notes with a 96-row count and no duplicate Item IDs. | No runner exists; reviewers must count the catalog/report source manually. |
| TD-008 | Spec: Code-Review Validation Matrix; Design control catalog shape | Confirm every review control preserves category information through either the stable catalog or category-bearing Item ID prefix. | manual | mandatory | Manual inspection notes showing category mapping for all 96 controls without changing report columns. | Covers stable IDs/categories while preserving the exact matrix. |
| TD-009 | Spec: Code-Review Validation Matrix | Confirm `Complies` is constrained to `Yes`, `No`, or `N/A`. | static | mandatory | `sdd-review` contract text defining only the allowed complies values. | Any additional value weakens downstream routing. |
| TD-010 | Spec: Platform control is irrelevant; Security catalog review-safe evidence | Confirm every `N/A` row requires both Evidence Location and Observations/Comments proving irrelevance. | static | mandatory | Report contract text requiring N/A evidence and comments, especially for platform-specific controls. | Prevents false passes for irrelevant platform controls. |
| TD-011 | Spec: Severity and Routing Semantics | Confirm blocking, critical, or explicitly blocking failures route back to `sdd-apply` and identify failed controls plus affected requirements. | static | mandatory | Review skill/routing contract text showing `next_recommended: sdd-apply` for blocking findings and required failure details. | Blocking review failures MUST prevent verify. |
| TD-012 | Spec: Severity and Routing Semantics | Confirm non-blocking findings may proceed to `sdd-verify` as warnings and are not silently dropped. | static | mandatory | Review envelope/report contract text showing non-blocking route to verify with warning/evidence summary. | Verify still consumes the report as evidence. |
| TD-013 | Spec: Review cannot safely run; Design routing | Confirm missing required artifacts, unknown changed files, unsafe workspace context, or persistence failure route to `resolve-blockers`. | static | mandatory | Review readiness/routing contract text listing these blocker cases. | This is separate from code-review blocking findings. |
| TD-014 | Spec: Verify consumes review evidence; Design file changes | Inspect `sdd-verify` contract updates to confirm verify consumes/cites `review-report.md` without owning or duplicating the full 96-control matrix. | static | mandatory | Verify skill/contract text requiring non-blocking review evidence and prohibiting full matrix ownership. | Keeps review and verify responsibilities separate. |
| TD-015 | Spec: Archive checks review readiness; Proposal success criteria | Inspect archive readiness contract to confirm archive requires both a non-blocking `review-report.md` and passing `verify-report.md`. | static | mandatory | Archive skill/OpenSpec/status contract text naming both artifacts as required gates. | Blocking review findings MUST prevent archive even if verify passed. |
| TD-016 | Spec: Security Boundary; Security applicability no-impact | Confirm review may cite security guideline IDs in `Standard`, but security applicability/design remain authoritative for applicability, required controls, and exceptions. | static | mandatory | Security catalog/review contract text preserving authority boundaries and avoiding duplicated guideline text. | Security design is not required for this no-impact change. |
| TD-017 | Spec: Catalog Boundary Preservation | Confirm security catalog updates provide cross-reference/evidence guidance only and do not duplicate or redefine security guideline authority inside `sdd-review`. | static | mandatory | Security catalog delta and review contract inspection showing references by guideline ID rather than copied guideline bodies. | Conflicts MUST resolve in favor of security applicability/design outputs. |
| TD-018 | Design: Sync compatibility | Inspect sync script behavior expectations to confirm adding new agent/skill files requires no script logic change because existing recursive sync copies them automatically. | manual | non-mandatory | Manual inspection notes for `scripts/sdd_init_agents.ps1` and `scripts/sdd_init_skills.ps1` recursive copy behavior. | Advisory compatibility check; a failed script-logic expectation should become a task/design risk before apply. |
| TD-019 | OpenSpec config testing constraints | Confirm verification plans report unavailable runtime checks explicitly and do not invent test, build, lint, typecheck, format, coverage, unit, integration, or e2e commands. | static | mandatory | `openspec/config.yaml` testing/verify sections plus verify/report text stating unavailable runtime tools. | All planned checks in this document are static or manual. |

## Security Control Coverage

Security applicability classifies this change as no-impact, so no `security-design.md` controls are mandatory. Security-relevant coverage for this change is limited to preserving catalog authority and review-safe cross-reference behavior.

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| Not applicable | No security-design controls are required because `securityImpact: false`. | No | `security-applicability.md` classification evidence; TD-016 verifies authority boundaries. | not-applicable | None |
| Catalog boundary | `sdd-review` may cite existing security guideline IDs but MUST NOT duplicate or redefine security guideline text, applicability, required controls, or exceptions. | Yes | TD-016, TD-017 | covered | None |
| N/A evidence | Platform-specific security review controls marked `N/A` MUST include evidence proving irrelevance and a comment explaining scope. | Yes | TD-010 | covered | None |

## No-Impact Assessment

Not applicable for the overall change: the SDD workflow, routing, persistence, and archive behavior have behavior/testability impact and require planned checks. Security design is not required because `security-applicability.md` classifies the change as `no-impact` with `securityImpact: false`.

## Evidence Expectations

- Mandatory cases require implementation, static/manual evidence, or a justified skip before verification can pass.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Static evidence should cite the exact Markdown contract, OpenSpec artifact, or configuration file inspected.
- Manual evidence should record the inspected files, the observed count or contract result, and any mismatch.
- Runtime command evidence is not expected because this repository has no configured runner, linter, typechecker, formatter, coverage, unit, integration, or e2e tool.
- Verification must report unavailable runtime checks explicitly instead of treating missing tools as passing evidence.

## Open Questions

None.
