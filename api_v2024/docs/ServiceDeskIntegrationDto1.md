# ServiceDeskIntegrationDto1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service Desk integration&#39;s name. The name must be unique. | 
**Description** | **string** | Service Desk integration&#39;s description. | 
**Type** | **string** | Service Desk integration types:  - ServiceNowSDIM - ServiceNow  | [default to "ServiceNowSDIM"]
**OwnerRef** | Pointer to [**OwnerDto**](OwnerDto.md) |  | [optional] 
**ClusterRef** | Pointer to [**SourceClusterDto**](SourceClusterDto.md) |  | [optional] 
**Cluster** | Pointer to **string** | Cluster ID for the Service Desk integration (replaced by clusterRef, retained for backward compatibility). | [optional] 
**ManagedSources** | Pointer to **[]string** | Source IDs for the Service Desk integration (replaced by provisioningConfig.managedSResourceRefs, but retained here for backward compatibility). | [optional] 
**ProvisioningConfig** | Pointer to [**ProvisioningConfig1**](ProvisioningConfig1.md) |  | [optional] 
**Attributes** | **map[string]interface{}** | Service Desk integration&#39;s attributes. Validation constraints enforced by the implementation. | 
**BeforeProvisioningRule** | Pointer to [**BeforeProvisioningRuleDto**](BeforeProvisioningRuleDto.md) |  | [optional] 

## Methods

### NewServiceDeskIntegrationDto1

`func NewServiceDeskIntegrationDto1(name string, description string, type_ string, attributes map[string]interface{}, ) *ServiceDeskIntegrationDto1`

NewServiceDeskIntegrationDto1 instantiates a new ServiceDeskIntegrationDto1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeskIntegrationDto1WithDefaults

`func NewServiceDeskIntegrationDto1WithDefaults() *ServiceDeskIntegrationDto1`

NewServiceDeskIntegrationDto1WithDefaults instantiates a new ServiceDeskIntegrationDto1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceDeskIntegrationDto1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeskIntegrationDto1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeskIntegrationDto1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceDeskIntegrationDto1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceDeskIntegrationDto1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceDeskIntegrationDto1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ServiceDeskIntegrationDto1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDeskIntegrationDto1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDeskIntegrationDto1) SetType(v string)`

SetType sets Type field to given value.


### GetOwnerRef

`func (o *ServiceDeskIntegrationDto1) GetOwnerRef() OwnerDto`

GetOwnerRef returns the OwnerRef field if non-nil, zero value otherwise.

### GetOwnerRefOk

`func (o *ServiceDeskIntegrationDto1) GetOwnerRefOk() (*OwnerDto, bool)`

GetOwnerRefOk returns a tuple with the OwnerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRef

`func (o *ServiceDeskIntegrationDto1) SetOwnerRef(v OwnerDto)`

SetOwnerRef sets OwnerRef field to given value.

### HasOwnerRef

`func (o *ServiceDeskIntegrationDto1) HasOwnerRef() bool`

HasOwnerRef returns a boolean if a field has been set.

### GetClusterRef

`func (o *ServiceDeskIntegrationDto1) GetClusterRef() SourceClusterDto`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *ServiceDeskIntegrationDto1) GetClusterRefOk() (*SourceClusterDto, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *ServiceDeskIntegrationDto1) SetClusterRef(v SourceClusterDto)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *ServiceDeskIntegrationDto1) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetCluster

`func (o *ServiceDeskIntegrationDto1) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ServiceDeskIntegrationDto1) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ServiceDeskIntegrationDto1) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ServiceDeskIntegrationDto1) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetManagedSources

`func (o *ServiceDeskIntegrationDto1) GetManagedSources() []string`

GetManagedSources returns the ManagedSources field if non-nil, zero value otherwise.

### GetManagedSourcesOk

`func (o *ServiceDeskIntegrationDto1) GetManagedSourcesOk() (*[]string, bool)`

GetManagedSourcesOk returns a tuple with the ManagedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSources

`func (o *ServiceDeskIntegrationDto1) SetManagedSources(v []string)`

SetManagedSources sets ManagedSources field to given value.

### HasManagedSources

`func (o *ServiceDeskIntegrationDto1) HasManagedSources() bool`

HasManagedSources returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *ServiceDeskIntegrationDto1) GetProvisioningConfig() ProvisioningConfig1`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ServiceDeskIntegrationDto1) GetProvisioningConfigOk() (*ProvisioningConfig1, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ServiceDeskIntegrationDto1) SetProvisioningConfig(v ProvisioningConfig1)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *ServiceDeskIntegrationDto1) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceDeskIntegrationDto1) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceDeskIntegrationDto1) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceDeskIntegrationDto1) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto1) GetBeforeProvisioningRule() BeforeProvisioningRuleDto`

GetBeforeProvisioningRule returns the BeforeProvisioningRule field if non-nil, zero value otherwise.

### GetBeforeProvisioningRuleOk

`func (o *ServiceDeskIntegrationDto1) GetBeforeProvisioningRuleOk() (*BeforeProvisioningRuleDto, bool)`

GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto1) SetBeforeProvisioningRule(v BeforeProvisioningRuleDto)`

SetBeforeProvisioningRule sets BeforeProvisioningRule field to given value.

### HasBeforeProvisioningRule

`func (o *ServiceDeskIntegrationDto1) HasBeforeProvisioningRule() bool`

HasBeforeProvisioningRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


