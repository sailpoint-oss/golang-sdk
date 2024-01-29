# ListAccounts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**PasswordRequired** | Pointer to **bool** |  | [optional] 
**PasswordProvided** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SsoMethod** | Pointer to **string** |  | [optional] 
**IdEncryption** | Pointer to **string** |  | [optional] 
**PasswordEncryption** | Pointer to **string** |  | [optional] 
**LastPasswdChange** | Pointer to **NullableString** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**DateDisabled** | Pointer to **NullableString** |  | [optional] 
**AccountServiceId** | Pointer to **int32** |  | [optional] 
**ServiceId** | Pointer to **int32** |  | [optional] 
**PendingPasswordRequestId** | Pointer to **NullableString** |  | [optional] 
**PasswordChangeStatus** | Pointer to **string** |  | [optional] 
**PasswordChangeResult** | Pointer to [**ListAccounts200ResponseInnerPasswordChangeResult**](ListAccounts200ResponseInnerPasswordChangeResult.md) |  | [optional] 

## Methods

### NewListAccounts200ResponseInner

`func NewListAccounts200ResponseInner() *ListAccounts200ResponseInner`

NewListAccounts200ResponseInner instantiates a new ListAccounts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccounts200ResponseInnerWithDefaults

`func NewListAccounts200ResponseInnerWithDefaults() *ListAccounts200ResponseInner`

NewListAccounts200ResponseInnerWithDefaults instantiates a new ListAccounts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAccounts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAccounts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAccounts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAccounts200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ListAccounts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAccounts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAccounts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListAccounts200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListAccounts200ResponseInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAccounts200ResponseInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAccounts200ResponseInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListAccounts200ResponseInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUsername

`func (o *ListAccounts200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListAccounts200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListAccounts200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListAccounts200ResponseInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPasswordRequired

`func (o *ListAccounts200ResponseInner) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *ListAccounts200ResponseInner) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *ListAccounts200ResponseInner) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *ListAccounts200ResponseInner) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetPasswordProvided

`func (o *ListAccounts200ResponseInner) GetPasswordProvided() bool`

GetPasswordProvided returns the PasswordProvided field if non-nil, zero value otherwise.

### GetPasswordProvidedOk

`func (o *ListAccounts200ResponseInner) GetPasswordProvidedOk() (*bool, bool)`

GetPasswordProvidedOk returns a tuple with the PasswordProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProvided

`func (o *ListAccounts200ResponseInner) SetPasswordProvided(v bool)`

SetPasswordProvided sets PasswordProvided field to given value.

### HasPasswordProvided

`func (o *ListAccounts200ResponseInner) HasPasswordProvided() bool`

HasPasswordProvided returns a boolean if a field has been set.

### GetApps

`func (o *ListAccounts200ResponseInner) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListAccounts200ResponseInner) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListAccounts200ResponseInner) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListAccounts200ResponseInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetSsoMethod

`func (o *ListAccounts200ResponseInner) GetSsoMethod() string`

GetSsoMethod returns the SsoMethod field if non-nil, zero value otherwise.

### GetSsoMethodOk

`func (o *ListAccounts200ResponseInner) GetSsoMethodOk() (*string, bool)`

GetSsoMethodOk returns a tuple with the SsoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoMethod

`func (o *ListAccounts200ResponseInner) SetSsoMethod(v string)`

SetSsoMethod sets SsoMethod field to given value.

### HasSsoMethod

`func (o *ListAccounts200ResponseInner) HasSsoMethod() bool`

HasSsoMethod returns a boolean if a field has been set.

### GetIdEncryption

`func (o *ListAccounts200ResponseInner) GetIdEncryption() string`

GetIdEncryption returns the IdEncryption field if non-nil, zero value otherwise.

### GetIdEncryptionOk

`func (o *ListAccounts200ResponseInner) GetIdEncryptionOk() (*string, bool)`

GetIdEncryptionOk returns a tuple with the IdEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdEncryption

`func (o *ListAccounts200ResponseInner) SetIdEncryption(v string)`

SetIdEncryption sets IdEncryption field to given value.

### HasIdEncryption

`func (o *ListAccounts200ResponseInner) HasIdEncryption() bool`

HasIdEncryption returns a boolean if a field has been set.

### GetPasswordEncryption

`func (o *ListAccounts200ResponseInner) GetPasswordEncryption() string`

GetPasswordEncryption returns the PasswordEncryption field if non-nil, zero value otherwise.

### GetPasswordEncryptionOk

`func (o *ListAccounts200ResponseInner) GetPasswordEncryptionOk() (*string, bool)`

GetPasswordEncryptionOk returns a tuple with the PasswordEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncryption

`func (o *ListAccounts200ResponseInner) SetPasswordEncryption(v string)`

SetPasswordEncryption sets PasswordEncryption field to given value.

### HasPasswordEncryption

`func (o *ListAccounts200ResponseInner) HasPasswordEncryption() bool`

HasPasswordEncryption returns a boolean if a field has been set.

### GetLastPasswdChange

`func (o *ListAccounts200ResponseInner) GetLastPasswdChange() string`

GetLastPasswdChange returns the LastPasswdChange field if non-nil, zero value otherwise.

### GetLastPasswdChangeOk

`func (o *ListAccounts200ResponseInner) GetLastPasswdChangeOk() (*string, bool)`

GetLastPasswdChangeOk returns a tuple with the LastPasswdChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswdChange

`func (o *ListAccounts200ResponseInner) SetLastPasswdChange(v string)`

SetLastPasswdChange sets LastPasswdChange field to given value.

### HasLastPasswdChange

`func (o *ListAccounts200ResponseInner) HasLastPasswdChange() bool`

HasLastPasswdChange returns a boolean if a field has been set.

### SetLastPasswdChangeNil

`func (o *ListAccounts200ResponseInner) SetLastPasswdChangeNil(b bool)`

 SetLastPasswdChangeNil sets the value for LastPasswdChange to be an explicit nil

### UnsetLastPasswdChange
`func (o *ListAccounts200ResponseInner) UnsetLastPasswdChange()`

UnsetLastPasswdChange ensures that no value is present for LastPasswdChange, not even an explicit nil
### GetServiceName

`func (o *ListAccounts200ResponseInner) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ListAccounts200ResponseInner) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ListAccounts200ResponseInner) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ListAccounts200ResponseInner) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetDateDisabled

`func (o *ListAccounts200ResponseInner) GetDateDisabled() string`

GetDateDisabled returns the DateDisabled field if non-nil, zero value otherwise.

### GetDateDisabledOk

`func (o *ListAccounts200ResponseInner) GetDateDisabledOk() (*string, bool)`

GetDateDisabledOk returns a tuple with the DateDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDisabled

`func (o *ListAccounts200ResponseInner) SetDateDisabled(v string)`

SetDateDisabled sets DateDisabled field to given value.

### HasDateDisabled

`func (o *ListAccounts200ResponseInner) HasDateDisabled() bool`

HasDateDisabled returns a boolean if a field has been set.

### SetDateDisabledNil

`func (o *ListAccounts200ResponseInner) SetDateDisabledNil(b bool)`

 SetDateDisabledNil sets the value for DateDisabled to be an explicit nil

### UnsetDateDisabled
`func (o *ListAccounts200ResponseInner) UnsetDateDisabled()`

UnsetDateDisabled ensures that no value is present for DateDisabled, not even an explicit nil
### GetAccountServiceId

`func (o *ListAccounts200ResponseInner) GetAccountServiceId() int32`

GetAccountServiceId returns the AccountServiceId field if non-nil, zero value otherwise.

### GetAccountServiceIdOk

`func (o *ListAccounts200ResponseInner) GetAccountServiceIdOk() (*int32, bool)`

GetAccountServiceIdOk returns a tuple with the AccountServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceId

`func (o *ListAccounts200ResponseInner) SetAccountServiceId(v int32)`

SetAccountServiceId sets AccountServiceId field to given value.

### HasAccountServiceId

`func (o *ListAccounts200ResponseInner) HasAccountServiceId() bool`

HasAccountServiceId returns a boolean if a field has been set.

### GetServiceId

`func (o *ListAccounts200ResponseInner) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListAccounts200ResponseInner) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListAccounts200ResponseInner) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListAccounts200ResponseInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPendingPasswordRequestId

`func (o *ListAccounts200ResponseInner) GetPendingPasswordRequestId() string`

GetPendingPasswordRequestId returns the PendingPasswordRequestId field if non-nil, zero value otherwise.

### GetPendingPasswordRequestIdOk

`func (o *ListAccounts200ResponseInner) GetPendingPasswordRequestIdOk() (*string, bool)`

GetPendingPasswordRequestIdOk returns a tuple with the PendingPasswordRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPasswordRequestId

`func (o *ListAccounts200ResponseInner) SetPendingPasswordRequestId(v string)`

SetPendingPasswordRequestId sets PendingPasswordRequestId field to given value.

### HasPendingPasswordRequestId

`func (o *ListAccounts200ResponseInner) HasPendingPasswordRequestId() bool`

HasPendingPasswordRequestId returns a boolean if a field has been set.

### SetPendingPasswordRequestIdNil

`func (o *ListAccounts200ResponseInner) SetPendingPasswordRequestIdNil(b bool)`

 SetPendingPasswordRequestIdNil sets the value for PendingPasswordRequestId to be an explicit nil

### UnsetPendingPasswordRequestId
`func (o *ListAccounts200ResponseInner) UnsetPendingPasswordRequestId()`

UnsetPendingPasswordRequestId ensures that no value is present for PendingPasswordRequestId, not even an explicit nil
### GetPasswordChangeStatus

`func (o *ListAccounts200ResponseInner) GetPasswordChangeStatus() string`

GetPasswordChangeStatus returns the PasswordChangeStatus field if non-nil, zero value otherwise.

### GetPasswordChangeStatusOk

`func (o *ListAccounts200ResponseInner) GetPasswordChangeStatusOk() (*string, bool)`

GetPasswordChangeStatusOk returns a tuple with the PasswordChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeStatus

`func (o *ListAccounts200ResponseInner) SetPasswordChangeStatus(v string)`

SetPasswordChangeStatus sets PasswordChangeStatus field to given value.

### HasPasswordChangeStatus

`func (o *ListAccounts200ResponseInner) HasPasswordChangeStatus() bool`

HasPasswordChangeStatus returns a boolean if a field has been set.

### GetPasswordChangeResult

`func (o *ListAccounts200ResponseInner) GetPasswordChangeResult() ListAccounts200ResponseInnerPasswordChangeResult`

GetPasswordChangeResult returns the PasswordChangeResult field if non-nil, zero value otherwise.

### GetPasswordChangeResultOk

`func (o *ListAccounts200ResponseInner) GetPasswordChangeResultOk() (*ListAccounts200ResponseInnerPasswordChangeResult, bool)`

GetPasswordChangeResultOk returns a tuple with the PasswordChangeResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeResult

`func (o *ListAccounts200ResponseInner) SetPasswordChangeResult(v ListAccounts200ResponseInnerPasswordChangeResult)`

SetPasswordChangeResult sets PasswordChangeResult field to given value.

### HasPasswordChangeResult

`func (o *ListAccounts200ResponseInner) HasPasswordChangeResult() bool`

HasPasswordChangeResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


