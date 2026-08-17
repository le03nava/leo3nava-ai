# Review Security Report: project-conventions-skill

## Verdict

| Field | Value |
| --- | --- |
| Change | `project-conventions-skill` |
| Status | success |
| Verdict | PASS |
| Blocking findings | 0 |
| Non-blocking warnings | 0 |
| Next recommendation | verify |
| JSON authority | `openspec/changes/project-conventions-skill/review-security-report.json` |
| Markdown authority | derived compatibility view |

## Source References

- Secure design: `openspec/changes/project-conventions-skill/design.md#secure-development-design`
- Test design: `openspec/changes/project-conventions-skill/test-design.md`
- Tasks/apply evidence: `openspec/changes/project-conventions-skill/tasks.md`, `openspec/changes/project-conventions-skill/apply-progress.md`
- Changed-file context: `skills/project-conventions/SKILL.md`, `skills/project-conventions/references/architecture.md`, `skills/project-conventions/references/code-style.md`, `skills/project-conventions/references/testing.md`, `skills/project-conventions/references/review-readiness.md`, `skills/project-conventions/references/secure-design.md`, `AGENTS.md`
- General review JSON: `openspec/changes/project-conventions-skill/review-report.json`
- General review Markdown compatibility: `openspec/changes/project-conventions-skill/review-report.md`
- Catalog JSON: `skills/sdd-review-security/references/security-guideline-catalog.operational.json` snapshot `security-guidelines-initial-user-snapshot-2026-06-30`
- Catalog human view: `skills/sdd-review-security/references/security-guideline-catalog.md`

## General Review Handoff

General review verdict: **PASS** with 0 blocking failures and 0 non-blocking findings. Changed files are documentation-only Markdown. secure-design.md provides narrative guidance with no SEC-* matrices, no Source ID rows, no scoring tables. No operational evidence applicable — documentation-only change with no runtime, deployment, or monitoring surface.

## Source Row Navigation

| Grouping field | Value | Expected | Validated | Blockers | Warnings | N/A | Notes |
| --- | --- | ---: | ---: | ---: | ---: | ---: | --- |
| controlDomain | identity-authentication | 10 | 10 | 0 | 0 | 10 | No authentication in documentation-only change |
| controlDomain | credential-secrets | 23 | 23 | 0 | 0 | 23 | No password/credential handling |
| controlDomain | observability-logging | 11 | 11 | 0 | 0 | 11 | No logging in documentation-only change |
| controlDomain | cryptography-data-protection | 8 | 8 | 0 | 0 | 8 | No cryptography in documentation-only change |
| controlDomain | database-access | 12 | 12 | 0 | 0 | 12 | No database access |
| controlDomain | secure-coding | 14 | 14 | 0 | 0 | 14 | No executable code in this change |
| controlDomain | session-management | 13 | 13 | 0 | 0 | 13 | No session management |
| controlDomain | safe-error-handling | 5 | 5 | 0 | 0 | 5 | No error handling code |
| controlDomain | file-handling | 12 | 12 | 0 | 0 | 12 | No file upload/download/generation |
| controlDomain | memory-safety | 6 | 6 | 0 | 0 | 6 | No memory management |
| controlDomain | input-validation | 16 | 16 | 0 | 0 | 16 | No user input processing |
| controlDomain | output-encoding | 5 | 5 | 0 | 0 | 5 | No output encoding |
| controlDomain | sensitive-data-protection | 9 | 9 | 0 | 0 | 9 | No sensitive data processing |
| controlDomain | authorization-access-control | 9 | 9 | 0 | 0 | 9 | No access control |
| controlDomain | pan-test-data | 2 | 2 | 0 | 0 | 2 | No PAN data |

## Source Row Summary

Expected Source ID count: 155. Validated Source ID count: 155. Coverage: complete. Exact once: true.

## Grouped Non-Applicability

| Grouping Field | Group | Source IDs / Count | Justification | Evidence Type | Evidence Location | Owner Phase | Route |
| --- | --- | --- | --- | --- | --- | --- | --- |
| corporateSection | All 15 sections (1-15) | 155 | Documentation-only Markdown skill files with no runtime behavior, no data processing, no authentication, no secrets, no network calls. Design classified as no-impact. | n/a-evidence | openspec/changes/project-conventions-skill/design.md#secure-development-design | review-security | verify |

