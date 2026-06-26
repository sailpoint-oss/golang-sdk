---
id: nerm-create-contributors-action-request
title: CreateContributorsActionRequest
pagination_label: CreateContributorsActionRequest
sidebar_label: CreateContributorsActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateContributorsActionRequest', 'NERMCreateContributorsActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-contributors-action-request
tags: ['SDK', 'Software Development Kit', 'CreateContributorsActionRequest', 'NERMCreateContributorsActionRequest']
---

# CreateContributorsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**ContributorsAction**](contributors-action) |  | [optional] 

## Methods

### NewCreateContributorsActionRequest

`func NewCreateContributorsActionRequest() *CreateContributorsActionRequest`

NewCreateContributorsActionRequest instantiates a new CreateContributorsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContributorsActionRequestWithDefaults

`func NewCreateContributorsActionRequestWithDefaults() *CreateContributorsActionRequest`

NewCreateContributorsActionRequestWithDefaults instantiates a new CreateContributorsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateContributorsActionRequest) GetWorkflowAction() ContributorsAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateContributorsActionRequest) GetWorkflowActionOk() (*ContributorsAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateContributorsActionRequest) SetWorkflowAction(v ContributorsAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateContributorsActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


