# \DelegationsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDelegationList**](DelegationsAPI.md#GetDelegationList) | **Get** /delegations | ネームサーバ申請候補の一覧取得
[**GetDelegationListCount**](DelegationsAPI.md#GetDelegationListCount) | **Get** /delegations/count | ネームサーバ申請候補の件数取得
[**PostDelegation**](DelegationsAPI.md#PostDelegation) | **Post** /delegations | ネームサーバ申請



## GetDelegationList

> GetDelegations GetDelegationList(ctx).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsRequested(keywordsRequested).KeywordsLabel(keywordsLabel).Execute()

ネームサーバ申請候補の一覧取得



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
	keywordsNetwork := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequested := openapiclient.DelegationsRequested(0) // DelegationsRequested |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.GetDelegationList(context.Background()).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsRequested(keywordsRequested).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.GetDelegationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelegationList`: GetDelegations
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.GetDelegationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsNetwork** | **[]string** |  | [default to []]
 **keywordsFavorite** | [**ZonesFavorite**](ZonesFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsRequested** | [**DelegationsRequested**](DelegationsRequested.md) |  | 
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetDelegations**](GetDelegations.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelegationListCount

> GetCount GetDelegationListCount(ctx).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsRequested(keywordsRequested).KeywordsLabel(keywordsLabel).Execute()

ネームサーバ申請候補の件数取得



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
	keywordsNetwork := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsRequested := openapiclient.DelegationsRequested(0) // DelegationsRequested |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.GetDelegationListCount(context.Background()).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsRequested(keywordsRequested).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.GetDelegationListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelegationListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.GetDelegationListCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegationListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsNetwork** | **[]string** |  | [default to []]
 **keywordsFavorite** | [**ZonesFavorite**](ZonesFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsRequested** | [**DelegationsRequested**](DelegationsRequested.md) |  | 
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


## PostDelegation

> AsyncResponse PostDelegation(ctx).PostDelegations(postDelegations).Execute()

ネームサーバ申請



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
	postDelegations := *openapiclient.NewPostDelegations([]string{"ZoneIds_example"}) // PostDelegations |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.PostDelegation(context.Background()).PostDelegations(postDelegations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.PostDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDelegation`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.PostDelegation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDelegationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDelegations** | [**PostDelegations**](PostDelegations.md) |  | 

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

