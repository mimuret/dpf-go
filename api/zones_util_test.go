package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetZoneByDomainName(t *testing.T) {
	zones := []Zone{
		{Id: "Z0000000000001", Name: "jp.", CommonConfigId: 1, ServiceCode: "dpfxxxxxxxxxx1", State: 1, Favorite: 1, Description: "", Network: *NewNullableString(nil)},
		{Id: "Z0000000000002", Name: "example.jp.", CommonConfigId: 1, ServiceCode: "dpfxxxxxxxxxx2", State: 1, Favorite: 1, Description: "", Network: *NewNullableString(nil)},
		{Id: "Z0000000000003", Name: "www.example.jp.", CommonConfigId: 1, ServiceCode: "dpfxxxxxxxxxx3", State: 1, Favorite: 1, Description: "", Network: *NewNullableString(nil)},
		{Id: "Z0000000000004", Name: "ex\\.ample.jp.", CommonConfigId: 1, ServiceCode: "dpfxxxxxxxxxx4", State: 1, Favorite: 1, Description: "", Network: *NewNullableString(nil)},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/zones" {
			// キーワードに含まれるゾーンをフィルタリングして返す（簡易実装）
			keywords := r.URL.Query()["_keywords_name[]"]
			var results []Zone
			for _, z := range zones {
				for _, k := range keywords {
					if z.Name == k {
						results = append(results, z)
						break
					}
				}
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(GetZones{
				RequestId: "test-req-id",
				Results:   results,
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
	client := NewAPIClient(cfg)

	tests := []struct {
		name       string
		domainName string
		ds         bool
		wantName   string
		wantErr    error
	}{
		{
			name:       "ロンゲストマッチ: サブドメインから最短の親",
			domainName: "sub.www.example.jp.",
			ds:         false,
			wantName:   "www.example.jp.",
		},
		{
			name:       "ロンゲストマッチ: 中間の親",
			domainName: "other.example.jp.",
			ds:         false,
			wantName:   "example.jp.",
		},
		{
			name:       "完全一致: ds=false",
			domainName: "example.jp.",
			ds:         false,
			wantName:   "example.jp.",
		},
		{
			name:       "完全一致除外: ds=true",
			domainName: "example.jp.",
			ds:         true,
			wantErr:    ErrZoneNotFound,
		},
		{
			name:       "完全一致除外: ds=true (サブドメイン)",
			domainName: "www.example.jp.",
			ds:         true,
			wantName:   "example.jp.",
		},
		{
			name:       "TLDマッチ (TLDは検索対象外のためマッチしない)",
			domainName: "test.jp.",
			ds:         false,
			wantErr:    ErrZoneNotFound,
		},
		{
			name:       "マッチなし",
			domainName: "example.com.",
			ds:         false,
			wantErr:    ErrZoneNotFound,
		},
		{
			name:       "エスケープされたドットを含むドメイン名",
			domainName: "www.ex\\.ample.jp.",
			ds:         false,
			wantName:   "ex\\.ample.jp.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zone, _, err := client.ZonesAPI.GetZoneByDomainName(context.Background(), tt.domainName, tt.ds)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if zone == nil {
				t.Fatal("expected result, got nil")
			}
			if zone.Name != tt.wantName {
				t.Errorf("expected %s, got %s", tt.wantName, zone.Name)
			}
		})
	}
}

func TestGetZoneByDomainName_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
	client := NewAPIClient(cfg)

	_, _, err := client.ZonesAPI.GetZoneByDomainName(context.Background(), "example.jp.", false)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
