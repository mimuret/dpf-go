# PostTsig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 先頭末尾がハイフン以外の[a-z0-9]とハイフンで構成された文字列が使用できます | 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]

## Methods

### NewPostTsig

`func NewPostTsig(name string, ) *PostTsig`

NewPostTsig instantiates a new PostTsig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTsigWithDefaults

`func NewPostTsigWithDefaults() *PostTsig`

NewPostTsigWithDefaults instantiates a new PostTsig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostTsig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostTsig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostTsig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostTsig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostTsig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostTsig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostTsig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


