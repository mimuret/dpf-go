# \ContractPartnersAPI

All URIs are relative to *https://api.dns-platform.jp/dpf/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractPartnerList**](ContractPartnersAPI.md#GetContractPartnerList) | **Get** /contracts/{ContractId}/contract_partners | DPF連携サービスの一覧取得



## GetContractPartnerList

> GetContractPartners GetContractPartnerList(ctx, contractId).Execute()

DPF連携サービスの一覧取得



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
	resp, r, err := apiClient.ContractPartnersAPI.GetContractPartnerList(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractPartnersAPI.GetContractPartnerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractPartnerList`: GetContractPartners
	fmt.Fprintf(os.Stdout, "Response from `ContractPartnersAPI.GetContractPartnerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractPartnerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContractPartners**](GetContractPartners.md)

### Authorization

[DPFViewer](../README.md#DPFViewer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

