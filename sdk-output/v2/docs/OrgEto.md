# OrgEto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**EnableExternalPasswordChange** | Pointer to **bool** |  | [optional] 
**EnableAutomaticPasswordReplay** | Pointer to **bool** |  | [optional] 
**EnableAutomationGeneration** | Pointer to **bool** |  | [optional] 
**KbaReqAnswers** | Pointer to **int32** |  | [optional] 
**KbaReqForAuthn** | Pointer to **int32** |  | [optional] 
**LockoutAttemptThreshold** | Pointer to **int32** |  | [optional] 
**LockoutTimeMinutes** | Pointer to **int32** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 
**Netmasks** | Pointer to **[]string** |  | [optional] 
**NotifyAuthenticationSettingChange** | Pointer to **bool** |  | [optional] 
**PasswordReplayState** | Pointer to **string** |  | [optional] 
**PreferredIdentityInviteTemplate** | Pointer to **string** |  | [optional] 
**RedirectPatterns** | Pointer to **[]string** |  | [optional] 
**SsoPartnerSource** | Pointer to **string** |  | [optional] 
**SystemNotificationEmails** | Pointer to **[]string** |  | [optional] 
**TrackAnalytics** | Pointer to **bool** |  | [optional] 
**UsageCertRequired** | Pointer to **bool** |  | [optional] 
**UsageCertText** | Pointer to **string** |  | [optional] 
**UsernameEmptyText** | Pointer to **string** |  | [optional] 
**UsernameLabel** | Pointer to **string** |  | [optional] 
**WhiteList** | Pointer to **bool** |  | [optional] 
**ApprovalConfig** | Pointer to [**ApprovalConfigEto**](ApprovalConfigEto.md) |  | [optional] 

## Methods

### NewOrgEto

`func NewOrgEto() *OrgEto`

NewOrgEto instantiates a new OrgEto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgEtoWithDefaults

`func NewOrgEtoWithDefaults() *OrgEto`

NewOrgEtoWithDefaults instantiates a new OrgEto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *OrgEto) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *OrgEto) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *OrgEto) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *OrgEto) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetEnableExternalPasswordChange

`func (o *OrgEto) GetEnableExternalPasswordChange() bool`

GetEnableExternalPasswordChange returns the EnableExternalPasswordChange field if non-nil, zero value otherwise.

### GetEnableExternalPasswordChangeOk

`func (o *OrgEto) GetEnableExternalPasswordChangeOk() (*bool, bool)`

GetEnableExternalPasswordChangeOk returns a tuple with the EnableExternalPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalPasswordChange

`func (o *OrgEto) SetEnableExternalPasswordChange(v bool)`

SetEnableExternalPasswordChange sets EnableExternalPasswordChange field to given value.

### HasEnableExternalPasswordChange

`func (o *OrgEto) HasEnableExternalPasswordChange() bool`

HasEnableExternalPasswordChange returns a boolean if a field has been set.

### GetEnableAutomaticPasswordReplay

`func (o *OrgEto) GetEnableAutomaticPasswordReplay() bool`

GetEnableAutomaticPasswordReplay returns the EnableAutomaticPasswordReplay field if non-nil, zero value otherwise.

### GetEnableAutomaticPasswordReplayOk

`func (o *OrgEto) GetEnableAutomaticPasswordReplayOk() (*bool, bool)`

GetEnableAutomaticPasswordReplayOk returns a tuple with the EnableAutomaticPasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPasswordReplay

`func (o *OrgEto) SetEnableAutomaticPasswordReplay(v bool)`

SetEnableAutomaticPasswordReplay sets EnableAutomaticPasswordReplay field to given value.

### HasEnableAutomaticPasswordReplay

`func (o *OrgEto) HasEnableAutomaticPasswordReplay() bool`

HasEnableAutomaticPasswordReplay returns a boolean if a field has been set.

### GetEnableAutomationGeneration

`func (o *OrgEto) GetEnableAutomationGeneration() bool`

