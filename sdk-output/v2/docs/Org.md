# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**ScriptName** | Pointer to **string** |  | [optional] 
**SsoDomain** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MaxRegisteredIdentities** | Pointer to **int32** |  | [optional] 
**IdentityCount** | Pointer to **int32** |  | [optional] 
**KbaReqForAuthn** | Pointer to **int32** |  | [optional] 
**KbaReqAnswers** | Pointer to **int32** |  | [optional] 
**LockoutAttemptThreshold** | Pointer to **int32** |  | [optional] 
**LockoutTimeMinutes** | Pointer to **int32** |  | [optional] 
**UsageCertRequired** | Pointer to **bool** |  | [optional] 
**AdminStrongAuthRequired** | Pointer to **bool** |  | [optional] 
**EnableExternalPasswordChange** | Pointer to **bool** |  | [optional] 
**EnablePasswordReplay** | Pointer to **bool** |  | [optional] 
**EnableAutomaticPasswordReplay** | Pointer to **bool** |  | [optional] 
**Netmasks** | Pointer to **[]string** |  | [optional] 
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**WhiteList** | Pointer to **bool** |  | [optional] 
**EmailTestMode** | Pointer to **bool** |  | [optional] 
**EmailTestAddress** | Pointer to **string** |  | [optional] 
**UsernameEmptyText** | Pointer to **string** |  | [optional] 
**UsernameLabel** | Pointer to **string** |  | [optional] 
**EnableAutomationGeneration** | Pointer to **bool** |  | [optional] 
**PasswordReplayState** | Pointer to **string** |  | [optional] 
**SystemNotificationConfig** | Pointer to [**OrgSystemNotificationConfig**](OrgSystemNotificationConfig.md) |  | [optional] 
**SystemNotificationEmails** | Pointer to **[]string** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 
**RedirectPatterns** | Pointer to **[]string** |  | [optional] 
**StyleHash** | Pointer to **string** |  | [optional] 
**ApprovalConfig** | Pointer to [**ApprovalConfigEto**](ApprovalConfigEto.md) |  | [optional] 
**SsoPartnerSource** | Pointer to **string** |  | [optional] 

## Methods

### NewOrg

`func NewOrg() *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Org) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Org) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Org) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Org) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Org) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Org) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Org) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Org) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Org) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *Org) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Org) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Org) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Org) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Org) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Org) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Org) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Org) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetScriptName

`func (o *Org) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *Org) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *Org) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *Org) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetSsoDomain

`func (o *Org) GetSsoDomain() string`

GetSsoDomain returns the SsoDomain field if non-nil, zero value otherwise.

### GetSsoDomainOk

`func (o *Org) GetSsoDomainOk() (*string, bool)`

GetSsoDomainOk returns a tuple with the SsoDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomain

`func (o *Org) SetSsoDomain(v string)`

SetSsoDomain sets SsoDomain field to given value.

### HasSsoDomain

`func (o *Org) HasSsoDomain() bool`

HasSsoDomain returns a boolean if a field has been set.

### GetStatus

`func (o *Org) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Org) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Org) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Org) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaxRegisteredIdentities

`func (o *Org) GetMaxRegisteredIdentities() int32`

GetMaxRegisteredIdentities returns the MaxRegisteredIdentities field if non-nil, zero value otherwise.

### GetMaxRegisteredIdentitiesOk

`func (o *Org) GetMaxRegisteredIdentitiesOk() (*int32, bool)`

GetMaxRegisteredIdentitiesOk returns a tuple with the MaxRegisteredIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegisteredIdentities

`func (o *Org) SetMaxRegisteredIdentities(v int32)`

SetMaxRegisteredIdentities sets MaxRegisteredIdentities field to given value.

### HasMaxRegisteredIdentities

`func (o *Org) HasMaxRegisteredIdentities() bool`

HasMaxRegisteredIdentities returns a boolean if a field has been set.

### GetIdentityCount

`func (o *Org) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *Org) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *Org) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *Org) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetKbaReqForAuthn

`func (o *Org) GetKbaReqForAuthn() int32`

GetKbaReqForAuthn returns the KbaReqForAuthn field if non-nil, zero value otherwise.

### GetKbaReqForAuthnOk

`func (o *Org) GetKbaReqForAuthnOk() (*int32, bool)`

GetKbaReqForAuthnOk returns a tuple with the KbaReqForAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqForAuthn

`func (o *Org) SetKbaReqForAuthn(v int32)`

SetKbaReqForAuthn sets KbaReqForAuthn field to given value.

