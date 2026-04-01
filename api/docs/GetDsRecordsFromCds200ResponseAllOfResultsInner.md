# GetDsRecordsFromCds200ResponseAllOfResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name | 
**Ttl** | **NullableInt32** | TTL | 
**Rrtype** | [**DsRecordsRrtype**](DsRecordsRrtype.md) |  | 
**Rdata** | [**[]RecordsRdataInner**](RecordsRdataInner.md) | レコードの値 | 

## Methods

### NewGetDsRecordsFromCds200ResponseAllOfResultsInner

`func NewGetDsRecordsFromCds200ResponseAllOfResultsInner(name string, ttl NullableInt32, rrtype DsRecordsRrtype, rdata []RecordsRdataInner, ) *GetDsRecordsFromCds200ResponseAllOfResultsInner`

NewGetDsRecordsFromCds200ResponseAllOfResultsInner instantiates a new GetDsRecordsFromCds200ResponseAllOfResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDsRecordsFromCds200ResponseAllOfResultsInnerWithDefaults

`func NewGetDsRecordsFromCds200ResponseAllOfResultsInnerWithDefaults() *GetDsRecordsFromCds200ResponseAllOfResultsInner`

NewGetDsRecordsFromCds200ResponseAllOfResultsInnerWithDefaults instantiates a new GetDsRecordsFromCds200ResponseAllOfResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetRrtype

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetRrtype() DsRecordsRrtype`

GetRrtype returns the Rrtype field if non-nil, zero value otherwise.

### GetRrtypeOk

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetRrtypeOk() (*DsRecordsRrtype, bool)`

GetRrtypeOk returns a tuple with the Rrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrtype

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) SetRrtype(v DsRecordsRrtype)`

SetRrtype sets Rrtype field to given value.


### GetRdata

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetRdata() []RecordsRdataInner`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) GetRdataOk() (*[]RecordsRdataInner, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *GetDsRecordsFromCds200ResponseAllOfResultsInner) SetRdata(v []RecordsRdataInner)`

SetRdata sets Rdata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


