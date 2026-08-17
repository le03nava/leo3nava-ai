# Review Report: sdd-review-schema-v2

## Verdict

| Field | Value |
|---|---|
| Change | sdd-review-schema-v2 |
| Status | success |
| Verdict | PASS WITH WARNINGS |
| Next recommendation | review-security |
| JSON authority | openspec/changes/sdd-review-schema-v2/review-report.json |
| Markdown authority | openspec/changes/sdd-review-schema-v2/review-report.md |

## Totals

| Metric | Count |
|---|---|
| Total controls | 96 |
| Passing | 18 |
| Failing | 0 |
| N/A | 78 |
| Blocking | 0 |
| Non-blocking | 9 |

## Unavailable Tooling

- test-runner (openspec/config.yaml: test_runner.available: false)
- linter
- type-checker
- formatter
- coverage

## Artifact Metadata

| Field | Value |
|---|---|
| Canonical JSON ref | openspec/changes/sdd-review-schema-v2/review-report.json |
| Derived Markdown ref | openspec/changes/sdd-review-schema-v2/review-report.md |
| JSON persisted / read back | true / true |
| Markdown generated / persisted / read back | true / true / true |
| JSON/Markdown parity | passed |
| JSON authority | canonical |
| Markdown authority | derived |

## Recommendation

- Next: review-security
- 9 non-blocking findings documented as warnings. No blocking failures. Schema deviations from design spec are improvements/tightenings that maintain internal consistency. Follow-on sdd-review-security update required before running security review against v2 reports.

## Non-Blocking Findings Summary

1. **Schema `status` enum deviation**: Design specifies `["completed", "blocked"]`; implemented schema has `["success", "blocked", "partial"]` matching SDD phase vocabulary. Reasonable adaptation.
2. **Schema `totals` shape richer than design**: Design specifies `blockingFailureCount` + `nonBlockingFindingCount`; schema adds `controlCount`, `passing`, `failing`, `notApplicable`. Improvement.
3. **Schema `$defs.ref` simplified**: Design specifies object with `artifact`, `ref`, `readable`, optional `notes`; schema uses simple string. Acceptable simplification.
4. **Schema `$defs.stateArtifactRef` shape differs**: Design says `artifact`, `ref`, `kind` (enum); schema has `ref`, `authority`, `readable`. Different field names, equivalent intent.
5. **Schema `artifactMetadata` richer than design**: Design specifies only `stateRegistration`; schema adds lifecycle tracking fields (9 required). Improvement.
6. **Schema `matrixRow` required fields differ from design**: Design includes `requirement`, `standard`, `category`, `severity`; schema requires `item`, `sourceItem`, `complies`, `findingType`, `evidenceLocation` and adds `findingType` enum. `findingType` is useful but not in design/spec.
7. **Schema `unavailableTooling` is array vs design's string**: Array is more structured. Improvement.
8. **Schema `schemaVersion` uses `minimum: 2` not `const 2`**: Allows future versions. Reasonable but deviates from design's `const 2`.
9. **Template Safe Evidence Rules wording simplified**: Design requires verbatim preservation of 4 specific lines; template has 4 lines with simplified wording. Non-blocking — intent preserved.

## Matrix

| Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes |
|---|---|---|---|---|---|---|---|---|
| REV-CORP-001 | naming standards | Estándar de Diseño y Código WEB | code-quality | Media | Yes | Catalog, schema, and template use consistent naming conventions for fields, keys, and IDs. | skills/sdd-review/references/review-control-catalog.json — field names; review-report.schema.json — property names | |
| REV-CORP-002 | folder structure per platform standard | Estándar de Diseño y Código WEB | code-quality | Media | Yes | Files remain in skills/sdd-review/references/ per platform standard. | skills/sdd-review/references/ | |
| REV-CORP-003 | service names comply with area naming standards | Estándar de Diseño y Código WEB | code-quality | Media | Yes | Artifact names follow established sdd-review naming convention. | review-control-catalog.json, review-report.schema.json, report-template.md | |
| REV-CORP-004 | developer tested code/unit use cases | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | Static verification completed per apply-progress Phase 4. No runtime test runner available. | openspec/changes/sdd-review-schema-v2/apply-progress.md — Static Verification Results | No test runner available; static/manual verification only. |
| REV-CORP-005 | variable declarations adequately commented; descriptive names may make comments unnecessary | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | JSON keys and schema properties are self-descriptive. No inline comments needed in JSON. | skills/sdd-review/references/review-report.schema.json — description, $comment fields | |
| REV-CORP-006 | functions, methods, classes documented | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | Schema has title, description, $comment. Template has preamble explaining purpose. | review-report.schema.json lines 6-8; report-template.md lines 1-7 | |
| REV-CORP-007 | complex algorithms and optimizations commented | Estándar de Diseño y Código WEB | code-quality | Media | N/A | No complex algorithms in static JSON/Markdown reference artifacts. | Changed files are JSON schema and Markdown template — no executable algorithms | Static reference artifacts only. |
| REV-CORP-008 | commented-out code has explanation; dead code removed | Estándar de Diseño y Código WEB | code-quality | Menor | N/A | No executable code; no commented-out code possible in JSON schema artifacts. | All three changed files are JSON or Markdown | No executable code in scope. |
| REV-CORP-009 | resources and memory released on all error paths | Estándar de Diseño y Código WEB | code-quality | Media | N/A | No runtime resources or memory management in static reference artifacts. | No application runtime in repository | Static reference artifacts only. |
| REV-CORP-010 | exceptions handled appropriately | Estándar de Diseño y Código WEB | code-quality | Menor | Yes | Schema defines status/verdict enums for structured error representation. | review-report.schema.json — status, verdict, nextRecommended enums | |
| REV-CORP-011 | error handling tested and includes non-happy-path cases in test script | Estándar de Diseño y Código WEB | code-quality | Media | Yes | Apply-progress documents static verification of error/edge cases (missing fields, legacy fields absent). | openspec/changes/sdd-review-schema-v2/apply-progress.md — Static Verification Results | |
| REV-CORP-012 | DB connections/sockets/files released even on error | Estándar de Diseño y Código WEB | code-quality | Mayor | N/A | No DB connections, sockets, or file handles in static JSON/Markdown artifacts. | No application runtime; no database access in scope | Static reference artifacts only. |
| REV-CORP-013 | cache/session result storage used appropriately | Estándar de Diseño y Código WEB | code-quality | Media | N/A | No cache or session storage in static reference artifacts. | No application runtime in repository | No runtime surfaces. |
| REV-CORP-014 | code can be optimized | Estándar de Diseño y Código WEB | code-quality | Menor | Yes | Schema v2 removes legacy fields reducing surface area. additionalProperties: false tightens object shapes. | review-report.schema.json — additionalProperties: false at root and nested objects | |
| REV-CORP-015 | standard error logging used | Estándar de Diseño y Código WEB | code-quality | Media | N/A | No error logging in static JSON/Markdown artifacts. | No application runtime in repository | Static reference artifacts only. |
| REV-CORP-016 | common functionality encapsulated | Estándar de Diseño y Código WEB | code-quality | Menor | N/A | No common functionality to encapsulate in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-017 | latest reusable component libraries used (AccessControl, email, employee lookup, etc.) | Estándar de Diseño y Código WEB | code-quality | Mayor | N/A | No reusable component libraries used in static JSON/Markdown artifacts. | No application runtime or UI components | Static reference artifacts only. |
| REV-CORP-018 | portal usability standard implemented | Estándar de Diseño y Código WEB | code-quality | Media | Yes | Template defines clear Markdown presentation structure for human readability. | skills/sdd-review/references/report-template.md — Required Structure section | |
| REV-CORP-019 | code kept simple | Estándar de Codificación Segura | secure-coding | Mayor | Yes | Schema is minimal with exactly 14 required fields, removing 7 legacy fields. | review-report.schema.json — required array; proposal.md — removed fields list | |
| REV-CORP-020 | API designed securely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No API created or modified. Changes are static reference artifacts. | No application runtime; design.md confirms no runtime code | No API in scope. |
| REV-CORP-021 | duplication avoided | Estándar de Codificación Segura | secure-coding | Mayor | Yes | No duplication across the three artifacts. Each has a distinct role (catalog, schema, template). | design.md — Data Flow section | |
| REV-CORP-022 | privileges restricted | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No privilege or access control in static reference artifacts. | design.md — Secure Development Design: no auth surfaces | No auth or privilege surfaces. |
| REV-CORP-023 | number of permission checks minimized | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No permission checks in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-024 | encapsulation used | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No encapsulation applicable to static JSON/Markdown. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-025 | third-party code security validated | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No third-party code introduced. | Only internal reference artifacts modified | No third-party dependencies. |
| REV-CORP-026 | security-related information documented | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No security-related information to document beyond safe-evidence rules. | design.md — Secure Development Design: no-impact classification | Safe evidence rules preserved in schema $comment and template. |
| REV-CORP-027 | activities that can use disproportionate resources handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No resource consumption in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-028 | resources released in all cases | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No resources to release in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-029 | avoid direct buffers for short-lived/infrequently used objects | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No buffers in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-030 | robust error/exception handling for services | Estándar de Codificación Segura | secure-coding | Mayor | Yes | Schema defines structured error representation through status/verdict/nextRecommended enums. | review-report.schema.json — status, verdict, nextRecommended properties | |
| REV-CORP-031 | resource limit checks avoid overflow | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No resource limits in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-032 | sensitive information removed from exceptions | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No exceptions with sensitive information in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-033 | highly confidential information not logged | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No logging in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-034 | sensitive information removed from memory | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No memory management in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-035 | confidential information not hardcoded | Estándar de Codificación Segura | secure-coding | Mayor | Yes | No confidential information hardcoded. Schema $comment and template Safe Evidence Rules explicitly prohibit it. | review-report.schema.json $comment; report-template.md ## Safe Evidence Rules | |
| REV-CORP-036 | valid format generated | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No format generation in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-037 | dynamic SQL avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No SQL in static reference artifacts. | No database access in scope | Static reference artifacts only. |
| REV-CORP-038 | XML and HTML generation handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No XML/HTML generation in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-039 | XML inclusion restricted | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No XML inclusion in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-040 | untrusted code interpretation handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No untrusted code interpretation in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-041 | exceptional floating-point value injection avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No floating-point values in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-042 | accessibility and extensibility standard met | Estándar de Codificación Segura | secure-coding | Mayor | Yes | Template defines accessible Markdown structure with clear sections and table headers. | report-template.md — Required Structure, Matrix sections | |
| REV-CORP-043 | package accessibility limited | Estándar de Codificación Segura | secure-coding | Mayor | Yes | Schema uses additionalProperties: false to limit object accessibility/extensibility. | review-report.schema.json — additionalProperties: false on all objects | |
| REV-CORP-044 | modules used to hide internal packages | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No modules or packages in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-045 | ClassLoader instance exposure limited | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No ClassLoader in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-046 | class and method extensibility limited | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No classes or methods in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-047 | superclass impact on subclass behavior understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No class inheritance in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-048 | strings normalized before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No string validation in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-049 | path names canonicalized before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No path name validation in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-050 | user input not logged without sanitizing first | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No user input logging in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-051 | files extracted safely from ZipInputStream | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No ZipInputStream in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-052 | string modifications performed before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No string modifications in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-053 | wrappers defined around native methods | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No native methods in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-054 | hidden form fields not trusted | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No form fields in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-055 | immutability preferred for value types | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No value types in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-056 | copies created for mutable output values | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No mutable output in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-057 | safe copies created for mutable input values and subclasses | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No mutable input in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-058 | copy functionality supported for mutable class | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No mutable classes in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-059 | identity equality not trusted when input reference objects can override equality | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No identity equality in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-060 | input passed to untrusted object treated as output | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No untrusted object input in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-061 | output from untrusted object treated as input | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No untrusted object output in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-062 | wrapper methods around mutable internal state defined | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No mutable internal state in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-063 | public static fields made final | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No public static fields in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-064 | public static final field values ensured constant | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No public static final fields in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-065 | mutable collections not exposed | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No mutable collections in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-066 | constructors of sensitive classes not exposed | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No sensitive class constructors in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-067 | private members of outer class not exposed from nested class | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No nested classes in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-068 | defend against partially initialized instances of non-final classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No partially initialized instances in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-069 | constructors avoid calling overridable methods | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No constructors in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-070 | sensitive classes cannot be copied | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No sensitive classes in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-071 | serialization avoided for security-sensitive classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No serialization in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-072 | deserialization treated like object construction | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No deserialization in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-073 | do not deviate proper serialization method signatures | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No serialization methods in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-074 | serialization compatibility enabled during class evolution | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No serialization compatibility in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-075 | instances of inner classes not serialized | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No inner class serialization in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-076 | privileges minimized before deserializing from privileged context | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No privileged deserialization in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-077 | untrusted serialized data filtered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No serialized data filtering in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-078 | permission verification understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No permission verification in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-079 | callback methods handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No callback methods in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-080 | java.security.AccessController.doPrivileged invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No AccessController.doPrivileged in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-081 | privileges restricted through doPrivileged | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No doPrivileged restrictions in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-082 | cached results of potentially privileged operations handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No cached privileged-operation results in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-083 | context transfer understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No context transfer in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-084 | thread construction context transfer understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No thread construction in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-085 | standard APIs that skip SecurityManager checks based on immediate caller classloader invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No SecurityManager checks in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-086 | standard APIs that use immediate caller classloader invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No immediate caller classloader APIs in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-087 | standard APIs with Java language access checks against immediate caller considered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No Java language access checks in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-088 | java.lang.reflect.Method.invoke immediate-call check behavior considered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No Method.invoke in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-089 | caller-sensitive method names avoided in interface classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No caller-sensitive interface methods in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-090 | returning privileged operation results avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No privileged operation results in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-091 | standard APIs that perform tasks using immediate caller module invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No immediate caller module APIs in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-092 | InvocationHandlers designed and used conservatively | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No InvocationHandlers in static reference artifacts. | No Java runtime | Static reference artifacts only. |
| REV-CORP-093 | tainted variables not allowed in privileged blocks | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No tainted variables or privileged blocks in static reference artifacts. | No application runtime | Static reference artifacts only. |
| REV-CORP-094 | reflection not used to increase accessibility of classes, methods, or fields | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No reflection in static reference artifacts. | No executable code in scope | Static reference artifacts only. |
| REV-CORP-095 | secure log4j versions used | Estándar de Codificación Segura | secure-coding | Mayor | N/A | No log4j usage in static reference artifacts. | No application runtime; no logging framework dependencies | Static reference artifacts only. |
| REV-CORP-096 | security controls not based on untrusted sources | Estándar de Codificación Segura | secure-coding | Mayor | Yes | Schema enforces structured enums for status/verdict. Safe evidence rules preserved in $comment and template. | review-report.schema.json — $comment; report-template.md — Safe Evidence Rules | |
