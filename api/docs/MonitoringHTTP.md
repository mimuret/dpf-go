# MonitoringHTTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | 監視名 | 
**Mtype** | **string** |  | 
**Description** | **string** | コメント | [default to ""]
**Props** | [**MonitoringHTTPProps**](MonitoringHTTPProps.md) |  | 

## Methods

### NewMonitoringHTTP

`func NewMonitoringHTTP(resourceName string, name string, mtype string, description string, props MonitoringHTTPProps, ) *MonitoringHTTP`

NewMonitoringHTTP instantiates a new MonitoringHTTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringHTTPWithDefaults

`func NewMonitoringHTTPWithDefaults() *MonitoringHTTP`

NewMonitoringHTTPWithDefaults instantiates a new MonitoringHTTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *MonitoringHTTP) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MonitoringHTTP) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *MonitoringHTTP) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *MonitoringHTTP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitoringHTTP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitoringHTTP) SetName(v string)`

SetName sets Name field to given value.


### GetMtype

`func (o *MonitoringHTTP) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *MonitoringHTTP) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *MonitoringHTTP) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetDescription

`func (o *MonitoringHTTP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitoringHTTP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitoringHTTP) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProps

`func (o *MonitoringHTTP) GetProps() MonitoringHTTPProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *MonitoringHTTP) GetPropsOk() (*MonitoringHTTPProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *MonitoringHTTP) SetProps(v MonitoringHTTPProps)`

SetProps sets Props field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


