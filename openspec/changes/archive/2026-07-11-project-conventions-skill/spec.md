# project-conventions-skill Specification

## Purpose

This spec defines the requirements and behavioral contracts for the `project-conventions-skill` change.
It covers the new `skills/project-conventions/SKILL.md` skill, its five reference files, the `AGENTS.md`
update, and the boundary contracts that separate project-conventions from sdd-review and sdd-review-security.

---

## Requirements

### Requirement: SKILL.md Activation Contract (SPEC-001)

`skills/project-conventions/SKILL.md` MUST include an **Activation Contract** section.

The Activation Contract MUST state:
- Load this skill when any SDD phase (`sdd-design`, `sdd-test-design`, `sdd-tasks`, `sdd-apply`) needs real project conventions.
- Load when the change touches architecture, naming, testing, security surfaces, or review-readiness.
- This is a supplemental skill injected by the orchestrator or skill-resolver; it MUST NOT be invoked directly by users.

#### Scenario: Phase agent loads project-conventions as supplemental skill

- GIVEN a `sdd-design` executor receives a change that touches naming or architecture
- WHEN the orchestrator injects `project-conventions` as a supplemental skill
- THEN the executor reads `SKILL.md` before generating `design.md`
- AND applies conventions found in the reference files

#### Scenario: User attempts to invoke project-conventions directly

- GIVEN a user sends a message that matches the skill trigger
- WHEN the system evaluates whether to load the skill
- THEN the system treats it as a supplemental skill only and does not run it as a standalone phase
- AND the orchestrator is the authoritative caller

---

### Requirement: SKILL.md Hard Rules (SPEC-002)

`skills/project-conventions/SKILL.md` MUST declare six non-negotiable hard rules.

The file MUST prohibit:
1. Using patterns not proven by existing files, config, or docs — unproven items MUST be marked `unknown` or `needs confirmation`.
2. Copying, materializing, or scoring the sdd-review 96-control matrix.
3. Creating SEC-* matrices, Source ID rows, or formal scoring tables.
4. Inventing test, lint, or build commands not present in the project.
5. Including secrets, tokens, PII, raw logs, or sensitive payloads as evidence.
6. Overriding `sdd-review-security`'s authority over `review-security-report.json`.

#### Scenario: Convention found but unconfirmed

- GIVEN project-conventions is loaded and a naming pattern is referenced in one file only
- WHEN the executor applies conventions
- THEN the item is marked `[needs confirmation]` in design.md
- AND the executor does not assert it as a hard rule

#### Scenario: Executor attempts to reproduce review scoring

- GIVEN project-conventions is loaded alongside sdd-design
- WHEN sdd-design produces design.md
- THEN no 96-control matrix row, SEC-* table, or Source ID row appears in the output
- AND sdd-review retains sole authority over that judgment

---

### Requirement: SKILL.md Decision Gates (SPEC-003)

`skills/project-conventions/SKILL.md` MUST include a **Decision Gates** section covering four cases.

The gates MUST be:
| Situation | Action |
|-----------|--------|
| Known convention found | Apply it |
| Contradicting conventions found | Report conflict; do not guess |
| Convention needed but missing | Block or request confirmation based on impact |
| Change touches architecture, testing, or security | Require evidence owner phase |

#### Scenario: Contradicting conventions found

- GIVEN two config files define conflicting naming conventions for the same domain
- WHEN project-conventions is applied
- THEN the executor reports the conflict in design.md risks section
- AND does not choose one arbitrarily

#### Scenario: Convention missing for high-impact area

- GIVEN the change touches auth and no auth convention exists in any reference file
- WHEN project-conventions evaluates the gate
- THEN the executor blocks or requests explicit confirmation before proceeding

---

### Requirement: SKILL.md Execution Steps (SPEC-004)

`skills/project-conventions/SKILL.md` MUST document five execution steps.

The steps MUST be:
1. Read relevant project files, config, and docs.
2. Extract conventions across: naming, structure, boundaries, testing, error handling, logging, docs/config.
3. Produce design commitments and expected evidence per SDD phase.
4. For SDD phases: propose an `## Implementation Quality Constraints` section for `design.md`.
5. Load `references/secure-design.md` when the change touches: data, auth, permissions, logs, exports, files, DB, config/secrets, external APIs, or operational behavior.

#### Scenario: Change touches data handling

- GIVEN project-conventions is loaded for a change that modifies a data-export endpoint
- WHEN step 5 is evaluated
- THEN `references/secure-design.md` is loaded
- AND secure-design constraints are included in the `## Implementation Quality Constraints` section

#### Scenario: Change is purely a UI label rename

- GIVEN project-conventions is loaded for a change with no data, auth, or security surface
- WHEN step 5 is evaluated
- THEN `references/secure-design.md` is NOT loaded
- AND the Implementation Quality Constraints section contains only style/naming conventions

