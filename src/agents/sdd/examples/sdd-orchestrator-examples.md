# SDD Orchestrator Examples

Long-form examples live here so `prompts/sdd/sdd-orchestrator.md` stays focused on contracts and hard gates.

## Example 1 - Preflight Dialog -> Cached Session State

**Trigger:** User says "use SDD to add dark mode".

```text
Orchestrator: Entry Routing classifies this as a new mutating SDD change.
Orchestrator: [ask-questions tool - 2 base groups in one call]
  Pace:      Interactive <- user picks this
  Artifacts: OpenSpec    <- user picks this

Orchestrator caches session state:
  {
    "execution_mode":    "interactive",
    "artifact_store":    "openspec",
    "delivery_strategy": null,
    "review_budget_lines": null,
    "chain_strategy":    null
  }

-> Runs backend-aware Init Guard:
  artifact_store.mode = openspec
  checks openspec/config.yaml
  -> Not found -> delegates sdd-init sub-agent with artifact_store.mode = openspec
  -> sdd-init writes openspec/config.yaml. Delegates sdd-explore.
-> Interactive proposal question round before sdd-propose:
  "Before I write the proposal, what business outcome should dark mode optimize for?
   Any brand/accessibility constraints, target user situations, or first-slice boundaries?"
  User answers, corrects assumptions, or explicitly skips/approves.
-> Delegates sdd-propose with exploration ref + proposal-shaping answers/skip decision.
-> Shows proposal summary to user, asks: "Continuar con spec y design?"
```

If the user asks `/sdd-status` instead, Entry Routing treats it as read-only and bypasses preflight:

```text
Orchestrator: produces status only; does not ask Pace/Artifacts, run sdd-init, or mutate artifacts.
```

Delivery decisions are intentionally deferred until `sdd-tasks` produces the Review Workload Forecast or the user explicitly asks to decide delivery earlier:

```json
{
  "delivery_strategy": "ask-on-risk | auto-chain | single-pr | exception-ok | null",
  "review_budget_lines": "number | null",
  "chain_strategy": "stacked-to-main | feature-branch-chain | null"
}
```

Equivalent init checks by mode:

| Mode | Init check |
| --- | --- |
| `engram` | `sdd-init/{project}` plus `sdd/{project}/testing-capabilities` |
| `openspec` | `openspec/config.yaml` |

## Example 2 - Phase Returns `partial` -> Gatekeeper Catches -> Re-run -> Success

**Trigger:** Automatic mode, `sdd-spec` returns.

```text
sdd-spec returns:
  status:           "partial"    <- Gatekeeper: status is not "success"
  artifacts:        [{ type: "spec", reference: resolved by artifact_store.mode }]
  next_recommended: "sdd-design"

Gatekeeper checklist:
  [ ] Status: success?           FAIL - status is "partial"
  [x] Artifact metadata valid?   PASS - envelope declares persisted/readable spec ref for the selected backend
  [x] No hallucination?          PASS - resolved reference/path is concrete; no full artifact body read needed on this path
  [ ] No drift?                  SKIP - blocked by status failure
  [ ] Routing coherence?         SKIP - blocked by status failure

Action: Re-run sdd-spec with feedback:
  "Previous run returned status: partial.
   Specific failure: spec sections 3 and 4 are incomplete (missing acceptance criteria).
   Complete those sections and return status: success."

sdd-spec (retry) returns:
  status:           "success"
  artifacts:        [{ type: "spec", reference: resolved by artifact_store.mode }]
  next_recommended: "sdd-design"

Gatekeeper checklist (retry):
  [x] Status: success?           PASS
  [x] Artifact metadata valid?   PASS
  [x] No hallucination?          PASS
  [x] No drift?                  PASS
  [x] Routing coherence?         PASS

State Persistence Gate:
  Write state update before launching design:
  schemaName: sdd.state
  currentPhase: spec
  completedPhases: [spec]
  nextRecommended: design
  Read-back: PASS - persisted state matches intended transition

-> Phase Gate PASS + State Persistence Gate PASS. Orchestrator continues to sdd-design automatically.
   sdd-design receives the spec ref/path and reads the full spec content as its own required input.
```
