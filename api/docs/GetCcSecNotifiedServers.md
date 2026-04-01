# GetCcSecNotifiedServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]NotifiedServer**](NotifiedServer.md) |  | [default to []]

## Methods

### NewGetCcSecNotifiedServers

`func NewGetCcSecNotifiedServers(requestId string, results []NotifiedServer, ) *GetCcSecNotifiedServers`

NewGetCcSecNotifiedServers instantiates a new GetCcSecNotifiedServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCcSecNotifiedServersWithDefaults

`func NewGetCcSecNotifiedServersWithDefaults() *GetCcSecNotifiedServers`

NewGetCcSecNotifiedServersWithDefaults instantiates a new GetCcSecNotifiedServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetCcSecNotifiedServers) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCcSecNotifiedServers) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCcSecNotifiedServers) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetCcSecNotifiedServers) GetResults() []NotifiedServer`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetCcSecNotifiedServers) GetResultsOk() (*[]NotifiedServer, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetCcSecNotifiedServers) SetResults(v []NotifiedServer)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


