# ConfigRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | ルール名 | 
**Description** | **string** | コメント | [default to ""]
**Methods** | [**[]ConfigRuleMethodsInner**](ConfigRuleMethodsInner.md) |  | 

## Methods

### NewConfigRulesInner

`func NewConfigRulesInner(resourceName string, name string, description string, methods []ConfigRuleMethodsInner, ) *ConfigRulesInner`

NewConfigRulesInner instantiates a new ConfigRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRulesInnerWithDefaults

`func NewConfigRulesInnerWithDefaults() *ConfigRulesInner`

NewConfigRulesInnerWithDefaults instantiates a new ConfigRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *ConfigRulesInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ConfigRulesInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ConfigRulesInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *ConfigRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigRulesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigRulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigRulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigRulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethods

`func (o *ConfigRulesInner) GetMethods() []ConfigRuleMethodsInner`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ConfigRulesInner) GetMethodsOk() (*[]ConfigRuleMethodsInner, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ConfigRulesInner) SetMethods(v []ConfigRuleMethodsInner)`

SetMethods sets Methods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


