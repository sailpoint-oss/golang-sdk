---
id: v1-business-application-ref
title: BusinessApplicationRef
pagination_label: BusinessApplicationRef
sidebar_label: BusinessApplicationRef
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'BusinessApplicationRef', 'V1BusinessApplicationRef'] 
slug: /tools/sdk/go/machineidentities/models/business-application-ref
tags: ['SDK', 'Software Development Kit', 'BusinessApplicationRef', 'V1BusinessApplicationRef']
---

# BusinessApplicationRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Reference type. | [optional] 
**Id** | Pointer to **string** | Business Application ID. | [optional] 
**Name** | Pointer to **NullableString** | Business Application display name. | [optional] 
**SanctionedStatus** | Pointer to **SanctionedStatus** |  | [optional] 
**CorrelationType** | Pointer to **string** | Whether the Business Application reference was manually assigned or automatically correlated. | [optional] 

## Methods

### NewBusinessApplicationRef

`func NewBusinessApplicationRef() *BusinessApplicationRef`

NewBusinessApplicationRef instantiates a new BusinessApplicationRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessApplicationRefWithDefaults

`func NewBusinessApplicationRefWithDefaults() *BusinessApplicationRef`

NewBusinessApplicationRefWithDefaults instantiates a new BusinessApplicationRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BusinessApplicationRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BusinessApplicationRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BusinessApplicationRef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BusinessApplicationRef) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *BusinessApplicationRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessApplicationRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessApplicationRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessApplicationRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BusinessApplicationRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessApplicationRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessApplicationRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessApplicationRef) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BusinessApplicationRef) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BusinessApplicationRef) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSanctionedStatus

`func (o *BusinessApplicationRef) GetSanctionedStatus() SanctionedStatus`

GetSanctionedStatus returns the SanctionedStatus field if non-nil, zero value otherwise.

### GetSanctionedStatusOk

`func (o *BusinessApplicationRef) GetSanctionedStatusOk() (*SanctionedStatus, bool)`

GetSanctionedStatusOk returns a tuple with the SanctionedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionedStatus

`func (o *BusinessApplicationRef) SetSanctionedStatus(v SanctionedStatus)`

SetSanctionedStatus sets SanctionedStatus field to given value.

### HasSanctionedStatus

`func (o *BusinessApplicationRef) HasSanctionedStatus() bool`

HasSanctionedStatus returns a boolean if a field has been set.

### GetCorrelationType

`func (o *BusinessApplicationRef) GetCorrelationType() string`

GetCorrelationType returns the CorrelationType field if non-nil, zero value otherwise.

### GetCorrelationTypeOk

`func (o *BusinessApplicationRef) GetCorrelationTypeOk() (*string, bool)`

GetCorrelationTypeOk returns a tuple with the CorrelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationType

`func (o *BusinessApplicationRef) SetCorrelationType(v string)`

SetCorrelationType sets CorrelationType field to given value.

### HasCorrelationType

`func (o *BusinessApplicationRef) HasCorrelationType() bool`

HasCorrelationType returns a boolean if a field has been set.


