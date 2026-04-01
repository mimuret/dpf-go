# PatchZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favorite** | Pointer to [**ZonesFavorite**](ZonesFavorite.md) |  | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]

## Methods

### NewPatchZone

`func NewPatchZone() *PatchZone`

NewPatchZone instantiates a new PatchZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchZoneWithDefaults

`func NewPatchZoneWithDefaults() *PatchZone`

NewPatchZoneWithDefaults instantiates a new PatchZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavorite

`func (o *PatchZone) GetFavorite() ZonesFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *PatchZone) GetFavoriteOk() (*ZonesFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *PatchZone) SetFavorite(v ZonesFavorite)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *PatchZone) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetDescription

`func (o *PatchZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


