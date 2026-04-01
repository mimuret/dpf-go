# ZoneHistoryText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CommittedAt** | **time.Time** |  | 
**Description** | **string** | コメント | [default to ""]
**Operator** | **NullableString** | 編集者 | 
**Text** | **string** |  | 

## Methods

### NewZoneHistoryText

`func NewZoneHistoryText(id int64, committedAt time.Time, description string, operator NullableString, text string, ) *ZoneHistoryText`

NewZoneHistoryText instantiates a new ZoneHistoryText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneHistoryTextWithDefaults

`func NewZoneHistoryTextWithDefaults() *ZoneHistoryText`

NewZoneHistoryTextWithDefaults instantiates a new ZoneHistoryText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneHistoryText) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneHistoryText) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneHistoryText) SetId(v int64)`

SetId sets Id field to given value.


### GetCommittedAt

`func (o *ZoneHistoryText) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *ZoneHistoryText) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *ZoneHistoryText) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.


### GetDescription

`func (o *ZoneHistoryText) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneHistoryText) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneHistoryText) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOperator

`func (o *ZoneHistoryText) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ZoneHistoryText) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ZoneHistoryText) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *ZoneHistoryText) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ZoneHistoryText) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetText

`func (o *ZoneHistoryText) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ZoneHistoryText) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ZoneHistoryText) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


