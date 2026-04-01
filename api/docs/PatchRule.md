# PatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ルール名 | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]

## Methods

### NewPatchRule

`func NewPatchRule() *PatchRule`

NewPatchRule instantiates a new PatchRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRuleWithDefaults

`func NewPatchRuleWithDefaults() *PatchRule`

NewPatchRuleWithDefaults instantiates a new PatchRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