GetEnableAutomationGeneration returns the EnableAutomationGeneration field if non-nil, zero value otherwise.

### GetEnableAutomationGenerationOk

`func (o *OrgEto) GetEnableAutomationGenerationOk() (*bool, bool)`

GetEnableAutomationGenerationOk returns a tuple with the EnableAutomationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomationGeneration

`func (o *OrgEto) SetEnableAutomationGeneration(v bool)`

SetEnableAutomationGeneration sets EnableAutomationGeneration field to given value.

### HasEnableAutomationGeneration

`func (o *OrgEto) HasEnableAutomationGeneration() bool`

HasEnableAutomationGeneration returns a boolean if a field has been set.

### GetKbaReqAnswers

`func (o *OrgEto) GetKbaReqAnswers() int32`

GetKbaReqAnswers returns the KbaReqAnswers field if non-nil, zero value otherwise.

### GetKbaReqAnswersOk

`func (o *OrgEto) GetKbaReqAnswersOk() (*int32, bool)`

GetKbaReqAnswersOk returns a tuple with the KbaReqAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqAnswers

`func (o *OrgEto) SetKbaReqAnswers(v int32)`

SetKbaReqAnswers sets KbaReqAnswers field to given value.

### HasKbaReqAnswers

`func (o *OrgEto) HasKbaReqAnswers() bool`

HasKbaReqAnswers returns a boolean if a field has been set.

### GetKbaReqForAuthn

`func (o *OrgEto) GetKbaReqForAuthn() int32`

GetKbaReqForAuthn returns the KbaReqForAuthn field if non-nil, zero value otherwise.

### GetKbaReqForAuthnOk

`func (o *OrgEto) GetKbaReqForAuthnOk() (*int32, bool)`

GetKbaReqForAuthnOk returns a tuple with the KbaReqForAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqForAuthn

`func (o *OrgEto) SetKbaReqForAuthn(v int32)`

SetKbaReqForAuthn sets KbaReqForAuthn field to given value.

### HasKbaReqForAuthn

`func (o *OrgEto) HasKbaReqForAuthn() bool`

HasKbaReqForAuthn returns a boolean if a field has been set.

### GetLockoutAttemptThreshold

`func (o *OrgEto) GetLockoutAttemptThreshold() int32`

GetLockoutAttemptThreshold returns the LockoutAttemptThreshold field if non-nil, zero value otherwise.

### GetLockoutAttemptThresholdOk

`func (o *OrgEto) GetLockoutAttemptThresholdOk() (*int32, bool)`

GetLockoutAttemptThresholdOk returns a tuple with the LockoutAttemptThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutAttemptThreshold

`func (o *OrgEto) SetLockoutAttemptThreshold(v int32)`

SetLockoutAttemptThreshold sets LockoutAttemptThreshold field to given value.

### HasLockoutAttemptThreshold

`func (o *OrgEto) HasLockoutAttemptThreshold() bool`

HasLockoutAttemptThreshold returns a boolean if a field has been set.

### GetLockoutTimeMinutes

`func (o *OrgEto) GetLockoutTimeMinutes() int32`

GetLockoutTimeMinutes returns the LockoutTimeMinutes field if non-nil, zero value otherwise.

### GetLockoutTimeMinutesOk

`func (o *OrgEto) GetLockoutTimeMinutesOk() (*int32, bool)`

GetLockoutTimeMinutesOk returns a tuple with the LockoutTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeMinutes

`func (o *OrgEto) SetLockoutTimeMinutes(v int32)`

SetLockoutTimeMinutes sets LockoutTimeMinutes field to given value.

### HasLockoutTimeMinutes

`func (o *OrgEto) HasLockoutTimeMinutes() bool`

HasLockoutTimeMinutes returns a boolean if a field has been set.

### GetLoginUrl

