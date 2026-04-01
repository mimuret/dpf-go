# GetZoneProxyHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]ZoneProxyHealth**](ZoneProxyHealth.md) |  | [default to []]

## Methods

### NewGetZoneProxyHealth

`func NewGetZoneProxyHealth(requestId string, results []ZoneProxyHealth, ) *GetZoneProxyHealth`

NewGetZoneProxyHealth instantiates a new GetZoneProxyHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneProxyHealthWithDefaults

`func NewGetZoneProxyHealthWithDefaults() *GetZoneProxyHealth`

NewGetZoneProxyHealthWithDefaults instantiates a new GetZoneProxyHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetZoneProxyHealth) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZoneProxyHealth) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZoneProxyHealth) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetZoneProxyHealth) GetResults() []ZoneProxyHealth`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetZoneProxyHealth) GetResultsOk() (*[]ZoneProxyHealth, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetZoneProxyHealth) SetResults(v []ZoneProxyHealth)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


