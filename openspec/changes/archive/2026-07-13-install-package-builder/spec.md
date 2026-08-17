# Install Package Builder — Specification

## Overview

This change introduces `install-package-builder`, a standalone skill that automates the end-to-end assembly of a versioned installation package for a change order. Given a change number, a version, an output path, and a list of repositories with metadata, the skill performs pre-flight checks, clones or updates repos, detects the tech stack, executes the appropriate build, classifies artifacts into the canonical folder hierarchy, generates a README.txt from a template, mirrors the structure into a Rollback folder, and produces a final zip. The skill is user-invocable and is NOT an SDD DAG phase.

---

## Functional Requirements

### Requirement: FR-001 — Inputs

The skill MUST accept the following inputs:

| Input | Type | Required | Description |
|-------|------|----------|-------------|
| `changeNumber` | string | Yes | Change order identifier (e.g. CHG0086767) |
| `version` | string | Yes | Package version number (e.g. 1) |
| `outputPath` | string | Yes | Directory where the package folder will be created |
| `repos` | list | Yes | List of repo objects: `url`, `branch`, `type` |
| `author` | string | Yes | Person responsible for the package |
| `userStories` | list | Yes | User story identifiers included in the change |
| `corrections` | list | No | Correction identifiers included in the change |

#### Scenario: Valid inputs provided

- GIVEN all required inputs are present and well-formed
- WHEN the skill is invoked
- THEN the pipeline proceeds to pre-flight checks without error

#### Scenario: Required input missing

- GIVEN `changeNumber` or `repos` is absent
- WHEN the skill is invoked
- THEN the skill MUST halt immediately with a clear message naming the missing field

---

### Requirement: FR-002 — Pre-flight Tool Check

The skill MUST verify that `git`, `mvn`, `npm`, `ng`, and `gradle` (or `gradlew`) are available in PATH before any repository work begins.

#### Scenario: All tools present

- GIVEN all required tools are installed and in PATH
- WHEN pre-flight runs
- THEN the pipeline proceeds to repository operations

#### Scenario: One or more tools missing

- GIVEN `mvn` is not installed
- WHEN pre-flight runs
- THEN the skill MUST halt with an error listing every missing tool by name; no repos are cloned

---

### Requirement: FR-003 — Repository Clone / Pull

The skill MUST clone each repository with `--depth 1` if it has not been cloned previously. If the local copy already exists, the skill MUST perform a `git pull` instead. The skill MUST support any GitHub organization.

#### Scenario: Repo not yet cloned

- GIVEN a repo URL that has not been cloned to the workspace
- WHEN processing begins for that repo
- THEN a shallow clone (`--depth 1`) of the specified branch is performed

#### Scenario: Repo already cloned

- GIVEN a local copy of the repo already exists
- WHEN processing begins for that repo
- THEN `git pull` is executed on the existing copy

---

### Requirement: FR-004 — Stack Detection and Build

The skill MUST detect the tech stack of each repo and execute the corresponding build command.

| Stack | Detection Signal | Build Command |
|-------|-----------------|---------------|
| java/maven | `pom.xml` present | `mvn clean package -DskipTests` |
| angular | `angular.json` present | `npm install && ng build --configuration=production` |
| react | `package.json` with `react` dep, no `angular.json` | `npm install && npm run build` |
| kotlin/android | `build.gradle.kts` or `build.gradle` with android plugin | `./gradlew assembleRelease` |
| react-native | `package.json` with `react-native` dep | `npm install && npx react-native bundle --platform android --dev false` |
| oracle-db | `*.sql` files present | No build — classify only |

#### Scenario: Java/Maven repo detected

- GIVEN a cloned repo contains `pom.xml`
- WHEN stack detection runs
- THEN `mvn clean package -DskipTests` is executed

#### Scenario: Oracle DB repo detected

- GIVEN a cloned repo contains `.sql` files and no build manifest
- WHEN stack detection runs
- THEN no build command is executed; SQL classification proceeds

#### Scenario: Unknown stack

- GIVEN a cloned repo matches none of the detection rules
- WHEN stack detection runs
- THEN the repo is marked as failed in the summary with reason "unrecognized stack"

---

### Requirement: FR-005 — Artifact Copy Rules

The skill MUST copy build outputs to the package folder using stack-specific rules.

| Stack | Source | Destination |
|-------|--------|-------------|
| java/maven | `target/*.ear` or `target/*.war` | `Install/AppDeploy/` |
| angular | `dist/` contents | `Install/AppDeploy/` |
| react | `dist/` contents | `Install/AppDeploy/` |
| kotlin/android | `app/build/outputs/**/*.apk` or `*.aab` | `Install/AppDeploy/` |
| react-native | `.apk` or bundle from build output | `Install/AppDeploy/` |
| oracle-db | `.sql` files | Classified via FR-006 |

#### Scenario: Maven artifact copied

