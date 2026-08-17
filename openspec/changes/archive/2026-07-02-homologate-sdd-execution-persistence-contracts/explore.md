# Exploration: Homologate SDD Execution and Persistence Contracts

## Current State

The repository is an AI agent/skill distribution. SDD behavior is defined mostly in Markdown skill contracts under `skills/sdd-*/SKILL.md` and shared contracts under `skills/_shared/*.md`; `openspec/config.yaml` confirms OpenSpec is the active artifact store for this change.

The core issue is real: persistence behavior is already centralized in `skills/_shared/persistence-contract.md`, but many phase skills still repeat mode-specific behavior in local `Execution and Persistence Contract` sections. `skills/_shared/sdd-phase-common.md` also duplicates retrieval and persistence rules in Sections B/C, so there are currently two shared places plus per-phase copies.

## Current Duplication / Overlap Map

| Area | Current owner(s) | Overlap observed |
| --- | --- | --- |
| Mode semantics (`engram`, `openspec`, `hybrid`, `none`) | `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-phase-common.md`, most `skills/sdd-*/SKILL.md` | Same mode behavior repeated with slightly different wording. |
| Engram naming and `capture_prompt: false` | `persistence-contract.md`, `engram-convention.md`, `sdd-phase-common.md`, per-phase persistence steps | Naming is mostly consistent, but authority is unclear. |
| OpenSpec paths | `openspec-convention.md`, `persistence-contract.md` resolver, per-phase steps | Paths are duplicated across resolver tables and phase skills. |
| Hybrid conflict/read-write policy | `persistence-contract.md`, selected per-phase sections | Per-phase wording repeats the policy but cannot stay complete as policy evolves. |
| Return envelope and routing tokens | `sdd-phase-common.md`, `sdd-status-contract.md`, per-phase output sections | Shared routing exists, while per-phase next-token rules are legitimately phase-specific. |
| Required phase dependencies | Per-phase skills, `persistence-contract.md` resolver, `sdd-status-contract.md` dependencies | Dependency readiness belongs partly in status, while phase-required inputs should stay in each phase artifact contract. |

Per-phase findings:

| Skill | Current local persistence shape | Keep phase-specific? | Move to shared? |
| --- | --- | --- | --- |
| `sdd-explore` | Full `Execution and Persistence Contract`; repeats all modes plus standalone/named-change nuances. | Named-change vs standalone, produced `explore`, next `propose`. | Mode behavior. |
| `sdd-propose` | Full section with all modes, optional explore/project context. | Inputs, proposal path/key, capabilities contract, next `spec`. | Mode rules and hybrid policy. |
| `sdd-spec` | Full section with all modes and multi-domain spec behavior. | Proposal required, domain spec fan-out, Engram concatenation, next `design`. | Generic mode restrictions. |
| `sdd-security-applicability` | Compact `## Persistence` table. | Proposal/spec inputs, produced `security-applicability`, next `design`. | Shared mode behavior can be referenced instead of restated. |
| `sdd-design` | Full section with all modes and security applicability dependency. | Required inputs, produced `design`, conditional next token. | Generic mode restrictions. |
| `sdd-security-design` | Compact `## Persistence` table. | Conditional artifact creation only when impacting, next `test-design`. | Shared mode behavior can be referenced. |
| `sdd-test-design` | Full section with all modes and dependencies. | Required inputs including conditional security design, produced `test-design`, next `tasks`. | Generic mode restrictions. |
| `sdd-tasks` | Full section with all modes and workload guard inputs. | Required planning inputs, produced `tasks`, next `apply`, workload forecast handoff. | Generic mode restrictions. |
| `sdd-apply` | Full section with implementation-specific persistence updates. | Task checkbox mutation, apply-progress, workspace guard interaction, next `apply`/`verify`. | Generic mode restrictions; keep mutation semantics local. |
| `sdd-verify` | No titled Execution section, but Hard Rules and steps define persistence. | Verification inputs, produced `verify-report`, verdict routing. | Generic persistence call/path rules. |
| `sdd-archive` | Full section with archive-specific mode behavior. | Spec merge, folder move, archive report, Engram lineage, next `none`. | Generic mode restrictions; keep archive semantics local. |
| `sdd-init` | Hard Rules define bootstrap persistence. | Project context/testing capabilities/registry bootstrap outputs. | Generic mode semantics, where applicable. |
| `sdd-new`, `sdd-onboard` | Orchestrator/coordinator workflow skills. | Coordination boundaries; they should not own phase persistence details. | Any mode semantics repeated in prompts should reference shared contracts. |

## Recommended Target Contract Boundaries

### 1. `skills/_shared/persistence-contract.md` becomes authoritative

Owns:
- artifact store mode resolution and mode roles;
- mode-specific read/write behavior;
- Engram/OpenSpec/hybrid/none common rules;
- hybrid conflict policy;
- artifact reference resolver;
- state persistence and read-back verification;
- sub-agent persistence expectations, including response ordering when persistence happens before final text.

### 2. `skills/_shared/sdd-phase-common.md` becomes execution/envelope glue only

Keep:
- executor boundary summary;
- skill loading protocol;
- Section D return envelope;
- routing token convention;
- artifact naming convention if it remains a routing/naming reminder;
- review workload guard.

