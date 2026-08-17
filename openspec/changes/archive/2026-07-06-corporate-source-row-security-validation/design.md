# Design: Corporate Source Row Security Validation

## Technical Approach

Add a source-row operational layer below the existing eight compact `SEC-*` controls. The compact controls remain the architecture, routing, and evidence summary layer; the source-row layer becomes the exhaustive audit inventory derived from `skills/_shared/security-guideline-catalog.md#full-corporate-guideline-snapshot`.

The implementation will update shared contracts first, then phase skills and agent prompts, so every downstream phase reads the same source-row schema and persistence semantics. Because this repository has no runtime test runner, validation is designed as static/manual artifact validation with explicit unavailable-tooling notes.

## Architecture Decisions

| Decision | Choice | Alternatives considered | Rationale |
| --- | --- | --- | --- |
| Layering | Keep eight compact `SEC-*` controls as the control layer and add source rows as operational evidence below them. | Replace compact controls with per-source controls. | Preserves stable review routing while enabling exhaustive corporate row coverage. |
| Source row storage | Store source-row schema and mapping rules in `skills/_shared/sdd-security-contract.md` and catalog mapping data in `skills/_shared/security-guideline-catalog.md`. | Duplicate mapping rules in each phase skill. | One shared source prevents drift across design, test-design, review-security, verify, and archive. |
| Range expansion | Expand ranges before validation and persist expanded IDs in artifacts/reports. | Keep ranges compressed and trust readers to expand mentally. | Missing/duplicate/unknown row checks require concrete IDs. |
| Mapping for unmapped IDs | Map every corporate Source ID to one or more existing compact `SEC-*` IDs using security concern semantics; do not add new compact controls. | Create new compact IDs for memory, crypto, output encoding, etc. | The proposal explicitly preserves the compact eight-control taxonomy. |
| Evidence policy | Evidence fields cite paths, sections, sanitized summaries, command summaries, or redacted placeholders only. | Copy raw evidence into reports. | Prevents leaking secrets, PII, PAN, tokens, connection strings, private keys, or confidential values. |

## Data Flow

```text
Full Corporate Guideline Snapshot
  -> source-row inventory + PCI alignment extraction
  -> range expansion and duplicate/unknown checks
  -> Source ID -> compact SEC-* mapping
  -> design.md#secure-development-design expected coverage
  -> test-design planned checks
  -> apply implementation evidence
  -> review-security source-row verdicts
  -> verify prerequisite consumption
  -> archive audit preservation
```

## File Changes

| File | Action | Description |
| --- | --- | --- |
| `skills/_shared/security-guideline-catalog.md` | Modify | Add source-row inventory/schema guidance, range expansion rules, source mapping table, and safe-evidence constraints. |
| `skills/_shared/sdd-security-contract.md` | Modify | Extend embedded secure-design and review-security contracts with source-row schema, traceability, routing, and persistence-compatible evidence semantics. |
| `skills/sdd-design/SKILL.md` | Modify | Require expected Source ID coverage below compact controls in `design.md#secure-development-design`. |
| `skills/sdd-test-design/SKILL.md` | Modify | Require static/manual source-row test planning and N/A evidence preservation. |
| `skills/sdd-review-security/SKILL.md` | Modify | Validate every corporate Source ID exactly once without duplicating the general 96-control review matrix. |
| `skills/sdd-verify/SKILL.md` | Modify | Consume non-blocking source-row security review evidence and block unresolved source-row blockers. |
| `skills/sdd-archive/SKILL.md` | Modify | Preserve source-row coverage, mappings, warnings, exceptions, and evidence refs in archive readiness/audit trail. |
| `agents/sdd/sdd-design.md`, `agents/sdd/sdd-test-design.md`, `agents/sdd/sdd-review-security.md`, `agents/sdd/sdd-verify.md`, `agents/sdd/sdd-archive.md` | Modify | Align executor prompts with source-row validation duties and active embedded secure-design authority. |
| `openspec/specs/sdd-security-guideline-catalog/spec.md` | Modify | Sync source-row inventory, mapping, and safe-evidence requirements at archive. |
| `openspec/specs/sdd-design-workflow/spec.md` | Modify | Sync secure design Source ID coverage requirement at archive. |
| `openspec/specs/sdd-test-design-workflow/spec.md` | Modify | Sync source-row test-planning requirement at archive. |
| `openspec/specs/sdd-review-security-workflow/spec.md` | Modify | Sync exhaustive source-row security review requirement at archive. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modify | Sync source-row persistence, verify, and archive requirements at archive. |

