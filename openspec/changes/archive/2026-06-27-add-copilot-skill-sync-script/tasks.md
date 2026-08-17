# Tasks: Add Opencode Skill Sync Script

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 40-70 |
| Review budget lines | 400 |
| 400-line budget risk | Low |
| Review budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: Low
400-line budget risk: Low

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Add the PowerShell sync script and verify manually | PR 1 | Single focused change; stacked-to-main is available but not needed for low risk. |

## Phase 1: Foundation

- [x] 1.1 Create `scripts/` if absent and add `scripts/sdd_init_opencode_skills.ps1` as the repo-local Windows entry point.
- [x] 1.2 In `scripts/sdd_init_opencode_skills.ps1`, resolve the repository root from `$PSScriptRoot` instead of caller working directory.
- [x] 1.3 In `scripts/sdd_init_opencode_skills.ps1`, validate `$env:USERPROFILE` is not null, empty, or whitespace before creating destination paths.

## Phase 2: Core Implementation

- [x] 2.1 In `scripts/sdd_init_opencode_skills.ps1`, build the source path as `<repo-root>\skills` and fail with a clear missing source error when it does not exist.
- [x] 2.2 In `scripts/sdd_init_opencode_skills.ps1`, build `%USERPROFILE%\.config\opencode\skills` using `Join-Path` and create it only after input validation succeeds.
- [x] 2.3 In `scripts/sdd_init_opencode_skills.ps1`, copy direct contents from `skills\*` into the opencode skills directory recursively with replacement enabled and no prompts.
- [x] 2.4 In `scripts/sdd_init_opencode_skills.ps1`, emit concise success output showing source and destination paths.

## Phase 3: Manual Smoke Verification

- [x] 3.1 Verify Scenario "Initial sync creates destination and copies skills" by running `powershell -ExecutionPolicy Bypass -File .\scripts\sdd_init_opencode_skills.ps1` and inspecting `%USERPROFILE%\.config\opencode\skills`.
- [x] 3.2 Verify Scenario "Existing destination receives all source contents" by rerunning the script against an existing destination.
- [x] 3.3 Verify Scenario "Rerun overwrites previously copied files" by editing a copied destination file, rerunning, and confirming it matches the source file.
- [x] 3.4 Verify Scenario "Missing USERPROFILE fails without destination changes" in an isolated PowerShell process with `USERPROFILE` cleared.
- [x] 3.5 Verify Scenario "Missing source directory fails without destination changes" using a disposable checkout or temporary fixture without `skills/`.

## Phase 4: Cleanup

- [x] 4.1 Confirm no application code or existing `skills/` content was modified.
- [x] 4.2 Confirm `openspec/changes/add-copilot-skill-sync-script/tasks.md` remains aligned with proposal, spec, design, and OpenSpec task rules.
