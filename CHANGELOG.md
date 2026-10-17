# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `perplexity`, a REST client generated with oapi-codegen from Perplexity's OpenAPI document (Perplexity AI API, spec version 1.0.0, OpenAPI 3.1.0, from docs.perplexity.ai). It covers all 19 paths: Sonar chat completions and their async variant, search, embeddings and contextualized embeddings, agents and agent files, skills, decisions, models and computer usage analytics.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from docs.perplexity.ai and regenerates the client. The spec is not committed.
- `make sec` skips semgrep's unsafe-deserialization-interface rule, which flags the `interface{}` values oapi-codegen generates for the spec's free-form objects.
- Added the `go-perplexity` agent skill: `claude plugin install go-perplexity@psyb0t` (Claude Code), `codex plugin add go-perplexity@psyb0t` (Codex), `openclaw skills install @psyb0t/go-perplexity` (ClawHub).
