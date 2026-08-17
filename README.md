# leo3nava_ai

Repository for reusable AI agent prompts, opencode skills, and Spec-Driven Development (SDD) workflow contracts. It is not an application runtime; it is a distribution and maintenance workspace for AI-assisted development instructions.

## Quick path

From the repository root, sync the current repository content into your Windows agent configurations:

```powershell
# Install or refresh skills
.\scripts\sdd_init_skills.ps1

# Install or refresh agents
.\scripts\sdd_init_agents.ps1
```

Expected destinations:

| Source | Destination |
|--------|-------------|
| `src/skills/` | `%USERPROFILE%\.config\opencode\skills` |
| `src/skills/` | `C:\Users\leo3n\.copilot\skills` |
| `src/agents/` | `%USERPROFILE%\.config\opencode\prompts` |
| `src/agents/` | `C:\Users\leo3n\.copilot\agents` |

The scripts only sync to tools whose base directories already exist. Missing destinations are skipped with an informational message, and existing destinations still receive the files. Destination subdirectories are created when needed and matching files are overwritten on rerun.

## Repository layout

| Path | Purpose |
|------|---------|
| `src/agents/` | Agent prompt definitions, including SDD phase agents. |
| `src/skills/` | Reusable opencode skills and shared SDD contracts. |
| `scripts/` | Windows PowerShell sync scripts for opencode and Copilot setup. |
| `tools-py/` | Python utility tools (report-exporter, token-monitor). |
| `openspec/` | Spec-driven documentation for repository capabilities and archived changes. |
| `LICENSE` | MIT license for this repository. |

## SDD workflow assets

The repository contains a full SDD workflow split into dedicated phase agents and skills:

- `sdd-init`
- `sdd-explore`
- `sdd-propose`
- `sdd-spec`
- `sdd-design`
- `sdd-test-design`
- `sdd-tasks`
- `sdd-apply`
- `sdd-review`
- `sdd-review-security`
- `sdd-verify`
- `sdd-archive`

Phase order: `explore? -> propose -> spec -> design -> test-design -> tasks -> apply -> review -> review-security -> verify -> archive`.

Manual post-archive utilities:

- `sdd-operational-doc` generates a Spanish operational handoff document from an archived SDD change. It is intentionally manual and is not part of the required SDD DAG.
- `sdd-technical-doc` generates a Spanish technical documentation Markdown file from an archived SDD change. It is intentionally manual, consumes archived evidence only, and is not part of the required SDD DAG, status, dependency, verify, or archive flow.

For new changes, `sdd-design` owns secure development design inside `design.md#secure-development-design`. Standalone `security-design.md` and `security-applicability.md` are legacy/archive compatibility data only and MUST NOT gate new-change routing. Previous standalone security validators have been removed from active workflow contracts; security review validates embedded rows against the shared catalog and artifact evidence.

The review phase writes canonical `review-report.json` plus a derived Markdown compatibility view: OpenSpec uses `openspec/changes/{change-name}/review-report.json` and `review-report.md`; Engram/hybrid uses `sdd/{change-name}/review-report.json` and `sdd/{change-name}/review`. Non-blocking general review routes to `sdd-review-security`, which writes canonical `review-security-report.json` plus derived Markdown before verification. Downstream phases may read derived Markdown for compatibility, but canonical review JSON and security-review JSON are authoritative when present; verification and archive consume review summaries without owning their matrices.

Shared contracts live under `src/skills/_shared/` and define persistence, status, security, language-domain, and executor-boundary rules used by the SDD agents.

## Requirements

- Windows PowerShell.
- `USERPROFILE` must be set for sync scripts.
- At least one supported base directory must exist to receive synced files: `%USERPROFILE%\.config\opencode` or `C:\Users\leo3n\.copilot`.
- No package manager, build system, or test runner is required by this repository.

## Verification

There is no application test runner configured. For script-related changes, verify behavior with isolated PowerShell smoke checks that avoid modifying the real user opencode directory unless that is the intended operation.

## License

MIT License. See `LICENSE` for details.
