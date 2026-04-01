# PatchCcPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsigId** | Pointer to **NullableInt64** |  | [optional] 
**Address** | Pointer to **string** | IPアドレス | [optional] 
**Enabled** | Pointer to [**CcPrimaryEnabled**](CcPrimaryEnabled.md) |  | [optional] 

## Methods

### NewPatchCcPrimary

`func NewPatchCcPrimary() *PatchCcPrimary`

NewPatchCcPrimary instantiates a new PatchCcPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCcPrimaryWithDefaults

`func NewPatchCcPrimaryWithDefaults() *PatchCcPrimary`

NewPatchCcPrimaryWithDefaults instantiates a new PatchCcPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsigId

`func (o *PatchCcPrimary) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *PatchCcPrimary) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *PatchCcPrimary) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.

### HasTsigId

`func (o *PatchCcPrimary) HasTsigId() bool`

HasTsigId returns a boolean if a field has been set.

### SetTsigIdNil

`func (o *PatchCcPrimary) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *PatchCcPrimary) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil
### GetAddress

`func (o *PatchCcPrimary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchCcPrimary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchCcPrimary) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchCcPrimary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchCcPrimary) GetEnabled() CcPrimaryEnabled`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchCcPrimary) GetEnabledOk() (*CcPrimaryEnabled, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchCcPrimary) SetEnabled(v CcPrimaryEnabled)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchCcPrimary) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


