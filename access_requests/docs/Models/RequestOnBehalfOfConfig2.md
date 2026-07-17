---
id: v1-request-on-behalf-of-config2
title: RequestOnBehalfOfConfig2
pagination_label: RequestOnBehalfOfConfig2
sidebar_label: RequestOnBehalfOfConfig2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RequestOnBehalfOfConfig2', 'V1RequestOnBehalfOfConfig2'] 
slug: /tools/sdk/go/accessrequests/models/request-on-behalf-of-config2
tags: ['SDK', 'Software Development Kit', 'RequestOnBehalfOfConfig2', 'V1RequestOnBehalfOfConfig2']
---

# RequestOnBehalfOfConfig2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRequestOnBehalfOfAnyoneByAnyone** | Pointer to **bool** | If this is true, anyone can request access for anyone. | [optional] [default to false]
**AllowRequestOnBehalfOfEmployeeByManager** | Pointer to **bool** | If this is true, a manager can request access for his or her direct reports. | [optional] [default to false]

## Methods

### NewRequestOnBehalfOfConfig2

`func NewRequestOnBehalfOfConfig2() *RequestOnBehalfOfConfig2`

NewRequestOnBehalfOfConfig2 instantiates a new RequestOnBehalfOfConfig2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestOnBehalfOfConfig2WithDefaults

`func NewRequestOnBehalfOfConfig2WithDefaults() *RequestOnBehalfOfConfig2`

NewRequestOnBehalfOfConfig2WithDefaults instantiates a new RequestOnBehalfOfConfig2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRequestOnBehalfOfAnyoneByAnyone

`func (o *RequestOnBehalfOfConfig2) GetAllowRequestOnBehalfOfAnyoneByAnyone() bool`

GetAllowRequestOnBehalfOfAnyoneByAnyone returns the AllowRequestOnBehalfOfAnyoneByAnyone field if non-nil, zero value otherwise.

### GetAllowRequestOnBehalfOfAnyoneByAnyoneOk

`func (o *RequestOnBehalfOfConfig2) GetAllowRequestOnBehalfOfAnyoneByAnyoneOk() (*bool, bool)`

GetAllowRequestOnBehalfOfAnyoneByAnyoneOk returns a tuple with the AllowRequestOnBehalfOfAnyoneByAnyone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRequestOnBehalfOfAnyoneByAnyone

`func (o *RequestOnBehalfOfConfig2) SetAllowRequestOnBehalfOfAnyoneByAnyone(v bool)`

SetAllowRequestOnBehalfOfAnyoneByAnyone sets AllowRequestOnBehalfOfAnyoneByAnyone field to given value.

### HasAllowRequestOnBehalfOfAnyoneByAnyone

`func (o *RequestOnBehalfOfConfig2) HasAllowRequestOnBehalfOfAnyoneByAnyone() bool`

HasAllowRequestOnBehalfOfAnyoneByAnyone returns a boolean if a field has been set.

### GetAllowRequestOnBehalfOfEmployeeByManager

`func (o *RequestOnBehalfOfConfig2) GetAllowRequestOnBehalfOfEmployeeByManager() bool`

GetAllowRequestOnBehalfOfEmployeeByManager returns the AllowRequestOnBehalfOfEmployeeByManager field if non-nil, zero value otherwise.

### GetAllowRequestOnBehalfOfEmployeeByManagerOk

`func (o *RequestOnBehalfOfConfig2) GetAllowRequestOnBehalfOfEmployeeByManagerOk() (*bool, bool)`

GetAllowRequestOnBehalfOfEmployeeByManagerOk returns a tuple with the AllowRequestOnBehalfOfEmployeeByManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRequestOnBehalfOfEmployeeByManager

`func (o *RequestOnBehalfOfConfig2) SetAllowRequestOnBehalfOfEmployeeByManager(v bool)`

SetAllowRequestOnBehalfOfEmployeeByManager sets AllowRequestOnBehalfOfEmployeeByManager field to given value.

### HasAllowRequestOnBehalfOfEmployeeByManager

`func (o *RequestOnBehalfOfConfig2) HasAllowRequestOnBehalfOfEmployeeByManager() bool`

HasAllowRequestOnBehalfOfEmployeeByManager returns a boolean if a field has been set.


