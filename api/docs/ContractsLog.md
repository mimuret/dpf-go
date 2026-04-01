# ContractsLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** |  | 
**LogType** | [**ContractsLogsLogType**](ContractsLogsLogType.md) |  | 
**Operator** | **NullableString** | 編集者 | 
**Operation** | [**ContractsLogsOperation**](ContractsLogsOperation.md) |  | 
**Target** | **string** |  | [default to ""]
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Status** | [**ContractsLogsStatus**](ContractsLogsStatus.md) |  | 

## Methods

### NewContractsLog

`func NewContractsLog(time time.Time, logType ContractsLogsLogType, operator NullableString, operation ContractsLogsOperation, target string, requestId string, status ContractsLogsStatus, ) *ContractsLog`

NewContractsLog instantiates a new ContractsLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsLogWithDefaults

`func NewContractsLogWithDefaults() *ContractsLog`

NewContractsLogWithDefaults instantiates a new ContractsLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ContractsLog) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ContractsLog) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ContractsLog) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetLogType

`func (o *ContractsLog) GetLogType() ContractsLogsLogType`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *ContractsLog) GetLogTypeOk() (*ContractsLogsLogType, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *ContractsLog) SetLogType(v ContractsLogsLogType)`

SetLogType sets LogType field to given value.


### GetOperator

`func (o *ContractsLog) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ContractsLog) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ContractsLog) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *ContractsLog) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ContractsLog) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOperation

`func (o *ContractsLog) GetOperation() ContractsLogsOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ContractsLog) GetOperationOk() (*ContractsLogsOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ContractsLog) SetOperation(v ContractsLogsOperation)`

SetOperation sets Operation field to given value.


### GetTarget

`func (o *ContractsLog) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContractsLog) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContractsLog) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetRequestId

`func (o *ContractsLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ContractsLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ContractsLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *ContractsLog) GetStatus() ContractsLogsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractsLog) GetStatusOk() (*ContractsLogsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractsLog) SetStatus(v ContractsLogsStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


