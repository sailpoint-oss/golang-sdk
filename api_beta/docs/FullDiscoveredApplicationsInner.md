# FullDiscoveredApplicationsInner

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
**AssociatedSources** | Pointer to **[]string** | List of associated sources related to this discovered application. | [optional] 

## Methods

### NewFullDiscoveredApplicationsInner

`func NewFullDiscoveredApplicationsInner() *FullDiscoveredApplicationsInner`

NewFullDiscoveredApplicationsInner instantiates a new FullDiscoveredApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDiscoveredApplicationsInnerWithDefaults

`func NewFullDiscoveredApplicationsInnerWithDefaults() *FullDiscoveredApplicationsInner`

NewFullDiscoveredApplicationsInnerWithDefaults instantiates a new FullDiscoveredApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullDiscoveredApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullDiscoveredApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullDiscoveredApplicationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullDiscoveredApplicationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullDiscoveredApplicationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullDiscoveredApplicationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullDiscoveredApplicationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullDiscoveredApplicationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDiscoverySource

`func (o *FullDiscoveredApplicationsInner) GetDiscoverySource() string`

GetDiscoverySource returns the DiscoverySource field if non-nil, zero value otherwise.

### GetDiscoverySourceOk

`func (o *FullDiscoveredApplicationsInner) GetDiscoverySourceOk() (*string, bool)`

GetDiscoverySourceOk returns a tuple with the DiscoverySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySource

`func (o *FullDiscoveredApplicationsInner) SetDiscoverySource(v string)`

SetDiscoverySource sets DiscoverySource field to given value.

### HasDiscoverySource

`func (o *FullDiscoveredApplicationsInner) HasDiscoverySource() bool`

HasDiscoverySource returns a boolean if a field has been set.

### GetDiscoveredVendor

`func (o *FullDiscoveredApplicationsInner) GetDiscoveredVendor() string`

GetDiscoveredVendor returns the DiscoveredVendor field if non-nil, zero value otherwise.

### GetDiscoveredVendorOk

`func (o *FullDiscoveredApplicationsInner) GetDiscoveredVendorOk() (*string, bool)`

GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVendor

`func (o *FullDiscoveredApplicationsInner) SetDiscoveredVendor(v string)`

SetDiscoveredVendor sets DiscoveredVendor field to given value.

### HasDiscoveredVendor

`func (o *FullDiscoveredApplicationsInner) HasDiscoveredVendor() bool`

HasDiscoveredVendor returns a boolean if a field has been set.

### GetDescription

`func (o *FullDiscoveredApplicationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullDiscoveredApplicationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullDiscoveredApplicationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullDiscoveredApplicationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecommendedConnectors

`func (o *FullDiscoveredApplicationsInner) GetRecommendedConnectors() []string`

GetRecommendedConnectors returns the RecommendedConnectors field if non-nil, zero value otherwise.

### GetRecommendedConnectorsOk

`func (o *FullDiscoveredApplicationsInner) GetRecommendedConnectorsOk() (*[]string, bool)`

GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedConnectors

`func (o *FullDiscoveredApplicationsInner) SetRecommendedConnectors(v []string)`

SetRecommendedConnectors sets RecommendedConnectors field to given value.

### HasRecommendedConnectors

`func (o *FullDiscoveredApplicationsInner) HasRecommendedConnectors() bool`

HasRecommendedConnectors returns a boolean if a field has been set.

### GetDiscoveredAt

`func (o *FullDiscoveredApplicationsInner) GetDiscoveredAt() time.Time`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *FullDiscoveredApplicationsInner) GetDiscoveredAtOk() (*time.Time, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *FullDiscoveredApplicationsInner) SetDiscoveredAt(v time.Time)`

SetDiscoveredAt sets DiscoveredAt field to given value.

### HasDiscoveredAt

`func (o *FullDiscoveredApplicationsInner) HasDiscoveredAt() bool`

HasDiscoveredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FullDiscoveredApplicationsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullDiscoveredApplicationsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullDiscoveredApplicationsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullDiscoveredApplicationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *FullDiscoveredApplicationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FullDiscoveredApplicationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FullDiscoveredApplicationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FullDiscoveredApplicationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssociatedSources

`func (o *FullDiscoveredApplicationsInner) GetAssociatedSources() []string`

GetAssociatedSources returns the AssociatedSources field if non-nil, zero value otherwise.

### GetAssociatedSourcesOk

`func (o *FullDiscoveredApplicationsInner) GetAssociatedSourcesOk() (*[]string, bool)`

GetAssociatedSourcesOk returns a tuple with the AssociatedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedSources

`func (o *FullDiscoveredApplicationsInner) SetAssociatedSources(v []string)`

SetAssociatedSources sets AssociatedSources field to given value.

### HasAssociatedSources

`func (o *FullDiscoveredApplicationsInner) HasAssociatedSources() bool`

HasAssociatedSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


