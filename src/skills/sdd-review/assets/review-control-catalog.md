# SDD Review Control Catalog

> **Derived/reference-only view**: `skills/sdd-review/assets/review-control-catalog.json` is the canonical source of truth for the stable 96-control corporate code-review checklist. This Markdown file is a human-readable reference view only.
>
> If this Markdown differs from `review-control-catalog.json`, JSON wins. Treat the mismatch as a derived-view repair issue: repair or regenerate this Markdown view from the JSON catalog instead of changing control semantics here.

The matrix in `review-report.md` uses the JSON control `id` as its `Item` value. The `Source Item` column preserves one-to-one mapping to the user's checklist items 1 through 96 and must not be added as an extra report matrix column.

## Manual Catalog Evidence

| Evidence | Result |
| --- | --- |
| Catalog source | Canonical JSON: `skills/sdd-review/assets/review-control-catalog.json`, based on user-provided 96-item corporate code-review checklist source items `1` through `96`. |
| Total controls | 96 rows. |
| Unique Item IDs | 96 unique IDs: `REV-CORP-001..REV-CORP-096`. |
| Source mapping | Each `REV-CORP-NNN` maps one-to-one to source item `N`. |
| Default reviewer | Every control uses `sdd-review`. |
| Default complies vocabulary | Every control uses `Yes`, `No`, or `N/A` to match the review report contract. |
| Report compatibility | Catalog fields map to the required report matrix columns without adding extra report columns. |
| Security authority | `Standard` may cite secure-coding sources, but this catalog does not replace security applicability/design authority. |
| Conflict policy | JSON wins; repair this Markdown view when it drifts. |

## Catalog

