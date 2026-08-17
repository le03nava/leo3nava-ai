# Design: Add SDD Review Phase

## Technical Approach

Introduce `sdd-review` as a first-class SDD executor and skill between apply and verify. The change updates phase routing, artifact resolution, shared persistence/status contracts, and downstream verify/archive gates so review evidence is durable but remains separate from verification and security-design authority.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Mandatory DAG slot | Route `tasks -> apply -> review -> verify -> archive`. | Keep `apply -> verify`; fold review into verify. | A separate gate makes the 96-control checklist auditable and prevents verify from owning code-review semantics. |
| Artifact identity | Create `review-report.md` in OpenSpec and use `sdd/{change-name}/review` for Engram/hybrid references. | Reuse `verify-report.md`; store only inline. | A distinct artifact keeps review evidence recoverable and compatible with current artifact-store contracts. |
| Control catalog shape | Maintain a stable 96-control review catalog with stable Item IDs/categories and render report rows with exactly the requested columns. | Add Category as an extra matrix column; rely on prose ordering. | Specs require exact report columns; categories can live in the catalog or Item ID prefix without changing the matrix. |
| Security boundary | Review rows may cite security guideline IDs in `Standard`, but applicability/design remain authoritative. | Let review redefine or override security controls. | Preserves the existing security contract and avoids conflicting sources of security truth. |

## Data Flow

```text
sdd-apply success
  -> sdd-review reads proposal/specs/security-applicability/design/test-design/tasks + changed-file context
  -> review-report.md
       | no blocking failures -> sdd-verify consumes summary/evidence
       | blocking failures    -> sdd-apply remediates
       | missing/unsafe input -> resolve-blockers
  -> sdd-archive requires non-blocking review-report.md + passing verify-report.md
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `agents/sdd/sdd-review.md` | Create | Executor prompt matching existing SDD agent pattern and loading `skills/sdd-review/SKILL.md`. |
| `skills/sdd-review/SKILL.md` | Create | Review phase contract: required inputs, 96-control matrix, verdict/routing, persistence, and return envelope. |
| `agents/sdd/sdd-orchestrator.md` | Modify | Add review to commands, flowchart, dependency graph, routing gates, state persistence expectations, and archive readiness. |
| `skills/_shared/sdd-status-contract.md` | Modify | Add native token `review`, `reviewReport` refs/paths/artifact state/dependency, and routing normalization. |
| `skills/_shared/persistence-contract.md` | Modify | Add resolver entry for OpenSpec `review-report.md` and Engram `sdd/{change-name}/review`. |
| `skills/_shared/openspec-convention.md` | Modify | Add `review-report.md` to directory structure, path table, read list, and archived artifact expectations. |
| `skills/_shared/security-guideline-catalog.md` | Modify | Add review-safe cross-reference/evidence guidance without duplicating guideline text. |
| `skills/sdd-apply/SKILL.md` | Modify | Change success routing from `verify` to `review` when all tasks are complete. |
| `skills/sdd-verify/SKILL.md` | Modify | Require non-blocking review evidence and consume the report without duplicating the full matrix. |
| `skills/sdd-archive/SKILL.md` | Modify | Require non-blocking `review-report.md` plus passing `verify-report.md`. |
| `README.md` | Modify | Document the new phase and order. |
| `scripts/sdd_init_agents.ps1`, `scripts/sdd_init_skills.ps1` | No logic change | Existing recursive sync copies the new agent/skill automatically. |

## Interfaces / Contracts

`review-report.md` MUST include: verdict, blocking summary, evidence summary, next recommendation, and a matrix with exactly these columns: Item, Artifact/Deliverable, Requirement, Reviewer, Standard, Severity, Complies, Affected Requirement, Evidence Location, Observations/Comments.

Control rules: all 96 controls appear once; Item IDs are stable and category-bearing (for example prefix or catalog mapping); `Complies` is `Yes`, `No`, or `N/A`; every `N/A` requires an evidence location and comment proving irrelevance.

Routing: apply success -> `review`; review non-blocking -> `verify`; review blocking -> `apply`; missing artifacts, unknown changed files, unsafe workspace, or persistence failure -> `resolve-blockers`; archive requires review + verify.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static contract inspection | Tokens, artifact refs, routing tables, exact matrix columns, archive gates. | Manual/OpenSpec diff inspection; no runtime runner is configured. |
| OpenSpec validation | Delta specs align with design and main contract conventions. | Read specs and affected Markdown contracts for requirement coverage. |
| Sync compatibility | New agent/skill copied by existing scripts. | Inspect recursive copy behavior; no script logic change expected. |

## Migration / Rollout

No data migration required. Roll out as stacked-to-main slices: shared contracts first, review executor/skill second, phase integration/docs last.

## Open Questions

None.

## Security Applicability Routing

`security-applicability.md` classifies this as `no-impact` with `securityImpact: false`. No `security-design.md` is required. The design preserves the stated boundary: review may cite guideline IDs, but security applicability/design remain authoritative. Next route: `test-design`.
