---
id: v2026-entitlement-dto
title: EntitlementDTO
pagination_label: EntitlementDTO
sidebar_label: EntitlementDTO
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'EntitlementDTO', 'V2026EntitlementDTO'] 
slug: /tools/sdk/go/v2026/models/entitlement-dto
tags: ['SDK', 'Software Development Kit', 'EntitlementDTO', 'V2026EntitlementDTO']
---

# EntitlementDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **NullableString** | Name of the Object | 
**Created** | Pointer to **SailPointTime** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **SailPointTime** | Last modification date of the Object | [optional] [readonly] 
**Attribute** | Pointer to **string** | Name of the entitlement attribute | [optional] 
**Uuid** | Pointer to **string** | Unique entitlement identifier within the database | [optional] 
**Value** | Pointer to **string** | Raw value of the entitlement | [optional] 
**Description** | Pointer to **string** | Entitlment description | [optional] 
**SourceSchemaObjectType** | Pointer to **string** | Schema objectType on the given application that maps to an Account Group | [optional] 
**Privileged** | Pointer to **bool** | Determines if this entitlement is privileged. | [optional] 
**IsGroup** | Pointer to **bool** | True when this object is used to represent a group attribute, otherwise it represents an account attribute. For the time being, the property is always true. | [optional] 
**CloudGoverned** | Pointer to **bool** | Determines if this entitlement is governed in the cloud. | [optional] 
**Requestable** | Pointer to **bool** | Determines if this entitlement is requestable. | [optional] 
**CloudEligible** | Pointer to **bool** | Determines if this entitlement is cloud eligible. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Entitlement attributes | [optional] 
**Source** | Pointer to [**EntitlementDTOAllOfSource**](entitlement-dto-all-of-source) |  | [optional] 
**Hash** | Pointer to **string** | Read-only calculated hash value of an entitlement | [optional] 
**DirectPermissions** | Pointer to [**[]PermissionDto**](permission-dto) |  | [optional] 
**InheritFrom** | Pointer to **[]string** | List of parent entitlements | [optional] 
**Segments** | Pointer to **[]string** | List of entitlement segments | [optional] 
**LastRefresh** | Pointer to **string** | Last time the entitlement was refreshed | [optional] 
**IdnServiceApp** | Pointer to **string** | IDN service application | [optional] 
**IdnExceptional** | Pointer to **string** | Informs whether an entitlement is a priviliged one. | [optional] 
**EntitlementitlementAggregated** | Pointer to **string** | Indicates whether an entitlement was aggregated | [optional] 
**SegmentStatus** | Pointer to **string** | Segment status (GLOBAL/LOCAL) | [optional] 
**Owner** | Pointer to [**OwnerReferenceDto**](owner-reference-dto) |  | [optional] 

## Methods

### NewEntitlementDTO

`func NewEntitlementDTO(name NullableString, ) *EntitlementDTO`

NewEntitlementDTO instantiates a new EntitlementDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDTOWithDefaults

`func NewEntitlementDTOWithDefaults() *EntitlementDTO`

NewEntitlementDTOWithDefaults instantiates a new EntitlementDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitlementDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementDTO) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *EntitlementDTO) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EntitlementDTO) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreated

`func (o *EntitlementDTO) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementDTO) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementDTO) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitlementDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *EntitlementDTO) GetModified() SailPointTime`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EntitlementDTO) GetModifiedOk() (*SailPointTime, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EntitlementDTO) SetModified(v SailPointTime)`

SetModified sets Modified field to given value.

### HasModified

`func (o *EntitlementDTO) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAttribute

`func (o *EntitlementDTO) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *EntitlementDTO) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *EntitlementDTO) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *EntitlementDTO) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetUuid

`func (o *EntitlementDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EntitlementDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EntitlementDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EntitlementDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValue

`func (o *EntitlementDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceSchemaObjectType

`func (o *EntitlementDTO) GetSourceSchemaObjectType() string`

GetSourceSchemaObjectType returns the SourceSchemaObjectType field if non-nil, zero value otherwise.

### GetSourceSchemaObjectTypeOk

`func (o *EntitlementDTO) GetSourceSchemaObjectTypeOk() (*string, bool)`

GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaObjectType

`func (o *EntitlementDTO) SetSourceSchemaObjectType(v string)`

SetSourceSchemaObjectType sets SourceSchemaObjectType field to given value.

### HasSourceSchemaObjectType

`func (o *EntitlementDTO) HasSourceSchemaObjectType() bool`

HasSourceSchemaObjectType returns a boolean if a field has been set.

### GetPrivileged

`func (o *EntitlementDTO) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *EntitlementDTO) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *EntitlementDTO) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *EntitlementDTO) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetIsGroup

`func (o *EntitlementDTO) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *EntitlementDTO) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *EntitlementDTO) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *EntitlementDTO) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetCloudGoverned

`func (o *EntitlementDTO) GetCloudGoverned() bool`

GetCloudGoverned returns the CloudGoverned field if non-nil, zero value otherwise.

### GetCloudGovernedOk

`func (o *EntitlementDTO) GetCloudGovernedOk() (*bool, bool)`

GetCloudGovernedOk returns a tuple with the CloudGoverned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGoverned

`func (o *EntitlementDTO) SetCloudGoverned(v bool)`

SetCloudGoverned sets CloudGoverned field to given value.

### HasCloudGoverned

`func (o *EntitlementDTO) HasCloudGoverned() bool`

HasCloudGoverned returns a boolean if a field has been set.

### GetRequestable

`func (o *EntitlementDTO) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *EntitlementDTO) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *EntitlementDTO) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *EntitlementDTO) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetCloudEligible

`func (o *EntitlementDTO) GetCloudEligible() bool`

GetCloudEligible returns the CloudEligible field if non-nil, zero value otherwise.

### GetCloudEligibleOk

`func (o *EntitlementDTO) GetCloudEligibleOk() (*bool, bool)`

GetCloudEligibleOk returns a tuple with the CloudEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEligible

`func (o *EntitlementDTO) SetCloudEligible(v bool)`

SetCloudEligible sets CloudEligible field to given value.

### HasCloudEligible

`func (o *EntitlementDTO) HasCloudEligible() bool`

HasCloudEligible returns a boolean if a field has been set.

### GetAttributes

`func (o *EntitlementDTO) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntitlementDTO) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntitlementDTO) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntitlementDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSource

`func (o *EntitlementDTO) GetSource() EntitlementDTOAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EntitlementDTO) GetSourceOk() (*EntitlementDTOAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EntitlementDTO) SetSource(v EntitlementDTOAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *EntitlementDTO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetHash

`func (o *EntitlementDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EntitlementDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EntitlementDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *EntitlementDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDirectPermissions

`func (o *EntitlementDTO) GetDirectPermissions() []PermissionDto`

GetDirectPermissions returns the DirectPermissions field if non-nil, zero value otherwise.

### GetDirectPermissionsOk

`func (o *EntitlementDTO) GetDirectPermissionsOk() (*[]PermissionDto, bool)`

GetDirectPermissionsOk returns a tuple with the DirectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPermissions

`func (o *EntitlementDTO) SetDirectPermissions(v []PermissionDto)`

SetDirectPermissions sets DirectPermissions field to given value.

### HasDirectPermissions

`func (o *EntitlementDTO) HasDirectPermissions() bool`

HasDirectPermissions returns a boolean if a field has been set.

### GetInheritFrom

`func (o *EntitlementDTO) GetInheritFrom() []string`

GetInheritFrom returns the InheritFrom field if non-nil, zero value otherwise.

### GetInheritFromOk

`func (o *EntitlementDTO) GetInheritFromOk() (*[]string, bool)`

GetInheritFromOk returns a tuple with the InheritFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritFrom

`func (o *EntitlementDTO) SetInheritFrom(v []string)`

SetInheritFrom sets InheritFrom field to given value.

### HasInheritFrom

`func (o *EntitlementDTO) HasInheritFrom() bool`

HasInheritFrom returns a boolean if a field has been set.

### GetSegments

`func (o *EntitlementDTO) GetSegments() []string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *EntitlementDTO) GetSegmentsOk() (*[]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *EntitlementDTO) SetSegments(v []string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *EntitlementDTO) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetLastRefresh

