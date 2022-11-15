# SearchDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**Description** | Pointer to **string** | The description of the access item | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Requestable** | Pointer to **bool** | Indicates if the access can be requested | [optional] 
**RequestCommentsRequired** | Pointer to **bool** | Indicates if comments are required when requesting access | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Entitlements** | Pointer to [**[]BaseEntitlement**](BaseEntitlement.md) |  | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements assigned to the identity | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Action** | Pointer to **string** | The action that was performed | [optional] 
**Stage** | Pointer to **string** | The current stage of the activity | [optional] 
**Origin** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** | The identity&#39;s status in SailPoint | [optional] 
**Requester** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**Recipient** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**Approvals** | Pointer to [**[]Approval**](Approval.md) |  | [optional] 
**OriginalRequests** | Pointer to [**[]OriginalRequest**](OriginalRequest.md) |  | [optional] 
**ExpansionItems** | Pointer to [**[]ExpansionItem**](ExpansionItem.md) |  | [optional] 
**AccountRequests** | Pointer to [**[]AccountRequest**](AccountRequest.md) |  | [optional] 
**Sources** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** | The ID of the account | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the account is disabled | [optional] 
**Locked** | Pointer to **bool** | Indicates if the account is locked | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**ManuallyCorrelated** | Pointer to **bool** | Indicates if the account has been manually correlated to an identity | [optional] 
**PasswordLastSet** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**EntitlementAttributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Identity** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Access** | Pointer to [**[]Access1**](Access1.md) | The list of access items assigned to the identity | [optional] 
**Uncorrelated** | Pointer to **bool** | Indicates if the account is not correlated to an identity | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**AvgDuration** | Pointer to **int32** |  | [optional] 
**ChangedAccounts** | Pointer to **int32** |  | [optional] 
**NextScheduled** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**StartTime** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**SourceOwner** | Pointer to **string** | John Doe | [optional] 
**Attribute** | Pointer to **string** | The name of the entitlement attribute | [optional] 
**Value** | Pointer to **string** | The value of the entitlement | [optional] 
**DisplayName** | Pointer to **string** | The display name of the identity | [optional] 
**IdentityCount** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | The type of event | [optional] 
**Actor** | Pointer to [**NameType**](NameType.md) |  | [optional] 
**Target** | Pointer to [**NameType**](NameType.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Objects** | Pointer to **[]string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**TechnicalName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** | The first name of the identity | [optional] 
**LastName** | Pointer to **string** | The last name of the identity | [optional] 
**Email** | Pointer to **string** | The identity&#39;s primary email address | [optional] 
**Phone** | Pointer to **string** | The phone number of the identity | [optional] 
**Inactive** | Pointer to **bool** | Indicates if the identity is inactive | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**IsManager** | Pointer to **bool** | Indicates if this identity is a manager of other identities | [optional] 
**IdentityProfile** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**ProcessingState** | Pointer to **NullableString** |  | [optional] 
**ProcessingDetails** | Pointer to [**ProcessingDetails**](ProcessingDetails.md) |  | [optional] 
**Accounts** | Pointer to [**[]BaseAccount**](BaseAccount.md) | List of accounts associated with the identity | [optional] 
**AccountCount** | Pointer to **int32** | Number of accounts associated with the identity | [optional] 
**Apps** | Pointer to [**[]App**](App.md) | The list of applications the identity has access to | [optional] 
**AppCount** | Pointer to **int32** | The number of applications the identity has access to | [optional] 
**AccessCount** | Pointer to **int32** | The number of access items assigned to the identity | [optional] 
**AccessProfileCount** | Pointer to **int32** |  | [optional] 
**RoleCount** | Pointer to **int32** | The number of roles assigned to the identity | [optional] 
**Owns** | Pointer to [**Owns**](Owns.md) |  | [optional] 
**AccessProfiles** | Pointer to [**[]Reference1**](Reference1.md) |  | [optional] 

## Methods

### NewSearchDocument

`func NewSearchDocument(id string, name string, type_ DocumentType, ) *SearchDocument`

NewSearchDocument instantiates a new SearchDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDocumentWithDefaults

`func NewSearchDocumentWithDefaults() *SearchDocument`

NewSearchDocumentWithDefaults instantiates a new SearchDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SearchDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SearchDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetDescription

`func (o *SearchDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *SearchDocument) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchDocument) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchDocument) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchDocument) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SearchDocument) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SearchDocument) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *SearchDocument) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SearchDocument) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SearchDocument) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SearchDocument) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *SearchDocument) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *SearchDocument) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *SearchDocument) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SearchDocument) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SearchDocument) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *SearchDocument) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *SearchDocument) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *SearchDocument) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetEnabled

`func (o *SearchDocument) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SearchDocument) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SearchDocument) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SearchDocument) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestable

