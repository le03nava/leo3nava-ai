# Tasks: SDD Init Opencode Prompts

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 25-60 |
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
| 1 | Add and manually verify `scripts/sdd_init_opencode_prompts.ps1` | PR 1 | Single reviewable script plus manual evidence; base `main`. |

## Phase 1: Script Foundation

- [x] 1.1 Create `scripts/sdd_init_opencode_prompts.ps1` following the PowerShell sync convention: `$ErrorActionPreference = 'Stop'`, `Split-Path -Parent $PSScriptRoot`, and `Join-Path`.
- [x] 1.2 Validate `$env:USERPROFILE` before filesystem changes; failure text must identify missing `USERPROFILE` for scenario “User home is unavailable”.
- [x] 1.3 Resolve source as `<repo root>\agents` from `$PSScriptRoot`, not current working directory, for scenarios “Run from repository root” and “Run from another working directory”.
- [x] 1.4 Validate `agents/` exists before creating destination; failure text must identify missing repository `agents` for scenario “Source directory is missing”.

## Phase 2: Synchronization Behavior

- [x] 2.1 Set destination to `$env:USERPROFILE\.config\opencode\prompts` and create it with `New-Item -ItemType Directory -Force` for scenario “Destination is missing”.
- [x] 2.2 Copy direct contents of repository `agents/` to the destination with recursive force copy so scenario “First sync copies agents” creates matching relative paths.
- [x] 2.3 Preserve unrelated destination files by avoiding destination deletion, satisfying scenario “Unrelated destination files are preserved”.
- [x] 2.4 Overwrite matching destination files on rerun, satisfying scenario “Rerun replaces copied files”.
- [x] 2.5 Keep scope limited to repository `agents/`; do not reference `skills`, Claude, or opencode installation/configuration paths for adapter scope scenarios.

## Phase 3: Manual Verification

- [x] 3.1 Run `powershell -ExecutionPolicy Bypass -File .\scripts\sdd_init_opencode_prompts.ps1` from repo root with a disposable `$env:USERPROFILE`; verify `.config\opencode\prompts\sdd` is populated.
- [x] 3.2 Run the script from a directory outside the repo and verify the same repository `agents/` source is copied.
- [x] 3.3 Modify a copied destination file, add an unrelated destination file, rerun, and verify the copied file is overwritten while unrelated content remains unchanged.
- [x] 3.4 In isolated PowerShell sessions, verify empty `USERPROFILE` and missing `agents/` fail before copy/destination mutation with clear errors.
