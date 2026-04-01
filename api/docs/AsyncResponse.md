# AsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**JobsUrl** | **string** |  | 

## Methods

### NewAsyncResponse

`func NewAsyncResponse(requestId string, jobsUrl string, ) *AsyncResponse`

NewAsyncResponse instantiates a new AsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncResponseWithDefaults

`func NewAsyncResponseWithDefaults() *AsyncResponse`

NewAsyncResponseWithDefaults instantiates a new AsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AsyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AsyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AsyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetJobsUrl

`func (o *AsyncResponse) GetJobsUrl() string`

GetJobsUrl returns the JobsUrl field if non-nil, zero value otherwise.

### GetJobsUrlOk

`func (o *AsyncResponse) GetJobsUrlOk() (*string, bool)`

GetJobsUrlOk returns a tuple with the JobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsUrl

`func (o *AsyncResponse) SetJobsUrl(v string)`

SetJobsUrl sets JobsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


