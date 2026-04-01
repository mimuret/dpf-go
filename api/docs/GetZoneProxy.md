# GetZoneProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**ZoneProxy**](ZoneProxy.md) |  | 

## Methods

### NewGetZoneProxy

`func NewGetZoneProxy(requestId string, result ZoneProxy, ) *GetZoneProxy`

NewGetZoneProxy instantiates a new GetZoneProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneProxyWithDefaults

`func NewGetZoneProxyWithDefaults() *GetZoneProxy`

NewGetZoneProxyWithDefaults instantiates a new GetZoneProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetZoneProxy) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZoneProxy) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZoneProxy) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetZoneProxy) GetResult() ZoneProxy`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneProxy) GetResultOk() (*ZoneProxy, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneProxy) SetResult(v ZoneProxy)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


