# \ZoneHistoriesAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZoneHistoryList**](ZoneHistoriesAPI.md#GetZoneHistoryList) | **Get** /zones/{ZoneId}/zone_histories | ゾーン反映履歴の一覧取得
[**GetZoneHistoryListCount**](ZoneHistoriesAPI.md#GetZoneHistoryListCount) | **Get** /zones/{ZoneId}/zone_histories/count | ゾーン反映履歴の件数取得
[**GetZoneHistoryText**](ZoneHistoriesAPI.md#GetZoneHistoryText) | **Get** /zones/{ZoneId}/zone_histories/{ZoneHistoryId}/text | ゾーン反映時のRFC1035形式のテキストの取得



## GetZoneHistoryList

> GetZoneHistories GetZoneHistoryList(ctx, zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).Order(order).Execute()

ゾーン反映履歴の一覧取得



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
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	order := openapiclient.SearchOrder("ASC") // SearchOrder | ソート順 (optional) (default to "ASC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneHistoriesAPI.GetZoneHistoryList(context.Background(), zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneHistoriesAPI.GetZoneHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneHistoryList`: GetZoneHistories
	fmt.Fprintf(os.Stdout, "Response from `ZoneHistoriesAPI.GetZoneHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneHistoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **order** | [**SearchOrder**](SearchOrder.md) | ソート順 | [default to &quot;ASC&quot;]

### Return type

[**GetZoneHistories**](GetZoneHistories.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneHistoryListCount

> GetCount GetZoneHistoryListCount(ctx, zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).Execute()

ゾーン反映履歴の件数取得



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
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneHistoriesAPI.GetZoneHistoryListCount(context.Background(), zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneHistoriesAPI.GetZoneHistoryListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneHistoryListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `ZoneHistoriesAPI.GetZoneHistoryListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneHistoryListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]

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


## GetZoneHistoryText

> GetZoneHistoriesText GetZoneHistoryText(ctx, zoneId, zoneHistoryId).Execute()

ゾーン反映時のRFC1035形式のテキストの取得



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
	zoneHistoryId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneHistoriesAPI.GetZoneHistoryText(context.Background(), zoneId, zoneHistoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneHistoriesAPI.GetZoneHistoryText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneHistoryText`: GetZoneHistoriesText
	fmt.Fprintf(os.Stdout, "Response from `ZoneHistoriesAPI.GetZoneHistoryText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**zoneHistoryId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneHistoryTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetZoneHistoriesText**](GetZoneHistoriesText.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

