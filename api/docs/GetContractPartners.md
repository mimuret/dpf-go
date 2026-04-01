# GetContractPartners

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]GetContractPartnersResultsInner**](GetContractPartnersResultsInner.md) |  | [default to []]

## Methods

### NewGetContractPartners

`func NewGetContractPartners(requestId string, results []GetContractPartnersResultsInner, ) *GetContractPartners`

NewGetContractPartners instantiates a new GetContractPartners object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractPartnersWithDefaults

`func NewGetContractPartnersWithDefaults() *GetContractPartners`

NewGetContractPartnersWithDefaults instantiates a new GetContractPartners object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetContractPartners) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetContractPartners) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetContractPartners) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetContractPartners) GetResults() []GetContractPartnersResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetContractPartners) GetResultsOk() (*[]GetContractPartnersResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetContractPartners) SetResults(v []GetContractPartnersResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


