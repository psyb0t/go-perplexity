package perplexity_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-perplexity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-key"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *perplexity.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := perplexity.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)

			return nil
		},
	)

	client, err := perplexity.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertBearer(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))
}

func TestListModels(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/models", r.URL.Path)
		assertBearer(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"object":"list","data":[{"id":"model-a","object":"model",` +
				`"created":1700000000,"owned_by":"perplexity"}]}`,
		))
	})

	res, err := client.ListModelsWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Data, 1)

	assert.Equal(t, "model-a", res.JSON200.Data[0].Id)
	assert.Equal(t, 1700000000, res.JSON200.Data[0].Created)
	assert.Equal(t, "perplexity", res.JSON200.Data[0].OwnedBy)
}

func TestChatCompletion(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/sonar", r.URL.Path)
		assertBearer(t, r)

		var body map[string]any

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "sonar", body["model"])

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"chat-1","model":"sonar","created":1,` +
				`"choices":[{"index":0,"delta":{"role":"assistant",` +
				`"content":""},"message":{"role":"assistant",` +
				`"content":"hi"}}]}`,
		))
	})

	var content perplexity.ChatMessageInput_Content
	require.NoError(t, content.FromChatMessageInputContent0("hello"))

	res, err := client.ChatCompletionsChatCompletionsPostWithResponse(
		context.Background(),
		perplexity.ApiChatCompletionsRequest{
			Model: perplexity.ApiChatCompletionsRequestModelSonar,
			Messages: []perplexity.ChatMessageInput{{
				Content: &content,
				Role:    perplexity.ChatMessageRoleUser,
			}},
		},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Choices, 1)
	require.NotNil(t, res.JSON200.Choices[0].Message.Content)

	text, err := res.JSON200.Choices[0].Message.Content.
		AsChatMessageOutputContent0()
	require.NoError(t, err)
	assert.Equal(t, "hi", text)
}

func TestSearch(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/search", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assertBearer(t, r)

		var body map[string]any

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "some topic", body["query"])
		assert.InDelta(t, 3.0, body["max_results"], 0)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"search-1","results":[{"title":"A page",` +
				`"url":"https://example.com/a","snippet":"text"}]}`,
		))
	})

	var query perplexity.ApiSearchRequest_Query
	require.NoError(t, query.FromApiSearchRequestQuery0("some topic"))

	maxResults := 3

	res, err := client.SearchSearchPostWithResponse(
		context.Background(),
		perplexity.ApiSearchRequest{Query: query, MaxResults: &maxResults},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Results, 1)

	assert.Equal(t, "search-1", res.JSON200.Id)
	assert.Equal(t, "A page", res.JSON200.Results[0].Title)
	assert.Equal(t, "https://example.com/a", res.JSON200.Results[0].Url)
}
