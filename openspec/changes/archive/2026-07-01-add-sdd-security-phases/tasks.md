# Tasks: Add SDD Security Phases

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 900-1300 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 -> PR 2 -> PR 3 |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: High
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Add shared security catalog/contract and token plumbing | PR 1 | Base `main`; includes status, persistence, OpenSpec path contracts. |
| 2 | Add security applicability/design agents and skills | PR 2 | Base PR 1 branch; depends on shared contracts. |
| 3 | Integrate downstream phases and docs | PR 3 | Base PR 2 branch; wires test-design, tasks, apply, verify, archive, README. |

## Phase 1: Shared Contracts and Token Plumbing

- [x] 1.1 Create `skills/_shared/security-guideline-catalog.md` with snapshot metadata, stable guideline IDs, taxonomy categories, mandatory flags, expected evidence, and audit notes.
- [x] 1.2 Create `skills/_shared/sdd-security-contract.md` defining applicability/design artifact schemas, classification values, evidence statuses, and approved exception fields.
- [x] 1.3 Update `skills/_shared/sdd-status-contract.md` with `security-applicability` and `security-design` tokens, dependency states, and `securityApplicability`/`securityDesign` refs.
- [x] 1.4 Update `skills/_shared/persistence-contract.md` and `skills/_shared/openspec-convention.md` with Engram keys and OpenSpec paths for both security artifacts.

## Phase 2: New Security Phase Executors

- [x] 2.1 Create `agents/sdd/sdd-security-applicability.md` enforcing always-run launch, blocking rules, and routing from the applicability spec scenarios.
- [x] 2.2 Create `skills/sdd-security-applicability/SKILL.md` to read proposal/specs/catalog and write `security-applicability.md` with no-impact or guideline mapping evidence.
- [x] 2.3 Create `agents/sdd/sdd-security-design.md` enforcing conditional execution after technical design for security-impacting changes.
- [x] 2.4 Create `skills/sdd-security-design/SKILL.md` mapping applicable guidelines to controls, mandatory evidence, residual risks, and approved exceptions.

## Phase 3: Orchestration and Downstream Integration

- [x] 3.1 Update `agents/sdd/sdd-orchestrator.md` with DAG `spec -> security-applicability -> design -> security-design? -> test-design` and readiness gates.
- [x] 3.2 Update `skills/sdd-design/SKILL.md` to require applicability before design and route to security-design or test-design.
- [x] 3.3 Update `skills/sdd-test-design/SKILL.md` to consume `security-design.md` when required and block uncovered mandatory controls.
- [x] 3.4 Update `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, and `skills/sdd-archive/SKILL.md` to carry security evidence and block archive on missing mandatory evidence without complete exceptions.
- [x] 3.5 Update `README.md` to list the two new SDD security phases and their order.

## Phase 4: Static Verification

- [x] 4.1 Manually verify YAML frontmatter and token/path consistency across new and modified agent/skill files.
- [x] 4.2 Validate spec scenarios: impact/no-impact routing, missing security-design blocker, archive blocker, and complete exception allowance.
