# QpsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | 集計対象月(YYYY-MM) | 
**Qps** | **int32** |  | 

## Methods

### NewQpsValue

`func NewQpsValue(month string, qps int32, ) *QpsValue`

NewQpsValue instantiates a new QpsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQpsValueWithDefaults

`func NewQpsValueWithDefaults() *QpsValue`

NewQpsValueWithDefaults instantiates a new QpsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *QpsValue) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *QpsValue) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *QpsValue) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetQps

`func (o *QpsValue) GetQps() int32`

GetQps returns the Qps field if non-nil, zero value otherwise.

### GetQpsOk

`func (o *QpsValue) GetQpsOk() (*int32, bool)`

GetQpsOk returns a tuple with the Qps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQps

`func (o *QpsValue) SetQps(v int32)`

SetQps sets Qps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


