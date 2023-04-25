# GetIdentity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**EncryptionKey** | Pointer to **map[string]interface{}** |  | [optional] 
**EncryptionCheck** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 
**PasswordResetSinceLastLogin** | Pointer to **bool** |  | [optional] 
**UsageCertAttested** | Pointer to **map[string]interface{}** |  | [optional] 
**UserFlags** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AltAuthVia** | Pointer to **string** |  | [optional] 
**AltAuthViaIntegrationData** | Pointer to **map[string]interface{}** |  | [optional] 
**KbaAnswers** | Pointer to **float32** |  | [optional] 
**DisablePasswordReset** | Pointer to **bool** |  | [optional] 
**PtaSourceId** | Pointer to **map[string]interface{}** |  | [optional] 
**SupportsPasswordPush** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Phone** | Pointer to **map[string]interface{}** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PersonalEmail** | Pointer to **map[string]interface{}** |  | [optional] 
**EmployeeNumber** | Pointer to **map[string]interface{}** |  | [optional] 
**RiskScore** | Pointer to **float32** |  | [optional] 
**FeatureFlags** | Pointer to **map[string]interface{}** |  | [optional] 
**Feature** | Pointer to **[]string** |  | [optional] 
**OrgEncryptionKey** | Pointer to **string** |  | [optional] 
**OrgEncryptionKeyId** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Org** | Pointer to [**GetIdentity200ResponseOrg**](GetIdentity200ResponseOrg.md) |  | [optional] 
**StepUpAuth** | Pointer to **bool** |  | [optional] 
**BxInstallPrompted** | Pointer to **bool** |  | [optional] 
**FederatedLogin** | Pointer to **bool** |  | [optional] 
**Auth** | Pointer to [**GetIdentity200ResponseAuth**](GetIdentity200ResponseAuth.md) |  | [optional] 
**OnNetwork** | Pointer to **bool** |  | [optional] 
**OnTrustedGeo** | Pointer to **bool** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIdentity200Response

`func NewGetIdentity200Response() *GetIdentity200Response`

NewGetIdentity200Response instantiates a new GetIdentity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentity200ResponseWithDefaults

`func NewGetIdentity200ResponseWithDefaults() *GetIdentity200Response`

NewGetIdentity200ResponseWithDefaults instantiates a new GetIdentity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIdentity200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIdentity200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIdentity200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetIdentity200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *GetIdentity200Response) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetIdentity200Response) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetIdentity200Response) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetIdentity200Response) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetUid

`func (o *GetIdentity200Response) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetIdentity200Response) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetIdentity200Response) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetIdentity200Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *GetIdentity200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIdentity200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIdentity200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIdentity200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetIdentity200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetIdentity200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetIdentity200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetIdentity200Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUuid

`func (o *GetIdentity200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetIdentity200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetIdentity200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetIdentity200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *GetIdentity200Response) GetEncryptionKey() map[string]interface{}`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *GetIdentity200Response) GetEncryptionKeyOk() (*map[string]interface{}, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *GetIdentity200Response) SetEncryptionKey(v map[string]interface{})`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *GetIdentity200Response) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionCheck

`func (o *GetIdentity200Response) GetEncryptionCheck() map[string]interface{}`

GetEncryptionCheck returns the EncryptionCheck field if non-nil, zero value otherwise.

### GetEncryptionCheckOk

`func (o *GetIdentity200Response) GetEncryptionCheckOk() (*map[string]interface{}, bool)`

GetEncryptionCheckOk returns a tuple with the EncryptionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCheck

`func (o *GetIdentity200Response) SetEncryptionCheck(v map[string]interface{})`

SetEncryptionCheck sets EncryptionCheck field to given value.

### HasEncryptionCheck

`func (o *GetIdentity200Response) HasEncryptionCheck() bool`

HasEncryptionCheck returns a boolean if a field has been set.

### GetStatus

