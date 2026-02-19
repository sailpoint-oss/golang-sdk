---
id: v2026-approval-config
title: ApprovalConfig
pagination_label: ApprovalConfig
sidebar_label: ApprovalConfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ApprovalConfig', 'V2026ApprovalConfig'] 
slug: /tools/sdk/go/v2026/models/approval-config
tags: ['SDK', 'Software Development Kit', 'ApprovalConfig', 'V2026ApprovalConfig']
---

# ApprovalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | Pointer to **string** | Approvers must be listed as a comma-separated string, with each entry representing an individual or group authorized to approve account creation or deletion requests. | [optional] 
**Comments** | Pointer to **string** | Specifies the approval status for an account creation or deletion request. Allowed values are APPROVAL, REJECTION, ALL, and OFF. | [optional] 

## Methods

### NewApprovalConfig

`func NewApprovalConfig() *ApprovalConfig`

NewApprovalConfig instantiates a new ApprovalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConfigWithDefaults

`func NewApprovalConfigWithDefaults() *ApprovalConfig`

NewApprovalConfigWithDefaults instantiates a new ApprovalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *ApprovalConfig) GetApprovers() string`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ApprovalConfig) GetApproversOk() (*string, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ApprovalConfig) SetApprovers(v string)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *ApprovalConfig) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetComments

`func (o *ApprovalConfig) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalConfig) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalConfig) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ApprovalConfig) HasComments() bool`

HasComments returns a boolean if a field has been set.


