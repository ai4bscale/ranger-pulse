# Contributing to Ranger-Pulse

Thanks for your interest. Where this project stands right now:

**Early architecture phase.** The agent/hub/plugin contracts described in
[`.specify/memory/constitution.md`](.specify/memory/constitution.md) and
the current spec under [`specs/`](specs/) are still being validated end
to end. Large PRs against a moving foundation are hard to review fairly,
so for now:

- **Issues and discussion are welcome** — bug reports, questions, and
  design feedback on the current spec help a lot.
- **PRs**: please open an issue first to discuss the change before
  investing time in code. Small fixes (typos, docs, obviously-correct
  bugs) are always fine to send directly.

This will open up to normal PR-based contribution once the MVP vertical
slice (agent → hub → UI, one plugin, end to end) lands and stabilizes.

## Development workflow

This project follows Spec-Driven Development — see
[`.specify/memory/constitution.md`](.specify/memory/constitution.md) for
why, and the [Spec Kit](https://github.com/github/spec-kit) docs for the
mechanics (`/speckit.specify`, `/speckit.plan`, `/speckit.tasks`).

## Code of Conduct

Be respectful, assume good faith, keep discussion technical. A formal
Code of Conduct document will be added before external PRs open up.
