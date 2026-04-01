# MonitoringHTTPProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | **int32** | 監視間隔（s） | [default to 30]
**Holdtime** | **int32** | 保留時間（s） | [default to 0]
**Timeout** | **int32** | タイムアウト（s） | [default to 5]
**Port** | **int32** | ポート番号,POST時に未指定の場合はHTTPの場合は80,HTTPSの場合は443 | 
**Https** | **bool** | HTTPS利用フラグ | [default to false]
**TlsSni** | **string** | TLS SNI値、未指定の場合、監視時にSNIとして、Hostヘッダーを利用 | 
**ResponseMatch** | **string** | レスポンスボディマッチ文字列 | [default to ""]
**Headers** | [**[]MonitoringPropsHeader**](MonitoringPropsHeader.md) | Hostヘッダー | 
**Path** | **string** | URLのPATH部(先頭の/が補完されて利用される) | 
**StatusCodes** | **[]string** | マッチするステータスコードの配列 | [default to []]

## Methods

### NewMonitoringHTTPProps

`func NewMonitoringHTTPProps(location MonitoringPropsLocation, interval int32, holdtime int32, timeout int32, port int32, https bool, tlsSni string, responseMatch string, headers []MonitoringPropsHeader, path string, statusCodes []string, ) *MonitoringHTTPProps`

NewMonitoringHTTPProps instantiates a new MonitoringHTTPProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringHTTPPropsWithDefaults

`func NewMonitoringHTTPPropsWithDefaults() *MonitoringHTTPProps`

NewMonitoringHTTPPropsWithDefaults instantiates a new MonitoringHTTPProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *MonitoringHTTPProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringHTTPProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringHTTPProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.


### GetInterval

`func (o *MonitoringHTTPProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MonitoringHTTPProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MonitoringHTTPProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetHoldtime

`func (o *MonitoringHTTPProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *MonitoringHTTPProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *MonitoringHTTPProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.


### GetTimeout

`func (o *MonitoringHTTPProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MonitoringHTTPProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MonitoringHTTPProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetPort

`func (o *MonitoringHTTPProps) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MonitoringHTTPProps) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MonitoringHTTPProps) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHttps

`func (o *MonitoringHTTPProps) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *MonitoringHTTPProps) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *MonitoringHTTPProps) SetHttps(v bool)`

SetHttps sets Https field to given value.


### GetTlsSni

`func (o *MonitoringHTTPProps) GetTlsSni() string`

GetTlsSni returns the TlsSni field if non-nil, zero value otherwise.

### GetTlsSniOk

`func (o *MonitoringHTTPProps) GetTlsSniOk() (*string, bool)`

GetTlsSniOk returns a tuple with the TlsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSni

`func (o *MonitoringHTTPProps) SetTlsSni(v string)`

SetTlsSni sets TlsSni field to given value.


### GetResponseMatch

`func (o *MonitoringHTTPProps) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *MonitoringHTTPProps) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *MonitoringHTTPProps) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetHeaders

`func (o *MonitoringHTTPProps) GetHeaders() []MonitoringPropsHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MonitoringHTTPProps) GetHeadersOk() (*[]MonitoringPropsHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MonitoringHTTPProps) SetHeaders(v []MonitoringPropsHeader)`

SetHeaders sets Headers field to given value.


### GetPath

`func (o *MonitoringHTTPProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MonitoringHTTPProps) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MonitoringHTTPProps) SetPath(v string)`

SetPath sets Path field to given value.


### GetStatusCodes

`func (o *MonitoringHTTPProps) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *MonitoringHTTPProps) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *MonitoringHTTPProps) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


