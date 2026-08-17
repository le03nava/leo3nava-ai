# SDD Verify Report Format

## Compliance Statuses

- ✅ `COMPLIANT`: covering test exists and passed.
- ❌ `FAILING`: covering test exists but failed.
- ❌ `UNTESTED`: no covering test found.
- ⚠️ `PARTIAL`: test passes but covers only part of the scenario.
- ⚠️ `WARNING`: non-mandatory planned case has no matching evidence.

Mandatory planned cases from `test-design.md` are verification-blocking: uncovered mandatory cases are ❌ `UNTESTED` and force verdict `FAIL`. Non-mandatory planned cases are advisory: uncovered non-mandatory cases are ⚠️ `WARNING` and MUST NOT fail verification by themselves.

## Report Template

~~~markdown
## Verification Report

**Change**: {change-name}
**Version**: {spec version or N/A}
**Mode**: {Strict TDD | Standard}

### Completeness
| Metric | Value |
|--------|-------|
| Tasks total | {N} |
| Tasks complete | {N} |
| Tasks incomplete | {N} |

### Build & Tests Execution
**Build**: ✅ Passed / ❌ Failed
```text
{build command and relevant output}
```

**Tests**: ✅ {N} passed / ❌ {N} failed / ⚠️ {N} skipped
```text
{test command and failure details}
```

**Coverage**: {N}% / threshold: {N}% → ✅ Above / ⚠️ Below / ➖ Not available

### Spec Compliance Matrix
| Requirement | Scenario | Test | Result |
|-------------|----------|------|--------|
| {REQ-01} | {Scenario} | `{file} > {test}` | ✅ COMPLIANT |
| {REQ-02} | {Scenario} | (none found) | ❌ UNTESTED |

**Compliance summary**: {N}/{total} scenarios compliant

### General Review Summary

Consumed from `review-report.md` / canonical `review-report.json`. Do not reproduce the full 96-control matrix.

**Verdict**: {PASS / PASS WITH WARNINGS / FAIL}
**Blocking failures**: {N} · **Non-blocking findings**: {N}
**Catalog**: `{review-control-catalog path}` — {N}/96 controls evaluated
**Top blocking findings**: {list or None}
**Source**: `{review-report path}`

### Security Review Summary

Consumed from `review-security-report.md` / canonical `review-security-report.json`. Do not reproduce the full 155-row matrix.

| Control Domain | Passing/Total | Blockers |
| --- | --- | --- |
| {control-domain} | {N}/{total} | {count or 0} |

**Overall**: {totals.passing}/{totals.sourceRowCount} passing · {totals.blockers} blockers · {totals.warnings} warnings · {totals.exceptions} exceptions
**Source**: `{review-security-report path}` — verdict: {PASS / PASS WITH WARNINGS / FAIL}

### Test-Design Coverage Matrix
| Case ID | Source | Severity | Expected Evidence | Observed Evidence | Result |
|---------|--------|----------|-------------------|-------------------|--------|
| {TD-01} | {Scenario/Risk} | mandatory | {expected evidence} | {test/apply evidence/justified skip} | ✅ COMPLIANT |
| {TD-02} | {Scenario/Risk} | mandatory | {expected evidence} | (none found) | ❌ UNTESTED |
| {TD-03} | {Scenario/Risk} | non-mandatory | {expected evidence} | (none found) | ⚠️ WARNING |

**Test-design summary**: {N}/{mandatory_total} mandatory cases covered; {M} non-mandatory warnings.

### Correctness (Static Evidence)
| Requirement | Status | Notes |
|------------|--------|-------|
| {Req name} | ✅ Implemented | {brief note} |

### Coherence (Design)
| Decision | Followed? | Notes |
|----------|-----------|-------|
| {Decision} | ✅ Yes | |

### Skipped / Degraded Dimensions
| Dimension | Status | Reason |
|-----------|--------|--------|
| {Spec correctness / Design coherence / Runtime evidence} | Full / Partial / Skipped | {why this dimension was degraded or skipped} |

### Issues Found
**CRITICAL**: {list or None}
**WARNING**: {list or None}
**SUGGESTION**: {list or None}

### Verdict
{PASS / PASS WITH WARNINGS / FAIL}
{one-line reason}
~~~

When Strict TDD is active, insert the TDD compliance, test layer distribution, changed-file coverage, and quality metrics sections from `strict-tdd-verify.md`.
