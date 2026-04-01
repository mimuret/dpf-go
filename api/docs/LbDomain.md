# LbDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CommonConfigId** | **int64** |  | 
**ServiceCode** | **string** |  | 
**State** | [**LbDomainState**](LbDomainState.md) |  | 
**Favorite** | [**LbDomainFavorite**](LbDomainFavorite.md) |  | 
**Name** | **string** | LBドメイン名 | 
**Description** | **string** | コメント | [default to ""]
**RuleResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 

## Methods

### NewLbDomain

`func NewLbDomain(id string, commonConfigId int64, serviceCode string, state LbDomainState, favorite LbDomainFavorite, name string, description string, ruleResourceName string, ) *LbDomain`

NewLbDomain instantiates a new LbDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbDomainWithDefaults

`func NewLbDomainWithDefaults() *LbDomain`

NewLbDomainWithDefaults instantiates a new LbDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LbDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbDomain) SetId(v string)`

SetId sets Id field to given value.


### GetCommonConfigId

`func (o *LbDomain) GetCommonConfigId() int64`

GetCommonConfigId returns the CommonConfigId field if non-nil, zero value otherwise.

### GetCommonConfigIdOk

`func (o *LbDomain) GetCommonConfigIdOk() (*int64, bool)`

GetCommonConfigIdOk returns a tuple with the CommonConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonConfigId

`func (o *LbDomain) SetCommonConfigId(v int64)`

SetCommonConfigId sets CommonConfigId field to given value.


### GetServiceCode

`func (o *LbDomain) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *LbDomain) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *LbDomain) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetState

`func (o *LbDomain) GetState() LbDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbDomain) GetStateOk() (*LbDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbDomain) SetState(v LbDomainState)`

SetState sets State field to given value.


### GetFavorite

`func (o *LbDomain) GetFavorite() LbDomainFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *LbDomain) GetFavoriteOk() (*LbDomainFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *LbDomain) SetFavorite(v LbDomainFavorite)`

SetFavorite sets Favorite field to given value.


### GetName

`func (o *LbDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbDomain) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LbDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbDomain) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRuleResourceName

`func (o *LbDomain) GetRuleResourceName() string`

GetRuleResourceName returns the RuleResourceName field if non-nil, zero value otherwise.

### GetRuleResourceNameOk

`func (o *LbDomain) GetRuleResourceNameOk() (*string, bool)`

GetRuleResourceNameOk returns a tuple with the RuleResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleResourceName

`func (o *LbDomain) SetRuleResourceName(v string)`

SetRuleResourceName sets RuleResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


