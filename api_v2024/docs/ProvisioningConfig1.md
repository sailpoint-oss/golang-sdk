# ProvisioningConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniversalManager** | Pointer to **bool** | Specifies whether this configuration is used to manage provisioning requests for all sources from the org.  If true, no managedResourceRefs are allowed. | [optional] [readonly] [default to false]
**ManagedResourceRefs** | Pointer to [**[]ProvisioningConfig1ManagedResourceRefsInner**](ProvisioningConfig1ManagedResourceRefsInner.md) | References to sources for the Service Desk integration template.  May only be specified if universalManager is false. | [optional] 
**PlanInitializerScript** | Pointer to [**ProvisioningConfig1PlanInitializerScript**](ProvisioningConfig1PlanInitializerScript.md) |  | [optional] 
**NoProvisioningRequests** | Pointer to **bool** | Name of an attribute that when true disables the saving of ProvisioningRequest objects whenever plans are sent through this integration. | [optional] [default to false]
**ProvisioningRequestExpiration** | Pointer to **int32** | When saving pending requests is enabled, this defines the number of hours the request is allowed to live before it is considered expired and no longer affects plan compilation. | [optional] 

## Methods

### NewProvisioningConfig1

`func NewProvisioningConfig1() *ProvisioningConfig1`

NewProvisioningConfig1 instantiates a new ProvisioningConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConfig1WithDefaults

`func NewProvisioningConfig1WithDefaults() *ProvisioningConfig1`

NewProvisioningConfig1WithDefaults instantiates a new ProvisioningConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniversalManager

`func (o *ProvisioningConfig1) GetUniversalManager() bool`

GetUniversalManager returns the UniversalManager field if non-nil, zero value otherwise.

### GetUniversalManagerOk

`func (o *ProvisioningConfig1) GetUniversalManagerOk() (*bool, bool)`

GetUniversalManagerOk returns a tuple with the UniversalManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalManager

`func (o *ProvisioningConfig1) SetUniversalManager(v bool)`

SetUniversalManager sets UniversalManager field to given value.

### HasUniversalManager

`func (o *ProvisioningConfig1) HasUniversalManager() bool`

HasUniversalManager returns a boolean if a field has been set.

### GetManagedResourceRefs

`func (o *ProvisioningConfig1) GetManagedResourceRefs() []ProvisioningConfig1ManagedResourceRefsInner`

GetManagedResourceRefs returns the ManagedResourceRefs field if non-nil, zero value otherwise.

### GetManagedResourceRefsOk

`func (o *ProvisioningConfig1) GetManagedResourceRefsOk() (*[]ProvisioningConfig1ManagedResourceRefsInner, bool)`

GetManagedResourceRefsOk returns a tuple with the ManagedResourceRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceRefs

`func (o *ProvisioningConfig1) SetManagedResourceRefs(v []ProvisioningConfig1ManagedResourceRefsInner)`

SetManagedResourceRefs sets ManagedResourceRefs field to given value.

### HasManagedResourceRefs

`func (o *ProvisioningConfig1) HasManagedResourceRefs() bool`

HasManagedResourceRefs returns a boolean if a field has been set.

### GetPlanInitializerScript

`func (o *ProvisioningConfig1) GetPlanInitializerScript() ProvisioningConfig1PlanInitializerScript`

GetPlanInitializerScript returns the PlanInitializerScript field if non-nil, zero value otherwise.

### GetPlanInitializerScriptOk

`func (o *ProvisioningConfig1) GetPlanInitializerScriptOk() (*ProvisioningConfig1PlanInitializerScript, bool)`

GetPlanInitializerScriptOk returns a tuple with the PlanInitializerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInitializerScript

`func (o *ProvisioningConfig1) SetPlanInitializerScript(v ProvisioningConfig1PlanInitializerScript)`

SetPlanInitializerScript sets PlanInitializerScript field to given value.

### HasPlanInitializerScript

`func (o *ProvisioningConfig1) HasPlanInitializerScript() bool`

HasPlanInitializerScript returns a boolean if a field has been set.

### GetNoProvisioningRequests

`func (o *ProvisioningConfig1) GetNoProvisioningRequests() bool`

GetNoProvisioningRequests returns the NoProvisioningRequests field if non-nil, zero value otherwise.

### GetNoProvisioningRequestsOk

`func (o *ProvisioningConfig1) GetNoProvisioningRequestsOk() (*bool, bool)`

GetNoProvisioningRequestsOk returns a tuple with the NoProvisioningRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProvisioningRequests

`func (o *ProvisioningConfig1) SetNoProvisioningRequests(v bool)`

SetNoProvisioningRequests sets NoProvisioningRequests field to given value.

### HasNoProvisioningRequests

`func (o *ProvisioningConfig1) HasNoProvisioningRequests() bool`

HasNoProvisioningRequests returns a boolean if a field has been set.

### GetProvisioningRequestExpiration

`func (o *ProvisioningConfig1) GetProvisioningRequestExpiration() int32`

GetProvisioningRequestExpiration returns the ProvisioningRequestExpiration field if non-nil, zero value otherwise.

### GetProvisioningRequestExpirationOk

`func (o *ProvisioningConfig1) GetProvisioningRequestExpirationOk() (*int32, bool)`

GetProvisioningRequestExpirationOk returns a tuple with the ProvisioningRequestExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningRequestExpiration

`func (o *ProvisioningConfig1) SetProvisioningRequestExpiration(v int32)`

SetProvisioningRequestExpiration sets ProvisioningRequestExpiration field to given value.

### HasProvisioningRequestExpiration

`func (o *ProvisioningConfig1) HasProvisioningRequestExpiration() bool`

HasProvisioningRequestExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


