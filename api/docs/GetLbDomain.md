# GetLbDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**LbDomain**](LbDomain.md) |  | 

## Methods

### NewGetLbDomain

`func NewGetLbDomain(requestId string, result LbDomain, ) *GetLbDomain`

NewGetLbDomain instantiates a new GetLbDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLbDomainWithDefaults

`func NewGetLbDomainWithDefaults() *GetLbDomain`

NewGetLbDomainWithDefaults instantiates a new GetLbDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetLbDomain) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetLbDomain) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetLbDomain) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetLbDomain) GetResult() LbDomain`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetLbDomain) GetResultOk() (*LbDomain, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetLbDomain) SetResult(v LbDomain)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


