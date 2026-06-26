---
id: nerm-audit-event-created-at
title: AuditEventCreatedAt
pagination_label: AuditEventCreatedAt
sidebar_label: AuditEventCreatedAt
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AuditEventCreatedAt', 'NERMAuditEventCreatedAt'] 
slug: /tools/sdk/go/nerm/models/audit-event-created-at
tags: ['SDK', 'Software Development Kit', 'AuditEventCreatedAt', 'NERMAuditEventCreatedAt']
---

# AuditEventCreatedAt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gt** | Pointer to **string** | Greater Than - search for events with a date greater than the value | [optional] 
**Lt** | Pointer to **string** | Less Than - search for events with a date less than the value | [optional] 
**Eq** | Pointer to **string** | Equal - search for events with a date equal to the value | [optional] 

## Methods

### NewAuditEventCreatedAt

`func NewAuditEventCreatedAt() *AuditEventCreatedAt`

NewAuditEventCreatedAt instantiates a new AuditEventCreatedAt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventCreatedAtWithDefaults

`func NewAuditEventCreatedAtWithDefaults() *AuditEventCreatedAt`

NewAuditEventCreatedAtWithDefaults instantiates a new AuditEventCreatedAt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGt

`func (o *AuditEventCreatedAt) GetGt() string`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *AuditEventCreatedAt) GetGtOk() (*string, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *AuditEventCreatedAt) SetGt(v string)`

SetGt sets Gt field to given value.

### HasGt

`func (o *AuditEventCreatedAt) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetLt

`func (o *AuditEventCreatedAt) GetLt() string`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *AuditEventCreatedAt) GetLtOk() (*string, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *AuditEventCreatedAt) SetLt(v string)`

SetLt sets Lt field to given value.

### HasLt

`func (o *AuditEventCreatedAt) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetEq

`func (o *AuditEventCreatedAt) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *AuditEventCreatedAt) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *AuditEventCreatedAt) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *AuditEventCreatedAt) HasEq() bool`

HasEq returns a boolean if a field has been set.


