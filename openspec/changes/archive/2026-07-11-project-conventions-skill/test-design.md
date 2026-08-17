# Test Design: Project Conventions Skill

## Overview

This change produces seven static Markdown/YAML deliverables: one SKILL.md, five reference files, and one AGENTS.md modification. There is no automated test runner (`strict_tdd: false`). All verification is manual inspection against spec requirements and design contracts. The security classification is **no-impact** — these files are read-only instruction content with no runtime behavior, no data processing, and no secrets.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | `openspec/changes/project-conventions-skill/proposal.md` | Scope, intent, non-goals |
| Spec | `openspec/changes/project-conventions-skill/spec.md` | Requirements SPEC-001 through SPEC-013 and all scenarios |
| Design | `openspec/changes/project-conventions-skill/design.md` | File changes, section contracts, sdd-design integration pattern |
| Secure Development Design | `openspec/changes/project-conventions-skill/design.md#secure-development-design` | No-impact classification, safe evidence policy, boundary confirmation |
| Testing Capabilities | `openspec/config.yaml` (testing section) | `strict_tdd: false` — no automated test runner; all checks are manual/static |

## Source ID Coverage Baseline

Changed-surface classification: **no-impact**. The change produces static documentation consumed by LLM agents at prompt-construction time. No runtime surfaces are touched (no auth, sessions, DB, files I/O, external APIs, logging, error handling). All 8 security categories are omitted as reviewable omissions; canonical validation belongs to `review-security-report.json`. No Source ID rows, SEC-* matrices, or scoring tables are required here.

Safe-evidence commitment: reference files must contain no secrets, tokens, or PII; `secure-design.md` may show patterns only via obviously-fake placeholder values. This is the only security-adjacent check that is mandatory for this change.

## Test Cases

