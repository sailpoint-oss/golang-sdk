# GetOrgSettings200Response

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
**SystemNotificationConfig** | Pointer to [**GetOrgSettings200ResponseSystemNotificationConfig**](GetOrgSettings200ResponseSystemNotificationConfig.md) |  | [optional] 
**SystemNotificationEmails** | Pointer to **[]string** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 
**RedirectPatterns** | Pointer to **[]string** |  | [optional] 
**StyleHash** | Pointer to **string** |  | [optional] 
**ApprovalConfig** | Pointer to [**GetOrgSettings200ResponseApprovalConfig**](GetOrgSettings200ResponseApprovalConfig.md) |  | [optional] 
**SsoPartnerSource** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOrgSettings200Response

`func NewGetOrgSettings200Response() *GetOrgSettings200Response`

NewGetOrgSettings200Response instantiates a new GetOrgSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgSettings200ResponseWithDefaults

`func NewGetOrgSettings200ResponseWithDefaults() *GetOrgSettings200Response`

NewGetOrgSettings200ResponseWithDefaults instantiates a new GetOrgSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrgSettings200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrgSettings200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrgSettings200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrgSettings200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOrgSettings200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrgSettings200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrgSettings200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrgSettings200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetOrgSettings200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrgSettings200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrgSettings200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrgSettings200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetOrgSettings200Response) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetOrgSettings200Response) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetOrgSettings200Response) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetOrgSettings200Response) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetOrgSettings200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetOrgSettings200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetOrgSettings200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetOrgSettings200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetScriptName

`func (o *GetOrgSettings200Response) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *GetOrgSettings200Response) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *GetOrgSettings200Response) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *GetOrgSettings200Response) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetSsoDomain

`func (o *GetOrgSettings200Response) GetSsoDomain() string`

GetSsoDomain returns the SsoDomain field if non-nil, zero value otherwise.

### GetSsoDomainOk

`func (o *GetOrgSettings200Response) GetSsoDomainOk() (*string, bool)`

GetSsoDomainOk returns a tuple with the SsoDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomain

`func (o *GetOrgSettings200Response) SetSsoDomain(v string)`

SetSsoDomain sets SsoDomain field to given value.

### HasSsoDomain

`func (o *GetOrgSettings200Response) HasSsoDomain() bool`

HasSsoDomain returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrgSettings200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrgSettings200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrgSettings200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrgSettings200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaxRegisteredIdentities

`func (o *GetOrgSettings200Response) GetMaxRegisteredIdentities() int32`

GetMaxRegisteredIdentities returns the MaxRegisteredIdentities field if non-nil, zero value otherwise.

### GetMaxRegisteredIdentitiesOk

`func (o *GetOrgSettings200Response) GetMaxRegisteredIdentitiesOk() (*int32, bool)`

GetMaxRegisteredIdentitiesOk returns a tuple with the MaxRegisteredIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegisteredIdentities

`func (o *GetOrgSettings200Response) SetMaxRegisteredIdentities(v int32)`

SetMaxRegisteredIdentities sets MaxRegisteredIdentities field to given value.

### HasMaxRegisteredIdentities

`func (o *GetOrgSettings200Response) HasMaxRegisteredIdentities() bool`

HasMaxRegisteredIdentities returns a boolean if a field has been set.

### GetIdentityCount

`func (o *GetOrgSettings200Response) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *GetOrgSettings200Response) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *GetOrgSettings200Response) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *GetOrgSettings200Response) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetKbaReqForAuthn

`func (o *GetOrgSettings200Response) GetKbaReqForAuthn() int32`

GetKbaReqForAuthn returns the KbaReqForAuthn field if non-nil, zero value otherwise.

### GetKbaReqForAuthnOk

`func (o *GetOrgSettings200Response) GetKbaReqForAuthnOk() (*int32, bool)`

GetKbaReqForAuthnOk returns a tuple with the KbaReqForAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqForAuthn

`func (o *GetOrgSettings200Response) SetKbaReqForAuthn(v int32)`

