# Security Review Report: token-monitor

**Verdict**: PASS WITH WARNINGS
**Status**: success
**Next recommended**: verify
**Generated**: 2026-07-13T15:43:53.822362Z

## Navigation

- [Security Review Summary](#security-review-summary)
- [Warnings](#warnings)
- [Source Row Matrix](#source-row-matrix)

## Security Review Summary

| Metric | Value |
|--------|-------|
| Source rows validated | 155 / 155 |
| Applicable | 0 |
| Not applicable | 155 |
| Blockers | 0 |
| Warnings | 1 |

### By Control Domain

| Control Domain | Total | N/A | Blockers |
|----------------|-------|-----|----------|
| authorization-access-control | 9 | 9 | 0 |
| credential-secrets | 23 | 23 | 0 |
| cryptography-data-protection | 8 | 8 | 0 |
| database-access | 12 | 12 | 0 |
| file-handling | 12 | 12 | 0 |
| identity-authentication | 10 | 10 | 0 |
| input-validation | 16 | 16 | 0 |
| memory-safety | 6 | 6 | 0 |
| observability-logging | 11 | 11 | 0 |
| output-encoding | 5 | 5 | 0 |
| pan-test-data | 2 | 2 | 0 |
| safe-error-handling | 5 | 5 | 0 |
| secure-coding | 14 | 14 | 0 |
| sensitive-data-protection | 9 | 9 | 0 |
| session-management | 13 | 13 | 0 |

## Warnings

- **W-SEC-001**: Header values from intercepted HTTP traffic are stored in SQLite without length validation. Parameterized queries prevent SQL injection, but excessively long header values could bloat the database.
  - Location: token-monitor/token_monitor.py (extract_headers, detect_agent)
  - Recommendation: Consider truncating header values to a reasonable max length (e.g., 256 chars) before persistence.

## Source Row Matrix

| Source ID | Applies | Complies | Finding | Justification (truncated) |
|-----------|---------|----------|---------|---------------------------|
| 1.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 1.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.13 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.14 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.15 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.16 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.17 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.18 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.19 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.20 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.21 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.22 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 2.23 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 3.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 4.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 5.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.13 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 6.14 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 7.13 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 8.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 8.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 8.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 8.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 8.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 9.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 10.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.10 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.11 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.12 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.13 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.14 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.15 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 11.16 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 12.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 12.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 12.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 12.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 12.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 13.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.3 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.4 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.5 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.6 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.7 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.8 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 14.9 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 15.1 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
| 15.2 | N/A | N/A | none | Local single-user developer tool. No authentication, no PII/... |