| ID | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| TD-001 | SPEC-001 / Design: file changes | `skills/project-conventions/SKILL.md` exists | manual | mandatory | File present at exact path; non-empty | Pure additive change |
| TD-002 | SPEC-001 | Activation Contract section present in SKILL.md | manual | mandatory | Section heading `## Activation Contract` present | |
| TD-003 | SPEC-001 | Activation Contract covers all three trigger conditions: (a) SDD phase needing conventions, (b) change touching architecture/naming/testing/security, (c) supplemental-only, not user-invocable | manual | mandatory | All three triggers readable in section text | |
| TD-004 | SPEC-002 | Hard Rules section present in SKILL.md | manual | mandatory | Section heading `## Hard Rules` present | |
| TD-005 | SPEC-002 | Hard Rule 1 — unproven patterns marked `unknown`/`needs confirmation` | manual | mandatory | Rule text present and explicit | |
| TD-006 | SPEC-002 | Hard Rule 2 — explicitly prohibits copying/materializing/scoring the 96-control matrix | manual | mandatory | Rule text references 96-control matrix prohibition | |
| TD-007 | SPEC-002 | Hard Rule 3 — explicitly prohibits SEC-* matrices, Source ID rows, formal scoring tables | manual | mandatory | Rule text references SEC-* and Source ID prohibition | |
| TD-008 | SPEC-002 | Hard Rule 4 — prohibits inventing test/lint/build commands not present in the project | manual | mandatory | Rule text present | |
| TD-009 | SPEC-002 | Hard Rule 5 — prohibits secrets, tokens, PII, raw logs, sensitive payloads as evidence | manual | mandatory | Rule text present | |
| TD-010 | SPEC-002 | Hard Rule 6 — prohibits overriding sdd-review-security's authority over review-security-report.json | manual | mandatory | Rule text present | |
| TD-011 | SPEC-003 | Decision Gates section present in SKILL.md | manual | mandatory | Section heading `## Decision Gates` present | |
| TD-012 | SPEC-003 | Decision Gates table has all 4 rows: known convention found / contradicting conventions / convention missing / change touches architecture/testing/security | manual | mandatory | 4-row table present with correct Situation and Action columns | |
| TD-013 | SPEC-004 | Execution Steps section present with exactly 5 numbered steps | manual | mandatory | Section heading `## Execution Steps` present; steps 1–5 listed | |
| TD-014 | SPEC-004 | Step 5 references `references/secure-design.md` as conditional load on security-adjacent surfaces | manual | mandatory | Step 5 text names the conditional surfaces (data, auth, permissions, logs, exports, files, DB, config/secrets, external APIs, operational behavior) | |
| TD-015 | SPEC-005 | Output Contract section present in SKILL.md | manual | mandatory | Section heading `## Output Contract` present | |
| TD-016 | SPEC-005 | Output Contract specifies: applied conventions summary, risks/gaps, expected evidence per phase (design/test-design/tasks/apply/review-security/verify), and prohibition on review scoring | manual | mandatory | All four fields present in section | |
| TD-017 | SPEC-004 / Design: frontmatter | SKILL.md frontmatter has all 4 required fields: `name`, `description`, `license`, `metadata` | static | mandatory | YAML frontmatter block contains exactly those four keys | |
| TD-018 | Design: section structure | SKILL.md section order matches prescribed sequence: Activation Contract → Hard Rules → Decision Gates → Execution Steps → Output Contract → References | manual | mandatory | Sections appear in that exact order | |
| TD-019 | SPEC-004 / Design | References section present in SKILL.md and points to all 5 reference files | manual | mandatory | Links or mentions for architecture.md, code-style.md, testing.md, review-readiness.md, secure-design.md | |
| TD-020 | SPEC-006 | `skills/project-conventions/references/architecture.md` exists | manual | mandatory | File present; non-empty | |
| TD-021 | SPEC-006 | architecture.md documents SDD executor model (phases are executor agents, not invoked skills) | manual | mandatory | Text present describing executor model | |
| TD-022 | SPEC-006 | architecture.md documents shared contracts location (`skills/_shared/`) | manual | mandatory | Path `skills/_shared/` mentioned | |
| TD-023 | SPEC-006 | architecture.md includes boundary statement: project-conventions guides; sdd-review judges | manual | mandatory | Exact boundary statement readable | |
| TD-024 | SPEC-006 | architecture.md references skill-resolver as supplemental injection mechanism | manual | mandatory | `skill-resolver` mentioned | |
| TD-025 | SPEC-007 | `skills/project-conventions/references/code-style.md` exists | manual | mandatory | File present; non-empty | |
| TD-026 | SPEC-007 | code-style.md states English artifacts rule explicitly (code, identifiers, comments, UI copy, docs, commit messages) | manual | mandatory | Explicit English rule present | |
| TD-027 | SPEC-007 | code-style.md documents YAML frontmatter schema with all 4 required fields | manual | mandatory | Fields `name`, `description`, `license`, `metadata` listed | |
| TD-028 | SPEC-007 | code-style.md documents section heading order | manual | mandatory | Prescribed section order documented | |
| TD-029 | SPEC-008 | `skills/project-conventions/references/testing.md` exists | manual | mandatory | File present; non-empty | |
| TD-030 | SPEC-008 | testing.md states `strict_tdd: false` | manual | mandatory | Exact text `strict_tdd: false` present | |
| TD-031 | SPEC-008 | testing.md states no automated test runner is configured | manual | mandatory | Explicit statement present | |
| TD-032 | SPEC-008 | testing.md references `test-design.md` as verification path | manual | mandatory | `test-design.md` mentioned | |
| TD-033 | SPEC-008 | testing.md includes update trigger note: file MUST be updated if strict_tdd or test commands change in `openspec/config.yaml` | manual | mandatory | Update trigger note present | |
| TD-034 | SPEC-009 | `skills/project-conventions/references/review-readiness.md` exists | manual | mandatory | File present; non-empty | |
| TD-035 | SPEC-009 | review-readiness.md states 400-line review budget | manual | mandatory | "400 lines" or "400-line" present | |
| TD-036 | SPEC-009 | review-readiness.md states auto-chain + stacked-to-main delivery strategy | manual | mandatory | Both terms present | |
| TD-037 | SPEC-009 | review-readiness.md references `chained-pr` skill for over-budget changes | manual | mandatory | `chained-pr` mentioned | |
| TD-038 | SPEC-009 | review-readiness.md states PR conventions: issue-first, conventional commits, no AI attribution | manual | mandatory | All three conventions present | |
| TD-039 | SPEC-010 | `skills/project-conventions/references/secure-design.md` exists | manual | mandatory | File present; non-empty | |
| TD-040 | SPEC-010 | secure-design.md contains all 9 surface areas (or confirmed 8 — see notes) | manual | mandatory | All surface area headings/sections present | Spec lists 9 in requirement text; test prompt says 8 — verify against final spec; both are satisfied if all areas from SPEC-010 appear |
| TD-041 | SPEC-010 | Each surface area uses narrative structure: design commitment / prohibited pattern / expected evidence / evidence owner / residual risk / exception policy | manual | mandatory | At least one of each element visible per area | |
| TD-042 | SPEC-010 | secure-design.md contains NO SEC-* matrices | static | mandatory | Grep for `SEC-` pattern returns no matches | Command: `rg "SEC-[A-Z0-9]+" skills/project-conventions/references/secure-design.md` |
| TD-043 | SPEC-010 | secure-design.md contains NO Source ID rows | static | mandatory | Grep for Source ID table patterns returns no matches | Command: `rg "Source ID" skills/project-conventions/references/secure-design.md` |
| TD-044 | SPEC-010 | secure-design.md contains NO Yes/No/N/A scoring columns | static | mandatory | No scoring table with Yes/No/N/A columns present | Inspect table structures |
| TD-045 | SPEC-011 | `AGENTS.md` contains `## Project Skills` section | manual | mandatory | Section heading present in file | |
| TD-046 | SPEC-011 | `## Project Skills` section registers `project-conventions` by name | manual | mandatory | `project-conventions` present in section | |
| TD-047 | SPEC-011 | Section describes skill as supplemental, NOT a required SDD DAG phase | manual | mandatory | Explicit "NOT a required SDD DAG phase" statement or equivalent present | |
| TD-048 | SPEC-013 / SPEC-002 | SKILL.md hard rules explicitly prohibit 96-control matrix duplication | manual | mandatory | Covered by TD-006 | Cross-check with TD-006 |
| TD-049 | SPEC-013 / SPEC-002 | SKILL.md hard rules explicitly prohibit SEC-* matrices and Source ID rows | manual | mandatory | Covered by TD-007 | Cross-check with TD-007 |
| TD-050 | SPEC-010 / SPEC-013 | No reference file contains any scoring table, matrix, or formal control catalog | static | mandatory | Grep for `SEC-`, `Source ID`, `\| Yes \|`, `\| No \|`, `\| N/A \|` patterns across all reference files returns no matches | Command: `rg "SEC-|Source ID" skills/project-conventions/references/` |
| TD-051 | Design: safe evidence policy | secure-design.md contains no real secrets, tokens, or PII — only placeholder examples (e.g., `<YOUR_SECRET_HERE>`) | manual | mandatory | Any secret-like content uses clearly-fake placeholders | |
| TD-052 | SPEC-012 / Design: integration pattern | `design.md` (this change) contains `## Implementation Quality Constraints` section | manual | mandatory | Section heading present in `openspec/changes/project-conventions-skill/design.md` | Validates the integration pattern demonstrated by this change itself |
| TD-053 | SPEC-012 / Design: integration pattern | `design.md` (this change) contains `## Secure Design Constraints` subsection or equivalent | manual | mandatory | Subsection or table present under Implementation Quality Constraints | May be a subsection or inline table |
| TD-054 | SPEC-012 | Secure Design Constraints table in design.md has at minimum: Data handling, Secrets/config, Logging/errors rows | manual | mandatory | All three row labels present | |
| TD-055 | SPEC-012 | Table columns match required schema: Area \| Project Secure Convention \| Design Commitment \| Evidence Owner | manual | mandatory | Column headers match exactly | |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| No-impact verification | Confirm `design.md#Operational Considerations` states "No aplica." | manual | Exact text `No aplica.` present in Operational Considerations section | N/A — static inspection |
| Artifact safety | No SDD artifact in this change contains secrets, tokens, or PII | manual | All files reviewed; `secure-design.md` uses placeholder values only | N/A — manual review |
| Boundary confirmation | Confirm no phase artifact creates or modifies `review-security-report.json` | manual | No `review-security-report.json` produced by this change | N/A — static inspection |