| Item ID | Source Item | Artifact/Deliverable | Requirement | Reviewer | Standard | Severity | Default Complies | Evidence Hint | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| REV-CORP-001 | 1 | Estándares de Nomenclatura | naming standards | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Evidencia 1 | Not provided |
| REV-CORP-002 | 2 | Folder structure | folder structure per platform standard | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Evidencia 5 | Not provided |
| REV-CORP-003 | 3 | Service names | service names comply with area naming standards | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Evidencia 8 | Not provided |
| REV-CORP-004 | 4 | Developer testing | developer tested code/unit use cases | sdd-review | Estándar de Diseño y Código WEB | Mayor | Yes | Evidencia 6 | Not provided |
| REV-CORP-005 | 5 | Variable declarations | variable declarations adequately commented; descriptive names may make comments unnecessary | sdd-review | Estándar de Diseño y Código WEB | Mayor | Yes | Evidencia 9 | Los nombres de las variables son muy descriptivas, no se ocupan comentar |
| REV-CORP-006 | 6 | Functions, methods, classes | functions, methods, classes documented | sdd-review | Estándar de Diseño y Código WEB | Mayor | Yes | Not provided | Not provided |
| REV-CORP-007 | 7 | Complex algorithms and optimizations | complex algorithms and optimizations commented | sdd-review | Estándar de Diseño y Código WEB | Media | N/A | Not provided | Not provided |
| REV-CORP-008 | 8 | Commented-out code | commented-out code has explanation; dead code removed | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A | Not provided | Not provided |
| REV-CORP-009 | 9 | Resources and memory | resources and memory released on all error paths | sdd-review | Estándar de Diseño y Código WEB | Media | N/A | Not provided | Not provided |
| REV-CORP-010 | 10 | Exceptions | exceptions handled appropriately | sdd-review | Estándar de Diseño y Código WEB | Menor | Yes | Evidencia 2 | common exception handler |
| REV-CORP-011 | 11 | Error handling tests | error handling tested and includes non-happy-path cases in test script | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Evidencia 7 | Not provided |
| REV-CORP-012 | 12 | DB connections/sockets/files | DB connections/sockets/files released even on error | sdd-review | Estándar de Diseño y Código WEB | Mayor | Yes | Not provided | note Spring transactions, only SP modified |
| REV-CORP-013 | 13 | Cache/session result storage | cache/session result storage used appropriately | sdd-review | Estándar de Diseño y Código WEB | Media | N/A | Not provided | Not provided |
| REV-CORP-014 | 14 | Code optimization | code can be optimized | sdd-review | Estándar de Diseño y Código WEB | Menor | Yes | Not provided | Not provided |
| REV-CORP-015 | 15 | Error logging | standard error logging used | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Not provided | Not provided |
| REV-CORP-016 | 16 | Common functionality | common functionality encapsulated | sdd-review | Estándar de Diseño y Código WEB | Menor | N/A | Not provided | Not provided |
| REV-CORP-017 | 17 | Reusable component libraries | latest reusable component libraries used (AccessControl, email, employee lookup, etc.) | sdd-review | Estándar de Diseño y Código WEB | Mayor | N/A | Not provided | Not provided |
| REV-CORP-018 | 18 | Portal usability standard | portal usability standard implemented | sdd-review | Estándar de Diseño y Código WEB | Media | Yes | Not provided | Not provided |
| REV-CORP-019 | 19 | Code simplicity | code kept simple | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | Not provided |
| REV-CORP-020 | 20 | API design | API designed securely | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | no new API created |
| REV-CORP-021 | 21 | Duplication | duplication avoided | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | Not provided |
| REV-CORP-022 | 22 | Privileges | privileges restricted | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Evidencia 4 | existing role privileges |
| REV-CORP-023 | 23 | Permission checks | number of permission checks minimized | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-024 | 24 | Encapsulation | encapsulation used | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-025 | 25 | Third-party code | third-party code security validated | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-026 | 26 | Security-related information | security-related information documented | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-027 | 27 | Disproportionate resources | activities that can use disproportionate resources handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-028 | 28 | Resources | resources released in all cases | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-029 | 29 | Direct buffers | avoid direct buffers for short-lived/infrequently used objects | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-030 | 30 | Service error/exception handling | robust error/exception handling for services | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Evidencia 2 | common exception handler |
| REV-CORP-031 | 31 | Resource limit checks | resource limit checks avoid overflow | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-032 | 32 | Exception content | sensitive information removed from exceptions | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-033 | 33 | Logging | highly confidential information not logged | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-034 | 34 | Memory | sensitive information removed from memory | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-035 | 35 | Hardcoded information | confidential information not hardcoded | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-036 | 36 | Format generation | valid format generated | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-037 | 37 | Dynamic SQL | dynamic SQL avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-038 | 38 | XML and HTML generation | XML and HTML generation handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-039 | 39 | XML inclusion | XML inclusion restricted | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-040 | 40 | Untrusted code interpretation | untrusted code interpretation handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-041 | 41 | Floating-point values | exceptional floating-point value injection avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-042 | 42 | Accessibility and extensibility | accessibility and extensibility standard met | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | Not provided |
| REV-CORP-043 | 43 | Package accessibility | package accessibility limited | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | Not provided |
| REV-CORP-044 | 44 | Modules | modules used to hide internal packages | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-045 | 45 | ClassLoader exposure | ClassLoader instance exposure limited | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-046 | 46 | Class and method extensibility | class and method extensibility limited | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-047 | 47 | Superclass/subclass behavior | superclass impact on subclass behavior understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-048 | 48 | Strings | strings normalized before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-049 | 49 | Path names | path names canonicalized before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-050 | 50 | User input logging | user input not logged without sanitizing first | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-051 | 51 | ZipInputStream extraction | files extracted safely from ZipInputStream | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-052 | 52 | String modifications | string modifications performed before validation | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-053 | 53 | Native methods | wrappers defined around native methods | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-054 | 54 | Hidden form fields | hidden form fields not trusted | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-055 | 55 | Value types | immutability preferred for value types | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-056 | 56 | Mutable output values | copies created for mutable output values | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-057 | 57 | Mutable input values and subclasses | safe copies created for mutable input values and subclasses | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-058 | 58 | Mutable class copy functionality | copy functionality supported for mutable class | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-059 | 59 | Identity equality | identity equality not trusted when input reference objects can override equality | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-060 | 60 | Input to untrusted object | input passed to untrusted object treated as output | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-061 | 61 | Output from untrusted object | output from untrusted object treated as input | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-062 | 62 | Mutable internal state | wrapper methods around mutable internal state defined | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-063 | 63 | Public static fields | public static fields made final | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-064 | 64 | Public static final fields | public static final field values ensured constant | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-065 | 65 | Mutable collections | mutable collections not exposed | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-066 | 66 | Sensitive class constructors | constructors of sensitive classes not exposed | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-067 | 67 | Nested class exposure | private members of outer class not exposed from nested class | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-068 | 68 | Partially initialized instances | defend against partially initialized instances of non-final classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-069 | 69 | Constructors | constructors avoid calling overridable methods | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-070 | 70 | Sensitive class copying | sensitive classes cannot be copied | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-071 | 71 | Security-sensitive serialization | serialization avoided for security-sensitive classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-072 | 72 | Deserialization | deserialization treated like object construction | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-073 | 73 | Serialization method signatures | do not deviate proper serialization method signatures | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-074 | 74 | Serialization compatibility | serialization compatibility enabled during class evolution | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-075 | 75 | Inner class serialization | instances of inner classes not serialized | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-076 | 76 | Privileged deserialization | privileges minimized before deserializing from privileged context | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-077 | 77 | Serialized data filtering | untrusted serialized data filtered | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-078 | 78 | Permission verification | permission verification understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-079 | 79 | Callback methods | callback methods handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-080 | 80 | AccessController.doPrivileged | java.security.AccessController.doPrivileged invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-081 | 81 | doPrivileged restrictions | privileges restricted through doPrivileged | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-082 | 82 | Cached privileged-operation results | cached results of potentially privileged operations handled carefully | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-083 | 83 | Context transfer | context transfer understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-084 | 84 | Thread construction context transfer | thread construction context transfer understood | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-085 | 85 | Caller classloader SecurityManager checks | standard APIs that skip SecurityManager checks based on immediate caller classloader invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-086 | 86 | Immediate caller classloader APIs | standard APIs that use immediate caller classloader invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-087 | 87 | Immediate caller Java access checks | standard APIs with Java language access checks against immediate caller considered | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-088 | 88 | Method.invoke immediate-call checks | java.lang.reflect.Method.invoke immediate-call check behavior considered | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-089 | 89 | Caller-sensitive interface methods | caller-sensitive method names avoided in interface classes | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-090 | 90 | Privileged operation results | returning privileged operation results avoided | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-091 | 91 | Immediate caller module APIs | standard APIs that perform tasks using immediate caller module invoked safely | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-092 | 92 | InvocationHandlers | InvocationHandlers designed and used conservatively | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-093 | 93 | Tainted variables | tainted variables not allowed in privileged blocks | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-094 | 94 | Reflection accessibility | reflection not used to increase accessibility of classes, methods, or fields | sdd-review | Estándar de Codificación Segura | Mayor | N/A | Not provided | Not provided |
| REV-CORP-095 | 95 | log4j versions | secure log4j versions used | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | No se modificó, se usa la misma versión ya aprobada. |
| REV-CORP-096 | 96 | Security controls | security controls not based on untrusted sources | sdd-review | Estándar de Codificación Segura | Mayor | Yes | Not provided | Not provided |
