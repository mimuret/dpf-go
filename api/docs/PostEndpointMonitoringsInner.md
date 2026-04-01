# PostEndpointMonitoringsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 紐付ける監視のリソース名(GET monitorings Schemaにおける resource_name) | 
**Enabled** | Pointer to **bool** | エンドポイントと監視の紐付け状態 | [optional] [default to false]

## Methods

### NewPostEndpointMonitoringsInner

`func NewPostEndpointMonitoringsInner(resourceName string, ) *PostEndpointMonitoringsInner`

NewPostEndpointMonitoringsInner instantiates a new PostEndpointMonitoringsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEndpointMonitoringsInnerWithDefaults

`func NewPostEndpointMonitoringsInnerWithDefaults() *PostEndpointMonitoringsInner`

NewPostEndpointMonitoringsInnerWithDefaults instantiates a new PostEndpointMonitoringsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PostEndpointMonitoringsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PostEndpointMonitoringsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PostEndpointMonitoringsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetEnabled

`func (o *PostEndpointMonitoringsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostEndpointMonitoringsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostEndpointMonitoringsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostEndpointMonitoringsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


