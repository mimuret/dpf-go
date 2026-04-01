# CcNoticeAccountPropsPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | 国番号（number指定時必須） | 
**Number** | **string** | ハイフンを除いた電話番号（country_code指定時必須） | 

## Methods

### NewCcNoticeAccountPropsPhone

`func NewCcNoticeAccountPropsPhone(countryCode string, number string, ) *CcNoticeAccountPropsPhone`

NewCcNoticeAccountPropsPhone instantiates a new CcNoticeAccountPropsPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcNoticeAccountPropsPhoneWithDefaults

`func NewCcNoticeAccountPropsPhoneWithDefaults() *CcNoticeAccountPropsPhone`

NewCcNoticeAccountPropsPhoneWithDefaults instantiates a new CcNoticeAccountPropsPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CcNoticeAccountPropsPhone) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CcNoticeAccountPropsPhone) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CcNoticeAccountPropsPhone) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetNumber

`func (o *CcNoticeAccountPropsPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CcNoticeAccountPropsPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CcNoticeAccountPropsPhone) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


