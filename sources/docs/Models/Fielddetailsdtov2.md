---
id: v1-fielddetailsdtov2
title: Fielddetailsdtov2
pagination_label: Fielddetailsdtov2
sidebar_label: Fielddetailsdtov2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Fielddetailsdtov2', 'V1Fielddetailsdtov2'] 
slug: /tools/sdk/go/sources/models/fielddetailsdtov2
tags: ['SDK', 'Software Development Kit', 'Fielddetailsdtov2', 'V1Fielddetailsdtov2']
---

# Fielddetailsdtov2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the attribute. | [optional] 
**Transform** | Pointer to **map[string]interface{}** | The transform to apply to the field | [optional] [default to {}]
**Attributes** | Pointer to **map[string]interface{}** | Attributes required for the transform | [optional] 
**IsRequired** | Pointer to **bool** | Flag indicating whether or not the attribute is required. | [optional] [readonly] [default to false]
**Type** | Pointer to **string** | The type of the attribute.  string: For text-based data.  int: For whole numbers.  long: For larger whole numbers.  date: For date and time values.  boolean: For true/false values.  secret: For sensitive data like passwords, which will be masked and encrypted.  | [optional] 
**IsMultiValued** | Pointer to **bool** | Flag indicating whether or not the attribute is multi-valued. | [optional] [default to false]

## Methods

### NewFielddetailsdtov2

`func NewFielddetailsdtov2() *Fielddetailsdtov2`

NewFielddetailsdtov2 instantiates a new Fielddetailsdtov2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFielddetailsdtov2WithDefaults

`func NewFielddetailsdtov2WithDefaults() *Fielddetailsdtov2`

NewFielddetailsdtov2WithDefaults instantiates a new Fielddetailsdtov2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Fielddetailsdtov2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fielddetailsdtov2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fielddetailsdtov2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Fielddetailsdtov2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransform

`func (o *Fielddetailsdtov2) GetTransform() map[string]interface{}`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *Fielddetailsdtov2) GetTransformOk() (*map[string]interface{}, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *Fielddetailsdtov2) SetTransform(v map[string]interface{})`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *Fielddetailsdtov2) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### GetAttributes

`func (o *Fielddetailsdtov2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Fielddetailsdtov2) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Fielddetailsdtov2) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Fielddetailsdtov2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsRequired

`func (o *Fielddetailsdtov2) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *Fielddetailsdtov2) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *Fielddetailsdtov2) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *Fielddetailsdtov2) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetType

`func (o *Fielddetailsdtov2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Fielddetailsdtov2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Fielddetailsdtov2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Fielddetailsdtov2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsMultiValued

`func (o *Fielddetailsdtov2) GetIsMultiValued() bool`

GetIsMultiValued returns the IsMultiValued field if non-nil, zero value otherwise.

### GetIsMultiValuedOk

`func (o *Fielddetailsdtov2) GetIsMultiValuedOk() (*bool, bool)`

GetIsMultiValuedOk returns a tuple with the IsMultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiValued

`func (o *Fielddetailsdtov2) SetIsMultiValued(v bool)`

SetIsMultiValued sets IsMultiValued field to given value.

### HasIsMultiValued

`func (o *Fielddetailsdtov2) HasIsMultiValued() bool`

HasIsMultiValued returns a boolean if a field has been set.


