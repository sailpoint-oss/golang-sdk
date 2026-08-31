---
id: v1-accessprofilemetadatabulkupdateresponse
title: Accessprofilemetadatabulkupdateresponse
pagination_label: Accessprofilemetadatabulkupdateresponse
sidebar_label: Accessprofilemetadatabulkupdateresponse
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Accessprofilemetadatabulkupdateresponse', 'V1Accessprofilemetadatabulkupdateresponse'] 
slug: /tools/sdk/go/accessprofiles/models/accessprofilemetadatabulkupdateresponse
tags: ['SDK', 'Software Development Kit', 'Accessprofilemetadatabulkupdateresponse', 'V1Accessprofilemetadatabulkupdateresponse']
---

# Accessprofilemetadatabulkupdateresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the task that is processing the bulk update. | [optional] 
**Type** | Pointer to **string** | Type of the object the bulk update applies to. | [optional] 
**Status** | Pointer to **string** | The status of the bulk update request. | [optional] 
**Created** | Pointer to **SailPointTime** | Time when the bulk update request was created | [optional] 

## Methods

### NewAccessprofilemetadatabulkupdateresponse

`func NewAccessprofilemetadatabulkupdateresponse() *Accessprofilemetadatabulkupdateresponse`

NewAccessprofilemetadatabulkupdateresponse instantiates a new Accessprofilemetadatabulkupdateresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessprofilemetadatabulkupdateresponseWithDefaults

`func NewAccessprofilemetadatabulkupdateresponseWithDefaults() *Accessprofilemetadatabulkupdateresponse`

NewAccessprofilemetadatabulkupdateresponseWithDefaults instantiates a new Accessprofilemetadatabulkupdateresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Accessprofilemetadatabulkupdateresponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Accessprofilemetadatabulkupdateresponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Accessprofilemetadatabulkupdateresponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Accessprofilemetadatabulkupdateresponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Accessprofilemetadatabulkupdateresponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Accessprofilemetadatabulkupdateresponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Accessprofilemetadatabulkupdateresponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Accessprofilemetadatabulkupdateresponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Accessprofilemetadatabulkupdateresponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Accessprofilemetadatabulkupdateresponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Accessprofilemetadatabulkupdateresponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Accessprofilemetadatabulkupdateresponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *Accessprofilemetadatabulkupdateresponse) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Accessprofilemetadatabulkupdateresponse) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Accessprofilemetadatabulkupdateresponse) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Accessprofilemetadatabulkupdateresponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


