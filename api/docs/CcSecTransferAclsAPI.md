# \CcSecTransferAclsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCcSecTransferAcl**](CcSecTransferAclsAPI.md#DeleteCcSecTransferAcl) | **Delete** /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId} | ゾーン転送ACLの削除
[**GetCcSecTransferAcl**](CcSecTransferAclsAPI.md#GetCcSecTransferAcl) | **Get** /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId} | ゾーン転送ACLの取得
[**GetCcSecTransferAclList**](CcSecTransferAclsAPI.md#GetCcSecTransferAclList) | **Get** /common_configs/{CommonConfigId}/cc_sec_transfer_acls | ゾーン転送ACLの一覧取得
[**PatchCcSecTransferAcl**](CcSecTransferAclsAPI.md#PatchCcSecTransferAcl) | **Patch** /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId} | ゾーン転送ACLの更新
[**PostCcSecTransferAcl**](CcSecTransferAclsAPI.md#PostCcSecTransferAcl) | **Post** /common_configs/{CommonConfigId}/cc_sec_transfer_acls | ゾーン転送ACLの作成



## DeleteCcSecTransferAcl

> AsyncResponse DeleteCcSecTransferAcl(ctx, commonConfigId, ccSecTransferAclId).Execute()

ゾーン転送ACLの削除



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
	ccSecTransferAclId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecTransferAclsAPI.DeleteCcSecTransferAcl(context.Background(), commonConfigId, ccSecTransferAclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecTransferAclsAPI.DeleteCcSecTransferAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCcSecTransferAcl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecTransferAclsAPI.DeleteCcSecTransferAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecTransferAclId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcSecTransferAclRequest struct via the builder pattern


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


## GetCcSecTransferAcl

> GetCcSecTransferAcl GetCcSecTransferAcl(ctx, commonConfigId, ccSecTransferAclId).Execute()

ゾーン転送ACLの取得



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
	ccSecTransferAclId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecTransferAclsAPI.GetCcSecTransferAcl(context.Background(), commonConfigId, ccSecTransferAclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecTransferAclsAPI.GetCcSecTransferAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcSecTransferAcl`: GetCcSecTransferAcl
	fmt.Fprintf(os.Stdout, "Response from `CcSecTransferAclsAPI.GetCcSecTransferAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecTransferAclId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcSecTransferAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCcSecTransferAcl**](GetCcSecTransferAcl.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCcSecTransferAclList

> GetCcSecTransferAcls GetCcSecTransferAclList(ctx, commonConfigId).Execute()

ゾーン転送ACLの一覧取得



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
	resp, r, err := apiClient.CcSecTransferAclsAPI.GetCcSecTransferAclList(context.Background(), commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecTransferAclsAPI.GetCcSecTransferAclList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcSecTransferAclList`: GetCcSecTransferAcls
	fmt.Fprintf(os.Stdout, "Response from `CcSecTransferAclsAPI.GetCcSecTransferAclList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcSecTransferAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCcSecTransferAcls**](GetCcSecTransferAcls.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCcSecTransferAcl

> AsyncResponse PatchCcSecTransferAcl(ctx, commonConfigId, ccSecTransferAclId).PatchCcSecTransferAcl(patchCcSecTransferAcl).Execute()

ゾーン転送ACLの更新



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
	ccSecTransferAclId := int64(789) // int64 | ID
	patchCcSecTransferAcl := *openapiclient.NewPatchCcSecTransferAcl() // PatchCcSecTransferAcl |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecTransferAclsAPI.PatchCcSecTransferAcl(context.Background(), commonConfigId, ccSecTransferAclId).PatchCcSecTransferAcl(patchCcSecTransferAcl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecTransferAclsAPI.PatchCcSecTransferAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCcSecTransferAcl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecTransferAclsAPI.PatchCcSecTransferAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccSecTransferAclId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCcSecTransferAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchCcSecTransferAcl** | [**PatchCcSecTransferAcl**](PatchCcSecTransferAcl.md) |  | 

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


## PostCcSecTransferAcl

> AsyncResponse PostCcSecTransferAcl(ctx, commonConfigId).PostCcSecTransferAcl(postCcSecTransferAcl).Execute()

ゾーン転送ACLの作成



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
	postCcSecTransferAcl := *openapiclient.NewPostCcSecTransferAcl("Network_example") // PostCcSecTransferAcl |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcSecTransferAclsAPI.PostCcSecTransferAcl(context.Background(), commonConfigId).PostCcSecTransferAcl(postCcSecTransferAcl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcSecTransferAclsAPI.PostCcSecTransferAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCcSecTransferAcl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcSecTransferAclsAPI.PostCcSecTransferAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCcSecTransferAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCcSecTransferAcl** | [**PostCcSecTransferAcl**](PostCcSecTransferAcl.md) |  | 

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

