# MonitoringPingProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**MonitoringPropsLocation**](MonitoringPropsLocation.md) |  | [default to MONITORINGPROPSLOCATION_ALL]
**Interval** | **int32** | 監視間隔（s） | [default to 30]
**Holdtime** | **int32** | 保留時間（s） | [default to 0]
**Timeout** | **int32** | タイムアウト（s） | [default to 5]

## Methods

### NewMonitoringPingProps

`func NewMonitoringPingProps(location MonitoringPropsLocation, interval int32, holdtime int32, timeout int32, ) *MonitoringPingProps`

NewMonitoringPingProps instantiates a new MonitoringPingProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringPingPropsWithDefaults

`func NewMonitoringPingPropsWithDefaults() *MonitoringPingProps`

NewMonitoringPingPropsWithDefaults instantiates a new MonitoringPingProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *MonitoringPingProps) GetLocation() MonitoringPropsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MonitoringPingProps) GetLocationOk() (*MonitoringPropsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MonitoringPingProps) SetLocation(v MonitoringPropsLocation)`

SetLocation sets Location field to given value.


### GetInterval

`func (o *MonitoringPingProps) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MonitoringPingProps) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MonitoringPingProps) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetHoldtime

`func (o *MonitoringPingProps) GetHoldtime() int32`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *MonitoringPingProps) GetHoldtimeOk() (*int32, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *MonitoringPingProps) SetHoldtime(v int32)`

SetHoldtime sets Holdtime field to given value.


### GetTimeout

`func (o *MonitoringPingProps) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MonitoringPingProps) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MonitoringPingProps) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


