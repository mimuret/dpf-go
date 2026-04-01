# SitesSitesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | 登録可能な文字列は[**こちら**](https://manual.iij.jp/dpf/help/19629152.html#DNS%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E7%99%BB%E9%8C%B2%E3%83%AB%E3%83%BC%E3%83%AB-%E3%83%9B%E3%82%B9%E3%83%88%E5%90%8D%E3%81%AE%E5%85%B1%E9%80%9A%E3%83%AB%E3%83%BC%E3%83%AB)のホスト名の共通ルールを参照してください。 | 
**Name** | **string** | サイト名 | 
**Rrtype** | [**SiteRrtype**](SiteRrtype.md) |  | 
**Description** | **string** | コメント | [default to ""]
**LiveStatus** | [**SiteLiveStatus**](SiteLiveStatus.md) |  | 
**Endpoints** | [**[]Endpoint**](Endpoint.md) |  | [default to []]

## Methods

### NewSitesSitesInner

`func NewSitesSitesInner(resourceName string, name string, rrtype SiteRrtype, description string, liveStatus SiteLiveStatus, endpoints []Endpoint, ) *SitesSitesInner`

NewSitesSitesInner instantiates a new SitesSitesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitesSitesInnerWithDefaults

`func NewSitesSitesInnerWithDefaults() *SitesSitesInner`

NewSitesSitesInnerWithDefaults instantiates a new SitesSitesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *SitesSitesInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SitesSitesInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SitesSitesInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetName

`func (o *SitesSitesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SitesSitesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SitesSitesInner) SetName(v string)`

SetName sets Name field to given value.


### GetRrtype

`func (o *SitesSitesInner) GetRrtype() SiteRrtype`

GetRrtype returns the Rrtype field if non-nil, zero value otherwise.

### GetRrtypeOk

`func (o *SitesSitesInner) GetRrtypeOk() (*SiteRrtype, bool)`

GetRrtypeOk returns a tuple with the Rrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrtype

`func (o *SitesSitesInner) SetRrtype(v SiteRrtype)`

SetRrtype sets Rrtype field to given value.


### GetDescription

`func (o *SitesSitesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SitesSitesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SitesSitesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLiveStatus

`func (o *SitesSitesInner) GetLiveStatus() SiteLiveStatus`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *SitesSitesInner) GetLiveStatusOk() (*SiteLiveStatus, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStatus

`func (o *SitesSitesInner) SetLiveStatus(v SiteLiveStatus)`

SetLiveStatus sets LiveStatus field to given value.


### GetEndpoints

`func (o *SitesSitesInner) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *SitesSitesInner) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *SitesSitesInner) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


