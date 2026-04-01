# CcNoticeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | アカウント名 | 
**Lang** | [**CcNoticeAccountLang**](CcNoticeAccountLang.md) |  | 
**Props** | [**CcNoticeAccountProps**](CcNoticeAccountProps.md) |  | 

## Methods

### NewCcNoticeAccount

`func NewCcNoticeAccount(resourceName string, name string, lang CcNoticeAccountLang, props CcNoticeAccountProps, ) *CcNoticeAccount`

NewCcNoticeAccount instantiates a new CcNoticeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcNoticeAccountWithDefaults

`func NewCcNoticeAccountWithDefaults() *CcNoticeAccount`

NewCcNoticeAccountWithDefaults instantiates a new CcNoticeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *CcNoticeAccount) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CcNoticeAccount) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CcNoticeAccount) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *CcNoticeAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CcNoticeAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CcNoticeAccount) SetName(v string)`

SetName sets Name field to given value.


### GetLang

`func (o *CcNoticeAccount) GetLang() CcNoticeAccountLang`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *CcNoticeAccount) GetLangOk() (*CcNoticeAccountLang, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *CcNoticeAccount) SetLang(v CcNoticeAccountLang)`

SetLang sets Lang field to given value.


### GetProps

`func (o *CcNoticeAccount) GetProps() CcNoticeAccountProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *CcNoticeAccount) GetPropsOk() (*CcNoticeAccountProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *CcNoticeAccount) SetProps(v CcNoticeAccountProps)`

SetProps sets Props field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


