# GetIdentity200ResponseOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ScriptName** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**NumQuestions** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MaxRegisteredUsers** | Pointer to **float32** |  | [optional] 
**Pod** | Pointer to **string** |  | [optional] 
**PwdResetPersonalPhone** | Pointer to **bool** |  | [optional] 
**PwdResetPersonalEmail** | Pointer to **bool** |  | [optional] 
**PwdResetKba** | Pointer to **bool** |  | [optional] 
**PwdResetEmail** | Pointer to **bool** |  | [optional] 
**PwdResetDuo** | Pointer to **bool** |  | [optional] 
**PwdResetPhoneMask** | Pointer to **bool** |  | [optional] 
**AuthErrorText** | Pointer to **map[string]interface{}** |  | [optional] 
**StrongAuthKba** | Pointer to **bool** |  | [optional] 
**StrongAuthPersonalPhone** | Pointer to **bool** |  | [optional] 
**StrongAuthPersonalEmail** | Pointer to **bool** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**KbaReqForAuthn** | Pointer to **float32** |  | [optional] 
**KbaReqAnswers** | Pointer to **float32** |  | [optional] 
**LockoutAttemptThreshold** | Pointer to **float32** |  | [optional] 
**LockoutTimeMinutes** | Pointer to **float32** |  | [optional] 
**UsageCertRequired** | Pointer to **bool** |  | [optional] 
**AdminStrongAuthRequired** | Pointer to **bool** |  | [optional] 
**EnableExternalPasswordChange** | Pointer to **bool** |  | [optional] 
**EnablePasswordReplay** | Pointer to **bool** |  | [optional] 
**EnableAutomaticPasswordReplay** | Pointer to **bool** |  | [optional] 
**NotifyAuthenticationSettingChange** | Pointer to **bool** |  | [optional] 
**Netmasks** | Pointer to **map[string]interface{}** |  | [optional] 
**CountryCodes** | Pointer to **map[string]interface{}** |  | [optional] 
**WhiteList** | Pointer to **bool** |  | [optional] 
**UsernameEmptyText** | Pointer to **map[string]interface{}** |  | [optional] 
**UsernameLabel** | Pointer to **map[string]interface{}** |  | [optional] 
**EnableAutomationGeneration** | Pointer to **bool** |  | [optional] 
**EmailTestMode** | Pointer to **bool** |  | [optional] 
**EmailTestAddress** | Pointer to **string** |  | [optional] 
**OrgType** | Pointer to **string** |  | [optional] 
**PasswordReplayState** | Pointer to **string** |  | [optional] 
**SystemNotificationConfig** | Pointer to **string** |  | [optional] 
**RedirectPatterns** | Pointer to **string** |  | [optional] 
**MaxClusterDebugHours** | Pointer to **string** |  | [optional] 
**BrandName** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailFromAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**StandardLogoUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**NarrowLogoUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**ActionButtonColor** | Pointer to **string** |  | [optional] 
**ActiveLinkColor** | Pointer to **string** |  | [optional] 
**NavigationColor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIdentity200ResponseOrg

`func NewGetIdentity200ResponseOrg() *GetIdentity200ResponseOrg`

NewGetIdentity200ResponseOrg instantiates a new GetIdentity200ResponseOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentity200ResponseOrgWithDefaults

`func NewGetIdentity200ResponseOrgWithDefaults() *GetIdentity200ResponseOrg`

NewGetIdentity200ResponseOrgWithDefaults instantiates a new GetIdentity200ResponseOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetIdentity200ResponseOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIdentity200ResponseOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIdentity200ResponseOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIdentity200ResponseOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScriptName

`func (o *GetIdentity200ResponseOrg) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *GetIdentity200ResponseOrg) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *GetIdentity200ResponseOrg) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *GetIdentity200ResponseOrg) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetMode

`func (o *GetIdentity200ResponseOrg) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetIdentity200ResponseOrg) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetIdentity200ResponseOrg) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetIdentity200ResponseOrg) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumQuestions

`func (o *GetIdentity200ResponseOrg) GetNumQuestions() float32`

GetNumQuestions returns the NumQuestions field if non-nil, zero value otherwise.

### GetNumQuestionsOk

`func (o *GetIdentity200ResponseOrg) GetNumQuestionsOk() (*float32, bool)`

GetNumQuestionsOk returns a tuple with the NumQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQuestions

`func (o *GetIdentity200ResponseOrg) SetNumQuestions(v float32)`

SetNumQuestions sets NumQuestions field to given value.

### HasNumQuestions

