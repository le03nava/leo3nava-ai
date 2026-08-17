# Design: Technical Documentation Agent

## Technical Approach

Create `sdd-technical-doc` as a manual, user-invocable, post-archive sibling utility to `sdd-operational-doc`. The utility will consume archived SDD evidence, apply a mandatory Spanish Markdown template, and produce exactly one technical documentation `.md` file. It will not participate in the active SDD DAG, status routing, verify gates, archive gates, or phase dependencies.

This maps to `sdd-technical-documentation-workflow` by making archived evidence the only source of truth and by enforcing explicit missing-data markers. It maps to `sdd-execution-persistence-contracts` by documenting the utility as archive-consuming and non-required across the skill, thin agent prompt, README, AGENTS, and specs.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Manual utility boundary | Implement `sdd-technical-doc` as a user-invocable skill plus thin agent prompt. | Add a new SDD phase after archive. | A new DAG phase would violate the proposal and make archive completion depend on a final deliverable that must remain optional/manual. |
| Evidence source | Read archived SDD artifacts only. | Read active change artifacts or infer from current code. | The requested document is a final archive-derived deliverable; active files may have drifted and inference risks fabricated product/runtime/security details. |
| Output language split | Keep skill/agent instructions in English and the generated template/document in neutral professional Spanish. | Write all artifacts in Spanish. | Repository skill contracts default to English, while the final technical-document artifact is explicitly required in Spanish. |
| Template asset | Store the required Spanish skeleton in `skills/sdd-technical-doc/assets/technical-document-template.md`. | Embed the full template in `SKILL.md`. | The template is long and output-specific; keeping it in `assets/` follows skill-creator guidance and keeps `SKILL.md` concise. |
| Validation style | Use deterministic instruction-level validation for references, inventory enums, markers, and section presence. | Depend on a runtime test framework. | This Markdown instruction repo has no package manifest, runner, linter, formatter, or type checker configured. |

## Data Flow

```text
Archived SDD change
  ├─ proposal/spec/design/test-design/tasks/review/security-review/verify/archive evidence
  └─ optional archived source/evidence refs
        │
        ▼
sdd-technical-doc manual utility
  ├─ resolve archive backend/path
  ├─ read Spanish template asset
  ├─ extract only review-safe archived facts
  ├─ filter final references
  ├─ validate object inventory enums
  └─ apply missing/inapplicable markers
        │
        ▼
One Spanish Markdown technical document
```

