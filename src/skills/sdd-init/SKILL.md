---
name: sdd-init
description: "Trigger: sdd init, iniciar sdd, openspec init. Initialize SDD context, testing capabilities, registry, and persistence."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

<!--
Copyright (c) 2026 gentleman-programming
Licensed under the MIT License.

Modificaciones bajo Copyright (c) 2026 FC-OXXO
Licensed under the MIT License.
-->

> **ORCHESTRATOR GATE**: Follow `skills/_shared/executor-boundary.md`.
> This skill is for the dedicated `sdd-init` executor only.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Activation Contract

Run this phase when the orchestrator/user asks to initialize SDD in a project. 
You are the phase executor: do the work yourself, do not delegate, and do not behave like the orchestrator.

## Phase Artifact Contract

Common backend mechanics: follow `skills/_shared/persistence-contract.md` through **Section B** (retrieval) and **Section C** (persistence) in `skills/_shared/sdd-phase-common.md`.

| Concern | Contract |
| --- | --- |
| Required inputs | User/orchestrator initialization request, resolved artifact-store mode, project root, and real repository files used for stack/tooling detection. |
| Produced artifacts | Project context `sdd-init/{project}`; testing capabilities `sdd/{project}/testing-capabilities`; `openspec/config.yaml`; local support registry `.atl/skill-registry.md`; Engram `skill-registry` when the selected mode supports each artifact. |
| Mutates | OpenSpec bootstrap directories/files only in OpenSpec-capable modes; Engram init/testing/registry observations only in Engram-capable modes; `.atl/skill-registry.md` only in modes that allow local support files. |
| Initialization semantics | Detect real stack, conventions, architecture, testing tools, strict TDD status, and persistence context before writing. Existing OpenSpec artifacts require explicit update handling instead of blind overwrite. |
| Local support artifact semantics | Build the skill registry from actual skill paths and persist it as `.atl/skill-registry.md`. |
| Conditional behavior | `engram` does not create `openspec/`; `openspec` does not call Engram. |
| Success routing | `next_recommended: sdd-new` or `sdd-explore`, according to the initialization envelope. |
| Block routing | `next_recommended: resolve-blockers` for unsafe paths, unresolved existing OpenSpec update decisions, unavailable persistence backend, or failed read-back verification. |

## Hard Rules

- Detect the real stack, conventions, architecture, testing tools, and persistence mode; never guess.
- Follow the Phase Artifact Contract for produced artifacts, mutation boundaries, and mode-specific write eligibility; follow `skills/_shared/persistence-contract.md` for common backend mechanics.
- Persist project context as `sdd-init/{project}` and testing capabilities separately as `sdd/{project}/testing-capabilities` when Engram-capable modes allow those writes.
- Persist project context and testing capabilities in `openspec/config.yaml` when OpenSpec-capable modes allow those writes.
- Build the skill registry following `skill-registry` rules and preserve the local support artifact rules from the Phase Artifact Contract.
- Save the registry to Engram as `skill-registry` when Engram-capable modes allow that write.
- Use `capture_prompt: false` for automated SDD/config saves when supported; omit it if the tool schema lacks it.
- If `openspec/` already exists, report what exists and ask before updating it.

## Decision Gates

| Input | Action |
|---|---|
| `mode=engram` | Save context and capabilities to Engram only. |
| `mode=openspec` | Create/update openspec bootstrap files only. |
| `openspec/` already exists | Report existing files and ask before updating OpenSpec artifacts. |
| strict TDD marker/config found | Use that value. |
| no marker/config but test runner exists | Default `strict_tdd: true`. |
| no test runner | Set `strict_tdd: false` and explain unavailable. |

## Execution Steps

1. Load supplemental skills according to `skills/_shared/skill-resolver.md` and the executor minimum in `skills/_shared/sdd-phase-common.md` Section A.
2. Inspect project files (`package.json`, `app.json`, `app.config.js`, `app.config.ts`, `metro.config.js`, `react-native.config.js`, `android/`, `ios/`, `go.mod`, `pyproject.toml`, `pom.xml`, `build.gradle`, `build.gradle.kts`, `settings.gradle`, `settings.gradle.kts`, `*.kt`, `*.kts`, CI, lint/test config) and summarize stack/conventions.
3. Detect test runner, test layers, coverage, linter, type checker, and formatter.
4. Resolve Strict TDD from agent marker, `openspec/config.yaml`, detected runner fallback, or no-runner fallback.
5. Initialize persistence for the resolved mode.
6. Build `.atl/skill-registry.md` using the skill-registry scan rules.
7. Persist project context, testing capabilities, and registry using the exact artifact names for the resolved mode.
8. Return the structured initialization envelope.

## Output Contract

Return the full Section D envelope from `skills/_shared/sdd-phase-common.md`: `status`, `executive_summary`, `detailed_report`, `artifacts`, `next_recommended`, `risks`, and `skill_resolution`.
Include project, stack, persistence mode, Strict TDD status, testing capability table, registry path, and next `/sdd-explore` or `/sdd-new` step.
The `artifacts` list MUST explicitly report each artifact as saved, skipped, or blocked:

- `sdd-init/{project}` Engram project context observation ID, when applicable.
- `sdd/{project}/testing-capabilities` Engram testing capabilities observation ID, when applicable.
- `skill-registry` Engram observation ID, when applicable.
- `.atl/skill-registry.md` path.
- `openspec/config.yaml` path, when applicable.

## References

- [references/init-details.md](references/init-details.md) — detection checklist, Engram payloads, config skeleton, and output templates.
- `../_shared/skill-resolver.md` — supplemental skill loading and `skill_resolution` protocol.
- `../_shared/sdd-phase-common.md` — phase retrieval, persistence, and return envelope.
- `../_shared/engram-convention.md` — Engram artifact naming.
- `../_shared/openspec-convention.md` — openspec layout and rules.
