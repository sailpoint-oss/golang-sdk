---
id: discovered-applications-inner
title: DiscoveredApplicationsInner
pagination_label: DiscoveredApplicationsInner
sidebar_label: DiscoveredApplicationsInner
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'DiscoveredApplicationsInner'] 
slug: /tools/sdk/go/beta/models/discovered-applications-inner
tags: ['SDK', 'Software Development Kit', 'DiscoveredApplicationsInner']
---

# DiscoveredApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** |  Pointer to **string** | Unique identifier for the discovered application. | [optional] 
**Name** |  Pointer to **string** | Name of the discovered application. | [optional] 
**DiscoverySource** |  Pointer to **string** | Source from which the application was discovered. | [optional] 
**DiscoveredVendor** |  Pointer to **string** | The vendor associated with the discovered application. | [optional] 
**Description** |  Pointer to **string** | A brief description of the discovered application. | [optional] 
**RecommendedConnectors** |  Pointer to **[]string** | List of recommended connectors for the application. | [optional] 
**DiscoveredTimestamp** |  Pointer to **time.Time** | The timestamp when the application was discovered, in ISO 8601 format. | [optional] 

## Methods

### NewDiscoveredApplicationsInner

`func NewDiscoveredApplicationsInner() *DiscoveredApplicationsInner`

NewDiscoveredApplicationsInner instantiates a new DiscoveredApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredApplicationsInnerWithDefaults

`func NewDiscoveredApplicationsInnerWithDefaults() *DiscoveredApplicationsInner`

NewDiscoveredApplicationsInnerWithDefaults instantiates a new DiscoveredApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveredApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredApplicationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveredApplicationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DiscoveredApplicationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveredApplicationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveredApplicationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveredApplicationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDiscoverySource

`func (o *DiscoveredApplicationsInner) GetDiscoverySource() string`

GetDiscoverySource returns the DiscoverySource field if non-nil, zero value otherwise.

### GetDiscoverySourceOk

`func (o *DiscoveredApplicationsInner) GetDiscoverySourceOk() (*string, bool)`

GetDiscoverySourceOk returns a tuple with the DiscoverySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySource

`func (o *DiscoveredApplicationsInner) SetDiscoverySource(v string)`

SetDiscoverySource sets DiscoverySource field to given value.

### HasDiscoverySource

`func (o *DiscoveredApplicationsInner) HasDiscoverySource() bool`

HasDiscoverySource returns a boolean if a field has been set.

### GetDiscoveredVendor

`func (o *DiscoveredApplicationsInner) GetDiscoveredVendor() string`

GetDiscoveredVendor returns the DiscoveredVendor field if non-nil, zero value otherwise.

### GetDiscoveredVendorOk

`func (o *DiscoveredApplicationsInner) GetDiscoveredVendorOk() (*string, bool)`

GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVendor

`func (o *DiscoveredApplicationsInner) SetDiscoveredVendor(v string)`

SetDiscoveredVendor sets DiscoveredVendor field to given value.

### HasDiscoveredVendor

`func (o *DiscoveredApplicationsInner) HasDiscoveredVendor() bool`

HasDiscoveredVendor returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveredApplicationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveredApplicationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveredApplicationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveredApplicationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecommendedConnectors

`func (o *DiscoveredApplicationsInner) GetRecommendedConnectors() []string`

GetRecommendedConnectors returns the RecommendedConnectors field if non-nil, zero value otherwise.

### GetRecommendedConnectorsOk

`func (o *DiscoveredApplicationsInner) GetRecommendedConnectorsOk() (*[]string, bool)`

GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedConnectors

`func (o *DiscoveredApplicationsInner) SetRecommendedConnectors(v []string)`

SetRecommendedConnectors sets RecommendedConnectors field to given value.

### HasRecommendedConnectors

`func (o *DiscoveredApplicationsInner) HasRecommendedConnectors() bool`

HasRecommendedConnectors returns a boolean if a field has been set.

### GetDiscoveredTimestamp

`func (o *DiscoveredApplicationsInner) GetDiscoveredTimestamp() time.Time`

GetDiscoveredTimestamp returns the DiscoveredTimestamp field if non-nil, zero value otherwise.

### GetDiscoveredTimestampOk

`func (o *DiscoveredApplicationsInner) GetDiscoveredTimestampOk() (*time.Time, bool)`

GetDiscoveredTimestampOk returns a tuple with the DiscoveredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTimestamp

`func (o *DiscoveredApplicationsInner) SetDiscoveredTimestamp(v time.Time)`

SetDiscoveredTimestamp sets DiscoveredTimestamp field to given value.

### HasDiscoveredTimestamp

`func (o *DiscoveredApplicationsInner) HasDiscoveredTimestamp() bool`

HasDiscoveredTimestamp returns a boolean if a field has been set.


[[Back to top]](#) 


