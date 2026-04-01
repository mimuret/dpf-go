# PostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name | 
**Ttl** | Pointer to **NullableInt32** | TTL | [optional] 
**Rrtype** | [**RecordsRrtypeWithoutSoa**](RecordsRrtypeWithoutSoa.md) |  | 
**Rdata** | [**[]RecordsRdataInner**](RecordsRdataInner.md) | レコードの値 | 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**Labels** | Pointer to **map[string]string** | ラベル | [optional] [default to {}]

## Methods

### NewPostRecord

`func NewPostRecord(name string, rrtype RecordsRrtypeWithoutSoa, rdata []RecordsRdataInner, ) *PostRecord`

NewPostRecord instantiates a new PostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRecordWithDefaults

`func NewPostRecordWithDefaults() *PostRecord`

NewPostRecordWithDefaults instantiates a new PostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostRecord) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *PostRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *PostRecord) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *PostRecord) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetRrtype

`func (o *PostRecord) GetRrtype() RecordsRrtypeWithoutSoa`

GetRrtype returns the Rrtype field if non-nil, zero value otherwise.

### GetRrtypeOk

`func (o *PostRecord) GetRrtypeOk() (*RecordsRrtypeWithoutSoa, bool)`

GetRrtypeOk returns a tuple with the Rrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrtype

`func (o *PostRecord) SetRrtype(v RecordsRrtypeWithoutSoa)`

SetRrtype sets Rrtype field to given value.


### GetRdata

`func (o *PostRecord) GetRdata() []RecordsRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PostRecord) GetRdataOk() (*[]RecordsRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PostRecord) SetRdata(v []RecordsRdataInner)`

SetRdata sets Rdata field to given value.


### GetDescription

`func (o *PostRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *PostRecord) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PostRecord) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PostRecord) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PostRecord) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


