# RoleMiningPotentialRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateDate** | Pointer to **[]string** | The creation date for a potential role | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements in a potential role. | [optional] 
**ExcludedEntitlements** | Pointer to **[]string** | The list of entitlement ids to be excluded. | [optional] 
**Id** | Pointer to **string** | Id of the potential role | [optional] 
**IdentityCount** | Pointer to **int32** | The number of identities in a potential role. | [optional] 
**IdentityDistribution** | Pointer to [**[]RoleMiningIdentityDistribution**](RoleMiningIdentityDistribution.md) | Identity attribute distribution | [optional] 
**IdentityIds** | Pointer to **[]string** | The list of ids in a potential role. | [optional] 
**ModifiedDate** | Pointer to **[]string** | The modified date for a potential role | [optional] 
**Name** | Pointer to **string** | Name of the potential role | [optional] 
**Type** | Pointer to [**RoleMiningRoleType**](RoleMiningRoleType.md) |  | [optional] 

## Methods

### NewRoleMiningPotentialRole

`func NewRoleMiningPotentialRole() *RoleMiningPotentialRole`

NewRoleMiningPotentialRole instantiates a new RoleMiningPotentialRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningPotentialRoleWithDefaults

`func NewRoleMiningPotentialRoleWithDefaults() *RoleMiningPotentialRole`

NewRoleMiningPotentialRoleWithDefaults instantiates a new RoleMiningPotentialRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateDate

`func (o *RoleMiningPotentialRole) GetCreateDate() []string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *RoleMiningPotentialRole) GetCreateDateOk() (*[]string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *RoleMiningPotentialRole) SetCreateDate(v []string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *RoleMiningPotentialRole) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *RoleMiningPotentialRole) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *RoleMiningPotentialRole) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *RoleMiningPotentialRole) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *RoleMiningPotentialRole) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetExcludedEntitlements

`func (o *RoleMiningPotentialRole) GetExcludedEntitlements() []string`

GetExcludedEntitlements returns the ExcludedEntitlements field if non-nil, zero value otherwise.

### GetExcludedEntitlementsOk

`func (o *RoleMiningPotentialRole) GetExcludedEntitlementsOk() (*[]string, bool)`

GetExcludedEntitlementsOk returns a tuple with the ExcludedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntitlements

`func (o *RoleMiningPotentialRole) SetExcludedEntitlements(v []string)`

SetExcludedEntitlements sets ExcludedEntitlements field to given value.

### HasExcludedEntitlements

`func (o *RoleMiningPotentialRole) HasExcludedEntitlements() bool`

HasExcludedEntitlements returns a boolean if a field has been set.

### GetId

`func (o *RoleMiningPotentialRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleMiningPotentialRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleMiningPotentialRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleMiningPotentialRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityCount

`func (o *RoleMiningPotentialRole) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *RoleMiningPotentialRole) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *RoleMiningPotentialRole) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *RoleMiningPotentialRole) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetIdentityDistribution

`func (o *RoleMiningPotentialRole) GetIdentityDistribution() []RoleMiningIdentityDistribution`

GetIdentityDistribution returns the IdentityDistribution field if non-nil, zero value otherwise.

### GetIdentityDistributionOk

`func (o *RoleMiningPotentialRole) GetIdentityDistributionOk() (*[]RoleMiningIdentityDistribution, bool)`

GetIdentityDistributionOk returns a tuple with the IdentityDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDistribution

`func (o *RoleMiningPotentialRole) SetIdentityDistribution(v []RoleMiningIdentityDistribution)`

SetIdentityDistribution sets IdentityDistribution field to given value.

### HasIdentityDistribution

`func (o *RoleMiningPotentialRole) HasIdentityDistribution() bool`

HasIdentityDistribution returns a boolean if a field has been set.

### GetIdentityIds

`func (o *RoleMiningPotentialRole) GetIdentityIds() []string`

GetIdentityIds returns the IdentityIds field if non-nil, zero value otherwise.

### GetIdentityIdsOk

`func (o *RoleMiningPotentialRole) GetIdentityIdsOk() (*[]string, bool)`

GetIdentityIdsOk returns a tuple with the IdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityIds

`func (o *RoleMiningPotentialRole) SetIdentityIds(v []string)`

SetIdentityIds sets IdentityIds field to given value.

### HasIdentityIds

`func (o *RoleMiningPotentialRole) HasIdentityIds() bool`

HasIdentityIds returns a boolean if a field has been set.

### GetModifiedDate

`func (o *RoleMiningPotentialRole) GetModifiedDate() []string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *RoleMiningPotentialRole) GetModifiedDateOk() (*[]string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *RoleMiningPotentialRole) SetModifiedDate(v []string)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *RoleMiningPotentialRole) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetName

`func (o *RoleMiningPotentialRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleMiningPotentialRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleMiningPotentialRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleMiningPotentialRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RoleMiningPotentialRole) GetType() RoleMiningRoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleMiningPotentialRole) GetTypeOk() (*RoleMiningRoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleMiningPotentialRole) SetType(v RoleMiningRoleType)`

SetType sets Type field to given value.

### HasType

`func (o *RoleMiningPotentialRole) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


