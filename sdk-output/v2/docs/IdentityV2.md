# IdentityV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**EncryptionCheck** | Pointer to **string** |  | [optional] 
**PasswordResetSinceLastLogin** | Pointer to **bool** |  | [optional] 
**UsageCertAttested** | Pointer to **time.Time** |  | [optional] 
**IdentityFlags** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AltAuthVia** | Pointer to **string** |  | [optional] 
**AltAuthViaIntegrationData** | Pointer to **string** |  | [optional] 
**KbaAnswers** | Pointer to **int32** |  | [optional] 
**DisablePasswordReset** | Pointer to **bool** |  | [optional] 
**PtaSourceId** | Pointer to **string** |  | [optional] 
**SupportsPasswordPush** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to [**DynamicSchemaEto**](DynamicSchemaEto.md) |  | [optional] 
**Role** | Pointer to **[]string** |  | [optional] 
**Phone** | Pointer to **string** | Work phone. | [optional] 
**Email** | Pointer to **string** | Work email. | [optional] 
**AltEmail** | Pointer to **string** |  | [optional] 
**AltPhone** | Pointer to **string** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityV2

`func NewIdentityV2() *IdentityV2`

NewIdentityV2 instantiates a new IdentityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityV2WithDefaults

`func NewIdentityV2WithDefaults() *IdentityV2`

NewIdentityV2WithDefaults instantiates a new IdentityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *IdentityV2) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IdentityV2) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IdentityV2) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IdentityV2) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetUid

`func (o *IdentityV2) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IdentityV2) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IdentityV2) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IdentityV2) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *IdentityV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *IdentityV2) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IdentityV2) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IdentityV2) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IdentityV2) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *IdentityV2) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IdentityV2) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IdentityV2) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IdentityV2) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdentityV2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdentityV2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdentityV2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdentityV2) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *IdentityV2) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *IdentityV2) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *IdentityV2) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *IdentityV2) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionCheck

`func (o *IdentityV2) GetEncryptionCheck() string`

GetEncryptionCheck returns the EncryptionCheck field if non-nil, zero value otherwise.

### GetEncryptionCheckOk

`func (o *IdentityV2) GetEncryptionCheckOk() (*string, bool)`

GetEncryptionCheckOk returns a tuple with the EncryptionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCheck

`func (o *IdentityV2) SetEncryptionCheck(v string)`

SetEncryptionCheck sets EncryptionCheck field to given value.

### HasEncryptionCheck

`func (o *IdentityV2) HasEncryptionCheck() bool`

HasEncryptionCheck returns a boolean if a field has been set.

### GetPasswordResetSinceLastLogin

`func (o *IdentityV2) GetPasswordResetSinceLastLogin() bool`

GetPasswordResetSinceLastLogin returns the PasswordResetSinceLastLogin field if non-nil, zero value otherwise.

### GetPasswordResetSinceLastLoginOk

`func (o *IdentityV2) GetPasswordResetSinceLastLoginOk() (*bool, bool)`

GetPasswordResetSinceLastLoginOk returns a tuple with the PasswordResetSinceLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetSinceLastLogin

`func (o *IdentityV2) SetPasswordResetSinceLastLogin(v bool)`

SetPasswordResetSinceLastLogin sets PasswordResetSinceLastLogin field to given value.

### HasPasswordResetSinceLastLogin

`func (o *IdentityV2) HasPasswordResetSinceLastLogin() bool`

HasPasswordResetSinceLastLogin returns a boolean if a field has been set.

### GetUsageCertAttested

`func (o *IdentityV2) GetUsageCertAttested() time.Time`

GetUsageCertAttested returns the UsageCertAttested field if non-nil, zero value otherwise.

### GetUsageCertAttestedOk

`func (o *IdentityV2) GetUsageCertAttestedOk() (*time.Time, bool)`

GetUsageCertAttestedOk returns a tuple with the UsageCertAttested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertAttested

`func (o *IdentityV2) SetUsageCertAttested(v time.Time)`

SetUsageCertAttested sets UsageCertAttested field to given value.

### HasUsageCertAttested

`func (o *IdentityV2) HasUsageCertAttested() bool`

HasUsageCertAttested returns a boolean if a field has been set.

### GetIdentityFlags

`func (o *IdentityV2) GetIdentityFlags() map[string]interface{}`

GetIdentityFlags returns the IdentityFlags field if non-nil, zero value otherwise.

### GetIdentityFlagsOk

`func (o *IdentityV2) GetIdentityFlagsOk() (*map[string]interface{}, bool)`

GetIdentityFlagsOk returns a tuple with the IdentityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityFlags

`func (o *IdentityV2) SetIdentityFlags(v map[string]interface{})`

SetIdentityFlags sets IdentityFlags field to given value.

### HasIdentityFlags

`func (o *IdentityV2) HasIdentityFlags() bool`

HasIdentityFlags returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAltAuthVia

`func (o *IdentityV2) GetAltAuthVia() string`

GetAltAuthVia returns the AltAuthVia field if non-nil, zero value otherwise.

### GetAltAuthViaOk

`func (o *IdentityV2) GetAltAuthViaOk() (*string, bool)`

GetAltAuthViaOk returns a tuple with the AltAuthVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAuthVia

`func (o *IdentityV2) SetAltAuthVia(v string)`

SetAltAuthVia sets AltAuthVia field to given value.

### HasAltAuthVia

`func (o *IdentityV2) HasAltAuthVia() bool`

HasAltAuthVia returns a boolean if a field has been set.

### GetAltAuthViaIntegrationData

`func (o *IdentityV2) GetAltAuthViaIntegrationData() string`

GetAltAuthViaIntegrationData returns the AltAuthViaIntegrationData field if non-nil, zero value otherwise.

### GetAltAuthViaIntegrationDataOk

`func (o *IdentityV2) GetAltAuthViaIntegrationDataOk() (*string, bool)`

GetAltAuthViaIntegrationDataOk returns a tuple with the AltAuthViaIntegrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAuthViaIntegrationData

`func (o *IdentityV2) SetAltAuthViaIntegrationData(v string)`

SetAltAuthViaIntegrationData sets AltAuthViaIntegrationData field to given value.

### HasAltAuthViaIntegrationData

`func (o *IdentityV2) HasAltAuthViaIntegrationData() bool`

HasAltAuthViaIntegrationData returns a boolean if a field has been set.

### GetKbaAnswers

`func (o *IdentityV2) GetKbaAnswers() int32`

GetKbaAnswers returns the KbaAnswers field if non-nil, zero value otherwise.

### GetKbaAnswersOk

`func (o *IdentityV2) GetKbaAnswersOk() (*int32, bool)`

GetKbaAnswersOk returns a tuple with the KbaAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbaAnswers

`func (o *IdentityV2) SetKbaAnswers(v int32)`

SetKbaAnswers sets KbaAnswers field to given value.

### HasKbaAnswers

`func (o *IdentityV2) HasKbaAnswers() bool`

HasKbaAnswers returns a boolean if a field has been set.

### GetDisablePasswordReset

`func (o *IdentityV2) GetDisablePasswordReset() bool`

GetDisablePasswordReset returns the DisablePasswordReset field if non-nil, zero value otherwise.

### GetDisablePasswordResetOk

`func (o *IdentityV2) GetDisablePasswordResetOk() (*bool, bool)`

GetDisablePasswordResetOk returns a tuple with the DisablePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePasswordReset

`func (o *IdentityV2) SetDisablePasswordReset(v bool)`

SetDisablePasswordReset sets DisablePasswordReset field to given value.

### HasDisablePasswordReset

`func (o *IdentityV2) HasDisablePasswordReset() bool`

HasDisablePasswordReset returns a boolean if a field has been set.

### GetPtaSourceId

`func (o *IdentityV2) GetPtaSourceId() string`

GetPtaSourceId returns the PtaSourceId field if non-nil, zero value otherwise.

### GetPtaSourceIdOk

`func (o *IdentityV2) GetPtaSourceIdOk() (*string, bool)`

GetPtaSourceIdOk returns a tuple with the PtaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtaSourceId

`func (o *IdentityV2) SetPtaSourceId(v string)`

SetPtaSourceId sets PtaSourceId field to given value.

### HasPtaSourceId

`func (o *IdentityV2) HasPtaSourceId() bool`

HasPtaSourceId returns a boolean if a field has been set.

### GetSupportsPasswordPush

`func (o *IdentityV2) GetSupportsPasswordPush() bool`

GetSupportsPasswordPush returns the SupportsPasswordPush field if non-nil, zero value otherwise.

### GetSupportsPasswordPushOk

`func (o *IdentityV2) GetSupportsPasswordPushOk() (*bool, bool)`

GetSupportsPasswordPushOk returns a tuple with the SupportsPasswordPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPasswordPush

`func (o *IdentityV2) SetSupportsPasswordPush(v bool)`

SetSupportsPasswordPush sets SupportsPasswordPush field to given value.

### HasSupportsPasswordPush

`func (o *IdentityV2) HasSupportsPasswordPush() bool`

HasSupportsPasswordPush returns a boolean if a field has been set.

### GetAttributes

`func (o *IdentityV2) GetAttributes() DynamicSchemaEto`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IdentityV2) GetAttributesOk() (*DynamicSchemaEto, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IdentityV2) SetAttributes(v DynamicSchemaEto)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IdentityV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRole

`func (o *IdentityV2) GetRole() []string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IdentityV2) GetRoleOk() (*[]string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IdentityV2) SetRole(v []string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IdentityV2) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPhone

`func (o *IdentityV2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *IdentityV2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *IdentityV2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *IdentityV2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAltEmail

`func (o *IdentityV2) GetAltEmail() string`

GetAltEmail returns the AltEmail field if non-nil, zero value otherwise.

### GetAltEmailOk

`func (o *IdentityV2) GetAltEmailOk() (*string, bool)`

GetAltEmailOk returns a tuple with the AltEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltEmail

`func (o *IdentityV2) SetAltEmail(v string)`

SetAltEmail sets AltEmail field to given value.

### HasAltEmail

`func (o *IdentityV2) HasAltEmail() bool`

HasAltEmail returns a boolean if a field has been set.

### GetAltPhone

`func (o *IdentityV2) GetAltPhone() string`

GetAltPhone returns the AltPhone field if non-nil, zero value otherwise.

### GetAltPhoneOk

`func (o *IdentityV2) GetAltPhoneOk() (*string, bool)`

GetAltPhoneOk returns a tuple with the AltPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhone

`func (o *IdentityV2) SetAltPhone(v string)`

SetAltPhone sets AltPhone field to given value.

### HasAltPhone

`func (o *IdentityV2) HasAltPhone() bool`

HasAltPhone returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *IdentityV2) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *IdentityV2) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *IdentityV2) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *IdentityV2) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