- GIVEN a successful Maven build produced a `.war` in `target/`
- WHEN artifact copy runs
- THEN the `.war` is placed in `Install/AppDeploy/`

#### Scenario: No artifact produced

- GIVEN a build succeeds but produces no expected output file
- WHEN artifact copy runs
- THEN the repo is marked with a warning "no artifact found" in the summary

---

### Requirement: FR-006 — SQL File Classification

The skill MUST classify each `.sql` file into the correct subfolder under `Install/DBObjects/` based on filename pattern, then content inspection as fallback.

| Rule (evaluated in order) | Destination |
|--------------------------|-------------|
| Filename contains `PKS` or `PKB` | `DBObjects/Packages/` |
| Filename contains `_PRC`, `SP_`, `PROC`, or `_PROC` | `DBObjects/Procedures/` |
| Filename contains `_TYPE` or `TYPE_` | `DBObjects/Types/` |
| Filename contains `_TAB` or `_TABLE`, OR content contains `CREATE TABLE` | `DBObjects/Tables/` |
| Filename is `execution.sql`, OR first DML keyword in content is `INSERT`, `UPDATE`, or `MERGE` | `DBObjects/Data/` |
| No rule matched | `DBObjects/Unclassified/` + warning logged |

Numeric prefix handling: if the filename already has an `NN_` prefix (e.g. `07_SP_FOO.sql`), it MUST be preserved. If absent, the skill MUST assign a sequential `NN_` prefix within the destination subfolder (e.g. `01_`, `02_`).

#### Scenario: PKB file classified

- GIVEN a file named `PKG_FOO_PKB.sql`
- WHEN classification runs
- THEN the file is placed in `DBObjects/Packages/`

#### Scenario: Unclassified SQL

- GIVEN a file named `misc_fix.sql` that matches no classification rule
- WHEN classification runs
- THEN the file is placed in `DBObjects/Unclassified/` and a warning is added to the summary

#### Scenario: Existing numeric prefix preserved

- GIVEN a file named `07_SP_FOO.sql`
- WHEN classification runs
- THEN the file is placed as `07_SP_FOO.sql` in `DBObjects/Procedures/` without prefix modification

#### Scenario: Missing numeric prefix assigned

- GIVEN two files `SP_FOO.sql` and `SP_BAR.sql` in the same destination subfolder with no existing prefix
- WHEN classification runs
- THEN the files are placed as `01_SP_FOO.sql` and `02_SP_BAR.sql` sequentially

---

### Requirement: FR-007 — Shell Scripts Copy

The skill MUST copy `.sh` files found in the repo root or a `scripts/` directory to `Install/ShellScripts/`.

#### Scenario: Shell scripts present at repo root

- GIVEN a cloned repo contains `deploy.sh` at the root
- WHEN artifact copy runs
- THEN `deploy.sh` is placed in `Install/ShellScripts/`

---

### Requirement: FR-008 — README.txt Generation

The skill MUST generate `Install/Instalacion/README.txt` by filling the template asset with:
- Application info per repo: name, version, repo URL, branch
- Environment table with all 24 OWEXX environments (host and containers)
- WebLogic installation steps: stop servers, undeploy, deploy, start, validate
- Control de Versiones: version, date, author, user stories list, corrections list
- Control de Revisiones STTI

#### Scenario: README generated

- GIVEN all required inputs are present and the template asset exists
- WHEN packaging completes
- THEN `Install/Instalacion/README.txt` is created with all sections populated

---

### Requirement: FR-009 — Rollback Folder Mirror

The skill MUST create a `Rollback/` directory that mirrors the `Install/` structure with empty placeholder files. The skill MUST generate `Rollback/Instalacion/README.txt` with rollback instructions placeholder. No build artifacts are copied into `Rollback/`.

#### Scenario: Rollback structure created

- GIVEN the Install/ folder is fully assembled
- WHEN the rollback mirror step runs
- THEN `Rollback/` contains the same subfolder structure with no artifacts

---

### Requirement: FR-010 — Final Zip

The skill MUST produce a zip file named `{CHG}_VERSION{N}.zip` at the configured `outputPath` containing the entire package folder.

#### Scenario: Zip produced

- GIVEN packaging completed successfully
- WHEN the zip step runs
- THEN `{CHG}_VERSION{N}.zip` is created at `outputPath`

---

### Requirement: FR-011 — Partial Failure Handling

The skill MUST continue processing remaining repositories if one repo fails. Failed repos MUST be logged and included in the summary report. The final package MUST contain all successfully processed artifacts.

#### Scenario: One repo fails, others succeed

- GIVEN three repos and the second fails to build
- WHEN the pipeline runs
- THEN repos one and three are processed; the summary lists repo two as failed with the error message

---

### Requirement: FR-012 — Summary Report

The skill MUST print a structured summary after completion containing: package path, zip path, repos processed (success/fail with counts), files placed per folder, and warnings (unclassified SQL, missing artifacts).

