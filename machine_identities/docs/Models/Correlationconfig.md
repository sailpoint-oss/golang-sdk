---
id: v1-correlationconfig
title: Correlationconfig
pagination_label: Correlationconfig
sidebar_label: Correlationconfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Correlationconfig', 'V1Correlationconfig'] 
slug: /tools/sdk/go/machineidentities/models/correlationconfig
tags: ['SDK', 'Software Development Kit', 'Correlationconfig', 'V1Correlationconfig']
---

# Correlationconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | System-generated unique ID of the correlation config. | 
**SourceId** | **string** | The source ID this config belongs to. | 
**ResourceId** | **string** | The source resource identifier for this config scope. | 
**Type** | **string** | The correlation config type. | 
**Attributes** | **map[string]interface{}** | JSON object of config attributes. May include syncPrimaryToMachineAccounts (boolean) on OWNER_PRIMARY only. | 
**Rules** | [**[]Correlationrule**](correlationrule) | The ordered set of correlation rules for this config. | 
**Created** | Pointer to **SailPointTime** | Creation date of the config. | [optional] [readonly] 
**Modified** | Pointer to **SailPointTime** | Last modification date of the config. | [optional] [readonly] 

## Methods

### NewCorrelationconfig

`func NewCorrelationconfig(id string, sourceId string, resourceId string, type_ string, attributes map[string]interface{}, rules []Correlationrule, ) *Correlationconfig`

NewCorrelationconfig instantiates a new Correlationconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelationconfigWithDefaults

`func NewCorrelationconfigWithDefaults() *Correlationconfig`

NewCorrelationconfigWithDefaults instantiates a new Correlationconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Correlationconfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Correlationconfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Correlationconfig) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *Correlationconfig) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Correlationconfig) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Correlationconfig) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetResourceId

`func (o *Correlationconfig) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Correlationconfig) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Correlationconfig) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetType

`func (o *Correlationconfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Correlationconfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Correlationconfig) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Correlationconfig) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Correlationconfig) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Correlationconfig) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetRules

`func (o *Correlationconfig) GetRules() []Correlationrule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Correlationconfig) GetRulesOk() (*[]Correlationrule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Correlationconfig) SetRules(v []Correlationrule)`

SetRules sets Rules field to given value.


### GetCreated

`func (o *Correlationconfig) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Correlationconfig) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Correlationconfig) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Correlationconfig) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Correlationconfig) GetModified() SailPointTime`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Correlationconfig) GetModifiedOk() (*SailPointTime, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Correlationconfig) SetModified(v SailPointTime)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Correlationconfig) HasModified() bool`

HasModified returns a boolean if a field has been set.


