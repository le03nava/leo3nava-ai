# Test Design: Technical Documentation Agent

## Overview

This change adds `sdd-technical-doc` as a manual, post-archive, archive-consuming utility for generating exactly one neutral professional Spanish Markdown technical document from archived SDD evidence. The repository is a Markdown instruction-contract repo with no runtime test runner, build command, linter, formatter, type checker, or coverage command configured, so validation is planned as static/manual artifact read-back and checklist evidence rather than executable tests.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/technical-documentation-agent/proposal.md` | Intent, scope, non-goals, required Spanish structure, reference restrictions, inventory enums, and success criteria. |
| Spec | `openspec/changes/technical-documentation-agent/specs/sdd-technical-documentation-workflow/spec.md` | Manual utility behavior, archive-only evidence, exact output count/language, document structure, marker, reference, and inventory requirements. |
| Spec | `openspec/changes/technical-documentation-agent/specs/sdd-execution-persistence-contracts/spec.md` | Non-DAG boundary, unchanged status/routing/verify/archive contracts, and archive-only persistence behavior. |
| Design | `openspec/changes/technical-documentation-agent/design.md` | Architecture decisions, file plan, contracts, operational considerations, rollout, and static/manual testing strategy. |
| Secure Development Design | `openspec/changes/technical-documentation-agent/design.md#secure-development-design` | Security-impacting changed-surface classification, sensitive data handling rules, file/output rules, logging/evidence safety rules, residual risks, and no planned exceptions. |
| Testing Capabilities | `openspec/config.yaml#testing` | Confirms no unit/integration/E2E runner, coverage command, linter, type checker, formatter, build command, or verify command is available. |
| Supplemental Skill | `C:\Users\leo3n\.config\opencode\skills\skill-creator\SKILL.md` | Frontmatter shape, concise LLM-first skill structure, asset placement, and registration expectations. |
| Supplemental Skill | `builtin:customize-opencode` | Applies because the change creates opencode agent/skill configuration; builtin content is injected by orchestration, with no readable file path provided. |

## Source ID Coverage Baseline

Corporate source-row coverage does not require a full Source ID matrix in this phase. The changed surface is instruction/template/documentation content for an archive-consuming manual utility, not application runtime code, authentication, database connectivity, or network services. Applicable narrative security rules are sensitive data handling, safe file/output handling, and sensitive logging/evidence safety from `design.md#secure-development-design`.

