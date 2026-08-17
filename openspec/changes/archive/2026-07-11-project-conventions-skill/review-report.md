# Review Report: project-conventions-skill

## Verdict

| Field | Value |
| --- | --- |
| Change | `project-conventions-skill` |
| Verdict | PASS |
| Blocking failures | 0 |
| Non-blocking findings | 0 |
| Next recommendation | review-security |

## Blocking Summary

| Item | Severity | Affected Requirement | Evidence Location | Required Follow-up |
| --- | --- | --- | --- | --- |
| None | — | — | — | — |

## Evidence Summary

- Inputs inspected: proposal, spec, design, test-design, tasks, apply-progress, and 7 changed files (SKILL.md, architecture.md, code-style.md, testing.md, review-readiness.md, secure-design.md, AGENTS.md).
- Catalog coverage: 96 unique review controls from canonical `skills/sdd-review/references/review-control-catalog.json` snapshot `sdd-review-control-catalog-2026-07-10-rev-corp-001-096`.
- Security boundary: review cites security guideline/source IDs where applicable; `design.md#secure-development-design`, `skills/sdd-review-security/references/security-guideline-catalog.operational.json`, and canonical `review-security-report.json` remain authoritative for new changes. Derived `review-security-report.md` and `skills/sdd-review-security/references/security-guideline-catalog.md` are human/audit compatibility views.
- Runtime checks: No automated tooling available. strict_tdd: false; no test runner, build, linter, type checker, or formatter configured. Change is documentation-only Markdown files.

## Review Matrix

