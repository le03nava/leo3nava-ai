# Design: Slim Secure Design Artifacts

## Technical Approach

Keep secure-design artifacts reviewable by moving exhaustive corporate Source ID inventory ownership to the shared catalog and exhaustive validation ownership to `review-security-report.md`. `design.md#secure-development-design` remains the active security contract, but it records compact `SEC-*` decisions, catalog identity, `expectedSourceIdCount: 155`, grouped coverage references, exact-once downstream validation rules, evidence owners, lifecycle state, and N/A/exception policy instead of materializing all 155 rows.

This change updates SDD instruction contracts and OpenSpec requirements only. The repository has no runtime application code, package manifest, or executable test runner; verification will rely on static/manual artifact checks.

## Architecture Decisions

### Decision: Catalog owns exhaustive inventory

**Choice**: Treat `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` as the authoritative 155 Source ID inventory.
**Alternatives considered**: Duplicating the full source matrix in every design or test-design artifact.
**Rationale**: Centralizing the inventory prevents drift and keeps design/test-design readable while preserving audit traceability through snapshot metadata and expected count.

### Decision: Design owns slim security contract by reference

**Choice**: Preserve compact `SEC-*` controls in design and add grouped source-row coverage metadata below them.
**Alternatives considered**: Making Source IDs replacement controls, or keeping standalone `security-design.md` / `security-applicability.md` in the active DAG.
**Rationale**: Compact controls remain the architectural security decision layer; source rows are operational evidence. New-change routing stays `design -> test-design`.

### Decision: Review-security owns exhaustive expansion

**Choice**: Require only `sdd-review-security` to expand and validate every expected Source ID exactly once in `review-security-report.md`.
**Alternatives considered**: Expanding rows in design, test-design, verify, and archive.
**Rationale**: One phase owning the full matrix reduces artifact bloat and prevents contradictory row verdicts across phases.

## Data Flow

```text
Catalog inventory (155 Source IDs)
        │
        ├─ referenced by design.md#secure-development-design
        │      compact SEC rows + grouped coverage refs + policies
        │
        ├─ consumed by test-design.md
        │      static/manual/automated evidence plan by group/mapping
        │
        └─ expanded by review-security-report.md
               exact-once Source ID matrix + verdicts + evidence
                       │
                       └─ verify/archive preserve summaries and links
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/_shared/sdd-security-contract.md` | Modify | Clarify artifact boundary: catalog inventory, slim design references, test-design planning, review-security expansion, verify/archive preservation. |
| `skills/_shared/security-guideline-catalog.md` | Modify | State authoritative ownership of the 155 Source ID inventory, compact mappings, grouped references, and exact-count contract. |
| `skills/sdd-design/SKILL.md` | Modify | Require slim source-row coverage metadata and prohibit exhaustive matrix duplication in design. |
| `skills/sdd-test-design/SKILL.md` | Modify | Consume slim design coverage and plan evidence without requiring design to carry all rows. |
| `skills/sdd-review-security/SKILL.md` | Modify | Make review-security the exclusive active phase that materializes the exhaustive 155-row matrix. |
| `skills/sdd-archive/SKILL.md` | Modify | Preserve source-row summaries, mappings, warnings, exceptions, and links without duplicating the full matrix. |
| `agents/sdd/sdd-design.md` | Modify | Align adapter prompt with slim design source-row contract. |
| `agents/sdd/sdd-test-design.md` | Modify | Align adapter prompt with grouped coverage planning. |
| `agents/sdd/sdd-review-security.md` | Modify | Align adapter prompt with exclusive full expansion ownership. |
| `agents/sdd/sdd-archive.md` | Modify | Align adapter prompt with archive preservation boundary. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modify | Sync accepted slim secure-design requirements at archive. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Modify | Sync accepted grouped source-row test planning requirements at archive. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modify | Sync accepted exhaustive review-security expansion requirements at archive. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modify | Sync accepted catalog inventory ownership requirements at archive. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modify | Sync accepted verify/archive source-row preservation requirements at archive. |

## Interfaces / Contracts

The active source-row boundary is:

```text
Catalog: authoritative inventory and compact mappings.
Design: compact SEC decisions + expected count + grouped coverage refs + policies.
Test-design: planned checks from slim coverage obligations.
Review-security: exhaustive exact-once Source ID expansion and validation.
Verify/archive: consume and preserve verdicts, summaries, warnings, exceptions, and links.
```

