# GetEndpointsResultsInner

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
**Monitorings** | [**[]EndpointMonitoringsMonitoringsInner**](EndpointMonitoringsMonitoringsInner.md) |  | [default to []]

## Methods

### NewGetEndpointsResultsInner

`func NewGetEndpointsResultsInner(resourceName string, name string, monitoringTarget string, description string, weight int32, manualFailback bool, manualFailover bool, enabled bool, liveStatus EndpointLiveStatus, readyStatus EndpointReadyStatus, rdata []EndpointRdataInner, monitorings []EndpointMonitoringsMonitoringsInner, ) *GetEndpointsResultsInner`

NewGetEndpointsResultsInner instantiates a new GetEndpointsResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsResultsInnerWithDefaults

`func NewGetEndpointsResultsInnerWithDefaults() *GetEndpointsResultsInner`

NewGetEndpointsResultsInnerWithDefaults instantiates a new GetEndpointsResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *GetEndpointsResultsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *GetEndpointsResultsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *GetEndpointsResultsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *GetEndpointsResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEndpointsResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEndpointsResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetMonitoringTarget

`func (o *GetEndpointsResultsInner) GetMonitoringTarget() string`

GetMonitoringTarget returns the MonitoringTarget field if non-nil, zero value otherwise.

### GetMonitoringTargetOk

`func (o *GetEndpointsResultsInner) GetMonitoringTargetOk() (*string, bool)`

GetMonitoringTargetOk returns a tuple with the MonitoringTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTarget

`func (o *GetEndpointsResultsInner) SetMonitoringTarget(v string)`

SetMonitoringTarget sets MonitoringTarget field to given value.


### GetDescription

`func (o *GetEndpointsResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEndpointsResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEndpointsResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWeight

`func (o *GetEndpointsResultsInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetEndpointsResultsInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetEndpointsResultsInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetManualFailback

`func (o *GetEndpointsResultsInner) GetManualFailback() bool`

GetManualFailback returns the ManualFailback field if non-nil, zero value otherwise.

### GetManualFailbackOk

`func (o *GetEndpointsResultsInner) GetManualFailbackOk() (*bool, bool)`

GetManualFailbackOk returns a tuple with the ManualFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailback

`func (o *GetEndpointsResultsInner) SetManualFailback(v bool)`

SetManualFailback sets ManualFailback field to given value.


### GetManualFailover

`func (o *GetEndpointsResultsInner) GetManualFailover() bool`

GetManualFailover returns the ManualFailover field if non-nil, zero value otherwise.

### GetManualFailoverOk

`func (o *GetEndpointsResultsInner) GetManualFailoverOk() (*bool, bool)`

GetManualFailoverOk returns a tuple with the ManualFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailover

`func (o *GetEndpointsResultsInner) SetManualFailover(v bool)`

SetManualFailover sets ManualFailover field to given value.


### GetEnabled

`func (o *GetEndpointsResultsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetEndpointsResultsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetEndpointsResultsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLiveStatus

`func (o *GetEndpointsResultsInner) GetLiveStatus() EndpointLiveStatus`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *GetEndpointsResultsInner) GetLiveStatusOk() (*EndpointLiveStatus, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStatus

`func (o *GetEndpointsResultsInner) SetLiveStatus(v EndpointLiveStatus)`

SetLiveStatus sets LiveStatus field to given value.


### GetReadyStatus

`func (o *GetEndpointsResultsInner) GetReadyStatus() EndpointReadyStatus`

GetReadyStatus returns the ReadyStatus field if non-nil, zero value otherwise.

### GetReadyStatusOk

`func (o *GetEndpointsResultsInner) GetReadyStatusOk() (*EndpointReadyStatus, bool)`

GetReadyStatusOk returns a tuple with the ReadyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyStatus

`func (o *GetEndpointsResultsInner) SetReadyStatus(v EndpointReadyStatus)`

SetReadyStatus sets ReadyStatus field to given value.


### GetRdata

`func (o *GetEndpointsResultsInner) GetRdata() []EndpointRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *GetEndpointsResultsInner) GetRdataOk() (*[]EndpointRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *GetEndpointsResultsInner) SetRdata(v []EndpointRdataInner)`

SetRdata sets Rdata field to given value.


### GetMonitorings

`func (o *GetEndpointsResultsInner) GetMonitorings() []EndpointMonitoringsMonitoringsInner`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *GetEndpointsResultsInner) GetMonitoringsOk() (*[]EndpointMonitoringsMonitoringsInner, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *GetEndpointsResultsInner) SetMonitorings(v []EndpointMonitoringsMonitoringsInner)`

SetMonitorings sets Monitorings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


