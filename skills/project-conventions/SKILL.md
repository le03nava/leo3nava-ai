---
name: project-conventions
description: >
  Load when a SDD phase or agent needs real project conventions for design, test-design,
  tasks, apply, or review-readiness. Trigger: supplemental skill injected by orchestrator
  when the change touches architecture, naming, testing, security surfaces, or delivery.
license: MIT
metadata:
  author: leo3nava_ai
  version: 1.0.0
---

## Activation Contract

Load this skill when:

- Any SDD phase (`sdd-design`, `sdd-test-design`, `sdd-tasks`, `sdd-apply`) needs real project conventions.
- The change touches architecture patterns, naming, testing infrastructure, security surfaces, or delivery conventions.
- It is loaded as a supplemental skill by the orchestrator via skill-resolver; it is not invoked directly by users.
- The change touches data, auth, permissions, logs, exports, files, DB, config/secrets, external APIs, or operational behavior; also load `references/secure-design.md`.

## Hard Rules

- Use only patterns proven by existing files, configuration, or docs in this repository. If a convention is unproven, mark it as `unknown` or `needs confirmation`.
- Never copy, materialize, or score the sdd-review 96-control matrix.
- Never create SEC-* matrices, Source ID rows, or formal scoring tables of any kind.
- Never invent test, lint, or build commands not present in the project configuration files.
- Evidence in output must remain safe: include file paths, section names, and summaries only; never include secrets, tokens, PII, raw logs, or sensitive payloads.
- `sdd-review-security` retains final authority over `review-security-report.json` and all formal security validation.

## Decision Gates

| Condition | Action |
|-----------|--------|
| Convention is known and proven by repo evidence | Apply it |
| Two or more contradicting conventions found | Report conflict; do not guess which applies |
| Convention is needed but not found in repo | Block or request confirmation depending on impact |
| Change touches architecture, testing, or security surfaces | Require evidence owner to be assigned to a downstream phase |
| Secure surface detected (data/auth/logs/secrets/APIs/files/DB) | Load `references/secure-design.md` before producing design commitments |

## Execution Steps

1. Read relevant project files, config, and docs (e.g., `openspec/config.yaml`, `AGENTS.md`, `skills/_shared/SKILL.md`, and applicable reference files in this skill).
2. Extract conventions across: naming, structure, boundaries, testing, error handling, logging, docs/config.
3. Produce design commitments and expected evidence per SDD phase.
4. For SDD phases: propose an `## Implementation Quality Constraints` section for `design.md`.
5. Load `references/secure-design.md` when the change touches: data, auth, permissions, logs, exports, files, DB, config/secrets, external APIs, or operational behavior.

## Output Contract

- `applied_conventions`: list of conventions applied with evidence source (file path or section).
- `gaps`: list of missing or contradicting conventions with impact assessment.
- `expected_evidence_by_phase`: object mapping phase names to required evidence.
  - `design`: implementation quality constraints documented.
  - `test-design`: convention checks planned.
  - `tasks`: convention-aligned task descriptions.
  - `apply`: convention-compliant code/files.
  - `review-security`: secure design evidence (if secure surface).
  - `verify`: convention compliance confirmed.
- `no_review_scoring`: always `true`; this output never includes review scores.

## References

- `references/architecture.md` — SDD executor model, shared contracts, boundaries.
- `references/code-style.md` — artifact language, YAML frontmatter schema, section order.
- `references/testing.md` — testing infrastructure state, verification path.
- `references/review-readiness.md` — delivery conventions, PR conventions.
- `references/secure-design.md` — secure design guidance per surface area.
