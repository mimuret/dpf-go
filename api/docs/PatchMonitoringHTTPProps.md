# PatchMonitoringHTTPProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [optional] [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | Pointer to **int32** | 監視間隔（s） | [optional] [default to 30]
**Holdtime** | Pointer to **int32** | 保留時間（s） | [optional] [default to 0]
**Timeout** | Pointer to **int32** | タイムアウト（s） | [optional] [default to 5]
**Port** | Pointer to **int32** | ポート番号,POST時に未指定の場合はHTTPの場合は80,HTTPSの場合は443 | [optional] 
**Https** | Pointer to **bool** | HTTPS利用フラグ | [optional] [default to false]
**TlsSni** | Pointer to **string** | TLS SNI値、未指定の場合、監視時にSNIとして、Hostヘッダーを利用 | [optional] 
**ResponseMatch** | Pointer to **string** | レスポンスボディマッチ文字列 | [optional] [default to ""]
**Headers** | Pointer to [**[]MonitoringPropsHeader**](MonitoringPropsHeader.md) | Hostヘッダー | [optional] 
**Path** | Pointer to **string** | URLのPATH部(先頭の/が補完されて利用される) | [optional] 
**StatusCodes** | Pointer to **[]string** | マッチするステータスコードの配列 | [optional] [default to []]

## Methods

### NewPatchMonitoringHTTPProps

`func NewPatchMonitoringHTTPProps() *PatchMonitoringHTTPProps`

NewPatchMonitoringHTTPProps instantiates a new PatchMonitoringHTTPProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMonitoringHTTPPropsWithDefaults

`func NewPatchMonitoringHTTPPropsWithDefaults() *PatchMonitoringHTTPProps`

NewPatchMonitoringHTTPPropsWithDefaults instantiates a new PatchMonitoringHTTPProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PatchMonitoringHTTPProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchMonitoringHTTPProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchMonitoringHTTPProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchMonitoringHTTPProps) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInterval

`func (o *PatchMonitoringHTTPProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchMonitoringHTTPProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchMonitoringHTTPProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchMonitoringHTTPProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetHoldtime

`func (o *PatchMonitoringHTTPProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *PatchMonitoringHTTPProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *PatchMonitoringHTTPProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *PatchMonitoringHTTPProps) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchMonitoringHTTPProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchMonitoringHTTPProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchMonitoringHTTPProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchMonitoringHTTPProps) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPort

`func (o *PatchMonitoringHTTPProps) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchMonitoringHTTPProps) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchMonitoringHTTPProps) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchMonitoringHTTPProps) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHttps

`func (o *PatchMonitoringHTTPProps) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *PatchMonitoringHTTPProps) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *PatchMonitoringHTTPProps) SetHttps(v bool)`

SetHttps sets Https field to given value.

### HasHttps

`func (o *PatchMonitoringHTTPProps) HasHttps() bool`

HasHttps returns a boolean if a field has been set.

### GetTlsSni

`func (o *PatchMonitoringHTTPProps) GetTlsSni() string`

GetTlsSni returns the TlsSni field if non-nil, zero value otherwise.

### GetTlsSniOk

`func (o *PatchMonitoringHTTPProps) GetTlsSniOk() (*string, bool)`

GetTlsSniOk returns a tuple with the TlsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSni

`func (o *PatchMonitoringHTTPProps) SetTlsSni(v string)`

SetTlsSni sets TlsSni field to given value.

### HasTlsSni

`func (o *PatchMonitoringHTTPProps) HasTlsSni() bool`

HasTlsSni returns a boolean if a field has been set.

### GetResponseMatch

`func (o *PatchMonitoringHTTPProps) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *PatchMonitoringHTTPProps) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *PatchMonitoringHTTPProps) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.

### HasResponseMatch

`func (o *PatchMonitoringHTTPProps) HasResponseMatch() bool`

HasResponseMatch returns a boolean if a field has been set.

### GetHeaders

`func (o *PatchMonitoringHTTPProps) GetHeaders() []MonitoringPropsHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatchMonitoringHTTPProps) GetHeadersOk() (*[]MonitoringPropsHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatchMonitoringHTTPProps) SetHeaders(v []MonitoringPropsHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatchMonitoringHTTPProps) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPath

`func (o *PatchMonitoringHTTPProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchMonitoringHTTPProps) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchMonitoringHTTPProps) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchMonitoringHTTPProps) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatusCodes

`func (o *PatchMonitoringHTTPProps) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *PatchMonitoringHTTPProps) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *PatchMonitoringHTTPProps) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *PatchMonitoringHTTPProps) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


