# Apply Evidence WU1: Shared Contract and Catalog

## Scope

WU1 updates only the shared security contract and catalog boundary for `slim-secure-design-artifacts`.

Completed tasks:

- 1.1 Update `skills/_shared/sdd-security-contract.md` to state catalog/design/test-design/review-security/verify/archive ownership, blocker routing, safe evidence, N/A, warning, and exception rules.
- 1.2 Update `skills/_shared/security-guideline-catalog.md` to declare the authoritative 155 Source ID inventory, snapshot metadata, expanded ranges, compact mappings, and expected count.
- 1.3 Record WU1 static evidence for TD-003, TD-005, TD-011, and compact security coverage.

Out of scope for this slice: phase skill changes, adapter prompt changes, source spec sync, review-security ownership edits, verify/archive edits, PR creation, and runtime tooling changes.

## Static / Manual Evidence

| Evidence ID | Test-design case / control | Evidence | Result |
| --- | --- | --- | --- |
| WU1-TD-003 | TD-003 catalog authority | `skills/_shared/security-guideline-catalog.md#snapshot-metadata` declares snapshot ID `security-guidelines-initial-user-snapshot-2026-06-30`, catalog version `1`, taxonomy version `1`, inventory authority for 155 concrete Source IDs, and expected count `155`. | Pass |
| WU1-TD-003 | TD-003 expanded ranges | `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` declares the catalog as authoritative, lists sections 1-15, and records `Inventory total: 155 expanded Source IDs`. | Pass |
| WU1-TD-005 | TD-005 compact mapping vocabulary | `skills/_shared/security-guideline-catalog.md#inventory-ownership-contract` and `skills/_shared/sdd-security-contract.md#artifact-ownership-boundary` restrict compact mappings to the eight existing `SEC-*` IDs and route missing/unknown mappings to blockers. | Pass |
| WU1-TD-011 | TD-011 grouped planning source | `skills/_shared/sdd-security-contract.md#designmdsecure-development-design-schema` requires slim design source-row coverage by catalog authority, expected count, grouped coverage references, compact mappings, owner phases, lifecycle state, exact-once downstream rule, safe-evidence policy, N/A policy, and exception policy. | Pass |
| WU1-BOUNDARY | Slim artifact boundary | `skills/_shared/sdd-security-contract.md#artifact-ownership-boundary` defines catalog, design, test-design, apply, review-security, verify, and archive ownership and prohibits design/test-design/apply/verify/archive from owning the full 155-row matrix. | Pass |
| WU1-ROUTING | SEC-ACCESS-001 blocker routing | `skills/_shared/sdd-security-contract.md#source-row-routing-and-persistence-semantics` preserves routes for missing/duplicate/unknown IDs, malformed rows, unknown mappings, unsafe evidence, unsupported N/A, implementation gaps, warnings-only results, and archive handoff. | Pass |
| WU1-SAFE | SEC-DATA/SECRET/LOG safe evidence | `skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls` and catalog source-row validation rules require paths, anchors, summaries, command summaries, sanitized observations, or redacted placeholders only. | Pass |

## Security Control Evidence

| Guideline | WU1 evidence |
| --- | --- |
| `SEC-AUTH-001` | Catalog remains authoritative for authentication Source IDs and compact mappings; review-security keeps exact-once validation ownership. |
| `SEC-SESS-001` | Session Source IDs remain catalog-owned and are consumed through grouped references and downstream security review. |
| `SEC-DATA-001` | Shared contract preserves safe-evidence rules and slim artifacts avoid raw sensitive data. |
| `SEC-SECRET-001` | Shared contract and catalog prohibit unsafe evidence and preserve blocker routing for secret/config evidence gaps. |
| `SEC-ACCESS-001` | Source-row routing remains denial-by-default for missing, malformed, duplicate, unknown, unmapped, unsafe, or unsupported N/A evidence. |
| `SEC-FILE-001` | File-based artifacts carry references and summaries; review-security owns exhaustive row materialization. |
| `SEC-DB-001` | Database Source IDs remain mapped in catalog inventory and traceable through compact mappings. |
| `SEC-LOG-001` | Logging/error evidence remains review-safe and warning evidence remains visible downstream. |

## Tooling Evidence

`openspec/config.yaml#testing` reports no runtime test runner, build command, coverage command, linter, type checker, or formatter. Those tools are unavailable and are not treated as passing evidence. WU1 validation used static/manual artifact read-back only.

## Changed Files Read-Back

- `skills/_shared/sdd-security-contract.md` — read back after edit; contains artifact ownership boundary, slim design source-row coverage rules, routing, safe evidence, N/A, warning, and exception semantics.
- `skills/_shared/security-guideline-catalog.md` — read back after edit; contains inventory authority, expected count, ownership contract, expanded section inventory, compact mapping vocabulary, and canonical section count total.
- `openspec/changes/slim-secure-design-artifacts/tasks.md` — read back after task checkbox update.
- `openspec/changes/slim-secure-design-artifacts/apply-progress.md` — read back after apply progress persistence.

## Deviations

None. WU1 follows the assigned shared contract/catalog boundary and does not implement WU2-WU4.
