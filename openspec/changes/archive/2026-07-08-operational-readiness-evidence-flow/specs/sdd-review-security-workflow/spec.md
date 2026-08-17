# Delta for sdd-review-security-workflow

## ADDED Requirements

### Requirement: Operational Evidence Leakage Review

`sdd-review-security` MUST validate that operational-readiness evidence does not leak restricted production data in ordinary SDD evidence, code, tests, fixtures, or examples. It MUST focus on leakage, secrets, and restricted production identifiers, while `sdd-review` owns general existence, traceability, and placeholder checks.

#### Scenario: Restricted production identifier is found

- GIVEN ordinary SDD evidence contains a production hostname, IP, port, SID/service name, credential, token, payload, full ID list, or generated file bytes
- WHEN security review evaluates readiness evidence
- THEN it MUST report unsafe evidence as blocking
- AND route to `resolve-blockers` unless implementation remediation is required.

#### Scenario: Operational document is the target

- GIVEN the user explicitly provides production operational values for final documentation
- WHEN security review evaluates ordinary SDD artifacts
- THEN those values MUST remain outside SDD evidence, code, tests, fixtures, and examples.

### Requirement: Safe Placeholder Security Boundary

Security review MUST accept exact `Pendiente de confirmar:` and exact `No aplica.` placeholders as safer than invented operational details when evidence is unavailable or inapplicable. It MUST NOT require disclosure of real operational data to pass.

#### Scenario: Missing value uses safe placeholder

- GIVEN operational data is unavailable
- WHEN security review validates the evidence boundary
- THEN exact placeholder usage MUST NOT be treated as a leakage failure.

#### Scenario: Placeholder hides required security evidence

- GIVEN a security obligation still requires proof of non-leakage
- WHEN only a placeholder is present
- THEN security review MUST require safe evidence of the boundary, not real operational values.
