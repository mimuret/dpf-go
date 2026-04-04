package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

// -- StringTokenProvider --

func TestStringTokenProvider(t *testing.T) {
	p := StringTokenProvider("my-token")
	got, err := p.Token()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "my-token" {
		t.Errorf("got %q, want %q", got, "my-token")
	}
}

// -- FileTokenProvider --

func TestFileTokenProvider(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "token")
	if err := os.WriteFile(path, []byte("  file-token\n"), 0600); err != nil {
		t.Fatal(err)
	}

	p := FileTokenProvider(path)
	got, err := p.Token()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "file-token" {
		t.Errorf("got %q, want %q", got, "file-token")
	}
}

func TestFileTokenProvider_ReadsOnEachCall(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "token")
	if err := os.WriteFile(path, []byte("token-v1"), 0600); err != nil {
		t.Fatal(err)
	}

	p := FileTokenProvider(path)
	got, _ := p.Token()
	if got != "token-v1" {
		t.Errorf("got %q, want %q", got, "token-v1")
	}

	if err := os.WriteFile(path, []byte("token-v2"), 0600); err != nil {
		t.Fatal(err)
	}
	got, _ = p.Token()
	if got != "token-v2" {
		t.Errorf("got %q, want %q", got, "token-v2")
	}
}

func TestFileTokenProvider_FileNotFound(t *testing.T) {
	p := FileTokenProvider("/nonexistent/path/token")
	_, err := p.Token()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

// -- tokenProviderTransport --

func TestTokenProviderTransport_InjectsAuthHeader(t *testing.T) {
	var gotAuth string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	tr := &tokenProviderTransport{
		base:     http.DefaultTransport,
		provider: StringTokenProvider("test-token"),
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	if _, err := client.Do(req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotAuth != "Bearer test-token" {
		t.Errorf("got %q, want %q", gotAuth, "Bearer test-token")
	}
}

func TestTokenProviderTransport_ProviderError(t *testing.T) {
	tr := &tokenProviderTransport{
		base:     http.DefaultTransport,
		provider: &errorTokenProvider{err: errors.New("token error")},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost", nil)
	if _, err := client.Do(req); err == nil {
		t.Error("expected error, got nil")
	}
}

type errorTokenProvider struct{ err error }

func (e *errorTokenProvider) Token() (string, error) { return "", e.err }

// -- rateLimitTransport --

func TestRateLimitTransport_EnforcesLimit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// 1 rps: 3リクエストが2秒以上かかることを確認
	tr := &rateLimitTransport{
		base:    http.DefaultTransport,
		limiter: rate.NewLimiter(1, 1),
	}
	client := &http.Client{Transport: tr}

	start := time.Now()
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
		if _, err := client.Do(req); err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
	}
	if elapsed := time.Since(start); elapsed < 2*time.Second {
		t.Errorf("3 requests at 1rps should take >= 2s, took %v", elapsed)
	}
}

func TestRateLimitTransport_ContextCancel(t *testing.T) {
	tr := &rateLimitTransport{
		base:    http.DefaultTransport,
		limiter: rate.NewLimiter(rate.Every(10*time.Second), 1),
	}
	// バーストを使い切る
	tr.limiter.Allow()

	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	req, _ := http.NewRequestWithContext(cancelCtx, http.MethodGet, "http://localhost", nil)
	client := &http.Client{Transport: tr}
	if _, err := client.Do(req); err == nil {
		t.Error("expected error from cancelled context, got nil")
	}
}

// -- findRateLimitTransport --

func TestFindRateLimitTransport_Direct(t *testing.T) {
	rl := &rateLimitTransport{base: http.DefaultTransport, limiter: rate.NewLimiter(1, 1)}
	if findRateLimitTransport(rl) != rl {
		t.Error("should find rateLimitTransport directly")
	}
}

func TestFindRateLimitTransport_Nested(t *testing.T) {
	rl := &rateLimitTransport{base: http.DefaultTransport, limiter: rate.NewLimiter(1, 1)}
	outer := &tokenProviderTransport{base: rl, provider: StringTokenProvider("t")}
	if findRateLimitTransport(outer) != rl {
		t.Error("should find rateLimitTransport through tokenProviderTransport")
	}
}

func TestFindRateLimitTransport_NotFound(t *testing.T) {
	tr := &tokenProviderTransport{base: http.DefaultTransport, provider: StringTokenProvider("t")}
	if findRateLimitTransport(tr) != nil {
		t.Error("should return nil when rateLimitTransport is not in chain")
	}
}

// -- WithRateLimit --

func TestWithRateLimit_AddsRateLimiter(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{Transport: http.DefaultTransport}
	WithRateLimit(5)(cfg)

	rl := findRateLimitTransport(cfg.HTTPClient.Transport)
	if rl == nil {
		t.Fatal("rateLimitTransport not found")
	}
	if rl.limiter.Limit() != 5 {
		t.Errorf("got limit %v, want 5", rl.limiter.Limit())
	}
}

func TestWithRateLimit_ReplacesExisting(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &rateLimitTransport{
			base:    http.DefaultTransport,
			limiter: rate.NewLimiter(1, 1),
		},
	}
	WithRateLimit(10)(cfg)

	rl := findRateLimitTransport(cfg.HTTPClient.Transport)
	if rl == nil {
		t.Fatal("rateLimitTransport not found")
	}
	if rl.limiter.Limit() != 10 {
		t.Errorf("got limit %v, want 10", rl.limiter.Limit())
	}
}

func TestWithRateLimit_WorksThroughCustomTransport(t *testing.T) {
	// カスタム transport でラップされた下の rateLimitTransport も差し替えられる
	rl := &rateLimitTransport{base: http.DefaultTransport, limiter: rate.NewLimiter(1, 1)}
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &tokenProviderTransport{
			base:     rl,
			provider: StringTokenProvider("wrap"),
		},
	}
	WithRateLimit(10)(cfg)

	found := findRateLimitTransport(cfg.HTTPClient.Transport)
	if found == nil {
		t.Fatal("rateLimitTransport not found")
	}
	if found.limiter.Limit() != 10 {
		t.Errorf("got limit %v, want 10", found.limiter.Limit())
	}
}