## Interfaces / Contracts

### Source Row Schema

Each source-row evidence object uses this contract:

```yaml
sourceId: "1.1"
corporateSection: "Authentication"
guidelineTextRef: "skills/_shared/security-guideline-catalog.md#source-1.1"
pciAlignment: "PCI Req 6.5.8, 6.5.10" # or "N/A"
mappedCompactGuidelineIds: ["SEC-AUTH-001"]
applies: "Yes" # Yes | No | N/A; design uses Yes or N/A only
complies: "planned" # planned | Yes | No | N/A, phase-specific
lifecycleStatus: "planned"
evidenceLocation: "openspec/changes/<change>/design.md#secure-development-design"
observations: "Review-safe summary only."
finding: "none" # none | blocker | warning
route: "test-design" # apply | resolve-blockers | verify | archive | test-design
```

### Range Expansion Strategy

- Parse dotted IDs as `{section}.{row}` numeric tuples.
- Expand a range only when every concrete ID exists in the immutable snapshot table.
- Validate after expansion: no missing IDs, no duplicates, no unknown IDs, and no source row without at least one compact mapping.
- Persist expanded IDs in source-row artifacts/reports; compressed ranges may appear only as human summaries.

### Mapping Strategy for Currently Unmapped IDs

- Map by security concern to the existing compact taxonomy: auth, sessions, data/PAN, secrets, access control, files, database/input validation, and sensitive logging.
- Rows that express cross-cutting secure coding, crypto, memory, or output encoding obligations map to the closest compact risk owner instead of creating new compact controls.
- A Source ID may map to multiple compact controls, but it appears exactly once in the source-row inventory and review-security matrix.
- Any future unmapped Source ID is a contract blocker routed to `resolve-blockers`, not a review warning.

### Routing Semantics

| Condition | Route |
| --- | --- |
| Missing/duplicate/unknown Source ID, malformed schema, missing compact mapping, unsafe evidence, missing `N/A` justification | `resolve-blockers` |
| Applicable row lacks implementation evidence and remediation is file/contract/prompt work | `apply` |
| Only warning rows remain and mandatory evidence is complete | `verify` |
| Verification confirms non-blocking review-security evidence | `archive` |

### Persistence Compatibility

Source-row semantics do not vary by backend. OpenSpec stores rows in established files; Engram uses existing SDD topic keys; hybrid writes both and applies the shared conflict policy; none mode returns inline evidence and reports recovery limits.

## Testing Strategy

| Layer | What to Test | Approach |
| --- | --- | --- |
| Static/manual | Source inventory expansion, exact-once Source ID coverage, valid compact mappings, safe evidence, and routing text. | Planned in `test-design.md`; inspect Markdown/YAML artifacts because no runner is configured. |
| Integration | Phase-to-phase traceability from design to review-security, verify, and archive. | Static/manual artifact checks across generated SDD files and shared contracts. |
| Runtime/E2E | Not available in this repo. | `openspec/config.yaml` reports no test/build/coverage runner; downstream evidence must state this explicitly. |

## Migration / Rollout

No data migration required. Rollout is a documentation/contract change in one SDD change, delivered in reviewable work units and archived through OpenSpec.

## Implementation Work Units

1. Shared catalog and security contract: source-row schema, inventory expansion, mapping, safe-evidence, routing, and persistence-neutral semantics.
2. Design/test-design phase contracts and prompts: expected Source ID coverage and planned checks.
3. Review-security contract and prompt: exhaustive source-row matrix, blockers, warnings, and evidence correlation.
4. Verify/archive contracts and source specs: consume/preserve source-row evidence without duplicating matrices.

