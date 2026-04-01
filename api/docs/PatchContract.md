# PatchContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favorite** | Pointer to [**ContractFavorite**](ContractFavorite.md) |  | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]

## Methods

### NewPatchContract

`func NewPatchContract() *PatchContract`

NewPatchContract instantiates a new PatchContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchContractWithDefaults

`func NewPatchContractWithDefaults() *PatchContract`

NewPatchContractWithDefaults instantiates a new PatchContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavorite

`func (o *PatchContract) GetFavorite() ContractFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *PatchContract) GetFavoriteOk() (*ContractFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *PatchContract) SetFavorite(v ContractFavorite)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *PatchContract) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetDescription

`func (o *PatchContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


