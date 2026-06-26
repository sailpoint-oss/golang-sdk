---
id: nerm-search200-response
title: Search200Response
pagination_label: Search200Response
sidebar_label: Search200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Search200Response', 'NERMSearch200Response'] 
slug: /tools/sdk/go/nerm/models/search200-response
tags: ['SDK', 'Software Development Kit', 'Search200Response', 'NERMSearch200Response']
---

# Search200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEvents** | Pointer to [**[]AuditEvent**](audit-event) |  | [optional] 

## Methods

### NewSearch200Response

`func NewSearch200Response() *Search200Response`

NewSearch200Response instantiates a new Search200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearch200ResponseWithDefaults

`func NewSearch200ResponseWithDefaults() *Search200Response`

NewSearch200ResponseWithDefaults instantiates a new Search200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEvents

`func (o *Search200Response) GetAuditEvents() []AuditEvent`

GetAuditEvents returns the AuditEvents field if non-nil, zero value otherwise.

### GetAuditEventsOk

`func (o *Search200Response) GetAuditEventsOk() (*[]AuditEvent, bool)`

GetAuditEventsOk returns a tuple with the AuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEvents

`func (o *Search200Response) SetAuditEvents(v []AuditEvent)`

SetAuditEvents sets AuditEvents field to given value.

### HasAuditEvents

`func (o *Search200Response) HasAuditEvents() bool`

HasAuditEvents returns a boolean if a field has been set.


