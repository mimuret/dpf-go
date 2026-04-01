# GetZoneHistoriesText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Result** | [**ZoneHistoryText**](ZoneHistoryText.md) |  | 

## Methods

### NewGetZoneHistoriesText

`func NewGetZoneHistoriesText(requestId string, result ZoneHistoryText, ) *GetZoneHistoriesText`

NewGetZoneHistoriesText instantiates a new GetZoneHistoriesText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneHistoriesTextWithDefaults

`func NewGetZoneHistoriesTextWithDefaults() *GetZoneHistoriesText`

NewGetZoneHistoriesTextWithDefaults instantiates a new GetZoneHistoriesText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetZoneHistoriesText) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZoneHistoriesText) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZoneHistoriesText) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResult

`func (o *GetZoneHistoriesText) GetResult() ZoneHistoryText`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneHistoriesText) GetResultOk() (*ZoneHistoryText, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneHistoriesText) SetResult(v ZoneHistoryText)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


