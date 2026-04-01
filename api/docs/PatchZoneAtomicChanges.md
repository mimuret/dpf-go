# PatchZoneAtomicChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]OverwriteRecordsInner**](OverwriteRecordsInner.md) |  | 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**OverwriteSoa** | Pointer to **bool** | 指定されたSOAレコードを取り込むためのフラグ | [optional] [default to true]
**OverwriteZoneApexNs** | Pointer to **bool** | 指定されたZone Apex NSレコードを取り込むためのフラグ | [optional] [default to true]

## Methods

### NewPatchZoneAtomicChanges

`func NewPatchZoneAtomicChanges(records []OverwriteRecordsInner, ) *PatchZoneAtomicChanges`

NewPatchZoneAtomicChanges instantiates a new PatchZoneAtomicChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchZoneAtomicChangesWithDefaults

`func NewPatchZoneAtomicChangesWithDefaults() *PatchZoneAtomicChanges`

NewPatchZoneAtomicChangesWithDefaults instantiates a new PatchZoneAtomicChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *PatchZoneAtomicChanges) GetRecords() []OverwriteRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PatchZoneAtomicChanges) GetRecordsOk() (*[]OverwriteRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PatchZoneAtomicChanges) SetRecords(v []OverwriteRecordsInner)`

SetRecords sets Records field to given value.


### GetDescription

`func (o *PatchZoneAtomicChanges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchZoneAtomicChanges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchZoneAtomicChanges) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchZoneAtomicChanges) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverwriteSoa

`func (o *PatchZoneAtomicChanges) GetOverwriteSoa() bool`

GetOverwriteSoa returns the OverwriteSoa field if non-nil, zero value otherwise.

### GetOverwriteSoaOk

`func (o *PatchZoneAtomicChanges) GetOverwriteSoaOk() (*bool, bool)`

GetOverwriteSoaOk returns a tuple with the OverwriteSoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteSoa

`func (o *PatchZoneAtomicChanges) SetOverwriteSoa(v bool)`

SetOverwriteSoa sets OverwriteSoa field to given value.

### HasOverwriteSoa

`func (o *PatchZoneAtomicChanges) HasOverwriteSoa() bool`

HasOverwriteSoa returns a boolean if a field has been set.

### GetOverwriteZoneApexNs

`func (o *PatchZoneAtomicChanges) GetOverwriteZoneApexNs() bool`

GetOverwriteZoneApexNs returns the OverwriteZoneApexNs field if non-nil, zero value otherwise.

### GetOverwriteZoneApexNsOk

`func (o *PatchZoneAtomicChanges) GetOverwriteZoneApexNsOk() (*bool, bool)`

GetOverwriteZoneApexNsOk returns a tuple with the OverwriteZoneApexNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteZoneApexNs

`func (o *PatchZoneAtomicChanges) SetOverwriteZoneApexNs(v bool)`

SetOverwriteZoneApexNs sets OverwriteZoneApexNs field to given value.

### HasOverwriteZoneApexNs

`func (o *PatchZoneAtomicChanges) HasOverwriteZoneApexNs() bool`

HasOverwriteZoneApexNs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


