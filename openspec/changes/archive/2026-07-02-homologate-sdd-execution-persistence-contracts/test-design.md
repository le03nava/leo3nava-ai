# Test Design: Homologate SDD Execution and Persistence Contracts

## Overview

This change is a Markdown-only SDD contract refactor with process/compatibility impact but no runtime test tooling. Planned evidence is static/manual review only: verify that shared persistence behavior has one authority, phase skills keep compact artifact contracts, duplicated contradictory persistence prose is removed, and existing artifact keys, OpenSpec paths, routing tokens, state fields, conditional artifacts, and backend behavior remain unchanged.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/homologate-sdd-execution-persistence-contracts/proposal.md` | Intent, scope, non-goals, affected areas, risks, rollback, and success criteria. |
| Spec | `openspec/changes/homologate-sdd-execution-persistence-contracts/specs/sdd-execution-persistence-contracts/spec.md` | Shared execution/persistence boundary, phase artifact contracts, conflict handling, and compatibility preservation. |
| Spec | `openspec/changes/homologate-sdd-execution-persistence-contracts/specs/sdd-security-applicability-workflow/spec.md` | Security applicability artifact/routing contract and no-impact routing compatibility. |
| Spec | `openspec/changes/homologate-sdd-execution-persistence-contracts/specs/sdd-security-design-workflow/spec.md` | Conditional security-design artifact behavior and delegated persistence boundary. |
| Spec | `openspec/changes/homologate-sdd-execution-persistence-contracts/specs/sdd-test-design-workflow/spec.md` | Mandatory test-design artifact contract and downstream consumption obligations. |
| Security Applicability | `openspec/changes/homologate-sdd-execution-persistence-contracts/security-applicability.md` | Confirms `classification: no-impact`, `securityImpact: false`, no applicable guidelines, and no required `security-design.md`. |
| Design | `openspec/changes/homologate-sdd-execution-persistence-contracts/design.md` | Architecture decisions, data flow, file-change plan, contract template, testing constraints, rollout risks, and security routing. |
| OpenSpec Config | `openspec/config.yaml` | Confirms no test runner, linter, formatter, type checker, or coverage command; applies static/manual verification only. |
| Security Design | Not required | Security applicability is no-impact, so there are no security-design controls to cover. |

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Authoritative Persistence Boundary / Design: Persistence authority | Verify `skills/_shared/persistence-contract.md` is the single detailed authority for artifact-store modes, backend read/write semantics, artifact resolution, hybrid conflict handling, state persistence, and persistence verification. | static | mandatory | Diff/manual review showing detailed mode semantics centralized in `skills/_shared/persistence-contract.md`. | No runtime command is required or available. |
| TD-002 | Spec: Shared persistence owns mode behavior / Proposal success criteria | Verify phase skills and shared execution files delegate common backend mechanics instead of redefining detailed `engram`, `openspec`, `hybrid`, or `none` mode behavior. | static | mandatory | Static search/manual review showing no duplicated contradictory phase-local persistence prose remains, and references point to the shared persistence authority. | Broad references are acceptable; competing algorithms are not. |
| TD-003 | Spec: Backend convention files remain scoped / Design: backend convention cleanup | Verify `skills/_shared/engram-convention.md` and OpenSpec convention text remain backend-specific references and do not compete with the shared persistence authority. | manual | mandatory | Contract review notes comparing backend convention wording with `persistence-contract.md`. | Backend-specific details may remain when scoped as implementation/reference details. |
| TD-004 | Spec: Execution Contract Boundary / Design: Execution boundary | Verify `skills/_shared/sdd-phase-common.md` keeps executor boundary, supplemental skill loading, Section D envelope, routing-token conventions, artifact naming reminders, and review workload guard, without duplicating detailed persistence algorithms. | static | mandatory | Manual diff review of `skills/_shared/sdd-phase-common.md` against the execution-boundary responsibilities. | Section D field names and routing mappings must remain stable. |
| TD-005 | Spec: Executor returns a stable envelope / Test-design downstream routing | Verify the Section D envelope fields and native/status routing token mapping are preserved, including `next_recommended` phase-agent/native normalization behavior. | manual | mandatory | Contract matrix showing envelope fields and routing token mappings remain unchanged. | Protects orchestrator compatibility. |
| TD-006 | Spec: Phase Artifact Contracts / Design: phase contract template | Verify each modified SDD phase skill includes a compact `## Phase Artifact Contract` or equivalent table naming common backend mechanics, required inputs, produced artifact, mutations, conditional behavior, success routing, and block routing. | static | mandatory | Manual review checklist for each changed `skills/sdd-*/SKILL.md`. | Phase-specific validation may remain outside the table. |
| TD-007 | Spec: Phase-specific mutation is preserved / Design: mutation-heavy phases | Verify `sdd-apply`, `sdd-archive`, and `sdd-init` keep explicit mutation semantics for existing artifacts, state updates, archive moves, local support artifacts, and task progress updates. | manual | mandatory | Contract comparison showing mutation obligations remain phase-local and were not lost during deduplication. | These mutations are phase behavior, not generic persistence mechanics. |
| TD-008 | Spec: Conflict and Ambiguity Resolution / Scenario: duplicated rule is homologated | Verify ambiguous or duplicated persistence rules are resolved by naming the authoritative owner and preserving existing behavior unless the rule was previously contradictory or undefined. | manual | mandatory | Review notes listing resolved duplicated rules and their authoritative owner. | Any behavior clarification must not redesign the DAG or backend formats. |
| TD-009 | Spec: Compatibility is preserved / Proposal non-goals / Design: Compatibility | Verify established Engram keys, OpenSpec paths, routing tokens, camelCase state/status fields, native state names, DAG order, and backend behavior are preserved without migration. | manual | mandatory | Compatibility matrix covering artifact keys, paths, routing tokens, state fields such as `nextRecommended`, `artifactRefs`, `testDesign`, `securityDesign`, and backend behavior. | This is the primary regression guard. |
| TD-010 | Spec: Security Applicability artifact/routing / Security applicability input | Verify `security-applicability.md` remains the phase artifact with classification, evidence, guideline mapping, risks, and routing recommendation, while persistence mode semantics are delegated. | static | mandatory | Static/manual review of `sdd-security-applicability` contract wording and artifact identity. | Must preserve no-impact and security-impacting routing behavior. |
| TD-011 | Spec: No-impact routing compatibility / Security Design conditional artifact behavior | Verify no-impact security applicability continues to skip `security-design.md`, and missing `security-design.md` is not treated as a blocker solely because persistence wording changed. | manual | mandatory | Manual routing review showing `securityImpact: false` routes to design/test-design without requiring `security-design.md`. | Applies across supported artifact-store modes. |
| TD-012 | Spec: Security Design artifact contract / conditional behavior | Verify `sdd-security-design` still creates `security-design.md` only for security-impacting changes and still maps applicable guidelines to controls, evidence, mandatory status, residual risks, and downstream obligations when required. | manual | mandatory | Contract review of `sdd-security-design` artifact contract and conditional behavior. | No security-design artifact is required for this change, but the workflow contract must remain intact. |
| TD-013 | Spec: test-design.md Artifact Contract | Verify `sdd-test-design` continues to require `test-design.md` for every change and maps spec scenarios, design risks, security controls when required, check type, severity, expected evidence, and no-impact assessments. | static | mandatory | Manual review of the `sdd-test-design` contract and this artifact. | Ensures this phase remains mandatory rather than optional. |
| TD-014 | Spec: Downstream Consumption | Verify `sdd-tasks`, `sdd-apply`, and `sdd-verify` still consume `test-design.md` as the test-planning source of truth, including blockers for omitted mandatory cases and deviation handling. | manual | mandatory | Cross-phase contract review showing downstream tasks/evidence/verification references to established `test-design.md` key/path. | Tasks must not route directly from design when test-design is required. |
| TD-015 | Design: Testing Strategy / OpenSpec config testing capabilities | Verify implementation and verification plans report unavailable runtime tests explicitly and do not invent test, lint, format, type-check, build, or coverage commands. | static | mandatory | Review of `openspec/config.yaml` testing section and downstream task/verify evidence plan. | Only static/manual contract checks are valid for this repository. |
| TD-016 | Proposal risk: Review overload / Design rollout | Verify task planning slices the Markdown-only implementation into reviewable work units under the configured 400 changed-line budget and preserves stacked-to-main delivery. | manual | non-mandatory | Task-plan review against `delivery.strategy: auto-chain`, `chain_strategy: stacked-to-main`, and `review_budget_lines: 400`. | Advisory for test design, blocking only if tasks exceed process constraints without escalation. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| Not applicable | No security controls are required because `security-applicability.md` classifies this Markdown-only contract change as `no-impact` with `securityImpact: false` and no applicable guidelines. | No | TD-011 verifies no-impact routing compatibility; TD-012 verifies conditional security-design behavior remains intact for future security-impacting changes. | not-applicable | None |

## No-Impact Security Assessment

Security applicability is no-impact. The change only updates SDD Markdown contracts, phase skill prose, shared persistence wording, and OpenSpec specification artifacts. It explicitly excludes runtime code, persistence backend redesign, key-format changes, authentication, sessions, sensitive data, secrets, permissions/access control, file handling behavior, database access, and sensitive logging. Therefore no `security-design.md` is required and there are no security controls or mandatory security evidence items to cover.

## Evidence Expectations

- Mandatory cases require static/manual evidence in implementation or verification artifacts; unavailable runtime tooling must be reported explicitly rather than treated as passing evidence.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Evidence must preserve established keys, paths, routing tokens, state/status field names, conditional security/test-design behavior, and backend behavior unless the spec explicitly authorizes a clarification.
- Static evidence may include changed-file review, structured contract matrices, and targeted text searches. It must not rely on invented test, lint, format, type-check, build, or coverage commands.

## Open Questions

- None.
