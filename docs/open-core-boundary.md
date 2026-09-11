# Open Core boundary

This repository (`ai4bscale/ranger-pulse`) is licensed under Apache 2.0
and contains everything needed to self-host Ranger-Pulse for free,
with no feature gates:

- The agent
- The Hub (storage, API, auth, RBAC)
- The base UI
- Reference/official plugins (e.g. `ssh-check`)

Per the working pricing hypothesis, self-hosting is free for one team.
Anything related to **paid usage beyond that** — issuing and validating
license keys, billing/subscription management, and any future hosted/
managed offering — lives in separate, private repositories
(`ai4bscale/ranger-pulse-billing` and successors), never inside this one.

This split is deliberate (see Article V of
[`.specify/memory/constitution.md`](../.specify/memory/constitution.md)):
the boundary is architectural and decided early, not retrofitted later by
crippling the open-source core.

This document will be updated as pricing and premium features firm up —
right now it describes intent, not a shipped mechanism.
