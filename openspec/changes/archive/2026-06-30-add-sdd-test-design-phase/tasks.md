# Tasks: Add Mandatory SDD Test Design Phase

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 650-900 |
| Review budget lines | 400 |
| 400-line budget risk | High |
| Review budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1: contract -> PR 2: routing -> PR 3: verification |
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
| 1 | Add executor, skill, and shared artifact contracts | PR 1 | Base `main`; creates phase surface without live routing. |
| 2 | Wire workflow routing, state/status, and downstream phase consumption | PR 2 | Base `main` after PR 1; enables `design -> test-design -> tasks`. |
| 3 | Update onboarding, registry artifacts, and static verification evidence | PR 3 | Base `main` after PR 2; documents and validates the workflow. |

## Phase 1: Phase Contract and Shared Artifacts

- [x] 1.1 Create `agents/sdd/sdd-test-design.md` with executor boundary, required skill-file read, no delegation, and no inline `skill()` calls.
- [x] 1.2 Create `skills/sdd-test-design/SKILL.md` defining inputs `proposal/spec/design`, `test-design.md` format, validation, persistence modes, and `next_recommended: tasks`.
- [x] 1.3 Update `skills/_shared/openspec-convention.md` to include `test-design.md` in active change structure, artifact paths, reading rules, and archive contents.
- [x] 1.4 Update `skills/_shared/sdd-phase-common.md` with `test-design` <-> `sdd-test-design` routing and artifact naming guidance.
- [x] 1.5 Update `skills/_shared/persistence-contract.md` with Engram/OpenSpec resolver rows, `artifactRefs.testDesign`, recovery fallback, and sub-agent context rules.

### Apply Progress — Slice 1 / Work Unit 1

- Completed tasks: 1.1, 1.2, 1.3, 1.4, 1.5.
- Scope boundary: PR 1 / Work Unit 1 only — executor, skill, and shared artifact contracts. Live orchestrator routing and downstream phase behavior remain pending for later slices.
- Mode: Standard; strict TDD is disabled and no runtime test runner is configured.
- Verification: static/file-contract review of touched Markdown artifacts only.

## Phase 2: Workflow Routing and Downstream Consumption

- [x] 2.1 Update `skills/_shared/sdd-status-contract.md` with `testDesign` fields, dependency states, artifact status, context files, and native/status token mapping.
- [x] 2.2 Update `agents/sdd/sdd-orchestrator.md` flowchart, DAG, routing table, launch metadata, state schema, gatekeeper, recovery, and status handling for `design -> test-design -> tasks`.
- [x] 2.3 Update `skills/sdd-design/SKILL.md` so successful design returns `next_recommended: test-design`.
- [x] 2.4 Update `skills/sdd-tasks/SKILL.md` to require `test-design.md`, block missing artifacts, and derive testing/evidence tasks from planned cases.
- [x] 2.5 Update `skills/sdd-apply/SKILL.md` to read `test-design.md`, follow planned checks, and document justified deviations in apply evidence.
- [x] 2.6 Update `skills/sdd-verify/SKILL.md` and `skills/sdd-verify/references/report-format.md` to fail uncovered mandatory cases and warn for uncovered non-mandatory cases.
- [x] 2.7 Update `skills/sdd-archive/SKILL.md` to require/archive `test-design.md` unless an explicit partial archive exception applies.

### Apply Progress — Slice 2 / Work Unit 2

- Completed tasks: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7.
- Cumulative completed tasks: 1.1, 1.2, 1.3, 1.4, 1.5, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7.
- Scope boundary: PR 2 / Work Unit 2 only — workflow routing, status/state contracts, downstream phase consumption, apply evidence, verify coverage semantics, and archive requirements. Onboarding, registry refresh, and final static verification tasks remain pending for Slice 3.
- Mode: Standard; strict TDD is disabled and no runtime test runner is configured.
- Verification: static/file-contract review of touched Markdown artifacts only; no runtime tests were available.

## Phase 3: Onboarding, Registry, and Static Verification

- [x] 3.1 Update `agents/sdd/sdd-onboard.md` and `skills/sdd-onboard/SKILL.md` with the new phase order and artifact list.
- [x] 3.2 Refresh generated skill registry artifacts if present so `sdd-test-design` is discoverable.
- [x] 3.3 Static-check all touched Markdown for `test-design`, `sdd-test-design`, and `testDesign` consistency across routing, state/status, persistence, archive, and downstream phase contracts.
- [x] 3.4 Static-check spec scenarios: successful design routes to test design, tasks block when missing, artifact contract exists, fail-vs-warn coverage semantics, continuation refs, and native-token compatibility.

### Apply Progress — Slice 3 / Work Unit 3

- Completed tasks: 3.1, 3.2, 3.3, 3.4.
- Cumulative completed tasks: 1.1, 1.2, 1.3, 1.4, 1.5, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 3.1, 3.2, 3.3, 3.4.
- Scope boundary: PR 3 / Work Unit 3 only — onboarding workflow narration, local generated skill registry refresh, and final static/file-contract verification evidence. No routing, downstream executor, archive, or runtime-test behavior was changed in this slice.
- Mode: Standard; strict TDD is disabled and no runtime test runner is configured.
- Registry refresh: verified parent path `.atl/` exists, then ran `gentle-ai skill-registry refresh --force` from the repository root. Result: `Skill registry refreshed (10 skills): C:\Leo\Proyectos\leo3nava_ai\leo3nava_ai\.atl\skill-registry.md`. The generated registry intentionally excludes `sdd-*`, `_shared`, and `skill-registry` phase skills per `skills/skill-registry/SKILL.md`; `sdd-test-design` discoverability for phase routing is provided by the dedicated agent/skill files and orchestrator/status contracts rather than the supplemental skill registry.
- Verification: static/file-contract review only; no runtime tests were available. Checked that onboarding now includes `design -> test-design -> tasks`, `test-design.md` in the artifact recap, and `sdd-test-design` in delegated phase order. Checked touched Markdown contracts for `test-design` native/artifact token, `sdd-test-design` phase/agent token, and `testDesign` persisted camelCase field usage across routing, state/status, persistence, archive, and downstream phase contracts. Checked spec scenario coverage against implementation contracts for successful design routing, missing-artifact task blocking, artifact contract, mandatory fail vs non-mandatory warning semantics, continuation refs, and native-token compatibility.
