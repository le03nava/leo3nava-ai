# Security Design: Remove SDD Security Applicability Executor

```yaml
schemaName: sdd.security-design
schemaVersion: 1
changeName: remove-sdd-security-applicability-executor
sourceApplicability: direct-classification-from-proposal-specs-design
classification: security-impacting
securityImpact: true
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  taxonomyVersion: 1
  source: skills/_shared/security-guideline-catalog.md
classificationRationale: >-
  This change is security-impacting at the SDD governance/workflow-contract level because it removes the repo-local
  security-applicability executor and moves new-change classification authority into mandatory security-design.md.
  It does not introduce runtime application behavior, secrets, authentication, sessions, database access, file handling,
  or logging behavior.
sourceRefs:
  - openspec/changes/remove-sdd-security-applicability-executor/proposal.md
  - openspec/changes/remove-sdd-security-applicability-executor/design.md
  - openspec/changes/remove-sdd-security-applicability-executor/specs/sdd-security-applicability-workflow/spec.md
  - openspec/changes/remove-sdd-security-applicability-executor/specs/sdd-execution-persistence-contracts/spec.md
  - openspec/changes/remove-sdd-security-applicability-executor/specs/sdd-security-design-workflow/spec.md
  - openspec/changes/remove-sdd-security-applicability-executor/specs/sdd-review-workflow/spec.md
categoryDecisionMatrix:
  - category: authentication
    decision: not-applicable
    severity: blocking
    rationale: No login, identity proofing, credential validation, MFA, impersonation, or account recovery behavior changes.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-AUTH-001
    sourceIds:
      - 1.1-1.10
      - 2.1-2.23
      - 6.3
      - 14.8
  - category: sessions
    decision: not-applicable
    severity: blocking
    rationale: No cookies, tokens, refresh flows, session lifetime, revocation, storage, or fixation controls are changed.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-SESS-001
    sourceIds:
      - 7.1-7.13
  - category: sensitive-data-pan
    decision: not-applicable
    severity: blocking
    rationale: No PAN, PII, confidential runtime data, retention, masking, transmission, or storage behavior is introduced.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-DATA-001
    sourceIds:
      - 4.1
      - 13.1-13.9
      - 15.1-15.2
  - category: secrets
    decision: not-applicable
    severity: blocking
    rationale: No API keys, passwords, certificates, tokens, encryption keys, or secret rotation behavior is added or modified.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-SECRET-001
    sourceIds:
      - 2.1
      - 4.2
      - 4.8
      - 5.5
      - 6.1
      - 13.5
  - category: permissions-access-control
    decision: not-applicable
    severity: blocking
    rationale: No runtime roles, ownership checks, authorization decisions, privilege boundaries, or protected resources are changed.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-ACCESS-001
    sourceIds:
      - 1.4
      - 6.2-6.4
      - 6.12
      - 13.1
      - 14.1-14.9
  - category: files
    decision: not-applicable
    severity: blocking
    rationale: No runtime uploads, downloads, generated exports, path handling, file metadata, or content validation behavior is introduced.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-FILE-001
    sourceIds:
      - 9.1-9.12
      - 14.7
  - category: database-access
    decision: not-applicable
    severity: blocking
    rationale: No queries, migrations, persistence, tenant isolation, reporting paths, or background database jobs are changed.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-DB-001
    sourceIds:
      - 5.1-5.12
      - 11.1-11.16
      - 12.2
  - category: sensitive-logging
    decision: not-applicable
    severity: blocking
    rationale: No runtime logs, traces, metrics, analytics, error reporting, or audit trail payloads are added or changed.
    evidenceRefs:
      - proposal.md Scope/Out of Scope
      - design.md Technical Approach
    guidelineIds:
      - SEC-LOG-001
    sourceIds:
      - 3.1-3.11
      - 8.1-8.5
controls:
  - controlId: SDD-GOV-001
    taxonomyCategory: workflow-governance
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs:
      - proposal.md Intent
      - design.md Interfaces / Contracts
      - specs/sdd-security-design-workflow/spec.md Requirement: Enriched Applicability Consumption
    requiredControl: >-
      New changes must use security-design.md as the mandatory classification authority and must not require or recreate
      security-applicability.md for classification.
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan static checks proving active routing reaches security-design and never launches sdd-security-applicability.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Changed agents, shared contracts, and specs remove active applicability executor references while preserving legacy data references.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Text-search and diff evidence proving launchable applicability agent/skill sources are absent and security-design remains mandatory.
    residualRisk: Stale global opencode copies outside this repository may remain launchable until a separate sync cleanup policy removes them.
    exception: null
  - controlId: SDD-GOV-002
    taxonomyCategory: workflow-governance
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs:
      - proposal.md Scope
      - design.md Architecture Decisions
      - specs/sdd-execution-persistence-contracts/spec.md Requirement: Mandatory Security Artifacts and Status
    requiredControl: >-
      Legacy security-applicability.md artifacts and state refs must remain readable as archive data only and must not weaken mandatory
      security-design.md or security review obligations for new changes.
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks for legacy/read-only wording in persistence, status, OpenSpec, security, and review contracts.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Preserve securityApplicability refs only as legacy/read-only fields and keep archive-only validator behavior explicit.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify no active dependency, successor, or launch mapping normalizes security-applicability into a runnable phase.
    residualRisk: Archived mentions can still appear in searches; active contracts must clearly distinguish legacy evidence from executable workflow.
    exception: null
  - controlId: SDD-GOV-003
    taxonomyCategory: evidence-hygiene
    mandatoryWhenApplicable: true
    operationalSeverity: conditional
    predicate: SDD evidence artifacts are edited or generated for this change.
    predicateResult: true
    sourceRefs:
      - proposal.md Scope
      - design.md Technical Approach
      - skills/_shared/security-guideline-catalog.md SEC-SECRET-001 and SEC-LOG-001 evidence hygiene expectations
    requiredControl: >-
      SDD artifacts must avoid embedding secrets, credentials, production identifiers, sensitive internal data, or unnecessary operational details.
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan manual/static artifact inspection for secret-like or sensitive-data examples in generated SDD evidence.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify changed artifacts contain only repository paths, workflow terms, and catalog IDs; no secrets or sensitive runtime data are introduced.
    residualRisk: Manual review may miss sensitive snippets if future edits add copied environment values; verification should include targeted secret/sensitive-data inspection.
    exception: null
notApplicableCatalogGuidelines:
  - guidelineId: SEC-AUTH-001
    taxonomyCategory: authentication
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no authentication behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-SESS-001
    taxonomyCategory: sessions
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no session or token behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-DATA-001
    taxonomyCategory: sensitive-data-pan
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no sensitive runtime data or PAN behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-SECRET-001
    taxonomyCategory: secrets
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable to runtime secret handling; retained as evidence-hygiene guidance for generated artifacts.
    evidenceStatus: not-applicable
  - guidelineId: SEC-ACCESS-001
    taxonomyCategory: permissions-access-control
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no runtime authorization or privilege boundary behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-FILE-001
    taxonomyCategory: files
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no runtime file handling behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-DB-001
    taxonomyCategory: database-access
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable; no database access behavior changes.
    evidenceStatus: not-applicable
  - guidelineId: SEC-LOG-001
    taxonomyCategory: sensitive-logging
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    rationale: Not applicable to runtime logging; retained as evidence-hygiene guidance for generated artifacts.
    evidenceStatus: not-applicable
carriedRisks:
  - risk: Stale global opencode copies may still expose the retired executor outside this repository.
    ownerPhase: apply
    evidenceExpectation: Document repo-local scope and stale-copy caveat; do not imply global cleanup is completed.
  - risk: Legacy archive mentions may confuse future searches or status readers.
    ownerPhase: verify
    evidenceExpectation: Verify active contracts label applicability as legacy/read-only data and keep new-change routing through security-design.
  - risk: Archive-only validator naming may be mistaken for active new-change validation.
    ownerPhase: verify
    evidenceExpectation: Verify validator references are archive-only and do not block new-change routing.
validation:
  validator: manual-static-security-design-review
  status: pass
  checkedAt: 2026-07-03T00:00:00Z
  notes: >-
    Proposal, four delta specs, technical design, catalog, and shared security contract were reviewed. The artifact classifies the
    change as security-impacting for workflow governance, marks all runtime catalog categories not-applicable with evidence,
    preserves mandatory downstream evidence expectations for governance controls, and sets nextRecommended to test-design.
nextRecommended: test-design
```

