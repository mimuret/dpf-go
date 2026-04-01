package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miekg/dns"
	"github.com/mimuret/dpf-go/api"
)

func RunImport(ctx context.Context, client *api.APIClient, zoneId string, zoneName string, zoneFile string) error {
	f, err := os.Open(zoneFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	zp := dns.NewZoneParser(f, zoneName, zoneFile)

	type recordKey struct {
		name   string
		rrtype api.RecordsRrtype
	}
	groups := make(map[recordKey][]string)
	ttls := make(map[recordKey]int32)

	for rr, ok := zp.Next(); ok; rr, ok = zp.Next() {
		header := rr.Header()
		key := recordKey{
			name:   dns.CanonicalName(header.Name),
			rrtype: api.RecordsRrtype(dns.TypeToString[header.Rrtype]),
		}

		parts := strings.SplitN(rr.String(), "\t", 5)
		if len(parts) < 5 {
			parts = strings.Fields(rr.String())
			if len(parts) < 5 {
				continue
			}
			groups[key] = append(groups[key], strings.Join(parts[4:], " "))
		} else {
			groups[key] = append(groups[key], parts[4])
		}
		ttls[key] = int32(header.Ttl)
	}

	if err := zp.Err(); err != nil {
		return fmt.Errorf("failed to parse zone file: %w", err)
	}

	var overwriteRecords []api.OverwriteRecordsInner
	for key, values := range groups {
		rdatas := make([]api.RecordsRdataInner, len(values))
		for i, v := range values {
			rd := api.NewRecordsRdataInner()
			rd.SetValue(v)
			rdatas[i] = *rd
		}

		ttlValue := ttls[key]
		ttl := api.NewNullableInt32(&ttlValue)
		overwriteRecords = append(overwriteRecords, *api.NewOverwriteRecordsInner(
			"", // description
			nil, // labels
			key.name,
			rdatas,
			key.rrtype,
			*ttl,
		))
	}

	patch := api.NewPatchZoneAtomicChanges(overwriteRecords)
	patch.SetOverwriteSoa(true)
	patch.SetOverwriteZoneApexNs(true)

	// API実行
	asyncRes, resp, err := client.ZonesAPI.PatchZoneAtomicChanges(ctx, zoneId).
		PatchZoneAtomicChanges(*patch).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to apply atomic changes: %w", err)
	}

	fmt.Printf("Atomic patch request accepted. Request ID: %s\n", asyncRes.GetRequestId())

	// ジョブ完了待機
	fmt.Println("Waiting for job completion...")
	job, _, err := client.JobsAPI.SyncWait(asyncRes, resp, nil)
	if err != nil {
		return fmt.Errorf("job failed: %w", err)
	}

	fmt.Printf("Job finished with status: %s\n", job.GetStatus())
	return nil
}

func main() {
	token := os.Getenv("DPF_TOKEN")
	zoneId := os.Getenv("DPF_ZONE_ID")
	zoneName := os.Getenv("DPF_ZONE_NAME")
	zoneFile := os.Getenv("ZONE_FILE")

	if token == "" || zoneId == "" || zoneName == "" || zoneFile == "" {
		log.Fatal("DPF_TOKEN, DPF_ZONE_ID, DPF_ZONE_NAME, and ZONE_FILE environment variables are required")
	}

	cfg := api.NewConfiguration()
	client := api.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)

	if err := RunImport(ctx, client, zoneId, zoneName, zoneFile); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Records imported from %s successfully\n", zoneFile)
}
