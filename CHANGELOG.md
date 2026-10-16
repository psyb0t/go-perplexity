# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `perplexity`, a REST client generated with oapi-codegen from Perplexity's OpenAPI document (Perplexity AI API, spec version 1.0.0, OpenAPI 3.1.0, from docs.perplexity.ai). It covers all 19 paths: Sonar chat completions and their async variant, search, embeddings and contextualized embeddings, agents and agent files, skills, decisions, models and computer usage analytics.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from docs.perplexity.ai and regenerates the client. The spec is not committed.
- `make sec` skips semgrep's unsafe-deserialization-interface rule, which flags the `interface{}` values oapi-codegen generates for the spec's free-form objects.
- Added the `go-perplexity` agent skill: `claude plugin install go-perplexity@psyb0t` (Claude Code), `codex plugin add go-perplexity@psyb0t` (Codex), `openclaw skills install @psyb0t/go-perplexity` (ClawHub).
