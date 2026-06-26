---
id: nerm-user1
title: User1
pagination_label: User1
sidebar_label: User1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'User1', 'NERMUser1'] 
slug: /tools/sdk/go/nerm/models/user1
tags: ['SDK', 'Software Development Kit', 'User1', 'NERMUser1']
---

# User1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The user name | 
**Email** | **string** | The user email | 
**Type** | **string** | The user type | [default to "NeprofileUser"]
**ProfileId** | Pointer to **string** | The user profile id. Not required for NeprofileUser | [optional] 
**Title** | Pointer to **string** | The user description | [optional] 
**Status** | Pointer to **string** | The user status | [optional] 
**Login** | **string** | The user login | 
**GroupStrings** | Pointer to **string** | The user group strings | [optional] 
**Locale** | Pointer to **string** | The locale the user prefers to use | [optional] 
**Password** | Pointer to **string** | The user password. Not required for NeprofileUser | [optional] 
**SailpointIdentityId** | Pointer to **string** | The SailPoint Identity ID associated with this user | [optional] 

## Methods

### NewUser1

`func NewUser1(name string, email string, type_ string, login string, ) *User1`

NewUser1 instantiates a new User1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUser1WithDefaults

`func NewUser1WithDefaults() *User1`

NewUser1WithDefaults instantiates a new User1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *User1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User1) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *User1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetType

`func (o *User1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User1) SetType(v string)`

SetType sets Type field to given value.


### GetProfileId

`func (o *User1) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *User1) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *User1) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *User1) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetTitle

`func (o *User1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *User1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogin

`func (o *User1) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User1) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User1) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetGroupStrings

`func (o *User1) GetGroupStrings() string`

GetGroupStrings returns the GroupStrings field if non-nil, zero value otherwise.

### GetGroupStringsOk

`func (o *User1) GetGroupStringsOk() (*string, bool)`

GetGroupStringsOk returns a tuple with the GroupStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupStrings

`func (o *User1) SetGroupStrings(v string)`

SetGroupStrings sets GroupStrings field to given value.

### HasGroupStrings

`func (o *User1) HasGroupStrings() bool`

HasGroupStrings returns a boolean if a field has been set.

### GetLocale

`func (o *User1) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User1) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User1) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *User1) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPassword

`func (o *User1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSailpointIdentityId

`func (o *User1) GetSailpointIdentityId() string`

GetSailpointIdentityId returns the SailpointIdentityId field if non-nil, zero value otherwise.

### GetSailpointIdentityIdOk

`func (o *User1) GetSailpointIdentityIdOk() (*string, bool)`

GetSailpointIdentityIdOk returns a tuple with the SailpointIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailpointIdentityId

`func (o *User1) SetSailpointIdentityId(v string)`

SetSailpointIdentityId sets SailpointIdentityId field to given value.

### HasSailpointIdentityId

`func (o *User1) HasSailpointIdentityId() bool`

HasSailpointIdentityId returns a boolean if a field has been set.