### HasKbaReqForAuthn

`func (o *Org) HasKbaReqForAuthn() bool`

HasKbaReqForAuthn returns a boolean if a field has been set.

### GetKbaReqAnswers

`func (o *Org) GetKbaReqAnswers() int32`

GetKbaReqAnswers returns the KbaReqAnswers field if non-nil, zero value otherwise.

### GetKbaReqAnswersOk

`func (o *Org) GetKbaReqAnswersOk() (*int32, bool)`

GetKbaReqAnswersOk returns a tuple with the KbaReqAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqAnswers

`func (o *Org) SetKbaReqAnswers(v int32)`

SetKbaReqAnswers sets KbaReqAnswers field to given value.

### HasKbaReqAnswers

`func (o *Org) HasKbaReqAnswers() bool`

HasKbaReqAnswers returns a boolean if a field has been set.

### GetLockoutAttemptThreshold

`func (o *Org) GetLockoutAttemptThreshold() int32`

GetLockoutAttemptThreshold returns the LockoutAttemptThreshold field if non-nil, zero value otherwise.

### GetLockoutAttemptThresholdOk

`func (o *Org) GetLockoutAttemptThresholdOk() (*int32, bool)`

GetLockoutAttemptThresholdOk returns a tuple with the LockoutAttemptThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutAttemptThreshold

`func (o *Org) SetLockoutAttemptThreshold(v int32)`

SetLockoutAttemptThreshold sets LockoutAttemptThreshold field to given value.

### HasLockoutAttemptThreshold

`func (o *Org) HasLockoutAttemptThreshold() bool`

HasLockoutAttemptThreshold returns a boolean if a field has been set.

### GetLockoutTimeMinutes

`func (o *Org) GetLockoutTimeMinutes() int32`

GetLockoutTimeMinutes returns the LockoutTimeMinutes field if non-nil, zero value otherwise.

### GetLockoutTimeMinutesOk

`func (o *Org) GetLockoutTimeMinutesOk() (*int32, bool)`

GetLockoutTimeMinutesOk returns a tuple with the LockoutTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeMinutes

`func (o *Org) SetLockoutTimeMinutes(v int32)`

SetLockoutTimeMinutes sets LockoutTimeMinutes field to given value.

### HasLockoutTimeMinutes

`func (o *Org) HasLockoutTimeMinutes() bool`

HasLockoutTimeMinutes returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *Org) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *Org) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *Org) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *Org) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetAdminStrongAuthRequired

`func (o *Org) GetAdminStrongAuthRequired() bool`

GetAdminStrongAuthRequired returns the AdminStrongAuthRequired field if non-nil, zero value otherwise.

### GetAdminStrongAuthRequiredOk

`func (o *Org) GetAdminStrongAuthRequiredOk() (*bool, bool)`

GetAdminStrongAuthRequiredOk returns a tuple with the AdminStrongAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStrongAuthRequired

`func (o *Org) SetAdminStrongAuthRequired(v bool)`

SetAdminStrongAuthRequired sets AdminStrongAuthRequired field to given value.

### HasAdminStrongAuthRequired

`func (o *Org) HasAdminStrongAuthRequired() bool`

HasAdminStrongAuthRequired returns a boolean if a field has been set.

### GetEnableExternalPasswordChange

`func (o *Org) GetEnableExternalPasswordChange() bool`

GetEnableExternalPasswordChange returns the EnableExternalPasswordChange field if non-nil, zero value otherwise.

### GetEnableExternalPasswordChangeOk

`func (o *Org) GetEnableExternalPasswordChangeOk() (*bool, bool)`

GetEnableExternalPasswordChangeOk returns a tuple with the EnableExternalPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalPasswordChange

`func (o *Org) SetEnableExternalPasswordChange(v bool)`

SetEnableExternalPasswordChange sets EnableExternalPasswordChange field to given value.

### HasEnableExternalPasswordChange

`func (o *Org) HasEnableExternalPasswordChange() bool`

HasEnableExternalPasswordChange returns a boolean if a field has been set.

### GetEnablePasswordReplay

`func (o *Org) GetEnablePasswordReplay() bool`

GetEnablePasswordReplay returns the EnablePasswordReplay field if non-nil, zero value otherwise.

### GetEnablePasswordReplayOk

`func (o *Org) GetEnablePasswordReplayOk() (*bool, bool)`

GetEnablePasswordReplayOk returns a tuple with the EnablePasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordReplay

`func (o *Org) SetEnablePasswordReplay(v bool)`

