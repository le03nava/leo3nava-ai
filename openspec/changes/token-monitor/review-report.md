# Review Report: token-monitor

## Verdict

| Field | Value |
| --- | --- |
| Change | `token-monitor` |
| Status | success |
| Verdict | PASS WITH WARNINGS |
| Next recommendation | verify |
| JSON authority | `openspec/changes/token-monitor/review-report.json` |
| Markdown authority | derived compatibility view |

## Totals

| Metric | Value |
| --- | --- |
| Total controls (96) | 96 |
| Passing | 18 |
| Failing | 0 |
| N/A | 78 |
| Blocking | 0 |
| Non-blocking | 2 |

## Unavailable Tooling

- Coverage runner (pytest-cov not wired to CI)
- Linter (no ruff/flake8 configured at repo level)
- Type checker (no mypy/pyright configured)
- Formatter (no black/ruff-format configured)

## Artifact Metadata

| Check | Result |
| --- | --- |
| Canonical JSON ref | `openspec/changes/token-monitor/review-report.json` |
| Derived Markdown ref | `openspec/changes/token-monitor/review-report.md` |
| JSON persisted / read back | true / true |
| Markdown generated / persisted / read back | true / true / true |
| JSON/Markdown parity | ok |
| JSON authority | canonical |
| Markdown authority | derived |

## Recommendation

- Next: `verify`
- All 22 mandatory automated test cases pass (24 tests total including 2 bonus integration-level tests). No blocking failures found. Two non-blocking warnings should be addressed before archive.

## Findings

### WARN-001 — Task 3.1 conftest.py unchecked (WARNING)

**Control**: REV-CORP-004 (developer testing)
**Description**: `tasks.md` shows task 3.1 (create `conftest.py` with shared fixtures) as unchecked `[ ]`. The test-design mandates fixtures be in `conftest.py`, but they are inlined in the test module. Tests pass regardless.
**Recommendation**: Create `conftest.py` with shared fixtures or mark task 3.1 as intentionally superseded.

### WARN-002 — Module-level side effect in token_monitor.py (WARNING)

**Control**: REV-CORP-014 (code optimization)
**Description**: Lines 247–250 of `token_monitor.py` execute `sys.argv` parsing at module import time to populate the `addons` list. This couples import behavior to runtime state. Tests work because `sys.argv[0]` is not `mitmdump`, but it's fragile.
**Recommendation**: Move addon creation into a guarded function or use mitmproxy's recommended addon loading pattern.

## Static Review Sign-offs

| Check | Result |
| --- | --- |
| TC-025: Parameterized SQL only | PASS — `storage.py` uses `?` placeholders exclusively; no f-string/%-format SQL |
| TC-026: No sensitive data logged/persisted | PASS — No request bodies, prompts, completions, or API keys stored |
| TC-027: pathlib in storage.py | PASS — All path ops use `pathlib.Path` |
| TC-028: pathlib in export.py | PASS — All path ops use `pathlib.Path` |

## Matrix