`func (o *GetIdentity200ResponseOrg) HasNumQuestions() bool`

HasNumQuestions returns a boolean if a field has been set.

### GetStatus

`func (o *GetIdentity200ResponseOrg) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIdentity200ResponseOrg) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIdentity200ResponseOrg) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIdentity200ResponseOrg) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaxRegisteredUsers

`func (o *GetIdentity200ResponseOrg) GetMaxRegisteredUsers() float32`

GetMaxRegisteredUsers returns the MaxRegisteredUsers field if non-nil, zero value otherwise.

### GetMaxRegisteredUsersOk

`func (o *GetIdentity200ResponseOrg) GetMaxRegisteredUsersOk() (*float32, bool)`

GetMaxRegisteredUsersOk returns a tuple with the MaxRegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegisteredUsers

`func (o *GetIdentity200ResponseOrg) SetMaxRegisteredUsers(v float32)`

SetMaxRegisteredUsers sets MaxRegisteredUsers field to given value.

### HasMaxRegisteredUsers

`func (o *GetIdentity200ResponseOrg) HasMaxRegisteredUsers() bool`

HasMaxRegisteredUsers returns a boolean if a field has been set.

### GetPod

`func (o *GetIdentity200ResponseOrg) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *GetIdentity200ResponseOrg) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *GetIdentity200ResponseOrg) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *GetIdentity200ResponseOrg) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPwdResetPersonalPhone

`func (o *GetIdentity200ResponseOrg) GetPwdResetPersonalPhone() bool`

GetPwdResetPersonalPhone returns the PwdResetPersonalPhone field if non-nil, zero value otherwise.

### GetPwdResetPersonalPhoneOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetPersonalPhoneOk() (*bool, bool)`

GetPwdResetPersonalPhoneOk returns a tuple with the PwdResetPersonalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetPersonalPhone

`func (o *GetIdentity200ResponseOrg) SetPwdResetPersonalPhone(v bool)`

SetPwdResetPersonalPhone sets PwdResetPersonalPhone field to given value.

### HasPwdResetPersonalPhone

`func (o *GetIdentity200ResponseOrg) HasPwdResetPersonalPhone() bool`

HasPwdResetPersonalPhone returns a boolean if a field has been set.

### GetPwdResetPersonalEmail

`func (o *GetIdentity200ResponseOrg) GetPwdResetPersonalEmail() bool`

GetPwdResetPersonalEmail returns the PwdResetPersonalEmail field if non-nil, zero value otherwise.

### GetPwdResetPersonalEmailOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetPersonalEmailOk() (*bool, bool)`

GetPwdResetPersonalEmailOk returns a tuple with the PwdResetPersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetPersonalEmail

`func (o *GetIdentity200ResponseOrg) SetPwdResetPersonalEmail(v bool)`

SetPwdResetPersonalEmail sets PwdResetPersonalEmail field to given value.

### HasPwdResetPersonalEmail

`func (o *GetIdentity200ResponseOrg) HasPwdResetPersonalEmail() bool`

HasPwdResetPersonalEmail returns a boolean if a field has been set.

### GetPwdResetKba

`func (o *GetIdentity200ResponseOrg) GetPwdResetKba() bool`

GetPwdResetKba returns the PwdResetKba field if non-nil, zero value otherwise.

### GetPwdResetKbaOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetKbaOk() (*bool, bool)`

GetPwdResetKbaOk returns a tuple with the PwdResetKba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetKba

`func (o *GetIdentity200ResponseOrg) SetPwdResetKba(v bool)`

SetPwdResetKba sets PwdResetKba field to given value.

### HasPwdResetKba

`func (o *GetIdentity200ResponseOrg) HasPwdResetKba() bool`

HasPwdResetKba returns a boolean if a field has been set.

### GetPwdResetEmail

`func (o *GetIdentity200ResponseOrg) GetPwdResetEmail() bool`

GetPwdResetEmail returns the PwdResetEmail field if non-nil, zero value otherwise.

### GetPwdResetEmailOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetEmailOk() (*bool, bool)`

GetPwdResetEmailOk returns a tuple with the PwdResetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetEmail

`func (o *GetIdentity200ResponseOrg) SetPwdResetEmail(v bool)`

SetPwdResetEmail sets PwdResetEmail field to given value.

### HasPwdResetEmail

`func (o *GetIdentity200ResponseOrg) HasPwdResetEmail() bool`

HasPwdResetEmail returns a boolean if a field has been set.

### GetPwdResetDuo

`func (o *GetIdentity200ResponseOrg) GetPwdResetDuo() bool`

