---
id: nerm-workflow-action-performers
title: WorkflowActionPerformers
pagination_label: WorkflowActionPerformers
sidebar_label: WorkflowActionPerformers
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowActionPerformers', 'NERMWorkflowActionPerformers'] 
slug: /tools/sdk/go/nerm/models/workflow-action-performers
tags: ['SDK', 'Software Development Kit', 'WorkflowActionPerformers', 'NERMWorkflowActionPerformers']
---

# WorkflowActionPerformers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the workflow action performer that was created. | [optional] 
**ContributorAttributeId** | Pointer to **string** | The id of the user attribute to perform the action. | [optional] 
**Contributors** | Pointer to **bool** | Set to allow profile contributor to perform the action. | [optional] [default to false]
**ContributorsManagerAttributeId** | Pointer to **string** | The id of the user attribute to perform the action. | [optional] 
**Owner** | Pointer to **bool** | Set to allow profile owner to perform the action. | [optional] [default to false]
**ProfilesContributorsAttributeId** | Pointer to **string** | The id of the profile attribute to perform the action. | [optional] 
**Requester** | Pointer to **bool** | Set to allow requester from the request to perform the action. | [optional] [default to false]
**RequestersManager** | Pointer to **bool** | Set to allow the requester's manager from the request to perform the action. | [optional] [default to false]
**WorkflowActionId** | Pointer to **string** | The id of the workflow action. | [optional] 

## Methods

### NewWorkflowActionPerformers

`func NewWorkflowActionPerformers() *WorkflowActionPerformers`

NewWorkflowActionPerformers instantiates a new WorkflowActionPerformers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionPerformersWithDefaults

`func NewWorkflowActionPerformersWithDefaults() *WorkflowActionPerformers`

NewWorkflowActionPerformersWithDefaults instantiates a new WorkflowActionPerformers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowActionPerformers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowActionPerformers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowActionPerformers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowActionPerformers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContributorAttributeId

`func (o *WorkflowActionPerformers) GetContributorAttributeId() string`

GetContributorAttributeId returns the ContributorAttributeId field if non-nil, zero value otherwise.

### GetContributorAttributeIdOk

`func (o *WorkflowActionPerformers) GetContributorAttributeIdOk() (*string, bool)`

GetContributorAttributeIdOk returns a tuple with the ContributorAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorAttributeId

`func (o *WorkflowActionPerformers) SetContributorAttributeId(v string)`

SetContributorAttributeId sets ContributorAttributeId field to given value.

### HasContributorAttributeId

`func (o *WorkflowActionPerformers) HasContributorAttributeId() bool`

HasContributorAttributeId returns a boolean if a field has been set.

### GetContributors

`func (o *WorkflowActionPerformers) GetContributors() bool`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *WorkflowActionPerformers) GetContributorsOk() (*bool, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *WorkflowActionPerformers) SetContributors(v bool)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *WorkflowActionPerformers) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetContributorsManagerAttributeId

`func (o *WorkflowActionPerformers) GetContributorsManagerAttributeId() string`

GetContributorsManagerAttributeId returns the ContributorsManagerAttributeId field if non-nil, zero value otherwise.

### GetContributorsManagerAttributeIdOk

`func (o *WorkflowActionPerformers) GetContributorsManagerAttributeIdOk() (*string, bool)`

GetContributorsManagerAttributeIdOk returns a tuple with the ContributorsManagerAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsManagerAttributeId

`func (o *WorkflowActionPerformers) SetContributorsManagerAttributeId(v string)`

SetContributorsManagerAttributeId sets ContributorsManagerAttributeId field to given value.

### HasContributorsManagerAttributeId

`func (o *WorkflowActionPerformers) HasContributorsManagerAttributeId() bool`

HasContributorsManagerAttributeId returns a boolean if a field has been set.

### GetOwner

`func (o *WorkflowActionPerformers) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowActionPerformers) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowActionPerformers) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowActionPerformers) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers) GetProfilesContributorsAttributeId() string`

GetProfilesContributorsAttributeId returns the ProfilesContributorsAttributeId field if non-nil, zero value otherwise.

### GetProfilesContributorsAttributeIdOk

`func (o *WorkflowActionPerformers) GetProfilesContributorsAttributeIdOk() (*string, bool)`

GetProfilesContributorsAttributeIdOk returns a tuple with the ProfilesContributorsAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers) SetProfilesContributorsAttributeId(v string)`

SetProfilesContributorsAttributeId sets ProfilesContributorsAttributeId field to given value.

### HasProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers) HasProfilesContributorsAttributeId() bool`

HasProfilesContributorsAttributeId returns a boolean if a field has been set.

### GetRequester

`func (o *WorkflowActionPerformers) GetRequester() bool`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *WorkflowActionPerformers) GetRequesterOk() (*bool, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *WorkflowActionPerformers) SetRequester(v bool)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *WorkflowActionPerformers) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRequestersManager

`func (o *WorkflowActionPerformers) GetRequestersManager() bool`

GetRequestersManager returns the RequestersManager field if non-nil, zero value otherwise.

### GetRequestersManagerOk

`func (o *WorkflowActionPerformers) GetRequestersManagerOk() (*bool, bool)`

GetRequestersManagerOk returns a tuple with the RequestersManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestersManager

`func (o *WorkflowActionPerformers) SetRequestersManager(v bool)`

SetRequestersManager sets RequestersManager field to given value.

### HasRequestersManager

`func (o *WorkflowActionPerformers) HasRequestersManager() bool`

HasRequestersManager returns a boolean if a field has been set.

### GetWorkflowActionId

`func (o *WorkflowActionPerformers) GetWorkflowActionId() string`

GetWorkflowActionId returns the WorkflowActionId field if non-nil, zero value otherwise.

### GetWorkflowActionIdOk

`func (o *WorkflowActionPerformers) GetWorkflowActionIdOk() (*string, bool)`

GetWorkflowActionIdOk returns a tuple with the WorkflowActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionId

`func (o *WorkflowActionPerformers) SetWorkflowActionId(v string)`

SetWorkflowActionId sets WorkflowActionId field to given value.

### HasWorkflowActionId

`func (o *WorkflowActionPerformers) HasWorkflowActionId() bool`

HasWorkflowActionId returns a boolean if a field has been set.


