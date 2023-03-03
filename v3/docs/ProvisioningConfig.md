# ProvisioningConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniversalManager** | Pointer to **bool** | Specifies whether this configuration is used to manage provisioning requests for all sources from the org.  If true, no managedResourceRefs are allowed. | [optional] [readonly] 
**ManagedResourceRefs** | Pointer to [**[]ProvisioningConfigManagedResourceRefsInner**](ProvisioningConfigManagedResourceRefsInner.md) | References to sources for the Service Desk integration template.  May only be specified if universalManager is false. | [optional] 
**PlanInitializerScript** | Pointer to [**ProvisioningConfigPlanInitializerScript**](ProvisioningConfigPlanInitializerScript.md) |  | [optional] 

## Methods

### NewProvisioningConfig

`func NewProvisioningConfig() *ProvisioningConfig`

NewProvisioningConfig instantiates a new ProvisioningConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConfigWithDefaults

`func NewProvisioningConfigWithDefaults() *ProvisioningConfig`

NewProvisioningConfigWithDefaults instantiates a new ProvisioningConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniversalManager

`func (o *ProvisioningConfig) GetUniversalManager() bool`

GetUniversalManager returns the UniversalManager field if non-nil, zero value otherwise.

### GetUniversalManagerOk

`func (o *ProvisioningConfig) GetUniversalManagerOk() (*bool, bool)`

GetUniversalManagerOk returns a tuple with the UniversalManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalManager

`func (o *ProvisioningConfig) SetUniversalManager(v bool)`

SetUniversalManager sets UniversalManager field to given value.

### HasUniversalManager

`func (o *ProvisioningConfig) HasUniversalManager() bool`

HasUniversalManager returns a boolean if a field has been set.

### GetManagedResourceRefs

`func (o *ProvisioningConfig) GetManagedResourceRefs() []ProvisioningConfigManagedResourceRefsInner`

GetManagedResourceRefs returns the ManagedResourceRefs field if non-nil, zero value otherwise.

### GetManagedResourceRefsOk

`func (o *ProvisioningConfig) GetManagedResourceRefsOk() (*[]ProvisioningConfigManagedResourceRefsInner, bool)`

GetManagedResourceRefsOk returns a tuple with the ManagedResourceRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceRefs

`func (o *ProvisioningConfig) SetManagedResourceRefs(v []ProvisioningConfigManagedResourceRefsInner)`

SetManagedResourceRefs sets ManagedResourceRefs field to given value.

### HasManagedResourceRefs

`func (o *ProvisioningConfig) HasManagedResourceRefs() bool`

HasManagedResourceRefs returns a boolean if a field has been set.

### GetPlanInitializerScript

`func (o *ProvisioningConfig) GetPlanInitializerScript() ProvisioningConfigPlanInitializerScript`

GetPlanInitializerScript returns the PlanInitializerScript field if non-nil, zero value otherwise.

### GetPlanInitializerScriptOk

`func (o *ProvisioningConfig) GetPlanInitializerScriptOk() (*ProvisioningConfigPlanInitializerScript, bool)`

GetPlanInitializerScriptOk returns a tuple with the PlanInitializerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInitializerScript

`func (o *ProvisioningConfig) SetPlanInitializerScript(v ProvisioningConfigPlanInitializerScript)`

SetPlanInitializerScript sets PlanInitializerScript field to given value.

### HasPlanInitializerScript

`func (o *ProvisioningConfig) HasPlanInitializerScript() bool`

HasPlanInitializerScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


