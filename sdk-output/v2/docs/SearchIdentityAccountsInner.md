# SearchIdentityAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**SearchIdentityAccountsInnerSource**](SearchIdentityAccountsInnerSource.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**ManuallyCorrelated** | Pointer to **bool** |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**PasswordLastSet** | Pointer to **time.Time** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**EntitlementAttributes** | Pointer to [**SearchIdentityAttributes**](SearchIdentityAttributes.md) |  | [optional] 

## Methods

### NewSearchIdentityAccountsInner

`func NewSearchIdentityAccountsInner() *SearchIdentityAccountsInner`

NewSearchIdentityAccountsInner instantiates a new SearchIdentityAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIdentityAccountsInnerWithDefaults

`func NewSearchIdentityAccountsInnerWithDefaults() *SearchIdentityAccountsInner`

NewSearchIdentityAccountsInnerWithDefaults instantiates a new SearchIdentityAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchIdentityAccountsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchIdentityAccountsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchIdentityAccountsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchIdentityAccountsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchIdentityAccountsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchIdentityAccountsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchIdentityAccountsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchIdentityAccountsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *SearchIdentityAccountsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SearchIdentityAccountsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SearchIdentityAccountsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SearchIdentityAccountsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSource

`func (o *SearchIdentityAccountsInner) GetSource() SearchIdentityAccountsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchIdentityAccountsInner) GetSourceOk() (*SearchIdentityAccountsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchIdentityAccountsInner) SetSource(v SearchIdentityAccountsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchIdentityAccountsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDisabled

`func (o *SearchIdentityAccountsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SearchIdentityAccountsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SearchIdentityAccountsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SearchIdentityAccountsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *SearchIdentityAccountsInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SearchIdentityAccountsInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SearchIdentityAccountsInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SearchIdentityAccountsInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *SearchIdentityAccountsInner) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *SearchIdentityAccountsInner) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *SearchIdentityAccountsInner) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *SearchIdentityAccountsInner) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetPrivileged

`func (o *SearchIdentityAccountsInner) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *SearchIdentityAccountsInner) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *SearchIdentityAccountsInner) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *SearchIdentityAccountsInner) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *SearchIdentityAccountsInner) GetPasswordLastSet() time.Time`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *SearchIdentityAccountsInner) GetPasswordLastSetOk() (*time.Time, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *SearchIdentityAccountsInner) SetPasswordLastSet(v time.Time)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *SearchIdentityAccountsInner) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### GetCreated

`func (o *SearchIdentityAccountsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchIdentityAccountsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchIdentityAccountsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchIdentityAccountsInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntitlementAttributes

`func (o *SearchIdentityAccountsInner) GetEntitlementAttributes() SearchIdentityAttributes`

GetEntitlementAttributes returns the EntitlementAttributes field if non-nil, zero value otherwise.

### GetEntitlementAttributesOk

`func (o *SearchIdentityAccountsInner) GetEntitlementAttributesOk() (*SearchIdentityAttributes, bool)`

GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAttributes

`func (o *SearchIdentityAccountsInner) SetEntitlementAttributes(v SearchIdentityAttributes)`

SetEntitlementAttributes sets EntitlementAttributes field to given value.

### HasEntitlementAttributes

`func (o *SearchIdentityAccountsInner) HasEntitlementAttributes() bool`

HasEntitlementAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


