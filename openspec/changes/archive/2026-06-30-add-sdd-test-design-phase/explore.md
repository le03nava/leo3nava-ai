# Exploration: Add SDD Test Design Phase

## Current State

The SDD workflow currently moves from `sdd-design` directly to `sdd-tasks`: `explore? -> proposal -> spec -> design -> tasks -> apply -> verify -> archive`. Specs define testable behavior, design defines implementation approach and a high-level testing strategy, and `sdd-tasks` derives implementation/testing tasks from specs + design. Verification later checks proposal/spec/design/tasks plus runtime evidence.

There is no dedicated artifact or phase for designing concrete test cases before implementation task planning. Test intent is currently split across spec scenarios, `design.md` testing strategy, task testing items, apply Strict TDD evidence, and verify compliance matrices.

## Affected Areas

- `agents/sdd/sdd-orchestrator.md` — central phase DAG, routing tokens, readiness gates, launch envelope, state persistence schema, recovery fallback, gatekeeper rules, and user-facing flow references all encode `design -> tasks`.
- `agents/sdd/sdd-onboard.md` and `skills/sdd-onboard/SKILL.md` — narrated walkthrough hardcodes the phase order and step numbering.
- `agents/sdd/*.md` — a new executor prompt would likely be needed, e.g. `sdd-test-design.md`, following the existing executor boundary pattern.
- `skills/sdd-test-design/SKILL.md` — a new phase skill would be needed to define the test-design artifact contract, dependencies, output format, persistence, and routing.
- `skills/sdd-design/SKILL.md` — successful design currently routes to `tasks`; it should route to the new test-design phase if the new phase becomes mandatory.
- `skills/sdd-tasks/SKILL.md` — dependencies should include test-design, and task generation should derive testing tasks from the test-design artifact rather than inventing them from only spec/design.
- `skills/sdd-apply/SKILL.md` — apply should read test-design when creating tests or following Strict TDD, and report deviations from planned test cases.
- `skills/sdd-verify/SKILL.md` and `skills/sdd-verify/references/report-format.md` if present — verification should compare implementation/runtime evidence against planned test cases in addition to specs/design/tasks.
- `skills/_shared/openspec-convention.md` — needs a path for the new artifact, likely `openspec/changes/{change-name}/test-design.md`.
- `skills/_shared/persistence-contract.md` — needs Engram/OpenSpec/Hybrid refs, state `artifactRefs`, state phase/token enum, and recovery rules for the new artifact.
- `skills/_shared/sdd-status-contract.md` — needs `artifactRefs`, `artifactPaths`, `contextFiles`, artifact status, dependencies, routing tokens, and dependency-state rules for the new phase.
- `skills/_shared/sdd-phase-common.md` — needs routing-token mapping and artifact naming guidance for `test-design`.
- `.atl/skill-registry.md` / generated registry artifacts, if present — should be refreshed after adding the skill so orchestrator skill resolution can discover supplemental skills while keeping fixed SDD phase skills out of supplemental matching.

## Approaches

1. **Mandatory phase between design and tasks** — Add `sdd-test-design` as a first-class phase after `sdd-design` and before `sdd-tasks`.
   - Pros: clear artifact boundary; tasks can be derived from spec + design + explicit test cases; verification can check planned coverage; matches the user's requested placement.
   - Cons: touches many contracts and state schemas; every status/routing enum must be updated consistently; existing active changes may need compatibility handling.
   - Effort: Medium/High.

2. **Section inside `design.md` only** — Expand `sdd-design` to include detailed test cases, then keep `sdd-tasks` unchanged structurally.
   - Pros: smaller implementation; no new token/state/status value; fewer compatibility risks.
   - Cons: weaker separation of concerns; test case design remains buried in technical design; tasks still depend on one large design artifact; does not satisfy the requested sub-agent/phase.
   - Effort: Low/Medium.

3. **Optional phase triggered by risk** — Add `sdd-test-design` but make it optional for high-risk or test-heavy changes only.
   - Pros: avoids process overhead for simple changes; useful for changes where explicit test planning matters.
   - Cons: more complex routing; tasks must support both dependency shapes; optionality may lead to skipped test design where it would help.
   - Effort: High.

## Recommendation

Use Approach 1: add a mandatory first-class `sdd-test-design` phase between `sdd-design` and `sdd-tasks`.

The phase should produce a concise `test-design.md` artifact that maps spec scenarios to planned test cases, test layers, fixtures/data, commands/evidence expectations, Strict TDD implications when enabled, and known manual/unavailable-test gaps. `sdd-tasks` should require this artifact and convert it into concrete implementation/testing tasks. `sdd-verify` should read it and report whether planned test cases were implemented/executed or intentionally skipped.

Recommended phase order:

```text
explore? -> proposal -> spec -> design -> test-design -> tasks -> apply -> verify -> archive
```

Recommended artifact naming:

- Phase/agent: `sdd-test-design`
- Engram key: `sdd/{change-name}/test-design`
- OpenSpec path: `openspec/changes/{change-name}/test-design.md`
- Status/state token: `test-design` if hyphenated tokens are accepted by native tooling; otherwise `testDesign` for persisted camelCase fields and `test-design` for phase envelopes. This needs a deliberate compatibility decision before implementation.

## Risks

- **CRITICAL compatibility risk**: Native dispatcher/status tooling may have a bounded token enum that currently excludes `test-design`. If the binary/runtime is external to this repo, Markdown-only changes may make prompts disagree with native status output.
- **State migration risk**: Existing SDD state schemas enumerate `currentPhase`, `completedPhases`, `artifactRefs`, and `nextRecommended`. All must be updated together or `/sdd-continue` can route incorrectly.
- **Artifact naming risk**: Hyphenated phase tokens are readable but may be harder for camelCase YAML/JSON fields. Decide the machine token convention once and apply it everywhere.
- **Scope creep risk**: Test design can become a full test plan document. Keep it compact and scenario-linked so it improves implementation tasks instead of adding process weight.
- **Duplication risk**: Specs already contain scenarios and design already contains testing strategy. The new artifact must define concrete test cases/evidence, not restate requirements.
- **Verification risk**: This repository has no runnable test runner in `openspec/config.yaml`; verification for this change will likely rely on static/file-contract checks unless a separate validation command exists.

## Ready for Proposal

Yes. The proposal should define the new phase as mandatory between `sdd-design` and `sdd-tasks`, with explicit compatibility work for routing/status/state/persistence contracts and a concise artifact contract for `test-design.md`.
