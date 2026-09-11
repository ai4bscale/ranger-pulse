# Feature Spec: MVP Vertical Slice — Agent → Hub → UI with one plugin

**Status:** Draft — ready for `/speckit.clarify` review
**Feature branch:** `001-mvp-vertical-slice`

## Why

We need one thing working end-to-end — one agent reporting one real,
useful piece of data, stored by the Hub, rendered by the UI, behind auth —
before adding any more plugins or polish. This slice proves the core
contracts (Article II, III, IV, IX of the constitution) actually hold
before we build on top of them.

## User Stories

1. **As a homelab operator**, I install the agent on a Linux machine and,
   within one collection interval, see in the Hub UI whether SSH is
   enabled on that machine and what state it's in — so I can catch a
   forgotten open SSH port without logging into every box by hand.
2. **As the Hub administrator**, I log in with a username/password and see
   a list of every agent that has reported in, with its last-seen time and
   status — so I know at a glance what's alive.
3. **As a second, read-only user**, I can view the same dashboard but
   cannot change agent configuration or user accounts — so I can share
   visibility with someone without giving them control.

## Functional Requirements

- **FR1** — The agent MUST discover plugin packages (collector binary +
  manifest) from a local plugins directory at startup and on a configurable
  refresh interval.
- **FR2** — The agent MUST execute each discovered plugin's collector on
  the interval declared in its manifest and capture its stdout as JSON.
- **FR3** — The agent MUST reject / quarantine a plugin whose output does
  not match the JSON schema declared in its own manifest, and log the
  failure rather than forwarding malformed data.
- **FR4** — The agent MUST send each successful collection result to the
  configured Hub over HTTPS, authenticated with a per-agent enrollment
  token, including agent ID, plugin ID, timestamp, and payload.
- **FR5** — The Hub MUST persist each report with enough metadata to query
  "latest report for agent X / plugin Y."
- **FR6** — The Hub MUST expose a documented REST API to: list agents,
  fetch an agent's latest report(s), and fetch a plugin's manifest.
- **FR7** — The Hub MUST require authentication for every API route except
  agent enrollment/ingestion (which uses the enrollment token instead).
- **FR8** — The Hub MUST support at least two roles: `admin` (full access,
  including user/agent management) and `viewer` (read-only dashboard
  access).
- **FR9** — The UI MUST provide a login screen and MUST NOT render any
  agent or report data to an unauthenticated session.
- **FR10** — The UI MUST list connected agents with last-seen time and a
  derived status (online / stale / offline based on a configurable
  threshold).
- **FR11** — The UI MUST render the `ssh-check` plugin's output using only
  the widget type and field hints declared in that plugin's manifest —
  no plugin-specific code in the UI (enforces constitution Article III).
- **FR12** — The reference `ssh-check` plugin MUST report at minimum:
  whether an SSH daemon is listening, on which port, and whether
  password authentication is enabled — as this is the concrete "pain"
  this slice exists to solve.

## Non-Goals (explicitly out of scope for this slice)

- Historical time-series charts / long-term retention policy.
- Alerting or notifications (Slack, email, webhooks).
- A plugin marketplace or remote plugin installation UI.
- Multi-tenant organizations (single Hub = single tenant for now).
- SSO / external identity providers (local accounts only).
- Additional plugins (network inventory, events/logs) — each gets its own
  spec once this slice is proven end-to-end.
- Any Open Core / paywall mechanics — this slice is 100% core/open.

## Acceptance Criteria (sample — expand per requirement in review)

- **Given** an agent with the `ssh-check` plugin installed, **when** the
  collection interval elapses, **then** the Hub has a report for that
  agent/plugin pair no older than one interval + transport latency.
- **Given** a viewer-role user, **when** they attempt to call an
  admin-only API route, **then** the Hub returns 403 and the UI does not
  expose the corresponding action.
- **Given** a plugin whose collector outputs JSON that fails its own
  declared schema, **when** the agent runs it, **then** the bad report is
  never forwarded to the Hub and an error is logged locally.

## Open Questions for `/speckit.clarify`

- Enrollment flow: manual token copy-paste vs. a short-lived
  install-script-generated token?
- Report storage for MVP: embedded (SQLite/BoltDB) vs. requiring an
  external DB — constitution favors lightweight/no-dependency (Article I),
  suggesting embedded for now.
- Transport: plain REST/JSON over HTTPS vs. gRPC from day one — REST is
  simpler to learn first; gRPC can come with Article IV's API versioning
  later without breaking this slice.
