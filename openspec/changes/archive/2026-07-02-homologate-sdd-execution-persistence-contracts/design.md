# Design: Homologate SDD Execution and Persistence Contracts

## Technical Approach

Make `skills/_shared/persistence-contract.md` the only detailed source for artifact-store mode behavior, backend read/write semantics, resolver behavior, hybrid conflict handling, and persistence verification. Keep `skills/_shared/sdd-phase-common.md` focused on executor mechanics and return envelopes. Each phase `SKILL.md` will keep a compact artifact contract that names its inputs, outputs, keys, paths, mutations, validation, and routing without restating backend algorithms.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Persistence authority | Centralize detailed persistence rules in `skills/_shared/persistence-contract.md`. | Keep duplicated mode prose in every phase. | Satisfies the spec boundary and prevents future drift while preserving existing behavior. |
| Execution boundary | Reduce `sdd-phase-common.md` persistence sections to pointers plus executor/envelope/routing/review rules. | Move all phase behavior into common. | Common stays reusable; phase-specific validation and routing stay local. |
| Phase contracts | Add a standard compact `## Phase Artifact Contract` table to phase skills. | Free-form phase prose. | Reviewers can verify artifact identity, mutations, and routing without rereading shared backend rules. |
| Compatibility | Preserve DAG, Engram keys, OpenSpec paths, routing tokens, camelCase state/status fields, and backend behavior. | Rename fields or simplify routing. | Specs explicitly require no migration and no consumer breakage. |

## Data Flow

```text
Orchestrator status
  ├─ artifact_store.mode ──→ persistence-contract.md ──→ backend resolver/read/write
  └─ nextRecommended/state ─→ sdd-phase-common.md/status contract ─→ phase launch

Phase SKILL.md
  ├─ reads Phase Artifact Contract for required inputs/outputs
  ├─ delegates common backend mechanics to persistence-contract.md
  └─ returns Section D envelope with next_recommended
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/_shared/persistence-contract.md` | Modify | Mark as authoritative; ensure mode semantics, resolver, hybrid policy, state persistence, and verification are complete and non-duplicative. |
| `skills/_shared/sdd-phase-common.md` | Modify | Replace detailed Section B/C backend algorithms with references to persistence authority; preserve Section A, Section D, naming, routing, and review workload guard. |
| `skills/_shared/engram-convention.md` | Modify | Reframe as backend-specific reference; remove claims that compete with shared persistence authority. |
| `skills/sdd-*/SKILL.md` | Modify | Replace repeated mode prose with standard phase artifact contracts. Preserve phase-local validation, mutations, and routing. |
| `skills/sdd-init/SKILL.md` | Modify | Keep initialization-specific local support artifact behavior explicit while delegating common mode mechanics. |
| `.atl/skill-registry.md` | Possible follow-up | Refresh only if skill metadata, triggers, or paths change; expected not required for prose-only contract edits. |
| `README.md` or docs indexes | Possible follow-up | Update only if they duplicate old persistence authority language; otherwise leave unchanged. |

## Interfaces / Contracts

Phase skills should use this structure:

```markdown
## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | `artifact-key` / OpenSpec path list |
| Produced artifact | `artifact-key` / OpenSpec path |
| Mutates | Existing artifact/state mutations, or `None` |
| Conditional behavior | Required/not-required rules |
| Success routing | `next_recommended: ...` |
| Block routing | `next_recommended: resolve-blockers` or phase-specific retry |
```

`sdd-apply`, `sdd-archive`, and `sdd-init` must keep mutation semantics explicit because those are phase behavior, not generic persistence mechanics.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static review | No duplicate contradictory persistence authority remains. | Grep for old mode prose and verify it delegates to `persistence-contract.md`. |
| Contract review | Artifact keys, paths, routing tokens, state/status field names are unchanged. | Manual matrix against `sdd-status-contract.md`, `openspec-convention.md`, and specs. |
| Runtime tests | Not available. | `openspec/config.yaml` reports no test runner, linter, formatter, or type checker; verify by structured document review. |

## Migration / Rollout

No data migration required. Roll out as stacked Markdown-only review slices: (1) shared persistence/common boundary, (2) backend convention cleanup, (3) core phase artifact contracts, (4) apply/archive/init mutation-heavy contracts, (5) optional registry/docs refresh. Each slice must target under 400 changed lines and preserve behavior.

## Open Questions

None.

## Security Applicability Routing

`security-applicability.md` classifies this change as `no-impact` with `securityImpact: false`, no taxonomy categories, no guidelines, no blocking unknowns, and no non-blocking risks. No `security-design.md` is required. Next route: `test-design`.
