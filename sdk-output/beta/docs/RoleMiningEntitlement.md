# RoleMiningEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementRef** | Pointer to [**RoleMiningEntitlementRef**](RoleMiningEntitlementRef.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the entitlement | [optional] 
**ApplicationName** | Pointer to **string** | Application name of the entitlement | [optional] 
**IdentityCount** | Pointer to **int32** | The number of identities with this entitlement in a role. | [optional] 
**Popularity** | Pointer to **int32** | The % popularity of this entitlement in a role. | [optional] 
**PopularityInOrg** | Pointer to **int32** | TThe % popularity of this entitlement in the org. | [optional] 

## Methods

### NewRoleMiningEntitlement

`func NewRoleMiningEntitlement() *RoleMiningEntitlement`

NewRoleMiningEntitlement instantiates a new RoleMiningEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningEntitlementWithDefaults

`func NewRoleMiningEntitlementWithDefaults() *RoleMiningEntitlement`

NewRoleMiningEntitlementWithDefaults instantiates a new RoleMiningEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementRef

`func (o *RoleMiningEntitlement) GetEntitlementRef() RoleMiningEntitlementRef`

GetEntitlementRef returns the EntitlementRef field if non-nil, zero value otherwise.

### GetEntitlementRefOk

`func (o *RoleMiningEntitlement) GetEntitlementRefOk() (*RoleMiningEntitlementRef, bool)`

GetEntitlementRefOk returns a tuple with the EntitlementRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementRef

`func (o *RoleMiningEntitlement) SetEntitlementRef(v RoleMiningEntitlementRef)`

SetEntitlementRef sets EntitlementRef field to given value.

### HasEntitlementRef

`func (o *RoleMiningEntitlement) HasEntitlementRef() bool`

HasEntitlementRef returns a boolean if a field has been set.

### GetName

`func (o *RoleMiningEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleMiningEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleMiningEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleMiningEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplicationName

`func (o *RoleMiningEntitlement) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *RoleMiningEntitlement) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *RoleMiningEntitlement) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *RoleMiningEntitlement) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetIdentityCount

`func (o *RoleMiningEntitlement) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *RoleMiningEntitlement) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *RoleMiningEntitlement) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *RoleMiningEntitlement) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetPopularity

`func (o *RoleMiningEntitlement) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *RoleMiningEntitlement) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *RoleMiningEntitlement) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *RoleMiningEntitlement) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPopularityInOrg

`func (o *RoleMiningEntitlement) GetPopularityInOrg() int32`

GetPopularityInOrg returns the PopularityInOrg field if non-nil, zero value otherwise.

### GetPopularityInOrgOk

`func (o *RoleMiningEntitlement) GetPopularityInOrgOk() (*int32, bool)`

GetPopularityInOrgOk returns a tuple with the PopularityInOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularityInOrg

`func (o *RoleMiningEntitlement) SetPopularityInOrg(v int32)`

SetPopularityInOrg sets PopularityInOrg field to given value.

### HasPopularityInOrg

`func (o *RoleMiningEntitlement) HasPopularityInOrg() bool`

HasPopularityInOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


