# \ZoneProxyAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZoneProxyHealthCheck**](ZoneProxyAPI.md#GetZoneProxyHealthCheck) | **Get** /zones/{ZoneId}/zone_proxy/health_check | プライマリネームサーバのヘルスチェック結果の取得
[**GetZoneProxyList**](ZoneProxyAPI.md#GetZoneProxyList) | **Get** /zones/{ZoneId}/zone_proxy | ゾーンプロキシ設定の取得
[**PatchZoneProxy**](ZoneProxyAPI.md#PatchZoneProxy) | **Patch** /zones/{ZoneId}/zone_proxy | ゾーンプロキシ設定の更新



## GetZoneProxyHealthCheck

> GetZoneProxyHealth GetZoneProxyHealthCheck(ctx, zoneId).Execute()

プライマリネームサーバのヘルスチェック結果の取得



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
	resp, r, err := apiClient.ZoneProxyAPI.GetZoneProxyHealthCheck(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneProxyAPI.GetZoneProxyHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneProxyHealthCheck`: GetZoneProxyHealth
	fmt.Fprintf(os.Stdout, "Response from `ZoneProxyAPI.GetZoneProxyHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneProxyHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetZoneProxyHealth**](GetZoneProxyHealth.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneProxyList

> GetZoneProxy GetZoneProxyList(ctx, zoneId).Execute()

ゾーンプロキシ設定の取得



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
	resp, r, err := apiClient.ZoneProxyAPI.GetZoneProxyList(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneProxyAPI.GetZoneProxyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneProxyList`: GetZoneProxy
	fmt.Fprintf(os.Stdout, "Response from `ZoneProxyAPI.GetZoneProxyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneProxyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetZoneProxy**](GetZoneProxy.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchZoneProxy

> AsyncResponse PatchZoneProxy(ctx, zoneId).PatchZoneProxy(patchZoneProxy).Execute()

ゾーンプロキシ設定の更新



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
	patchZoneProxy := *openapiclient.NewPatchZoneProxy(openapiclient.ZoneProxyEnabled(0)) // PatchZoneProxy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneProxyAPI.PatchZoneProxy(context.Background(), zoneId).PatchZoneProxy(patchZoneProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneProxyAPI.PatchZoneProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchZoneProxy`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZoneProxyAPI.PatchZoneProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchZoneProxy** | [**PatchZoneProxy**](PatchZoneProxy.md) |  | 

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

