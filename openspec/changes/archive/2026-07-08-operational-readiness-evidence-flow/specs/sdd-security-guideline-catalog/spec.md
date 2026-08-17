# Delta for sdd-security-guideline-catalog

## ADDED Requirements

### Requirement: Operational Safe-Evidence Policy

The security guideline catalog MUST define safe-evidence policy for operational readiness. It MUST cover logs, monitoring mechanisms, administration, reprocessing, ownership, unresolved gaps, and final documentation boundaries. Monitoring evidence MUST be mechanism-oriented and MUST NOT be limited to SQL-only checks.

#### Scenario: Monitoring evidence is categorized

- GIVEN a change requires operational monitoring evidence
- WHEN the catalog is consulted
- THEN it MUST support monitoring mechanisms such as dashboards, alerts, jobs, traces, scripts, or documented manual checks
- AND it MUST NOT require SQL-only evidence.

#### Scenario: Operational gap is safe

- GIVEN operational evidence is unavailable
- WHEN catalog policy is applied
- THEN exact `Pendiente de confirmar:` MUST be an accepted safe state for ordinary SDD evidence.

### Requirement: Restricted Operational Data Classification

The catalog MUST classify production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, generated file bytes, and equivalent environment-specific operational details as restricted for ordinary SDD evidence, code, tests, fixtures, and examples. Final operational documentation MAY include user-provided restricted operational values.

#### Scenario: Restricted data appears in evidence

- GIVEN an SDD artifact cites restricted operational data
- WHEN catalog-based review-security validation runs
- THEN the evidence MUST be unsafe unless it is scoped to final user-provided operational documentation.

#### Scenario: Safe evidence is enough

- GIVEN a field is inapplicable
- WHEN catalog policy is applied
- THEN exact `No aplica.` with optional rationale MUST satisfy the safe-evidence state.