`func (o *GetIdentity200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIdentity200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIdentity200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIdentity200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPending

`func (o *GetIdentity200Response) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GetIdentity200Response) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GetIdentity200Response) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *GetIdentity200Response) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPasswordResetSinceLastLogin

`func (o *GetIdentity200Response) GetPasswordResetSinceLastLogin() bool`

GetPasswordResetSinceLastLogin returns the PasswordResetSinceLastLogin field if non-nil, zero value otherwise.

### GetPasswordResetSinceLastLoginOk

`func (o *GetIdentity200Response) GetPasswordResetSinceLastLoginOk() (*bool, bool)`

GetPasswordResetSinceLastLoginOk returns a tuple with the PasswordResetSinceLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetSinceLastLogin

`func (o *GetIdentity200Response) SetPasswordResetSinceLastLogin(v bool)`

SetPasswordResetSinceLastLogin sets PasswordResetSinceLastLogin field to given value.

### HasPasswordResetSinceLastLogin

`func (o *GetIdentity200Response) HasPasswordResetSinceLastLogin() bool`

HasPasswordResetSinceLastLogin returns a boolean if a field has been set.

### GetUsageCertAttested

`func (o *GetIdentity200Response) GetUsageCertAttested() map[string]interface{}`

GetUsageCertAttested returns the UsageCertAttested field if non-nil, zero value otherwise.

### GetUsageCertAttestedOk

`func (o *GetIdentity200Response) GetUsageCertAttestedOk() (*map[string]interface{}, bool)`

GetUsageCertAttestedOk returns a tuple with the UsageCertAttested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertAttested

`func (o *GetIdentity200Response) SetUsageCertAttested(v map[string]interface{})`

SetUsageCertAttested sets UsageCertAttested field to given value.

### HasUsageCertAttested

`func (o *GetIdentity200Response) HasUsageCertAttested() bool`

HasUsageCertAttested returns a boolean if a field has been set.

### GetUserFlags

`func (o *GetIdentity200Response) GetUserFlags() map[string]interface{}`

GetUserFlags returns the UserFlags field if non-nil, zero value otherwise.

### GetUserFlagsOk

`func (o *GetIdentity200Response) GetUserFlagsOk() (*map[string]interface{}, bool)`

GetUserFlagsOk returns a tuple with the UserFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlags

`func (o *GetIdentity200Response) SetUserFlags(v map[string]interface{})`

SetUserFlags sets UserFlags field to given value.

### HasUserFlags

`func (o *GetIdentity200Response) HasUserFlags() bool`

HasUserFlags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIdentity200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIdentity200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIdentity200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIdentity200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAltAuthVia

`func (o *GetIdentity200Response) GetAltAuthVia() string`

GetAltAuthVia returns the AltAuthVia field if non-nil, zero value otherwise.

### GetAltAuthViaOk

`func (o *GetIdentity200Response) GetAltAuthViaOk() (*string, bool)`

GetAltAuthViaOk returns a tuple with the AltAuthVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAuthVia

`func (o *GetIdentity200Response) SetAltAuthVia(v string)`

SetAltAuthVia sets AltAuthVia field to given value.

### HasAltAuthVia

`func (o *GetIdentity200Response) HasAltAuthVia() bool`

HasAltAuthVia returns a boolean if a field has been set.

### GetAltAuthViaIntegrationData

`func (o *GetIdentity200Response) GetAltAuthViaIntegrationData() map[string]interface{}`

GetAltAuthViaIntegrationData returns the AltAuthViaIntegrationData field if non-nil, zero value otherwise.

### GetAltAuthViaIntegrationDataOk

`func (o *GetIdentity200Response) GetAltAuthViaIntegrationDataOk() (*map[string]interface{}, bool)`

GetAltAuthViaIntegrationDataOk returns a tuple with the AltAuthViaIntegrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAuthViaIntegrationData

`func (o *GetIdentity200Response) SetAltAuthViaIntegrationData(v map[string]interface{})`

