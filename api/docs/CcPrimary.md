# CcPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Address** | **string** | IPアドレス | 
**TsigId** | **NullableInt64** |  | 
**Enabled** | [**CcPrimaryEnabled**](CcPrimaryEnabled.md) |  | 

## Methods

### NewCcPrimary

`func NewCcPrimary(id int64, address string, tsigId NullableInt64, enabled CcPrimaryEnabled, ) *CcPrimary`

NewCcPrimary instantiates a new CcPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcPrimaryWithDefaults

`func NewCcPrimaryWithDefaults() *CcPrimary`

NewCcPrimaryWithDefaults instantiates a new CcPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CcPrimary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcPrimary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcPrimary) SetId(v int64)`

SetId sets Id field to given value.


### GetAddress

`func (o *CcPrimary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CcPrimary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CcPrimary) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTsigId

`func (o *CcPrimary) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *CcPrimary) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *CcPrimary) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.


### SetTsigIdNil

`func (o *CcPrimary) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *CcPrimary) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil
### GetEnabled

`func (o *CcPrimary) GetEnabled() CcPrimaryEnabled`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CcPrimary) GetEnabledOk() (*CcPrimaryEnabled, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CcPrimary) SetEnabled(v CcPrimaryEnabled)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


