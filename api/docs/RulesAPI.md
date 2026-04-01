# \RulesAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRule**](RulesAPI.md#DeleteRule) | **Delete** /lb_domains/{LbDomainId}/rules/{RuleResourceName} | ルールの削除
[**GetRule**](RulesAPI.md#GetRule) | **Get** /lb_domains/{LbDomainId}/rules/{RuleResourceName} | ルールの取得
[**GetRuleList**](RulesAPI.md#GetRuleList) | **Get** /lb_domains/{LbDomainId}/rules | ルールの一覧取得
[**PatchRule**](RulesAPI.md#PatchRule) | **Patch** /lb_domains/{LbDomainId}/rules/{RuleResourceName} | ルールの更新
[**PostRule**](RulesAPI.md#PostRule) | **Post** /lb_domains/{LbDomainId}/rules | ルールの作成



## DeleteRule

> AsyncResponse DeleteRule(ctx, lbDomainId, ruleResourceName).Execute()

ルールの削除



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
	ruleResourceName := "ruleResourceName_example" // string | GET rules Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRule(context.Background(), lbDomainId, ruleResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRule`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetRule

> GetRule GetRule(ctx, lbDomainId, ruleResourceName).Execute()

ルールの取得



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
	ruleResourceName := "ruleResourceName_example" // string | GET rules Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRule(context.Background(), lbDomainId, ruleResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRule`: GetRule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRule**](GetRule.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleList

> GetRules GetRuleList(ctx, lbDomainId).Execute()

ルールの一覧取得



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
	resp, r, err := apiClient.RulesAPI.GetRuleList(context.Background(), lbDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleList`: GetRules
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRules**](GetRules.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRule

> AsyncResponse PatchRule(ctx, lbDomainId, ruleResourceName).PatchRule(patchRule).Execute()

ルールの更新



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
	ruleResourceName := "ruleResourceName_example" // string | GET rules Schemaにおける resource_name
	patchRule := *openapiclient.NewPatchRule() // PatchRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.PatchRule(context.Background(), lbDomainId, ruleResourceName).PatchRule(patchRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.PatchRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRule`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.PatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchRule** | [**PatchRule**](PatchRule.md) |  | 

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


## PostRule

> AsyncResponse PostRule(ctx, lbDomainId).PostRule(postRule).Execute()

ルールの作成



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
	postRule := *openapiclient.NewPostRule("Name_example") // PostRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.PostRule(context.Background(), lbDomainId).PostRule(postRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.PostRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRule`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.PostRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRule** | [**PostRule**](PostRule.md) |  | 

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

