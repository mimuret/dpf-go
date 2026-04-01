# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | エンドポイント名 | 
**MonitoringTarget** | **string** | 監視ターゲット（IPアドレス形式かホスト名形式） | [default to ""]
**Description** | **string** | コメント | [default to ""]
**Weight** | **int32** | weight | 
**ManualFailback** | **bool** | 手動切り戻し設定値 | [default to false]
**ManualFailover** | **bool** | 手動切り離し設定値 | [default to false]
**Enabled** | **bool** | 状態 | [default to false]
**LiveStatus** | [**EndpointLiveStatus**](EndpointLiveStatus.md) |  | 
**ReadyStatus** | [**EndpointReadyStatus**](EndpointReadyStatus.md) |  | 
**Rdata** | [**[]EndpointRdataInner**](EndpointRdataInner.md) | RDATA | 

## Methods

### NewEndpoint

`func NewEndpoint(resourceName string, name string, monitoringTarget string, description string, weight int32, manualFailback bool, manualFailover bool, enabled bool, liveStatus EndpointLiveStatus, readyStatus EndpointReadyStatus, rdata []EndpointRdataInner, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *Endpoint) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Endpoint) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Endpoint) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *Endpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Endpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Endpoint) SetName(v string)`

SetName sets Name field to given value.


### GetMonitoringTarget

`func (o *Endpoint) GetMonitoringTarget() string`

GetMonitoringTarget returns the MonitoringTarget field if non-nil, zero value otherwise.

### GetMonitoringTargetOk

`func (o *Endpoint) GetMonitoringTargetOk() (*string, bool)`

GetMonitoringTargetOk returns a tuple with the MonitoringTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTarget

`func (o *Endpoint) SetMonitoringTarget(v string)`

SetMonitoringTarget sets MonitoringTarget field to given value.


### GetDescription

`func (o *Endpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Endpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Endpoint) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWeight

`func (o *Endpoint) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Endpoint) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Endpoint) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetManualFailback

`func (o *Endpoint) GetManualFailback() bool`

GetManualFailback returns the ManualFailback field if non-nil, zero value otherwise.

### GetManualFailbackOk

`func (o *Endpoint) GetManualFailbackOk() (*bool, bool)`

GetManualFailbackOk returns a tuple with the ManualFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailback

`func (o *Endpoint) SetManualFailback(v bool)`

SetManualFailback sets ManualFailback field to given value.


### GetManualFailover

`func (o *Endpoint) GetManualFailover() bool`

GetManualFailover returns the ManualFailover field if non-nil, zero value otherwise.

### GetManualFailoverOk

`func (o *Endpoint) GetManualFailoverOk() (*bool, bool)`

GetManualFailoverOk returns a tuple with the ManualFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailover

`func (o *Endpoint) SetManualFailover(v bool)`

SetManualFailover sets ManualFailover field to given value.


### GetEnabled

`func (o *Endpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Endpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Endpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLiveStatus

`func (o *Endpoint) GetLiveStatus() EndpointLiveStatus`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *Endpoint) GetLiveStatusOk() (*EndpointLiveStatus, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStatus

`func (o *Endpoint) SetLiveStatus(v EndpointLiveStatus)`

SetLiveStatus sets LiveStatus field to given value.


### GetReadyStatus

`func (o *Endpoint) GetReadyStatus() EndpointReadyStatus`

GetReadyStatus returns the ReadyStatus field if non-nil, zero value otherwise.

### GetReadyStatusOk

`func (o *Endpoint) GetReadyStatusOk() (*EndpointReadyStatus, bool)`

GetReadyStatusOk returns a tuple with the ReadyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyStatus

`func (o *Endpoint) SetReadyStatus(v EndpointReadyStatus)`

SetReadyStatus sets ReadyStatus field to given value.


### GetRdata

`func (o *Endpoint) GetRdata() []EndpointRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *Endpoint) GetRdataOk() (*[]EndpointRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *Endpoint) SetRdata(v []EndpointRdataInner)`

SetRdata sets Rdata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


