## Change Archived

**Change**: test-cases-json
**Archived to**: `openspec/changes/archive/2026-07-11-test-cases-json/`

### Specs Synced
| Domain | Action | Details |
|--------|--------|---------|
| sdd-test-design-workflow | Updated | ADDED requirement: test-cases.json produced before test-design.md; JSON-first semantics merged from change delta
| sdd-execution-persistence-contracts | Updated | ADDED requirement: test-cases artifact reference and resolver row

### Archive Contents
- proposal.md ✅
- specs/ ✅
- design.md ✅
- design.md#secure-development-design ✅
- test-design.md ✅
- review-report.json ✅ (canonical, non-blocking)
- review-report.md ✅ (derived compatibility view)
- review-security-report.json ✅ (canonical, non-blocking)
- review-security-report.md ✅ (derived compatibility view)
- tasks.md ✅ (17/17 tasks complete)
- test-cases.json ✅ (32/32 verified)
- verify-report.md ✅ (PASS WITH WARNINGS)

### Source of Truth Updated
The following specs now reflect the new behavior:
- `openspec/specs/sdd-test-design-workflow/spec.md`
- `openspec/specs/sdd-execution-persistence-contracts/spec.md`

### Archive Audit Trail and Validation
- Artifact store mode: openspec
- Canonical review-report.json: `openspec/changes/test-cases-json/review-report.json` ✅
- Derived review-report.md: `openspec/changes/test-cases-json/review-report.md` ✅
- Canonical review-security-report.json: `openspec/changes/test-cases-json/review-security-report.json` ✅
- Derived review-security-report.md: `openspec/changes/test-cases-json/review-security-report.md` ✅
- Verify report: `openspec/changes/test-cases-json/verify-report.md` (PASS WITH WARNINGS — non-blocking) ✅
- Test cases canonical JSON: `openspec/changes/test-cases-json/test-cases.json` (32 cases verified) ✅
- Tasks: `openspec/changes/test-cases-json/tasks.md` (17/17 checked) ✅

### Source-Row Security Summary (preserved references)
- Catalog snapshot: `security-guidelines-initial-user-snapshot-2026-06-30`
- Expected Source ID count: 155
- Validated Source ID count: 155
- Coverage status: complete
- Exact-once coverage: true
- Safe evidence / N/A evidence: no-impact N/A evidence at `design.md#secure-development-design`

### Notes and Warnings
- Verification verdict: PASS WITH WARNINGS. Warning carried forward: repository tooling (test runner/linter/typechecker/formatter/coverage) unavailable and recorded explicitly in verify-report.md. This is non-blocking for archive.
- General review: PASS WITH WARNINGS (1 non-blocking finding: schema field drift) — recorded and carried forward.

### Archive Operation
- Moved `openspec/changes/test-cases-json/` → `openspec/changes/archive/2026-07-11-test-cases-json/`
- Verified archive folder contains proposal, specs, design, test-design, tasks, review-report.json, review-security-report.json, test-cases.json, and verify-report.md.
- Main specs merged: updated `openspec/specs/sdd-test-design-workflow/spec.md` and `openspec/specs/sdd-execution-persistence-contracts/spec.md` per delta. No destructive removals detected.

### Final Decision
Archive completed. next_recommended: none
