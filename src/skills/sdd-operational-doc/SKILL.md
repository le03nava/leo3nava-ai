---
name: sdd-operational-doc
description: "Trigger: operational doc, documento operativo, post-archive SDD. Generate an operations document from archived SDD evidence."
disable-model-invocation: true
user-invocable: true
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
---

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`. The generated operational document defaults to neutral professional Spanish because this artifact is a Spanish operational handoff document.

## Activation Contract

Run manually after an SDD change has been archived and the user asks for an operational document, operations manual, production handoff, or `sdd-operational-doc`.

This utility is manual and post-archive only. It is not a required SDD DAG phase, does not create archive readiness, and must not modify SDD DAG state, rerun phases, re-archive, backfill SDD evidence, or claim missing evidence as fact.

## Inputs

Accept from the user or launch context:
- Archived change name, or an archive folder path.
- Artifact store mode when known: `openspec | engram`.
- Optional explicit product metadata: product, application owner, platform, IT project lead, developer, support roles, retention policy, and target output path.
- Optional explicit final-document operational values. These values may be used only in the generated operational document and must never be copied back into SDD evidence, archive reports, review reports, verify reports, examples, or tasks.

## Evidence Sources

Use the selected backend through `skills/_shared/persistence-contract.md`.

Archived operational evidence is useful input when present, but it is not required for generation. Before drafting any section, read archived operational evidence/status, refs, unresolved gaps, warning carry-forward, and manual-document handoff notes from archive artifacts when present. Use proposal, specs, design, test-design, tasks, review, review-security, and verify artifacts only to support or clarify the archived handoff, not to invent missing operational facts.

| Mode | Read from |
| --- | --- |
| `openspec` | `openspec/changes/archive/YYYY-MM-DD-{change-name}/` |
| `engram` | `sdd/{change-name}/archive-report` and linked artifact observations |

When reading an OpenSpec archive, inspect available `proposal.md`, `specs/**/spec.md`, `design.md`, `test-design.md`, `tasks.md`, canonical `review-report.json` plus derived `review-report.md`, canonical `review-security-report.json` plus derived `review-security-report.md`, `verify-report.md`, and any archive report/evidence files.

## Operational Evidence Boundary

| Data class | Allowed in SDD evidence/archive sources | Allowed in final operational document |
| --- | --- | --- |
| Safe summaries, artifact paths, section anchors, evidence states, unresolved gaps | Yes | Yes |
| Production hostnames, IPs, ports, SID/service names, credentials, tokens, sensitive payloads, full ID lists, generated file bytes, PAN, PII, confidential values, or equivalent restricted operational data | No | Only when explicitly user-provided for this final document |

If final restricted values are provided by the user, include them only in the generated document section that needs them. Do not update, rewrite, or suggest updates to SDD artifacts with those values. If values are missing or unsafe to disclose, write the exact unresolved marker `Pendiente de confirmar:` followed by the missing item.

## Hard Rules

- Use `assets/operational-document-template.md` as the mandatory output structure.
- Keep every numbered section from 1 through 9, including sub-sections.
- If a section does not apply, write exactly `No aplica.` in that section unless the template requires a more specific sentence.
- Do not invent owners, systems, SQL, jobs, endpoints, schedules, backup commands, retention windows, thresholds, log paths, escalation roles, environment identifiers, generated files, payloads, or production values.
- If required operational evidence is absent but the section concept applies, write `Pendiente de confirmar:` followed by the missing evidence.
- Diagrams R1, R2, R3, and R4 are mandatory. Use Mermaid diagrams. If evidence is incomplete, include a conservative diagram with explicit `Pendiente de confirmar:` nodes.
- Section 4 must describe execution of the product/entrypoint. Do not write a detailed installation procedure.
- Section 4.1 writes exactly `No aplica.` unless the change is BI, ETL, ODI, or WebService load related.
- Section 5 writes exactly `No aplica.` unless UI, reports, dashboards, or frontend assets are touched.
- Section 6.4 R4 may use only `Escalar a Soporte` as an escalation leaf.
- Backup, quarantine, restart, monitoring, and reprocess instructions must be operationally concrete or explicitly marked pending.
- Preserve safe-evidence boundaries: summarize review/security evidence and link/cite refs; do not copy full review matrices.
- Never copy production hostnames, IPs, ports, SID/service names, credentials, tokens, sensitive payloads, full ID lists, generated file bytes, or equivalent restricted operational data from SDD evidence or examples. Final restricted values require explicit user input for the final document only.

## Decision Gates

| Condition | Action |
| --- | --- |
| Archived change cannot be found | Stop and ask for the archive path or change name. |
| Archived operational evidence is absent | Continue as a marked draft. Use `No aplica.` for inapplicable absent content and `Pendiente de confirmar:` for missing applicable content; do not backfill SDD evidence or treat the draft as archive readiness. |
| Hybrid backends differ materially | Stop and ask which backend is authoritative. |
| Product metadata is missing | Generate with `Pendiente de confirmar:` fields; do not block. |
| Operational commands are missing | Do not fabricate; include pending placeholders with the exact missing command type. |
| User provides final restricted values | Use them only in the final operational document output; do not persist them to SDD artifacts or examples. |
| User provided target path | Write the document there after reading existing file first if it exists. |
| No target path provided | Return the document inline and suggest a path under the archived change folder. |

## Execution Steps

1. Resolve the archived change and selected backend.
2. Read the template at `skills/sdd-operational-doc/assets/operational-document-template.md`.
3. Read archived operational evidence when present: operational status, refs, unresolved gaps, warnings, and manual-document handoff notes. Missing archived operational evidence is not a blocker.
4. Read archived SDD artifacts only as supporting context and collect only safe evidence needed for the template.
5. Derive entrypoints, touched objects, systems, schedules, UI impact, technologies, monitoring signals, logs, backup/reprocess operations, and support gaps from archived evidence. If evidence is absent and the content is inapplicable, write `No aplica.`; if evidence is absent but the content appears applicable, use `Pendiente de confirmar:` instead of inference.
6. Generate the complete document in neutral professional Spanish.
7. If writing to disk, use a safe Markdown path such as `openspec/changes/archive/YYYY-MM-DD-{change-name}/operational-document.md` unless the user provided another path.
8. Return a concise summary with output location, sections marked `No aplica.`, and fields marked `Pendiente de confirmar:`.

## Output Contract

Return:
- Generated document path, or `inline only`.
- Evidence sources read.
- Count/list of `Pendiente de confirmar:` items.
- Confirmation that all sections 1 through 9 and diagrams R1-R4 are present.

## References

- `skills/sdd-operational-doc/assets/operational-document-template.md` — mandatory document skeleton.
- `skills/_shared/persistence-contract.md` — backend resolution policy.
- `skills/_shared/openspec-convention.md` — OpenSpec archive path conventions.
- `skills/_shared/language-domain-contract.md` — artifact language rules.
