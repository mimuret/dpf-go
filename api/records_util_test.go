package api

import (
	"context"
	"encoding/json"
	"errors"
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

			if tt.expectNil {
				if !errors.Is(err, ErrRecordNotFound) {
					t.Fatalf("expected ErrRecordNotFound, got %v", err)
				}
				if record != nil {
					t.Errorf("expected nil result, got %s", record.Id)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
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

func TestUpsertRecord(t *testing.T) {
	zoneID := "Z0000000000001"
	recordID := "R0000000000001"
	existingRecord := Record{Id: recordID, Name: "www.example.jp.", Rrtype: RECORDSRRTYPE_A}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-Id", "test-req-id")

		// GetRecordByName (GetRecordList)
		if r.Method == http.MethodGet && r.URL.Path == "/zones/"+zoneID+"/records" {
			name := r.URL.Query().Get("_keywords_name[]")
			var results []Record
			if name == "www.example.jp." {
				results = append(results, existingRecord)
			}
			_ = json.NewEncoder(w).Encode(GetRecords{
				RequestId: "test-req-id",
				Results:   results,
			})
			return
		}

		// PostRecord
		if r.Method == http.MethodPost && r.URL.Path == "/zones/"+zoneID+"/records" {
			var post PostRecord
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				t.Errorf("failed to decode post body: %v", err)
			}
			_ = json.NewEncoder(w).Encode(AsyncResponse{
				RequestId: "test-req-id",
				JobsUrl:   "https://api.dns-platform.jp/dpf/v1/jobs/test-req-id",
			})
			return
		}

		// PatchRecord
		if r.Method == http.MethodPatch && r.URL.Path == "/zones/"+zoneID+"/records/"+recordID {
			var patch PatchRecord
			if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
				t.Errorf("failed to decode patch body: %v", err)
			}
			_ = json.NewEncoder(w).Encode(AsyncResponse{
				RequestId: "test-req-id",
				JobsUrl:   "https://api.dns-platform.jp/dpf/v1/jobs/test-req-id",
			})
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: ts.URL}}
	client := NewAPIClient(cfg)

	t.Run("新規作成 (PostRecord)", func(t *testing.T) {
		post := NewPostRecord("new.example.jp.", []RecordsRdataInner{{Value: PtrString("192.0.2.1")}}, RECORDSRRTYPEWITHOUTSOA_A)
		res, _, err := client.RecordsAPI.UpsertRecord(context.Background(), zoneID, *post)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.RequestId != "test-req-id" {
			t.Errorf("expected RequestId test-req-id, got %s", res.RequestId)
		}
	})

	t.Run("更新 (PatchRecord)", func(t *testing.T) {
		post := NewPostRecord("www.example.jp.", []RecordsRdataInner{{Value: PtrString("192.0.2.2")}}, RECORDSRRTYPEWITHOUTSOA_A)
		res, _, err := client.RecordsAPI.UpsertRecord(context.Background(), zoneID, *post)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.RequestId != "test-req-id" {
			t.Errorf("expected RequestId test-req-id, got %s", res.RequestId)
		}
	})
}
