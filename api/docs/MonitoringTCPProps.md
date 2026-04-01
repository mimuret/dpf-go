# MonitoringTCPProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | **int32** | 監視間隔（s） | [default to 30]
**Holdtime** | **int32** | 保留時間（s） | [default to 0]
**Timeout** | **int32** | タイムアウト（s） | [default to 5]
**Port** | **int32** | ポート番号 | 
**TlsEnabled** | **bool** | TLS監視利用フラグ | [default to false]
**TlsSni** | **string** | TLS SNI値、未指定の場合、監視時にSNIとして、Hostヘッダーを利用 | 

## Methods

### NewMonitoringTCPProps

`func NewMonitoringTCPProps(location MonitoringPropsLocation, interval int32, holdtime int32, timeout int32, port int32, tlsEnabled bool, tlsSni string, ) *MonitoringTCPProps`

NewMonitoringTCPProps instantiates a new MonitoringTCPProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTCPPropsWithDefaults

`func NewMonitoringTCPPropsWithDefaults() *MonitoringTCPProps`

NewMonitoringTCPPropsWithDefaults instantiates a new MonitoringTCPProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *MonitoringTCPProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringTCPProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringTCPProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.


### GetInterval

`func (o *MonitoringTCPProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MonitoringTCPProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MonitoringTCPProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetHoldtime

`func (o *MonitoringTCPProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *MonitoringTCPProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *MonitoringTCPProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.


### GetTimeout

`func (o *MonitoringTCPProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MonitoringTCPProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MonitoringTCPProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetPort

`func (o *MonitoringTCPProps) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MonitoringTCPProps) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MonitoringTCPProps) SetPort(v int32)`

SetPort sets Port field to given value.


### GetTlsEnabled

`func (o *MonitoringTCPProps) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *MonitoringTCPProps) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *MonitoringTCPProps) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.


### GetTlsSni

`func (o *MonitoringTCPProps) GetTlsSni() string`

GetTlsSni returns the TlsSni field if non-nil, zero value otherwise.

### GetTlsSniOk

`func (o *MonitoringTCPProps) GetTlsSniOk() (*string, bool)`

GetTlsSniOk returns a tuple with the TlsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSni

`func (o *MonitoringTCPProps) SetTlsSni(v string)`

SetTlsSni sets TlsSni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


