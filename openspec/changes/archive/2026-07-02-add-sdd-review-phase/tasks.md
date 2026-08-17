# Tasks: Add SDD Review Phase

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
| 1 | Shared status, persistence, and OpenSpec contracts | PR 1 | Base: main; enables routing. |
| 2 | Review executor, skill, report contract, and catalog | PR 2 | Base: PR 1; owns matrix. |
| 3 | Apply/verify/archive/orchestrator/README integration | PR 3 | Base: PR 2; completes adoption. |

## Phase 1: Shared Contracts / PR 1

- [x] 1.1 Update `skills/_shared/sdd-status-contract.md` with native token `review`, `reviewReport` refs/paths/state, and `apply -> review -> verify -> archive` routing for TD-001, TD-002, TD-004.
- [x] 1.2 Update `skills/_shared/persistence-contract.md` with OpenSpec `openspec/changes/{change-name}/review-report.md`, Engram `sdd/{change-name}/review`, and missing-evidence blockers for TD-003, TD-004, TD-013.
- [x] 1.3 Update `skills/_shared/openspec-convention.md` with `review-report.md` tree, path table, read list, and archive expectations for TD-003, TD-015.
- [x] 1.4 Record static evidence that source specs remain archive sync targets: `openspec/specs/sdd-review-workflow/spec.md`, `sdd-execution-persistence-contracts`, and `sdd-security-guideline-catalog`.

### PR 1 Static Evidence

| Task | Evidence |
|------|----------|
| 1.1 | `skills/_shared/sdd-status-contract.md` now includes native token `review`, `reviewReport` refs/paths/context/artifact state, a `review` dependency state, launch mapping to `sdd-review`, and the `apply -> review -> verify -> archive` routing order. |
| 1.2 | `skills/_shared/persistence-contract.md` now resolves review evidence as Engram `sdd/{change-name}/review` and OpenSpec `openspec/changes/{change-name}/review-report.md`, and treats missing/ambiguous/blocking/unreadable review evidence plus unsafe review inputs as blockers. |
| 1.3 | `skills/_shared/openspec-convention.md` now lists `review-report.md` in the change tree, artifact path table, read list, and archive preservation expectations. |
| 1.4 | Archive sync targets remain the main specs at `openspec/specs/sdd-review-workflow/spec.md`, `openspec/specs/sdd-execution-persistence-contracts/spec.md`, and `openspec/specs/sdd-security-guideline-catalog/spec.md` when matching deltas exist under the active change. Static inspection found existing source specs for `sdd-execution-persistence-contracts` and `sdd-security-guideline-catalog`; `sdd-review-workflow` is a new domain expected to be created by archive sync from this change's delta. |

## Phase 2: Review Executor and Catalog / PR 2

- [x] 2.1 Create `agents/sdd/sdd-review.md` using existing SDD executor prompt conventions and `skills/sdd-review/SKILL.md` loading.
- [x] 2.2 Create `skills/sdd-review/SKILL.md` with required inputs, readiness blockers, persistence rules, and return routing for TD-011, TD-012, TD-013.
- [x] 2.3 Add the `review-report.md` contract/template with verdict, blocking summary, evidence summary, next recommendation, and exact matrix header for TD-005, TD-006.
- [x] 2.4 Add the stable 96-control catalog/mapping in `skills/sdd-review/SKILL.md`; manually evidence 96 unique Item IDs and category preservation for TD-007, TD-008.
- [x] 2.5 Constrain `Complies` to `Yes`, `No`, or `N/A`; require N/A evidence location and comments for TD-009, TD-010.
- [x] 2.6 Update `skills/_shared/security-guideline-catalog.md` with review-safe cross-reference/evidence guidance without duplicating authority for TD-016, TD-017.

### PR 2 Static Evidence

