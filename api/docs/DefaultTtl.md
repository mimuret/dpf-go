# DefaultTtl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int32** | TTL | 
**State** | [**DefaultTtlState**](DefaultTtlState.md) |  | 
**Operator** | **NullableString** | 編集者 | 

## Methods

### NewDefaultTtl

`func NewDefaultTtl(value int32, state DefaultTtlState, operator NullableString, ) *DefaultTtl`

NewDefaultTtl instantiates a new DefaultTtl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultTtlWithDefaults

`func NewDefaultTtlWithDefaults() *DefaultTtl`

NewDefaultTtlWithDefaults instantiates a new DefaultTtl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DefaultTtl) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DefaultTtl) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DefaultTtl) SetValue(v int32)`

SetValue sets Value field to given value.


### GetState

`func (o *DefaultTtl) GetState() DefaultTtlState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DefaultTtl) GetStateOk() (*DefaultTtlState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DefaultTtl) SetState(v DefaultTtlState)`

SetState sets State field to given value.


### GetOperator

`func (o *DefaultTtl) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DefaultTtl) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DefaultTtl) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *DefaultTtl) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *DefaultTtl) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


