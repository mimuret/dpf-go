# PostMonitoringPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | [optional] 
**Name** | **string** | 監視名 | 
**Mtype** | **string** |  | 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**Props** | Pointer to [**PatchMonitoringPingProps**](PatchMonitoringPingProps.md) |  | [optional] 

## Methods

### NewPostMonitoringPing

`func NewPostMonitoringPing(name string, mtype string, ) *PostMonitoringPing`

NewPostMonitoringPing instantiates a new PostMonitoringPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMonitoringPingWithDefaults

`func NewPostMonitoringPingWithDefaults() *PostMonitoringPing`

NewPostMonitoringPingWithDefaults instantiates a new PostMonitoringPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *PostMonitoringPing) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PostMonitoringPing) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PostMonitoringPing) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PostMonitoringPing) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetName

`func (o *PostMonitoringPing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostMonitoringPing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostMonitoringPing) SetName(v string)`

SetName sets Name field to given value.


### GetMtype

`func (o *PostMonitoringPing) GetMtype() string`

GetMtype returns the Mtype field if non-nil, zero value otherwise.

### GetMtypeOk

`func (o *PostMonitoringPing) GetMtypeOk() (*string, bool)`

GetMtypeOk returns a tuple with the Mtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtype

`func (o *PostMonitoringPing) SetMtype(v string)`

SetMtype sets Mtype field to given value.


### GetDescription

`func (o *PostMonitoringPing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostMonitoringPing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostMonitoringPing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostMonitoringPing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProps

`func (o *PostMonitoringPing) GetProps() PatchMonitoringPingProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *PostMonitoringPing) GetPropsOk() (*PatchMonitoringPingProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *PostMonitoringPing) SetProps(v PatchMonitoringPingProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *PostMonitoringPing) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