## Tradeoffs and Risks

| Risk / Tradeoff | Impact | Mitigation |
| --- | --- | --- |
| Source-row matrices are large. | Reviewer fatigue. | Keep compact summary first; source rows are operational detail below it. |
| Some rows are cross-cutting and imperfectly map to eight compact controls. | Mapping disputes. | Use row-to-many mappings and document rationale; missing mapping blocks. |
| Static/manual validation can miss parser bugs. | Lower automation confidence. | Make exact schema and expected checks explicit for future automation. |
| Evidence may leak sensitive values. | Security/regulatory exposure. | Safe-evidence policy blocks unsafe evidence and requires redacted references. |

## Open Questions

None.

## Secure Development Design

```yaml
schemaName: sdd.embedded-secure-design
schemaVersion: 1
changeName: corporate-source-row-security-validation
classification: security-impacting
securityImpact: true
securityImpactRationale: "The change modifies SDD security validation contracts, evidence routing, and archive gates for every corporate security Source ID."
sourceInputs:
  proposal: openspec/changes/corporate-source-row-security-validation/proposal.md
  specs:
    - openspec/changes/corporate-source-row-security-validation/specs/sdd-security-guideline-catalog/spec.md
    - openspec/changes/corporate-source-row-security-validation/specs/sdd-design-workflow/spec.md
    - openspec/changes/corporate-source-row-security-validation/specs/sdd-test-design-workflow/spec.md
    - openspec/changes/corporate-source-row-security-validation/specs/sdd-review-security-workflow/spec.md
    - openspec/changes/corporate-source-row-security-validation/specs/sdd-execution-persistence-contracts/spec.md
  design: openspec/changes/corporate-source-row-security-validation/design.md#secure-development-design
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
    rationale: "The change must map and validate all authentication/password/auth-url Source IDs through compact auth coverage."
    evidenceRefs: ["proposal.md lines 8-13", "specs/sdd-design-workflow/spec.md#secure-design-source-id-coverage"]
  - category: sessions
    guidelineId: SEC-SESS-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Session Source IDs 7.1-7.13 must be expanded, mapped, and validated below compact session coverage."
    evidenceRefs: ["skills/_shared/security-guideline-catalog.md#session-management", "specs/sdd-review-security-workflow/spec.md#exhaustive-source-row-security-review"]
  - category: sensitive-data-pan
    guidelineId: SEC-DATA-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "The source-row layer must preserve data/PAN, crypto, memory, and output-encoding obligations with safe evidence."
    evidenceRefs: ["proposal.md lines 11-12", "specs/sdd-security-guideline-catalog/spec.md#safe-source-row-evidence"]
  - category: secrets
    guidelineId: SEC-SECRET-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Evidence handling and mappings must prevent secret leakage and validate secret-related Source IDs."
    evidenceRefs: ["proposal.md lines 11-12", "skills/_shared/sdd-security-contract.md#safe-evidence-rules-for-mandatory-security-controls"]
  - category: permissions-access-control
    guidelineId: SEC-ACCESS-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Routing must deny unsafe progression for missing mappings, missing evidence, and unsupported N/A rows."
    evidenceRefs: ["proposal.md line 32", "specs/sdd-security-guideline-catalog/spec.md#shared-security-contract-source-row-schema"]
  - category: files
    guidelineId: SEC-FILE-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "File-handling Source IDs 9.1-9.12 and 14.7 require explicit mapping and future evidence coverage."
    evidenceRefs: ["skills/_shared/security-guideline-catalog.md#file-handling", "specs/sdd-security-guideline-catalog/spec.md#compact-sec-mapping-coverage"]
  - category: database-access
    guidelineId: SEC-DB-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Database, input-validation, and output-to-command/query Source IDs must map to compact database-access validation."
    evidenceRefs: ["skills/_shared/security-guideline-catalog.md#databases", "skills/_shared/security-guideline-catalog.md#input-validation"]
  - category: sensitive-logging
    guidelineId: SEC-LOG-001
    applies: Yes
    decision: applicable
    lifecycleStatus: planned
    rationale: "Logging and error-handling Source IDs must be validated while preserving review-safe observations."
    evidenceRefs: ["skills/_shared/security-guideline-catalog.md#access-and-activity-logging", "skills/_shared/security-guideline-catalog.md#error-handling"]
controls:
  - guidelineId: SEC-AUTH-001
    taxonomyCategory: authentication
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.1-1.10", "2.1-2.23", "4.3", "4.7", "6.3", "6.5", "6.14", "14.8"]
    requiredControl: "Catalog and phase contracts must expand, map, and validate authentication/password Source IDs under compact auth coverage."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "This design records source-row mapping and downstream evidence expectations."
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check that auth Source IDs are expanded, mapped, and covered exactly once."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Shared catalog/contract and phase prompt updates cite auth Source ID handling."
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: "Verify consumes non-blocking review-security source-row evidence."
    residualRisk: "Large auth/password row set may increase review effort; compact summary stays first."
    exception: null
  - guidelineId: SEC-SESS-001
    taxonomyCategory: sessions
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["7.1-7.13", "4.7", "6.5", "6.14"]
    requiredControl: "Session Source IDs must be expanded and mapped to compact session evidence without replacing the compact control."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check covers session source-row expansion and mapping."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Catalog mapping and review-security guidance include session source rows."
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: "Archive preserves review-security source-row verdicts."
    residualRisk: none
    exception: null
  - guidelineId: SEC-DATA-001
    taxonomyCategory: sensitive-data-pan
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["4.1", "4.3-4.7", "6.5-6.11", "10.1-10.6", "12.1", "12.3-12.5", "13.1-13.9", "15.1-15.2"]
    requiredControl: "Source-row evidence must preserve sensitive data/PAN obligations and use only review-safe references or redacted placeholders."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "Safe-evidence and N/A evidence policy are documented in this design."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Shared security contract blocks unsafe evidence and unknown data sensitivity."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Security review rejects evidence containing secrets, PAN, PII, tokens, connection strings, private keys, or confidential values."
    residualRisk: "Static/manual evidence cannot prove runtime behavior; verify must report unavailable tooling."
    exception: null
  - guidelineId: SEC-SECRET-001
    taxonomyCategory: secrets
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["2.1", "4.2", "4.5-4.8", "5.5", "6.1", "6.13", "13.5"]
    requiredControl: "Contracts must forbid raw secret evidence and require secret-related Source IDs to map to compact secret coverage."
    expectedEvidence:
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Security contract and review-security skill include safe-evidence secret restrictions."
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual evidence check verifies secret-safe evidence wording."
      - type: verification-evidence
        ownerPhase: verify
        status: planned
        detail: "Verify cites review-security verdict and does not reproduce sensitive evidence."
    residualRisk: none
    exception: null
  - guidelineId: SEC-ACCESS-001
    taxonomyCategory: permissions-access-control
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["1.4", "5.1", "6.2-6.5", "6.12-6.14", "13.1", "14.1-14.9"]
    requiredControl: "Missing source-row coverage, mapping, N/A evidence, or safe evidence must deny progression by routing to apply or resolve-blockers."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "Routing table defines denial-by-default behavior for source-row blockers."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Phase skills implement source-row blocker routing semantics."
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: "Archive refuses unresolved mandatory source-row blockers."
    residualRisk: none
    exception: null
  - guidelineId: SEC-FILE-001
    taxonomyCategory: files
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["9.1-9.12", "14.7"]
    requiredControl: "File-handling Source IDs must be mapped and validated in the source-row matrix for future file-related changes."
    expectedEvidence:
      - type: test-design-check
        ownerPhase: test-design
        status: planned
        detail: "Static/manual check validates file Source ID mapping and N/A evidence expectations."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Catalog mapping table includes file source rows and cross-map 14.7."
    residualRisk: none
    exception: null
  - guidelineId: SEC-DB-001
    taxonomyCategory: database-access
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["5.1-5.12", "6.5-6.10", "11.1-11.16", "12.2"]
    requiredControl: "Database/input-validation Source IDs must map to compact DB coverage and require evidence for query/input safety."
    expectedEvidence:
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Catalog and security contract include DB/input-validation mappings."
      - type: verification-evidence
        ownerPhase: review-security
        status: planned
        detail: "Security review fails applicable DB rows that lack corroborating evidence."
    residualRisk: none
    exception: null
  - guidelineId: SEC-LOG-001
    taxonomyCategory: sensitive-logging
    mandatoryWhenApplicable: true
    operationalSeverity: blocking
    sourceRefs: ["3.1-3.11", "8.1-8.5"]
    requiredControl: "Logging/error Source IDs must validate audit usefulness while blocking sensitive evidence leakage."
    expectedEvidence:
      - type: design-control
        ownerPhase: design
        status: planned
        detail: "Safe-evidence policy requires sanitized observations only."
      - type: implementation-reference
        ownerPhase: apply
        status: planned
        detail: "Review-security/verify/archive contracts cite source-row evidence without raw sensitive payloads."
      - type: verification-evidence
        ownerPhase: archive
        status: planned
        detail: "Archive preserves warning/blocker summaries and safe evidence references."
    residualRisk: none
    exception: null
sourceRowCoverage:
  schema: corporate-source-row-operational-layer
  exactSourceIdCount: 155
  coverageRule: "Every Source ID from the preserved snapshot is expanded and represented exactly once in source-row inventory/review matrices. Each row maps to one or more compact SEC-* IDs."
  safeEvidencePolicy: "Use paths, section refs, sanitized summaries, command summaries, or redacted placeholders only; never include secrets, PII, PAN, tokens, connection strings, private keys, or confidential values."
  notApplicablePolicy: "N/A requires evidence and justification proving category, platform, API, data class, or workflow irrelevance; unsupported N/A blocks with route resolve-blockers."
  groups:
    - corporateSection: "1. Authentication"
      pciAlignment: "PCI Req 6.5.8, 6.5.10"
      sourceIds: ["1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1.10"]
      mappedCompactGuidelineIds: ["SEC-AUTH-001", "SEC-ACCESS-001"]
    - corporateSection: "2. Passwords"
      pciAlignment: "PCI Req 6.3.1"
      sourceIds: ["2.1", "2.2", "2.3", "2.4", "2.5", "2.6", "2.7", "2.8", "2.9", "2.10", "2.11", "2.12", "2.13", "2.14", "2.15", "2.16", "2.17", "2.18", "2.19", "2.20", "2.21", "2.22", "2.23"]
      mappedCompactGuidelineIds: ["SEC-AUTH-001", "SEC-SECRET-001"]
    - corporateSection: "3. Access and Activity Logging"
      pciAlignment: "N/A"
      sourceIds: ["3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "3.7", "3.8", "3.9", "3.10", "3.11"]
      mappedCompactGuidelineIds: ["SEC-LOG-001"]
    - corporateSection: "4. Cryptography"
      pciAlignment: "PCI Req 6 - 6.5.3"
      sourceIds: ["4.1", "4.2", "4.3", "4.4", "4.5", "4.6", "4.7", "4.8"]
      mappedCompactGuidelineIds: ["SEC-DATA-001", "SEC-SECRET-001", "SEC-AUTH-001", "SEC-SESS-001"]
    - corporateSection: "5. Databases"
      pciAlignment: "N/A"
      sourceIds: ["5.1", "5.2", "5.3", "5.4", "5.5", "5.6", "5.7", "5.8", "5.9", "5.10", "5.11", "5.12"]
      mappedCompactGuidelineIds: ["SEC-DB-001", "SEC-ACCESS-001", "SEC-SECRET-001"]
    - corporateSection: "6. Coding"
      pciAlignment: "PCI Req 6.5.8, 6.5.9"
      sourceIds: ["6.1", "6.2", "6.3", "6.4", "6.5", "6.6", "6.7", "6.8", "6.9", "6.10", "6.11", "6.12", "6.13", "6.14"]
      mappedCompactGuidelineIds: ["SEC-SECRET-001", "SEC-ACCESS-001", "SEC-AUTH-001", "SEC-DATA-001", "SEC-DB-001", "SEC-SESS-001", "SEC-FILE-001", "SEC-LOG-001"]
    - corporateSection: "7. Session Management"
      pciAlignment: "PCI Req 6.5.9, 6.5.10"
      sourceIds: ["7.1", "7.2", "7.3", "7.4", "7.5", "7.6", "7.7", "7.8", "7.9", "7.10", "7.11", "7.12", "7.13"]
      mappedCompactGuidelineIds: ["SEC-SESS-001"]
    - corporateSection: "8. Error Handling"
      pciAlignment: "PCI Req 6.3.c, 6.5.5"
      sourceIds: ["8.1", "8.2", "8.3", "8.4", "8.5"]
      mappedCompactGuidelineIds: ["SEC-LOG-001"]
    - corporateSection: "9. File Handling"
      pciAlignment: "PCI Req 6.5.8"
      sourceIds: ["9.1", "9.2", "9.3", "9.4", "9.5", "9.6", "9.7", "9.8", "9.9", "9.10", "9.11", "9.12"]
      mappedCompactGuidelineIds: ["SEC-FILE-001"]
    - corporateSection: "10. Memory Management"
      pciAlignment: "PCI Req 6.5.1"
      sourceIds: ["10.1", "10.2", "10.3", "10.4", "10.5", "10.6"]
      mappedCompactGuidelineIds: ["SEC-DATA-001", "SEC-SECRET-001"]
    - corporateSection: "11. Input Validation"
      pciAlignment: "PCI Req 6.5.1, 6.5.7, 6.5.8, 6.5.9"
      sourceIds: ["11.1", "11.2", "11.3", "11.4", "11.5", "11.6", "11.7", "11.8", "11.9", "11.10", "11.11", "11.12", "11.13", "11.14", "11.15", "11.16"]
      mappedCompactGuidelineIds: ["SEC-DB-001"]
    - corporateSection: "12. Output Encoding"
      pciAlignment: "PCI Req 6.5.9"
      sourceIds: ["12.1", "12.2", "12.3", "12.4", "12.5"]
      mappedCompactGuidelineIds: ["SEC-DATA-001", "SEC-DB-001"]
    - corporateSection: "13. Data Protection"
      pciAlignment: "PCI Req 6.3.c, 6.5.4"
      sourceIds: ["13.1", "13.2", "13.3", "13.4", "13.5", "13.6", "13.7", "13.8", "13.9"]
      mappedCompactGuidelineIds: ["SEC-DATA-001", "SEC-ACCESS-001", "SEC-SECRET-001"]
    - corporateSection: "14. Access Control"
      pciAlignment: "N/A"
      sourceIds: ["14.1", "14.2", "14.3", "14.4", "14.5", "14.6", "14.7", "14.8", "14.9"]
      mappedCompactGuidelineIds: ["SEC-ACCESS-001", "SEC-FILE-001", "SEC-AUTH-001"]
    - corporateSection: "15. PAN"
      pciAlignment: "PCI Req 6.4"
      sourceIds: ["15.1", "15.2"]
      mappedCompactGuidelineIds: ["SEC-DATA-001"]
traceability:
  pattern: "Source ID -> mapped compact SEC-* -> design control -> test-design check -> apply implementation reference -> review-security row verdict -> verify prerequisite evidence -> archive preserved audit trail"
  downstreamOwners: ["test-design", "tasks", "apply", "review-security", "verify", "archive"]
notApplicableGuidelines: []
exceptions: null
carriedRisks:
  - risk: "Large source-row evidence can be noisy for reviewers."
    ownerPhase: tasks
    expectedEvidence: "Tasks should split implementation into reviewable work units and keep compact summaries before detailed matrices."
validation:
  method: design.md#secure-development-design static/manual review
  status: pass
  checkedAt: "2026-07-06"
  notes: "Design includes all eight compact SEC rows exactly once, expected source-row coverage for 155 Source IDs, safe-evidence policy, N/A policy, routing semantics, persistence compatibility, and downstream traceability. Runtime tooling is unavailable per openspec/config.yaml."
archiveGateNotes:
  - "Archive must preserve source-row coverage, mappings, warnings, exceptions, and evidence references without requiring standalone security-design artifacts."
nextRecommended: test-design
```

