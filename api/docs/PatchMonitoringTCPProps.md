# PatchMonitoringTCPProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [optional] [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | Pointer to **int32** | 監視間隔（s） | [optional] [default to 30]
**Holdtime** | Pointer to **int32** | 保留時間（s） | [optional] [default to 0]
**Timeout** | Pointer to **int32** | タイムアウト（s） | [optional] [default to 5]
**Port** | Pointer to **int32** | ポート番号 | [optional] 
**TlsEnabled** | Pointer to **bool** | TLS監視利用フラグ | [optional] [default to false]
**TlsSni** | Pointer to **string** | TLS SNI値、未指定の場合、監視時にSNIとして、Hostヘッダーを利用 | [optional] 

## Methods

### NewPatchMonitoringTCPProps

`func NewPatchMonitoringTCPProps() *PatchMonitoringTCPProps`

NewPatchMonitoringTCPProps instantiates a new PatchMonitoringTCPProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMonitoringTCPPropsWithDefaults

`func NewPatchMonitoringTCPPropsWithDefaults() *PatchMonitoringTCPProps`

NewPatchMonitoringTCPPropsWithDefaults instantiates a new PatchMonitoringTCPProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PatchMonitoringTCPProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchMonitoringTCPProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchMonitoringTCPProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchMonitoringTCPProps) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInterval

`func (o *PatchMonitoringTCPProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchMonitoringTCPProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchMonitoringTCPProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchMonitoringTCPProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetHoldtime

`func (o *PatchMonitoringTCPProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *PatchMonitoringTCPProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *PatchMonitoringTCPProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *PatchMonitoringTCPProps) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchMonitoringTCPProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchMonitoringTCPProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchMonitoringTCPProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchMonitoringTCPProps) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPort

`func (o *PatchMonitoringTCPProps) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchMonitoringTCPProps) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchMonitoringTCPProps) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchMonitoringTCPProps) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *PatchMonitoringTCPProps) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *PatchMonitoringTCPProps) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *PatchMonitoringTCPProps) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *PatchMonitoringTCPProps) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetTlsSni

`func (o *PatchMonitoringTCPProps) GetTlsSni() string`

GetTlsSni returns the TlsSni field if non-nil, zero value otherwise.

### GetTlsSniOk

`func (o *PatchMonitoringTCPProps) GetTlsSniOk() (*string, bool)`

GetTlsSniOk returns a tuple with the TlsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSni

`func (o *PatchMonitoringTCPProps) SetTlsSni(v string)`

SetTlsSni sets TlsSni field to given value.

### HasTlsSni

`func (o *PatchMonitoringTCPProps) HasTlsSni() bool`

HasTlsSni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


