---
name: install-package-builder
description: "Trigger: install package, paquete instalacion, build package, CHG deploy. Build a versioned installation package from git repos for a change order number."
user-invocable: true
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
---

## Activation Contract

Run when the user wants to build an installation package for a change order. Accept these inputs:
- changeNumber (required): e.g. CHG0086767
- version (default 1): package version number
- outputPath (required): directory where the package will be created
- repos (required): list of objects with { url, branch, type } where type is one of: java, angular, react, kotlin, react-native, oracle-db
- author (required): name for README control section
- userStories (optional): list of user story IDs
- corrections (optional): list of correction descriptions

This skill is standalone and user-invocable. It is NOT an SDD DAG phase and must NOT change SDD state.

## Pre-flight Check

Before any other work, verify required tools are installed by running `git --version`, `mvn --version`, `npm --version`. If any of these fail, STOP and report which tools are missing. Also run `ng version` and `gradle --version` (or `./gradlew --version`) as warnings only — warn if missing but do not stop.

## Input Schema

Input object:
- `changeNumber`: string, required
- `version`: string|number, default `1`
- `outputPath`: absolute path string, required
- `repos`: required list of `{ url, branch, type }`
- `author`: string, required
- `userStories`: optional list of strings
- `corrections`: optional list of strings

## Execution Steps

1. **Pre-flight**: Verify required tools (git, mvn, npm). Warn if ng or gradle missing.
2. **Clone/Pull**: For each repo, run `git clone --depth 1 -b {branch} {url} {tempDir}/{repoName}`. If already cloned, run `git pull`.
3. **Detect stack**: Inspect cloned repo root for stack signal files (see Stack Detection table).
4. **Build**: Run the build command for the detected stack (see Build Commands table).
5. **Classify SQL**: For oracle-db repos, classify each .sql file by filename pattern into the correct DBObjects subfolder (see SQL Classification Rules).
6. **Assemble package**: Copy compiled artifacts to `{outputPath}/{changeNumber}_VERSION{version}/Install/` following the Artifact Copy Rules.
7. **Generate README**: Fill readme-template.txt with change metadata and write to `Install/Instalacion/README.txt`. Read assets/environments.yaml for the environment table.
8. **Create Rollback structure**: Mirror Install/ folder structure under `Rollback/` with empty subfolders and a placeholder README.
9. **Zip and report**: Zip the package folder to `{outputPath}/{changeNumber}_VERSION{version}.zip`. Print the Summary Report.

## Stack Detection

| Signal File | Detected Stack |
|---|---|
| pom.xml | java |
| angular.json | angular |
| build.gradle or build.gradle.kts containing "com.android" | kotlin |
| package.json with "react-native" in dependencies | react-native |
| package.json with "react" in dependencies (no angular.json) | react |
| Only *.sql files, no build file | oracle-db |

## Build Commands

| Stack | Command | Artifact Location |
|---|---|---|
| java | `mvn clean package -DskipTests` | `target/*.ear` or `target/*.war` |
| angular | `npm install && ng build --configuration=production` | `dist/` |
| react | `npm install && npm run build` | `build/` or `dist/` |
| kotlin | `./gradlew assembleRelease` | `app/build/outputs/apk/release/*.apk` or `app/build/outputs/bundle/release/*.aab` |
| react-native | `npm install && npx react-native bundle --platform android --dev false --entry-file index.js --bundle-output android/app/src/main/assets/index.android.bundle` | `.apk` from Gradle or bundle file |
| oracle-db | (no build) | .sql files classified directly |

## SQL Classification Rules

Evaluate each .sql filename in this exact precedence order — stop at first match:

1. Filename contains `PKS` or `PKB` → `DBObjects/Packages/`
2. Filename contains `_PRC`, `SP_`, `PROC`, or `_PROC` → `DBObjects/Procedures/`
3. Filename contains `_TYPE` or `TYPE_` → `DBObjects/Types/`
4. Filename contains `_TAB` or `_TABLE`, OR file content begins with `CREATE TABLE` → `DBObjects/Tables/`
5. Filename is `execution.sql`, OR first non-comment SQL statement is INSERT/UPDATE/MERGE → `DBObjects/Data/`
6. No match → `DBObjects/Unclassified/` + log warning

**Numeric prefix rule**: If filename already starts with two digits and underscore (e.g. `07_`), preserve it. Otherwise, auto-assign the next sequential `NN_` within that subfolder.

## Artifact Copy Rules

| Stack | Source | Destination |
|---|---|---|
| java | `target/*.ear`, `target/*.war` | `Install/AppDeploy/` |
| angular | `dist/**` | `Install/AppDeploy/` |
| react | `build/**` or `dist/**` | `Install/AppDeploy/` |
| kotlin | `app/build/outputs/apk/**/*.apk`, `*.aab` | `Install/AppDeploy/` |
| react-native | output `.apk` or bundle | `Install/AppDeploy/` |
| oracle-db | classified .sql files | `Install/DBObjects/{subfolder}/` |
| .sh files | repo root or `scripts/` dir | `Install/ShellScripts/` |

## Hard Rules

- Do NOT deploy to any server or environment.
- Do NOT store, log, or accept credentials or secrets.
- Do NOT modify source repositories after cloning (read-only).
- Clean up temp clone directories after assembly.
- Re-running with the same inputs must overwrite the previous output cleanly (idempotent).
- If one repo fails to build, log the error and continue with remaining repos.

## Output Contract

After completion, print a structured Summary Report:
- Package path and zip path
- Per-repo result: success or failure with error message
- Files placed per folder (counts)
- Warnings (unclassified SQL files, missing expected artifacts, missing optional tools)
