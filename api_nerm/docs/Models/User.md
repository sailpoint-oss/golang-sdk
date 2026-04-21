---
id: nerm-user
title: User
pagination_label: User
sidebar_label: User
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'User', 'NERMUser'] 
slug: /tools/sdk/go/nerm/models/user
tags: ['SDK', 'Software Development Kit', 'User', 'NERMUser']
---

# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the object to retrieve or update | [optional] 
**Uid** | Pointer to **string** | UID of the user | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Email** | Pointer to **string** | The email | [optional] 
**Type** | Pointer to **string** | Type of user | [optional] [default to "NeprofileUser"]
**Title** | Pointer to **string** | The title | [optional] 
**Status** | Pointer to **string** | Status of the user | [optional] 
**Login** | Pointer to **string** | The login | [optional] 
**LastLogin** | Pointer to **SailPointTime** | When the user last logged in | [optional] [readonly] 
**CookiesAcceptedAt** | Pointer to **SailPointTime** | When cookies were accepted | [optional] [readonly] 
**PreferredLanguage** | Pointer to **string** | The locale the user prefers to use | [optional] 
**Locale** | Pointer to **string** | The locale the user prefers to use | [optional] 
**GroupStrings** | Pointer to **string** | Group strings configured on the customer's Active Directory configuration, provided by the IDP at the moment on authentication. | [optional] 
**SailpointIdentityId** | Pointer to **string** | The identity ID of the user in ISC | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *User) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *User) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *User) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *User) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *User) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogin

`func (o *User) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *User) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() SailPointTime`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*SailPointTime, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v SailPointTime)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetCookiesAcceptedAt

`func (o *User) GetCookiesAcceptedAt() SailPointTime`

GetCookiesAcceptedAt returns the CookiesAcceptedAt field if non-nil, zero value otherwise.

### GetCookiesAcceptedAtOk

`func (o *User) GetCookiesAcceptedAtOk() (*SailPointTime, bool)`

GetCookiesAcceptedAtOk returns a tuple with the CookiesAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesAcceptedAt

`func (o *User) SetCookiesAcceptedAt(v SailPointTime)`

SetCookiesAcceptedAt sets CookiesAcceptedAt field to given value.

### HasCookiesAcceptedAt

`func (o *User) HasCookiesAcceptedAt() bool`

HasCookiesAcceptedAt returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *User) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *User) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *User) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *User) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetLocale

`func (o *User) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *User) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetGroupStrings

`func (o *User) GetGroupStrings() string`

GetGroupStrings returns the GroupStrings field if non-nil, zero value otherwise.

### GetGroupStringsOk

`func (o *User) GetGroupStringsOk() (*string, bool)`

GetGroupStringsOk returns a tuple with the GroupStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupStrings

`func (o *User) SetGroupStrings(v string)`

SetGroupStrings sets GroupStrings field to given value.

### HasGroupStrings

`func (o *User) HasGroupStrings() bool`

HasGroupStrings returns a boolean if a field has been set.

### GetSailpointIdentityId

`func (o *User) GetSailpointIdentityId() string`

GetSailpointIdentityId returns the SailpointIdentityId field if non-nil, zero value otherwise.

### GetSailpointIdentityIdOk

`func (o *User) GetSailpointIdentityIdOk() (*string, bool)`

GetSailpointIdentityIdOk returns a tuple with the SailpointIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailpointIdentityId

`func (o *User) SetSailpointIdentityId(v string)`

SetSailpointIdentityId sets SailpointIdentityId field to given value.

### HasSailpointIdentityId

`func (o *User) HasSailpointIdentityId() bool`

HasSailpointIdentityId returns a boolean if a field has been set.


