# GetTsigsCommonConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | 処理の問い合わせの際のキーになる文字列 | 
**Results** | [**[]CommonConfig**](CommonConfig.md) |  | [default to []]

## Methods

### NewGetTsigsCommonConfigs

`func NewGetTsigsCommonConfigs(requestId string, results []CommonConfig, ) *GetTsigsCommonConfigs`

NewGetTsigsCommonConfigs instantiates a new GetTsigsCommonConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTsigsCommonConfigsWithDefaults

`func NewGetTsigsCommonConfigsWithDefaults() *GetTsigsCommonConfigs`

NewGetTsigsCommonConfigsWithDefaults instantiates a new GetTsigsCommonConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetTsigsCommonConfigs) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetTsigsCommonConfigs) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetTsigsCommonConfigs) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResults

`func (o *GetTsigsCommonConfigs) GetResults() []CommonConfig`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetTsigsCommonConfigs) GetResultsOk() (*[]CommonConfig, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetTsigsCommonConfigs) SetResults(v []CommonConfig)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


