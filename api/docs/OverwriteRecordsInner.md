# OverwriteRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name | 
**Ttl** | **NullableInt32** | TTL | 
**Rrtype** | [**RecordsRrtype**](RecordsRrtype.md) |  | 
**Rdata** | [**[]RecordsRdataInner**](RecordsRdataInner.md) | レコードの値 | 
**Description** | **string** | コメント | [default to ""]
**Labels** | **map[string]string** | ラベル | [default to {}]

## Methods

### NewOverwriteRecordsInner

`func NewOverwriteRecordsInner(name string, ttl NullableInt32, rrtype RecordsRrtype, rdata []RecordsRdataInner, description string, labels map[string]string, ) *OverwriteRecordsInner`

NewOverwriteRecordsInner instantiates a new OverwriteRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverwriteRecordsInnerWithDefaults

`func NewOverwriteRecordsInnerWithDefaults() *OverwriteRecordsInner`

NewOverwriteRecordsInnerWithDefaults instantiates a new OverwriteRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OverwriteRecordsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverwriteRecordsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverwriteRecordsInner) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *OverwriteRecordsInner) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *OverwriteRecordsInner) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *OverwriteRecordsInner) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *OverwriteRecordsInner) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *OverwriteRecordsInner) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetRrtype

`func (o *OverwriteRecordsInner) GetRrtype() RecordsRrtype`

GetRrtype returns the Rrtype field if non-nil, zero value otherwise.

### GetRrtypeOk

`func (o *OverwriteRecordsInner) GetRrtypeOk() (*RecordsRrtype, bool)`

GetRrtypeOk returns a tuple with the Rrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrtype

`func (o *OverwriteRecordsInner) SetRrtype(v RecordsRrtype)`

SetRrtype sets Rrtype field to given value.


### GetRdata

`func (o *OverwriteRecordsInner) GetRdata() []RecordsRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *OverwriteRecordsInner) GetRdataOk() (*[]RecordsRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *OverwriteRecordsInner) SetRdata(v []RecordsRdataInner)`

SetRdata sets Rdata field to given value.


### GetDescription

`func (o *OverwriteRecordsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OverwriteRecordsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OverwriteRecordsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabels

`func (o *OverwriteRecordsInner) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OverwriteRecordsInner) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OverwriteRecordsInner) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


