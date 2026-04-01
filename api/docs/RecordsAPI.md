# \RecordsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecord**](RecordsAPI.md#DeleteRecord) | **Delete** /zones/{ZoneId}/records/{RecordId} | レコードの削除
[**DeleteRecordChanges**](RecordsAPI.md#DeleteRecordChanges) | **Delete** /zones/{ZoneId}/records/{RecordId}/changes | 編集中レコードの取消
[**GetDsRecordsFromCds**](RecordsAPI.md#GetDsRecordsFromCds) | **Get** /zones/{ZoneId}/records/{RecordId}/ds_records | CDS経由のDSレコードの取得
[**GetRecord**](RecordsAPI.md#GetRecord) | **Get** /zones/{ZoneId}/records/{RecordId} | レコードの取得
[**GetRecordCurrents**](RecordsAPI.md#GetRecordCurrents) | **Get** /zones/{ZoneId}/records/currents | DNS反映済レコードの一覧取得
[**GetRecordCurrentsCount**](RecordsAPI.md#GetRecordCurrentsCount) | **Get** /zones/{ZoneId}/records/currents/count | DNS反映済レコードの件数取得
[**GetRecordDiffs**](RecordsAPI.md#GetRecordDiffs) | **Get** /zones/{ZoneId}/records/diffs | レコードの編集差分の一覧取得
[**GetRecordDiffsCount**](RecordsAPI.md#GetRecordDiffsCount) | **Get** /zones/{ZoneId}/records/diffs/count | レコードの編集差分の件数取得
[**GetRecordList**](RecordsAPI.md#GetRecordList) | **Get** /zones/{ZoneId}/records | レコードの一覧取得
[**GetRecordListCount**](RecordsAPI.md#GetRecordListCount) | **Get** /zones/{ZoneId}/records/count | レコードの件数取得
[**PatchRecord**](RecordsAPI.md#PatchRecord) | **Patch** /zones/{ZoneId}/records/{RecordId} | レコードの更新
[**PostRecord**](RecordsAPI.md#PostRecord) | **Post** /zones/{ZoneId}/records | レコードの作成
[**PutRecord**](RecordsAPI.md#PutRecord) | **Put** /zones/{ZoneId}/records | レコードの一括更新



## DeleteRecord

> AsyncResponse DeleteRecord(ctx, zoneId, recordId).Execute()

レコードの削除



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
	recordId := "recordId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.DeleteRecord(context.Background(), zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.DeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecord`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.DeleteRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**recordId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[DPFOperator](../README.md#DPFOperator)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordChanges

> AsyncResponse DeleteRecordChanges(ctx, zoneId, recordId).Execute()

編集中レコードの取消



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
	recordId := "recordId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.DeleteRecordChanges(context.Background(), zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.DeleteRecordChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecordChanges`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.DeleteRecordChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**recordId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[DPFOperator](../README.md#DPFOperator)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsRecordsFromCds

> GetDsRecordsFromCds200Response GetDsRecordsFromCds(ctx, zoneId, recordId).Execute()

CDS経由のDSレコードの取得



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
	recordId := "recordId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetDsRecordsFromCds(context.Background(), zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetDsRecordsFromCds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDsRecordsFromCds`: GetDsRecordsFromCds200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetDsRecordsFromCds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**recordId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDsRecordsFromCdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDsRecordsFromCds200Response**](GetDsRecordsFromCds200Response.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecord

> GetRecord GetRecord(ctx, zoneId, recordId).Execute()

レコードの取得



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
	recordId := "recordId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecord(context.Background(), zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecord`: GetRecord
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**recordId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRecord**](GetRecord.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordCurrents

> GetRecords GetRecordCurrents(ctx, zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()

DNS反映済レコードの一覧取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordCurrents(context.Background(), zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordCurrents`: GetRecords
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordCurrents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordCurrentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetRecords**](GetRecords.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordCurrentsCount

> GetCount GetRecordCurrentsCount(ctx, zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()

DNS反映済レコードの件数取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordCurrentsCount(context.Background(), zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordCurrentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordCurrentsCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordCurrentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordCurrentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

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


## GetRecordDiffs

> GetRecordsDiffs GetRecordDiffs(ctx, zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()

レコードの編集差分の一覧取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordDiffs(context.Background(), zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordDiffs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordDiffs`: GetRecordsDiffs
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordDiffs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordDiffsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetRecordsDiffs**](GetRecordsDiffs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordDiffsCount

> GetCount GetRecordDiffsCount(ctx, zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()

レコードの編集差分の件数取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordDiffsCount(context.Background(), zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordDiffsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordDiffsCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordDiffsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordDiffsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

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


## GetRecordList

> GetRecords GetRecordList(ctx, zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()

レコードの一覧取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordList(context.Background(), zoneId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordList`: GetRecords
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetRecords**](GetRecords.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordListCount

> GetCount GetRecordListCount(ctx, zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()

レコードの件数取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsTtl := int32(56) // int32 |  (optional)
	keywordsRrtype := openapiclient.RecordsRrtype("A") // RecordsRrtype |  (optional)
	keywordsRdata := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsOperator := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecordListCount(context.Background(), zoneId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsTtl(keywordsTtl).KeywordsRrtype(keywordsRrtype).KeywordsRdata(keywordsRdata).KeywordsDescription(keywordsDescription).KeywordsOperator(keywordsOperator).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecordListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecordListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsTtl** | **int32** |  | 
 **keywordsRrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
 **keywordsRdata** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsOperator** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

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


## PatchRecord

> AsyncResponse PatchRecord(ctx, zoneId, recordId).PatchRecord(patchRecord).Execute()

レコードの更新



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
	recordId := "recordId_example" // string | ID
	patchRecord := *openapiclient.NewPatchRecord() // PatchRecord |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.PatchRecord(context.Background(), zoneId, recordId).PatchRecord(patchRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.PatchRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRecord`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.PatchRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 
**recordId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchRecord** | [**PatchRecord**](PatchRecord.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[DPFOperator](../README.md#DPFOperator)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRecord

> AsyncResponse PostRecord(ctx, zoneId).PostRecord(postRecord).Execute()

レコードの作成



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
	postRecord := *openapiclient.NewPostRecord("Name_example", openapiclient.RecordsRrtypeWithoutSoa("A"), []openapiclient.RecordsRdataInner{*openapiclient.NewRecordsRdataInner()}) // PostRecord |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.PostRecord(context.Background(), zoneId).PostRecord(postRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.PostRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRecord`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.PostRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRecord** | [**PostRecord**](PostRecord.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[DPFOperator](../README.md#DPFOperator)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRecord

> AsyncResponse PutRecord(ctx, zoneId).PutRecord(putRecord).Execute()

レコードの一括更新



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
	putRecord := *openapiclient.NewPutRecord([]openapiclient.OverwriteRecordsInner{*openapiclient.NewOverwriteRecordsInner("Name_example", NullableInt32(123), openapiclient.RecordsRrtype("A"), []openapiclient.RecordsRdataInner{*openapiclient.NewRecordsRdataInner()}, "Description_example", map[string]string{"key": "Inner_example"})}) // PutRecord |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.PutRecord(context.Background(), zoneId).PutRecord(putRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.PutRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRecord`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.PutRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRecord** | [**PutRecord**](PutRecord.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[DPFOperator](../README.md#DPFOperator)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

