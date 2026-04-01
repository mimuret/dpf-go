# \LbDomainsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLbDomain**](LbDomainsAPI.md#GetLbDomain) | **Get** /lb_domains/{LbDomainId} | LBドメインの取得
[**GetLbDomainContract**](LbDomainsAPI.md#GetLbDomainContract) | **Get** /lb_domains/{LbDomainId}/contract | LBドメインに紐付くDPF契約情報の取得
[**GetLbDomainLabels**](LbDomainsAPI.md#GetLbDomainLabels) | **Get** /lb_domains/{LbDomainId}/labels | LBドメインのラベル一覧取得
[**GetLbDomainList**](LbDomainsAPI.md#GetLbDomainList) | **Get** /lb_domains | LBドメインの一覧取得
[**GetLbDomainListCount**](LbDomainsAPI.md#GetLbDomainListCount) | **Get** /lb_domains/count | LBドメインの件数取得
[**PatchLbDomain**](LbDomainsAPI.md#PatchLbDomain) | **Patch** /lb_domains/{LbDomainId} | LBドメインの更新
[**PutLbDomainLabels**](LbDomainsAPI.md#PutLbDomainLabels) | **Put** /lb_domains/{LbDomainId}/labels | LBドメインのラベル一括更新



## GetLbDomain

> GetLbDomain GetLbDomain(ctx, lbDomainId).Execute()

LBドメインの取得



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
	resp, r, err := apiClient.LbDomainsAPI.GetLbDomain(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.GetLbDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomain`: GetLbDomain
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.GetLbDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLbDomain**](GetLbDomain.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLbDomainContract

> GetContract GetLbDomainContract(ctx, lbDomainId).Execute()

LBドメインに紐付くDPF契約情報の取得



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
	resp, r, err := apiClient.LbDomainsAPI.GetLbDomainContract(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.GetLbDomainContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainContract`: GetContract
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.GetLbDomainContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainContractRequest struct via the builder pattern


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


## GetLbDomainLabels

> GetLbDomainLabels GetLbDomainLabels(ctx, lbDomainId).Execute()

LBドメインのラベル一覧取得



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
	resp, r, err := apiClient.LbDomainsAPI.GetLbDomainLabels(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.GetLbDomainLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainLabels`: GetLbDomainLabels
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.GetLbDomainLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLbDomainLabels**](GetLbDomainLabels.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLbDomainList

> GetLbDomains GetLbDomainList(ctx).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()

LBドメインの一覧取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsFavorite := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsAPI.GetLbDomainList(context.Background()).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.GetLbDomainList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainList`: GetLbDomains
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.GetLbDomainList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainListRequest struct via the builder pattern


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


## GetLbDomainListCount

> GetCount GetLbDomainListCount(ctx).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()

LBドメインの件数取得



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
	keywordsName := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsFavorite := []int32{int32(123)} // []int32 |  (optional) (default to [])
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsAPI.GetLbDomainListCount(context.Background()).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.GetLbDomainListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLbDomainListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.GetLbDomainListCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLbDomainListCountRequest struct via the builder pattern


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


## PatchLbDomain

> AsyncResponse PatchLbDomain(ctx, lbDomainId).PatchLbDomain(patchLbDomain).Execute()

LBドメインの更新



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
	patchLbDomain := *openapiclient.NewPatchLbDomain() // PatchLbDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsAPI.PatchLbDomain(context.Background(), lbDomainId).PatchLbDomain(patchLbDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.PatchLbDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLbDomain`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.PatchLbDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLbDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchLbDomain** | [**PatchLbDomain**](PatchLbDomain.md) |  | 

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


## PutLbDomainLabels

> AsyncResponse PutLbDomainLabels(ctx, lbDomainId).Body(body).Execute()

LBドメインのラベル一括更新



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
	body := LbDomainLabels(987) // LbDomainLabels |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LbDomainsAPI.PutLbDomainLabels(context.Background(), lbDomainId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LbDomainsAPI.PutLbDomainLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLbDomainLabels`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `LbDomainsAPI.PutLbDomainLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLbDomainLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **LbDomainLabels** |  | 

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