SetAltAuthViaIntegrationData sets AltAuthViaIntegrationData field to given value.

### HasAltAuthViaIntegrationData

`func (o *GetIdentity200Response) HasAltAuthViaIntegrationData() bool`

HasAltAuthViaIntegrationData returns a boolean if a field has been set.

### GetKbaAnswers

`func (o *GetIdentity200Response) GetKbaAnswers() float32`

GetKbaAnswers returns the KbaAnswers field if non-nil, zero value otherwise.

### GetKbaAnswersOk

`func (o *GetIdentity200Response) GetKbaAnswersOk() (*float32, bool)`

GetKbaAnswersOk returns a tuple with the KbaAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaAnswers

`func (o *GetIdentity200Response) SetKbaAnswers(v float32)`

SetKbaAnswers sets KbaAnswers field to given value.

### HasKbaAnswers

`func (o *GetIdentity200Response) HasKbaAnswers() bool`

HasKbaAnswers returns a boolean if a field has been set.

### GetDisablePasswordReset

`func (o *GetIdentity200Response) GetDisablePasswordReset() bool`

GetDisablePasswordReset returns the DisablePasswordReset field if non-nil, zero value otherwise.

### GetDisablePasswordResetOk

`func (o *GetIdentity200Response) GetDisablePasswordResetOk() (*bool, bool)`

GetDisablePasswordResetOk returns a tuple with the DisablePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePasswordReset

`func (o *GetIdentity200Response) SetDisablePasswordReset(v bool)`

SetDisablePasswordReset sets DisablePasswordReset field to given value.

### HasDisablePasswordReset

`func (o *GetIdentity200Response) HasDisablePasswordReset() bool`

HasDisablePasswordReset returns a boolean if a field has been set.

### GetPtaSourceId

`func (o *GetIdentity200Response) GetPtaSourceId() map[string]interface{}`

GetPtaSourceId returns the PtaSourceId field if non-nil, zero value otherwise.

### GetPtaSourceIdOk

`func (o *GetIdentity200Response) GetPtaSourceIdOk() (*map[string]interface{}, bool)`

GetPtaSourceIdOk returns a tuple with the PtaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtaSourceId

`func (o *GetIdentity200Response) SetPtaSourceId(v map[string]interface{})`

SetPtaSourceId sets PtaSourceId field to given value.

### HasPtaSourceId

`func (o *GetIdentity200Response) HasPtaSourceId() bool`

HasPtaSourceId returns a boolean if a field has been set.

### GetSupportsPasswordPush

`func (o *GetIdentity200Response) GetSupportsPasswordPush() bool`

GetSupportsPasswordPush returns the SupportsPasswordPush field if non-nil, zero value otherwise.

### GetSupportsPasswordPushOk

`func (o *GetIdentity200Response) GetSupportsPasswordPushOk() (*bool, bool)`

GetSupportsPasswordPushOk returns a tuple with the SupportsPasswordPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPasswordPush

`func (o *GetIdentity200Response) SetSupportsPasswordPush(v bool)`

SetSupportsPasswordPush sets SupportsPasswordPush field to given value.

### HasSupportsPasswordPush

`func (o *GetIdentity200Response) HasSupportsPasswordPush() bool`

HasSupportsPasswordPush returns a boolean if a field has been set.

### GetAttributes

`func (o *GetIdentity200Response) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetIdentity200Response) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetIdentity200Response) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetIdentity200Response) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetExternalId

`func (o *GetIdentity200Response) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetIdentity200Response) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetIdentity200Response) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetIdentity200Response) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRole

`func (o *GetIdentity200Response) GetRole() []map[string]interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetIdentity200Response) GetRoleOk() (*[]map[string]interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetIdentity200Response) SetRole(v []map[string]interface{})`

SetRole sets Role field to given value.

### HasRole

