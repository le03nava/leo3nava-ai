# sdd-test-design-workflow Specification

## Purpose

Define the mandatory `sdd-test-design` phase that produces `test-design.md` after technical design with narrative secure development design and before task planning in the SDD workflow.

## Requirements

### Requirement: Mandatory Phase Order

The SDD workflow MUST run `sdd-test-design` after `sdd-design` succeeds for every new change. `sdd-design` MUST include embedded secure development design, and `sdd-tasks` MUST be blocked until `test-design.md` exists or the phase explicitly blocks.

#### Scenario: Design routes through test design

- GIVEN an SDD change has completed `sdd-design`
- WHEN the phase returns a successful envelope
- THEN the next recommended phase MUST be `test-design` or `sdd-test-design`
- AND `security-design` MUST NOT be the direct successor for new changes.

#### Scenario: Tasks requested too early

- GIVEN `test-design.md` is missing for an active change
- WHEN `sdd-tasks` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing `sdd-test-design` phase or artifact.

#### Scenario: Test design requires embedded secure design

- GIVEN `design.md` lacks `## Secure Development Design` for a new active change
- WHEN `sdd-test-design` is launched
- THEN the launch MUST be rejected or returned as blocked
- AND the blocker MUST name the missing embedded secure design section.

### Requirement: test-design.md Artifact Contract

The `sdd-test-design` phase MUST create `test-cases.json` with all cases set to `status: pending` as the first output. It MUST then derive `test-design.md` from that JSON as a human-readable view. `test-design.md` is read-only for humans and MUST NOT be treated as source of truth; downstream phases MUST consume `test-cases.json` for case data. The artifact MUST map spec scenarios, design risks, and applicable narrative category rules from `design.md#secure-development-design` to planned automated, manual, or static checks; mark each case as mandatory or non-mandatory; state expected evidence; and document justified no-impact assessments from design classification. It MUST NOT require design to include YAML, schema fields, compact controls, Source IDs, matrices, or all-row `N/A` evidence. Common persistence behavior MUST remain delegated to the shared persistence authority.

#### Scenario: JSON is produced before Markdown

- GIVEN `sdd-test-design` is ready to emit its outputs
- WHEN the phase runs
- THEN `test-cases.json` MUST be written and verified before `test-design.md` is generated
- AND `test-design.md` MUST be derived solely from the content of `test-cases.json`

#### Scenario: test-design.md reflects JSON content faithfully

- GIVEN `test-cases.json` has been persisted with N cases
- WHEN `test-design.md` is derived
- THEN it MUST list all N cases with their type, mandatory flag, and expected evidence
- AND it MUST NOT contain case data that is absent from the JSON

#### Scenario: Markdown is not authoritative

- GIVEN `test-design.md` and `test-cases.json` both exist for a change
- WHEN a downstream phase resolves test case data
- THEN it MUST read `test-cases.json` as the source of truth
- AND it MUST NOT treat manually edited Markdown as an authoritative update

#### Scenario: Behavior-impacting change

- GIVEN specs or design describe behavior, contracts, routing, or compatibility changes
- WHEN `sdd-test-design` runs
- THEN `test-cases.json` MUST list planned checks linked to those inputs
- AND each check MUST include type, mandatory flag, and expected evidence

#### Scenario: Applicable security controls are consumed

- GIVEN `design.md` lists applicable secure development category rules
- WHEN `sdd-test-design` runs
- THEN `test-cases.json` MUST include checks or justified non-test evidence for mandatory rules
- AND blocked security obligations MUST remain blockers

#### Scenario: No-impact assessment is handled

- GIVEN design classifies the change as no security impact with changed-surface rationale
- WHEN `sdd-test-design` runs
- THEN it MUST cite that assessment in `test-cases.json`
- AND it MUST still produce `test-cases.json` and derived `test-design.md` without requiring
  YAML, schema, matrix, or all-row `N/A` content

#### Scenario: No-impact change

- GIVEN the change has no behavior, security, or testability impact
- WHEN `sdd-test-design` runs
- THEN `test-cases.json` MUST state a no-impact assessment
- AND downstream phases MUST treat the artifact as complete rather than absent

