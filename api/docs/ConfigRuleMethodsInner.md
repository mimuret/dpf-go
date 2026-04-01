# ConfigRuleMethodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | 優先度（親メソッドが failover の場合のみ指定） | [optional] 
**Method** | [**ConfigRuleMethodsInnerMethod**](ConfigRuleMethodsInnerMethod.md) |  | 

## Methods

### NewConfigRuleMethodsInner

`func NewConfigRuleMethodsInner(method ConfigRuleMethodsInnerMethod, ) *ConfigRuleMethodsInner`

NewConfigRuleMethodsInner instantiates a new ConfigRuleMethodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRuleMethodsInnerWithDefaults

`func NewConfigRuleMethodsInnerWithDefaults() *ConfigRuleMethodsInner`

NewConfigRuleMethodsInnerWithDefaults instantiates a new ConfigRuleMethodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *ConfigRuleMethodsInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConfigRuleMethodsInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConfigRuleMethodsInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConfigRuleMethodsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMethod

`func (o *ConfigRuleMethodsInner) GetMethod() ConfigRuleMethodsInnerMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConfigRuleMethodsInner) GetMethodOk() (*ConfigRuleMethodsInnerMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConfigRuleMethodsInner) SetMethod(v ConfigRuleMethodsInnerMethod)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


