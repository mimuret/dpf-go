# \CcPrimariesAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCcPrimary**](CcPrimariesAPI.md#DeleteCcPrimary) | **Delete** /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId} | プライマリネームサーバ設定の削除
[**GetCcPrimary**](CcPrimariesAPI.md#GetCcPrimary) | **Get** /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId} | プライマリネームサーバ設定の取得
[**GetCcPrimaryList**](CcPrimariesAPI.md#GetCcPrimaryList) | **Get** /common_configs/{CommonConfigId}/cc_primaries | プライマリネームサーバ設定の一覧取得
[**PatchCcPrimary**](CcPrimariesAPI.md#PatchCcPrimary) | **Patch** /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId} | プライマリネームサーバ設定の更新
[**PostCcPrimary**](CcPrimariesAPI.md#PostCcPrimary) | **Post** /common_configs/{CommonConfigId}/cc_primaries | プライマリネームサーバ設定の作成



## DeleteCcPrimary

> AsyncResponse DeleteCcPrimary(ctx, commonConfigId, ccPrimaryId).Execute()

プライマリネームサーバ設定の削除



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
	ccPrimaryId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcPrimariesAPI.DeleteCcPrimary(context.Background(), commonConfigId, ccPrimaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcPrimariesAPI.DeleteCcPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCcPrimary`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcPrimariesAPI.DeleteCcPrimary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccPrimaryId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcPrimaryRequest struct via the builder pattern


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


## GetCcPrimary

> GetCcPrimary GetCcPrimary(ctx, commonConfigId, ccPrimaryId).Execute()

プライマリネームサーバ設定の取得



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
	ccPrimaryId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcPrimariesAPI.GetCcPrimary(context.Background(), commonConfigId, ccPrimaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcPrimariesAPI.GetCcPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcPrimary`: GetCcPrimary
	fmt.Fprintf(os.Stdout, "Response from `CcPrimariesAPI.GetCcPrimary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccPrimaryId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCcPrimary**](GetCcPrimary.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCcPrimaryList

> GetCcPrimaries GetCcPrimaryList(ctx, commonConfigId).Execute()

プライマリネームサーバ設定の一覧取得



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
	resp, r, err := apiClient.CcPrimariesAPI.GetCcPrimaryList(context.Background(), commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcPrimariesAPI.GetCcPrimaryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcPrimaryList`: GetCcPrimaries
	fmt.Fprintf(os.Stdout, "Response from `CcPrimariesAPI.GetCcPrimaryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcPrimaryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCcPrimaries**](GetCcPrimaries.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCcPrimary

> AsyncResponse PatchCcPrimary(ctx, commonConfigId, ccPrimaryId).PatchCcPrimary(patchCcPrimary).Execute()

プライマリネームサーバ設定の更新



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
	ccPrimaryId := int64(789) // int64 | ID
	patchCcPrimary := *openapiclient.NewPatchCcPrimary() // PatchCcPrimary |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcPrimariesAPI.PatchCcPrimary(context.Background(), commonConfigId, ccPrimaryId).PatchCcPrimary(patchCcPrimary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcPrimariesAPI.PatchCcPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCcPrimary`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcPrimariesAPI.PatchCcPrimary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccPrimaryId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCcPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchCcPrimary** | [**PatchCcPrimary**](PatchCcPrimary.md) |  | 

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


## PostCcPrimary

> AsyncResponse PostCcPrimary(ctx, commonConfigId).PostCcPrimary(postCcPrimary).Execute()

プライマリネームサーバ設定の作成



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
	postCcPrimary := *openapiclient.NewPostCcPrimary("Address_example") // PostCcPrimary |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcPrimariesAPI.PostCcPrimary(context.Background(), commonConfigId).PostCcPrimary(postCcPrimary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcPrimariesAPI.PostCcPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCcPrimary`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcPrimariesAPI.PostCcPrimary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCcPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCcPrimary** | [**PostCcPrimary**](PostCcPrimary.md) |  | 

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

