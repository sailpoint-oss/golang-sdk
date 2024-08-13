# SlimDiscoveredApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the discovered application. | [optional] 
**Name** | Pointer to **string** | Name of the discovered application. | [optional] 
**DiscoverySource** | Pointer to **string** | Source from which the application was discovered. | [optional] 
**DiscoveredVendor** | Pointer to **string** | The vendor associated with the discovered application. | [optional] 
**Description** | Pointer to **string** | A brief description of the discovered application. | [optional] 
**RecommendedConnectors** | Pointer to **[]string** | List of recommended connectors for the application. | [optional] 
**DiscoveredAt** | Pointer to **time.Time** | The timestamp when the application was last received via an entitlement aggregation invocation  or a manual csv upload, in ISO 8601 format. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the application was first discovered, in ISO 8601 format. | [optional] 
**Status** | Pointer to **string** | The status of an application within the discovery source.  By default this field is set to \&quot;ACTIVE\&quot; when the application is discovered.  If an application has been deleted from within the discovery source, the status will be set to \&quot;INACTIVE\&quot;. | [optional] 

## Methods

### NewSlimDiscoveredApplicationsInner

`func NewSlimDiscoveredApplicationsInner() *SlimDiscoveredApplicationsInner`

NewSlimDiscoveredApplicationsInner instantiates a new SlimDiscoveredApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimDiscoveredApplicationsInnerWithDefaults

`func NewSlimDiscoveredApplicationsInnerWithDefaults() *SlimDiscoveredApplicationsInner`

NewSlimDiscoveredApplicationsInnerWithDefaults instantiates a new SlimDiscoveredApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimDiscoveredApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimDiscoveredApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimDiscoveredApplicationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlimDiscoveredApplicationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SlimDiscoveredApplicationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlimDiscoveredApplicationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlimDiscoveredApplicationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlimDiscoveredApplicationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDiscoverySource

`func (o *SlimDiscoveredApplicationsInner) GetDiscoverySource() string`

GetDiscoverySource returns the DiscoverySource field if non-nil, zero value otherwise.

### GetDiscoverySourceOk

`func (o *SlimDiscoveredApplicationsInner) GetDiscoverySourceOk() (*string, bool)`

GetDiscoverySourceOk returns a tuple with the DiscoverySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySource

`func (o *SlimDiscoveredApplicationsInner) SetDiscoverySource(v string)`

SetDiscoverySource sets DiscoverySource field to given value.

### HasDiscoverySource

`func (o *SlimDiscoveredApplicationsInner) HasDiscoverySource() bool`

HasDiscoverySource returns a boolean if a field has been set.

### GetDiscoveredVendor

`func (o *SlimDiscoveredApplicationsInner) GetDiscoveredVendor() string`

GetDiscoveredVendor returns the DiscoveredVendor field if non-nil, zero value otherwise.

### GetDiscoveredVendorOk

`func (o *SlimDiscoveredApplicationsInner) GetDiscoveredVendorOk() (*string, bool)`

GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVendor

`func (o *SlimDiscoveredApplicationsInner) SetDiscoveredVendor(v string)`

SetDiscoveredVendor sets DiscoveredVendor field to given value.

### HasDiscoveredVendor

`func (o *SlimDiscoveredApplicationsInner) HasDiscoveredVendor() bool`

HasDiscoveredVendor returns a boolean if a field has been set.

### GetDescription

`func (o *SlimDiscoveredApplicationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlimDiscoveredApplicationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlimDiscoveredApplicationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlimDiscoveredApplicationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecommendedConnectors

`func (o *SlimDiscoveredApplicationsInner) GetRecommendedConnectors() []string`

GetRecommendedConnectors returns the RecommendedConnectors field if non-nil, zero value otherwise.

### GetRecommendedConnectorsOk

`func (o *SlimDiscoveredApplicationsInner) GetRecommendedConnectorsOk() (*[]string, bool)`

GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedConnectors

`func (o *SlimDiscoveredApplicationsInner) SetRecommendedConnectors(v []string)`

SetRecommendedConnectors sets RecommendedConnectors field to given value.

### HasRecommendedConnectors

`func (o *SlimDiscoveredApplicationsInner) HasRecommendedConnectors() bool`

HasRecommendedConnectors returns a boolean if a field has been set.

### GetDiscoveredAt

`func (o *SlimDiscoveredApplicationsInner) GetDiscoveredAt() time.Time`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *SlimDiscoveredApplicationsInner) GetDiscoveredAtOk() (*time.Time, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *SlimDiscoveredApplicationsInner) SetDiscoveredAt(v time.Time)`

SetDiscoveredAt sets DiscoveredAt field to given value.

### HasDiscoveredAt

`func (o *SlimDiscoveredApplicationsInner) HasDiscoveredAt() bool`

HasDiscoveredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SlimDiscoveredApplicationsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlimDiscoveredApplicationsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlimDiscoveredApplicationsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SlimDiscoveredApplicationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SlimDiscoveredApplicationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlimDiscoveredApplicationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlimDiscoveredApplicationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlimDiscoveredApplicationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


