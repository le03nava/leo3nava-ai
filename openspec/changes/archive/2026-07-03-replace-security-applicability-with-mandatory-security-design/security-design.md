# Security Design: Mandatory Security Design for SDD Changes

```yaml
schemaName: sdd.security-design
schemaVersion: 1
changeName: replace-security-applicability-with-mandatory-security-design
classification: security-impacting
securityImpact: true
securityImpactRationale: >-
  This change does not add an application runtime, user authentication, sessions,
  PAN processing, file uploads, database access, or runtime logging. It is still
  security-impacting because it redesigns SDD security governance, mandatory
  evidence gates, security matrix ownership, and archive readiness behavior.
sourceInputs:
  proposal: openspec/changes/replace-security-applicability-with-mandatory-security-design/proposal.md
  design: openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md
  specs:
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-review-security-workflow/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-security-applicability-workflow/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-security-design-workflow/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-test-design-workflow/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-review-workflow/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-execution-persistence-contracts/spec.md
    - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-security-guideline-catalog/spec.md
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  catalogVersion: 1
  taxonomyVersion: 1
  source: skills/_shared/security-guideline-catalog.md
  compactGuidelineCount: 8
taxonomyEvaluation:
  - category: authentication
    guidelineId: SEC-AUTH-001
    applies: N/A
    decision: not-applicable
    lifecycleStatus: not-applicable
    rationale: Change does not modify login, identity proofing, MFA, credential validation, impersonation, or account recovery behavior.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#technical-approach
  - category: sessions
    guidelineId: SEC-SESS-001
    applies: N/A
    decision: not-applicable
    lifecycleStatus: not-applicable
    rationale: Change does not modify cookies, tokens, refresh flows, session lifetimes, revocation, or fixation protections.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#technical-approach
  - category: sensitive-data-pan
    guidelineId: SEC-DATA-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: Security evidence artifacts and review reports must not copy PAN, PII, secrets, or unnecessary sensitive values into repo-persisted SDD evidence.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-review-security-workflow/spec.md
      - skills/_shared/security-guideline-catalog.md
  - category: secrets
    guidelineId: SEC-SECRET-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: The redesigned evidence gates must require references to secret/config locations without committing, logging, or reproducing secret values in artifacts.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#interfaces--contracts
      - skills/_shared/security-guideline-catalog.md
  - category: permissions-access-control
    guidelineId: SEC-ACCESS-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: The change modifies workflow authorization-like governance gates by making security design and security review mandatory before verify/archive.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/proposal.md#scope
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#technical-approach
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-execution-persistence-contracts/spec.md
  - category: files
    guidelineId: SEC-FILE-001
    applies: N/A
    decision: not-applicable
    lifecycleStatus: not-applicable
    rationale: Change edits repository Markdown/contracts/scripts only; it does not introduce uploads, downloads, generated exports, path handling, or user-provided file processing.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#file-changes
  - category: database-access
    guidelineId: SEC-DB-001
    applies: N/A
    decision: not-applicable
    lifecycleStatus: not-applicable
    rationale: Repository has no database runtime and this change does not add queries, migrations, persistence paths, tenant isolation, or data access behavior.
    evidenceRefs:
      - openspec/config.yaml
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/design.md#technical-approach
  - category: sensitive-logging
    guidelineId: SEC-LOG-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: Security review and archive evidence are audit-like artifacts; they must preserve useful evidence without exposing secrets, PAN, credentials, or unnecessary sensitive context.
    evidenceRefs:
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-review-security-workflow/spec.md
      - openspec/changes/replace-security-applicability-with-mandatory-security-design/specs/sdd-security-guideline-catalog/spec.md
controls:
  - guidelineId: SEC-DATA-001
    taxonomyCategory: sensitive-data-pan
    applies: Yes
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: [4.1, 13.1-13.9, 15.1-15.2]
    requiredControl: >-
      Update mandatory security-design, security-review, verification, and archive contracts so evidence records use artifact paths,
      section references, command summaries, or redacted examples instead of copying PAN, PII, credentials, or confidential values.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: This security design records sensitive-evidence minimization as an applicable governance control.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Updated shared security contract, catalog guidance, security-design/review-security skills, and validator text reject raw sensitive values in evidence fields.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Static/manual checks confirm matrix and evidence examples use paths/sections/redacted placeholders instead of real sensitive data.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify report cites static grep/manual inspection for sensitive-data leakage in changed SDD artifacts.
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: Archive gate confirms mandatory evidence is verified or exception-approved before archive.
    residualRisk: Low residual risk that future free-text evidence may include sensitive examples; mitigated by validator guidance and security review.
    exception: null
  - guidelineId: SEC-SECRET-001
    taxonomyCategory: secrets
    applies: Yes
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: [2.1, 4.2, 4.8, 5.5, 6.1, 13.5]
    requiredControl: >-
      Contracts and review-security guidance must require evidence references for secret handling and must prohibit committing,
      echoing, or logging actual secret values in SDD artifacts, examples, validators, review reports, or archive reports.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: This artifact marks secret exposure prevention as applicable to SDD evidence governance.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Shared contract/catalog and phase skills instruct agents to cite secret source/config names or redacted placeholders only.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Planned static/manual check searches changed artifacts for obvious secret-looking literals and verifies examples are placeholders.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify report records no committed secrets or only approved redacted placeholders in changed SDD artifacts.
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: Archive blocks if mandatory secret-safety evidence is missing without approved exception.
    residualRisk: Low residual risk from human-authored evidence snippets; mitigated by explicit artifact rules and review-security validation.
    exception: null
  - guidelineId: SEC-ACCESS-001
    taxonomyCategory: permissions-access-control
    applies: Yes
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: [1.4, 6.2-6.4, 6.12, 13.1, 14.1-14.9]
    requiredControl: >-
      The active SDD DAG, status contract, persistence contract, skills, and validators must enforce denial-by-default phase progression:
      new changes cannot skip mandatory security-design or review-security before verify/archive, and blockers route back to apply or resolve-blockers.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: This artifact records mandatory security gates as the access-control analogue for SDD workflow governance.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Updated orchestrator, phase agents/skills, status/persistence contracts, and validators enforce security-design and review-security dependencies.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Static routing tests/checklists cover spec -> design -> security-design -> test-design and review -> review-security -> verify.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify report proves no direct design-to-test-design or review-to-verify route remains for new changes.
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: Archive requires non-blocking security-design/security-review evidence or complete approved exceptions.
    residualRisk: Medium residual risk of stale adapter prompt copies retaining old routing; mitigated by cross-file grep verification and stacked rollout.
    exception: null
  - guidelineId: SEC-LOG-001
    taxonomyCategory: sensitive-logging
    applies: Yes
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: [3.1-3.11, 8.1-8.5]
    requiredControl: >-
      Security review and archive reports must be useful audit artifacts while avoiding sensitive payloads; observations should cite evidence locations,
      summaries, and redacted values rather than raw secrets, credentials, PAN, or confidential operational context.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: This artifact classifies security evidence/reporting as applicable audit-like logging governance.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Review-security skill/report contract and catalog guidance require evidence location plus observations without sensitive value disclosure.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Planned static/manual checks inspect report templates and examples for redaction and evidence-location requirements.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify report records static/manual inspection that changed report templates do not require sensitive values.
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: Archive gate confirms review-security report is present and non-blocking without sensitive evidence leakage.
    residualRisk: Low residual risk that future reports over-share context; mitigated by review-security validation and archive gate notes.
    exception: null
notApplicableGuidelines:
  - guidelineId: SEC-AUTH-001
    taxonomyCategory: authentication
    applies: N/A
    lifecycleStatus: not-applicable
    rationale: No authentication, identity proofing, credential validation, MFA, impersonation, or account recovery behavior is introduced or modified.
  - guidelineId: SEC-SESS-001
    taxonomyCategory: sessions
    applies: N/A
    lifecycleStatus: not-applicable
    rationale: No cookies, bearer tokens, refresh tokens, server-side sessions, logout flow, or session lifetime behavior is introduced or modified.
  - guidelineId: SEC-FILE-001
    taxonomyCategory: files
    applies: N/A
    lifecycleStatus: not-applicable
    rationale: Repository artifact edits do not create upload/download, generated export, path traversal, user-provided filename, or file authorization behavior.
  - guidelineId: SEC-DB-001
    taxonomyCategory: database-access
    applies: N/A
    lifecycleStatus: not-applicable
    rationale: The repository has no database runtime; this change does not add SQL/ORM queries, migrations, tenant boundaries, or persistence access paths.
exceptions: null
carriedRisks:
  - risk: Adapter prompt copies or status examples may retain old conditional security-applicability wording.
    ownerPhase: apply
    expectedEvidence: Cross-file static grep/read-back checks for removed active routing and mandatory security-design/review-security wording.
  - risk: No executable project test runner exists, so validation will rely on PowerShell/static/manual checks.
    ownerPhase: test-design
    expectedEvidence: Test design must make unavailable runtime testing explicit and define static/manual evidence instead.
validation:
  status: manual-pending
  validator: scripts/validate_security_design.ps1
  checkedAt: 2026-07-03T00:00:00-06:00
  notes: Validator is planned by this change and does not exist yet; this artifact was manually checked against catalog IDs, taxonomy categories, matrix vocabulary, lifecycle statuses, evidence owner phases, and nextRecommended.
archiveGateNotes:
  - New changes must not archive unless security-design.md exists and preserves catalog identity plus all compact SEC-* rows.
  - Applicable mandatory controls block archive until implementation and verification evidence are verified or a complete approved exception is recorded.
  - N/A rows remain archive-readable only when they include evidence and rationale proving why the guideline is outside this workflow redesign scope.
  - review-security-report.md must be present and non-blocking before verify/archive for new changes.
  - Legacy security-applicability.md artifacts remain readable only for archived or old changes and must not become active routing authority for new changes.
nextRecommended: test-design
```

