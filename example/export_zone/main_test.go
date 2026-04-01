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

func TestExportZone(t *testing.T) {
	zoneId := "Z0123456789012"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet && r.URL.Path == "/zones/"+zoneId+"/records/currents" {
			res := api.GetRecords{
				Results: []api.Record{
					{
						Name:   "www.example.jp.",
						Rrtype: "A",
						Ttl:    *api.NewNullableInt32(api.PtrInt32(3600)),
						Rdata: []api.RecordsRdataInner{
							{Value: api.PtrString("192.168.0.1")},
						},
					},
				},
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

	outputFile := "test_export.zone"
	defer func() { _ = os.Remove(outputFile) }()

	if err := RunExport(ctx, client, zoneId, outputFile); err != nil {
		t.Fatalf("RunExport failed: %v", err)
	}

	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	if len(content) == 0 {
		t.Error("output file is empty")
	}
}
