# NotifiedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Address** | **string** | IPアドレス | 
**TsigId** | **NullableInt64** |  | 

## Methods

### NewNotifiedServer

`func NewNotifiedServer(id int64, address string, tsigId NullableInt64, ) *NotifiedServer`

NewNotifiedServer instantiates a new NotifiedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifiedServerWithDefaults

`func NewNotifiedServerWithDefaults() *NotifiedServer`

NewNotifiedServerWithDefaults instantiates a new NotifiedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotifiedServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotifiedServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotifiedServer) SetId(v int64)`

SetId sets Id field to given value.


### GetAddress

`func (o *NotifiedServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NotifiedServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NotifiedServer) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTsigId

`func (o *NotifiedServer) GetTsigId() int64`

GetTsigId returns the TsigId field if non-nil, zero value otherwise.

### GetTsigIdOk

`func (o *NotifiedServer) GetTsigIdOk() (*int64, bool)`

GetTsigIdOk returns a tuple with the TsigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigId

`func (o *NotifiedServer) SetTsigId(v int64)`

SetTsigId sets TsigId field to given value.


### SetTsigIdNil

`func (o *NotifiedServer) SetTsigIdNil(b bool)`

 SetTsigIdNil sets the value for TsigId to be an explicit nil

### UnsetTsigId
`func (o *NotifiedServer) UnsetTsigId()`

UnsetTsigId ensures that no value is present for TsigId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