| Item | Requirement | Standard | Category | Severity | Complies | Finding | Evidence Location | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| REV-CORP-001 | naming standards | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | token-monitor/storage.py, token-monitor/token_monitor.py, token-monitor/export.py | snake_case naming throughout; PEP 8 compliant |
| REV-CORP-002 | folder structure per platform standard | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | token-monitor/ folder structure | Flat module layout with tests/ subdirectory per design |
| REV-CORP-003 | service names comply with area naming standards | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | token-monitor/storage.py, token-monitor/token_monitor.py, token-monitor/export.py | Module names match design file table |
| REV-CORP-004 | developer tested code/unit use cases | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | non-blocking: task 3.1 conftest.py unchecked | 24 tests passing (pytest output); tasks.md task 3.1 unchecked | WARN-001: Fixtures inlined in test module instead of conftest.py |
| REV-CORP-005 | variable declarations adequately commented; descriptive names may make comments unnecessary | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | | token-monitor/storage.py, token-monitor/token_monitor.py | Descriptive variable names; docstrings on public functions |
| REV-CORP-006 | functions, methods, classes documented | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | | token-monitor/storage.py, token-monitor/token_monitor.py, token-monitor/export.py | All public functions have docstrings |
| REV-CORP-007 | complex algorithms and optimizations commented | Estándar de Diseño y Código WEB | code-quality | Media | N/A | | No complex algorithms in implementation | Simple JSON parsing and SQLite inserts only |
| REV-CORP-008 | commented-out code has explanation; dead code removed | Estándar de Diseño y Código WEB | code-quality | Menor | N/A | | No commented-out code in any source file | Clean codebase |
| REV-CORP-009 | resources and memory released on all error paths | Estándar de Diseño y Código WEB | code-quality | Media | N/A | | No manual memory/resource allocation beyond SQLite connection | SQLite connection closed in done(); Python GC handles rest |
| REV-CORP-010 | exceptions handled appropriately | Estándar de Diseño y Código WEB | code-quality | Menor | Yes | | token-monitor/storage.py lines 79-116; token-monitor/token_monitor.py lines 213-216 | All exceptions caught, logged to stderr, proxy continues |
| REV-CORP-011 | error handling tested and includes non-happy-path cases in test script | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | tests/test_token_monitor.py TC-003, TC-005, TC-019 | Error paths tested: malformed JSON, missing usage, DB write failure |
| REV-CORP-012 | DB connections/sockets/files released even on error | Estándar de Diseño y Código WEB | code-quality | Mayor | Yes | | token-monitor/storage.py insert_event try/finally implicit via commit; token_monitor.py done() closes conn | SQLite connection managed with explicit close |
| REV-CORP-013 | cache/session result storage used appropriately | Estándar de Diseño y Código WEB | code-quality | Media | N/A | | No cache or session storage used | Local tool, no caching layer |
| REV-CORP-014 | code can be optimized | Estándar de Diseño y Código WEB | code-quality | Menor | Yes | non-blocking: module-level side effect | token-monitor/token_monitor.py lines 247-250 | WARN-002: sys.argv parsing at import time |
| REV-CORP-015 | standard error logging used | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | token-monitor/token_monitor.py LOGGER; token-monitor/storage.py print to stderr | Consistent stderr logging for errors/warnings |
| REV-CORP-016 | common functionality encapsulated | Estándar de Diseño y Código WEB | code-quality | Menor | N/A | | No duplicated common functionality across modules | Each module has distinct responsibility |
| REV-CORP-017 | latest reusable component libraries used (AccessControl, email, employee lookup, etc.) | Estándar de Diseño y Código WEB | code-quality | Mayor | N/A | | No enterprise component libraries applicable to local CLI tool | Standalone developer tool |
| REV-CORP-018 | portal usability standard implemented | Estándar de Diseño y Código WEB | code-quality | Media | Yes | | CLI uses argparse with --help; README documents usage | Standard Python CLI usability |
| REV-CORP-019 | code kept simple | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | token-monitor/storage.py, token-monitor/token_monitor.py, token-monitor/export.py | Simple, focused modules; single responsibility per file |
| REV-CORP-020 | API designed securely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No new API created (local tool, no HTTP server) | Proxy intercepts only, does not expose an API |
| REV-CORP-021 | duplication avoided | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | No duplicated logic between modules | Extraction, storage, export cleanly separated |
| REV-CORP-022 | privileges restricted | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | token-monitor/storage.py uses user-owned SQLite; localhost only | Single-user local tool, no privilege escalation |
| REV-CORP-023 | number of permission checks minimized | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No permission/auth system in local tool | Out of scope per proposal |
| REV-CORP-024 | encapsulation used | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | Python modules with clear public/private separation via underscore prefix | No Java-style encapsulation needed; Python conventions followed |
| REV-CORP-025 | third-party code security validated | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | mitmproxy pinned to 10.4.2 in requirements.txt; well-known maintained project | No custom third-party security validation needed for mitmproxy |
| REV-CORP-026 | security-related information documented | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No security-related information to document beyond design.md#secure-development-design | Classification: no-impact |
| REV-CORP-027 | activities that can use disproportionate resources handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No unbounded resource consumption; SSE buffers released per-flow (TC-006) | Buffer cleanup verified by test |
| REV-CORP-028 | resources released in all cases | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | Python with GC; SQLite connection explicitly closed | No manual resource management beyond conn.close() |
| REV-CORP-029 | avoid direct buffers for short-lived/infrequently used objects | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No direct buffer allocation (Python, not Java/C) | Not applicable to Python runtime |
| REV-CORP-030 | robust error/exception handling for services | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | token-monitor/storage.py lines 79-116; token-monitor/token_monitor.py lines 213-216 | All exceptions caught and logged; proxy never crashes |
| REV-CORP-031 | resource limit checks avoid overflow | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No integer overflow risk; Python has arbitrary precision integers | Not applicable to Python |
| REV-CORP-032 | sensitive information removed from exceptions | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | Exceptions logged to stderr only with generic message; no sensitive data exposed | Error messages contain only table name and SQLite error |
| REV-CORP-033 | highly confidential information not logged | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No confidential information logged; only token counts and agent labels (TC-026 static review) | Request bodies never persisted |
| REV-CORP-034 | sensitive information removed from memory | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No sensitive data in memory; only integer counts and string labels | Not applicable |
| REV-CORP-035 | confidential information not hardcoded | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No hardcoded credentials in any source file | Only DEFAULT_DB_PATH hardcoded (non-sensitive) |
| REV-CORP-036 | valid format generated | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No format generation beyond JSON/CSV output (standard library) | Uses json.dumps and csv.DictWriter safely |
| REV-CORP-037 | dynamic SQL avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | token-monitor/storage.py uses parameterized ? placeholders only (TC-025) | No dynamic SQL construction |
| REV-CORP-038 | XML and HTML generation handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No XML or HTML generation in this tool | Not applicable |
| REV-CORP-039 | XML inclusion restricted | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No XML processing | Not applicable |
| REV-CORP-040 | untrusted code interpretation handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No eval/exec or untrusted code interpretation | Not applicable |
| REV-CORP-041 | exceptional floating-point value injection avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No floating-point arithmetic in security context | Not applicable |
| REV-CORP-042 | accessibility and extensibility standard met | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | Modular design; functions are independently testable | Pure functions for extraction, class for addon lifecycle |
| REV-CORP-043 | package accessibility limited | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | Private helpers prefixed with underscore | _header_get, _safe_response_json, _calculate_duration_ms, _fetch_rows, _write_jsonl, _write_csv |
| REV-CORP-044 | modules used to hide internal packages | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Java module system; Python packages with clear boundaries | Not applicable to Python |
| REV-CORP-045 | ClassLoader instance exposure limited | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No ClassLoader (Java-specific) | Not applicable to Python |
| REV-CORP-046 | class and method extensibility limited | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Java-style class extensibility concern | Not applicable to Python |
| REV-CORP-047 | superclass impact on subclass behavior understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No inheritance hierarchy | Not applicable |
| REV-CORP-048 | strings normalized before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No string validation gates | Not applicable |
| REV-CORP-049 | path names canonicalized before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | Path handling uses pathlib which normalizes paths (TC-027, TC-028) | pathlib handles canonicalization |
| REV-CORP-050 | user input not logged without sanitizing first | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No user input logged; only token counts and agent labels (TC-026) | Not applicable |
| REV-CORP-051 | files extracted safely from ZipInputStream | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No ZIP file handling | Not applicable |
| REV-CORP-052 | string modifications performed before validation | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No string modification before validation pattern | Not applicable |
| REV-CORP-053 | wrappers defined around native methods | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No native methods (Python, not Java/C) | Not applicable |
| REV-CORP-054 | hidden form fields not trusted | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No HTML forms | Not applicable |
| REV-CORP-055 | immutability preferred for value types | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Java-style value types | Not applicable to Python |
| REV-CORP-056 | copies created for mutable output values | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No mutable output value security concern in Python tool | Not applicable |
| REV-CORP-057 | safe copies created for mutable input values and subclasses | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No security-sensitive mutable input handling | Not applicable |
| REV-CORP-058 | copy functionality supported for mutable class | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No mutable class copy pattern | Not applicable |
| REV-CORP-059 | identity equality not trusted when input reference objects can override equality | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No identity equality security concern | Not applicable |
| REV-CORP-060 | input passed to untrusted object treated as output | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No untrusted object input/output boundary | Not applicable |
| REV-CORP-061 | output from untrusted object treated as input | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No untrusted object output handling | Not applicable |
| REV-CORP-062 | wrapper methods around mutable internal state defined | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No mutable internal state exposure concern | Not applicable |
| REV-CORP-063 | public static fields made final | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No public static fields (Java-specific) | Not applicable to Python |
| REV-CORP-064 | public static final field values ensured constant | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No public static final fields (Java-specific) | Not applicable to Python |
| REV-CORP-065 | mutable collections not exposed | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No mutable collection exposure security concern | Not applicable |
| REV-CORP-066 | constructors of sensitive classes not exposed | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No sensitive class constructors | Not applicable |
| REV-CORP-067 | private members of outer class not exposed from nested class | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No nested class private member exposure | Not applicable |
| REV-CORP-068 | defend against partially initialized instances of non-final classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No partially initialized instance concern (Python) | Not applicable |
| REV-CORP-069 | constructors avoid calling overridable methods | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No constructor calling overridable methods pattern | Not applicable |
| REV-CORP-070 | sensitive classes cannot be copied | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No sensitive class copy prevention needed | Not applicable |
| REV-CORP-071 | serialization avoided for security-sensitive classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No serialization of security-sensitive classes | Not applicable |
| REV-CORP-072 | deserialization treated like object construction | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No deserialization of untrusted data | Not applicable |
| REV-CORP-073 | do not deviate proper serialization method signatures | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No serialization methods | Not applicable |
| REV-CORP-074 | serialization compatibility enabled during class evolution | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No serialization compatibility concern | Not applicable |
| REV-CORP-075 | instances of inner classes not serialized | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No inner class serialization | Not applicable |
| REV-CORP-076 | privileges minimized before deserializing from privileged context | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No privileged deserialization | Not applicable |
| REV-CORP-077 | untrusted serialized data filtered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No untrusted serialized data filtering | Not applicable |
| REV-CORP-078 | permission verification understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No permission verification system | Not applicable |
| REV-CORP-079 | callback methods handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No security-sensitive callback methods | Not applicable |
| REV-CORP-080 | java.security.AccessController.doPrivileged invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Java AccessController (Java-specific) | Not applicable to Python |
| REV-CORP-081 | privileges restricted through doPrivileged | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No doPrivileged (Java-specific) | Not applicable to Python |
| REV-CORP-082 | cached results of potentially privileged operations handled carefully | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No cached privileged operation results | Not applicable |
| REV-CORP-083 | context transfer understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No security context transfer | Not applicable |
| REV-CORP-084 | thread construction context transfer understood | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No thread construction context transfer | Not applicable |
| REV-CORP-085 | standard APIs that skip SecurityManager checks based on immediate caller classloader invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No SecurityManager (Java-specific) | Not applicable to Python |
| REV-CORP-086 | standard APIs that use immediate caller classloader invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No immediate caller classloader (Java-specific) | Not applicable to Python |
| REV-CORP-087 | standard APIs with Java language access checks against immediate caller considered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Java language access checks | Not applicable to Python |
| REV-CORP-088 | java.lang.reflect.Method.invoke immediate-call check behavior considered | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No Method.invoke (Java-specific) | Not applicable to Python |
| REV-CORP-089 | caller-sensitive method names avoided in interface classes | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No caller-sensitive interface methods (Java-specific) | Not applicable to Python |
| REV-CORP-090 | returning privileged operation results avoided | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No privileged operation results returned | Not applicable |
| REV-CORP-091 | standard APIs that perform tasks using immediate caller module invoked safely | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No immediate caller module APIs (Java-specific) | Not applicable to Python |
| REV-CORP-092 | InvocationHandlers designed and used conservatively | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No InvocationHandlers (Java-specific) | Not applicable to Python |
| REV-CORP-093 | tainted variables not allowed in privileged blocks | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No privileged blocks with tainted variables | Not applicable |
| REV-CORP-094 | reflection not used to increase accessibility of classes, methods, or fields | Estándar de Codificación Segura | secure-coding | Mayor | N/A | | No reflection used to bypass access controls | Not applicable |
| REV-CORP-095 | secure log4j versions used | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | Python project; no log4j dependency | Uses Python logging module only |
| REV-CORP-096 | security controls not based on untrusted sources | Estándar de Codificación Segura | secure-coding | Mayor | Yes | | No security controls based on untrusted sources; tool is local-only | No auth/trust decisions made |
