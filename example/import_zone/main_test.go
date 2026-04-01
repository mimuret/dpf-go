package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mimuret/dpf-go/api"
)

func TestImportZone(t *testing.T) {
	zoneId := "Z0123456789012"
	zoneName := "example.jp."
	reqID := "782d746ac3cb46499b31708fa80e8660"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// 1. PatchZoneAtomicChanges
		if r.Method == http.MethodPatch && r.URL.Path == "/zones/"+zoneId+"/atomic_changes" {
			res := api.AsyncResponse{
				RequestId: reqID,
				JobsUrl:   "http://localhost/jobs/" + reqID,
			}
			w.WriteHeader(http.StatusAccepted)
			_ = json.NewEncoder(w).Encode(res)
			return
		}

		// 2. SyncWait Polling
		if r.Method == http.MethodGet && r.URL.Path == "/jobs/"+reqID {
			res := api.GetJobs{
				RequestId: reqID,
				Status:    "SUCCESSFUL",
			}
			_ = json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := api.NewConfiguration()
	cfg.Servers = api.ServerConfigurations{{URL: ts.URL}}
	client := api.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, "test-token")

	zoneFile := "test_import.zone"
	err := os.WriteFile(zoneFile, []byte("www.example.jp. 3600 IN A 192.168.0.1"), 0644)
	if err != nil {
		t.Fatalf("failed to create test zone file: %v", err)
	}
	defer func() { _ = os.Remove(zoneFile) }()

	if err := RunImport(ctx, client, zoneId, zoneName, zoneFile); err != nil {
		t.Fatalf("RunImport failed: %v", err)
	}
}
