# Test Design: Install Package Builder

## Overview

This change introduces a standalone LLM skill (`install-package-builder`) with no automated test runner. All tests are **static** — they verify that the skill files, asset files, reference files, adapter prompt, and registration entries exist with the correct structure and content. No pytest, no jest, no executable pipeline tests exist; verification is done by reading files and inspecting content. The testing surface is limited to artifact presence, frontmatter correctness, section completeness, YAML validity, marker presence, and registration entries.

## Inputs

| Artifact | Reference | Used For |
| --- | --- | --- |
| Proposal | openspec/changes/install-package-builder/proposal.md | Scope, success criteria, environment count (24 OWEXX), pre-flight requirements |
| Spec | openspec/changes/install-package-builder/spec.md | FR-001–FR-012 requirements and acceptance scenarios defining pipeline behavior |
| Design | openspec/changes/install-package-builder/design.md | Architecture decisions (skill-as-instructions, YAML env data, template README), 9 pipeline stages, file changes |
| Secure Development Design | openspec/changes/install-package-builder/design.md#secure-development-design | Classification: No security impact — no SEC categories applicable |
| Testing Capabilities | Orchestrator session context | No automated test runner available — all tests are static (file existence + content inspection) |

## Source ID Coverage Baseline

No corporate source-row coverage applies to this change. The Secure Development Design section in `design.md` classifies this change as **No security impact** — no authentication, session, data, secrets, logging, or input-validation surfaces are introduced. No SEC category rules are applicable. `review-security-report.json` retains authority for exhaustive source-row validation and `N/A` decisions.

## Test Cases

| ID | Group | Source | Check | Type | Severity | Expected Evidence | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| TC-A-001 | A | design: File Changes | `src/skills/install-package-builder/SKILL.md` exists | static | mandatory | File present at path | Blocker for all other Group A checks |
| TC-A-002 | A | NFR-004, skill conventions | SKILL.md frontmatter has: name, description, user-invocable: true, license, metadata.author, metadata.version | static | mandatory | All 6 frontmatter fields present | user-invocable must be true (not false) |
| TC-A-003 | A | FR-001–FR-012 pipeline coverage | SKILL.md contains all 7 required sections: Activation Contract, Pre-flight Check, Input Schema, Execution Steps, SQL Classification Rules, Hard Rules, Output Contract | static | mandatory | All section headings found | Absence of any section is a blocker |
| TC-A-004 | A | skill conventions | SKILL.md description starts with "Trigger:" and is ≤250 chars | static | mandatory | description value passes both checks | Match pattern: `description: "Trigger:.*"` |
| TC-A-005 | A | design: pipeline stages order | Execution Steps section lists 9 stages in order: preflight, clone, detect, build, classify, assemble, readme, zip, report | static | mandatory | All 9 stage names appear in document order | Order matters — verify relative position in file |
| TC-A-006 | A | spec: FR-006 | SQL Classification Rules section lists 6 rules in precedence order: Packages, Procedures, Types, Tables, Data, Unclassified | static | mandatory | All 6 rule categories appear in order | Precedence order determines classification outcome |
| TC-B-001 | B | design: File Changes | `src/skills/install-package-builder/assets/readme-template.txt` exists | static | mandatory | File present at path | Blocker for TC-B-002 and TC-B-003 |
| TC-B-002 | B | spec: FR-008 | readme-template.txt contains all 12 required markers: {{CHG_NUMBER}}, {{VERSION}}, {{DATE}}, {{AUTHOR}}, {{APP_NAME}}, {{APP_VERSION}}, {{REPO_URL}}, {{BRANCH}}, {{ENVIRONMENT_TABLE}}, {{INSTALL_STEPS}}, {{USER_STORIES}}, {{CORRECTIONS}} | static | mandatory | All 12 markers found in file | Missing marker = README generation silently omits data |
| TC-B-003 | B | spec: FR-008 | readme-template.txt contains 4 required sections: Información General, installation steps section, Control de Versiones del documento, Control de Revisiones de la Plantilla STTI | static | mandatory | All 4 section headings found | Section headings enable operator navigation |
| TC-B-004 | B | design: File Changes, Decision: YAML | `src/skills/install-package-builder/assets/environments.yaml` exists | static | mandatory | File present at path | Proposal originally said .json; design chose YAML |
| TC-B-005 | B | proposal, spec: FR-008 | environments.yaml is valid YAML with exactly 24 environment entries | static | mandatory | YAML parses without error; entry count == 24 | Count mismatch means README table is incomplete |
| TC-B-006 | B | spec: FR-008 | environments.yaml contains all 24 named environments (Desarrollo through Villa Hermosa) | static | mandatory | All 24 names found in file | Full list: Desarrollo, QA, Aguascalientes, Azcapotzalco, Chihuahua, Culiacán, Escobedo, Guadalupe, La Paz, León, Merida, Mexicali, Obregón, Puebla, Querétaro, Reynosa, Saltillo, San Martín, Tampico, Tijuana, Tlajomulco, Toluca, Veracruz, Villa Hermosa |
| TC-C-001 | C | design: File Changes | `src/skills/install-package-builder/references/sql-classification.md` exists | static | mandatory | File present at path | Human-readable reference for rule maintenance |
| TC-C-002 | C | spec: FR-006 | sql-classification.md documents all 6 classification rules with their filename patterns | static | mandatory | All 6 rule patterns documented: PKS/PKB, _PRC/SP_/PROC/_PROC, _TYPE/TYPE_, _TAB/_TABLE/CREATE TABLE, execution.sql/INSERT/UPDATE/MERGE, Unclassified | Reference must stay in sync with SKILL.md rules |
| TC-D-001 | D | design: File Changes | `agents/sdd/install-package-builder.md` (repo copy) exists | static | mandatory | File present at path | Repo copy enables version-controlled adapter history |
| TC-D-002 | D | design: Skill-as-Instructions Pattern | Adapter prompt contains instruction to read/load SKILL.md before starting | static | mandatory | Instruction referencing SKILL.md found in file | Without this, agent executes without skill context |
| TC-D-003 | D | design: Decision — logic in SKILL.md only | Adapter prompt does NOT contain pipeline logic (git/mvn/npm commands, SQL classification rules) | static | mandatory | No pipeline logic found in adapter prompt | Logic duplication causes drift and maintenance risk |
| TC-E-001 | E | design: File Changes — AGENTS.md Modify | AGENTS.md contains an entry for `install-package-builder` | static | mandatory | 'install-package-builder' found in AGENTS.md | Must document as manual utility, not SDD DAG phase |
| TC-E-002 | E | design: File Changes — opencode.json Modify | `~/.config/opencode/opencode.json` contains agent entry for install-package-builder with prompt field pointing to prompts/sdd/ | static | mandatory | Agent entry with correct prompt path found | Absence means skill is not discoverable by opencode |