#### Scenario: Summary printed after success

- GIVEN all repos built successfully
- WHEN the pipeline ends
- THEN a summary is printed with package path, zip path, all repos listed as success, and file counts per folder

---

## Non-Functional Requirements

### Requirement: NFR-001 — No Deployment

The skill MUST NOT deploy any artifact to any server, WebLogic instance, or environment.

#### Scenario: Post-build step attempted

- GIVEN the skill has assembled the package
- WHEN the pipeline ends
- THEN no deployment command is executed; output is limited to local file system writes

---

### Requirement: NFR-002 — No Credential Storage

The skill MUST NOT store, log, or persist any credentials, tokens, or secrets. Standard OS-level git credential helpers are the only accepted authentication mechanism.

#### Scenario: Private repo cloned

- GIVEN a repo requires authentication
- WHEN git clone executes
- THEN the skill relies on the OS git credential helper; no credentials are captured or stored by the skill

---

### Requirement: NFR-003 — Read-Only on Source Repos

The skill MUST NOT commit, push, or modify any file in the cloned source repositories. Git operations after clone MUST be limited to read-only (`git pull`, `git log`).

#### Scenario: Artifact copy does not touch source

- GIVEN a repo is cloned and built
- WHEN artifact copy runs
- THEN no `git commit` or `git push` is executed in the cloned repo

---

### Requirement: NFR-004 — Standalone, Not a DAG Phase

The skill MUST be user-invocable directly and MUST NOT be treated as a required SDD DAG phase.

#### Scenario: Skill invoked directly

- GIVEN a user provides valid inputs
- WHEN the skill is invoked outside of any SDD orchestration
- THEN the full pipeline runs without requiring SDD context

---

### Requirement: NFR-005 — Canonical Package Structure

The output package MUST exactly match the folder layout defined in FR-005 through FR-009. No extra top-level folders are created. No artifacts are placed outside their designated destination.

#### Scenario: Package layout validated

- GIVEN the pipeline completes
- WHEN the output folder is inspected
- THEN only `Install/AppDeploy/`, `Install/DBObjects/`, `Install/ShellScripts/`, `Install/Instalacion/`, `Rollback/` (mirrored) are present

---

## Acceptance Scenarios

### AS-001 — Happy Path: Java + DB Repos, All Tools Present

- GIVEN all pre-flight tools are installed, a Java/Maven repo and an Oracle DB repo are provided
- WHEN the skill runs to completion
- THEN a correctly structured `{CHG}_VERSION{N}.zip` is produced with `.war` in `AppDeploy/`, SQL files classified in `DBObjects/`, `README.txt` present, and Rollback mirror empty

### AS-002 — Missing Tool

- GIVEN `mvn` is not installed on the host
- WHEN the skill begins
- THEN pre-flight fails immediately with an error listing `mvn` as missing; no repos are cloned; no output folder is created

### AS-003 — One Repo Build Fails

- GIVEN three repos and the Angular repo fails `ng build`
- WHEN the pipeline runs
- THEN the Java and DB repos complete; the summary report lists the Angular repo as failed with its error; the package is zipped with available artifacts

### AS-004 — Unclassified SQL File

- GIVEN an Oracle DB repo contains `misc_patch.sql` matching no classification rule
- WHEN SQL classification runs
- THEN the file is placed in `DBObjects/Unclassified/` and the summary report includes a warning for that file

### AS-005 — Angular Repo Only

- GIVEN a single Angular repo is provided
- WHEN the pipeline runs
- THEN `dist/` contents are copied to `Install/AppDeploy/`; no `DBObjects/` folder is created; the package and zip are valid

### AS-006 — SQL File with Existing Numeric Prefix

- GIVEN a file named `07_SP_FOO.sql`
- WHEN SQL classification runs
- THEN the file is placed in `DBObjects/Procedures/` as `07_SP_FOO.sql` with the prefix unchanged

### AS-007 — SQL File Without Numeric Prefix

- GIVEN a file named `SP_FOO.sql` with no existing prefix, and it is the first file placed in `DBObjects/Procedures/`
- WHEN SQL classification runs
- THEN the file is placed as `01_SP_FOO.sql`

### AS-008 — Rollback Folder Structure

- GIVEN the pipeline completes successfully
- WHEN the Rollback mirror step runs
- THEN `Rollback/` contains the same subfolder hierarchy as `Install/` but with no artifact files; `Rollback/Instalacion/README.txt` exists with a rollback instructions placeholder

---

## Out of Scope

- Deploying packages to any server or WebLogic instance
- Generating rollback artifact content (structure is mirrored; content is the user's responsibility)
- Retrieving change order data from ServiceNow or any ITSM system (repo list is provided as input)
- Credentials or secrets management beyond standard OS-level git credential helpers
- Supporting build stacks not listed in FR-004
- Validating business logic or correctness of SQL file content beyond classification rules
