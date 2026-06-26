---
id: nerm-create-workflow-page-request
title: CreateWorkflowPageRequest
pagination_label: CreateWorkflowPageRequest
sidebar_label: CreateWorkflowPageRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateWorkflowPageRequest', 'NERMCreateWorkflowPageRequest'] 
slug: /tools/sdk/go/nerm/models/create-workflow-page-request
tags: ['SDK', 'Software Development Kit', 'CreateWorkflowPageRequest', 'NERMCreateWorkflowPageRequest']
---

# CreateWorkflowPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**WorkflowPage**](workflow-page) |  | [optional] 

## Methods

### NewCreateWorkflowPageRequest

`func NewCreateWorkflowPageRequest() *CreateWorkflowPageRequest`

NewCreateWorkflowPageRequest instantiates a new CreateWorkflowPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowPageRequestWithDefaults

`func NewCreateWorkflowPageRequestWithDefaults() *CreateWorkflowPageRequest`

NewCreateWorkflowPageRequestWithDefaults instantiates a new CreateWorkflowPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CreateWorkflowPageRequest) GetPage() WorkflowPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CreateWorkflowPageRequest) GetPageOk() (*WorkflowPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CreateWorkflowPageRequest) SetPage(v WorkflowPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *CreateWorkflowPageRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


