# \EndpointsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEndpoint**](EndpointsAPI.md#DeleteEndpoint) | **Delete** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints/{EndpointResourceName} | エンドポイントの削除
[**GetEndpoint**](EndpointsAPI.md#GetEndpoint) | **Get** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints/{EndpointResourceName} | エンドポイントの取得
[**GetEndpointList**](EndpointsAPI.md#GetEndpointList) | **Get** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints | エンドポイントの一覧取得
[**PatchEndpoint**](EndpointsAPI.md#PatchEndpoint) | **Patch** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints/{EndpointResourceName} | エンドポイントの更新
[**PostEndpoint**](EndpointsAPI.md#PostEndpoint) | **Post** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints | エンドポイントの作成
[**PostEndpointFailback**](EndpointsAPI.md#PostEndpointFailback) | **Post** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints/{EndpointResourceName}/failback | エンドポイントの手動切り戻し
[**PostEndpointFailover**](EndpointsAPI.md#PostEndpointFailover) | **Post** /lb_domains/{LbDomainId}/sites/{SiteResourceName}/endpoints/{EndpointResourceName}/failover | エンドポイントの手動切り離し



## DeleteEndpoint

> AsyncResponse DeleteEndpoint(ctx, lbDomainId, siteResourceName, endpointResourceName).Execute()

エンドポイントの削除



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	endpointResourceName := "endpointResourceName_example" // string | GET endpoints Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.DeleteEndpoint(context.Background(), lbDomainId, siteResourceName, endpointResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEndpoint`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.DeleteEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 
**endpointResourceName** | **string** | GET endpoints Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


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


## GetEndpoint

> GetEndpoint GetEndpoint(ctx, lbDomainId, siteResourceName, endpointResourceName).Execute()

エンドポイントの取得



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	endpointResourceName := "endpointResourceName_example" // string | GET endpoints Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpoint(context.Background(), lbDomainId, siteResourceName, endpointResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoint`: GetEndpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 
**endpointResourceName** | **string** | GET endpoints Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetEndpoint**](GetEndpoint.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointList

> GetEndpoints GetEndpointList(ctx, lbDomainId, siteResourceName).Execute()

エンドポイントの一覧取得



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointList(context.Background(), lbDomainId, siteResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointList`: GetEndpoints
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpoints**](GetEndpoints.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEndpoint

> AsyncResponse PatchEndpoint(ctx, lbDomainId, siteResourceName, endpointResourceName).PatchEndpoint(patchEndpoint).Execute()

エンドポイントの更新



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	endpointResourceName := "endpointResourceName_example" // string | GET endpoints Schemaにおける resource_name
	patchEndpoint := *openapiclient.NewPatchEndpoint() // PatchEndpoint |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PatchEndpoint(context.Background(), lbDomainId, siteResourceName, endpointResourceName).PatchEndpoint(patchEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PatchEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEndpoint`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PatchEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 
**endpointResourceName** | **string** | GET endpoints Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchEndpoint** | [**PatchEndpoint**](PatchEndpoint.md) |  | 

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


## PostEndpoint

> AsyncResponse PostEndpoint(ctx, lbDomainId, siteResourceName).PostEndpoint(postEndpoint).Execute()

エンドポイントの作成



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	postEndpoint := *openapiclient.NewPostEndpoint("Name_example", int32(123), []openapiclient.EndpointRdataInner{*openapiclient.NewEndpointRdataInner("Value_example")}) // PostEndpoint |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PostEndpoint(context.Background(), lbDomainId, siteResourceName).PostEndpoint(postEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PostEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEndpoint`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PostEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postEndpoint** | [**PostEndpoint**](PostEndpoint.md) |  | 

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


## PostEndpointFailback

> AsyncResponse PostEndpointFailback(ctx, lbDomainId, siteResourceName, endpointResourceName).PostManualFailback(postManualFailback).Execute()

エンドポイントの手動切り戻し



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	endpointResourceName := "endpointResourceName_example" // string | GET endpoints Schemaにおける resource_name
	postManualFailback := *openapiclient.NewPostManualFailback() // PostManualFailback |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PostEndpointFailback(context.Background(), lbDomainId, siteResourceName, endpointResourceName).PostManualFailback(postManualFailback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PostEndpointFailback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEndpointFailback`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PostEndpointFailback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 
**endpointResourceName** | **string** | GET endpoints Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEndpointFailbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postManualFailback** | [**PostManualFailback**](PostManualFailback.md) |  | 

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


## PostEndpointFailover

> AsyncResponse PostEndpointFailover(ctx, lbDomainId, siteResourceName, endpointResourceName).Execute()

エンドポイントの手動切り離し



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
	siteResourceName := "siteResourceName_example" // string | GET sites Schemaにおける resource_name
	endpointResourceName := "endpointResourceName_example" // string | GET endpoints Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PostEndpointFailover(context.Background(), lbDomainId, siteResourceName, endpointResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PostEndpointFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEndpointFailover`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PostEndpointFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**siteResourceName** | **string** | GET sites Schemaにおける resource_name | 
**endpointResourceName** | **string** | GET endpoints Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEndpointFailoverRequest struct via the builder pattern


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