`func (o *SearchDocument) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *SearchDocument) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *SearchDocument) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *SearchDocument) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *SearchDocument) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *SearchDocument) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *SearchDocument) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *SearchDocument) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetOwner

`func (o *SearchDocument) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SearchDocument) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SearchDocument) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SearchDocument) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSource

`func (o *SearchDocument) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchDocument) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchDocument) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchDocument) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEntitlements

`func (o *SearchDocument) GetEntitlements() []BaseEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *SearchDocument) GetEntitlementsOk() (*[]BaseEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *SearchDocument) SetEntitlements(v []BaseEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *SearchDocument) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *SearchDocument) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *SearchDocument) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *SearchDocument) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *SearchDocument) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetTags

`func (o *SearchDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAction

`func (o *SearchDocument) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SearchDocument) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SearchDocument) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SearchDocument) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStage

`func (o *SearchDocument) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *SearchDocument) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *SearchDocument) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *SearchDocument) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetOrigin

`func (o *SearchDocument) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SearchDocument) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SearchDocument) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SearchDocument) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *SearchDocument) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *SearchDocument) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetStatus

`func (o *SearchDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequester

`func (o *SearchDocument) GetRequester() Source1`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *SearchDocument) GetRequesterOk() (*Source1, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *SearchDocument) SetRequester(v Source1)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *SearchDocument) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRecipient

`func (o *SearchDocument) GetRecipient() Source1`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SearchDocument) GetRecipientOk() (*Source1, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SearchDocument) SetRecipient(v Source1)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *SearchDocument) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *SearchDocument) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *SearchDocument) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *SearchDocument) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *SearchDocument) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetErrors

`func (o *SearchDocument) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SearchDocument) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SearchDocument) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SearchDocument) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *SearchDocument) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *SearchDocument) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *SearchDocument) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SearchDocument) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SearchDocument) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SearchDocument) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *SearchDocument) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *SearchDocument) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetApprovals

`func (o *SearchDocument) GetApprovals() []Approval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *SearchDocument) GetApprovalsOk() (*[]Approval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *SearchDocument) SetApprovals(v []Approval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *SearchDocument) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetOriginalRequests

`func (o *SearchDocument) GetOriginalRequests() []OriginalRequest`

GetOriginalRequests returns the OriginalRequests field if non-nil, zero value otherwise.

### GetOriginalRequestsOk

`func (o *SearchDocument) GetOriginalRequestsOk() (*[]OriginalRequest, bool)`

GetOriginalRequestsOk returns a tuple with the OriginalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequests

`func (o *SearchDocument) SetOriginalRequests(v []OriginalRequest)`

SetOriginalRequests sets OriginalRequests field to given value.

### HasOriginalRequests

`func (o *SearchDocument) HasOriginalRequests() bool`

HasOriginalRequests returns a boolean if a field has been set.

### GetExpansionItems

`func (o *SearchDocument) GetExpansionItems() []ExpansionItem`

GetExpansionItems returns the ExpansionItems field if non-nil, zero value otherwise.

### GetExpansionItemsOk

`func (o *SearchDocument) GetExpansionItemsOk() (*[]ExpansionItem, bool)`

GetExpansionItemsOk returns a tuple with the ExpansionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansionItems

`func (o *SearchDocument) SetExpansionItems(v []ExpansionItem)`

SetExpansionItems sets ExpansionItems field to given value.

### HasExpansionItems

`func (o *SearchDocument) HasExpansionItems() bool`

HasExpansionItems returns a boolean if a field has been set.

### GetAccountRequests

`func (o *SearchDocument) GetAccountRequests() []AccountRequest`

GetAccountRequests returns the AccountRequests field if non-nil, zero value otherwise.

### GetAccountRequestsOk

`func (o *SearchDocument) GetAccountRequestsOk() (*[]AccountRequest, bool)`

GetAccountRequestsOk returns a tuple with the AccountRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequests

`func (o *SearchDocument) SetAccountRequests(v []AccountRequest)`

SetAccountRequests sets AccountRequests field to given value.

### HasAccountRequests

`func (o *SearchDocument) HasAccountRequests() bool`

HasAccountRequests returns a boolean if a field has been set.

### GetSources

`func (o *SearchDocument) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SearchDocument) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SearchDocument) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SearchDocument) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAccountId

`func (o *SearchDocument) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SearchDocument) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SearchDocument) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SearchDocument) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDisabled

`func (o *SearchDocument) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SearchDocument) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SearchDocument) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SearchDocument) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *SearchDocument) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SearchDocument) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SearchDocument) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SearchDocument) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPrivileged

`func (o *SearchDocument) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *SearchDocument) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *SearchDocument) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *SearchDocument) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *SearchDocument) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *SearchDocument) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *SearchDocument) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *SearchDocument) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *SearchDocument) GetPasswordLastSet() time.Time`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *SearchDocument) GetPasswordLastSetOk() (*time.Time, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *SearchDocument) SetPasswordLastSet(v time.Time)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *SearchDocument) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### SetPasswordLastSetNil

`func (o *SearchDocument) SetPasswordLastSetNil(b bool)`

 SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil

### UnsetPasswordLastSet
`func (o *SearchDocument) UnsetPasswordLastSet()`

UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
### GetEntitlementAttributes

`func (o *SearchDocument) GetEntitlementAttributes() map[string]interface{}`

GetEntitlementAttributes returns the EntitlementAttributes field if non-nil, zero value otherwise.

### GetEntitlementAttributesOk

`func (o *SearchDocument) GetEntitlementAttributesOk() (*map[string]interface{}, bool)`

GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAttributes

`func (o *SearchDocument) SetEntitlementAttributes(v map[string]interface{})`

SetEntitlementAttributes sets EntitlementAttributes field to given value.

### HasEntitlementAttributes

`func (o *SearchDocument) HasEntitlementAttributes() bool`

HasEntitlementAttributes returns a boolean if a field has been set.

### SetEntitlementAttributesNil

`func (o *SearchDocument) SetEntitlementAttributesNil(b bool)`

 SetEntitlementAttributesNil sets the value for EntitlementAttributes to be an explicit nil

### UnsetEntitlementAttributes
`func (o *SearchDocument) UnsetEntitlementAttributes()`

UnsetEntitlementAttributes ensures that no value is present for EntitlementAttributes, not even an explicit nil
### GetAttributes

`func (o *SearchDocument) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SearchDocument) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SearchDocument) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SearchDocument) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIdentity

`func (o *SearchDocument) GetIdentity() DisplayReference`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SearchDocument) GetIdentityOk() (*DisplayReference, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SearchDocument) SetIdentity(v DisplayReference)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SearchDocument) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAccess

`func (o *SearchDocument) GetAccess() []Access1`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SearchDocument) GetAccessOk() (*[]Access1, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SearchDocument) SetAccess(v []Access1)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SearchDocument) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetUncorrelated

`func (o *SearchDocument) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *SearchDocument) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *SearchDocument) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *SearchDocument) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetDuration

`func (o *SearchDocument) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchDocument) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchDocument) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SearchDocument) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAvgDuration

`func (o *SearchDocument) GetAvgDuration() int32`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *SearchDocument) GetAvgDurationOk() (*int32, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *SearchDocument) SetAvgDuration(v int32)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *SearchDocument) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetChangedAccounts

`func (o *SearchDocument) GetChangedAccounts() int32`

GetChangedAccounts returns the ChangedAccounts field if non-nil, zero value otherwise.

### GetChangedAccountsOk

`func (o *SearchDocument) GetChangedAccountsOk() (*int32, bool)`

GetChangedAccountsOk returns a tuple with the ChangedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAccounts

`func (o *SearchDocument) SetChangedAccounts(v int32)`

SetChangedAccounts sets ChangedAccounts field to given value.

### HasChangedAccounts

`func (o *SearchDocument) HasChangedAccounts() bool`

HasChangedAccounts returns a boolean if a field has been set.

### GetNextScheduled

`func (o *SearchDocument) GetNextScheduled() time.Time`

GetNextScheduled returns the NextScheduled field if non-nil, zero value otherwise.

### GetNextScheduledOk

`func (o *SearchDocument) GetNextScheduledOk() (*time.Time, bool)`

GetNextScheduledOk returns a tuple with the NextScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduled

`func (o *SearchDocument) SetNextScheduled(v time.Time)`

SetNextScheduled sets NextScheduled field to given value.

### HasNextScheduled

`func (o *SearchDocument) HasNextScheduled() bool`

HasNextScheduled returns a boolean if a field has been set.

### SetNextScheduledNil

`func (o *SearchDocument) SetNextScheduledNil(b bool)`

 SetNextScheduledNil sets the value for NextScheduled to be an explicit nil

### UnsetNextScheduled
`func (o *SearchDocument) UnsetNextScheduled()`

UnsetNextScheduled ensures that no value is present for NextScheduled, not even an explicit nil
### GetStartTime

`func (o *SearchDocument) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchDocument) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchDocument) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SearchDocument) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *SearchDocument) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *SearchDocument) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetSourceOwner

`func (o *SearchDocument) GetSourceOwner() string`

GetSourceOwner returns the SourceOwner field if non-nil, zero value otherwise.

### GetSourceOwnerOk

`func (o *SearchDocument) GetSourceOwnerOk() (*string, bool)`

GetSourceOwnerOk returns a tuple with the SourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwner

`func (o *SearchDocument) SetSourceOwner(v string)`

SetSourceOwner sets SourceOwner field to given value.

### HasSourceOwner

`func (o *SearchDocument) HasSourceOwner() bool`

HasSourceOwner returns a boolean if a field has been set.

### GetAttribute

`func (o *SearchDocument) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SearchDocument) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SearchDocument) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *SearchDocument) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *SearchDocument) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchDocument) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchDocument) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchDocument) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchDocument) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchDocument) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchDocument) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SearchDocument) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIdentityCount

`func (o *SearchDocument) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *SearchDocument) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *SearchDocument) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *SearchDocument) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetType

`func (o *SearchDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActor

`func (o *SearchDocument) GetActor() NameType`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SearchDocument) GetActorOk() (*NameType, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SearchDocument) SetActor(v NameType)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SearchDocument) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTarget

`func (o *SearchDocument) GetTarget() NameType`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SearchDocument) GetTargetOk() (*NameType, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SearchDocument) SetTarget(v NameType)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SearchDocument) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStack

`func (o *SearchDocument) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *SearchDocument) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *SearchDocument) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *SearchDocument) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetIpAddress

