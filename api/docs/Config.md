# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitorings** | [**[]Monitoring**](Monitoring.md) |  | [default to []]
**Sites** | [**[]ConfigSitesInner**](ConfigSitesInner.md) |  | [default to []]
**Rules** | [**[]ConfigRulesInner**](ConfigRulesInner.md) |  | 

## Methods

### NewConfig

`func NewConfig(monitorings []Monitoring, sites []ConfigSitesInner, rules []ConfigRulesInner, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorings

`func (o *Config) GetMonitorings() []Monitoring`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *Config) GetMonitoringsOk() (*[]Monitoring, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *Config) SetMonitorings(v []Monitoring)`

SetMonitorings sets Monitorings field to given value.


### GetSites

`func (o *Config) GetSites() []ConfigSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Config) GetSitesOk() (*[]ConfigSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *Config) SetSites(v []ConfigSitesInner)`

SetSites sets Sites field to given value.


### GetRules

`func (o *Config) GetRules() []ConfigRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Config) GetRulesOk() (*[]ConfigRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Config) SetRules(v []ConfigRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


