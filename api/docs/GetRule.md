# GetRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**Rule**](Rule.md) |  | 

## Methods

### NewGetRule

`func NewGetRule(requestId string, result Rule, ) *GetRule`

NewGetRule instantiates a new GetRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleWithDefaults

`func NewGetRuleWithDefaults() *GetRule`

NewGetRuleWithDefaults instantiates a new GetRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRule) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRule) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRule) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetRule) GetResult() Rule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRule) GetResultOk() (*Rule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRule) SetResult(v Rule)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