`func (o *SearchDocument) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SearchDocument) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SearchDocument) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SearchDocument) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDetails

`func (o *SearchDocument) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SearchDocument) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SearchDocument) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SearchDocument) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetObjects

`func (o *SearchDocument) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *SearchDocument) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *SearchDocument) SetObjects(v []string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *SearchDocument) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOperation

`func (o *SearchDocument) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SearchDocument) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SearchDocument) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SearchDocument) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTechnicalName

`func (o *SearchDocument) GetTechnicalName() string`

GetTechnicalName returns the TechnicalName field if non-nil, zero value otherwise.

### GetTechnicalNameOk

`func (o *SearchDocument) GetTechnicalNameOk() (*string, bool)`

GetTechnicalNameOk returns a tuple with the TechnicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalName

`func (o *SearchDocument) SetTechnicalName(v string)`

SetTechnicalName sets TechnicalName field to given value.

### HasTechnicalName

`func (o *SearchDocument) HasTechnicalName() bool`

HasTechnicalName returns a boolean if a field has been set.

### GetFirstName

`func (o *SearchDocument) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SearchDocument) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SearchDocument) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SearchDocument) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SearchDocument) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SearchDocument) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SearchDocument) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SearchDocument) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *SearchDocument) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SearchDocument) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SearchDocument) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SearchDocument) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *SearchDocument) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SearchDocument) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SearchDocument) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SearchDocument) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetInactive

`func (o *SearchDocument) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *SearchDocument) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *SearchDocument) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *SearchDocument) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetProtected

`func (o *SearchDocument) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *SearchDocument) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *SearchDocument) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *SearchDocument) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *SearchDocument) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *SearchDocument) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *SearchDocument) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *SearchDocument) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetManager

`func (o *SearchDocument) GetManager() DisplayReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *SearchDocument) GetManagerOk() (*DisplayReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *SearchDocument) SetManager(v DisplayReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *SearchDocument) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetIsManager

`func (o *SearchDocument) GetIsManager() bool`

GetIsManager returns the IsManager field if non-nil, zero value otherwise.

### GetIsManagerOk

`func (o *SearchDocument) GetIsManagerOk() (*bool, bool)`

GetIsManagerOk returns a tuple with the IsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManager

`func (o *SearchDocument) SetIsManager(v bool)`

SetIsManager sets IsManager field to given value.

### HasIsManager

`func (o *SearchDocument) HasIsManager() bool`

HasIsManager returns a boolean if a field has been set.

### GetIdentityProfile

`func (o *SearchDocument) GetIdentityProfile() Reference1`

GetIdentityProfile returns the IdentityProfile field if non-nil, zero value otherwise.

### GetIdentityProfileOk

`func (o *SearchDocument) GetIdentityProfileOk() (*Reference1, bool)`

GetIdentityProfileOk returns a tuple with the IdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProfile

`func (o *SearchDocument) SetIdentityProfile(v Reference1)`

SetIdentityProfile sets IdentityProfile field to given value.

### HasIdentityProfile

`func (o *SearchDocument) HasIdentityProfile() bool`

HasIdentityProfile returns a boolean if a field has been set.

### GetProcessingState

`func (o *SearchDocument) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *SearchDocument) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *SearchDocument) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *SearchDocument) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *SearchDocument) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *SearchDocument) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetProcessingDetails

`func (o *SearchDocument) GetProcessingDetails() ProcessingDetails`

GetProcessingDetails returns the ProcessingDetails field if non-nil, zero value otherwise.

### GetProcessingDetailsOk

`func (o *SearchDocument) GetProcessingDetailsOk() (*ProcessingDetails, bool)`

GetProcessingDetailsOk returns a tuple with the ProcessingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDetails

`func (o *SearchDocument) SetProcessingDetails(v ProcessingDetails)`

SetProcessingDetails sets ProcessingDetails field to given value.

### HasProcessingDetails

`func (o *SearchDocument) HasProcessingDetails() bool`

HasProcessingDetails returns a boolean if a field has been set.

### GetAccounts

`func (o *SearchDocument) GetAccounts() []BaseAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SearchDocument) GetAccountsOk() (*[]BaseAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SearchDocument) SetAccounts(v []BaseAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SearchDocument) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAccountCount

