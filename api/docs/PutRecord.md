# PutRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]OverwriteRecordsInner**](OverwriteRecordsInner.md) |  | 
**OverwriteSoa** | Pointer to **bool** | 指定されたSOAレコードを取り込むためのフラグ | [optional] [default to true]
**OverwriteZoneApexNs** | Pointer to **bool** | 指定されたZone Apex NSレコードを取り込むためのフラグ | [optional] [default to true]

## Methods

### NewPutRecord

`func NewPutRecord(records []OverwriteRecordsInner, ) *PutRecord`

NewPutRecord instantiates a new PutRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRecordWithDefaults

`func NewPutRecordWithDefaults() *PutRecord`

NewPutRecordWithDefaults instantiates a new PutRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *PutRecord) GetRecords() []OverwriteRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PutRecord) GetRecordsOk() (*[]OverwriteRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PutRecord) SetRecords(v []OverwriteRecordsInner)`

SetRecords sets Records field to given value.


### GetOverwriteSoa

`func (o *PutRecord) GetOverwriteSoa() bool`

GetOverwriteSoa returns the OverwriteSoa field if non-nil, zero value otherwise.

### GetOverwriteSoaOk

`func (o *PutRecord) GetOverwriteSoaOk() (*bool, bool)`

GetOverwriteSoaOk returns a tuple with the OverwriteSoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteSoa

`func (o *PutRecord) SetOverwriteSoa(v bool)`

SetOverwriteSoa sets OverwriteSoa field to given value.

### HasOverwriteSoa

`func (o *PutRecord) HasOverwriteSoa() bool`

HasOverwriteSoa returns a boolean if a field has been set.

### GetOverwriteZoneApexNs

`func (o *PutRecord) GetOverwriteZoneApexNs() bool`

GetOverwriteZoneApexNs returns the OverwriteZoneApexNs field if non-nil, zero value otherwise.

### GetOverwriteZoneApexNsOk

`func (o *PutRecord) GetOverwriteZoneApexNsOk() (*bool, bool)`

GetOverwriteZoneApexNsOk returns a tuple with the OverwriteZoneApexNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteZoneApexNs

`func (o *PutRecord) SetOverwriteZoneApexNs(v bool)`

SetOverwriteZoneApexNs sets OverwriteZoneApexNs field to given value.

### HasOverwriteZoneApexNs

`func (o *PutRecord) HasOverwriteZoneApexNs() bool`

HasOverwriteZoneApexNs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


