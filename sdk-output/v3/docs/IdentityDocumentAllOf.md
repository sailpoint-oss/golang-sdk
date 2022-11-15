# IdentityDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The first name of the identity | [optional] 
**LastName** | Pointer to **string** | The last name of the identity | [optional] 
**DisplayName** | Pointer to **string** | The display name of the identity | [optional] 
**Email** | Pointer to **string** | The identity&#39;s primary email address | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Phone** | Pointer to **string** | The phone number of the identity | [optional] 
**Inactive** | Pointer to **bool** | Indicates if the identity is inactive | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | The identity&#39;s status in SailPoint | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**IsManager** | Pointer to **bool** | Indicates if this identity is a manager of other identities | [optional] 
**IdentityProfile** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**ProcessingState** | Pointer to **NullableString** |  | [optional] 
**ProcessingDetails** | Pointer to [**ProcessingDetails**](ProcessingDetails.md) |  | [optional] 
**Accounts** | Pointer to [**[]BaseAccount**](BaseAccount.md) | List of accounts associated with the identity | [optional] 
**AccountCount** | Pointer to **int32** | Number of accounts associated with the identity | [optional] 
**Apps** | Pointer to [**[]App**](App.md) | The list of applications the identity has access to | [optional] 
**AppCount** | Pointer to **int32** | The number of applications the identity has access to | [optional] 
**Access** | Pointer to [**[]Access1**](Access1.md) | The list of access items assigned to the identity | [optional] 
**AccessCount** | Pointer to **int32** | The number of access items assigned to the identity | [optional] 
**AccessProfileCount** | Pointer to **int32** | The number of access profiles assigned to the identity | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements assigned to the identity | [optional] 
**RoleCount** | Pointer to **int32** | The number of roles assigned to the identity | [optional] 
**Owns** | Pointer to [**Owns**](Owns.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIdentityDocumentAllOf

`func NewIdentityDocumentAllOf() *IdentityDocumentAllOf`

NewIdentityDocumentAllOf instantiates a new IdentityDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDocumentAllOfWithDefaults

`func NewIdentityDocumentAllOfWithDefaults() *IdentityDocumentAllOf`

NewIdentityDocumentAllOfWithDefaults instantiates a new IdentityDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *IdentityDocumentAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentityDocumentAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentityDocumentAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentityDocumentAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *IdentityDocumentAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentityDocumentAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentityDocumentAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentityDocumentAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityDocumentAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityDocumentAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityDocumentAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityDocumentAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityDocumentAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityDocumentAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityDocumentAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityDocumentAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreated

`func (o *IdentityDocumentAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentityDocumentAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentityDocumentAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentityDocumentAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *IdentityDocumentAllOf) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IdentityDocumentAllOf) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *IdentityDocumentAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IdentityDocumentAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IdentityDocumentAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IdentityDocumentAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *IdentityDocumentAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *IdentityDocumentAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *IdentityDocumentAllOf) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *IdentityDocumentAllOf) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *IdentityDocumentAllOf) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *IdentityDocumentAllOf) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *IdentityDocumentAllOf) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *IdentityDocumentAllOf) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetPhone

`func (o *IdentityDocumentAllOf) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *IdentityDocumentAllOf) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *IdentityDocumentAllOf) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *IdentityDocumentAllOf) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetInactive

`func (o *IdentityDocumentAllOf) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *IdentityDocumentAllOf) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *IdentityDocumentAllOf) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *IdentityDocumentAllOf) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetProtected

`func (o *IdentityDocumentAllOf) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *IdentityDocumentAllOf) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *IdentityDocumentAllOf) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *IdentityDocumentAllOf) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityDocumentAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityDocumentAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityDocumentAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityDocumentAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *IdentityDocumentAllOf) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *IdentityDocumentAllOf) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *IdentityDocumentAllOf) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *IdentityDocumentAllOf) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetManager

`func (o *IdentityDocumentAllOf) GetManager() DisplayReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *IdentityDocumentAllOf) GetManagerOk() (*DisplayReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *IdentityDocumentAllOf) SetManager(v DisplayReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *IdentityDocumentAllOf) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetIsManager

`func (o *IdentityDocumentAllOf) GetIsManager() bool`

GetIsManager returns the IsManager field if non-nil, zero value otherwise.

### GetIsManagerOk

`func (o *IdentityDocumentAllOf) GetIsManagerOk() (*bool, bool)`

GetIsManagerOk returns a tuple with the IsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManager

`func (o *IdentityDocumentAllOf) SetIsManager(v bool)`

SetIsManager sets IsManager field to given value.

### HasIsManager

`func (o *IdentityDocumentAllOf) HasIsManager() bool`

HasIsManager returns a boolean if a field has been set.

### GetIdentityProfile

`func (o *IdentityDocumentAllOf) GetIdentityProfile() Reference1`

GetIdentityProfile returns the IdentityProfile field if non-nil, zero value otherwise.

### GetIdentityProfileOk

`func (o *IdentityDocumentAllOf) GetIdentityProfileOk() (*Reference1, bool)`

GetIdentityProfileOk returns a tuple with the IdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProfile

`func (o *IdentityDocumentAllOf) SetIdentityProfile(v Reference1)`

SetIdentityProfile sets IdentityProfile field to given value.

### HasIdentityProfile

`func (o *IdentityDocumentAllOf) HasIdentityProfile() bool`

HasIdentityProfile returns a boolean if a field has been set.

### GetSource

`func (o *IdentityDocumentAllOf) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IdentityDocumentAllOf) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IdentityDocumentAllOf) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *IdentityDocumentAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAttributes

`func (o *IdentityDocumentAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IdentityDocumentAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IdentityDocumentAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IdentityDocumentAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProcessingState

`func (o *IdentityDocumentAllOf) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *IdentityDocumentAllOf) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *IdentityDocumentAllOf) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *IdentityDocumentAllOf) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *IdentityDocumentAllOf) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *IdentityDocumentAllOf) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetProcessingDetails

`func (o *IdentityDocumentAllOf) GetProcessingDetails() ProcessingDetails`

GetProcessingDetails returns the ProcessingDetails field if non-nil, zero value otherwise.

### GetProcessingDetailsOk

`func (o *IdentityDocumentAllOf) GetProcessingDetailsOk() (*ProcessingDetails, bool)`

GetProcessingDetailsOk returns a tuple with the ProcessingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDetails

`func (o *IdentityDocumentAllOf) SetProcessingDetails(v ProcessingDetails)`

SetProcessingDetails sets ProcessingDetails field to given value.

### HasProcessingDetails

`func (o *IdentityDocumentAllOf) HasProcessingDetails() bool`

HasProcessingDetails returns a boolean if a field has been set.

### GetAccounts

`func (o *IdentityDocumentAllOf) GetAccounts() []BaseAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *IdentityDocumentAllOf) GetAccountsOk() (*[]BaseAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *IdentityDocumentAllOf) SetAccounts(v []BaseAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *IdentityDocumentAllOf) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAccountCount

`func (o *IdentityDocumentAllOf) GetAccountCount() int32`

GetAccountCount returns the AccountCount field if non-nil, zero value otherwise.

### GetAccountCountOk

`func (o *IdentityDocumentAllOf) GetAccountCountOk() (*int32, bool)`

GetAccountCountOk returns a tuple with the AccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCount

`func (o *IdentityDocumentAllOf) SetAccountCount(v int32)`

SetAccountCount sets AccountCount field to given value.

### HasAccountCount

`func (o *IdentityDocumentAllOf) HasAccountCount() bool`

HasAccountCount returns a boolean if a field has been set.

### GetApps

`func (o *IdentityDocumentAllOf) GetApps() []App`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *IdentityDocumentAllOf) GetAppsOk() (*[]App, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *IdentityDocumentAllOf) SetApps(v []App)`

SetApps sets Apps field to given value.

### HasApps

`func (o *IdentityDocumentAllOf) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAppCount

`func (o *IdentityDocumentAllOf) GetAppCount() int32`

GetAppCount returns the AppCount field if non-nil, zero value otherwise.

### GetAppCountOk

`func (o *IdentityDocumentAllOf) GetAppCountOk() (*int32, bool)`

GetAppCountOk returns a tuple with the AppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCount

`func (o *IdentityDocumentAllOf) SetAppCount(v int32)`

SetAppCount sets AppCount field to given value.

### HasAppCount

`func (o *IdentityDocumentAllOf) HasAppCount() bool`

HasAppCount returns a boolean if a field has been set.

### GetAccess

`func (o *IdentityDocumentAllOf) GetAccess() []Access1`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *IdentityDocumentAllOf) GetAccessOk() (*[]Access1, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *IdentityDocumentAllOf) SetAccess(v []Access1)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *IdentityDocumentAllOf) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessCount

`func (o *IdentityDocumentAllOf) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *IdentityDocumentAllOf) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *IdentityDocumentAllOf) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *IdentityDocumentAllOf) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetAccessProfileCount

`func (o *IdentityDocumentAllOf) GetAccessProfileCount() int32`

GetAccessProfileCount returns the AccessProfileCount field if non-nil, zero value otherwise.

### GetAccessProfileCountOk

`func (o *IdentityDocumentAllOf) GetAccessProfileCountOk() (*int32, bool)`

GetAccessProfileCountOk returns a tuple with the AccessProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileCount

`func (o *IdentityDocumentAllOf) SetAccessProfileCount(v int32)`

SetAccessProfileCount sets AccessProfileCount field to given value.

### HasAccessProfileCount

`func (o *IdentityDocumentAllOf) HasAccessProfileCount() bool`

HasAccessProfileCount returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *IdentityDocumentAllOf) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *IdentityDocumentAllOf) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *IdentityDocumentAllOf) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *IdentityDocumentAllOf) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetRoleCount

`func (o *IdentityDocumentAllOf) GetRoleCount() int32`

GetRoleCount returns the RoleCount field if non-nil, zero value otherwise.

### GetRoleCountOk

`func (o *IdentityDocumentAllOf) GetRoleCountOk() (*int32, bool)`

GetRoleCountOk returns a tuple with the RoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCount

`func (o *IdentityDocumentAllOf) SetRoleCount(v int32)`

SetRoleCount sets RoleCount field to given value.

### HasRoleCount

`func (o *IdentityDocumentAllOf) HasRoleCount() bool`

HasRoleCount returns a boolean if a field has been set.

### GetOwns

`func (o *IdentityDocumentAllOf) GetOwns() Owns`

GetOwns returns the Owns field if non-nil, zero value otherwise.

### GetOwnsOk

`func (o *IdentityDocumentAllOf) GetOwnsOk() (*Owns, bool)`

GetOwnsOk returns a tuple with the Owns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwns

`func (o *IdentityDocumentAllOf) SetOwns(v Owns)`

SetOwns sets Owns field to given value.

### HasOwns

`func (o *IdentityDocumentAllOf) HasOwns() bool`

HasOwns returns a boolean if a field has been set.

### GetTags

`func (o *IdentityDocumentAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IdentityDocumentAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IdentityDocumentAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IdentityDocumentAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


