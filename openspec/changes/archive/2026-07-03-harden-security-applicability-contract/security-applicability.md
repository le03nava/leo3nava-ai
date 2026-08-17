# Security Applicability: Harden Security Applicability Contract

```yaml
schemaName: sdd.security-applicability
schemaVersion: 1
changeName: harden-security-applicability-contract
classification: security-impacting
securityImpact: true
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  taxonomyVersion: 1
taxonomyCategories:
  - authentication
  - sessions
  - sensitive-data-pan
  - secrets
  - permissions-access-control
  - files
  - database-access
  - sensitive-logging
applicableGuidelines:
  - SEC-AUTH-001
  - SEC-SESS-001
  - SEC-DATA-001
  - SEC-SECRET-001
  - SEC-ACCESS-001
  - SEC-FILE-001
  - SEC-DB-001
  - SEC-LOG-001
categoryDecisionMatrix:
  - category: authentication
    decision: applicable
    severity: blocking
    rationale: The change hardens the shared applicability contract that determines when authentication obligations become downstream controls.
    evidenceRefs: [proposal.md#scope, specs/sdd-security-applicability-workflow/spec.md#complete-category-decision-matrix, security-design.md#controls]
    guidelineIds: [SEC-AUTH-001]
    sourceIds: [1.1-1.10, 2.1-2.23, 6.3, 14.8]
  - category: sessions
    decision: applicable
    severity: blocking
    rationale: The security-design workflow must preserve session obligations when enriched applicability metadata marks session changes applicable.
    evidenceRefs: [specs/sdd-security-design-workflow/spec.md#enriched-applicability-consumption, security-design.md#controls]
    guidelineIds: [SEC-SESS-001]
    sourceIds: [7.1-7.13]
  - category: sensitive-data-pan
    decision: applicable
    severity: blocking
    rationale: The hardened no-impact proof distinguishes absent evidence from affirmative sensitive-data and PAN no-impact decisions.
    evidenceRefs: [proposal.md#problem, specs/sdd-security-applicability-workflow/spec.md#explicit-no-impact-proof, security-design.md#controls]
    guidelineIds: [SEC-DATA-001]
    sourceIds: [4.1, 13.1-13.9, 15.1-15.2]
  - category: secrets
    decision: applicable
    severity: blocking
    rationale: The override and validator rules must prevent configuration from weakening secret-handling source coverage or blocking obligations.
    evidenceRefs: [specs/sdd-security-applicability-workflow/spec.md#supported-applicability-overrides, security-design.md#controls]
    guidelineIds: [SEC-SECRET-001]
    sourceIds: [2.1, 4.2, 4.8, 5.5, 6.1, 13.5]
  - category: permissions-access-control
    decision: applicable
    severity: blocking
    rationale: The change governs authorization and least-privilege applicability routing through enriched downstream security-design consumption.
    evidenceRefs: [specs/sdd-security-design-workflow/spec.md#enriched-applicability-consumption, security-design.md#controls]
    guidelineIds: [SEC-ACCESS-001]
    sourceIds: [1.4, 6.2-6.4, 6.12, 13.1, 14.1-14.9]
  - category: files
    decision: applicable
    severity: blocking
    rationale: Complete category evaluation and source validation must preserve file-handling obligations for future upload, download, path, or generated-file changes.
    evidenceRefs: [specs/sdd-security-applicability-workflow/spec.md#complete-category-decision-matrix, security-design.md#controls]
    guidelineIds: [SEC-FILE-001]
    sourceIds: [9.1-9.12, 14.7]
  - category: database-access
    decision: applicable
    severity: blocking
    rationale: The validator formalizes database guideline and Source ID resolution while staying scoped to static Markdown and YAML references.
    evidenceRefs: [design.md#architecture-decisions, specs/sdd-security-guideline-catalog/spec.md#catalog-validator-contract, security-design.md#controls]
    guidelineIds: [SEC-DB-001]
    sourceIds: [5.1-5.12, 11.1-11.16, 12.2]
  - category: sensitive-logging
    decision: applicable
    severity: blocking
    rationale: Operational severity and validation metadata affect whether logging and redaction obligations remain blocking or downstream-visible.
    evidenceRefs: [specs/sdd-security-guideline-catalog/spec.md#operational-severity-vocabulary, security-design.md#controls]
    guidelineIds: [SEC-LOG-001]
    sourceIds: [3.1-3.11, 8.1-8.5]
noImpactProof:
  status: not-applicable
  reason: Classification is security-impacting with securityImpact true; all supported categories are intentionally evaluated as applicable for this governance-contract change.
  evidenceRefs: [security-applicability.md#classification-rationale, security-design.md#controls]
overridesApplied: []
validation:
  validator: scripts/validate_security_applicability.ps1
  status: pass
  checkedAt: 2026-07-03T08:32:13-06:00
  notes: Manual remediation run for review blocker REV-CORP-096; command evidence is recorded in tasks.md.
evidenceSummary:
  - Proposal intent requires audit-grade security decisions with explicit category evaluation, source coverage, operational blocking metadata, overrides, and static validation.
  - Proposal scope modifies the shared security applicability schema, the security guideline catalog, the security design workflow, and validator expectations.
  - Applicability workflow spec requires every supported security taxonomy category to be evaluated and blocks unknown blocking decisions.
  - Guideline catalog spec makes SEC-* to Source ID mappings formal corporate source coverage and defines blocking/conditional/advisory operational severity.
  - Security design workflow spec consumes enriched applicability fields and translates applicable obligations into downstream controls, evidence, risks, or approved exceptions.
designChangingUnknowns: []
nonBlockingRisks:
  - Exact validator implementation path is deferred to technical design; proposal lists it as TBD.
  - Current catalog has Source IDs in the full corporate snapshot, but compact SEC-* rows do not yet contain formal source coverage columns; this change is intended to add that mapping.
nextRecommended: design
```

