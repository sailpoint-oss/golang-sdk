---
id: nerm-attribute1-validations-attributes
title: Attribute1ValidationsAttributes
pagination_label: Attribute1ValidationsAttributes
sidebar_label: Attribute1ValidationsAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Attribute1ValidationsAttributes', 'NERMAttribute1ValidationsAttributes'] 
slug: /tools/sdk/go/nerm/models/attribute1-validations-attributes
tags: ['SDK', 'Software Development Kit', 'Attribute1ValidationsAttributes', 'NERMAttribute1ValidationsAttributes']
---

# Attribute1ValidationsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationMethod** | Pointer to **string** | The type of validation to be applied | [optional] 
**Value** | Pointer to **string** | The value of the validator | [optional] 
**Destroy** | Pointer to **bool** | If the validator should be removed | [optional] 

## Methods

### NewAttribute1ValidationsAttributes

`func NewAttribute1ValidationsAttributes() *Attribute1ValidationsAttributes`

NewAttribute1ValidationsAttributes instantiates a new Attribute1ValidationsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttribute1ValidationsAttributesWithDefaults

`func NewAttribute1ValidationsAttributesWithDefaults() *Attribute1ValidationsAttributes`

NewAttribute1ValidationsAttributesWithDefaults instantiates a new Attribute1ValidationsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationMethod

`func (o *Attribute1ValidationsAttributes) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *Attribute1ValidationsAttributes) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *Attribute1ValidationsAttributes) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.

### HasValidationMethod

`func (o *Attribute1ValidationsAttributes) HasValidationMethod() bool`

HasValidationMethod returns a boolean if a field has been set.

### GetValue

`func (o *Attribute1ValidationsAttributes) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Attribute1ValidationsAttributes) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Attribute1ValidationsAttributes) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Attribute1ValidationsAttributes) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDestroy

`func (o *Attribute1ValidationsAttributes) GetDestroy() bool`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *Attribute1ValidationsAttributes) GetDestroyOk() (*bool, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *Attribute1ValidationsAttributes) SetDestroy(v bool)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *Attribute1ValidationsAttributes) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.


