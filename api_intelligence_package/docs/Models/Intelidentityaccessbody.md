---
id: v1-intelidentityaccessbody
title: Intelidentityaccessbody
pagination_label: Intelidentityaccessbody
sidebar_label: Intelidentityaccessbody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityaccessbody', 'V1Intelidentityaccessbody'] 
slug: /tools/sdk/go/intelligencepackage/models/intelidentityaccessbody
tags: ['SDK', 'Software Development Kit', 'Intelidentityaccessbody', 'V1Intelidentityaccessbody']
---

# Intelidentityaccessbody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]Intelaccessaccountwire**](intelaccessaccountwire) | Accounts for the identity in camelCase wire format from Shelby List Accounts. | 
**PrivilegedAccessItems** | [**[]Intelprivilegedaccessitemwire**](intelprivilegedaccessitemwire) | Privileged access items for the identity returned by SDS Search. | 

## Methods

### NewIntelidentityaccessbody

`func NewIntelidentityaccessbody(accounts []Intelaccessaccountwire, privilegedAccessItems []Intelprivilegedaccessitemwire, ) *Intelidentityaccessbody`

NewIntelidentityaccessbody instantiates a new Intelidentityaccessbody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityaccessbodyWithDefaults

`func NewIntelidentityaccessbodyWithDefaults() *Intelidentityaccessbody`

NewIntelidentityaccessbodyWithDefaults instantiates a new Intelidentityaccessbody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *Intelidentityaccessbody) GetAccounts() []Intelaccessaccountwire`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Intelidentityaccessbody) GetAccountsOk() (*[]Intelaccessaccountwire, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Intelidentityaccessbody) SetAccounts(v []Intelaccessaccountwire)`

SetAccounts sets Accounts field to given value.


### GetPrivilegedAccessItems

`func (o *Intelidentityaccessbody) GetPrivilegedAccessItems() []Intelprivilegedaccessitemwire`

GetPrivilegedAccessItems returns the PrivilegedAccessItems field if non-nil, zero value otherwise.

### GetPrivilegedAccessItemsOk

`func (o *Intelidentityaccessbody) GetPrivilegedAccessItemsOk() (*[]Intelprivilegedaccessitemwire, bool)`

GetPrivilegedAccessItemsOk returns a tuple with the PrivilegedAccessItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedAccessItems

`func (o *Intelidentityaccessbody) SetPrivilegedAccessItems(v []Intelprivilegedaccessitemwire)`

SetPrivilegedAccessItems sets PrivilegedAccessItems field to given value.