| Task | Evidence |
|------|----------|
| 2.1 | `agents/sdd/sdd-review.md` follows existing SDD executor prompt conventions: executor boundary, no delegation, no inline Skill tool, and `%USERPROFILE%/.config/opencode/skills/sdd-review/SKILL.md` phase skill loading. |
| 2.2 | `skills/sdd-review/SKILL.md` defines required inputs, produced artifact identities, readiness blockers for missing artifacts/unknown changed files/unsafe workspace/invalid catalog/persistence failure, and routing to `verify`, `apply`, or `resolve-blockers` for TD-011 through TD-013. |
| 2.3 | `skills/sdd-review/references/report-template.md` defines required verdict, blocking summary, evidence summary, next recommendation, and the exact matrix header: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. |
| 2.4 | `skills/sdd-review/SKILL.md` references the stable corporate checklist catalog at `skills/sdd-review/references/control-catalog.md`. Manual inspection evidence: catalog contains exactly 96 unique Item IDs (`REV-CORP-001..REV-CORP-096`) mapped one-to-one to the user's corporate code-review checklist source items 1 through 96, with catalog fields for Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, default Complies, evidence hints, and notes. |
| 2.5 | `skills/sdd-review/SKILL.md` and `skills/sdd-review/references/report-template.md` constrain `Complies` to `Yes`, `No`, or `N/A`; `N/A` rows require Evidence Location plus Observations/Comments proving irrelevance. |
| 2.6 | `skills/_shared/security-guideline-catalog.md` now adds review-phase cross-reference guidance and review-safe evidence types while preserving security applicability/design and the catalog as the authority for guideline IDs, mandatory evidence, and exceptions. |

## Phase 3: Workflow Integration and Evidence / PR 3

- [x] 3.1 Update `agents/sdd/sdd-orchestrator.md` commands, flowchart, dependency graph, state persistence, routing, and archive readiness for TD-001, TD-002, TD-015.
- [x] 3.2 Update `skills/sdd-apply/SKILL.md` so completed apply recommends `review`, not direct `verify`, for TD-001.
- [x] 3.3 Update `skills/sdd-verify/SKILL.md` to require/cite non-blocking `review-report.md` without owning the full matrix for TD-014, TD-019.
- [x] 3.4 Update `skills/sdd-archive/SKILL.md` to require non-blocking review plus passing verify before archive for TD-015.
- [x] 3.5 Update `README.md` with the new SDD phase order and review artifact.
- [x] 3.6 Manually inspect `scripts/sdd_init_agents.ps1` and `scripts/sdd_init_skills.ps1`; record no-logic-change recursive sync evidence for TD-018.
- [x] 3.7 Record static/manual evidence for TD-001 through TD-019 and explicitly report unavailable runtime checks from `openspec/config.yaml`.

### PR 3 Static and Manual Evidence

