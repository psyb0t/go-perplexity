# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `perplexity_test`, and use only the exported API.
- **TST2** Client tests run against `httptest.NewServer` and check the method, path, `Authorization: Bearer` header, request body where there is one, and the decoded `JSON200`. They never call the live API.
- **TST3** `example_test.go` holds the README usage example, kept identical in substance. It needs a real key, so it has no `// Output:` line and only has to compile.
- **TST4** Test data uses placeholders (`model-a`, `test-key`, `example.com`), never real keys or real account data.
- **TST5** `make test` runs `go test -race` and must pass.
