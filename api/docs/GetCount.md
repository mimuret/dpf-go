# GetCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**GetCountResult**](GetCountResult.md) |  | 

## Methods

### NewGetCount

`func NewGetCount(requestId string, result GetCountResult, ) *GetCount`

NewGetCount instantiates a new GetCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountWithDefaults

`func NewGetCountWithDefaults() *GetCount`

NewGetCountWithDefaults instantiates a new GetCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetCount) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCount) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCount) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetCount) GetResult() GetCountResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCount) GetResultOk() (*GetCountResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCount) SetResult(v GetCountResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


