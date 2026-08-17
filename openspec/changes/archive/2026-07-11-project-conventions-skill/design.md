# Design: Project Conventions Skill

## Technical Approach

Ship a new supplemental skill at `skills/project-conventions/` that SDD phase executors can load via skill-resolver injection. The skill exposes five compact reference files covering architecture, code-style, testing, review-readiness, and secure-design conventions. When injected, the phase agent reads the SKILL.md, follows its execution steps, and appends an `## Implementation Quality Constraints` section to its output artifact (e.g., design.md). No existing sdd-design, sdd-review, or sdd-review-security skill contracts are modified.

## Architecture Decisions

### Decision: Supplemental skill (not embedded in sdd-design)

**Choice**: Standalone supplemental skill loaded via skill-resolver
**Alternatives considered**: Embedding conventions directly in sdd-design SKILL.md; a shared-contract file in `skills/_shared/`
**Rationale**: Keeps sdd-design generic and reusable across projects. Skill-resolver already supports supplemental injection — no new mechanism needed. Conventions are project-specific, not shared infrastructure.

### Decision: Reference files as separate Markdown (not inline in SKILL.md)

**Choice**: Five reference files under `skills/project-conventions/references/`
**Alternatives considered**: One monolithic SKILL.md with all content; YAML/JSON config files
**Rationale**: Markdown references are human-readable, independently updatable, and match the pattern used by sdd-review-security references. Keeps SKILL.md focused on activation/execution contract.

### Decision: Narrative secure-design only (no scoring)

**Choice**: `secure-design.md` uses prose rule structure per surface area
**Alternatives considered**: Control tables, SEC-* matrices, Source ID rows
**Rationale**: Scoring authority belongs exclusively to sdd-review and sdd-review-security. project-conventions guides; review phases judge.

## Data Flow

```
Orchestrator
    │
    ├── resolves skill from registry
    │
    ▼
Phase Agent (e.g., sdd-design)
    │
    ├── reads skills/project-conventions/SKILL.md
    │       │
    │       ├── Step 1-4: reads references/{architecture,code-style,testing,review-readiness}.md
    │       │
    │       └── Step 5 (conditional): reads references/secure-design.md
    │
    ▼
Output: design.md with ## Implementation Quality Constraints section
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `skills/project-conventions/SKILL.md` | Create | Main skill activation contract with frontmatter, hard rules, decision gates, execution steps, output contract, references |
| `skills/project-conventions/references/architecture.md` | Create | SDD executor model, shared contracts, boundary statement, supplemental skill injection |
| `skills/project-conventions/references/code-style.md` | Create | English artifacts, YAML frontmatter schema, section heading order, Markdown table conventions |
| `skills/project-conventions/references/testing.md` | Create | strict_tdd: false, no automated runner, verification via test-design.md, update trigger |
| `skills/project-conventions/references/review-readiness.md` | Create | 400-line budget, auto-chain, stacked-to-main, chained-pr reference, PR conventions |
| `skills/project-conventions/references/secure-design.md` | Create | 9 surface areas with narrative rules (6-part structure), no SEC-* matrices |
| `AGENTS.md` | Modify | Add `## Project Skills` section registering project-conventions |

## Interfaces / Contracts

### SKILL.md Frontmatter Schema

```yaml
---
name: project-conventions
description: "Project-specific conventions for SDD phases. Trigger: supplemental skill loaded when project conventions are needed."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
  delegate_only: true
---
```

### SKILL.md Section Structure

1. **Activation Contract** — trigger conditions (SDD phase needs conventions, change touches architecture/naming/testing/security/review-readiness)
2. **Hard Rules** — 6 prohibitions per SPEC-002
3. **Decision Gates** — 4-row table per SPEC-003
4. **Execution Steps** — 5 numbered steps per SPEC-004
5. **Output Contract** — structured output fields per SPEC-005
6. **References** — links to all 5 reference files

### sdd-design Integration Pattern

When project-conventions is injected as supplemental skill into sdd-design:

- The executor adds `## Implementation Quality Constraints` after `## Testing Strategy` in design.md
- The section lists applied project conventions (code-style, testing, review-readiness)
- An optional `## Secure Design Constraints` subsection uses a table with columns: Area | Project Secure Convention | Design Commitment | Evidence Owner
- Minimum: at least the rows applicable to the change's touched surfaces
- This is GUIDANCE — sdd-review retains sole judgment authority

