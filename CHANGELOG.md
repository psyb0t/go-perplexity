# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.7.0, 2026-11-25

- Removed the field `Error` (`*string`) from `ApiSearchPage`. Pages with no copy at the `snapshot_as_of` time are now omitted from the results.
- Updated docs for `ApiSearchPage.LastUpdated`: it is omitted on `snapshot_as_of` searches.
- Updated docs for `ApiSearchRequest.SnapshotAsOf`: it cannot be combined with `search_type` `people`, and a response can have fewer than `max_results` results, or none.
- Updated docs for `ContextRequest.SnapshotAsOf`: it cannot be combined with the html and links formats.
- Updated docs for `ContextRequest.TimeoutMs`: the 95000 maximum with criteria is no longer stated.
- The regenerated client matches the upstream spec (version 1.0.0).

## v0.6.2, 2026-11-24

- Added the monitor agent target types `MonitorAgentTarget` and `MonitorAgentTargetRequest` (type `agent`), with `MonitorAgentFilters`, `MonitorAgentFiltersRequest`, `MonitorAgentProfileReference` and the type constants `MonitorAgentTargetTypeAgent`, `MonitorAgentTargetRequestTypeAgent` and `MonitorAgentProfileReferenceTypeCustom`.
- Added `MonitorAgentUpdateChange` (type `agent_update`) and `MonitorAgentSource` for agent monitor findings, with the constant `MonitorAgentUpdateChangeTypeAgentUpdate`.
- Added the accessors `AsMonitorAgentTarget`, `FromMonitorAgentTarget` and `MergeMonitorAgentTarget` on `MonitorTarget`, the matching `...MonitorAgentTargetRequest` methods on `MonitorTargetRequest`, and the `...MonitorAgentUpdateChange` methods on `MonitorChange`.
- Added the field `SnapshotAsOf` (`*string`) to `ApiSearchRequest` and `ContextRequest`.
- Added the field `Error` (`*string`) to `ApiSearchPage`.
- Updated docs for monitor types, context metadata origin and the cache and max age request fields.
- The regenerated client matches the upstream spec (version 1.0.0).

## v0.6.1, 2026-11-21

- Added the type `DecisionToolCallDetails` with the field `Invocation` (`int64`), the billable operations for one grounding tool.
- Added the field `Usage.ToolCallsDetails` (`*map[string]DecisionToolCallDetails`), set with `mode: "web_grounding"`, keyed by tool name.
- Updated docs for `DecisionsRequest.MaxStaleness`: the API accepts it but ignores it.
- Updated docs for `DecisionsRequest.Mode`: a search costs $0.001 and a successful URL fetch costs $0.0005.
- Updated docs for `DecisionsRequest.Model`: `pplx-decider-v1-27b` is accepted, and `pplx-decider-v1.1-27b` is an alias of it.
- The regenerated client matches the upstream spec (version 1.0.0).

## v0.6.0, 2026-11-20

- Added a context endpoint: `Client.Context`, `Client.ContextWithBody`, `ClientWithResponses.ContextWithResponse` and `ContextWithBodyWithResponse`, with `ContextDocument`, `ContextCost` and `ContextErrorResponse` types.
- Added gateway endpoints: `GatewayCreateChatCompletion`, `GatewayCreateMessage`, `GatewayCreateResponse`, `GatewayListModels` and `GatewayRetrieveModel`, each with `WithBody` and `WithResponse` variants where applicable.
- Added `ApiSearchRequest.DisplayServerTime` and `ApiSearchErrorResponse`.
- Added `EmbeddingsRequest_Input` helpers `FromEmbeddingsRequestInput2`, `FromEmbeddingsRequestInput3`, `MergeEmbeddingsRequestInput2` and `MergeEmbeddingsRequestInput3`.
- Breaking: `EmbeddingsResponse.Data` changed from `*[]EmbeddingObject` to `*[]EmbeddingsResponse_Data_Item`.
- Breaking: removed `SearchSearchPostResponse.JSON422` and `GetJSON422`.
- Added `ClientInterface` and `ClientWithResponsesInterface` methods for the new endpoints, so custom implementations of these interfaces must add them.
- Fixed generation: new `scripts/patch-spec.sh` runs from `gen.go` and points the stream chunk `service_tier` and `usage` at `OrchestratorChatServiceTier` and `OrchestratorChatCompletionUsage`, so `OrchestratorChatCreateChatCompletionStreamResponse.Usage` is now `*OrchestratorChatCompletionUsage`.
- Added the dependency `github.com/oapi-codegen/nullable` v1.1.0 (vendored).

