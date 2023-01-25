# UpdateOrgSettingsRequest

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
**ApprovalConfig** | Pointer to [**GetOrgSettings200ResponseApprovalConfig**](GetOrgSettings200ResponseApprovalConfig.md) |  | [optional] 

## Methods

### NewUpdateOrgSettingsRequest

`func NewUpdateOrgSettingsRequest() *UpdateOrgSettingsRequest`

NewUpdateOrgSettingsRequest instantiates a new UpdateOrgSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgSettingsRequestWithDefaults

`func NewUpdateOrgSettingsRequestWithDefaults() *UpdateOrgSettingsRequest`

NewUpdateOrgSettingsRequestWithDefaults instantiates a new UpdateOrgSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *UpdateOrgSettingsRequest) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *UpdateOrgSettingsRequest) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *UpdateOrgSettingsRequest) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *UpdateOrgSettingsRequest) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetEnableExternalPasswordChange

`func (o *UpdateOrgSettingsRequest) GetEnableExternalPasswordChange() bool`

GetEnableExternalPasswordChange returns the EnableExternalPasswordChange field if non-nil, zero value otherwise.

### GetEnableExternalPasswordChangeOk

`func (o *UpdateOrgSettingsRequest) GetEnableExternalPasswordChangeOk() (*bool, bool)`

GetEnableExternalPasswordChangeOk returns a tuple with the EnableExternalPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalPasswordChange

`func (o *UpdateOrgSettingsRequest) SetEnableExternalPasswordChange(v bool)`

SetEnableExternalPasswordChange sets EnableExternalPasswordChange field to given value.

### HasEnableExternalPasswordChange

`func (o *UpdateOrgSettingsRequest) HasEnableExternalPasswordChange() bool`

HasEnableExternalPasswordChange returns a boolean if a field has been set.

### GetEnableAutomaticPasswordReplay

`func (o *UpdateOrgSettingsRequest) GetEnableAutomaticPasswordReplay() bool`

GetEnableAutomaticPasswordReplay returns the EnableAutomaticPasswordReplay field if non-nil, zero value otherwise.

### GetEnableAutomaticPasswordReplayOk

`func (o *UpdateOrgSettingsRequest) GetEnableAutomaticPasswordReplayOk() (*bool, bool)`

GetEnableAutomaticPasswordReplayOk returns a tuple with the EnableAutomaticPasswordReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPasswordReplay

`func (o *UpdateOrgSettingsRequest) SetEnableAutomaticPasswordReplay(v bool)`

SetEnableAutomaticPasswordReplay sets EnableAutomaticPasswordReplay field to given value.

### HasEnableAutomaticPasswordReplay

`func (o *UpdateOrgSettingsRequest) HasEnableAutomaticPasswordReplay() bool`

HasEnableAutomaticPasswordReplay returns a boolean if a field has been set.

### GetEnableAutomationGeneration

`func (o *UpdateOrgSettingsRequest) GetEnableAutomationGeneration() bool`

GetEnableAutomationGeneration returns the EnableAutomationGeneration field if non-nil, zero value otherwise.

### GetEnableAutomationGenerationOk

`func (o *UpdateOrgSettingsRequest) GetEnableAutomationGenerationOk() (*bool, bool)`

GetEnableAutomationGenerationOk returns a tuple with the EnableAutomationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomationGeneration

`func (o *UpdateOrgSettingsRequest) SetEnableAutomationGeneration(v bool)`

SetEnableAutomationGeneration sets EnableAutomationGeneration field to given value.

### HasEnableAutomationGeneration

`func (o *UpdateOrgSettingsRequest) HasEnableAutomationGeneration() bool`

HasEnableAutomationGeneration returns a boolean if a field has been set.

### GetKbaReqAnswers

`func (o *UpdateOrgSettingsRequest) GetKbaReqAnswers() int32`

GetKbaReqAnswers returns the KbaReqAnswers field if non-nil, zero value otherwise.

### GetKbaReqAnswersOk

`func (o *UpdateOrgSettingsRequest) GetKbaReqAnswersOk() (*int32, bool)`

GetKbaReqAnswersOk returns a tuple with the KbaReqAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqAnswers

