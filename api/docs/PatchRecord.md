# PatchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **NullableInt32** | TTL | [optional] 
**Rdata** | Pointer to [**[]RecordsRdataInner**](RecordsRdataInner.md) | レコードの値 | [optional] 
**Labels** | Pointer to **map[string]string** | ラベル | [optional] [default to {}]
**Description** | Pointer to **string** | コメント | [optional] [default to ""]

## Methods

### NewPatchRecord

`func NewPatchRecord() *PatchRecord`

NewPatchRecord instantiates a new PatchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRecordWithDefaults

`func NewPatchRecordWithDefaults() *PatchRecord`

NewPatchRecordWithDefaults instantiates a new PatchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *PatchRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PatchRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PatchRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PatchRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *PatchRecord) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *PatchRecord) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetRdata

`func (o *PatchRecord) GetRdata() []RecordsRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PatchRecord) GetRdataOk() (*[]RecordsRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PatchRecord) SetRdata(v []RecordsRdataInner)`

SetRdata sets Rdata field to given value.

### HasRdata

`func (o *PatchRecord) HasRdata() bool`

HasRdata returns a boolean if a field has been set.

### GetLabels

`func (o *PatchRecord) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PatchRecord) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PatchRecord) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PatchRecord) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *PatchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


