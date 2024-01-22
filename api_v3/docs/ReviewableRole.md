# ReviewableRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id for the Role | [optional] 
**Name** | Pointer to **string** | The name of the Role | [optional] 
**Description** | Pointer to **string** | Information about the Role | [optional] 
**Privileged** | Pointer to **bool** | Indicates if the entitlement is a privileged entitlement | [optional] 
**Owner** | Pointer to [**NullableIdentityReferenceWithNameAndEmail**](IdentityReferenceWithNameAndEmail.md) |  | [optional] 
**Revocable** | Pointer to **bool** | Indicates whether the Role can be revoked or requested | [optional] 
**EndDate** | Pointer to **time.Time** | The date when a user&#39;s access expires. | [optional] 
**AccessProfiles** | Pointer to [**[]ReviewableAccessProfile**](ReviewableAccessProfile.md) | The list of Access Profiles associated with this Role | [optional] 

## Methods

### NewReviewableRole

`func NewReviewableRole() *ReviewableRole`

NewReviewableRole instantiates a new ReviewableRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewableRoleWithDefaults

`func NewReviewableRoleWithDefaults() *ReviewableRole`

NewReviewableRoleWithDefaults instantiates a new ReviewableRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewableRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewableRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewableRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReviewableRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReviewableRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewableRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewableRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReviewableRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ReviewableRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReviewableRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReviewableRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReviewableRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivileged

`func (o *ReviewableRole) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *ReviewableRole) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *ReviewableRole) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *ReviewableRole) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetOwner

`func (o *ReviewableRole) GetOwner() IdentityReferenceWithNameAndEmail`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReviewableRole) GetOwnerOk() (*IdentityReferenceWithNameAndEmail, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReviewableRole) SetOwner(v IdentityReferenceWithNameAndEmail)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ReviewableRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ReviewableRole) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ReviewableRole) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetRevocable

`func (o *ReviewableRole) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *ReviewableRole) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *ReviewableRole) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *ReviewableRole) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.

### GetEndDate

`func (o *ReviewableRole) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReviewableRole) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReviewableRole) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReviewableRole) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAccessProfiles

`func (o *ReviewableRole) GetAccessProfiles() []ReviewableAccessProfile`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *ReviewableRole) GetAccessProfilesOk() (*[]ReviewableAccessProfile, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *ReviewableRole) SetAccessProfiles(v []ReviewableAccessProfile)`

SetAccessProfiles sets AccessProfiles field to given value.

### HasAccessProfiles

`func (o *ReviewableRole) HasAccessProfiles() bool`

HasAccessProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


