# PutConfigRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | ルール名 | 
**Description** | **string** | コメント | 
**Methods** | [**[]ConfigRuleMethodsInner**](ConfigRuleMethodsInner.md) |  | 

## Methods

### NewPutConfigRulesInner

`func NewPutConfigRulesInner(resourceName string, name string, description string, methods []ConfigRuleMethodsInner, ) *PutConfigRulesInner`

NewPutConfigRulesInner instantiates a new PutConfigRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutConfigRulesInnerWithDefaults

`func NewPutConfigRulesInnerWithDefaults() *PutConfigRulesInner`

NewPutConfigRulesInnerWithDefaults instantiates a new PutConfigRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PutConfigRulesInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PutConfigRulesInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PutConfigRulesInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *PutConfigRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutConfigRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutConfigRulesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PutConfigRulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutConfigRulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutConfigRulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethods

`func (o *PutConfigRulesInner) GetMethods() []ConfigRuleMethodsInner`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *PutConfigRulesInner) GetMethodsOk() (*[]ConfigRuleMethodsInner, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *PutConfigRulesInner) SetMethods(v []ConfigRuleMethodsInner)`

SetMethods sets Methods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


