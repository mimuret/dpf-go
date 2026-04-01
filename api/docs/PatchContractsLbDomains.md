# PatchContractsLbDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonConfigId** | **int64** |  | 
**LbDomainIds** | **[]string** | GET lb_domains Schemaにおける id の配列 | 

## Methods

### NewPatchContractsLbDomains

`func NewPatchContractsLbDomains(commonConfigId int64, lbDomainIds []string, ) *PatchContractsLbDomains`

NewPatchContractsLbDomains instantiates a new PatchContractsLbDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchContractsLbDomainsWithDefaults

`func NewPatchContractsLbDomainsWithDefaults() *PatchContractsLbDomains`

NewPatchContractsLbDomainsWithDefaults instantiates a new PatchContractsLbDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonConfigId

`func (o *PatchContractsLbDomains) GetCommonConfigId() int64`

GetCommonConfigId returns the CommonConfigId field if non-nil, zero value otherwise.

### GetCommonConfigIdOk

`func (o *PatchContractsLbDomains) GetCommonConfigIdOk() (*int64, bool)`

GetCommonConfigIdOk returns a tuple with the CommonConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonConfigId

`func (o *PatchContractsLbDomains) SetCommonConfigId(v int64)`

SetCommonConfigId sets CommonConfigId field to given value.


### GetLbDomainIds

`func (o *PatchContractsLbDomains) GetLbDomainIds() []string`

GetLbDomainIds returns the LbDomainIds field if non-nil, zero value otherwise.

### GetLbDomainIdsOk

`func (o *PatchContractsLbDomains) GetLbDomainIdsOk() (*[]string, bool)`

GetLbDomainIdsOk returns a tuple with the LbDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbDomainIds

`func (o *PatchContractsLbDomains) SetLbDomainIds(v []string)`

SetLbDomainIds sets LbDomainIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