#### Scenario: Persistence boundary is delegated

- GIVEN `sdd-test-design` writes or resolves `test-cases.json` and `test-design.md`
  in any supported artifact-store mode
- WHEN backend behavior is required
- THEN it MUST follow the shared persistence authority
- AND it MUST keep the mandatory artifact contract local to the phase

### Requirement: Check Types and Severity

Test-design cases MUST support automated checks, manual checks, and static/file-contract checks. Mandatory cases MUST be treated as verification-blocking; non-mandatory cases MUST be advisory and reported as warnings when uncovered.

#### Scenario: Mandatory case uncovered

- GIVEN `test-design.md` marks a case as mandatory
- WHEN `sdd-verify` finds no matching implementation, execution, or justified skip evidence
- THEN verification MUST fail or report a blocking issue for that case

#### Scenario: Non-mandatory case uncovered

- GIVEN `test-design.md` marks a case as non-mandatory
- WHEN `sdd-verify` finds no matching evidence
- THEN verification MUST report a warning
- AND verification MUST NOT fail solely because of that uncovered non-mandatory case

### Requirement: Downstream Consumption

Downstream phases MUST use `test-design.md` as the test-planning source of truth and MUST preserve links to embedded secure development design evidence. `sdd-tasks` MUST derive implementation and verification tasks from it; `sdd-apply` SHOULD follow planned cases or report justified deviations; `sdd-verify` MUST compare evidence against both test cases and security evidence obligations. Updates to persistence contract wording MUST NOT weaken this downstream consumption requirement or route tasks directly from design without the required test-design artifact.

#### Scenario: Verify compares security evidence

- GIVEN test design references embedded secure development design evidence
- WHEN `sdd-verify` runs
- THEN it MUST evaluate matching evidence or justified exceptions
- AND missing mandatory evidence MUST remain blocking.

#### Scenario: Tasks derive testing work

- GIVEN `test-design.md` contains planned checks
- WHEN `sdd-tasks` creates implementation tasks
- THEN tasks MUST include corresponding testing or evidence work
- AND omitted mandatory cases MUST be reported as blockers

#### Scenario: Apply deviates from plan

- GIVEN implementation cannot follow a planned case exactly
- WHEN `sdd-apply` records evidence
- THEN it SHOULD document the deviation and replacement evidence
- AND `sdd-verify` MUST evaluate that justification against the planned case

#### Scenario: Persistence refactor preserves consumption

- GIVEN common persistence rules move to a shared authority
- WHEN downstream phases resolve test-planning inputs
- THEN they MUST still consume `test-design.md` from the established key or path
- AND they MUST NOT treat the artifact as optional.

### Requirement: Continuation, State, and Status Contracts

SDD state, status, continuation, and artifact refs MUST include `test-design.md` and the `sdd-test-design` dependency edge. Continuation MUST recover from persisted state without routing from design directly to tasks.

#### Scenario: Continuation after design

- GIVEN persisted state has design complete and test design incomplete
- WHEN continuation or status is requested
- THEN the next recommended phase MUST be `test-design` or `sdd-test-design`
- AND status MUST show tasks blocked by test design

#### Scenario: Artifact refs persisted

- GIVEN `test-design.md` has been created
- WHEN state or status is generated
- THEN artifact refs and paths MUST include the test-design artifact
- AND downstream launches MUST receive that ref when available

### Requirement: Native Token Compatibility

The workflow MUST remain compatible with native dispatcher and status-token handling. Implementations MUST define a bounded mapping between phase-agent token `sdd-test-design`, native/status token `test-design`, and any persisted field names required by native tooling.

#### Scenario: Native dispatcher supports test-design

- GIVEN native status or dispatcher tooling accepts a bounded `test-design` token
- WHEN the orchestrator routes after design
- THEN it MUST use that native token consistently in state and status

#### Scenario: Native dispatcher lacks test-design token

- GIVEN native tooling rejects or omits `test-design`
- WHEN the workflow is updated
- THEN design MUST specify compatibility handling before apply
- AND prompts MUST NOT claim unsupported native status behavior without a mapping

