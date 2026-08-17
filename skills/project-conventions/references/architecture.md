## SDD Executor Model

Each SDD phase is executed by a dedicated executor agent. Phases execute their own skill instructions directly and do not launch sub-agents. The orchestrator coordinates phase order, injects supplemental context, and validates phase outcomes.

## Shared Contracts

Common logic and cross-phase contracts live in `skills/_shared/`. Phase and project skills should reference shared contracts instead of duplicating them.

## Supplemental Skills

Project-specific skills, including `project-conventions`, are injected by the orchestrator through the skill-resolver pattern. Phase executors consume them from `## Skills to load before work` when provided. Core SDD phase skills are not part of the supplemental skill registry.

## Boundary Rule

`project-conventions` guides design decisions, conventions, and expected evidence. `sdd-review` is the sole owner of 96-control matrix judgment. `sdd-review-security` is the sole owner of `review-security-report.json` and formal security validation.

## Phase DAG

The SDD phase order is:

`explore → propose → spec → design → test-design → tasks → apply → review → review-security → verify → archive`

Do not run dependent phases in parallel.