`func (o *UpdateOrgSettingsRequest) SetKbaReqAnswers(v int32)`

SetKbaReqAnswers sets KbaReqAnswers field to given value.

### HasKbaReqAnswers

`func (o *UpdateOrgSettingsRequest) HasKbaReqAnswers() bool`

HasKbaReqAnswers returns a boolean if a field has been set.

### GetKbaReqForAuthn

`func (o *UpdateOrgSettingsRequest) GetKbaReqForAuthn() int32`

GetKbaReqForAuthn returns the KbaReqForAuthn field if non-nil, zero value otherwise.

### GetKbaReqForAuthnOk

`func (o *UpdateOrgSettingsRequest) GetKbaReqForAuthnOk() (*int32, bool)`

GetKbaReqForAuthnOk returns a tuple with the KbaReqForAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaReqForAuthn

`func (o *UpdateOrgSettingsRequest) SetKbaReqForAuthn(v int32)`

SetKbaReqForAuthn sets KbaReqForAuthn field to given value.

### HasKbaReqForAuthn

`func (o *UpdateOrgSettingsRequest) HasKbaReqForAuthn() bool`

HasKbaReqForAuthn returns a boolean if a field has been set.

### GetLockoutAttemptThreshold

`func (o *UpdateOrgSettingsRequest) GetLockoutAttemptThreshold() int32`

GetLockoutAttemptThreshold returns the LockoutAttemptThreshold field if non-nil, zero value otherwise.

### GetLockoutAttemptThresholdOk

`func (o *UpdateOrgSettingsRequest) GetLockoutAttemptThresholdOk() (*int32, bool)`

GetLockoutAttemptThresholdOk returns a tuple with the LockoutAttemptThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutAttemptThreshold

`func (o *UpdateOrgSettingsRequest) SetLockoutAttemptThreshold(v int32)`

SetLockoutAttemptThreshold sets LockoutAttemptThreshold field to given value.

### HasLockoutAttemptThreshold

`func (o *UpdateOrgSettingsRequest) HasLockoutAttemptThreshold() bool`

HasLockoutAttemptThreshold returns a boolean if a field has been set.

### GetLockoutTimeMinutes

`func (o *UpdateOrgSettingsRequest) GetLockoutTimeMinutes() int32`

GetLockoutTimeMinutes returns the LockoutTimeMinutes field if non-nil, zero value otherwise.

### GetLockoutTimeMinutesOk

`func (o *UpdateOrgSettingsRequest) GetLockoutTimeMinutesOk() (*int32, bool)`

GetLockoutTimeMinutesOk returns a tuple with the LockoutTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeMinutes

`func (o *UpdateOrgSettingsRequest) SetLockoutTimeMinutes(v int32)`

SetLockoutTimeMinutes sets LockoutTimeMinutes field to given value.

### HasLockoutTimeMinutes

`func (o *UpdateOrgSettingsRequest) HasLockoutTimeMinutes() bool`

HasLockoutTimeMinutes returns a boolean if a field has been set.

### GetLoginUrl

