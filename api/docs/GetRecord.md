# GetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**Record**](Record.md) |  | 

## Methods

### NewGetRecord

`func NewGetRecord(requestId string, result Record, ) *GetRecord`

NewGetRecord instantiates a new GetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordWithDefaults

`func NewGetRecordWithDefaults() *GetRecord`

NewGetRecordWithDefaults instantiates a new GetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRecord) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRecord) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRecord) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetRecord) GetResult() Record`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecord) GetResultOk() (*Record, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecord) SetResult(v Record)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


