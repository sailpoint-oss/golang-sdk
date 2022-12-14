# AccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFlags** | Pointer to **[]string** |  | [optional] 
**Cn** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DistinguishedName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**IdNowDescription** | Pointer to **string** |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**MemberOf** | Pointer to **[]string** |  | [optional] 
**MsDSPrincipalName** | Pointer to **string** |  | [optional] 
**MsNPAllowDialin** | Pointer to **string** |  | [optional] 
**ObjectClass** | Pointer to **[]string** |  | [optional] 
**ObjectSid** | Pointer to **string** |  | [optional] 
**Objectguid** | Pointer to **string** |  | [optional] 
**PasswordLastSet** | Pointer to **int32** |  | [optional] 
**PrimaryGroupID** | Pointer to **string** |  | [optional] 
**PwdLastSet** | Pointer to **string** |  | [optional] 
**SAMAccountName** | Pointer to **string** |  | [optional] 
**Sn** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountAttributes

`func NewAccountAttributes() *AccountAttributes`

NewAccountAttributes instantiates a new AccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAttributesWithDefaults

`func NewAccountAttributesWithDefaults() *AccountAttributes`

NewAccountAttributesWithDefaults instantiates a new AccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFlags

`func (o *AccountAttributes) GetAccountFlags() []string`

GetAccountFlags returns the AccountFlags field if non-nil, zero value otherwise.

### GetAccountFlagsOk

`func (o *AccountAttributes) GetAccountFlagsOk() (*[]string, bool)`

GetAccountFlagsOk returns a tuple with the AccountFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFlags

`func (o *AccountAttributes) SetAccountFlags(v []string)`

SetAccountFlags sets AccountFlags field to given value.

### HasAccountFlags

`func (o *AccountAttributes) HasAccountFlags() bool`

HasAccountFlags returns a boolean if a field has been set.

### GetCn

`func (o *AccountAttributes) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *AccountAttributes) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *AccountAttributes) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *AccountAttributes) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccountAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccountAttributes) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *AccountAttributes) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *AccountAttributes) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *AccountAttributes) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *AccountAttributes) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetDomain

`func (o *AccountAttributes) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AccountAttributes) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AccountAttributes) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AccountAttributes) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGivenName

`func (o *AccountAttributes) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AccountAttributes) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AccountAttributes) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AccountAttributes) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetIdNowDescription

`func (o *AccountAttributes) GetIdNowDescription() string`

GetIdNowDescription returns the IdNowDescription field if non-nil, zero value otherwise.

### GetIdNowDescriptionOk

`func (o *AccountAttributes) GetIdNowDescriptionOk() (*string, bool)`

GetIdNowDescriptionOk returns a tuple with the IdNowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNowDescription

`func (o *AccountAttributes) SetIdNowDescription(v string)`

SetIdNowDescription sets IdNowDescription field to given value.

### HasIdNowDescription

`func (o *AccountAttributes) HasIdNowDescription() bool`

HasIdNowDescription returns a boolean if a field has been set.

### GetMail

`func (o *AccountAttributes) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *AccountAttributes) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *AccountAttributes) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *AccountAttributes) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMemberOf

`func (o *AccountAttributes) GetMemberOf() []string`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *AccountAttributes) GetMemberOfOk() (*[]string, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *AccountAttributes) SetMemberOf(v []string)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *AccountAttributes) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetMsDSPrincipalName

`func (o *AccountAttributes) GetMsDSPrincipalName() string`

GetMsDSPrincipalName returns the MsDSPrincipalName field if non-nil, zero value otherwise.

### GetMsDSPrincipalNameOk

`func (o *AccountAttributes) GetMsDSPrincipalNameOk() (*string, bool)`

GetMsDSPrincipalNameOk returns a tuple with the MsDSPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDSPrincipalName

`func (o *AccountAttributes) SetMsDSPrincipalName(v string)`

SetMsDSPrincipalName sets MsDSPrincipalName field to given value.

### HasMsDSPrincipalName

`func (o *AccountAttributes) HasMsDSPrincipalName() bool`

HasMsDSPrincipalName returns a boolean if a field has been set.

### GetMsNPAllowDialin

`func (o *AccountAttributes) GetMsNPAllowDialin() string`

GetMsNPAllowDialin returns the MsNPAllowDialin field if non-nil, zero value otherwise.

### GetMsNPAllowDialinOk

