# RoleMiningSessionDraftRoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Draft role description | [optional] 
**IdentityIds** | Pointer to **[]string** | The list of identities for this role mining session. | [optional] 
**EntitlementIds** | Pointer to **[]string** | The list of entitlement ids for this role mining session. | [optional] 
**ExcludedEntitlements** | Pointer to **[]string** | The list of excluded entitlement ids. | [optional] 
**Modified** | Pointer to **time.Time** | Last modified date | [optional] 
**Name** | Pointer to **string** | Name of the draft role | [optional] 
**Type** | Pointer to [**RoleMiningRoleType**](RoleMiningRoleType.md) |  | [optional] 

## Methods

### NewRoleMiningSessionDraftRoleDto

`func NewRoleMiningSessionDraftRoleDto() *RoleMiningSessionDraftRoleDto`

NewRoleMiningSessionDraftRoleDto instantiates a new RoleMiningSessionDraftRoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningSessionDraftRoleDtoWithDefaults

`func NewRoleMiningSessionDraftRoleDtoWithDefaults() *RoleMiningSessionDraftRoleDto`

NewRoleMiningSessionDraftRoleDtoWithDefaults instantiates a new RoleMiningSessionDraftRoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleMiningSessionDraftRoleDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleMiningSessionDraftRoleDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleMiningSessionDraftRoleDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleMiningSessionDraftRoleDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentityIds

`func (o *RoleMiningSessionDraftRoleDto) GetIdentityIds() []string`

GetIdentityIds returns the IdentityIds field if non-nil, zero value otherwise.

### GetIdentityIdsOk

`func (o *RoleMiningSessionDraftRoleDto) GetIdentityIdsOk() (*[]string, bool)`

GetIdentityIdsOk returns a tuple with the IdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityIds

`func (o *RoleMiningSessionDraftRoleDto) SetIdentityIds(v []string)`

SetIdentityIds sets IdentityIds field to given value.

### HasIdentityIds

`func (o *RoleMiningSessionDraftRoleDto) HasIdentityIds() bool`

HasIdentityIds returns a boolean if a field has been set.

### GetEntitlementIds

`func (o *RoleMiningSessionDraftRoleDto) GetEntitlementIds() []string`

GetEntitlementIds returns the EntitlementIds field if non-nil, zero value otherwise.

### GetEntitlementIdsOk

`func (o *RoleMiningSessionDraftRoleDto) GetEntitlementIdsOk() (*[]string, bool)`

GetEntitlementIdsOk returns a tuple with the EntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementIds

`func (o *RoleMiningSessionDraftRoleDto) SetEntitlementIds(v []string)`

SetEntitlementIds sets EntitlementIds field to given value.

### HasEntitlementIds

`func (o *RoleMiningSessionDraftRoleDto) HasEntitlementIds() bool`

HasEntitlementIds returns a boolean if a field has been set.

### GetExcludedEntitlements

`func (o *RoleMiningSessionDraftRoleDto) GetExcludedEntitlements() []string`

GetExcludedEntitlements returns the ExcludedEntitlements field if non-nil, zero value otherwise.

### GetExcludedEntitlementsOk

`func (o *RoleMiningSessionDraftRoleDto) GetExcludedEntitlementsOk() (*[]string, bool)`

GetExcludedEntitlementsOk returns a tuple with the ExcludedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntitlements

`func (o *RoleMiningSessionDraftRoleDto) SetExcludedEntitlements(v []string)`

SetExcludedEntitlements sets ExcludedEntitlements field to given value.

### HasExcludedEntitlements

`func (o *RoleMiningSessionDraftRoleDto) HasExcludedEntitlements() bool`

HasExcludedEntitlements returns a boolean if a field has been set.

### GetModified

`func (o *RoleMiningSessionDraftRoleDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleMiningSessionDraftRoleDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleMiningSessionDraftRoleDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *RoleMiningSessionDraftRoleDto) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *RoleMiningSessionDraftRoleDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleMiningSessionDraftRoleDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleMiningSessionDraftRoleDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleMiningSessionDraftRoleDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RoleMiningSessionDraftRoleDto) GetType() RoleMiningRoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleMiningSessionDraftRoleDto) GetTypeOk() (*RoleMiningRoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleMiningSessionDraftRoleDto) SetType(v RoleMiningRoleType)`

SetType sets Type field to given value.

### HasType

`func (o *RoleMiningSessionDraftRoleDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