// -- NewDPFAPIClient --

func TestNewDPFAPIClient_NoDefaultRateLimit(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &tokenProviderTransport{
			base:     http.DefaultTransport,
			provider: StringTokenProvider("tok"),
		},
	}
	if findRateLimitTransport(cfg.HTTPClient.Transport) != nil {
		t.Error("should not have rateLimitTransport by default")
	}
}

func TestNewDPFAPIClient_DefaultUserAgent(t *testing.T) {
	client := NewDPFAPIClient(StringTokenProvider("tok"))
	if client.cfg.UserAgent != "dpf-go-client" {
		t.Errorf("got %q, want %q", client.cfg.UserAgent, "dpf-go-client")
	}
}

func TestNewDPFAPIClient_WithEndpoint(t *testing.T) {
	client := NewDPFAPIClient(StringTokenProvider("tok"), WithEndpoint("https://example.com"))
	if len(client.cfg.Servers) == 0 || client.cfg.Servers[0].URL != "https://example.com" {
		t.Errorf("endpoint not set correctly: %v", client.cfg.Servers)
	}
}

func TestNewDPFAPIClient_WithUserAgent(t *testing.T) {
	client := NewDPFAPIClient(StringTokenProvider("tok"), WithUserAgent("myapp/1.0"))
	if client.cfg.UserAgent != "myapp/1.0" {
		t.Errorf("got %q, want %q", client.cfg.UserAgent, "myapp/1.0")
	}
}

func TestNewDPFAPIClient_InjectsToken(t *testing.T) {
	var gotAuth string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewDPFAPIClient(StringTokenProvider("my-token"), WithEndpoint(ts.URL))
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	if _, err := client.cfg.HTTPClient.Do(req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotAuth != "Bearer my-token" {
		t.Errorf("got %q, want %q", gotAuth, "Bearer my-token")
	}
}

func TestNewDPFAPIClient_WithRateLimit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewDPFAPIClient(StringTokenProvider("tok"),
		WithEndpoint(ts.URL),
		WithRateLimit(1),
	)

	start := time.Now()
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
		if _, err := client.cfg.HTTPClient.Do(req); err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
	}
	if elapsed := time.Since(start); elapsed < 2*time.Second {
		t.Errorf("3 requests at 1rps should take >= 2s, took %v", elapsed)
	}
}

// -- NewClientFromEnv --

func TestNewClientFromEnv_Token(t *testing.T) {
	t.Setenv("DPF_API_TOKEN", "env-token")
	t.Setenv("DPF_API_TOKEN_FILE", "")
	t.Setenv("DPF_API_ENDPOINT", "")
	t.Setenv("DPF_API_USER_AGENT", "")

	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
}

func TestNewClientFromEnv_TokenFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "token")
	if err := os.WriteFile(path, []byte("file-token"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DPF_API_TOKEN", "")
	t.Setenv("DPF_API_TOKEN_FILE", path)

	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
}

func TestNewClientFromEnv_NoToken(t *testing.T) {
	t.Setenv("DPF_API_TOKEN", "")
	t.Setenv("DPF_API_TOKEN_FILE", "")

	if _, err := NewClientFromEnv(); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewClientFromEnv_WithEndpointAndUserAgent(t *testing.T) {
	t.Setenv("DPF_API_TOKEN", "tok")
	t.Setenv("DPF_API_ENDPOINT", "https://custom.example.com")
	t.Setenv("DPF_API_USER_AGENT", "custom-agent")

	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client.cfg.Servers[0].URL != "https://custom.example.com" {
		t.Errorf("endpoint not set: %v", client.cfg.Servers[0].URL)
	}
	if client.cfg.UserAgent != "custom-agent" {
		t.Errorf("user agent not set: %v", client.cfg.UserAgent)
	}
}
