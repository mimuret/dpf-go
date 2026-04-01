# PostRuleMethodMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | [optional] 
**Mtype** | **string** | メソッド種別 | 
**ParentResourceName** | Pointer to **string** | 親メソッドのリソース名（mtype が failover/exit_site/exit_sorry の場合は必須） | [optional] 
**Enabled** | Pointer to **bool** | 状態 | [optional] [default to false]
**SiteResourceName** | Pointer to **string** | 紐付けされているサイトのリソース名（mtype が exit_site の場合のみ指定） | [optional] 

## Methods

### NewPostRuleMethodMethod

`func NewPostRuleMethodMethod(mtype string, ) *PostRuleMethodMethod`

NewPostRuleMethodMethod instantiates a new PostRuleMethodMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRuleMethodMethodWithDefaults

`func NewPostRuleMethodMethodWithDefaults() *PostRuleMethodMethod`

NewPostRuleMethodMethodWithDefaults instantiates a new PostRuleMethodMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PostRuleMethodMethod) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PostRuleMethodMethod) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PostRuleMethodMethod) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PostRuleMethodMethod) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetMtype

`func (o *PostRuleMethodMethod) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *PostRuleMethodMethod) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *PostRuleMethodMethod) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetParentResourceName

`func (o *PostRuleMethodMethod) GetParentResourceName() string`

GetParentResourceName returns the ParentResourceName field if non-nil, zero value otherwise.

### GetParentResourceNameOk

`func (o *PostRuleMethodMethod) GetParentResourceNameOk() (*string, bool)`

GetParentResourceNameOk returns a tuple with the ParentResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceName

`func (o *PostRuleMethodMethod) SetParentResourceName(v string)`

SetParentResourceName sets ParentResourceName field to given value.

### HasParentResourceName

`func (o *PostRuleMethodMethod) HasParentResourceName() bool`

HasParentResourceName returns a boolean if a field has been set.

### GetEnabled

`func (o *PostRuleMethodMethod) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostRuleMethodMethod) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostRuleMethodMethod) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostRuleMethodMethod) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSiteResourceName

`func (o *PostRuleMethodMethod) GetSiteResourceName() string`

GetSiteResourceName returns the SiteResourceName field if non-nil, zero value otherwise.

### GetSiteResourceNameOk

`func (o *PostRuleMethodMethod) GetSiteResourceNameOk() (*string, bool)`

GetSiteResourceNameOk returns a tuple with the SiteResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteResourceName

`func (o *PostRuleMethodMethod) SetSiteResourceName(v string)`

SetSiteResourceName sets SiteResourceName field to given value.

### HasSiteResourceName

`func (o *PostRuleMethodMethod) HasSiteResourceName() bool`

HasSiteResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


