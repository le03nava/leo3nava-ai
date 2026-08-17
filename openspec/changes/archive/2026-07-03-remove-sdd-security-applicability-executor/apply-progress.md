# Apply Progress: Remove SDD Security Applicability Executor

## Slice

| Field | Value |
| --- | --- |
| Current slice boundary | WU2 / PR2 stacked-to-main: Preserve legacy data compatibility, align review/source specs, and capture final static verification evidence for tasks 2.x, 3.x, and 4.x. |
| Previous slice boundary | WU1 / PR1 stacked-to-main: Remove launchable executor/skill and active routing authority. |
| Mode | Standard apply; strict TDD disabled. |
| Artifact store | OpenSpec |
| Runtime tooling | Unavailable per `openspec/config.yaml`; not reported as passed. |

## Completed Tasks

- [x] 1.1 Delete `agents/sdd/sdd-security-applicability.md`; verify repo-local executor absence with TD-002 and SDD-GOV-001 evidence.
- [x] 1.2 Delete `skills/sdd-security-applicability/SKILL.md` and remove the empty directory if present; verify no replacement launchable skill is added (TD-002).
- [x] 1.3 Update `agents/sdd/sdd-orchestrator.md` so `security-applicability` is legacy/archive data only and never an active DAG successor or launch mapping (TD-001, TD-003, TD-005, SDD-GOV-001).
- [x] 2.1 Update shared persistence/status/security/OpenSpec contracts for legacy/read-only applicability refs (TD-004, TD-006, SDD-GOV-002).
- [x] 2.2 Verify `sdd-security-design` and related shared contracts keep direct classification from proposal/specs/design and never consult the retired executor or artifact (TD-007, TD-008, SDD-GOV-001).
- [x] 2.3 Confirm `scripts/validate_security_applicability.ps1` remains archive-only and is not introduced as a new-change gate (TD-012, SDD-GOV-002).
- [x] 3.1 Update `skills/sdd-review/references/report-template.md` for optional legacy applicability evidence (TD-009, TD-010).
- [x] 3.2 Sync active source specs under `openspec/specs/**/spec.md` with the delta requirements without editing archives (TD-011, SDD-GOV-002).
- [x] 3.3 Document repo-local scope and stale `%USERPROFILE%` copy caveat without claiming global cleanup (TD-013).
- [x] 4.1 Capture targeted search/read-back evidence for deleted launch files, absent new `security-applicability.md`, mandatory `security-design.md`, and no active successor mapping (TD-001..TD-007, TD-009..TD-012).
- [x] 4.2 Inspect changed artifacts for secrets, credentials, production identifiers, sensitive runtime data, and copied environment values (TD-014, SDD-GOV-003).
- [x] 4.3 State runtime tests, coverage, lint, typecheck, format, and build tooling are unavailable rather than passed (TD-015).

## Pending Tasks

None. All implementation tasks are complete and persisted in `tasks.md`.

## Files Changed

| File | Action | What Was Done |
| --- | --- | --- |
| `agents/sdd/sdd-security-applicability.md` | Deleted in WU1 | Removed the repo-local launchable security applicability executor prompt. |
| `skills/sdd-security-applicability/SKILL.md` | Deleted in WU1 | Removed the repo-local launchable security applicability skill. |
| `skills/sdd-security-applicability/` | Removed in WU1 if empty | Empty skill directory removed from the working tree. |
| `agents/sdd/sdd-orchestrator.md` | Modified in WU1 | Reworded the naming contract so `sdd-security-applicability` is a retired archived-data token only and removed it from active phase context. |
| `skills/_shared/persistence-contract.md` | Confirmed/modified | Preserves `security-applicability` resolver rows only as legacy/read-only data and keeps new state active refs on `securityDesign` / `securityReviewReport`. |
| `skills/_shared/sdd-status-contract.md` | Modified | Removed the legacy token-to-agent launch mapping by changing `security-applicability` to no launch target and preserving `securityApplicability` only as legacy/read-only status data. |
| `skills/_shared/sdd-security-contract.md` | Confirmed | Defines `security-applicability.md` as legacy/read-only and mandatory `security-design.md` as new-change classification authority. |
| `skills/_shared/openspec-convention.md` | Modified | Replaced the phase-table entry with a legacy data row and kept OpenSpec path readability without launch authority. |
| `skills/sdd-security-design/SKILL.md` | Confirmed | Keeps classification direct from proposal/specs/design; legacy applicability is allowed only when explicitly identified as old/archive compatibility context. |
| `scripts/validate_security_applicability.ps1` | Confirmed | Header states the validator is legacy/archive-only and MUST NOT block new-change routing or phase success. |
| `skills/sdd-review/references/report-template.md` | Modified | Marks legacy security applicability as optional old/archive evidence and keeps `security-design.md` / `review-security-report.md` as new-change authorities. |
| `README.md` | Modified | Documents repo-local removal scope and the stale `%USERPROFILE%` destination-copy caveat. |
| `openspec/specs/sdd-security-applicability-workflow/spec.md` | Modified | Synced source requirements so no executor/skill is offered and legacy artifact resolution remains data-only. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modified | Synced source requirements for legacy token/data refs that never normalize into a runnable successor. |
| `openspec/specs/sdd-security-design-workflow/spec.md` | Modified | Synced source requirements for direct security-design classification and missing applicability not proving no-impact. |
| `openspec/specs/sdd-review-workflow/spec.md` | Modified | Synced source requirements for optional legacy applicability evidence and mandatory security design/review authority. |
| `openspec/changes/remove-sdd-security-applicability-executor/tasks.md` | Modified | Marked all tasks 1.1 through 4.3 complete. |
| `openspec/changes/remove-sdd-security-applicability-executor/apply-progress.md` | Updated | Merged WU1 and WU2 evidence without dropping prior completed work. |