Validation therefore uses static/manual checks over the planned `SKILL.md`, Spanish template asset, thin agent prompt, README/AGENTS docs, and delta specs. Test-design does not require design YAML, schema fields, compact controls, Source ID matrices, machine-readable applicability fields, exhaustive `N/A` rows, or a full source-row catalog; exhaustive expansion and omitted-row validation remain owned by `review-security-report.md`. Runtime/build/lint/type/format/coverage tooling is unavailable and must be reported as unavailable evidence, never as passing evidence.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | Spec: Manual Post-Archive Utility; Design: Manual utility boundary | Confirm `sdd-technical-doc` is documented as manual, user-invocable, post-archive, and archive-consuming, not as an SDD phase. | static | mandatory | `skills/sdd-technical-doc/SKILL.md`, `agents/sdd/sdd-technical-doc.md`, `README.md`, `AGENTS.md`, and specs contain manual boundary wording. | Blocks if any artifact makes it required for active SDD flow. |
| TD-002 | Spec: Archive flow is unchanged; Persistence delta | Confirm no DAG/status/successor/dependency/verify/archive requirement is added for `sdd-technical-doc`. | static | mandatory | Static read-back of SDD phase docs/status references shows phase order unchanged and no required `sdd-technical-doc` token in status, verify, archive, or dependency contracts. | The utility may be listed only as a manual post-archive utility. |
| TD-003 | Proposal: In Scope; skill-creator guidance | Validate new skill frontmatter uses required fields and concise YAML-safe one-line description. | static | mandatory | `skills/sdd-technical-doc/SKILL.md` has `name`, one-line quoted `description`, `license`, `metadata.author`, and `metadata.version`. | Use `skill-creator` fallback because `docs/skill-style-guide.md` is unavailable. |
| TD-004 | Design: Template asset decision; skill-creator guidance | Validate `SKILL.md` remains concise and LLM-first, with long Spanish template content moved to `assets/technical-document-template.md`. | static | mandatory | `SKILL.md` uses imperative runtime sections and does not inline the full document template. | Target concise body; hard cap from supplemental skill is 1000 tokens unless project style later overrides. |
| TD-005 | Design: Agent prompt | Validate `agents/sdd/sdd-technical-doc.md` is a thin executor prompt that reads/uses the skill, performs the phase itself, and does not delegate to sub-agents. | static | mandatory | Agent prompt names the skill path, says not to delegate, and reinforces execution by the current agent. | Compare intent with existing thin manual utility pattern where useful, without modifying that utility. |
| TD-006 | Proposal/Design: Archive-only boundary | Confirm skill and agent refuse active-change generation, rerunning phases, status mutation, archive readiness mutation, and evidence backfill. | static | mandatory | Skill/agent instructions state archived evidence is the source of truth and generation must not mutate SDD state. | Includes final-document-only values not being copied back into SDD artifacts. |
| TD-007 | Spec: Manual generation succeeds | Confirm the utility contract produces exactly one Markdown technical documentation file in neutral professional Spanish. | static | mandatory | Skill output contract and template instructions require one `.md` output only and Spanish neutral professional content. | Multiple files, non-Markdown outputs, or mixed-language final document rules block. |
| TD-008 | Proposal: Required document structure; Spec: Mandatory Spanish Document Structure | Validate the Spanish template preserves the required section skeleton from `Identificacion/Identificación del Producto` through sections 1, 2, 4, 5, 6, 7, and 8. | manual | mandatory | Checklist evidence from `skills/sdd-technical-doc/assets/technical-document-template.md` against proposal lines 33-96. | Accent normalization is acceptable only if meaning and required structure are preserved. |
| TD-009 | Proposal/Spec: BI/POS conditional sections | Validate BI-only sections 2.2, 4.3, and 5 and POS-only sections 2.7 and 6 remain conditional and render `No aplica.` when not applicable. | manual | mandatory | Template and skill marker rules explicitly state conditional platform sections use `No aplica.` when archived evidence does not prove applicability. | Must not infer BI/POS platform from file names alone unless archived evidence proves it. |
| TD-010 | Proposal/Spec: Static placeholders | Validate sections 7 and 8 preserve static approval-signature and revision-control placeholders. | static | mandatory | Template contains static placeholder tables for `7 Firmas de aprobacion/aprobación` and `8 Control de revisiones`. | Placeholders may be static even when archive evidence lacks approvers/revisions. |
| TD-011 | Spec: Archived Evidence Only | Validate all inapplicable sections use exactly `No aplica.`. | static | mandatory | Skill and template instructions contain the exact literal marker `No aplica.` for inapplicable content. | Do not substitute variants such as `N/A` in the final Spanish document. |
| TD-012 | Spec: Evidence is missing; Design: Data Flow | Validate applicable but missing data uses exactly `Información no disponible en la evidencia archivada.` and is not inferred. | static | mandatory | Skill and template instructions contain the exact literal marker and non-invention rule for missing applicable data. | Applies to product, owner, runtime, database, security, integration, object, and reference details. |
| TD-013 | Proposal/Spec: References and Object Inventory | Validate final references are restricted to FCTI, story, `.pks`, `.pkb`, and `.sql` package/table sources created or modified by development. | static | mandatory | Skill/template reference rules list only allowed functional/development source categories. | Archive evidence can be cited internally as safe evidence, but final document references are restricted. |
| TD-014 | Proposal/Spec: Installer script exclusion | Validate installer scripts are excluded from final references even if present in archived evidence. | static | mandatory | Skill/template contain an explicit installer-script exclusion rule for final references. | Examples must not include installer scripts as acceptable final references. |
| TD-015 | Proposal/Spec: Inventory table columns | Validate section 1.2 inventory table has required columns `Tipo`, `Esquema.Objeto`, `Operacion`, and `Descripcion funcional`. | static | mandatory | Template includes the required pipe table headers and skill requires one row per touched development object when evidence exists. | Header spelling should follow the Spanish template contract; accents may be neutralized only if consistent with final artifact style. |
| TD-016 | Proposal/Spec: Inventory enum constraints | Validate `Tipo` and `Operacion` allowed values are explicitly constrained and unsupported values are corrected only from archive evidence or marked unavailable. | static | mandatory | Skill/template list `Tipo` enum values and `Operacion` enum values exactly as specified. | Invalid enum mappings must not be invented. |
| TD-017 | Secure Development Design: Sensitive Data Handling Rules | Validate skill/template/examples do not contain or encourage copying secrets, credentials, tokens, private keys, connection strings, PAN, PII, confidential client values, raw payloads, production hostnames/IPs/ports/SID values, full ID lists, or generated bytes. | static | mandatory | Changed files contain safe placeholders and explicit restricted-data refusal rules only. | Review-security owns deeper source-row expansion; this case plans safe-evidence coverage. |
| TD-018 | Secure Development Design: Sensitive Data Handling Rules | Validate final-document-only values supplied by the user are allowed only in the generated final document and are not backfilled into SDD artifacts, examples, tests, or archived design evidence. | static | mandatory | Skill includes final-document-only boundary and backfill prohibition. | Blocks if instructions tell agents to update archived SDD evidence with user-supplied final values. |
| TD-019 | Secure Development Design: Files and Generated Document Rules | Validate safe path/output guidance: exactly one Markdown output, read existing content before overwrite decisions, prefer archive-local safe paths, and avoid silent archive mutation. | static | mandatory | Skill output section contains path-safety and overwrite guidance. | If no output path is supplied, inline return or a safe archive-local suggestion is acceptable. |
| TD-020 | Secure Development Design: Sensitive Logging and Evidence Safety Rules | Validate command outputs, summaries, examples, and evidence references use paths/sections/sanitized descriptions rather than raw restricted values. | static | mandatory | Skill and agent instruct sanitized evidence summaries only. | Missing markers are safe evidence and should not be treated as leaks. |
| TD-021 | Proposal: README/AGENTS docs | Validate `README.md` and `AGENTS.md` list `sdd-technical-doc` as a manual post-archive utility without altering SDD phase order. | static | mandatory | Documentation section beside `sdd-operational-doc`; no required-DAG phrasing. | Project instruction says manual post-archive utilities must not become required DAG phases. |
| TD-022 | Design: Testing Strategy; OpenSpec testing config | Validate verify evidence reports unavailable runtime/build/lint/type/format/coverage tooling explicitly. | manual | mandatory | Verify phase checklist cites `openspec/config.yaml#testing` and reports unavailable tools as unavailable, not passing. | No executable command is required or available for this Markdown repo. |
| TD-023 | Proposal: No generated technical document in this change | Confirm implementation does not generate the final technical documentation deliverable during this SDD change. | static | non-mandatory | Changed files include skill/template/agent/docs/spec artifacts only, not an actual customer technical-document output for this change. | This is out of scope per proposal; accidental generated deliverables should be removed unless explicitly requested later. |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Manual post-archive boundary | Confirm utility reads archived evidence and does not mutate SDD state, status, archive readiness, verify gates, or phase order. | static | `SKILL.md`, agent prompt, README, AGENTS, and specs use manual/archive-only wording. | No runtime or integration runner is available. |
| Archive source resolution | Confirm skill resolves archived OpenSpec paths or archived backend refs and refuses active-change inference. | static | Skill execution steps reference archived change folders/refs and archived evidence only. | No executable resolver test is available. |
| Final document output | Confirm output is one Markdown file in neutral professional Spanish, or inline when no target path is supplied. | static | Skill output contract and Spanish template asset. | No formatter or Markdown linter is available. |
| Missing/inapplicable markers | Confirm exact markers `No aplica.` and `Información no disponible en la evidencia archivada.` are used for the correct cases. | static | Skill/template marker rules and template placeholders. | No snapshot/golden-file runner is available. |
| Reference filtering | Confirm final references exclude installer scripts and include only FCTI/story/package/table source files allowed by the spec. | manual | Static checklist against skill/template rules. | No automated file classifier exists. |
| Inventory validation | Confirm object inventory columns and enum constraints are documented and invalid values are not invented. | manual | Static checklist against template headers and skill validation rules. | No schema validator exists. |
| Restricted data boundary | Confirm evidence and examples are sanitized and final-document-only values are not backfilled into SDD artifacts. | static | Changed-file read-back shows no raw restricted example values and explicit safe-evidence instructions. | No secret scanner is configured; review-security should perform manual/static inspection. |
| Documentation discoverability | Confirm README/AGENTS present the utility next to manual post-archive utilities and do not change the canonical SDD phase list. | static | Documentation diff/read-back evidence. | No docs lint command is available. |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-SDH-001` | Extract only facts present in archived evidence or explicit user-provided final-document context; do not invent product, owner, runtime, database, security, integration, object, or reference details. | Yes | TD-006, TD-011, TD-012, TD-018 | covered | None |
| `SEC-SDH-002` | Do not copy raw secrets, credentials, tokens, private keys, connection strings, PAN, PII, confidential client values, raw payloads, production hostnames/IPs/ports/SID values, full ID lists, or generated bytes into ordinary SDD artifacts or examples. | Yes | TD-017, TD-020 | covered | None |
| `SEC-FILE-001` | Produce exactly one Spanish Markdown technical document per invocation. | Yes | TD-007, TD-019 | covered | None |
| `SEC-FILE-002` | Read existing content before overwrite decisions and constrain output to a safe Markdown path, preferably under the archived change folder. | Yes | TD-019 | covered | None |
| `SEC-FILE-003` | Exclude installer scripts from final references even when archived evidence contains them. | Yes | TD-013, TD-014 | covered | None |
| `SEC-EVID-001` | Summarize evidence by path, section, or sanitized description only; do not include raw sensitive payloads or restricted operational/client values in outputs, examples, or summaries. | Yes | TD-017, TD-020 | covered | None |
| `SEC-EVID-002` | Treat `No aplica.` and `Información no disponible en la evidencia archivada.` as safe markers for inapplicable or missing applicable content. | Yes | TD-011, TD-012 | covered | None |
| `SEC-EVID-003` | Keep final-document-only values supplied by the user out of SDD artifacts, tests, fixtures, examples, and archived design artifacts. | Yes | TD-018 | covered | None |

## No-Impact Assessment

Not applicable. This change has behavior and testability impact because it adds a new manual archive-consuming utility, Spanish template contract, thin agent prompt, documentation updates, specs, and security-sensitive evidence-handling rules.

## Evidence Expectations

- Mandatory cases require implementation, read-back validation, static/manual evidence, or a justified skip before verification can pass.
- Non-mandatory cases should be reported as warnings when uncovered, but they do not block verification by themselves.
- Security validation evidence should cite embedded `design.md#secure-development-design` narrative rules, owner phase, and planned static/manual evidence.
- Applicable narrative category rules require planned safe evidence. Exhaustive `N/A` source coverage and missed-applicable validation remain owned by `review-security-report.md`.
- Warning-only source coverage must be preserved with expected observation evidence and may proceed only when mandatory evidence is complete.
- Test-design consumes narrative design rules only and MUST NOT require design YAML, schema fields, compact matrices, Source ID matrices, machine-readable applicability fields, or the full 155-row Source ID matrix; exhaustive row materialization belongs only to `review-security-report.md`.
- No-impact routing is valid only when justified by changed-surface classification inside mandatory `design.md#secure-development-design`; absence of standalone `security-design.md` is not a blocker for new changes.
- Runtime tests, build commands, linters, type checkers, formatters, and coverage commands are unavailable for this repository and must be reported as unavailable evidence, not treated as passed checks.

## Open Questions

None.
