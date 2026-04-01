# PatchMonitoringPingProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [optional] [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | Pointer to **int32** | 監視間隔（s） | [optional] [default to 30]
**Holdtime** | Pointer to **int32** | 保留時間（s） | [optional] [default to 0]
**Timeout** | Pointer to **int32** | タイムアウト（s） | [optional] [default to 5]

## Methods

### NewPatchMonitoringPingProps

`func NewPatchMonitoringPingProps() *PatchMonitoringPingProps`

NewPatchMonitoringPingProps instantiates a new PatchMonitoringPingProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMonitoringPingPropsWithDefaults

`func NewPatchMonitoringPingPropsWithDefaults() *PatchMonitoringPingProps`

NewPatchMonitoringPingPropsWithDefaults instantiates a new PatchMonitoringPingProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PatchMonitoringPingProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchMonitoringPingProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchMonitoringPingProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchMonitoringPingProps) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInterval

`func (o *PatchMonitoringPingProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchMonitoringPingProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchMonitoringPingProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchMonitoringPingProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetHoldtime

`func (o *PatchMonitoringPingProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *PatchMonitoringPingProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *PatchMonitoringPingProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *PatchMonitoringPingProps) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchMonitoringPingProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchMonitoringPingProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchMonitoringPingProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchMonitoringPingProps) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


