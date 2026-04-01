# GetContracts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]Contract**](Contract.md) |  | [default to []]

## Methods

### NewGetContracts

`func NewGetContracts(requestId string, results []Contract, ) *GetContracts`

NewGetContracts instantiates a new GetContracts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractsWithDefaults

`func NewGetContractsWithDefaults() *GetContracts`

NewGetContractsWithDefaults instantiates a new GetContracts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetContracts) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetContracts) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetContracts) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetContracts) GetResults() []Contract`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetContracts) GetResultsOk() (*[]Contract, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetContracts) SetResults(v []Contract)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


