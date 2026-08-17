## Artifact Language

All generated artifacts are written in English: code, identifiers, comments, UI copy, documentation, commit messages, and PR descriptions. Reply text to users follows the conversation language configured by persona and runtime instructions.

## YAML Frontmatter Schema

Skill files use YAML frontmatter with required fields:

- `name`
- `description`
- `license`
- `metadata`
  - `author`
  - `version`

## Section Heading Order

Skill documents follow this exact section order:

1. `Activation Contract`
2. `Hard Rules`
3. `Decision Gates`
4. `Execution Steps`
5. `Output Contract`
6. `References`

## Markdown Conventions

- Use level-2 headings (`##`) for main sections.
- Use level-3 headings (`###`) for subsections.
- Use pipe tables for Decision Gates and quality constraints.
- Use numbered lists for Execution Steps.
- Use bullet lists for Hard Rules and Output Contract fields.

## Frontmatter Note

Non-invocable or shared-only skills may add `user-invocable: false` and `disable-model-invocation: true` when runtime support is available.
