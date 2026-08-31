---
id: v1-accessprofilemetadatabulkupdatebyfilterrequest
title: Accessprofilemetadatabulkupdatebyfilterrequest
pagination_label: Accessprofilemetadatabulkupdatebyfilterrequest
sidebar_label: Accessprofilemetadatabulkupdatebyfilterrequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Accessprofilemetadatabulkupdatebyfilterrequest', 'V1Accessprofilemetadatabulkupdatebyfilterrequest'] 
slug: /tools/sdk/go/accessprofiles/models/accessprofilemetadatabulkupdatebyfilterrequest
tags: ['SDK', 'Software Development Kit', 'Accessprofilemetadatabulkupdatebyfilterrequest', 'V1Accessprofilemetadatabulkupdatebyfilterrequest']
---

# Accessprofilemetadatabulkupdatebyfilterrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **created**: *gt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **requestable**: *eq*  **source.id**: *eq, in*  Supported composite operators are *and, or* | 
**Operation** | **string** | The operation to be performed | 
**ReplaceScope** | **string** | The choice of update scope. **ATTRIBUTE** replaces only the values of the attributes named in `values`, and **ALL** replaces every metadata attribute on the access profile. | 
**Values** | [**[]AccessprofilemetadatabulkupdatebyidrequestValuesInner**](accessprofilemetadatabulkupdatebyidrequest-values-inner) | The metadata to be updated, including attribute key and value. | 

## Methods

### NewAccessprofilemetadatabulkupdatebyfilterrequest

`func NewAccessprofilemetadatabulkupdatebyfilterrequest(filters string, operation string, replaceScope string, values []AccessprofilemetadatabulkupdatebyidrequestValuesInner, ) *Accessprofilemetadatabulkupdatebyfilterrequest`

NewAccessprofilemetadatabulkupdatebyfilterrequest instantiates a new Accessprofilemetadatabulkupdatebyfilterrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessprofilemetadatabulkupdatebyfilterrequestWithDefaults

`func NewAccessprofilemetadatabulkupdatebyfilterrequestWithDefaults() *Accessprofilemetadatabulkupdatebyfilterrequest`

NewAccessprofilemetadatabulkupdatebyfilterrequestWithDefaults instantiates a new Accessprofilemetadatabulkupdatebyfilterrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetFilters() string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetFiltersOk() (*string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) SetFilters(v string)`

SetFilters sets Filters field to given value.


### GetOperation

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReplaceScope

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetReplaceScope() string`

GetReplaceScope returns the ReplaceScope field if non-nil, zero value otherwise.

### GetReplaceScopeOk

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetReplaceScopeOk() (*string, bool)`

GetReplaceScopeOk returns a tuple with the ReplaceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceScope

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) SetReplaceScope(v string)`

SetReplaceScope sets ReplaceScope field to given value.


### GetValues

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetValues() []AccessprofilemetadatabulkupdatebyidrequestValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) GetValuesOk() (*[]AccessprofilemetadatabulkupdatebyidrequestValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Accessprofilemetadatabulkupdatebyfilterrequest) SetValues(v []AccessprofilemetadatabulkupdatebyidrequestValuesInner)`

SetValues sets Values field to given value.