## Classification Rationale

This change is **security-impacting**. It does not add an application runtime capability, but it changes the governance contract that decides whether future SDD changes require security controls, which guideline IDs are authoritative, which corporate Source IDs prove coverage, and when applicability must block.

Because the proposal and specs explicitly modify category evaluation, source mapping, severity metadata, validator behavior, and downstream security-design consumption, a defect in this change could weaken or incorrectly route controls across every supported taxonomy category. Under the current catalog, the safest auditable mapping is therefore all supported categories and their compact `SEC-*` guidelines.

## Guideline Mapping

| Guideline ID | Category | Mandatory When Applicable | Evidence | Source Coverage Evidence |
| --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | Yes | The applicability workflow must evaluate every category; governance changes affect whether authentication changes map to required controls. | Source `1.1`-`1.10`, `2.1`-`2.23`, `6.3`, `14.8` |
| `SEC-SESS-001` | `sessions` | Yes | Security design must consume enriched applicability and translate session obligations when applicable. | Source `7.1`-`7.13` |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | The catalog spec makes source coverage formal for sensitive data/PAN guidelines and no-impact proof cannot be inferred from absent evidence. | Source `4.1`, `13.1`-`13.9`, `15.1`-`15.2` |
| `SEC-SECRET-001` | `secrets` | Yes | The validator and source coverage rules must preserve secret-handling obligations when future changes touch credentials, keys, tokens, or configuration secrets. | Source `2.1`, `4.2`, `4.8`, `5.5`, `6.1`, `13.5` |
| `SEC-ACCESS-001` | `permissions-access-control` | Yes | The change governs category decisions and downstream control translation for authorization, least privilege, ownership, and protected-resource access. | Source `1.4`, `6.2`-`6.4`, `6.12`, `13.1`, `14.1`-`14.9` |
| `SEC-FILE-001` | `files` | Yes | Complete category evaluation and formal source mapping must preserve file-handling controls when uploads, downloads, paths, or generated files are touched. | Source `9.1`-`9.12`, `14.7` |
| `SEC-DB-001` | `database-access` | Yes | Catalog validation must keep database guideline IDs, categories, and source refs resolvable for future database-access changes. | Source `5.1`-`5.12`, `11.1`-`11.16`, `12.2` |
| `SEC-LOG-001` | `sensitive-logging` | Yes | Operational severity and validator behavior affect whether logging/redaction obligations become blocking risks or downstream security-design controls. | Source `3.1`-`3.11`, `8.1`-`8.5` |

## No-Impact Evidence

Not applicable because this artifact is explicitly `security-impacting` with `securityImpact: true`. The hardened `noImpactProof` field is present only to state that no-impact routing is unavailable for this change; all supported categories are evaluated as applicable in the decision matrix.

## Blocking Unknowns

None. The proposal and specs provide enough evidence to classify the change as security-impacting. The remaining validator implementation path is a technical design detail, not a classification blocker.

## Security-Design Risks

- Technical design must prevent the validator from becoming broader than static Markdown/YAML structure and reference validation, matching the proposal mitigation.
- Technical design must define exact compact `SEC-*` to corporate Source ID mappings without creating false audit confidence.
- Technical design must preserve no-impact routing compatibility so enriched applicability fields do not require `security-design.md` for valid no-impact changes.
