---
id: nerm-workflow-action-performers1
title: WorkflowActionPerformers1
pagination_label: WorkflowActionPerformers1
sidebar_label: WorkflowActionPerformers1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowActionPerformers1', 'NERMWorkflowActionPerformers1'] 
slug: /tools/sdk/go/nerm/models/workflow-action-performers1
tags: ['SDK', 'Software Development Kit', 'WorkflowActionPerformers1', 'NERMWorkflowActionPerformers1']
---

# WorkflowActionPerformers1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContributorAttributeId** | Pointer to **string** | Specify the id of the user attribute to perform the action. | [optional] 
**Contributors** | Pointer to **bool** | Set to true to allow profile contributor to perform the action. | [optional] [default to false]
**ContributorsManagerAttributeId** | Pointer to **string** | Specify the id of the user attribute to perform the action. | [optional] 
**Owner** | Pointer to **bool** | Set to true to allow profile owner to perform the action. | [optional] [default to false]
**ProfilesContributorsAttributeId** | Pointer to **string** | Specify the id of the profile attribute to perform the action. | [optional] 
**Requester** | Pointer to **bool** | Set to true to allow requester from the request to perform the action. | [optional] [default to false]
**RequestersManager** | Pointer to **bool** | Set to true to allow the requester's manager from the request to perform the action. | [optional] [default to false]
**WorkflowActionId** | Pointer to **string** | Specify the id of the workflow action you would like to create the workflow action performer/s for. | [optional] 

## Methods

### NewWorkflowActionPerformers1

`func NewWorkflowActionPerformers1() *WorkflowActionPerformers1`

NewWorkflowActionPerformers1 instantiates a new WorkflowActionPerformers1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionPerformers1WithDefaults

`func NewWorkflowActionPerformers1WithDefaults() *WorkflowActionPerformers1`

NewWorkflowActionPerformers1WithDefaults instantiates a new WorkflowActionPerformers1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributorAttributeId

`func (o *WorkflowActionPerformers1) GetContributorAttributeId() string`

GetContributorAttributeId returns the ContributorAttributeId field if non-nil, zero value otherwise.

### GetContributorAttributeIdOk

`func (o *WorkflowActionPerformers1) GetContributorAttributeIdOk() (*string, bool)`

GetContributorAttributeIdOk returns a tuple with the ContributorAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorAttributeId

`func (o *WorkflowActionPerformers1) SetContributorAttributeId(v string)`

SetContributorAttributeId sets ContributorAttributeId field to given value.

### HasContributorAttributeId

`func (o *WorkflowActionPerformers1) HasContributorAttributeId() bool`

HasContributorAttributeId returns a boolean if a field has been set.

### GetContributors

`func (o *WorkflowActionPerformers1) GetContributors() bool`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *WorkflowActionPerformers1) GetContributorsOk() (*bool, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *WorkflowActionPerformers1) SetContributors(v bool)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *WorkflowActionPerformers1) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetContributorsManagerAttributeId

`func (o *WorkflowActionPerformers1) GetContributorsManagerAttributeId() string`

GetContributorsManagerAttributeId returns the ContributorsManagerAttributeId field if non-nil, zero value otherwise.

### GetContributorsManagerAttributeIdOk

`func (o *WorkflowActionPerformers1) GetContributorsManagerAttributeIdOk() (*string, bool)`

GetContributorsManagerAttributeIdOk returns a tuple with the ContributorsManagerAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsManagerAttributeId

`func (o *WorkflowActionPerformers1) SetContributorsManagerAttributeId(v string)`

SetContributorsManagerAttributeId sets ContributorsManagerAttributeId field to given value.

### HasContributorsManagerAttributeId

`func (o *WorkflowActionPerformers1) HasContributorsManagerAttributeId() bool`

HasContributorsManagerAttributeId returns a boolean if a field has been set.

### GetOwner

`func (o *WorkflowActionPerformers1) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowActionPerformers1) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowActionPerformers1) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowActionPerformers1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers1) GetProfilesContributorsAttributeId() string`

GetProfilesContributorsAttributeId returns the ProfilesContributorsAttributeId field if non-nil, zero value otherwise.

### GetProfilesContributorsAttributeIdOk

`func (o *WorkflowActionPerformers1) GetProfilesContributorsAttributeIdOk() (*string, bool)`

GetProfilesContributorsAttributeIdOk returns a tuple with the ProfilesContributorsAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers1) SetProfilesContributorsAttributeId(v string)`

SetProfilesContributorsAttributeId sets ProfilesContributorsAttributeId field to given value.

### HasProfilesContributorsAttributeId

`func (o *WorkflowActionPerformers1) HasProfilesContributorsAttributeId() bool`

HasProfilesContributorsAttributeId returns a boolean if a field has been set.

### GetRequester

`func (o *WorkflowActionPerformers1) GetRequester() bool`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *WorkflowActionPerformers1) GetRequesterOk() (*bool, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *WorkflowActionPerformers1) SetRequester(v bool)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *WorkflowActionPerformers1) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRequestersManager

`func (o *WorkflowActionPerformers1) GetRequestersManager() bool`

GetRequestersManager returns the RequestersManager field if non-nil, zero value otherwise.

### GetRequestersManagerOk

`func (o *WorkflowActionPerformers1) GetRequestersManagerOk() (*bool, bool)`

GetRequestersManagerOk returns a tuple with the RequestersManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestersManager

`func (o *WorkflowActionPerformers1) SetRequestersManager(v bool)`

SetRequestersManager sets RequestersManager field to given value.

### HasRequestersManager

`func (o *WorkflowActionPerformers1) HasRequestersManager() bool`

HasRequestersManager returns a boolean if a field has been set.

### GetWorkflowActionId

`func (o *WorkflowActionPerformers1) GetWorkflowActionId() string`

GetWorkflowActionId returns the WorkflowActionId field if non-nil, zero value otherwise.

### GetWorkflowActionIdOk

`func (o *WorkflowActionPerformers1) GetWorkflowActionIdOk() (*string, bool)`

GetWorkflowActionIdOk returns a tuple with the WorkflowActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionId

`func (o *WorkflowActionPerformers1) SetWorkflowActionId(v string)`

SetWorkflowActionId sets WorkflowActionId field to given value.

### HasWorkflowActionId

`func (o *WorkflowActionPerformers1) HasWorkflowActionId() bool`

HasWorkflowActionId returns a boolean if a field has been set.


