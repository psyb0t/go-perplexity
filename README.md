# go-perplexity

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-perplexity.svg)](https://pkg.go.dev/github.com/psyb0t/go-perplexity)
[![CI](https://github.com/psyb0t/go-perplexity/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-perplexity/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-perplexity/badges/version.svg)](https://github.com/psyb0t/go-perplexity/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-perplexity/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-perplexity/badges/importers.svg)](https://github.com/psyb0t/go-perplexity/blob/badges/importers.md)

Unofficial Go client for the [Perplexity API](https://docs.perplexity.ai), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Perplexity publishes. I am not affiliated with Perplexity.

The package is `perplexity`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-perplexity
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-perplexity"
)

func main() {
	auth := perplexity.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("PERPLEXITY_API_KEY"))

		return nil
	})

	client, err := perplexity.NewClientWithResponses("https://api.perplexity.ai", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.ListModelsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

The full API is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-perplexity); from inside your module, `go doc github.com/psyb0t/go-perplexity` prints the same.

Method names come from the spec's operation IDs. Some are long because the spec derives them from the route: the Sonar chat endpoint is `ChatCompletionsChatCompletionsPost`, search is `SearchSearchPost`, and the async Sonar endpoints carry an `AsyncChatCompletions` prefix and suffix. The list models response type is `ListModelsResponse2`, because `ListModelsResponse` is already a schema name in the spec.

## Auth

Perplexity authenticates with a bearer token. Create an API key in the API settings of a Perplexity account and send it in the `Authorization` header, as in the request editor above. The base URL is `https://api.perplexity.ai`.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://docs.perplexity.ai/openapi.json` to `api/openapi.json`.
- oapi-codegen runs with `config.yaml`. The config maps `number` to `float64`, because the spec gives most numbers no format and oapi-codegen would otherwise use `float32`.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-perplexity) teaches a coding agent how to use go-perplexity when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-perplexity@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-perplexity@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-perplexity
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Perplexity's published API specification. The specification itself is not redistributed here.