GetPwdResetDuo returns the PwdResetDuo field if non-nil, zero value otherwise.

### GetPwdResetDuoOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetDuoOk() (*bool, bool)`

GetPwdResetDuoOk returns a tuple with the PwdResetDuo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetDuo

`func (o *GetIdentity200ResponseOrg) SetPwdResetDuo(v bool)`

SetPwdResetDuo sets PwdResetDuo field to given value.

### HasPwdResetDuo

`func (o *GetIdentity200ResponseOrg) HasPwdResetDuo() bool`

HasPwdResetDuo returns a boolean if a field has been set.

### GetPwdResetPhoneMask

`func (o *GetIdentity200ResponseOrg) GetPwdResetPhoneMask() bool`

GetPwdResetPhoneMask returns the PwdResetPhoneMask field if non-nil, zero value otherwise.

### GetPwdResetPhoneMaskOk

`func (o *GetIdentity200ResponseOrg) GetPwdResetPhoneMaskOk() (*bool, bool)`

GetPwdResetPhoneMaskOk returns a tuple with the PwdResetPhoneMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdResetPhoneMask

`func (o *GetIdentity200ResponseOrg) SetPwdResetPhoneMask(v bool)`

SetPwdResetPhoneMask sets PwdResetPhoneMask field to given value.

### HasPwdResetPhoneMask

`func (o *GetIdentity200ResponseOrg) HasPwdResetPhoneMask() bool`

HasPwdResetPhoneMask returns a boolean if a field has been set.

### GetAuthErrorText

`func (o *GetIdentity200ResponseOrg) GetAuthErrorText() map[string]interface{}`

GetAuthErrorText returns the AuthErrorText field if non-nil, zero value otherwise.

### GetAuthErrorTextOk

`func (o *GetIdentity200ResponseOrg) GetAuthErrorTextOk() (*map[string]interface{}, bool)`

GetAuthErrorTextOk returns a tuple with the AuthErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthErrorText

`func (o *GetIdentity200ResponseOrg) SetAuthErrorText(v map[string]interface{})`

SetAuthErrorText sets AuthErrorText field to given value.

### HasAuthErrorText

`func (o *GetIdentity200ResponseOrg) HasAuthErrorText() bool`

HasAuthErrorText returns a boolean if a field has been set.

### GetStrongAuthKba

`func (o *GetIdentity200ResponseOrg) GetStrongAuthKba() bool`

GetStrongAuthKba returns the StrongAuthKba field if non-nil, zero value otherwise.

### GetStrongAuthKbaOk

`func (o *GetIdentity200ResponseOrg) GetStrongAuthKbaOk() (*bool, bool)`

GetStrongAuthKbaOk returns a tuple with the StrongAuthKba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAuthKba

`func (o *GetIdentity200ResponseOrg) SetStrongAuthKba(v bool)`

SetStrongAuthKba sets StrongAuthKba field to given value.

### HasStrongAuthKba

`func (o *GetIdentity200ResponseOrg) HasStrongAuthKba() bool`

HasStrongAuthKba returns a boolean if a field has been set.

### GetStrongAuthPersonalPhone

`func (o *GetIdentity200ResponseOrg) GetStrongAuthPersonalPhone() bool`

GetStrongAuthPersonalPhone returns the StrongAuthPersonalPhone field if non-nil, zero value otherwise.

### GetStrongAuthPersonalPhoneOk

`func (o *GetIdentity200ResponseOrg) GetStrongAuthPersonalPhoneOk() (*bool, bool)`

GetStrongAuthPersonalPhoneOk returns a tuple with the StrongAuthPersonalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAuthPersonalPhone

`func (o *GetIdentity200ResponseOrg) SetStrongAuthPersonalPhone(v bool)`

SetStrongAuthPersonalPhone sets StrongAuthPersonalPhone field to given value.

### HasStrongAuthPersonalPhone

`func (o *GetIdentity200ResponseOrg) HasStrongAuthPersonalPhone() bool`

HasStrongAuthPersonalPhone returns a boolean if a field has been set.

### GetStrongAuthPersonalEmail

`func (o *GetIdentity200ResponseOrg) GetStrongAuthPersonalEmail() bool`

GetStrongAuthPersonalEmail returns the StrongAuthPersonalEmail field if non-nil, zero value otherwise.

### GetStrongAuthPersonalEmailOk

`func (o *GetIdentity200ResponseOrg) GetStrongAuthPersonalEmailOk() (*bool, bool)`

GetStrongAuthPersonalEmailOk returns a tuple with the StrongAuthPersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAuthPersonalEmail

`func (o *GetIdentity200ResponseOrg) SetStrongAuthPersonalEmail(v bool)`

SetStrongAuthPersonalEmail sets StrongAuthPersonalEmail field to given value.

### HasStrongAuthPersonalEmail

`func (o *GetIdentity200ResponseOrg) HasStrongAuthPersonalEmail() bool`

HasStrongAuthPersonalEmail returns a boolean if a field has been set.

### GetIntegrations

`func (o *GetIdentity200ResponseOrg) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *GetIdentity200ResponseOrg) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *GetIdentity200ResponseOrg) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *GetIdentity200ResponseOrg) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetProductName

`func (o *GetIdentity200ResponseOrg) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *GetIdentity200ResponseOrg) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *GetIdentity200ResponseOrg) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *GetIdentity200ResponseOrg) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetKbaReqForAuthn

`func (o *GetIdentity200ResponseOrg) GetKbaReqForAuthn() float32`

GetKbaReqForAuthn returns the KbaReqForAuthn field if non-nil, zero value otherwise.

### GetKbaReqForAuthnOk

`func (o *GetIdentity200ResponseOrg) GetKbaReqForAuthnOk() (*float32, bool)`

GetKbaReqForAuthnOk returns a tuple with the KbaReqForAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqForAuthn

`func (o *GetIdentity200ResponseOrg) SetKbaReqForAuthn(v float32)`

SetKbaReqForAuthn sets KbaReqForAuthn field to given value.

### HasKbaReqForAuthn

`func (o *GetIdentity200ResponseOrg) HasKbaReqForAuthn() bool`

HasKbaReqForAuthn returns a boolean if a field has been set.

### GetKbaReqAnswers

`func (o *GetIdentity200ResponseOrg) GetKbaReqAnswers() float32`

GetKbaReqAnswers returns the KbaReqAnswers field if non-nil, zero value otherwise.

### GetKbaReqAnswersOk

`func (o *GetIdentity200ResponseOrg) GetKbaReqAnswersOk() (*float32, bool)`

GetKbaReqAnswersOk returns a tuple with the KbaReqAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqAnswers

`func (o *GetIdentity200ResponseOrg) SetKbaReqAnswers(v float32)`

SetKbaReqAnswers sets KbaReqAnswers field to given value.

### HasKbaReqAnswers

`func (o *GetIdentity200ResponseOrg) HasKbaReqAnswers() bool`

HasKbaReqAnswers returns a boolean if a field has been set.

### GetLockoutAttemptThreshold

`func (o *GetIdentity200ResponseOrg) GetLockoutAttemptThreshold() float32`

GetLockoutAttemptThreshold returns the LockoutAttemptThreshold field if non-nil, zero value otherwise.

### GetLockoutAttemptThresholdOk

`func (o *GetIdentity200ResponseOrg) GetLockoutAttemptThresholdOk() (*float32, bool)`

GetLockoutAttemptThresholdOk returns a tuple with the LockoutAttemptThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutAttemptThreshold

`func (o *GetIdentity200ResponseOrg) SetLockoutAttemptThreshold(v float32)`

SetLockoutAttemptThreshold sets LockoutAttemptThreshold field to given value.

### HasLockoutAttemptThreshold

`func (o *GetIdentity200ResponseOrg) HasLockoutAttemptThreshold() bool`

HasLockoutAttemptThreshold returns a boolean if a field has been set.

### GetLockoutTimeMinutes

`func (o *GetIdentity200ResponseOrg) GetLockoutTimeMinutes() float32`

GetLockoutTimeMinutes returns the LockoutTimeMinutes field if non-nil, zero value otherwise.

### GetLockoutTimeMinutesOk

`func (o *GetIdentity200ResponseOrg) GetLockoutTimeMinutesOk() (*float32, bool)`

GetLockoutTimeMinutesOk returns a tuple with the LockoutTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeMinutes

`func (o *GetIdentity200ResponseOrg) SetLockoutTimeMinutes(v float32)`

SetLockoutTimeMinutes sets LockoutTimeMinutes field to given value.

### HasLockoutTimeMinutes

`func (o *GetIdentity200ResponseOrg) HasLockoutTimeMinutes() bool`

HasLockoutTimeMinutes returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *GetIdentity200ResponseOrg) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *GetIdentity200ResponseOrg) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *GetIdentity200ResponseOrg) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *GetIdentity200ResponseOrg) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetAdminStrongAuthRequired

`func (o *GetIdentity200ResponseOrg) GetAdminStrongAuthRequired() bool`

