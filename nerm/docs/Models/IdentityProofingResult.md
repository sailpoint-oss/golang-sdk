---
id: nerm-identity-proofing-result
title: IdentityProofingResult
pagination_label: IdentityProofingResult
sidebar_label: IdentityProofingResult
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IdentityProofingResult', 'NERMIdentityProofingResult'] 
slug: /tools/sdk/go/nerm/models/identity-proofing-result
tags: ['SDK', 'Software Development Kit', 'IdentityProofingResult', 'NERMIdentityProofingResult']
---

# IdentityProofingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IdentityProofingActionId** | Pointer to **string** |  | [optional] 
**WorkflowSessionId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**ProofingWorkflow** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**ProofingAttributes** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **SailPointTime** |  | [optional] 
**UpdatedAt** | Pointer to **SailPointTime** |  | [optional] 

## Methods

### NewIdentityProofingResult

`func NewIdentityProofingResult() *IdentityProofingResult`

NewIdentityProofingResult instantiates a new IdentityProofingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProofingResultWithDefaults

`func NewIdentityProofingResultWithDefaults() *IdentityProofingResult`

NewIdentityProofingResultWithDefaults instantiates a new IdentityProofingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProofingResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProofingResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProofingResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProofingResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProofingActionId

`func (o *IdentityProofingResult) GetIdentityProofingActionId() string`

GetIdentityProofingActionId returns the IdentityProofingActionId field if non-nil, zero value otherwise.

### GetIdentityProofingActionIdOk

`func (o *IdentityProofingResult) GetIdentityProofingActionIdOk() (*string, bool)`

GetIdentityProofingActionIdOk returns a tuple with the IdentityProofingActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProofingActionId

`func (o *IdentityProofingResult) SetIdentityProofingActionId(v string)`

SetIdentityProofingActionId sets IdentityProofingActionId field to given value.

### HasIdentityProofingActionId

`func (o *IdentityProofingResult) HasIdentityProofingActionId() bool`

HasIdentityProofingActionId returns a boolean if a field has been set.

### GetWorkflowSessionId

`func (o *IdentityProofingResult) GetWorkflowSessionId() string`

GetWorkflowSessionId returns the WorkflowSessionId field if non-nil, zero value otherwise.

### GetWorkflowSessionIdOk

`func (o *IdentityProofingResult) GetWorkflowSessionIdOk() (*string, bool)`

GetWorkflowSessionIdOk returns a tuple with the WorkflowSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSessionId

`func (o *IdentityProofingResult) SetWorkflowSessionId(v string)`

SetWorkflowSessionId sets WorkflowSessionId field to given value.

### HasWorkflowSessionId

`func (o *IdentityProofingResult) HasWorkflowSessionId() bool`

HasWorkflowSessionId returns a boolean if a field has been set.

### GetProfileId

`func (o *IdentityProofingResult) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IdentityProofingResult) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *IdentityProofingResult) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *IdentityProofingResult) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProofingWorkflow

`func (o *IdentityProofingResult) GetProofingWorkflow() string`

GetProofingWorkflow returns the ProofingWorkflow field if non-nil, zero value otherwise.

### GetProofingWorkflowOk

`func (o *IdentityProofingResult) GetProofingWorkflowOk() (*string, bool)`

GetProofingWorkflowOk returns a tuple with the ProofingWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofingWorkflow

`func (o *IdentityProofingResult) SetProofingWorkflow(v string)`

SetProofingWorkflow sets ProofingWorkflow field to given value.

### HasProofingWorkflow

`func (o *IdentityProofingResult) HasProofingWorkflow() bool`

HasProofingWorkflow returns a boolean if a field has been set.

### GetResult

`func (o *IdentityProofingResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *IdentityProofingResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *IdentityProofingResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *IdentityProofingResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetProofingAttributes

`func (o *IdentityProofingResult) GetProofingAttributes() map[string]string`

GetProofingAttributes returns the ProofingAttributes field if non-nil, zero value otherwise.

### GetProofingAttributesOk

`func (o *IdentityProofingResult) GetProofingAttributesOk() (*map[string]string, bool)`

GetProofingAttributesOk returns a tuple with the ProofingAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofingAttributes

`func (o *IdentityProofingResult) SetProofingAttributes(v map[string]string)`

SetProofingAttributes sets ProofingAttributes field to given value.

### HasProofingAttributes

`func (o *IdentityProofingResult) HasProofingAttributes() bool`

HasProofingAttributes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentityProofingResult) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityProofingResult) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityProofingResult) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityProofingResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityProofingResult) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityProofingResult) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityProofingResult) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityProofingResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


