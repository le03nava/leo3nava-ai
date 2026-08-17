# Design: Add Opencode Skill Sync Script

## Technical Approach

Add one repo-local Windows PowerShell entry point, `scripts/sdd_init_opencode_skills.ps1`, that copies this repository's `skills/` directory contents into the current user's opencode skills directory: `%USERPROFILE%\.config\opencode\skills`.

The script resolves paths from its own location with `$PSScriptRoot`, validates required inputs before creating or copying anything, creates the destination when valid, and performs a recursive replacement copy so reruns refresh existing files deterministically.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
|---|---|---|---|
| Script location | Create `scripts/sdd_init_opencode_skills.ps1` | Put the script beside `skills/` or under OpenSpec | `scripts/` is the expected utility location from the proposal; it keeps runtime tooling separate from skill content and SDD artifacts. |
| Repository path resolution | Resolve repo root from `$PSScriptRoot` as the parent of `scripts/` | Depend on current working directory | Invocation should work from any shell location. `$PSScriptRoot` binds behavior to the script file, not caller state. |
| Destination path | Build `%USERPROFILE%\.config\opencode\skills` with `Join-Path` after validating `$env:USERPROFILE` | Hard-code user paths or infer from `$HOME` | The spec explicitly targets `USERPROFILE`; validating it gives clearer Windows behavior and error messages. |
| Copy behavior | Copy `skills\*` recursively with replacement enabled | Mirror/delete destination first, copy selected skills, or prompt | The requirement is direct all-content copy without filtering or prompts. Replacement makes reruns deterministic while avoiding destructive deletion of unrelated destination files. |
| Error model | Fail fast with clear `throw`/terminating errors before destination changes | Best-effort warnings | Missing `USERPROFILE` or source `skills/` MUST stop before copy or destination creation. Terminating errors are easiest to detect manually. |

## Data Flow

```text
scripts/sdd_init_opencode_skills.ps1
  ├─ $PSScriptRoot
  ├─ repo root = parent of scripts/
  ├─ source = repo root/skills
  ├─ destination = $env:USERPROFILE/.config/opencode/skills
  └─ Copy-Item source/* → destination -Recurse -Force
```

Validation order:

1. Check `$env:USERPROFILE` is not null, empty, or whitespace.
2. Resolve source `skills/` relative to `$PSScriptRoot`.
3. Check source directory exists.
4. Create destination directory if needed.
5. Copy direct source contents recursively into destination with replacement.

This order preserves the spec requirement that missing inputs do not create or modify the opencode destination.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `scripts/sdd_init_opencode_skills.ps1` | Create | PowerShell utility that validates inputs, creates `%USERPROFILE%\.config\opencode\skills`, and recursively copies `skills\*` with replacement. |

No application code changes are required. The repository currently has `skills/` content and no detected `scripts/` directory or `.ps1` files, so the implementation will introduce the scripts directory as needed.

## Interfaces / Contracts

The script exposes a file-based CLI contract:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\sdd_init_opencode_skills.ps1
```

Runtime contract:

- Input: repository `skills/` directory and `$env:USERPROFILE`.
- Output: `%USERPROFILE%\.config\opencode\skills` exists and contains the direct contents of repository `skills/`.
- Failure: clear terminating error mentioning `USERPROFILE` when missing, or source `skills` when missing.
- No interactive prompts.

## Testing Strategy

The OpenSpec config reports no package manifest, test runner, linter, formatter, type checker, coverage command, or automated test layers. Automated verification is therefore unavailable for this change.

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Script smoke | Normal sync creates destination and copies content | Manual command: `powershell -ExecutionPolicy Bypass -File .\scripts\sdd_init_opencode_skills.ps1`, then inspect `%USERPROFILE%\.config\opencode\skills`. |
| Script smoke | Rerun replaces existing files without prompt | Modify a copied destination file, rerun command, confirm it matches source. |
| Failure path | Missing `USERPROFILE` fails before copying | In an isolated PowerShell process, clear `USERPROFILE`, run the script, and confirm a clear error. |
| Failure path | Missing source fails before destination changes | Temporarily run from a copied repo fixture without `skills/`, or rename only in a disposable checkout, and confirm clear error. |

## Migration / Rollout

No migration required. Users opt in by running the new script. Existing opencode skill files with matching relative paths are overwritten during execution as specified.

## Open Questions

None.
