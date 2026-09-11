# Ranger-Pulse

**Know what your infrastructure is doing.**

Ranger-Pulse is a lightweight, plugin-based monitoring and posture-check
engine for homelabs and small fleets — written in Go. Instead of a
one-size-fits-all metrics wall, each plugin answers one concrete question
("is SSH exposed?", "what's on my network?", "what changed in the logs?")
and reports it through a common agent → hub → UI pipeline.

> Status: early-stage / pre-alpha. This is an active learning project as
> much as a product — expect things to move fast and break.

## Why

Most homelab monitoring tools are either full observability stacks
(Prometheus + Grafana: powerful, heavy to run and configure) or single-
purpose dashboards (uptime-only, metrics-only). Ranger-Pulse aims for the
middle: a small core that any plugin — official or community — can extend,
without needing to touch the core or the UI.

## Architecture (short version)

- **Agent** — a single static binary that runs on each monitored machine.
  Discovers installed plugins, executes their collectors on schedule, and
  reports results to the Hub.
- **Hub** — the central service. Authenticates agents and users, stores
  reports, exposes a documented API, and enforces roles/permissions.
- **Plugin** — a collector binary + a declarative manifest. The manifest
  describes the collector's output schema *and* how the UI should render
  it (widget type, fields, thresholds) — plugins never ship executable UI
  code. See [`docs/plugin-manifest.md`](docs/plugin-manifest.md).
- **UI** — a generic renderer driven entirely by plugin manifests, plus
  the admin surface (users, roles, agents).

Full rationale lives in [`.specify/memory/constitution.md`](.specify/memory/constitution.md).

## Status of this repo

This project follows **Spec-Driven Development** (via
[GitHub Spec Kit](https://github.com/github/spec-kit)). Before there's
code, there's a spec:

- [`.specify/memory/constitution.md`](.specify/memory/constitution.md) — project principles
- [`specs/001-mvp-vertical-slice/spec.md`](specs/001-mvp-vertical-slice/spec.md) — current feature spec

## License

Ranger-Pulse core (this repo) is licensed under [Apache 2.0](LICENSE).
The project follows an **open-core** model — see
[`docs/open-core-boundary.md`](docs/open-core-boundary.md) for what is and
isn't covered by this license.

## Contributing

Not yet open for external contributions while the core architecture is
being validated — see [`CONTRIBUTING.md`](CONTRIBUTING.md) for the
current stance and how that will evolve.
