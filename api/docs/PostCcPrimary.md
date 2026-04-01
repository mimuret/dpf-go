# PostCcPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsigId** | Pointer to **NullableInt64** |  | [optional] 
**Address** | **string** | IPアドレス | 

## Methods

### NewPostCcPrimary

`func NewPostCcPrimary(address string, ) *PostCcPrimary`

NewPostCcPrimary instantiates a new PostCcPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCcPrimaryWithDefaults

`func NewPostCcPrimaryWithDefaults() *PostCcPrimary`

NewPostCcPrimaryWithDefaults instantiates a new PostCcPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsigId

`func (o *PostCcPrimary) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *PostCcPrimary) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *PostCcPrimary) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.

### HasTsigId

`func (o *PostCcPrimary) HasTsigId() bool`

HasTsigId returns a boolean if a field has been set.

### SetTsigIdNil

`func (o *PostCcPrimary) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *PostCcPrimary) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil
### GetAddress

`func (o *PostCcPrimary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostCcPrimary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostCcPrimary) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