GetAdminStrongAuthRequired returns the AdminStrongAuthRequired field if non-nil, zero value otherwise.

### GetAdminStrongAuthRequiredOk

`func (o *GetIdentity200ResponseOrg) GetAdminStrongAuthRequiredOk() (*bool, bool)`

GetAdminStrongAuthRequiredOk returns a tuple with the AdminStrongAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStrongAuthRequired

`func (o *GetIdentity200ResponseOrg) SetAdminStrongAuthRequired(v bool)`

SetAdminStrongAuthRequired sets AdminStrongAuthRequired field to given value.

### HasAdminStrongAuthRequired

`func (o *GetIdentity200ResponseOrg) HasAdminStrongAuthRequired() bool`

HasAdminStrongAuthRequired returns a boolean if a field has been set.

### GetEnableExternalPasswordChange

`func (o *GetIdentity200ResponseOrg) GetEnableExternalPasswordChange() bool`

GetEnableExternalPasswordChange returns the EnableExternalPasswordChange field if non-nil, zero value otherwise.

### GetEnableExternalPasswordChangeOk

`func (o *GetIdentity200ResponseOrg) GetEnableExternalPasswordChangeOk() (*bool, bool)`

GetEnableExternalPasswordChangeOk returns a tuple with the EnableExternalPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalPasswordChange

`func (o *GetIdentity200ResponseOrg) SetEnableExternalPasswordChange(v bool)`

SetEnableExternalPasswordChange sets EnableExternalPasswordChange field to given value.

### HasEnableExternalPasswordChange

`func (o *GetIdentity200ResponseOrg) HasEnableExternalPasswordChange() bool`

HasEnableExternalPasswordChange returns a boolean if a field has been set.

### GetEnablePasswordReplay

`func (o *GetIdentity200ResponseOrg) GetEnablePasswordReplay() bool`

GetEnablePasswordReplay returns the EnablePasswordReplay field if non-nil, zero value otherwise.

### GetEnablePasswordReplayOk

`func (o *GetIdentity200ResponseOrg) GetEnablePasswordReplayOk() (*bool, bool)`

GetEnablePasswordReplayOk returns a tuple with the EnablePasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordReplay

`func (o *GetIdentity200ResponseOrg) SetEnablePasswordReplay(v bool)`

SetEnablePasswordReplay sets EnablePasswordReplay field to given value.

### HasEnablePasswordReplay

`func (o *GetIdentity200ResponseOrg) HasEnablePasswordReplay() bool`

HasEnablePasswordReplay returns a boolean if a field has been set.

### GetEnableAutomaticPasswordReplay

`func (o *GetIdentity200ResponseOrg) GetEnableAutomaticPasswordReplay() bool`

GetEnableAutomaticPasswordReplay returns the EnableAutomaticPasswordReplay field if non-nil, zero value otherwise.

### GetEnableAutomaticPasswordReplayOk

`func (o *GetIdentity200ResponseOrg) GetEnableAutomaticPasswordReplayOk() (*bool, bool)`

GetEnableAutomaticPasswordReplayOk returns a tuple with the EnableAutomaticPasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPasswordReplay

`func (o *GetIdentity200ResponseOrg) SetEnableAutomaticPasswordReplay(v bool)`

SetEnableAutomaticPasswordReplay sets EnableAutomaticPasswordReplay field to given value.

### HasEnableAutomaticPasswordReplay

`func (o *GetIdentity200ResponseOrg) HasEnableAutomaticPasswordReplay() bool`

HasEnableAutomaticPasswordReplay returns a boolean if a field has been set.

### GetNotifyAuthenticationSettingChange

`func (o *GetIdentity200ResponseOrg) GetNotifyAuthenticationSettingChange() bool`

GetNotifyAuthenticationSettingChange returns the NotifyAuthenticationSettingChange field if non-nil, zero value otherwise.

### GetNotifyAuthenticationSettingChangeOk

`func (o *GetIdentity200ResponseOrg) GetNotifyAuthenticationSettingChangeOk() (*bool, bool)`

GetNotifyAuthenticationSettingChangeOk returns a tuple with the NotifyAuthenticationSettingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAuthenticationSettingChange

`func (o *GetIdentity200ResponseOrg) SetNotifyAuthenticationSettingChange(v bool)`

SetNotifyAuthenticationSettingChange sets NotifyAuthenticationSettingChange field to given value.

### HasNotifyAuthenticationSettingChange

`func (o *GetIdentity200ResponseOrg) HasNotifyAuthenticationSettingChange() bool`

HasNotifyAuthenticationSettingChange returns a boolean if a field has been set.

