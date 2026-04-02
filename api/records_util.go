// Package api provides a client for the DPF API.
//
// This package includes utilities for record management.
package api

import (
	"context"
	"errors"
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
//   - *Record: 見つかったレコード。
//   - *http.Response: APIリクエストの HTTP レスポンス。
//   - error: APIの呼び出しでエラーが発生した場合、またはレコードが見つからない場合 (ErrRecordNotFound)。
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

	return nil, httpResp, ErrRecordNotFound
}

// UpsertRecord は、指定されたレコード（name+rrtype）がすでに存在する場合は更新 (PatchRecord)、
// 存在しない場合には新規作成 (PostRecord) を行います。
//
// 引数:
//   - ctx: コンテキスト。
//   - zoneId: ゾーンのID。
//   - postRecord: 作成または更新するレコードの情報。
//
// 戻り値:
//   - *AsyncResponse: 非同期リクエストのレスポンス（ジョブURLを含む）。
//   - *http.Response: APIリクエストの HTTP レスポンス。
//   - error: APIの呼び出しでエラーが発生した場合。
func (a *RecordsAPIService) UpsertRecord(ctx context.Context, zoneId string, postRecord PostRecord) (asyncResponse *AsyncResponse, httpResp *http.Response, err error) {
	ctx, span := a.client.tracer.Start(ctx, "RecordsAPIService.UpsertRecord")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	// 既存のレコードを検索
	targetRecord, httpResp, err := a.GetRecordByName(ctx, zoneId, postRecord.Name, RecordsRrtype(postRecord.Rrtype))
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			// 存在しない場合は PostRecord で新規作成
			return a.PostRecord(ctx, zoneId).
				PostRecord(postRecord).
				Execute()
		}
		return nil, httpResp, err
	}

	// すでに存在する場合は PatchRecord で更新
	patchRecord := NewPatchRecord()
	if postRecord.HasDescription() {
		patchRecord.SetDescription(postRecord.GetDescription())
	}
	if postRecord.HasLabels() {
		patchRecord.SetLabels(postRecord.GetLabels())
	}
	patchRecord.SetRdata(postRecord.GetRdata())
	if postRecord.Ttl.IsSet() {
		if v := postRecord.Ttl.Get(); v != nil {
			patchRecord.SetTtl(*v)
		} else {
			patchRecord.SetTtlNil()
		}
	}

	return a.PatchRecord(ctx, zoneId, targetRecord.Id).
		PatchRecord(*patchRecord).
		Execute()
}
