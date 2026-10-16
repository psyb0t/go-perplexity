// Package perplexity is a Go client for the Perplexity REST API, generated
// with oapi-codegen from the OpenAPI document Perplexity publishes.
package perplexity

//go:generate bash scripts/fetch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.json