SetKbaReqForAuthn sets KbaReqForAuthn field to given value.

### HasKbaReqForAuthn

`func (o *GetOrgSettings200Response) HasKbaReqForAuthn() bool`

HasKbaReqForAuthn returns a boolean if a field has been set.

### GetKbaReqAnswers

`func (o *GetOrgSettings200Response) GetKbaReqAnswers() int32`

GetKbaReqAnswers returns the KbaReqAnswers field if non-nil, zero value otherwise.

### GetKbaReqAnswersOk

`func (o *GetOrgSettings200Response) GetKbaReqAnswersOk() (*int32, bool)`

GetKbaReqAnswersOk returns a tuple with the KbaReqAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqAnswers

`func (o *GetOrgSettings200Response) SetKbaReqAnswers(v int32)`

SetKbaReqAnswers sets KbaReqAnswers field to given value.

### HasKbaReqAnswers

`func (o *GetOrgSettings200Response) HasKbaReqAnswers() bool`

HasKbaReqAnswers returns a boolean if a field has been set.

### GetLockoutAttemptThreshold

`func (o *GetOrgSettings200Response) GetLockoutAttemptThreshold() int32`

GetLockoutAttemptThreshold returns the LockoutAttemptThreshold field if non-nil, zero value otherwise.

### GetLockoutAttemptThresholdOk

`func (o *GetOrgSettings200Response) GetLockoutAttemptThresholdOk() (*int32, bool)`

GetLockoutAttemptThresholdOk returns a tuple with the LockoutAttemptThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutAttemptThreshold

`func (o *GetOrgSettings200Response) SetLockoutAttemptThreshold(v int32)`

SetLockoutAttemptThreshold sets LockoutAttemptThreshold field to given value.

### HasLockoutAttemptThreshold

`func (o *GetOrgSettings200Response) HasLockoutAttemptThreshold() bool`

HasLockoutAttemptThreshold returns a boolean if a field has been set.

### GetLockoutTimeMinutes

`func (o *GetOrgSettings200Response) GetLockoutTimeMinutes() int32`

GetLockoutTimeMinutes returns the LockoutTimeMinutes field if non-nil, zero value otherwise.

### GetLockoutTimeMinutesOk

`func (o *GetOrgSettings200Response) GetLockoutTimeMinutesOk() (*int32, bool)`

GetLockoutTimeMinutesOk returns a tuple with the LockoutTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeMinutes

`func (o *GetOrgSettings200Response) SetLockoutTimeMinutes(v int32)`

SetLockoutTimeMinutes sets LockoutTimeMinutes field to given value.

### HasLockoutTimeMinutes

`func (o *GetOrgSettings200Response) HasLockoutTimeMinutes() bool`

HasLockoutTimeMinutes returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *GetOrgSettings200Response) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *GetOrgSettings200Response) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *GetOrgSettings200Response) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *GetOrgSettings200Response) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetAdminStrongAuthRequired

`func (o *GetOrgSettings200Response) GetAdminStrongAuthRequired() bool`

GetAdminStrongAuthRequired returns the AdminStrongAuthRequired field if non-nil, zero value otherwise.

### GetAdminStrongAuthRequiredOk

`func (o *GetOrgSettings200Response) GetAdminStrongAuthRequiredOk() (*bool, bool)`

GetAdminStrongAuthRequiredOk returns a tuple with the AdminStrongAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStrongAuthRequired

`func (o *GetOrgSettings200Response) SetAdminStrongAuthRequired(v bool)`

SetAdminStrongAuthRequired sets AdminStrongAuthRequired field to given value.

### HasAdminStrongAuthRequired

`func (o *GetOrgSettings200Response) HasAdminStrongAuthRequired() bool`

HasAdminStrongAuthRequired returns a boolean if a field has been set.

### GetEnableExternalPasswordChange

`func (o *GetOrgSettings200Response) GetEnableExternalPasswordChange() bool`

GetEnableExternalPasswordChange returns the EnableExternalPasswordChange field if non-nil, zero value otherwise.

