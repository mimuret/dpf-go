# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**ErrorType** | **string** |  | 
**ErrorMessage** | **string** |  | 
**ErrorDetails** | [**[]ErrorDetailsInner**](ErrorDetailsInner.md) |  | 

## Methods

### NewBadRequest

`func NewBadRequest(requestId string, errorType string, errorMessage string, errorDetails []ErrorDetailsInner, ) *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *BadRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BadRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BadRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrorType

`func (o *BadRequest) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *BadRequest) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *BadRequest) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorMessage

`func (o *BadRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BadRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BadRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetErrorDetails

`func (o *BadRequest) GetErrorDetails() []ErrorDetailsInner`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *BadRequest) GetErrorDetailsOk() (*[]ErrorDetailsInner, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *BadRequest) SetErrorDetails(v []ErrorDetailsInner)`

SetErrorDetails sets ErrorDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


