# Tsig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Algorithm** | [**TsigsAlgorithm**](TsigsAlgorithm.md) |  | 
**Secret** | **string** |  | 
**Description** | **string** | コメント | [default to ""]

## Methods

### NewTsig

`func NewTsig(id int64, name string, algorithm TsigsAlgorithm, secret string, description string, ) *Tsig`

NewTsig instantiates a new Tsig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTsigWithDefaults

`func NewTsigWithDefaults() *Tsig`

NewTsigWithDefaults instantiates a new Tsig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tsig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tsig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tsig) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Tsig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tsig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tsig) SetName(v string)`

SetName sets Name field to given value.


### GetAlgorithm

`func (o *Tsig) GetAlgorithm() TsigsAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Tsig) GetAlgorithmOk() (*TsigsAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Tsig) SetAlgorithm(v TsigsAlgorithm)`

SetAlgorithm sets Algorithm field to given value.


### GetSecret

`func (o *Tsig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Tsig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Tsig) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetDescription

`func (o *Tsig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tsig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tsig) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


