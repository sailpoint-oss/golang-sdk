# SearchIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Synced** | Pointer to **time.Time** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**LauncherAccount**](LauncherAccount.md) |  | [optional] 
**IdentityProfile** | Pointer to [**LauncherAccount**](LauncherAccount.md) |  | [optional] 
**Manager** | Pointer to [**SearchIdentityManager**](SearchIdentityManager.md) |  | [optional] 
**Attributes** | Pointer to [**SearchIdentityAttributes**](SearchIdentityAttributes.md) |  | [optional] 
**ProcessingState** | Pointer to **string** |  | [optional] 
**ProcessingDetails** | Pointer to [**SearchIdentityProcessingDetails**](SearchIdentityProcessingDetails.md) |  | [optional] 
**AccountCount** | Pointer to **int32** |  | [optional] 
**Accounts** | Pointer to [**[]SearchIdentityAccountsInner**](SearchIdentityAccountsInner.md) |  | [optional] 
**AppCount** | Pointer to **int32** |  | [optional] 
**Apps** | Pointer to [**[]SearchIdentityAppsInner**](SearchIdentityAppsInner.md) |  | [optional] 
**AccessCount** | Pointer to **int32** |  | [optional] 
**EntitlementCount** | Pointer to **int32** |  | [optional] 
**RoleCount** | Pointer to **int32** |  | [optional] 
**AccessProfileCount** | Pointer to **int32** |  | [optional] 
**Access** | Pointer to [**[]SearchIdentityAccessInner**](SearchIdentityAccessInner.md) |  | [optional] 

## Methods

### NewSearchIdentity

`func NewSearchIdentity() *SearchIdentity`

NewSearchIdentity instantiates a new SearchIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIdentityWithDefaults

`func NewSearchIdentityWithDefaults() *SearchIdentity`

NewSearchIdentityWithDefaults instantiates a new SearchIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *SearchIdentity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SearchIdentity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SearchIdentity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SearchIdentity) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SearchIdentity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SearchIdentity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SearchIdentity) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SearchIdentity) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SearchIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *SearchIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *SearchIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SearchIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SearchIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SearchIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreated

`func (o *SearchIdentity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchIdentity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchIdentity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchIdentity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *SearchIdentity) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SearchIdentity) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SearchIdentity) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SearchIdentity) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSynced

`func (o *SearchIdentity) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SearchIdentity) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SearchIdentity) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *SearchIdentity) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetPhone

`func (o *SearchIdentity) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SearchIdentity) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SearchIdentity) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SearchIdentity) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetInactive

`func (o *SearchIdentity) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *SearchIdentity) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *SearchIdentity) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *SearchIdentity) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetStatus

`func (o *SearchIdentity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchIdentity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchIdentity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchIdentity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *SearchIdentity) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *SearchIdentity) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *SearchIdentity) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *SearchIdentity) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetSource

`func (o *SearchIdentity) GetSource() LauncherAccount`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchIdentity) GetSourceOk() (*LauncherAccount, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchIdentity) SetSource(v LauncherAccount)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchIdentity) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIdentityProfile

`func (o *SearchIdentity) GetIdentityProfile() LauncherAccount`

GetIdentityProfile returns the IdentityProfile field if non-nil, zero value otherwise.

### GetIdentityProfileOk

`func (o *SearchIdentity) GetIdentityProfileOk() (*LauncherAccount, bool)`

GetIdentityProfileOk returns a tuple with the IdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProfile

`func (o *SearchIdentity) SetIdentityProfile(v LauncherAccount)`

SetIdentityProfile sets IdentityProfile field to given value.

### HasIdentityProfile

`func (o *SearchIdentity) HasIdentityProfile() bool`

HasIdentityProfile returns a boolean if a field has been set.

### GetManager

`func (o *SearchIdentity) GetManager() SearchIdentityManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *SearchIdentity) GetManagerOk() (*SearchIdentityManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *SearchIdentity) SetManager(v SearchIdentityManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *SearchIdentity) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetAttributes

`func (o *SearchIdentity) GetAttributes() SearchIdentityAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SearchIdentity) GetAttributesOk() (*SearchIdentityAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SearchIdentity) SetAttributes(v SearchIdentityAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SearchIdentity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProcessingState

`func (o *SearchIdentity) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *SearchIdentity) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *SearchIdentity) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *SearchIdentity) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### GetProcessingDetails

