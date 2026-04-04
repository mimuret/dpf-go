package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/miekg/dns"
	"github.com/mimuret/dpf-go/api"
)

func RunExport(ctx context.Context, client *api.APIClient, zoneId string, outputFile string) error {
	// 全レコードを取得 (ExecuteAllを使用)
	res, _, err := client.RecordsAPI.GetRecordCurrents(ctx, zoneId).ExecuteAll()
	if err != nil {
		return fmt.Errorf("failed to get records: %w", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() { _ = f.Close() }()

	for _, record := range res.Results {
		ttl := uint32(record.GetTtl())
		for _, rd := range record.Rdata {
			rrStr := fmt.Sprintf("%s %d IN %s %s", record.Name, ttl, record.Rrtype, rd.GetValue())
			rr, err := dns.NewRR(rrStr)
			if err != nil {
				log.Printf("Warning: failed to parse RR %s: %v", rrStr, err)
				continue
			}
			if _, err := fmt.Fprintln(f, rr.String()); err != nil {
				return fmt.Errorf("failed to write record: %w", err)
			}
		}
	}
	return nil
}

func main() {
	token := os.Getenv("DPF_API_TOKEN")
	zoneId := os.Getenv("DPF_ZONE_ID")
	outputFile := os.Getenv("OUTPUT_FILE")

	if token == "" || zoneId == "" || outputFile == "" {
		log.Fatal("DPF_API_TOKEN, DPF_ZONE_ID, and OUTPUT_FILE environment variables are required")
	}

	cfg := api.NewConfiguration()
	client := api.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)

	if err := RunExport(ctx, client, zoneId, outputFile); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Records exported to %s successfully\n", outputFile)
}
