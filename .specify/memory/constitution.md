# Project Constitution

**Project:** Ranger-Pulse — a pluggable, agent-based monitoring & posture-check
engine for homelabs, written in Go.

This document defines the non-negotiable principles for this project.
Every spec, plan, and task must comply with these articles. Changes to this
file require explicit, deliberate revision — not a side effect of a feature.

---

## Article I — Lightweight by Default

The agent is the thing running on every monitored machine, including
resource-constrained hardware (Raspberry Pi class). It MUST:
- Ship as a single static binary, no runtime dependencies.
- Keep idle resource usage low enough to run alongside real workloads
  (target: comparable to other Go-based lightweight agents in this space,
  not a JVM- or Python-based collector).
- Never block or degrade the host system it monitors.

## Article II — Plugins Are the Extension Point, Not an Afterthought

The core (agent + hub) MUST remain plugin-agnostic. It knows how to:
load a plugin package, execute its collector, transport its output, store it,
and expose it — and nothing about what any specific plugin measures.
All domain logic (SSH checks, network scans, log parsing, etc.) lives in
plugins, including the ones shipped as "official" first-party plugins.

## Article III — No Third-Party Code Executes in the UI

A plugin package consists of:
1. A **collector** (compiled binary or script) — executes on the agent side,
   on infrastructure the operator already trusts.
2. A **manifest** (declarative JSON/YAML) — read by the Hub and UI to know
   the collector's output schema and how to render it.

The UI MUST render exclusively from the declarative manifest (widget type,
fields, thresholds, labels). The UI MUST NOT execute, `eval`, or load
arbitrary code shipped by a plugin. This is a security boundary, not a
style preference, and it may not be relaxed by any future feature.

## Article IV — API-First

The Hub exposes a versioned, documented API (REST and/or gRPC). The
first-party UI is a client of that API like any other — no data or
capability may exist in the UI that isn't reachable through the API.
This keeps a CLI, a future mobile client, or third-party integrations
possible without touching the core.

## Article V — Open Core Boundary Is Explicit

- **Core (agent, hub, plugin runtime, base UI, reference plugins for
  services/network/events)**: permissive open-source license (MIT or
  Apache 2.0). Fully functional and self-hostable with no feature gates.
- **Premium** (future): fleet management at scale, SSO/SAML, compliance
  exports, hosted/managed Hub, premium plugins. These live in separate
  modules/repos from day one — never as a flag hidden inside core code.
  The constitution does not commit to *what* is premium yet, only that
  the boundary is architectural, decided early, and never retrofitted by
  crippling the open core.

## Article VI — Learning-Lab Scope Discipline

This project exists primarily to learn Git/GitHub workflows, CI/CD,
release engineering, and Go — not to out-build existing tools on day one.
Every spec MUST define a Non-Goals section. Features outside the current
vertical slice are deferred, not silently absorbed into scope.

## Article VII — Observable & Testable Core

Core components (agent report pipeline, Hub ingestion, auth/RBAC) require
automated tests before being considered done. UI polish and additional
plugins may follow a lighter bar, but the trust boundary (Article III) and
the auth boundary (Article IX) are always covered by tests.

## Article VIII — Everything Ships Through CI

No artifact (binary, container image, release) is published by hand.
GitHub Actions builds, tests, and — once a spec/plan says a slice is
release-ready — packages it. This is deliberate: the CI/CD pipeline is
itself a first-class learning goal of this project, not incidental tooling.

## Article IX — Auth Is Not Optional, Even in the MVP

Because this handles security-relevant data (SSH exposure, network
inventory) and is meant to run on real infrastructure, the Hub MUST require
authentication and a minimal role model (at least admin / read-only) from
the first shippable slice — not bolted on after "it works."

---

*Ratified for the MVP vertical slice described in
`specs/001-mvp-vertical-slice/spec.md`. Revise via explicit constitution
update, not silently through a feature spec.*