Compact matrix summary:

| Guideline | Applies / lifecycle | Rationale | Secure design decision / control | Evidence owner / expected evidence | Residual risk / exception |
| --- | --- | --- | --- | --- | --- |
| `SEC-AUTH-001` | `Yes` / `planned` | Auth/password Source IDs must be expanded and mapped. | Add source-row auth coverage below compact auth control. | test-design/apply/review-security/verify/archive source-row evidence. | Review effort risk; no exception. |
| `SEC-SESS-001` | `Yes` / `planned` | Session Source IDs 7.1-7.13 need exact coverage. | Add session source-row mapping and evidence checks. | test-design/apply/archive evidence. | None / no exception. |
| `SEC-DATA-001` | `Yes` / `planned` | Data/PAN, crypto, memory, and output rows need safe evidence. | Block unsafe evidence and unsupported N/A. | design/apply/review-security evidence. | Static/manual limitation; no exception. |
| `SEC-SECRET-001` | `Yes` / `planned` | Secret-related rows and evidence must never expose raw values. | Require redacted references and secret-safe observations. | apply/test-design/verify evidence. | None / no exception. |
| `SEC-ACCESS-001` | `Yes` / `planned` | Routing must deny progression for missing mappings/evidence. | Define source-row blocker routing. | design/apply/archive evidence. | None / no exception. |
| `SEC-FILE-001` | `Yes` / `planned` | File source rows must be mapped for future file changes. | Add file source-row coverage. | test-design/apply evidence. | None / no exception. |
| `SEC-DB-001` | `Yes` / `planned` | DB/input/output-to-query rows need compact DB mapping. | Add DB source-row mapping and corroboration rules. | apply/review-security evidence. | None / no exception. |
| `SEC-LOG-001` | `Yes` / `planned` | Logging/error rows must be validated safely. | Preserve audit usefulness without sensitive payloads. | design/apply/archive evidence. | None / no exception. |

