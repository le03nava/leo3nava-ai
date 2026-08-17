# sdd-design-workflow Specification

## Purpose

Define `sdd-design` as the mandatory technical design phase that owns narrative secure development planning for new changes.

## Requirements

### Requirement: Embedded Secure Development Design

`sdd-design` MUST persist `design.md` for every new active change. The artifact MUST include a narrative, rule-based `## Secure Development Design` section that is the active design-time security planning authority. The section MUST classify the changed security surface from proposal, specs, design context, and known code/artifact impact, then include human-readable subsections only for applicable security categories such as Authentication, Sessions, Sensitive Data/PAN, Secrets, Permissions/Access Control, Files, Database Access, or Sensitive Logging. It MUST cite catalog context in prose. It MUST NOT use YAML, JSON, schema blocks, control tables, compact matrices, Source ID matrices, machine-readable applicability fields, exhaustive applicability inventories, or per-row `N/A` bookkeeping.

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

### Requirement: Secure Development Design Category Rule Contract

Applicable secure design content MUST be category development rules, not rows. Each applicable category MUST state implementation obligations, prohibited unsafe patterns, evidence owner, expected evidence, residual risk, exception handling, and safe-evidence policy. Omitted categories are reviewable omissions, not design-passing `N/A` rows.

#### Scenario: Applicable guideline has obligations

- GIVEN a guideline applies to the change
- WHEN `sdd-design` records that category
- THEN it MUST dictate implementation rules and downstream evidence expectations
- AND unresolved mandatory evidence MUST remain visible to test design, review-security, verify, and archive.

#### Scenario: Exception is required

- GIVEN mandatory evidence cannot be planned
- WHEN the category rule records an exception
- THEN it MUST include approver, accepted-risk rationale, mitigation or follow-up, and evidence gap.

### Requirement: Direct Routing to Test Design

For new changes, successful `sdd-design` MUST route directly to `sdd-test-design`; it MUST NOT require, launch, or produce a standalone security-design phase or artifact.

#### Scenario: Design routes to test design

- GIVEN `design.md` includes the secure development section
- WHEN the phase returns a successful envelope
- THEN `next_recommended` MUST be `sdd-test-design` or `test-design`
- AND no standalone security-design artifact MUST be required for the new change.

### Requirement: Secure Design Source ID Coverage

`sdd-design` MUST use catalog context only as human-readable category guidance. It MAY cite catalog snapshot identity/path and category names, but MUST NOT copy, expand, validate, or require Source IDs, per-Source-ID coverage, compact mappings, lifecycle fields, or `N/A` evidence. Exhaustive compact-control, Source ID, applicability, and `N/A` validation belong to canonical `review-security-report.json`, with Markdown as a derived compatibility view.

#### Scenario: Source IDs are not planned in design

- GIVEN proposal and specs require security validation
- WHEN `sdd-design` succeeds
- THEN design completion MUST NOT require Source ID rows, Source ID coverage, or per-row `N/A` evidence
- AND exhaustive validation ownership MUST point to canonical `review-security-report.json`.

### Requirement: Design Preserves Compact Summary

`sdd-design` MUST NOT preserve a compact control summary as a required matrix for new active changes. It MUST preserve security intent through narrative applicable category rules. It MUST NOT duplicate the 96-control review matrix, replace category rules with per-source controls, require all compact controls, materialize the 155-row security matrix, or encode machine-readable applicability fields.

#### Scenario: Category rules replace compact matrix

- GIVEN a change has multiple corporate Source IDs
- WHEN reviewers inspect secure design
- THEN applicable category development rules MUST be visible
- AND compact or Source ID matrix evidence MUST NOT be required in design.

### Requirement: Operational Readiness Planning

`sdd-design` MAY include `## Operational Considerations` when the change design makes operational behavior applicable. The section MUST be design-owned and conditional, not a mandatory readiness gate. When present, it SHOULD cover applicable logs/errors, monitoring, administration operations, reprocessing/recovery, and backup/retention/cleanup/generated-artifact concerns. It MUST define expected downstream evidence only for applicable concerns and MUST use safe evidence, `Pendiente de confirmar:`, or `No aplica.` without inventing data.

#### Scenario: Applicable operational behavior is designed

- GIVEN proposal and specs show operational behavior applies
- WHEN `sdd-design` succeeds
- THEN `design.md` SHOULD include `## Operational Considerations`
- AND it MUST describe expected evidence for applicable concerns only.

#### Scenario: Operational behavior does not apply

- GIVEN operational considerations are out of scope
- WHEN `design.md` is written
- THEN the section MAY be omitted or state `No aplica.` with rationale
- AND downstream phases MUST NOT require readiness completeness.

#### Scenario: Production data is unavailable

- GIVEN an applicable operational field lacks user-provided data
- WHEN design records the consideration
- THEN it MUST use `Pendiente de confirmar:` instead of inventing data.

### Requirement: Operational Evidence Safety in Design

Design MUST keep safe SDD evidence separate from final operational documentation. It MUST NOT include production hostnames, IPs, ports, SID/service names, credentials, tokens, payloads, full ID lists, or generated file bytes unless the user explicitly provided them for final operational documentation. This safety boundary applies only to operational evidence that the design includes or references.

#### Scenario: Restricted identifier is needed later

- GIVEN production connection details are needed for handoff
- WHEN design lacks explicit user-provided values
- THEN design MUST record the gap safely
- AND final documentation MAY receive those values only from the user.
