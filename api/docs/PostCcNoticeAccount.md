# PostCcNoticeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | [optional] 
**Name** | **string** | アカウント名 | 
**Lang** | [**CcNoticeAccountLang**](CcNoticeAccountLang.md) |  | 
**Props** | Pointer to [**CcNoticeAccountProps**](CcNoticeAccountProps.md) |  | [optional] 

## Methods

### NewPostCcNoticeAccount

`func NewPostCcNoticeAccount(name string, lang CcNoticeAccountLang, ) *PostCcNoticeAccount`

NewPostCcNoticeAccount instantiates a new PostCcNoticeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCcNoticeAccountWithDefaults

`func NewPostCcNoticeAccountWithDefaults() *PostCcNoticeAccount`

NewPostCcNoticeAccountWithDefaults instantiates a new PostCcNoticeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PostCcNoticeAccount) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PostCcNoticeAccount) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PostCcNoticeAccount) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PostCcNoticeAccount) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetName

`func (o *PostCcNoticeAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCcNoticeAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCcNoticeAccount) SetName(v string)`

SetName sets Name field to given value.


### GetLang

`func (o *PostCcNoticeAccount) GetLang() CcNoticeAccountLang`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *PostCcNoticeAccount) GetLangOk() (*CcNoticeAccountLang, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *PostCcNoticeAccount) SetLang(v CcNoticeAccountLang)`

SetLang sets Lang field to given value.


### GetProps

`func (o *PostCcNoticeAccount) GetProps() CcNoticeAccountProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *PostCcNoticeAccount) GetPropsOk() (*CcNoticeAccountProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *PostCcNoticeAccount) SetProps(v CcNoticeAccountProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *PostCcNoticeAccount) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


