# Security Design: Harden Security Applicability Contract

```yaml
schemaName: sdd.security-design
schemaVersion: 1
changeName: harden-security-applicability-contract
sourceApplicability: openspec/changes/harden-security-applicability-contract/security-applicability.md
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  taxonomyVersion: 1
  source: skills/_shared/security-guideline-catalog.md
controls:
  - guidelineId: SEC-AUTH-001
    taxonomyCategory: authentication
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.1-1.10", "2.1-2.23", "6.3", "14.8"]
    requiredControl: Authentication applicability decisions must preserve explicit category evaluation, formal Source ID coverage, blocking severity semantics, and downstream security-design consumption for authentication-governed changes.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve catalog snapshot identity, SEC-AUTH-001, authentication category, mandatory flag, operational severity, and Source refs in this security design.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan static/manual checks that validator rules reject unknown authentication category decisions, invalid severity values, and unresolved blocking authentication obligations.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Update the applicability contract, catalog, and validator so authentication mappings and Source refs are represented without weakening no-impact routing.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify security-applicability artifacts can cite SEC-AUTH-001 and authentication Source refs and block when required evidence is missing.
    residualRisk: Static validation proves structure and references, not runtime authentication behavior in future application changes.
    exception: null
  - guidelineId: SEC-SESS-001
    taxonomyCategory: sessions
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["7.1-7.13"]
    requiredControl: Session-related applicability obligations must remain represented as blocking downstream controls whenever session or token changes are classified applicable.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-SESS-001, sessions category, mandatory flag, operational severity, and Source refs in the control model.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks proving enriched applicability fields do not drop session obligations during security-design, tasks, verify, or archive handoff.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Update downstream phase contracts to consume enriched applicability metadata only when securityImpact is true.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify session guideline refs survive applicability-to-security-design translation and remain archive-visible.
    residualRisk: Future session implementation details still require change-specific security design; this control only hardens governance routing.
    exception: null
  - guidelineId: SEC-DATA-001
    taxonomyCategory: sensitive-data-pan
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["4.1", "13.1-13.9", "15.1-15.2"]
    requiredControl: Sensitive-data/PAN applicability must require positive no-impact proof or explicit downstream controls; missing evidence must not be treated as no-impact.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-DATA-001, sensitive-data-pan category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks for complete category matrix rows, explicit no-impact proof, valid Source refs, and routing consistency for data/PAN decisions.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Add schema/template and validator support that distinguishes absent data evidence from affirmative no-impact evidence.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify no-impact examples require every supported category to be not-applicable with rationale and evidence refs.
    residualRisk: Corporate Source ID mappings are being formalized in this change; incorrect mappings could create false audit confidence if not reviewed.
    exception: null
  - guidelineId: SEC-SECRET-001
    taxonomyCategory: secrets
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["2.1", "4.2", "4.8", "5.5", "6.1", "13.5"]
    requiredControl: Secret-handling applicability must preserve source coverage and prevent config overrides from weakening required categories, source checks, or blocking obligations.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-SECRET-001, secrets category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks that safe rules.security-applicability overrides may add prompts or stricter validation but cannot weaken source coverage or blocking severity.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Implement override handling in shared contract/applicability skill/validator so unsafe weakening is rejected or ignored in favor of the base contract.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify unsafe override examples fail or preserve stricter base requirements, with no secret-related category bypass.
    residualRisk: The validator cannot prove real secret storage/rotation behavior for later application changes; it only preserves governance obligations.
    exception: null
  - guidelineId: SEC-ACCESS-001
    taxonomyCategory: permissions-access-control
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.4", "6.2-6.4", "6.12", "13.1", "14.1-14.9"]
    requiredControl: Access-control applicability must keep authorization and least-privilege obligations blocking when applicable, including no-impact compatibility for unrelated changes with complete proof.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-ACCESS-001, permissions-access-control category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks for valid no-impact routing: complete not-applicable matrix proof skips security-design, while incomplete proof blocks rather than silently skipping.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Update downstream compatibility notes so enriched fields do not make security-design mandatory for valid no-impact changes.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify access-control guideline refs and no-impact compatibility remain understandable to tasks, verify, and archive.
    residualRisk: Additive compatibility depends on downstream phase instructions being updated consistently across skills and adapter prompt copies.
    exception: null
  - guidelineId: SEC-FILE-001
    taxonomyCategory: files
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["9.1-9.12", "14.7"]
    requiredControl: File-handling applicability must remain covered by complete category decision rows, formal Source ID validation, and downstream evidence preservation.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-FILE-001, files category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks that every supported category appears exactly once and unknown file decisions with blocking severity prevent phase success.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Update catalog metadata and validator reference checks so file Source IDs resolve within the recorded snapshot.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify duplicate/missing file category rows and invalid file Source refs fail static validation.
    residualRisk: No runtime file upload/download behavior exists in this repository; evidence will be static contract/validator evidence.
    exception: null
  - guidelineId: SEC-DB-001
    taxonomyCategory: database-access
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["5.1-5.12", "11.1-11.16", "12.2"]
    requiredControl: Database-access applicability must formalize Source ID coverage and keep validator scope limited to static Markdown/YAML structure and reference validation.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-DB-001, database-access category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan static validator checks for schema fields, matrix completeness, guideline/source validity, severity vocabulary, override safety, and routing mismatch only.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Implement scripts/validate_security_applicability.ps1 as a repository-local static validator, not a broad runtime security scanner.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify validator results report structural/reference failures and do not claim runtime database security behavior.
    residualRisk: Validator scope must remain intentionally narrow; expanding into runtime scanning would exceed the design and review budget.
    exception: null
  - guidelineId: SEC-LOG-001
    taxonomyCategory: sensitive-logging
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["3.1-3.11", "8.1-8.5"]
    requiredControl: Logging/redaction applicability must preserve operational severity and downstream evidence expectations so advisory or blocking logging obligations remain visible through archive.
    expectedEvidence:
      - type: design-control
        ownerPhase: security-design
        status: planned
        detail: Preserve SEC-LOG-001, sensitive-logging category, mandatory flag, operational severity, and Source refs.
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: Plan checks that severity vocabulary is limited to blocking, conditional, and advisory and does not reuse review labels.
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: Update catalog and security contract prose so operational severity is distinct from review severity labels and remains downstream-visible.
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: Verify invalid severity labels fail validation and logging guideline evidence remains archive-readable.
    residualRisk: Static governance controls cannot prove future runtime log redaction; future security-impacting changes still need change-specific evidence.
    exception: null
carriedRisks:
  - risk: Exact validator implementation path was deferred by applicability and resolved in technical design as scripts/validate_security_applicability.ps1.
    ownerPhase: apply
    evidenceExpectation: Implement and reference the PowerShell validator; verify it is static and scoped to Markdown/YAML structure and references.
    status: planned
  - risk: Compact SEC-* rows do not yet contain formal source coverage columns; this change adds those mappings from the full corporate snapshot.
    ownerPhase: apply
    evidenceExpectation: Add formal Source ID coverage per compact SEC-* guideline and verify artifact refs resolve within snapshot security-guidelines-initial-user-snapshot-2026-06-30.
    status: planned
  - risk: Enriched applicability fields could accidentally force security-design for valid no-impact changes.
    ownerPhase: verify
    evidenceExpectation: Verify a valid no-impact artifact with explicit proof skips security-design while invalid no-impact proof blocks routing.
    status: planned
nextRecommended: test-design
```

