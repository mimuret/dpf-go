# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | name | 
**Ttl** | **NullableInt32** | TTL | 
**Rrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
**Rdata** | [**[]RecordsRdataInner**](RecordsRdataInner.md) | レコードの値 | 
**Labels** | **map[string]string** | ラベル | [default to {}]
**State** | [**RecordsState**](RecordsState.md) |  | 
**Description** | **string** | コメント | [default to ""]
**Operator** | **NullableString** | 編集者 | 

## Methods

### NewRecord

`func NewRecord(id string, name string, ttl NullableInt32, rrtype RecordsRrtype, rdata []RecordsRdataInner, labels map[string]string, state RecordsState, description string, operator NullableString, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Record) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Record) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Record) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Record) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Record) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Record) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *Record) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Record) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Record) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *Record) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *Record) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetRrtype

`func (o *Record) GetRrtype() RecordsRrtype`

GetRrtype returns the Rrtype field if non-nil, zero value otherwise.

### GetRrtypeOk

`func (o *Record) GetRrtypeOk() (*RecordsRrtype, bool)`

GetRrtypeOk returns a tuple with the Rrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrtype

`func (o *Record) SetRrtype(v RecordsRrtype)`

SetRrtype sets Rrtype field to given value.


### GetRdata

`func (o *Record) GetRdata() []RecordsRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *Record) GetRdataOk() (*[]RecordsRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *Record) SetRdata(v []RecordsRdataInner)`

SetRdata sets Rdata field to given value.


### GetLabels

`func (o *Record) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Record) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Record) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetState

`func (o *Record) GetState() RecordsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Record) GetStateOk() (*RecordsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Record) SetState(v RecordsState)`

SetState sets State field to given value.


### GetDescription

`func (o *Record) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Record) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Record) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOperator

`func (o *Record) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Record) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Record) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *Record) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *Record) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


