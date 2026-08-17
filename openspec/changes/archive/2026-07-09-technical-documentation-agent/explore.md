## Exploration: technical-documentation-agent

### Current State

The repository already has a matching manual post-archive pattern in `sdd-operational-doc`:

- `skills/sdd-operational-doc/SKILL.md` defines a user-invocable manual utility that consumes archived SDD evidence and does not participate in the SDD DAG.
- `agents/sdd/sdd-operational-doc.md` is a thin executor prompt that loads the installed skill, forbids delegation, and preserves the post-archive boundary.
- `skills/sdd-operational-doc/assets/operational-document-template.md` provides the mandatory Spanish output skeleton for the generated operational document.
- `AGENTS.md` and `README.md` explicitly document that `sdd-operational-doc` is manual and not a required SDD phase.

The OpenSpec project config describes this repository as an AI agent/skill distribution with no executable test runner. Existing SDD rules require generated technical artifacts to default to English unless the target artifact explicitly requires Spanish. For this change, the skill/agent instructions should remain English, while the generated final technical documentation Markdown must be neutral professional Spanish by explicit requirement.

### Affected Areas

- `skills/sdd-technical-doc/SKILL.md` — likely new user-invocable skill defining the archive-only technical documentation workflow, evidence rules, and output contract.
- `skills/sdd-technical-doc/assets/technical-document-template.md` — likely new mandatory Spanish template containing the exact user-provided section structure, BI/POS conditional sections, inventory table, signatures, and revision-control placeholders.
- `agents/sdd/sdd-technical-doc.md` — likely new manual executor prompt, mirroring the `sdd-operational-doc` agent boundary.
- `AGENTS.md` — add one line documenting that this utility is manual/post-archive and not a required SDD DAG phase.
- `README.md` — add the new utility under manual post-archive utilities.
- `.atl/skill-registry.md` or generated registry artifacts — may need refresh after implementation if the project uses the registry for skill discovery.
- `openspec/changes/technical-documentation-agent/**` — normal SDD artifacts for this change.

### Approaches

1. **New sibling manual utility skill and agent** — create `sdd-technical-doc` as a peer of `sdd-operational-doc`.
   - Pros: Reuses the proven manual post-archive shape; keeps DAG untouched; isolates technical-document-specific references, inventory rules, and Spanish template.
   - Cons: Requires a new template asset and careful naming to avoid confusion with SDD design/technical design phases.
   - Effort: Medium.

2. **Extend `sdd-operational-doc` with a technical-document mode** — add a second document type to the existing operational utility.
   - Pros: Fewer top-level files.
   - Cons: Mixes two document contracts with different section structures and evidence filters; raises risk of operational rules leaking into technical documentation or vice versa.
   - Effort: Medium with higher maintenance risk.

3. **Add a required SDD phase** — integrate technical documentation generation into the SDD flow.
   - Pros: Could enforce documentation generation automatically.
   - Cons: Directly violates the user requirement and existing project instruction: this must be manual/post-archive and must not modify DAG/flows.
   - Effort: High and not acceptable.

### Recommendation

Use Approach 1: a new sibling manual utility, likely named `sdd-technical-doc`, with a thin user-invocable agent and one template asset. This preserves the existing SDD DAG, keeps archive consumption explicit, and avoids overloading `sdd-operational-doc` with a second document contract.

The new skill should require the executor to:

- Resolve an already archived change by name or archive folder path.
- Read only archived SDD evidence from the selected backend, following the persistence contract and hybrid conflict policy.
- Generate exactly one `.md` technical documentation file in neutral professional Spanish.
- Use the user-provided section structure exactly, including BI/POS conditional sections and static placeholders for signatures/revision control.
- Write `No aplica.` when a section does not apply.
- Write a clear unavailable-information marker when required information is not present in archived evidence, e.g. `Informacion no disponible en la evidencia archivada.`
- Never invent functional behavior, owners, object operations, dependencies, scripts, SQL details, or business references.
- Restrict references to functional documents and development sources only: FCTI, story, and `.pks` / `.pkb` / `.sql` package/table sources created or modified.
- Explicitly exclude installer scripts from references, even if they appear in archived evidence.
- Always include the mandatory inventory table under `Alcance` with columns `Tipo`, `Esquema.Objeto`, `Operacion`, and `Descripcion funcional`.
- Validate inventory enum values: `Tipo` must be one of `TABLE`, `INDEX`, `PACKAGE SPEC`, `PACKAGE BODY`, `PROCEDURE`, `FUNCTION`, `TRIGGER`, `JOB ControlM`, `VIEW`, `SEQUENCE`, `GRANT`; `Operacion` must be one of `CREATE`, `ALTER`, `REPLACE`, `DROP`.

### Risks

- The exact technical-document section structure was not present in this launch envelope. Proposal/spec should capture the template verbatim before implementation.
- Archive-only evidence can be incomplete. The skill must prefer explicit unavailable-information text over inference.
- Installer scripts may contain tempting object and deployment details, but the requirement forbids referencing them. The skill may inspect archived evidence to understand changed files, but final references must exclude installer-script refs.
- `sdd-technical-doc` naming should avoid implying a normal SDD design artifact or required phase.
- The injected `customize-opencode` supplemental skill is built-in and has no readable `SKILL.md` file path in this executor context; implementation should still respect opencode agent/skill conventions verified from existing local files.

### Ready for Proposal

Yes. The proposal should define a new manual post-archive utility and explicitly include the user-provided Spanish technical documentation template as the output contract. It should keep DAG/flow changes out of scope.