When evidence proves a section is out of scope, the generated section writes exactly `No aplica.`. When the section applies but archived evidence lacks a required value, the generated section writes `Información no disponible en la evidencia archivada.` and does not infer or fabricate the value.

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/sdd-technical-doc/SKILL.md` | Create | User-invocable archive-consuming skill with activation contract, evidence rules, marker rules, reference filtering, enum validation, output contract, and safe-evidence boundaries. |
| `skills/sdd-technical-doc/assets/technical-document-template.md` | Create | Mandatory Spanish technical-document skeleton preserving all requested sections, conditional BI/POS sections, static approval/revision placeholders, and inventory table headers. |
| `agents/sdd/sdd-technical-doc.md` | Create | Thin executor prompt that reads the skill, does not delegate, and reinforces the manual post-archive boundary. |
| `AGENTS.md` | Modify | Add `sdd-technical-doc` beside `sdd-operational-doc` as a manual post-archive utility, not a required SDD DAG phase. |
| `README.md` | Modify | List `sdd-technical-doc` under manual post-archive utilities and keep phase order unchanged. |
| `openspec/specs/sdd-technical-documentation-workflow/spec.md` | Create during archive | Source spec synchronized from the change delta to define the new capability. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modify during archive | Source spec synchronized from the delta to preserve the non-DAG manual boundary. |

## Interfaces / Contracts

The skill will accept an archived change name or archive folder path, artifact store mode when known, and an optional target output path. It will read OpenSpec archives from `openspec/changes/archive/YYYY-MM-DD-{change-name}/` when using file-backed evidence, or archived Engram/hybrid refs according to the shared persistence contract when applicable.

The generated document contract is one Markdown file in neutral professional Spanish. Final references are restricted to functional/development sources: FCTI, story, and `.pks`, `.pkb`, or `.sql` package/table sources created or modified by the development. Installer scripts must be excluded from final references even if archived evidence contains them.

Section `1.2 Alcance` must include one object inventory row per object touched by development. `Tipo` is limited to `TABLE`, `INDEX`, `PACKAGE SPEC`, `PACKAGE BODY`, `PROCEDURE`, `FUNCTION`, `TRIGGER`, `JOB ControlM`, `VIEW`, `SEQUENCE`, or `GRANT`. `Operacion` is limited to `CREATE`, `ALTER`, `REPLACE`, or `DROP`. Rows may be corrected only from archived evidence; otherwise invalid or unsupported values use `Información no disponible en la evidencia archivada.` rather than invented mappings.

## Operational Considerations

### Strategy

The utility is operationally passive: it reads archived evidence and writes or returns one Markdown document only when manually invoked. It must not mutate SDD state, rerun phases, modify archive readiness, or backfill missing SDD evidence. If no output path is supplied, it should return the document inline or suggest a safe path under the archive folder.

### Evidence Plan

| Category | Expected safe evidence | Owner phase | Status |
| --- | --- | --- | --- |
| Manual boundary | Skill, thin agent, README, and AGENTS state that the utility is post-archive/manual only. | apply/review | planned |
| Archive source resolution | Paths or artifact refs to archived SDD evidence; no raw restricted values. | apply/review | planned |
| Final document output | Generated `.md` path or inline output summary. | utility execution | planned |
| Missing/inapplicable data | Literal `No aplica.` or `Información no disponible en la evidencia archivada.` markers in generated content. | utility execution/review | planned |

### Restricted Data Boundary

Restricted operational/client values stay out of ordinary SDD evidence and implementation examples. The skill may summarize archived evidence and cite paths/sections, but it must not leak credentials, tokens, connection strings, private operational identifiers, PAN/PII, raw payloads, or confidential client values. Final-document-only values are not backfilled into SDD artifacts.

### Unresolved Gaps

- None blocking for design. Actual generated documents may contain `Información no disponible en la evidencia archivada.` when archive evidence lacks required details.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static/manual | Skill frontmatter and concise LLM-first structure. | Inspect against `skill-creator` fallback rules because `docs/skill-style-guide.md` is not present. |
| Static/manual | Template preserves required Spanish sections and table headers. | Manual checklist against proposal/spec structure. |
| Static/manual | Agent prompt remains thin and non-delegating. | Compare with `agents/sdd/sdd-operational-doc.md` pattern. |
| Static/manual | README/AGENTS do not alter phase order. | Confirm `sdd-technical-doc` appears only under manual post-archive utilities. |
| Static/manual | Reference filtering and inventory enum rules are explicit. | Inspect `SKILL.md` and template instructions for installer exclusion and allowed enum values. |

No automated unit, integration, E2E, coverage, linter, formatter, or type-check command is available in `openspec/config.yaml`. Verification must therefore report unavailable runtime tests explicitly and rely on artifact read-back plus manual/static checklist evidence.

## Migration / Rollout

No migration required. Rollout is additive: install/sync the new skill and agent through existing repository sync scripts. Existing SDD DAG state and archived changes require no data migration.

## Open Questions

None.

## Secure Development Design

### Classification and Changed Surface

Classification: security-impacting. The change adds an archive-consuming documentation utility that reads SDD evidence and emits a final Markdown document. It does not add application runtime code, authentication/session behavior, database access, network services, or deployment gates. The changed surface is documentation-generation instructions, template content, a thin opencode agent prompt, README/AGENTS documentation, and source spec synchronization at archive time.

The applicable security concerns are sensitive data handling, file/output handling, and sensitive logging/evidence safety. Secrets and access-control categories are not designed as runtime integrations here because no credential store, auth workflow, permission model, or service call is added; omitted categories remain reviewable by `sdd-review-security` rather than proven by design-time `N/A` rows.

### Sensitive Data Handling Rules

- The utility must extract only facts present in archived evidence or explicit user-provided final-document context.
- It must not invent product, owner, runtime, database, security, integration, object, or reference details.
- It must not copy raw secrets, credentials, tokens, private keys, connection strings, PAN, PII, confidential client values, raw payloads, production hostnames/IPs/ports/SID values, full ID lists, or generated bytes from archived evidence into ordinary SDD artifacts or examples.
- Evidence owners: apply must implement the rule in the skill/template; review-security must validate that generated instructions preserve the safe-evidence boundary; archive must preserve the final specs and reports.
- Expected evidence: changed-file refs to `skills/sdd-technical-doc/SKILL.md`, the Spanish template asset, and docs showing explicit marker and non-invention rules.
- Residual risk: archived evidence may contain unsafe values. The mitigation is instruction-level refusal to copy restricted values unless the user explicitly supplies final-document-only values for the generated document, and those values must not be backfilled.

### Files and Generated Document Rules

- The utility may produce exactly one Spanish Markdown technical document per invocation.
- If a target path is provided, implementation must read existing content first before overwrite decisions and keep output constrained to a safe Markdown path, preferably under the archived change folder.
- Installer scripts must be excluded from final references, even when present in archived evidence.
- Evidence owners: apply implements output-path and reference-filtering instructions; review validates no DAG/status/archive artifacts are changed as part of generation; verify confirms manual/static checks.
- Residual risk: a user-provided target path could point outside the archive. The skill must treat path choice as explicit user intent and avoid silent archive mutation beyond writing the requested final document.

### Sensitive Logging and Evidence Safety Rules

- The skill and agent must summarize evidence sources by path, section, or sanitized description only.
- Command outputs, review summaries, and generated-document summaries must not include raw sensitive payloads or restricted operational/client values.
- Missing applicable content must use `Información no disponible en la evidencia archivada.`; inapplicable content must use `No aplica.`. These markers are safe evidence and must not be treated as leaked data.
- Evidence owners: test-design plans manual/static checks; apply records changed files and safe summaries; review-security validates unsafe-evidence absence; verify carries any warning forward.
- Residual risk: the generated final document may intentionally contain final-document-only values supplied by the user. Those values belong only to the final manual document and must never be copied back to SDD evidence, tests, fixtures, examples, or archived design artifacts.

### Exception and Evidence Policy

No exceptions are planned. Any future exception to safe-evidence handling or required documentation markers must include an approver, approval date, accepted-risk rationale, mitigation or follow-up, and the exact evidence gap before it can be considered archive-ready.
