# Design: Install Package Builder

## Technical Approach

The skill follows the same LLM-driven agent pattern as `sdd-operational-doc` and `sdd-technical-doc`: a `SKILL.md` file instructs the LLM what to do, and the LLM executes sequential bash tool calls to accomplish each pipeline stage. There is no runtime process, daemon, or orchestrator — each bash invocation is independent and stateless. The pipeline stages execute in strict order: preflight → clone → detect → build → classify → assemble → readme → zip → report.

## Architecture Decisions

### Decision: Skill-as-Instructions Pattern (No Runtime Code)

**Choice**: Encode all logic in SKILL.md as LLM instructions; the agent interprets and executes bash commands directly.
**Alternatives considered**: CLI script (bash/Python), standalone binary.
**Rationale**: Matches existing repo conventions (`sdd-operational-doc`, `sdd-technical-doc`). Zero runtime dependencies beyond the build tools themselves. The LLM handles branching logic (stack detection, SQL classification) without requiring a parser.

### Decision: YAML for Environment Data

**Choice**: `environments.yaml` instead of JSON.
**Alternatives considered**: JSON, embedded in SKILL.md, external fetch.
**Rationale**: YAML is more readable for the 24-environment dataset with nested containers. The LLM can parse YAML inline without tooling. Proposal mentioned `.json` but YAML better serves human editability and the spec's maintainability goal.

### Decision: Template-Based README Generation

**Choice**: `readme-template.txt` with `{{MARKER}}` placeholders filled by the LLM.
**Alternatives considered**: Programmatic generation from scratch, Mustache/Handlebars engine.
**Rationale**: The LLM performs string substitution natively. A plain-text template is human-editable, versionable, and requires no template engine dependency.

### Decision: SQL Classification as Ordered Rules in SKILL.md

**Choice**: Encode classification as a precedence-ordered decision table in SKILL.md instructions.
**Alternatives considered**: External classifier script, regex config file.
**Rationale**: Rules are simple string matches on filenames with one content-inspection fallback. Encoding in SKILL.md keeps the skill self-contained and matches the zero-runtime-code pattern.

## Data Flow

```
User Inputs (changeNumber, repos[], outputPath, ...)
       │
       ▼
┌─────────────────┐
│  1. Preflight   │── verify git, mvn, npm, ng, gradle
└────────┬────────┘
         ▼
┌─────────────────┐
│  2. Clone/Pull  │── git clone --depth 1 (or git pull) per repo
└────────┬────────┘
         ▼
┌─────────────────┐
│  3. Detect      │── check pom.xml, angular.json, build.gradle, package.json, *.sql
└────────┬────────┘
         ▼
┌─────────────────┐
│  4. Build       │── mvn / npm+ng / npm+build / gradlew / (none for oracle-db)
└────────┬────────┘
         ▼
┌─────────────────┐
│  5. Classify    │── SQL files → Packages/Procedures/Types/Tables/Data/Unclassified
└────────┬────────┘
         ▼
┌─────────────────┐
│  6. Assemble    │── copy artifacts to Install/AppDeploy/, DBObjects/, ShellScripts/
└────────┬────────┘
         ▼
┌─────────────────┐
│  7. README      │── fill readme-template.txt → Install/Instalacion/README.txt
└────────┬────────┘
         ▼
┌─────────────────┐
│  8. Mirror      │── create Rollback/ structure (empty)
└────────┬────────┘
         ▼
┌─────────────────┐
│  9. Zip+Report  │── zip package, print summary
└────────┘
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `src/skills/install-package-builder/SKILL.md` | Create | Main skill instructions — pipeline logic, classification rules, hard rules |
| `src/skills/install-package-builder/assets/readme-template.txt` | Create | README.txt template with {{MARKERS}} |
| `src/skills/install-package-builder/assets/environments.yaml` | Create | 24 OWEXX environments (name, host, containers) |
| `src/skills/install-package-builder/references/sql-classification.md` | Create | SQL classification rule reference for human readers |
| `agents/sdd/install-package-builder.md` | Create | Adapter prompt (repo copy) |
| `C:/Users/leo3n/.config/opencode/prompts/sdd/install-package-builder.md` | Create | Adapter prompt (actual loaded by opencode) |
| `C:/Users/leo3n/.config/opencode/opencode.json` | Modify | Add agent entry for `install-package-builder` |
| `AGENTS.md` | Modify | Document the new skill as a manual utility |

## Interfaces / Contracts

### Input Contract (accepted by SKILL.md)

```yaml
changeNumber: "CHG0086767"        # required
version: "1"                      # required, default "1"
outputPath: "C:/output"           # required, absolute path
author: "Leo Nava"                # required
repos:                            # required, at least one
  - url: "https://github.com/org/repo.git"
    branch: "main"
    type: "java"                  # java|angular|react|kotlin|react-native|oracle-db
