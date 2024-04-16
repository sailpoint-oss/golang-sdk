---
id: approval-dto
title: ApprovalDto
pagination_label: ApprovalDto
sidebar_label: ApprovalDto
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'ApprovalDto'] 
slug: /tools/sdk/go/beta/models/approval-dto
tags: ['SDK', 'Software Development Kit', 'ApprovalDto']
---

# ApprovalDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** |  Pointer to **map[string]interface{}** | Object representation of a comment on the approval | [optional] 
**ApprovedBy** |  Pointer to [**[]ApprovalIdentity**](approval-identity) | An array of identities who have approved the approval | [optional] 
**RejectedBy** |  Pointer to [**[]ApprovalIdentity**](approval-identity) | An array of identities who have rejected the approval | [optional] 
**ReassignFrom** |  Pointer to [**ApprovalIdentity**](approval-identity) |  | [optional] 
**ReassignTo** |  Pointer to [**ApprovalIdentity**](approval-identity) |  | [optional] 
**AdditionalAttributes** |  Pointer to **map[string]interface{}** | Any additional attributes that the approval request may need | [optional] 

## Methods

### NewApprovalDto

`func NewApprovalDto() *ApprovalDto`

NewApprovalDto instantiates a new ApprovalDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalDtoWithDefaults

`func NewApprovalDtoWithDefaults() *ApprovalDto`

NewApprovalDtoWithDefaults instantiates a new ApprovalDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *ApprovalDto) GetComments() map[string]interface{}`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalDto) GetCommentsOk() (*map[string]interface{}, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalDto) SetComments(v map[string]interface{})`

SetComments sets Comments field to given value.

### HasComments

`func (o *ApprovalDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetApprovedBy

`func (o *ApprovalDto) GetApprovedBy() []ApprovalIdentity`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *ApprovalDto) GetApprovedByOk() (*[]ApprovalIdentity, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *ApprovalDto) SetApprovedBy(v []ApprovalIdentity)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *ApprovalDto) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetRejectedBy

`func (o *ApprovalDto) GetRejectedBy() []ApprovalIdentity`

GetRejectedBy returns the RejectedBy field if non-nil, zero value otherwise.

### GetRejectedByOk

`func (o *ApprovalDto) GetRejectedByOk() (*[]ApprovalIdentity, bool)`

GetRejectedByOk returns a tuple with the RejectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBy

`func (o *ApprovalDto) SetRejectedBy(v []ApprovalIdentity)`

SetRejectedBy sets RejectedBy field to given value.

### HasRejectedBy

`func (o *ApprovalDto) HasRejectedBy() bool`

HasRejectedBy returns a boolean if a field has been set.

### GetReassignFrom

`func (o *ApprovalDto) GetReassignFrom() ApprovalIdentity`

GetReassignFrom returns the ReassignFrom field if non-nil, zero value otherwise.

### GetReassignFromOk

`func (o *ApprovalDto) GetReassignFromOk() (*ApprovalIdentity, bool)`

GetReassignFromOk returns a tuple with the ReassignFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignFrom

`func (o *ApprovalDto) SetReassignFrom(v ApprovalIdentity)`

SetReassignFrom sets ReassignFrom field to given value.

### HasReassignFrom

`func (o *ApprovalDto) HasReassignFrom() bool`

HasReassignFrom returns a boolean if a field has been set.

### GetReassignTo

`func (o *ApprovalDto) GetReassignTo() ApprovalIdentity`

GetReassignTo returns the ReassignTo field if non-nil, zero value otherwise.

### GetReassignToOk

`func (o *ApprovalDto) GetReassignToOk() (*ApprovalIdentity, bool)`

GetReassignToOk returns a tuple with the ReassignTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignTo

`func (o *ApprovalDto) SetReassignTo(v ApprovalIdentity)`

SetReassignTo sets ReassignTo field to given value.

### HasReassignTo

`func (o *ApprovalDto) HasReassignTo() bool`

HasReassignTo returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *ApprovalDto) GetAdditionalAttributes() map[string]interface{}`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *ApprovalDto) GetAdditionalAttributesOk() (*map[string]interface{}, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *ApprovalDto) SetAdditionalAttributes(v map[string]interface{})`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *ApprovalDto) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to top]](#) 


