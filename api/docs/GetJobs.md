# GetJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Status** | **string** | 詳細説明は[**こちら**](#tag/jobs) | 
**ResourcesUrl** | Pointer to **string** | status: SUCCESSFULの場合のみ、レスポンスに含まれます | [optional] 
**ErrorType** | Pointer to **string** | status: FAILEDの場合のみ、レスポンスに含まれます | [optional] 
**ErrorMessage** | Pointer to **string** | status: FAILEDの場合のみ、レスポンスに含まれます | [optional] 

## Methods

### NewGetJobs

`func NewGetJobs(requestId string, status string, ) *GetJobs`

NewGetJobs instantiates a new GetJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobsWithDefaults

`func NewGetJobsWithDefaults() *GetJobs`

NewGetJobsWithDefaults instantiates a new GetJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetJobs) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetJobs) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetJobs) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *GetJobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobs) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResourcesUrl

`func (o *GetJobs) GetResourcesUrl() string`

GetResourcesUrl returns the ResourcesUrl field if non-nil, zero value otherwise.

### GetResourcesUrlOk

`func (o *GetJobs) GetResourcesUrlOk() (*string, bool)`

GetResourcesUrlOk returns a tuple with the ResourcesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesUrl

`func (o *GetJobs) SetResourcesUrl(v string)`

SetResourcesUrl sets ResourcesUrl field to given value.

### HasResourcesUrl

`func (o *GetJobs) HasResourcesUrl() bool`

HasResourcesUrl returns a boolean if a field has been set.

### GetErrorType

`func (o *GetJobs) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *GetJobs) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *GetJobs) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *GetJobs) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetJobs) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetJobs) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetJobs) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetJobs) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