## Security Matrix

| Guideline ID | Category | Applies | Lifecycle Status | Required Control | Evidence Owner | Evidence Location | Observations |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | N/A | `not-applicable` | None; no authentication behavior is changed. | `security-design` | `design.md#technical-approach` | Workflow contracts do not touch login, MFA, credential validation, impersonation, or recovery. |
| `SEC-SESS-001` | `sessions` | N/A | `not-applicable` | None; no session or token behavior is changed. | `security-design` | `design.md#technical-approach` | No cookies, bearer tokens, refresh tokens, logout, revocation, or fixation controls are in scope. |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | `planned` | Evidence artifacts must use paths, sections, summaries, or redacted examples instead of copying PAN, PII, credentials, or confidential values. | `security-design`, `apply`, `test-design`, `verify`, `archive` | This file `controls[SEC-DATA-001]`; planned contract/skill/validator updates | Applies to governance evidence, not to application data processing. |
| `SEC-SECRET-001` | `secrets` | Yes | `planned` | SDD artifacts, examples, validators, and reports must cite secret/config references without committing, echoing, or logging actual secret values. | `security-design`, `apply`, `test-design`, `verify`, `archive` | This file `controls[SEC-SECRET-001]`; planned shared contract/catalog/skill guidance | Applies because evidence gates can accidentally create committed secret exposure. |
| `SEC-ACCESS-001` | `permissions-access-control` | Yes | `planned` | Enforce denial-by-default workflow progression: mandatory security-design and review-security before verify/archive, with blockers routing back to apply or resolve-blockers. | `security-design`, `apply`, `test-design`, `verify`, `archive` | `proposal.md#scope`; `design.md#technical-approach`; `specs/sdd-execution-persistence-contracts/spec.md` | This is the primary security governance control for the workflow redesign. |
| `SEC-FILE-001` | `files` | N/A | `not-applicable` | None; no upload/download/generated-export/user-file path behavior is changed. | `security-design` | `design.md#file-changes` | Repository Markdown/script edits are not the catalog's file-handling threat surface. |
| `SEC-DB-001` | `database-access` | N/A | `not-applicable` | None; no database access behavior is changed. | `security-design` | `openspec/config.yaml`; `design.md#technical-approach` | Repository has no database runtime or query/migration layer. |
| `SEC-LOG-001` | `sensitive-logging` | Yes | `planned` | Security review and archive reports must provide audit evidence without exposing secrets, PAN, credentials, or unnecessary sensitive context. | `security-design`, `apply`, `test-design`, `verify`, `archive` | This file `controls[SEC-LOG-001]`; planned review-security report contract | Applies to audit-like SDD reports/evidence, not runtime application logging. |

