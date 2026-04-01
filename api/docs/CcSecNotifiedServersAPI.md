# \CcSecNotifiedServersAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCcSecNotifiedServer**](CcSecNotifiedServersAPI.md#DeleteCcSecNotifiedServer) | **Delete** /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId} | DNS NOTIFY設定の削除
[**GetCcSecNotifiedServer**](CcSecNotifiedServersAPI.md#GetCcSecNotifiedServer) | **Get** /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId} | DNS NOTIFY設定の取得
[**GetCcSecNotifiedServerList**](CcSecNotifiedServersAPI.md#GetCcSecNotifiedServerList) | **Get** /common_configs/{CommonConfigId}/cc_sec_notified_servers | DNS NOTIFY設定の一覧取得
[**PatchCcSecNotifiedServer**](CcSecNotifiedServersAPI.md#PatchCcSecNotifiedServer) | **Patch** /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId} | DNS NOTIFY設定の更新
[**PostCcSecNotifiedServer**](CcSecNotifiedServersAPI.md#PostCcSecNotifiedServer) | **Post** /common_configs/{CommonConfigId}/cc_sec_notified_servers | DNS NOTIFY設定の作成



## DeleteCcSecNotifiedServer

> AsyncResponse DeleteCcSecNotifiedServer(ctx, commonConfigId, ccSecNotifiedServerId).Execute()

DNS NOTIFY設定の削除



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccSecNotifiedServerId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecNotifiedServersAPI.DeleteCcSecNotifiedServer(context.Background(), commonConfigId, ccSecNotifiedServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecNotifiedServersAPI.DeleteCcSecNotifiedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCcSecNotifiedServer`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecNotifiedServersAPI.DeleteCcSecNotifiedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecNotifiedServerId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcSecNotifiedServerRequest struct via the builder pattern


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


## GetCcSecNotifiedServer

> GetCcSecNotifiedServer GetCcSecNotifiedServer(ctx, commonConfigId, ccSecNotifiedServerId).Execute()

DNS NOTIFY設定の取得



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccSecNotifiedServerId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecNotifiedServersAPI.GetCcSecNotifiedServer(context.Background(), commonConfigId, ccSecNotifiedServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecNotifiedServersAPI.GetCcSecNotifiedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcSecNotifiedServer`: GetCcSecNotifiedServer
	fmt.Fprintf(os.Stdout, "Response from `CcSecNotifiedServersAPI.GetCcSecNotifiedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecNotifiedServerId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcSecNotifiedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCcSecNotifiedServer**](GetCcSecNotifiedServer.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCcSecNotifiedServerList

> GetCcSecNotifiedServers GetCcSecNotifiedServerList(ctx, commonConfigId).Execute()

DNS NOTIFY設定の一覧取得



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecNotifiedServersAPI.GetCcSecNotifiedServerList(context.Background(), commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecNotifiedServersAPI.GetCcSecNotifiedServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcSecNotifiedServerList`: GetCcSecNotifiedServers
	fmt.Fprintf(os.Stdout, "Response from `CcSecNotifiedServersAPI.GetCcSecNotifiedServerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcSecNotifiedServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCcSecNotifiedServers**](GetCcSecNotifiedServers.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCcSecNotifiedServer

> AsyncResponse PatchCcSecNotifiedServer(ctx, commonConfigId, ccSecNotifiedServerId).PatchCcSecNotifiedServer(patchCcSecNotifiedServer).Execute()

DNS NOTIFY設定の更新



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccSecNotifiedServerId := int64(789) // int64 | ID
	patchCcSecNotifiedServer := *openapiclient.NewPatchCcSecNotifiedServer() // PatchCcSecNotifiedServer |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecNotifiedServersAPI.PatchCcSecNotifiedServer(context.Background(), commonConfigId, ccSecNotifiedServerId).PatchCcSecNotifiedServer(patchCcSecNotifiedServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecNotifiedServersAPI.PatchCcSecNotifiedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCcSecNotifiedServer`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecNotifiedServersAPI.PatchCcSecNotifiedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecNotifiedServerId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCcSecNotifiedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchCcSecNotifiedServer** | [**PatchCcSecNotifiedServer**](PatchCcSecNotifiedServer.md) |  | 

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


## PostCcSecNotifiedServer

> AsyncResponse PostCcSecNotifiedServer(ctx, commonConfigId).PostCcSecNotifiedServer(postCcSecNotifiedServer).Execute()

DNS NOTIFY設定の作成



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	postCcSecNotifiedServer := *openapiclient.NewPostCcSecNotifiedServer("Address_example") // PostCcSecNotifiedServer |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecNotifiedServersAPI.PostCcSecNotifiedServer(context.Background(), commonConfigId).PostCcSecNotifiedServer(postCcSecNotifiedServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecNotifiedServersAPI.PostCcSecNotifiedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCcSecNotifiedServer`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecNotifiedServersAPI.PostCcSecNotifiedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCcSecNotifiedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCcSecNotifiedServer** | [**PostCcSecNotifiedServer**](PostCcSecNotifiedServer.md) |  | 

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

