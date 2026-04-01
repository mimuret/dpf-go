# \ConfigAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigList**](ConfigAPI.md#GetConfigList) | **Get** /lb_domains/{LbDomainId}/config | 設定の一覧取得
[**PutConfig**](ConfigAPI.md#PutConfig) | **Put** /lb_domains/{LbDomainId}/config | 設定の一括更新



## GetConfigList

> GetConfig GetConfigList(ctx, lbDomainId).Execute()

設定の一覧取得



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
	resp, r, err := apiClient.ConfigAPI.GetConfigList(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigList`: GetConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetConfigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConfig**](GetConfig.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfig

> AsyncResponse PutConfig(ctx, lbDomainId).PutConfig(putConfig).Execute()

設定の一括更新



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
	putConfig := *openapiclient.NewPutConfig([]openapiclient.PutConfigMonitoringsInner{*openapiclient.NewPutConfigMonitoringsInner("ResourceName_example", "Name_example", openapiclient.MonitoringMtype("ping"), "Description_example", *openapiclient.NewPutConfigMonitoringsInnerProps())}, []openapiclient.PutConfigSitesInner{*openapiclient.NewPutConfigSitesInner("ResourceName_example", "Name_example", openapiclient.SiteRrtype("A"), "Description_example", []openapiclient.PutConfigSitesInnerEndpointsInner{*openapiclient.NewPutConfigSitesInnerEndpointsInner("ResourceName_example", "Name_example", "MonitoringTarget_example", "Description_example", int32(123), false, false, false, []openapiclient.EndpointRdataInner{*openapiclient.NewEndpointRdataInner("Value_example")}, []openapiclient.PutConfigSitesInnerEndpointsInnerMonitoringsInner{*openapiclient.NewPutConfigSitesInnerEndpointsInnerMonitoringsInner("ResourceName_example", false)})})}, []openapiclient.PutConfigRulesInner{*openapiclient.NewPutConfigRulesInner("ResourceName_example", "Name_example", "Description_example", []openapiclient.ConfigRuleMethodsInner{*openapiclient.NewConfigRuleMethodsInner(*openapiclient.NewConfigRuleMethodsInnerMethod("ResourceName_example", "Mtype_example", false, "LiveStatus_example", "ReadyStatus_example", []openapiclient.ConfigRuleMethodsInner{*openapiclient.NewConfigRuleMethodsInner(*openapiclient.NewConfigRuleMethodsInnerMethod("ResourceName_example", "Mtype_example", false, "LiveStatus_example", "ReadyStatus_example", []openapiclient.ConfigRuleMethodsInner{}))}))})}) // PutConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.PutConfig(context.Background(), lbDomainId).PutConfig(putConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.PutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.PutConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putConfig** | [**PutConfig**](PutConfig.md) |  | 

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

