# Archive Report: project-conventions-skill

Status: success

## Executive summary

The `project-conventions-skill` SDD change is archived on 2026-07-11. Review and security review passed (review-report.json: PASS; review-security-report.json: PASS). All change artifacts and the full audit trail were moved into the archive directory. Permanent deliverables (skill files under `skills/project-conventions/` and the updated `AGENTS.md`) remain in place.

## What was archived

- All SDD artifacts produced during the change lifecycle: proposal, spec, design, test-design, tasks, apply-progress, review reports (JSON + Markdown), review-security reports (JSON + Markdown), verify report, and the canonical state file.

Archive location: `openspec/changes/archive/2026-07-11-project-conventions-skill/`

## Verification summary

- General review: PASS (96 controls, 0 failures) — `openspec/changes/archive/2026-07-11-project-conventions-skill/review-report.json`
- Security review: PASS (155 source rows validated, 0 blockers) — `openspec/changes/archive/2026-07-11-project-conventions-skill/review-security-report.json`
- Verify report: PASS — `openspec/changes/archive/2026-07-11-project-conventions-skill/verify-report.md`

No blockers were present and tooling that was unavailable (test runner, build, linter, etc.) was documented as unavailable in the review evidence. Those unavailable tooling notes are preserved in the archived review artifacts.

## Risks and notes carried forward

- This change introduces documentation-only skill files under `skills/project-conventions/`. They are permanent deliverables and were NOT moved to the archive. Consumers should treat these as authoritative guidance for future SDD phases.
- No unsafe evidence, secrets, or sensitive data were found in the archived artifacts.
- Unavailable automated tooling is documented and intentionally not treated as passed.

## Archive contents

Files moved to archive (canonical audit trail):

- proposal.md
- spec.md
- design.md
- test-design.md
- tasks.md
- apply-progress.md
- review-report.json
- review-report.md
- review-security-report.json
- review-security-report.md
- verify-report.md
- state.yaml (archived state with archive metadata)

## Artifacts

- type: archive-report
  path: openspec/changes/archive/2026-07-11-project-conventions-skill/archive-report.md
  persisted: true
  readable: true

## Next recommended

none

## Skill resolution

Permanent skill files created by this change (left in place):

- skills/project-conventions/SKILL.md
- skills/project-conventions/references/architecture.md
- skills/project-conventions/references/code-style.md
- skills/project-conventions/references/testing.md
- skills/project-conventions/references/review-readiness.md
- skills/project-conventions/references/secure-design.md
- AGENTS.md (modified)

Signed-off-by: Archive executor