| Item | Artifact/Deliverable | Requirement | Reviewer | Standard | Severity | Complies | Affected Requirement | Evidence Location | Observations/Comments |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| REV-CORP-001 | Estándares de Nomenclatura | naming standards | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | SPEC-007 | skills/project-conventions/references/code-style.md | English artifact naming convention documented; skill file naming follows established pattern |
| REV-CORP-002 | Folder structure | folder structure per platform standard | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | SPEC-006 | skills/project-conventions/references/architecture.md | Follows skills/{name}/references/ pattern consistent with existing skills |
| REV-CORP-003 | Service names | service names comply with area naming standards | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | SPEC-001 | skills/project-conventions/SKILL.md frontmatter name field | Skill named project-conventions per skill-creator conventions |
| REV-CORP-004 | Developer testing | developer tested code/unit use cases | sdd-review | Estándar de Diseño y Código WEB | Mayor | Yes | SPEC-008 | openspec/changes/project-conventions-skill/test-design.md | Manual verification planned per test-design; strict_tdd: false |
| REV-CORP-005 | Variable declarations | variable declarations adequately commented; descriptive names may make comments unnecessary | sdd-review | Estándar de Diseño y Código WEB | Mayor | N/A |  | Change is documentation-only Markdown; no code variables declared | No executable code in this change; only Markdown skill files |
| REV-CORP-006 | Functions, methods, classes | functions, methods, classes documented | sdd-review | Estándar de Diseño y Código WEB | Mayor | N/A |  | Change is documentation-only Markdown; no functions/methods/classes | No executable code in this change |
| REV-CORP-007 | Complex algorithms and optimizations | complex algorithms and optimizations commented | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no algorithms | No executable code in this change |
| REV-CORP-008 | Commented-out code | commented-out code has explanation; dead code removed | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A |  | Documentation-only change; no source code | No executable code in this change |
| REV-CORP-009 | Resources and memory | resources and memory released on all error paths | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no resource allocation | No executable code in this change |
| REV-CORP-010 | Exceptions | exceptions handled appropriately | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A |  | Documentation-only change; no exception handling | No executable code in this change |
| REV-CORP-011 | Error handling tests | error handling tested and includes non-happy-path cases in test script | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no error handling code to test | No executable code; verification is manual per test-design |
| REV-CORP-012 | DB connections/sockets/files | DB connections/sockets/files released even on error | sdd-review | Estándar de Diseño y Código WEB | Mayor | N/A |  | Documentation-only change; no DB/socket/file handles | No executable code in this change |
| REV-CORP-013 | Cache/session result storage | cache/session result storage used appropriately | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no cache/session usage | No executable code in this change |
| REV-CORP-014 | Code optimization | code can be optimized | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A |  | Documentation-only change; no code to optimize | No executable code in this change |
| REV-CORP-015 | Error logging | standard error logging used | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no logging implementation | No executable code in this change |
| REV-CORP-016 | Common functionality | common functionality encapsulated | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A |  | Documentation-only change; no shared code | No executable code in this change |
| REV-CORP-017 | Reusable component libraries | latest reusable component libraries used | sdd-review | Estándar de Diseño y Código WEB | Mayor | N/A |  | Documentation-only change; no component libraries | No executable code in this change |
| REV-CORP-018 | Portal usability standard | portal usability standard implemented | sdd-review | Estándar de Diseño y Código WEB | Media | N/A |  | Documentation-only change; no UI portal | No UI or portal in this change |
| REV-CORP-019 | Code simplicity | code kept simple | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-002 | skills/project-conventions/SKILL.md | Skill structure is minimal and follows established patterns; documentation is concise |
| REV-CORP-020 | API design | API designed securely | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no API created | No API endpoints in this change |
| REV-CORP-021 | Duplication | duplication avoided | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-002, SPEC-013 | skills/project-conventions/SKILL.md Hard Rules | Hard rules explicitly prohibit duplicating 96-control matrix; references are separate files avoiding content duplication |
| REV-CORP-022 | Privileges | privileges restricted | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-013 | skills/project-conventions/SKILL.md Hard Rules, Activation Contract | Skill is non-user-invocable, supplemental only; cannot override review-security authority |
| REV-CORP-023 | Permission checks | number of permission checks minimized | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no permission checks in code | No executable code in this change |
| REV-CORP-024 | Encapsulation | encapsulation used | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-006 | skills/project-conventions/references/ directory structure | Concerns separated into distinct reference files; SKILL.md encapsulates activation logic |
| REV-CORP-025 | Third-party code | third-party code security validated | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no third-party code introduced | No dependencies added |
| REV-CORP-026 | Security-related information | security-related information documented | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-010 | skills/project-conventions/references/secure-design.md | 9 security surface areas documented with narrative rules including evidence expectations |
| REV-CORP-027 | Disproportionate resources | activities that can use disproportionate resources handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no resource-intensive operations | No executable code in this change |
| REV-CORP-028 | Resources | resources released in all cases | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no resource allocation | No executable code in this change |
| REV-CORP-029 | Direct buffers | avoid direct buffers for short-lived/infrequently used objects | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no buffer usage | No executable code in this change |
| REV-CORP-030 | Service error/exception handling | robust error/exception handling for services | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no service code | No executable code in this change |
| REV-CORP-031 | Resource limit checks | resource limit checks avoid overflow | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no resource limits | No executable code in this change |
| REV-CORP-032 | Exception content | sensitive information removed from exceptions | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no exceptions | No executable code in this change |
| REV-CORP-033 | Logging | highly confidential information not logged | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no logging code | No executable code in this change |
| REV-CORP-034 | Memory | sensitive information removed from memory | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no memory management | No executable code in this change |
| REV-CORP-035 | Hardcoded information | confidential information not hardcoded | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-002, SPEC-010 | skills/project-conventions/SKILL.md Hard Rules; skills/project-conventions/references/secure-design.md section 5 | Hard rule 5 prohibits secrets/tokens/PII in evidence; secure-design.md section 5 has no-exception policy for hardcoded secrets |
| REV-CORP-036 | Format generation | valid format generated | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-007 | skills/project-conventions/SKILL.md frontmatter; skills/project-conventions/references/code-style.md | Valid YAML frontmatter and well-formed Markdown throughout all files |
| REV-CORP-037 | Dynamic SQL | dynamic SQL avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no SQL | No executable code in this change |
| REV-CORP-038 | XML and HTML generation | XML and HTML generation handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no XML/HTML generation | No executable code in this change |
| REV-CORP-039 | XML inclusion | XML inclusion restricted | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no XML processing | No executable code in this change |
| REV-CORP-040 | Untrusted code interpretation | untrusted code interpretation handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no code interpretation | No executable code in this change |
| REV-CORP-041 | Floating-point values | exceptional floating-point value injection avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no floating-point operations | No executable code in this change |
| REV-CORP-042 | Accessibility and extensibility | accessibility and extensibility standard met | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-001, SPEC-004 | skills/project-conventions/SKILL.md Execution Steps; Decision Gates | Skill is extensible via reference files; execution steps allow conditional loading |
| REV-CORP-043 | Package accessibility | package accessibility limited | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-001 | skills/project-conventions/SKILL.md Activation Contract | Skill is non-user-invocable; orchestrator-only access |
| REV-CORP-044 | Modules | modules used to hide internal packages | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no Java/module system | No executable code in this change |
| REV-CORP-045 | ClassLoader exposure | ClassLoader instance exposure limited | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no Java ClassLoader | No executable code in this change |
| REV-CORP-046 | Class and method extensibility | class and method extensibility limited | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no classes | No executable code in this change |
| REV-CORP-047 | Superclass/subclass behavior | superclass impact on subclass behavior understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no inheritance | No executable code in this change |
| REV-CORP-048 | Strings | strings normalized before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no string validation code | No executable code in this change |
| REV-CORP-049 | Path names | path names canonicalized before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no path validation code | No executable code in this change |
| REV-CORP-050 | User input logging | user input not logged without sanitizing first | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no logging code | No executable code in this change |
| REV-CORP-051 | ZipInputStream extraction | files extracted safely from ZipInputStream | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no zip handling | No executable code in this change |
| REV-CORP-052 | String modifications | string modifications performed before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no string modification code | No executable code in this change |
| REV-CORP-053 | Native methods | wrappers defined around native methods | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no native methods | No executable code in this change |
| REV-CORP-054 | Hidden form fields | hidden form fields not trusted | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no form fields | No executable code or UI in this change |
| REV-CORP-055 | Value types | immutability preferred for value types | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no value types | No executable code in this change |
| REV-CORP-056 | Mutable output values | copies created for mutable output values | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no mutable values | No executable code in this change |
| REV-CORP-057 | Mutable input values and subclasses | safe copies created for mutable input values and subclasses | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no mutable inputs | No executable code in this change |
| REV-CORP-058 | Mutable class copy functionality | copy functionality supported for mutable class | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no mutable classes | No executable code in this change |
| REV-CORP-059 | Identity equality | identity equality not trusted when input reference objects can override equality | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no equality checks | No executable code in this change |
| REV-CORP-060 | Input to untrusted object | input passed to untrusted object treated as output | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no object interaction | No executable code in this change |
| REV-CORP-061 | Output from untrusted object | output from untrusted object treated as input | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no untrusted objects | No executable code in this change |
| REV-CORP-062 | Mutable internal state | wrapper methods around mutable internal state defined | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no internal state | No executable code in this change |
| REV-CORP-063 | Public static fields | public static fields made final | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no static fields | No executable code in this change |
| REV-CORP-064 | Public static final fields | public static final field values ensured constant | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no static final fields | No executable code in this change |
| REV-CORP-065 | Mutable collections | mutable collections not exposed | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no collections | No executable code in this change |
| REV-CORP-066 | Sensitive class constructors | constructors of sensitive classes not exposed | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no classes | No executable code in this change |
| REV-CORP-067 | Nested class exposure | private members of outer class not exposed from nested class | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no nested classes | No executable code in this change |
| REV-CORP-068 | Partially initialized instances | defend against partially initialized instances of non-final classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no class instantiation | No executable code in this change |
| REV-CORP-069 | Constructors | constructors avoid calling overridable methods | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no constructors | No executable code in this change |
| REV-CORP-070 | Sensitive class copying | sensitive classes cannot be copied | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no sensitive classes | No executable code in this change |
| REV-CORP-071 | Security-sensitive serialization | serialization avoided for security-sensitive classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no serialization | No executable code in this change |
| REV-CORP-072 | Deserialization | deserialization treated like object construction | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no deserialization | No executable code in this change |
| REV-CORP-073 | Serialization method signatures | do not deviate proper serialization method signatures | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no serialization methods | No executable code in this change |
| REV-CORP-074 | Serialization compatibility | serialization compatibility enabled during class evolution | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no serialization | No executable code in this change |
| REV-CORP-075 | Inner class serialization | instances of inner classes not serialized | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no inner classes | No executable code in this change |
| REV-CORP-076 | Privileged deserialization | privileges minimized before deserializing from privileged context | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no deserialization | No executable code in this change |
| REV-CORP-077 | Serialized data filtering | untrusted serialized data filtered | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no serialized data | No executable code in this change |
| REV-CORP-078 | Permission verification | permission verification understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no permission verification code | No executable code in this change |
| REV-CORP-079 | Callback methods | callback methods handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no callbacks | No executable code in this change |
| REV-CORP-080 | AccessController.doPrivileged | java.security.AccessController.doPrivileged invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no Java SecurityManager | No executable code in this change |
| REV-CORP-081 | doPrivileged restrictions | privileges restricted through doPrivileged | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no doPrivileged | No executable code in this change |
| REV-CORP-082 | Cached privileged-operation results | cached results of potentially privileged operations handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no privileged operations | No executable code in this change |
| REV-CORP-083 | Context transfer | context transfer understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no context transfer | No executable code in this change |
| REV-CORP-084 | Thread construction context transfer | thread construction context transfer understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no threading | No executable code in this change |
| REV-CORP-085 | Caller classloader SecurityManager checks | standard APIs that skip SecurityManager checks based on immediate caller classloader invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no SecurityManager | No executable code in this change |
| REV-CORP-086 | Immediate caller classloader APIs | standard APIs that use immediate caller classloader invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no classloader APIs | No executable code in this change |
| REV-CORP-087 | Immediate caller Java access checks | standard APIs with Java language access checks against immediate caller considered | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no Java access checks | No executable code in this change |
| REV-CORP-088 | Method.invoke immediate-call checks | java.lang.reflect.Method.invoke immediate-call check behavior considered | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no reflection | No executable code in this change |
| REV-CORP-089 | Caller-sensitive interface methods | caller-sensitive method names avoided in interface classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no interfaces | No executable code in this change |
| REV-CORP-090 | Privileged operation results | returning privileged operation results avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no privileged results | No executable code in this change |
| REV-CORP-091 | Immediate caller module APIs | standard APIs that perform tasks using immediate caller module invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no module APIs | No executable code in this change |
| REV-CORP-092 | InvocationHandlers | InvocationHandlers designed and used conservatively | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no InvocationHandlers | No executable code in this change |
| REV-CORP-093 | Tainted variables | tainted variables not allowed in privileged blocks | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no privileged blocks | No executable code in this change |
| REV-CORP-094 | Reflection accessibility | reflection not used to increase accessibility of classes, methods, or fields | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no reflection | No executable code in this change |
| REV-CORP-095 | log4j versions | secure log4j versions used | sdd-review | Estándar de Codificación Segura | Mayor | N/A |  | Documentation-only change; no log4j dependency | No dependencies added in this change |
| REV-CORP-096 | Security controls | security controls not based on untrusted sources | sdd-review | Estándar de Codificación Segura | Mayor | Yes | SPEC-013 | skills/project-conventions/SKILL.md Hard Rules; skills/project-conventions/references/architecture.md Boundary Rule | Security validation authority explicitly delegated to sdd-review-security; project-conventions cannot override |