## Evidence Expectations

| Guideline ID | Evidence Type | Owner Phase | Status | Detail |
| --- | --- | --- | --- | --- |
| `SEC-DATA-001` | `design-control` | `security-design` | `planned` | Security design records sensitive-evidence minimization as an applicable governance control. |
| `SEC-DATA-001` | `implementation-reference` | `apply` | `planned` | Shared contract, catalog guidance, phase skills, and validator text reject raw sensitive values in evidence fields. |
| `SEC-DATA-001` | `test-design-check` | `test-design` | `planned` | Static/manual checks confirm evidence examples use paths, sections, summaries, or redacted placeholders. |
| `SEC-DATA-001` | `verification-evidence` | `verify` | `planned` | Verify report cites static/manual inspection for sensitive-data leakage in changed SDD artifacts. |
| `SEC-DATA-001` | `verification-evidence` | `archive` | `planned` | Archive confirms mandatory evidence is verified or exception-approved. |
| `SEC-SECRET-001` | `design-control` | `security-design` | `planned` | Secret exposure prevention is classified as applicable to SDD evidence governance. |
| `SEC-SECRET-001` | `implementation-reference` | `apply` | `planned` | Guidance instructs agents to cite secret/config names or redacted placeholders only. |
| `SEC-SECRET-001` | `test-design-check` | `test-design` | `planned` | Static/manual check searches changed artifacts for obvious secret-looking literals and verifies placeholder examples. |
| `SEC-SECRET-001` | `verification-evidence` | `verify` | `planned` | Verify report records no committed secrets or only approved redacted placeholders. |
| `SEC-SECRET-001` | `verification-evidence` | `archive` | `planned` | Archive blocks if mandatory secret-safety evidence is missing without approved exception. |
| `SEC-ACCESS-001` | `design-control` | `security-design` | `planned` | Mandatory security gates are recorded as the access-control analogue for SDD workflow governance. |
| `SEC-ACCESS-001` | `implementation-reference` | `apply` | `planned` | Orchestrator, phase agents/skills, status/persistence contracts, and validators enforce dependencies. |
| `SEC-ACCESS-001` | `test-design-check` | `test-design` | `planned` | Static routing checks cover `spec -> design -> security-design -> test-design` and `review -> review-security -> verify`. |
| `SEC-ACCESS-001` | `verification-evidence` | `verify` | `planned` | Verify proves no direct design-to-test-design or review-to-verify route remains for new changes. |
| `SEC-ACCESS-001` | `verification-evidence` | `archive` | `planned` | Archive requires non-blocking security-design/security-review evidence or complete approved exceptions. |
| `SEC-LOG-001` | `design-control` | `security-design` | `planned` | Security evidence/reporting is classified as applicable audit-like logging governance. |
| `SEC-LOG-001` | `implementation-reference` | `apply` | `planned` | Review-security skill/report contract and catalog guidance require evidence locations plus safe observations. |
| `SEC-LOG-001` | `test-design-check` | `test-design` | `planned` | Static/manual checks inspect report templates and examples for redaction and evidence-location requirements. |
| `SEC-LOG-001` | `verification-evidence` | `verify` | `planned` | Verify records static/manual inspection that changed report templates do not require sensitive values. |
| `SEC-LOG-001` | `verification-evidence` | `archive` | `planned` | Archive confirms review-security report is present, non-blocking, and free of sensitive evidence leakage. |

