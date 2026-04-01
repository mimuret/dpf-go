# GetRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]Record**](Record.md) |  | [default to []]

## Methods

### NewGetRecords

`func NewGetRecords(requestId string, results []Record, ) *GetRecords`

NewGetRecords instantiates a new GetRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordsWithDefaults

`func NewGetRecordsWithDefaults() *GetRecords`

NewGetRecordsWithDefaults instantiates a new GetRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRecords) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRecords) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRecords) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetRecords) GetResults() []Record`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetRecords) GetResultsOk() (*[]Record, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetRecords) SetResults(v []Record)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