### GetNetmasks

`func (o *GetIdentity200ResponseOrg) GetNetmasks() map[string]interface{}`

GetNetmasks returns the Netmasks field if non-nil, zero value otherwise.

### GetNetmasksOk

`func (o *GetIdentity200ResponseOrg) GetNetmasksOk() (*map[string]interface{}, bool)`

GetNetmasksOk returns a tuple with the Netmasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmasks

`func (o *GetIdentity200ResponseOrg) SetNetmasks(v map[string]interface{})`

SetNetmasks sets Netmasks field to given value.

### HasNetmasks

`func (o *GetIdentity200ResponseOrg) HasNetmasks() bool`

HasNetmasks returns a boolean if a field has been set.

### GetCountryCodes

`func (o *GetIdentity200ResponseOrg) GetCountryCodes() map[string]interface{}`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *GetIdentity200ResponseOrg) GetCountryCodesOk() (*map[string]interface{}, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *GetIdentity200ResponseOrg) SetCountryCodes(v map[string]interface{})`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *GetIdentity200ResponseOrg) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetWhiteList

`func (o *GetIdentity200ResponseOrg) GetWhiteList() bool`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *GetIdentity200ResponseOrg) GetWhiteListOk() (*bool, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *GetIdentity200ResponseOrg) SetWhiteList(v bool)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *GetIdentity200ResponseOrg) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetUsernameEmptyText

`func (o *GetIdentity200ResponseOrg) GetUsernameEmptyText() map[string]interface{}`

GetUsernameEmptyText returns the UsernameEmptyText field if non-nil, zero value otherwise.

### GetUsernameEmptyTextOk

`func (o *GetIdentity200ResponseOrg) GetUsernameEmptyTextOk() (*map[string]interface{}, bool)`

GetUsernameEmptyTextOk returns a tuple with the UsernameEmptyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameEmptyText

`func (o *GetIdentity200ResponseOrg) SetUsernameEmptyText(v map[string]interface{})`

SetUsernameEmptyText sets UsernameEmptyText field to given value.

### HasUsernameEmptyText

`func (o *GetIdentity200ResponseOrg) HasUsernameEmptyText() bool`

HasUsernameEmptyText returns a boolean if a field has been set.

### GetUsernameLabel

`func (o *GetIdentity200ResponseOrg) GetUsernameLabel() map[string]interface{}`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *GetIdentity200ResponseOrg) GetUsernameLabelOk() (*map[string]interface{}, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *GetIdentity200ResponseOrg) SetUsernameLabel(v map[string]interface{})`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *GetIdentity200ResponseOrg) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetEnableAutomationGeneration

`func (o *GetIdentity200ResponseOrg) GetEnableAutomationGeneration() bool`

GetEnableAutomationGeneration returns the EnableAutomationGeneration field if non-nil, zero value otherwise.

### GetEnableAutomationGenerationOk

`func (o *GetIdentity200ResponseOrg) GetEnableAutomationGenerationOk() (*bool, bool)`

GetEnableAutomationGenerationOk returns a tuple with the EnableAutomationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomationGeneration

`func (o *GetIdentity200ResponseOrg) SetEnableAutomationGeneration(v bool)`

SetEnableAutomationGeneration sets EnableAutomationGeneration field to given value.

### HasEnableAutomationGeneration

`func (o *GetIdentity200ResponseOrg) HasEnableAutomationGeneration() bool`

HasEnableAutomationGeneration returns a boolean if a field has been set.

### GetEmailTestMode

`func (o *GetIdentity200ResponseOrg) GetEmailTestMode() bool`

GetEmailTestMode returns the EmailTestMode field if non-nil, zero value otherwise.

### GetEmailTestModeOk

`func (o *GetIdentity200ResponseOrg) GetEmailTestModeOk() (*bool, bool)`

GetEmailTestModeOk returns a tuple with the EmailTestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestMode

`func (o *GetIdentity200ResponseOrg) SetEmailTestMode(v bool)`

SetEmailTestMode sets EmailTestMode field to given value.

### HasEmailTestMode

`func (o *GetIdentity200ResponseOrg) HasEmailTestMode() bool`

HasEmailTestMode returns a boolean if a field has been set.

### GetEmailTestAddress

`func (o *GetIdentity200ResponseOrg) GetEmailTestAddress() string`

GetEmailTestAddress returns the EmailTestAddress field if non-nil, zero value otherwise.

### GetEmailTestAddressOk

`func (o *GetIdentity200ResponseOrg) GetEmailTestAddressOk() (*string, bool)`

