package api

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/codes"

	"github.com/miekg/dns"
)

// GetZoneByDomainName は、指定されたドメイン名に対応するゾーンを取得します。
// domainName をベースに親ドメインを辿り、ロンゲストマッチするゾーンを返します。
// ds フラグが true の場合は、完全一致するドメイン名のゾーンを除外します。
// APIエラーの場合は error を返し、ゾーンが見つからない場合は ErrZoneNotFound を返します。
func (a *ZonesAPIService) GetZoneByDomainName(ctx context.Context, domainName string, ds bool) (bestMatch *Zone, httpResp *http.Response, err error) {
	ctx, span := a.client.tracer.Start(ctx, "ZonesAPIService.GetZoneByDomainName")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	domainName = dns.CanonicalName(domainName)

	offsets := dns.Split(domainName)
	// offsets for "www.example.jp." -> [0, 4, 12]
	// parents: ["www.example.jp.", "example.jp."]
	var parents []string
	for i := 0; i < len(offsets)-1; i++ {
		parents = append(parents, domainName[offsets[i]:])
	}

	req := a.GetZoneList(ctx).KeywordsName(parents)
	res, httpResp, err := req.ExecuteAll()
	if err != nil {
		return nil, httpResp, err
	}

	if res != nil {
		for i := range res.Results {
			z := &res.Results[i]
			if ds && z.Name == domainName {
				continue
			}

			if dns.IsSubDomain(z.Name, domainName) {
				if bestMatch == nil || len(z.Name) > len(bestMatch.Name) {
					bestMatch = z
				}
			}
		}
	}

	if bestMatch == nil {
		return nil, httpResp, ErrZoneNotFound
	}

	return bestMatch, httpResp, nil
}
