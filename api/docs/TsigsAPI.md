# \TsigsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTsig**](TsigsAPI.md#DeleteTsig) | **Delete** /contracts/{ContractId}/tsigs/{TsigId} | TSIG鍵の削除
[**GetTsig**](TsigsAPI.md#GetTsig) | **Get** /contracts/{ContractId}/tsigs/{TsigId} | TSIG鍵の取得
[**GetTsigCommonConfigs**](TsigsAPI.md#GetTsigCommonConfigs) | **Get** /contracts/{ContractId}/tsigs/{TsigId}/common_configs | TSIG鍵を利用している共通設定の一覧取得
[**GetTsigCommonConfigsCount**](TsigsAPI.md#GetTsigCommonConfigsCount) | **Get** /contracts/{ContractId}/tsigs/{TsigId}/common_configs/count | TSIG鍵を利用している共通設定の件数取得
[**GetTsigList**](TsigsAPI.md#GetTsigList) | **Get** /contracts/{ContractId}/tsigs | TSIG鍵の一覧取得
[**GetTsigListCount**](TsigsAPI.md#GetTsigListCount) | **Get** /contracts/{ContractId}/tsigs/count | TSIG鍵の件数取得
[**PatchTsig**](TsigsAPI.md#PatchTsig) | **Patch** /contracts/{ContractId}/tsigs/{TsigId} | TSIG鍵の更新
[**PostTsig**](TsigsAPI.md#PostTsig) | **Post** /contracts/{ContractId}/tsigs | TSIG鍵の作成



## DeleteTsig

> AsyncResponse DeleteTsig(ctx, contractId, tsigId).Execute()

TSIG鍵の削除



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
	tsigId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.DeleteTsig(context.Background(), contractId, tsigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.DeleteTsig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTsig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.DeleteTsig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**tsigId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTsigRequest struct via the builder pattern


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


## GetTsig

> GetTsig GetTsig(ctx, contractId, tsigId).Execute()

TSIG鍵の取得



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
	tsigId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.GetTsig(context.Background(), contractId, tsigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.GetTsig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsig`: GetTsig
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.GetTsig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**tsigId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTsigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTsig**](GetTsig.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTsigCommonConfigs

> GetTsigsCommonConfigs GetTsigCommonConfigs(ctx, contractId, tsigId).Offset(offset).Limit(limit).Execute()

TSIG鍵を利用している共通設定の一覧取得



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
	tsigId := int64(789) // int64 | ID
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.GetTsigCommonConfigs(context.Background(), contractId, tsigId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.GetTsigCommonConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsigCommonConfigs`: GetTsigsCommonConfigs
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.GetTsigCommonConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**tsigId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTsigCommonConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**GetTsigsCommonConfigs**](GetTsigsCommonConfigs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTsigCommonConfigsCount

> GetCount GetTsigCommonConfigsCount(ctx, contractId, tsigId).Execute()

TSIG鍵を利用している共通設定の件数取得



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
	tsigId := int64(789) // int64 | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.GetTsigCommonConfigsCount(context.Background(), contractId, tsigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.GetTsigCommonConfigsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsigCommonConfigsCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.GetTsigCommonConfigsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**tsigId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTsigCommonConfigsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetTsigList

> GetTsigs GetTsigList(ctx, contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()

TSIG鍵の一覧取得



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
	resp, r, err := apiClient.TsigsAPI.GetTsigList(context.Background(), contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.GetTsigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsigList`: GetTsigs
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.GetTsigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTsigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]

### Return type

[**GetTsigs**](GetTsigs.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTsigListCount

> GetCount GetTsigListCount(ctx, contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()

TSIG鍵の件数取得



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
	resp, r, err := apiClient.TsigsAPI.GetTsigListCount(context.Background(), contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsName(keywordsName).KeywordsDescription(keywordsDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.GetTsigListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsigListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.GetTsigListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTsigListCountRequest struct via the builder pattern


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


## PatchTsig

> AsyncResponse PatchTsig(ctx, contractId, tsigId).PatchTsig(patchTsig).Execute()

TSIG鍵の更新



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
	tsigId := int64(789) // int64 | ID
	patchTsig := *openapiclient.NewPatchTsig("Description_example") // PatchTsig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.PatchTsig(context.Background(), contractId, tsigId).PatchTsig(patchTsig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.PatchTsig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTsig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.PatchTsig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 
**tsigId** | **int64** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTsigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchTsig** | [**PatchTsig**](PatchTsig.md) |  | 

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


## PostTsig

> AsyncResponse PostTsig(ctx, contractId).PostTsig(postTsig).Execute()

TSIG鍵の作成



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
	postTsig := *openapiclient.NewPostTsig("Name_example") // PostTsig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigsAPI.PostTsig(context.Background(), contractId).PostTsig(postTsig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigsAPI.PostTsig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTsig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `TsigsAPI.PostTsig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTsigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postTsig** | [**PostTsig**](PostTsig.md) |  | 

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

