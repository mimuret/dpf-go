# dpf-go

`dpf-go` は IIJ DNSプラットフォーム（DPF）API のための Go 言語用クライアントライブラリです。

## 特徴

- **最新の API サポート**: OpenAPI Generator を用いて最新の API 定義から自動生成されています。
- **自動ページネーション**: `ExecuteAll()` メソッドにより、大量のレコードやゾーン情報を自動的にページネーションして全件取得できます。
- **堅牢なエラーハンドリング**: 詳細なエラー情報のパースをサポートしており、トラブルシューティングを容易にします。
- **便利なユーティリティ**:
  - `GetRecordByName`: ドメイン名とタイプによる正確なレコード検索。
  - `GetZoneByDomainName`: 最適なゾーン（最長一致）の自動検索。
  - `SyncWait`: 非同期ジョブの完了待機。
- **OpenTelemetry による可観測性**: 標準でトレーシングをサポートしており、API 呼び出しの追跡が可能です。

## インストール

```bash
go get github.com/mimuret/dpf-go
```

## クイックスタート

以下のコードは、API を使用してゾーンの一覧を取得する簡単な例です。

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mimuret/dpf-go/api"
)

func main() {
	// 設定とクライアントの初期化
	cfg := api.NewConfiguration()
	client := api.NewAPIClient(cfg)

	// IIJ ID アクセストークンの設定
	token := os.Getenv("DPF_API_TOKEN")
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)

	// 全てのゾーンを取得 (ExecuteAll はページネーションを自動で行います)
	zones, _, err := client.ZonesAPI.GetZoneList(ctx).ExecuteAll()
	if err != nil {
		log.Fatalf("failed to get zones: %v", err)
	}

	for _, zone := range zones.Results {
		fmt.Printf("Zone: %s (ID: %s)\n", zone.Name, zone.Id)
	}
}
```

## 便利なユーティリティの利用

### ジョブの実行完了を待機する (`SyncWait`)

DPF API の更新系リクエストは非同期で行われます。`SyncWait` を使うことで、ジョブの完了（成功または失敗）までポーリングを伴う待機が可能です。

```go
res, resp, err := client.RecordsAPI.PatchZoneAtomicChanges(ctx, zoneId).
    PatchZoneAtomicChangesRequest(payload).
    Execute()

// ジョブの完了を待機
job, resp, err := client.JobsAPI.SyncWait(ctx, res, resp, err)
if err != nil {
    log.Fatalf("job failed: %v", err)
}
fmt.Printf("Job status: %s\n", job.Status)
```

## 開発者向け

### コードの生成

API 定義 (`api/swagger.json`) からコードを再生成するには、`Makefile` を使用します。これには Docker と `openapi-generator-cli` が必要です。

```bash
make generate
```
