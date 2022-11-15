# RoleMiningPotentialRoleSummaryDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityDistribution** | Pointer to [**[]RoleMiningIdentityDistribution**](RoleMiningIdentityDistribution.md) | Identity attribute distribution | [optional] 
**PotentialRoleRef** | Pointer to [**RoleMiningPotentialRoleRef**](RoleMiningPotentialRoleRef.md) |  | [optional] 
**IdentityCount** | Pointer to **int32** | The number of identities in a potential role. | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements in a potential role. | [optional] 
**IdentityGroupStatus** | Pointer to **string** | The status for this identity group which can be \&quot;REQUESTED\&quot; or \&quot;OBTAINED\&quot; | [optional] 
**ProvisionState** | Pointer to [**RoleMiningPotentialRoleProvisionState**](RoleMiningPotentialRoleProvisionState.md) |  | [optional] 
**RoleId** | Pointer to **string** | ID of the provisioned role in IIQ or IDN.  Null if this potential role has not been provisioned. | [optional] 
**Density** | Pointer to **int32** | The density metric (0-100) of this potential role. Higher density values indicate higher similarity amongst the identities. | [optional] 
**Freshness** | Pointer to **int32** | The freshness metric (0-100) of this potential role. Higher freshness values indicate this potential role is more distinctive compared to existing roles. | [optional] 
**Quality** | Pointer to **int32** | The quality metric (0-100) of this potential role. Higher quality values indicate this potential role has high density and freshness. | [optional] 

## Methods

### NewRoleMiningPotentialRoleSummaryDistribution

`func NewRoleMiningPotentialRoleSummaryDistribution() *RoleMiningPotentialRoleSummaryDistribution`

NewRoleMiningPotentialRoleSummaryDistribution instantiates a new RoleMiningPotentialRoleSummaryDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningPotentialRoleSummaryDistributionWithDefaults

`func NewRoleMiningPotentialRoleSummaryDistributionWithDefaults() *RoleMiningPotentialRoleSummaryDistribution`

NewRoleMiningPotentialRoleSummaryDistributionWithDefaults instantiates a new RoleMiningPotentialRoleSummaryDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityDistribution

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityDistribution() []RoleMiningIdentityDistribution`

GetIdentityDistribution returns the IdentityDistribution field if non-nil, zero value otherwise.

### GetIdentityDistributionOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityDistributionOk() (*[]RoleMiningIdentityDistribution, bool)`

GetIdentityDistributionOk returns a tuple with the IdentityDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDistribution

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetIdentityDistribution(v []RoleMiningIdentityDistribution)`

SetIdentityDistribution sets IdentityDistribution field to given value.

### HasIdentityDistribution

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasIdentityDistribution() bool`

HasIdentityDistribution returns a boolean if a field has been set.

### GetPotentialRoleRef

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetPotentialRoleRef() RoleMiningPotentialRoleRef`

GetPotentialRoleRef returns the PotentialRoleRef field if non-nil, zero value otherwise.

### GetPotentialRoleRefOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetPotentialRoleRefOk() (*RoleMiningPotentialRoleRef, bool)`

GetPotentialRoleRefOk returns a tuple with the PotentialRoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialRoleRef

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetPotentialRoleRef(v RoleMiningPotentialRoleRef)`

SetPotentialRoleRef sets PotentialRoleRef field to given value.

### HasPotentialRoleRef

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasPotentialRoleRef() bool`

HasPotentialRoleRef returns a boolean if a field has been set.

### GetIdentityCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetIdentityGroupStatus

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityGroupStatus() string`

GetIdentityGroupStatus returns the IdentityGroupStatus field if non-nil, zero value otherwise.

### GetIdentityGroupStatusOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetIdentityGroupStatusOk() (*string, bool)`

GetIdentityGroupStatusOk returns a tuple with the IdentityGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityGroupStatus

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetIdentityGroupStatus(v string)`

SetIdentityGroupStatus sets IdentityGroupStatus field to given value.

### HasIdentityGroupStatus

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasIdentityGroupStatus() bool`

HasIdentityGroupStatus returns a boolean if a field has been set.

### GetProvisionState

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetProvisionState() RoleMiningPotentialRoleProvisionState`

GetProvisionState returns the ProvisionState field if non-nil, zero value otherwise.

### GetProvisionStateOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetProvisionStateOk() (*RoleMiningPotentialRoleProvisionState, bool)`

GetProvisionStateOk returns a tuple with the ProvisionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionState

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetProvisionState(v RoleMiningPotentialRoleProvisionState)`

SetProvisionState sets ProvisionState field to given value.

### HasProvisionState

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasProvisionState() bool`

HasProvisionState returns a boolean if a field has been set.

### GetRoleId

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetDensity

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetDensity() int32`

GetDensity returns the Density field if non-nil, zero value otherwise.

### GetDensityOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetDensityOk() (*int32, bool)`

GetDensityOk returns a tuple with the Density field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDensity

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetDensity(v int32)`

SetDensity sets Density field to given value.

### HasDensity

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasDensity() bool`

HasDensity returns a boolean if a field has been set.

### GetFreshness

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetFreshness() int32`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetFreshnessOk() (*int32, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetFreshness(v int32)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetQuality

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *RoleMiningPotentialRoleSummaryDistribution) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *RoleMiningPotentialRoleSummaryDistribution) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *RoleMiningPotentialRoleSummaryDistribution) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


