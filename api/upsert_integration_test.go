package api

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestIntegrationUpsertRecord(t *testing.T) {
	token := os.Getenv("DPF_TOKEN")
	zoneId := os.Getenv("DPF_ZONE_ID")
	if token == "" || zoneId == "" {
		t.Skip("DPF_TOKEN and DPF_ZONE_ID are required for integration tests")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAccessToken, token)

	// テスト用のレコード名。既存のレコードと被らないように注意
	name := "upsert-it-test.example.jp."

	t.Run("新規作成および更新", func(t *testing.T) {
		// 1. 新規作成を意識した Upsert
		postBody := NewPostRecord(name, []RecordsRdataInner{{Value: PtrString("192.0.2.1")}}, RECORDSRRTYPEWITHOUTSOA_A)
		postBody.SetDescription("Initial creation")

		res, resp, err := client.RecordsAPI.UpsertRecord(ctx, zoneId, *postBody)
		if err != nil {
			t.Fatalf("failed to upsert record (post): %v", err)
		}
		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("expected 202 Accepted, got %d", resp.StatusCode)
		}

		// ジョブの完了を待機
		_, _, err = client.JobsAPI.SyncWait(res, resp, err)
		if err != nil {
			t.Fatalf("SyncWait for creation failed: %v", err)
		}

		// 作成されたことを確認
		found, _, err := client.RecordsAPI.GetRecordByName(ctx, zoneId, name, RECORDSRRTYPE_A)
		if err != nil {
			t.Fatalf("failed to get created record: %v", err)
		}
		if found == nil {
			t.Fatal("record not found after creation")
		}
		if found.Description != "Initial creation" {
			t.Errorf("expected description 'Initial creation', got '%s'", found.Description)
		}

		// 2. 更新 (Patch) を意識した Upsert
		postBody.SetDescription("Updated description")
		postBody.Rdata = []RecordsRdataInner{{Value: PtrString("192.0.2.2")}}

		res, resp, err = client.RecordsAPI.UpsertRecord(ctx, zoneId, *postBody)
		if err != nil {
			t.Fatalf("failed to upsert record (patch): %v", err)
		}
		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("expected 202 Accepted, got %d", resp.StatusCode)
		}

		// ジョブの完了を待機
		_, _, err = client.JobsAPI.SyncWait(res, resp, err)
		if err != nil {
			t.Fatalf("SyncWait for update failed: %v", err)
		}

		// 更新されたことを確認
		found, _, err = client.RecordsAPI.GetRecordByName(ctx, zoneId, name, RECORDSRRTYPE_A)
		if err != nil {
			t.Fatalf("failed to get updated record: %v", err)
		}
		if found == nil {
			t.Fatal("record missing after update")
		}
		if found.Description != "Updated description" {
			t.Errorf("expected description 'Updated description', got '%s'", found.Description)
		}
		if len(found.Rdata) != 1 || found.Rdata[0].GetValue() != "192.0.2.2" {
			t.Errorf("expected Rdata 192.0.2.2, got %v", found.Rdata)
		}

		// 3. 後始末 (レコードの削除)
		delRes, delResp, err := client.RecordsAPI.DeleteRecord(ctx, zoneId, found.Id).Execute()
		if err != nil {
			t.Errorf("failed to delete record: %v", err)
		} else {
			_, _, _ = client.JobsAPI.SyncWait(delRes, delResp, nil)
		}
	})
}
