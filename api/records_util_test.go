package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRecordByName(t *testing.T) {
	zoneID := "Z0000000000001"
	records := []Record{
		{Id: "R0000000000001", Name: "example.jp.", Rrtype: RECORDSRRTYPE_A},
		{Id: "R0000000000002", Name: "www.example.jp.", Rrtype: RECORDSRRTYPE_A},
		{Id: "R0000000000003", Name: "www.example.jp.", Rrtype: RECORDSRRTYPE_CNAME},
		{Id: "R0000000000004", Name: "ww.example.jp.", Rrtype: RECORDSRRTYPE_A},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/zones/" + zoneID + "/records"
		if r.URL.Path == expectedPath {
			// APIのKeywordsNameは部分一致としてシミュレート
			keywords := r.URL.Query()["_keywords_name[]"]
			rrtype := r.URL.Query().Get("_keywords_rrtype[]")

			var results []Record
			for _, rec := range records {
				nameMatch := false
				if len(keywords) == 0 {
					nameMatch = true
				} else {
					for _, k := range keywords {
						if strings.Contains(rec.Name, k) {
							nameMatch = true
							break
						}
					}
				}

				rrtypeMatch := rrtype == "" || string(rec.Rrtype) == rrtype

				if nameMatch && rrtypeMatch {
					results = append(results, rec)
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Request-Id", "test-req-id")
			_ = json.NewEncoder(w).Encode(GetRecords{
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
		name             string
		targetDomainName string
		rrtype           RecordsRrtype
		wantID           string
		expectNil        bool
	}{
		{
			name:             "完全一致: Aレコード",
			targetDomainName: "www.example.jp.",
			rrtype:           RECORDSRRTYPE_A,
			wantID:           "R0000000000002",
		},
		{
			name:             "完全一致: CNAMEレコード",
			targetDomainName: "www.example.jp.",
			rrtype:           RECORDSRRTYPE_CNAME,
			wantID:           "R0000000000003",
		},
		{
			name:             "ドットなし指定でも正規化されて一致する",
			targetDomainName: "example.jp",
			rrtype:           RECORDSRRTYPE_A,
			wantID:           "R0000000000001",
		},
		{
			name:             "部分一致するが完全一致がない場合",
			targetDomainName: "w.example.jp.",
			rrtype:           RECORDSRRTYPE_A,
			expectNil:        true,
		},
		{
			name:             "存在しないレコード",
			targetDomainName: "none.example.jp.",
			rrtype:           RECORDSRRTYPE_A,
			expectNil:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record, _, err := client.RecordsAPI.GetRecordByName(context.Background(), zoneID, tt.targetDomainName, tt.rrtype)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.expectNil {
				if record != nil {
					t.Errorf("expected nil result, got %s", record.Id)
				}
				return
			}

			if record == nil {
				t.Fatal("expected result, got nil")
			}
			if record.Id != tt.wantID {
				t.Errorf("expected ID %s, got %s", tt.wantID, record.Id)
			}
		})
	}
}

func TestGetRecordByName_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
	client := NewAPIClient(cfg)

	_, _, err := client.RecordsAPI.GetRecordByName(context.Background(), "Z0000000000001", "example.jp.", RECORDSRRTYPE_A)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
