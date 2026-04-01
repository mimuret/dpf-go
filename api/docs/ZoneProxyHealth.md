# ZoneProxyHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPアドレス | 
**Status** | [**ZoneProxyStatus**](ZoneProxyStatus.md) |  | 
**TsigName** | **string** |  | 
**Enabled** | [**ZoneProxyEnabled**](ZoneProxyEnabled.md) |  | 

## Methods

### NewZoneProxyHealth

`func NewZoneProxyHealth(address string, status ZoneProxyStatus, tsigName string, enabled ZoneProxyEnabled, ) *ZoneProxyHealth`

NewZoneProxyHealth instantiates a new ZoneProxyHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneProxyHealthWithDefaults

`func NewZoneProxyHealthWithDefaults() *ZoneProxyHealth`

NewZoneProxyHealthWithDefaults instantiates a new ZoneProxyHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ZoneProxyHealth) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneProxyHealth) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneProxyHealth) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetStatus

`func (o *ZoneProxyHealth) GetStatus() ZoneProxyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneProxyHealth) GetStatusOk() (*ZoneProxyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneProxyHealth) SetStatus(v ZoneProxyStatus)`

SetStatus sets Status field to given value.


### GetTsigName

`func (o *ZoneProxyHealth) GetTsigName() string`

GetTsigName returns the TsigName field if non-nil, zero value otherwise.

### GetTsigNameOk

`func (o *ZoneProxyHealth) GetTsigNameOk() (*string, bool)`

GetTsigNameOk returns a tuple with the TsigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigName

`func (o *ZoneProxyHealth) SetTsigName(v string)`

SetTsigName sets TsigName field to given value.


### GetEnabled

`func (o *ZoneProxyHealth) GetEnabled() ZoneProxyEnabled`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneProxyHealth) GetEnabledOk() (*ZoneProxyEnabled, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneProxyHealth) SetEnabled(v ZoneProxyEnabled)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


