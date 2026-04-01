# PutConfigSitesInnerEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | エンドポイント名 | 
**MonitoringTarget** | **string** | 監視ターゲット（IPアドレス形式かホスト名形式） | 
**Description** | **string** | コメント | 
**Weight** | **int32** | weight | 
**ManualFailback** | **bool** | 手動切り戻し設定値 | 
**ManualFailover** | **bool** | 手動切り離し設定値 | 
**Enabled** | **bool** | 状態 | 
**Rdata** | [**[]EndpointRdataInner**](EndpointRdataInner.md) | RDATA | 
**Monitorings** | [**[]PutConfigSitesInnerEndpointsInnerMonitoringsInner**](PutConfigSitesInnerEndpointsInnerMonitoringsInner.md) |  | 

## Methods

### NewPutConfigSitesInnerEndpointsInner

`func NewPutConfigSitesInnerEndpointsInner(resourceName string, name string, monitoringTarget string, description string, weight int32, manualFailback bool, manualFailover bool, enabled bool, rdata []EndpointRdataInner, monitorings []PutConfigSitesInnerEndpointsInnerMonitoringsInner, ) *PutConfigSitesInnerEndpointsInner`

NewPutConfigSitesInnerEndpointsInner instantiates a new PutConfigSitesInnerEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutConfigSitesInnerEndpointsInnerWithDefaults

`func NewPutConfigSitesInnerEndpointsInnerWithDefaults() *PutConfigSitesInnerEndpointsInner`

NewPutConfigSitesInnerEndpointsInnerWithDefaults instantiates a new PutConfigSitesInnerEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PutConfigSitesInnerEndpointsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PutConfigSitesInnerEndpointsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PutConfigSitesInnerEndpointsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *PutConfigSitesInnerEndpointsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutConfigSitesInnerEndpointsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutConfigSitesInnerEndpointsInner) SetName(v string)`

SetName sets Name field to given value.


### GetMonitoringTarget

`func (o *PutConfigSitesInnerEndpointsInner) GetMonitoringTarget() string`

GetMonitoringTarget returns the MonitoringTarget field if non-nil, zero value otherwise.

### GetMonitoringTargetOk

`func (o *PutConfigSitesInnerEndpointsInner) GetMonitoringTargetOk() (*string, bool)`

GetMonitoringTargetOk returns a tuple with the MonitoringTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTarget

`func (o *PutConfigSitesInnerEndpointsInner) SetMonitoringTarget(v string)`

SetMonitoringTarget sets MonitoringTarget field to given value.


### GetDescription

`func (o *PutConfigSitesInnerEndpointsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutConfigSitesInnerEndpointsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutConfigSitesInnerEndpointsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWeight

`func (o *PutConfigSitesInnerEndpointsInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PutConfigSitesInnerEndpointsInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PutConfigSitesInnerEndpointsInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetManualFailback

`func (o *PutConfigSitesInnerEndpointsInner) GetManualFailback() bool`

GetManualFailback returns the ManualFailback field if non-nil, zero value otherwise.

### GetManualFailbackOk

`func (o *PutConfigSitesInnerEndpointsInner) GetManualFailbackOk() (*bool, bool)`

GetManualFailbackOk returns a tuple with the ManualFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailback

`func (o *PutConfigSitesInnerEndpointsInner) SetManualFailback(v bool)`

SetManualFailback sets ManualFailback field to given value.


### GetManualFailover

`func (o *PutConfigSitesInnerEndpointsInner) GetManualFailover() bool`

GetManualFailover returns the ManualFailover field if non-nil, zero value otherwise.

### GetManualFailoverOk

`func (o *PutConfigSitesInnerEndpointsInner) GetManualFailoverOk() (*bool, bool)`

GetManualFailoverOk returns a tuple with the ManualFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailover

`func (o *PutConfigSitesInnerEndpointsInner) SetManualFailover(v bool)`

SetManualFailover sets ManualFailover field to given value.


### GetEnabled

`func (o *PutConfigSitesInnerEndpointsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutConfigSitesInnerEndpointsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutConfigSitesInnerEndpointsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRdata

`func (o *PutConfigSitesInnerEndpointsInner) GetRdata() []EndpointRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PutConfigSitesInnerEndpointsInner) GetRdataOk() (*[]EndpointRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PutConfigSitesInnerEndpointsInner) SetRdata(v []EndpointRdataInner)`

SetRdata sets Rdata field to given value.


### GetMonitorings

`func (o *PutConfigSitesInnerEndpointsInner) GetMonitorings() []PutConfigSitesInnerEndpointsInnerMonitoringsInner`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *PutConfigSitesInnerEndpointsInner) GetMonitoringsOk() (*[]PutConfigSitesInnerEndpointsInnerMonitoringsInner, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *PutConfigSitesInnerEndpointsInner) SetMonitorings(v []PutConfigSitesInnerEndpointsInnerMonitoringsInner)`

SetMonitorings sets Monitorings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