## Blockers and Warnings

### Blockers

None.

### Warnings

None.

### Unsafe Evidence Rejections

None.

### Warning Carry-Forward

None.

## Exceptions

None.

## Unavailable Tooling

| Tool | Available | Evidence |
| --- | --- | --- |
| testRunner | No | strict_tdd: false; no test runner configured |
| build | No | No build tool; documentation-only Markdown |
| linter | No | No Markdown linter configured |
| typeChecker | No | No type checker for Markdown |
| formatter | No | No formatter configured |
| coverage | No | No coverage tool configured |

## Artifact Metadata

| Check | Result |
| --- | --- |
| Canonical JSON ref | `openspec/changes/project-conventions-skill/review-security-report.json` |
| Derived Markdown ref | `openspec/changes/project-conventions-skill/review-security-report.md` |
| JSON persisted/read back | true / true |
| Markdown generated/read back | true / true |
| JSON/Markdown parity | consistent |
| JSON authority | canonical |
| Markdown authority | derived |

## Recommendation

- Next recommendation: `verify`
- Follow-up: Security review passed with no blockers. All 155 source rows validated as N/A for this documentation-only change. Proceed to verification phase.

## Full Source Row Matrix


| Source ID | Corporate Section | Control Domain | Applies | Complies | Lifecycle Status | Evidence Type | Finding | Owner Phase | Route |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1.1 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.2 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.3 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.4 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.5 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.6 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.7 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.8 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.9 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 1.10 | 1. Authentication | identity-authentication | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.1 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.2 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.3 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.4 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.5 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.6 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.7 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.8 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.9 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.10 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.11 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.12 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.13 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.14 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.15 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.16 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.17 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.18 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.19 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.20 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.21 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.22 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 2.23 | 2. Passwords | credential-secrets | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.1 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.2 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.3 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.4 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.5 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.6 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.7 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.8 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.9 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.10 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 3.11 | 3. Access and Activity Logging | observability-logging | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.1 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.2 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.3 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.4 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.5 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.6 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.7 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 4.8 | 4. Cryptography | cryptography-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.1 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.2 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.3 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.4 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.5 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.6 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.7 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.8 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.9 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.10 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.11 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 5.12 | 5. Databases | database-access | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.1 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.2 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.3 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.4 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.5 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.6 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.7 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.8 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.9 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.10 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.11 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.12 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.13 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 6.14 | 6. Coding | secure-coding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.1 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.2 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.3 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.4 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.5 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.6 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.7 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.8 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.9 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.10 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.11 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.12 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 7.13 | 7. Session Management | session-management | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 8.1 | 8. Error Handling | safe-error-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 8.2 | 8. Error Handling | safe-error-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 8.3 | 8. Error Handling | safe-error-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 8.4 | 8. Error Handling | safe-error-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 8.5 | 8. Error Handling | safe-error-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.1 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.2 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.3 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.4 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.5 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.6 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.7 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.8 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.9 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.10 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.11 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 9.12 | 9. File Handling | file-handling | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.1 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.2 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.3 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.4 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.5 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 10.6 | 10. Memory Management | memory-safety | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.1 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.2 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.3 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.4 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.5 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.6 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.7 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.8 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.9 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.10 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.11 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.12 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.13 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.14 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.15 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 11.16 | 11. Input Validation | input-validation | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 12.1 | 12. Output Encoding | output-encoding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 12.2 | 12. Output Encoding | output-encoding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 12.3 | 12. Output Encoding | output-encoding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 12.4 | 12. Output Encoding | output-encoding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 12.5 | 12. Output Encoding | output-encoding | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.1 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.2 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.3 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.4 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.5 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.6 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.7 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.8 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 13.9 | 13. Data Protection | sensitive-data-protection | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.1 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.2 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.3 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.4 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.5 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.6 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.7 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.8 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 14.9 | 14. Access Control | authorization-access-control | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 15.1 | 15. PAN â€” Primary Account Number | pan-test-data | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |
| 15.2 | 15. PAN â€” Primary Account Number | pan-test-data | N/A | N/A | not-applicable | n/a-evidence | none | review-security | verify |