### AGENTS.md Addition

```markdown
## Project Skills

- `project-conventions` — supplemental skill loaded by SDD phases when project conventions are needed (architecture, code-style, testing, review-readiness, secure-design). NOT a required SDD DAG phase.
```

## Operational Considerations

No aplica. This change creates static Markdown documentation files (skill instructions and reference content). No runtime behavior, no services, no monitoring, no administration operations.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Manual verification | SKILL.md frontmatter schema validity | Verify 4 required YAML fields present (name, description, license, metadata) |
| Manual verification | Section order compliance | Verify sections appear in prescribed order |
| Manual verification | Hard rules completeness | Verify all 6 prohibitions from SPEC-002 present |
| Manual verification | Reference file presence | Verify all 5 files exist and are non-empty |
| Manual verification | No SEC-* matrices in any file | Grep for SEC-*, Source ID, scoring table patterns |
| Manual verification | AGENTS.md update | Verify `## Project Skills` section present with correct content |

No automated test runner is available (strict_tdd: false).

## Migration / Rollout

No migration required. Pure additive change — new files only, plus an append to AGENTS.md.

## Open Questions

None. All requirements are fully specified.

## Secure Development Design

### Classification and Changed Surface

**Classification**: no-impact

**Changed artifacts**: New Markdown files (`skills/project-conventions/SKILL.md`, 5 reference files) and an append to `AGENTS.md`. All are static documentation/instruction content.

**Touched runtime surfaces**: None. These files are read by LLM agents at prompt-construction time; they are not executed, do not process user input, do not store data, do not authenticate, and do not communicate with external services.

**Untouched runtime surfaces**: Authentication, sessions, databases, file I/O, APIs, logging, error handling — none of these exist in this repository (it's an instruction/skill distribution repo, not an application runtime).

**Why no security category applies**: The change produces read-only Markdown text consumed by AI agents. No inputs are processed, no outputs are served to users, no data is stored, no secrets are handled at runtime, no external APIs are called. The only security-adjacent content is the `secure-design.md` reference file itself, which is a guidance document containing no actual secrets or sensitive data.

**Omitted categories**: All 8 security categories (authentication, sessions, sensitive-data-pan, secrets, permissions-access-control, files, database-access, sensitive-logging) are omitted as reviewable omissions for canonical `review-security-report.json` validation.

### Safe Evidence Policy

Despite no-impact classification, the following design-time commitments apply to the content of the produced files:

- **No secrets or tokens** in any reference file — `secure-design.md` uses placeholder values only (e.g., `<YOUR_SECRET_HERE>`) when showing examples
- **No PII** in any file — all content is generic instruction text
- **Output contract** in SKILL.md specifies path/summary references for evidence, never raw payloads
- **Exception policy**: If a reference file ever needs to demonstrate a pattern involving secrets, it MUST use obviously-fake placeholder values

### Boundary Confirmation

- project-conventions is read-only guidance; it does not process runtime data
- sdd-review-security retains authority over all formal security validation via `review-security-report.json`
- No phase other than sdd-review produces or modifies the 96-control matrix evaluation

### Exception and Evidence Policy

No exceptions are planned. No sensitive evidence is produced by this change.

## Implementation Quality Constraints

**Project Conventions Applied:**
- English artifacts only (all generated files in English)
- YAML frontmatter with name, description, license, metadata
- Section order: Activation Contract → Hard Rules → Decision Gates → Execution Steps → Output Contract → References
- No automated test runner (strict_tdd: false) — verification via test-design.md
- Review budget: 400 lines, stacked-to-main delivery

**Secure Design Constraints:**

| Area | Project Secure Convention | Design Commitment | Evidence Owner |
|------|--------------------------|-------------------|----------------|
| Data handling | Artifacts are public documentation; no PII/secrets | Reference files contain no secrets, tokens, or PII; examples use placeholders | design/apply |
| Secrets/config | No hardcoded secrets in any SDD artifact | secure-design.md uses placeholder values only | apply/review-security |
| Logging/errors | No sensitive payloads in evidence | Output contract specifies path/summary refs only | apply/verify |