SetEnablePasswordReplay sets EnablePasswordReplay field to given value.

### HasEnablePasswordReplay

`func (o *Org) HasEnablePasswordReplay() bool`

HasEnablePasswordReplay returns a boolean if a field has been set.

### GetEnableAutomaticPasswordReplay

`func (o *Org) GetEnableAutomaticPasswordReplay() bool`

GetEnableAutomaticPasswordReplay returns the EnableAutomaticPasswordReplay field if non-nil, zero value otherwise.

### GetEnableAutomaticPasswordReplayOk

`func (o *Org) GetEnableAutomaticPasswordReplayOk() (*bool, bool)`

GetEnableAutomaticPasswordReplayOk returns a tuple with the EnableAutomaticPasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPasswordReplay

`func (o *Org) SetEnableAutomaticPasswordReplay(v bool)`

SetEnableAutomaticPasswordReplay sets EnableAutomaticPasswordReplay field to given value.

### HasEnableAutomaticPasswordReplay

`func (o *Org) HasEnableAutomaticPasswordReplay() bool`

HasEnableAutomaticPasswordReplay returns a boolean if a field has been set.

### GetNetmasks

`func (o *Org) GetNetmasks() []string`

GetNetmasks returns the Netmasks field if non-nil, zero value otherwise.

### GetNetmasksOk

`func (o *Org) GetNetmasksOk() (*[]string, bool)`

GetNetmasksOk returns a tuple with the Netmasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmasks

`func (o *Org) SetNetmasks(v []string)`

SetNetmasks sets Netmasks field to given value.

### HasNetmasks

`func (o *Org) HasNetmasks() bool`

HasNetmasks returns a boolean if a field has been set.

### GetCountryCodes

`func (o *Org) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *Org) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *Org) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *Org) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetWhiteList

`func (o *Org) GetWhiteList() bool`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *Org) GetWhiteListOk() (*bool, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *Org) SetWhiteList(v bool)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *Org) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetEmailTestMode

`func (o *Org) GetEmailTestMode() bool`

GetEmailTestMode returns the EmailTestMode field if non-nil, zero value otherwise.

### GetEmailTestModeOk

`func (o *Org) GetEmailTestModeOk() (*bool, bool)`

GetEmailTestModeOk returns a tuple with the EmailTestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestMode

`func (o *Org) SetEmailTestMode(v bool)`

SetEmailTestMode sets EmailTestMode field to given value.

### HasEmailTestMode

`func (o *Org) HasEmailTestMode() bool`

HasEmailTestMode returns a boolean if a field has been set.

### GetEmailTestAddress

`func (o *Org) GetEmailTestAddress() string`

GetEmailTestAddress returns the EmailTestAddress field if non-nil, zero value otherwise.

### GetEmailTestAddressOk

`func (o *Org) GetEmailTestAddressOk() (*string, bool)`

GetEmailTestAddressOk returns a tuple with the EmailTestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestAddress

`func (o *Org) SetEmailTestAddress(v string)`

SetEmailTestAddress sets EmailTestAddress field to given value.

### HasEmailTestAddress

`func (o *Org) HasEmailTestAddress() bool`

HasEmailTestAddress returns a boolean if a field has been set.

### GetUsernameEmptyText

`func (o *Org) GetUsernameEmptyText() string`

GetUsernameEmptyText returns the UsernameEmptyText field if non-nil, zero value otherwise.

### GetUsernameEmptyTextOk

`func (o *Org) GetUsernameEmptyTextOk() (*string, bool)`

GetUsernameEmptyTextOk returns a tuple with the UsernameEmptyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameEmptyText

`func (o *Org) SetUsernameEmptyText(v string)`

SetUsernameEmptyText sets UsernameEmptyText field to given value.

### HasUsernameEmptyText

`func (o *Org) HasUsernameEmptyText() bool`

HasUsernameEmptyText returns a boolean if a field has been set.

### GetUsernameLabel

`func (o *Org) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *Org) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *Org) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *Org) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetEnableAutomationGeneration

`func (o *Org) GetEnableAutomationGeneration() bool`

GetEnableAutomationGeneration returns the EnableAutomationGeneration field if non-nil, zero value otherwise.

### GetEnableAutomationGenerationOk

`func (o *Org) GetEnableAutomationGenerationOk() (*bool, bool)`

GetEnableAutomationGenerationOk returns a tuple with the EnableAutomationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomationGeneration

`func (o *Org) SetEnableAutomationGeneration(v bool)`

