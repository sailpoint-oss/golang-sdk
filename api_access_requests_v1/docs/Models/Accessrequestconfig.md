---
id: v1-accessrequestconfig
title: Accessrequestconfig
pagination_label: Accessrequestconfig
sidebar_label: Accessrequestconfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Accessrequestconfig', 'V1Accessrequestconfig'] 
slug: /tools/sdk/go/accessrequestsv1/models/accessrequestconfig
tags: ['SDK', 'Software Development Kit', 'Accessrequestconfig', 'V1Accessrequestconfig']
---

# Accessrequestconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalsMustBeExternal** | Pointer to **bool** | If this is true, approvals must be processed by an external system. Also, if this is true, it blocks Request Center access requests and returns an error for any user who isn't an org admin. | [optional] [default to false]
**ReauthorizationEnabled** | Pointer to **bool** | If this is true, reauthorization will be enforced for appropriately configured access items. Enablement of this feature is currently in a limited state. | [optional] [default to false]
**RequestOnBehalfOfConfig** | Pointer to [**Requestonbehalfofconfig**](requestonbehalfofconfig) |  | [optional] 
**EntitlementRequestConfig** | Pointer to [**Entitlementrequestconfig**](entitlementrequestconfig) |  | [optional] 
**GovGroupVisibilityEnabled** | Pointer to **bool** | If this is true, requesters and requested-for users will be able to see the names of governance group members when a request is awaiting the group's approval. Up to the first 10 members of the group will be listed. | [optional] [default to false]

## Methods

### NewAccessrequestconfig

`func NewAccessrequestconfig() *Accessrequestconfig`

NewAccessrequestconfig instantiates a new Accessrequestconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessrequestconfigWithDefaults

`func NewAccessrequestconfigWithDefaults() *Accessrequestconfig`

NewAccessrequestconfigWithDefaults instantiates a new Accessrequestconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalsMustBeExternal

`func (o *Accessrequestconfig) GetApprovalsMustBeExternal() bool`

GetApprovalsMustBeExternal returns the ApprovalsMustBeExternal field if non-nil, zero value otherwise.

### GetApprovalsMustBeExternalOk

`func (o *Accessrequestconfig) GetApprovalsMustBeExternalOk() (*bool, bool)`

GetApprovalsMustBeExternalOk returns a tuple with the ApprovalsMustBeExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalsMustBeExternal

`func (o *Accessrequestconfig) SetApprovalsMustBeExternal(v bool)`

SetApprovalsMustBeExternal sets ApprovalsMustBeExternal field to given value.

### HasApprovalsMustBeExternal

`func (o *Accessrequestconfig) HasApprovalsMustBeExternal() bool`

HasApprovalsMustBeExternal returns a boolean if a field has been set.

### GetReauthorizationEnabled

`func (o *Accessrequestconfig) GetReauthorizationEnabled() bool`

GetReauthorizationEnabled returns the ReauthorizationEnabled field if non-nil, zero value otherwise.

### GetReauthorizationEnabledOk

`func (o *Accessrequestconfig) GetReauthorizationEnabledOk() (*bool, bool)`

GetReauthorizationEnabledOk returns a tuple with the ReauthorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthorizationEnabled

`func (o *Accessrequestconfig) SetReauthorizationEnabled(v bool)`

SetReauthorizationEnabled sets ReauthorizationEnabled field to given value.

### HasReauthorizationEnabled

`func (o *Accessrequestconfig) HasReauthorizationEnabled() bool`

HasReauthorizationEnabled returns a boolean if a field has been set.

### GetRequestOnBehalfOfConfig

`func (o *Accessrequestconfig) GetRequestOnBehalfOfConfig() Requestonbehalfofconfig`

GetRequestOnBehalfOfConfig returns the RequestOnBehalfOfConfig field if non-nil, zero value otherwise.

### GetRequestOnBehalfOfConfigOk

`func (o *Accessrequestconfig) GetRequestOnBehalfOfConfigOk() (*Requestonbehalfofconfig, bool)`

GetRequestOnBehalfOfConfigOk returns a tuple with the RequestOnBehalfOfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOnBehalfOfConfig

`func (o *Accessrequestconfig) SetRequestOnBehalfOfConfig(v Requestonbehalfofconfig)`

SetRequestOnBehalfOfConfig sets RequestOnBehalfOfConfig field to given value.

### HasRequestOnBehalfOfConfig

`func (o *Accessrequestconfig) HasRequestOnBehalfOfConfig() bool`

HasRequestOnBehalfOfConfig returns a boolean if a field has been set.

### GetEntitlementRequestConfig

`func (o *Accessrequestconfig) GetEntitlementRequestConfig() Entitlementrequestconfig`

GetEntitlementRequestConfig returns the EntitlementRequestConfig field if non-nil, zero value otherwise.

### GetEntitlementRequestConfigOk

`func (o *Accessrequestconfig) GetEntitlementRequestConfigOk() (*Entitlementrequestconfig, bool)`

GetEntitlementRequestConfigOk returns a tuple with the EntitlementRequestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementRequestConfig

`func (o *Accessrequestconfig) SetEntitlementRequestConfig(v Entitlementrequestconfig)`

SetEntitlementRequestConfig sets EntitlementRequestConfig field to given value.

### HasEntitlementRequestConfig

`func (o *Accessrequestconfig) HasEntitlementRequestConfig() bool`

HasEntitlementRequestConfig returns a boolean if a field has been set.

### GetGovGroupVisibilityEnabled

`func (o *Accessrequestconfig) GetGovGroupVisibilityEnabled() bool`

GetGovGroupVisibilityEnabled returns the GovGroupVisibilityEnabled field if non-nil, zero value otherwise.

### GetGovGroupVisibilityEnabledOk

`func (o *Accessrequestconfig) GetGovGroupVisibilityEnabledOk() (*bool, bool)`

GetGovGroupVisibilityEnabledOk returns a tuple with the GovGroupVisibilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovGroupVisibilityEnabled

`func (o *Accessrequestconfig) SetGovGroupVisibilityEnabled(v bool)`

SetGovGroupVisibilityEnabled sets GovGroupVisibilityEnabled field to given value.

### HasGovGroupVisibilityEnabled

`func (o *Accessrequestconfig) HasGovGroupVisibilityEnabled() bool`

HasGovGroupVisibilityEnabled returns a boolean if a field has been set.


