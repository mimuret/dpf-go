# GetContractLabels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**ContractLabels**](ContractLabels.md) |  | 

## Methods

### NewGetContractLabels200Response

`func NewGetContractLabels200Response(requestId string, result ContractLabels, ) *GetContractLabels200Response`

NewGetContractLabels200Response instantiates a new GetContractLabels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractLabels200ResponseWithDefaults

`func NewGetContractLabels200ResponseWithDefaults() *GetContractLabels200Response`

NewGetContractLabels200ResponseWithDefaults instantiates a new GetContractLabels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetContractLabels200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetContractLabels200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetContractLabels200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetContractLabels200Response) GetResult() ContractLabels`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetContractLabels200Response) GetResultOk() (*ContractLabels, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetContractLabels200Response) SetResult(v ContractLabels)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


