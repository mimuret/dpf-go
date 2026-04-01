# ZonesLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** |  | 
**LogType** | [**ZonesLogsType**](ZonesLogsType.md) |  | 
**Operator** | **NullableString** | 編集者 | 
**Operation** | [**ZonesLogsOperation**](ZonesLogsOperation.md) |  | 
**Target** | **string** |  | [default to ""]
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Status** | [**ZonesLogsStatus**](ZonesLogsStatus.md) |  | 

## Methods

### NewZonesLog

`func NewZonesLog(time time.Time, logType ZonesLogsType, operator NullableString, operation ZonesLogsOperation, target string, requestId string, status ZonesLogsStatus, ) *ZonesLog`

NewZonesLog instantiates a new ZonesLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonesLogWithDefaults

`func NewZonesLogWithDefaults() *ZonesLog`

NewZonesLogWithDefaults instantiates a new ZonesLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ZonesLog) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ZonesLog) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ZonesLog) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetLogType

`func (o *ZonesLog) GetLogType() ZonesLogsType`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *ZonesLog) GetLogTypeOk() (*ZonesLogsType, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *ZonesLog) SetLogType(v ZonesLogsType)`

SetLogType sets LogType field to given value.


### GetOperator

`func (o *ZonesLog) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ZonesLog) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ZonesLog) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *ZonesLog) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ZonesLog) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOperation

`func (o *ZonesLog) GetOperation() ZonesLogsOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ZonesLog) GetOperationOk() (*ZonesLogsOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ZonesLog) SetOperation(v ZonesLogsOperation)`

SetOperation sets Operation field to given value.


### GetTarget

`func (o *ZonesLog) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ZonesLog) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ZonesLog) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetRequestId

`func (o *ZonesLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ZonesLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ZonesLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *ZonesLog) GetStatus() ZonesLogsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZonesLog) GetStatusOk() (*ZonesLogsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZonesLog) SetStatus(v ZonesLogsStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


