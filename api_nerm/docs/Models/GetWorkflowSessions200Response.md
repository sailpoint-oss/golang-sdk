---
id: nerm-get-workflow-sessions200-response
title: GetWorkflowSessions200Response
pagination_label: GetWorkflowSessions200Response
sidebar_label: GetWorkflowSessions200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetWorkflowSessions200Response', 'NERMGetWorkflowSessions200Response'] 
slug: /tools/sdk/go/nerm/models/get-workflow-sessions200-response
tags: ['SDK', 'Software Development Kit', 'GetWorkflowSessions200Response', 'NERMGetWorkflowSessions200Response']
---

# GetWorkflowSessions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowSessions** | Pointer to [**[]WorkflowSession**](workflow-session) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetWorkflowSessions200Response

`func NewGetWorkflowSessions200Response() *GetWorkflowSessions200Response`

NewGetWorkflowSessions200Response instantiates a new GetWorkflowSessions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowSessions200ResponseWithDefaults

`func NewGetWorkflowSessions200ResponseWithDefaults() *GetWorkflowSessions200Response`

NewGetWorkflowSessions200ResponseWithDefaults instantiates a new GetWorkflowSessions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowSessions

`func (o *GetWorkflowSessions200Response) GetWorkflowSessions() []WorkflowSession`

GetWorkflowSessions returns the WorkflowSessions field if non-nil, zero value otherwise.

### GetWorkflowSessionsOk

`func (o *GetWorkflowSessions200Response) GetWorkflowSessionsOk() (*[]WorkflowSession, bool)`

GetWorkflowSessionsOk returns a tuple with the WorkflowSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSessions

`func (o *GetWorkflowSessions200Response) SetWorkflowSessions(v []WorkflowSession)`

SetWorkflowSessions sets WorkflowSessions field to given value.

### HasWorkflowSessions

`func (o *GetWorkflowSessions200Response) HasWorkflowSessions() bool`

HasWorkflowSessions returns a boolean if a field has been set.

### GetMetadata

`func (o *GetWorkflowSessions200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetWorkflowSessions200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetWorkflowSessions200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetWorkflowSessions200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


