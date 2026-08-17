## Review Budget

The canonical review budget for this project is **400 changed lines per PR**.

## Delivery Strategy

- Default strategy: `auto-chain`
- Chain mode: `stacked-to-main`

When a change is split, each PR branch targets the previous PR branch in the stack, not `main` directly.

## Chained PR Requirement

If the estimated or actual diff exceeds 400 lines, the change must be split into chained work units using the `chained-pr` skill:

- Skill path: `skills/chained-pr/SKILL.md`
- Goal: keep each PR reviewable within the budget while preserving delivery continuity.

## PR Conventions

- **Issue-first**: a GitHub issue must exist (or be created) before opening a PR.
  - Use: `skills/issue-creation/SKILL.md`
- **Conventional commits** are required:
  - `feat:`
  - `fix:`
  - `docs:`
  - `refactor:`
  - `test:`
  - `chore:`
- **No AI attribution** in commit messages.
  - Do not add `Co-Authored-By`.
  - Do not add any AI-assistant attribution lines.
- **Branch naming** should follow `{type}/{issue-number}-{slug}`.

## PR Readiness Checklist

Before creating a PR, verify all of the following:

1. Working tree status is clean or intentionally staged (`git status`).
2. The proposed code delta is reviewed (`git diff`).
3. Remote tracking is correct for the branch and base.
4. Recent commit history is coherent (`git log --oneline -10`).
5. The diff against the base branch is correct and review-scoped.
