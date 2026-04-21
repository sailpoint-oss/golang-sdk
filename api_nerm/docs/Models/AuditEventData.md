---
id: nerm-audit-event-data
title: AuditEventData
pagination_label: AuditEventData
sidebar_label: AuditEventData
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AuditEventData', 'NERMAuditEventData'] 
slug: /tools/sdk/go/nerm/models/audit-event-data
tags: ['SDK', 'Software Development Kit', 'AuditEventData', 'NERMAuditEventData']
---

# AuditEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | The profile id associated with the event | [optional] 
**WorkflowId** | Pointer to **string** | The workflow id associated with the event | [optional] 
**WorkflowName** | Pointer to **string** | The workflow name associated with the event | [optional] 
**WorkflowUid** | Pointer to **string** | The workflow uid associated with the event | [optional] 
**ProfileTypeId** | Pointer to **string** | The profile type associated with the event | [optional] 

## Methods

### NewAuditEventData

`func NewAuditEventData() *AuditEventData`

NewAuditEventData instantiates a new AuditEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventDataWithDefaults

`func NewAuditEventDataWithDefaults() *AuditEventData`

NewAuditEventDataWithDefaults instantiates a new AuditEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *AuditEventData) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AuditEventData) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AuditEventData) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AuditEventData) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AuditEventData) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AuditEventData) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AuditEventData) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AuditEventData) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowName

`func (o *AuditEventData) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *AuditEventData) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *AuditEventData) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *AuditEventData) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetWorkflowUid

`func (o *AuditEventData) GetWorkflowUid() string`

GetWorkflowUid returns the WorkflowUid field if non-nil, zero value otherwise.

### GetWorkflowUidOk

`func (o *AuditEventData) GetWorkflowUidOk() (*string, bool)`

GetWorkflowUidOk returns a tuple with the WorkflowUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowUid

`func (o *AuditEventData) SetWorkflowUid(v string)`

SetWorkflowUid sets WorkflowUid field to given value.

### HasWorkflowUid

`func (o *AuditEventData) HasWorkflowUid() bool`

HasWorkflowUid returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *AuditEventData) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *AuditEventData) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *AuditEventData) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *AuditEventData) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.


