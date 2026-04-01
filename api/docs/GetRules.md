# GetRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]Rule**](Rule.md) |  | [default to []]

## Methods

### NewGetRules

`func NewGetRules(requestId string, results []Rule, ) *GetRules`

NewGetRules instantiates a new GetRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesWithDefaults

`func NewGetRulesWithDefaults() *GetRules`

NewGetRulesWithDefaults instantiates a new GetRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRules) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRules) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRules) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetRules) GetResults() []Rule`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetRules) GetResultsOk() (*[]Rule, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetRules) SetResults(v []Rule)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


