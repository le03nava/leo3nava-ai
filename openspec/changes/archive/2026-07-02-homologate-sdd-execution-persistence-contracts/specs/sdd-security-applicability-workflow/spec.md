# Delta for sdd-security-applicability-workflow

## MODIFIED Requirements

### Requirement: Artifact and Routing Contract

The phase MUST produce `security-applicability.md` with classification, evidence, guideline mapping, risks, and routing recommendation. Security-impacting changes MUST route to `sdd-security-design`; no-impact changes MUST skip security design and continue normal design workflow. The phase artifact contract MUST preserve the existing artifact key/path and routing behavior while delegating common artifact-store mode semantics, artifact resolution, and persistence verification to the shared persistence authority.
(Previously: The requirement defined the artifact and conditional routing, but did not state the shared persistence boundary.)

#### Scenario: Artifact drives conditional routing

- GIVEN `security-applicability.md` marks a change as security-impacting
- WHEN the orchestrator computes next phases
- THEN it MUST require `sdd-security-design`
- AND design-related successors MUST receive the applicability artifact reference.

#### Scenario: Persistence boundary is delegated

- GIVEN `sdd-security-applicability` writes or reads its artifact in any artifact-store mode
- WHEN the executor resolves backend behavior
- THEN it MUST use the shared persistence authority for mode semantics and resolver behavior
- AND it MUST keep `security-applicability.md` as the phase artifact contract.

#### Scenario: No-impact routing compatibility is preserved

- GIVEN `security-applicability.md` records no-impact evidence
- WHEN downstream routing is computed
- THEN the workflow MUST continue to skip `sdd-security-design`
- AND missing `security-design.md` MUST NOT become a blocker solely because persistence wording changed.
