# CcNoticeAccountProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mail** | **string** | メールアドレス | 
**Phone** | [**CcNoticeAccountPropsPhone**](CcNoticeAccountPropsPhone.md) |  | 

## Methods

### NewCcNoticeAccountProps

`func NewCcNoticeAccountProps(mail string, phone CcNoticeAccountPropsPhone, ) *CcNoticeAccountProps`

NewCcNoticeAccountProps instantiates a new CcNoticeAccountProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcNoticeAccountPropsWithDefaults

`func NewCcNoticeAccountPropsWithDefaults() *CcNoticeAccountProps`

NewCcNoticeAccountPropsWithDefaults instantiates a new CcNoticeAccountProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMail

`func (o *CcNoticeAccountProps) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *CcNoticeAccountProps) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *CcNoticeAccountProps) SetMail(v string)`

SetMail sets Mail field to given value.


### GetPhone

`func (o *CcNoticeAccountProps) GetPhone() CcNoticeAccountPropsPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CcNoticeAccountProps) GetPhoneOk() (*CcNoticeAccountPropsPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CcNoticeAccountProps) SetPhone(v CcNoticeAccountPropsPhone)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


