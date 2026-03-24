---
id: full-discovered-applications
title: FullDiscoveredApplications
pagination_label: FullDiscoveredApplications
sidebar_label: FullDiscoveredApplications
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'FullDiscoveredApplications', 'FullDiscoveredApplications'] 
slug: /tools/sdk/go/v3/models/full-discovered-applications
tags: ['SDK', 'Software Development Kit', 'FullDiscoveredApplications', 'FullDiscoveredApplications']
---

# FullDiscoveredApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the discovered application. | [optional] 
**Name** | Pointer to **string** | Name of the discovered application. | [optional] 
**DiscoverySource** | Pointer to **string** | Source from which the application was discovered. | [optional] 
**DiscoveredVendor** | Pointer to **string** | The vendor associated with the discovered application. | [optional] 
**Description** | Pointer to **string** | A brief description of the discovered application. | [optional] 
**RecommendedConnectors** | Pointer to **[]string** | List of recommended connectors for the application. | [optional] 
**DiscoveredAt** | Pointer to **SailPointTime** | The timestamp when the application was last received via an entitlement aggregation invocation  or a manual csv upload, in ISO 8601 format. | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The timestamp when the application was first discovered, in ISO 8601 format. | [optional] 
**Status** | Pointer to **string** | The status of an application within the discovery source.  By default this field is set to \"ACTIVE\" when the application is discovered.  If an application has been deleted from within the discovery source, the status will be set to \"INACTIVE\". | [optional] 
**AssociatedSources** | Pointer to **[]string** | List of associated sources related to this discovered application. | [optional] 
**RiskScore** | Pointer to **int32** | The risk score of the application ranging from 0-100, 100 being highest risk. | [optional] 
**IsBusiness** | Pointer to **bool** | Indicates whether the application is used for business purposes. | [optional] [default to true]
**TotalSigninsCount** | Pointer to **int32** | The total number of sign-in accounts for the application. | [optional] 
**RiskLevel** | Pointer to **string** | The risk level of the application. | [optional] 

## Methods

### NewFullDiscoveredApplications

`func NewFullDiscoveredApplications() *FullDiscoveredApplications`

NewFullDiscoveredApplications instantiates a new FullDiscoveredApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDiscoveredApplicationsWithDefaults

`func NewFullDiscoveredApplicationsWithDefaults() *FullDiscoveredApplications`

NewFullDiscoveredApplicationsWithDefaults instantiates a new FullDiscoveredApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullDiscoveredApplications) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullDiscoveredApplications) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullDiscoveredApplications) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullDiscoveredApplications) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullDiscoveredApplications) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullDiscoveredApplications) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullDiscoveredApplications) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullDiscoveredApplications) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDiscoverySource

`func (o *FullDiscoveredApplications) GetDiscoverySource() string`

GetDiscoverySource returns the DiscoverySource field if non-nil, zero value otherwise.

### GetDiscoverySourceOk

`func (o *FullDiscoveredApplications) GetDiscoverySourceOk() (*string, bool)`

GetDiscoverySourceOk returns a tuple with the DiscoverySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySource

`func (o *FullDiscoveredApplications) SetDiscoverySource(v string)`

SetDiscoverySource sets DiscoverySource field to given value.

### HasDiscoverySource

`func (o *FullDiscoveredApplications) HasDiscoverySource() bool`

HasDiscoverySource returns a boolean if a field has been set.

### GetDiscoveredVendor

`func (o *FullDiscoveredApplications) GetDiscoveredVendor() string`

GetDiscoveredVendor returns the DiscoveredVendor field if non-nil, zero value otherwise.

### GetDiscoveredVendorOk

`func (o *FullDiscoveredApplications) GetDiscoveredVendorOk() (*string, bool)`

GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredVendor

`func (o *FullDiscoveredApplications) SetDiscoveredVendor(v string)`

SetDiscoveredVendor sets DiscoveredVendor field to given value.

### HasDiscoveredVendor

`func (o *FullDiscoveredApplications) HasDiscoveredVendor() bool`

HasDiscoveredVendor returns a boolean if a field has been set.

### GetDescription

