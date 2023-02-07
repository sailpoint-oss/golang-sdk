# RoleMiningSessionParametersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinNumIdentitiesInPotentialRole** | Pointer to **int32** | Minimum number of identities in a potential role | [optional] 
**Name** | Pointer to **string** | The session&#39;s saved name | [optional] 
**PruneThreshold** | Pointer to **int32** | The prune threshold to be used or null to calculate prescribedPruneThreshold | [optional] 
**Saved** | Pointer to **bool** | The session&#39;s saved status | [optional] 
**Scope** | Pointer to [**RoleMiningSessionScope**](RoleMiningSessionScope.md) |  | [optional] 
**Type** | Pointer to [**RoleMiningRoleType**](RoleMiningRoleType.md) |  | [optional] 

## Methods

### NewRoleMiningSessionParametersDto

`func NewRoleMiningSessionParametersDto() *RoleMiningSessionParametersDto`

NewRoleMiningSessionParametersDto instantiates a new RoleMiningSessionParametersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningSessionParametersDtoWithDefaults

`func NewRoleMiningSessionParametersDtoWithDefaults() *RoleMiningSessionParametersDto`

NewRoleMiningSessionParametersDtoWithDefaults instantiates a new RoleMiningSessionParametersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinNumIdentitiesInPotentialRole

`func (o *RoleMiningSessionParametersDto) GetMinNumIdentitiesInPotentialRole() int32`

GetMinNumIdentitiesInPotentialRole returns the MinNumIdentitiesInPotentialRole field if non-nil, zero value otherwise.

### GetMinNumIdentitiesInPotentialRoleOk

`func (o *RoleMiningSessionParametersDto) GetMinNumIdentitiesInPotentialRoleOk() (*int32, bool)`

GetMinNumIdentitiesInPotentialRoleOk returns a tuple with the MinNumIdentitiesInPotentialRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumIdentitiesInPotentialRole

`func (o *RoleMiningSessionParametersDto) SetMinNumIdentitiesInPotentialRole(v int32)`

SetMinNumIdentitiesInPotentialRole sets MinNumIdentitiesInPotentialRole field to given value.

### HasMinNumIdentitiesInPotentialRole

`func (o *RoleMiningSessionParametersDto) HasMinNumIdentitiesInPotentialRole() bool`

HasMinNumIdentitiesInPotentialRole returns a boolean if a field has been set.

### GetName

`func (o *RoleMiningSessionParametersDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleMiningSessionParametersDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleMiningSessionParametersDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleMiningSessionParametersDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPruneThreshold

`func (o *RoleMiningSessionParametersDto) GetPruneThreshold() int32`

GetPruneThreshold returns the PruneThreshold field if non-nil, zero value otherwise.

### GetPruneThresholdOk

`func (o *RoleMiningSessionParametersDto) GetPruneThresholdOk() (*int32, bool)`

GetPruneThresholdOk returns a tuple with the PruneThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneThreshold

`func (o *RoleMiningSessionParametersDto) SetPruneThreshold(v int32)`

SetPruneThreshold sets PruneThreshold field to given value.

### HasPruneThreshold

`func (o *RoleMiningSessionParametersDto) HasPruneThreshold() bool`

HasPruneThreshold returns a boolean if a field has been set.

### GetSaved

`func (o *RoleMiningSessionParametersDto) GetSaved() bool`

GetSaved returns the Saved field if non-nil, zero value otherwise.

### GetSavedOk

`func (o *RoleMiningSessionParametersDto) GetSavedOk() (*bool, bool)`

GetSavedOk returns a tuple with the Saved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaved

`func (o *RoleMiningSessionParametersDto) SetSaved(v bool)`

SetSaved sets Saved field to given value.

### HasSaved

`func (o *RoleMiningSessionParametersDto) HasSaved() bool`

HasSaved returns a boolean if a field has been set.

### GetScope

`func (o *RoleMiningSessionParametersDto) GetScope() RoleMiningSessionScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RoleMiningSessionParametersDto) GetScopeOk() (*RoleMiningSessionScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RoleMiningSessionParametersDto) SetScope(v RoleMiningSessionScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RoleMiningSessionParametersDto) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *RoleMiningSessionParametersDto) GetType() RoleMiningRoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleMiningSessionParametersDto) GetTypeOk() (*RoleMiningRoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleMiningSessionParametersDto) SetType(v RoleMiningRoleType)`

SetType sets Type field to given value.

### HasType

`func (o *RoleMiningSessionParametersDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


