## Secure Design Narrative Guidance

This reference defines narrative secure-design expectations for SDD phases.

Hard boundary:

- No control-ID matrix formats
- No source-row catalog rows
- No scored checklist tables
- No row-based scoring catalogs in this file

### 1) Changed-Surface Classification

- design commitment: The design must explicitly classify touched surfaces: inputs, outputs, storage, logs, exports, auth, sessions, permissions, files, database, config/secrets, and external APIs.
- prohibited unsafe pattern: Leaving surface scope implicit, or treating a surface as "not applicable" without evidence.
- expected safe evidence: `design.md` contains a changed-surface classification list or table. Any uncertain surface is marked `unknown` and flagged for confirmation.
- evidence owner phase: design
- residual risk: Misclassified surfaces can escape deeper review and produce blind spots.
- exception policy: If a surface is intentionally out of scope, document why and the decision boundary in `design.md`.

### 2) Data Sensitivity

- design commitment: Every data element touched by the change is classified as `public`, `internal`, `confidential`, `PII`, `PAN`, `secrets`, `tokens`, or `unknown`.
- prohibited unsafe pattern: Treating unclassified data as public by default; logging or exposing confidential/PII/PAN/secrets/tokens.
- expected safe evidence: `design.md` lists data elements with sensitivity labels. Unknown classifications are flagged and block assumption.
- evidence owner phase: design, review-security
- residual risk: Classification can drift from real usage as implementation evolves.
- exception policy: Any sensitivity downgrade (for example, `PII` to `internal`) requires explicit approval documented in `design.md`.

### 3) Access Control

- design commitment: The design documents authentication mechanism, authorization boundaries, ownership or tenant boundaries, and privilege checks for state-changing or data-exposing operations.
- prohibited unsafe pattern: Implicit authorization assumptions ("trusted caller"), missing ownership boundaries, or undocumented privilege escalation paths.
- expected safe evidence: `design.md` defines auth/authz boundaries; `apply-progress.md` records where checks are implemented.
- evidence owner phase: design, apply, review-security
- residual risk: Edge cases such as concurrent ownership changes may still be under-tested.
- exception policy: Read-only public data may declare `no auth required` only with explicit justification in `design.md`.

### 4) Input/Output Handling

- design commitment: External inputs are validated, sanitized, and canonicalized before use. File/path inputs are protected against traversal. Generated artifacts avoid unsafe content.
- prohibited unsafe pattern: Trusting caller-provided file paths without canonicalization, executing shell commands with unescaped input, or generating outputs from unsanitized content.
- expected safe evidence: `design.md` identifies input sources and handling approach; `apply-progress.md` notes where validation/canonicalization is applied.
- evidence owner phase: design, apply, verify
- residual risk: Canonicalization and encoding behavior may vary by environment.
- exception policy: Internal-only controlled inputs may use reduced validation only when justified in `design.md`.

### 5) Secrets and Config

- design commitment: No hardcoded secrets, tokens, API keys, or connection strings in source files, artifacts, or SDD evidence. Sensitive config values use environment variables or secret managers.
- prohibited unsafe pattern: Hardcoded secrets in code or artifacts; tokens/connection strings in `design.md`, `apply-progress.md`, or other SDD documents; logging secret values.
- expected safe evidence: SDD artifacts use placeholders (for example, `<SECRET_VALUE>`) where secret-like values would appear. `apply-progress.md` confirms no hardcoded secrets.
- evidence owner phase: apply, review-security
- residual risk: Future maintenance can reintroduce hardcoded secrets if review discipline degrades.
- exception policy: None. Hardcoded secret exceptions are not allowed.

### 6) Logging and Error Handling

- design commitment: Errors are actionable for operators without exposing sensitive data. Logs use sanitized summaries only.
- prohibited unsafe pattern: Logging raw request/response payloads, exposing stack traces to end users, or returning internal identifiers in external error messages.
- expected safe evidence: `design.md` defines logging and error-handling approach; `apply-progress.md` confirms sanitized patterns.
- evidence owner phase: design, apply, verify
- residual risk: Debug modes can expose excessive detail if enabled in production.
- exception policy: Verbose debug logging is allowed only behind a gated flag and documented in operational runbooks.

### 7) Dependencies and External Interfaces

- design commitment: New dependencies and external APIs are assessed for maintenance health, license compatibility, and known vulnerabilities. Network calls define timeout and retry behavior.
- prohibited unsafe pattern: Adding unmaintained or unknown-license dependencies, making network calls without timeout, or trusting external responses without validation.
- expected safe evidence: `design.md` lists dependency assessments; `apply-progress.md` notes timeout/retry/validation implementation.
- evidence owner phase: design, apply, review-security
- residual risk: Supply-chain risk remains continuous and cannot be eliminated by one change.
- exception policy: Existing unchanged dependencies do not require reassessment unless version changes.

### 8) Operational Security

- design commitment: SDD artifacts must not include production secrets, real customer data, or sensitive payloads. Security-relevant operations include monitoring/alerting and safe recovery behavior.
- prohibited unsafe pattern: Committing real secrets/PII into `design.md`, `apply-progress.md`, `test-design.md`, or archive artifacts; silent failures without monitoring hooks.
- expected safe evidence: SDD artifacts contain sanitized examples only; `design.md` documents monitoring approach when applicable; archive outputs remain sanitized.
- evidence owner phase: apply, review-security, archive
- residual risk: Operational runbooks and real deployments can drift over time.
- exception policy: When realistic fixtures are needed, use explicitly labeled synthetic data generation.

### 9) Evidence Expectations by Phase

- design commitment: Every SDD phase contributes explicit security evidence appropriate to its scope.
- prohibited unsafe pattern: Deferring all security evidence to review phases or leaving phase expectations implicit.
- expected safe evidence:
  - `design`: changed-surface classification, data sensitivity list, auth/authz approach, and secure design constraints where applicable.
  - `test-design`: planned security-relevant checks (input validation, auth boundaries, logging behavior).
  - `tasks`: explicit security implementation tasks (for example, validation and authorization tasks).
  - `apply`: `apply-progress.md` maps implemented security commitments and confirms no secrets in artifacts.
  - `review-security`: `review-security-report.json` remains the sole formal source-row-first security validation authority.
  - `verify`: verification confirms evidence consistency across design, test-design, tasks, apply, and review-security.
  - `archive`: archived evidence remains sanitized and excludes sensitive data.
- evidence owner phase: design, test-design, tasks, apply, review-security, verify, archive
- residual risk: Evidence completeness may vary if upstream phases under-specify security scope.
- exception policy: Phase-level omissions must be explicitly documented with rationale and remediation plan before closure.