`func (o *SearchIdentity) GetProcessingDetails() SearchIdentityProcessingDetails`

GetProcessingDetails returns the ProcessingDetails field if non-nil, zero value otherwise.

### GetProcessingDetailsOk

`func (o *SearchIdentity) GetProcessingDetailsOk() (*SearchIdentityProcessingDetails, bool)`

GetProcessingDetailsOk returns a tuple with the ProcessingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDetails

`func (o *SearchIdentity) SetProcessingDetails(v SearchIdentityProcessingDetails)`

SetProcessingDetails sets ProcessingDetails field to given value.

### HasProcessingDetails

`func (o *SearchIdentity) HasProcessingDetails() bool`

HasProcessingDetails returns a boolean if a field has been set.

### GetAccountCount

`func (o *SearchIdentity) GetAccountCount() int32`

GetAccountCount returns the AccountCount field if non-nil, zero value otherwise.

### GetAccountCountOk

`func (o *SearchIdentity) GetAccountCountOk() (*int32, bool)`

GetAccountCountOk returns a tuple with the AccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCount

`func (o *SearchIdentity) SetAccountCount(v int32)`

SetAccountCount sets AccountCount field to given value.

### HasAccountCount

`func (o *SearchIdentity) HasAccountCount() bool`

HasAccountCount returns a boolean if a field has been set.

### GetAccounts

`func (o *SearchIdentity) GetAccounts() []SearchIdentityAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SearchIdentity) GetAccountsOk() (*[]SearchIdentityAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SearchIdentity) SetAccounts(v []SearchIdentityAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SearchIdentity) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAppCount

`func (o *SearchIdentity) GetAppCount() int32`

GetAppCount returns the AppCount field if non-nil, zero value otherwise.

### GetAppCountOk

`func (o *SearchIdentity) GetAppCountOk() (*int32, bool)`

GetAppCountOk returns a tuple with the AppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCount

`func (o *SearchIdentity) SetAppCount(v int32)`

SetAppCount sets AppCount field to given value.

### HasAppCount

`func (o *SearchIdentity) HasAppCount() bool`

HasAppCount returns a boolean if a field has been set.

### GetApps

`func (o *SearchIdentity) GetApps() []SearchIdentityAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SearchIdentity) GetAppsOk() (*[]SearchIdentityAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SearchIdentity) SetApps(v []SearchIdentityAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SearchIdentity) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAccessCount

`func (o *SearchIdentity) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *SearchIdentity) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *SearchIdentity) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *SearchIdentity) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *SearchIdentity) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *SearchIdentity) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *SearchIdentity) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *SearchIdentity) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetRoleCount

`func (o *SearchIdentity) GetRoleCount() int32`

GetRoleCount returns the RoleCount field if non-nil, zero value otherwise.

### GetRoleCountOk

`func (o *SearchIdentity) GetRoleCountOk() (*int32, bool)`

GetRoleCountOk returns a tuple with the RoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCount

`func (o *SearchIdentity) SetRoleCount(v int32)`

SetRoleCount sets RoleCount field to given value.

### HasRoleCount

`func (o *SearchIdentity) HasRoleCount() bool`

HasRoleCount returns a boolean if a field has been set.

### GetAccessProfileCount

`func (o *SearchIdentity) GetAccessProfileCount() int32`

GetAccessProfileCount returns the AccessProfileCount field if non-nil, zero value otherwise.

### GetAccessProfileCountOk

`func (o *SearchIdentity) GetAccessProfileCountOk() (*int32, bool)`

GetAccessProfileCountOk returns a tuple with the AccessProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileCount

`func (o *SearchIdentity) SetAccessProfileCount(v int32)`

SetAccessProfileCount sets AccessProfileCount field to given value.

### HasAccessProfileCount

`func (o *SearchIdentity) HasAccessProfileCount() bool`

HasAccessProfileCount returns a boolean if a field has been set.

### GetAccess

`func (o *SearchIdentity) GetAccess() []SearchIdentityAccessInner`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SearchIdentity) GetAccessOk() (*[]SearchIdentityAccessInner, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SearchIdentity) SetAccess(v []SearchIdentityAccessInner)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SearchIdentity) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


