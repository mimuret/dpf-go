# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ServiceCode** | **string** |  | 
**Name** | **string** | name | 
**Network** | **NullableString** |  | 
**Description** | **string** | コメント | [default to ""]
**DelegationRequestedAt** | **time.Time** |  | 

## Methods

### NewDelegation

`func NewDelegation(id string, serviceCode string, name string, network NullableString, description string, delegationRequestedAt time.Time, ) *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Delegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation) SetId(v string)`

SetId sets Id field to given value.


### GetServiceCode

`func (o *Delegation) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *Delegation) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *Delegation) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetName

`func (o *Delegation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Delegation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Delegation) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *Delegation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Delegation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Delegation) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### SetNetworkNil

`func (o *Delegation) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *Delegation) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetDescription

`func (o *Delegation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Delegation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Delegation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDelegationRequestedAt

`func (o *Delegation) GetDelegationRequestedAt() time.Time`

GetDelegationRequestedAt returns the DelegationRequestedAt field if non-nil, zero value otherwise.

### GetDelegationRequestedAtOk

`func (o *Delegation) GetDelegationRequestedAtOk() (*time.Time, bool)`

GetDelegationRequestedAtOk returns a tuple with the DelegationRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationRequestedAt

`func (o *Delegation) SetDelegationRequestedAt(v time.Time)`

SetDelegationRequestedAt sets DelegationRequestedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