## v0.5.0, 2026-11-17

- Added the monitors endpoints. New methods on `Client`: `CreateMonitor`, `CreateMonitorWithBody`, `ListMonitors`, `GetMonitor`, `UpdateMonitor`, `UpdateMonitorWithBody`, `DeleteMonitor`, `PauseMonitor`, `ResumeMonitor`, `RunMonitor`, `ListMonitorRuns`, `GetMonitorRun` and `TestMonitorWebhook`. `ClientWithResponses` has the matching `...WithResponse` methods.
- Added the monitor types `Monitor`, `MonitorRun`, `MonitorChange`, `MonitorDiff`, `MonitorList`, `MonitorObject`, `CreateMonitorRequest`, `CreateMonitorParams`, `ListMonitorsParams`, `ListMonitorRunsParams` and the response types such as `CreateMonitorResponse` and `GetMonitorRunResponse`. Targets are page targets, and schedules are interval or cron schedules.
- Added the output item types `SandboxReadFileOutputItem`, `SandboxWriteFileOutputItem`, `ShareFileOutputItem` and `SkillLoadedOutputItem`, with the `OutputItem` methods `As...`, `From...` and `Merge...` for each.
- Added the field `LocateSpan.SearchResult`.
- Breaking for types that implement the interfaces: `ClientInterface` and `ClientWithResponsesInterface` gained the monitor methods, so mocks and wrappers must add them. Callers of `Client` and `ClientWithResponses` are unaffected.
- No exported identifier was removed. The regenerated client matches the upstream spec (version 1.0.0).

## v0.4.0, 2026-11-14

- Removed the optional fields `ApiSearchRequest.MaxTokens` (`*int`, JSON key `max_tokens`) and `ApiSearchRequest.MaxTokensPerPage` (`*int`, JSON key `max_tokens_per_page`). Callers that set either field must drop it.
- No other exported identifier changed. The regenerated client matches the upstream spec (version 1.0.0).

## v0.3.6, 2026-11-13

- Added the `locate` question type for decisions. New types `LocateQuestion` (`Instructions`, `Type`), `LocateAnswer` (`Spans`, `Type`) and `LocateSpan` (`Start`, `End`, `Text`, `Score`). A `locate` answer lists the verbatim spans of `state` that match the instructions, with offsets in Unicode characters.
- Added the enum types `LocateQuestionType` and `LocateAnswerType`, each with the constant `...TypeLocate` and a `Valid` method.
- Added the union methods `Question.AsLocateQuestion`, `Question.FromLocateQuestion`, `Question.MergeLocateQuestion`, `Answer.AsLocateAnswer`, `Answer.FromLocateAnswer` and `Answer.MergeLocateAnswer`. `Question.ValueByDiscriminator` and `Answer.ValueByDiscriminator` now handle the `locate` type.
- The doc comments on `DecisionsRequest` and `Usage` and the decisions endpoint methods now name the models `pplx-decider-v1.3` and `pplx-decider-flash-v1.3`. The `pplx-decider-latest` aliases now select the v1.3 models, and only the v1.3 models serve `locate` questions.
- `DecisionsRequest.Questions`'s doc comment now says `locate` questions are not accepted with `mode` set to `web_grounding`. Existing callers are unaffected.

## v0.3.5, 2026-11-12

- `ResponsesRequest.MaxOutputTokens`: the doc comment no longer says the field is required for `anthropic/*` models. It now says the field is an optional total output-token budget across the run, including for Anthropic models. When neither the request nor its effort level or saved profile sets a budget, each generation can produce up to 128,000 output tokens, capped by the model's supported limit.
- Raised the minimum Go version in `go.mod` from 1.26.10 to 1.26.11, which fixes GO-2026-6859 in `net/http`. The development image and the setup reference follow.
- No exported identifier was added, removed or changed. The client change is a comment-only update matching the upstream spec (version 1.0.0).

