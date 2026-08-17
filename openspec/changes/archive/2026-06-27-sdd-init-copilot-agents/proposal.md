# Proposal: SDD Init Opencode Prompts

## Intent

Create a Windows-friendly initialization script that refreshes the current user's opencode prompts from this repository's `agents/` source directory. The script must be rerunnable, avoid machine-specific hardcoded user paths, and replace existing destination files when repository prompt sources change.

## Scope

### In Scope
- Add a Windows script named `sdd_init_opencode_prompts.ps1` for syncing repository `agents/` contents as opencode prompts.
- Resolve the destination from the current user's home environment, targeting `%USERPROFILE%\.config\opencode\prompts` or the PowerShell equivalent.
- Create the destination directory when missing and overwrite existing destination files on rerun.
- Fail with clear errors when the user-home environment value or repository `agents/` source is unavailable.

### Out of Scope
- Implementing the script in this proposal phase.
- Syncing repository skills, Claude prompts, or other adapter folders.
- Installing or configuring opencode itself.
- Cross-platform shell support beyond Windows.

## Capabilities

### New Capabilities
- `opencode-prompt-sync`: Defines repo-local behavior for initializing or refreshing opencode prompts from repository `agents/` into the current Windows user's opencode prompts directory.

### Modified Capabilities
- None.

## Approach

Add a small Windows script under `scripts/` that computes its repository-relative `agents/` source, reads the user home from a Windows environment variable instead of hardcoding a user-specific home path, ensures `%USERPROFILE%\.config\opencode\prompts` exists, then copies source contents with replacement enabled.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `scripts/` | New | Add the `sdd_init_opencode_prompts.ps1` Windows sync script. |
| `agents/` | Read-only | Source content for opencode prompt synchronization. |
| `openspec/specs/opencode-prompt-sync/spec.md` | New | Future capability spec for expected sync behavior. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Overwriting local opencode prompt edits | Medium | Document deterministic replacement and scope it only to matching copied contents under `.config\opencode\prompts`. |
| Script run from unexpected working directory | Medium | Resolve source relative to script location, not current shell directory. |
| Missing or unusual home environment | Low | Fail fast with a clear `USERPROFILE` error before copying. |

## Rollback Plan

Remove the new script and the `opencode-prompt-sync` delta spec. If the script was already run, restore `%USERPROFILE%\.config\opencode\prompts` from backup or manually remove copied prompt files.

## Dependencies

- Repository `agents/` directory must exist.
- Windows user-home environment variable must be available.

## Success Criteria

- [ ] The proposal defines intent, scope, non-goals, assumptions, risks, and rollback.
- [ ] The next spec phase can create `opencode-prompt-sync` requirements without guessing source or destination behavior.
- [ ] Implementation remains deferred until apply phase.
