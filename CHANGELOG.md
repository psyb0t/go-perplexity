# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.3.0, 2026-10-28

- Added the admin audit events endpoint. New methods: `Client.ListAdminAuditEventsV1AdminAuditEventsGet` and `ClientWithResponses.ListAdminAuditEventsV1AdminAuditEventsGetWithResponse`, plus `NewListAdminAuditEventsV1AdminAuditEventsGetRequest` and `ParseListAdminAuditEventsV1AdminAuditEventsGetResponse`. Filter and page with `ListAdminAuditEventsV1AdminAuditEventsGetParams`.
- Added the types `AuditEvent`, `AuditEventsResponse`, `AdminApiErrorResponse` and `ListAdminAuditEventsV1AdminAuditEventsGetResponse`.
- Breaking: `ListBatchesParams.After` and `ListFilesParams.After` are renamed to `ListBatchesParams.AfterId` and `ListFilesParams.AfterId`. The query parameter changes from `after` to `after_id`. Callers must rename the field.
- Added methods to `ClientInterface` and `ClientWithResponsesInterface` for the new endpoint. Types that implement these interfaces must add them.
- The doc comments on `ComputerUsageResponse.Categories` and `ComputerUsageV2Response.Categories` now say the list holds the available category labels, and that the credit source axis is optional for credit usage.

## v0.2.4, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

## v0.2.3, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.2.2, 2026-10-26

- `DecisionsRequest`'s doc comment now says the grouped answering of questions applies to `pplx-decider-v1.2`. `pplx-decider-flash-v1.2` answers each question separately and counts the `state` again for every question.
- No exported identifier was added, removed or changed. This is a comment-only update in the generated client, matching the upstream spec (version 1.0.0).

## v0.2.1, 2026-10-25

- Added optional fields `PeopleSearchTool.MaxTokens` and `PeopleSearchTool.MaxTokensPerPage` (both `*int32`, JSON keys `max_tokens` and `max_tokens_per_page`). They cap the total tokens of search context and the tokens extracted per result page. Existing callers are unaffected.

## v0.2.0, 2026-10-22

- Added the Batch API. New `Client` and `ClientWithResponses` methods: `CreateBatch`, `CreateBatchWithBody`, `GetBatch`, `ListBatches`, `CancelBatch`, plus the matching `...WithResponse` methods. New types include `Batch`, `BatchStatus`, `BatchRequestCounts`, `BatchUsage`, `BatchList`, `CreateBatchRequest`, `CreateBatchParams` and `ListBatchesParams`.
- Added batch file endpoints: `CreateFileWithBody`, `ListFiles`, `GetFile`, `GetFileContent` and `DeleteFile`, with their `...WithResponse` methods, `New...Request` builders and `Parse...Response` functions. New types include `BatchFile`, `BatchFileList`, `BatchFileDeleted`, `BatchFilePurpose`, `ListFilesParams` and `CreateFileMultipartBody`.
- Breaking: `ResponsesRequest.Preset` and `ResponsesResponse.Preset` are replaced by `ResponsesRequest.Effort` and `ResponsesResponse.Effort`. The JSON key changes from `preset` to `effort`. The values and versioning (`name`, `name@latest`, `name@<version>`) work as before. Doc comments on `MaxSteps`, `Model` and `Profile` now say "effort level" instead of "preset".
- Added methods to `ClientInterface` and `ClientWithResponsesInterface` for every new endpoint. Types that implement these interfaces must add them.
- Raised the minimum Go version in `go.mod` from 1.26.9 to 1.26.10, which includes the fix for GO-2026-6809 in `crypto/x509`. The development image and the setup reference follow.

## v0.1.5, 2026-10-20

- `ChoiceAnswer.Confidence`'s doc comment now says it is computed from `probabilities` as `(p_top - 1/n) / (1 - 1/n)`. It is 0 when every option is equally likely and 1 when one option has all the probability or there is only one option.
- `ScoreAnswer.Confidence`'s doc comment now says it is computed from `probabilities` as `1 - spread / baseline`, with a floor of 0, and is 1 for a single level.
- `DecisionsRequest`'s doc comment now says the `state` plus any one question must stay under 256K input tokens. Questions that do not fit together are answered in groups, and each group counts the `state` again. The body limit is still 32 MiB.
- `Usage`'s doc comment now says input tokens are billed at the selected model's rate: $0.02 per million for `pplx-decider-v1.2` and $0.01 for `pplx-decider-flash-v1.2`. Output tokens are free.
- `ToolCallDetails.TokenUsage`'s doc comment no longer says advisor entries represent individual invocations.
- No exported identifier was added, removed or changed. All of the above are comment-only updates in the generated client, matching the upstream spec.

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
