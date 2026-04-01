# CcSecTransferAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Network** | **string** | IPアドレス/プレフィックス長 | 
**TsigId** | **NullableInt64** |  | 

## Methods

### NewCcSecTransferAcl

`func NewCcSecTransferAcl(id int64, network string, tsigId NullableInt64, ) *CcSecTransferAcl`

NewCcSecTransferAcl instantiates a new CcSecTransferAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcSecTransferAclWithDefaults

`func NewCcSecTransferAclWithDefaults() *CcSecTransferAcl`

NewCcSecTransferAclWithDefaults instantiates a new CcSecTransferAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CcSecTransferAcl) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcSecTransferAcl) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcSecTransferAcl) SetId(v int64)`

SetId sets Id field to given value.


### GetNetwork

`func (o *CcSecTransferAcl) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CcSecTransferAcl) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CcSecTransferAcl) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTsigId

`func (o *CcSecTransferAcl) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *CcSecTransferAcl) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *CcSecTransferAcl) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.


### SetTsigIdNil

`func (o *CcSecTransferAcl) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *CcSecTransferAcl) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