`func (o *SearchDocument) GetAccountCount() int32`

GetAccountCount returns the AccountCount field if non-nil, zero value otherwise.

### GetAccountCountOk

`func (o *SearchDocument) GetAccountCountOk() (*int32, bool)`

GetAccountCountOk returns a tuple with the AccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCount

`func (o *SearchDocument) SetAccountCount(v int32)`

SetAccountCount sets AccountCount field to given value.

### HasAccountCount

`func (o *SearchDocument) HasAccountCount() bool`

HasAccountCount returns a boolean if a field has been set.

### GetApps

`func (o *SearchDocument) GetApps() []App`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SearchDocument) GetAppsOk() (*[]App, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SearchDocument) SetApps(v []App)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SearchDocument) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAppCount

`func (o *SearchDocument) GetAppCount() int32`

GetAppCount returns the AppCount field if non-nil, zero value otherwise.

### GetAppCountOk

`func (o *SearchDocument) GetAppCountOk() (*int32, bool)`

GetAppCountOk returns a tuple with the AppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCount

`func (o *SearchDocument) SetAppCount(v int32)`

SetAppCount sets AppCount field to given value.

### HasAppCount

`func (o *SearchDocument) HasAppCount() bool`

HasAppCount returns a boolean if a field has been set.

### GetAccessCount

`func (o *SearchDocument) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *SearchDocument) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *SearchDocument) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *SearchDocument) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetAccessProfileCount

`func (o *SearchDocument) GetAccessProfileCount() int32`

GetAccessProfileCount returns the AccessProfileCount field if non-nil, zero value otherwise.

### GetAccessProfileCountOk

`func (o *SearchDocument) GetAccessProfileCountOk() (*int32, bool)`

GetAccessProfileCountOk returns a tuple with the AccessProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileCount

`func (o *SearchDocument) SetAccessProfileCount(v int32)`

SetAccessProfileCount sets AccessProfileCount field to given value.

### HasAccessProfileCount

`func (o *SearchDocument) HasAccessProfileCount() bool`

HasAccessProfileCount returns a boolean if a field has been set.

### GetRoleCount

`func (o *SearchDocument) GetRoleCount() int32`

GetRoleCount returns the RoleCount field if non-nil, zero value otherwise.

### GetRoleCountOk

`func (o *SearchDocument) GetRoleCountOk() (*int32, bool)`

GetRoleCountOk returns a tuple with the RoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCount

`func (o *SearchDocument) SetRoleCount(v int32)`

SetRoleCount sets RoleCount field to given value.

### HasRoleCount

`func (o *SearchDocument) HasRoleCount() bool`

HasRoleCount returns a boolean if a field has been set.

### GetOwns

`func (o *SearchDocument) GetOwns() Owns`

GetOwns returns the Owns field if non-nil, zero value otherwise.

### GetOwnsOk

`func (o *SearchDocument) GetOwnsOk() (*Owns, bool)`

GetOwnsOk returns a tuple with the Owns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwns

`func (o *SearchDocument) SetOwns(v Owns)`

SetOwns sets Owns field to given value.

### HasOwns

`func (o *SearchDocument) HasOwns() bool`

HasOwns returns a boolean if a field has been set.

### GetAccessProfiles

`func (o *SearchDocument) GetAccessProfiles() []Reference1`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *SearchDocument) GetAccessProfilesOk() (*[]Reference1, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *SearchDocument) SetAccessProfiles(v []Reference1)`

SetAccessProfiles sets AccessProfiles field to given value.

### HasAccessProfiles

`func (o *SearchDocument) HasAccessProfiles() bool`

HasAccessProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


