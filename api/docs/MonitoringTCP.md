# MonitoringTCP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | 監視名 | 
**Mtype** | **string** |  | 
**Description** | **string** | コメント | [default to ""]
**Props** | [**MonitoringTCPProps**](MonitoringTCPProps.md) |  | 

## Methods

### NewMonitoringTCP

`func NewMonitoringTCP(resourceName string, name string, mtype string, description string, props MonitoringTCPProps, ) *MonitoringTCP`

NewMonitoringTCP instantiates a new MonitoringTCP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTCPWithDefaults

`func NewMonitoringTCPWithDefaults() *MonitoringTCP`

NewMonitoringTCPWithDefaults instantiates a new MonitoringTCP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *MonitoringTCP) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MonitoringTCP) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *MonitoringTCP) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *MonitoringTCP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitoringTCP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitoringTCP) SetName(v string)`

SetName sets Name field to given value.


### GetMtype

`func (o *MonitoringTCP) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *MonitoringTCP) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *MonitoringTCP) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetDescription

`func (o *MonitoringTCP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitoringTCP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitoringTCP) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProps

`func (o *MonitoringTCP) GetProps() MonitoringTCPProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *MonitoringTCP) GetPropsOk() (*MonitoringTCPProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *MonitoringTCP) SetProps(v MonitoringTCPProps)`

SetProps sets Props field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