## Control Matrix

| Guideline ID | Category | Mandatory | Required Control | Evidence Owners | Residual Risk | Exception |
| --- | --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `authentication` | Yes | Preserve explicit authentication category evaluation, formal Source ID coverage, blocking severity, and downstream consumption. | `security-design`, `test-design`, `apply`, `verify` | Static validation does not prove future runtime authentication behavior. | None |
| `SEC-SESS-001` | `sessions` | Yes | Keep session/token obligations represented as blocking downstream controls whenever applicable. | `security-design`, `test-design`, `apply`, `verify` | Future session implementation details still need change-specific controls. | None |
| `SEC-DATA-001` | `sensitive-data-pan` | Yes | Require positive no-impact proof or explicit downstream controls; never infer no-impact from absent evidence. | `security-design`, `test-design`, `apply`, `verify` | Incorrect Source ID mappings could create false audit confidence. | None |
| `SEC-SECRET-001` | `secrets` | Yes | Preserve source coverage and reject/ignore config overrides that weaken required categories, source checks, or blocking obligations. | `security-design`, `test-design`, `apply`, `verify` | Validator cannot prove real secret storage/rotation behavior for later changes. | None |
| `SEC-ACCESS-001` | `permissions-access-control` | Yes | Keep authorization obligations blocking when applicable and preserve no-impact compatibility when proof is complete. | `security-design`, `test-design`, `apply`, `verify` | Downstream phase instructions must be updated consistently. | None |
| `SEC-FILE-001` | `files` | Yes | Require complete category rows, formal Source ID validation, and downstream evidence preservation for file-handling decisions. | `security-design`, `test-design`, `apply`, `verify` | Evidence is static because the repository has no runtime file feature. | None |
| `SEC-DB-001` | `database-access` | Yes | Formalize database Source ID coverage while keeping validator scope limited to static structure and references. | `security-design`, `test-design`, `apply`, `verify` | Validator scope expansion would exceed this design and review budget. | None |
| `SEC-LOG-001` | `sensitive-logging` | Yes | Preserve logging/redaction severity and evidence expectations through archive. | `security-design`, `test-design`, `apply`, `verify` | Static governance controls cannot prove future runtime log redaction. | None |

## Evidence Expectations

