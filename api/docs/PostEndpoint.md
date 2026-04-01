# PostEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | [optional] 
**Name** | **string** | エンドポイント名 | 
**MonitoringTarget** | Pointer to **string** | 監視ターゲット（IPアドレス形式かホスト名形式） | [optional] [default to ""]
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**Weight** | **int32** | weight | 
**ManualFailback** | Pointer to **bool** | 手動切り戻し設定値 | [optional] [default to false]
**ManualFailover** | Pointer to **bool** | 手動切り離し設定値 | [optional] [default to false]
**Enabled** | Pointer to **bool** | 状態 | [optional] [default to false]
**Rdata** | [**[]EndpointRdataInner**](EndpointRdataInner.md) | RDATA | 
**Monitorings** | Pointer to [**[]PostEndpointMonitoringsInner**](PostEndpointMonitoringsInner.md) |  | [optional] [default to []]

## Methods

### NewPostEndpoint

`func NewPostEndpoint(name string, weight int32, rdata []EndpointRdataInner, ) *PostEndpoint`

NewPostEndpoint instantiates a new PostEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEndpointWithDefaults

`func NewPostEndpointWithDefaults() *PostEndpoint`

NewPostEndpointWithDefaults instantiates a new PostEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PostEndpoint) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PostEndpoint) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PostEndpoint) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PostEndpoint) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetName

`func (o *PostEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetMonitoringTarget

`func (o *PostEndpoint) GetMonitoringTarget() string`

GetMonitoringTarget returns the MonitoringTarget field if non-nil, zero value otherwise.

### GetMonitoringTargetOk

`func (o *PostEndpoint) GetMonitoringTargetOk() (*string, bool)`

GetMonitoringTargetOk returns a tuple with the MonitoringTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTarget

`func (o *PostEndpoint) SetMonitoringTarget(v string)`

SetMonitoringTarget sets MonitoringTarget field to given value.

### HasMonitoringTarget

`func (o *PostEndpoint) HasMonitoringTarget() bool`

HasMonitoringTarget returns a boolean if a field has been set.

### GetDescription

`func (o *PostEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *PostEndpoint) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PostEndpoint) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PostEndpoint) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetManualFailback

`func (o *PostEndpoint) GetManualFailback() bool`

GetManualFailback returns the ManualFailback field if non-nil, zero value otherwise.

### GetManualFailbackOk

`func (o *PostEndpoint) GetManualFailbackOk() (*bool, bool)`

GetManualFailbackOk returns a tuple with the ManualFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailback

`func (o *PostEndpoint) SetManualFailback(v bool)`

SetManualFailback sets ManualFailback field to given value.

### HasManualFailback

`func (o *PostEndpoint) HasManualFailback() bool`

HasManualFailback returns a boolean if a field has been set.

### GetManualFailover

`func (o *PostEndpoint) GetManualFailover() bool`

GetManualFailover returns the ManualFailover field if non-nil, zero value otherwise.

### GetManualFailoverOk

`func (o *PostEndpoint) GetManualFailoverOk() (*bool, bool)`

GetManualFailoverOk returns a tuple with the ManualFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailover

`func (o *PostEndpoint) SetManualFailover(v bool)`

SetManualFailover sets ManualFailover field to given value.

### HasManualFailover

`func (o *PostEndpoint) HasManualFailover() bool`

HasManualFailover returns a boolean if a field has been set.

### GetEnabled

`func (o *PostEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostEndpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRdata

`func (o *PostEndpoint) GetRdata() []EndpointRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PostEndpoint) GetRdataOk() (*[]EndpointRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PostEndpoint) SetRdata(v []EndpointRdataInner)`

SetRdata sets Rdata field to given value.


### GetMonitorings

`func (o *PostEndpoint) GetMonitorings() []PostEndpointMonitoringsInner`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *PostEndpoint) GetMonitoringsOk() (*[]PostEndpointMonitoringsInner, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *PostEndpoint) SetMonitorings(v []PostEndpointMonitoringsInner)`

SetMonitorings sets Monitorings field to given value.

### HasMonitorings

`func (o *PostEndpoint) HasMonitorings() bool`

HasMonitorings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


