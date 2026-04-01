# GetDsRecordsFromCds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]GetDsRecordsFromCds200ResponseAllOfResultsInner**](GetDsRecordsFromCds200ResponseAllOfResultsInner.md) |  | [default to []]

## Methods

### NewGetDsRecordsFromCds200Response

`func NewGetDsRecordsFromCds200Response(requestId string, results []GetDsRecordsFromCds200ResponseAllOfResultsInner, ) *GetDsRecordsFromCds200Response`

NewGetDsRecordsFromCds200Response instantiates a new GetDsRecordsFromCds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDsRecordsFromCds200ResponseWithDefaults

`func NewGetDsRecordsFromCds200ResponseWithDefaults() *GetDsRecordsFromCds200Response`

NewGetDsRecordsFromCds200ResponseWithDefaults instantiates a new GetDsRecordsFromCds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetDsRecordsFromCds200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetDsRecordsFromCds200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetDsRecordsFromCds200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetDsRecordsFromCds200Response) GetResults() []GetDsRecordsFromCds200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetDsRecordsFromCds200Response) GetResultsOk() (*[]GetDsRecordsFromCds200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetDsRecordsFromCds200Response) SetResults(v []GetDsRecordsFromCds200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


