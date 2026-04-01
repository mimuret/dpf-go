# PutConfigMonitoringsInnerProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | 監視元ロケーション（指定可能な監視種別はping,tcp,http） | [optional] 
**Interval** | Pointer to **int32** | 監視間隔（s）（指定可能な監視種別はping,tcp,http） | [optional] 
**Holdtime** | Pointer to **int32** | 保留時間（s）（指定可能な監視種別はping,tcp,http） | [optional] 
**Timeout** | Pointer to **int32** | タイムアウト（s）（指定可能な監視種別はping,tcp,http） | [optional] 
**Port** | Pointer to **int32** | ポート番号（指定可能な監視種別はtcp,http） | [optional] 
**TlsEnabled** | Pointer to **bool** | TLS監視（指定可能な監視種別はtcp） | [optional] 
**TlsSni** | Pointer to **string** | TLS SNI（指定可能な監視種別はtcp,http） | [optional] 
**ResponseMatch** | Pointer to **string** | レスポンスボディマッチ文字列（指定可能な監視種別はhttp） | [optional] 
**Https** | Pointer to **bool** | HTTPS利用（指定可能な監視種別はhttp） | [optional] 
**Headers** | Pointer to [**[]PutConfigMonitoringsInnerPropsHeadersInner**](PutConfigMonitoringsInnerPropsHeadersInner.md) | Hostヘッダー（指定可能な監視種別はhttp） | [optional] 
**Path** | Pointer to **string** | パス（先頭に「/」は不要。指定可能な監視種別はhttp） | [optional] 
**StatusCodes** | Pointer to **[]string** | ステータスコード（0-9の3桁の文字列のみ指定可能。指定可能な監視種別はhttp） | [optional] 
**Result** | Pointer to **string** | 固定値（指定可能な監視種別はstatic） | [optional] 

## Methods

### NewPutConfigMonitoringsInnerProps

`func NewPutConfigMonitoringsInnerProps() *PutConfigMonitoringsInnerProps`

NewPutConfigMonitoringsInnerProps instantiates a new PutConfigMonitoringsInnerProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutConfigMonitoringsInnerPropsWithDefaults

`func NewPutConfigMonitoringsInnerPropsWithDefaults() *PutConfigMonitoringsInnerProps`

NewPutConfigMonitoringsInnerPropsWithDefaults instantiates a new PutConfigMonitoringsInnerProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PutConfigMonitoringsInnerProps) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PutConfigMonitoringsInnerProps) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PutConfigMonitoringsInnerProps) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PutConfigMonitoringsInnerProps) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInterval

`func (o *PutConfigMonitoringsInnerProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PutConfigMonitoringsInnerProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PutConfigMonitoringsInnerProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PutConfigMonitoringsInnerProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetHoldtime

`func (o *PutConfigMonitoringsInnerProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *PutConfigMonitoringsInnerProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *PutConfigMonitoringsInnerProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *PutConfigMonitoringsInnerProps) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetTimeout

`func (o *PutConfigMonitoringsInnerProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PutConfigMonitoringsInnerProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PutConfigMonitoringsInnerProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PutConfigMonitoringsInnerProps) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPort

`func (o *PutConfigMonitoringsInnerProps) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PutConfigMonitoringsInnerProps) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PutConfigMonitoringsInnerProps) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PutConfigMonitoringsInnerProps) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *PutConfigMonitoringsInnerProps) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *PutConfigMonitoringsInnerProps) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *PutConfigMonitoringsInnerProps) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *PutConfigMonitoringsInnerProps) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetTlsSni

`func (o *PutConfigMonitoringsInnerProps) GetTlsSni() string`

GetTlsSni returns the TlsSni field if non-nil, zero value otherwise.

### GetTlsSniOk

`func (o *PutConfigMonitoringsInnerProps) GetTlsSniOk() (*string, bool)`

GetTlsSniOk returns a tuple with the TlsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSni

`func (o *PutConfigMonitoringsInnerProps) SetTlsSni(v string)`

SetTlsSni sets TlsSni field to given value.

### HasTlsSni

`func (o *PutConfigMonitoringsInnerProps) HasTlsSni() bool`

HasTlsSni returns a boolean if a field has been set.

### GetResponseMatch

`func (o *PutConfigMonitoringsInnerProps) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *PutConfigMonitoringsInnerProps) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *PutConfigMonitoringsInnerProps) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.

### HasResponseMatch

`func (o *PutConfigMonitoringsInnerProps) HasResponseMatch() bool`

HasResponseMatch returns a boolean if a field has been set.

### GetHttps

`func (o *PutConfigMonitoringsInnerProps) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *PutConfigMonitoringsInnerProps) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *PutConfigMonitoringsInnerProps) SetHttps(v bool)`

SetHttps sets Https field to given value.

### HasHttps

`func (o *PutConfigMonitoringsInnerProps) HasHttps() bool`

HasHttps returns a boolean if a field has been set.

### GetHeaders

`func (o *PutConfigMonitoringsInnerProps) GetHeaders() []PutConfigMonitoringsInnerPropsHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PutConfigMonitoringsInnerProps) GetHeadersOk() (*[]PutConfigMonitoringsInnerPropsHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PutConfigMonitoringsInnerProps) SetHeaders(v []PutConfigMonitoringsInnerPropsHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PutConfigMonitoringsInnerProps) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPath

`func (o *PutConfigMonitoringsInnerProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PutConfigMonitoringsInnerProps) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PutConfigMonitoringsInnerProps) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PutConfigMonitoringsInnerProps) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatusCodes

`func (o *PutConfigMonitoringsInnerProps) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *PutConfigMonitoringsInnerProps) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *PutConfigMonitoringsInnerProps) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *PutConfigMonitoringsInnerProps) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### GetResult

`func (o *PutConfigMonitoringsInnerProps) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PutConfigMonitoringsInnerProps) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PutConfigMonitoringsInnerProps) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *PutConfigMonitoringsInnerProps) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