`func (o *EntitlementDTO) GetLastRefresh() string`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *EntitlementDTO) GetLastRefreshOk() (*string, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *EntitlementDTO) SetLastRefresh(v string)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *EntitlementDTO) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### GetIdnServiceApp

`func (o *EntitlementDTO) GetIdnServiceApp() string`

GetIdnServiceApp returns the IdnServiceApp field if non-nil, zero value otherwise.

### GetIdnServiceAppOk

`func (o *EntitlementDTO) GetIdnServiceAppOk() (*string, bool)`

GetIdnServiceAppOk returns a tuple with the IdnServiceApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnServiceApp

`func (o *EntitlementDTO) SetIdnServiceApp(v string)`

SetIdnServiceApp sets IdnServiceApp field to given value.

### HasIdnServiceApp

`func (o *EntitlementDTO) HasIdnServiceApp() bool`

HasIdnServiceApp returns a boolean if a field has been set.

### GetIdnExceptional

`func (o *EntitlementDTO) GetIdnExceptional() string`

GetIdnExceptional returns the IdnExceptional field if non-nil, zero value otherwise.

### GetIdnExceptionalOk

`func (o *EntitlementDTO) GetIdnExceptionalOk() (*string, bool)`

GetIdnExceptionalOk returns a tuple with the IdnExceptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnExceptional

`func (o *EntitlementDTO) SetIdnExceptional(v string)`

SetIdnExceptional sets IdnExceptional field to given value.

### HasIdnExceptional

`func (o *EntitlementDTO) HasIdnExceptional() bool`

HasIdnExceptional returns a boolean if a field has been set.

### GetEntitlementitlementAggregated

`func (o *EntitlementDTO) GetEntitlementitlementAggregated() string`

GetEntitlementitlementAggregated returns the EntitlementitlementAggregated field if non-nil, zero value otherwise.

### GetEntitlementitlementAggregatedOk

`func (o *EntitlementDTO) GetEntitlementitlementAggregatedOk() (*string, bool)`

GetEntitlementitlementAggregatedOk returns a tuple with the EntitlementitlementAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementitlementAggregated

`func (o *EntitlementDTO) SetEntitlementitlementAggregated(v string)`

SetEntitlementitlementAggregated sets EntitlementitlementAggregated field to given value.

### HasEntitlementitlementAggregated

`func (o *EntitlementDTO) HasEntitlementitlementAggregated() bool`

HasEntitlementitlementAggregated returns a boolean if a field has been set.

### GetSegmentStatus

`func (o *EntitlementDTO) GetSegmentStatus() string`

GetSegmentStatus returns the SegmentStatus field if non-nil, zero value otherwise.

### GetSegmentStatusOk

`func (o *EntitlementDTO) GetSegmentStatusOk() (*string, bool)`

GetSegmentStatusOk returns a tuple with the SegmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStatus

`func (o *EntitlementDTO) SetSegmentStatus(v string)`

SetSegmentStatus sets SegmentStatus field to given value.

### HasSegmentStatus

`func (o *EntitlementDTO) HasSegmentStatus() bool`

HasSegmentStatus returns a boolean if a field has been set.

### GetOwner

`func (o *EntitlementDTO) GetOwner() OwnerReferenceDto`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EntitlementDTO) GetOwnerOk() (*OwnerReferenceDto, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EntitlementDTO) SetOwner(v OwnerReferenceDto)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EntitlementDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