### Requirement: Source Row Test Planning

`sdd-test-design` MUST plan static, manual, or automated checks from applicable narrative category rules and changed-surface context in `design.md#secure-development-design`. It MUST consume catalog snapshot/context only as needed to understand categories, evidence owners, residual risks, exceptions, and safe-evidence policy. It MUST NOT require design to carry YAML, schema fields, compact controls, Source IDs, matrices, or `N/A` rows for omitted controls, and MUST preserve `sdd-review-security` as the exhaustive machine-readable validation owner.

#### Scenario: Applicable category receives checks

- GIVEN design defines Sensitive Data/PAN development rules
- WHEN test design is produced
- THEN `test-design.md` MUST include planned checks or justified non-test evidence
- AND each check MUST cite the narrative rule or category context.

#### Scenario: No runtime runner exists

- GIVEN repository testing capabilities report no runner
- WHEN source-row checks are planned
- THEN test design MUST use static or manual evidence where appropriate
- AND it MUST report unavailable runtime automation explicitly.

#### Scenario: Mandatory coverage lacks planned evidence

- GIVEN an applicable mandatory category rule has no check or evidence plan
- WHEN `sdd-test-design` validates readiness
- THEN the phase MUST block
- AND the blocker MUST name the affected category rule and missing evidence plan.

### Requirement: Source Row N/A and Warning Evidence

`sdd-test-design` MUST preserve evidence expectations for applicable narrative rules, warning-only coverage, residual risks, and exceptions from design. It MUST NOT create or require exhaustive `N/A` evidence for every non-applicable compact control or Source ID. Exhaustive `N/A` decisions, compact/Source ID matrices, and missed-applicable validation MUST remain owned by canonical `review-security-report.json`, with Markdown as a derived compatibility view.

#### Scenario: Omitted coverage remains reviewable

- GIVEN design omits a category as non-applicable
- WHEN test design maps security evidence
- THEN it MUST rely on design classification context only
- AND review-security MUST later validate whether the omission was safe.

#### Scenario: Warning-only coverage is tracked

- GIVEN source coverage has warning classification only
- WHEN test design records checks
- THEN it MUST preserve the warning and expected evidence
- AND it MUST allow downstream routing when no blockers exist.

### Requirement: Operational Readiness Test Planning

`sdd-test-design` MUST derive operational checks only from `design.md#operational-considerations` or equivalent design evidence. When design omits operational considerations or marks them `No aplica.`, test design MUST NOT create mandatory readiness checks. When considerations exist, checks SHOULD validate applicable logs/errors, monitoring, administration operations, reprocessing/recovery, backup/retention/cleanup/generated artifacts, evidence traceability, and safe-evidence boundaries. Missing runtime tooling MUST NOT prevent static or documentary checks.

#### Scenario: Design includes operational considerations

- GIVEN design includes applicable operational considerations
- WHEN `test-design.md` is produced
- THEN it MUST list checks for the applicable concerns
- AND each mandatory check MUST state expected evidence.

#### Scenario: Design has no operational considerations

- GIVEN design omits the section or states `No aplica.`
- WHEN test design is produced
- THEN no mandatory operational-readiness checks MUST be added
- AND downstream phases MUST consume the design decision.

#### Scenario: No runtime runner exists

- GIVEN repository testing capabilities report no executable runner
- WHEN operational checks are applicable
- THEN test design MUST use static or documentary evidence where appropriate.

### Requirement: Operational Data Safety Checks

Test design MUST plan restricted-data checks only when operational evidence exists or is planned. These checks MUST verify ordinary SDD evidence, code, tests, fixtures, and examples do not contain restricted production operational data, while distinguishing safe SDD evidence from final operational documentation inputs.

#### Scenario: Restricted data check is required

- GIVEN operational evidence references infrastructure or generated artifacts
- WHEN test design records validation work
- THEN it MUST include checks for restricted identifiers, secrets, payloads, full ID lists, and generated bytes.

#### Scenario: No operational evidence exists

- GIVEN design states operational considerations do not apply
- WHEN test design validates scope
- THEN it MUST NOT require operational data safety checks solely for readiness completeness.
