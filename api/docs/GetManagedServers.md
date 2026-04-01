# GetManagedServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | **[]string** |  | [default to []]

## Methods

### NewGetManagedServers

`func NewGetManagedServers(requestId string, results []string, ) *GetManagedServers`

NewGetManagedServers instantiates a new GetManagedServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedServersWithDefaults

`func NewGetManagedServersWithDefaults() *GetManagedServers`

NewGetManagedServersWithDefaults instantiates a new GetManagedServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetManagedServers) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetManagedServers) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetManagedServers) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetManagedServers) GetResults() []string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetManagedServers) GetResultsOk() (*[]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetManagedServers) SetResults(v []string)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