## Operational Evidence Summary

- Summary: No operational evidence applicable. This change creates documentation-only skill files with no runtime, deployment, or monitoring surface.
- Items: None

## Changed-File / Security Handoff

- Summary: Changed files are Markdown documentation only. secure-design.md provides narrative guidance for downstream phases but contains no executable code, secrets, or sensitive data. No security-relevant runtime behavior introduced. Handoff to review-security for validation of secure-design boundary compliance and narrative rule completeness.
- Items:
  - skills/project-conventions/SKILL.md — skill activation contract, hard rules prohibit matrix/scoring duplication
  - skills/project-conventions/references/secure-design.md — narrative secure-design guidance, 9 areas, no SEC-* matrices
  - AGENTS.md — registers skill as supplemental only

## Matrix Validation

| Check | Result |
| --- | --- |
| Catalog controls | 96/96 |
| Unique Item IDs | true |
| Complete Item ID sequence | true |
| Complete source item sequence | true |
| Allowed Complies vocabulary | Yes, No, N/A |
| N/A rows have evidence and rationale | true |
| Blocking counts match verdict | true |
| Derived Markdown generated | true |
| Derived Markdown read back | true |
| Derived Markdown parity | true |
| Safe evidence checked | true |

## Recommendation

- Next recommendation: `review-security`
- Follow-up: No blocking failures. Proceed to security review for validation of secure-design boundary compliance and narrative completeness in secure-design.md.