## v0.3.4, 2026-11-05

- Added web grounding for decisions. New optional request fields `DecisionsRequest.Mode` (type `*DecisionsRequestMode`, constant `DecisionsRequestModeWebGrounding`, with a `Valid` method) and `DecisionsRequest.MaxStaleness` (`*int64`, seconds). With `mode` set to `web_grounding`, the API searches the web before the model answers and fetches the page for a URL-only `state`. `MaxStaleness` limits the age of the fetched page and is ignored without `Mode`.
- Added optional response fields `DecisionsResponse.SearchResults` (`*[]DecisionSearchResult`) and `DecisionsResponse.FetchedResults` (`*[]DecisionFetchedResult`). They hold the search evidence and fetched page text, and `usage.input_tokens` counts both.
- Added the types `DecisionSearchResult` (`Url`, `Title`, `Snippet`, `Date`, `LastUpdated`) and `DecisionFetchedResult` (`Url`, `Title`, `Snippet`).
- `DecisionsRequest.Questions`'s doc comment now says at least one question, with at most 100 when `mode` is `web_grounding`. More returns `400` with `param` `questions`.
- The doc comments on `DecisionsRequest.State`, `DecisionsRequest_State` and the decisions endpoint methods on `Client`, `ClientWithResponses` and their interfaces describe the new mode. Existing callers are unaffected.

## v0.3.3, 2026-11-03

- `ComputerUsageResponse.Categories` and `ComputerUsageV2Response.Categories`: the doc comments no longer describe the Credit Source axis as optional paid/promo data. They say to use the list to see which breakdowns are available, and that the `credit_usage` example uses the Model axis.
- `ComputerUsageCategoryCount`: the doc comment now gives a model name for the `credit_usage` Model axis as the example, instead of paid/promo.
- `ComputerUsageResponse.HasMore`: the doc comment now says it reports whether more complete buckets remain in the requested time window, and not whether the data has synced.
- `ComputerUsageResponse.NextPage`, `ComputerUsageV2Response.NextPage` and the `Page` field of `GetComputerUsageAnalyticsParams` and its v2 counterpart: the doc comments now call the cursor opaque and say to keep the other query parameters unchanged to continue the same report.
- `ComputerUsageV2Bucket.Results`, `ComputerUsageV2Response.Data`, `ComputerUsageV2Result` and `ComputerUsageV2Result.Count`: the doc comments now say usage is for the selected dataset, with credits for `credit_usage` and queries for `query_volume`, instead of credits only.
- The per-member usage endpoint doc comments on the `Client` and `ClientWithResponses` types and their interfaces say `credit_usage` returns each member's total credits and a Model breakdown, and drop the legacy Credit Source text.
- No exported identifier was added, removed or changed. This is a comment-only update in the generated client, matching the upstream spec (version 1.0.0).

## v0.3.2, 2026-10-29

- `ResponsesRequest.MaxSteps`: the doc comment now gives the maximum allowed value as 500 instead of 100.
- No exported identifier was added, removed or changed. This is a comment-only update in the generated client, matching the upstream spec (version 1.0.0).

## v0.3.1, 2026-10-29

- The doc comments on `ClientInterface.ListAdminAuditEventsV1AdminAuditEventsGet`, `Client.ListAdminAuditEventsV1AdminAuditEventsGet`, `ClientWithResponsesInterface.ListAdminAuditEventsV1AdminAuditEventsGetWithResponse` and `ClientWithResponses.ListAdminAuditEventsV1AdminAuditEventsGetWithResponse` now say the audit log API key is created under "Enterprise APIs, Audit API" instead of "Admin API", and link to the "Audit API guide" instead of the "Audit Log API guide".
- No exported identifier was added, removed or changed. This is a comment-only update in the generated client, matching the upstream spec (version 1.0.0).

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
