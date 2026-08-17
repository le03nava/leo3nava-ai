# Tasks: Project Conventions Skill

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | ~380–450 lines across 7 files |
| Review budget lines | 400 |
| 400-line budget risk | Medium |
| Review budget risk | Medium |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 → PR 2 (see work units below) |
| Delivery strategy | auto-chain |
| Chain strategy | stacked-to-main |
| Size exception | none |

Decision needed before apply: No
Chained PRs recommended: Yes
Chain strategy: stacked-to-main
Size exception: none
Review budget lines: 400
Review budget risk: Medium
400-line budget risk: Medium

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | SKILL.md + architecture.md + code-style.md + testing.md | PR 1 | Base: main; ~220–260 lines |
| 2 | review-readiness.md + secure-design.md + AGENTS.md update | PR 2 | Base: PR 1 branch; ~160–190 lines |

---

## Phase 1: Foundation — Skill Directory and Core SKILL.md

- [x] 1.1 Create directory `skills/project-conventions/references/` (ensure path exists before writing any file)
- [x] 1.2 Write `skills/project-conventions/SKILL.md` with YAML frontmatter containing exactly: `name`, `description`, `disable-model-invocation`, `user-invocable: false`, `license`, `metadata` (author, version, delegate_only) — satisfies TD-017
- [x] 1.3 Add `## Activation Contract` section to SKILL.md: state all 3 trigger conditions (SDD phase needing conventions; change touches architecture/naming/testing/security/review-readiness; supplemental only, not user-invocable) — satisfies TD-002, TD-003
- [x] 1.4 Add `## Hard Rules` section with all 6 prohibitions: (1) unproven patterns → mark `unknown`/`needs confirmation`; (2) no 96-control matrix copy/materialize/scoring; (3) no SEC-* matrices/Source ID rows/scoring tables; (4) no invented test/lint/build commands; (5) no secrets/tokens/PII/raw logs as evidence; (6) no override of sdd-review-security authority — satisfies TD-004–TD-010
- [x] 1.5 Add `## Decision Gates` section with 4-row table (known convention / contradicting conventions / convention missing / change touches architecture+testing+security) — satisfies TD-011, TD-012
- [x] 1.6 Add `## Execution Steps` section with exactly 5 numbered steps; Step 5 must name conditional surfaces for `references/secure-design.md` load (data, auth, permissions, logs, exports, files, DB, config/secrets, external APIs, operational behavior) — satisfies TD-013, TD-014
- [x] 1.7 Add `## Output Contract` section: applied conventions summary, risks/gaps, expected evidence per phase (design/test-design/tasks/apply/review-security/verify), explicit prohibition on review scoring — satisfies TD-015, TD-016
- [x] 1.8 Add `## References` section pointing to all 5 reference files by relative path — satisfies TD-018, TD-019
- [x] 1.9 Verify SKILL.md section order matches: Activation Contract → Hard Rules → Decision Gates → Execution Steps → Output Contract → References — satisfies TD-018

## Phase 2: Reference Files

- [x] 2.1 Write `skills/project-conventions/references/architecture.md`: document SDD executor model (phases are executor agents, not invoked skills); shared contracts live in `skills/_shared/`; boundary rule (project-conventions guides; sdd-review judges); no parallel dependent planning phases; supplemental skills injected via skill-resolver only — satisfies TD-020–TD-024
- [x] 2.2 Write `skills/project-conventions/references/code-style.md`: English artifacts rule (code, identifiers, comments, UI copy, docs, commit messages); YAML frontmatter schema with all 4 required fields; section heading order; Markdown table conventions for gates and constraints — satisfies TD-025–TD-028
- [x] 2.3 Write `skills/project-conventions/references/testing.md`: `strict_tdd: false` (exact text); no automated test runner; verification via `test-design.md`; update trigger note (file MUST be updated if strict_tdd or test commands change in `openspec/config.yaml`); if strict_tdd becomes true, sdd-apply follows `strict-tdd.md` — satisfies TD-029–TD-033
- [x] 2.4 Write `skills/project-conventions/references/review-readiness.md`: 400-line review budget; auto-chain + stacked-to-main delivery strategy; changes over 400 lines must use `chained-pr` skill; PR conventions (issue-first, conventional commits, no AI attribution) — satisfies TD-034–TD-038
- [x] 2.5 Write `skills/project-conventions/references/secure-design.md` with all 9 surface areas using narrative 6-part structure (design commitment / prohibited pattern / expected safe evidence / evidence owner / residual risk / exception policy): (1) changed-surface classification; (2) data sensitivity levels; (3) access control; (4) input/output handling; (5) secrets/config; (6) logging/error handling; (7) dependencies/external interfaces; (8) operational security; (9) evidence expectations per phase — satisfies TD-039–TD-041
- [x] 2.6 Ensure `secure-design.md` contains no SEC-* patterns, no "Source ID" rows, no Yes/No/N/A scoring columns, and all secret-like examples use obviously-fake placeholders (e.g., `<YOUR_SECRET_HERE>`) — satisfies TD-042–TD-044, TD-051

## Phase 3: Integration Update

- [x] 3.1 Append `## Project Skills` section to `AGENTS.md` registering `project-conventions` with: its trigger (supplemental skill loaded by SDD phases when conventions are needed), and an explicit note that it is NOT a required SDD DAG phase — satisfies TD-045–TD-047

## Phase 4: Verification

- [ ] 4.1 Verify all 7 files exist and are non-empty: `SKILL.md`, `references/architecture.md`, `references/code-style.md`, `references/testing.md`, `references/review-readiness.md`, `references/secure-design.md`, `AGENTS.md` — satisfies TD-001, TD-020, TD-025, TD-029, TD-034, TD-039, TD-045
- [ ] 4.2 Static grep: `rg "SEC-[A-Z0-9]+" skills/project-conventions/references/secure-design.md` → must return no matches — satisfies TD-042
- [ ] 4.3 Static grep: `rg "Source ID" skills/project-conventions/references/secure-design.md` → must return no matches — satisfies TD-043
- [ ] 4.4 Static grep across all reference files: `rg "SEC-|Source ID" skills/project-conventions/references/` → must return no matches — satisfies TD-050
- [ ] 4.5 Manual check: `design.md#Operational Considerations` contains exact text `No aplica.` — satisfies operational check
- [ ] 4.6 Manual check: no `review-security-report.json` was created or modified by this change — satisfies TD-049 boundary check
- [ ] 4.7 Manual check: `design.md` contains `## Implementation Quality Constraints` with `## Secure Design Constraints` table (Area | Project Secure Convention | Design Commitment | Evidence Owner) with at minimum Data handling, Secrets/config, Logging/errors rows — satisfies TD-052–TD-055