## Classification Summary

This change is **security-impacting** because it changes the governance location of security classification for new SDD changes: `security-design.md` becomes mandatory and authoritative, while `security-applicability.md` becomes legacy/read-only data only.

The change is **not runtime security-impacting**. It does not add or modify authentication, sessions, sensitive data/PAN handling, secrets, runtime authorization, file handling, database access, or runtime logging.

## Control Matrix

| Control ID | Category | Mandatory | Required Control | Evidence Owners | Residual Risk | Exception |
| --- | --- | --- | --- | --- | --- | --- |
| `SDD-GOV-001` | `workflow-governance` | Yes | New changes use `security-design.md` as mandatory classification authority and do not recreate `security-applicability.md`. | `test-design`, `apply`, `verify` | Stale global opencode copies may remain outside repo scope. | None |
| `SDD-GOV-002` | `workflow-governance` | Yes | Legacy applicability artifacts remain readable only as archive data and do not weaken mandatory security design/review. | `test-design`, `apply`, `verify` | Archived mentions may confuse search unless active wording is explicit. | None |
| `SDD-GOV-003` | `evidence-hygiene` | Yes | Generated SDD evidence avoids secrets, credentials, production identifiers, and unnecessary sensitive details. | `test-design`, `verify` | Future manual edits could accidentally paste sensitive snippets. | None |

