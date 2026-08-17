# Exploration: remove-sdd-security-applicability-executor

## Current State

The repository has already moved new-change security classification into mandatory `sdd-security-design`. Active workflow documentation in `README.md`, `agents/sdd/sdd-orchestrator.md`, `skills/_shared/sdd-status-contract.md`, `skills/_shared/sdd-security-contract.md`, `skills/_shared/openspec-convention.md`, and `openspec/specs/*` says `security-applicability.md` is legacy/read-only and MUST NOT gate new-change routing.

However, two runnable repo-local assets still exist:

- `agents/sdd/sdd-security-applicability.md` — an executor prompt that still instructs a launched executor to read `%USERPROFILE%/.config/opencode/skills/sdd-security-applicability/SKILL.md`.
- `skills/sdd-security-applicability/SKILL.md` — a deprecated executor skill that returns redirects or legacy summaries and includes the old artifact template.

Those files no longer provide required new-change behavior. Their continued presence creates an attractive nuisance: status/token contracts say the phase is legacy-only, but sync scripts still copy it as a usable agent/skill.

## Findings

### Can the executor and skill be deleted safely?

Yes, with coordinated contract cleanup. Deleting these repo files is safe for the active new-change DAG because current source contracts already route `spec -> design -> security-design -> test-design` and classify in `security-design.md`.

Deletion is not safe as a standalone file removal unless references that imply a runnable executor/skill are updated. The highest-risk references are:

- `agents/sdd/sdd-orchestrator.md` — maps `security-applicability` to `sdd-security-applicability` in compatibility text and lists it in required context.
- `skills/_shared/sdd-status-contract.md` — currently includes a mapping row from `security-applicability` to `sdd-security-applicability`.
- `skills/_shared/openspec-convention.md` and `skills/_shared/persistence-contract.md` — mention the legacy artifact path/key; these should remain data references but avoid naming a runnable skill as the owner.
- `openspec/specs/sdd-security-applicability-workflow/spec.md` — still has a scenario that says `sdd-security-applicability` reads/writes through shared persistence; this should become data-only compatibility wording.
- `skills/sdd-review/references/report-template.md` — input list still includes `security-applicability`; should mark it legacy/optional or remove from the default new-change input set.

Archived OpenSpec change folders contain historical mentions of the old phase and should remain unchanged as audit trail.

### Which docs/scripts/contracts must change?

The proposal should target active source contracts only:

| File | Required change |
| --- | --- |
| `agents/sdd/sdd-security-applicability.md` | Delete. |
| `skills/sdd-security-applicability/SKILL.md` | Delete, and remove the now-empty directory if applicable. |
| `agents/sdd/sdd-orchestrator.md` | Remove launch mapping/context for `sdd-security-applicability`; keep legacy `securityApplicability` artifact refs as data-only compatibility. |
| `skills/_shared/sdd-status-contract.md` | Remove any active token-to-agent mapping for `security-applicability`; preserve `securityApplicability` fields as legacy/read-only status data. |
| `skills/_shared/persistence-contract.md` | Keep resolver rows for old `security-applicability.md` artifacts, but describe them as data-only archive refs, not phase outputs. |
| `skills/_shared/openspec-convention.md` | Keep the legacy file path in archive/read sections, but remove skill ownership wording. |
| `skills/_shared/sdd-security-contract.md` | Keep the legacy schema as the canonical readability contract for archived artifacts. Remove wording that implies routing through an executor. |
| `skills/_shared/security-guideline-catalog.md` | Keep legacy citation/readability language; avoid mentioning compatibility readers as phase executors. |
| `openspec/specs/sdd-security-applicability-workflow/spec.md` | Rewrite as a legacy artifact readability/data compatibility spec or fold its requirements into security/design/persistence specs later. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` and `openspec/specs/sdd-security-design-workflow/spec.md` | Verify they continue to forbid active `security-applicability` dependencies and preserve mandatory `security-design`. |
| `skills/sdd-review/references/report-template.md` | Remove `security-applicability` from the default input list or label it legacy/optional. |

`README.md` already presents the new phase list without `sdd-security-applicability` and only mentions `security-applicability.md` as legacy read-only evidence; no required change was found there.

### Should legacy artifact readability remain?

Yes. Preserve `security-applicability.md` readability as data-only compatibility through shared contracts and validators. Archived changes still contain real `security-applicability.md` and state entries with `securityApplicability`; readers should be able to understand historical classification, no-impact proof, catalog identity, and validation evidence without rerunning a phase.

The compatibility boundary should be explicit:

- keep schema vocabulary in `skills/_shared/sdd-security-contract.md`;
- keep `securityApplicability` fields in status/state as legacy/read-only;
- keep OpenSpec archive path conventions for old artifacts;
- do not expose a launchable executor or skill for new or old changes.

### Should `scripts/validate_security_applicability.ps1` stay?

Recommendation: keep it for now as a legacy/archive-only validator.

| Option | Pros | Cons | Effort |
| --- | --- | --- | --- |
| Keep validator with stronger legacy-only labeling | Preserves objective validation for archived `security-applicability.md`; matches current spec line that legacy validation MAY remain; avoids weakening audit readability. | Leaves one tool named after a retired artifact; users may confuse it with new-change validation unless docs stay explicit. | Low |
| Remove validator now | Removes all executable support for the retired phase/artifact. | Archived applicability artifacts lose repo-local validation; `validation.validator: scripts/validate_security_applicability.ps1` inside old artifacts becomes non-runnable historical metadata; specs and security contract need broader migration wording. | Medium |

Keeping the validator does not imply keeping the executor. Its header already says it MUST NOT block new-change routing and points new changes to `scripts/validate_security_design.ps1`.

## Risks

- **Archived OpenSpec audit trail:** Archived files mention the removed executor/skill as historical facts. They should not be edited, but search results will continue to show those mentions.
- **Status token normalization:** Consumers may still normalize legacy `security-applicability` state tokens. After removing the runnable executor, normalization must not launch a missing phase; it should treat the token as legacy data and route active work to `design` or `security-design`.
- **Sync scripts:** `scripts/sdd_init_agents.ps1` and `scripts/sdd_init_skills.ps1` copy all repo `agents/` and `skills/` contents to opencode/Copilot destinations, but do not delete stale destination files. After repo deletion, stale global copies may remain until manually cleaned or sync semantics change. Per repo-only scope, this exploration did not inspect or modify global copies.
- **Spec wording drift:** `openspec/specs/sdd-security-applicability-workflow/spec.md` currently names `sdd-security-applicability` in a persistence scenario. If left unchanged, it implies the executor can still exist.
- **Validator confusion:** Keeping `validate_security_applicability.ps1` is useful for archives but must stay clearly separated from mandatory `validate_security_design.ps1` for new changes.

## Recommendation

Proceed to proposal. Remove the repo-local `sdd-security-applicability` executor and skill, keep `security-applicability.md` as a legacy data artifact, and keep `scripts/validate_security_applicability.ps1` as an archive-only validator unless a later change intentionally migrates old artifacts to non-runnable validation metadata.

The implementation should update active source contracts so no new-change or compatibility path implies a runnable `sdd-security-applicability` executor still exists. Do not edit archived OpenSpec changes except through normal archive mechanics for this new change.

## Open Questions

None blocking. One proposal decision remains: whether to add cleanup behavior to sync scripts for stale destination files. The safer default for this change is no, because sync-script deletion semantics are broader than removing the repo-local executor/skill and could affect user-owned global files.

## Ready for Proposal

Yes. Suggested next phase: `propose`.