## Security Control Coverage

| Guideline | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| Safe evidence policy | No real secrets in any produced file | Yes | TD-051: manual inspection of secure-design.md and all reference files for placeholder-only patterns | covered | None |
| No SEC-* matrices | secure-design.md must not contain SEC-* or Source ID rows | Yes | TD-042, TD-043, TD-050: static grep | covered | None |
| No scoring tables | No Yes/No/N/A scoring columns in any file | Yes | TD-044, TD-050: static and manual inspection | covered | None |
| Boundary — review-security authority | review-security-report.json not created/modified | Yes | TD-049, operational check: no such file produced | covered | None |
| All 8 security categories | Omitted (no-impact classification) | No — reviewable omissions | Owned by canonical review-security-report.json | not-applicable | No-impact classification per design.md#secure-development-design |

## Evidence Expectations

- Mandatory cases require manual or static inspection confirmation before `sdd-verify` can report success.
- Non-mandatory cases are advisory; they do not block verification.
- Security validation evidence: `design.md#secure-development-design` classifies this change as no-impact with a safe-evidence policy. Mandatory controls cover only content safety (no secrets/PII) and anti-matrix boundaries. Exhaustive source-row validation remains with `review-security-report.json`.
- Unavailable tooling: no automated test runner is configured (`strict_tdd: false`). All evidence is manual/static inspection. This is a reported constraint; it is not passing evidence on its own.
- No-impact classification is justified by changed-surface analysis in `design.md#secure-development-design`. Absence of a standalone `security-design.md` is not a blocker.
- Runtime tests, coverage commands, linters, type checkers, and formatters are all unavailable for this repository.

## Open Questions

- None. All requirements are fully specified per design.md.
