# ServiceDeskIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **string** | Name of the Object | 
**Created** | Pointer to **time.Time** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Last modification date of the Object | [optional] [readonly] 
**Description** | **string** | Description of the Service Desk integration | 
**Type** | **string** | Service Desk integration types  - ServiceNowSDIM - ServiceNow  | [default to "ServiceNowSDIM"]
**OwnerRef** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) | Reference to the identity that is the owner of this Service Desk integration | [optional] 
**ClusterRef** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) | Reference to the source cluster for this Service Desk integration | [optional] 
**Cluster** | Pointer to **string** | ID of the cluster for the Service Desk integration (replaced by clusterRef, retained for backward compatibility) | [optional] 
**ManagedSources** | Pointer to **[]string** | Source IDs for the Service Desk integration (replaced by provisioningConfig.managedSResourceRefs, but retained here for backward compatibility) | [optional] 
**ProvisioningConfig** | Pointer to [**ProvisioningConfig**](ProvisioningConfig.md) |  | [optional] 
**Attributes** | **map[string]interface{}** | Attributes of the Service Desk integration.  Validation constraints enforced by the implementation. | 
**BeforeProvisioningRule** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) | Reference to beforeProvisioningRule for this Service Desk integration | [optional] 

## Methods

### NewServiceDeskIntegrationDto

`func NewServiceDeskIntegrationDto(name string, description string, type_ string, attributes map[string]interface{}, ) *ServiceDeskIntegrationDto`

NewServiceDeskIntegrationDto instantiates a new ServiceDeskIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeskIntegrationDtoWithDefaults

`func NewServiceDeskIntegrationDtoWithDefaults() *ServiceDeskIntegrationDto`

NewServiceDeskIntegrationDtoWithDefaults instantiates a new ServiceDeskIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceDeskIntegrationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceDeskIntegrationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceDeskIntegrationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceDeskIntegrationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceDeskIntegrationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeskIntegrationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeskIntegrationDto) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *ServiceDeskIntegrationDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceDeskIntegrationDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceDeskIntegrationDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceDeskIntegrationDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ServiceDeskIntegrationDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ServiceDeskIntegrationDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ServiceDeskIntegrationDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ServiceDeskIntegrationDto) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceDeskIntegrationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceDeskIntegrationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceDeskIntegrationDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ServiceDeskIntegrationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDeskIntegrationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDeskIntegrationDto) SetType(v string)`

SetType sets Type field to given value.


### GetOwnerRef

`func (o *ServiceDeskIntegrationDto) GetOwnerRef() BaseReferenceDto`

GetOwnerRef returns the OwnerRef field if non-nil, zero value otherwise.

### GetOwnerRefOk

`func (o *ServiceDeskIntegrationDto) GetOwnerRefOk() (*BaseReferenceDto, bool)`

GetOwnerRefOk returns a tuple with the OwnerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRef

`func (o *ServiceDeskIntegrationDto) SetOwnerRef(v BaseReferenceDto)`

SetOwnerRef sets OwnerRef field to given value.

### HasOwnerRef

`func (o *ServiceDeskIntegrationDto) HasOwnerRef() bool`

HasOwnerRef returns a boolean if a field has been set.

### GetClusterRef

`func (o *ServiceDeskIntegrationDto) GetClusterRef() BaseReferenceDto`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *ServiceDeskIntegrationDto) GetClusterRefOk() (*BaseReferenceDto, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *ServiceDeskIntegrationDto) SetClusterRef(v BaseReferenceDto)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *ServiceDeskIntegrationDto) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetCluster

`func (o *ServiceDeskIntegrationDto) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ServiceDeskIntegrationDto) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ServiceDeskIntegrationDto) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ServiceDeskIntegrationDto) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetManagedSources

`func (o *ServiceDeskIntegrationDto) GetManagedSources() []string`

GetManagedSources returns the ManagedSources field if non-nil, zero value otherwise.

### GetManagedSourcesOk

`func (o *ServiceDeskIntegrationDto) GetManagedSourcesOk() (*[]string, bool)`

GetManagedSourcesOk returns a tuple with the ManagedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSources

`func (o *ServiceDeskIntegrationDto) SetManagedSources(v []string)`

SetManagedSources sets ManagedSources field to given value.

### HasManagedSources

`func (o *ServiceDeskIntegrationDto) HasManagedSources() bool`

HasManagedSources returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *ServiceDeskIntegrationDto) GetProvisioningConfig() ProvisioningConfig`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ServiceDeskIntegrationDto) GetProvisioningConfigOk() (*ProvisioningConfig, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ServiceDeskIntegrationDto) SetProvisioningConfig(v ProvisioningConfig)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *ServiceDeskIntegrationDto) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceDeskIntegrationDto) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceDeskIntegrationDto) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceDeskIntegrationDto) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto) GetBeforeProvisioningRule() BaseReferenceDto`

GetBeforeProvisioningRule returns the BeforeProvisioningRule field if non-nil, zero value otherwise.

### GetBeforeProvisioningRuleOk

`func (o *ServiceDeskIntegrationDto) GetBeforeProvisioningRuleOk() (*BaseReferenceDto, bool)`

GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto) SetBeforeProvisioningRule(v BaseReferenceDto)`

SetBeforeProvisioningRule sets BeforeProvisioningRule field to given value.

### HasBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto) HasBeforeProvisioningRule() bool`

HasBeforeProvisioningRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


