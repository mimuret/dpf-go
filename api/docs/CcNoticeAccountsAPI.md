# \CcNoticeAccountsAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCcNoticeAccount**](CcNoticeAccountsAPI.md#DeleteCcNoticeAccount) | **Delete** /common_configs/{CommonConfigId}/cc_notice_accounts/{CcNoticeAccountResourceName} | 通知先アカウント設定の削除
[**GetCcNoticeAccount**](CcNoticeAccountsAPI.md#GetCcNoticeAccount) | **Get** /common_configs/{CommonConfigId}/cc_notice_accounts/{CcNoticeAccountResourceName} | 通知先アカウントの取得
[**GetCcNoticeAccountList**](CcNoticeAccountsAPI.md#GetCcNoticeAccountList) | **Get** /common_configs/{CommonConfigId}/cc_notice_accounts | 通知先アカウント設定の一覧取得
[**PatchCcNoticeAccount**](CcNoticeAccountsAPI.md#PatchCcNoticeAccount) | **Patch** /common_configs/{CommonConfigId}/cc_notice_accounts/{CcNoticeAccountResourceName} | 通知先アカウント設定の更新
[**PostCcNoticeAccount**](CcNoticeAccountsAPI.md#PostCcNoticeAccount) | **Post** /common_configs/{CommonConfigId}/cc_notice_accounts | 通知先アカウント設定の作成



## DeleteCcNoticeAccount

> AsyncResponse DeleteCcNoticeAccount(ctx, commonConfigId, ccNoticeAccountResourceName).Execute()

通知先アカウント設定の削除



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccNoticeAccountResourceName := "ccNoticeAccountResourceName_example" // string | GET cc_notice_accounts Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcNoticeAccountsAPI.DeleteCcNoticeAccount(context.Background(), commonConfigId, ccNoticeAccountResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcNoticeAccountsAPI.DeleteCcNoticeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCcNoticeAccount`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcNoticeAccountsAPI.DeleteCcNoticeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccNoticeAccountResourceName** | **string** | GET cc_notice_accounts Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCcNoticeAccountRequest struct via the builder pattern


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


## GetCcNoticeAccount

> GetCcNoticeAccount GetCcNoticeAccount(ctx, commonConfigId, ccNoticeAccountResourceName).Execute()

通知先アカウントの取得



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccNoticeAccountResourceName := "ccNoticeAccountResourceName_example" // string | GET cc_notice_accounts Schemaにおける resource_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcNoticeAccountsAPI.GetCcNoticeAccount(context.Background(), commonConfigId, ccNoticeAccountResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcNoticeAccountsAPI.GetCcNoticeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcNoticeAccount`: GetCcNoticeAccount
	fmt.Fprintf(os.Stdout, "Response from `CcNoticeAccountsAPI.GetCcNoticeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccNoticeAccountResourceName** | **string** | GET cc_notice_accounts Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcNoticeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCcNoticeAccount**](GetCcNoticeAccount.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCcNoticeAccountList

> GetCcNoticeAccounts GetCcNoticeAccountList(ctx, commonConfigId).Execute()

通知先アカウント設定の一覧取得



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcNoticeAccountsAPI.GetCcNoticeAccountList(context.Background(), commonConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcNoticeAccountsAPI.GetCcNoticeAccountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCcNoticeAccountList`: GetCcNoticeAccounts
	fmt.Fprintf(os.Stdout, "Response from `CcNoticeAccountsAPI.GetCcNoticeAccountList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCcNoticeAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCcNoticeAccounts**](GetCcNoticeAccounts.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCcNoticeAccount

> AsyncResponse PatchCcNoticeAccount(ctx, commonConfigId, ccNoticeAccountResourceName).PatchCcNoticeAccount(patchCcNoticeAccount).Execute()

通知先アカウント設定の更新



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	ccNoticeAccountResourceName := "ccNoticeAccountResourceName_example" // string | GET cc_notice_accounts Schemaにおける resource_name
	patchCcNoticeAccount := *openapiclient.NewPatchCcNoticeAccount() // PatchCcNoticeAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcNoticeAccountsAPI.PatchCcNoticeAccount(context.Background(), commonConfigId, ccNoticeAccountResourceName).PatchCcNoticeAccount(patchCcNoticeAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcNoticeAccountsAPI.PatchCcNoticeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCcNoticeAccount`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcNoticeAccountsAPI.PatchCcNoticeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 
**ccNoticeAccountResourceName** | **string** | GET cc_notice_accounts Schemaにおける resource_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCcNoticeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchCcNoticeAccount** | [**PatchCcNoticeAccount**](PatchCcNoticeAccount.md) |  | 

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


## PostCcNoticeAccount

> AsyncResponse PostCcNoticeAccount(ctx, commonConfigId).PostCcNoticeAccount(postCcNoticeAccount).Execute()

通知先アカウント設定の作成



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
	commonConfigId := int64(789) // int64 | GET common_configs Schemaにおける id
	postCcNoticeAccount := *openapiclient.NewPostCcNoticeAccount("Name_example", openapiclient.CcNoticeAccountLang("ja")) // PostCcNoticeAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CcNoticeAccountsAPI.PostCcNoticeAccount(context.Background(), commonConfigId).PostCcNoticeAccount(postCcNoticeAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CcNoticeAccountsAPI.PostCcNoticeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCcNoticeAccount`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CcNoticeAccountsAPI.PostCcNoticeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commonConfigId** | **int64** | GET common_configs Schemaにおける id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCcNoticeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCcNoticeAccount** | [**PostCcNoticeAccount**](PostCcNoticeAccount.md) |  | 

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

