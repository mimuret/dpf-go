# PostCcSecTransferAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsigId** | Pointer to **NullableInt64** |  | [optional] 
**Network** | **string** | IPアドレス/プレフィックス長 | 

## Methods

### NewPostCcSecTransferAcl

`func NewPostCcSecTransferAcl(network string, ) *PostCcSecTransferAcl`

NewPostCcSecTransferAcl instantiates a new PostCcSecTransferAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCcSecTransferAclWithDefaults

`func NewPostCcSecTransferAclWithDefaults() *PostCcSecTransferAcl`

NewPostCcSecTransferAclWithDefaults instantiates a new PostCcSecTransferAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsigId

`func (o *PostCcSecTransferAcl) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *PostCcSecTransferAcl) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *PostCcSecTransferAcl) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.

### HasTsigId

`func (o *PostCcSecTransferAcl) HasTsigId() bool`

HasTsigId returns a boolean if a field has been set.

### SetTsigIdNil

`func (o *PostCcSecTransferAcl) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *PostCcSecTransferAcl) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil
### GetNetwork

`func (o *PostCcSecTransferAcl) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PostCcSecTransferAcl) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PostCcSecTransferAcl) SetNetwork(v string)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


