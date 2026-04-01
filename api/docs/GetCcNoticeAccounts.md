# GetCcNoticeAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]CcNoticeAccount**](CcNoticeAccount.md) |  | [default to []]

## Methods

### NewGetCcNoticeAccounts

`func NewGetCcNoticeAccounts(requestId string, results []CcNoticeAccount, ) *GetCcNoticeAccounts`

NewGetCcNoticeAccounts instantiates a new GetCcNoticeAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCcNoticeAccountsWithDefaults

`func NewGetCcNoticeAccountsWithDefaults() *GetCcNoticeAccounts`

NewGetCcNoticeAccountsWithDefaults instantiates a new GetCcNoticeAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetCcNoticeAccounts) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCcNoticeAccounts) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCcNoticeAccounts) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetCcNoticeAccounts) GetResults() []CcNoticeAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetCcNoticeAccounts) GetResultsOk() (*[]CcNoticeAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetCcNoticeAccounts) SetResults(v []CcNoticeAccount)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


