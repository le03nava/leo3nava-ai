# Apply Evidence: Corporate Source Row Security Validation — WU4 Static Evidence

schemaName: sdd.apply-evidence
schemaVersion: 1
changeName: corporate-source-row-security-validation
artifactStore: openspec
workUnit: WU4 — verify/archive/source spec sync/final static evidence
evidenceType: apply-time-static-manual
formalVerifyEvidence: false
verdict: PASS WITH WARNINGS
nextRecommended: review

## Scope

This non-formal apply evidence records WU4 static/manual evidence for tasks 4.1-4.4. Full SDD `sdd-verify` still requires non-blocking `review-report.md` and `review-security-report.md` after apply/review phases; this artifact is not a formal verify report and does not bypass those gates.

## Completeness

| Area | Status | Evidence |
| --- | --- | --- |
| Tasks | Pass | `openspec/changes/corporate-source-row-security-validation/tasks.md` marks 13/13 tasks complete, preserving WU1-WU3 evidence and adding WU4 evidence. |
| Verify contract | Pass | `skills/sdd-verify/SKILL.md` and `agents/sdd/sdd-verify.md` consume non-blocking review-security source-row verdicts, preserve warning summaries, block unresolved source blockers, and report unavailable tooling. |
| Archive contract | Pass | `skills/sdd-archive/SKILL.md` and `agents/sdd/sdd-archive.md` preserve source-row coverage, compact mappings, warnings, exceptions, safe evidence refs, `N/A` status, and legacy-artifact non-requirement. |
| Source specs | Pass | Source specs under `openspec/specs/` now include catalog, design, test-design, review-security, verify/archive, and persistence source-row requirements. |

## Source-Row Evidence Summary

| Check | Result |
| --- | --- |
| Expanded coverage | Pass — contracts preserve the 155 expanded Source ID universe across 15 corporate sections and require exact-once validation. |
| Compact mapping | Pass — every source row must map to one or more existing compact `SEC-*` IDs; no replacement controls were added. |
| Safe evidence | Pass — contracts allow only review-safe paths, section refs, sanitized summaries, command summaries, or redacted placeholders and block unsafe values. |
| N/A policy | Pass — `N/A` requires evidence plus justification; unsupported `N/A` routes to `resolve-blockers`. |
| Review boundary | Pass — verify/archive cite review-security source-row verdicts and warnings without duplicating the full source-row matrix or 96-control general review matrix. |
| Legacy validator | Pass — no active gate requires standalone `security-design.md` or `scripts/validate_security_design.ps1`. |

## Tooling Availability

| Tool | Availability | Evidence Treatment |
| --- | --- | --- |
| Runtime tests | Unavailable | Reported unavailable; not marked passed. |
| Build command | Unavailable | Reported unavailable; not marked passed. |
| Linter | Unavailable | Reported unavailable; not marked passed. |
| Type checker | Unavailable | Reported unavailable; not marked passed. |
| Formatter | Unavailable | Reported unavailable; not marked passed. |
| Coverage | Unavailable | Reported unavailable; not marked passed. |

Evidence source: `openspec/config.yaml#testing` declares no runtime test runner, build command, linter, type checker, formatter, or coverage command.

## Warnings

- Static/manual evidence only. This repository has no configured runtime/build/lint/type/format/coverage runner.
- Full archive readiness still depends on downstream non-blocking general review, non-blocking security review, and a formal verify phase consuming those reports.
- This artifact intentionally uses a non-formal apply-evidence filename so native status tooling does not treat it as `sdd-verify` output.

## CRITICAL Findings

None for WU4 static/manual apply evidence.
