# Opencode Skill Sync Specification

## Purpose

Define repo-local behavior for initializing or refreshing opencode skills from this repository's `skills/` source into the current Windows user's opencode skills directory.

## Requirements

### Requirement: Sync Repository Skills to Opencode

The system MUST provide a Windows-invoked sync operation at `scripts/sdd_init_opencode_skills.ps1` that copies the direct contents of the repository `skills/` directory into `%USERPROFILE%\.config\opencode\skills`. The system MUST create the destination directory when it does not exist. The system SHALL copy all direct source contents without prompting for selection or filtering.

#### Scenario: Initial sync creates destination and copies skills

- GIVEN `$env:USERPROFILE` is set and the repository `skills/` directory exists
- AND `%USERPROFILE%\.config\opencode\skills` does not exist
- WHEN the sync operation is run
- THEN `%USERPROFILE%\.config\opencode\skills` MUST exist
- AND the direct contents of repository `skills/` MUST be present in the destination

#### Scenario: Existing destination receives all source contents

- GIVEN `$env:USERPROFILE` is set and the repository `skills/` directory exists
- AND `%USERPROFILE%\.config\opencode\skills` already exists
- WHEN the sync operation is run
- THEN the direct contents of repository `skills/` MUST be copied into `%USERPROFILE%\.config\opencode\skills`

### Requirement: Fail Fast on Missing Inputs

The system MUST fail before copying when required inputs are unavailable. It MUST report a clear error when `$env:USERPROFILE` is empty or missing. It MUST report a clear error when the repository `skills/` source directory is missing. It MUST NOT create or modify the opencode skills destination after detecting either error.

#### Scenario: Missing USERPROFILE fails without destination changes

- GIVEN `$env:USERPROFILE` is empty or missing
- WHEN the sync operation is run
- THEN the operation MUST fail with a clear `USERPROFILE` error
- AND the operation MUST NOT copy any skills

#### Scenario: Missing source directory fails without destination changes

- GIVEN `$env:USERPROFILE` is set
- AND the repository `skills/` directory is missing
- WHEN the sync operation is run
- THEN the operation MUST fail with a clear missing source error
- AND the operation MUST NOT copy any skills

### Requirement: Deterministic Reruns

The system MUST make reruns deterministic by replacing existing destination files with the current repository `skills/` contents. The system SHALL allow the operation to be run repeatedly without requiring manual cleanup.

#### Scenario: Rerun overwrites previously copied files

- GIVEN `$env:USERPROFILE` is set and the repository `skills/` directory exists
- AND `%USERPROFILE%\.config\opencode\skills` contains a file with the same relative path as a source file
- WHEN the sync operation is run
- THEN the destination file MUST be replaced by the current source file
- AND the operation MUST complete without an interactive prompt
