# \DnssecAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDnssecList**](DnssecAPI.md#GetDnssecList) | **Get** /zones/{ZoneId}/dnssec | DNSSEC情報の取得
[**PatchDnssec**](DnssecAPI.md#PatchDnssec) | **Patch** /zones/{ZoneId}/dnssec | DNSSEC情報の更新
[**PatchDnssecKskRollover**](DnssecAPI.md#PatchDnssecKskRollover) | **Patch** /zones/{ZoneId}/dnssec/ksk_rollover | KSKロールオーバーの開始



## GetDnssecList

> GetDnssec GetDnssecList(ctx, zoneId).Execute()

DNSSEC情報の取得



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
	resp, r, err := apiClient.DnssecAPI.GetDnssecList(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnssecAPI.GetDnssecList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnssecList`: GetDnssec
	fmt.Fprintf(os.Stdout, "Response from `DnssecAPI.GetDnssecList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnssecListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDnssec**](GetDnssec.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDnssec

> AsyncResponse PatchDnssec(ctx, zoneId).PatchDnssec(patchDnssec).Execute()

DNSSEC情報の更新



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
	patchDnssec := *openapiclient.NewPatchDnssec(openapiclient.DnssecEnabled(0)) // PatchDnssec |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnssecAPI.PatchDnssec(context.Background(), zoneId).PatchDnssec(patchDnssec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnssecAPI.PatchDnssec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDnssec`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `DnssecAPI.PatchDnssec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDnssecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDnssec** | [**PatchDnssec**](PatchDnssec.md) |  | 

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


## PatchDnssecKskRollover

> AsyncResponse PatchDnssecKskRollover(ctx, zoneId).Execute()

KSKロールオーバーの開始



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
	resp, r, err := apiClient.DnssecAPI.PatchDnssecKskRollover(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnssecAPI.PatchDnssecKskRollover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDnssecKskRollover`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `DnssecAPI.PatchDnssecKskRollover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDnssecKskRolloverRequest struct via the builder pattern


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

