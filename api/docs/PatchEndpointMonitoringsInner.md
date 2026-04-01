# PatchEndpointMonitoringsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 紐付ける監視のリソース名(GET monitorings Schemaにおける resource_name) | 
**Enabled** | Pointer to **bool** | エンドポイントと監視の紐付け状態 | [optional] 

## Methods

### NewPatchEndpointMonitoringsInner

`func NewPatchEndpointMonitoringsInner(resourceName string, ) *PatchEndpointMonitoringsInner`

NewPatchEndpointMonitoringsInner instantiates a new PatchEndpointMonitoringsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchEndpointMonitoringsInnerWithDefaults

`func NewPatchEndpointMonitoringsInnerWithDefaults() *PatchEndpointMonitoringsInner`

NewPatchEndpointMonitoringsInnerWithDefaults instantiates a new PatchEndpointMonitoringsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PatchEndpointMonitoringsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PatchEndpointMonitoringsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PatchEndpointMonitoringsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetEnabled

`func (o *PatchEndpointMonitoringsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchEndpointMonitoringsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchEndpointMonitoringsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchEndpointMonitoringsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


