# Delta for sdd-design-workflow

## MODIFIED Requirements

### Requirement: Embedded Secure Development Design

`sdd-design` MUST persist `design.md` for every new active change. It MUST include a narrative, rule-based `## Secure Development Design` section. The section MUST classify the changed security surface, then include human-readable subsections only for applicable categories touched by the design, such as Authentication, Sessions, Sensitive Data/PAN, Secrets, Permissions/Access Control, Files, Database Access, or Sensitive Logging. Each subsection MUST dictate development rules, evidence owners, residual risks, exceptions, and safe-evidence policy. New active designs MUST NOT use YAML, JSON, schema blocks, control tables, compact or Source ID matrices, exhaustive applicability inventories, or per-row `N/A` bookkeeping.
(Previously: design evaluated compact rows and allowed matrix-style applicability evidence.)

#### Scenario: Design contains applicable category rules

- GIVEN proposal and specs are readable
- WHEN `sdd-design` succeeds
- THEN `design.md` MUST include `## Secure Development Design`
- AND the section MUST contain narrative rule subsections only for applicable categories.

#### Scenario: No-impact change records rationale

- GIVEN no security category applies
- WHEN `design.md` is written
- THEN it MUST record no-impact classification rationale and changed-surface inventory
- AND it MUST NOT include YAML, schema, matrix, or all-row `N/A` content.

### Requirement: Secure Development Design Row Contract

Applicable secure design content MUST be category development rules, not rows. Each applicable category MUST state implementation obligations, prohibited unsafe patterns, evidence owner, expected evidence, residual risk, exception handling, and safe-evidence policy. Omitted categories are reviewable omissions, not design-passing `N/A` rows.
(Previously: applicable and N/A security content used row-like fields.)

#### Scenario: Applicable category has obligations

- GIVEN a category applies to the change
- WHEN `sdd-design` records that category
- THEN it MUST dictate implementation rules and downstream evidence expectations
- AND unresolved mandatory evidence MUST remain visible to test design, review-security, verify, and archive.

### Requirement: Secure Design Source ID Coverage

`sdd-design` MUST use catalog context only as human-readable category guidance. It MAY cite catalog snapshot identity/path and category names, but MUST NOT copy, expand, validate, or require Source IDs, per-Source-ID coverage, or `N/A` evidence. Exhaustive compact-control, Source ID, applicability, and `N/A` validation belong to `review-security-report.md`.
(Previously: design summarized applicable mappings and Source ID context.)

#### Scenario: Source IDs are not planned in design

- GIVEN a Source ID, section, or compact mapping may be relevant
- WHEN design is persisted
- THEN design completion MUST NOT require Source ID rows, Source ID coverage, or per-row `N/A` evidence
- AND exhaustive validation ownership MUST point to `review-security-report.md`.

### Requirement: Design Preserves Compact Summary

`sdd-design` MUST NOT preserve a compact control summary as a required matrix for new active changes. It MUST preserve security intent through narrative applicable category rules. It MUST NOT duplicate the 96-control review matrix, replace category rules with per-source controls, require all compact controls, materialize the 155-row security matrix, or encode machine-readable applicability fields.
(Previously: design preserved a readable compact `SEC-*` summary.)

#### Scenario: Category rules replace compact matrix

- GIVEN a change touches authentication or sensitive data
- WHEN reviewers inspect secure design
- THEN applicable category development rules MUST be visible
- AND compact or Source ID matrix evidence MUST NOT be required in design.
