# \CommonConfigsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommonConfig**](CommonConfigsAPI.md#DeleteCommonConfig) | **Delete** /contracts/{ContractId}/common_configs/{CommonConfigId} | 共通設定の削除
[**GetCommonConfig**](CommonConfigsAPI.md#GetCommonConfig) | **Get** /contracts/{ContractId}/common_configs/{CommonConfigId} | 共通設定の取得
[**GetCommonConfigList**](CommonConfigsAPI.md#GetCommonConfigList) | **Get** /contracts/{ContractId}/common_configs | 共通設定の一覧取得
[**GetCommonConfigListCount**](CommonConfigsAPI.md#GetCommonConfigListCount) | **Get** /contracts/{ContractId}/common_configs/count | 共通設定の件数取得
[**PatchCommonConfig**](CommonConfigsAPI.md#PatchCommonConfig) | **Patch** /contracts/{ContractId}/common_configs/{CommonConfigId} | 共通設定の更新
[**PatchCommonConfigDefault**](CommonConfigsAPI.md#PatchCommonConfigDefault) | **Patch** /contracts/{ContractId}/common_configs/default | 初期適用される共通設定の更新
[**PatchCommonConfigManagedDns**](CommonConfigsAPI.md#PatchCommonConfigManagedDns) | **Patch** /contracts/{ContractId}/common_configs/{CommonConfigId}/managed_dns | マネージドDNSサーバの状態更新
[**PostCommonConfig**](CommonConfigsAPI.md#PostCommonConfig) | **Post** /contracts/{ContractId}/common_configs | 共通設定の作成
[**PostCommonConfigCopy**](CommonConfigsAPI.md#PostCommonConfigCopy) | **Post** /contracts/{ContractId}/common_configs/{CommonConfigId}/copy | 共通設定のコピー



## DeleteCommonConfig

> AsyncResponse DeleteCommonConfig(ctx, contractId, commonConfigId).Execute()

共通設定の削除



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
	contractId := "contractId_example" // string | ID
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.DeleteCommonConfig(context.Background(), contractId, commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.DeleteCommonConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCommonConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.DeleteCommonConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommonConfigRequest struct via the builder pattern


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


## GetCommonConfig

> GetCommonConfig GetCommonConfig(ctx, contractId, commonConfigId).Execute()

共通設定の取得



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
	contractId := "contractId_example" // string | ID
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.GetCommonConfig(context.Background(), contractId, commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.GetCommonConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommonConfig`: GetCommonConfig
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.GetCommonConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCommonConfig**](GetCommonConfig.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonConfigList

> GetCommonConfigs GetCommonConfigList(ctx, contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()

共通設定の一覧取得



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
	contractId := "contractId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.GetCommonConfigList(context.Background(), contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.GetCommonConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommonConfigList`: GetCommonConfigs
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.GetCommonConfigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]

### Return type

[**GetCommonConfigs**](GetCommonConfigs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonConfigListCount

> GetCount GetCommonConfigListCount(ctx, contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()

共通設定の件数取得



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
	contractId := "contractId_example" // string | ID
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.GetCommonConfigListCount(context.Background(), contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.GetCommonConfigListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommonConfigListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.GetCommonConfigListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonConfigListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]

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


## PatchCommonConfig

> AsyncResponse PatchCommonConfig(ctx, contractId, commonConfigId).PatchCommonConfig(patchCommonConfig).Execute()

共通設定の更新



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
	contractId := "contractId_example" // string | ID
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	patchCommonConfig := *openapiclient.NewPatchCommonConfig() // PatchCommonConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.PatchCommonConfig(context.Background(), contractId, commonConfigId).PatchCommonConfig(patchCommonConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.PatchCommonConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCommonConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.PatchCommonConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommonConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchCommonConfig** | [**PatchCommonConfig**](PatchCommonConfig.md) |  | 

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


## PatchCommonConfigDefault

> AsyncResponse PatchCommonConfigDefault(ctx, contractId).PatchDefaultCommonConfig(patchDefaultCommonConfig).Execute()

初期適用される共通設定の更新



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
	contractId := "contractId_example" // string | ID
	patchDefaultCommonConfig := *openapiclient.NewPatchDefaultCommonConfig(int64(123)) // PatchDefaultCommonConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.PatchCommonConfigDefault(context.Background(), contractId).PatchDefaultCommonConfig(patchDefaultCommonConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.PatchCommonConfigDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCommonConfigDefault`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.PatchCommonConfigDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommonConfigDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDefaultCommonConfig** | [**PatchDefaultCommonConfig**](PatchDefaultCommonConfig.md) |  | 

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


## PatchCommonConfigManagedDns

> AsyncResponse PatchCommonConfigManagedDns(ctx, contractId, commonConfigId).PatchManagedDns(patchManagedDns).Execute()

マネージドDNSサーバの状態更新



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
	contractId := "contractId_example" // string | ID
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	patchManagedDns := *openapiclient.NewPatchManagedDns(openapiclient.ManagedDnsEnabled(0)) // PatchManagedDns |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.PatchCommonConfigManagedDns(context.Background(), contractId, commonConfigId).PatchManagedDns(patchManagedDns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.PatchCommonConfigManagedDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCommonConfigManagedDns`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.PatchCommonConfigManagedDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommonConfigManagedDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchManagedDns** | [**PatchManagedDns**](PatchManagedDns.md) |  | 

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


## PostCommonConfig

> AsyncResponse PostCommonConfig(ctx, contractId).PostCommonConfig(postCommonConfig).Execute()

共通設定の作成



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
	contractId := "contractId_example" // string | ID
	postCommonConfig := *openapiclient.NewPostCommonConfig("Name_example") // PostCommonConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.PostCommonConfig(context.Background(), contractId).PostCommonConfig(postCommonConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.PostCommonConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommonConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.PostCommonConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommonConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCommonConfig** | [**PostCommonConfig**](PostCommonConfig.md) |  | 

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


## PostCommonConfigCopy

> AsyncResponse PostCommonConfigCopy(ctx, contractId, commonConfigId).PostCommonConfig(postCommonConfig).Execute()

共通設定のコピー



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
	contractId := "contractId_example" // string | ID
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	postCommonConfig := *openapiclient.NewPostCommonConfig("Name_example") // PostCommonConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonConfigsAPI.PostCommonConfigCopy(context.Background(), contractId, commonConfigId).PostCommonConfig(postCommonConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonConfigsAPI.PostCommonConfigCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommonConfigCopy`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CommonConfigsAPI.PostCommonConfigCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommonConfigCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCommonConfig** | [**PostCommonConfig**](PostCommonConfig.md) |  | 

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

