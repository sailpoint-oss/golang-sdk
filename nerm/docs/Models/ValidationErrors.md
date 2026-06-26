---
id: nerm-validation-errors
title: ValidationErrors
pagination_label: ValidationErrors
sidebar_label: ValidationErrors
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ValidationErrors', 'NERMValidationErrors'] 
slug: /tools/sdk/go/nerm/models/validation-errors
tags: ['SDK', 'Software Development Kit', 'ValidationErrors', 'NERMValidationErrors']
---

# ValidationErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewValidationErrors

`func NewValidationErrors() *ValidationErrors`

NewValidationErrors instantiates a new ValidationErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorsWithDefaults

`func NewValidationErrorsWithDefaults() *ValidationErrors`

NewValidationErrorsWithDefaults instantiates a new ValidationErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ValidationErrors) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationErrors) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationErrors) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ValidationErrors) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *ValidationErrors) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationErrors) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationErrors) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidationErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


