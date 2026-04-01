# GetZoneLabels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**ZoneLabels**](ZoneLabels.md) |  | 

## Methods

### NewGetZoneLabels200Response

`func NewGetZoneLabels200Response(requestId string, result ZoneLabels, ) *GetZoneLabels200Response`

NewGetZoneLabels200Response instantiates a new GetZoneLabels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneLabels200ResponseWithDefaults

`func NewGetZoneLabels200ResponseWithDefaults() *GetZoneLabels200Response`

NewGetZoneLabels200ResponseWithDefaults instantiates a new GetZoneLabels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetZoneLabels200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZoneLabels200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZoneLabels200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetZoneLabels200Response) GetResult() ZoneLabels`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneLabels200Response) GetResultOk() (*ZoneLabels, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneLabels200Response) SetResult(v ZoneLabels)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


