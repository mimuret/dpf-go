# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ServiceCode** | **string** |  | 
**State** | [**ContractState**](ContractState.md) |  | 
**Favorite** | [**ContractFavorite**](ContractFavorite.md) |  | 
**Plan** | [**ContractPlan**](ContractPlan.md) |  | 
**Description** | **string** | コメント | [default to ""]

## Methods

### NewContract

`func NewContract(id string, serviceCode string, state ContractState, favorite ContractFavorite, plan ContractPlan, description string, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v string)`

SetId sets Id field to given value.


### GetServiceCode

`func (o *Contract) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *Contract) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *Contract) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetState

`func (o *Contract) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Contract) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Contract) SetState(v ContractState)`

SetState sets State field to given value.


### GetFavorite

`func (o *Contract) GetFavorite() ContractFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *Contract) GetFavoriteOk() (*ContractFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *Contract) SetFavorite(v ContractFavorite)`

SetFavorite sets Favorite field to given value.


### GetPlan

`func (o *Contract) GetPlan() ContractPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Contract) GetPlanOk() (*ContractPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Contract) SetPlan(v ContractPlan)`

SetPlan sets Plan field to given value.


### GetDescription

`func (o *Contract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Contract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Contract) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


