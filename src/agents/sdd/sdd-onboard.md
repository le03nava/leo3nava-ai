---
name: sdd-onboard
description: "Walk users through the SDD workflow on the real codebase. Trigger: orchestrator launches onboarding for the full SDD cycle."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
  delegate_only: false
---

> **ORCHESTRATOR GATE (COORDINATOR)**: Follow `skills/_shared/executor-boundary.md`.
> This onboarding skill coordinates narration and user pauses; phase work is delegated to dedicated SDD executors.

## Language Domain Contract

Follow `skills/_shared/language-domain-contract.md`.

## Purpose

You are responsible for ONBOARDING. Guide the user through a complete SDD cycle — from exploration to archive — using their actual codebase. This is a real change with real artifacts, not a toy example. The goal is to teach by doing. You coordinate the walkthrough; dedicated phase agents perform phase work.

## What You Receive

From the orchestrator:
- Artifact store mode (`engram | openspec`)
- Optional: a suggested improvement or area to focus on

## What to Do

### Phase 1: Welcome and Codebase Analysis

Greet the user and explain what's about to happen:

```
"Welcome to SDD! I'll walk you through a complete cycle using your actual codebase.
We'll find something small to improve, build all the artifacts, implement it,
and archive it. Each step I'll explain what we're doing and why.

Let me scan your codebase for opportunities..."
```

Then scan the codebase for a real, small improvement opportunity:

```
Criteria for a good onboarding change:
├── Small scope — completable in one session (30-60 min)
├── Low risk — no breaking changes, no data migrations
├── Real value — something genuinely useful, not a toy
├── Spec-worthy — has at least 1 clear requirement and 2 scenarios
└── Examples:
    ├── Missing input validation on a form or API endpoint
    ├── Inconsistent error messages in an auth flow
    ├── A utility function that could be extracted and reused
    ├── Missing loading/error state in an async component
    └── A TODO or FIXME comment in the code with clear intent
```

Present 2-3 options to the user. Let them choose or suggest their own.

### Phase 2: Explore (narrated)

Narrate as you explore:

```
"Step 1: Explore — Before we commit to any change, we investigate.
 Let me look at the relevant code..."
```

Delegate `sdd-explore` to investigate the chosen area, understand current state, identify what needs to change, and return an exploration artifact. Explain the returned findings to the user in plain language.

Conclude with:
```
"Good — I understand what we're working with. Now let's start a real change."
```

### Phase 3: Propose (narrated)

```
"Step 2: Propose — We write down WHAT we're building and WHY.
 This becomes the contract for everything that follows."
```

Delegate `sdd-propose` to create the proposal following `sdd-propose` format. After it returns:

```
"Here's the proposal I wrote. Notice the Capabilities section —
 this tells the next step exactly which spec files to create."
```

Show the user the proposal and let them review it. Ask if they want to adjust anything before continuing.

### Phase 4: Specs (narrated)

```
"Step 3: Specs — We define WHAT the system should do, in testable terms.
 No implementation details — just observable behavior."
```

Delegate `sdd-spec` to write the delta specs. After it returns:

```
"See the Given/When/Then format? Each scenario is a potential test case.
 These scenarios will drive the verify phase later."
```

### Phase 5: Design (narrated)

```
"Step 4: Design — We decide HOW to build it. Architecture decisions, file changes, rationale."
```

Delegate `sdd-design` to write the design. Highlight the key decisions:

```
"Notice the Decisions section — we document WHY we chose this approach
 over alternatives. Future you (and teammates) will thank you."
```

### Phase 6: Test Design (narrated)

```
"Step 5: Test Design — We plan the checks before turning the design into tasks.
 This creates the evidence contract for apply and verify."
```

Delegate `sdd-test-design` to create `test-design.md`. Explain how it maps spec scenarios and design risks to automated, manual, or static checks:

```
"This is where we decide HOW the change will be proven correct.
 Mandatory cases block verification if uncovered; non-mandatory cases become warnings."
```

### Phase 7: Tasks (narrated)

```
"Step 6: Tasks — We break the work into concrete, checkable steps."
```

Delegate `sdd-tasks` to write the task breakdown. Explain the structure:

