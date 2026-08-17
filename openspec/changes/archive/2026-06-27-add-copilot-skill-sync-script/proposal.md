# Proposal: Add Opencode Skill Sync Script

## Intent

Provide a repo-local Windows utility that initializes opencode skills by copying this repository's `skills/` contents into the user's opencode skills directory. This removes the manual copy step and makes local skill setup repeatable.

## Scope

### In Scope
- Add `scripts/sdd_init_opencode_skills.ps1` as the Windows sync entry point.
- Copy the direct contents of repository `skills/` into `%USERPROFILE%\.config\opencode\skills`.
- Replace existing files at the destination to make reruns deterministic.
- Fail fast when `$env:USERPROFILE` is empty or the source `skills/` directory is missing.

### Out of Scope
- Cross-platform shell scripts.
- Selective skill filtering or interactive prompts.
- Changes to existing skill content, opencode prompt adapters, or Claude adapters.

## Capabilities

### New Capabilities
- `opencode-skill-sync`: Covers initializing or refreshing opencode's local skills directory from this repository's `skills/` source.

### Modified Capabilities
- None.

## Approach

Create a PowerShell script that resolves the repository root relative to the script location, validates `USERPROFILE`, builds paths with `Join-Path`, creates `%USERPROFILE%\.config\opencode\skills` if missing, and copies `skills\*` directly into that destination with replacement enabled.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `scripts/sdd_init_opencode_skills.ps1` | New | PowerShell sync script for opencode skills. |
| `%USERPROFILE%\.config\opencode\skills` | Modified at runtime | Destination files are overwritten when the script is run. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Existing opencode skills are overwritten | High | Make replacement behavior explicit in docs/specs and keep script deterministic. |
| PowerShell execution policy blocks direct execution | Medium | Support documented invocation with `powershell -ExecutionPolicy Bypass -File ...`. |
| Internal runtime skills are copied too | Medium | Preserve requested all-skills behavior; defer filtering as future work. |

## Rollback Plan

Revert the added script. Users can manually restore `%USERPROFILE%\.config\opencode\skills` from their own backups or delete copied skill folders before rerunning a future corrected script.

## Dependencies

- Windows PowerShell.
- Repository `skills/` directory.
- `$env:USERPROFILE` pointing to the target user profile.

## Success Criteria

- [ ] Running the script creates `%USERPROFILE%\.config\opencode\skills` when absent.
- [ ] Running the script copies repository `skills/` contents directly into the opencode skills directory.
- [ ] Rerunning the script replaces existing destination files without requiring manual cleanup.
- [ ] Missing `USERPROFILE` or source `skills/` fails with a clear error.
