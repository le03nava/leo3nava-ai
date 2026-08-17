## Current Testing State

As of this specification date, `strict_tdd: false`. No automated test runner is configured in `openspec/config.yaml`.

## Verification Path

Functional verification relies on manual checks defined in each change's `test-design.md`. There are no `pytest`, `go test`, `npm test`, or equivalent commands to run during apply.

## Static Checks

Linting, formatting, and type-checking tools are not configured. Convention compliance is verified through manual inspection.

## SDD Apply Behavior

Because `strict_tdd: false`, `sdd-apply` does not run a test suite after each task. It implements assigned tasks and documents completion evidence manually.

## Update Trigger

⚠️ This file must be updated when `strict_tdd` or `test_command` values change in `openspec/config.yaml`. If `strict_tdd` becomes `true`, `sdd-apply` must follow `strict-tdd.md` and a test command must be specified.

## Planned Tests Notation

Test cases planned in `test-design.md` are verified during `sdd-verify` by manual inspection of apply-progress and implementation evidence, not by running a test suite.
