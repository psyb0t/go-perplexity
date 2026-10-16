# Go code

Refines `~/.claude/rule-details/coding-style.md`.

- **GO1** `client.gen.go` is oapi-codegen output. Never edit it by hand. Change `config.yaml`, or add a patch script under `scripts/` run from `gen.go`, then `make generate`.
- **GO2** A spec problem that breaks oapi-codegen is fixed in a patch step run by `go:generate` (jq for the JSON spec), never in the generated file. The spec is downloaded into the gitignored `api/` directory and never committed.
- **GO3** The whole spec is generated. No include-tags or path trimming.
- **GO4** `number` maps to `float64` in `config.yaml`. Keep it.
- **GO5** Hand-written Go (tests, examples) passes `make lint` with zero issues. Lines are at most 80 columns, the `.golangci.yml` limit; split calls one argument per line rather than adding `//nolint`.
- **GO6** `make sec` passes. `scripts/sec.sh` excludes only semgrep's unsafe-deserialization-interface rule, which the generated `interface{}` fields for free-form objects trip. Do not widen it. A gitleaks hit on a sample value from the vendor spec is allowlisted for that exact value in the generated file only.