### GetEnableExternalPasswordChangeOk

`func (o *GetOrgSettings200Response) GetEnableExternalPasswordChangeOk() (*bool, bool)`

GetEnableExternalPasswordChangeOk returns a tuple with the EnableExternalPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalPasswordChange

`func (o *GetOrgSettings200Response) SetEnableExternalPasswordChange(v bool)`

SetEnableExternalPasswordChange sets EnableExternalPasswordChange field to given value.

### HasEnableExternalPasswordChange

`func (o *GetOrgSettings200Response) HasEnableExternalPasswordChange() bool`

HasEnableExternalPasswordChange returns a boolean if a field has been set.

### GetEnablePasswordReplay

`func (o *GetOrgSettings200Response) GetEnablePasswordReplay() bool`

GetEnablePasswordReplay returns the EnablePasswordReplay field if non-nil, zero value otherwise.

### GetEnablePasswordReplayOk

`func (o *GetOrgSettings200Response) GetEnablePasswordReplayOk() (*bool, bool)`

GetEnablePasswordReplayOk returns a tuple with the EnablePasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordReplay

`func (o *GetOrgSettings200Response) SetEnablePasswordReplay(v bool)`

SetEnablePasswordReplay sets EnablePasswordReplay field to given value.

### HasEnablePasswordReplay

`func (o *GetOrgSettings200Response) HasEnablePasswordReplay() bool`

HasEnablePasswordReplay returns a boolean if a field has been set.

### GetEnableAutomaticPasswordReplay

`func (o *GetOrgSettings200Response) GetEnableAutomaticPasswordReplay() bool`

GetEnableAutomaticPasswordReplay returns the EnableAutomaticPasswordReplay field if non-nil, zero value otherwise.

### GetEnableAutomaticPasswordReplayOk

`func (o *GetOrgSettings200Response) GetEnableAutomaticPasswordReplayOk() (*bool, bool)`

GetEnableAutomaticPasswordReplayOk returns a tuple with the EnableAutomaticPasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPasswordReplay

`func (o *GetOrgSettings200Response) SetEnableAutomaticPasswordReplay(v bool)`

SetEnableAutomaticPasswordReplay sets EnableAutomaticPasswordReplay field to given value.

### HasEnableAutomaticPasswordReplay

`func (o *GetOrgSettings200Response) HasEnableAutomaticPasswordReplay() bool`

HasEnableAutomaticPasswordReplay returns a boolean if a field has been set.

### GetNetmasks

`func (o *GetOrgSettings200Response) GetNetmasks() []string`

GetNetmasks returns the Netmasks field if non-nil, zero value otherwise.

### GetNetmasksOk

`func (o *GetOrgSettings200Response) GetNetmasksOk() (*[]string, bool)`

GetNetmasksOk returns a tuple with the Netmasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmasks

`func (o *GetOrgSettings200Response) SetNetmasks(v []string)`

SetNetmasks sets Netmasks field to given value.

### HasNetmasks

`func (o *GetOrgSettings200Response) HasNetmasks() bool`

HasNetmasks returns a boolean if a field has been set.

### GetCountryCodes

