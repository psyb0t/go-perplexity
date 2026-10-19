# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.4, 2026-10-19

- `ResponsesResponse.Preset` is a new field. It names the preset and version that ran the response, as `name@<version>`. It is present only when the request used a preset, and streams resumed with `starting_after` omit it.
- `ResponsesRequest.Preset`'s doc comment now says it accepts `name`, `name@latest`, or `name@<version>`. `name` and `name@latest` run the active version of that preset; `name@<version>` runs that exact version. Preset versions are opaque, so do not parse or order them.
- Sending a versioned value back as `ResponsesRequest.Preset` reruns that exact version, including any bug fix patches applied since it last ran. A preset value without a version means the run predates preset versions or ran without one, and reruns the active version.

## v0.1.3, 2026-10-17

- `DecisionsRequest.Model`'s doc comment, and the matching text on `ClientInterface.DecisionsV1DecisionsPost`, `DecisionsV1DecisionsPostWithBody`, `ClientWithResponsesInterface.DecisionsV1DecisionsPostWithResponse` and `DecisionsV1DecisionsPostWithBodyWithResponse`, now name `pplx-decider-v1.2` and `pplx-decider-flash-v1.2` as the decision models, plus the aliases `pplx-decider-latest`, `perplexity/pplx-decider-latest`, `pplx-decider-flash-latest` and `perplexity/pplx-decider-flash-latest`. The old names `pplx-decider-v1.1-27b` and `pplx-decider-v1-27b` are still documented as accepted, and now map to `pplx-decider-v1.2`.
- `DecisionsRequest.State` and `DecisionsRequest_State`'s doc comments now say images can appear anywhere in `state`, not only inside an array, and that each question allows up to 8 images.
- `DecisionsRequest`'s doc comment now gives the input limit as 256K tokens instead of 262,144 tokens.
- `ResponsesRequest.Model`'s doc comment updates its example model id from `anthropic/claude-sonnet-4-6` to `anthropic/claude-sonnet-5-5`.
- No exported identifier was added, removed or changed. All of the above are comment-only updates in the generated client, matching the upstream spec.

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
