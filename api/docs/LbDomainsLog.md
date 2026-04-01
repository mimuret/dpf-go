# LbDomainsLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** |  | 
**LogType** | [**LbDomainsLogsType**](LbDomainsLogsType.md) |  | 
**Operator** | **NullableString** | 編集者 | 
**Operation** | [**LbDomainsLogsOperation**](LbDomainsLogsOperation.md) |  | 
**Target** | **string** |  | [default to ""]
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Status** | [**LbDomainsLogsStatus**](LbDomainsLogsStatus.md) |  | 

## Methods

### NewLbDomainsLog

`func NewLbDomainsLog(time time.Time, logType LbDomainsLogsType, operator NullableString, operation LbDomainsLogsOperation, target string, requestId string, status LbDomainsLogsStatus, ) *LbDomainsLog`

NewLbDomainsLog instantiates a new LbDomainsLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbDomainsLogWithDefaults

`func NewLbDomainsLogWithDefaults() *LbDomainsLog`

NewLbDomainsLogWithDefaults instantiates a new LbDomainsLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *LbDomainsLog) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LbDomainsLog) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LbDomainsLog) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetLogType

`func (o *LbDomainsLog) GetLogType() LbDomainsLogsType`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LbDomainsLog) GetLogTypeOk() (*LbDomainsLogsType, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LbDomainsLog) SetLogType(v LbDomainsLogsType)`

SetLogType sets LogType field to given value.


### GetOperator

`func (o *LbDomainsLog) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LbDomainsLog) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LbDomainsLog) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *LbDomainsLog) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *LbDomainsLog) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOperation

`func (o *LbDomainsLog) GetOperation() LbDomainsLogsOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LbDomainsLog) GetOperationOk() (*LbDomainsLogsOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LbDomainsLog) SetOperation(v LbDomainsLogsOperation)`

SetOperation sets Operation field to given value.


### GetTarget

`func (o *LbDomainsLog) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LbDomainsLog) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LbDomainsLog) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetRequestId

`func (o *LbDomainsLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LbDomainsLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LbDomainsLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *LbDomainsLog) GetStatus() LbDomainsLogsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LbDomainsLog) GetStatusOk() (*LbDomainsLogsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LbDomainsLog) SetStatus(v LbDomainsLogsStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


