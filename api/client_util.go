package api

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"golang.org/x/time/rate"
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

// unwrappable is implemented by transports that wrap another RoundTripper.
type unwrappable interface {
	Unwrap() http.RoundTripper
}

// findRateLimitTransport walks the transport chain and returns the first *rateLimitTransport found.
func findRateLimitTransport(t http.RoundTripper) *rateLimitTransport {
	for t != nil {
		if rl, ok := t.(*rateLimitTransport); ok {
			return rl
		}
		u, ok := t.(unwrappable)
		if !ok {
			return nil
		}
		t = u.Unwrap()
	}
	return nil
}

// rateLimitTransport is an http.RoundTripper that enforces a rate limit on requests.
type rateLimitTransport struct {
	base    http.RoundTripper
	limiter *rate.Limiter
}

func (t *rateLimitTransport) Unwrap() http.RoundTripper { return t.base }

func (t *rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}
	return t.base.RoundTrip(req)
}

// tokenProviderTransport is an http.RoundTripper that injects the Authorization header dynamically.
type tokenProviderTransport struct {
	base     http.RoundTripper
	provider TokenProvider
}

func (t *tokenProviderTransport) Unwrap() http.RoundTripper { return t.base }

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

// WithRateLimit sets the maximum number of requests per second.
// If NewDPFAPIClient's default rate limiter is already set, it is replaced.
func WithRateLimit(rps rate.Limit) ClientOption {
	return func(cfg *Configuration) {
		burst := 1
		if rps != rate.Inf && rps > 1 {
			burst = int(rps)
		}
		// Walk the chain to find and replace an existing rateLimitTransport
		if cfg.HTTPClient != nil {
			if rl := findRateLimitTransport(cfg.HTTPClient.Transport); rl != nil {
				rl.limiter = rate.NewLimiter(rps, burst)
				return
			}
		}
		// Otherwise wrap the existing transport
		if cfg.HTTPClient == nil {
			cfg.HTTPClient = &http.Client{}
		}
		base := cfg.HTTPClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		cfg.HTTPClient.Transport = &rateLimitTransport{
			base:    base,
			limiter: rate.NewLimiter(rps, burst),
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

// NewClientFromEnv は環境変数から設定を読み込み、APIClient を生成して返します。
// 認証トークンは DPF_API_TOKEN (トークン文字列) または DPF_API_TOKEN_FILE (トークンファイルパス) で指定します。
// どちらも未設定の場合はエラーを返します。
// オプションとして DPF_API_ENDPOINT でエンドポイント URL、DPF_API_USER_AGENT で User-Agent を指定できます。
func NewClientFromEnv() (*APIClient, error) {
	var tokenProvider TokenProvider
	if os.Getenv("DPF_API_TOKEN") != "" {
		tokenProvider = StringTokenProvider(os.Getenv("DPF_API_TOKEN"))
	} else if os.Getenv("DPF_API_TOKEN_FILE") != "" {
		tokenProvider = FileTokenProvider(os.Getenv("DPF_API_TOKEN_FILE"))
	} else {
		return nil, errors.New("DPF_API_TOKEN or DPF_API_TOKEN_FILE is required")
	}
	var opts []ClientOption
	if os.Getenv("DPF_API_ENDPOINT") != "" {
		opts = append(opts, WithEndpoint(os.Getenv("DPF_API_ENDPOINT")))
	}
	if os.Getenv("DPF_API_USER_AGENT") != "" {
		opts = append(opts, WithUserAgent(os.Getenv("DPF_API_USER_AGENT")))
	}
	return NewDPFAPIClient(tokenProvider, opts...), nil
}
