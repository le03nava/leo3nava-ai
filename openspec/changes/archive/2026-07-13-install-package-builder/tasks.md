# Tasks: Install Package Builder

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 350–500 |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | null |
| Chain strategy | pending |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: pending
Size exception: none
Review budget lines: 400
Review budget risk: Medium
400-line budget risk: Medium

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | All 8 deliverables (new files + registrations) | PR 1 | Additive only; no existing logic changed; single reviewable unit is sufficient |

---

## Phase 1: Skill Core (T-01)

- [x] T-01-A: Create `src/skills/install-package-builder/SKILL.md` with correct YAML frontmatter: `name`, `description` starting with "Trigger:" (≤250 chars), `user-invocable: true`, `license: MIT`, `metadata.author`, `metadata.version` → satisfies TC-A-001, TC-A-002, TC-A-004
- [x] T-01-B: Add `## Activation Contract` section to SKILL.md → satisfies TC-A-003
- [x] T-01-C: Add `## Pre-flight Check` section listing tools: git, mvn, npm, ng, gradle (or gradlew) → satisfies TC-A-003 (FR-002)
- [x] T-01-D: Add `## Input Schema` section documenting 7 inputs: changeNumber, version, outputPath, repos[], author, userStories[], corrections[] → satisfies TC-A-003 (FR-001)
- [x] T-01-E: Add `## Execution Steps` section with all 9 pipeline stages in order: preflight → clone → detect → build → classify → assemble → readme → zip → report → satisfies TC-A-003, TC-A-005
- [x] T-01-F: Add `## SQL Classification Rules` section with 6 rules in precedence order: Packages (PKS/PKB), Procedures (_PRC/SP_/PROC/_PROC), Types (_TYPE/TYPE_), Tables (_TAB/_TABLE/CREATE TABLE), Data (execution.sql/INSERT/UPDATE/MERGE), Unclassified + numeric prefix handling rule → satisfies TC-A-003, TC-A-006 (FR-006)
- [x] T-01-G: Add `## Hard Rules` section: no deploy, no secrets/credentials stored, read-only on source repos, idempotent output → satisfies TC-A-003 (NFR-001, NFR-002, NFR-003)
- [x] T-01-H: Add `## Output Contract` section describing the canonical folder layout (`Install/AppDeploy/`, `Install/DBObjects/`, `Install/ShellScripts/`, `Install/Instalacion/README.txt`, `Rollback/`, zip) → satisfies TC-A-003

## Phase 2: Asset Files (T-02, T-03)

- [x] T-02-A: Create `src/skills/install-package-builder/assets/readme-template.txt` with 4 required sections: Información General, installation steps, Control de Versiones del documento, Control de Revisiones de la Plantilla STTI → satisfies TC-B-001, TC-B-003
- [x] T-02-B: Ensure readme-template.txt contains all 12 markers: `{{CHG_NUMBER}}`, `{{VERSION}}`, `{{DATE}}`, `{{AUTHOR}}`, `{{APP_NAME}}`, `{{APP_VERSION}}`, `{{REPO_URL}}`, `{{BRANCH}}`, `{{ENVIRONMENT_TABLE}}`, `{{INSTALL_STEPS}}`, `{{USER_STORIES}}`, `{{CORRECTIONS}}` → satisfies TC-B-002
- [x] T-03-A: Create `src/skills/install-package-builder/assets/environments.yaml` with exactly 24 OWEXX environment entries, each containing: name, host, containers (web list and api list) using the real host values provided in the design → satisfies TC-B-004, TC-B-005, TC-B-006

## Phase 3: Reference File (T-04)

- [x] T-04-A: Create `src/skills/install-package-builder/references/sql-classification.md` with all 6 classification rules in precedence order and filename pattern examples for each rule (PKS/PKB, _PRC/SP_/PROC/_PROC, _TYPE/TYPE_, _TAB/_TABLE/CREATE TABLE, execution.sql/INSERT/UPDATE/MERGE, Unclassified) → satisfies TC-C-001, TC-C-002

## Phase 4: Adapter Prompts (T-05, T-06)

- [x] T-05-A: Create `agents/sdd/install-package-builder.md` with instruction to read SKILL.md before starting and input forwarding; no pipeline logic in this file → satisfies TC-D-001, TC-D-002, TC-D-003
- [x] T-06-A: Create `C:/Users/leo3n/.config/opencode/prompts/sdd/install-package-builder.md` with identical content as T-05-A → mirrors TC-D-001 check for actual loaded prompt

## Phase 5: Registration (T-07, T-08)

- [x] T-07-A: Edit `AGENTS.md` — add `install-package-builder` entry following the existing manual utilities pattern; document as user-invocable skill, not an SDD DAG phase → satisfies TC-E-001
- [x] T-08-A: Edit `C:/Users/leo3n/.config/opencode/opencode.json` — add agent entry: `{ "description": "...", "prompt": "{file:C:/Users/leo3n/.config/opencode/prompts/sdd/install-package-builder.md}" }` → satisfies TC-E-002

## Phase 6: Static Verification Evidence

- [x] V-01: Confirm `src/skills/install-package-builder/SKILL.md` exists and all frontmatter fields are present — evidence for TC-A-001, TC-A-002
- [x] V-02: Confirm SKILL.md description starts with "Trigger:" — evidence for TC-A-004
- [x] V-03: Confirm all 7 required section headings are present in SKILL.md — evidence for TC-A-003
- [x] V-04: Confirm Execution Steps lists all 9 stages in document order — evidence for TC-A-005
- [x] V-05: Confirm SQL Classification Rules lists all 6 rules in precedence order — evidence for TC-A-006
- [x] V-06: Confirm readme-template.txt exists with all 12 markers and 4 section headings — evidence for TC-B-001, TC-B-002, TC-B-003
- [x] V-07: Confirm environments.yaml exists, parses as valid YAML, has exactly 24 entries with all 24 names — evidence for TC-B-004, TC-B-005, TC-B-006
- [x] V-08: Confirm sql-classification.md exists with all 6 rules and patterns — evidence for TC-C-001, TC-C-002
- [x] V-09: Confirm adapter prompt exists, references SKILL.md, contains no pipeline logic — evidence for TC-D-001, TC-D-002, TC-D-003
- [x] V-10: Confirm AGENTS.md contains `install-package-builder` — evidence for TC-E-001
- [x] V-11: Confirm opencode.json agent entry with correct prompt path — evidence for TC-E-002
- [x] V-12: Confirm `design.md#Operational Considerations` contains "No aplica." marker — operational evidence (static)

## Test Case Coverage Map

| Task(s) | Test Cases Satisfied |
|---------|---------------------|
| T-01-A | TC-A-001, TC-A-002, TC-A-004 |
| T-01-B–H | TC-A-003 |
| T-01-E | TC-A-005 |
| T-01-F | TC-A-006 |
| T-02-A | TC-B-001, TC-B-003 |
| T-02-B | TC-B-002 |
| T-03-A | TC-B-004, TC-B-005, TC-B-006 |
| T-04-A | TC-C-001, TC-C-002 |
| T-05-A, T-06-A | TC-D-001, TC-D-002, TC-D-003 |
| T-07-A | TC-E-001 |
| T-08-A | TC-E-002 |
| V-01–V-12 | All 19 mandatory cases (static verification pass) |
