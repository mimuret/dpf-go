# PutConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitorings** | [**[]PutConfigMonitoringsInner**](PutConfigMonitoringsInner.md) |  | 
**Sites** | [**[]PutConfigSitesInner**](PutConfigSitesInner.md) |  | 
**Rules** | [**[]PutConfigRulesInner**](PutConfigRulesInner.md) |  | 

## Methods

### NewPutConfig

`func NewPutConfig(monitorings []PutConfigMonitoringsInner, sites []PutConfigSitesInner, rules []PutConfigRulesInner, ) *PutConfig`

NewPutConfig instantiates a new PutConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutConfigWithDefaults

`func NewPutConfigWithDefaults() *PutConfig`

NewPutConfigWithDefaults instantiates a new PutConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorings

`func (o *PutConfig) GetMonitorings() []PutConfigMonitoringsInner`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *PutConfig) GetMonitoringsOk() (*[]PutConfigMonitoringsInner, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *PutConfig) SetMonitorings(v []PutConfigMonitoringsInner)`

SetMonitorings sets Monitorings field to given value.


### GetSites

`func (o *PutConfig) GetSites() []PutConfigSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PutConfig) GetSitesOk() (*[]PutConfigSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PutConfig) SetSites(v []PutConfigSitesInner)`

SetSites sets Sites field to given value.


### GetRules

`func (o *PutConfig) GetRules() []PutConfigRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PutConfig) GetRulesOk() (*[]PutConfigRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PutConfig) SetRules(v []PutConfigRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


