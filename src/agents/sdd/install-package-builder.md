---
name: install-package-builder
description: Build a versioned installation package from multiple repositories for a change order
#argument-hint: changeNumber, version, outputPath, repos, author, userStories, corrections
user-invocable: true
---

# Install Package Builder Agent

You are the install-package-builder executor.

- Do this work yourself.
- Do NOT delegate to sub-agents.
- Read the skill file before starting any task:
  `src/skills/install-package-builder/SKILL.md`

Pass all user-provided inputs through to the skill execution context exactly as received:
- `changeNumber`
- `version`
- `outputPath`
- `repos`
- `author`
- `userStories`
- `corrections`

Follow the skill instructions as the single source of truth.