## Catalog Category Decisions

| Category | Guideline | Decision | Rationale |
| --- | --- | --- | --- |
| `authentication` | `SEC-AUTH-001` | Not applicable | No identity, credential, login, recovery, MFA, or impersonation behavior changes. |
| `sessions` | `SEC-SESS-001` | Not applicable | No cookie, token, session lifecycle, revocation, renewal, or fixation behavior changes. |
| `sensitive-data-pan` | `SEC-DATA-001` | Not applicable | No PAN, PII, confidential data, masking, retention, transmission, or storage behavior changes. |
| `secrets` | `SEC-SECRET-001` | Not applicable to runtime | No runtime secret handling changes; evidence hygiene still requires avoiding secret leakage in artifacts. |
| `permissions-access-control` | `SEC-ACCESS-001` | Not applicable | No role, ownership, authorization, privilege, or protected-resource behavior changes. |
| `files` | `SEC-FILE-001` | Not applicable | No runtime upload, download, generated export, path, or file validation behavior changes. |
| `database-access` | `SEC-DB-001` | Not applicable | No query, migration, persistence, tenant isolation, or reporting behavior changes. |
| `sensitive-logging` | `SEC-LOG-001` | Not applicable to runtime | No runtime log/error/trace behavior changes; evidence hygiene still requires avoiding sensitive artifact content. |

## Evidence Expectations

| Control ID | Evidence Type | Owner Phase | Status | Detail |
| --- | --- | --- | --- | --- |
| `SDD-GOV-001` | `test-design-check` | `test-design` | `planned` | Plan static checks proving active routing reaches security-design and never launches `sdd-security-applicability`. |
| `SDD-GOV-001` | `implementation-reference` | `apply` | `planned` | Cite changed agents, shared contracts, and specs that remove active applicability executor references. |
| `SDD-GOV-001` | `verification-evidence` | `verify` | `planned` | Record text-search and diff evidence for deleted launch sources and mandatory security-design routing. |
| `SDD-GOV-002` | `test-design-check` | `test-design` | `planned` | Plan checks that legacy applicability refs are read-only/archive data only. |
| `SDD-GOV-002` | `implementation-reference` | `apply` | `planned` | Cite preservation of legacy refs without active dependency or successor mapping. |
| `SDD-GOV-002` | `verification-evidence` | `verify` | `planned` | Verify no routing/status contract normalizes applicability into a runnable phase. |
| `SDD-GOV-003` | `test-design-check` | `test-design` | `planned` | Plan artifact inspection for secret-like or sensitive-data examples. |
| `SDD-GOV-003` | `verification-evidence` | `verify` | `planned` | Verify changed artifacts contain only repository paths, workflow terms, and catalog IDs. |

## Carried Applicability Risks

- Stale global opencode copies may still expose the retired executor outside repo scope. Owner: `apply`. Expected evidence: document repo-local scope and stale-copy caveat.
- Legacy archive mentions may confuse future searches or status readers. Owner: `verify`. Expected evidence: active contracts label applicability as legacy/read-only data and keep new-change routing through security-design.
- Archive-only validator naming may be mistaken for active new-change validation. Owner: `verify`. Expected evidence: validator references are archive-only and do not block new-change routing.

## Archive Gate Notes

- Mandatory governance controls block archive until evidence is verified or a complete approved exception is recorded.
- Runtime catalog categories are not applicable for this change, but their N/A rationale must remain archive-readable.
- `security-applicability.md` must not be created for this new change.
- `nextRecommended` is `test-design`.
