# \LogsLbDomainsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLbDomainsLogList**](LogsLbDomainsAPI.md#GetLbDomainsLogList) | **Get** /lb_domains/{LbDomainId}/logs | LBドメイン操作ログの一覧取得
[**GetLbDomainsLogListCount**](LogsLbDomainsAPI.md#GetLbDomainsLogListCount) | **Get** /lb_domains/{LbDomainId}/logs/count | LBドメイン操作ログの件数取得



## GetLbDomainsLogList

> GetLbDomainsLogs GetLbDomainsLogList(ctx, lbDomainId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

LBドメイン操作ログの一覧取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	lbDomainId := "lbDomainId_example" // string | GET lb_domains Schemaにおける id
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	order := openapiclient.SearchOrder("ASC") // SearchOrder | ソート順 (optional) (default to "ASC")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := []openapiclient.LbDomainsLogsType{openapiclient.LbDomainsLogsType("service")} // []LbDomainsLogsType |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := []openapiclient.LbDomainsLogsOperation{openapiclient.LbDomainsLogsOperation("update_lb_domain_description")} // []LbDomainsLogsOperation |  (optional) (default to [])
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := []openapiclient.LbDomainsLogsStatus{openapiclient.LbDomainsLogsStatus("start")} // []LbDomainsLogsStatus |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLbDomainsAPI.GetLbDomainsLogList(context.Background(), lbDomainId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLbDomainsAPI.GetLbDomainsLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainsLogList`: GetLbDomainsLogs
	fmt.Fprintf(os.Stdout, "Response from `LogsLbDomainsAPI.GetLbDomainsLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainsLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **startDate** | **string** | 開始日 | [default to &quot;&quot;]
 **endDate** | **string** | 終了日 | [default to &quot;&quot;]
 **timeZone** | **string** | タイムゾーン | [default to &quot;+00:00&quot;]
 **order** | [**SearchOrder**](SearchOrder.md) | ソート順 | [default to &quot;ASC&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsLogType** | [**[]LbDomainsLogsType**](LbDomainsLogsType.md) |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**[]LbDomainsLogsOperation**](LbDomainsLogsOperation.md) |  | [default to []]
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**[]LbDomainsLogsStatus**](LbDomainsLogsStatus.md) |  | [default to []]

### Return type

[**GetLbDomainsLogs**](GetLbDomainsLogs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLbDomainsLogListCount

> GetCount GetLbDomainsLogListCount(ctx, lbDomainId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

LBドメイン操作ログの件数取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	lbDomainId := "lbDomainId_example" // string | GET lb_domains Schemaにおける id
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := []openapiclient.LbDomainsLogsType{openapiclient.LbDomainsLogsType("service")} // []LbDomainsLogsType |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := []openapiclient.LbDomainsLogsOperation{openapiclient.LbDomainsLogsOperation("update_lb_domain_description")} // []LbDomainsLogsOperation |  (optional) (default to [])
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := []openapiclient.LbDomainsLogsStatus{openapiclient.LbDomainsLogsStatus("start")} // []LbDomainsLogsStatus |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLbDomainsAPI.GetLbDomainsLogListCount(context.Background(), lbDomainId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLbDomainsAPI.GetLbDomainsLogListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainsLogListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `LogsLbDomainsAPI.GetLbDomainsLogListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainsLogListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **startDate** | **string** | 開始日 | [default to &quot;&quot;]
 **endDate** | **string** | 終了日 | [default to &quot;&quot;]
 **timeZone** | **string** | タイムゾーン | [default to &quot;+00:00&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsLogType** | [**[]LbDomainsLogsType**](LbDomainsLogsType.md) |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**[]LbDomainsLogsOperation**](LbDomainsLogsOperation.md) |  | [default to []]
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**[]LbDomainsLogsStatus**](LbDomainsLogsStatus.md) |  | [default to []]

### Return type

[**GetCount**](GetCount.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

