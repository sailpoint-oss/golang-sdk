---
id: v1-accessprofilemetadatabulkupdatebyidrequest
title: Accessprofilemetadatabulkupdatebyidrequest
pagination_label: Accessprofilemetadatabulkupdatebyidrequest
sidebar_label: Accessprofilemetadatabulkupdatebyidrequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Accessprofilemetadatabulkupdatebyidrequest', 'V1Accessprofilemetadatabulkupdatebyidrequest'] 
slug: /tools/sdk/go/accessprofiles/models/accessprofilemetadatabulkupdatebyidrequest
tags: ['SDK', 'Software Development Kit', 'Accessprofilemetadatabulkupdatebyidrequest', 'V1Accessprofilemetadatabulkupdatebyidrequest']
---

# Accessprofilemetadatabulkupdatebyidrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessProfiles** | **[]string** | The IDs of the access profiles to update. | 
**Operation** | **string** | The operation to be performed | 
**ReplaceScope** | **string** | The choice of update scope. **ATTRIBUTE** replaces only the values of the attributes named in `values`, and **ALL** replaces every metadata attribute on the access profile. | 
**Values** | [**[]AccessprofilemetadatabulkupdatebyidrequestValuesInner**](accessprofilemetadatabulkupdatebyidrequest-values-inner) | The metadata to be updated, including attribute key and value. | 

## Methods

### NewAccessprofilemetadatabulkupdatebyidrequest

`func NewAccessprofilemetadatabulkupdatebyidrequest(accessProfiles []string, operation string, replaceScope string, values []AccessprofilemetadatabulkupdatebyidrequestValuesInner, ) *Accessprofilemetadatabulkupdatebyidrequest`

NewAccessprofilemetadatabulkupdatebyidrequest instantiates a new Accessprofilemetadatabulkupdatebyidrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessprofilemetadatabulkupdatebyidrequestWithDefaults

`func NewAccessprofilemetadatabulkupdatebyidrequestWithDefaults() *Accessprofilemetadatabulkupdatebyidrequest`

NewAccessprofilemetadatabulkupdatebyidrequestWithDefaults instantiates a new Accessprofilemetadatabulkupdatebyidrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessProfiles

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetAccessProfiles() []string`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetAccessProfilesOk() (*[]string, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *Accessprofilemetadatabulkupdatebyidrequest) SetAccessProfiles(v []string)`

SetAccessProfiles sets AccessProfiles field to given value.


### GetOperation

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Accessprofilemetadatabulkupdatebyidrequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReplaceScope

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetReplaceScope() string`

GetReplaceScope returns the ReplaceScope field if non-nil, zero value otherwise.

### GetReplaceScopeOk

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetReplaceScopeOk() (*string, bool)`

GetReplaceScopeOk returns a tuple with the ReplaceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceScope

`func (o *Accessprofilemetadatabulkupdatebyidrequest) SetReplaceScope(v string)`

SetReplaceScope sets ReplaceScope field to given value.


### GetValues

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetValues() []AccessprofilemetadatabulkupdatebyidrequestValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Accessprofilemetadatabulkupdatebyidrequest) GetValuesOk() (*[]AccessprofilemetadatabulkupdatebyidrequestValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Accessprofilemetadatabulkupdatebyidrequest) SetValues(v []AccessprofilemetadatabulkupdatebyidrequestValuesInner)`

SetValues sets Values field to given value.



