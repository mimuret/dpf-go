# PatchCcNoticeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | アカウント名 | [optional] 
**Lang** | Pointer to [**CcNoticeAccountLang**](CcNoticeAccountLang.md) |  | [optional] 
**Props** | Pointer to [**CcNoticeAccountProps**](CcNoticeAccountProps.md) |  | [optional] 

## Methods

### NewPatchCcNoticeAccount

`func NewPatchCcNoticeAccount() *PatchCcNoticeAccount`

NewPatchCcNoticeAccount instantiates a new PatchCcNoticeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCcNoticeAccountWithDefaults

`func NewPatchCcNoticeAccountWithDefaults() *PatchCcNoticeAccount`

NewPatchCcNoticeAccountWithDefaults instantiates a new PatchCcNoticeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchCcNoticeAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchCcNoticeAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchCcNoticeAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchCcNoticeAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLang

`func (o *PatchCcNoticeAccount) GetLang() CcNoticeAccountLang`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *PatchCcNoticeAccount) GetLangOk() (*CcNoticeAccountLang, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *PatchCcNoticeAccount) SetLang(v CcNoticeAccountLang)`

SetLang sets Lang field to given value.

### HasLang

`func (o *PatchCcNoticeAccount) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetProps

`func (o *PatchCcNoticeAccount) GetProps() CcNoticeAccountProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *PatchCcNoticeAccount) GetPropsOk() (*CcNoticeAccountProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *PatchCcNoticeAccount) SetProps(v CcNoticeAccountProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *PatchCcNoticeAccount) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


