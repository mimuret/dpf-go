# Dnssec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | [**DnssecEnabled**](DnssecEnabled.md) |  | 
**State** | [**DnssecState**](DnssecState.md) |  | 
**DsState** | [**DnssecDsState**](DnssecDsState.md) |  | 

## Methods

### NewDnssec

`func NewDnssec(enabled DnssecEnabled, state DnssecState, dsState DnssecDsState, ) *Dnssec`

NewDnssec instantiates a new Dnssec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecWithDefaults

`func NewDnssecWithDefaults() *Dnssec`

NewDnssecWithDefaults instantiates a new Dnssec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Dnssec) GetEnabled() DnssecEnabled`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Dnssec) GetEnabledOk() (*DnssecEnabled, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Dnssec) SetEnabled(v DnssecEnabled)`

SetEnabled sets Enabled field to given value.


### GetState

`func (o *Dnssec) GetState() DnssecState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dnssec) GetStateOk() (*DnssecState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dnssec) SetState(v DnssecState)`

SetState sets State field to given value.


### GetDsState

`func (o *Dnssec) GetDsState() DnssecDsState`

GetDsState returns the DsState field if non-nil, zero value otherwise.

### GetDsStateOk

`func (o *Dnssec) GetDsStateOk() (*DnssecDsState, bool)`

GetDsStateOk returns a tuple with the DsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsState

`func (o *Dnssec) SetDsState(v DnssecDsState)`

SetDsState sets DsState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