## Operational Considerations Checks

| Category | Planned Check | Type | Expected Evidence | Unavailable Tooling Note |
| --- | --- | --- | --- | --- |
| Operational Considerations | design.md states "No aplica." for operational considerations | static | Read design.md#Operational Considerations — confirm "No aplica." marker is present | No runtime monitoring tooling applicable |

## Security Control Coverage

| Guideline ID | Required Control | Mandatory | Planned Check or Evidence | Status | Exception |
| --- | --- | --- | --- | --- | --- |
| SEC-AUTH | Authentication logic | No | design.md#secure-development-design confirms no auth introduced; no check planned | not-applicable | Classification: No security impact |
| SEC-SESSION | Session management | No | design.md#secure-development-design confirms no sessions introduced | not-applicable | Classification: No security impact |
| SEC-DATA | PII/confidential data | No | design.md confirms environments.yaml contains only internal infrastructure names; no PII | not-applicable | Classification: No security impact |
| SEC-SECRET | Secrets handling | No | design.md confirms NFR-002 prohibits credential input/storage; no secrets in artifacts | not-applicable | Classification: No security impact |
| SEC-LOG | Logging infrastructure | No | design.md confirms no logging infrastructure created; summary report contains only paths/counts | not-applicable | Classification: No security impact |
| SEC-INPUT | Untrusted input validation | No | design.md confirms inputs come from invoking user via LLM conversation, not external untrusted sources | not-applicable | Classification: No security impact |

## Evidence Expectations

- All 19 test cases are **mandatory**. Each must produce a passing static check (file existence confirmed or content pattern found) before verification is complete.
- No automated test runner is available. All evidence is produced by reading files and inspecting content.
- Security validation evidence: `design.md#secure-development-design` classifies this change as No security impact. All SEC categories are not-applicable with documented rationale. `review-security-report.json` retains authority for exhaustive source-row validation.
- The "No aplica." operational considerations marker in design.md is the planned static evidence that no operational checks are omitted.
- Missing tooling (pytest, jest, coverage, linters) is not passing evidence — it is an explicit constraint; all checks are planned as static instead.

## Open Questions

- None. All technical decisions are resolved. Design is complete and unambiguous.