---

### Requirement: SKILL.md Output Contract (SPEC-005)

`skills/project-conventions/SKILL.md` MUST define an **Output Contract** section.

The output contract MUST specify:
- Summary of applied conventions.
- Risks or gaps (missing or contradicting conventions).
- Expected evidence per phase: design, test-design, tasks, apply, review-security, verify.
- Prohibition: NO review scoring in any output.

#### Scenario: Output produced after applying conventions

- GIVEN project-conventions has read all relevant reference files
- WHEN the executor writes its output
- THEN the output contains applied conventions, risks/gaps, and evidence expectations per phase
- AND contains no scores, ratings, or compliance verdicts

---

### Requirement: references/architecture.md Content (SPEC-006)

`skills/project-conventions/references/architecture.md` MUST document:
- SDD executor model: phases are executor agents, not invoked skills.
- Shared contracts live in `skills/_shared/`.
- Boundary rule: project-conventions guides; sdd-review judges.
- No parallel dependent planning phases.
- Supplemental skills are injected via skill-resolver, not hardcoded in phase skills.

#### Scenario: Phase agent reads architecture reference

- GIVEN sdd-design loads project-conventions and reads `references/architecture.md`
- WHEN generating design.md
- THEN the executor applies the boundary rule and does not embed review judgments
- AND supplemental skill injection is via skill-resolver only

---

### Requirement: references/code-style.md Content (SPEC-007)

`skills/project-conventions/references/code-style.md` MUST document:
- All generated artifacts in English: code, identifiers, comments, UI copy, docs, commit messages.
- Skill frontmatter MUST use YAML with fixed fields: `name`, `description`, `license`, `metadata`.
- Section headings follow the fixed order: Activation Contract, Hard Rules, Decision Gates, Execution Steps, Output Contract, References.
- Markdown conventions for skills: consistent heading levels, table format for gates and constraints.

#### Scenario: New skill is created following code-style reference

- GIVEN a developer reads `references/code-style.md` before creating a new skill
- WHEN the skill file is written
- THEN frontmatter includes all four required YAML fields
- AND sections appear in the prescribed order

---

### Requirement: references/testing.md Content (SPEC-008)

`skills/project-conventions/references/testing.md` MUST document:
- `strict_tdd: false` — no automated test runner is configured as of the spec date.
- Verification relies on manual checks defined in `test-design.md`.
- An explicit note: this file MUST be updated if `strict_tdd` or test commands change in `openspec/config.yaml`.
- If `strict_tdd` becomes true, SDD apply MUST follow `strict-tdd.md`.

#### Scenario: Executor reads testing reference before writing tasks

- GIVEN project-conventions is loaded for a change requiring test evidence
- WHEN the executor reads `references/testing.md`
- THEN it applies manual-verification strategy to task definitions
- AND does not invent automated test commands

#### Scenario: strict_tdd changes in the future

- GIVEN `openspec/config.yaml` is updated to set `strict_tdd: true`
- WHEN the next SDD apply cycle runs
- THEN `references/testing.md` MUST be updated before that cycle proceeds
- AND `strict-tdd.md` becomes the authoritative testing contract

---

### Requirement: references/review-readiness.md Content (SPEC-009)

`skills/project-conventions/references/review-readiness.md` MUST document:
- Review budget: 400 lines per PR.
- Delivery strategy: auto-chain, stacked-to-main.
- Changes over 400 lines MUST use `chained-pr` skill.
- PR conventions: issue-first, conventional commits, no AI attribution in commit messages.

#### Scenario: Change exceeds 400-line budget

- GIVEN a change produces more than 400 lines of diff
- WHEN review-readiness constraints are applied
- THEN the executor flags the change for chained-pr processing
- AND the delivery plan references `chained-pr` skill

#### Scenario: Commit message is written

- GIVEN an executor produces a commit as part of sdd-apply
- WHEN review-readiness is active
- THEN the commit message follows conventional commits format
- AND contains no AI attribution ("Co-Authored-By" or similar)

---

### Requirement: references/secure-design.md Content (SPEC-010)

`skills/project-conventions/references/secure-design.md` MUST document all nine surface areas using narrative rules only.

Each rule MUST use this structure where applicable:
- design commitment
- prohibited unsafe pattern
- expected safe evidence
- evidence owner phase
- residual risk
- exception policy

