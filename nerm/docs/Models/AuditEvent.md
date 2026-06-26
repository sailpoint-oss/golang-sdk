---
id: nerm-audit-event
title: AuditEvent
pagination_label: AuditEvent
sidebar_label: AuditEvent
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AuditEvent', 'NERMAuditEvent'] 
slug: /tools/sdk/go/nerm/models/audit-event
tags: ['SDK', 'Software Development Kit', 'AuditEvent', 'NERMAuditEvent']
---

# AuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**AuditEventCreatedAt**](audit-event-created-at) |  | [optional] 
**SubjectType** | Pointer to **string** | Categorization of audit event. | [optional] 
**Type** | Pointer to **string** | The type of audit event | [optional] 
**SubjectId** | Pointer to **string** | Identifier of the subject | [optional] 
**Data** | Pointer to [**AuditEventData**](audit-event-data) |  | [optional] 

## Methods

### NewAuditEvent

`func NewAuditEvent() *AuditEvent`

NewAuditEvent instantiates a new AuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventWithDefaults

`func NewAuditEventWithDefaults() *AuditEvent`

NewAuditEventWithDefaults instantiates a new AuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuditEvent) GetCreatedAt() AuditEventCreatedAt`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditEvent) GetCreatedAtOk() (*AuditEventCreatedAt, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditEvent) SetCreatedAt(v AuditEventCreatedAt)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSubjectType

`func (o *AuditEvent) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *AuditEvent) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *AuditEvent) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *AuditEvent) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetType

`func (o *AuditEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubjectId

`func (o *AuditEvent) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *AuditEvent) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *AuditEvent) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *AuditEvent) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetData

`func (o *AuditEvent) GetData() AuditEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditEvent) GetDataOk() (*AuditEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditEvent) SetData(v AuditEventData)`

SetData sets Data field to given value.

### HasData

`func (o *AuditEvent) HasData() bool`

HasData returns a boolean if a field has been set.