`func (o *FullDiscoveredApplications) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullDiscoveredApplications) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullDiscoveredApplications) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullDiscoveredApplications) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecommendedConnectors

`func (o *FullDiscoveredApplications) GetRecommendedConnectors() []string`

GetRecommendedConnectors returns the RecommendedConnectors field if non-nil, zero value otherwise.

### GetRecommendedConnectorsOk

`func (o *FullDiscoveredApplications) GetRecommendedConnectorsOk() (*[]string, bool)`

GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedConnectors

`func (o *FullDiscoveredApplications) SetRecommendedConnectors(v []string)`

SetRecommendedConnectors sets RecommendedConnectors field to given value.

### HasRecommendedConnectors

`func (o *FullDiscoveredApplications) HasRecommendedConnectors() bool`

HasRecommendedConnectors returns a boolean if a field has been set.

### GetDiscoveredAt

`func (o *FullDiscoveredApplications) GetDiscoveredAt() SailPointTime`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *FullDiscoveredApplications) GetDiscoveredAtOk() (*SailPointTime, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *FullDiscoveredApplications) SetDiscoveredAt(v SailPointTime)`

SetDiscoveredAt sets DiscoveredAt field to given value.

### HasDiscoveredAt

`func (o *FullDiscoveredApplications) HasDiscoveredAt() bool`

HasDiscoveredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FullDiscoveredApplications) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullDiscoveredApplications) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullDiscoveredApplications) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullDiscoveredApplications) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *FullDiscoveredApplications) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FullDiscoveredApplications) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FullDiscoveredApplications) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FullDiscoveredApplications) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssociatedSources

`func (o *FullDiscoveredApplications) GetAssociatedSources() []string`

GetAssociatedSources returns the AssociatedSources field if non-nil, zero value otherwise.

### GetAssociatedSourcesOk

`func (o *FullDiscoveredApplications) GetAssociatedSourcesOk() (*[]string, bool)`

GetAssociatedSourcesOk returns a tuple with the AssociatedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedSources

`func (o *FullDiscoveredApplications) SetAssociatedSources(v []string)`

SetAssociatedSources sets AssociatedSources field to given value.

### HasAssociatedSources

`func (o *FullDiscoveredApplications) HasAssociatedSources() bool`

HasAssociatedSources returns a boolean if a field has been set.

### GetRiskScore

`func (o *FullDiscoveredApplications) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *FullDiscoveredApplications) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *FullDiscoveredApplications) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *FullDiscoveredApplications) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetIsBusiness

`func (o *FullDiscoveredApplications) GetIsBusiness() bool`

GetIsBusiness returns the IsBusiness field if non-nil, zero value otherwise.

### GetIsBusinessOk

`func (o *FullDiscoveredApplications) GetIsBusinessOk() (*bool, bool)`

GetIsBusinessOk returns a tuple with the IsBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusiness

`func (o *FullDiscoveredApplications) SetIsBusiness(v bool)`

SetIsBusiness sets IsBusiness field to given value.

### HasIsBusiness

`func (o *FullDiscoveredApplications) HasIsBusiness() bool`

HasIsBusiness returns a boolean if a field has been set.

### GetTotalSigninsCount

`func (o *FullDiscoveredApplications) GetTotalSigninsCount() int32`

GetTotalSigninsCount returns the TotalSigninsCount field if non-nil, zero value otherwise.

### GetTotalSigninsCountOk

`func (o *FullDiscoveredApplications) GetTotalSigninsCountOk() (*int32, bool)`

GetTotalSigninsCountOk returns a tuple with the TotalSigninsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSigninsCount

`func (o *FullDiscoveredApplications) SetTotalSigninsCount(v int32)`

SetTotalSigninsCount sets TotalSigninsCount field to given value.

### HasTotalSigninsCount

`func (o *FullDiscoveredApplications) HasTotalSigninsCount() bool`

HasTotalSigninsCount returns a boolean if a field has been set.

### GetRiskLevel

`func (o *FullDiscoveredApplications) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *FullDiscoveredApplications) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *FullDiscoveredApplications) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *FullDiscoveredApplications) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.