The nine areas MUST be:
1. **Changed-surface classification** — inputs, outputs, storage, logs, exports, auth, sessions, permissions, files, database, config/secrets, external APIs.
2. **Data sensitivity levels** — public / internal / confidential / PII / PAN / secrets / tokens / unknown; if unknown, MUST block or ask.
3. **Access control** — document auth, authorization, ownership/tenant boundaries, privilege checks.
4. **Input/output handling** — validation, sanitization, encoding, canonicalization, file/path safety.
5. **Secrets/config** — no hardcoded secrets, no tokens in logs, no connection strings in SDD evidence.
6. **Logging/error handling** — useful errors without leakage; sanitized log summaries.
7. **Dependencies/external interfaces** — risks from new libraries, external APIs, network calls, retries/timeouts.
8. **Operational security** — safe evidence, monitoring mechanism, recovery/reprocessing, final-document-only boundaries.
9. **Evidence expectations per phase** — design, test-design, tasks, apply, review-security, verify, archive.

The file MUST NOT contain SEC-* matrices, Source ID rows, or Yes/No/N/A scoring columns.

#### Scenario: Change touches auth surface

- GIVEN `references/secure-design.md` is loaded for a change modifying authentication
- WHEN the executor processes area 3 (Access control)
- THEN design.md includes narrative design commitments for auth and authorization boundaries
- AND no scoring table or SEC-* row appears in the output

#### Scenario: Data sensitivity is unknown

- GIVEN a new data field is introduced and its sensitivity classification is not documented
- WHEN area 2 (Data sensitivity) is evaluated
- THEN the executor blocks or requests confirmation before proceeding
- AND does not assume a default sensitivity level

#### Scenario: Secrets appear in SDD evidence

- GIVEN an executor is writing design evidence that references a config value
- WHEN area 5 (Secrets/config) rule is applied
- THEN the evidence contains only key names and descriptions
- AND no actual secret, token, or connection string value appears in any SDD artifact

---

### Requirement: AGENTS.md Project Skills section (SPEC-011)

`AGENTS.md` MUST be updated to include a `## Project Skills` section.

The section MUST:
- Register the `project-conventions` skill by name.
- State its trigger: supplemental skill loaded by SDD phases when project conventions are needed.
- Note it is NOT a required SDD DAG phase.

#### Scenario: New contributor reads AGENTS.md

- GIVEN a developer opens `AGENTS.md` to understand available skills
- WHEN they reach the `## Project Skills` section
- THEN they find `project-conventions` listed with its trigger and a note that it is supplemental only
- AND they do not mistake it for a mandatory SDD phase

---

### Requirement: sdd-design integration pattern (SPEC-012)

When `project-conventions` is injected as a supplemental skill into `sdd-design`, `sdd-design` SHOULD add an `## Implementation Quality Constraints` section to `design.md`.

This section MUST document project-specific commitments sourced from project-conventions.

An optional `## Secure Design Constraints` subsection MAY be included using this table format:

| Area | Project Secure Convention | Design Commitment | Evidence Owner |
|------|--------------------------|-------------------|----------------|
| Data handling | … | … | … |
| Access control | … | … | … |
| Logging/errors | … | … | … |

This section is GUIDANCE for design authors. sdd-review retains sole authority to judge against the 96-control matrix.

#### Scenario: sdd-design produces design.md with quality constraints

- GIVEN project-conventions is injected and the change touches naming and testing
- WHEN sdd-design writes design.md
- THEN design.md contains `## Implementation Quality Constraints` with project-sourced commitments
- AND the section contains no review verdicts or compliance scores

#### Scenario: sdd-design produces design.md with secure design constraints

- GIVEN project-conventions is injected and `references/secure-design.md` is loaded
- WHEN sdd-design writes design.md
- THEN design.md MAY include `## Secure Design Constraints` as a narrative table
- AND sdd-review still applies the 96-control matrix independently at review time

---

### Requirement: Boundary Verification (SPEC-013)

The boundary between project-conventions and the review phases MUST be explicitly enforced.

The boundary MUST confirm:
- `project-conventions` guides design decisions and expected evidence.
- `sdd-design` documents implementation commitments.
- `sdd-test-design` plans checks against those commitments.
- `sdd-tasks` creates implementation tasks aligned to those commitments.
- `sdd-apply` implements with conventions in mind.
- `sdd-review` judges against the 96-control matrix and is the sole owner of that judgment.
- `sdd-review-security` judges against the Source ID catalog and is the sole owner of `review-security-report.json`.

No phase other than `sdd-review` MUST produce or modify the 96-control matrix evaluation.
No phase other than `sdd-review-security` MUST produce or modify `review-security-report.json`.

#### Scenario: sdd-apply inadvertently scores review controls

- GIVEN project-conventions is active during sdd-apply
- WHEN an executor writes implementation notes
- THEN no 96-control evaluation or partial review score appears in the apply output
- AND review artifacts remain exclusively owned by sdd-review

#### Scenario: sdd-design generates review-security-report.json

- GIVEN project-conventions loads `references/secure-design.md` during sdd-design
- WHEN design.md is written
- THEN `review-security-report.json` is NOT created or modified
- AND sdd-review-security remains the sole owner of that artifact
