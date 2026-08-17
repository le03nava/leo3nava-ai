# Design: SDD Init Opencode Prompts

## Technical Approach

Add one Windows PowerShell script at `scripts/sdd_init_opencode_prompts.ps1` to refresh opencode prompt files from the repository `agents/` directory into `%USERPROFILE%\.config\opencode\prompts`.

The script follows the existing PowerShell sync convention: `$ErrorActionPreference = 'Stop'`, `$PSScriptRoot`-based repo resolution, `$env:USERPROFILE` validation, `Join-Path`, `New-Item -Force`, `Copy-Item -Recurse -Force`, and concise success output.

The repository currently has `agents/` with SDD prompt source Markdown files under `agents/sdd/`, so implementation is not blocked. Runtime still fails fast if that source directory is missing in another checkout.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
|---|---|---|---|
| Script path and filename | Create `scripts/sdd_init_opencode_prompts.ps1` | Extensionless file; place under `agents/` | Existing repo pattern uses PowerShell scripts in `scripts/` with `.ps1`. The `.ps1` extension is the Windows executable contract and makes invocation unambiguous. |
| Source resolution | Resolve repo root as parent of `$PSScriptRoot`, then source as `agents` | Resolve from current working directory | The spec requires command-location resolution so the script works even when launched from outside the repository root. |
| Destination resolution | Validate `$env:USERPROFILE`, then target `.config\opencode\prompts` below it | Hardcode `C:\Users\...` or use unrelated adapter paths | Matches `%USERPROFILE%` semantics and avoids machine-specific paths. |
| Copy semantics | Copy direct contents of `agents/` into destination with `Copy-Item -Recurse -Force` | Delete destination first; copy all adapter folders; prompt before overwrite | `-Force` overwrites matching relative paths on rerun while preserving unrelated destination files. Scope remains only repository `agents/` content. |
| Failure model | Throw terminating errors before destination creation/copy when home or source is missing | Warn and continue | Clear fail-fast errors satisfy missing-home and missing-source requirements and prevent partial sync. |

## Data Flow

```text
scripts/sdd_init_opencode_prompts.ps1
  ├─ validate $env:USERPROFILE
  ├─ repo root = parent of $PSScriptRoot
  ├─ source = repo root/agents
  ├─ validate source exists
  ├─ destination = $env:USERPROFILE/.config/opencode/prompts
  ├─ create destination directory
  └─ copy source/* → destination with -Recurse -Force
```

Validation order is important: missing `USERPROFILE` or missing source stops before `New-Item` and before any `Copy-Item` operation.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `scripts/sdd_init_opencode_prompts.ps1` | Create | PowerShell sync utility for repository `agents/` to `%USERPROFILE%\.config\opencode\prompts`. |

No existing script is modified. No repository skills, Claude prompts, or other adapter folders are synchronized.

## Interfaces / Contracts

CLI contract:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\sdd_init_opencode_prompts.ps1
```

Runtime contract:

- Input source: repository `agents/` directory relative to the script location.
- Destination: `$env:USERPROFILE\.config\opencode\prompts`.
- Existing matching destination files are overwritten.
- Destination files without a matching source relative path are preserved.
- Failures are terminating errors that mention the missing `USERPROFILE` or missing repository `agents` source.
- No opencode installation/configuration is attempted.

## Testing Strategy

`openspec/config.yaml` reports no executable test runner, linter, formatter, type checker, coverage command, or automated test layers. Verification should use manual PowerShell/static checks.

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Static inspection | Path resolution and scope | Confirm script uses `$PSScriptRoot`, source `agents`, destination `.config\opencode\prompts`, and does not reference `skills`, Claude, or unrelated adapter paths. |
| Manual smoke | Destination creation and copy | In PowerShell, set `$env:USERPROFILE` to a disposable temp home, run the script, then inspect `.config\opencode\prompts\sdd`. |
| Manual rerun | Overwrite matching files, preserve unrelated files | Modify a copied destination agent, add an unrelated destination file, rerun, confirm copied file matches source and unrelated file remains. |
| Manual failure | Missing home/source fail clearly | Run in an isolated PowerShell process with empty `USERPROFILE`; separately test from a disposable checkout without `agents/`. |

## Migration / Rollout

No migration required. Users opt in by running the new script. Reruns deterministically refresh matching opencode prompt files.

## Open Questions

None.
