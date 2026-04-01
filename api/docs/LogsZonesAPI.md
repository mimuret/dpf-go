# \LogsZonesAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZonesLogList**](LogsZonesAPI.md#GetZonesLogList) | **Get** /zones/{ZoneId}/logs | ゾーン操作ログの一覧取得
[**GetZonesLogListCount**](LogsZonesAPI.md#GetZonesLogListCount) | **Get** /zones/{ZoneId}/logs/count | ゾーン操作ログの件数取得



## GetZonesLogList

> GetZonesLogs GetZonesLogList(ctx, zoneId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

ゾーン操作ログの一覧取得



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
	zoneId := "zoneId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	order := openapiclient.SearchOrder("ASC") // SearchOrder | ソート順 (optional) (default to "ASC")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := openapiclient.ZonesLogsType("service") // ZonesLogsType |  (optional)
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := openapiclient.ZonesLogsOperation("updating_default_ttl") // ZonesLogsOperation |  (optional)
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := openapiclient.ZonesLogsStatus("start") // ZonesLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsZonesAPI.GetZonesLogList(context.Background(), zoneId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsZonesAPI.GetZonesLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZonesLogList`: GetZonesLogs
	fmt.Fprintf(os.Stdout, "Response from `LogsZonesAPI.GetZonesLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesLogListRequest struct via the builder pattern


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
 **keywordsLogType** | [**ZonesLogsType**](ZonesLogsType.md) |  | 
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**ZonesLogsOperation**](ZonesLogsOperation.md) |  | 
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**ZonesLogsStatus**](ZonesLogsStatus.md) |  | 

### Return type

[**GetZonesLogs**](GetZonesLogs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZonesLogListCount

> GetCount GetZonesLogListCount(ctx, zoneId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

ゾーン操作ログの件数取得



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
	zoneId := "zoneId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := openapiclient.ZonesLogsType("service") // ZonesLogsType |  (optional)
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := openapiclient.ZonesLogsOperation("updating_default_ttl") // ZonesLogsOperation |  (optional)
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := openapiclient.ZonesLogsStatus("start") // ZonesLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsZonesAPI.GetZonesLogListCount(context.Background(), zoneId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsZonesAPI.GetZonesLogListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZonesLogListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `LogsZonesAPI.GetZonesLogListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesLogListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **startDate** | **string** | 開始日 | [default to &quot;&quot;]
 **endDate** | **string** | 終了日 | [default to &quot;&quot;]
 **timeZone** | **string** | タイムゾーン | [default to &quot;+00:00&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsLogType** | [**ZonesLogsType**](ZonesLogsType.md) |  | 
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**ZonesLogsOperation**](ZonesLogsOperation.md) |  | 
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**ZonesLogsStatus**](ZonesLogsStatus.md) |  | 

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

