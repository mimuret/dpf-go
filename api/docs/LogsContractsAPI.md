# \LogsContractsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractsLogList**](LogsContractsAPI.md#GetContractsLogList) | **Get** /contracts/{ContractId}/logs | DPF契約操作ログの一覧取得
[**GetContractsLogListCount**](LogsContractsAPI.md#GetContractsLogListCount) | **Get** /contracts/{ContractId}/logs/count | DPF契約操作ログの件数取得



## GetContractsLogList

> GetContractsLogs GetContractsLogList(ctx, contractId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

DPF契約操作ログの一覧取得



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
	contractId := "contractId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	order := openapiclient.SearchOrder("ASC") // SearchOrder | ソート順 (optional) (default to "ASC")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := openapiclient.ContractsLogsLogType("service") // ContractsLogsLogType |  (optional)
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := openapiclient.ContractsLogsOperation("add_cc_primary") // ContractsLogsOperation |  (optional)
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := openapiclient.ContractsLogsStatus("start") // ContractsLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsContractsAPI.GetContractsLogList(context.Background(), contractId).Type_(type_).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Order(order).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsContractsAPI.GetContractsLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsLogList`: GetContractsLogs
	fmt.Fprintf(os.Stdout, "Response from `LogsContractsAPI.GetContractsLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsLogListRequest struct via the builder pattern


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
 **keywordsLogType** | [**ContractsLogsLogType**](ContractsLogsLogType.md) |  | 
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**ContractsLogsOperation**](ContractsLogsOperation.md) |  | 
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**ContractsLogsStatus**](ContractsLogsStatus.md) |  | 

### Return type

[**GetContractsLogs**](GetContractsLogs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsLogListCount

> GetCount GetContractsLogListCount(ctx, contractId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()

DPF契約操作ログの件数取得



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
	contractId := "contractId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	startDate := "startDate_example" // string | 開始日 (optional) (default to "")
	endDate := "endDate_example" // string | 終了日 (optional) (default to "")
	timeZone := "timeZone_example" // string | タイムゾーン (optional) (default to "+00:00")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLogType := openapiclient.ContractsLogsLogType("service") // ContractsLogsLogType |  (optional)
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperation := openapiclient.ContractsLogsOperation("add_cc_primary") // ContractsLogsOperation |  (optional)
	keywordsTarget := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDetail := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequestId := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsStatus := openapiclient.ContractsLogsStatus("start") // ContractsLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsContractsAPI.GetContractsLogListCount(context.Background(), contractId).Type_(type_).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).KeywordsFullText(keywordsFullText).KeywordsLogType(keywordsLogType).KeywordsOperator(keywordsOperator).KeywordsOperation(keywordsOperation).KeywordsTarget(keywordsTarget).KeywordsDetail(keywordsDetail).KeywordsRequestId(keywordsRequestId).KeywordsStatus(keywordsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsContractsAPI.GetContractsLogListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsLogListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `LogsContractsAPI.GetContractsLogListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsLogListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **startDate** | **string** | 開始日 | [default to &quot;&quot;]
 **endDate** | **string** | 終了日 | [default to &quot;&quot;]
 **timeZone** | **string** | タイムゾーン | [default to &quot;+00:00&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsLogType** | [**ContractsLogsLogType**](ContractsLogsLogType.md) |  | 
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsOperation** | [**ContractsLogsOperation**](ContractsLogsOperation.md) |  | 
 **keywordsTarget** | **[]string** |  | [default to []]
 **keywordsDetail** | **[]string** |  | [default to []]
 **keywordsRequestId** | **[]string** |  | [default to []]
 **keywordsStatus** | [**ContractsLogsStatus**](ContractsLogsStatus.md) |  | 

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

