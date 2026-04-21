---
id: nerm-create-ldap-action-request
title: CreateLdapActionRequest
pagination_label: CreateLdapActionRequest
sidebar_label: CreateLdapActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateLdapActionRequest', 'NERMCreateLdapActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-ldap-action-request
tags: ['SDK', 'Software Development Kit', 'CreateLdapActionRequest', 'NERMCreateLdapActionRequest']
---

# CreateLdapActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**LdapAction**](ldap-action) |  | [optional] 

## Methods

### NewCreateLdapActionRequest

`func NewCreateLdapActionRequest() *CreateLdapActionRequest`

NewCreateLdapActionRequest instantiates a new CreateLdapActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapActionRequestWithDefaults

`func NewCreateLdapActionRequestWithDefaults() *CreateLdapActionRequest`

NewCreateLdapActionRequestWithDefaults instantiates a new CreateLdapActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateLdapActionRequest) GetWorkflowAction() LdapAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateLdapActionRequest) GetWorkflowActionOk() (*LdapAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateLdapActionRequest) SetWorkflowAction(v LdapAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateLdapActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