No new runtime API, CLI, or data contract is introduced.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static/manual contract review | Design and skill/spec wording preserves slim artifact boundary, catalog ownership, expected count 155, exact-once rule, evidence owners, and no legacy standalone artifact dependency. | Manual diff inspection and artifact read-back. |
| Runtime unit/integration/e2e | Not applicable; repository has no detected test runner, package manifest, build command, or runtime application. | Record unavailable tooling explicitly in test-design and verify evidence. |
| Security review | Full Source ID expansion remains owned by `review-security-report.md`, not design. | `sdd-review-security` validates against catalog, design, test-design, apply evidence, changed files, and `review-report.md`. |

## Migration / Rollout

No data migration required. Rollout is documentation/contract-only and backward compatible with archived changes: legacy standalone `security-design.md` and `security-applicability.md` remain readable compatibility artifacts but are not active dependencies for new changes.

## Open Questions

None.

## Secure Development Design

```yaml
schemaName: sdd.embedded-secure-design
schemaVersion: 1
changeName: slim-secure-design-artifacts
classification: security-impacting
securityImpact: true
securityImpactRationale: "This change modifies secure development workflow contracts, security evidence boundaries, and corporate Source ID validation ownership. It does not change runtime application behavior, but it directly affects how mandatory security evidence is planned, reviewed, verified, and archived."
sourceInputs:
  proposal: openspec/changes/slim-secure-design-artifacts/proposal.md
  specs:
    - openspec/changes/slim-secure-design-artifacts/specs/sdd-design-workflow/spec.md
    - openspec/changes/slim-secure-design-artifacts/specs/sdd-test-design-workflow/spec.md
    - openspec/changes/slim-secure-design-artifacts/specs/sdd-review-security-workflow/spec.md
    - openspec/changes/slim-secure-design-artifacts/specs/sdd-security-guideline-catalog/spec.md
    - openspec/changes/slim-secure-design-artifacts/specs/sdd-execution-persistence-contracts/spec.md
  design: openspec/changes/slim-secure-design-artifacts/design.md#secure-development-design
catalog:
  snapshotId: security-guidelines-initial-user-snapshot-2026-06-30
  catalogVersion: 1
  taxonomyVersion: 1
  source: skills/_shared/security-guideline-catalog.md
taxonomyEvaluation:
  - category: authentication
    guidelineId: SEC-AUTH-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change alters validation contracts for authentication-related corporate Source IDs without changing runtime auth behavior."
    evidenceRefs: ["proposal.md#approach", "specs/sdd-design-workflow/spec.md#secure-design-source-id-coverage", "skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory"]
  - category: sessions
    guidelineId: SEC-SESS-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change alters how session-related Source ID coverage is referenced, planned, and reviewed."
    evidenceRefs: ["specs/sdd-review-security-workflow/spec.md#exhaustive-source-row-security-review", "skills/_shared/security-guideline-catalog.md#compact-mapping-coverage-by-guideline"]
  - category: sensitive-data-pan
    guidelineId: SEC-DATA-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change defines safe-evidence policy for data/PAN-related security evidence and avoids copying sensitive values into artifacts."
    evidenceRefs: ["skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls", "openspec/config.yaml#context"]
  - category: secrets
    guidelineId: SEC-SECRET-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change preserves secret-safe evidence requirements for design, review-security, verify, and archive artifacts."
    evidenceRefs: ["skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls"]
  - category: permissions-access-control
    guidelineId: SEC-ACCESS-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change affects denial-by-default workflow gating and source-row blocker routing semantics."
    evidenceRefs: ["specs/sdd-security-guideline-catalog/spec.md#shared-security-contract-source-row-schema"]
  - category: files
    guidelineId: SEC-FILE-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change updates file-based SDD artifact contracts and archive/report evidence boundaries."
    evidenceRefs: ["openspec/changes/slim-secure-design-artifacts/proposal.md#affected-areas"]
  - category: database-access
    guidelineId: SEC-DB-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Database-access Source IDs remain in the catalog inventory and must be preserved through grouped design references and review-security expansion."
    evidenceRefs: ["skills/_shared/security-guideline-catalog.md#expanded-source-inventory-and-compact-mappings"]
  - category: sensitive-logging
    guidelineId: SEC-LOG-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The change preserves review-safe logging/error evidence requirements and prevents unsafe details in reports."
    evidenceRefs: ["skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls"]
controls:
  - guidelineId: SEC-AUTH-001
    taxonomyCategory: authentication
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.1-1.10", "2.1-2.23", "4.3", "4.7", "6.3", "6.5", "6.14", "14.8"]
    requiredControl: "Design/test-design reference authentication Source IDs by catalog snapshot and grouped coverage; review-security expands and validates each concrete Source ID exactly once."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check that authentication groups and mappings are planned from catalog references, not copied into design."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Exact-once expanded Source ID validation in review-security-report.md."
    residualRisk: none
    exception: null
  - guidelineId: SEC-SESS-001
    taxonomyCategory: sessions
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["4.7", "6.5", "6.14", "7.1-7.13"]
    requiredControl: "Session Source ID coverage remains catalog-owned and is planned by group; exhaustive verdicts are deferred to review-security."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check for session mapping coverage and unavailable runtime tooling note."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Review-security expands and validates session Source IDs exactly once."
    residualRisk: none
    exception: null
  - guidelineId: SEC-DATA-001
    taxonomyCategory: sensitive-data-pan
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["4.1", "4.3-4.7", "6.5-6.11", "10.1-10.6", "12.1", "12.3-12.5", "13.1-13.9", "15.1-15.2"]
    requiredControl: "Artifacts cite safe paths, sections, and summaries only; design avoids raw PAN, PII, credentials, tokens, connection strings, or confidential values."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "This embedded section documents safe-evidence policy and grouped coverage by reference."
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: "Verify consumes non-blocking review-security evidence and preserves warnings without duplicating the full matrix."
    residualRisk: none
    exception: null
  - guidelineId: SEC-SECRET-001
    taxonomyCategory: secrets
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["2.1", "4.2", "4.5-4.8", "5.5", "6.1", "6.13", "10.6", "13.5"]
    requiredControl: "Evidence references secret/config handling requirements without reproducing secret values; unsafe evidence blocks source-row validation."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Plan static/manual safe-evidence checks for secret-related Source ID groups."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Reject unsafe evidence and route blockers according to security contract."
    residualRisk: none
    exception: null
  - guidelineId: SEC-ACCESS-001
    taxonomyCategory: permissions-access-control
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.4", "5.1", "6.2-6.5", "6.12-6.14", "13.1", "14.1-14.9"]
    requiredControl: "Workflow routing remains denial-by-default for missing, duplicate, malformed, unmapped, unsafe, or unsupported N/A source-row evidence."
    expectedEvidence:
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Skill/prompt/spec changes preserve blocker routing and source-row ownership boundaries."
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: "Archive preserves security review verdicts and refuses unresolved source-row blockers."
    residualRisk: none
    exception: null
  - guidelineId: SEC-FILE-001
    taxonomyCategory: files
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["6.5", "6.14", "9.1-9.12", "14.7"]
    requiredControl: "File-based OpenSpec and skill artifacts carry references and summaries only; review-security owns exhaustive row materialization."
    expectedEvidence:
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Modified Markdown contracts align design, test-design, review-security, verify, and archive boundaries."
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: "Read-back confirms persisted artifacts contain required sections and no legacy standalone dependency."
    residualRisk: none
    exception: null
  - guidelineId: SEC-DB-001
    taxonomyCategory: database-access
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["5.1-5.12", "6.5-6.10", "11.1-11.16", "12.2"]
    requiredControl: "Database Source IDs remain covered by catalog mappings and are validated through review-security expansion even though this change does not alter runtime database code."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check that database coverage stays mapped and traceable by catalog reference."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Review-security produces row-level database Source ID verdicts when validating the implemented contract changes."
    residualRisk: none
    exception: null
  - guidelineId: SEC-LOG-001
    taxonomyCategory: sensitive-logging
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["3.1-3.11", "6.5", "6.14", "8.1-8.5"]
    requiredControl: "Generated evidence remains review-safe and avoids raw secrets, PAN, credentials, tokens, or unnecessary sensitive operational context."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "Safe-evidence policy is stated in this design and inherited by downstream phases."
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: "Archive preserves warnings, exceptions, and evidence links without copying full matrices."
    residualRisk: none
    exception: null
sourceRowCoverage:
  schema: corporate-source-row-operational-layer
  inventoryAuthority: skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory
  expectedSourceIdCount: 155
  expectedSourceUniverse: "All concrete Source IDs in the catalog operational inventory; design uses grouped references and does not duplicate the full expanded matrix."
  coverageRule: "Review-security MUST expand catalog ranges, validate every expected Source ID exactly once, reject missing/duplicate/unknown Source IDs, and write the exhaustive matrix only to review-security-report.md."
  validCompactGuidelineIds: [SEC-AUTH-001, SEC-SESS-001, SEC-DATA-001, SEC-SECRET-001, SEC-ACCESS-001, SEC-FILE-001, SEC-DB-001, SEC-LOG-001]
  compactMappingRule: "Every Source ID MUST map to one or more valid compact SEC IDs from the catalog; missing, empty, or unknown mappings route to resolve-blockers."
  lifecycleStatus: planned
  evidenceOwners: [design, test-design, apply, review-security, verify, archive]
  downstreamTraceability: "Source ID -> compact SEC-* -> design control -> test-design check -> apply evidence -> review-security row verdict -> verify prerequisite -> archive audit trail."
  safeEvidencePolicy: "Use artifact paths, section anchors, changed-file refs, command summaries, sanitized observations, or redacted placeholders only; never include raw secrets, credentials, PAN, PII, tokens, private keys, connection strings, or confidential values."
  notApplicablePolicy: "N/A requires evidence plus justification proving irrelevance by category, platform, API, data class, or workflow; unsupported N/A blocks design/test-design/review-security readiness."
  exceptionPolicy: "Missing mandatory evidence can proceed only with a complete approved exception containing approver, approvedAt, accepted-risk rationale, mitigation/follow-up, and evidence gap."
  fullExpansionOwner: review-security
  groups:
    - corporateSection: "1. Authentication"
      pciAlignment: "PCI Req 6.5.8, 6.5.10"
      sourceIdRef: "1.1-1.10"
      count: 10
      mappedCompactGuidelineIds: [SEC-AUTH-001, SEC-ACCESS-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "2. Passwords"
      pciAlignment: "PCI Req 6.3.1"
      sourceIdRef: "2.1-2.23"
      count: 23
      mappedCompactGuidelineIds: [SEC-AUTH-001, SEC-SECRET-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "3. Access and Activity Logging"
      pciAlignment: "N/A"
      sourceIdRef: "3.1-3.11"
      count: 11
      mappedCompactGuidelineIds: [SEC-LOG-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "4. Cryptography"
      pciAlignment: "PCI Req 6 - 6.5.3"
      sourceIdRef: "4.1-4.8"
      count: 8
      mappedCompactGuidelineIds: [SEC-DATA-001, SEC-SECRET-001, SEC-AUTH-001, SEC-SESS-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "5. Databases"
      pciAlignment: "N/A"
      sourceIdRef: "5.1-5.12"
      count: 12
      mappedCompactGuidelineIds: [SEC-DB-001, SEC-ACCESS-001, SEC-SECRET-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "6. Coding"
      pciAlignment: "PCI Req 6.5.8, 6.5.9"
      sourceIdRef: "6.1-6.14"
      count: 14
      mappedCompactGuidelineIds: [SEC-SECRET-001, SEC-ACCESS-001, SEC-AUTH-001, SEC-DATA-001, SEC-DB-001, SEC-SESS-001, SEC-FILE-001, SEC-LOG-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "7. Session Management"
      pciAlignment: "PCI Req 6.5.9, 6.5.10"
      sourceIdRef: "7.1-7.13"
      count: 13
      mappedCompactGuidelineIds: [SEC-SESS-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "8. Error Handling"
      pciAlignment: "PCI Req 6.3.c, 6.5.5"
      sourceIdRef: "8.1-8.5"
      count: 5
      mappedCompactGuidelineIds: [SEC-LOG-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "9. File Handling"
      pciAlignment: "PCI Req 6.5.8"
      sourceIdRef: "9.1-9.12"
      count: 12
      mappedCompactGuidelineIds: [SEC-FILE-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "10. Memory Management"
      pciAlignment: "PCI Req 6.5.1"
      sourceIdRef: "10.1-10.6"
      count: 6
      mappedCompactGuidelineIds: [SEC-DATA-001, SEC-SECRET-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "11. Input Validation"
      pciAlignment: "PCI Req 6.5.1, 6.5.7, 6.5.8, 6.5.9"
      sourceIdRef: "11.1-11.16"
      count: 16
      mappedCompactGuidelineIds: [SEC-DB-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "12. Output Encoding"
      pciAlignment: "PCI Req 6.5.9"
      sourceIdRef: "12.1-12.5"
      count: 5
      mappedCompactGuidelineIds: [SEC-DATA-001, SEC-DB-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "13. Data Protection"
      pciAlignment: "PCI Req 6.3.c, 6.5.4"
      sourceIdRef: "13.1-13.9"
      count: 9
      mappedCompactGuidelineIds: [SEC-DATA-001, SEC-ACCESS-001, SEC-SECRET-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "14. Access Control"
      pciAlignment: "N/A"
      sourceIdRef: "14.1-14.9"
      count: 9
      mappedCompactGuidelineIds: [SEC-ACCESS-001, SEC-FILE-001, SEC-AUTH-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
    - corporateSection: "15. PAN - Primary Account Number"
      pciAlignment: "PCI Req 6.4"
      sourceIdRef: "15.1-15.2"
      count: 2
      mappedCompactGuidelineIds: [SEC-DATA-001]
      lifecycleStatus: planned
      evidenceOwners: [test-design, apply, review-security, verify, archive]
notApplicableGuidelines: []
exceptions: null
carriedRisks: []
validation:
  method: design.md#secure-development-design static/manual review
  status: pass
  checkedAt: "2026-07-06T00:00:00Z"
  notes: "Design includes all eight compact SEC controls, cites catalog snapshot/version, declares expectedSourceIdCount 155, preserves grouped source-row coverage by reference, assigns evidence owners, states exact-once review-security expansion, and avoids full 155-row or 96-control matrix duplication. Runtime test tooling is unavailable in openspec/config.yaml and is not treated as passing evidence."
archiveGateNotes:
  - "Archive must preserve catalog snapshot, expected count, compact mappings, review-security verdict link, warnings, exceptions, and safe evidence references without requiring standalone security-design.md or security-applicability.md."
nextRecommended: test-design
```

