# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CommonConfigId** | **int64** |  | 
**ServiceCode** | **string** |  | 
**State** | [**ZonesState**](ZonesState.md) |  | 
**Favorite** | [**ZonesFavorite**](ZonesFavorite.md) |  | 
**Name** | **string** | name | 
**Network** | **NullableString** |  | 
**Description** | **string** | コメント | [default to ""]

## Methods

### NewZone

`func NewZone(id string, commonConfigId int64, serviceCode string, state ZonesState, favorite ZonesFavorite, name string, network NullableString, description string, ) *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v string)`

SetId sets Id field to given value.


### GetCommonConfigId

`func (o *Zone) GetCommonConfigId() int64`

GetCommonConfigId returns the CommonConfigId field if non-nil, zero value otherwise.

### GetCommonConfigIdOk

`func (o *Zone) GetCommonConfigIdOk() (*int64, bool)`

GetCommonConfigIdOk returns a tuple with the CommonConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonConfigId

`func (o *Zone) SetCommonConfigId(v int64)`

SetCommonConfigId sets CommonConfigId field to given value.


### GetServiceCode

`func (o *Zone) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *Zone) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *Zone) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetState

`func (o *Zone) GetState() ZonesState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Zone) GetStateOk() (*ZonesState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Zone) SetState(v ZonesState)`

SetState sets State field to given value.


### GetFavorite

`func (o *Zone) GetFavorite() ZonesFavorite`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *Zone) GetFavoriteOk() (*ZonesFavorite, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *Zone) SetFavorite(v ZonesFavorite)`

SetFavorite sets Favorite field to given value.


### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *Zone) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Zone) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Zone) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### SetNetworkNil

`func (o *Zone) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *Zone) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetDescription

`func (o *Zone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Zone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Zone) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


