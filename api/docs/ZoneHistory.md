# ZoneHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CommittedAt** | **time.Time** |  | 
**Description** | **string** | コメント | [default to ""]
**Operator** | **NullableString** | 編集者 | 

## Methods

### NewZoneHistory

`func NewZoneHistory(id int64, committedAt time.Time, description string, operator NullableString, ) *ZoneHistory`

NewZoneHistory instantiates a new ZoneHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneHistoryWithDefaults

`func NewZoneHistoryWithDefaults() *ZoneHistory`

NewZoneHistoryWithDefaults instantiates a new ZoneHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneHistory) SetId(v int64)`

SetId sets Id field to given value.


### GetCommittedAt

`func (o *ZoneHistory) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *ZoneHistory) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *ZoneHistory) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.


### GetDescription

`func (o *ZoneHistory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneHistory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneHistory) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOperator

`func (o *ZoneHistory) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ZoneHistory) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ZoneHistory) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *ZoneHistory) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ZoneHistory) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