`func (o *GetIdentity200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPhone

`func (o *GetIdentity200Response) GetPhone() map[string]interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetIdentity200Response) GetPhoneOk() (*map[string]interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetIdentity200Response) SetPhone(v map[string]interface{})`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetIdentity200Response) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *GetIdentity200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetIdentity200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetIdentity200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetIdentity200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPersonalEmail

`func (o *GetIdentity200Response) GetPersonalEmail() map[string]interface{}`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *GetIdentity200Response) GetPersonalEmailOk() (*map[string]interface{}, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *GetIdentity200Response) SetPersonalEmail(v map[string]interface{})`

SetPersonalEmail sets PersonalEmail field to given value.

### HasPersonalEmail

`func (o *GetIdentity200Response) HasPersonalEmail() bool`

HasPersonalEmail returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *GetIdentity200Response) GetEmployeeNumber() map[string]interface{}`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *GetIdentity200Response) GetEmployeeNumberOk() (*map[string]interface{}, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *GetIdentity200Response) SetEmployeeNumber(v map[string]interface{})`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *GetIdentity200Response) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetRiskScore

`func (o *GetIdentity200Response) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *GetIdentity200Response) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *GetIdentity200Response) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *GetIdentity200Response) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *GetIdentity200Response) GetFeatureFlags() map[string]interface{}`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *GetIdentity200Response) GetFeatureFlagsOk() (*map[string]interface{}, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *GetIdentity200Response) SetFeatureFlags(v map[string]interface{})`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *GetIdentity200Response) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetFeature

`func (o *GetIdentity200Response) GetFeature() []string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *GetIdentity200Response) GetFeatureOk() (*[]string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *GetIdentity200Response) SetFeature(v []string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *GetIdentity200Response) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetOrgEncryptionKey

`func (o *GetIdentity200Response) GetOrgEncryptionKey() string`

GetOrgEncryptionKey returns the OrgEncryptionKey field if non-nil, zero value otherwise.

### GetOrgEncryptionKeyOk

`func (o *GetIdentity200Response) GetOrgEncryptionKeyOk() (*string, bool)`

GetOrgEncryptionKeyOk returns a tuple with the OrgEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEncryptionKey

`func (o *GetIdentity200Response) SetOrgEncryptionKey(v string)`

SetOrgEncryptionKey sets OrgEncryptionKey field to given value.

### HasOrgEncryptionKey

`func (o *GetIdentity200Response) HasOrgEncryptionKey() bool`

HasOrgEncryptionKey returns a boolean if a field has been set.

### GetOrgEncryptionKeyId

`func (o *GetIdentity200Response) GetOrgEncryptionKeyId() string`

GetOrgEncryptionKeyId returns the OrgEncryptionKeyId field if non-nil, zero value otherwise.

### GetOrgEncryptionKeyIdOk

`func (o *GetIdentity200Response) GetOrgEncryptionKeyIdOk() (*string, bool)`

GetOrgEncryptionKeyIdOk returns a tuple with the OrgEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEncryptionKeyId

`func (o *GetIdentity200Response) SetOrgEncryptionKeyId(v string)`

SetOrgEncryptionKeyId sets OrgEncryptionKeyId field to given value.

### HasOrgEncryptionKeyId

`func (o *GetIdentity200Response) HasOrgEncryptionKeyId() bool`

HasOrgEncryptionKeyId returns a boolean if a field has been set.

### GetMeta

`func (o *GetIdentity200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetIdentity200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetIdentity200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetIdentity200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOrg

`func (o *GetIdentity200Response) GetOrg() GetIdentity200ResponseOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GetIdentity200Response) GetOrgOk() (*GetIdentity200ResponseOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GetIdentity200Response) SetOrg(v GetIdentity200ResponseOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GetIdentity200Response) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStepUpAuth

`func (o *GetIdentity200Response) GetStepUpAuth() bool`

