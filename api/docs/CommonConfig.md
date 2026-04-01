# CommonConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** | name | 
**ManagedDnsEnabled** | [**ManagedDnsEnabled**](ManagedDnsEnabled.md) |  | 
**Default** | [**CommonConfigDefault**](CommonConfigDefault.md) |  | 
**Description** | **string** | コメント | [default to ""]

## Methods

### NewCommonConfig

`func NewCommonConfig(id int64, name string, managedDnsEnabled ManagedDnsEnabled, default_ CommonConfigDefault, description string, ) *CommonConfig`

NewCommonConfig instantiates a new CommonConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonConfigWithDefaults

`func NewCommonConfigWithDefaults() *CommonConfig`

NewCommonConfigWithDefaults instantiates a new CommonConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonConfig) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CommonConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonConfig) SetName(v string)`

SetName sets Name field to given value.


### GetManagedDnsEnabled

`func (o *CommonConfig) GetManagedDnsEnabled() ManagedDnsEnabled`

GetManagedDnsEnabled returns the ManagedDnsEnabled field if non-nil, zero value otherwise.

### GetManagedDnsEnabledOk

`func (o *CommonConfig) GetManagedDnsEnabledOk() (*ManagedDnsEnabled, bool)`

GetManagedDnsEnabledOk returns a tuple with the ManagedDnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDnsEnabled

`func (o *CommonConfig) SetManagedDnsEnabled(v ManagedDnsEnabled)`

SetManagedDnsEnabled sets ManagedDnsEnabled field to given value.


### GetDefault

`func (o *CommonConfig) GetDefault() CommonConfigDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CommonConfig) GetDefaultOk() (*CommonConfigDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CommonConfig) SetDefault(v CommonConfigDefault)`

SetDefault sets Default field to given value.


### GetDescription

`func (o *CommonConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonConfig) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