| Case | Evidence |
|------|----------|
| TD-001 | `agents/sdd/sdd-orchestrator.md` now shows apply success flowing to `sdd-review` in the flowchart and status routing table, and `skills/sdd-apply/SKILL.md` now returns `next_recommended: review` when all implementation tasks are complete or `applyState` is `all_done`. No direct apply-to-verify success route remains in `skills/sdd-apply/SKILL.md`. |
| TD-002 | `agents/sdd/sdd-orchestrator.md` dependency graph is now `apply -> review -> verify -> archive`; status routing, phase readiness, state schema tokens, context table, and recovery rows include `review` / `review-report` between apply and verify. `skills/_shared/sdd-status-contract.md` already contains the shared `review` token and DAG from PR 1. |
| TD-003 | `skills/_shared/openspec-convention.md` and `skills/_shared/persistence-contract.md` from PR 1 define OpenSpec review persistence as `openspec/changes/{change-name}/review-report.md`; `skills/sdd-review/SKILL.md` from PR 2 produces that artifact. |
| TD-004 | `skills/_shared/persistence-contract.md` from PR 1 resolves review as Engram/hybrid key `sdd/{change-name}/review` and OpenSpec `openspec/changes/{change-name}/review-report.md`; `skills/_shared/sdd-status-contract.md` exposes `reviewReport` refs/paths/context/artifact status. |
| TD-005 | `skills/sdd-review/references/report-template.md` from PR 2 includes verdict, blocking summary, evidence summary, next recommendation, and a matrix section. |
| TD-006 | `skills/sdd-review/references/report-template.md` from PR 2 uses the exact matrix header: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments. |
| TD-007 | PR 2 manual catalog evidence recorded exactly 96 unique Item IDs (`REV-CORP-001..REV-CORP-096`) in `skills/sdd-review/references/control-catalog.md`; this PR preserved the catalog unchanged. |
| TD-008 | PR 2 manual catalog evidence recorded category preservation through the stable catalog without adding a matrix column; this PR preserved the catalog unchanged. |
| TD-009 | `skills/sdd-review/SKILL.md` and `skills/sdd-review/references/report-template.md` from PR 2 constrain `Complies` to `Yes`, `No`, or `N/A`; this PR preserved that contract. |
| TD-010 | `skills/sdd-review/SKILL.md` and `skills/sdd-review/references/report-template.md` from PR 2 require Evidence Location plus Observations/Comments for every `N/A`; this PR preserved that contract. |
| TD-011 | `skills/sdd-review/SKILL.md` from PR 2 routes blocking or critical failed controls to `next_recommended: apply` and requires failed controls plus affected requirements. |
| TD-012 | `skills/sdd-review/SKILL.md` from PR 2 routes reports with only non-blocking findings to `next_recommended: verify` and keeps warnings in the persisted report. |
| TD-013 | `skills/sdd-review/SKILL.md` from PR 2 routes missing required artifacts, unknown changed files, unsafe workspace context, invalid catalog shape, or persistence failure to `resolve-blockers`; `skills/_shared/persistence-contract.md` also records downstream missing/ambiguous review evidence blockers. |
| TD-014 | `skills/sdd-verify/SKILL.md` now requires non-blocking `review-report.md` / `sdd/{change-name}/review`, cites verdict/blocking/evidence summary, blocks missing/ambiguous/unreadable/blocking review evidence, and explicitly forbids duplicating or owning the full 96-control matrix. |
| TD-015 | `agents/sdd/sdd-orchestrator.md` archive readiness now requires non-blocking review-report plus passing verify-report; `skills/sdd-archive/SKILL.md` now blocks missing/ambiguous/blocking review evidence and requires non-blocking review plus passing verification before archive. |
| TD-016 | `skills/_shared/security-guideline-catalog.md` from PR 2 allows review-phase cross-references while preserving security applicability/design/catalog authority; `skills/sdd-review/SKILL.md` keeps the same security boundary. |
| TD-017 | `skills/_shared/security-guideline-catalog.md` from PR 2 provides review-safe evidence guidance and cross-reference behavior without duplicating or redefining security guideline authority; this PR did not modify the corporate/security catalog authority text. |
| TD-018 | Manual inspection found `scripts/sdd_init_agents.ps1` copies the top-level `agents` tree using `Get-ChildItem -LiteralPath $sourceAgentsPath -Force | Copy-Item -Destination $target.DestinationPath -Recurse -Force`; `scripts/sdd_init_skills.ps1` copies the top-level `skills` tree with the same recursive `Copy-Item -Recurse -Force` pattern. Adding `agents/sdd/sdd-review.md` and `skills/sdd-review/` needs no script logic change. |
| TD-019 | `openspec/config.yaml` explicitly reports no runtime test runner, unit/integration/e2e tool, coverage command, linter, type checker, or formatter. Runtime checks are unavailable by configuration: `testing.test_runner.available: false`, `command: ""`, quality commands are empty, and verify notes require reporting unavailable runtime tests instead of treating missing tools as passing evidence. `skills/sdd-verify/SKILL.md` now preserves this by requiring unavailable runtime evidence to be reported explicitly and by not inventing commands. |

### PR 3 Runtime Check Availability

No runtime runner, build, linter, typechecker, formatter, coverage, unit, integration, or e2e command is configured in `openspec/config.yaml`. This apply batch therefore used static/manual evidence only, as required by `test-design.md` and the launch testing capabilities.