| Guideline ID | Evidence Type | Owner Phase | Status | Detail |
| --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `test-design-check` | `test-design` | `planned` | Check unknown authentication decisions, invalid severity values, and unresolved blocking obligations fail. |
| `SEC-AUTH-001` | `implementation-reference` | `apply` | `planned` | Update applicability contract, catalog, and validator to preserve auth category and Source refs. |
| `SEC-AUTH-001` | `verification-evidence` | `verify` | `planned` | Verify SEC-AUTH-001 refs are valid and blocking evidence gaps block. |
| `SEC-SESS-001` | `test-design-check` | `test-design` | `planned` | Check session obligations survive security-design and downstream handoff. |
| `SEC-SESS-001` | `implementation-reference` | `apply` | `planned` | Update downstream contracts to consume enriched metadata only when securityImpact is true. |
| `SEC-SESS-001` | `verification-evidence` | `verify` | `planned` | Verify session guideline evidence remains archive-visible. |
| `SEC-DATA-001` | `test-design-check` | `test-design` | `planned` | Check matrix completeness, no-impact proof, Source refs, and routing consistency. |
| `SEC-DATA-001` | `implementation-reference` | `apply` | `planned` | Add schema/template and validator support for positive no-impact proof. |
| `SEC-DATA-001` | `verification-evidence` | `verify` | `planned` | Verify no-impact examples require every supported category to be not-applicable with rationale and evidence. |
| `SEC-SECRET-001` | `test-design-check` | `test-design` | `planned` | Check supported overrides cannot weaken source coverage or blocking severity. |
| `SEC-SECRET-001` | `implementation-reference` | `apply` | `planned` | Implement safe override handling in contract/applicability/validator artifacts. |
| `SEC-SECRET-001` | `verification-evidence` | `verify` | `planned` | Verify unsafe overrides fail or preserve stricter base requirements. |
| `SEC-ACCESS-001` | `test-design-check` | `test-design` | `planned` | Check valid no-impact proof skips security-design and incomplete proof blocks. |
| `SEC-ACCESS-001` | `implementation-reference` | `apply` | `planned` | Add downstream compatibility notes for enriched applicability fields. |
| `SEC-ACCESS-001` | `verification-evidence` | `verify` | `planned` | Verify no-impact compatibility remains understandable to tasks, verify, and archive. |
| `SEC-FILE-001` | `test-design-check` | `test-design` | `planned` | Check every category appears exactly once and blocking unknowns prevent success. |
| `SEC-FILE-001` | `implementation-reference` | `apply` | `planned` | Add catalog metadata and validator checks so file Source IDs resolve. |
| `SEC-FILE-001` | `verification-evidence` | `verify` | `planned` | Verify duplicate/missing file rows and invalid Source refs fail. |
| `SEC-DB-001` | `test-design-check` | `test-design` | `planned` | Check validator scope covers schema, matrix, refs, severity, overrides, and routing only. |
| `SEC-DB-001` | `implementation-reference` | `apply` | `planned` | Create `scripts/validate_security_applicability.ps1` as a static validator. |
| `SEC-DB-001` | `verification-evidence` | `verify` | `planned` | Verify validator failures do not claim runtime database security behavior. |
| `SEC-LOG-001` | `test-design-check` | `test-design` | `planned` | Check severity vocabulary allows only `blocking`, `conditional`, and `advisory`. |
| `SEC-LOG-001` | `implementation-reference` | `apply` | `planned` | Separate operational severity from review labels in catalog and contract prose. |
| `SEC-LOG-001` | `verification-evidence` | `verify` | `planned` | Verify invalid severity labels fail and logging evidence remains archive-readable. |

## Carried Applicability Risks

| Risk | Owner Phase | Evidence Expectation | Status |
| --- | --- | --- | --- |
| Exact validator implementation path was deferred by applicability and resolved in design as `scripts/validate_security_applicability.ps1`. | `apply` | Implement and reference the PowerShell validator; verify it remains static and scoped to Markdown/YAML structure and references. | `planned` |
| Compact `SEC-*` rows do not yet contain formal source coverage columns; this change adds those mappings from the full corporate snapshot. | `apply` | Add formal Source ID coverage per compact guideline and verify refs resolve within snapshot `security-guidelines-initial-user-snapshot-2026-06-30`. | `planned` |
| Enriched applicability fields could accidentally force security-design for valid no-impact changes. | `verify` | Verify valid no-impact proof skips security-design while invalid no-impact proof blocks routing. | `planned` |

## Exception Handling

No approved exceptions are recorded for this security design. Any future exception that attempts to satisfy missing mandatory evidence must include all shared contract fields: `status: exception-approved`, `guidelineId`, `approver`, `approvedAt`, `acceptedRiskRationale`, `mitigationOrFollowUp`, and `evidenceGap`.

## Archive Gate Notes

- Mandatory applicable controls block archive until evidence is verified or a complete approved exception is recorded.
- Incomplete exceptions do not satisfy archive readiness.
- Archive evidence must preserve catalog snapshot identity, guideline IDs, taxonomy categories, source references, operational severity, expected evidence status, and residual risks.
- Valid no-impact artifacts remain compatible with downstream routing: they skip `security-design.md` only when explicit no-impact proof and validation metadata are complete.