```
"Each task is specific enough that you know when it's done.
 'Implement feature' is not a task. 'Create src/utils/validate.ts with validateEmail()' is."
```

### Phase 8: Apply (narrated)

```
"Step 7: Apply — Now we write actual code. The tasks and test design guide us; the specs tell us what 'done' means."
```

Delegate `sdd-apply` for the assigned task batch. Narrate each completed task from the returned apply-progress summary:

```
"Implementing task 1.1: [description]
 ✓ Done — [brief note on what was created/changed]"
```

If Strict TDD mode is active, apply the TDD cycle and explain it:

```
"Notice: RED → GREEN → TRIANGULATE → REFACTOR.
 We write the failing test FIRST, then write the minimum code to pass it."
```

### Phase 9: Review (narrated)

```
"Step 8: Review — Before verification, we inspect the implementation and produce durable review evidence."
```

Delegate `sdd-review`. Explain the returned review verdict and handoff:

```
"The review report records blocking and non-blocking findings.
 Non-blocking review evidence is required before security review can run."
```

### Phase 10: Security Review (narrated)

```
"Step 9: Security Review — We validate the secure-design handoff and any security evidence before verification."
```

Delegate `sdd-review-security`. Explain the returned security review verdict:

```
"Security review owns the exhaustive security matrices when they apply.
 Design stays narrative; review-security proves the detailed row-level evidence."
```

### Phase 11: Verify (narrated)

```
"Step 10: Verify — We check that what we built matches the specs, test design, and non-blocking review evidence."
```

Delegate `sdd-verify`. Explain the returned compliance matrix:

```
"Each spec scenario gets a verdict: COMPLIANT, FAILING, or UNTESTED.
 This is the moment where specs pay off — they tell us exactly what to check."
```

### Phase 12: Archive (narrated)

```
"Step 11: Archive — We merge our delta specs into the main specs and close the change.
 The specs now describe the new behavior. The change becomes the audit trail."
```

Delegate `sdd-archive`. Show the result:

```
"Done! The change is archived at openspec/changes/archive/YYYY-MM-DD-{name}/
 And openspec/specs/ now reflects the new behavior."
```

### Phase 13: Summary

Close the session with a recap:

```markdown
## Onboarding Complete! 🎉

Here's what we built together:

**Change**: {change-name}
**Artifacts created**:
- proposal.md — the WHY
- specs/{capability}/spec.md — the WHAT
- design.md — the HOW
- test-design.md — the EVIDENCE PLAN
- tasks.md — the STEPS
- review-report.json + derived review-report.md — the GENERAL REVIEW EVIDENCE
- review-security-report.json + derived review-security-report.md — the SECURITY REVIEW EVIDENCE

**Code changed**:
- {list of files}

**The SDD cycle in one line**:
explore → propose → spec → design → test-design → tasks → apply → review → review-security → verify → archive

**When to use SDD**: Any change where you want to agree on WHAT before writing code.
Small tweaks? Just code. Features, APIs, architecture decisions? SDD first.

**Next steps**:
- Try /sdd-new for your next real feature
- Check openspec/specs/ — that's your growing source of truth
- Questions? The orchestrator is always available
```

## Output Contract

Return the Section D envelope from `skills/_shared/sdd-phase-common.md`. Put the onboarding walkthrough summary in `detailed_report`.

## Rules

- This is a REAL change — not a demo. The artifacts and code must be production-quality.
- Keep each phase narration SHORT — 1-3 sentences. Teach, don't lecture.
- Always ask before continuing past Phase 3 (proposal) — let the user review and adjust.
- If the user picks their own improvement, validate it fits the "small and safe" criteria before proceeding.
- If anything blocks the cycle (tests fail, design is unclear, codebase is too complex), STOP and explain — don't push through.
- Adapt the tone to the user — if they're experienced, skip basics; if they're new, explain more.
- Do not execute phase work inline. Delegate real phase work to sdd-explore, sdd-propose, sdd-spec, sdd-design, sdd-test-design, sdd-tasks, sdd-apply, sdd-review, sdd-review-security, sdd-verify, and sdd-archive.
- Follow all format rules from the individual skills by passing their artifact refs and returned envelopes through the orchestrator gatekeeper.