`func (o *OrgEto) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *OrgEto) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *OrgEto) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *OrgEto) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetNetmasks

`func (o *OrgEto) GetNetmasks() []string`

GetNetmasks returns the Netmasks field if non-nil, zero value otherwise.

### GetNetmasksOk

`func (o *OrgEto) GetNetmasksOk() (*[]string, bool)`

GetNetmasksOk returns a tuple with the Netmasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmasks

`func (o *OrgEto) SetNetmasks(v []string)`

SetNetmasks sets Netmasks field to given value.

### HasNetmasks

`func (o *OrgEto) HasNetmasks() bool`

HasNetmasks returns a boolean if a field has been set.

### GetNotifyAuthenticationSettingChange

`func (o *OrgEto) GetNotifyAuthenticationSettingChange() bool`

GetNotifyAuthenticationSettingChange returns the NotifyAuthenticationSettingChange field if non-nil, zero value otherwise.

### GetNotifyAuthenticationSettingChangeOk

`func (o *OrgEto) GetNotifyAuthenticationSettingChangeOk() (*bool, bool)`

GetNotifyAuthenticationSettingChangeOk returns a tuple with the NotifyAuthenticationSettingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAuthenticationSettingChange

`func (o *OrgEto) SetNotifyAuthenticationSettingChange(v bool)`

SetNotifyAuthenticationSettingChange sets NotifyAuthenticationSettingChange field to given value.

### HasNotifyAuthenticationSettingChange

`func (o *OrgEto) HasNotifyAuthenticationSettingChange() bool`

HasNotifyAuthenticationSettingChange returns a boolean if a field has been set.

### GetPasswordReplayState

`func (o *OrgEto) GetPasswordReplayState() string`

GetPasswordReplayState returns the PasswordReplayState field if non-nil, zero value otherwise.

### GetPasswordReplayStateOk

`func (o *OrgEto) GetPasswordReplayStateOk() (*string, bool)`

GetPasswordReplayStateOk returns a tuple with the PasswordReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReplayState

`func (o *OrgEto) SetPasswordReplayState(v string)`

SetPasswordReplayState sets PasswordReplayState field to given value.

### HasPasswordReplayState

`func (o *OrgEto) HasPasswordReplayState() bool`

HasPasswordReplayState returns a boolean if a field has been set.

### GetPreferredIdentityInviteTemplate

`func (o *OrgEto) GetPreferredIdentityInviteTemplate() string`

GetPreferredIdentityInviteTemplate returns the PreferredIdentityInviteTemplate field if non-nil, zero value otherwise.

### GetPreferredIdentityInviteTemplateOk

`func (o *OrgEto) GetPreferredIdentityInviteTemplateOk() (*string, bool)`

GetPreferredIdentityInviteTemplateOk returns a tuple with the PreferredIdentityInviteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIdentityInviteTemplate

`func (o *OrgEto) SetPreferredIdentityInviteTemplate(v string)`

SetPreferredIdentityInviteTemplate sets PreferredIdentityInviteTemplate field to given value.

### HasPreferredIdentityInviteTemplate

`func (o *OrgEto) HasPreferredIdentityInviteTemplate() bool`

HasPreferredIdentityInviteTemplate returns a boolean if a field has been set.

### GetRedirectPatterns

`func (o *OrgEto) GetRedirectPatterns() []string`

GetRedirectPatterns returns the RedirectPatterns field if non-nil, zero value otherwise.

### GetRedirectPatternsOk

`func (o *OrgEto) GetRedirectPatternsOk() (*[]string, bool)`

GetRedirectPatternsOk returns a tuple with the RedirectPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPatterns

`func (o *OrgEto) SetRedirectPatterns(v []string)`

SetRedirectPatterns sets RedirectPatterns field to given value.

### HasRedirectPatterns

`func (o *OrgEto) HasRedirectPatterns() bool`

HasRedirectPatterns returns a boolean if a field has been set.

### GetSsoPartnerSource

`func (o *OrgEto) GetSsoPartnerSource() string`

GetSsoPartnerSource returns the SsoPartnerSource field if non-nil, zero value otherwise.

### GetSsoPartnerSourceOk

`func (o *OrgEto) GetSsoPartnerSourceOk() (*string, bool)`

GetSsoPartnerSourceOk returns a tuple with the SsoPartnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoPartnerSource

`func (o *OrgEto) SetSsoPartnerSource(v string)`

SetSsoPartnerSource sets SsoPartnerSource field to given value.

### HasSsoPartnerSource

`func (o *OrgEto) HasSsoPartnerSource() bool`

HasSsoPartnerSource returns a boolean if a field has been set.

### GetSystemNotificationEmails

`func (o *OrgEto) GetSystemNotificationEmails() []string`

GetSystemNotificationEmails returns the SystemNotificationEmails field if non-nil, zero value otherwise.

### GetSystemNotificationEmailsOk

`func (o *OrgEto) GetSystemNotificationEmailsOk() (*[]string, bool)`

GetSystemNotificationEmailsOk returns a tuple with the SystemNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationEmails

`func (o *OrgEto) SetSystemNotificationEmails(v []string)`

SetSystemNotificationEmails sets SystemNotificationEmails field to given value.

### HasSystemNotificationEmails

`func (o *OrgEto) HasSystemNotificationEmails() bool`

HasSystemNotificationEmails returns a boolean if a field has been set.

### GetTrackAnalytics

`func (o *OrgEto) GetTrackAnalytics() bool`

GetTrackAnalytics returns the TrackAnalytics field if non-nil, zero value otherwise.

### GetTrackAnalyticsOk

`func (o *OrgEto) GetTrackAnalyticsOk() (*bool, bool)`

GetTrackAnalyticsOk returns a tuple with the TrackAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAnalytics

`func (o *OrgEto) SetTrackAnalytics(v bool)`

SetTrackAnalytics sets TrackAnalytics field to given value.

### HasTrackAnalytics

`func (o *OrgEto) HasTrackAnalytics() bool`

HasTrackAnalytics returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *OrgEto) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *OrgEto) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *OrgEto) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *OrgEto) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetUsageCertText

`func (o *OrgEto) GetUsageCertText() string`

GetUsageCertText returns the UsageCertText field if non-nil, zero value otherwise.

### GetUsageCertTextOk

`func (o *OrgEto) GetUsageCertTextOk() (*string, bool)`

GetUsageCertTextOk returns a tuple with the UsageCertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertText

`func (o *OrgEto) SetUsageCertText(v string)`

SetUsageCertText sets UsageCertText field to given value.

### HasUsageCertText

`func (o *OrgEto) HasUsageCertText() bool`

HasUsageCertText returns a boolean if a field has been set.

### GetUsernameEmptyText

`func (o *OrgEto) GetUsernameEmptyText() string`

GetUsernameEmptyText returns the UsernameEmptyText field if non-nil, zero value otherwise.

### GetUsernameEmptyTextOk

`func (o *OrgEto) GetUsernameEmptyTextOk() (*string, bool)`

GetUsernameEmptyTextOk returns a tuple with the UsernameEmptyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameEmptyText

`func (o *OrgEto) SetUsernameEmptyText(v string)`

SetUsernameEmptyText sets UsernameEmptyText field to given value.

### HasUsernameEmptyText

`func (o *OrgEto) HasUsernameEmptyText() bool`

HasUsernameEmptyText returns a boolean if a field has been set.

### GetUsernameLabel

`func (o *OrgEto) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *OrgEto) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *OrgEto) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *OrgEto) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetWhiteList

`func (o *OrgEto) GetWhiteList() bool`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *OrgEto) GetWhiteListOk() (*bool, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *OrgEto) SetWhiteList(v bool)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *OrgEto) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *OrgEto) GetApprovalConfig() ApprovalConfigEto`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *OrgEto) GetApprovalConfigOk() (*ApprovalConfigEto, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *OrgEto) SetApprovalConfig(v ApprovalConfigEto)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *OrgEto) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


