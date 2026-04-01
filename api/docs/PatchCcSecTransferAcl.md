# PatchCcSecTransferAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsigId** | Pointer to **NullableInt64** |  | [optional] 
**Network** | Pointer to **string** | IPアドレス/プレフィックス長 | [optional] 

## Methods

### NewPatchCcSecTransferAcl

`func NewPatchCcSecTransferAcl() *PatchCcSecTransferAcl`

NewPatchCcSecTransferAcl instantiates a new PatchCcSecTransferAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCcSecTransferAclWithDefaults

`func NewPatchCcSecTransferAclWithDefaults() *PatchCcSecTransferAcl`

NewPatchCcSecTransferAclWithDefaults instantiates a new PatchCcSecTransferAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsigId

`func (o *PatchCcSecTransferAcl) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *PatchCcSecTransferAcl) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *PatchCcSecTransferAcl) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.

### HasTsigId

`func (o *PatchCcSecTransferAcl) HasTsigId() bool`

HasTsigId returns a boolean if a field has been set.

### SetTsigIdNil

`func (o *PatchCcSecTransferAcl) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *PatchCcSecTransferAcl) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil
### GetNetwork

`func (o *PatchCcSecTransferAcl) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PatchCcSecTransferAcl) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PatchCcSecTransferAcl) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PatchCcSecTransferAcl) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


