// Package api provides a client for the DPF API.
//
// This package includes utilities for record management.
package api

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/codes"

	"github.com/miekg/dns"
)

// GetRecordByName は、指定されたレコード名とRRTYPEに完全一致するレコードを取得します。
//
// APIの検索 (KeywordsName) は部分一致も含めて取得されるため、この関数内で完全一致のフィルタリングを行います。
//
// 引数:
//   - ctx: コンテキスト。
//   - zoneId: ゾーンのID。
//   - targetDomainName: レコード名（例: "www.example.jp."）。
//   - rrtype: レコードタイプ（例: RECORDSRRTYPE_A）。
//
// 戻り値:
//   - *Record: 見つかったレコード。見つからない場合は nil。
//   - *http.Response: APIリクエストの HTTP レスポンス。
//   - error: APIの呼び出しでエラーが発生した場合。
func (a *RecordsAPIService) GetRecordByName(ctx context.Context, zoneId string, targetDomainName string, rrtype RecordsRrtype) (targetRecord *Record, httpResp *http.Response, err error) {
	ctx, span := a.client.tracer.Start(ctx, "RecordsAPIService.GetRecordByName")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	targetDomainName = dns.CanonicalName(targetDomainName)

	req := a.GetRecordList(ctx, zoneId).
		KeywordsName([]string{targetDomainName}).
		KeywordsRrtype(rrtype)

	res, httpResp, err := req.ExecuteAll()
	if err != nil {
		return nil, httpResp, err
	}

	if res != nil {
		for i := range res.Results {
			r := &res.Results[i]
			if r.Name == targetDomainName && r.Rrtype == rrtype {
				return r, httpResp, nil
			}
		}
	}

	return nil, httpResp, nil
}