Move or replace with references:
- current Sections B/C retrieval and persistence details should point to `persistence-contract.md` instead of restating `mem_search`, `mem_get_observation`, `mem_save`, OpenSpec, hybrid, and none behavior.

### 3. Backend convention files stay backend-specific

- `openspec-convention.md`: filesystem layout, concrete paths, writing rules, archive layout, delta spec mechanics.
- `engram-convention.md`: Engram naming and operational reference. Consider updating its note that critical calls are inlined directly in each skill; after this change, critical behavior should be centralized through `persistence-contract.md`, not copied into every phase skill.

### 4. Per-phase skills keep a compact `Phase Artifact Contract`

Replace each local `Execution and Persistence Contract` / `Persistence` section with a compact contract like:

| Field | Meaning |
| --- | --- |
| `shared_contract` | `skills/_shared/persistence-contract.md` plus `skills/_shared/sdd-phase-common.md` envelope. |
| `required_inputs` | Phase-specific artifacts or launch fields. |
| `produces` | Artifact type and whether it is mandatory/optional/conditional. |
| `engram_key` | `sdd/{change-name}/{artifact}` when applicable. |
| `openspec_path` | Exact file/folder path from `openspec-convention.md`. |
| `mutation` | Only for apply/archive/init or any phase that updates existing artifacts. |
| `next_recommended` | Phase-specific success/blocked routing. |

## Shared Artifact Matrix Decision

Do **not** create a new shared artifact matrix file for the first migration. Extend `skills/_shared/persistence-contract.md` instead.

Rationale:
- It already contains `## Artifact Reference Resolver`, which is effectively the artifact matrix.
- A separate matrix would create a third shared location and continue the same authority problem under a new filename.
- Reviewers can validate one authoritative persistence surface more easily.

Recommended refinement: expand the existing resolver into an `Artifact Matrix and Resolver` with columns for artifact, produced by, required by/readers, Engram key, OpenSpec path, `none` behavior, and special notes. If that table becomes too large later, split only after the migration proves the boundary and make `persistence-contract.md` the explicit index pointing to the split file.

## Risks

- **Behavior drift during partial migration**: if only some phase skills are updated, old copied rules may conflict with the shared contract.
- **Over-centralization**: phase-specific semantics such as `sdd-apply` checkbox mutation or `sdd-archive` folder movement must not be flattened into generic persistence rules.
- **Status/persistence coupling**: `sdd-status-contract.md` owns readiness and dependency states; persistence contract should reference status behavior but not duplicate all readiness logic.
- **Adapter sync risk**: this repo syncs skills into `%USERPROFILE%\.config\opencode\skills`; implementation must account for source skill files and adapter-installed copies without treating the installed copy as the source of truth.
- **Review size**: touching all phase skills plus shared contracts could exceed the 400-line review budget unless split into reviewable work units.

## Migration Strategy

1. Update shared contracts first:
   - make `persistence-contract.md` explicitly authoritative;
   - reduce `sdd-phase-common.md` Sections B/C to references or very thin pointers;
   - align `engram-convention.md` wording so it is reference documentation, not a competing authority.
2. Update phase skills in slices:
   - planning phases: `sdd-explore`, `sdd-propose`, `sdd-spec`, `sdd-security-applicability`, `sdd-design`, `sdd-security-design`, `sdd-test-design`, `sdd-tasks`;
   - execution/closure phases: `sdd-apply`, `sdd-verify`, `sdd-archive`;
   - bootstrap/coordinator skills: `sdd-init`, `sdd-new`, `sdd-onboard` as needed.
3. For each phase, replace repeated mode prose with a `Phase Artifact Contract`; keep validation, decision gates, artifact format, and routing rules local.
4. Verify with repository-level text checks:
   - only shared contracts should define mode semantics in detail;
   - phase skills should not contain repeated phrases like `Follow BOTH conventions`, `Never create or modify SDD/OpenSpec files`, or `Never force openspec/ creation` except where phase-specific and justified.
5. Sync/adapter follow-up should be handled by the repo's existing skill sync workflow, not by editing installed user copies as the source.

## Candidate Change Scope

In scope:
- Normalize SDD skill documentation contracts only.
- Make `persistence-contract.md` the authoritative persistence contract.
- Simplify `sdd-phase-common.md` to execution, loading, envelope, routing, and review guard responsibilities.
- Replace per-phase persistence sections with compact phase artifact contracts.
- Preserve all current artifact keys, OpenSpec paths, next tokens, dependency requirements, and security/test-design DAG rules.
- Update any directly related reference wording in `engram-convention.md` if it contradicts the new authority boundary.

Non-goals:
- No runtime implementation changes outside skill/shared Markdown contracts.
- No changes to OpenSpec generated artifacts beyond this change's SDD files.
- No redesign of the SDD DAG, security workflow, test-design workflow, review workload guard, or native status schema.
- No migration from OpenSpec to Engram or change in default artifact store behavior.
- No editing of installed adapter copies as source-of-truth implementation; source changes belong in repo `skills/`.

## Ready for Proposal

Yes. The proposal should define the contract boundary refactor, require compatibility preservation for artifact keys/paths/routing tokens, and plan implementation as chained reviewable slices because the likely diff touches many Markdown contracts.
