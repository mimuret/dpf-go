package api

import (
	"net/http"
	"os"
	"strings"
)

// TokenProvider is an interface that provides the current access token.
// Implementations can support static tokens, file-based tokens,
// or background-renewed tokens (e.g., HashiCorp Vault).
type TokenProvider interface {
	Token() (string, error)
}

// StringTokenProvider is a TokenProvider that always returns a static token.
type StringTokenProvider string

func (s StringTokenProvider) Token() (string, error) {
	return string(s), nil
}

// FileTokenProvider is a TokenProvider that reads the token from a file on each call.
type FileTokenProvider string

func (f FileTokenProvider) Token() (string, error) {
	data, err := os.ReadFile(string(f))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// tokenProviderTransport is an http.RoundTripper that injects the Authorization header dynamically.
type tokenProviderTransport struct {
	base     http.RoundTripper
	provider TokenProvider
}

func (t *tokenProviderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.provider.Token()
	if err != nil {
		return nil, err
	}
	req = req.Clone(req.Context())
	req.Header.Set("Authorization", "Bearer "+token)
	return t.base.RoundTrip(req)
}

// ClientOption is a functional option for NewClient.
type ClientOption func(*Configuration)

// WithEndpoint sets the API endpoint URL.
func WithEndpoint(endpoint string) ClientOption {
	return func(cfg *Configuration) {
		cfg.Servers = ServerConfigurations{
			{
				URL:         endpoint,
				Description: "API endpoint",
			},
		}
	}
}

// WithUserAgent sets the User-Agent header.
func WithUserAgent(userAgent string) ClientOption {
	return func(cfg *Configuration) {
		cfg.UserAgent = userAgent
	}
}

// NewDPFAPIClient creates a new APIClient with the given TokenProvider and optional configuration.
// The token is injected via the Authorization header on every request.
func NewDPFAPIClient(provider TokenProvider, opts ...ClientOption) *APIClient {
	cfg := NewConfiguration()
	cfg.UserAgent = "dpf-go-client"
	cfg.HTTPClient = &http.Client{
		Transport: &tokenProviderTransport{
			base:     http.DefaultTransport,
			provider: provider,
		},
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return NewAPIClient(cfg)
}
