# PatchLbDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favorite** | Pointer to [**LbDomainFavorite**](LbDomainFavorite.md) |  | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**RuleResourceName** | Pointer to **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | [optional] 

## Methods

### NewPatchLbDomain

`func NewPatchLbDomain() *PatchLbDomain`

NewPatchLbDomain instantiates a new PatchLbDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLbDomainWithDefaults

`func NewPatchLbDomainWithDefaults() *PatchLbDomain`

NewPatchLbDomainWithDefaults instantiates a new PatchLbDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavorite

`func (o *PatchLbDomain) GetFavorite() LbDomainFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *PatchLbDomain) GetFavoriteOk() (*LbDomainFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *PatchLbDomain) SetFavorite(v LbDomainFavorite)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *PatchLbDomain) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetDescription

`func (o *PatchLbDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchLbDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchLbDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchLbDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleResourceName

`func (o *PatchLbDomain) GetRuleResourceName() string`

GetRuleResourceName returns the RuleResourceName field if non-nil, zero value otherwise.

### GetRuleResourceNameOk

`func (o *PatchLbDomain) GetRuleResourceNameOk() (*string, bool)`

GetRuleResourceNameOk returns a tuple with the RuleResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleResourceName

`func (o *PatchLbDomain) SetRuleResourceName(v string)`

SetRuleResourceName sets RuleResourceName field to given value.

### HasRuleResourceName

`func (o *PatchLbDomain) HasRuleResourceName() bool`

HasRuleResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


