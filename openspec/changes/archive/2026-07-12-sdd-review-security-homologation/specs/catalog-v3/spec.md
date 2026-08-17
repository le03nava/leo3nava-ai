# Delta for catalog-v3

## Purpose

Delta spec for `security-guideline-catalog.operational.json`. Covers the v3 structural cleanup: header field removal, `schemaVersion` bump, and per-row field renames/removals. No content is lost — all 155 rows and their `guidelineText` values are preserved.

## ADDED Requirements

### Requirement: CATALOG-001 — Schema Version

The catalog MUST declare `schemaVersion` equal to `3` at the top-level header.

#### Scenario: Valid v3 catalog is loaded

- GIVEN a catalog file at `skills/sdd-review-security/references/security-guideline-catalog.operational.json`
- WHEN a consumer reads the top-level `schemaVersion` field
- THEN the value MUST equal `3` (integer)

#### Scenario: Pre-v3 catalog is detected

- GIVEN a catalog file where `schemaVersion` is less than `3`
- WHEN a consumer validates the catalog
- THEN the consumer MUST reject it as incompatible

---

### Requirement: CATALOG-002 — generatedHumanViewRef Removed from Header

The catalog header MUST NOT contain a `generatedHumanViewRef` field.

#### Scenario: Header is validated for removed field

- GIVEN a v3 catalog header
- WHEN a validator checks for `generatedHumanViewRef`
- THEN the field MUST NOT be present

---

### Requirement: CATALOG-003 — reporting Object Removed from Header

The catalog header MUST NOT contain a `reporting` object.

#### Scenario: Header is validated for removed object

- GIVEN a v3 catalog header
- WHEN a validator checks for `reporting`
- THEN the key MUST NOT be present at any nesting level of the header

---

### Requirement: CATALOG-004 — authority Object Removed from Header

The catalog header MUST NOT contain an `authority` object.

#### Scenario: Header is validated for removed object

- GIVEN a v3 catalog header
- WHEN a validator checks for `authority`
- THEN the key MUST NOT be present at any nesting level of the header

---

### Requirement: CATALOG-005 — ownerPhase Field Present on Every sourceRow

Every `sourceRow` in the catalog MUST have an `ownerPhase` field (renamed from `defaultOwnerPhase`).

#### Scenario: All rows have ownerPhase

- GIVEN a v3 catalog with 155 sourceRows
- WHEN a validator iterates all `sourceRows`
- THEN every row MUST contain a non-null `ownerPhase` field

#### Scenario: No row carries the old defaultOwnerPhase name

- GIVEN a v3 catalog
- WHEN a validator counts rows where `has("defaultOwnerPhase")` is true
- THEN the count MUST equal `0`

---

### Requirement: CATALOG-006 — route Field Present on Every sourceRow

Every `sourceRow` in the catalog MUST have a `route` field (renamed from `defaultRoute`).

#### Scenario: All rows have route

- GIVEN a v3 catalog with 155 sourceRows
- WHEN a validator iterates all `sourceRows`
- THEN every row MUST contain a non-null `route` field

#### Scenario: No row carries the old defaultRoute name

- GIVEN a v3 catalog
- WHEN a validator counts rows where `has("defaultRoute")` is true
- THEN the count MUST equal `0`

---

### Requirement: CATALOG-007 — guidelineRefs Removed from sourceRows

No `sourceRow` MAY contain a `guidelineRefs` field.

#### Scenario: Rows are checked for guidelineRefs

- GIVEN a v3 catalog
- WHEN a validator scans all `sourceRows`
- THEN no row MUST contain the key `guidelineRefs`

---

### Requirement: CATALOG-008 — evidenceExpectation Removed from sourceRows

No `sourceRow` MAY contain an `evidenceExpectation` field.

#### Scenario: Rows are checked for evidenceExpectation

- GIVEN a v3 catalog
- WHEN a validator scans all `sourceRows`
- THEN no row MUST contain the key `evidenceExpectation`

---

### Requirement: CATALOG-009 — defaultOwnerPhase Removed from sourceRows

No `sourceRow` MAY contain a `defaultOwnerPhase` field.

#### Scenario: Legacy field absence is confirmed

- GIVEN a v3 catalog
- WHEN a validator counts rows with `defaultOwnerPhase`
- THEN the count MUST equal `0`

---

### Requirement: CATALOG-010 — defaultRoute Removed from sourceRows

No `sourceRow` MAY contain a `defaultRoute` field.

#### Scenario: Legacy field absence is confirmed

- GIVEN a v3 catalog
- WHEN a validator counts rows with `defaultRoute`
- THEN the count MUST equal `0`

---

### Requirement: CATALOG-011 — All 155 sourceRows Preserved

All 155 `sourceRows` MUST be present in the v3 catalog with `guidelineText` and all content fields intact.

#### Scenario: Row count is verified

- GIVEN a v3 catalog
- WHEN a consumer counts `sourceRows`
- THEN the count MUST equal `155`

#### Scenario: guidelineText survives the migration

- GIVEN any sourceRow in the v3 catalog
- WHEN a consumer reads `guidelineText`
- THEN the value MUST be non-empty and MUST match the pre-migration value for that `sourceId`

---

### Requirement: CATALOG-012 — vocabulary Does Not Contain safeEvidence or forbiddenEvidence

The catalog `vocabulary` section MUST NOT contain `safeEvidence` or `forbiddenEvidence` keys.

#### Scenario: vocabulary is validated for absent keys

- GIVEN a v3 catalog with a `vocabulary` section
- WHEN a validator checks for `safeEvidence` and `forbiddenEvidence`
- THEN neither key MUST be present (no-op confirmation — keys are not currently present)
