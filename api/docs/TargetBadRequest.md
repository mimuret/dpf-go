# TargetBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**ErrorType** | **string** |  | 
**ErrorMessage** | **string** |  | 
**ErrorDetails** | [**[]TargetErrorDetailsInner**](TargetErrorDetailsInner.md) |  | 

## Methods

### NewTargetBadRequest

`func NewTargetBadRequest(requestId string, errorType string, errorMessage string, errorDetails []TargetErrorDetailsInner, ) *TargetBadRequest`

NewTargetBadRequest instantiates a new TargetBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetBadRequestWithDefaults

`func NewTargetBadRequestWithDefaults() *TargetBadRequest`

NewTargetBadRequestWithDefaults instantiates a new TargetBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TargetBadRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TargetBadRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TargetBadRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrorType

`func (o *TargetBadRequest) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *TargetBadRequest) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *TargetBadRequest) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorMessage

`func (o *TargetBadRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TargetBadRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TargetBadRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetErrorDetails

`func (o *TargetBadRequest) GetErrorDetails() []TargetErrorDetailsInner`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *TargetBadRequest) GetErrorDetailsOk() (*[]TargetErrorDetailsInner, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *TargetBadRequest) SetErrorDetails(v []TargetErrorDetailsInner)`

SetErrorDetails sets ErrorDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


