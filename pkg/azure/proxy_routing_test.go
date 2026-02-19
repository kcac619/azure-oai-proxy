package azure

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestCompletionsRequestRoutesToAzureCompletionsEndpoint(t *testing.T) {
	originalEndpoint := AzureOpenAIEndpoint
	AzureOpenAIEndpoint = "https://example.openai.azure.com"
	defer func() { AzureOpenAIEndpoint = originalEndpoint }()

	body := []byte(`{"model":"gpt-5.2-codex","prompt":"hello"}`)
	req := httptest.NewRequest("POST", "http://localhost:11437/v1/completions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	makeDirector()(req)

	if got, want := req.URL.Path, "/openai/deployments/gpt-5.2-codex/completions"; got != want {
		t.Fatalf("unexpected proxied path: got %q want %q", got, want)
	}
}

func TestChatCompletionsForGPT52CodexAutoConvertsToResponses(t *testing.T) {
	originalEndpoint := AzureOpenAIEndpoint
	AzureOpenAIEndpoint = "https://example.openai.azure.com"
	defer func() { AzureOpenAIEndpoint = originalEndpoint }()

	body := []byte(`{"model":"gpt-5.2-codex","messages":[{"role":"user","content":"hello"}]}`)
	req := httptest.NewRequest("POST", "http://localhost:11437/v1/chat/completions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	makeDirector()(req)

	if got, want := req.URL.Path, "/openai/v1/responses"; got != want {
		t.Fatalf("unexpected proxied path: got %q want %q", got, want)
	}
}
