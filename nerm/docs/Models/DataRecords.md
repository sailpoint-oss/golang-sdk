---
id: nerm-data-records
title: DataRecords
pagination_label: DataRecords
sidebar_label: DataRecords
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DataRecords', 'NERMDataRecords'] 
slug: /tools/sdk/go/nerm/models/data-records
tags: ['SDK', 'Software Development Kit', 'DataRecords', 'NERMDataRecords']
---

# DataRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterRecordId** | Pointer to **string** | The id of the master record | [optional] 

## Methods

### NewDataRecords

`func NewDataRecords() *DataRecords`

NewDataRecords instantiates a new DataRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRecordsWithDefaults

`func NewDataRecordsWithDefaults() *DataRecords`

NewDataRecordsWithDefaults instantiates a new DataRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterRecordId

`func (o *DataRecords) GetMasterRecordId() string`

GetMasterRecordId returns the MasterRecordId field if non-nil, zero value otherwise.

### GetMasterRecordIdOk

`func (o *DataRecords) GetMasterRecordIdOk() (*string, bool)`

GetMasterRecordIdOk returns a tuple with the MasterRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRecordId

`func (o *DataRecords) SetMasterRecordId(v string)`

SetMasterRecordId sets MasterRecordId field to given value.

### HasMasterRecordId

`func (o *DataRecords) HasMasterRecordId() bool`

HasMasterRecordId returns a boolean if a field has been set.


