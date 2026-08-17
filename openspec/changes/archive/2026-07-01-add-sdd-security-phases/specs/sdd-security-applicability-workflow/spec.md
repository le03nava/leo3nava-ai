# sdd-security-applicability-workflow Specification

## Purpose

Define the always-run security applicability phase that classifies whether an SDD change has security impact and records guideline mapping or no-impact evidence.

## Requirements

### Requirement: Always-Run Applicability Classification

The SDD workflow MUST run `sdd-security-applicability` after `sdd-spec` succeeds and before technical design proceeds. The phase MUST classify the change as security-impacting or no-impact using available proposal and spec evidence.

#### Scenario: Security-impacting change classified

- GIVEN an SDD change affects authentication, sessions, PAN, secrets, permissions, files, database access, or sensitive logging
- WHEN `sdd-security-applicability` runs
- THEN it MUST mark the change as security-impacting
- AND it MUST identify applicable guideline categories.

#### Scenario: No-impact change classified

- GIVEN an SDD change has no plausible security impact
- WHEN `sdd-security-applicability` runs
- THEN it MUST record explicit no-impact evidence
- AND downstream phases MUST treat the applicability artifact as complete.

### Requirement: Blocking and Risk Rules

The phase MUST block only when missing information could change security design decisions across authentication, sessions, sensitive data or PAN, secrets, permissions or access control, files, database access, or sensitive logging. Minor evidence gaps SHOULD continue as risks for `sdd-security-design`.

#### Scenario: Design-changing information is missing

- GIVEN security-relevant scope is ambiguous in a way that could change required controls
- WHEN applicability is evaluated
- THEN the phase MUST return blocked
- AND the blocker MUST name the missing decision area.

#### Scenario: Minor evidence gap exists

- GIVEN a security impact is known but a non-decisive detail is incomplete
- WHEN applicability is evaluated
- THEN the phase SHOULD continue
- AND it MUST record the gap as a security-design risk.

### Requirement: Artifact and Routing Contract

The phase MUST produce `security-applicability.md` with classification, evidence, guideline mapping, risks, and routing recommendation. Security-impacting changes MUST route to `sdd-security-design`; no-impact changes MUST skip security design and continue normal design workflow.

#### Scenario: Artifact drives conditional routing

- GIVEN `security-applicability.md` marks a change as security-impacting
- WHEN the orchestrator computes next phases
- THEN it MUST require `sdd-security-design`
- AND design-related successors MUST receive the applicability artifact reference.
