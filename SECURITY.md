# Security Policy

Ranger-Pulse collects security-relevant data about the machines it
monitors (e.g. whether SSH is exposed). Taking vulnerability reports
seriously matters more than usual for a project like this.

## Reporting a vulnerability

Please **do not** open a public GitHub issue for a suspected security
vulnerability. Instead, use GitHub's private vulnerability reporting for
this repository (Security tab → "Report a vulnerability"), or email
security@ai4bscale.dev (placeholder — update once the domain/inbox exist).

Include what you'd normally include in a report: affected component
(agent, hub, UI, a specific plugin), steps to reproduce, and potential
impact. We'll acknowledge reports and follow up as this policy matures
alongside the project.

## Scope

At this stage (pre-alpha, no released binaries), the most relevant
reports are:
- Anything that would let a plugin's manifest execute code in the UI
  (this violates Article III of the project constitution and is treated
  as critical).
- Anything that would let an unauthenticated request reach agent or
  report data through the Hub API.