userStories:                      # required (may be empty list)
  - "US-12345"
corrections:                      # optional
  - "INC001234"
```

### Output Contract

```
{outputPath}/{CHG}_VERSION{N}/
├── Install/
│   ├── AppDeploy/              ← .war, .ear, dist/, .apk, .aab
│   ├── DBObjects/
│   │   ├── Packages/           ← PKS/PKB files
│   │   ├── Procedures/         ← _PRC/SP_/PROC files
│   │   ├── Types/              ← _TYPE/TYPE_ files
│   │   ├── Tables/             ← _TAB/_TABLE/CREATE TABLE files
│   │   ├── Data/               ← execution.sql, INSERT/UPDATE/MERGE files
│   │   └── Unclassified/       ← fallback
│   ├── ShellScripts/           ← .sh files
│   └── Instalacion/
│       └── README.txt          ← generated from template
├── Rollback/                   ← mirror of Install/ structure, empty
└── (zip: {CHG}_VERSION{N}.zip)
```

## Operational Considerations

No aplica. This skill produces local filesystem artifacts only. It has no runtime deployment, no servers, no monitoring, no scheduled jobs. The user invokes it ad-hoc and inspects the output manually.

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Manual | Full pipeline execution with real repos | Invoke the skill with a test change number and verify output structure |
| Manual | SQL classification edge cases | Provide SQL files hitting each rule + unclassified fallback |
| Manual | Partial failure handling | Provide a repo with invalid URL and verify continuation |
| Manual | Pre-flight failure | Remove `mvn` from PATH temporarily, verify halt |

No automated test infrastructure applies — this is an LLM skill, not executable code. Verification is manual invocation with known inputs and output inspection.

## Migration / Rollout

No migration required. The skill is additive: new files, new agent entry. No existing behavior is modified.

## Open Questions

None. All technical decisions are resolved from the proposal and spec.

## Secure Development Design

### Classification and Changed Surface

**Classification: No security impact.**

Changed artifacts: new SKILL.md, new asset files (readme-template.txt, environments.yaml), new adapter prompt, modified opencode.json (agent registration), modified AGENTS.md.

**Touched runtime surfaces**: The skill instructs the LLM to execute `git clone`, `mvn`, `npm`, `ng`, `gradle`, file copy, and zip commands on the local host. All operations are local filesystem writes to a user-specified `outputPath`.

**Untouched runtime surfaces**: No authentication system, no session management, no database access, no API endpoints, no user data processing, no logging infrastructure, no configuration/secrets stores, no network services.

**Why no security category applies**:
- **SEC-AUTH / SEC-SESSION / SEC-ACCESS**: No authentication, session, or authorization logic is introduced. Git credential delegation uses existing OS credential helpers — the skill neither accepts nor stores credentials.
- **SEC-DATA**: No PII, PAN, or confidential data is processed. Environment hostnames in `environments.yaml` are internal infrastructure names already present in existing README artifacts distributed to operations teams.
- **SEC-SECRET**: No secrets are stored, logged, or embedded. The skill explicitly prohibits credential input (NFR-002).
- **SEC-LOG**: No logging infrastructure is created or modified. The summary report contains only file paths, success/fail status, and folder counts.
- **SEC-INPUT**: No user-facing input validation surface — inputs come from the invoking user via the LLM conversation, not from untrusted external sources.

Omitted categories are reviewable omissions for `review-security-report.json`.

### Exception and Evidence Policy

No exceptions are planned. No security controls are applicable to this change. Safe-evidence rules are satisfied trivially: no secrets, credentials, PII, or confidential values exist in any artifact produced by this design.
