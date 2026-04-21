---
id: nerm-ldap-action
title: LdapAction
pagination_label: LdapAction
sidebar_label: LdapAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'LdapAction', 'NERMLdapAction'] 
slug: /tools/sdk/go/nerm/models/ldap-action
tags: ['SDK', 'Software Development Kit', 'LdapAction', 'NERMLdapAction']
---

# LdapAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**StoreType** | **string** | the type of store. | 
**LdapActionUserRolesAttributes** | Pointer to [**LdapActionLdapActionUserRolesAttributes**](ldap-action-ldap-action-user-roles-attributes) |  | [optional] 
**WorkflowActionValueBuildersAttributes** | Pointer to [**LdapActionWorkflowActionValueBuildersAttributes**](ldap-action-workflow-action-value-builders-attributes) |  | [optional] 
**WorkflowActionDirectoryGroupsAttributes** | Pointer to [**LdapActionWorkflowActionDirectoryGroupsAttributes**](ldap-action-workflow-action-directory-groups-attributes) |  | [optional] 

## Methods

### NewLdapAction

`func NewLdapAction(workflowId string, description string, storeType string, ) *LdapAction`

NewLdapAction instantiates a new LdapAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapActionWithDefaults

`func NewLdapActionWithDefaults() *LdapAction`

NewLdapActionWithDefaults instantiates a new LdapAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *LdapAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *LdapAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *LdapAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *LdapAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *LdapAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *LdapAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *LdapAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *LdapAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStoreType

`func (o *LdapAction) GetStoreType() string`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *LdapAction) GetStoreTypeOk() (*string, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *LdapAction) SetStoreType(v string)`

SetStoreType sets StoreType field to given value.


### GetLdapActionUserRolesAttributes

`func (o *LdapAction) GetLdapActionUserRolesAttributes() LdapActionLdapActionUserRolesAttributes`

GetLdapActionUserRolesAttributes returns the LdapActionUserRolesAttributes field if non-nil, zero value otherwise.

### GetLdapActionUserRolesAttributesOk

`func (o *LdapAction) GetLdapActionUserRolesAttributesOk() (*LdapActionLdapActionUserRolesAttributes, bool)`

GetLdapActionUserRolesAttributesOk returns a tuple with the LdapActionUserRolesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapActionUserRolesAttributes

`func (o *LdapAction) SetLdapActionUserRolesAttributes(v LdapActionLdapActionUserRolesAttributes)`

SetLdapActionUserRolesAttributes sets LdapActionUserRolesAttributes field to given value.

### HasLdapActionUserRolesAttributes

`func (o *LdapAction) HasLdapActionUserRolesAttributes() bool`

HasLdapActionUserRolesAttributes returns a boolean if a field has been set.

### GetWorkflowActionValueBuildersAttributes

`func (o *LdapAction) GetWorkflowActionValueBuildersAttributes() LdapActionWorkflowActionValueBuildersAttributes`

GetWorkflowActionValueBuildersAttributes returns the WorkflowActionValueBuildersAttributes field if non-nil, zero value otherwise.

### GetWorkflowActionValueBuildersAttributesOk

`func (o *LdapAction) GetWorkflowActionValueBuildersAttributesOk() (*LdapActionWorkflowActionValueBuildersAttributes, bool)`

GetWorkflowActionValueBuildersAttributesOk returns a tuple with the WorkflowActionValueBuildersAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionValueBuildersAttributes

`func (o *LdapAction) SetWorkflowActionValueBuildersAttributes(v LdapActionWorkflowActionValueBuildersAttributes)`

SetWorkflowActionValueBuildersAttributes sets WorkflowActionValueBuildersAttributes field to given value.

### HasWorkflowActionValueBuildersAttributes

`func (o *LdapAction) HasWorkflowActionValueBuildersAttributes() bool`

HasWorkflowActionValueBuildersAttributes returns a boolean if a field has been set.

### GetWorkflowActionDirectoryGroupsAttributes

`func (o *LdapAction) GetWorkflowActionDirectoryGroupsAttributes() LdapActionWorkflowActionDirectoryGroupsAttributes`

GetWorkflowActionDirectoryGroupsAttributes returns the WorkflowActionDirectoryGroupsAttributes field if non-nil, zero value otherwise.

### GetWorkflowActionDirectoryGroupsAttributesOk

`func (o *LdapAction) GetWorkflowActionDirectoryGroupsAttributesOk() (*LdapActionWorkflowActionDirectoryGroupsAttributes, bool)`

GetWorkflowActionDirectoryGroupsAttributesOk returns a tuple with the WorkflowActionDirectoryGroupsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionDirectoryGroupsAttributes

`func (o *LdapAction) SetWorkflowActionDirectoryGroupsAttributes(v LdapActionWorkflowActionDirectoryGroupsAttributes)`

SetWorkflowActionDirectoryGroupsAttributes sets WorkflowActionDirectoryGroupsAttributes field to given value.

### HasWorkflowActionDirectoryGroupsAttributes

`func (o *LdapAction) HasWorkflowActionDirectoryGroupsAttributes() bool`

HasWorkflowActionDirectoryGroupsAttributes returns a boolean if a field has been set.


