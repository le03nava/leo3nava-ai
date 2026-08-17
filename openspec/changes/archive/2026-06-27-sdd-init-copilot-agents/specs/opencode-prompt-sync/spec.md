# Opencode Prompt Sync Specification

## Purpose

Defines Windows-only behavior for initializing or refreshing opencode prompt files from this repository's `agents/` directory into the current user's opencode prompts directory.

## Requirements

### Requirement: Repository Prompt Source Resolution

The sync command at `scripts/sdd_init_opencode_prompts.ps1` MUST resolve the source `agents/` directory from the repository location containing the command, not from the shell's current working directory. The source MUST be limited to repository `agents/` contents, which are installed as opencode prompts.

#### Scenario: Run from repository root

- GIVEN the command is launched from the repository root
- WHEN source resolution runs
- THEN the source MUST resolve to the repository `agents/` directory

#### Scenario: Run from another working directory

- GIVEN the command is launched while the shell current directory is outside the repository
- WHEN source resolution runs
- THEN the source MUST still resolve to the repository `agents/` directory associated with the command location

### Requirement: Windows User Destination Resolution

The sync command MUST derive the destination from the current Windows user's home environment and target `.config\opencode\prompts` below that home. It MUST support `%USERPROFILE%` semantics or the PowerShell equivalent and MUST NOT hardcode a user-specific home path.

#### Scenario: User home is available

- GIVEN the current user's home environment value is available
- WHEN destination resolution runs
- THEN the destination MUST be `<home>\.config\opencode\prompts`

#### Scenario: User home is unavailable

- GIVEN no usable current-user home environment value is available
- WHEN destination resolution runs
- THEN the command MUST fail before copying
- AND the failure MUST clearly identify the missing home environment value

### Requirement: Destination Directory Creation

The sync command MUST create the opencode prompts destination directory when it does not exist.

#### Scenario: Destination is missing

- GIVEN `<home>\.config\opencode\prompts` does not exist
- WHEN the sync command runs
- THEN the command MUST create the destination directory
- AND it MUST continue with agent file synchronization

### Requirement: Prompt File Synchronization Semantics

The sync command MUST copy repository `agents/` contents into the opencode prompts destination and MUST replace destination files that correspond to copied source files on rerun. It MUST NOT delete unrelated files in the destination unless those files are replaced by copied source files with the same relative path.

#### Scenario: First sync copies prompts

- GIVEN repository `agents/` contains prompt source files
- AND the destination exists
- WHEN the sync command runs
- THEN each source file MUST exist at the matching relative path under the destination

#### Scenario: Rerun replaces copied files

- GIVEN a destination file exists at the same relative path as a repository `agents/` source file
- WHEN the sync command runs again
- THEN the destination file MUST be overwritten with the repository version

#### Scenario: Unrelated destination files are preserved

- GIVEN the destination contains a file with no matching relative path in repository `agents/`
- WHEN the sync command runs
- THEN that unrelated file MUST remain present and unchanged

### Requirement: Missing Source Failure

The sync command MUST fail clearly when the repository `agents/` source directory is missing or unavailable, and MUST NOT create or modify destination prompt files in that case.

#### Scenario: Source directory is missing

- GIVEN the repository `agents/` directory is unavailable
- WHEN the sync command runs
- THEN the command MUST fail before synchronization
- AND the failure MUST clearly identify the missing repository `agents/` source

### Requirement: Adapter Scope Boundaries

The sync command MUST synchronize only repository `agents/` contents into opencode prompts. It MUST NOT sync repository `skills/`, Claude prompts, other adapter folders, or install or configure external editor tooling.

#### Scenario: Other adapter folders exist

- GIVEN the repository contains skills, prompts, or adapter-specific folders outside `agents/`
- WHEN the sync command runs
- THEN those folders MUST NOT be copied to `.config\opencode\prompts`

#### Scenario: Editor tooling is not configured

- GIVEN opencode is not installed or configured
- WHEN the sync command runs
- THEN the command MUST NOT attempt to install or configure opencode