SetEnableAutomationGeneration sets EnableAutomationGeneration field to given value.

### HasEnableAutomationGeneration

`func (o *Org) HasEnableAutomationGeneration() bool`

HasEnableAutomationGeneration returns a boolean if a field has been set.

### GetPasswordReplayState

`func (o *Org) GetPasswordReplayState() string`

GetPasswordReplayState returns the PasswordReplayState field if non-nil, zero value otherwise.

### GetPasswordReplayStateOk

`func (o *Org) GetPasswordReplayStateOk() (*string, bool)`

GetPasswordReplayStateOk returns a tuple with the PasswordReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReplayState

`func (o *Org) SetPasswordReplayState(v string)`

SetPasswordReplayState sets PasswordReplayState field to given value.

### HasPasswordReplayState

`func (o *Org) HasPasswordReplayState() bool`

HasPasswordReplayState returns a boolean if a field has been set.

### GetSystemNotificationConfig

`func (o *Org) GetSystemNotificationConfig() OrgSystemNotificationConfig`

GetSystemNotificationConfig returns the SystemNotificationConfig field if non-nil, zero value otherwise.

### GetSystemNotificationConfigOk

`func (o *Org) GetSystemNotificationConfigOk() (*OrgSystemNotificationConfig, bool)`

GetSystemNotificationConfigOk returns a tuple with the SystemNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationConfig

`func (o *Org) SetSystemNotificationConfig(v OrgSystemNotificationConfig)`

SetSystemNotificationConfig sets SystemNotificationConfig field to given value.

### HasSystemNotificationConfig

`func (o *Org) HasSystemNotificationConfig() bool`

HasSystemNotificationConfig returns a boolean if a field has been set.

### GetSystemNotificationEmails

`func (o *Org) GetSystemNotificationEmails() []string`

GetSystemNotificationEmails returns the SystemNotificationEmails field if non-nil, zero value otherwise.

### GetSystemNotificationEmailsOk

`func (o *Org) GetSystemNotificationEmailsOk() (*[]string, bool)`

GetSystemNotificationEmailsOk returns a tuple with the SystemNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationEmails

`func (o *Org) SetSystemNotificationEmails(v []string)`

SetSystemNotificationEmails sets SystemNotificationEmails field to given value.

### HasSystemNotificationEmails

`func (o *Org) HasSystemNotificationEmails() bool`

HasSystemNotificationEmails returns a boolean if a field has been set.

### GetLoginUrl

`func (o *Org) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *Org) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *Org) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *Org) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetRedirectPatterns

`func (o *Org) GetRedirectPatterns() []string`

GetRedirectPatterns returns the RedirectPatterns field if non-nil, zero value otherwise.

### GetRedirectPatternsOk

`func (o *Org) GetRedirectPatternsOk() (*[]string, bool)`

GetRedirectPatternsOk returns a tuple with the RedirectPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPatterns

`func (o *Org) SetRedirectPatterns(v []string)`

SetRedirectPatterns sets RedirectPatterns field to given value.

### HasRedirectPatterns

`func (o *Org) HasRedirectPatterns() bool`

HasRedirectPatterns returns a boolean if a field has been set.

### GetStyleHash

`func (o *Org) GetStyleHash() string`

GetStyleHash returns the StyleHash field if non-nil, zero value otherwise.

### GetStyleHashOk

`func (o *Org) GetStyleHashOk() (*string, bool)`

GetStyleHashOk returns a tuple with the StyleHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleHash

`func (o *Org) SetStyleHash(v string)`

SetStyleHash sets StyleHash field to given value.

### HasStyleHash

`func (o *Org) HasStyleHash() bool`

HasStyleHash returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *Org) GetApprovalConfig() ApprovalConfigEto`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *Org) GetApprovalConfigOk() (*ApprovalConfigEto, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *Org) SetApprovalConfig(v ApprovalConfigEto)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *Org) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.

### GetSsoPartnerSource

`func (o *Org) GetSsoPartnerSource() string`

GetSsoPartnerSource returns the SsoPartnerSource field if non-nil, zero value otherwise.

### GetSsoPartnerSourceOk

`func (o *Org) GetSsoPartnerSourceOk() (*string, bool)`

GetSsoPartnerSourceOk returns a tuple with the SsoPartnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoPartnerSource

`func (o *Org) SetSsoPartnerSource(v string)`

SetSsoPartnerSource sets SsoPartnerSource field to given value.

### HasSsoPartnerSource

`func (o *Org) HasSsoPartnerSource() bool`

HasSsoPartnerSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


