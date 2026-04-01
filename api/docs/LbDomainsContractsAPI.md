# \LbDomainsContractsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractsLbDomainList**](LbDomainsContractsAPI.md#GetContractsLbDomainList) | **Get** /contracts/{ContractId}/lb_domains | DPF契約に紐付くLBドメインの一覧取得
[**GetContractsLbDomainListCount**](LbDomainsContractsAPI.md#GetContractsLbDomainListCount) | **Get** /contracts/{ContractId}/lb_domains/count | DPF契約に紐付くLBドメインの件数取得
[**PatchContractsLbDomainCommonConfigs**](LbDomainsContractsAPI.md#PatchContractsLbDomainCommonConfigs) | **Patch** /contracts/{ContractId}/lb_domains/common_configs | DPF契約に紐付くLBドメインの共通設定の更新



## GetContractsLbDomainList

> GetLbDomains GetContractsLbDomainList(ctx, contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()

DPF契約に紐付くLBドメインの一覧取得



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
	keywordsServiceCode := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsFavorite := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsContractsAPI.GetContractsLbDomainList(context.Background(), contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsContractsAPI.GetContractsLbDomainList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsLbDomainList`: GetLbDomains
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsContractsAPI.GetContractsLbDomainList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsLbDomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsState** | **[]int32** |  | [default to []]
 **keywordsFavorite** | **[]int32** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsCommonConfigId** | **[]int64** |  | [default to []]
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetLbDomains**](GetLbDomains.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsLbDomainListCount

> GetCount GetContractsLbDomainListCount(ctx, contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()

DPF契約に紐付くLBドメインの件数取得



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
	keywordsServiceCode := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsFavorite := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsContractsAPI.GetContractsLbDomainListCount(context.Background(), contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsContractsAPI.GetContractsLbDomainListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsLbDomainListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsContractsAPI.GetContractsLbDomainListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsLbDomainListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsState** | **[]int32** |  | [default to []]
 **keywordsFavorite** | **[]int32** |  | [default to []]
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsCommonConfigId** | **[]int64** |  | [default to []]
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


## PatchContractsLbDomainCommonConfigs

> AsyncResponse PatchContractsLbDomainCommonConfigs(ctx, contractId).PatchContractsLbDomains(patchContractsLbDomains).Execute()

DPF契約に紐付くLBドメインの共通設定の更新



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
	patchContractsLbDomains := *openapiclient.NewPatchContractsLbDomains(int64(123), []string{"LbDomainIds_example"}) // PatchContractsLbDomains |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsContractsAPI.PatchContractsLbDomainCommonConfigs(context.Background(), contractId).PatchContractsLbDomains(patchContractsLbDomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsContractsAPI.PatchContractsLbDomainCommonConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContractsLbDomainCommonConfigs`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsContractsAPI.PatchContractsLbDomainCommonConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContractsLbDomainCommonConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchContractsLbDomains** | [**PatchContractsLbDomains**](PatchContractsLbDomains.md) |  | 

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