`func (o *AccountAttributes) GetMsNPAllowDialinOk() (*string, bool)`

GetMsNPAllowDialinOk returns a tuple with the MsNPAllowDialin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsNPAllowDialin

`func (o *AccountAttributes) SetMsNPAllowDialin(v string)`

SetMsNPAllowDialin sets MsNPAllowDialin field to given value.

### HasMsNPAllowDialin

`func (o *AccountAttributes) HasMsNPAllowDialin() bool`

HasMsNPAllowDialin returns a boolean if a field has been set.

### GetObjectClass

`func (o *AccountAttributes) GetObjectClass() []string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AccountAttributes) GetObjectClassOk() (*[]string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AccountAttributes) SetObjectClass(v []string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AccountAttributes) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectSid

`func (o *AccountAttributes) GetObjectSid() string`

GetObjectSid returns the ObjectSid field if non-nil, zero value otherwise.

### GetObjectSidOk

`func (o *AccountAttributes) GetObjectSidOk() (*string, bool)`

GetObjectSidOk returns a tuple with the ObjectSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSid

`func (o *AccountAttributes) SetObjectSid(v string)`

SetObjectSid sets ObjectSid field to given value.

### HasObjectSid

`func (o *AccountAttributes) HasObjectSid() bool`

HasObjectSid returns a boolean if a field has been set.

### GetObjectguid

`func (o *AccountAttributes) GetObjectguid() string`

GetObjectguid returns the Objectguid field if non-nil, zero value otherwise.

### GetObjectguidOk

`func (o *AccountAttributes) GetObjectguidOk() (*string, bool)`

GetObjectguidOk returns a tuple with the Objectguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectguid

`func (o *AccountAttributes) SetObjectguid(v string)`

SetObjectguid sets Objectguid field to given value.

### HasObjectguid

`func (o *AccountAttributes) HasObjectguid() bool`

HasObjectguid returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *AccountAttributes) GetPasswordLastSet() int32`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *AccountAttributes) GetPasswordLastSetOk() (*int32, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *AccountAttributes) SetPasswordLastSet(v int32)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *AccountAttributes) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### GetPrimaryGroupID

`func (o *AccountAttributes) GetPrimaryGroupID() string`

GetPrimaryGroupID returns the PrimaryGroupID field if non-nil, zero value otherwise.

### GetPrimaryGroupIDOk

`func (o *AccountAttributes) GetPrimaryGroupIDOk() (*string, bool)`

GetPrimaryGroupIDOk returns a tuple with the PrimaryGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroupID

`func (o *AccountAttributes) SetPrimaryGroupID(v string)`

SetPrimaryGroupID sets PrimaryGroupID field to given value.

### HasPrimaryGroupID

`func (o *AccountAttributes) HasPrimaryGroupID() bool`

HasPrimaryGroupID returns a boolean if a field has been set.

### GetPwdLastSet

`func (o *AccountAttributes) GetPwdLastSet() string`

GetPwdLastSet returns the PwdLastSet field if non-nil, zero value otherwise.

### GetPwdLastSetOk

`func (o *AccountAttributes) GetPwdLastSetOk() (*string, bool)`

GetPwdLastSetOk returns a tuple with the PwdLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdLastSet

`func (o *AccountAttributes) SetPwdLastSet(v string)`

SetPwdLastSet sets PwdLastSet field to given value.

### HasPwdLastSet

`func (o *AccountAttributes) HasPwdLastSet() bool`

HasPwdLastSet returns a boolean if a field has been set.

### GetSAMAccountName

`func (o *AccountAttributes) GetSAMAccountName() string`

GetSAMAccountName returns the SAMAccountName field if non-nil, zero value otherwise.

### GetSAMAccountNameOk

`func (o *AccountAttributes) GetSAMAccountNameOk() (*string, bool)`

GetSAMAccountNameOk returns a tuple with the SAMAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMAccountName

`func (o *AccountAttributes) SetSAMAccountName(v string)`

SetSAMAccountName sets SAMAccountName field to given value.

### HasSAMAccountName

`func (o *AccountAttributes) HasSAMAccountName() bool`

HasSAMAccountName returns a boolean if a field has been set.

### GetSn

`func (o *AccountAttributes) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *AccountAttributes) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *AccountAttributes) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *AccountAttributes) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *AccountAttributes) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *AccountAttributes) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *AccountAttributes) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *AccountAttributes) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


