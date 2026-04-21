---
id: nerm-create-review-action-request
title: CreateReviewActionRequest
pagination_label: CreateReviewActionRequest
sidebar_label: CreateReviewActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateReviewActionRequest', 'NERMCreateReviewActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-review-action-request
tags: ['SDK', 'Software Development Kit', 'CreateReviewActionRequest', 'NERMCreateReviewActionRequest']
---

# CreateReviewActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**ReviewAction**](review-action) |  | [optional] 

## Methods

### NewCreateReviewActionRequest

`func NewCreateReviewActionRequest() *CreateReviewActionRequest`

NewCreateReviewActionRequest instantiates a new CreateReviewActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReviewActionRequestWithDefaults

`func NewCreateReviewActionRequestWithDefaults() *CreateReviewActionRequest`

NewCreateReviewActionRequestWithDefaults instantiates a new CreateReviewActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateReviewActionRequest) GetWorkflowAction() ReviewAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateReviewActionRequest) GetWorkflowActionOk() (*ReviewAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateReviewActionRequest) SetWorkflowAction(v ReviewAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateReviewActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