## Static Evidence

| Case / Control | Evidence |
| --- | --- |
| TD-001 / TD-005 / SDD-GOV-001 | `agents/sdd/sdd-orchestrator.md` naming contract states `sdd-security-applicability` is retired old/archive data only and must not be launched, mapped to an agent/skill, or emitted as a new-change successor. Targeted grep of active `skills/` and `agents/` for token-to-agent launch mappings returned no matches. |
| TD-002 / SDD-GOV-001 | `glob` read-back found no `agents/sdd/sdd-security-applicability.md` and no `skills/sdd-security-applicability/SKILL.md`. No replacement launchable executor or skill was added. |
| TD-003 | `glob` read-back found no `openspec/changes/remove-sdd-security-applicability-executor/security-applicability.md`; `openspec/changes/remove-sdd-security-applicability-executor/security-design.md` remains present and mandatory. |
| TD-004 / TD-006 / SDD-GOV-002 | `skills/_shared/persistence-contract.md`, `skills/_shared/sdd-status-contract.md`, `skills/_shared/sdd-security-contract.md`, and `skills/_shared/openspec-convention.md` preserve applicability refs as legacy/read-only data only and keep active dependencies on mandatory security design/security review refs. |
| TD-007 / TD-008 / SDD-GOV-001 | `skills/sdd-security-design/SKILL.md` requires proposal/spec/design inputs, treats legacy applicability only as explicitly old/archive compatibility context, says no-impact still creates `security-design.md`, and forbids consuming legacy applicability matrix fields as mandatory new-change inputs. |
| TD-009 / TD-010 | `skills/sdd-review/references/report-template.md` now lists `security-design` in required new-change inputs and labels `security-applicability` as optional legacy old/archive compatibility evidence only; security design and security review remain new-change authorities. |
| TD-011 | WU2 did not intentionally edit `openspec/changes/archive/**`; archive paths in `git status --short` are pre-existing approved dirty workspace context from the prior archived change. Active source spec edits were limited to `openspec/specs/**/spec.md`. |
| TD-012 / SDD-GOV-002 | `scripts/validate_security_applicability.ps1` header says it is retained only for explicit legacy/archive compatibility checks and MUST NOT block new-change routing or phase success. Targeted grep found no new-change gate invocation. |
| TD-013 | `README.md` states removal is repo-local and stale copies under `%USERPROFILE%` destinations may remain until a separate sync cleanup policy handles destination deletion. |
| TD-014 / SDD-GOV-003 | Targeted sensitive-pattern inspection of changed Markdown artifacts found only workflow/catalog terms and negative examples such as “secrets”, “tokens”, and “credentials”; no raw secret values, keys, production identifiers, environment values, or sensitive runtime data were introduced. |
| TD-015 | `openspec/config.yaml` has `testing.test_runner.available: false`, empty test/build commands, and no coverage/lint/typecheck/format commands. No runtime test, lint, typecheck, format, build, or coverage command was run or reported as passing. |

## Tooling Notes

- Runtime tests: unavailable per `openspec/config.yaml` (`testing.test_runner.available: false`, empty command).
- Coverage: unavailable per `openspec/config.yaml`.
- Linter: unavailable per `openspec/config.yaml`.
- Type checker: unavailable per `openspec/config.yaml`.
- Formatter: unavailable per `openspec/config.yaml`.
- Build command: unavailable per `openspec/config.yaml` (`rules.verify.build_command: ""`).
- `rg` command was unavailable in this PowerShell environment; targeted repository searches used available file/search tools and `git` read-only inspection instead.

## Deviations

None. WU2 stayed within the assigned boundary, preserved WU1 evidence, did not edit global `%USERPROFILE%` files, and did not intentionally edit archived change folders.

## Workload / PR Boundary

- Mode: stacked PR slice.
- Current work unit: WU2 / PR2 stacked-to-main.
- Boundary: starts after WU1 deletion/routing changes; ends with shared contracts, review template, active source specs, tasks, and apply-progress synchronized.
- Estimated review budget impact: WU2 is documentation/contract/spec synchronization and final evidence; no runtime code or generated artifacts beyond the active OpenSpec apply artifacts.

## Status

12/12 tasks complete. Ready for `sdd-review`.
