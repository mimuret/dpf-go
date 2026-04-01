# DsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rrset** | **string** |  | 
**TransitedAt** | **time.Time** |  | 

## Methods

### NewDsRecord

`func NewDsRecord(rrset string, transitedAt time.Time, ) *DsRecord`

NewDsRecord instantiates a new DsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsRecordWithDefaults

`func NewDsRecordWithDefaults() *DsRecord`

NewDsRecordWithDefaults instantiates a new DsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRrset

`func (o *DsRecord) GetRrset() string`

GetRrset returns the Rrset field if non-nil, zero value otherwise.

### GetRrsetOk

`func (o *DsRecord) GetRrsetOk() (*string, bool)`

GetRrsetOk returns a tuple with the Rrset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrset

`func (o *DsRecord) SetRrset(v string)`

SetRrset sets Rrset field to given value.


### GetTransitedAt

`func (o *DsRecord) GetTransitedAt() time.Time`

GetTransitedAt returns the TransitedAt field if non-nil, zero value otherwise.

### GetTransitedAtOk

`func (o *DsRecord) GetTransitedAtOk() (*time.Time, bool)`

GetTransitedAtOk returns a tuple with the TransitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitedAt

`func (o *DsRecord) SetTransitedAt(v time.Time)`

SetTransitedAt sets TransitedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


