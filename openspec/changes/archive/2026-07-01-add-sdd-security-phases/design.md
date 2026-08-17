# Design: Add SDD Security Phases

## Technical Approach

Insert a mandatory security triage edge into the SDD DAG and make security design a conditional successor of technical design:

```text
proposal -> spec -> security-applicability -> design -> security-design? -> test-design -> tasks -> apply -> verify -> archive
```

`sdd-security-applicability` is an always-run planning phase. It reads proposal/specs plus the in-repo guideline catalog, writes `security-applicability.md`, and routes either to `design` for no-impact or still to `design` with security mappings for impacting changes. After technical design, the orchestrator runs `sdd-security-design` only when applicability says `securityImpact: true`; otherwise missing `security-design.md` is not a blocker. `sdd-test-design` consumes security controls/evidence when present.

## Architecture Decisions

| Decision | Choice | Rationale |
| --- | --- | --- |
| Applicability before design | Always run after specs, before technical design | Security impact can change architecture decisions; late discovery would make design/tasks unreliable. |
| Security design after technical design | Conditional successor of `sdd-design` | Security controls need concrete architecture context, but no-impact changes avoid unnecessary ceremony. |
| Catalog location | Create `skills/_shared/security-guideline-catalog.md` plus `skills/_shared/sdd-security-contract.md` | Shared files keep prompts compact, auditable, and reusable across applicability, design, verify, and archive. |
| Evidence model | Mandatory-by-guideline with approved exceptions | Archive can enforce evidence deterministically without blocking non-mandatory or no-impact work. |

## Routing, State, and Persistence

Update bounded tokens and state/status fields to include `security-applicability` and `security-design`, with phase-agent aliases `sdd-security-applicability` and `sdd-security-design`. Persist artifact refs in camelCase state/status fields: `securityApplicability` and `securityDesign`. Engram keys are `sdd/{change}/security-applicability` and `sdd/{change}/security-design`; OpenSpec paths are `openspec/changes/{change}/security-applicability.md` and `openspec/changes/{change}/security-design.md`.

State persistence remains orchestrator-owned: phase success must be gatechecked, state updated, read back, then routed. `tasks` readiness requires `test-design`; `test-design` readiness requires `security-design` only when applicability is impacting.

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `agents/sdd/sdd-security-applicability.md` | Create | Executor prompt for applicability phase. |
| `agents/sdd/sdd-security-design.md` | Create | Executor prompt for conditional security design. |
| `skills/sdd-security-applicability/SKILL.md` | Create | Reads proposal/spec/catalog; writes applicability artifact and routing recommendation. |
| `skills/sdd-security-design/SKILL.md` | Create | Maps applicable guidelines to controls, evidence, risks, and exceptions. |
| `skills/_shared/security-guideline-catalog.md` | Create | In-repo snapshot preserving source metadata, ids, source text/summaries, taxonomy, mandatory evidence. |
| `skills/_shared/sdd-security-contract.md` | Create | Shared artifact schemas, classification values, exception fields, and evidence statuses. |
| `agents/sdd/sdd-orchestrator.md` | Modify | DAG, status routing, dependency gates, launch protocol, and interactive/auto flow. |
| `skills/_shared/sdd-status-contract.md` | Modify | Tokens, status schema, dependency states, artifact refs/paths. |
| `skills/_shared/persistence-contract.md` | Modify | Resolver table, minimum state schema, Engram/OpenSpec refs. |
| `skills/_shared/openspec-convention.md` | Modify | New artifact paths in change folders. |
| `skills/sdd-design/SKILL.md` | Modify | Require existing applicability artifact before design and route to security-design/test-design appropriately. |
| `skills/sdd-test-design/SKILL.md` | Modify | Read `security-design.md` when required and plan checks/evidence per controls. |
| `skills/sdd-tasks/SKILL.md`, `skills/sdd-apply/SKILL.md`, `skills/sdd-verify/SKILL.md`, `skills/sdd-archive/SKILL.md` | Modify | Carry security evidence through tasks, implementation evidence, verification matrices, and archive blockers. |
| `README.md` | Modify | List the two new SDD phases. |

## Interfaces / Contracts

`security-applicability.md` MUST include: `classification` (`security-impacting` or `no-impact`), evidence summary, touched taxonomy categories, guideline IDs, design-changing unknowns, non-blocking risks, and `next_recommended`.

`security-design.md` MUST include: applicable guideline ID, category, mandatory flag, required control, expected evidence type, owner phase (`design`, `test-design`, `apply`, `verify`, `archive`), status, residual risk, and optional exception.

Catalog entries MUST preserve: source snapshot metadata, stable ID, category, source text or faithful summary, mandatory-when-applicable flag, expected evidence, and audit notes. Exception evidence MUST include approver, guideline ID, accepted-risk rationale, and mitigation or follow-up.

## Testing Strategy

No executable runner exists. Verification is static/manual: read generated/modified files, validate YAML frontmatter, token mappings, DAG text, state schema fields, artifact path consistency, and OpenSpec deltas. Add mandatory `test-design.md` cases for security-impacting and no-impact routing, archive blocking, and exception completeness.

## Migration / Rollout

No data migration required. Roll out by adding shared contracts/catalog first, then new agents/skills, then orchestrator/status/persistence/archive consumers. Existing active changes without `security-applicability.md` should be treated as legacy and continue unless restarted under the new DAG; new changes must follow the security phases.

## Open Questions

None
