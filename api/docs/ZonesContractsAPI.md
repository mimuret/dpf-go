# \ZonesContractsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractsZoneList**](ZonesContractsAPI.md#GetContractsZoneList) | **Get** /contracts/{ContractId}/zones | DPF契約に紐付くゾーンの一覧取得
[**GetContractsZoneListCount**](ZonesContractsAPI.md#GetContractsZoneListCount) | **Get** /contracts/{ContractId}/zones/count | DPF契約に紐付くゾーンの件数取得
[**PatchContractsZoneCommonConfigs**](ZonesContractsAPI.md#PatchContractsZoneCommonConfigs) | **Patch** /contracts/{ContractId}/zones/common_configs | DPF契約に紐付くゾーンの共通設定の更新



## GetContractsZoneList

> GetZones GetContractsZoneList(ctx, contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()

DPF契約に紐付くゾーンの一覧取得



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
	keywordsNetwork := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := openapiclient.ZonesState(1) // ZonesState |  (optional)
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsZoneProxyEnabled := openapiclient.ZoneProxyEnabled(0) // ZoneProxyEnabled |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesContractsAPI.GetContractsZoneList(context.Background(), contractId).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesContractsAPI.GetContractsZoneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsZoneList`: GetZones
	fmt.Fprintf(os.Stdout, "Response from `ZonesContractsAPI.GetContractsZoneList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsNetwork** | **[]string** |  | [default to []]
 **keywordsState** | [**ZonesState**](ZonesState.md) |  | 
 **keywordsFavorite** | [**ZonesFavorite**](ZonesFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsCommonConfigId** | **[]int64** |  | [default to []]
 **keywordsZoneProxyEnabled** | [**ZoneProxyEnabled**](ZoneProxyEnabled.md) |  | 
 **keywordsLabel** | **[]string** |  | [default to []]

### Return type

[**GetZones**](GetZones.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsZoneListCount

> GetCount GetContractsZoneListCount(ctx, contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()

DPF契約に紐付くゾーンの件数取得



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
	keywordsNetwork := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsState := openapiclient.ZonesState(1) // ZonesState |  (optional)
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsZoneProxyEnabled := openapiclient.ZoneProxyEnabled(0) // ZoneProxyEnabled |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesContractsAPI.GetContractsZoneListCount(context.Background(), contractId).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesContractsAPI.GetContractsZoneListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsZoneListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `ZonesContractsAPI.GetContractsZoneListCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsZoneListCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SearchType**](SearchType.md) |  | [default to &quot;AND&quot;]
 **keywordsFullText** | **[]string** |  | [default to []]
 **keywordsServiceCode** | **[]string** |  | [default to []]
 **keywordsName** | **[]string** |  | [default to []]
 **keywordsNetwork** | **[]string** |  | [default to []]
 **keywordsState** | [**ZonesState**](ZonesState.md) |  | 
 **keywordsFavorite** | [**ZonesFavorite**](ZonesFavorite.md) |  | 
 **keywordsDescription** | **[]string** |  | [default to []]
 **keywordsCommonConfigId** | **[]int64** |  | [default to []]
 **keywordsZoneProxyEnabled** | [**ZoneProxyEnabled**](ZoneProxyEnabled.md) |  | 
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


## PatchContractsZoneCommonConfigs

> AsyncResponse PatchContractsZoneCommonConfigs(ctx, contractId).PatchContractsZones(patchContractsZones).Execute()

DPF契約に紐付くゾーンの共通設定の更新



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
	patchContractsZones := *openapiclient.NewPatchContractsZones(int64(123), []string{"ZoneIds_example"}) // PatchContractsZones |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesContractsAPI.PatchContractsZoneCommonConfigs(context.Background(), contractId).PatchContractsZones(patchContractsZones).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesContractsAPI.PatchContractsZoneCommonConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContractsZoneCommonConfigs`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesContractsAPI.PatchContractsZoneCommonConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContractsZoneCommonConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchContractsZones** | [**PatchContractsZones**](PatchContractsZones.md) |  | 

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

