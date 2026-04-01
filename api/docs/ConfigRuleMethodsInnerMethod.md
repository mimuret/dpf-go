# ConfigRuleMethodsInnerMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Mtype** | **string** | メソッド種別 | 
**ParentResourceName** | Pointer to **string** | 親メソッドのリソース名 | [optional] 
**Enabled** | **bool** | 状態 | [default to false]
**SiteResourceName** | Pointer to **string** | 紐付けされているサイトのリソース名（mtype が exit_site の場合のみ指定） | [optional] 
**LiveStatus** | **string** | Live Status | 
**ReadyStatus** | **string** | Ready Status | 
**Methods** | [**[]ConfigRuleMethodsInner**](ConfigRuleMethodsInner.md) |  | 

## Methods

### NewConfigRuleMethodsInnerMethod

`func NewConfigRuleMethodsInnerMethod(resourceName string, mtype string, enabled bool, liveStatus string, readyStatus string, methods []ConfigRuleMethodsInner, ) *ConfigRuleMethodsInnerMethod`

NewConfigRuleMethodsInnerMethod instantiates a new ConfigRuleMethodsInnerMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRuleMethodsInnerMethodWithDefaults

`func NewConfigRuleMethodsInnerMethodWithDefaults() *ConfigRuleMethodsInnerMethod`

NewConfigRuleMethodsInnerMethodWithDefaults instantiates a new ConfigRuleMethodsInnerMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *ConfigRuleMethodsInnerMethod) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ConfigRuleMethodsInnerMethod) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ConfigRuleMethodsInnerMethod) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetMtype

`func (o *ConfigRuleMethodsInnerMethod) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *ConfigRuleMethodsInnerMethod) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *ConfigRuleMethodsInnerMethod) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetParentResourceName

`func (o *ConfigRuleMethodsInnerMethod) GetParentResourceName() string`

GetParentResourceName returns the ParentResourceName field if non-nil, zero value otherwise.

### GetParentResourceNameOk

`func (o *ConfigRuleMethodsInnerMethod) GetParentResourceNameOk() (*string, bool)`

GetParentResourceNameOk returns a tuple with the ParentResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceName

`func (o *ConfigRuleMethodsInnerMethod) SetParentResourceName(v string)`

SetParentResourceName sets ParentResourceName field to given value.

### HasParentResourceName

`func (o *ConfigRuleMethodsInnerMethod) HasParentResourceName() bool`

HasParentResourceName returns a boolean if a field has been set.

### GetEnabled

`func (o *ConfigRuleMethodsInnerMethod) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigRuleMethodsInnerMethod) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigRuleMethodsInnerMethod) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSiteResourceName

`func (o *ConfigRuleMethodsInnerMethod) GetSiteResourceName() string`

GetSiteResourceName returns the SiteResourceName field if non-nil, zero value otherwise.

### GetSiteResourceNameOk

`func (o *ConfigRuleMethodsInnerMethod) GetSiteResourceNameOk() (*string, bool)`

GetSiteResourceNameOk returns a tuple with the SiteResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteResourceName

`func (o *ConfigRuleMethodsInnerMethod) SetSiteResourceName(v string)`

SetSiteResourceName sets SiteResourceName field to given value.

### HasSiteResourceName

`func (o *ConfigRuleMethodsInnerMethod) HasSiteResourceName() bool`

HasSiteResourceName returns a boolean if a field has been set.

### GetLiveStatus

`func (o *ConfigRuleMethodsInnerMethod) GetLiveStatus() string`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *ConfigRuleMethodsInnerMethod) GetLiveStatusOk() (*string, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStatus

`func (o *ConfigRuleMethodsInnerMethod) SetLiveStatus(v string)`

SetLiveStatus sets LiveStatus field to given value.


### GetReadyStatus

`func (o *ConfigRuleMethodsInnerMethod) GetReadyStatus() string`

GetReadyStatus returns the ReadyStatus field if non-nil, zero value otherwise.

### GetReadyStatusOk

`func (o *ConfigRuleMethodsInnerMethod) GetReadyStatusOk() (*string, bool)`

GetReadyStatusOk returns a tuple with the ReadyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyStatus

`func (o *ConfigRuleMethodsInnerMethod) SetReadyStatus(v string)`

SetReadyStatus sets ReadyStatus field to given value.


### GetMethods

`func (o *ConfigRuleMethodsInnerMethod) GetMethods() []ConfigRuleMethodsInner`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ConfigRuleMethodsInnerMethod) GetMethodsOk() (*[]ConfigRuleMethodsInner, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ConfigRuleMethodsInnerMethod) SetMethods(v []ConfigRuleMethodsInner)`

SetMethods sets Methods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


