---
id: v1-accessprofilemetadatabulkupdatebyidrequest-values-inner
title: AccessprofilemetadatabulkupdatebyidrequestValuesInner
pagination_label: AccessprofilemetadatabulkupdatebyidrequestValuesInner
sidebar_label: AccessprofilemetadatabulkupdatebyidrequestValuesInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccessprofilemetadatabulkupdatebyidrequestValuesInner', 'V1AccessprofilemetadatabulkupdatebyidrequestValuesInner'] 
slug: /tools/sdk/go/accessprofiles/models/accessprofilemetadatabulkupdatebyidrequest-values-inner
tags: ['SDK', 'Software Development Kit', 'AccessprofilemetadatabulkupdatebyidrequestValuesInner', 'V1AccessprofilemetadatabulkupdatebyidrequestValuesInner']
---

# AccessprofilemetadatabulkupdatebyidrequestValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The technical name of the metadata attribute. | 
**Values** | **[]string** | The values of the attribute to be updated. | 
**ObjectType** | Pointer to **string** | The type of the metadata attribute. Set to `custom` for custom metadata attributes, which require a suite license. | [optional] 

## Methods

### NewAccessprofilemetadatabulkupdatebyidrequestValuesInner

`func NewAccessprofilemetadatabulkupdatebyidrequestValuesInner(attribute string, values []string, ) *AccessprofilemetadatabulkupdatebyidrequestValuesInner`

NewAccessprofilemetadatabulkupdatebyidrequestValuesInner instantiates a new AccessprofilemetadatabulkupdatebyidrequestValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessprofilemetadatabulkupdatebyidrequestValuesInnerWithDefaults

`func NewAccessprofilemetadatabulkupdatebyidrequestValuesInnerWithDefaults() *AccessprofilemetadatabulkupdatebyidrequestValuesInner`

NewAccessprofilemetadatabulkupdatebyidrequestValuesInnerWithDefaults instantiates a new AccessprofilemetadatabulkupdatebyidrequestValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetValues

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) SetValues(v []string)`

SetValues sets Values field to given value.


### SetValuesNil

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetObjectType

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AccessprofilemetadatabulkupdatebyidrequestValuesInner) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


