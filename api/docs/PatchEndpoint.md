# PatchEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | エンドポイント名 | [optional] 
**MonitoringTarget** | Pointer to **string** | 監視ターゲット（IPアドレス形式かホスト名形式） | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**Weight** | Pointer to **int32** | weight | [optional] 
**ManualFailback** | Pointer to **bool** | 手動切り戻し設定値 | [optional] 
**ManualFailover** | Pointer to **bool** | 手動切り離し設定値 | [optional] 
**Enabled** | Pointer to **bool** | 状態 | [optional] 
**Rdata** | Pointer to [**[]EndpointRdataInner**](EndpointRdataInner.md) | RDATA | [optional] 
**Monitorings** | Pointer to [**[]PatchEndpointMonitoringsInner**](PatchEndpointMonitoringsInner.md) |  | [optional] 

## Methods

### NewPatchEndpoint

`func NewPatchEndpoint() *PatchEndpoint`

NewPatchEndpoint instantiates a new PatchEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchEndpointWithDefaults

`func NewPatchEndpointWithDefaults() *PatchEndpoint`

NewPatchEndpointWithDefaults instantiates a new PatchEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMonitoringTarget

`func (o *PatchEndpoint) GetMonitoringTarget() string`

GetMonitoringTarget returns the MonitoringTarget field if non-nil, zero value otherwise.

### GetMonitoringTargetOk

`func (o *PatchEndpoint) GetMonitoringTargetOk() (*string, bool)`

GetMonitoringTargetOk returns a tuple with the MonitoringTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTarget

`func (o *PatchEndpoint) SetMonitoringTarget(v string)`

SetMonitoringTarget sets MonitoringTarget field to given value.

### HasMonitoringTarget

`func (o *PatchEndpoint) HasMonitoringTarget() bool`

HasMonitoringTarget returns a boolean if a field has been set.

### GetDescription

`func (o *PatchEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *PatchEndpoint) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchEndpoint) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchEndpoint) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchEndpoint) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetManualFailback

`func (o *PatchEndpoint) GetManualFailback() bool`

GetManualFailback returns the ManualFailback field if non-nil, zero value otherwise.

### GetManualFailbackOk

`func (o *PatchEndpoint) GetManualFailbackOk() (*bool, bool)`

GetManualFailbackOk returns a tuple with the ManualFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailback

`func (o *PatchEndpoint) SetManualFailback(v bool)`

SetManualFailback sets ManualFailback field to given value.

### HasManualFailback

`func (o *PatchEndpoint) HasManualFailback() bool`

HasManualFailback returns a boolean if a field has been set.

### GetManualFailover

`func (o *PatchEndpoint) GetManualFailover() bool`

GetManualFailover returns the ManualFailover field if non-nil, zero value otherwise.

### GetManualFailoverOk

`func (o *PatchEndpoint) GetManualFailoverOk() (*bool, bool)`

GetManualFailoverOk returns a tuple with the ManualFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailover

`func (o *PatchEndpoint) SetManualFailover(v bool)`

SetManualFailover sets ManualFailover field to given value.

### HasManualFailover

`func (o *PatchEndpoint) HasManualFailover() bool`

HasManualFailover returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchEndpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRdata

`func (o *PatchEndpoint) GetRdata() []EndpointRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PatchEndpoint) GetRdataOk() (*[]EndpointRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PatchEndpoint) SetRdata(v []EndpointRdataInner)`

SetRdata sets Rdata field to given value.

### HasRdata

`func (o *PatchEndpoint) HasRdata() bool`

HasRdata returns a boolean if a field has been set.

### GetMonitorings

`func (o *PatchEndpoint) GetMonitorings() []PatchEndpointMonitoringsInner`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *PatchEndpoint) GetMonitoringsOk() (*[]PatchEndpointMonitoringsInner, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *PatchEndpoint) SetMonitorings(v []PatchEndpointMonitoringsInner)`

SetMonitorings sets Monitorings field to given value.

### HasMonitorings

`func (o *PatchEndpoint) HasMonitorings() bool`

HasMonitorings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