`func (o *GetOrgSettings200Response) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *GetOrgSettings200Response) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *GetOrgSettings200Response) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *GetOrgSettings200Response) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetWhiteList

`func (o *GetOrgSettings200Response) GetWhiteList() bool`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *GetOrgSettings200Response) GetWhiteListOk() (*bool, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *GetOrgSettings200Response) SetWhiteList(v bool)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *GetOrgSettings200Response) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetEmailTestMode

`func (o *GetOrgSettings200Response) GetEmailTestMode() bool`

GetEmailTestMode returns the EmailTestMode field if non-nil, zero value otherwise.

### GetEmailTestModeOk

`func (o *GetOrgSettings200Response) GetEmailTestModeOk() (*bool, bool)`

GetEmailTestModeOk returns a tuple with the EmailTestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestMode

`func (o *GetOrgSettings200Response) SetEmailTestMode(v bool)`

SetEmailTestMode sets EmailTestMode field to given value.

### HasEmailTestMode

`func (o *GetOrgSettings200Response) HasEmailTestMode() bool`

HasEmailTestMode returns a boolean if a field has been set.

### GetEmailTestAddress

`func (o *GetOrgSettings200Response) GetEmailTestAddress() string`

GetEmailTestAddress returns the EmailTestAddress field if non-nil, zero value otherwise.

### GetEmailTestAddressOk

`func (o *GetOrgSettings200Response) GetEmailTestAddressOk() (*string, bool)`

GetEmailTestAddressOk returns a tuple with the EmailTestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestAddress

`func (o *GetOrgSettings200Response) SetEmailTestAddress(v string)`

SetEmailTestAddress sets EmailTestAddress field to given value.

### HasEmailTestAddress

`func (o *GetOrgSettings200Response) HasEmailTestAddress() bool`

HasEmailTestAddress returns a boolean if a field has been set.

### GetUsernameEmptyText

`func (o *GetOrgSettings200Response) GetUsernameEmptyText() string`

GetUsernameEmptyText returns the UsernameEmptyText field if non-nil, zero value otherwise.

### GetUsernameEmptyTextOk

`func (o *GetOrgSettings200Response) GetUsernameEmptyTextOk() (*string, bool)`

GetUsernameEmptyTextOk returns a tuple with the UsernameEmptyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameEmptyText

`func (o *GetOrgSettings200Response) SetUsernameEmptyText(v string)`

SetUsernameEmptyText sets UsernameEmptyText field to given value.

### HasUsernameEmptyText

`func (o *GetOrgSettings200Response) HasUsernameEmptyText() bool`

HasUsernameEmptyText returns a boolean if a field has been set.

### GetUsernameLabel

`func (o *GetOrgSettings200Response) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *GetOrgSettings200Response) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *GetOrgSettings200Response) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *GetOrgSettings200Response) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetEnableAutomationGeneration

`func (o *GetOrgSettings200Response) GetEnableAutomationGeneration() bool`

GetEnableAutomationGeneration returns the EnableAutomationGeneration field if non-nil, zero value otherwise.

### GetEnableAutomationGenerationOk

`func (o *GetOrgSettings200Response) GetEnableAutomationGenerationOk() (*bool, bool)`

GetEnableAutomationGenerationOk returns a tuple with the EnableAutomationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomationGeneration

`func (o *GetOrgSettings200Response) SetEnableAutomationGeneration(v bool)`

SetEnableAutomationGeneration sets EnableAutomationGeneration field to given value.

### HasEnableAutomationGeneration

`func (o *GetOrgSettings200Response) HasEnableAutomationGeneration() bool`

HasEnableAutomationGeneration returns a boolean if a field has been set.

### GetPasswordReplayState

`func (o *GetOrgSettings200Response) GetPasswordReplayState() string`

GetPasswordReplayState returns the PasswordReplayState field if non-nil, zero value otherwise.

### GetPasswordReplayStateOk

`func (o *GetOrgSettings200Response) GetPasswordReplayStateOk() (*string, bool)`

GetPasswordReplayStateOk returns a tuple with the PasswordReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReplayState

`func (o *GetOrgSettings200Response) SetPasswordReplayState(v string)`

SetPasswordReplayState sets PasswordReplayState field to given value.

### HasPasswordReplayState

`func (o *GetOrgSettings200Response) HasPasswordReplayState() bool`

HasPasswordReplayState returns a boolean if a field has been set.

### GetSystemNotificationConfig

`func (o *GetOrgSettings200Response) GetSystemNotificationConfig() GetOrgSettings200ResponseSystemNotificationConfig`

GetSystemNotificationConfig returns the SystemNotificationConfig field if non-nil, zero value otherwise.

### GetSystemNotificationConfigOk

`func (o *GetOrgSettings200Response) GetSystemNotificationConfigOk() (*GetOrgSettings200ResponseSystemNotificationConfig, bool)`

