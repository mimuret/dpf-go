# \ContractsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContract**](ContractsAPI.md#GetContract) | **Get** /contracts/{ContractId} | DPF契約情報の取得
[**GetContractLabels**](ContractsAPI.md#GetContractLabels) | **Get** /contracts/{ContractId}/labels | DPF契約情報のラベル一覧取得
[**GetContractList**](ContractsAPI.md#GetContractList) | **Get** /contracts | DPF契約情報の一覧取得
[**GetContractListCount**](ContractsAPI.md#GetContractListCount) | **Get** /contracts/count | DPF契約情報の件数取得
[**PatchContract**](ContractsAPI.md#PatchContract) | **Patch** /contracts/{ContractId} | DPF契約情報の更新
[**PutContractLabels**](ContractsAPI.md#PutContractLabels) | **Put** /contracts/{ContractId}/labels | DPF契約情報のラベル一括更新



## GetContract

> GetContract GetContract(ctx, contractId).Execute()

DPF契約情報の取得



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContract(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContract`: GetContract
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContract**](GetContract.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractLabels

> GetContractLabels200Response GetContractLabels(ctx, contractId).Execute()

DPF契約情報のラベル一覧取得



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContractLabels(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractLabels`: GetContractLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContractLabels200Response**](GetContractLabels200Response.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractList

> GetContracts GetContractList(ctx).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsPlan(keywordsPlan).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()

DPF契約情報の一覧取得



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
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsServiceCode := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsPlan := openapiclient.ContractPlan(1) // ContractPlan |  (optional)
	keywordsState := openapiclient.ContractState(1) // ContractState |  (optional)
	keywordsFavorite := openapiclient.ContractFavorite(1) // ContractFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContractList(context.Background()).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsPlan(keywordsPlan).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractList`: GetContracts
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsPlan** | [**ContractPlan**](ContractPlan.md) |  | 
 **keywordsState** | [**ContractState**](ContractState.md) |  | 
 **keywordsFavorite** | [**ContractFavorite**](ContractFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetContracts**](GetContracts.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractListCount

> GetCount GetContractListCount(ctx).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsPlan(keywordsPlan).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()

DPF契約情報の件数取得



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
	type_ := openapiclient.SearchType("AND") // SearchType |  (optional) (default to "AND")
	keywordsFullText := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsServiceCode := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsPlan := openapiclient.ContractPlan(1) // ContractPlan |  (optional)
	keywordsState := openapiclient.ContractState(1) // ContractState |  (optional)
	keywordsFavorite := openapiclient.ContractFavorite(1) // ContractFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContractListCount(context.Background()).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsPlan(keywordsPlan).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractListCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsPlan** | [**ContractPlan**](ContractPlan.md) |  | 
 **keywordsState** | [**ContractState**](ContractState.md) |  | 
 **keywordsFavorite** | [**ContractFavorite**](ContractFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

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


## PatchContract

> AsyncResponse PatchContract(ctx, contractId).PatchContract(patchContract).Execute()

DPF契約情報の更新



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
	patchContract := *openapiclient.NewPatchContract() // PatchContract |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PatchContract(context.Background(), contractId).PatchContract(patchContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PatchContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContract`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PatchContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchContract** | [**PatchContract**](PatchContract.md) |  | 

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


## PutContractLabels

> AsyncResponse PutContractLabels(ctx, contractId).ContractLabels(contractLabels).Execute()

DPF契約情報のラベル一括更新



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
	contractLabels := *openapiclient.NewContractLabels(map[string]string{"key": "Inner_example"}) // ContractLabels |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PutContractLabels(context.Background(), contractId).ContractLabels(contractLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PutContractLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutContractLabels`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PutContractLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutContractLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contractLabels** | [**ContractLabels**](ContractLabels.md) |  | 

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

