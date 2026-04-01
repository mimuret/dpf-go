# \ZonesAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZoneChanges**](ZonesAPI.md#DeleteZoneChanges) | **Delete** /zones/{ZoneId}/changes | 編集中レコードの一括取消
[**GetZone**](ZonesAPI.md#GetZone) | **Get** /zones/{ZoneId} | ゾーンの取得
[**GetZoneContract**](ZonesAPI.md#GetZoneContract) | **Get** /zones/{ZoneId}/contract | ゾーンに紐付くDPF契約情報の取得
[**GetZoneLabels**](ZonesAPI.md#GetZoneLabels) | **Get** /zones/{ZoneId}/labels | ゾーンのラベル一覧取得
[**GetZoneList**](ZonesAPI.md#GetZoneList) | **Get** /zones | ゾーンの一覧取得
[**GetZoneListCount**](ZonesAPI.md#GetZoneListCount) | **Get** /zones/count | ゾーンの件数取得
[**GetZoneManagedDnsServers**](ZonesAPI.md#GetZoneManagedDnsServers) | **Get** /zones/{ZoneId}/managed_dns_servers | マネージドDNSサーバの一覧取得
[**PatchZone**](ZonesAPI.md#PatchZone) | **Patch** /zones/{ZoneId} | ゾーンの更新
[**PatchZoneAtomicChanges**](ZonesAPI.md#PatchZoneAtomicChanges) | **Patch** /zones/{ZoneId}/atomic_changes | レコードの一括更新とゾーン反映
[**PatchZoneChanges**](ZonesAPI.md#PatchZoneChanges) | **Patch** /zones/{ZoneId}/changes | 編集中レコードのゾーン反映
[**PutZoneLabels**](ZonesAPI.md#PutZoneLabels) | **Put** /zones/{ZoneId}/labels | ゾーンのラベル一括更新



## DeleteZoneChanges

> AsyncResponse DeleteZoneChanges(ctx, zoneId).Execute()

編集中レコードの一括取消



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
	zoneId := "zoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.DeleteZoneChanges(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.DeleteZoneChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteZoneChanges`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.DeleteZoneChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneChangesRequest struct via the builder pattern


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


## GetZone

> GetZone GetZone(ctx, zoneId).Execute()

ゾーンの取得



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
	zoneId := "zoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZone(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZone`: GetZone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetZone**](GetZone.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneContract

> GetContract GetZoneContract(ctx, zoneId).Execute()

ゾーンに紐付くDPF契約情報の取得



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
	zoneId := "zoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZoneContract(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZoneContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneContract`: GetContract
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZoneContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneContractRequest struct via the builder pattern


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


## GetZoneLabels

> GetZoneLabels200Response GetZoneLabels(ctx, zoneId).Execute()

ゾーンのラベル一覧取得



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
	zoneId := "zoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZoneLabels(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZoneLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneLabels`: GetZoneLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZoneLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetZoneLabels200Response**](GetZoneLabels200Response.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneList

> GetZones GetZoneList(ctx).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()

ゾーンの一覧取得



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
	keywordsState := openapiclient.ZonesState(1) // ZonesState |  (optional)
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsZoneProxyEnabled := openapiclient.ZoneProxyEnabled(0) // ZoneProxyEnabled |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZoneList(context.Background()).Type_(type_).Offset(offset).Limit(limit).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZoneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneList`: GetZones
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneListRequest struct via the builder pattern


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


## GetZoneListCount

> GetCount GetZoneListCount(ctx).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()

ゾーンの件数取得



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
	keywordsState := openapiclient.ZonesState(1) // ZonesState |  (optional)
	keywordsFavorite := openapiclient.ZonesFavorite(1) // ZonesFavorite |  (optional)
	keywordsDescription := []string{"Inner_example"} // []string |  (optional) (default to [])
	keywordsCommonConfigId := []int64{int64(123)} // []int64 |  (optional) (default to [])
	keywordsZoneProxyEnabled := openapiclient.ZoneProxyEnabled(0) // ZoneProxyEnabled |  (optional)
	keywordsLabel := []string{"Inner_example"} // []string |  (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZoneListCount(context.Background()).Type_(type_).KeywordsFullText(keywordsFullText).KeywordsServiceCode(keywordsServiceCode).KeywordsName(keywordsName).KeywordsNetwork(keywordsNetwork).KeywordsState(keywordsState).KeywordsFavorite(keywordsFavorite).KeywordsDescription(keywordsDescription).KeywordsCommonConfigId(keywordsCommonConfigId).KeywordsZoneProxyEnabled(keywordsZoneProxyEnabled).KeywordsLabel(keywordsLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZoneListCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneListCount`: GetCount
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZoneListCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneListCountRequest struct via the builder pattern


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


## GetZoneManagedDnsServers

> GetManagedServers GetZoneManagedDnsServers(ctx, zoneId).Execute()

マネージドDNSサーバの一覧取得



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
	zoneId := "zoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZoneManagedDnsServers(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZoneManagedDnsServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneManagedDnsServers`: GetManagedServers
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZoneManagedDnsServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneManagedDnsServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetManagedServers**](GetManagedServers.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchZone

> AsyncResponse PatchZone(ctx, zoneId).PatchZone(patchZone).Execute()

ゾーンの更新



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
	zoneId := "zoneId_example" // string | ID
	patchZone := *openapiclient.NewPatchZone() // PatchZone |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.PatchZone(context.Background(), zoneId).PatchZone(patchZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PatchZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchZone`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PatchZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchZone** | [**PatchZone**](PatchZone.md) |  | 

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


## PatchZoneAtomicChanges

> AsyncResponse PatchZoneAtomicChanges(ctx, zoneId).PatchZoneAtomicChanges(patchZoneAtomicChanges).Execute()

レコードの一括更新とゾーン反映



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
	zoneId := "zoneId_example" // string | ID
	patchZoneAtomicChanges := *openapiclient.NewPatchZoneAtomicChanges([]openapiclient.OverwriteRecordsInner{*openapiclient.NewOverwriteRecordsInner("Name_example", NullableInt32(123), openapiclient.RecordsRrtype("A"), []openapiclient.RecordsRdataInner{*openapiclient.NewRecordsRdataInner()}, "Description_example", map[string]string{"key": "Inner_example"})}) // PatchZoneAtomicChanges |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.PatchZoneAtomicChanges(context.Background(), zoneId).PatchZoneAtomicChanges(patchZoneAtomicChanges).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PatchZoneAtomicChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchZoneAtomicChanges`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PatchZoneAtomicChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneAtomicChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchZoneAtomicChanges** | [**PatchZoneAtomicChanges**](PatchZoneAtomicChanges.md) |  | 

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


## PatchZoneChanges

> AsyncResponse PatchZoneChanges(ctx, zoneId).PatchZoneCommit(patchZoneCommit).Execute()

編集中レコードのゾーン反映



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
	zoneId := "zoneId_example" // string | ID
	patchZoneCommit := *openapiclient.NewPatchZoneCommit() // PatchZoneCommit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.PatchZoneChanges(context.Background(), zoneId).PatchZoneCommit(patchZoneCommit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PatchZoneChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchZoneChanges`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PatchZoneChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchZoneCommit** | [**PatchZoneCommit**](PatchZoneCommit.md) |  | 

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


## PutZoneLabels

> AsyncResponse PutZoneLabels(ctx, zoneId).ZoneLabels(zoneLabels).Execute()

ゾーンのラベル一括更新



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
	zoneId := "zoneId_example" // string | ID
	zoneLabels := *openapiclient.NewZoneLabels(map[string]string{"key": "Inner_example"}) // ZoneLabels |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.PutZoneLabels(context.Background(), zoneId).ZoneLabels(zoneLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PutZoneLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutZoneLabels`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PutZoneLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneLabels** | [**ZoneLabels**](ZoneLabels.md) |  | 

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