GetSystemNotificationConfigOk returns a tuple with the SystemNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationConfig

`func (o *GetOrgSettings200Response) SetSystemNotificationConfig(v GetOrgSettings200ResponseSystemNotificationConfig)`

SetSystemNotificationConfig sets SystemNotificationConfig field to given value.

### HasSystemNotificationConfig

`func (o *GetOrgSettings200Response) HasSystemNotificationConfig() bool`

HasSystemNotificationConfig returns a boolean if a field has been set.

### GetSystemNotificationEmails

`func (o *GetOrgSettings200Response) GetSystemNotificationEmails() []string`

GetSystemNotificationEmails returns the SystemNotificationEmails field if non-nil, zero value otherwise.

### GetSystemNotificationEmailsOk

`func (o *GetOrgSettings200Response) GetSystemNotificationEmailsOk() (*[]string, bool)`

GetSystemNotificationEmailsOk returns a tuple with the SystemNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationEmails

`func (o *GetOrgSettings200Response) SetSystemNotificationEmails(v []string)`

SetSystemNotificationEmails sets SystemNotificationEmails field to given value.

### HasSystemNotificationEmails

`func (o *GetOrgSettings200Response) HasSystemNotificationEmails() bool`

HasSystemNotificationEmails returns a boolean if a field has been set.

### GetLoginUrl

`func (o *GetOrgSettings200Response) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *GetOrgSettings200Response) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *GetOrgSettings200Response) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *GetOrgSettings200Response) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetRedirectPatterns

`func (o *GetOrgSettings200Response) GetRedirectPatterns() []string`

GetRedirectPatterns returns the RedirectPatterns field if non-nil, zero value otherwise.

### GetRedirectPatternsOk

`func (o *GetOrgSettings200Response) GetRedirectPatternsOk() (*[]string, bool)`

GetRedirectPatternsOk returns a tuple with the RedirectPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPatterns

`func (o *GetOrgSettings200Response) SetRedirectPatterns(v []string)`

SetRedirectPatterns sets RedirectPatterns field to given value.

### HasRedirectPatterns

`func (o *GetOrgSettings200Response) HasRedirectPatterns() bool`

HasRedirectPatterns returns a boolean if a field has been set.

### GetStyleHash

`func (o *GetOrgSettings200Response) GetStyleHash() string`

GetStyleHash returns the StyleHash field if non-nil, zero value otherwise.

### GetStyleHashOk

`func (o *GetOrgSettings200Response) GetStyleHashOk() (*string, bool)`

GetStyleHashOk returns a tuple with the StyleHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleHash

`func (o *GetOrgSettings200Response) SetStyleHash(v string)`

SetStyleHash sets StyleHash field to given value.

### HasStyleHash

`func (o *GetOrgSettings200Response) HasStyleHash() bool`

HasStyleHash returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *GetOrgSettings200Response) GetApprovalConfig() GetOrgSettings200ResponseApprovalConfig`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *GetOrgSettings200Response) GetApprovalConfigOk() (*GetOrgSettings200ResponseApprovalConfig, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *GetOrgSettings200Response) SetApprovalConfig(v GetOrgSettings200ResponseApprovalConfig)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *GetOrgSettings200Response) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.

### GetSsoPartnerSource

`func (o *GetOrgSettings200Response) GetSsoPartnerSource() string`

GetSsoPartnerSource returns the SsoPartnerSource field if non-nil, zero value otherwise.

### GetSsoPartnerSourceOk

`func (o *GetOrgSettings200Response) GetSsoPartnerSourceOk() (*string, bool)`

GetSsoPartnerSourceOk returns a tuple with the SsoPartnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoPartnerSource

`func (o *GetOrgSettings200Response) SetSsoPartnerSource(v string)`

SetSsoPartnerSource sets SsoPartnerSource field to given value.

### HasSsoPartnerSource

`func (o *GetOrgSettings200Response) HasSsoPartnerSource() bool`

HasSsoPartnerSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