GetEmailTestAddressOk returns a tuple with the EmailTestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTestAddress

`func (o *GetIdentity200ResponseOrg) SetEmailTestAddress(v string)`

SetEmailTestAddress sets EmailTestAddress field to given value.

### HasEmailTestAddress

`func (o *GetIdentity200ResponseOrg) HasEmailTestAddress() bool`

HasEmailTestAddress returns a boolean if a field has been set.

### GetOrgType

`func (o *GetIdentity200ResponseOrg) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *GetIdentity200ResponseOrg) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *GetIdentity200ResponseOrg) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *GetIdentity200ResponseOrg) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### GetPasswordReplayState

`func (o *GetIdentity200ResponseOrg) GetPasswordReplayState() string`

GetPasswordReplayState returns the PasswordReplayState field if non-nil, zero value otherwise.

### GetPasswordReplayStateOk

`func (o *GetIdentity200ResponseOrg) GetPasswordReplayStateOk() (*string, bool)`

GetPasswordReplayStateOk returns a tuple with the PasswordReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReplayState

`func (o *GetIdentity200ResponseOrg) SetPasswordReplayState(v string)`

SetPasswordReplayState sets PasswordReplayState field to given value.

### HasPasswordReplayState

`func (o *GetIdentity200ResponseOrg) HasPasswordReplayState() bool`

HasPasswordReplayState returns a boolean if a field has been set.

### GetSystemNotificationConfig

`func (o *GetIdentity200ResponseOrg) GetSystemNotificationConfig() string`

GetSystemNotificationConfig returns the SystemNotificationConfig field if non-nil, zero value otherwise.

### GetSystemNotificationConfigOk

`func (o *GetIdentity200ResponseOrg) GetSystemNotificationConfigOk() (*string, bool)`

GetSystemNotificationConfigOk returns a tuple with the SystemNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationConfig

`func (o *GetIdentity200ResponseOrg) SetSystemNotificationConfig(v string)`

SetSystemNotificationConfig sets SystemNotificationConfig field to given value.

### HasSystemNotificationConfig

`func (o *GetIdentity200ResponseOrg) HasSystemNotificationConfig() bool`

HasSystemNotificationConfig returns a boolean if a field has been set.

### GetRedirectPatterns

`func (o *GetIdentity200ResponseOrg) GetRedirectPatterns() string`

GetRedirectPatterns returns the RedirectPatterns field if non-nil, zero value otherwise.

### GetRedirectPatternsOk

`func (o *GetIdentity200ResponseOrg) GetRedirectPatternsOk() (*string, bool)`

GetRedirectPatternsOk returns a tuple with the RedirectPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPatterns

`func (o *GetIdentity200ResponseOrg) SetRedirectPatterns(v string)`

SetRedirectPatterns sets RedirectPatterns field to given value.

### HasRedirectPatterns

`func (o *GetIdentity200ResponseOrg) HasRedirectPatterns() bool`

HasRedirectPatterns returns a boolean if a field has been set.

### GetMaxClusterDebugHours

`func (o *GetIdentity200ResponseOrg) GetMaxClusterDebugHours() string`

GetMaxClusterDebugHours returns the MaxClusterDebugHours field if non-nil, zero value otherwise.

### GetMaxClusterDebugHoursOk

`func (o *GetIdentity200ResponseOrg) GetMaxClusterDebugHoursOk() (*string, bool)`

GetMaxClusterDebugHoursOk returns a tuple with the MaxClusterDebugHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterDebugHours

`func (o *GetIdentity200ResponseOrg) SetMaxClusterDebugHours(v string)`

SetMaxClusterDebugHours sets MaxClusterDebugHours field to given value.

### HasMaxClusterDebugHours

`func (o *GetIdentity200ResponseOrg) HasMaxClusterDebugHours() bool`

HasMaxClusterDebugHours returns a boolean if a field has been set.

### GetBrandName

`func (o *GetIdentity200ResponseOrg) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *GetIdentity200ResponseOrg) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *GetIdentity200ResponseOrg) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *GetIdentity200ResponseOrg) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetLogo

