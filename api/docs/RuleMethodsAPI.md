# \RuleMethodsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRuleMethod**](RuleMethodsAPI.md#DeleteRuleMethod) | **Delete** /lb_domains/{LbDomainId}/rules/{RuleResourceName}/rule_methods/{RuleMethodResourceName} | ルールメソッドの削除
[**GetRuleMethod**](RuleMethodsAPI.md#GetRuleMethod) | **Get** /lb_domains/{LbDomainId}/rules/{RuleResourceName}/rule_methods/{RuleMethodResourceName} | ルールメソッドの取得
[**GetRuleMethodList**](RuleMethodsAPI.md#GetRuleMethodList) | **Get** /lb_domains/{LbDomainId}/rules/{RuleResourceName}/rule_methods | ルールメソッドの一覧取得
[**PatchRuleMethod**](RuleMethodsAPI.md#PatchRuleMethod) | **Patch** /lb_domains/{LbDomainId}/rules/{RuleResourceName}/rule_methods/{RuleMethodResourceName} | ルールメソッドの更新
[**PostRuleMethod**](RuleMethodsAPI.md#PostRuleMethod) | **Post** /lb_domains/{LbDomainId}/rules/{RuleResourceName}/rule_methods | ルールメソッドの作成



## DeleteRuleMethod

> AsyncResponse DeleteRuleMethod(ctx, lbDomainId, ruleResourceName, ruleMethodResourceName).Execute()

ルールメソッドの削除



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
	ruleMethodResourceName := "ruleMethodResourceName_example" // string | GET rule_methods Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleMethodsAPI.DeleteRuleMethod(context.Background(), lbDomainId, ruleResourceName, ruleMethodResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleMethodsAPI.DeleteRuleMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleMethod`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RuleMethodsAPI.DeleteRuleMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 
**ruleMethodResourceName** | **string** | GET rule_methods Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleMethodRequest struct via the builder pattern


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


## GetRuleMethod

> GetRuleMethod GetRuleMethod(ctx, lbDomainId, ruleResourceName, ruleMethodResourceName).Execute()

ルールメソッドの取得



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
	ruleMethodResourceName := "ruleMethodResourceName_example" // string | GET rule_methods Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleMethodsAPI.GetRuleMethod(context.Background(), lbDomainId, ruleResourceName, ruleMethodResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleMethodsAPI.GetRuleMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleMethod`: GetRuleMethod
	fmt.Fprintf(os.Stdout, "Response from `RuleMethodsAPI.GetRuleMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 
**ruleMethodResourceName** | **string** | GET rule_methods Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRuleMethod**](GetRuleMethod.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleMethodList

> GetRuleMethods GetRuleMethodList(ctx, lbDomainId, ruleResourceName).Execute()

ルールメソッドの一覧取得



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
	resp, r, err := apiClient.RuleMethodsAPI.GetRuleMethodList(context.Background(), lbDomainId, ruleResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleMethodsAPI.GetRuleMethodList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleMethodList`: GetRuleMethods
	fmt.Fprintf(os.Stdout, "Response from `RuleMethodsAPI.GetRuleMethodList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleMethodListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRuleMethods**](GetRuleMethods.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRuleMethod

> AsyncResponse PatchRuleMethod(ctx, lbDomainId, ruleResourceName, ruleMethodResourceName).PatchRuleMethod(patchRuleMethod).Execute()

ルールメソッドの更新



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
	ruleMethodResourceName := "ruleMethodResourceName_example" // string | GET rule_methods Schemaにおける resource_name
	patchRuleMethod := *openapiclient.NewPatchRuleMethod() // PatchRuleMethod |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleMethodsAPI.PatchRuleMethod(context.Background(), lbDomainId, ruleResourceName, ruleMethodResourceName).PatchRuleMethod(patchRuleMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleMethodsAPI.PatchRuleMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRuleMethod`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RuleMethodsAPI.PatchRuleMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 
**ruleMethodResourceName** | **string** | GET rule_methods Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRuleMethod** | [**PatchRuleMethod**](PatchRuleMethod.md) |  | 

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


## PostRuleMethod

> AsyncResponse PostRuleMethod(ctx, lbDomainId, ruleResourceName).PostRuleMethod(postRuleMethod).Execute()

ルールメソッドの作成



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
	postRuleMethod := *openapiclient.NewPostRuleMethod(*openapiclient.NewPostRuleMethodMethod("Mtype_example")) // PostRuleMethod |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleMethodsAPI.PostRuleMethod(context.Background(), lbDomainId, ruleResourceName).PostRuleMethod(postRuleMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleMethodsAPI.PostRuleMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRuleMethod`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `RuleMethodsAPI.PostRuleMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbDomainId** | **string** | GET lb_domains Schemaにおける id | 
**ruleResourceName** | **string** | GET rules Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRuleMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRuleMethod** | [**PostRuleMethod**](PostRuleMethod.md) |  | 

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

