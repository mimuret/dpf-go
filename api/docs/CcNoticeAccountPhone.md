# CcNoticeAccountPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | 国番号（number指定時必須） | [optional] 
**Number** | Pointer to **string** | ハイフンを除いた電話番号（country_code指定時必須） | [optional] 

## Methods

### NewCcNoticeAccountPhone

`func NewCcNoticeAccountPhone() *CcNoticeAccountPhone`

NewCcNoticeAccountPhone instantiates a new CcNoticeAccountPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcNoticeAccountPhoneWithDefaults

`func NewCcNoticeAccountPhoneWithDefaults() *CcNoticeAccountPhone`

NewCcNoticeAccountPhoneWithDefaults instantiates a new CcNoticeAccountPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CcNoticeAccountPhone) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CcNoticeAccountPhone) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CcNoticeAccountPhone) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CcNoticeAccountPhone) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetNumber

`func (o *CcNoticeAccountPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CcNoticeAccountPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CcNoticeAccountPhone) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CcNoticeAccountPhone) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


