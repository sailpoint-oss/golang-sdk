---
id: object-mapping-bulk-create-request
title: ObjectMappingBulkCreateRequest
pagination_label: ObjectMappingBulkCreateRequest
sidebar_label: ObjectMappingBulkCreateRequest
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'ObjectMappingBulkCreateRequest'] 
slug: /tools/sdk/go/v3/models/object-mapping-bulk-create-request
tags: ['SDK', 'Software Development Kit', 'ObjectMappingBulkCreateRequest']
---

# ObjectMappingBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewObjectMappings** |  [**[]ObjectMappingRequest**](object-mapping-request) |  | 

## Methods

### NewObjectMappingBulkCreateRequest

`func NewObjectMappingBulkCreateRequest(newObjectMappings []ObjectMappingRequest, ) *ObjectMappingBulkCreateRequest`

NewObjectMappingBulkCreateRequest instantiates a new ObjectMappingBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMappingBulkCreateRequestWithDefaults

`func NewObjectMappingBulkCreateRequestWithDefaults() *ObjectMappingBulkCreateRequest`

NewObjectMappingBulkCreateRequestWithDefaults instantiates a new ObjectMappingBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewObjectMappings

`func (o *ObjectMappingBulkCreateRequest) GetNewObjectMappings() []ObjectMappingRequest`

GetNewObjectMappings returns the NewObjectMappings field if non-nil, zero value otherwise.

### GetNewObjectMappingsOk

`func (o *ObjectMappingBulkCreateRequest) GetNewObjectMappingsOk() (*[]ObjectMappingRequest, bool)`

GetNewObjectMappingsOk returns a tuple with the NewObjectMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewObjectMappings

`func (o *ObjectMappingBulkCreateRequest) SetNewObjectMappings(v []ObjectMappingRequest)`

SetNewObjectMappings sets NewObjectMappings field to given value.



[[Back to top]](#) 


