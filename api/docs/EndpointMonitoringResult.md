# EndpointMonitoringResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | エンドポイントと監視の紐付けの成功しているかどうかの状態 | [default to "unknown"]

## Methods

### NewEndpointMonitoringResult

`func NewEndpointMonitoringResult(result string, ) *EndpointMonitoringResult`

NewEndpointMonitoringResult instantiates a new EndpointMonitoringResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointMonitoringResultWithDefaults

`func NewEndpointMonitoringResultWithDefaults() *EndpointMonitoringResult`

NewEndpointMonitoringResultWithDefaults instantiates a new EndpointMonitoringResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *EndpointMonitoringResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EndpointMonitoringResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EndpointMonitoringResult) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