Source-row traceability summary:

| Corporate sections | Source IDs | Compact mapping | Downstream evidence path |
| --- | --- | --- | --- |
| Authentication, Passwords | `1.1`-`1.10`, `2.1`-`2.23` | `SEC-AUTH-001`, `SEC-ACCESS-001`, `SEC-SECRET-001` | design -> test-design -> apply -> review-security -> verify -> archive |
| Logging, Error Handling | `3.1`-`3.11`, `8.1`-`8.5` | `SEC-LOG-001` | design -> test-design -> apply -> review-security -> verify -> archive |
| Cryptography, Data Protection, PAN, Memory, Output Encoding | `4.1`-`4.8`, `10.1`-`10.6`, `12.1`-`12.5`, `13.1`-`13.9`, `15.1`-`15.2` | `SEC-DATA-001`, `SEC-SECRET-001`, contextual `SEC-AUTH-001`/`SEC-SESS-001`/`SEC-DB-001` | design -> test-design -> apply -> review-security -> verify -> archive |
| Databases, Coding, Input Validation | `5.1`-`5.12`, `6.1`-`6.14`, `11.1`-`11.16` | `SEC-DB-001`, `SEC-ACCESS-001`, `SEC-DATA-001`, `SEC-SECRET-001`, cross-cutting compact controls | design -> test-design -> apply -> review-security -> verify -> archive |
| Session, File, Access Control | `7.1`-`7.13`, `9.1`-`9.12`, `14.1`-`14.9` | `SEC-SESS-001`, `SEC-FILE-001`, `SEC-ACCESS-001`, contextual `SEC-AUTH-001` | design -> test-design -> apply -> review-security -> verify -> archive |
