---
id: v1-businessapplicationref
title: Businessapplicationref
pagination_label: Businessapplicationref
sidebar_label: Businessapplicationref
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Businessapplicationref', 'V1Businessapplicationref'] 
slug: /tools/sdk/go/machineidentities/models/businessapplicationref
tags: ['SDK', 'Software Development Kit', 'Businessapplicationref', 'V1Businessapplicationref']
---

# Businessapplicationref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Reference type. | [optional] 
**Id** | Pointer to **string** | Business Application ID. | [optional] 
**Name** | Pointer to **NullableString** | Business Application display name. | [optional] 
**SanctionedStatus** | Pointer to **Sanctionedstatus** |  | [optional] 
**CorrelationType** | Pointer to **string** | Whether the Business Application reference was manually assigned or automatically correlated. | [optional] 

## Methods

### NewBusinessapplicationref

`func NewBusinessapplicationref() *Businessapplicationref`

NewBusinessapplicationref instantiates a new Businessapplicationref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessapplicationrefWithDefaults

`func NewBusinessapplicationrefWithDefaults() *Businessapplicationref`

NewBusinessapplicationrefWithDefaults instantiates a new Businessapplicationref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Businessapplicationref) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Businessapplicationref) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Businessapplicationref) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Businessapplicationref) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Businessapplicationref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Businessapplicationref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Businessapplicationref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Businessapplicationref) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Businessapplicationref) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Businessapplicationref) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Businessapplicationref) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Businessapplicationref) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Businessapplicationref) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Businessapplicationref) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSanctionedStatus

`func (o *Businessapplicationref) GetSanctionedStatus() Sanctionedstatus`

GetSanctionedStatus returns the SanctionedStatus field if non-nil, zero value otherwise.

### GetSanctionedStatusOk

`func (o *Businessapplicationref) GetSanctionedStatusOk() (*Sanctionedstatus, bool)`

GetSanctionedStatusOk returns a tuple with the SanctionedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionedStatus

`func (o *Businessapplicationref) SetSanctionedStatus(v Sanctionedstatus)`

SetSanctionedStatus sets SanctionedStatus field to given value.

### HasSanctionedStatus

`func (o *Businessapplicationref) HasSanctionedStatus() bool`

HasSanctionedStatus returns a boolean if a field has been set.

### GetCorrelationType

`func (o *Businessapplicationref) GetCorrelationType() string`

GetCorrelationType returns the CorrelationType field if non-nil, zero value otherwise.

### GetCorrelationTypeOk

`func (o *Businessapplicationref) GetCorrelationTypeOk() (*string, bool)`

GetCorrelationTypeOk returns a tuple with the CorrelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationType

`func (o *Businessapplicationref) SetCorrelationType(v string)`

SetCorrelationType sets CorrelationType field to given value.

### HasCorrelationType

`func (o *Businessapplicationref) HasCorrelationType() bool`

HasCorrelationType returns a boolean if a field has been set.


