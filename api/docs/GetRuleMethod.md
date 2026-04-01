# GetRuleMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**RuleMethod**](RuleMethod.md) |  | 

## Methods

### NewGetRuleMethod

`func NewGetRuleMethod(requestId string, result RuleMethod, ) *GetRuleMethod`

NewGetRuleMethod instantiates a new GetRuleMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleMethodWithDefaults

`func NewGetRuleMethodWithDefaults() *GetRuleMethod`

NewGetRuleMethodWithDefaults instantiates a new GetRuleMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRuleMethod) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRuleMethod) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRuleMethod) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetRuleMethod) GetResult() RuleMethod`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRuleMethod) GetResultOk() (*RuleMethod, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRuleMethod) SetResult(v RuleMethod)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


