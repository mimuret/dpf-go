# \MonitoringsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitoring**](MonitoringsAPI.md#DeleteMonitoring) | **Delete** /lb_domains/{LbDomainId}/monitorings/{MonitoringResourceName} | 監視の削除
[**GetMonitoring**](MonitoringsAPI.md#GetMonitoring) | **Get** /lb_domains/{LbDomainId}/monitorings/{MonitoringResourceName} | 監視の取得
[**GetMonitoringList**](MonitoringsAPI.md#GetMonitoringList) | **Get** /lb_domains/{LbDomainId}/monitorings | 監視の一覧取得
[**PatchMonitoring**](MonitoringsAPI.md#PatchMonitoring) | **Patch** /lb_domains/{LbDomainId}/monitorings/{MonitoringResourceName} | 監視の更新
[**PostMonitoring**](MonitoringsAPI.md#PostMonitoring) | **Post** /lb_domains/{LbDomainId}/monitorings | 監視の作成



## DeleteMonitoring

> AsyncResponse DeleteMonitoring(ctx, lbDomainId, monitoringResourceName).Execute()

監視の削除



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
	monitoringResourceName := "monitoringResourceName_example" // string | GET monitorings Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringsAPI.DeleteMonitoring(context.Background(), lbDomainId, monitoringResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringsAPI.DeleteMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMonitoring`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitoringsAPI.DeleteMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**monitoringResourceName** | **string** | GET monitorings Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitoringRequest struct via the builder pattern


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


## GetMonitoring

> GetMonitoring GetMonitoring(ctx, lbDomainId, monitoringResourceName).Execute()

監視の取得



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
	monitoringResourceName := "monitoringResourceName_example" // string | GET monitorings Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringsAPI.GetMonitoring(context.Background(), lbDomainId, monitoringResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringsAPI.GetMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitoring`: GetMonitoring
	fmt.Fprintf(os.Stdout, "Response from `MonitoringsAPI.GetMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**monitoringResourceName** | **string** | GET monitorings Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetMonitoring**](GetMonitoring.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringList

> GetMonitorings GetMonitoringList(ctx, lbDomainId).Execute()

監視の一覧取得



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringsAPI.GetMonitoringList(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringsAPI.GetMonitoringList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitoringList`: GetMonitorings
	fmt.Fprintf(os.Stdout, "Response from `MonitoringsAPI.GetMonitoringList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMonitorings**](GetMonitorings.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMonitoring

> AsyncResponse PatchMonitoring(ctx, lbDomainId, monitoringResourceName).PatchMonitoring(patchMonitoring).Execute()

監視の更新



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
	monitoringResourceName := "monitoringResourceName_example" // string | GET monitorings Schemaにおける resource_name
	patchMonitoring := *openapiclient.NewPatchMonitoring() // PatchMonitoring |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringsAPI.PatchMonitoring(context.Background(), lbDomainId, monitoringResourceName).PatchMonitoring(patchMonitoring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringsAPI.PatchMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMonitoring`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitoringsAPI.PatchMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**monitoringResourceName** | **string** | GET monitorings Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchMonitoring** | [**PatchMonitoring**](PatchMonitoring.md) |  | 

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


## PostMonitoring

> AsyncResponse PostMonitoring(ctx, lbDomainId).PostMonitoring(postMonitoring).Execute()

監視の作成



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
	postMonitoring := *openapiclient.NewPostMonitoring("Name_example", "Mtype_example") // PostMonitoring |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringsAPI.PostMonitoring(context.Background(), lbDomainId).PostMonitoring(postMonitoring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringsAPI.PostMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMonitoring`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitoringsAPI.PostMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMonitoring** | [**PostMonitoring**](PostMonitoring.md) |  | 

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

