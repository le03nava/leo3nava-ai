# Project Instructions

Use this file to document repository-specific instructions for AI coding agents.

## Manual SDD Utilities

- `sdd-operational-doc` is a manual post-archive utility. It generates a Spanish operational handoff document from an archived SDD change and must not be treated as a required SDD DAG phase.
- `sdd-technical-doc` is a manual post-archive utility. It generates a Spanish technical documentation Markdown file from an archived SDD change and must not be treated as a required SDD DAG phase.
- `install-package-builder` is a manual user-invocable utility. It builds a versioned installation package from git repositories for a change order and must not be treated as a required SDD DAG phase.

## Project Skills

- `project-conventions` is a supplemental skill loaded by SDD phases (sdd-design, sdd-test-design, sdd-tasks, sdd-apply) when the change requires real project conventions. It provides architecture patterns, code style, testing infrastructure state, delivery conventions, and secure design guidance. It is NOT a required SDD DAG phase and must not be treated as one. Location: `skills/project-conventions/SKILL.md`.
