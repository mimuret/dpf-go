# GetZonesLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]ZonesLog**](ZonesLog.md) |  | [default to []]

## Methods

### NewGetZonesLogs

`func NewGetZonesLogs(requestId string, results []ZonesLog, ) *GetZonesLogs`

NewGetZonesLogs instantiates a new GetZonesLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZonesLogsWithDefaults

`func NewGetZonesLogsWithDefaults() *GetZonesLogs`

NewGetZonesLogsWithDefaults instantiates a new GetZonesLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetZonesLogs) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZonesLogs) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZonesLogs) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetZonesLogs) GetResults() []ZonesLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetZonesLogs) GetResultsOk() (*[]ZonesLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetZonesLogs) SetResults(v []ZonesLog)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


