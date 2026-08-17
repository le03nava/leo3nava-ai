# Design: Remove SDD Security Applicability Executor

## Technical Approach

Remove the repo-local `sdd-security-applicability` executor surface while preserving old `security-applicability.md` artifacts as legacy, read-only evidence. The active workflow remains `spec -> design -> security-design -> test-design`; security classification for new changes is owned by mandatory `security-design.md`, not by a separate applicability phase.

This change is documentation/instruction-contract work only. It updates source prompts, shared contracts, and OpenSpec specs so launchers, status readers, reviewers, and archive logic agree that applicability is data compatibility, not runnable workflow.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Remove launch sources | Delete `agents/sdd/sdd-security-applicability.md` and `skills/sdd-security-applicability/SKILL.md`. | Keep them as deprecated redirectors. | A present agent/skill is still discoverable and may be synced as active. Deletion is the only unambiguous repo-local signal that the executor no longer exists. |
| Keep legacy artifact references | Preserve `securityApplicability` state/status fields and OpenSpec path references as `legacy/read-only`. | Remove every mention of applicability. | Archives and old state may still contain `security-applicability.md`; readers must not lose audit evidence or fail old-folder inspection. |
| Move classification authority | Require new-change classification, no-impact proof, controls, and evidence ownership in `security-design.md`. | Recreate a lightweight applicability classifier. | The approved DAG makes `security-design` mandatory for all new changes, including no-impact changes, so a separate classifier would reintroduce the conflict. |
| Preserve archive-only validation | Keep `scripts/validate_security_applicability.ps1` as a legacy/archive validator. | Delete or rename the validator now. | The proposal explicitly keeps archive readability. The script already states it must not block new-change routing and only validates historical Markdown/YAML contracts. |
| Do not edit archived folders | Leave `openspec/changes/archive/**` untouched. | Rewrite archived copies to remove old wording. | Archives are audit trails. Active source contracts should clarify behavior without mutating historical records. |

## Data Flow

```text
New change:
spec -> design -> security-design -> test-design -> tasks
                 └─ owns classification, matrix, controls, N/A proof

Legacy/archive read:
archived security-applicability.md -> compatibility reader/status display only
                                  └─ no runnable successor mapping
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `agents/sdd/sdd-security-applicability.md` | Delete | Remove repo-local launchable executor prompt. |
| `skills/sdd-security-applicability/SKILL.md` | Delete | Remove repo-local launchable executor skill and empty directory if left empty. |
| `agents/sdd/sdd-orchestrator.md` | Modify | Remove active applicability launch mapping and keep only legacy/archive wording where needed. |
| `skills/_shared/persistence-contract.md` | Modify | Mark security applicability resolver/state refs as legacy/read-only and exclude it from active phase tokens. |
| `skills/_shared/sdd-status-contract.md` | Modify | Ensure status/dependency rules never emit `security-applicability` as a new-change successor. |
| `skills/_shared/sdd-security-contract.md` | Modify | Define `security-applicability.md` as legacy/read-only and `security-design.md` as mandatory classification authority. |
| `skills/_shared/openspec-convention.md` | Modify | Preserve legacy path readability while documenting new-change absence. |
| `skills/sdd-review/references/report-template.md` | Modify | Treat applicability as optional legacy evidence; keep `security-design.md` and security review as authorities. |
| `openspec/specs/**/spec.md` | Modify | Sync final capability requirements during archive from this change's delta specs. |
| `scripts/validate_security_applicability.ps1` | Preserve | Keep as archive-only validator; do not use as a new-change gate. |
| `%USERPROFILE%/.config/opencode/**` | Out of scope | Stale global copies may remain until external sync/deletion is handled separately. |

## Interfaces / Contracts

- New-change routing tokens MUST exclude active `security-applicability` successors.
- `artifactRefs.securityApplicability` and `security-applicability.md` paths MAY remain only as legacy data refs.
- New changes MUST NOT produce `security-applicability.md`.
- New changes MUST produce `security-design.md`, including `classification`, `securityImpact`, taxonomy evaluation, controls or N/A rationale, validation metadata, and `nextRecommended: test-design`.
- Archive readers MAY validate old applicability artifacts with `scripts/validate_security_applicability.ps1`; that validator MUST NOT block new-change routing.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static contract review | Deleted launch files and active contract wording. | Inspect repository paths and grep for active launch references. |
| OpenSpec validation | Delta specs align with proposal and design. | Manual review of modified requirement scenarios; no runtime runner exists. |
| Archive compatibility | Archived folders and validator remain intact. | Confirm `openspec/changes/archive/**` is unchanged and validator text remains archive-only. |

No package manifest, build file, linter, or test runner is configured in `openspec/config.yaml`; verification must report unavailable runtime tests explicitly and rely on read-back, diff inspection, and targeted text searches.

## Migration / Rollout

No data migration required. Rollout is a repo-local contract update plus source deletion. Existing archives remain readable. Global opencode copies under `%USERPROFILE%` are outside this repository scope and may remain stale until a separate sync cleanup policy is approved.

## Open Questions

None.

## Security Applicability Routing

This design is security-impacting at the workflow-contract level because it changes where security classification is owned. Next route: `security-design`. The next phase must classify this change in mandatory `security-design.md`; no `security-applicability.md` should be created.