`func (o *UpdateOrgSettingsRequest) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *UpdateOrgSettingsRequest) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *UpdateOrgSettingsRequest) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *UpdateOrgSettingsRequest) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetNetmasks

`func (o *UpdateOrgSettingsRequest) GetNetmasks() []string`

GetNetmasks returns the Netmasks field if non-nil, zero value otherwise.

### GetNetmasksOk

`func (o *UpdateOrgSettingsRequest) GetNetmasksOk() (*[]string, bool)`

GetNetmasksOk returns a tuple with the Netmasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmasks

`func (o *UpdateOrgSettingsRequest) SetNetmasks(v []string)`

SetNetmasks sets Netmasks field to given value.

### HasNetmasks

`func (o *UpdateOrgSettingsRequest) HasNetmasks() bool`

HasNetmasks returns a boolean if a field has been set.

### GetNotifyAuthenticationSettingChange

`func (o *UpdateOrgSettingsRequest) GetNotifyAuthenticationSettingChange() bool`

GetNotifyAuthenticationSettingChange returns the NotifyAuthenticationSettingChange field if non-nil, zero value otherwise.

### GetNotifyAuthenticationSettingChangeOk

`func (o *UpdateOrgSettingsRequest) GetNotifyAuthenticationSettingChangeOk() (*bool, bool)`

GetNotifyAuthenticationSettingChangeOk returns a tuple with the NotifyAuthenticationSettingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAuthenticationSettingChange

`func (o *UpdateOrgSettingsRequest) SetNotifyAuthenticationSettingChange(v bool)`

SetNotifyAuthenticationSettingChange sets NotifyAuthenticationSettingChange field to given value.

### HasNotifyAuthenticationSettingChange

`func (o *UpdateOrgSettingsRequest) HasNotifyAuthenticationSettingChange() bool`

HasNotifyAuthenticationSettingChange returns a boolean if a field has been set.

### GetPasswordReplayState

`func (o *UpdateOrgSettingsRequest) GetPasswordReplayState() string`

GetPasswordReplayState returns the PasswordReplayState field if non-nil, zero value otherwise.

### GetPasswordReplayStateOk

`func (o *UpdateOrgSettingsRequest) GetPasswordReplayStateOk() (*string, bool)`

GetPasswordReplayStateOk returns a tuple with the PasswordReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReplayState

`func (o *UpdateOrgSettingsRequest) SetPasswordReplayState(v string)`

SetPasswordReplayState sets PasswordReplayState field to given value.

### HasPasswordReplayState

`func (o *UpdateOrgSettingsRequest) HasPasswordReplayState() bool`

HasPasswordReplayState returns a boolean if a field has been set.

### GetPreferredIdentityInviteTemplate

`func (o *UpdateOrgSettingsRequest) GetPreferredIdentityInviteTemplate() string`

GetPreferredIdentityInviteTemplate returns the PreferredIdentityInviteTemplate field if non-nil, zero value otherwise.

### GetPreferredIdentityInviteTemplateOk

`func (o *UpdateOrgSettingsRequest) GetPreferredIdentityInviteTemplateOk() (*string, bool)`

GetPreferredIdentityInviteTemplateOk returns a tuple with the PreferredIdentityInviteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIdentityInviteTemplate

`func (o *UpdateOrgSettingsRequest) SetPreferredIdentityInviteTemplate(v string)`

SetPreferredIdentityInviteTemplate sets PreferredIdentityInviteTemplate field to given value.

### HasPreferredIdentityInviteTemplate

`func (o *UpdateOrgSettingsRequest) HasPreferredIdentityInviteTemplate() bool`

HasPreferredIdentityInviteTemplate returns a boolean if a field has been set.

### GetRedirectPatterns

`func (o *UpdateOrgSettingsRequest) GetRedirectPatterns() []string`

GetRedirectPatterns returns the RedirectPatterns field if non-nil, zero value otherwise.

### GetRedirectPatternsOk

`func (o *UpdateOrgSettingsRequest) GetRedirectPatternsOk() (*[]string, bool)`

GetRedirectPatternsOk returns a tuple with the RedirectPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPatterns

`func (o *UpdateOrgSettingsRequest) SetRedirectPatterns(v []string)`

SetRedirectPatterns sets RedirectPatterns field to given value.

### HasRedirectPatterns

`func (o *UpdateOrgSettingsRequest) HasRedirectPatterns() bool`

HasRedirectPatterns returns a boolean if a field has been set.

### GetSsoPartnerSource

`func (o *UpdateOrgSettingsRequest) GetSsoPartnerSource() string`

GetSsoPartnerSource returns the SsoPartnerSource field if non-nil, zero value otherwise.

### GetSsoPartnerSourceOk

`func (o *UpdateOrgSettingsRequest) GetSsoPartnerSourceOk() (*string, bool)`

GetSsoPartnerSourceOk returns a tuple with the SsoPartnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoPartnerSource

`func (o *UpdateOrgSettingsRequest) SetSsoPartnerSource(v string)`

SetSsoPartnerSource sets SsoPartnerSource field to given value.

### HasSsoPartnerSource

`func (o *UpdateOrgSettingsRequest) HasSsoPartnerSource() bool`

HasSsoPartnerSource returns a boolean if a field has been set.

### GetSystemNotificationEmails

`func (o *UpdateOrgSettingsRequest) GetSystemNotificationEmails() []string`

GetSystemNotificationEmails returns the SystemNotificationEmails field if non-nil, zero value otherwise.

### GetSystemNotificationEmailsOk

`func (o *UpdateOrgSettingsRequest) GetSystemNotificationEmailsOk() (*[]string, bool)`

GetSystemNotificationEmailsOk returns a tuple with the SystemNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemNotificationEmails

`func (o *UpdateOrgSettingsRequest) SetSystemNotificationEmails(v []string)`

SetSystemNotificationEmails sets SystemNotificationEmails field to given value.

### HasSystemNotificationEmails

`func (o *UpdateOrgSettingsRequest) HasSystemNotificationEmails() bool`

HasSystemNotificationEmails returns a boolean if a field has been set.

### GetTrackAnalytics

`func (o *UpdateOrgSettingsRequest) GetTrackAnalytics() bool`

GetTrackAnalytics returns the TrackAnalytics field if non-nil, zero value otherwise.

### GetTrackAnalyticsOk

`func (o *UpdateOrgSettingsRequest) GetTrackAnalyticsOk() (*bool, bool)`

GetTrackAnalyticsOk returns a tuple with the TrackAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAnalytics

`func (o *UpdateOrgSettingsRequest) SetTrackAnalytics(v bool)`

SetTrackAnalytics sets TrackAnalytics field to given value.

### HasTrackAnalytics

`func (o *UpdateOrgSettingsRequest) HasTrackAnalytics() bool`

HasTrackAnalytics returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *UpdateOrgSettingsRequest) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *UpdateOrgSettingsRequest) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *UpdateOrgSettingsRequest) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *UpdateOrgSettingsRequest) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetUsageCertText

`func (o *UpdateOrgSettingsRequest) GetUsageCertText() string`

GetUsageCertText returns the UsageCertText field if non-nil, zero value otherwise.

### GetUsageCertTextOk

`func (o *UpdateOrgSettingsRequest) GetUsageCertTextOk() (*string, bool)`

GetUsageCertTextOk returns a tuple with the UsageCertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertText

`func (o *UpdateOrgSettingsRequest) SetUsageCertText(v string)`

SetUsageCertText sets UsageCertText field to given value.

### HasUsageCertText

`func (o *UpdateOrgSettingsRequest) HasUsageCertText() bool`

HasUsageCertText returns a boolean if a field has been set.

### GetUsernameEmptyText

`func (o *UpdateOrgSettingsRequest) GetUsernameEmptyText() string`

GetUsernameEmptyText returns the UsernameEmptyText field if non-nil, zero value otherwise.

### GetUsernameEmptyTextOk

`func (o *UpdateOrgSettingsRequest) GetUsernameEmptyTextOk() (*string, bool)`

GetUsernameEmptyTextOk returns a tuple with the UsernameEmptyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameEmptyText

`func (o *UpdateOrgSettingsRequest) SetUsernameEmptyText(v string)`

SetUsernameEmptyText sets UsernameEmptyText field to given value.

### HasUsernameEmptyText

`func (o *UpdateOrgSettingsRequest) HasUsernameEmptyText() bool`

HasUsernameEmptyText returns a boolean if a field has been set.

### GetUsernameLabel

`func (o *UpdateOrgSettingsRequest) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *UpdateOrgSettingsRequest) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *UpdateOrgSettingsRequest) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *UpdateOrgSettingsRequest) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetWhiteList

`func (o *UpdateOrgSettingsRequest) GetWhiteList() bool`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *UpdateOrgSettingsRequest) GetWhiteListOk() (*bool, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *UpdateOrgSettingsRequest) SetWhiteList(v bool)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *UpdateOrgSettingsRequest) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *UpdateOrgSettingsRequest) GetApprovalConfig() GetOrgSettings200ResponseApprovalConfig`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *UpdateOrgSettingsRequest) GetApprovalConfigOk() (*GetOrgSettings200ResponseApprovalConfig, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *UpdateOrgSettingsRequest) SetApprovalConfig(v GetOrgSettings200ResponseApprovalConfig)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *UpdateOrgSettingsRequest) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


