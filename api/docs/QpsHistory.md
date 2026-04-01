# QpsHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCode** | **string** |  | 
**Name** | **NullableString** |  | 
**Values** | [**[]QpsValue**](QpsValue.md) |  | 

## Methods

### NewQpsHistory

`func NewQpsHistory(serviceCode string, name NullableString, values []QpsValue, ) *QpsHistory`

NewQpsHistory instantiates a new QpsHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQpsHistoryWithDefaults

`func NewQpsHistoryWithDefaults() *QpsHistory`

NewQpsHistoryWithDefaults instantiates a new QpsHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCode

`func (o *QpsHistory) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *QpsHistory) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *QpsHistory) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetName

`func (o *QpsHistory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QpsHistory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QpsHistory) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *QpsHistory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QpsHistory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValues

`func (o *QpsHistory) GetValues() []QpsValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *QpsHistory) GetValuesOk() (*[]QpsValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *QpsHistory) SetValues(v []QpsValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


