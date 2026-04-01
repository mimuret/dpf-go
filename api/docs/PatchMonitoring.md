# PatchMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 監視名 | [optional] 
**Description** | Pointer to **string** | コメント | [optional] [default to ""]
**Props** | Pointer to [**PatchMonitoringStaticProps**](PatchMonitoringStaticProps.md) |  | [optional] 

## Methods

### NewPatchMonitoring

`func NewPatchMonitoring() *PatchMonitoring`

NewPatchMonitoring instantiates a new PatchMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMonitoringWithDefaults

`func NewPatchMonitoringWithDefaults() *PatchMonitoring`

NewPatchMonitoringWithDefaults instantiates a new PatchMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchMonitoring) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchMonitoring) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchMonitoring) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchMonitoring) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchMonitoring) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchMonitoring) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchMonitoring) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchMonitoring) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProps

`func (o *PatchMonitoring) GetProps() PatchMonitoringStaticProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *PatchMonitoring) GetPropsOk() (*PatchMonitoringStaticProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *PatchMonitoring) SetProps(v PatchMonitoringStaticProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *PatchMonitoring) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


