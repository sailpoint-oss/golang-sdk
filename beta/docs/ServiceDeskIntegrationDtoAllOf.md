# ServiceDeskIntegrationDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewServiceDeskIntegrationDtoAllOf

`func NewServiceDeskIntegrationDtoAllOf(description string, type_ string, attributes map[string]interface{}, ) *ServiceDeskIntegrationDtoAllOf`

NewServiceDeskIntegrationDtoAllOf instantiates a new ServiceDeskIntegrationDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeskIntegrationDtoAllOfWithDefaults

`func NewServiceDeskIntegrationDtoAllOfWithDefaults() *ServiceDeskIntegrationDtoAllOf`

NewServiceDeskIntegrationDtoAllOfWithDefaults instantiates a new ServiceDeskIntegrationDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceDeskIntegrationDtoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceDeskIntegrationDtoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ServiceDeskIntegrationDtoAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDeskIntegrationDtoAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetOwnerRef

`func (o *ServiceDeskIntegrationDtoAllOf) GetOwnerRef() BaseReferenceDto`

GetOwnerRef returns the OwnerRef field if non-nil, zero value otherwise.

### GetOwnerRefOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetOwnerRefOk() (*BaseReferenceDto, bool)`

GetOwnerRefOk returns a tuple with the OwnerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRef

`func (o *ServiceDeskIntegrationDtoAllOf) SetOwnerRef(v BaseReferenceDto)`

SetOwnerRef sets OwnerRef field to given value.

### HasOwnerRef

`func (o *ServiceDeskIntegrationDtoAllOf) HasOwnerRef() bool`

HasOwnerRef returns a boolean if a field has been set.

### GetClusterRef

`func (o *ServiceDeskIntegrationDtoAllOf) GetClusterRef() BaseReferenceDto`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetClusterRefOk() (*BaseReferenceDto, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *ServiceDeskIntegrationDtoAllOf) SetClusterRef(v BaseReferenceDto)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *ServiceDeskIntegrationDtoAllOf) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetCluster

`func (o *ServiceDeskIntegrationDtoAllOf) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ServiceDeskIntegrationDtoAllOf) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ServiceDeskIntegrationDtoAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetManagedSources

`func (o *ServiceDeskIntegrationDtoAllOf) GetManagedSources() []string`

GetManagedSources returns the ManagedSources field if non-nil, zero value otherwise.

### GetManagedSourcesOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetManagedSourcesOk() (*[]string, bool)`

GetManagedSourcesOk returns a tuple with the ManagedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSources

`func (o *ServiceDeskIntegrationDtoAllOf) SetManagedSources(v []string)`

SetManagedSources sets ManagedSources field to given value.

### HasManagedSources

`func (o *ServiceDeskIntegrationDtoAllOf) HasManagedSources() bool`

HasManagedSources returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *ServiceDeskIntegrationDtoAllOf) GetProvisioningConfig() ProvisioningConfig`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetProvisioningConfigOk() (*ProvisioningConfig, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ServiceDeskIntegrationDtoAllOf) SetProvisioningConfig(v ProvisioningConfig)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *ServiceDeskIntegrationDtoAllOf) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceDeskIntegrationDtoAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceDeskIntegrationDtoAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDtoAllOf) GetBeforeProvisioningRule() BaseReferenceDto`

GetBeforeProvisioningRule returns the BeforeProvisioningRule field if non-nil, zero value otherwise.

### GetBeforeProvisioningRuleOk

`func (o *ServiceDeskIntegrationDtoAllOf) GetBeforeProvisioningRuleOk() (*BaseReferenceDto, bool)`

GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDtoAllOf) SetBeforeProvisioningRule(v BaseReferenceDto)`

SetBeforeProvisioningRule sets BeforeProvisioningRule field to given value.

### HasBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDtoAllOf) HasBeforeProvisioningRule() bool`

HasBeforeProvisioningRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


