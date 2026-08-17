# Proposal: Project Conventions Skill

## Intent

SDD design artifacts (design.md, test-design.md, tasks.md) currently reference generic conventions.
No mechanism exists for SDD phase agents to load repo-specific naming, architecture, testing, or
secure-design commitments. The result: conventions are either implicit (invisible to reviewers) or
copy-pasted into every change, creating drift. This change ships `skills/project-conventions/SKILL.md`
as a loadable supplemental skill so SDD phases can reference real, version-controlled project facts
without coupling into sdd-review's 96-control matrix or sdd-review-security's catalog.

## Scope

### In Scope

- `skills/project-conventions/SKILL.md` — activation contract, hard rules, decision gates, execution steps, output contract
- `skills/project-conventions/references/architecture.md` — SDD executor model, shared contracts, boundary patterns
- `skills/project-conventions/references/code-style.md` — English artifacts rule, Markdown/YAML conventions, frontmatter schema
- `skills/project-conventions/references/testing.md` — strict_tdd: false, no automated runner, verification strategy
- `skills/project-conventions/references/review-readiness.md` — 400-line budget, auto-chain, stacked-to-main, PR readiness
- `skills/project-conventions/references/secure-design.md` — narrative secure-design guidelines per surface (inputs, outputs, storage, logs, auth, secrets, external APIs, operational security)
- Update `AGENTS.md` to add a `## Project Skills` section registering project-conventions

### Out of Scope

- No changes to sdd-review, sdd-review-security, or their catalogs
- No SEC-* matrices, no Source ID rows, no scoring tables in any reference file
- No automated test infrastructure (strict_tdd remains false, openspec/config.yaml unchanged)
- No changes to sdd-design SKILL.md skill contract (skill injection already works via skill-resolver)

## Capabilities

### New Capabilities

- `project-conventions-skill`: Loadable supplemental skill that surfaces repo-specific conventions (architecture, code-style, testing, review-readiness, secure-design) to SDD phase agents at design and apply time

### Modified Capabilities

None — existing specs describe the SDD workflow phases; this adds a supporting skill artifact only.

## Approach

Follow skill-creator pattern exactly: YAML frontmatter (`name`, `description` with trigger keywords,
`license`, `metadata`) + fixed section headings (Activation Contract, Hard Rules, Decision Gates,
Execution Steps, Output Contract, References).

Reference files are compact authoritative Markdown. Where evidence is missing or unconfirmed, items
are marked `[needs confirmation]`. `secure-design.md` uses a narrative rule structure
(`design-commitment / prohibited-pattern / expected-evidence / evidence-owner / residual-risk /
exception-policy`) — NO scoring rows. SKILL.md loads secure-design.md conditionally when the
change touches data, auth, logs, exports, files, DB, config/secrets, external APIs, or operational
behavior.

sdd-design integration: document in SKILL.md how project-conventions should be applied when injected
as a supplemental skill (adds `## Implementation Quality Constraints` section with optional
`## Secure Design Constraints` subsection to design.md); no edits to sdd-design SKILL.md required.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/project-conventions/SKILL.md` | New | Main skill activation contract |
| `skills/project-conventions/references/` | New | 5 reference files (architecture, code-style, testing, review-readiness, secure-design) |
| `AGENTS.md` | Modified | Add `## Project Skills` section registering project-conventions |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| skill-creator pattern not fully aligned | Medium | Read skill-creator/SKILL.md during spec; align frontmatter/sections before writing |
| Boundary drift (someone adds review scoring to project-conventions) | Low | Hard rule in SKILL.md explicitly prohibits SEC-* matrices and scoring |
| testing.md becomes stale when strict_tdd changes | Low | Add explicit "update this file when strict_tdd changes" note in testing.md |

## Rollback Plan

Delete `skills/project-conventions/` and revert the `AGENTS.md` `## Project Skills` section.
No other files are modified; no existing skill contract changes. Zero downstream breakage.

## Dependencies

- `skills/skill-creator/SKILL.md` — must be read during spec to confirm frontmatter and section schema

## Success Criteria

- [ ] `skills/project-conventions/SKILL.md` passes skill-creator frontmatter schema (name, description with trigger, license, metadata)
- [ ] All 5 reference files present and non-empty; no SEC-* matrices or scoring tables in any of them
- [ ] `secure-design.md` uses narrative rules only (design-commitment / prohibited-pattern / expected-evidence)
- [ ] `AGENTS.md` `## Project Skills` section registers project-conventions with trigger keywords
- [ ] sdd-design can load project-conventions as supplemental skill without changes to its own SKILL.md
