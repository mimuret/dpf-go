# \DefaultTtlAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDefaultTtlChanges**](DefaultTtlAPI.md#DeleteDefaultTtlChanges) | **Delete** /zones/{ZoneId}/default_ttl/changes | 編集中デフォルトTTLの取消
[**GetDefaultTtlDiffs**](DefaultTtlAPI.md#GetDefaultTtlDiffs) | **Get** /zones/{ZoneId}/default_ttl/diffs | デフォルトTTLの編集差分の取得
[**GetDefaultTtlList**](DefaultTtlAPI.md#GetDefaultTtlList) | **Get** /zones/{ZoneId}/default_ttl | デフォルトTTLの取得
[**PatchDefaultTtl**](DefaultTtlAPI.md#PatchDefaultTtl) | **Patch** /zones/{ZoneId}/default_ttl | デフォルトTTLの更新



## DeleteDefaultTtlChanges

> AsyncResponse DeleteDefaultTtlChanges(ctx, zoneId).Execute()

編集中デフォルトTTLの取消



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultTtlAPI.DeleteDefaultTtlChanges(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultTtlAPI.DeleteDefaultTtlChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDefaultTtlChanges`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultTtlAPI.DeleteDefaultTtlChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultTtlChangesRequest struct via the builder pattern


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


## GetDefaultTtlDiffs

> GetDefaultTtlDiffs GetDefaultTtlDiffs(ctx, zoneId).Execute()

デフォルトTTLの編集差分の取得



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultTtlAPI.GetDefaultTtlDiffs(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultTtlAPI.GetDefaultTtlDiffs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTtlDiffs`: GetDefaultTtlDiffs
	fmt.Fprintf(os.Stdout, "Response from `DefaultTtlAPI.GetDefaultTtlDiffs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTtlDiffsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultTtlDiffs**](GetDefaultTtlDiffs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTtlList

> GetDefaultTtl GetDefaultTtlList(ctx, zoneId).Execute()

デフォルトTTLの取得



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultTtlAPI.GetDefaultTtlList(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultTtlAPI.GetDefaultTtlList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTtlList`: GetDefaultTtl
	fmt.Fprintf(os.Stdout, "Response from `DefaultTtlAPI.GetDefaultTtlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTtlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultTtl**](GetDefaultTtl.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDefaultTtl

> AsyncResponse PatchDefaultTtl(ctx, zoneId).PatchDefaultTtl(patchDefaultTtl).Execute()

デフォルトTTLの更新



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
	patchDefaultTtl := *openapiclient.NewPatchDefaultTtl(int32(123)) // PatchDefaultTtl |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultTtlAPI.PatchDefaultTtl(context.Background(), zoneId).PatchDefaultTtl(patchDefaultTtl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultTtlAPI.PatchDefaultTtl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDefaultTtl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultTtlAPI.PatchDefaultTtl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDefaultTtlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDefaultTtl** | [**PatchDefaultTtl**](PatchDefaultTtl.md) |  | 

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

