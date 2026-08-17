# Proposal: Install Package Builder

## Intent

Manual assembly of installation packages for change orders (e.g., CHG0086767) is error-prone, slow, and depends on tribal knowledge. Developers must clone repos, run builds, sort SQL by object type, copy artifacts to the correct folder hierarchy, and write a README by hand — every time. A single mistake (wrong artifact path, missing SQL prefix, outdated README) causes deployment failures in production. This change introduces a standalone automated agent that executes the full pipeline from a change order number to a ready-to-ship zip package.

## Scope

### In Scope
- Standalone `install-package-builder` skill/agent (not an SDD DAG phase)
- Clone or pull all git repositories associated with a change order
- Tech-stack detection and build execution: Java/Maven, Angular, React, Kotlin/Android, React Native, Oracle DB (SQL only — no build)
- SQL file classification into Packages, Procedures, Types, Tables, Data subfolders with numeric prefix handling
- Output folder structure: `{CHG}_VERSION{N}/Install/` and `Rollback/` mirror
- README.txt generation from a template (Información General, Pasos de instalación, Control de Versiones, Control de Revisiones)
- Pre-flight tool check (git, mvn, npm, ng, gradlew)
- Partial-failure handling: one repo failure reports and continues
- Final zip of the completed package
- Reference asset with the 24 OWEXX environments (host + containers)

### Out of Scope
- Deploying packages to any server or WebLogic instance
- Rollback artifact content generation (structure is mirrored, content is user responsibility)
- Change order data retrieval from ServiceNow or any ITSM system (repo list is provided as input)
- Authentication/secrets management for private registries beyond standard git credentials

## Capabilities

### New Capabilities
- `install-package-builder`: End-to-end pipeline that takes a change order number + repo list and produces a versioned, zipped installation package with classified artifacts and a generated README.txt

### Modified Capabilities
- None

## Approach

Implement as a skill following the `sdd-operational-doc` / `sdd-technical-doc` pattern:
- Skill source at `src/skills/install-package-builder/SKILL.md`
- Environment reference asset at `src/skills/install-package-builder/assets/environments.json`
- README template at `src/skills/install-package-builder/assets/readme-template.txt`
- Adapter prompt at `agents/sdd/install-package-builder.md` (copy) and `~/.config/opencode/prompts/sdd/install-package-builder.md`
- Agent entry in `~/.config/opencode/opencode.json`
- Entry in project `AGENTS.md`

Pipeline execution uses bash tool calls sequentially per repo; SQL classification is string-matching logic encoded in the skill instructions.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `src/skills/install-package-builder/` | New | Skill definition, assets, README template |
| `agents/sdd/install-package-builder.md` | New | Adapter prompt (copy) |
| `~/.config/opencode/prompts/sdd/install-package-builder.md` | New | Actual agent prompt |
| `~/.config/opencode/opencode.json` | Modified | Register new agent |
| `AGENTS.md` | Modified | Document new skill |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Build tool not installed on agent host | Med | Pre-flight check fails fast with clear message listing missing tools |
| Repo clone fails (auth, wrong URL) | Med | Report per-repo failure, continue with remaining repos |
| SQL classification misses edge cases | Low | Rules are explicit; unclassified files are copied to a `DBObjects/Unclassified/` fallback and reported |
| README template diverges from real format | Low | Template stored as versioned asset; user can override via skill input |
| Large monorepos slow down the pipeline | Low | Shallow clone (`--depth 1`) by default |

## Rollback Plan

The skill creates no permanent side effects on the source repos or any server. If the output package is wrong: delete the generated `{CHG}_VERSION{N}/` folder and re-run with corrected inputs. No state to revert.

## Dependencies

- `git` available on agent host
- At least one of: `mvn`, `npm`/`ng`, `./gradlew` depending on the repos in the change order
- Write access to the configured output directory

## Success Criteria

- [ ] Given a change order number and a list of repo URLs, the agent produces a correctly structured `{CHG}_VERSION{N}.zip` without manual steps
- [ ] SQL files are classified into the correct subfolders with numeric prefixes preserved or assigned
- [ ] README.txt is generated with correct environment data for all 24 OWEXX environments
- [ ] A single failing repo does not abort the entire pipeline; failures are reported in the summary
- [ ] Pre-flight check catches missing tools before any work begins
