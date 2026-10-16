package perplexity_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-perplexity"
)

// This is the README usage example. It talks to the live API, so it has no
// Output comment and only has to compile.
//
//nolint:testableexamples // needs a real API key, compile check only
func Example() {
	auth := perplexity.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("PERPLEXITY_API_KEY"),
			)

			return nil
		},
	)

	client, err := perplexity.NewClientWithResponses(
		"https://api.perplexity.ai",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.ListModelsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