Compact matrix summary:

| Guideline | Applies / lifecycle | Rationale | Secure design decision / control | Evidence owner / expected evidence | Residual risk / exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `Yes` / `planned` | Auth Source IDs are affected through validation-contract changes. | Catalog-owned grouped references; review-security expands exact-once. | test-design static/manual plan; review-security row verdicts. | none / none |
| `SEC-SESS-001` | `Yes` / `planned` | Session Source IDs are affected through coverage and review boundary changes. | Session rows stay catalog-owned; exhaustive verdicts live in review-security. | test-design plan; review-security exact-once validation. | none / none |
| `SEC-DATA-001` | `Yes` / `planned` | Safe-evidence policy affects sensitive-data/PAN evidence. | Use paths/sections/summaries only; no sensitive raw values in artifacts. | design policy; verify consumes review-security evidence. | none / none |
| `SEC-SECRET-001` | `Yes` / `planned` | Secret-safe evidence rules are preserved and clarified. | Never reproduce secret values; unsafe evidence blocks. | test-design safe-evidence checks; review-security rejection rules. | none / none |
| `SEC-ACCESS-001` | `Yes` / `planned` | Routing gates and blocker semantics affect access-control evidence. | Denial-by-default routing for missing/malformed/unsafe source-row evidence. | apply contract changes; archive blocker checks. | none / none |
| `SEC-FILE-001` | `Yes` / `planned` | File-based artifacts and reports are directly modified. | Store slim references in design; full matrix only in review-security. | apply modified Markdown; verify read-back. | none / none |
| `SEC-DB-001` | `Yes` / `planned` | DB Source IDs remain in the required catalog universe. | Preserve DB mappings by catalog reference and review-security expansion. | test-design static/manual check; review-security row verdicts. | none / none |
| `SEC-LOG-001` | `Yes` / `planned` | Logging/error evidence must remain review-safe. | Evidence links and observations must avoid sensitive operational details. | design policy; archive preserves safe links. | none / none |

Source-row coverage summary:

| Boundary | Design Decision |
| --- | --- |
| Inventory authority | `skills/_shared/security-guideline-catalog.md#corporate-source-row-operational-inventory` owns the 155 concrete Source IDs and compact mappings. |
| Expected universe | `expectedSourceIdCount: 155`; grouped references above cover sections 1-15 with counts totaling 155. |
| Exact-once rule | `sdd-review-security` expands catalog ranges and validates each expected Source ID exactly once in `review-security-report.md`. |
| Valid mappings | Only the eight catalog compact IDs are valid mappings. Unknown, missing, or empty mappings block. |
| Evidence owners | Design declares contract; test-design plans checks; apply implements contract edits; review-security expands/validates; verify consumes non-blocking verdicts; archive preserves summaries and links. |
| N/A / exception policy | `N/A` requires positive irrelevance evidence and justification; mandatory evidence gaps require complete approved exceptions. |
| Safe-evidence policy | Cite paths, sections, command summaries, sanitized observations, or redacted placeholders only. |