GetStepUpAuth returns the StepUpAuth field if non-nil, zero value otherwise.

### GetStepUpAuthOk

`func (o *GetIdentity200Response) GetStepUpAuthOk() (*bool, bool)`

GetStepUpAuthOk returns a tuple with the StepUpAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUpAuth

`func (o *GetIdentity200Response) SetStepUpAuth(v bool)`

SetStepUpAuth sets StepUpAuth field to given value.

### HasStepUpAuth

`func (o *GetIdentity200Response) HasStepUpAuth() bool`

HasStepUpAuth returns a boolean if a field has been set.

### GetBxInstallPrompted

`func (o *GetIdentity200Response) GetBxInstallPrompted() bool`

GetBxInstallPrompted returns the BxInstallPrompted field if non-nil, zero value otherwise.

### GetBxInstallPromptedOk

`func (o *GetIdentity200Response) GetBxInstallPromptedOk() (*bool, bool)`

GetBxInstallPromptedOk returns a tuple with the BxInstallPrompted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBxInstallPrompted

`func (o *GetIdentity200Response) SetBxInstallPrompted(v bool)`

SetBxInstallPrompted sets BxInstallPrompted field to given value.

### HasBxInstallPrompted

`func (o *GetIdentity200Response) HasBxInstallPrompted() bool`

HasBxInstallPrompted returns a boolean if a field has been set.

### GetFederatedLogin

`func (o *GetIdentity200Response) GetFederatedLogin() bool`

GetFederatedLogin returns the FederatedLogin field if non-nil, zero value otherwise.

### GetFederatedLoginOk

`func (o *GetIdentity200Response) GetFederatedLoginOk() (*bool, bool)`

GetFederatedLoginOk returns a tuple with the FederatedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedLogin

`func (o *GetIdentity200Response) SetFederatedLogin(v bool)`

SetFederatedLogin sets FederatedLogin field to given value.

### HasFederatedLogin

`func (o *GetIdentity200Response) HasFederatedLogin() bool`

HasFederatedLogin returns a boolean if a field has been set.

### GetAuth

`func (o *GetIdentity200Response) GetAuth() GetIdentity200ResponseAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *GetIdentity200Response) GetAuthOk() (*GetIdentity200ResponseAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *GetIdentity200Response) SetAuth(v GetIdentity200ResponseAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *GetIdentity200Response) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetOnNetwork

`func (o *GetIdentity200Response) GetOnNetwork() bool`

GetOnNetwork returns the OnNetwork field if non-nil, zero value otherwise.

### GetOnNetworkOk

`func (o *GetIdentity200Response) GetOnNetworkOk() (*bool, bool)`

GetOnNetworkOk returns a tuple with the OnNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnNetwork

`func (o *GetIdentity200Response) SetOnNetwork(v bool)`

SetOnNetwork sets OnNetwork field to given value.

### HasOnNetwork

`func (o *GetIdentity200Response) HasOnNetwork() bool`

HasOnNetwork returns a boolean if a field has been set.

### GetOnTrustedGeo

`func (o *GetIdentity200Response) GetOnTrustedGeo() bool`

GetOnTrustedGeo returns the OnTrustedGeo field if non-nil, zero value otherwise.

### GetOnTrustedGeoOk

`func (o *GetIdentity200Response) GetOnTrustedGeoOk() (*bool, bool)`

GetOnTrustedGeoOk returns a tuple with the OnTrustedGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTrustedGeo

`func (o *GetIdentity200Response) SetOnTrustedGeo(v bool)`

SetOnTrustedGeo sets OnTrustedGeo field to given value.

### HasOnTrustedGeo

`func (o *GetIdentity200Response) HasOnTrustedGeo() bool`

HasOnTrustedGeo returns a boolean if a field has been set.

### GetLoginUrl

`func (o *GetIdentity200Response) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *GetIdentity200Response) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *GetIdentity200Response) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *GetIdentity200Response) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