`func (o *GetIdentity200ResponseOrg) GetLogo() map[string]interface{}`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *GetIdentity200ResponseOrg) GetLogoOk() (*map[string]interface{}, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *GetIdentity200ResponseOrg) SetLogo(v map[string]interface{})`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *GetIdentity200ResponseOrg) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetEmailFromAddress

`func (o *GetIdentity200ResponseOrg) GetEmailFromAddress() map[string]interface{}`

GetEmailFromAddress returns the EmailFromAddress field if non-nil, zero value otherwise.

### GetEmailFromAddressOk

`func (o *GetIdentity200ResponseOrg) GetEmailFromAddressOk() (*map[string]interface{}, bool)`

GetEmailFromAddressOk returns a tuple with the EmailFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromAddress

`func (o *GetIdentity200ResponseOrg) SetEmailFromAddress(v map[string]interface{})`

SetEmailFromAddress sets EmailFromAddress field to given value.

### HasEmailFromAddress

`func (o *GetIdentity200ResponseOrg) HasEmailFromAddress() bool`

HasEmailFromAddress returns a boolean if a field has been set.

### GetStandardLogoUrl

`func (o *GetIdentity200ResponseOrg) GetStandardLogoUrl() map[string]interface{}`

GetStandardLogoUrl returns the StandardLogoUrl field if non-nil, zero value otherwise.

### GetStandardLogoUrlOk

`func (o *GetIdentity200ResponseOrg) GetStandardLogoUrlOk() (*map[string]interface{}, bool)`

GetStandardLogoUrlOk returns a tuple with the StandardLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLogoUrl

`func (o *GetIdentity200ResponseOrg) SetStandardLogoUrl(v map[string]interface{})`

SetStandardLogoUrl sets StandardLogoUrl field to given value.

### HasStandardLogoUrl

`func (o *GetIdentity200ResponseOrg) HasStandardLogoUrl() bool`

HasStandardLogoUrl returns a boolean if a field has been set.

### GetNarrowLogoUrl

`func (o *GetIdentity200ResponseOrg) GetNarrowLogoUrl() map[string]interface{}`

GetNarrowLogoUrl returns the NarrowLogoUrl field if non-nil, zero value otherwise.

### GetNarrowLogoUrlOk

`func (o *GetIdentity200ResponseOrg) GetNarrowLogoUrlOk() (*map[string]interface{}, bool)`

GetNarrowLogoUrlOk returns a tuple with the NarrowLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrowLogoUrl

`func (o *GetIdentity200ResponseOrg) SetNarrowLogoUrl(v map[string]interface{})`

SetNarrowLogoUrl sets NarrowLogoUrl field to given value.

### HasNarrowLogoUrl

`func (o *GetIdentity200ResponseOrg) HasNarrowLogoUrl() bool`

HasNarrowLogoUrl returns a boolean if a field has been set.

### GetActionButtonColor

`func (o *GetIdentity200ResponseOrg) GetActionButtonColor() string`

GetActionButtonColor returns the ActionButtonColor field if non-nil, zero value otherwise.

### GetActionButtonColorOk

`func (o *GetIdentity200ResponseOrg) GetActionButtonColorOk() (*string, bool)`

GetActionButtonColorOk returns a tuple with the ActionButtonColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionButtonColor

`func (o *GetIdentity200ResponseOrg) SetActionButtonColor(v string)`

SetActionButtonColor sets ActionButtonColor field to given value.

### HasActionButtonColor

`func (o *GetIdentity200ResponseOrg) HasActionButtonColor() bool`

HasActionButtonColor returns a boolean if a field has been set.

### GetActiveLinkColor

`func (o *GetIdentity200ResponseOrg) GetActiveLinkColor() string`

GetActiveLinkColor returns the ActiveLinkColor field if non-nil, zero value otherwise.

### GetActiveLinkColorOk

`func (o *GetIdentity200ResponseOrg) GetActiveLinkColorOk() (*string, bool)`

GetActiveLinkColorOk returns a tuple with the ActiveLinkColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLinkColor

`func (o *GetIdentity200ResponseOrg) SetActiveLinkColor(v string)`

SetActiveLinkColor sets ActiveLinkColor field to given value.

### HasActiveLinkColor

`func (o *GetIdentity200ResponseOrg) HasActiveLinkColor() bool`

HasActiveLinkColor returns a boolean if a field has been set.

### GetNavigationColor

`func (o *GetIdentity200ResponseOrg) GetNavigationColor() string`

GetNavigationColor returns the NavigationColor field if non-nil, zero value otherwise.

### GetNavigationColorOk

`func (o *GetIdentity200ResponseOrg) GetNavigationColorOk() (*string, bool)`

GetNavigationColorOk returns a tuple with the NavigationColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationColor

`func (o *GetIdentity200ResponseOrg) SetNavigationColor(v string)`

SetNavigationColor sets NavigationColor field to given value.

### HasNavigationColor

`func (o *GetIdentity200ResponseOrg) HasNavigationColor() bool`

HasNavigationColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


