# EndpointMonitoringsMonitoringsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | 監視名 | 
**Mtype** | **string** |  | 
**Description** | **string** | コメント | [default to ""]
**Props** | [**MonitoringStaticProps**](MonitoringStaticProps.md) |  | 
**Enabled** | **bool** | エンドポイントと監視の紐付け状態 | [default to false]
**Result** | **string** | エンドポイントと監視の紐付けの成功しているかどうかの状態 | [default to "unknown"]

## Methods

### NewEndpointMonitoringsMonitoringsInner

`func NewEndpointMonitoringsMonitoringsInner(resourceName string, name string, mtype string, description string, props MonitoringStaticProps, enabled bool, result string, ) *EndpointMonitoringsMonitoringsInner`

NewEndpointMonitoringsMonitoringsInner instantiates a new EndpointMonitoringsMonitoringsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointMonitoringsMonitoringsInnerWithDefaults

`func NewEndpointMonitoringsMonitoringsInnerWithDefaults() *EndpointMonitoringsMonitoringsInner`

NewEndpointMonitoringsMonitoringsInnerWithDefaults instantiates a new EndpointMonitoringsMonitoringsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *EndpointMonitoringsMonitoringsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *EndpointMonitoringsMonitoringsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *EndpointMonitoringsMonitoringsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *EndpointMonitoringsMonitoringsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointMonitoringsMonitoringsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointMonitoringsMonitoringsInner) SetName(v string)`

SetName sets Name field to given value.


### GetMtype

`func (o *EndpointMonitoringsMonitoringsInner) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *EndpointMonitoringsMonitoringsInner) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *EndpointMonitoringsMonitoringsInner) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetDescription

`func (o *EndpointMonitoringsMonitoringsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointMonitoringsMonitoringsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointMonitoringsMonitoringsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProps

`func (o *EndpointMonitoringsMonitoringsInner) GetProps() MonitoringStaticProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *EndpointMonitoringsMonitoringsInner) GetPropsOk() (*MonitoringStaticProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *EndpointMonitoringsMonitoringsInner) SetProps(v MonitoringStaticProps)`

SetProps sets Props field to given value.


### GetEnabled

`func (o *EndpointMonitoringsMonitoringsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EndpointMonitoringsMonitoringsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EndpointMonitoringsMonitoringsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetResult

`func (o *EndpointMonitoringsMonitoringsInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EndpointMonitoringsMonitoringsInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EndpointMonitoringsMonitoringsInner) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