## Non-Applicable Rationale

| Guideline ID | Category | Rationale | Evidence Location |
| --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | The workflow redesign does not modify any login, identity proofing, credential validation, MFA, impersonation, or account recovery behavior. | `design.md#technical-approach` |
| `SEC-SESS-001` | `sessions` | The workflow redesign does not modify cookies, tokens, refresh flows, session lifetime, revocation, logout, or fixation prevention. | `design.md#technical-approach` |
| `SEC-FILE-001` | `files` | The change edits repository artifacts and planned validators; it does not add upload/download flows, generated exports, user-controlled paths, or file authorization behavior. | `design.md#file-changes` |
| `SEC-DB-001` | `database-access` | The repository has no database runtime and this change does not add SQL/ORM access, migrations, tenant boundaries, reports, or background data jobs. | `openspec/config.yaml`; `design.md#technical-approach` |

## Exceptions

None.

## Carried Applicability Risks

| Risk | Owner Phase | Expected Evidence |
| --- | --- | --- |
| Adapter prompt copies or status examples may retain old conditional security-applicability wording. | `apply` | Cross-file static grep/read-back checks for removed active routing and mandatory security-design/review-security wording. |
| No executable project test runner exists, so validation will rely on PowerShell/static/manual checks. | `test-design` | Test design must make unavailable runtime testing explicit and define static/manual evidence instead. |

## Archive Gate Notes

- New changes must not archive unless `security-design.md` exists and preserves catalog identity plus all compact `SEC-*` rows.
- Applicable mandatory controls (`SEC-DATA-001`, `SEC-SECRET-001`, `SEC-ACCESS-001`, `SEC-LOG-001`) block archive until implementation and verification evidence are verified or a complete approved exception is recorded.
- `N/A` rows remain archive-readable only when they include evidence and rationale proving why the guideline is outside this workflow redesign scope.
- `review-security-report.md` must be present and non-blocking before verify/archive for new changes.
- Legacy `security-applicability.md` artifacts remain readable only for archived or old changes and must not become active routing authority for new changes.
- Artifact-local `nextRecommended` is `test-design`.
